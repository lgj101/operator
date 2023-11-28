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

// GetTenantLogReportOKCode is the HTTP code returned for type GetTenantLogReportOK
const GetTenantLogReportOKCode int = 200

/*
GetTenantLogReportOK A successful response.

swagger:response getTenantLogReportOK
*/
type GetTenantLogReportOK struct {

	/*
	  In: Body
	*/
	Payload *models.TenantLogReport `json:"body,omitempty"`
}

// NewGetTenantLogReportOK creates GetTenantLogReportOK with default headers values
func NewGetTenantLogReportOK() *GetTenantLogReportOK {

	return &GetTenantLogReportOK{}
}

// WithPayload adds the payload to the get tenant log report o k response
func (o *GetTenantLogReportOK) WithPayload(payload *models.TenantLogReport) *GetTenantLogReportOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tenant log report o k response
func (o *GetTenantLogReportOK) SetPayload(payload *models.TenantLogReport) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTenantLogReportOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetTenantLogReportDefault Generic error response.

swagger:response getTenantLogReportDefault
*/
type GetTenantLogReportDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTenantLogReportDefault creates GetTenantLogReportDefault with default headers values
func NewGetTenantLogReportDefault(code int) *GetTenantLogReportDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTenantLogReportDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get tenant log report default response
func (o *GetTenantLogReportDefault) WithStatusCode(code int) *GetTenantLogReportDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get tenant log report default response
func (o *GetTenantLogReportDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get tenant log report default response
func (o *GetTenantLogReportDefault) WithPayload(payload *models.Error) *GetTenantLogReportDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tenant log report default response
func (o *GetTenantLogReportDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTenantLogReportDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
