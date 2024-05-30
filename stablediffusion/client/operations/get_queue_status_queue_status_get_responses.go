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

// GetQueueStatusQueueStatusGetReader is a Reader for the GetQueueStatusQueueStatusGet structure.
type GetQueueStatusQueueStatusGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueueStatusQueueStatusGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQueueStatusQueueStatusGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQueueStatusQueueStatusGetOK creates a GetQueueStatusQueueStatusGetOK with default headers values
func NewGetQueueStatusQueueStatusGetOK() *GetQueueStatusQueueStatusGetOK {
	return &GetQueueStatusQueueStatusGetOK{}
}

/*
GetQueueStatusQueueStatusGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetQueueStatusQueueStatusGetOK struct {
	Payload *models.Estimation
}

// IsSuccess returns true when this get queue status queue status get o k response has a 2xx status code
func (o *GetQueueStatusQueueStatusGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get queue status queue status get o k response has a 3xx status code
func (o *GetQueueStatusQueueStatusGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get queue status queue status get o k response has a 4xx status code
func (o *GetQueueStatusQueueStatusGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get queue status queue status get o k response has a 5xx status code
func (o *GetQueueStatusQueueStatusGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get queue status queue status get o k response a status code equal to that given
func (o *GetQueueStatusQueueStatusGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get queue status queue status get o k response
func (o *GetQueueStatusQueueStatusGetOK) Code() int {
	return 200
}

func (o *GetQueueStatusQueueStatusGetOK) Error() string {
	return fmt.Sprintf("[GET /queue/status][%d] getQueueStatusQueueStatusGetOK  %+v", 200, o.Payload)
}

func (o *GetQueueStatusQueueStatusGetOK) String() string {
	return fmt.Sprintf("[GET /queue/status][%d] getQueueStatusQueueStatusGetOK  %+v", 200, o.Payload)
}

func (o *GetQueueStatusQueueStatusGetOK) GetPayload() *models.Estimation {
	return o.Payload
}

func (o *GetQueueStatusQueueStatusGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Estimation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
