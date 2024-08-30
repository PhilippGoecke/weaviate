//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package resolver

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	raftImpl "github.com/hashicorp/raft"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/cluster/log"
	"github.com/weaviate/weaviate/cluster/utils"
)

const (
	// RaftTcpMaxPool controls how many connections raft transport will pool
	raftTcpMaxPool = 3
	// RaftTcpTimeout is used to apply I/O deadlines.
	raftTcpTimeout = 10 * time.Second
)

type raft struct {
	// Resolver allows the raft to also be used to resolve node-ids to ip addresses.
	Resolver

	// RaftPort is the configured RAFT port in the cluster that the resolver will append to the node id.
	RaftPort int
	// IsLocalCluster is the cluster running on a single host machine. This is necessary to ensure that we don't use the
	// same port multiple time when we only have a single underlying machine.
	IsLocalCluster bool
	// NodeNameToPortMap maps a given node name ot a given port. This is useful when running locally so that we can
	// keep in memory which node uses which port.
	NodeNameToPortMap map[string]int

	nodesLock        sync.Mutex
	notResolvedNodes map[raftImpl.ServerID]struct{}
}

func NewRaft(cfg RaftConfig) *raft {
	return &raft{
		Resolver:          cfg.NodeToAddress,
		RaftPort:          cfg.RaftPort,
		IsLocalCluster:    cfg.IsLocalHost,
		NodeNameToPortMap: cfg.NodeNameToPortMap,
		notResolvedNodes:  make(map[raftImpl.ServerID]struct{}),
	}
}

// ServerAddr resolves server ID to a RAFT address
func (a *raft) ServerAddr(id raftImpl.ServerID) (raftImpl.ServerAddress, error) {
	addr := ""
	err := backoff.Retry(func() error {
		// Get the address from the node id
		addr = a.Resolver.NodeAddress(string(id))

		// Update the internal notResolvedNodes if the addr if empty, otherwise delete it from the map
		a.nodesLock.Lock()
		defer a.nodesLock.Unlock()
		if addr == "" {
			a.notResolvedNodes[id] = struct{}{}
			return fmt.Errorf("could not resolve server id %s", id)
		}
		return nil
	}, utils.ConstantBackoff(3, 3*time.Second))

	if err != nil {
		return "", err
	}
	delete(a.notResolvedNodes, id)

	// If we are not running a local cluster we can immediately return, otherwise we need to lookup the port of the node
	// as we can't use the default raft port locally.
	if !a.IsLocalCluster {
		return raftImpl.ServerAddress(fmt.Sprintf("%s:%d", addr, a.RaftPort)), nil
	}
	return raftImpl.ServerAddress(fmt.Sprintf("%s:%d", addr, a.NodeNameToPortMap[string(id)])), nil
}

// NewTCPTransport returns a new raft.NetworkTransportConfig that utilizes
// this resolver to resolve addresses based on server IDs.
// This is particularly crucial as K8s assigns new IPs on each node restart.
func (a *raft) NewTCPTransport(
	bindAddr string,
	advertise net.Addr,
	maxPool int,
	timeout time.Duration,
	logger *logrus.Logger,
) (*raftImpl.NetworkTransport, error) {
	cfg := &raftImpl.NetworkTransportConfig{
		ServerAddressProvider: a,
		MaxPool:               raftTcpMaxPool,
		Timeout:               raftTcpTimeout,
		Logger:                log.NewHCLogrusLogger("raft-net", logger),
	}
	return raftImpl.NewTCPTransportWithConfig(bindAddr, advertise, cfg)
}

func (a *raft) NotResolvedNodes() map[raftImpl.ServerID]struct{} {
	return a.notResolvedNodes
}
