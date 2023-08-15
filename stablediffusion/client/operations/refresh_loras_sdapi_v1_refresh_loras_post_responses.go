// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

// RefreshLorasSdapiV1RefreshLorasPostReader is a Reader for the RefreshLorasSdapiV1RefreshLorasPost structure.
type RefreshLorasSdapiV1RefreshLorasPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshLorasSdapiV1RefreshLorasPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshLorasSdapiV1RefreshLorasPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewRefreshLorasSdapiV1RefreshLorasPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRefreshLorasSdapiV1RefreshLorasPostOK creates a RefreshLorasSdapiV1RefreshLorasPostOK with default headers values
func NewRefreshLorasSdapiV1RefreshLorasPostOK() *RefreshLorasSdapiV1RefreshLorasPostOK {
	return &RefreshLorasSdapiV1RefreshLorasPostOK{}
}

/*
RefreshLorasSdapiV1RefreshLorasPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type RefreshLorasSdapiV1RefreshLorasPostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this refresh loras sdapi v1 refresh loras post o k response has a 2xx status code
func (o *RefreshLorasSdapiV1RefreshLorasPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this refresh loras sdapi v1 refresh loras post o k response has a 3xx status code
func (o *RefreshLorasSdapiV1RefreshLorasPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh loras sdapi v1 refresh loras post o k response has a 4xx status code
func (o *RefreshLorasSdapiV1RefreshLorasPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh loras sdapi v1 refresh loras post o k response has a 5xx status code
func (o *RefreshLorasSdapiV1RefreshLorasPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh loras sdapi v1 refresh loras post o k response a status code equal to that given
func (o *RefreshLorasSdapiV1RefreshLorasPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the refresh loras sdapi v1 refresh loras post o k response
func (o *RefreshLorasSdapiV1RefreshLorasPostOK) Code() int {
	return 200
}

func (o *RefreshLorasSdapiV1RefreshLorasPostOK) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/refresh-loras][%d] refreshLorasSdapiV1RefreshLorasPostOK  %+v", 200, o.Payload)
}

func (o *RefreshLorasSdapiV1RefreshLorasPostOK) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/refresh-loras][%d] refreshLorasSdapiV1RefreshLorasPostOK  %+v", 200, o.Payload)
}

func (o *RefreshLorasSdapiV1RefreshLorasPostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RefreshLorasSdapiV1RefreshLorasPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshLorasSdapiV1RefreshLorasPostInternalServerError creates a RefreshLorasSdapiV1RefreshLorasPostInternalServerError with default headers values
func NewRefreshLorasSdapiV1RefreshLorasPostInternalServerError() *RefreshLorasSdapiV1RefreshLorasPostInternalServerError {
	return &RefreshLorasSdapiV1RefreshLorasPostInternalServerError{}
}

/*
RefreshLorasSdapiV1RefreshLorasPostInternalServerError describes a response with status code 500, with default header values.

HTTPException
*/
type RefreshLorasSdapiV1RefreshLorasPostInternalServerError struct {
	Payload *models.HTTPException
}

// IsSuccess returns true when this refresh loras sdapi v1 refresh loras post internal server error response has a 2xx status code
func (o *RefreshLorasSdapiV1RefreshLorasPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh loras sdapi v1 refresh loras post internal server error response has a 3xx status code
func (o *RefreshLorasSdapiV1RefreshLorasPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh loras sdapi v1 refresh loras post internal server error response has a 4xx status code
func (o *RefreshLorasSdapiV1RefreshLorasPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh loras sdapi v1 refresh loras post internal server error response has a 5xx status code
func (o *RefreshLorasSdapiV1RefreshLorasPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this refresh loras sdapi v1 refresh loras post internal server error response a status code equal to that given
func (o *RefreshLorasSdapiV1RefreshLorasPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the refresh loras sdapi v1 refresh loras post internal server error response
func (o *RefreshLorasSdapiV1RefreshLorasPostInternalServerError) Code() int {
	return 500
}

func (o *RefreshLorasSdapiV1RefreshLorasPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/refresh-loras][%d] refreshLorasSdapiV1RefreshLorasPostInternalServerError  %+v", 500, o.Payload)
}

func (o *RefreshLorasSdapiV1RefreshLorasPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/refresh-loras][%d] refreshLorasSdapiV1RefreshLorasPostInternalServerError  %+v", 500, o.Payload)
}

func (o *RefreshLorasSdapiV1RefreshLorasPostInternalServerError) GetPayload() *models.HTTPException {
	return o.Payload
}

func (o *RefreshLorasSdapiV1RefreshLorasPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
