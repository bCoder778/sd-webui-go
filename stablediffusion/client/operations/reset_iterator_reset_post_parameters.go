// Code generated by go-swagger; DO NOT EDIT.

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

	"github.com/bCoder778/sd-webui-go/stablediffusion/models"
)

// NewResetIteratorResetPostParams creates a new ResetIteratorResetPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetIteratorResetPostParams() *ResetIteratorResetPostParams {
	return &ResetIteratorResetPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetIteratorResetPostParamsWithTimeout creates a new ResetIteratorResetPostParams object
// with the ability to set a timeout on a request.
func NewResetIteratorResetPostParamsWithTimeout(timeout time.Duration) *ResetIteratorResetPostParams {
	return &ResetIteratorResetPostParams{
		timeout: timeout,
	}
}

// NewResetIteratorResetPostParamsWithContext creates a new ResetIteratorResetPostParams object
// with the ability to set a context for a request.
func NewResetIteratorResetPostParamsWithContext(ctx context.Context) *ResetIteratorResetPostParams {
	return &ResetIteratorResetPostParams{
		Context: ctx,
	}
}

// NewResetIteratorResetPostParamsWithHTTPClient creates a new ResetIteratorResetPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetIteratorResetPostParamsWithHTTPClient(client *http.Client) *ResetIteratorResetPostParams {
	return &ResetIteratorResetPostParams{
		HTTPClient: client,
	}
}

/*
ResetIteratorResetPostParams contains all the parameters to send to the API endpoint

	for the reset iterator reset post operation.

	Typically these are written to a http.Request.
*/
type ResetIteratorResetPostParams struct {

	// Body.
	Body *models.ResetBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset iterator reset post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetIteratorResetPostParams) WithDefaults() *ResetIteratorResetPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset iterator reset post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetIteratorResetPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset iterator reset post params
func (o *ResetIteratorResetPostParams) WithTimeout(timeout time.Duration) *ResetIteratorResetPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset iterator reset post params
func (o *ResetIteratorResetPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset iterator reset post params
func (o *ResetIteratorResetPostParams) WithContext(ctx context.Context) *ResetIteratorResetPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset iterator reset post params
func (o *ResetIteratorResetPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset iterator reset post params
func (o *ResetIteratorResetPostParams) WithHTTPClient(client *http.Client) *ResetIteratorResetPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset iterator reset post params
func (o *ResetIteratorResetPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reset iterator reset post params
func (o *ResetIteratorResetPostParams) WithBody(body *models.ResetBody) *ResetIteratorResetPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reset iterator reset post params
func (o *ResetIteratorResetPostParams) SetBody(body *models.ResetBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ResetIteratorResetPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
