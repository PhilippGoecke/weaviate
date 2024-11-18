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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// SchemaObjectsShardsGetOKCode is the HTTP code returned for type SchemaObjectsShardsGetOK
const SchemaObjectsShardsGetOKCode int = 200

/*
SchemaObjectsShardsGetOK Found the status of the shards, returned as body

swagger:response schemaObjectsShardsGetOK
*/
type SchemaObjectsShardsGetOK struct {
	/*
	  In: Body
	*/
	Payload models.ShardStatusList `json:"body,omitempty"`
}

// NewSchemaObjectsShardsGetOK creates SchemaObjectsShardsGetOK with default headers values
func NewSchemaObjectsShardsGetOK() *SchemaObjectsShardsGetOK {
	return &SchemaObjectsShardsGetOK{}
}

// WithPayload adds the payload to the schema objects shards get o k response
func (o *SchemaObjectsShardsGetOK) WithPayload(payload models.ShardStatusList) *SchemaObjectsShardsGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects shards get o k response
func (o *SchemaObjectsShardsGetOK) SetPayload(payload models.ShardStatusList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsShardsGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ShardStatusList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SchemaObjectsShardsGetUnauthorizedCode is the HTTP code returned for type SchemaObjectsShardsGetUnauthorized
const SchemaObjectsShardsGetUnauthorizedCode int = 401

/*
SchemaObjectsShardsGetUnauthorized Unauthorized or invalid credentials.

swagger:response schemaObjectsShardsGetUnauthorized
*/
type SchemaObjectsShardsGetUnauthorized struct{}

// NewSchemaObjectsShardsGetUnauthorized creates SchemaObjectsShardsGetUnauthorized with default headers values
func NewSchemaObjectsShardsGetUnauthorized() *SchemaObjectsShardsGetUnauthorized {
	return &SchemaObjectsShardsGetUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaObjectsShardsGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaObjectsShardsGetForbiddenCode is the HTTP code returned for type SchemaObjectsShardsGetForbidden
const SchemaObjectsShardsGetForbiddenCode int = 403

/*
SchemaObjectsShardsGetForbidden Forbidden

swagger:response schemaObjectsShardsGetForbidden
*/
type SchemaObjectsShardsGetForbidden struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsShardsGetForbidden creates SchemaObjectsShardsGetForbidden with default headers values
func NewSchemaObjectsShardsGetForbidden() *SchemaObjectsShardsGetForbidden {
	return &SchemaObjectsShardsGetForbidden{}
}

// WithPayload adds the payload to the schema objects shards get forbidden response
func (o *SchemaObjectsShardsGetForbidden) WithPayload(payload *models.ErrorResponse) *SchemaObjectsShardsGetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects shards get forbidden response
func (o *SchemaObjectsShardsGetForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsShardsGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsShardsGetNotFoundCode is the HTTP code returned for type SchemaObjectsShardsGetNotFound
const SchemaObjectsShardsGetNotFoundCode int = 404

/*
SchemaObjectsShardsGetNotFound This class does not exist

swagger:response schemaObjectsShardsGetNotFound
*/
type SchemaObjectsShardsGetNotFound struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsShardsGetNotFound creates SchemaObjectsShardsGetNotFound with default headers values
func NewSchemaObjectsShardsGetNotFound() *SchemaObjectsShardsGetNotFound {
	return &SchemaObjectsShardsGetNotFound{}
}

// WithPayload adds the payload to the schema objects shards get not found response
func (o *SchemaObjectsShardsGetNotFound) WithPayload(payload *models.ErrorResponse) *SchemaObjectsShardsGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects shards get not found response
func (o *SchemaObjectsShardsGetNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsShardsGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsShardsGetInternalServerErrorCode is the HTTP code returned for type SchemaObjectsShardsGetInternalServerError
const SchemaObjectsShardsGetInternalServerErrorCode int = 500

/*
SchemaObjectsShardsGetInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaObjectsShardsGetInternalServerError
*/
type SchemaObjectsShardsGetInternalServerError struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsShardsGetInternalServerError creates SchemaObjectsShardsGetInternalServerError with default headers values
func NewSchemaObjectsShardsGetInternalServerError() *SchemaObjectsShardsGetInternalServerError {
	return &SchemaObjectsShardsGetInternalServerError{}
}

// WithPayload adds the payload to the schema objects shards get internal server error response
func (o *SchemaObjectsShardsGetInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaObjectsShardsGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects shards get internal server error response
func (o *SchemaObjectsShardsGetInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsShardsGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
