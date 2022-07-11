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

package stick_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetStickTablesOKCode is the HTTP code returned for type GetStickTablesOK
const GetStickTablesOKCode int = 200

/*GetStickTablesOK Successful operation

swagger:response getStickTablesOK
*/
type GetStickTablesOK struct {

	/*
	  In: Body
	*/
	Payload models.StickTables `json:"body,omitempty"`
}

// NewGetStickTablesOK creates GetStickTablesOK with default headers values
func NewGetStickTablesOK() *GetStickTablesOK {

	return &GetStickTablesOK{}
}

// WithPayload adds the payload to the get stick tables o k response
func (o *GetStickTablesOK) WithPayload(payload models.StickTables) *GetStickTablesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stick tables o k response
func (o *GetStickTablesOK) SetPayload(payload models.StickTables) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStickTablesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.StickTables{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetStickTablesDefault General Error

swagger:response getStickTablesDefault
*/
type GetStickTablesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStickTablesDefault creates GetStickTablesDefault with default headers values
func NewGetStickTablesDefault(code int) *GetStickTablesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetStickTablesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get stick tables default response
func (o *GetStickTablesDefault) WithStatusCode(code int) *GetStickTablesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get stick tables default response
func (o *GetStickTablesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get stick tables default response
func (o *GetStickTablesDefault) WithConfigurationVersion(configurationVersion string) *GetStickTablesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get stick tables default response
func (o *GetStickTablesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get stick tables default response
func (o *GetStickTablesDefault) WithPayload(payload *models.Error) *GetStickTablesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stick tables default response
func (o *GetStickTablesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStickTablesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
