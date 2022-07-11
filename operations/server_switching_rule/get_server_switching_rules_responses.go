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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetServerSwitchingRulesOKCode is the HTTP code returned for type GetServerSwitchingRulesOK
const GetServerSwitchingRulesOKCode int = 200

/*GetServerSwitchingRulesOK Successful operation

swagger:response getServerSwitchingRulesOK
*/
type GetServerSwitchingRulesOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetServerSwitchingRulesOKBody `json:"body,omitempty"`
}

// NewGetServerSwitchingRulesOK creates GetServerSwitchingRulesOK with default headers values
func NewGetServerSwitchingRulesOK() *GetServerSwitchingRulesOK {

	return &GetServerSwitchingRulesOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get server switching rules o k response
func (o *GetServerSwitchingRulesOK) WithConfigurationVersion(configurationVersion string) *GetServerSwitchingRulesOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get server switching rules o k response
func (o *GetServerSwitchingRulesOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get server switching rules o k response
func (o *GetServerSwitchingRulesOK) WithPayload(payload *GetServerSwitchingRulesOKBody) *GetServerSwitchingRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server switching rules o k response
func (o *GetServerSwitchingRulesOK) SetPayload(payload *GetServerSwitchingRulesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerSwitchingRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetServerSwitchingRulesDefault General Error

swagger:response getServerSwitchingRulesDefault
*/
type GetServerSwitchingRulesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerSwitchingRulesDefault creates GetServerSwitchingRulesDefault with default headers values
func NewGetServerSwitchingRulesDefault(code int) *GetServerSwitchingRulesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServerSwitchingRulesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) WithStatusCode(code int) *GetServerSwitchingRulesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) WithConfigurationVersion(configurationVersion string) *GetServerSwitchingRulesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) WithPayload(payload *models.Error) *GetServerSwitchingRulesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerSwitchingRulesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
