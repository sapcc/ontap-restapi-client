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

// NetworkEthernetPortDeleteReader is a Reader for the NetworkEthernetPortDelete structure.
type NetworkEthernetPortDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetPortDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetPortDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewNetworkEthernetPortDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetPortDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetPortDeleteOK creates a NetworkEthernetPortDeleteOK with default headers values
func NewNetworkEthernetPortDeleteOK() *NetworkEthernetPortDeleteOK {
	return &NetworkEthernetPortDeleteOK{}
}

/*
NetworkEthernetPortDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetPortDeleteOK struct {
}

// IsSuccess returns true when this network ethernet port delete o k response has a 2xx status code
func (o *NetworkEthernetPortDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet port delete o k response has a 3xx status code
func (o *NetworkEthernetPortDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet port delete o k response has a 4xx status code
func (o *NetworkEthernetPortDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet port delete o k response has a 5xx status code
func (o *NetworkEthernetPortDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet port delete o k response a status code equal to that given
func (o *NetworkEthernetPortDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network ethernet port delete o k response
func (o *NetworkEthernetPortDeleteOK) Code() int {
	return 200
}

func (o *NetworkEthernetPortDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /network/ethernet/ports/{uuid}][%d] networkEthernetPortDeleteOK", 200)
}

func (o *NetworkEthernetPortDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /network/ethernet/ports/{uuid}][%d] networkEthernetPortDeleteOK", 200)
}

func (o *NetworkEthernetPortDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkEthernetPortDeleteAccepted creates a NetworkEthernetPortDeleteAccepted with default headers values
func NewNetworkEthernetPortDeleteAccepted() *NetworkEthernetPortDeleteAccepted {
	return &NetworkEthernetPortDeleteAccepted{}
}

/*
NetworkEthernetPortDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type NetworkEthernetPortDeleteAccepted struct {
}

// IsSuccess returns true when this network ethernet port delete accepted response has a 2xx status code
func (o *NetworkEthernetPortDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet port delete accepted response has a 3xx status code
func (o *NetworkEthernetPortDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet port delete accepted response has a 4xx status code
func (o *NetworkEthernetPortDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet port delete accepted response has a 5xx status code
func (o *NetworkEthernetPortDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet port delete accepted response a status code equal to that given
func (o *NetworkEthernetPortDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the network ethernet port delete accepted response
func (o *NetworkEthernetPortDeleteAccepted) Code() int {
	return 202
}

func (o *NetworkEthernetPortDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /network/ethernet/ports/{uuid}][%d] networkEthernetPortDeleteAccepted", 202)
}

func (o *NetworkEthernetPortDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /network/ethernet/ports/{uuid}][%d] networkEthernetPortDeleteAccepted", 202)
}

func (o *NetworkEthernetPortDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkEthernetPortDeleteDefault creates a NetworkEthernetPortDeleteDefault with default headers values
func NewNetworkEthernetPortDeleteDefault(code int) *NetworkEthernetPortDeleteDefault {
	return &NetworkEthernetPortDeleteDefault{
		_statusCode: code,
	}
}

/*
	NetworkEthernetPortDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1376858 | Port already has an interface bound. |
| 1966189 | Port is the home port or current port of an interface. |
| 1966302 | This interface group is hosting VLAN interfaces that must be deleted before running this command. |
| 1967105 | Cannot delete a physical port. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkEthernetPortDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ethernet port delete default response has a 2xx status code
func (o *NetworkEthernetPortDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ethernet port delete default response has a 3xx status code
func (o *NetworkEthernetPortDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ethernet port delete default response has a 4xx status code
func (o *NetworkEthernetPortDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ethernet port delete default response has a 5xx status code
func (o *NetworkEthernetPortDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ethernet port delete default response a status code equal to that given
func (o *NetworkEthernetPortDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ethernet port delete default response
func (o *NetworkEthernetPortDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetPortDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ethernet/ports/{uuid}][%d] network_ethernet_port_delete default %s", o._statusCode, payload)
}

func (o *NetworkEthernetPortDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ethernet/ports/{uuid}][%d] network_ethernet_port_delete default %s", o._statusCode, payload)
}

func (o *NetworkEthernetPortDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetPortDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
