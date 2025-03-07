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

// ClusterLdapDeleteReader is a Reader for the ClusterLdapDelete structure.
type ClusterLdapDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterLdapDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterLdapDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterLdapDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterLdapDeleteOK creates a ClusterLdapDeleteOK with default headers values
func NewClusterLdapDeleteOK() *ClusterLdapDeleteOK {
	return &ClusterLdapDeleteOK{}
}

/*
ClusterLdapDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ClusterLdapDeleteOK struct {
}

// IsSuccess returns true when this cluster ldap delete o k response has a 2xx status code
func (o *ClusterLdapDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ldap delete o k response has a 3xx status code
func (o *ClusterLdapDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ldap delete o k response has a 4xx status code
func (o *ClusterLdapDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ldap delete o k response has a 5xx status code
func (o *ClusterLdapDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ldap delete o k response a status code equal to that given
func (o *ClusterLdapDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster ldap delete o k response
func (o *ClusterLdapDeleteOK) Code() int {
	return 200
}

func (o *ClusterLdapDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/authentication/cluster/ldap][%d] clusterLdapDeleteOK", 200)
}

func (o *ClusterLdapDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/authentication/cluster/ldap][%d] clusterLdapDeleteOK", 200)
}

func (o *ClusterLdapDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterLdapDeleteDefault creates a ClusterLdapDeleteDefault with default headers values
func NewClusterLdapDeleteDefault(code int) *ClusterLdapDeleteDefault {
	return &ClusterLdapDeleteDefault{
		_statusCode: code,
	}
}

/*
ClusterLdapDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterLdapDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster ldap delete default response has a 2xx status code
func (o *ClusterLdapDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster ldap delete default response has a 3xx status code
func (o *ClusterLdapDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster ldap delete default response has a 4xx status code
func (o *ClusterLdapDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster ldap delete default response has a 5xx status code
func (o *ClusterLdapDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster ldap delete default response a status code equal to that given
func (o *ClusterLdapDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster ldap delete default response
func (o *ClusterLdapDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ClusterLdapDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/authentication/cluster/ldap][%d] cluster_ldap_delete default %s", o._statusCode, payload)
}

func (o *ClusterLdapDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/authentication/cluster/ldap][%d] cluster_ldap_delete default %s", o._statusCode, payload)
}

func (o *ClusterLdapDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterLdapDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
