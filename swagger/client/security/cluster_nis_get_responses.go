// Code generated by go-swagger; DO NOT EDIT.

package security

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

// ClusterNisGetReader is a Reader for the ClusterNisGet structure.
type ClusterNisGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNisGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNisGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNisGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNisGetOK creates a ClusterNisGetOK with default headers values
func NewClusterNisGetOK() *ClusterNisGetOK {
	return &ClusterNisGetOK{}
}

/*
ClusterNisGetOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNisGetOK struct {
	Payload *models.ClusterNisService
}

// IsSuccess returns true when this cluster nis get o k response has a 2xx status code
func (o *ClusterNisGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster nis get o k response has a 3xx status code
func (o *ClusterNisGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster nis get o k response has a 4xx status code
func (o *ClusterNisGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster nis get o k response has a 5xx status code
func (o *ClusterNisGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster nis get o k response a status code equal to that given
func (o *ClusterNisGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster nis get o k response
func (o *ClusterNisGetOK) Code() int {
	return 200
}

func (o *ClusterNisGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/nis][%d] clusterNisGetOK %s", 200, payload)
}

func (o *ClusterNisGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/nis][%d] clusterNisGetOK %s", 200, payload)
}

func (o *ClusterNisGetOK) GetPayload() *models.ClusterNisService {
	return o.Payload
}

func (o *ClusterNisGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNisService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNisGetDefault creates a ClusterNisGetDefault with default headers values
func NewClusterNisGetDefault(code int) *ClusterNisGetDefault {
	return &ClusterNisGetDefault{
		_statusCode: code,
	}
}

/*
ClusterNisGetDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterNisGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster nis get default response has a 2xx status code
func (o *ClusterNisGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster nis get default response has a 3xx status code
func (o *ClusterNisGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster nis get default response has a 4xx status code
func (o *ClusterNisGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster nis get default response has a 5xx status code
func (o *ClusterNisGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster nis get default response a status code equal to that given
func (o *ClusterNisGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster nis get default response
func (o *ClusterNisGetDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNisGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/nis][%d] cluster_nis_get default %s", o._statusCode, payload)
}

func (o *ClusterNisGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/nis][%d] cluster_nis_get default %s", o._statusCode, payload)
}

func (o *ClusterNisGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNisGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
