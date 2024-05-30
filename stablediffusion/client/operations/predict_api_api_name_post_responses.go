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

// PredictAPIAPINamePostReader is a Reader for the PredictAPIAPINamePost structure.
type PredictAPIAPINamePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PredictAPIAPINamePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPredictAPIAPINamePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewPredictAPIAPINamePostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPredictAPIAPINamePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPredictAPIAPINamePostOK creates a PredictAPIAPINamePostOK with default headers values
func NewPredictAPIAPINamePostOK() *PredictAPIAPINamePostOK {
	return &PredictAPIAPINamePostOK{}
}

/*
PredictAPIAPINamePostOK describes a response with status code 200, with default header values.

Successful Response
*/
type PredictAPIAPINamePostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this predict Api Api name post o k response has a 2xx status code
func (o *PredictAPIAPINamePostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this predict Api Api name post o k response has a 3xx status code
func (o *PredictAPIAPINamePostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this predict Api Api name post o k response has a 4xx status code
func (o *PredictAPIAPINamePostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this predict Api Api name post o k response has a 5xx status code
func (o *PredictAPIAPINamePostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this predict Api Api name post o k response a status code equal to that given
func (o *PredictAPIAPINamePostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the predict Api Api name post o k response
func (o *PredictAPIAPINamePostOK) Code() int {
	return 200
}

func (o *PredictAPIAPINamePostOK) Error() string {
	return fmt.Sprintf("[POST /api/{api_name}][%d] predictApiApiNamePostOK  %+v", 200, o.Payload)
}

func (o *PredictAPIAPINamePostOK) String() string {
	return fmt.Sprintf("[POST /api/{api_name}][%d] predictApiApiNamePostOK  %+v", 200, o.Payload)
}

func (o *PredictAPIAPINamePostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PredictAPIAPINamePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPredictAPIAPINamePostUnprocessableEntity creates a PredictAPIAPINamePostUnprocessableEntity with default headers values
func NewPredictAPIAPINamePostUnprocessableEntity() *PredictAPIAPINamePostUnprocessableEntity {
	return &PredictAPIAPINamePostUnprocessableEntity{}
}

/*
PredictAPIAPINamePostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type PredictAPIAPINamePostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this predict Api Api name post unprocessable entity response has a 2xx status code
func (o *PredictAPIAPINamePostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this predict Api Api name post unprocessable entity response has a 3xx status code
func (o *PredictAPIAPINamePostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this predict Api Api name post unprocessable entity response has a 4xx status code
func (o *PredictAPIAPINamePostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this predict Api Api name post unprocessable entity response has a 5xx status code
func (o *PredictAPIAPINamePostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this predict Api Api name post unprocessable entity response a status code equal to that given
func (o *PredictAPIAPINamePostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the predict Api Api name post unprocessable entity response
func (o *PredictAPIAPINamePostUnprocessableEntity) Code() int {
	return 422
}

func (o *PredictAPIAPINamePostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/{api_name}][%d] predictApiApiNamePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PredictAPIAPINamePostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/{api_name}][%d] predictApiApiNamePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PredictAPIAPINamePostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *PredictAPIAPINamePostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPredictAPIAPINamePostInternalServerError creates a PredictAPIAPINamePostInternalServerError with default headers values
func NewPredictAPIAPINamePostInternalServerError() *PredictAPIAPINamePostInternalServerError {
	return &PredictAPIAPINamePostInternalServerError{}
}

/*
PredictAPIAPINamePostInternalServerError describes a response with status code 500, with default header values.

HTTPException
*/
type PredictAPIAPINamePostInternalServerError struct {
	Payload *models.HTTPException
}

// IsSuccess returns true when this predict Api Api name post internal server error response has a 2xx status code
func (o *PredictAPIAPINamePostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this predict Api Api name post internal server error response has a 3xx status code
func (o *PredictAPIAPINamePostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this predict Api Api name post internal server error response has a 4xx status code
func (o *PredictAPIAPINamePostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this predict Api Api name post internal server error response has a 5xx status code
func (o *PredictAPIAPINamePostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this predict Api Api name post internal server error response a status code equal to that given
func (o *PredictAPIAPINamePostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the predict Api Api name post internal server error response
func (o *PredictAPIAPINamePostInternalServerError) Code() int {
	return 500
}

func (o *PredictAPIAPINamePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/{api_name}][%d] predictApiApiNamePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PredictAPIAPINamePostInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/{api_name}][%d] predictApiApiNamePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PredictAPIAPINamePostInternalServerError) GetPayload() *models.HTTPException {
	return o.Payload
}

func (o *PredictAPIAPINamePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
