// Code generated by go-swagger; DO NOT EDIT.

package eks

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

// NewListEKSRegionsParams creates a new ListEKSRegionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListEKSRegionsParams() *ListEKSRegionsParams {
	return &ListEKSRegionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListEKSRegionsParamsWithTimeout creates a new ListEKSRegionsParams object
// with the ability to set a timeout on a request.
func NewListEKSRegionsParamsWithTimeout(timeout time.Duration) *ListEKSRegionsParams {
	return &ListEKSRegionsParams{
		timeout: timeout,
	}
}

// NewListEKSRegionsParamsWithContext creates a new ListEKSRegionsParams object
// with the ability to set a context for a request.
func NewListEKSRegionsParamsWithContext(ctx context.Context) *ListEKSRegionsParams {
	return &ListEKSRegionsParams{
		Context: ctx,
	}
}

// NewListEKSRegionsParamsWithHTTPClient creates a new ListEKSRegionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListEKSRegionsParamsWithHTTPClient(client *http.Client) *ListEKSRegionsParams {
	return &ListEKSRegionsParams{
		HTTPClient: client,
	}
}

/* ListEKSRegionsParams contains all the parameters to send to the API endpoint
   for the list e k s regions operation.

   Typically these are written to a http.Request.
*/
type ListEKSRegionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list e k s regions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEKSRegionsParams) WithDefaults() *ListEKSRegionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list e k s regions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEKSRegionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list e k s regions params
func (o *ListEKSRegionsParams) WithTimeout(timeout time.Duration) *ListEKSRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list e k s regions params
func (o *ListEKSRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list e k s regions params
func (o *ListEKSRegionsParams) WithContext(ctx context.Context) *ListEKSRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list e k s regions params
func (o *ListEKSRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list e k s regions params
func (o *ListEKSRegionsParams) WithHTTPClient(client *http.Client) *ListEKSRegionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list e k s regions params
func (o *ListEKSRegionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListEKSRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
