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

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContentCode is the HTTP code returned for type DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent
const DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContentCode int = 204

/*DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent Successful operation

swagger:response deleteServicesHaproxyRuntimeAclFileEntriesIdNoContent
*/
type DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent struct {
}

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent creates DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent with default headers values
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent() *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent {

	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequestCode is the HTTP code returned for type DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest
const DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequestCode int = 400

/*DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest Bad request

swagger:response deleteServicesHaproxyRuntimeAclFileEntriesIdBadRequest
*/
type DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest creates DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest with default headers values
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest() *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest {

	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the delete services haproxy runtime Acl file entries Id bad request response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) WithConfigurationVersion(configurationVersion string) *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete services haproxy runtime Acl file entries Id bad request response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete services haproxy runtime Acl file entries Id bad request response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) WithPayload(payload *models.Error) *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete services haproxy runtime Acl file entries Id bad request response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFoundCode is the HTTP code returned for type DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound
const DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFoundCode int = 404

/*DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound The specified resource was not found

swagger:response deleteServicesHaproxyRuntimeAclFileEntriesIdNotFound
*/
type DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound creates DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound with default headers values
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound() *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound {

	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete services haproxy runtime Acl file entries Id not found response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) WithConfigurationVersion(configurationVersion string) *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete services haproxy runtime Acl file entries Id not found response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete services haproxy runtime Acl file entries Id not found response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) WithPayload(payload *models.Error) *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete services haproxy runtime Acl file entries Id not found response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault General Error

swagger:response deleteServicesHaproxyRuntimeAclFileEntriesIdDefault
*/
type DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDDefault creates DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault with default headers values
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDDefault(code int) *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete services haproxy runtime ACL file entries ID default response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) WithStatusCode(code int) *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete services haproxy runtime ACL file entries ID default response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete services haproxy runtime ACL file entries ID default response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) WithConfigurationVersion(configurationVersion string) *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete services haproxy runtime ACL file entries ID default response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete services haproxy runtime ACL file entries ID default response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) WithPayload(payload *models.Error) *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete services haproxy runtime ACL file entries ID default response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
