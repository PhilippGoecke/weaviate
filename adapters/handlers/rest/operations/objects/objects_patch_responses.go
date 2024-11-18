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

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsPatchNoContentCode is the HTTP code returned for type ObjectsPatchNoContent
const ObjectsPatchNoContentCode int = 204

/*
ObjectsPatchNoContent Successfully applied. No content provided.

swagger:response objectsPatchNoContent
*/
type ObjectsPatchNoContent struct{}

// NewObjectsPatchNoContent creates ObjectsPatchNoContent with default headers values
func NewObjectsPatchNoContent() *ObjectsPatchNoContent {
	return &ObjectsPatchNoContent{}
}

// WriteResponse to the client
func (o *ObjectsPatchNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ObjectsPatchBadRequestCode is the HTTP code returned for type ObjectsPatchBadRequest
const ObjectsPatchBadRequestCode int = 400

/*
ObjectsPatchBadRequest The patch-JSON is malformed.

swagger:response objectsPatchBadRequest
*/
type ObjectsPatchBadRequest struct{}

// NewObjectsPatchBadRequest creates ObjectsPatchBadRequest with default headers values
func NewObjectsPatchBadRequest() *ObjectsPatchBadRequest {
	return &ObjectsPatchBadRequest{}
}

// WriteResponse to the client
func (o *ObjectsPatchBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ObjectsPatchUnauthorizedCode is the HTTP code returned for type ObjectsPatchUnauthorized
const ObjectsPatchUnauthorizedCode int = 401

/*
ObjectsPatchUnauthorized Unauthorized or invalid credentials.

swagger:response objectsPatchUnauthorized
*/
type ObjectsPatchUnauthorized struct{}

// NewObjectsPatchUnauthorized creates ObjectsPatchUnauthorized with default headers values
func NewObjectsPatchUnauthorized() *ObjectsPatchUnauthorized {
	return &ObjectsPatchUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsPatchUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsPatchForbiddenCode is the HTTP code returned for type ObjectsPatchForbidden
const ObjectsPatchForbiddenCode int = 403

/*
ObjectsPatchForbidden Forbidden

swagger:response objectsPatchForbidden
*/
type ObjectsPatchForbidden struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsPatchForbidden creates ObjectsPatchForbidden with default headers values
func NewObjectsPatchForbidden() *ObjectsPatchForbidden {
	return &ObjectsPatchForbidden{}
}

// WithPayload adds the payload to the objects patch forbidden response
func (o *ObjectsPatchForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsPatchForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects patch forbidden response
func (o *ObjectsPatchForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsPatchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsPatchNotFoundCode is the HTTP code returned for type ObjectsPatchNotFound
const ObjectsPatchNotFoundCode int = 404

/*
ObjectsPatchNotFound Successful query result but no resource was found.

swagger:response objectsPatchNotFound
*/
type ObjectsPatchNotFound struct{}

// NewObjectsPatchNotFound creates ObjectsPatchNotFound with default headers values
func NewObjectsPatchNotFound() *ObjectsPatchNotFound {
	return &ObjectsPatchNotFound{}
}

// WriteResponse to the client
func (o *ObjectsPatchNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ObjectsPatchUnprocessableEntityCode is the HTTP code returned for type ObjectsPatchUnprocessableEntity
const ObjectsPatchUnprocessableEntityCode int = 422

/*
ObjectsPatchUnprocessableEntity The patch-JSON is valid but unprocessable.

swagger:response objectsPatchUnprocessableEntity
*/
type ObjectsPatchUnprocessableEntity struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsPatchUnprocessableEntity creates ObjectsPatchUnprocessableEntity with default headers values
func NewObjectsPatchUnprocessableEntity() *ObjectsPatchUnprocessableEntity {
	return &ObjectsPatchUnprocessableEntity{}
}

// WithPayload adds the payload to the objects patch unprocessable entity response
func (o *ObjectsPatchUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ObjectsPatchUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects patch unprocessable entity response
func (o *ObjectsPatchUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsPatchUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsPatchInternalServerErrorCode is the HTTP code returned for type ObjectsPatchInternalServerError
const ObjectsPatchInternalServerErrorCode int = 500

/*
ObjectsPatchInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsPatchInternalServerError
*/
type ObjectsPatchInternalServerError struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsPatchInternalServerError creates ObjectsPatchInternalServerError with default headers values
func NewObjectsPatchInternalServerError() *ObjectsPatchInternalServerError {
	return &ObjectsPatchInternalServerError{}
}

// WithPayload adds the payload to the objects patch internal server error response
func (o *ObjectsPatchInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsPatchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects patch internal server error response
func (o *ObjectsPatchInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsPatchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
