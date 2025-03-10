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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// MultiAdminVerifyApprovalGroupDeleteReader is a Reader for the MultiAdminVerifyApprovalGroupDelete structure.
type MultiAdminVerifyApprovalGroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiAdminVerifyApprovalGroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiAdminVerifyApprovalGroupDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMultiAdminVerifyApprovalGroupDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMultiAdminVerifyApprovalGroupDeleteOK creates a MultiAdminVerifyApprovalGroupDeleteOK with default headers values
func NewMultiAdminVerifyApprovalGroupDeleteOK() *MultiAdminVerifyApprovalGroupDeleteOK {
	return &MultiAdminVerifyApprovalGroupDeleteOK{}
}

/*
MultiAdminVerifyApprovalGroupDeleteOK describes a response with status code 200, with default header values.

OK
*/
type MultiAdminVerifyApprovalGroupDeleteOK struct {
}

// IsSuccess returns true when this multi admin verify approval group delete o k response has a 2xx status code
func (o *MultiAdminVerifyApprovalGroupDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this multi admin verify approval group delete o k response has a 3xx status code
func (o *MultiAdminVerifyApprovalGroupDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this multi admin verify approval group delete o k response has a 4xx status code
func (o *MultiAdminVerifyApprovalGroupDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this multi admin verify approval group delete o k response has a 5xx status code
func (o *MultiAdminVerifyApprovalGroupDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this multi admin verify approval group delete o k response a status code equal to that given
func (o *MultiAdminVerifyApprovalGroupDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the multi admin verify approval group delete o k response
func (o *MultiAdminVerifyApprovalGroupDeleteOK) Code() int {
	return 200
}

func (o *MultiAdminVerifyApprovalGroupDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/multi-admin-verify/approval-groups/{owner.uuid}/{name}][%d] multiAdminVerifyApprovalGroupDeleteOK", 200)
}

func (o *MultiAdminVerifyApprovalGroupDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/multi-admin-verify/approval-groups/{owner.uuid}/{name}][%d] multiAdminVerifyApprovalGroupDeleteOK", 200)
}

func (o *MultiAdminVerifyApprovalGroupDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMultiAdminVerifyApprovalGroupDeleteDefault creates a MultiAdminVerifyApprovalGroupDeleteDefault with default headers values
func NewMultiAdminVerifyApprovalGroupDeleteDefault(code int) *MultiAdminVerifyApprovalGroupDeleteDefault {
	return &MultiAdminVerifyApprovalGroupDeleteDefault{
		_statusCode: code,
	}
}

/*
	MultiAdminVerifyApprovalGroupDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262319 | Cannot delete an approval group that is in the list of global approval groups. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type MultiAdminVerifyApprovalGroupDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this multi admin verify approval group delete default response has a 2xx status code
func (o *MultiAdminVerifyApprovalGroupDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this multi admin verify approval group delete default response has a 3xx status code
func (o *MultiAdminVerifyApprovalGroupDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this multi admin verify approval group delete default response has a 4xx status code
func (o *MultiAdminVerifyApprovalGroupDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this multi admin verify approval group delete default response has a 5xx status code
func (o *MultiAdminVerifyApprovalGroupDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this multi admin verify approval group delete default response a status code equal to that given
func (o *MultiAdminVerifyApprovalGroupDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the multi admin verify approval group delete default response
func (o *MultiAdminVerifyApprovalGroupDeleteDefault) Code() int {
	return o._statusCode
}

func (o *MultiAdminVerifyApprovalGroupDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/multi-admin-verify/approval-groups/{owner.uuid}/{name}][%d] multi_admin_verify_approval_group_delete default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyApprovalGroupDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/multi-admin-verify/approval-groups/{owner.uuid}/{name}][%d] multi_admin_verify_approval_group_delete default %s", o._statusCode, payload)
}

func (o *MultiAdminVerifyApprovalGroupDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MultiAdminVerifyApprovalGroupDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
