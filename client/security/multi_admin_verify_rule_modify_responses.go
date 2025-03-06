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

// MultiAdminVerifyRuleModifyReader is a Reader for the MultiAdminVerifyRuleModify structure.
type MultiAdminVerifyRuleModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiAdminVerifyRuleModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiAdminVerifyRuleModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMultiAdminVerifyRuleModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMultiAdminVerifyRuleModifyOK creates a MultiAdminVerifyRuleModifyOK with default headers values
func NewMultiAdminVerifyRuleModifyOK() *MultiAdminVerifyRuleModifyOK {
	return &MultiAdminVerifyRuleModifyOK{}
}

/*
MultiAdminVerifyRuleModifyOK describes a response with status code 200, with default header values.

OK
*/
type MultiAdminVerifyRuleModifyOK struct {
}

// IsSuccess returns true when this multi admin verify rule modify o k response has a 2xx status code
func (o *MultiAdminVerifyRuleModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this multi admin verify rule modify o k response has a 3xx status code
func (o *MultiAdminVerifyRuleModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this multi admin verify rule modify o k response has a 4xx status code
func (o *MultiAdminVerifyRuleModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this multi admin verify rule modify o k response has a 5xx status code
func (o *MultiAdminVerifyRuleModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this multi admin verify rule modify o k response a status code equal to that given
func (o *MultiAdminVerifyRuleModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the multi admin verify rule modify o k response
func (o *MultiAdminVerifyRuleModifyOK) Code() int {
	return 200
}

func (o *MultiAdminVerifyRuleModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/multi-admin-verify/rules/{owner.uuid}/{operation}][%d] multiAdminVerifyRuleModifyOK", 200)
}

func (o *MultiAdminVerifyRuleModifyOK) String() string {
	return fmt.Sprintf("[PATCH /security/multi-admin-verify/rules/{owner.uuid}/{operation}][%d] multiAdminVerifyRuleModifyOK", 200)
}

func (o *MultiAdminVerifyRuleModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMultiAdminVerifyRuleModifyDefault creates a MultiAdminVerifyRuleModifyDefault with default headers values
func NewMultiAdminVerifyRuleModifyDefault(code int) *MultiAdminVerifyRuleModifyDefault {
	return &MultiAdminVerifyRuleModifyDefault{
		_statusCode: code,
	}
}

/*
	MultiAdminVerifyRuleModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262310 | System rules cannot be deleted or have their query modified. |
| 262311 | Value must be greater than zero. |
| 262312 | Number of required approvers must be less than the total number of unique approvers in the approval-groups. |
| 262313 | Number of unique approvers in the approval-groups must be greater than the number of required approvers. |
| 262316 | Value must be in the range one second to two weeks. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type MultiAdminVerifyRuleModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this multi admin verify rule modify default response has a 2xx status code
func (o *MultiAdminVerifyRuleModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this multi admin verify rule modify default response has a 3xx status code
func (o *MultiAdminVerifyRuleModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this multi admin verify rule modify default response has a 4xx status code
func (o *MultiAdminVerifyRuleModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this multi admin verify rule modify default response has a 5xx status code
func (o *MultiAdminVerifyRuleModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this multi admin verify rule modify default response a status code equal to that given
func (o *MultiAdminVerifyRuleModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the multi admin verify rule modify default response
func (o *MultiAdminVerifyRuleModifyDefault) Code() int {
	return o._statusCode
}

func (o *MultiAdminVerifyRuleModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/multi-admin-verify/rules/{owner.uuid}/{operation}][%d] multi_admin_verify_rule_modify default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyRuleModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/multi-admin-verify/rules/{owner.uuid}/{operation}][%d] multi_admin_verify_rule_modify default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyRuleModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MultiAdminVerifyRuleModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
