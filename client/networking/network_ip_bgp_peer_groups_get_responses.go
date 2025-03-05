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

// NetworkIPBgpPeerGroupsGetReader is a Reader for the NetworkIPBgpPeerGroupsGet structure.
type NetworkIPBgpPeerGroupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPBgpPeerGroupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPBgpPeerGroupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPBgpPeerGroupsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPBgpPeerGroupsGetOK creates a NetworkIPBgpPeerGroupsGetOK with default headers values
func NewNetworkIPBgpPeerGroupsGetOK() *NetworkIPBgpPeerGroupsGetOK {
	return &NetworkIPBgpPeerGroupsGetOK{}
}

/*
NetworkIPBgpPeerGroupsGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPBgpPeerGroupsGetOK struct {
	Payload *models.BgpPeerGroupResponse
}

// IsSuccess returns true when this network Ip bgp peer groups get o k response has a 2xx status code
func (o *NetworkIPBgpPeerGroupsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network Ip bgp peer groups get o k response has a 3xx status code
func (o *NetworkIPBgpPeerGroupsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network Ip bgp peer groups get o k response has a 4xx status code
func (o *NetworkIPBgpPeerGroupsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network Ip bgp peer groups get o k response has a 5xx status code
func (o *NetworkIPBgpPeerGroupsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network Ip bgp peer groups get o k response a status code equal to that given
func (o *NetworkIPBgpPeerGroupsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network Ip bgp peer groups get o k response
func (o *NetworkIPBgpPeerGroupsGetOK) Code() int {
	return 200
}

func (o *NetworkIPBgpPeerGroupsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups][%d] networkIpBgpPeerGroupsGetOK %s", 200, payload)
}

func (o *NetworkIPBgpPeerGroupsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups][%d] networkIpBgpPeerGroupsGetOK %s", 200, payload)
}

func (o *NetworkIPBgpPeerGroupsGetOK) GetPayload() *models.BgpPeerGroupResponse {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BgpPeerGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkIPBgpPeerGroupsGetDefault creates a NetworkIPBgpPeerGroupsGetDefault with default headers values
func NewNetworkIPBgpPeerGroupsGetDefault(code int) *NetworkIPBgpPeerGroupsGetDefault {
	return &NetworkIPBgpPeerGroupsGetDefault{
		_statusCode: code,
	}
}

/*
NetworkIPBgpPeerGroupsGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkIPBgpPeerGroupsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ip bgp peer groups get default response has a 2xx status code
func (o *NetworkIPBgpPeerGroupsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ip bgp peer groups get default response has a 3xx status code
func (o *NetworkIPBgpPeerGroupsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ip bgp peer groups get default response has a 4xx status code
func (o *NetworkIPBgpPeerGroupsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ip bgp peer groups get default response has a 5xx status code
func (o *NetworkIPBgpPeerGroupsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ip bgp peer groups get default response a status code equal to that given
func (o *NetworkIPBgpPeerGroupsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ip bgp peer groups get default response
func (o *NetworkIPBgpPeerGroupsGetDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPBgpPeerGroupsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups][%d] network_ip_bgp_peer_groups_get default %s", o._statusCode, payload)
}

func (o *NetworkIPBgpPeerGroupsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ip/bgp/peer-groups][%d] network_ip_bgp_peer_groups_get default %s", o._statusCode, payload)
}

func (o *NetworkIPBgpPeerGroupsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
