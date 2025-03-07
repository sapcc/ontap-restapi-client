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

// ClusterModifyReader is a Reader for the ClusterModify structure.
type ClusterModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewClusterModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterModifyOK creates a ClusterModifyOK with default headers values
func NewClusterModifyOK() *ClusterModifyOK {
	return &ClusterModifyOK{}
}

/*
ClusterModifyOK describes a response with status code 200, with default header values.

OK
*/
type ClusterModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cluster modify o k response has a 2xx status code
func (o *ClusterModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster modify o k response has a 3xx status code
func (o *ClusterModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster modify o k response has a 4xx status code
func (o *ClusterModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster modify o k response has a 5xx status code
func (o *ClusterModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster modify o k response a status code equal to that given
func (o *ClusterModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster modify o k response
func (o *ClusterModifyOK) Code() int {
	return 200
}

func (o *ClusterModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster][%d] clusterModifyOK %s", 200, payload)
}

func (o *ClusterModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster][%d] clusterModifyOK %s", 200, payload)
}

func (o *ClusterModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterModifyAccepted creates a ClusterModifyAccepted with default headers values
func NewClusterModifyAccepted() *ClusterModifyAccepted {
	return &ClusterModifyAccepted{}
}

/*
ClusterModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ClusterModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cluster modify accepted response has a 2xx status code
func (o *ClusterModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster modify accepted response has a 3xx status code
func (o *ClusterModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster modify accepted response has a 4xx status code
func (o *ClusterModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster modify accepted response has a 5xx status code
func (o *ClusterModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster modify accepted response a status code equal to that given
func (o *ClusterModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cluster modify accepted response
func (o *ClusterModifyAccepted) Code() int {
	return 202
}

func (o *ClusterModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster][%d] clusterModifyAccepted %s", 202, payload)
}

func (o *ClusterModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster][%d] clusterModifyAccepted %s", 202, payload)
}

func (o *ClusterModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterModifyDefault creates a ClusterModifyDefault with default headers values
func NewClusterModifyDefault(code int) *ClusterModifyDefault {
	return &ClusterModifyDefault{
		_statusCode: code,
	}
}

/*
	ClusterModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655414 | Failed to the remove Active Directory machine. |
| 655431 | Username cannot be empty. |
| 655432 | Password cannot be empty. |
| 655562 | The NetBIOS name cannot be longer than 15 characters. |
| 655563 | NetBIOS name contains characters that are not allowed. |
| 655914 | Failed to create the Active Directory machine account. |
| 655915 | A CIFS server for this SVM already exists. Having both a CIFS server and an Active Directory account for the same SVM is not supported. Use the \\\"vserver cifs delete\\\" command to delete the existing CIFS server and try the command again. |
| 656464 | Failed to create the Active Directory machine account. Reason: Invalid Credentials. |
| 656465 | Failed to create the Active Directory machine account. Reason: An account with this name already exists. |
| 656466 | Failed to create the Active Directory machine account. Reason: Unable to connect to any domain controllers. |
| 656467 | Failed to create the Active Directory machine account. Reason: Organizational-Unit not found. |
| 656478 | Failed to create the Active Directory machine account. Reason: KDC has no support for encryption type. |
| 656483 | Active Directory account creation for the admin SVM requires an effective cluster version of 9.16.0 or later. |
| 3604491 | Updating timezone failed. |
| 3604520 | Internal error. System state is not correct to read or change timezone. |
| 8847361 | Too many DNS domains provided. |
| 8847362 | Too many name servers provided. |
| 8847400 | The "dns_domains" field is required when "name_servers" is specified. |
| 9240587 | A name must be provided. |
| 9240588 | The name is too long. |
| 12451843 | Certificate does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ClusterModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster modify default response has a 2xx status code
func (o *ClusterModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster modify default response has a 3xx status code
func (o *ClusterModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster modify default response has a 4xx status code
func (o *ClusterModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster modify default response has a 5xx status code
func (o *ClusterModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster modify default response a status code equal to that given
func (o *ClusterModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster modify default response
func (o *ClusterModifyDefault) Code() int {
	return o._statusCode
}

func (o *ClusterModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster][%d] cluster_modify default %s", o._statusCode, payload)
}

func (o *ClusterModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster][%d] cluster_modify default %s", o._statusCode, payload)
}

func (o *ClusterModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
