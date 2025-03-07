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

// NetworkIPBgpPeerGroupsCreateReader is a Reader for the NetworkIPBgpPeerGroupsCreate structure.
type NetworkIPBgpPeerGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPBgpPeerGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNetworkIPBgpPeerGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPBgpPeerGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPBgpPeerGroupsCreateCreated creates a NetworkIPBgpPeerGroupsCreateCreated with default headers values
func NewNetworkIPBgpPeerGroupsCreateCreated() *NetworkIPBgpPeerGroupsCreateCreated {
	return &NetworkIPBgpPeerGroupsCreateCreated{}
}

/*
NetworkIPBgpPeerGroupsCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NetworkIPBgpPeerGroupsCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this network Ip bgp peer groups create created response has a 2xx status code
func (o *NetworkIPBgpPeerGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network Ip bgp peer groups create created response has a 3xx status code
func (o *NetworkIPBgpPeerGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network Ip bgp peer groups create created response has a 4xx status code
func (o *NetworkIPBgpPeerGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this network Ip bgp peer groups create created response has a 5xx status code
func (o *NetworkIPBgpPeerGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this network Ip bgp peer groups create created response a status code equal to that given
func (o *NetworkIPBgpPeerGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the network Ip bgp peer groups create created response
func (o *NetworkIPBgpPeerGroupsCreateCreated) Code() int {
	return 201
}

func (o *NetworkIPBgpPeerGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /network/ip/bgp/peer-groups][%d] networkIpBgpPeerGroupsCreateCreated", 201)
}

func (o *NetworkIPBgpPeerGroupsCreateCreated) String() string {
	return fmt.Sprintf("[POST /network/ip/bgp/peer-groups][%d] networkIpBgpPeerGroupsCreateCreated", 201)
}

func (o *NetworkIPBgpPeerGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewNetworkIPBgpPeerGroupsCreateDefault creates a NetworkIPBgpPeerGroupsCreateDefault with default headers values
func NewNetworkIPBgpPeerGroupsCreateDefault(code int) *NetworkIPBgpPeerGroupsCreateDefault {
	return &NetworkIPBgpPeerGroupsCreateDefault{
		_statusCode: code,
	}
}

/*
	NetworkIPBgpPeerGroupsCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1376963 | Duplicate IP address is specified. |
| 1966133 | Since masking an address with a netmask represents an entire IP subnet, the masked and unmasked IP addresses cannot be the same. |
| 1966267 | IPv6 addresses must have a prefix length of 64. |
| 1966269 | IPv4 addresses must have a netmask length between 1 and 32. |
| 1967082 | IPspace name and UUID must match if both are given. |
| 1967155 | The specified local.port.name does not match the location.port.name for the specified local.interface. |
| 1967156 | The specified local.port.node.name does not match the location.port.node.name for the specified local.interface. |
| 1967157 | The specified local.port.uuid does not match the location.port.uuid for the specified local.interface. |
| 1967158 | The specified local.interface.name does not exist in the associated IPspace. local.ip.address and local.ip.netmask are required to create a new LIF. |
| 1967159 | local.interface does not support management-bgp service. |
| 1967160 | The specified local.interface.name does not match the specified interface name of local.interface.uuid. |
| 1967161 | The specified local.interface.uuid does not exist in the specified IPspace. |
| 1967162 | Either local.interface or local.ip and local.port are required to specify a local LIF. |
| 1967163 | The specified local.port.name does not match the specified port name of local.port.uuid. |
| 1967164 | The specified local.port.node.name does not match the specified node name of local.port.uuid. |
| 1967165 | The specified local.port does not exist. |
| 1967166 | ipspace.uuid or ipspace.name must be provided with local.interface.name together to identify a LIF. |
| 1967167 | Internal error. Failed to update BGP configuration for node. Retry the command, if necessary. |
| 1967168 | Internal error. Failed to create a VIP port for IPspace on node. Retry the command, if necessary. |
| 1967169 | Internal error. BGP configuration changed during the operation. Retry the command, if necessary. |
| 1967170 | Internal error. VIP port configuration changed during the operation. Retry the command, if necessary. |
| 1967171 | Internal error. Fail to access or update BGP peer group. Retry the command, if necessary. |
| 1967172 | Peer group could not be updated because IPspace does not exist. Retry the command, if necessary. |
| 1967173 | The specified local.ip.address does not match the address for the specified local.interface. |
| 1967174 | The specified local.ip.netmask does not match the netmask for the specified local.interface. |
| 1967176 | The specified local.interface.name does not exist in the associated IPspace. local.port.name, local.port.node.name, or local.port.uuid is required to create a new LIF. |
| 1967177 | Internal error. Failed to access the local interface. Retry the command, if necessary. |
| 1967178 | The IPv6 address specified with local.ip.address is not supported because it is link-local, multicast, v4-compatible, v4-mapped, loopback or "::". |
| 1967179 | The IPv4 address specified with local.ip.address is not supported because it is multicast, loopback or 0.0.0.0. |
| 1967187 | Configuring 4 bytes peer.asn requires an effective cluster version of 9.9.1 or later. |
| 1967188 | Configuring peer address as a next hop requires an effective cluster version of 9.9.1 or later. |
| 1967189 | The parameter peer.asn can't be zero. |
| 1967192 | Configuring peer.md5_enabled requires an effective cluster version of 9.16.1 or later. |
| 1967193 | Configuring peer.md5_enabled requires the peer.md5_secret parameter. |
| 1967194 | Configuring peer.md5_secret requires that peer.md5_enabled is set to true. |
| 53281985 | Internal error. Failed to update BGP peer group because BGP LIF moved during the operation. Wait a few minutes and try the command again. |
| 53282006 | BGP peer group could not be updated to use a peer address because the value provided is not a valid peer address. If necessary, try the command again with a routable host address. |
| 53282007 | BGP peer group could not be updated to use a peer address because the address represents a different address family to the address of the associated BGP LIF. If necessary, try the command again with a matching address family. |
| 53282018 | Failed to create BGP peer group because an existing peer group has already established a BGP session between LIF and peer address. If necessary, try the command again with a different BGP LIF or a different peer address. |
| 53282021 | IPsec must be enabled to use TCP MD5 in BGP sessions with a BGP peer-group. |
| 53282024 | peer.md5_secret should be 1-80 bytes. |
| 53282025 | peer.md5_secret should have an even length. |
| 53282026 | The hex format of peer.md5_secret should have only hex string starting with '0x'. |
| 53282027 | peer.md5_secret should be 1-79 characters. |
| 53282028 | peer.md5_secret supports only printable ASCII characters such as numbers, alphabets, or special characters. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkIPBgpPeerGroupsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ip bgp peer groups create default response has a 2xx status code
func (o *NetworkIPBgpPeerGroupsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ip bgp peer groups create default response has a 3xx status code
func (o *NetworkIPBgpPeerGroupsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ip bgp peer groups create default response has a 4xx status code
func (o *NetworkIPBgpPeerGroupsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ip bgp peer groups create default response has a 5xx status code
func (o *NetworkIPBgpPeerGroupsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ip bgp peer groups create default response a status code equal to that given
func (o *NetworkIPBgpPeerGroupsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ip bgp peer groups create default response
func (o *NetworkIPBgpPeerGroupsCreateDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPBgpPeerGroupsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ip/bgp/peer-groups][%d] network_ip_bgp_peer_groups_create default %s", o._statusCode, payload)
}

func (o *NetworkIPBgpPeerGroupsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ip/bgp/peer-groups][%d] network_ip_bgp_peer_groups_create default %s", o._statusCode, payload)
}

func (o *NetworkIPBgpPeerGroupsCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
