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

// RembgRemoveRembgPostReader is a Reader for the RembgRemoveRembgPost structure.
type RembgRemoveRembgPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RembgRemoveRembgPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRembgRemoveRembgPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewRembgRemoveRembgPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRembgRemoveRembgPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRembgRemoveRembgPostOK creates a RembgRemoveRembgPostOK with default headers values
func NewRembgRemoveRembgPostOK() *RembgRemoveRembgPostOK {
	return &RembgRemoveRembgPostOK{}
}

/*
RembgRemoveRembgPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type RembgRemoveRembgPostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this rembg remove rembg post o k response has a 2xx status code
func (o *RembgRemoveRembgPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rembg remove rembg post o k response has a 3xx status code
func (o *RembgRemoveRembgPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rembg remove rembg post o k response has a 4xx status code
func (o *RembgRemoveRembgPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rembg remove rembg post o k response has a 5xx status code
func (o *RembgRemoveRembgPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rembg remove rembg post o k response a status code equal to that given
func (o *RembgRemoveRembgPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rembg remove rembg post o k response
func (o *RembgRemoveRembgPostOK) Code() int {
	return 200
}

func (o *RembgRemoveRembgPostOK) Error() string {
	return fmt.Sprintf("[POST /rembg][%d] rembgRemoveRembgPostOK  %+v", 200, o.Payload)
}

func (o *RembgRemoveRembgPostOK) String() string {
	return fmt.Sprintf("[POST /rembg][%d] rembgRemoveRembgPostOK  %+v", 200, o.Payload)
}

func (o *RembgRemoveRembgPostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RembgRemoveRembgPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRembgRemoveRembgPostUnprocessableEntity creates a RembgRemoveRembgPostUnprocessableEntity with default headers values
func NewRembgRemoveRembgPostUnprocessableEntity() *RembgRemoveRembgPostUnprocessableEntity {
	return &RembgRemoveRembgPostUnprocessableEntity{}
}

/*
RembgRemoveRembgPostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type RembgRemoveRembgPostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this rembg remove rembg post unprocessable entity response has a 2xx status code
func (o *RembgRemoveRembgPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rembg remove rembg post unprocessable entity response has a 3xx status code
func (o *RembgRemoveRembgPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rembg remove rembg post unprocessable entity response has a 4xx status code
func (o *RembgRemoveRembgPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this rembg remove rembg post unprocessable entity response has a 5xx status code
func (o *RembgRemoveRembgPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this rembg remove rembg post unprocessable entity response a status code equal to that given
func (o *RembgRemoveRembgPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the rembg remove rembg post unprocessable entity response
func (o *RembgRemoveRembgPostUnprocessableEntity) Code() int {
	return 422
}

func (o *RembgRemoveRembgPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /rembg][%d] rembgRemoveRembgPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RembgRemoveRembgPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /rembg][%d] rembgRemoveRembgPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RembgRemoveRembgPostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *RembgRemoveRembgPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRembgRemoveRembgPostInternalServerError creates a RembgRemoveRembgPostInternalServerError with default headers values
func NewRembgRemoveRembgPostInternalServerError() *RembgRemoveRembgPostInternalServerError {
	return &RembgRemoveRembgPostInternalServerError{}
}

/*
RembgRemoveRembgPostInternalServerError describes a response with status code 500, with default header values.

HTTPException
*/
type RembgRemoveRembgPostInternalServerError struct {
	Payload *models.HTTPException
}

// IsSuccess returns true when this rembg remove rembg post internal server error response has a 2xx status code
func (o *RembgRemoveRembgPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rembg remove rembg post internal server error response has a 3xx status code
func (o *RembgRemoveRembgPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rembg remove rembg post internal server error response has a 4xx status code
func (o *RembgRemoveRembgPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this rembg remove rembg post internal server error response has a 5xx status code
func (o *RembgRemoveRembgPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this rembg remove rembg post internal server error response a status code equal to that given
func (o *RembgRemoveRembgPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the rembg remove rembg post internal server error response
func (o *RembgRemoveRembgPostInternalServerError) Code() int {
	return 500
}

func (o *RembgRemoveRembgPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /rembg][%d] rembgRemoveRembgPostInternalServerError  %+v", 500, o.Payload)
}

func (o *RembgRemoveRembgPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /rembg][%d] rembgRemoveRembgPostInternalServerError  %+v", 500, o.Payload)
}

func (o *RembgRemoveRembgPostInternalServerError) GetPayload() *models.HTTPException {
	return o.Payload
}

func (o *RembgRemoveRembgPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
