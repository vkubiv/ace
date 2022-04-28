// Code generated by go-swagger; DO NOT EDIT.

// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0
//

package operations

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

	"github.com/trustbloc/ace/pkg/client/csh/models"
)

// NewPostExtractParams creates a new PostExtractParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExtractParams() *PostExtractParams {
	return &PostExtractParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExtractParamsWithTimeout creates a new PostExtractParams object
// with the ability to set a timeout on a request.
func NewPostExtractParamsWithTimeout(timeout time.Duration) *PostExtractParams {
	return &PostExtractParams{
		timeout: timeout,
	}
}

// NewPostExtractParamsWithContext creates a new PostExtractParams object
// with the ability to set a context for a request.
func NewPostExtractParamsWithContext(ctx context.Context) *PostExtractParams {
	return &PostExtractParams{
		Context: ctx,
	}
}

// NewPostExtractParamsWithHTTPClient creates a new PostExtractParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExtractParamsWithHTTPClient(client *http.Client) *PostExtractParams {
	return &PostExtractParams{
		HTTPClient: client,
	}
}

/* PostExtractParams contains all the parameters to send to the API endpoint
   for the post extract operation.

   Typically these are written to a http.Request.
*/
type PostExtractParams struct {

	// Request.
	Request []models.Query

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post extract params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExtractParams) WithDefaults() *PostExtractParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post extract params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExtractParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post extract params
func (o *PostExtractParams) WithTimeout(timeout time.Duration) *PostExtractParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post extract params
func (o *PostExtractParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post extract params
func (o *PostExtractParams) WithContext(ctx context.Context) *PostExtractParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post extract params
func (o *PostExtractParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post extract params
func (o *PostExtractParams) WithHTTPClient(client *http.Client) *PostExtractParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post extract params
func (o *PostExtractParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post extract params
func (o *PostExtractParams) WithRequest(request []models.Query) *PostExtractParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post extract params
func (o *PostExtractParams) SetRequest(request []models.Query) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostExtractParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
