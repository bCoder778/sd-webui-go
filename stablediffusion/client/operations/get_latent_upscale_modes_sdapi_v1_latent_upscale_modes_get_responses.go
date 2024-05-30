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

// GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetReader is a Reader for the GetLatentUpscaleModesSdapiV1LatentUpscaleModesGet structure.
type GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK creates a GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK with default headers values
func NewGetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK() *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK {
	return &GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK{}
}

/*
GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK struct {
	Payload []*models.LatentUpscalerModeItem
}

// IsSuccess returns true when this get latent upscale modes sdapi v1 latent upscale modes get o k response has a 2xx status code
func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get latent upscale modes sdapi v1 latent upscale modes get o k response has a 3xx status code
func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latent upscale modes sdapi v1 latent upscale modes get o k response has a 4xx status code
func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latent upscale modes sdapi v1 latent upscale modes get o k response has a 5xx status code
func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get latent upscale modes sdapi v1 latent upscale modes get o k response a status code equal to that given
func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get latent upscale modes sdapi v1 latent upscale modes get o k response
func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK) Code() int {
	return 200
}

func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK) Error() string {
	return fmt.Sprintf("[GET /sdapi/v1/latent-upscale-modes][%d] getLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK  %+v", 200, o.Payload)
}

func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK) String() string {
	return fmt.Sprintf("[GET /sdapi/v1/latent-upscale-modes][%d] getLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK  %+v", 200, o.Payload)
}

func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK) GetPayload() []*models.LatentUpscalerModeItem {
	return o.Payload
}

func (o *GetLatentUpscaleModesSdapiV1LatentUpscaleModesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
