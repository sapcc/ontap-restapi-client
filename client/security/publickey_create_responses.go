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

	"github.com/sapcc/ontap-restapi/models"
)

// PublickeyCreateReader is a Reader for the PublickeyCreate structure.
type PublickeyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublickeyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPublickeyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublickeyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublickeyCreateCreated creates a PublickeyCreateCreated with default headers values
func NewPublickeyCreateCreated() *PublickeyCreateCreated {
	return &PublickeyCreateCreated{}
}

/*
PublickeyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type PublickeyCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this publickey create created response has a 2xx status code
func (o *PublickeyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this publickey create created response has a 3xx status code
func (o *PublickeyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this publickey create created response has a 4xx status code
func (o *PublickeyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this publickey create created response has a 5xx status code
func (o *PublickeyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this publickey create created response a status code equal to that given
func (o *PublickeyCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the publickey create created response
func (o *PublickeyCreateCreated) Code() int {
	return 201
}

func (o *PublickeyCreateCreated) Error() string {
	return fmt.Sprintf("[POST /security/authentication/publickeys][%d] publickeyCreateCreated", 201)
}

func (o *PublickeyCreateCreated) String() string {
	return fmt.Sprintf("[POST /security/authentication/publickeys][%d] publickeyCreateCreated", 201)
}

func (o *PublickeyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewPublickeyCreateDefault creates a PublickeyCreateDefault with default headers values
func NewPublickeyCreateDefault(code int) *PublickeyCreateDefault {
	return &PublickeyCreateDefault{
		_statusCode: code,
	}
}

/*
	PublickeyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5832705 | Public key already exists for the given user and application. |
| 5832707 | Failed to generate fingerprint for the public key. |
| 5832722 | The public key cannot be associated with this user on the SVM because a login method using the given application and authentication method does not exist for this user. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type PublickeyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this publickey create default response has a 2xx status code
func (o *PublickeyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this publickey create default response has a 3xx status code
func (o *PublickeyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this publickey create default response has a 4xx status code
func (o *PublickeyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this publickey create default response has a 5xx status code
func (o *PublickeyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this publickey create default response a status code equal to that given
func (o *PublickeyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the publickey create default response
func (o *PublickeyCreateDefault) Code() int {
	return o._statusCode
}

func (o *PublickeyCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/publickeys][%d] publickey_create default %s", o._statusCode, payload)
}

func (o *PublickeyCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/authentication/publickeys][%d] publickey_create default %s", o._statusCode, payload)
}

func (o *PublickeyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PublickeyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
