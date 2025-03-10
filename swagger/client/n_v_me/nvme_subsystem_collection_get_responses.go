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

// NvmeSubsystemCollectionGetReader is a Reader for the NvmeSubsystemCollectionGet structure.
type NvmeSubsystemCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeSubsystemCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemCollectionGetOK creates a NvmeSubsystemCollectionGetOK with default headers values
func NewNvmeSubsystemCollectionGetOK() *NvmeSubsystemCollectionGetOK {
	return &NvmeSubsystemCollectionGetOK{}
}

/*
NvmeSubsystemCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeSubsystemCollectionGetOK struct {
	Payload *models.NvmeSubsystemResponse
}

// IsSuccess returns true when this nvme subsystem collection get o k response has a 2xx status code
func (o *NvmeSubsystemCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme subsystem collection get o k response has a 3xx status code
func (o *NvmeSubsystemCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme subsystem collection get o k response has a 4xx status code
func (o *NvmeSubsystemCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme subsystem collection get o k response has a 5xx status code
func (o *NvmeSubsystemCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme subsystem collection get o k response a status code equal to that given
func (o *NvmeSubsystemCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme subsystem collection get o k response
func (o *NvmeSubsystemCollectionGetOK) Code() int {
	return 200
}

func (o *NvmeSubsystemCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems][%d] nvmeSubsystemCollectionGetOK %s", 200, payload)
}

func (o *NvmeSubsystemCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems][%d] nvmeSubsystemCollectionGetOK %s", 200, payload)
}

func (o *NvmeSubsystemCollectionGetOK) GetPayload() *models.NvmeSubsystemResponse {
	return o.Payload
}

func (o *NvmeSubsystemCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeSubsystemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeSubsystemCollectionGetDefault creates a NvmeSubsystemCollectionGetDefault with default headers values
func NewNvmeSubsystemCollectionGetDefault(code int) *NvmeSubsystemCollectionGetDefault {
	return &NvmeSubsystemCollectionGetDefault{
		_statusCode: code,
	}
}

/*
NvmeSubsystemCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type NvmeSubsystemCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme subsystem collection get default response has a 2xx status code
func (o *NvmeSubsystemCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme subsystem collection get default response has a 3xx status code
func (o *NvmeSubsystemCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme subsystem collection get default response has a 4xx status code
func (o *NvmeSubsystemCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme subsystem collection get default response has a 5xx status code
func (o *NvmeSubsystemCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme subsystem collection get default response a status code equal to that given
func (o *NvmeSubsystemCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme subsystem collection get default response
func (o *NvmeSubsystemCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NvmeSubsystemCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems][%d] nvme_subsystem_collection_get default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems][%d] nvme_subsystem_collection_get default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
