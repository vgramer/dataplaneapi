// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies LLC
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

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateACLHandlerFunc turns a function with the right signature into a create Acl handler
type CreateACLHandlerFunc func(CreateACLParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateACLHandlerFunc) Handle(params CreateACLParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateACLHandler interface for that can handle valid create Acl params
type CreateACLHandler interface {
	Handle(CreateACLParams, interface{}) middleware.Responder
}

// NewCreateACL creates a new http.Handler for the create Acl operation
func NewCreateACL(ctx *middleware.Context, handler CreateACLHandler) *CreateACL {
	return &CreateACL{Context: ctx, Handler: handler}
}

/*CreateACL swagger:route POST /services/haproxy/configuration/acls ACL createAcl

Add a new ACL line

Adds a new ACL line of the specified type in the specified parent.

*/
type CreateACL struct {
	Context *middleware.Context
	Handler CreateACLHandler
}

func (o *CreateACL) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateACLParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}