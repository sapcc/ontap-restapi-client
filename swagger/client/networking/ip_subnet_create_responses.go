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

// IPSubnetCreateReader is a Reader for the IPSubnetCreate structure.
type IPSubnetCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPSubnetCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIPSubnetCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIPSubnetCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIPSubnetCreateCreated creates a IPSubnetCreateCreated with default headers values
func NewIPSubnetCreateCreated() *IPSubnetCreateCreated {
	return &IPSubnetCreateCreated{}
}

/*
IPSubnetCreateCreated describes a response with status code 201, with default header values.

Created
*/
type IPSubnetCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.IPSubnetResponse
}

// IsSuccess returns true when this ip subnet create created response has a 2xx status code
func (o *IPSubnetCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ip subnet create created response has a 3xx status code
func (o *IPSubnetCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ip subnet create created response has a 4xx status code
func (o *IPSubnetCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ip subnet create created response has a 5xx status code
func (o *IPSubnetCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ip subnet create created response a status code equal to that given
func (o *IPSubnetCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ip subnet create created response
func (o *IPSubnetCreateCreated) Code() int {
	return 201
}

func (o *IPSubnetCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ip/subnets][%d] ipSubnetCreateCreated %s", 201, payload)
}

func (o *IPSubnetCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ip/subnets][%d] ipSubnetCreateCreated %s", 201, payload)
}

func (o *IPSubnetCreateCreated) GetPayload() *models.IPSubnetResponse {
	return o.Payload
}

func (o *IPSubnetCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.IPSubnetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIPSubnetCreateDefault creates a IPSubnetCreateDefault with default headers values
func NewIPSubnetCreateDefault(code int) *IPSubnetCreateDefault {
	return &IPSubnetCreateDefault{
		_statusCode: code,
	}
}

/*
	IPSubnetCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1377658 | Invalid gateway for subnet in IPspace. |
| 1377659 | Subnet would overlap with existing subnet named in IPspace. |
| 1377660 | A subnet with the name already exists in the IPspace. |
| 1377661 | Subnet in IPspace cannot use subnet address because that address is already used by another subnet in the same IPspace. |
| 1377662 | The IP range address is not within the subnet in IPspace. |
| 1377663 | The specified IP address range of subnet in IPspace contains an address already in use by a LIF. |
| 1377664 | The specified IP address range of subnet in IPspace contains an address already in use by the Service Processor. |
| 1377673 | The addresses provided must have the same address family. |
| 1377675 | The netmask of the interface did not match the netmask of the subnet. |
| 1377681 | Cannot update LIF associations for LIF. The broadcast domain of the LIF does not match the broadcast domain of the subnet. |
| 1377682 | IPv6 is not enabled in the cluster. |
| 1966269 | IPv4 Addresses must have a prefix length between 1 and 32. |
| 1967082 | The specified ipspace.name does not match the IPspace name of specified ipspace.uuid |
| 53282568 | The subnet.address must be specified together with subnet.netmask. |
| 53282569 | The specified subnet.netmask is not valid. |
| 53282570 | Each pair of ranges must have ip_ranges.start less than or equal to ip_ranges.end. |
| 53282571 | The ip_ranges.start and ip_ranges.end fields must have the same number of items. |
| 53282573 | Broadcast domain is a required parameter. The broadcast_domain.name and/or the broadcast_domain.uuid must be specified. |
| 53282574 | The specified broadcast_domain and ipspace parameters do not match. |
| 53282575 | Operation might have left configuration in an inconsistent state. Unable to set UUID for created entry. |
| 53282576 | The specified ipspace.uuid is invalid. |
| 53282577 | The specified broadcast_domain.uuid is invalid. |
| 53282578 | The specified broadcast_domain.name does not match the IPspace name of specified broadcast_domain.uuid |
| 53282579 | Missing the ipspace.name or ipspace.uuid parameter. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IPSubnetCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ip subnet create default response has a 2xx status code
func (o *IPSubnetCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ip subnet create default response has a 3xx status code
func (o *IPSubnetCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ip subnet create default response has a 4xx status code
func (o *IPSubnetCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ip subnet create default response has a 5xx status code
func (o *IPSubnetCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ip subnet create default response a status code equal to that given
func (o *IPSubnetCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ip subnet create default response
func (o *IPSubnetCreateDefault) Code() int {
	return o._statusCode
}

func (o *IPSubnetCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ip/subnets][%d] ip_subnet_create default %s", o._statusCode, payload)
}

func (o *IPSubnetCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ip/subnets][%d] ip_subnet_create default %s", o._statusCode, payload)
}

func (o *IPSubnetCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IPSubnetCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
