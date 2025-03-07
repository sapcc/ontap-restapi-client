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

// AccountPasswordCreateReader is a Reader for the AccountPasswordCreate structure.
type AccountPasswordCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountPasswordCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAccountPasswordCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountPasswordCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountPasswordCreateCreated creates a AccountPasswordCreateCreated with default headers values
func NewAccountPasswordCreateCreated() *AccountPasswordCreateCreated {
	return &AccountPasswordCreateCreated{}
}

/*
AccountPasswordCreateCreated describes a response with status code 201, with default header values.

Created
*/
type AccountPasswordCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this account password create created response has a 2xx status code
func (o *AccountPasswordCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account password create created response has a 3xx status code
func (o *AccountPasswordCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account password create created response has a 4xx status code
func (o *AccountPasswordCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this account password create created response has a 5xx status code
func (o *AccountPasswordCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this account password create created response a status code equal to that given
func (o *AccountPasswordCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the account password create created response
func (o *AccountPasswordCreateCreated) Code() int {
	return 201
}

func (o *AccountPasswordCreateCreated) Error() string {
	return fmt.Sprintf("[POST /security/authentication/password][%d] accountPasswordCreateCreated", 201)
}

func (o *AccountPasswordCreateCreated) String() string {
	return fmt.Sprintf("[POST /security/authentication/password][%d] accountPasswordCreateCreated", 201)
}

func (o *AccountPasswordCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewAccountPasswordCreateDefault creates a AccountPasswordCreateDefault with default headers values
func NewAccountPasswordCreateDefault(code int) *AccountPasswordCreateDefault {
	return &AccountPasswordCreateDefault{
		_statusCode: code,
	}
}

/*
	AccountPasswordCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 7077918 | The password cannot contain the username. |
| 7077919 | The minimum length for new password does not meet the policy. |
| 7077920 | The new password must have both letters and numbers. |
| 7077921 | The minimum number of special characters required do not meet the policy. |
| 7077924 | The new password must be different than last N passwords. |
| 7077925 | The new password must be different to the old password. |
| 7077940 | The password exceeds maximum supported length. |
| 7077941 | Defined password composition exceeds the maximum password length of 128 characters. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AccountPasswordCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this account password create default response has a 2xx status code
func (o *AccountPasswordCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this account password create default response has a 3xx status code
func (o *AccountPasswordCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this account password create default response has a 4xx status code
func (o *AccountPasswordCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this account password create default response has a 5xx status code
func (o *AccountPasswordCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this account password create default response a status code equal to that given
func (o *AccountPasswordCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the account password create default response
func (o *AccountPasswordCreateDefault) Code() int {
	return o._statusCode
}

func (o *AccountPasswordCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/password][%d] account_password_create default %s", o._statusCode, payload)
}

func (o *AccountPasswordCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/password][%d] account_password_create default %s", o._statusCode, payload)
}

func (o *AccountPasswordCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AccountPasswordCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
