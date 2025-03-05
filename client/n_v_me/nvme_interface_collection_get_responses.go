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

	"github.com/sapcc/ontap-restapi/models"
)

// NvmeInterfaceCollectionGetReader is a Reader for the NvmeInterfaceCollectionGet structure.
type NvmeInterfaceCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeInterfaceCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeInterfaceCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeInterfaceCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeInterfaceCollectionGetOK creates a NvmeInterfaceCollectionGetOK with default headers values
func NewNvmeInterfaceCollectionGetOK() *NvmeInterfaceCollectionGetOK {
	return &NvmeInterfaceCollectionGetOK{}
}

/*
NvmeInterfaceCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeInterfaceCollectionGetOK struct {
	Payload *models.NvmeInterfaceResponse
}

// IsSuccess returns true when this nvme interface collection get o k response has a 2xx status code
func (o *NvmeInterfaceCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme interface collection get o k response has a 3xx status code
func (o *NvmeInterfaceCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme interface collection get o k response has a 4xx status code
func (o *NvmeInterfaceCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme interface collection get o k response has a 5xx status code
func (o *NvmeInterfaceCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme interface collection get o k response a status code equal to that given
func (o *NvmeInterfaceCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme interface collection get o k response
func (o *NvmeInterfaceCollectionGetOK) Code() int {
	return 200
}

func (o *NvmeInterfaceCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/interfaces][%d] nvmeInterfaceCollectionGetOK %s", 200, payload)
}

func (o *NvmeInterfaceCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/interfaces][%d] nvmeInterfaceCollectionGetOK %s", 200, payload)
}

func (o *NvmeInterfaceCollectionGetOK) GetPayload() *models.NvmeInterfaceResponse {
	return o.Payload
}

func (o *NvmeInterfaceCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeInterfaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeInterfaceCollectionGetDefault creates a NvmeInterfaceCollectionGetDefault with default headers values
func NewNvmeInterfaceCollectionGetDefault(code int) *NvmeInterfaceCollectionGetDefault {
	return &NvmeInterfaceCollectionGetDefault{
		_statusCode: code,
	}
}

/*
NvmeInterfaceCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type NvmeInterfaceCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme interface collection get default response has a 2xx status code
func (o *NvmeInterfaceCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme interface collection get default response has a 3xx status code
func (o *NvmeInterfaceCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme interface collection get default response has a 4xx status code
func (o *NvmeInterfaceCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme interface collection get default response has a 5xx status code
func (o *NvmeInterfaceCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme interface collection get default response a status code equal to that given
func (o *NvmeInterfaceCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme interface collection get default response
func (o *NvmeInterfaceCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NvmeInterfaceCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/interfaces][%d] nvme_interface_collection_get default %s", o._statusCode, payload)
}

func (o *NvmeInterfaceCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/interfaces][%d] nvme_interface_collection_get default %s", o._statusCode, payload)
}

func (o *NvmeInterfaceCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeInterfaceCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
