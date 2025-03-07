// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NvmeNamespaceGetReader is a Reader for the NvmeNamespaceGet structure.
type NvmeNamespaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeNamespaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeNamespaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeNamespaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeNamespaceGetOK creates a NvmeNamespaceGetOK with default headers values
func NewNvmeNamespaceGetOK() *NvmeNamespaceGetOK {
	return &NvmeNamespaceGetOK{}
}

/*
NvmeNamespaceGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeNamespaceGetOK struct {
	Payload *models.NvmeNamespace
}

// IsSuccess returns true when this nvme namespace get o k response has a 2xx status code
func (o *NvmeNamespaceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme namespace get o k response has a 3xx status code
func (o *NvmeNamespaceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme namespace get o k response has a 4xx status code
func (o *NvmeNamespaceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme namespace get o k response has a 5xx status code
func (o *NvmeNamespaceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme namespace get o k response a status code equal to that given
func (o *NvmeNamespaceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme namespace get o k response
func (o *NvmeNamespaceGetOK) Code() int {
	return 200
}

func (o *NvmeNamespaceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/namespaces/{uuid}][%d] nvmeNamespaceGetOK %s", 200, payload)
}

func (o *NvmeNamespaceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/namespaces/{uuid}][%d] nvmeNamespaceGetOK %s", 200, payload)
}

func (o *NvmeNamespaceGetOK) GetPayload() *models.NvmeNamespace {
	return o.Payload
}

func (o *NvmeNamespaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeNamespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeNamespaceGetDefault creates a NvmeNamespaceGetDefault with default headers values
func NewNvmeNamespaceGetDefault(code int) *NvmeNamespaceGetDefault {
	return &NvmeNamespaceGetDefault{
		_statusCode: code,
	}
}

/*
	NvmeNamespaceGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 72090006 | The specified namespace was not found. |
| 72090007 | The specified namespace was not found. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeNamespaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme namespace get default response has a 2xx status code
func (o *NvmeNamespaceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme namespace get default response has a 3xx status code
func (o *NvmeNamespaceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme namespace get default response has a 4xx status code
func (o *NvmeNamespaceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme namespace get default response has a 5xx status code
func (o *NvmeNamespaceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme namespace get default response a status code equal to that given
func (o *NvmeNamespaceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme namespace get default response
func (o *NvmeNamespaceGetDefault) Code() int {
	return o._statusCode
}

func (o *NvmeNamespaceGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/namespaces/{uuid}][%d] nvme_namespace_get default %s", o._statusCode, payload)
}

func (o *NvmeNamespaceGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/namespaces/{uuid}][%d] nvme_namespace_get default %s", o._statusCode, payload)
}

func (o *NvmeNamespaceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeNamespaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
