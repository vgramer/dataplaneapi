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

package userlist

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

	"github.com/haproxytech/client-native/v4/models"
)

// GetUserlistsHandlerFunc turns a function with the right signature into a get userlists handler
type GetUserlistsHandlerFunc func(GetUserlistsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserlistsHandlerFunc) Handle(params GetUserlistsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetUserlistsHandler interface for that can handle valid get userlists params
type GetUserlistsHandler interface {
	Handle(GetUserlistsParams, interface{}) middleware.Responder
}

// NewGetUserlists creates a new http.Handler for the get userlists operation
func NewGetUserlists(ctx *middleware.Context, handler GetUserlistsHandler) *GetUserlists {
	return &GetUserlists{Context: ctx, Handler: handler}
}

/* GetUserlists swagger:route GET /services/haproxy/configuration/userlists Userlist getUserlists

Return an array of userlists

Returns an array of all configured userlists.

*/
type GetUserlists struct {
	Context *middleware.Context
	Handler GetUserlistsHandler
}

func (o *GetUserlists) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserlistsParams()
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

// GetUserlistsOKBody get userlists o k body
//
// swagger:model GetUserlistsOKBody
type GetUserlistsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Userlists `json:"data"`
}

// Validate validates this get userlists o k body
func (o *GetUserlistsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserlistsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getUserlistsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getUserlistsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getUserlistsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get userlists o k body based on the context it is used
func (o *GetUserlistsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserlistsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getUserlistsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getUserlistsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserlistsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserlistsOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserlistsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
