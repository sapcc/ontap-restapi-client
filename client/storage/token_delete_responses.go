// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// TokenDeleteReader is a Reader for the TokenDelete structure.
type TokenDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokenDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTokenDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTokenDeleteOK creates a TokenDeleteOK with default headers values
func NewTokenDeleteOK() *TokenDeleteOK {
	return &TokenDeleteOK{}
}

/*
TokenDeleteOK describes a response with status code 200, with default header values.

OK
*/
type TokenDeleteOK struct {
}

// IsSuccess returns true when this token delete o k response has a 2xx status code
func (o *TokenDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this token delete o k response has a 3xx status code
func (o *TokenDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this token delete o k response has a 4xx status code
func (o *TokenDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this token delete o k response has a 5xx status code
func (o *TokenDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this token delete o k response a status code equal to that given
func (o *TokenDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the token delete o k response
func (o *TokenDeleteOK) Code() int {
	return 200
}

func (o *TokenDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/file/clone/tokens/{node.uuid}/{uuid}][%d] tokenDeleteOK", 200)
}

func (o *TokenDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /storage/file/clone/tokens/{node.uuid}/{uuid}][%d] tokenDeleteOK", 200)
}

func (o *TokenDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTokenDeleteDefault creates a TokenDeleteDefault with default headers values
func NewTokenDeleteDefault(code int) *TokenDeleteDefault {
	return &TokenDeleteDefault{
		_statusCode: code,
	}
}

/*
	TokenDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 13565958 | Failed to get information about token `uuid` for node `node.name`. |
| 13565961 | Failed to delete token for node `node.name`. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type TokenDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this token delete default response has a 2xx status code
func (o *TokenDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this token delete default response has a 3xx status code
func (o *TokenDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this token delete default response has a 4xx status code
func (o *TokenDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this token delete default response has a 5xx status code
func (o *TokenDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this token delete default response a status code equal to that given
func (o *TokenDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the token delete default response
func (o *TokenDeleteDefault) Code() int {
	return o._statusCode
}

func (o *TokenDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/file/clone/tokens/{node.uuid}/{uuid}][%d] token_delete default %s", o._statusCode, payload)
}

func (o *TokenDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/file/clone/tokens/{node.uuid}/{uuid}][%d] token_delete default %s", o._statusCode, payload)
}

func (o *TokenDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TokenDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
