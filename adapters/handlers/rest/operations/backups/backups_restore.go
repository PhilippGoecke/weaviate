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

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// BackupsRestoreHandlerFunc turns a function with the right signature into a backups restore handler
type BackupsRestoreHandlerFunc func(BackupsRestoreParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn BackupsRestoreHandlerFunc) Handle(params BackupsRestoreParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// BackupsRestoreHandler interface for that can handle valid backups restore params
type BackupsRestoreHandler interface {
	Handle(BackupsRestoreParams, *models.Principal) middleware.Responder
}

// NewBackupsRestore creates a new http.Handler for the backups restore operation
func NewBackupsRestore(ctx *middleware.Context, handler BackupsRestoreHandler) *BackupsRestore {
	return &BackupsRestore{Context: ctx, Handler: handler}
}

/*
	BackupsRestore swagger:route POST /backups/{backend}/{id}/restore backups backupsRestore

# Start a restoration process

Starts a process of restoring a backup for a set of collections. <br/><br/>Any backup can be restored to any machine, as long as the number of nodes between source and target are identical.<br/><br/>Requrements:<br/><br/>- None of the collections to be restored already exist on the target restoration node(s).<br/>- The node names of the backed-up collections' must match those of the target restoration node(s).
*/
type BackupsRestore struct {
	Context *middleware.Context
	Handler BackupsRestoreHandler
}

func (o *BackupsRestore) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	Params := NewBackupsRestoreParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)
}
