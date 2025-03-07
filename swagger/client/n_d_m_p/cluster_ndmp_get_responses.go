// Code generated by go-swagger; DO NOT EDIT.

package n_d_m_p

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

// ClusterNdmpGetReader is a Reader for the ClusterNdmpGet structure.
type ClusterNdmpGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNdmpGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNdmpGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNdmpGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNdmpGetOK creates a ClusterNdmpGetOK with default headers values
func NewClusterNdmpGetOK() *ClusterNdmpGetOK {
	return &ClusterNdmpGetOK{}
}

/*
ClusterNdmpGetOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNdmpGetOK struct {
	Payload *models.ClusterNdmpProperties
}

// IsSuccess returns true when this cluster ndmp get o k response has a 2xx status code
func (o *ClusterNdmpGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ndmp get o k response has a 3xx status code
func (o *ClusterNdmpGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ndmp get o k response has a 4xx status code
func (o *ClusterNdmpGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ndmp get o k response has a 5xx status code
func (o *ClusterNdmpGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ndmp get o k response a status code equal to that given
func (o *ClusterNdmpGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster ndmp get o k response
func (o *ClusterNdmpGetOK) Code() int {
	return 200
}

func (o *ClusterNdmpGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp][%d] clusterNdmpGetOK %s", 200, payload)
}

func (o *ClusterNdmpGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp][%d] clusterNdmpGetOK %s", 200, payload)
}

func (o *ClusterNdmpGetOK) GetPayload() *models.ClusterNdmpProperties {
	return o.Payload
}

func (o *ClusterNdmpGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNdmpProperties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNdmpGetDefault creates a ClusterNdmpGetDefault with default headers values
func NewClusterNdmpGetDefault(code int) *ClusterNdmpGetDefault {
	return &ClusterNdmpGetDefault{
		_statusCode: code,
	}
}

/*
ClusterNdmpGetDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterNdmpGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster ndmp get default response has a 2xx status code
func (o *ClusterNdmpGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster ndmp get default response has a 3xx status code
func (o *ClusterNdmpGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster ndmp get default response has a 4xx status code
func (o *ClusterNdmpGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster ndmp get default response has a 5xx status code
func (o *ClusterNdmpGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster ndmp get default response a status code equal to that given
func (o *ClusterNdmpGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster ndmp get default response
func (o *ClusterNdmpGetDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNdmpGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp][%d] cluster_ndmp_get default %s", o._statusCode, payload)
}

func (o *ClusterNdmpGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp][%d] cluster_ndmp_get default %s", o._statusCode, payload)
}

func (o *ClusterNdmpGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNdmpGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
