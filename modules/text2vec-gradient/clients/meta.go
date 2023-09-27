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

package clients

func (v *gradient) MetaInfo() (map[string]interface{}, error) {
	return map[string]interface{}{
		"name":              " Gradient Module",
		"documentationHref": "https://cloud..com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings",
	}, nil
}
