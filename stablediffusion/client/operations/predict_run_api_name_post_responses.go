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

// PredictRunAPINamePostReader is a Reader for the PredictRunAPINamePost structure.
type PredictRunAPINamePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PredictRunAPINamePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPredictRunAPINamePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewPredictRunAPINamePostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPredictRunAPINamePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPredictRunAPINamePostOK creates a PredictRunAPINamePostOK with default headers values
func NewPredictRunAPINamePostOK() *PredictRunAPINamePostOK {
	return &PredictRunAPINamePostOK{}
}

/*
PredictRunAPINamePostOK describes a response with status code 200, with default header values.

Successful Response
*/
type PredictRunAPINamePostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this predict run Api name post o k response has a 2xx status code
func (o *PredictRunAPINamePostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this predict run Api name post o k response has a 3xx status code
func (o *PredictRunAPINamePostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this predict run Api name post o k response has a 4xx status code
func (o *PredictRunAPINamePostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this predict run Api name post o k response has a 5xx status code
func (o *PredictRunAPINamePostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this predict run Api name post o k response a status code equal to that given
func (o *PredictRunAPINamePostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the predict run Api name post o k response
func (o *PredictRunAPINamePostOK) Code() int {
	return 200
}

func (o *PredictRunAPINamePostOK) Error() string {
	return fmt.Sprintf("[POST /run/{api_name}][%d] predictRunApiNamePostOK  %+v", 200, o.Payload)
}

func (o *PredictRunAPINamePostOK) String() string {
	return fmt.Sprintf("[POST /run/{api_name}][%d] predictRunApiNamePostOK  %+v", 200, o.Payload)
}

func (o *PredictRunAPINamePostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PredictRunAPINamePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPredictRunAPINamePostUnprocessableEntity creates a PredictRunAPINamePostUnprocessableEntity with default headers values
func NewPredictRunAPINamePostUnprocessableEntity() *PredictRunAPINamePostUnprocessableEntity {
	return &PredictRunAPINamePostUnprocessableEntity{}
}

/*
PredictRunAPINamePostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type PredictRunAPINamePostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this predict run Api name post unprocessable entity response has a 2xx status code
func (o *PredictRunAPINamePostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this predict run Api name post unprocessable entity response has a 3xx status code
func (o *PredictRunAPINamePostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this predict run Api name post unprocessable entity response has a 4xx status code
func (o *PredictRunAPINamePostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this predict run Api name post unprocessable entity response has a 5xx status code
func (o *PredictRunAPINamePostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this predict run Api name post unprocessable entity response a status code equal to that given
func (o *PredictRunAPINamePostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the predict run Api name post unprocessable entity response
func (o *PredictRunAPINamePostUnprocessableEntity) Code() int {
	return 422
}

func (o *PredictRunAPINamePostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /run/{api_name}][%d] predictRunApiNamePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PredictRunAPINamePostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /run/{api_name}][%d] predictRunApiNamePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PredictRunAPINamePostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *PredictRunAPINamePostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPredictRunAPINamePostInternalServerError creates a PredictRunAPINamePostInternalServerError with default headers values
func NewPredictRunAPINamePostInternalServerError() *PredictRunAPINamePostInternalServerError {
	return &PredictRunAPINamePostInternalServerError{}
}

/*
PredictRunAPINamePostInternalServerError describes a response with status code 500, with default header values.

HTTPException
*/
type PredictRunAPINamePostInternalServerError struct {
	Payload *models.HTTPException
}

// IsSuccess returns true when this predict run Api name post internal server error response has a 2xx status code
func (o *PredictRunAPINamePostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this predict run Api name post internal server error response has a 3xx status code
func (o *PredictRunAPINamePostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this predict run Api name post internal server error response has a 4xx status code
func (o *PredictRunAPINamePostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this predict run Api name post internal server error response has a 5xx status code
func (o *PredictRunAPINamePostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this predict run Api name post internal server error response a status code equal to that given
func (o *PredictRunAPINamePostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the predict run Api name post internal server error response
func (o *PredictRunAPINamePostInternalServerError) Code() int {
	return 500
}

func (o *PredictRunAPINamePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /run/{api_name}][%d] predictRunApiNamePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PredictRunAPINamePostInternalServerError) String() string {
	return fmt.Sprintf("[POST /run/{api_name}][%d] predictRunApiNamePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PredictRunAPINamePostInternalServerError) GetPayload() *models.HTTPException {
	return o.Payload
}

func (o *PredictRunAPINamePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
