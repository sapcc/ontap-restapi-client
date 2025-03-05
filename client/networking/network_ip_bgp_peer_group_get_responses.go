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

// NetworkIPBgpPeerGroupGetReader is a Reader for the NetworkIPBgpPeerGroupGet structure.
type NetworkIPBgpPeerGroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPBgpPeerGroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPBgpPeerGroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPBgpPeerGroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPBgpPeerGroupGetOK creates a NetworkIPBgpPeerGroupGetOK with default headers values
func NewNetworkIPBgpPeerGroupGetOK() *NetworkIPBgpPeerGroupGetOK {
	return &NetworkIPBgpPeerGroupGetOK{}
}

/*
NetworkIPBgpPeerGroupGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPBgpPeerGroupGetOK struct {
	Payload *models.BgpPeerGroup
}

// IsSuccess returns true when this network Ip bgp peer group get o k response has a 2xx status code
func (o *NetworkIPBgpPeerGroupGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network Ip bgp peer group get o k response has a 3xx status code
func (o *NetworkIPBgpPeerGroupGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network Ip bgp peer group get o k response has a 4xx status code
func (o *NetworkIPBgpPeerGroupGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network Ip bgp peer group get o k response has a 5xx status code
func (o *NetworkIPBgpPeerGroupGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network Ip bgp peer group get o k response a status code equal to that given
func (o *NetworkIPBgpPeerGroupGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network Ip bgp peer group get o k response
func (o *NetworkIPBgpPeerGroupGetOK) Code() int {
	return 200
}

func (o *NetworkIPBgpPeerGroupGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups/{uuid}][%d] networkIpBgpPeerGroupGetOK %s", 200, payload)
}

func (o *NetworkIPBgpPeerGroupGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups/{uuid}][%d] networkIpBgpPeerGroupGetOK %s", 200, payload)
}

func (o *NetworkIPBgpPeerGroupGetOK) GetPayload() *models.BgpPeerGroup {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BgpPeerGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkIPBgpPeerGroupGetDefault creates a NetworkIPBgpPeerGroupGetDefault with default headers values
func NewNetworkIPBgpPeerGroupGetDefault(code int) *NetworkIPBgpPeerGroupGetDefault {
	return &NetworkIPBgpPeerGroupGetDefault{
		_statusCode: code,
	}
}

/*
NetworkIPBgpPeerGroupGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkIPBgpPeerGroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ip bgp peer group get default response has a 2xx status code
func (o *NetworkIPBgpPeerGroupGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ip bgp peer group get default response has a 3xx status code
func (o *NetworkIPBgpPeerGroupGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ip bgp peer group get default response has a 4xx status code
func (o *NetworkIPBgpPeerGroupGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ip bgp peer group get default response has a 5xx status code
func (o *NetworkIPBgpPeerGroupGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ip bgp peer group get default response a status code equal to that given
func (o *NetworkIPBgpPeerGroupGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ip bgp peer group get default response
func (o *NetworkIPBgpPeerGroupGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPBgpPeerGroupGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups/{uuid}][%d] network_ip_bgp_peer_group_get default %s", o._statusCode, payload)
}

func (o *NetworkIPBgpPeerGroupGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups/{uuid}][%d] network_ip_bgp_peer_group_get default %s", o._statusCode, payload)
}

func (o *NetworkIPBgpPeerGroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
