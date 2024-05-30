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

// GetUpscalersSdapiV1UpscalersGetReader is a Reader for the GetUpscalersSdapiV1UpscalersGet structure.
type GetUpscalersSdapiV1UpscalersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUpscalersSdapiV1UpscalersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUpscalersSdapiV1UpscalersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUpscalersSdapiV1UpscalersGetOK creates a GetUpscalersSdapiV1UpscalersGetOK with default headers values
func NewGetUpscalersSdapiV1UpscalersGetOK() *GetUpscalersSdapiV1UpscalersGetOK {
	return &GetUpscalersSdapiV1UpscalersGetOK{}
}

/*
GetUpscalersSdapiV1UpscalersGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetUpscalersSdapiV1UpscalersGetOK struct {
	Payload []*models.UpscalerItem
}

// IsSuccess returns true when this get upscalers sdapi v1 upscalers get o k response has a 2xx status code
func (o *GetUpscalersSdapiV1UpscalersGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get upscalers sdapi v1 upscalers get o k response has a 3xx status code
func (o *GetUpscalersSdapiV1UpscalersGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upscalers sdapi v1 upscalers get o k response has a 4xx status code
func (o *GetUpscalersSdapiV1UpscalersGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upscalers sdapi v1 upscalers get o k response has a 5xx status code
func (o *GetUpscalersSdapiV1UpscalersGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get upscalers sdapi v1 upscalers get o k response a status code equal to that given
func (o *GetUpscalersSdapiV1UpscalersGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get upscalers sdapi v1 upscalers get o k response
func (o *GetUpscalersSdapiV1UpscalersGetOK) Code() int {
	return 200
}

func (o *GetUpscalersSdapiV1UpscalersGetOK) Error() string {
	return fmt.Sprintf("[GET /sdapi/v1/upscalers][%d] getUpscalersSdapiV1UpscalersGetOK  %+v", 200, o.Payload)
}

func (o *GetUpscalersSdapiV1UpscalersGetOK) String() string {
	return fmt.Sprintf("[GET /sdapi/v1/upscalers][%d] getUpscalersSdapiV1UpscalersGetOK  %+v", 200, o.Payload)
}

func (o *GetUpscalersSdapiV1UpscalersGetOK) GetPayload() []*models.UpscalerItem {
	return o.Payload
}

func (o *GetUpscalersSdapiV1UpscalersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
