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

package tcp_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetTCPRequestRulesOKCode is the HTTP code returned for type GetTCPRequestRulesOK
const GetTCPRequestRulesOKCode int = 200

/*GetTCPRequestRulesOK Successful operation

swagger:response getTcpRequestRulesOK
*/
type GetTCPRequestRulesOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetTCPRequestRulesOKBody `json:"body,omitempty"`
}

// NewGetTCPRequestRulesOK creates GetTCPRequestRulesOK with default headers values
func NewGetTCPRequestRulesOK() *GetTCPRequestRulesOK {

	return &GetTCPRequestRulesOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get Tcp request rules o k response
func (o *GetTCPRequestRulesOK) WithConfigurationVersion(configurationVersion string) *GetTCPRequestRulesOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get Tcp request rules o k response
func (o *GetTCPRequestRulesOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get Tcp request rules o k response
func (o *GetTCPRequestRulesOK) WithPayload(payload *GetTCPRequestRulesOKBody) *GetTCPRequestRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Tcp request rules o k response
func (o *GetTCPRequestRulesOK) SetPayload(payload *GetTCPRequestRulesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPRequestRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetTCPRequestRulesDefault General Error

swagger:response getTcpRequestRulesDefault
*/
type GetTCPRequestRulesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTCPRequestRulesDefault creates GetTCPRequestRulesDefault with default headers values
func NewGetTCPRequestRulesDefault(code int) *GetTCPRequestRulesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTCPRequestRulesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get TCP request rules default response
func (o *GetTCPRequestRulesDefault) WithStatusCode(code int) *GetTCPRequestRulesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get TCP request rules default response
func (o *GetTCPRequestRulesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get TCP request rules default response
func (o *GetTCPRequestRulesDefault) WithConfigurationVersion(configurationVersion string) *GetTCPRequestRulesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get TCP request rules default response
func (o *GetTCPRequestRulesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get TCP request rules default response
func (o *GetTCPRequestRulesDefault) WithPayload(payload *models.Error) *GetTCPRequestRulesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get TCP request rules default response
func (o *GetTCPRequestRulesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTCPRequestRulesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
