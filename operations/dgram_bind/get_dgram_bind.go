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

package dgram_bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v3/models"
)

// GetDgramBindHandlerFunc turns a function with the right signature into a get dgram bind handler
type GetDgramBindHandlerFunc func(GetDgramBindParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDgramBindHandlerFunc) Handle(params GetDgramBindParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDgramBindHandler interface for that can handle valid get dgram bind params
type GetDgramBindHandler interface {
	Handle(GetDgramBindParams, interface{}) middleware.Responder
}

// NewGetDgramBind creates a new http.Handler for the get dgram bind operation
func NewGetDgramBind(ctx *middleware.Context, handler GetDgramBindHandler) *GetDgramBind {
	return &GetDgramBind{Context: ctx, Handler: handler}
}

/* GetDgramBind swagger:route GET /services/haproxy/configuration/dgram_binds/{name} DgramBind getDgramBind

Return one dgram bind

Returns one dgram bind configuration by it's name in the specified log forward.

*/
type GetDgramBind struct {
	Context *middleware.Context
	Handler GetDgramBindHandler
}

func (o *GetDgramBind) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDgramBindParams()
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

// GetDgramBindOKBody get dgram bind o k body
//
// swagger:model GetDgramBindOKBody
type GetDgramBindOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.DgramBind `json:"data,omitempty"`
}

// Validate validates this get dgram bind o k body
func (o *GetDgramBindOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDgramBindOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDgramBindOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getDgramBindOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get dgram bind o k body based on the context it is used
func (o *GetDgramBindOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDgramBindOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDgramBindOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getDgramBindOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDgramBindOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDgramBindOKBody) UnmarshalBinary(b []byte) error {
	var res GetDgramBindOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
