// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Operator
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package operator_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lgj101/operator/models"
)

// PutTenantYAMLCreatedCode is the HTTP code returned for type PutTenantYAMLCreated
const PutTenantYAMLCreatedCode int = 201

/*
PutTenantYAMLCreated A successful response.

swagger:response putTenantYAMLCreated
*/
type PutTenantYAMLCreated struct {
}

// NewPutTenantYAMLCreated creates PutTenantYAMLCreated with default headers values
func NewPutTenantYAMLCreated() *PutTenantYAMLCreated {

	return &PutTenantYAMLCreated{}
}

// WriteResponse to the client
func (o *PutTenantYAMLCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*
PutTenantYAMLDefault Generic error response.

swagger:response putTenantYAMLDefault
*/
type PutTenantYAMLDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutTenantYAMLDefault creates PutTenantYAMLDefault with default headers values
func NewPutTenantYAMLDefault(code int) *PutTenantYAMLDefault {
	if code <= 0 {
		code = 500
	}

	return &PutTenantYAMLDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put tenant y a m l default response
func (o *PutTenantYAMLDefault) WithStatusCode(code int) *PutTenantYAMLDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put tenant y a m l default response
func (o *PutTenantYAMLDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put tenant y a m l default response
func (o *PutTenantYAMLDefault) WithPayload(payload *models.Error) *PutTenantYAMLDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put tenant y a m l default response
func (o *PutTenantYAMLDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTenantYAMLDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
