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

package service_discovery

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

// GetAWSRegionHandlerFunc turns a function with the right signature into a get a w s region handler
type GetAWSRegionHandlerFunc func(GetAWSRegionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAWSRegionHandlerFunc) Handle(params GetAWSRegionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAWSRegionHandler interface for that can handle valid get a w s region params
type GetAWSRegionHandler interface {
	Handle(GetAWSRegionParams, interface{}) middleware.Responder
}

// NewGetAWSRegion creates a new http.Handler for the get a w s region operation
func NewGetAWSRegion(ctx *middleware.Context, handler GetAWSRegionHandler) *GetAWSRegion {
	return &GetAWSRegion{Context: ctx, Handler: handler}
}

/* GetAWSRegion swagger:route GET /service_discovery/aws/{id} ServiceDiscovery getAWSRegion

Return an AWS region

Return one AWS Region configuration by it's id.

*/
type GetAWSRegion struct {
	Context *middleware.Context
	Handler GetAWSRegionHandler
}

func (o *GetAWSRegion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAWSRegionParams()
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

// GetAWSRegionOKBody get a w s region o k body
//
// swagger:model GetAWSRegionOKBody
type GetAWSRegionOKBody struct {

	// data
	Data *models.AwsRegion `json:"data,omitempty"`
}

// Validate validates this get a w s region o k body
func (o *GetAWSRegionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAWSRegionOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAWSRegionOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getAWSRegionOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get a w s region o k body based on the context it is used
func (o *GetAWSRegionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAWSRegionOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAWSRegionOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getAWSRegionOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAWSRegionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAWSRegionOKBody) UnmarshalBinary(b []byte) error {
	var res GetAWSRegionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
