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

// IPSubnetGetReader is a Reader for the IPSubnetGet structure.
type IPSubnetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPSubnetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPSubnetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIPSubnetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIPSubnetGetOK creates a IPSubnetGetOK with default headers values
func NewIPSubnetGetOK() *IPSubnetGetOK {
	return &IPSubnetGetOK{}
}

/*
IPSubnetGetOK describes a response with status code 200, with default header values.

OK
*/
type IPSubnetGetOK struct {
	Payload *models.IPSubnet
}

// IsSuccess returns true when this ip subnet get o k response has a 2xx status code
func (o *IPSubnetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ip subnet get o k response has a 3xx status code
func (o *IPSubnetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ip subnet get o k response has a 4xx status code
func (o *IPSubnetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ip subnet get o k response has a 5xx status code
func (o *IPSubnetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ip subnet get o k response a status code equal to that given
func (o *IPSubnetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ip subnet get o k response
func (o *IPSubnetGetOK) Code() int {
	return 200
}

func (o *IPSubnetGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/subnets/{uuid}][%d] ipSubnetGetOK %s", 200, payload)
}

func (o *IPSubnetGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/subnets/{uuid}][%d] ipSubnetGetOK %s", 200, payload)
}

func (o *IPSubnetGetOK) GetPayload() *models.IPSubnet {
	return o.Payload
}

func (o *IPSubnetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPSubnet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIPSubnetGetDefault creates a IPSubnetGetDefault with default headers values
func NewIPSubnetGetDefault(code int) *IPSubnetGetDefault {
	return &IPSubnetGetDefault{
		_statusCode: code,
	}
}

/*
IPSubnetGetDefault describes a response with status code -1, with default header values.

Error
*/
type IPSubnetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ip subnet get default response has a 2xx status code
func (o *IPSubnetGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ip subnet get default response has a 3xx status code
func (o *IPSubnetGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ip subnet get default response has a 4xx status code
func (o *IPSubnetGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ip subnet get default response has a 5xx status code
func (o *IPSubnetGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ip subnet get default response a status code equal to that given
func (o *IPSubnetGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ip subnet get default response
func (o *IPSubnetGetDefault) Code() int {
	return o._statusCode
}

func (o *IPSubnetGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/subnets/{uuid}][%d] ip_subnet_get default %s", o._statusCode, payload)
}

func (o *IPSubnetGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/subnets/{uuid}][%d] ip_subnet_get default %s", o._statusCode, payload)
}

func (o *IPSubnetGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IPSubnetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
