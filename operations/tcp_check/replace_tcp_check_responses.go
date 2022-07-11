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

package tcp_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// ReplaceTCPCheckOKCode is the HTTP code returned for type ReplaceTCPCheckOK
const ReplaceTCPCheckOKCode int = 200

/*ReplaceTCPCheckOK TCP check replaced

swagger:response replaceTcpCheckOK
*/
type ReplaceTCPCheckOK struct {

	/*
	  In: Body
	*/
	Payload *models.TCPCheck `json:"body,omitempty"`
}

// NewReplaceTCPCheckOK creates ReplaceTCPCheckOK with default headers values
func NewReplaceTCPCheckOK() *ReplaceTCPCheckOK {

	return &ReplaceTCPCheckOK{}
}

// WithPayload adds the payload to the replace Tcp check o k response
func (o *ReplaceTCPCheckOK) WithPayload(payload *models.TCPCheck) *ReplaceTCPCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp check o k response
func (o *ReplaceTCPCheckOK) SetPayload(payload *models.TCPCheck) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceTCPCheckAcceptedCode is the HTTP code returned for type ReplaceTCPCheckAccepted
const ReplaceTCPCheckAcceptedCode int = 202

/*ReplaceTCPCheckAccepted Configuration change accepted and reload requested

swagger:response replaceTcpCheckAccepted
*/
type ReplaceTCPCheckAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.TCPCheck `json:"body,omitempty"`
}

// NewReplaceTCPCheckAccepted creates ReplaceTCPCheckAccepted with default headers values
func NewReplaceTCPCheckAccepted() *ReplaceTCPCheckAccepted {

	return &ReplaceTCPCheckAccepted{}
}

// WithReloadID adds the reloadId to the replace Tcp check accepted response
func (o *ReplaceTCPCheckAccepted) WithReloadID(reloadID string) *ReplaceTCPCheckAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace Tcp check accepted response
func (o *ReplaceTCPCheckAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace Tcp check accepted response
func (o *ReplaceTCPCheckAccepted) WithPayload(payload *models.TCPCheck) *ReplaceTCPCheckAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp check accepted response
func (o *ReplaceTCPCheckAccepted) SetPayload(payload *models.TCPCheck) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPCheckAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceTCPCheckBadRequestCode is the HTTP code returned for type ReplaceTCPCheckBadRequest
const ReplaceTCPCheckBadRequestCode int = 400

/*ReplaceTCPCheckBadRequest Bad request

swagger:response replaceTcpCheckBadRequest
*/
type ReplaceTCPCheckBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPCheckBadRequest creates ReplaceTCPCheckBadRequest with default headers values
func NewReplaceTCPCheckBadRequest() *ReplaceTCPCheckBadRequest {

	return &ReplaceTCPCheckBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace Tcp check bad request response
func (o *ReplaceTCPCheckBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceTCPCheckBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Tcp check bad request response
func (o *ReplaceTCPCheckBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Tcp check bad request response
func (o *ReplaceTCPCheckBadRequest) WithPayload(payload *models.Error) *ReplaceTCPCheckBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp check bad request response
func (o *ReplaceTCPCheckBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPCheckBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceTCPCheckNotFoundCode is the HTTP code returned for type ReplaceTCPCheckNotFound
const ReplaceTCPCheckNotFoundCode int = 404

/*ReplaceTCPCheckNotFound The specified resource was not found

swagger:response replaceTcpCheckNotFound
*/
type ReplaceTCPCheckNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPCheckNotFound creates ReplaceTCPCheckNotFound with default headers values
func NewReplaceTCPCheckNotFound() *ReplaceTCPCheckNotFound {

	return &ReplaceTCPCheckNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace Tcp check not found response
func (o *ReplaceTCPCheckNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceTCPCheckNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace Tcp check not found response
func (o *ReplaceTCPCheckNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace Tcp check not found response
func (o *ReplaceTCPCheckNotFound) WithPayload(payload *models.Error) *ReplaceTCPCheckNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp check not found response
func (o *ReplaceTCPCheckNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPCheckNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceTCPCheckDefault General Error

swagger:response replaceTcpCheckDefault
*/
type ReplaceTCPCheckDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPCheckDefault creates ReplaceTCPCheckDefault with default headers values
func NewReplaceTCPCheckDefault(code int) *ReplaceTCPCheckDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceTCPCheckDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace TCP check default response
func (o *ReplaceTCPCheckDefault) WithStatusCode(code int) *ReplaceTCPCheckDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace TCP check default response
func (o *ReplaceTCPCheckDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace TCP check default response
func (o *ReplaceTCPCheckDefault) WithConfigurationVersion(configurationVersion string) *ReplaceTCPCheckDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace TCP check default response
func (o *ReplaceTCPCheckDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace TCP check default response
func (o *ReplaceTCPCheckDefault) WithPayload(payload *models.Error) *ReplaceTCPCheckDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace TCP check default response
func (o *ReplaceTCPCheckDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPCheckDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
