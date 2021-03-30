// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/api/utils/apiclient/models"
)

// GetUsersForProjectReader is a Reader for the GetUsersForProject structure.
type GetUsersForProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersForProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersForProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsersForProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersForProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetUsersForProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUsersForProjectOK creates a GetUsersForProjectOK with default headers values
func NewGetUsersForProjectOK() *GetUsersForProjectOK {
	return &GetUsersForProjectOK{}
}

/*GetUsersForProjectOK handles this case with default header values.

User
*/
type GetUsersForProjectOK struct {
	Payload []*models.User
}

func (o *GetUsersForProjectOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProjectOK  %+v", 200, o.Payload)
}

func (o *GetUsersForProjectOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *GetUsersForProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersForProjectUnauthorized creates a GetUsersForProjectUnauthorized with default headers values
func NewGetUsersForProjectUnauthorized() *GetUsersForProjectUnauthorized {
	return &GetUsersForProjectUnauthorized{}
}

/*GetUsersForProjectUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetUsersForProjectUnauthorized struct {
}

func (o *GetUsersForProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProjectUnauthorized ", 401)
}

func (o *GetUsersForProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersForProjectForbidden creates a GetUsersForProjectForbidden with default headers values
func NewGetUsersForProjectForbidden() *GetUsersForProjectForbidden {
	return &GetUsersForProjectForbidden{}
}

/*GetUsersForProjectForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetUsersForProjectForbidden struct {
}

func (o *GetUsersForProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProjectForbidden ", 403)
}

func (o *GetUsersForProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersForProjectDefault creates a GetUsersForProjectDefault with default headers values
func NewGetUsersForProjectDefault(code int) *GetUsersForProjectDefault {
	return &GetUsersForProjectDefault{
		_statusCode: code,
	}
}

/*GetUsersForProjectDefault handles this case with default header values.

errorResponse
*/
type GetUsersForProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get users for project default response
func (o *GetUsersForProjectDefault) Code() int {
	return o._statusCode
}

func (o *GetUsersForProjectDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/users][%d] getUsersForProject default  %+v", o._statusCode, o.Payload)
}

func (o *GetUsersForProjectDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUsersForProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}