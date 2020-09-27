// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListVSphereFoldersNoCredentialsReader is a Reader for the ListVSphereFoldersNoCredentials structure.
type ListVSphereFoldersNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVSphereFoldersNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVSphereFoldersNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVSphereFoldersNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVSphereFoldersNoCredentialsOK creates a ListVSphereFoldersNoCredentialsOK with default headers values
func NewListVSphereFoldersNoCredentialsOK() *ListVSphereFoldersNoCredentialsOK {
	return &ListVSphereFoldersNoCredentialsOK{}
}

/*ListVSphereFoldersNoCredentialsOK handles this case with default header values.

VSphereFolder
*/
type ListVSphereFoldersNoCredentialsOK struct {
	Payload []*models.VSphereFolder
}

func (o *ListVSphereFoldersNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/folders][%d] listVSphereFoldersNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListVSphereFoldersNoCredentialsOK) GetPayload() []*models.VSphereFolder {
	return o.Payload
}

func (o *ListVSphereFoldersNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVSphereFoldersNoCredentialsDefault creates a ListVSphereFoldersNoCredentialsDefault with default headers values
func NewListVSphereFoldersNoCredentialsDefault(code int) *ListVSphereFoldersNoCredentialsDefault {
	return &ListVSphereFoldersNoCredentialsDefault{
		_statusCode: code,
	}
}

/*ListVSphereFoldersNoCredentialsDefault handles this case with default header values.

errorResponse
*/
type ListVSphereFoldersNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v sphere folders no credentials default response
func (o *ListVSphereFoldersNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListVSphereFoldersNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/folders][%d] listVSphereFoldersNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListVSphereFoldersNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVSphereFoldersNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}