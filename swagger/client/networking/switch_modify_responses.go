// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// SwitchModifyReader is a Reader for the SwitchModify structure.
type SwitchModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwitchModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSwitchModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSwitchModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSwitchModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSwitchModifyOK creates a SwitchModifyOK with default headers values
func NewSwitchModifyOK() *SwitchModifyOK {
	return &SwitchModifyOK{}
}

/*
SwitchModifyOK describes a response with status code 200, with default header values.

OK
*/
type SwitchModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this switch modify o k response has a 2xx status code
func (o *SwitchModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this switch modify o k response has a 3xx status code
func (o *SwitchModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this switch modify o k response has a 4xx status code
func (o *SwitchModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this switch modify o k response has a 5xx status code
func (o *SwitchModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this switch modify o k response a status code equal to that given
func (o *SwitchModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the switch modify o k response
func (o *SwitchModifyOK) Code() int {
	return 200
}

func (o *SwitchModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches/{name}][%d] switchModifyOK %s", 200, payload)
}

func (o *SwitchModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches/{name}][%d] switchModifyOK %s", 200, payload)
}

func (o *SwitchModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SwitchModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwitchModifyAccepted creates a SwitchModifyAccepted with default headers values
func NewSwitchModifyAccepted() *SwitchModifyAccepted {
	return &SwitchModifyAccepted{}
}

/*
SwitchModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SwitchModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this switch modify accepted response has a 2xx status code
func (o *SwitchModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this switch modify accepted response has a 3xx status code
func (o *SwitchModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this switch modify accepted response has a 4xx status code
func (o *SwitchModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this switch modify accepted response has a 5xx status code
func (o *SwitchModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this switch modify accepted response a status code equal to that given
func (o *SwitchModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the switch modify accepted response
func (o *SwitchModifyAccepted) Code() int {
	return 202
}

func (o *SwitchModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches/{name}][%d] switchModifyAccepted %s", 202, payload)
}

func (o *SwitchModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches/{name}][%d] switchModifyAccepted %s", 202, payload)
}

func (o *SwitchModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SwitchModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwitchModifyDefault creates a SwitchModifyDefault with default headers values
func NewSwitchModifyDefault(code int) *SwitchModifyDefault {
	return &SwitchModifyDefault{
		_statusCode: code,
	}
}

/*
	SwitchModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 12517379 | Settings updated, but the IP address \"{address}\" is not reachable. Verify that the address is valid or check the network path. |
| 12517381 | Settings updated, but the SNMP validation request timed out. Verify that the \"snmp.user\" parameter is valid. |
| 12517383 | Settings updated, but the SNMP validation request timed out. Verify that the \"snmp.user\" parameter is valid (i.e., the SNMPv3 user exists in ONTAP and on the remote switch). If the \"snmp.user\" parameter is valid, verify that the SNMPv3 user's credentials are the same both in ONTAP as well as in the remote switch. If a custom engine-id was provided for the SNMPv3 user, ensure it is the same as that of the remote switch. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SwitchModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this switch modify default response has a 2xx status code
func (o *SwitchModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this switch modify default response has a 3xx status code
func (o *SwitchModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this switch modify default response has a 4xx status code
func (o *SwitchModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this switch modify default response has a 5xx status code
func (o *SwitchModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this switch modify default response a status code equal to that given
func (o *SwitchModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the switch modify default response
func (o *SwitchModifyDefault) Code() int {
	return o._statusCode
}

func (o *SwitchModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches/{name}][%d] switch_modify default %s", o._statusCode, payload)
}

func (o *SwitchModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches/{name}][%d] switch_modify default %s", o._statusCode, payload)
}

func (o *SwitchModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwitchModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
