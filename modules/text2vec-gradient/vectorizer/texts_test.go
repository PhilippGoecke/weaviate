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

package vectorizer

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// as used in the nearText searcher
func TestVectorizingTexts(t *testing.T) {
	type testCase struct {
		name                  string
		input                 []string
		expectedClientCall    string
		expectedApiEndpoint   string
		expectedProjectID     string
		expectedEndpointID    string
		expectedGradientModel string
		gradientModel         string
	}

	tests := []testCase{
		{
			name:                  "single word",
			input:                 []string{"hello"},
			gradientModel:         "large",
			expectedGradientModel: "large",
			expectedClientCall:    "hello",
		},
		{
			name:                  "multiple words",
			input:                 []string{"hello world, this is me!"},
			gradientModel:         "large",
			expectedGradientModel: "large",
			expectedClientCall:    "hello world, this is me!",
		},
		{
			name:                  "multiple sentences (joined with a dot)",
			input:                 []string{"this is sentence 1", "and here's number 2"},
			gradientModel:         "large",
			expectedGradientModel: "large",
			expectedClientCall:    "this is sentence 1. and here's number 2",
		},
		{
			name:                  "multiple sentences already containing a dot",
			input:                 []string{"this is sentence 1.", "and here's number 2"},
			gradientModel:         "large",
			expectedGradientModel: "large",
			expectedClientCall:    "this is sentence 1. and here's number 2",
		},
		{
			name:                  "multiple sentences already containing a question mark",
			input:                 []string{"this is sentence 1?", "and here's number 2"},
			gradientModel:         "large",
			expectedGradientModel: "large",
			expectedClientCall:    "this is sentence 1? and here's number 2",
		},
		{
			name:                  "multiple sentences already containing an exclamation mark",
			input:                 []string{"this is sentence 1!", "and here's number 2"},
			gradientModel:         "large",
			expectedGradientModel: "large",
			expectedClientCall:    "this is sentence 1! and here's number 2",
		},
		{
			name:                  "multiple sentences already containing comma",
			input:                 []string{"this is sentence 1,", "and here's number 2"},
			gradientModel:         "large",
			expectedGradientModel: "large",
			expectedClientCall:    "this is sentence 1, and here's number 2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := &fakeClient{}

			v := New(client)

			settings := &fakeSettings{
				apiEndpoint: "",
				projectID:   "",
				endpointID:  "",
			}
			vec, err := v.Texts(context.Background(), test.input, settings)

			require.Nil(t, err)
			assert.Equal(t, []float32{0.1, 1.1, 2.1, 3.1}, vec)
			assert.Equal(t, test.expectedClientCall, strings.Join(client.lastInput, ","))
			assert.Equal(t, test.expectedApiEndpoint, client.lastConfig.ApiEndpoint)
			assert.Equal(t, test.expectedProjectID, client.lastConfig.ProjectID)
			assert.Equal(t, test.expectedEndpointID, client.lastConfig.Model)
		})
	}
}
