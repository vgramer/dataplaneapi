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

package spoe_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetSpoeTransactionParams creates a new GetSpoeTransactionParams object
//
// There are no default values defined in the spec.
func NewGetSpoeTransactionParams() GetSpoeTransactionParams {

	return GetSpoeTransactionParams{}
}

// GetSpoeTransactionParams contains all the bound params for the get spoe transaction operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSpoeTransaction
type GetSpoeTransactionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Transaction id
	  Required: true
	  In: path
	*/
	ID string
	/*Spoe file name
	  Required: true
	  In: query
	*/
	Spoe string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSpoeTransactionParams() beforehand.
func (o *GetSpoeTransactionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qSpoe, qhkSpoe, _ := qs.GetOK("spoe")
	if err := o.bindSpoe(qSpoe, qhkSpoe, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetSpoeTransactionParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	return nil
}

// bindSpoe binds and validates parameter Spoe from query.
func (o *GetSpoeTransactionParams) bindSpoe(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("spoe", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("spoe", "query", raw); err != nil {
		return err
	}
	o.Spoe = raw

	return nil
}
