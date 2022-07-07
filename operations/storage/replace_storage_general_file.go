// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceStorageGeneralFileHandlerFunc turns a function with the right signature into a replace storage general file handler
type ReplaceStorageGeneralFileHandlerFunc func(ReplaceStorageGeneralFileParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceStorageGeneralFileHandlerFunc) Handle(params ReplaceStorageGeneralFileParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceStorageGeneralFileHandler interface for that can handle valid replace storage general file params
type ReplaceStorageGeneralFileHandler interface {
	Handle(ReplaceStorageGeneralFileParams, interface{}) middleware.Responder
}

// NewReplaceStorageGeneralFile creates a new http.Handler for the replace storage general file operation
func NewReplaceStorageGeneralFile(ctx *middleware.Context, handler ReplaceStorageGeneralFileHandler) *ReplaceStorageGeneralFile {
	return &ReplaceStorageGeneralFile{Context: ctx, Handler: handler}
}

/* ReplaceStorageGeneralFile swagger:route PUT /services/haproxy/storage/general/{name} Storage replaceStorageGeneralFile

Replace contents of a managed general use file on disk

Replaces the contents of a managed general use file on disk

*/
type ReplaceStorageGeneralFile struct {
	Context *middleware.Context
	Handler ReplaceStorageGeneralFileHandler
}

func (o *ReplaceStorageGeneralFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewReplaceStorageGeneralFileParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
