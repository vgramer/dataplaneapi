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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetLogForwardsOKCode is the HTTP code returned for type GetLogForwardsOK
const GetLogForwardsOKCode int = 200

/*GetLogForwardsOK Successful operation

swagger:response getLogForwardsOK
*/
type GetLogForwardsOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetLogForwardsOKBody `json:"body,omitempty"`
}

// NewGetLogForwardsOK creates GetLogForwardsOK with default headers values
func NewGetLogForwardsOK() *GetLogForwardsOK {

	return &GetLogForwardsOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get log forwards o k response
func (o *GetLogForwardsOK) WithConfigurationVersion(configurationVersion string) *GetLogForwardsOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get log forwards o k response
func (o *GetLogForwardsOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get log forwards o k response
func (o *GetLogForwardsOK) WithPayload(payload *GetLogForwardsOKBody) *GetLogForwardsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get log forwards o k response
func (o *GetLogForwardsOK) SetPayload(payload *GetLogForwardsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogForwardsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetLogForwardsDefault General Error

swagger:response getLogForwardsDefault
*/
type GetLogForwardsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLogForwardsDefault creates GetLogForwardsDefault with default headers values
func NewGetLogForwardsDefault(code int) *GetLogForwardsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLogForwardsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get log forwards default response
func (o *GetLogForwardsDefault) WithStatusCode(code int) *GetLogForwardsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get log forwards default response
func (o *GetLogForwardsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get log forwards default response
func (o *GetLogForwardsDefault) WithConfigurationVersion(configurationVersion string) *GetLogForwardsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get log forwards default response
func (o *GetLogForwardsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get log forwards default response
func (o *GetLogForwardsDefault) WithPayload(payload *models.Error) *GetLogForwardsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get log forwards default response
func (o *GetLogForwardsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogForwardsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
