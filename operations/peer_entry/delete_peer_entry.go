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

package peer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeletePeerEntryHandlerFunc turns a function with the right signature into a delete peer entry handler
type DeletePeerEntryHandlerFunc func(DeletePeerEntryParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePeerEntryHandlerFunc) Handle(params DeletePeerEntryParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeletePeerEntryHandler interface for that can handle valid delete peer entry params
type DeletePeerEntryHandler interface {
	Handle(DeletePeerEntryParams, interface{}) middleware.Responder
}

// NewDeletePeerEntry creates a new http.Handler for the delete peer entry operation
func NewDeletePeerEntry(ctx *middleware.Context, handler DeletePeerEntryHandler) *DeletePeerEntry {
	return &DeletePeerEntry{Context: ctx, Handler: handler}
}

/* DeletePeerEntry swagger:route DELETE /services/haproxy/configuration/peer_entries/{name} PeerEntry deletePeerEntry

Delete a peer_entry

Deletes a peer entry configuration by it's name in the specified peer section.

*/
type DeletePeerEntry struct {
	Context *middleware.Context
	Handler DeletePeerEntryHandler
}

func (o *DeletePeerEntry) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeletePeerEntryParams()
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
