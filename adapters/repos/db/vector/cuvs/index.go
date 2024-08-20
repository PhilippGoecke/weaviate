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

//go:build cuvs

package cuvs_index

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"

	"io"
	"sync"

	"github.com/pkg/errors"
	cuvs "github.com/rapidsai/cuvs/go"
	"github.com/rapidsai/cuvs/go/cagra"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/adapters/repos/db/helpers"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv"
	"github.com/weaviate/weaviate/adapters/repos/db/vector/common"
	"github.com/weaviate/weaviate/adapters/repos/db/vector/hnsw/distancer"
	schemaConfig "github.com/weaviate/weaviate/entities/schema/config"
	cuvsEnt "github.com/weaviate/weaviate/entities/vectorindex/cuvs"
)

type VectorIndex interface {
	// Dump(labels ...string)
	Add(id uint64, vector []float32) error
	AddBatch(ctx context.Context, id []uint64, vector [][]float32) error
	Delete(id ...uint64) error
	SearchByVector(vector []float32, k int, allow helpers.AllowList) ([]uint64, []float32, error)
	// SearchByVectorDistance(vector []float32, dist float32,
	// 	maxLimit int64, allow helpers.AllowList) ([]uint64, []float32, error)
	// UpdateUserConfig(updated schemaconfig.VectorIndexConfig, callback func()) error
	// Drop(ctx context.Context) error
	// Shutdown(ctx context.Context) error
	// Flush() error
	// SwitchCommitLogs(ctx context.Context) error
	// ListFiles(ctx context.Context, basePath string) ([]string, error)
	PostStartup()
	// Compressed() bool
	// ValidateBeforeInsert(vector []float32) error
	// DistanceBetweenVectors(x, y []float32) (float32, error)
	// ContainsNode(id uint64) bool
	// DistancerProvider() distancer.Provider
	AlreadyIndexed() uint64
	// QueryVectorDistancer(queryVector []float32) common.QueryVectorDistancer
}

type cuvs_index struct {
	sync.Mutex
	id              string
	targetVector    string
	dims            int32
	store           *lsmkv.Store
	logger          logrus.FieldLogger
	distanceMetric  cuvs.Distance
	cuvsIndex       *cagra.CagraIndex
	cuvsIndexParams *cagra.IndexParams
	dlpackTensor    *cuvs.Tensor[float32]
	idCuvsIdMap     map[uint32]uint64
	cuvsResource    *cuvs.Resource
	cuvsExpandCount uint64

	// rescore             int64
	// bq                  compressionhelpers.BinaryQuantizer

	// pqResults *common.PqMaxPool
	// pool      *pools

	// compression string
	// bqCache     cache.Cache[uint64]
	count uint64
}

func New(cfg Config, uc cuvsEnt.UserConfig, store *lsmkv.Store) (*cuvs_index, error) {
	if err := cfg.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid config")
	}

	logger := cfg.Logger
	if logger == nil {
		l := logrus.New()
		l.Out = io.Discard
		logger = l
	}

	res, err := cuvs.NewResource(nil)

	if err != nil {
		return nil, err
	}

	cuvsIndexParams, err := cagra.CreateIndexParams()

	if err != nil {
		return nil, err
	}

	cuvsIndex, err := cagra.CreateIndex()

	if err != nil {
		return nil, fmt.Errorf("create cuvs index: %w", err)
	}

	index := &cuvs_index{
		id:              cfg.ID,
		targetVector:    cfg.TargetVector,
		logger:          logger,
		distanceMetric:  cuvs.DistanceL2, //TODO: make configurable
		cuvsIndex:       cuvsIndex,
		cuvsIndexParams: cuvsIndexParams,
		cuvsResource:    &res,
		dlpackTensor:    nil,
		idCuvsIdMap:     make(map[uint32]uint64),
	}

	// if err := index.initBuckets(context.Background()); err != nil {
	// 	return nil, fmt.Errorf("init cuvs index buckets: %w", err)
	// }

	return index, nil

}

func byteSliceFromFloat32Slice(vector []float32, slice []byte) []byte {
	for i := range vector {
		binary.LittleEndian.PutUint32(slice[i*4:], math.Float32bits(vector[i]))
	}
	return slice
}

func (index *cuvs_index) Add(id uint64, vector []float32) error {
	index.logger.Debug("adding single")
	return index.AddBatch(context.Background(), []uint64{id}, [][]float32{vector})
	// index.Lock()
	// defer index.Unlock()

	// if index.cuvsIndex == nil {
	// 	return errors.New("cuvs index is nil")
	// }

	// if len(vector) != int(index.dims) {
	// 	return errors.Errorf("insert called with a vector of the wrong size")
	// }

	// slice := make([]byte, len(vector)*4)

	// // store in bucket
	// idBytes := make([]byte, 8)
	// binary.BigEndian.PutUint64(idBytes, id)
	// index.store.Bucket(index.getBucketName()).Put(idBytes, byteSliceFromFloat32Slice(vector, slice))

	// // vector = index.normalized(vector)

	// index.idCuvsIdMap[uint32(index.count)] = id

	// index.count += 1

	// index.dlpackTensor.Expand(index.cuvsResource, [][]float32{vector})

	// err := cagra.BuildIndex(*index.cuvsResource, index.cuvsIndexParams, index.dlpackTensor, index.cuvsIndex)

	// if err != nil {
	// 	return err
	// }

	// return nil
}

func (index *cuvs_index) AddBatch(ctx context.Context, id []uint64, vector [][]float32) error {
	index.Lock()
	defer index.Unlock()

	index.logger.Debug("adding batch, batch size: ", len(id))
	index.logger.Debug("adding batch, batch dimension: ", len(vector[0]))
	// if len(vector[0]) != 128 {
	// 	panic("length not 128")
	// }

	if err := ctx.Err(); err != nil {
		return err
	}

	if index.cuvsIndex == nil {
		return errors.New("cuvs index is nil")
	}

	// store in bucket
	// for i := range id {
	// 	slice := make([]byte, len(vector)*4)
	// 	idBytes := make([]byte, 8)
	// 	binary.BigEndian.PutUint64(idBytes, id[i])
	// 	index.store.Bucket(index.getBucketName()).Put(idBytes, byteSliceFromFloat32Slice(vector[i], slice))
	// }

	for i := range id {
		index.idCuvsIdMap[uint32(index.count)] = id[i]
		index.count += 1
	}

	// index.dlpackTensor.Expand(index.cuvsResource, vector)
	if index.dlpackTensor == nil {
		tensor, err := cuvs.NewTensor(false, vector)
		if err != nil {
			return err
		}
		_, err = tensor.ToDevice(index.cuvsResource)
		if err != nil {
			return err
		}
		index.dlpackTensor = &tensor
	} else {
		println("getShape:" + strconv.FormatInt(index.dlpackTensor.GetShape()[1], 10))
		_, err := index.dlpackTensor.Expand(index.cuvsResource, vector)
		if err != nil {
			return err
		}
	}

	err := cagra.BuildIndex(*index.cuvsResource, index.cuvsIndexParams, index.dlpackTensor, index.cuvsIndex)

	if err != nil {
		return err
	}

	return nil

}

func (index *cuvs_index) Delete(ids ...uint64) error {
	return nil
}

func (index *cuvs_index) SearchByVector(vector [][]float32, k int, allow helpers.AllowList) ([]uint64, []float32, error) {
	index.Lock()
	defer index.Unlock()

	tensor, err := cuvs.NewTensor(false, vector)

	if err != nil {
		return nil, nil, err
	}

	_, err = tensor.ToDevice(index.cuvsResource)

	if err != nil {
		return nil, nil, err
	}

	queries, err := cuvs.NewTensor(true, vector)

	if err != nil {
		return nil, nil, err
	}

	NeighborsDataset := make([][]uint32, 1)
	for i := range NeighborsDataset {
		NeighborsDataset[i] = make([]uint32, k)
	}
	DistancesDataset := make([][]float32, 1)
	for i := range DistancesDataset {
		DistancesDataset[i] = make([]float32, k)
	}

	neighbors, err := cuvs.NewTensor(true, NeighborsDataset)

	if err != nil {
		return nil, nil, err
	}

	distances, err := cuvs.NewTensor(true, DistancesDataset)

	if err != nil {
		return nil, nil, err
	}

	_, err = queries.ToDevice(index.cuvsResource)

	if err != nil {
		return nil, nil, err
	}

	_, err = neighbors.ToDevice(index.cuvsResource)

	if err != nil {
		return nil, nil, err
	}
	_, err = distances.ToDevice(index.cuvsResource)

	if err != nil {
		return nil, nil, err
	}

	params, err := cagra.CreateSearchParams()

	cagra.SearchIndex(*index.cuvsResource, params, index.cuvsIndex, &queries, &neighbors, &distances)

	neighbors.ToHost(index.cuvsResource)
	distances.ToHost(index.cuvsResource)

	neighborsSlice, err := neighbors.GetArray()

	if err != nil {
		return nil, nil, err
	}

	distancesSlice, err := distances.GetArray()

	if err != nil {
		return nil, nil, err
	}

	neighborsResultSlice := make([]uint64, k)

	for i := range neighborsSlice[0] {
		neighborsResultSlice[i] = index.idCuvsIdMap[neighborsSlice[0][i]]
	}

	return neighborsResultSlice, distancesSlice[0], nil

}

func (index *cuvs_index) initBuckets(ctx context.Context) error {
	if err := index.store.CreateOrLoadBucket(ctx, index.getBucketName(),
		// lsmkv.WithForceCompation(forceCompaction),
		lsmkv.WithUseBloomFilter(false),
		lsmkv.WithCalcCountNetAdditions(false),
	); err != nil {
		return fmt.Errorf("Create or load flat vectors bucket: %w", err)
	}

	return nil
}

func (index *cuvs_index) getBucketName() string {
	if index.targetVector != "" {
		return fmt.Sprintf("%s_%s", helpers.VectorsBucketLSM, index.targetVector)
	}
	return helpers.VectorsBucketLSM
}
func float32SliceFromByteSlice(vector []byte, slice []float32) []float32 {
	for i := range slice {
		slice[i] = math.Float32frombits(binary.LittleEndian.Uint32(vector[i*4:]))
	}
	return slice
}

func (index *cuvs_index) AlreadyIndexed() uint64 {
	return index.count
}

func (index *cuvs_index) PostStartup() {

	// cursor := index.store.Bucket(index.getBucketName()).Cursor()
	// defer cursor.Close()

	// // The initial size of 10k is chosen fairly arbitrarily. The cost of growing
	// // this slice dynamically should be quite cheap compared to other operations
	// // involved here, e.g. disk reads.
	// ids := make([]uint64, 0, 10_000)
	// vectors := make([][]float32, 0, 10_000)
	// maxID := uint64(0)

	// for key, v := cursor.First(); key != nil; key, v = cursor.Next() {
	// 	id := binary.BigEndian.Uint64(key)
	// 	// vecs = append(vecs, vec{
	// 	// 	id:  id,
	// 	// 	vec: uint64SliceFromByteSlice(v, make([]uint64, len(v)/8)),
	// 	// })
	// 	ids = append(ids, id)
	// 	vectors = append(vectors, float32SliceFromByteSlice(v, make([]float32, len(v)/4)))
	// 	if id > maxID {
	// 		maxID = id
	// 	}
	// }

	// index.AddBatch(context.Background(), ids, vectors)
}

func (index *cuvs_index) Dump(labels ...string) {

}

// searchbyvectordistance
func (index *cuvs_index) SearchByVectorDistance(vector []float32, dist float32, maxLimit int64, allowList helpers.AllowList) ([]uint64, []float32, error) {
	return []uint64{}, []float32{}, nil
}

func (index *cuvs_index) maxLimit() int64 {
	return 1_000_000_000
}

func (index *cuvs_index) UpdateUserConfig(updated schemaConfig.VectorIndexConfig, callback func()) error {
	return nil
}

func (index *cuvs_index) Drop(ctx context.Context) error {
	return nil
}

func (index *cuvs_index) Shutdown(ctx context.Context) error {
	return nil
}

func (index *cuvs_index) Flush() error {
	return nil
}

func (index *cuvs_index) SwitchCommitLogs(ctx context.Context) error {
	return nil
}

func (index *cuvs_index) ListFiles(ctx context.Context, basePath string) ([]string, error) {
	return []string{}, nil
}

func (index *cuvs_index) Compressed() bool {
	return false
}

func (index *cuvs_index) ValidateBeforeInsert(vector []float32) error {
	return nil
}

func (index *cuvs_index) DistanceBetweenVectors(x, y []float32) (float32, error) {
	return 0, nil
}

func (index *cuvs_index) ContainsNode(id uint64) bool {
	return false
}

func (index *cuvs_index) DistancerProvider() distancer.Provider {
	return nil
}

func (index *cuvs_index) QueryVectorDistancer(queryVector []float32) common.QueryVectorDistancer {
	distFunc := func(nodeID uint64) (float32, error) {

		return 0, nil
	}
	return common.QueryVectorDistancer{DistanceFunc: distFunc}
}
