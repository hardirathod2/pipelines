// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/go_http_client/run_model"
)

// ReadArtifactReader is a Reader for the ReadArtifact structure.
type ReadArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReadArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewReadArtifactDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadArtifactOK creates a ReadArtifactOK with default headers values
func NewReadArtifactOK() *ReadArtifactOK {
	return &ReadArtifactOK{}
}

/*ReadArtifactOK handles this case with default header values.

A successful response.
*/
type ReadArtifactOK struct {
	Payload *run_model.APIReadArtifactResponse
}

func (o *ReadArtifactOK) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/runs/{run_id}/nodes/{node_id}/artifacts/{artifact_name}:read][%d] readArtifactOK  %+v", 200, o.Payload)
}

func (o *ReadArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.APIReadArtifactResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadArtifactDefault creates a ReadArtifactDefault with default headers values
func NewReadArtifactDefault(code int) *ReadArtifactDefault {
	return &ReadArtifactDefault{
		_statusCode: code,
	}
}

/*ReadArtifactDefault handles this case with default header values.

ReadArtifactDefault read artifact default
*/
type ReadArtifactDefault struct {
	_statusCode int

	Payload *run_model.APIStatus
}

// Code gets the status code for the read artifact default response
func (o *ReadArtifactDefault) Code() int {
	return o._statusCode
}

func (o *ReadArtifactDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/runs/{run_id}/nodes/{node_id}/artifacts/{artifact_name}:read][%d] ReadArtifact default  %+v", o._statusCode, o.Payload)
}

func (o *ReadArtifactDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
