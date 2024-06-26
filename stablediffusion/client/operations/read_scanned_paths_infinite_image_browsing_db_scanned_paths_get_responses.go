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

// ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetReader is a Reader for the ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGet structure.
type ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK creates a ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK with default headers values
func NewReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK() *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK {
	return &ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK{}
}

/*
ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK describes a response with status code 200, with default header values.

Successful Response
*/
type ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK struct {
	Payload []*models.ScannedPathModel
}

// IsSuccess returns true when this read scanned paths infinite image browsing db scanned paths get o k response has a 2xx status code
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read scanned paths infinite image browsing db scanned paths get o k response has a 3xx status code
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read scanned paths infinite image browsing db scanned paths get o k response has a 4xx status code
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read scanned paths infinite image browsing db scanned paths get o k response has a 5xx status code
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read scanned paths infinite image browsing db scanned paths get o k response a status code equal to that given
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read scanned paths infinite image browsing db scanned paths get o k response
func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK) Code() int {
	return 200
}

func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK) Error() string {
	return fmt.Sprintf("[GET /infinite_image_browsing/db/scanned_paths][%d] readScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK  %+v", 200, o.Payload)
}

func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK) String() string {
	return fmt.Sprintf("[GET /infinite_image_browsing/db/scanned_paths][%d] readScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK  %+v", 200, o.Payload)
}

func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK) GetPayload() []*models.ScannedPathModel {
	return o.Payload
}

func (o *ReadScannedPathsInfiniteImageBrowsingDbScannedPathsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
