//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package modgradient

import (
	"github.com/weaviate/weaviate/entities/modulecapabilities"
	"github.com/weaviate/weaviate/modules/text2vec-gradient/neartext"
)

func (m *GradientModule) initNearText() error {
	m.searcher = neartext.NewSearcher(m.vectorizer)
	m.graphqlProvider = neartext.New(m.nearTextTransformer)
	return nil
}

func (m *GradientModule) Arguments() map[string]modulecapabilities.GraphQLArgument {
	return m.graphqlProvider.Arguments()
}

func (m *GradientModule) VectorSearches() map[string]modulecapabilities.VectorForParams {
	return m.searcher.VectorSearches()
}

var (
	_ = modulecapabilities.GraphQLArguments(New())
	_ = modulecapabilities.Searcher(New())
)
