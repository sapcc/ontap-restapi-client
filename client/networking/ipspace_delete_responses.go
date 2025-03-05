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

	"github.com/sapcc/ontap-restapi/models"
)

// IpspaceDeleteReader is a Reader for the IpspaceDelete structure.
type IpspaceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpspaceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpspaceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpspaceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpspaceDeleteOK creates a IpspaceDeleteOK with default headers values
func NewIpspaceDeleteOK() *IpspaceDeleteOK {
	return &IpspaceDeleteOK{}
}

/*
IpspaceDeleteOK describes a response with status code 200, with default header values.

OK
*/
type IpspaceDeleteOK struct {
}

// IsSuccess returns true when this ipspace delete o k response has a 2xx status code
func (o *IpspaceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipspace delete o k response has a 3xx status code
func (o *IpspaceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipspace delete o k response has a 4xx status code
func (o *IpspaceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipspace delete o k response has a 5xx status code
func (o *IpspaceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipspace delete o k response a status code equal to that given
func (o *IpspaceDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipspace delete o k response
func (o *IpspaceDeleteOK) Code() int {
	return 200
}

func (o *IpspaceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /network/ipspaces/{uuid}][%d] ipspaceDeleteOK", 200)
}

func (o *IpspaceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /network/ipspaces/{uuid}][%d] ipspaceDeleteOK", 200)
}

func (o *IpspaceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpspaceDeleteDefault creates a IpspaceDeleteDefault with default headers values
func NewIpspaceDeleteDefault(code int) *IpspaceDeleteDefault {
	return &IpspaceDeleteDefault{
		_statusCode: code,
	}
}

/*
	IpspaceDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1966309 | The IPspace cannot be removed because it is being used by the audit log forwarding feature. |
| 1966333 | IPspace can't be removed because it has SVMs assigned. |
| 1969180 | Cannot delete an IPspace with an existing broadcast domain. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IpspaceDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ipspace delete default response has a 2xx status code
func (o *IpspaceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipspace delete default response has a 3xx status code
func (o *IpspaceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipspace delete default response has a 4xx status code
func (o *IpspaceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipspace delete default response has a 5xx status code
func (o *IpspaceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipspace delete default response a status code equal to that given
func (o *IpspaceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipspace delete default response
func (o *IpspaceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpspaceDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ipspaces/{uuid}][%d] ipspace_delete default %s", o._statusCode, payload)
}

func (o *IpspaceDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ipspaces/{uuid}][%d] ipspace_delete default %s", o._statusCode, payload)
}

func (o *IpspaceDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IpspaceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
