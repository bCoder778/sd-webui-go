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

// TrainEmbeddingSdapiV1TrainEmbeddingPostReader is a Reader for the TrainEmbeddingSdapiV1TrainEmbeddingPost structure.
type TrainEmbeddingSdapiV1TrainEmbeddingPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTrainEmbeddingSdapiV1TrainEmbeddingPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewTrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTrainEmbeddingSdapiV1TrainEmbeddingPostOK creates a TrainEmbeddingSdapiV1TrainEmbeddingPostOK with default headers values
func NewTrainEmbeddingSdapiV1TrainEmbeddingPostOK() *TrainEmbeddingSdapiV1TrainEmbeddingPostOK {
	return &TrainEmbeddingSdapiV1TrainEmbeddingPostOK{}
}

/*
TrainEmbeddingSdapiV1TrainEmbeddingPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type TrainEmbeddingSdapiV1TrainEmbeddingPostOK struct {
	Payload *models.TrainResponse
}

// IsSuccess returns true when this train embedding sdapi v1 train embedding post o k response has a 2xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this train embedding sdapi v1 train embedding post o k response has a 3xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this train embedding sdapi v1 train embedding post o k response has a 4xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this train embedding sdapi v1 train embedding post o k response has a 5xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this train embedding sdapi v1 train embedding post o k response a status code equal to that given
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the train embedding sdapi v1 train embedding post o k response
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostOK) Code() int {
	return 200
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostOK) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/train/embedding][%d] trainEmbeddingSdapiV1TrainEmbeddingPostOK  %+v", 200, o.Payload)
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostOK) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/train/embedding][%d] trainEmbeddingSdapiV1TrainEmbeddingPostOK  %+v", 200, o.Payload)
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostOK) GetPayload() *models.TrainResponse {
	return o.Payload
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrainResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity creates a TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity with default headers values
func NewTrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity() *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity {
	return &TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity{}
}

/*
TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this train embedding sdapi v1 train embedding post unprocessable entity response has a 2xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this train embedding sdapi v1 train embedding post unprocessable entity response has a 3xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this train embedding sdapi v1 train embedding post unprocessable entity response has a 4xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this train embedding sdapi v1 train embedding post unprocessable entity response has a 5xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this train embedding sdapi v1 train embedding post unprocessable entity response a status code equal to that given
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the train embedding sdapi v1 train embedding post unprocessable entity response
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity) Code() int {
	return 422
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/train/embedding][%d] trainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/train/embedding][%d] trainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError creates a TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError with default headers values
func NewTrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError() *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError {
	return &TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError{}
}

/*
TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError describes a response with status code 500, with default header values.

HTTPException
*/
type TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError struct {
	Payload *models.HTTPException
}

// IsSuccess returns true when this train embedding sdapi v1 train embedding post internal server error response has a 2xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this train embedding sdapi v1 train embedding post internal server error response has a 3xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this train embedding sdapi v1 train embedding post internal server error response has a 4xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this train embedding sdapi v1 train embedding post internal server error response has a 5xx status code
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this train embedding sdapi v1 train embedding post internal server error response a status code equal to that given
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the train embedding sdapi v1 train embedding post internal server error response
func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError) Code() int {
	return 500
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /sdapi/v1/train/embedding][%d] trainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError  %+v", 500, o.Payload)
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /sdapi/v1/train/embedding][%d] trainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError  %+v", 500, o.Payload)
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError) GetPayload() *models.HTTPException {
	return o.Payload
}

func (o *TrainEmbeddingSdapiV1TrainEmbeddingPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
