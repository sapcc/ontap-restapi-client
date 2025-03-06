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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NetworkIPRoutesGetReader is a Reader for the NetworkIPRoutesGet structure.
type NetworkIPRoutesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPRoutesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPRoutesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPRoutesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPRoutesGetOK creates a NetworkIPRoutesGetOK with default headers values
func NewNetworkIPRoutesGetOK() *NetworkIPRoutesGetOK {
	return &NetworkIPRoutesGetOK{}
}

/*
NetworkIPRoutesGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPRoutesGetOK struct {
	Payload *models.NetworkRouteResponse
}

// IsSuccess returns true when this network Ip routes get o k response has a 2xx status code
func (o *NetworkIPRoutesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network Ip routes get o k response has a 3xx status code
func (o *NetworkIPRoutesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network Ip routes get o k response has a 4xx status code
func (o *NetworkIPRoutesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network Ip routes get o k response has a 5xx status code
func (o *NetworkIPRoutesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network Ip routes get o k response a status code equal to that given
func (o *NetworkIPRoutesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network Ip routes get o k response
func (o *NetworkIPRoutesGetOK) Code() int {
	return 200
}

func (o *NetworkIPRoutesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/routes][%d] networkIpRoutesGetOK %s", 200, payload)
}

func (o *NetworkIPRoutesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/routes][%d] networkIpRoutesGetOK %s", 200, payload)
}

func (o *NetworkIPRoutesGetOK) GetPayload() *models.NetworkRouteResponse {
	return o.Payload
}

func (o *NetworkIPRoutesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkRouteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkIPRoutesGetDefault creates a NetworkIPRoutesGetDefault with default headers values
func NewNetworkIPRoutesGetDefault(code int) *NetworkIPRoutesGetDefault {
	return &NetworkIPRoutesGetDefault{
		_statusCode: code,
	}
}

/*
NetworkIPRoutesGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkIPRoutesGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ip routes get default response has a 2xx status code
func (o *NetworkIPRoutesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ip routes get default response has a 3xx status code
func (o *NetworkIPRoutesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ip routes get default response has a 4xx status code
func (o *NetworkIPRoutesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ip routes get default response has a 5xx status code
func (o *NetworkIPRoutesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ip routes get default response a status code equal to that given
func (o *NetworkIPRoutesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ip routes get default response
func (o *NetworkIPRoutesGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPRoutesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/routes][%d] network_ip_routes_get default %s", o._statusCode, payload)
}

func (o *NetworkIPRoutesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/routes][%d] network_ip_routes_get default %s", o._statusCode, payload)
}

func (o *NetworkIPRoutesGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPRoutesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
