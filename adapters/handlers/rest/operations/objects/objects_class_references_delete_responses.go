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

// ObjectsClassReferencesDeleteNoContentCode is the HTTP code returned for type ObjectsClassReferencesDeleteNoContent
const ObjectsClassReferencesDeleteNoContentCode int = 204

/*
ObjectsClassReferencesDeleteNoContent Successfully deleted.

swagger:response objectsClassReferencesDeleteNoContent
*/
type ObjectsClassReferencesDeleteNoContent struct{}

// NewObjectsClassReferencesDeleteNoContent creates ObjectsClassReferencesDeleteNoContent with default headers values
func NewObjectsClassReferencesDeleteNoContent() *ObjectsClassReferencesDeleteNoContent {
	return &ObjectsClassReferencesDeleteNoContent{}
}

// WriteResponse to the client
func (o *ObjectsClassReferencesDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ObjectsClassReferencesDeleteBadRequestCode is the HTTP code returned for type ObjectsClassReferencesDeleteBadRequest
const ObjectsClassReferencesDeleteBadRequestCode int = 400

/*
ObjectsClassReferencesDeleteBadRequest Malformed request.

swagger:response objectsClassReferencesDeleteBadRequest
*/
type ObjectsClassReferencesDeleteBadRequest struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassReferencesDeleteBadRequest creates ObjectsClassReferencesDeleteBadRequest with default headers values
func NewObjectsClassReferencesDeleteBadRequest() *ObjectsClassReferencesDeleteBadRequest {
	return &ObjectsClassReferencesDeleteBadRequest{}
}

// WithPayload adds the payload to the objects class references delete bad request response
func (o *ObjectsClassReferencesDeleteBadRequest) WithPayload(payload *models.ErrorResponse) *ObjectsClassReferencesDeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class references delete bad request response
func (o *ObjectsClassReferencesDeleteBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassReferencesDeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassReferencesDeleteUnauthorizedCode is the HTTP code returned for type ObjectsClassReferencesDeleteUnauthorized
const ObjectsClassReferencesDeleteUnauthorizedCode int = 401

/*
ObjectsClassReferencesDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response objectsClassReferencesDeleteUnauthorized
*/
type ObjectsClassReferencesDeleteUnauthorized struct{}

// NewObjectsClassReferencesDeleteUnauthorized creates ObjectsClassReferencesDeleteUnauthorized with default headers values
func NewObjectsClassReferencesDeleteUnauthorized() *ObjectsClassReferencesDeleteUnauthorized {
	return &ObjectsClassReferencesDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsClassReferencesDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsClassReferencesDeleteForbiddenCode is the HTTP code returned for type ObjectsClassReferencesDeleteForbidden
const ObjectsClassReferencesDeleteForbiddenCode int = 403

/*
ObjectsClassReferencesDeleteForbidden Forbidden

swagger:response objectsClassReferencesDeleteForbidden
*/
type ObjectsClassReferencesDeleteForbidden struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassReferencesDeleteForbidden creates ObjectsClassReferencesDeleteForbidden with default headers values
func NewObjectsClassReferencesDeleteForbidden() *ObjectsClassReferencesDeleteForbidden {
	return &ObjectsClassReferencesDeleteForbidden{}
}

// WithPayload adds the payload to the objects class references delete forbidden response
func (o *ObjectsClassReferencesDeleteForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsClassReferencesDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class references delete forbidden response
func (o *ObjectsClassReferencesDeleteForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassReferencesDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassReferencesDeleteNotFoundCode is the HTTP code returned for type ObjectsClassReferencesDeleteNotFound
const ObjectsClassReferencesDeleteNotFoundCode int = 404

/*
ObjectsClassReferencesDeleteNotFound Successful query result but no resource was found.

swagger:response objectsClassReferencesDeleteNotFound
*/
type ObjectsClassReferencesDeleteNotFound struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassReferencesDeleteNotFound creates ObjectsClassReferencesDeleteNotFound with default headers values
func NewObjectsClassReferencesDeleteNotFound() *ObjectsClassReferencesDeleteNotFound {
	return &ObjectsClassReferencesDeleteNotFound{}
}

// WithPayload adds the payload to the objects class references delete not found response
func (o *ObjectsClassReferencesDeleteNotFound) WithPayload(payload *models.ErrorResponse) *ObjectsClassReferencesDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class references delete not found response
func (o *ObjectsClassReferencesDeleteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassReferencesDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassReferencesDeleteUnprocessableEntityCode is the HTTP code returned for type ObjectsClassReferencesDeleteUnprocessableEntity
const ObjectsClassReferencesDeleteUnprocessableEntityCode int = 422

/*
ObjectsClassReferencesDeleteUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?

swagger:response objectsClassReferencesDeleteUnprocessableEntity
*/
type ObjectsClassReferencesDeleteUnprocessableEntity struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassReferencesDeleteUnprocessableEntity creates ObjectsClassReferencesDeleteUnprocessableEntity with default headers values
func NewObjectsClassReferencesDeleteUnprocessableEntity() *ObjectsClassReferencesDeleteUnprocessableEntity {
	return &ObjectsClassReferencesDeleteUnprocessableEntity{}
}

// WithPayload adds the payload to the objects class references delete unprocessable entity response
func (o *ObjectsClassReferencesDeleteUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ObjectsClassReferencesDeleteUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class references delete unprocessable entity response
func (o *ObjectsClassReferencesDeleteUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassReferencesDeleteUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassReferencesDeleteInternalServerErrorCode is the HTTP code returned for type ObjectsClassReferencesDeleteInternalServerError
const ObjectsClassReferencesDeleteInternalServerErrorCode int = 500

/*
ObjectsClassReferencesDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsClassReferencesDeleteInternalServerError
*/
type ObjectsClassReferencesDeleteInternalServerError struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassReferencesDeleteInternalServerError creates ObjectsClassReferencesDeleteInternalServerError with default headers values
func NewObjectsClassReferencesDeleteInternalServerError() *ObjectsClassReferencesDeleteInternalServerError {
	return &ObjectsClassReferencesDeleteInternalServerError{}
}

// WithPayload adds the payload to the objects class references delete internal server error response
func (o *ObjectsClassReferencesDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsClassReferencesDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class references delete internal server error response
func (o *ObjectsClassReferencesDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassReferencesDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
