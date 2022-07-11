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

package log_forward

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v4/models"
)

// GetLogForwardHandlerFunc turns a function with the right signature into a get log forward handler
type GetLogForwardHandlerFunc func(GetLogForwardParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLogForwardHandlerFunc) Handle(params GetLogForwardParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetLogForwardHandler interface for that can handle valid get log forward params
type GetLogForwardHandler interface {
	Handle(GetLogForwardParams, interface{}) middleware.Responder
}

// NewGetLogForward creates a new http.Handler for the get log forward operation
func NewGetLogForward(ctx *middleware.Context, handler GetLogForwardHandler) *GetLogForward {
	return &GetLogForward{Context: ctx, Handler: handler}
}

/* GetLogForward swagger:route GET /services/haproxy/configuration/log_forwards/{name} LogForward getLogForward

Return a log forward

Returns one log forward configuration by it's name.

*/
type GetLogForward struct {
	Context *middleware.Context
	Handler GetLogForwardHandler
}

func (o *GetLogForward) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetLogForwardParams()
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

// GetLogForwardOKBody get log forward o k body
//
// swagger:model GetLogForwardOKBody
type GetLogForwardOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.LogForward `json:"data,omitempty"`
}

// Validate validates this get log forward o k body
func (o *GetLogForwardOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLogForwardOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLogForwardOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLogForwardOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get log forward o k body based on the context it is used
func (o *GetLogForwardOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLogForwardOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLogForwardOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLogForwardOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLogForwardOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLogForwardOKBody) UnmarshalBinary(b []byte) error {
	var res GetLogForwardOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
