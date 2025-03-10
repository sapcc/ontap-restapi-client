// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ClusterPeerGetReader is a Reader for the ClusterPeerGet structure.
type ClusterPeerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterPeerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterPeerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterPeerGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterPeerGetOK creates a ClusterPeerGetOK with default headers values
func NewClusterPeerGetOK() *ClusterPeerGetOK {
	return &ClusterPeerGetOK{}
}

/*
ClusterPeerGetOK describes a response with status code 200, with default header values.

OK
*/
type ClusterPeerGetOK struct {
	Payload *models.ClusterPeer
}

// IsSuccess returns true when this cluster peer get o k response has a 2xx status code
func (o *ClusterPeerGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster peer get o k response has a 3xx status code
func (o *ClusterPeerGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster peer get o k response has a 4xx status code
func (o *ClusterPeerGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster peer get o k response has a 5xx status code
func (o *ClusterPeerGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster peer get o k response a status code equal to that given
func (o *ClusterPeerGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster peer get o k response
func (o *ClusterPeerGetOK) Code() int {
	return 200
}

func (o *ClusterPeerGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/peers/{uuid}][%d] clusterPeerGetOK %s", 200, payload)
}

func (o *ClusterPeerGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/peers/{uuid}][%d] clusterPeerGetOK %s", 200, payload)
}

func (o *ClusterPeerGetOK) GetPayload() *models.ClusterPeer {
	return o.Payload
}

func (o *ClusterPeerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterPeer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterPeerGetDefault creates a ClusterPeerGetDefault with default headers values
func NewClusterPeerGetDefault(code int) *ClusterPeerGetDefault {
	return &ClusterPeerGetDefault{
		_statusCode: code,
	}
}

/*
ClusterPeerGetDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterPeerGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster peer get default response has a 2xx status code
func (o *ClusterPeerGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster peer get default response has a 3xx status code
func (o *ClusterPeerGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster peer get default response has a 4xx status code
func (o *ClusterPeerGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster peer get default response has a 5xx status code
func (o *ClusterPeerGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster peer get default response a status code equal to that given
func (o *ClusterPeerGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster peer get default response
func (o *ClusterPeerGetDefault) Code() int {
	return o._statusCode
}

func (o *ClusterPeerGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/peers/{uuid}][%d] cluster_peer_get default %s", o._statusCode, payload)
}

func (o *ClusterPeerGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/peers/{uuid}][%d] cluster_peer_get default %s", o._statusCode, payload)
}

func (o *ClusterPeerGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterPeerGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
