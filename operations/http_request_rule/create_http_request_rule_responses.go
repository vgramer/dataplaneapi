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

package http_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateHTTPRequestRuleCreatedCode is the HTTP code returned for type CreateHTTPRequestRuleCreated
const CreateHTTPRequestRuleCreatedCode int = 201

/*CreateHTTPRequestRuleCreated HTTP Request Rule created

swagger:response createHttpRequestRuleCreated
*/
type CreateHTTPRequestRuleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.HTTPRequestRule `json:"body,omitempty"`
}

// NewCreateHTTPRequestRuleCreated creates CreateHTTPRequestRuleCreated with default headers values
func NewCreateHTTPRequestRuleCreated() *CreateHTTPRequestRuleCreated {

	return &CreateHTTPRequestRuleCreated{}
}

// WithPayload adds the payload to the create Http request rule created response
func (o *CreateHTTPRequestRuleCreated) WithPayload(payload *models.HTTPRequestRule) *CreateHTTPRequestRuleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Http request rule created response
func (o *CreateHTTPRequestRuleCreated) SetPayload(payload *models.HTTPRequestRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHTTPRequestRuleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateHTTPRequestRuleAcceptedCode is the HTTP code returned for type CreateHTTPRequestRuleAccepted
const CreateHTTPRequestRuleAcceptedCode int = 202

/*CreateHTTPRequestRuleAccepted Configuration change accepted and reload requested

swagger:response createHttpRequestRuleAccepted
*/
type CreateHTTPRequestRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.HTTPRequestRule `json:"body,omitempty"`
}

// NewCreateHTTPRequestRuleAccepted creates CreateHTTPRequestRuleAccepted with default headers values
func NewCreateHTTPRequestRuleAccepted() *CreateHTTPRequestRuleAccepted {

	return &CreateHTTPRequestRuleAccepted{}
}

// WithReloadID adds the reloadId to the create Http request rule accepted response
func (o *CreateHTTPRequestRuleAccepted) WithReloadID(reloadID string) *CreateHTTPRequestRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create Http request rule accepted response
func (o *CreateHTTPRequestRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create Http request rule accepted response
func (o *CreateHTTPRequestRuleAccepted) WithPayload(payload *models.HTTPRequestRule) *CreateHTTPRequestRuleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Http request rule accepted response
func (o *CreateHTTPRequestRuleAccepted) SetPayload(payload *models.HTTPRequestRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHTTPRequestRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateHTTPRequestRuleBadRequestCode is the HTTP code returned for type CreateHTTPRequestRuleBadRequest
const CreateHTTPRequestRuleBadRequestCode int = 400

/*CreateHTTPRequestRuleBadRequest Bad request

swagger:response createHttpRequestRuleBadRequest
*/
type CreateHTTPRequestRuleBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateHTTPRequestRuleBadRequest creates CreateHTTPRequestRuleBadRequest with default headers values
func NewCreateHTTPRequestRuleBadRequest() *CreateHTTPRequestRuleBadRequest {

	return &CreateHTTPRequestRuleBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the create Http request rule bad request response
func (o *CreateHTTPRequestRuleBadRequest) WithConfigurationVersion(configurationVersion string) *CreateHTTPRequestRuleBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create Http request rule bad request response
func (o *CreateHTTPRequestRuleBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create Http request rule bad request response
func (o *CreateHTTPRequestRuleBadRequest) WithPayload(payload *models.Error) *CreateHTTPRequestRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Http request rule bad request response
func (o *CreateHTTPRequestRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHTTPRequestRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateHTTPRequestRuleConflictCode is the HTTP code returned for type CreateHTTPRequestRuleConflict
const CreateHTTPRequestRuleConflictCode int = 409

/*CreateHTTPRequestRuleConflict The specified resource already exists

swagger:response createHttpRequestRuleConflict
*/
type CreateHTTPRequestRuleConflict struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateHTTPRequestRuleConflict creates CreateHTTPRequestRuleConflict with default headers values
func NewCreateHTTPRequestRuleConflict() *CreateHTTPRequestRuleConflict {

	return &CreateHTTPRequestRuleConflict{}
}

// WithConfigurationVersion adds the configurationVersion to the create Http request rule conflict response
func (o *CreateHTTPRequestRuleConflict) WithConfigurationVersion(configurationVersion string) *CreateHTTPRequestRuleConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create Http request rule conflict response
func (o *CreateHTTPRequestRuleConflict) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create Http request rule conflict response
func (o *CreateHTTPRequestRuleConflict) WithPayload(payload *models.Error) *CreateHTTPRequestRuleConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Http request rule conflict response
func (o *CreateHTTPRequestRuleConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHTTPRequestRuleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateHTTPRequestRuleDefault General Error

swagger:response createHttpRequestRuleDefault
*/
type CreateHTTPRequestRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateHTTPRequestRuleDefault creates CreateHTTPRequestRuleDefault with default headers values
func NewCreateHTTPRequestRuleDefault(code int) *CreateHTTPRequestRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateHTTPRequestRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create HTTP request rule default response
func (o *CreateHTTPRequestRuleDefault) WithStatusCode(code int) *CreateHTTPRequestRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create HTTP request rule default response
func (o *CreateHTTPRequestRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create HTTP request rule default response
func (o *CreateHTTPRequestRuleDefault) WithConfigurationVersion(configurationVersion string) *CreateHTTPRequestRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create HTTP request rule default response
func (o *CreateHTTPRequestRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create HTTP request rule default response
func (o *CreateHTTPRequestRuleDefault) WithPayload(payload *models.Error) *CreateHTTPRequestRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create HTTP request rule default response
func (o *CreateHTTPRequestRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHTTPRequestRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
