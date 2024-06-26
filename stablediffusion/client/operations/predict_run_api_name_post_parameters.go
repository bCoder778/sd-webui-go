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

// NewPredictRunAPINamePostParams creates a new PredictRunAPINamePostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPredictRunAPINamePostParams() *PredictRunAPINamePostParams {
	return &PredictRunAPINamePostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPredictRunAPINamePostParamsWithTimeout creates a new PredictRunAPINamePostParams object
// with the ability to set a timeout on a request.
func NewPredictRunAPINamePostParamsWithTimeout(timeout time.Duration) *PredictRunAPINamePostParams {
	return &PredictRunAPINamePostParams{
		timeout: timeout,
	}
}

// NewPredictRunAPINamePostParamsWithContext creates a new PredictRunAPINamePostParams object
// with the ability to set a context for a request.
func NewPredictRunAPINamePostParamsWithContext(ctx context.Context) *PredictRunAPINamePostParams {
	return &PredictRunAPINamePostParams{
		Context: ctx,
	}
}

// NewPredictRunAPINamePostParamsWithHTTPClient creates a new PredictRunAPINamePostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPredictRunAPINamePostParamsWithHTTPClient(client *http.Client) *PredictRunAPINamePostParams {
	return &PredictRunAPINamePostParams{
		HTTPClient: client,
	}
}

/*
PredictRunAPINamePostParams contains all the parameters to send to the API endpoint

	for the predict run api name post operation.

	Typically these are written to a http.Request.
*/
type PredictRunAPINamePostParams struct {

	// APIName.
	APIName string

	// Body.
	Body *models.PredictBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the predict run api name post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PredictRunAPINamePostParams) WithDefaults() *PredictRunAPINamePostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the predict run api name post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PredictRunAPINamePostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the predict run api name post params
func (o *PredictRunAPINamePostParams) WithTimeout(timeout time.Duration) *PredictRunAPINamePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the predict run api name post params
func (o *PredictRunAPINamePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the predict run api name post params
func (o *PredictRunAPINamePostParams) WithContext(ctx context.Context) *PredictRunAPINamePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the predict run api name post params
func (o *PredictRunAPINamePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the predict run api name post params
func (o *PredictRunAPINamePostParams) WithHTTPClient(client *http.Client) *PredictRunAPINamePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the predict run api name post params
func (o *PredictRunAPINamePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIName adds the aPIName to the predict run api name post params
func (o *PredictRunAPINamePostParams) WithAPIName(aPIName string) *PredictRunAPINamePostParams {
	o.SetAPIName(aPIName)
	return o
}

// SetAPIName adds the apiName to the predict run api name post params
func (o *PredictRunAPINamePostParams) SetAPIName(aPIName string) {
	o.APIName = aPIName
}

// WithBody adds the body to the predict run api name post params
func (o *PredictRunAPINamePostParams) WithBody(body *models.PredictBody) *PredictRunAPINamePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the predict run api name post params
func (o *PredictRunAPINamePostParams) SetBody(body *models.PredictBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PredictRunAPINamePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param api_name
	if err := r.SetPathParam("api_name", o.APIName); err != nil {
		return err
	}
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
