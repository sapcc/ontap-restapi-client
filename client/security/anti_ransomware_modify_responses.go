// Code generated by go-swagger; DO NOT EDIT.

package security

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

// AntiRansomwareModifyReader is a Reader for the AntiRansomwareModify structure.
type AntiRansomwareModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AntiRansomwareModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAntiRansomwareModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAntiRansomwareModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAntiRansomwareModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAntiRansomwareModifyOK creates a AntiRansomwareModifyOK with default headers values
func NewAntiRansomwareModifyOK() *AntiRansomwareModifyOK {
	return &AntiRansomwareModifyOK{}
}

/*
AntiRansomwareModifyOK describes a response with status code 200, with default header values.

OK
*/
type AntiRansomwareModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this anti ransomware modify o k response has a 2xx status code
func (o *AntiRansomwareModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this anti ransomware modify o k response has a 3xx status code
func (o *AntiRansomwareModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this anti ransomware modify o k response has a 4xx status code
func (o *AntiRansomwareModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this anti ransomware modify o k response has a 5xx status code
func (o *AntiRansomwareModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this anti ransomware modify o k response a status code equal to that given
func (o *AntiRansomwareModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the anti ransomware modify o k response
func (o *AntiRansomwareModifyOK) Code() int {
	return 200
}

func (o *AntiRansomwareModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/anti-ransomware][%d] antiRansomwareModifyOK %s", 200, payload)
}

func (o *AntiRansomwareModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/anti-ransomware][%d] antiRansomwareModifyOK %s", 200, payload)
}

func (o *AntiRansomwareModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *AntiRansomwareModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAntiRansomwareModifyAccepted creates a AntiRansomwareModifyAccepted with default headers values
func NewAntiRansomwareModifyAccepted() *AntiRansomwareModifyAccepted {
	return &AntiRansomwareModifyAccepted{}
}

/*
AntiRansomwareModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AntiRansomwareModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this anti ransomware modify accepted response has a 2xx status code
func (o *AntiRansomwareModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this anti ransomware modify accepted response has a 3xx status code
func (o *AntiRansomwareModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this anti ransomware modify accepted response has a 4xx status code
func (o *AntiRansomwareModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this anti ransomware modify accepted response has a 5xx status code
func (o *AntiRansomwareModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this anti ransomware modify accepted response a status code equal to that given
func (o *AntiRansomwareModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the anti ransomware modify accepted response
func (o *AntiRansomwareModifyAccepted) Code() int {
	return 202
}

func (o *AntiRansomwareModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/anti-ransomware][%d] antiRansomwareModifyAccepted %s", 202, payload)
}

func (o *AntiRansomwareModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/anti-ransomware][%d] antiRansomwareModifyAccepted %s", 202, payload)
}

func (o *AntiRansomwareModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *AntiRansomwareModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAntiRansomwareModifyDefault creates a AntiRansomwareModifyDefault with default headers values
func NewAntiRansomwareModifyDefault(code int) *AntiRansomwareModifyDefault {
	return &AntiRansomwareModifyDefault{
		_statusCode: code,
	}
}

/*
	AntiRansomwareModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 203161684 | Autonomous Ransomware Protection (ARP) package file must be either a .tar or .tgz file. |
| 203161685 | Autonomous Ransomware Protection (ARP) package file download failed. |
| 203161686 | Autonomous Ransomware Protection (ARP) package file unzip failed. |
| 203161687 | Autonomous Ransomware Protection (ARP) package file untar failed. |
| 203161688 | Autonomous Ransomware Protection (ARP) package archive invalid. No ARP version file found. |
| 203161689 | Internal error. Filesystem error while installing Autonomous Ransomware Protection (ARP) package file. |
| 203161690 | Autonomous Ransomware Protection (ARP) package update requires an effective cluster version of 9.16.1 or later. |
| 203161691 | Autonomous Ransomware Protection (ARP) package file is invalid. The checksum verification of the ARP configuration files failed. |
| 203161692 | Autonomous Ransomware Protection (ARP) package update failed because one or more nodes are not healthy. |
| 203161694 | Failed to update the Autonomous Ransomware Protection (ARP) package file because the `uri` property was not specified. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AntiRansomwareModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this anti ransomware modify default response has a 2xx status code
func (o *AntiRansomwareModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this anti ransomware modify default response has a 3xx status code
func (o *AntiRansomwareModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this anti ransomware modify default response has a 4xx status code
func (o *AntiRansomwareModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this anti ransomware modify default response has a 5xx status code
func (o *AntiRansomwareModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this anti ransomware modify default response a status code equal to that given
func (o *AntiRansomwareModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the anti ransomware modify default response
func (o *AntiRansomwareModifyDefault) Code() int {
	return o._statusCode
}

func (o *AntiRansomwareModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/anti-ransomware][%d] anti_ransomware_modify default %s", o._statusCode, payload)
}

func (o *AntiRansomwareModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/anti-ransomware][%d] anti_ransomware_modify default %s", o._statusCode, payload)
}

func (o *AntiRansomwareModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AntiRansomwareModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
