// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bCoder778/sd-webui-go/stablediffusion/models"
)

// APIInfoInfoGetReader is a Reader for the APIInfoInfoGet structure.
type APIInfoInfoGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIInfoInfoGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIInfoInfoGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewAPIInfoInfoGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAPIInfoInfoGetOK creates a APIInfoInfoGetOK with default headers values
func NewAPIInfoInfoGetOK() *APIInfoInfoGetOK {
	return &APIInfoInfoGetOK{}
}

/*
APIInfoInfoGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type APIInfoInfoGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this api info info get o k response has a 2xx status code
func (o *APIInfoInfoGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this api info info get o k response has a 3xx status code
func (o *APIInfoInfoGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api info info get o k response has a 4xx status code
func (o *APIInfoInfoGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this api info info get o k response has a 5xx status code
func (o *APIInfoInfoGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this api info info get o k response a status code equal to that given
func (o *APIInfoInfoGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the api info info get o k response
func (o *APIInfoInfoGetOK) Code() int {
	return 200
}

func (o *APIInfoInfoGetOK) Error() string {
	return fmt.Sprintf("[GET /info][%d] apiInfoInfoGetOK  %+v", 200, o.Payload)
}

func (o *APIInfoInfoGetOK) String() string {
	return fmt.Sprintf("[GET /info][%d] apiInfoInfoGetOK  %+v", 200, o.Payload)
}

func (o *APIInfoInfoGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *APIInfoInfoGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIInfoInfoGetUnprocessableEntity creates a APIInfoInfoGetUnprocessableEntity with default headers values
func NewAPIInfoInfoGetUnprocessableEntity() *APIInfoInfoGetUnprocessableEntity {
	return &APIInfoInfoGetUnprocessableEntity{}
}

/*
APIInfoInfoGetUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type APIInfoInfoGetUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this api info info get unprocessable entity response has a 2xx status code
func (o *APIInfoInfoGetUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api info info get unprocessable entity response has a 3xx status code
func (o *APIInfoInfoGetUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api info info get unprocessable entity response has a 4xx status code
func (o *APIInfoInfoGetUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this api info info get unprocessable entity response has a 5xx status code
func (o *APIInfoInfoGetUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this api info info get unprocessable entity response a status code equal to that given
func (o *APIInfoInfoGetUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the api info info get unprocessable entity response
func (o *APIInfoInfoGetUnprocessableEntity) Code() int {
	return 422
}

func (o *APIInfoInfoGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /info][%d] apiInfoInfoGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *APIInfoInfoGetUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /info][%d] apiInfoInfoGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *APIInfoInfoGetUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *APIInfoInfoGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
