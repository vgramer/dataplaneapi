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

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceBackendSwitchingRuleHandlerFunc turns a function with the right signature into a replace backend switching rule handler
type ReplaceBackendSwitchingRuleHandlerFunc func(ReplaceBackendSwitchingRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceBackendSwitchingRuleHandlerFunc) Handle(params ReplaceBackendSwitchingRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceBackendSwitchingRuleHandler interface for that can handle valid replace backend switching rule params
type ReplaceBackendSwitchingRuleHandler interface {
	Handle(ReplaceBackendSwitchingRuleParams, interface{}) middleware.Responder
}

// NewReplaceBackendSwitchingRule creates a new http.Handler for the replace backend switching rule operation
func NewReplaceBackendSwitchingRule(ctx *middleware.Context, handler ReplaceBackendSwitchingRuleHandler) *ReplaceBackendSwitchingRule {
	return &ReplaceBackendSwitchingRule{Context: ctx, Handler: handler}
}

/* ReplaceBackendSwitchingRule swagger:route PUT /services/haproxy/configuration/backend_switching_rules/{index} BackendSwitchingRule replaceBackendSwitchingRule

Replace a Backend Switching Rule

Replaces a Backend Switching Rule configuration by it's index in the specified frontend.

*/
type ReplaceBackendSwitchingRule struct {
	Context *middleware.Context
	Handler ReplaceBackendSwitchingRuleHandler
}

func (o *ReplaceBackendSwitchingRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewReplaceBackendSwitchingRuleParams()
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
