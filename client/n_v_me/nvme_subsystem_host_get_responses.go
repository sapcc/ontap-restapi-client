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

// NvmeSubsystemHostGetReader is a Reader for the NvmeSubsystemHostGet structure.
type NvmeSubsystemHostGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemHostGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeSubsystemHostGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemHostGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemHostGetOK creates a NvmeSubsystemHostGetOK with default headers values
func NewNvmeSubsystemHostGetOK() *NvmeSubsystemHostGetOK {
	return &NvmeSubsystemHostGetOK{}
}

/*
NvmeSubsystemHostGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeSubsystemHostGetOK struct {
	Payload *models.NvmeSubsystemHost
}

// IsSuccess returns true when this nvme subsystem host get o k response has a 2xx status code
func (o *NvmeSubsystemHostGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme subsystem host get o k response has a 3xx status code
func (o *NvmeSubsystemHostGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme subsystem host get o k response has a 4xx status code
func (o *NvmeSubsystemHostGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme subsystem host get o k response has a 5xx status code
func (o *NvmeSubsystemHostGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme subsystem host get o k response a status code equal to that given
func (o *NvmeSubsystemHostGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme subsystem host get o k response
func (o *NvmeSubsystemHostGetOK) Code() int {
	return 200
}

func (o *NvmeSubsystemHostGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems/{subsystem.uuid}/hosts/{nqn}][%d] nvmeSubsystemHostGetOK %s", 200, payload)
}

func (o *NvmeSubsystemHostGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems/{subsystem.uuid}/hosts/{nqn}][%d] nvmeSubsystemHostGetOK %s", 200, payload)
}

func (o *NvmeSubsystemHostGetOK) GetPayload() *models.NvmeSubsystemHost {
	return o.Payload
}

func (o *NvmeSubsystemHostGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeSubsystemHost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeSubsystemHostGetDefault creates a NvmeSubsystemHostGetDefault with default headers values
func NewNvmeSubsystemHostGetDefault(code int) *NvmeSubsystemHostGetDefault {
	return &NvmeSubsystemHostGetDefault{
		_statusCode: code,
	}
}

/*
	NvmeSubsystemHostGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 72090001 | The NVMe subsystem does not exist. |
| 72090004 | The NVMe subsystem host does not exist. |
| 72090022 | The NVMe subsystem host does not exist in the specified subsystem. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeSubsystemHostGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme subsystem host get default response has a 2xx status code
func (o *NvmeSubsystemHostGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme subsystem host get default response has a 3xx status code
func (o *NvmeSubsystemHostGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme subsystem host get default response has a 4xx status code
func (o *NvmeSubsystemHostGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme subsystem host get default response has a 5xx status code
func (o *NvmeSubsystemHostGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme subsystem host get default response a status code equal to that given
func (o *NvmeSubsystemHostGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme subsystem host get default response
func (o *NvmeSubsystemHostGetDefault) Code() int {
	return o._statusCode
}

func (o *NvmeSubsystemHostGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems/{subsystem.uuid}/hosts/{nqn}][%d] nvme_subsystem_host_get default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemHostGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nvme/subsystems/{subsystem.uuid}/hosts/{nqn}][%d] nvme_subsystem_host_get default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemHostGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemHostGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
