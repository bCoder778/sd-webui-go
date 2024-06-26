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

// ShutdownAppInfiniteImageBrowsingShutdownPostReader is a Reader for the ShutdownAppInfiniteImageBrowsingShutdownPost structure.
type ShutdownAppInfiniteImageBrowsingShutdownPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShutdownAppInfiniteImageBrowsingShutdownPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShutdownAppInfiniteImageBrowsingShutdownPostOK creates a ShutdownAppInfiniteImageBrowsingShutdownPostOK with default headers values
func NewShutdownAppInfiniteImageBrowsingShutdownPostOK() *ShutdownAppInfiniteImageBrowsingShutdownPostOK {
	return &ShutdownAppInfiniteImageBrowsingShutdownPostOK{}
}

/*
ShutdownAppInfiniteImageBrowsingShutdownPostOK describes a response with status code 200, with default header values.

Successful Response
*/
type ShutdownAppInfiniteImageBrowsingShutdownPostOK struct {
	Payload interface{}
}

// IsSuccess returns true when this shutdown app infinite image browsing shutdown post o k response has a 2xx status code
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this shutdown app infinite image browsing shutdown post o k response has a 3xx status code
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this shutdown app infinite image browsing shutdown post o k response has a 4xx status code
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this shutdown app infinite image browsing shutdown post o k response has a 5xx status code
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this shutdown app infinite image browsing shutdown post o k response a status code equal to that given
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the shutdown app infinite image browsing shutdown post o k response
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostOK) Code() int {
	return 200
}

func (o *ShutdownAppInfiniteImageBrowsingShutdownPostOK) Error() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/shutdown][%d] shutdownAppInfiniteImageBrowsingShutdownPostOK  %+v", 200, o.Payload)
}

func (o *ShutdownAppInfiniteImageBrowsingShutdownPostOK) String() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/shutdown][%d] shutdownAppInfiniteImageBrowsingShutdownPostOK  %+v", 200, o.Payload)
}

func (o *ShutdownAppInfiniteImageBrowsingShutdownPostOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ShutdownAppInfiniteImageBrowsingShutdownPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError creates a ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError with default headers values
func NewShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError() *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError {
	return &ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError{}
}

/*
ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError describes a response with status code 500, with default header values.

HTTPException
*/
type ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError struct {
	Payload *models.HTTPException
}

// IsSuccess returns true when this shutdown app infinite image browsing shutdown post internal server error response has a 2xx status code
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this shutdown app infinite image browsing shutdown post internal server error response has a 3xx status code
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this shutdown app infinite image browsing shutdown post internal server error response has a 4xx status code
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this shutdown app infinite image browsing shutdown post internal server error response has a 5xx status code
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this shutdown app infinite image browsing shutdown post internal server error response a status code equal to that given
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the shutdown app infinite image browsing shutdown post internal server error response
func (o *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError) Code() int {
	return 500
}

func (o *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/shutdown][%d] shutdownAppInfiniteImageBrowsingShutdownPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /infinite_image_browsing/shutdown][%d] shutdownAppInfiniteImageBrowsingShutdownPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError) GetPayload() *models.HTTPException {
	return o.Payload
}

func (o *ShutdownAppInfiniteImageBrowsingShutdownPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
