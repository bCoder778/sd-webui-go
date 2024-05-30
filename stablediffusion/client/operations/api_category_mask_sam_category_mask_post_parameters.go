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

// NewAPICategoryMaskSamCategoryMaskPostParams creates a new APICategoryMaskSamCategoryMaskPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPICategoryMaskSamCategoryMaskPostParams() *APICategoryMaskSamCategoryMaskPostParams {
	return &APICategoryMaskSamCategoryMaskPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPICategoryMaskSamCategoryMaskPostParamsWithTimeout creates a new APICategoryMaskSamCategoryMaskPostParams object
// with the ability to set a timeout on a request.
func NewAPICategoryMaskSamCategoryMaskPostParamsWithTimeout(timeout time.Duration) *APICategoryMaskSamCategoryMaskPostParams {
	return &APICategoryMaskSamCategoryMaskPostParams{
		timeout: timeout,
	}
}

// NewAPICategoryMaskSamCategoryMaskPostParamsWithContext creates a new APICategoryMaskSamCategoryMaskPostParams object
// with the ability to set a context for a request.
func NewAPICategoryMaskSamCategoryMaskPostParamsWithContext(ctx context.Context) *APICategoryMaskSamCategoryMaskPostParams {
	return &APICategoryMaskSamCategoryMaskPostParams{
		Context: ctx,
	}
}

// NewAPICategoryMaskSamCategoryMaskPostParamsWithHTTPClient creates a new APICategoryMaskSamCategoryMaskPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPICategoryMaskSamCategoryMaskPostParamsWithHTTPClient(client *http.Client) *APICategoryMaskSamCategoryMaskPostParams {
	return &APICategoryMaskSamCategoryMaskPostParams{
		HTTPClient: client,
	}
}

/*
APICategoryMaskSamCategoryMaskPostParams contains all the parameters to send to the API endpoint

	for the api category mask sam category mask post operation.

	Typically these are written to a http.Request.
*/
type APICategoryMaskSamCategoryMaskPostParams struct {

	// Body.
	Body *models.BodyAPICategoryMaskSamCategoryMaskPost

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the api category mask sam category mask post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APICategoryMaskSamCategoryMaskPostParams) WithDefaults() *APICategoryMaskSamCategoryMaskPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the api category mask sam category mask post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APICategoryMaskSamCategoryMaskPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the api category mask sam category mask post params
func (o *APICategoryMaskSamCategoryMaskPostParams) WithTimeout(timeout time.Duration) *APICategoryMaskSamCategoryMaskPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api category mask sam category mask post params
func (o *APICategoryMaskSamCategoryMaskPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api category mask sam category mask post params
func (o *APICategoryMaskSamCategoryMaskPostParams) WithContext(ctx context.Context) *APICategoryMaskSamCategoryMaskPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api category mask sam category mask post params
func (o *APICategoryMaskSamCategoryMaskPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api category mask sam category mask post params
func (o *APICategoryMaskSamCategoryMaskPostParams) WithHTTPClient(client *http.Client) *APICategoryMaskSamCategoryMaskPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api category mask sam category mask post params
func (o *APICategoryMaskSamCategoryMaskPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the api category mask sam category mask post params
func (o *APICategoryMaskSamCategoryMaskPostParams) WithBody(body *models.BodyAPICategoryMaskSamCategoryMaskPost) *APICategoryMaskSamCategoryMaskPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the api category mask sam category mask post params
func (o *APICategoryMaskSamCategoryMaskPostParams) SetBody(body *models.BodyAPICategoryMaskSamCategoryMaskPost) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *APICategoryMaskSamCategoryMaskPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
