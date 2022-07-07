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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// CreateStorageMapFileMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var CreateStorageMapFileMaxParseMemory int64 = 32 << 20

// NewCreateStorageMapFileParams creates a new CreateStorageMapFileParams object
//
// There are no default values defined in the spec.
func NewCreateStorageMapFileParams() CreateStorageMapFileParams {

	return CreateStorageMapFileParams{}
}

// CreateStorageMapFileParams contains all the bound params for the create storage map file operation
// typically these are obtained from a http.Request
//
// swagger:parameters createStorageMapFile
type CreateStorageMapFileParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The map file contents
	  In: formData
	*/
	FileUpload io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateStorageMapFileParams() beforehand.
func (o *CreateStorageMapFileParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(CreateStorageMapFileMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	fileUpload, fileUploadHeader, err := r.FormFile("file_upload")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "fileUpload", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindFileUpload(fileUpload, fileUploadHeader); err != nil {
		res = append(res, err)
	} else {
		o.FileUpload = &runtime.File{Data: fileUpload, Header: fileUploadHeader}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFileUpload binds file parameter FileUpload.
//
// The only supported validations on files are MinLength and MaxLength
func (o *CreateStorageMapFileParams) bindFileUpload(file multipart.File, header *multipart.FileHeader) error {
	return nil
}
