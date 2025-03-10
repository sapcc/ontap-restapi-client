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

// NetworkEthernetBroadcastDomainModifyReader is a Reader for the NetworkEthernetBroadcastDomainModify structure.
type NetworkEthernetBroadcastDomainModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetBroadcastDomainModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetBroadcastDomainModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetBroadcastDomainModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetBroadcastDomainModifyOK creates a NetworkEthernetBroadcastDomainModifyOK with default headers values
func NewNetworkEthernetBroadcastDomainModifyOK() *NetworkEthernetBroadcastDomainModifyOK {
	return &NetworkEthernetBroadcastDomainModifyOK{}
}

/*
NetworkEthernetBroadcastDomainModifyOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetBroadcastDomainModifyOK struct {
}

// IsSuccess returns true when this network ethernet broadcast domain modify o k response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet broadcast domain modify o k response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet broadcast domain modify o k response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet broadcast domain modify o k response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet broadcast domain modify o k response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network ethernet broadcast domain modify o k response
func (o *NetworkEthernetBroadcastDomainModifyOK) Code() int {
	return 200
}

func (o *NetworkEthernetBroadcastDomainModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /network/ethernet/broadcast-domains/{uuid}][%d] networkEthernetBroadcastDomainModifyOK", 200)
}

func (o *NetworkEthernetBroadcastDomainModifyOK) String() string {
	return fmt.Sprintf("[PATCH /network/ethernet/broadcast-domains/{uuid}][%d] networkEthernetBroadcastDomainModifyOK", 200)
}

func (o *NetworkEthernetBroadcastDomainModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkEthernetBroadcastDomainModifyDefault creates a NetworkEthernetBroadcastDomainModifyDefault with default headers values
func NewNetworkEthernetBroadcastDomainModifyDefault(code int) *NetworkEthernetBroadcastDomainModifyDefault {
	return &NetworkEthernetBroadcastDomainModifyDefault{
		_statusCode: code,
	}
}

/*
	NetworkEthernetBroadcastDomainModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1376484 | Cannot change the MTU of a VLAN to be greater than the MTU of the port hosting it. |
| 1377267 | The specified IPspace does not exist. |
| 1377269 | Failed to lookup the specified IPspace. |
| 1377560 | Broadcast domain already exists in specified IPspace. |
| 1377575 | Remove associated subnets before deleting this broadcast domain. |
| 1377605 | Moving the system-generated broadcast domain to another IPspace is not supported. |
| 1377609 | Updates are partially complete. Updating broadcast domain attributes on some or all of the ports in the broadcast domain have failed. |
| 1966460 | The specified MTU is either too large or too small. |
| 1967082 | The specified ipspace.name does not match the IPspace name of ipspace.uuid. |
| 1967150 | The specified ipspace.uuid is not valid. |
| 1967151 | The specified ipspace.uuid and ipspace.name do not match. |
| 1967152 | Patching IPspace for a broadcast domain requires an effective cluster version of 9.7 or later. |
| 53280884 | The MTU of the broadcast domain cannot be modified on this platform. |
| 53282013 | Broadcast domain cannot be renamed because the name is reserved by the system. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkEthernetBroadcastDomainModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ethernet broadcast domain modify default response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ethernet broadcast domain modify default response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ethernet broadcast domain modify default response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ethernet broadcast domain modify default response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ethernet broadcast domain modify default response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ethernet broadcast domain modify default response
func (o *NetworkEthernetBroadcastDomainModifyDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetBroadcastDomainModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/broadcast-domains/{uuid}][%d] network_ethernet_broadcast_domain_modify default %s", o._statusCode, payload)
}

func (o *NetworkEthernetBroadcastDomainModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/broadcast-domains/{uuid}][%d] network_ethernet_broadcast_domain_modify default %s", o._statusCode, payload)
}

func (o *NetworkEthernetBroadcastDomainModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
