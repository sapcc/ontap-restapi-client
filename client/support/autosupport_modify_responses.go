// Code generated by go-swagger; DO NOT EDIT.

package support

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

// AutosupportModifyReader is a Reader for the AutosupportModify structure.
type AutosupportModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutosupportModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutosupportModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutosupportModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutosupportModifyOK creates a AutosupportModifyOK with default headers values
func NewAutosupportModifyOK() *AutosupportModifyOK {
	return &AutosupportModifyOK{}
}

/*
AutosupportModifyOK describes a response with status code 200, with default header values.

OK
*/
type AutosupportModifyOK struct {
}

// IsSuccess returns true when this autosupport modify o k response has a 2xx status code
func (o *AutosupportModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this autosupport modify o k response has a 3xx status code
func (o *AutosupportModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autosupport modify o k response has a 4xx status code
func (o *AutosupportModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this autosupport modify o k response has a 5xx status code
func (o *AutosupportModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this autosupport modify o k response a status code equal to that given
func (o *AutosupportModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the autosupport modify o k response
func (o *AutosupportModifyOK) Code() int {
	return 200
}

func (o *AutosupportModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /support/autosupport][%d] autosupportModifyOK", 200)
}

func (o *AutosupportModifyOK) String() string {
	return fmt.Sprintf("[PATCH /support/autosupport][%d] autosupportModifyOK", 200)
}

func (o *AutosupportModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAutosupportModifyDefault creates a AutosupportModifyDefault with default headers values
func NewAutosupportModifyDefault(code int) *AutosupportModifyDefault {
	return &AutosupportModifyDefault{
		_statusCode: code,
	}
}

/*
	AutosupportModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8650862 | The SMTP mail host provided cannot be empty |
| 8650863 | A maximum of 5 SMTP mail hosts can be provided |
| 8650864 | A maximum of 5 email addresses can be provided |
| 8650865 | A maximum of 5 partner email addresses can be provided |
| 8650880 | Automatic update requires AutoSupport to be enabled |
| 8650881 | Automatic update requires AutoSupport to use the HTTPS transport |
| 8650882 | Automatic update requires AutoSupport OnDemand to be enabled |
| 8650886 | The provided parameter requires an effective cluster version of ONTAP 9.15.1 or later |
| 8650887 | Username or password is not allowed in the AutoSupport url and put-url fields |
| 8650889 | The provided parameter requires an effective cluster version of ONTAP 9.16.1 or later |
| 53149727 | The proxy URI provided is invalid |
| 53149728 | The SMTP mail host URI provided is invalid |
| 53149732 | The proxy URI provided is invalid. IPv6 addresses must be enclosed within square brackets |
| 53149737 | The proxy URI provided specifies an unsupported scheme |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AutosupportModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this autosupport modify default response has a 2xx status code
func (o *AutosupportModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this autosupport modify default response has a 3xx status code
func (o *AutosupportModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this autosupport modify default response has a 4xx status code
func (o *AutosupportModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this autosupport modify default response has a 5xx status code
func (o *AutosupportModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this autosupport modify default response a status code equal to that given
func (o *AutosupportModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the autosupport modify default response
func (o *AutosupportModifyDefault) Code() int {
	return o._statusCode
}

func (o *AutosupportModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /support/autosupport][%d] autosupport_modify default %s", o._statusCode, payload)
}

func (o *AutosupportModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /support/autosupport][%d] autosupport_modify default %s", o._statusCode, payload)
}

func (o *AutosupportModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutosupportModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
