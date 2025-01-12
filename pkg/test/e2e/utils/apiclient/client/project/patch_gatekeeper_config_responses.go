// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// PatchGatekeeperConfigReader is a Reader for the PatchGatekeeperConfig structure.
type PatchGatekeeperConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGatekeeperConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchGatekeeperConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchGatekeeperConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchGatekeeperConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchGatekeeperConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchGatekeeperConfigOK creates a PatchGatekeeperConfigOK with default headers values
func NewPatchGatekeeperConfigOK() *PatchGatekeeperConfigOK {
	return &PatchGatekeeperConfigOK{}
}

/* PatchGatekeeperConfigOK describes a response with status code 200, with default header values.

GatekeeperConfig
*/
type PatchGatekeeperConfigOK struct {
	Payload *models.GatekeeperConfig
}

func (o *PatchGatekeeperConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfigOK  %+v", 200, o.Payload)
}
func (o *PatchGatekeeperConfigOK) GetPayload() *models.GatekeeperConfig {
	return o.Payload
}

func (o *PatchGatekeeperConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatekeeperConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGatekeeperConfigUnauthorized creates a PatchGatekeeperConfigUnauthorized with default headers values
func NewPatchGatekeeperConfigUnauthorized() *PatchGatekeeperConfigUnauthorized {
	return &PatchGatekeeperConfigUnauthorized{}
}

/* PatchGatekeeperConfigUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchGatekeeperConfigUnauthorized struct {
}

func (o *PatchGatekeeperConfigUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfigUnauthorized ", 401)
}

func (o *PatchGatekeeperConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGatekeeperConfigForbidden creates a PatchGatekeeperConfigForbidden with default headers values
func NewPatchGatekeeperConfigForbidden() *PatchGatekeeperConfigForbidden {
	return &PatchGatekeeperConfigForbidden{}
}

/* PatchGatekeeperConfigForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchGatekeeperConfigForbidden struct {
}

func (o *PatchGatekeeperConfigForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfigForbidden ", 403)
}

func (o *PatchGatekeeperConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGatekeeperConfigDefault creates a PatchGatekeeperConfigDefault with default headers values
func NewPatchGatekeeperConfigDefault(code int) *PatchGatekeeperConfigDefault {
	return &PatchGatekeeperConfigDefault{
		_statusCode: code,
	}
}

/* PatchGatekeeperConfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchGatekeeperConfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch gatekeeper config default response
func (o *PatchGatekeeperConfigDefault) Code() int {
	return o._statusCode
}

func (o *PatchGatekeeperConfigDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config][%d] patchGatekeeperConfig default  %+v", o._statusCode, o.Payload)
}
func (o *PatchGatekeeperConfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchGatekeeperConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
