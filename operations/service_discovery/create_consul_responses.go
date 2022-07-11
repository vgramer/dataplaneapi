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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateConsulCreatedCode is the HTTP code returned for type CreateConsulCreated
const CreateConsulCreatedCode int = 201

/*CreateConsulCreated Consul created

swagger:response createConsulCreated
*/
type CreateConsulCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Consul `json:"body,omitempty"`
}

// NewCreateConsulCreated creates CreateConsulCreated with default headers values
func NewCreateConsulCreated() *CreateConsulCreated {

	return &CreateConsulCreated{}
}

// WithPayload adds the payload to the create consul created response
func (o *CreateConsulCreated) WithPayload(payload *models.Consul) *CreateConsulCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create consul created response
func (o *CreateConsulCreated) SetPayload(payload *models.Consul) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConsulCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateConsulBadRequestCode is the HTTP code returned for type CreateConsulBadRequest
const CreateConsulBadRequestCode int = 400

/*CreateConsulBadRequest Bad request

swagger:response createConsulBadRequest
*/
type CreateConsulBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateConsulBadRequest creates CreateConsulBadRequest with default headers values
func NewCreateConsulBadRequest() *CreateConsulBadRequest {

	return &CreateConsulBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create consul bad request response
func (o *CreateConsulBadRequest) WithConfigurationVersion(configurationVersion string) *CreateConsulBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create consul bad request response
func (o *CreateConsulBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create consul bad request response
func (o *CreateConsulBadRequest) WithPayload(payload *models.Error) *CreateConsulBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create consul bad request response
func (o *CreateConsulBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConsulBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateConsulConflictCode is the HTTP code returned for type CreateConsulConflict
const CreateConsulConflictCode int = 409

/*CreateConsulConflict The specified resource already exists

swagger:response createConsulConflict
*/
type CreateConsulConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateConsulConflict creates CreateConsulConflict with default headers values
func NewCreateConsulConflict() *CreateConsulConflict {

	return &CreateConsulConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create consul conflict response
func (o *CreateConsulConflict) WithConfigurationVersion(configurationVersion string) *CreateConsulConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create consul conflict response
func (o *CreateConsulConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create consul conflict response
func (o *CreateConsulConflict) WithPayload(payload *models.Error) *CreateConsulConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create consul conflict response
func (o *CreateConsulConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConsulConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateConsulDefault General Error

swagger:response createConsulDefault
*/
type CreateConsulDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateConsulDefault creates CreateConsulDefault with default headers values
func NewCreateConsulDefault(code int) *CreateConsulDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateConsulDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create consul default response
func (o *CreateConsulDefault) WithStatusCode(code int) *CreateConsulDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create consul default response
func (o *CreateConsulDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create consul default response
func (o *CreateConsulDefault) WithConfigurationVersion(configurationVersion string) *CreateConsulDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create consul default response
func (o *CreateConsulDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create consul default response
func (o *CreateConsulDefault) WithPayload(payload *models.Error) *CreateConsulDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create consul default response
func (o *CreateConsulDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConsulDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
