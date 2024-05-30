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

// AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostReader is a Reader for the AddCustomTagInfiniteImageBrowsingDbAddCustomTagPost structure.
type AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewAddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK creates a AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK with default headers values
func NewAddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK() *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK {
	return &AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK{}
}

/*
AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this add custom tag infinite image browsing db add custom tag post o k response has a 2xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add custom tag infinite image browsing db add custom tag post o k response has a 3xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add custom tag infinite image browsing db add custom tag post o k response has a 4xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add custom tag infinite image browsing db add custom tag post o k response has a 5xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add custom tag infinite image browsing db add custom tag post o k response a status code equal to that given
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add custom tag infinite image browsing db add custom tag post o k response
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK) Code() int {
	return 200
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK) Error() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/db/add_custom_tag][%d] addCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK  %+v", 200, o.Payload)
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK) String() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/db/add_custom_tag][%d] addCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK  %+v", 200, o.Payload)
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity creates a AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity with default headers values
func NewAddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity() *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity {
	return &AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity{}
}

/*
AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this add custom tag infinite image browsing db add custom tag post unprocessable entity response has a 2xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add custom tag infinite image browsing db add custom tag post unprocessable entity response has a 3xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add custom tag infinite image browsing db add custom tag post unprocessable entity response has a 4xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this add custom tag infinite image browsing db add custom tag post unprocessable entity response has a 5xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this add custom tag infinite image browsing db add custom tag post unprocessable entity response a status code equal to that given
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the add custom tag infinite image browsing db add custom tag post unprocessable entity response
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity) Code() int {
	return 422
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/db/add_custom_tag][%d] addCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/db/add_custom_tag][%d] addCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError creates a AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError with default headers values
func NewAddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError() *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError {
	return &AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError{}
}

/*
AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError describes a response with status code 500, with default header values.

HTTPException
*/
type AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError struct {
	Payload *models.HTTPException
}

// IsSuccess returns true when this add custom tag infinite image browsing db add custom tag post internal server error response has a 2xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add custom tag infinite image browsing db add custom tag post internal server error response has a 3xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add custom tag infinite image browsing db add custom tag post internal server error response has a 4xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add custom tag infinite image browsing db add custom tag post internal server error response has a 5xx status code
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add custom tag infinite image browsing db add custom tag post internal server error response a status code equal to that given
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add custom tag infinite image browsing db add custom tag post internal server error response
func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError) Code() int {
	return 500
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/db/add_custom_tag][%d] addCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError  %+v", 500, o.Payload)
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/db/add_custom_tag][%d] addCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError  %+v", 500, o.Payload)
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError) GetPayload() *models.HTTPException {
	return o.Payload
}

func (o *AddCustomTagInfiniteImageBrowsingDbAddCustomTagPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
