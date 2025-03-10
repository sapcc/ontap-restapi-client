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

// AccountTotpDeleteReader is a Reader for the AccountTotpDelete structure.
type AccountTotpDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountTotpDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountTotpDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountTotpDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountTotpDeleteOK creates a AccountTotpDeleteOK with default headers values
func NewAccountTotpDeleteOK() *AccountTotpDeleteOK {
	return &AccountTotpDeleteOK{}
}

/*
AccountTotpDeleteOK describes a response with status code 200, with default header values.

OK
*/
type AccountTotpDeleteOK struct {
}

// IsSuccess returns true when this account totp delete o k response has a 2xx status code
func (o *AccountTotpDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account totp delete o k response has a 3xx status code
func (o *AccountTotpDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account totp delete o k response has a 4xx status code
func (o *AccountTotpDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account totp delete o k response has a 5xx status code
func (o *AccountTotpDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account totp delete o k response a status code equal to that given
func (o *AccountTotpDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account totp delete o k response
func (o *AccountTotpDeleteOK) Code() int {
	return 200
}

func (o *AccountTotpDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/login/totps/{owner.uuid}/{account.name}][%d] accountTotpDeleteOK", 200)
}

func (o *AccountTotpDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/login/totps/{owner.uuid}/{account.name}][%d] accountTotpDeleteOK", 200)
}

func (o *AccountTotpDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountTotpDeleteDefault creates a AccountTotpDeleteDefault with default headers values
func NewAccountTotpDeleteDefault(code int) *AccountTotpDeleteDefault {
	return &AccountTotpDeleteDefault{
		_statusCode: code,
	}
}

/*
AccountTotpDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type AccountTotpDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this account totp delete default response has a 2xx status code
func (o *AccountTotpDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this account totp delete default response has a 3xx status code
func (o *AccountTotpDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this account totp delete default response has a 4xx status code
func (o *AccountTotpDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this account totp delete default response has a 5xx status code
func (o *AccountTotpDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this account totp delete default response a status code equal to that given
func (o *AccountTotpDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the account totp delete default response
func (o *AccountTotpDeleteDefault) Code() int {
	return o._statusCode
}

func (o *AccountTotpDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/login/totps/{owner.uuid}/{account.name}][%d] account_totp_delete default %s", o._statusCode, payload)
}

func (o *AccountTotpDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/login/totps/{owner.uuid}/{account.name}][%d] account_totp_delete default %s", o._statusCode, payload)
}

func (o *AccountTotpDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AccountTotpDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
