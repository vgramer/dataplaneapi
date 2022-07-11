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

// DeleteAWSRegionNoContentCode is the HTTP code returned for type DeleteAWSRegionNoContent
const DeleteAWSRegionNoContentCode int = 204

/*DeleteAWSRegionNoContent Resource deleted

swagger:response deleteAWSRegionNoContent
*/
type DeleteAWSRegionNoContent struct {
}

// NewDeleteAWSRegionNoContent creates DeleteAWSRegionNoContent with default headers values
func NewDeleteAWSRegionNoContent() *DeleteAWSRegionNoContent {

	return &DeleteAWSRegionNoContent{}
}

// WriteResponse to the client
func (o *DeleteAWSRegionNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteAWSRegionNotFoundCode is the HTTP code returned for type DeleteAWSRegionNotFound
const DeleteAWSRegionNotFoundCode int = 404

/*DeleteAWSRegionNotFound The specified resource was not found

swagger:response deleteAWSRegionNotFound
*/
type DeleteAWSRegionNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAWSRegionNotFound creates DeleteAWSRegionNotFound with default headers values
func NewDeleteAWSRegionNotFound() *DeleteAWSRegionNotFound {

	return &DeleteAWSRegionNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete a w s region not found response
func (o *DeleteAWSRegionNotFound) WithConfigurationVersion(configurationVersion string) *DeleteAWSRegionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete a w s region not found response
func (o *DeleteAWSRegionNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete a w s region not found response
func (o *DeleteAWSRegionNotFound) WithPayload(payload *models.Error) *DeleteAWSRegionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete a w s region not found response
func (o *DeleteAWSRegionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAWSRegionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteAWSRegionDefault General Error

swagger:response deleteAWSRegionDefault
*/
type DeleteAWSRegionDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAWSRegionDefault creates DeleteAWSRegionDefault with default headers values
func NewDeleteAWSRegionDefault(code int) *DeleteAWSRegionDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteAWSRegionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete a w s region default response
func (o *DeleteAWSRegionDefault) WithStatusCode(code int) *DeleteAWSRegionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete a w s region default response
func (o *DeleteAWSRegionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete a w s region default response
func (o *DeleteAWSRegionDefault) WithConfigurationVersion(configurationVersion string) *DeleteAWSRegionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete a w s region default response
func (o *DeleteAWSRegionDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete a w s region default response
func (o *DeleteAWSRegionDefault) WithPayload(payload *models.Error) *DeleteAWSRegionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete a w s region default response
func (o *DeleteAWSRegionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAWSRegionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
