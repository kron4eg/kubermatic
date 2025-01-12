// Code generated by go-swagger; DO NOT EDIT.

package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGatesParams creates a new GatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGatesParams() *GatesParams {
	return &GatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGatesParamsWithTimeout creates a new GatesParams object
// with the ability to set a timeout on a request.
func NewGatesParamsWithTimeout(timeout time.Duration) *GatesParams {
	return &GatesParams{
		timeout: timeout,
	}
}

// NewGatesParamsWithContext creates a new GatesParams object
// with the ability to set a context for a request.
func NewGatesParamsWithContext(ctx context.Context) *GatesParams {
	return &GatesParams{
		Context: ctx,
	}
}

// NewGatesParamsWithHTTPClient creates a new GatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGatesParamsWithHTTPClient(client *http.Client) *GatesParams {
	return &GatesParams{
		HTTPClient: client,
	}
}

/* GatesParams contains all the parameters to send to the API endpoint
   for the gates operation.

   Typically these are written to a http.Request.
*/
type GatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GatesParams) WithDefaults() *GatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gates params
func (o *GatesParams) WithTimeout(timeout time.Duration) *GatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gates params
func (o *GatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gates params
func (o *GatesParams) WithContext(ctx context.Context) *GatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gates params
func (o *GatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gates params
func (o *GatesParams) WithHTTPClient(client *http.Client) *GatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gates params
func (o *GatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
