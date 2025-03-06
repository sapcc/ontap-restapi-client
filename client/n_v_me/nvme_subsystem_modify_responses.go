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

// NvmeSubsystemModifyReader is a Reader for the NvmeSubsystemModify structure.
type NvmeSubsystemModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeSubsystemModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemModifyOK creates a NvmeSubsystemModifyOK with default headers values
func NewNvmeSubsystemModifyOK() *NvmeSubsystemModifyOK {
	return &NvmeSubsystemModifyOK{}
}

/*
NvmeSubsystemModifyOK describes a response with status code 200, with default header values.

OK
*/
type NvmeSubsystemModifyOK struct {
}

// IsSuccess returns true when this nvme subsystem modify o k response has a 2xx status code
func (o *NvmeSubsystemModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme subsystem modify o k response has a 3xx status code
func (o *NvmeSubsystemModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme subsystem modify o k response has a 4xx status code
func (o *NvmeSubsystemModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme subsystem modify o k response has a 5xx status code
func (o *NvmeSubsystemModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme subsystem modify o k response a status code equal to that given
func (o *NvmeSubsystemModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme subsystem modify o k response
func (o *NvmeSubsystemModifyOK) Code() int {
	return 200
}

func (o *NvmeSubsystemModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/nvme/subsystems/{uuid}][%d] nvmeSubsystemModifyOK", 200)
}

func (o *NvmeSubsystemModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/nvme/subsystems/{uuid}][%d] nvmeSubsystemModifyOK", 200)
}

func (o *NvmeSubsystemModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNvmeSubsystemModifyDefault creates a NvmeSubsystemModifyDefault with default headers values
func NewNvmeSubsystemModifyDefault(code int) *NvmeSubsystemModifyDefault {
	return &NvmeSubsystemModifyDefault{
		_statusCode: code,
	}
}

/*
	NvmeSubsystemModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 72090001 | The NVMe subsystem does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeSubsystemModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme subsystem modify default response has a 2xx status code
func (o *NvmeSubsystemModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme subsystem modify default response has a 3xx status code
func (o *NvmeSubsystemModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme subsystem modify default response has a 4xx status code
func (o *NvmeSubsystemModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme subsystem modify default response has a 5xx status code
func (o *NvmeSubsystemModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme subsystem modify default response a status code equal to that given
func (o *NvmeSubsystemModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme subsystem modify default response
func (o *NvmeSubsystemModifyDefault) Code() int {
	return o._statusCode
}

func (o *NvmeSubsystemModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nvme/subsystems/{uuid}][%d] nvme_subsystem_modify default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nvme/subsystems/{uuid}][%d] nvme_subsystem_modify default %s", o._statusCode, payload)
}

func (o *NvmeSubsystemModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
