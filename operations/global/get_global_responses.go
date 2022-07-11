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

package global

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetGlobalOKCode is the HTTP code returned for type GetGlobalOK
const GetGlobalOKCode int = 200

/*GetGlobalOK Successful operation

swagger:response getGlobalOK
*/
type GetGlobalOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetGlobalOKBody `json:"body,omitempty"`
}

// NewGetGlobalOK creates GetGlobalOK with default headers values
func NewGetGlobalOK() *GetGlobalOK {

	return &GetGlobalOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get global o k response
func (o *GetGlobalOK) WithConfigurationVersion(configurationVersion string) *GetGlobalOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get global o k response
func (o *GetGlobalOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get global o k response
func (o *GetGlobalOK) WithPayload(payload *GetGlobalOKBody) *GetGlobalOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get global o k response
func (o *GetGlobalOK) SetPayload(payload *GetGlobalOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGlobalOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetGlobalDefault General Error

swagger:response getGlobalDefault
*/
type GetGlobalDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetGlobalDefault creates GetGlobalDefault with default headers values
func NewGetGlobalDefault(code int) *GetGlobalDefault {
	if code <= 0 {
		code = 500
	}

	return &GetGlobalDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get global default response
func (o *GetGlobalDefault) WithStatusCode(code int) *GetGlobalDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get global default response
func (o *GetGlobalDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get global default response
func (o *GetGlobalDefault) WithConfigurationVersion(configurationVersion string) *GetGlobalDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get global default response
func (o *GetGlobalDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get global default response
func (o *GetGlobalDefault) WithPayload(payload *models.Error) *GetGlobalDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get global default response
func (o *GetGlobalDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGlobalDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
