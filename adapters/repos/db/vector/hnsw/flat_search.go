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

package hnsw

import (
	"github.com/weaviate/weaviate/adapters/repos/db/helpers"
	"github.com/weaviate/weaviate/adapters/repos/db/priorityqueue"
	"github.com/weaviate/weaviate/adapters/repos/db/vector/compressionhelpers"
)

func (h *hnsw) flatSearch(queryVector []float32, k, limit int,
	allowList helpers.AllowList,
) ([]uint64, []float32, error) {
	if !h.shouldRescore() {
		limit = k
	}
	results := priorityqueue.NewMax[any](limit)

	h.RLock()
	nodeSize := uint64(len(h.nodes))
	h.RUnlock()

	var compressorDistancer compressionhelpers.CompressorDistancer
	if h.compressed.Load() {
		distancer, returnFn := h.compressor.NewDistancer(queryVector)
		defer returnFn()
		compressorDistancer = distancer
	}

	it := allowList.Iterator()
	for candidate, ok := it.Next(); ok; candidate, ok = it.Next() {
		// Hot fix for https://github.com/weaviate/weaviate/issues/1937
		// this if statement mitigates the problem but it doesn't resolve the issue
		if candidate >= nodeSize {
			h.logger.WithField("action", "flatSearch").
				Debugf("trying to get candidate: %v but we only have: %v elements.",
					candidate, nodeSize)
			continue
		}

		h.shardedNodeLocks.RLock(candidate)
		c := h.nodes[candidate]
		h.shardedNodeLocks.RUnlock(candidate)

		if c == nil || h.hasTombstone(candidate) {
			continue
		}

		dist, err := h.distToNode(compressorDistancer, candidate, queryVector)
		if err != nil {
			return nil, nil, err
		}

		if results.Len() < limit {
			results.Insert(candidate, dist)
		} else if results.Top().Dist > dist {
			results.Pop()
			results.Insert(candidate, dist)
		}
	}

	if h.shouldRescore() {
		compressorDistancer, fn := h.compressor.NewDistancer(queryVector)
		h.rescore(results, k, compressorDistancer)
		fn()
	}

	ids := make([]uint64, results.Len())
	dists := make([]float32, results.Len())

	// results is ordered in reverse, we need to flip the order before presenting
	// to the user!
	i := len(ids) - 1
	for results.Len() > 0 {
		res := results.Pop()
		ids[i] = res.ID
		dists[i] = res.Dist
		i--
	}

	return ids, dists, nil
}
