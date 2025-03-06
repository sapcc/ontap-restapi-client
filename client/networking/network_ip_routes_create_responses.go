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

// NetworkIPRoutesCreateReader is a Reader for the NetworkIPRoutesCreate structure.
type NetworkIPRoutesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPRoutesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNetworkIPRoutesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPRoutesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPRoutesCreateCreated creates a NetworkIPRoutesCreateCreated with default headers values
func NewNetworkIPRoutesCreateCreated() *NetworkIPRoutesCreateCreated {
	return &NetworkIPRoutesCreateCreated{}
}

/*
NetworkIPRoutesCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NetworkIPRoutesCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.NetworkRouteResponse
}

// IsSuccess returns true when this network Ip routes create created response has a 2xx status code
func (o *NetworkIPRoutesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network Ip routes create created response has a 3xx status code
func (o *NetworkIPRoutesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network Ip routes create created response has a 4xx status code
func (o *NetworkIPRoutesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this network Ip routes create created response has a 5xx status code
func (o *NetworkIPRoutesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this network Ip routes create created response a status code equal to that given
func (o *NetworkIPRoutesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the network Ip routes create created response
func (o *NetworkIPRoutesCreateCreated) Code() int {
	return 201
}

func (o *NetworkIPRoutesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ip/routes][%d] networkIpRoutesCreateCreated %s", 201, payload)
}

func (o *NetworkIPRoutesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ip/routes][%d] networkIpRoutesCreateCreated %s", 201, payload)
}

func (o *NetworkIPRoutesCreateCreated) GetPayload() *models.NetworkRouteResponse {
	return o.Payload
}

func (o *NetworkIPRoutesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.NetworkRouteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkIPRoutesCreateDefault creates a NetworkIPRoutesCreateDefault with default headers values
func NewNetworkIPRoutesCreateDefault(code int) *NetworkIPRoutesCreateDefault {
	return &NetworkIPRoutesCreateDefault{
		_statusCode: code,
	}
}

/*
	NetworkIPRoutesCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1966265 | The destination and gateway must belong to the same IP address family. |
| 1966345 | Duplicate route exists. |
| 1966505 | The gateway address cannot be the same as a LIF address. A LIF is already configured with that IP address. |
| 1967080 | The destination.address is missing. |
| 1967081 | The specified SVM must exist in the specified IPspace. |
| 1967082 | The specified ipspace.uuid and ipspace.name refer to different IPspaces. |
| 1967146 | The specified svm.name is not valid. |
| 2621462 | The specified SVM does not exist. |
| 2621574 | This operation is not permitted on an SVM that is configured as the destination of a MetroCluster SVM relationship. |
| 2621706 | The specified UUID and name refer to different SVMs. |
| 53282375 | The specified destination is invalid. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkIPRoutesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ip routes create default response has a 2xx status code
func (o *NetworkIPRoutesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ip routes create default response has a 3xx status code
func (o *NetworkIPRoutesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ip routes create default response has a 4xx status code
func (o *NetworkIPRoutesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ip routes create default response has a 5xx status code
func (o *NetworkIPRoutesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ip routes create default response a status code equal to that given
func (o *NetworkIPRoutesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ip routes create default response
func (o *NetworkIPRoutesCreateDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPRoutesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ip/routes][%d] network_ip_routes_create default %s", o._statusCode, payload)
}

func (o *NetworkIPRoutesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/ip/routes][%d] network_ip_routes_create default %s", o._statusCode, payload)
}

func (o *NetworkIPRoutesCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPRoutesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
