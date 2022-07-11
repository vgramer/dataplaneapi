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

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetHAProxyConfigurationOKCode is the HTTP code returned for type GetHAProxyConfigurationOK
const GetHAProxyConfigurationOKCode int = 200

/*GetHAProxyConfigurationOK Operation successful

swagger:response getHAProxyConfigurationOK
*/
type GetHAProxyConfigurationOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetHAProxyConfigurationOKBody `json:"body,omitempty"`
}

// NewGetHAProxyConfigurationOK creates GetHAProxyConfigurationOK with default headers values
func NewGetHAProxyConfigurationOK() *GetHAProxyConfigurationOK {

	return &GetHAProxyConfigurationOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get h a proxy configuration o k response
func (o *GetHAProxyConfigurationOK) WithConfigurationVersion(configurationVersion string) *GetHAProxyConfigurationOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get h a proxy configuration o k response
func (o *GetHAProxyConfigurationOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get h a proxy configuration o k response
func (o *GetHAProxyConfigurationOK) WithPayload(payload *GetHAProxyConfigurationOKBody) *GetHAProxyConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get h a proxy configuration o k response
func (o *GetHAProxyConfigurationOK) SetPayload(payload *GetHAProxyConfigurationOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHAProxyConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetHAProxyConfigurationDefault General Error

swagger:response getHAProxyConfigurationDefault
*/
type GetHAProxyConfigurationDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHAProxyConfigurationDefault creates GetHAProxyConfigurationDefault with default headers values
func NewGetHAProxyConfigurationDefault(code int) *GetHAProxyConfigurationDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHAProxyConfigurationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get h a proxy configuration default response
func (o *GetHAProxyConfigurationDefault) WithStatusCode(code int) *GetHAProxyConfigurationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get h a proxy configuration default response
func (o *GetHAProxyConfigurationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get h a proxy configuration default response
func (o *GetHAProxyConfigurationDefault) WithConfigurationVersion(configurationVersion string) *GetHAProxyConfigurationDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get h a proxy configuration default response
func (o *GetHAProxyConfigurationDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get h a proxy configuration default response
func (o *GetHAProxyConfigurationDefault) WithPayload(payload *models.Error) *GetHAProxyConfigurationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get h a proxy configuration default response
func (o *GetHAProxyConfigurationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHAProxyConfigurationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
