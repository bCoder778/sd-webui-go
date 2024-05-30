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

// GetMetadataSdExtraNetworksMetadataGetReader is a Reader for the GetMetadataSdExtraNetworksMetadataGet structure.
type GetMetadataSdExtraNetworksMetadataGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetadataSdExtraNetworksMetadataGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetadataSdExtraNetworksMetadataGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetMetadataSdExtraNetworksMetadataGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMetadataSdExtraNetworksMetadataGetOK creates a GetMetadataSdExtraNetworksMetadataGetOK with default headers values
func NewGetMetadataSdExtraNetworksMetadataGetOK() *GetMetadataSdExtraNetworksMetadataGetOK {
	return &GetMetadataSdExtraNetworksMetadataGetOK{}
}

/*
GetMetadataSdExtraNetworksMetadataGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetMetadataSdExtraNetworksMetadataGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get metadata sd extra networks metadata get o k response has a 2xx status code
func (o *GetMetadataSdExtraNetworksMetadataGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get metadata sd extra networks metadata get o k response has a 3xx status code
func (o *GetMetadataSdExtraNetworksMetadataGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metadata sd extra networks metadata get o k response has a 4xx status code
func (o *GetMetadataSdExtraNetworksMetadataGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get metadata sd extra networks metadata get o k response has a 5xx status code
func (o *GetMetadataSdExtraNetworksMetadataGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get metadata sd extra networks metadata get o k response a status code equal to that given
func (o *GetMetadataSdExtraNetworksMetadataGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get metadata sd extra networks metadata get o k response
func (o *GetMetadataSdExtraNetworksMetadataGetOK) Code() int {
	return 200
}

func (o *GetMetadataSdExtraNetworksMetadataGetOK) Error() string {
	return fmt.Sprintf("[GET /sd_extra_networks/metadata][%d] getMetadataSdExtraNetworksMetadataGetOK  %+v", 200, o.Payload)
}

func (o *GetMetadataSdExtraNetworksMetadataGetOK) String() string {
	return fmt.Sprintf("[GET /sd_extra_networks/metadata][%d] getMetadataSdExtraNetworksMetadataGetOK  %+v", 200, o.Payload)
}

func (o *GetMetadataSdExtraNetworksMetadataGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetMetadataSdExtraNetworksMetadataGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetadataSdExtraNetworksMetadataGetUnprocessableEntity creates a GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity with default headers values
func NewGetMetadataSdExtraNetworksMetadataGetUnprocessableEntity() *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity {
	return &GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity{}
}

/*
GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity struct {
	Payload *models.HTTPValidationError
}

// IsSuccess returns true when this get metadata sd extra networks metadata get unprocessable entity response has a 2xx status code
func (o *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get metadata sd extra networks metadata get unprocessable entity response has a 3xx status code
func (o *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metadata sd extra networks metadata get unprocessable entity response has a 4xx status code
func (o *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get metadata sd extra networks metadata get unprocessable entity response has a 5xx status code
func (o *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get metadata sd extra networks metadata get unprocessable entity response a status code equal to that given
func (o *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get metadata sd extra networks metadata get unprocessable entity response
func (o *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity) Code() int {
	return 422
}

func (o *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /sd_extra_networks/metadata][%d] getMetadataSdExtraNetworksMetadataGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /sd_extra_networks/metadata][%d] getMetadataSdExtraNetworksMetadataGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity) GetPayload() *models.HTTPValidationError {
	return o.Payload
}

func (o *GetMetadataSdExtraNetworksMetadataGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
