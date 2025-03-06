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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NvmeSubsystemHostCollectionGetReader is a Reader for the NvmeSubsystemHostCollectionGet structure.
type NvmeSubsystemHostCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemHostCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeSubsystemHostCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemHostCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemHostCollectionGetOK creates a NvmeSubsystemHostCollectionGetOK with default headers values
func NewNvmeSubsystemHostCollectionGetOK() *NvmeSubsystemHostCollectionGetOK {
	return &NvmeSubsystemHostCollectionGetOK{}
}

/*
NvmeSubsystemHostCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeSubsystemHostCollectionGetOK struct {
	Payload *models.NvmeSubsystemHostResponse
}

// IsSuccess returns true when this nvme subsystem host collection get o k response has a 2xx status code
func (o *NvmeSubsystemHostCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme subsystem host collection get o k response has a 3xx status code
func (o *NvmeSubsystemHostCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme subsystem host collection get o k response has a 4xx status code
func (o *NvmeSubsystemHostCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme subsystem host collection get o k response has a 5xx status code
func (o *NvmeSubsystemHostCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme subsystem host collection get o k response a status code equal to that given
func (o *NvmeSubsystemHostCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme subsystem host collection get o k response
func (o *NvmeSubsystemHostCollectionGetOK) Code() int {
	return 200
}

func (o *NvmeSubsystemHostCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems/{subsystem.uuid}/hosts][%d] nvmeSubsystemHostCollectionGetOK %s", 200, payload)
}

func (o *NvmeSubsystemHostCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems/{subsystem.uuid}/hosts][%d] nvmeSubsystemHostCollectionGetOK %s", 200, payload)
}

func (o *NvmeSubsystemHostCollectionGetOK) GetPayload() *models.NvmeSubsystemHostResponse {
	return o.Payload
}

func (o *NvmeSubsystemHostCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeSubsystemHostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeSubsystemHostCollectionGetDefault creates a NvmeSubsystemHostCollectionGetDefault with default headers values
func NewNvmeSubsystemHostCollectionGetDefault(code int) *NvmeSubsystemHostCollectionGetDefault {
	return &NvmeSubsystemHostCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	NvmeSubsystemHostCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 72090001 | The NVMe subsystem does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeSubsystemHostCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme subsystem host collection get default response has a 2xx status code
func (o *NvmeSubsystemHostCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme subsystem host collection get default response has a 3xx status code
func (o *NvmeSubsystemHostCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme subsystem host collection get default response has a 4xx status code
func (o *NvmeSubsystemHostCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme subsystem host collection get default response has a 5xx status code
func (o *NvmeSubsystemHostCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme subsystem host collection get default response a status code equal to that given
func (o *NvmeSubsystemHostCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme subsystem host collection get default response
func (o *NvmeSubsystemHostCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NvmeSubsystemHostCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems/{subsystem.uuid}/hosts][%d] nvme_subsystem_host_collection_get default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemHostCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems/{subsystem.uuid}/hosts][%d] nvme_subsystem_host_collection_get default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemHostCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemHostCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
