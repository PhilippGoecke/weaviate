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

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupConfig Backup custom configuration
//
// swagger:model BackupConfig
type BackupConfig struct {

	// Desired CPU core utilization ranging from 1%-80%
	// Maximum: 80
	// Minimum: 1
	CPUPercentage int64 `json:"CPUPercentage,omitempty"`

	// Weaviate will attempt to come close the specified size, with a minimum of 2MB, default of 128MB, and a maximum of 512MB
	// Maximum: 512
	// Minimum: 2
	ChunkSize int64 `json:"ChunkSize,omitempty"`

	// compression level used by compression algorithm
	// Enum: [DefaultCompression BestSpeed BestCompression]
	CompressionLevel string `json:"CompressionLevel,omitempty"`

	// S3 endpoint, e.g. s3.amazonaws.com
	Endpoint string `json:"Endpoint,omitempty"`

	// Name of the S3 bucket
	S3Bucket string `json:"S3Bucket,omitempty"`

	// Path within the bucket
	S3Path string `json:"S3Path,omitempty"`
}

// Validate validates this backup config
func (m *BackupConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUPercentage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChunkSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompressionLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupConfig) validateCPUPercentage(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUPercentage) { // not required
		return nil
	}

	if err := validate.MinimumInt("CPUPercentage", "body", m.CPUPercentage, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("CPUPercentage", "body", m.CPUPercentage, 80, false); err != nil {
		return err
	}

	return nil
}

func (m *BackupConfig) validateChunkSize(formats strfmt.Registry) error {
	if swag.IsZero(m.ChunkSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("ChunkSize", "body", m.ChunkSize, 2, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("ChunkSize", "body", m.ChunkSize, 512, false); err != nil {
		return err
	}

	return nil
}

var backupConfigTypeCompressionLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DefaultCompression","BestSpeed","BestCompression"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupConfigTypeCompressionLevelPropEnum = append(backupConfigTypeCompressionLevelPropEnum, v)
	}
}

const (

	// BackupConfigCompressionLevelDefaultCompression captures enum value "DefaultCompression"
	BackupConfigCompressionLevelDefaultCompression string = "DefaultCompression"

	// BackupConfigCompressionLevelBestSpeed captures enum value "BestSpeed"
	BackupConfigCompressionLevelBestSpeed string = "BestSpeed"

	// BackupConfigCompressionLevelBestCompression captures enum value "BestCompression"
	BackupConfigCompressionLevelBestCompression string = "BestCompression"
)

// prop value enum
func (m *BackupConfig) validateCompressionLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, backupConfigTypeCompressionLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BackupConfig) validateCompressionLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.CompressionLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateCompressionLevelEnum("CompressionLevel", "body", m.CompressionLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup config based on context it is used
func (m *BackupConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupConfig) UnmarshalBinary(b []byte) error {
	var res BackupConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
