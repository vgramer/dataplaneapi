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

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v3/models"
)

// GetServerSwitchingRulesHandlerFunc turns a function with the right signature into a get server switching rules handler
type GetServerSwitchingRulesHandlerFunc func(GetServerSwitchingRulesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetServerSwitchingRulesHandlerFunc) Handle(params GetServerSwitchingRulesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetServerSwitchingRulesHandler interface for that can handle valid get server switching rules params
type GetServerSwitchingRulesHandler interface {
	Handle(GetServerSwitchingRulesParams, interface{}) middleware.Responder
}

// NewGetServerSwitchingRules creates a new http.Handler for the get server switching rules operation
func NewGetServerSwitchingRules(ctx *middleware.Context, handler GetServerSwitchingRulesHandler) *GetServerSwitchingRules {
	return &GetServerSwitchingRules{Context: ctx, Handler: handler}
}

/* GetServerSwitchingRules swagger:route GET /services/haproxy/configuration/server_switching_rules ServerSwitchingRule getServerSwitchingRules

Return an array of all Server Switching Rules

Returns all Backend Switching Rules that are configured in specified backend.

*/
type GetServerSwitchingRules struct {
	Context *middleware.Context
	Handler GetServerSwitchingRulesHandler
}

func (o *GetServerSwitchingRules) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetServerSwitchingRulesParams()
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

// GetServerSwitchingRulesOKBody get server switching rules o k body
//
// swagger:model GetServerSwitchingRulesOKBody
type GetServerSwitchingRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.ServerSwitchingRules `json:"data"`
}

// Validate validates this get server switching rules o k body
func (o *GetServerSwitchingRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerSwitchingRulesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServerSwitchingRulesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getServerSwitchingRulesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getServerSwitchingRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get server switching rules o k body based on the context it is used
func (o *GetServerSwitchingRulesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerSwitchingRulesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getServerSwitchingRulesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getServerSwitchingRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServerSwitchingRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServerSwitchingRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetServerSwitchingRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
