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

	"github.com/sapcc/ontap-restapi-client/models"
)

// ClusterAccountAdProxyGetReader is a Reader for the ClusterAccountAdProxyGet structure.
type ClusterAccountAdProxyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterAccountAdProxyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterAccountAdProxyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterAccountAdProxyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterAccountAdProxyGetOK creates a ClusterAccountAdProxyGetOK with default headers values
func NewClusterAccountAdProxyGetOK() *ClusterAccountAdProxyGetOK {
	return &ClusterAccountAdProxyGetOK{}
}

/*
ClusterAccountAdProxyGetOK describes a response with status code 200, with default header values.

OK
*/
type ClusterAccountAdProxyGetOK struct {
	Payload *models.ClusterAdProxy
}

// IsSuccess returns true when this cluster account ad proxy get o k response has a 2xx status code
func (o *ClusterAccountAdProxyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster account ad proxy get o k response has a 3xx status code
func (o *ClusterAccountAdProxyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster account ad proxy get o k response has a 4xx status code
func (o *ClusterAccountAdProxyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster account ad proxy get o k response has a 5xx status code
func (o *ClusterAccountAdProxyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster account ad proxy get o k response a status code equal to that given
func (o *ClusterAccountAdProxyGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster account ad proxy get o k response
func (o *ClusterAccountAdProxyGetOK) Code() int {
	return 200
}

func (o *ClusterAccountAdProxyGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/ad-proxy][%d] clusterAccountAdProxyGetOK %s", 200, payload)
}

func (o *ClusterAccountAdProxyGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/ad-proxy][%d] clusterAccountAdProxyGetOK %s", 200, payload)
}

func (o *ClusterAccountAdProxyGetOK) GetPayload() *models.ClusterAdProxy {
	return o.Payload
}

func (o *ClusterAccountAdProxyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterAdProxy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterAccountAdProxyGetDefault creates a ClusterAccountAdProxyGetDefault with default headers values
func NewClusterAccountAdProxyGetDefault(code int) *ClusterAccountAdProxyGetDefault {
	return &ClusterAccountAdProxyGetDefault{
		_statusCode: code,
	}
}

/*
ClusterAccountAdProxyGetDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterAccountAdProxyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster account ad proxy get default response has a 2xx status code
func (o *ClusterAccountAdProxyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster account ad proxy get default response has a 3xx status code
func (o *ClusterAccountAdProxyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster account ad proxy get default response has a 4xx status code
func (o *ClusterAccountAdProxyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster account ad proxy get default response has a 5xx status code
func (o *ClusterAccountAdProxyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster account ad proxy get default response a status code equal to that given
func (o *ClusterAccountAdProxyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster account ad proxy get default response
func (o *ClusterAccountAdProxyGetDefault) Code() int {
	return o._statusCode
}

func (o *ClusterAccountAdProxyGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/ad-proxy][%d] cluster_account_ad_proxy_get default %s", o._statusCode, payload)
}

func (o *ClusterAccountAdProxyGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/cluster/ad-proxy][%d] cluster_account_ad_proxy_get default %s", o._statusCode, payload)
}

func (o *ClusterAccountAdProxyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterAccountAdProxyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
