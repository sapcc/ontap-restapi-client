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

	"github.com/sapcc/ontap-restapi/models"
)

// ClusterNtpServersModifyReader is a Reader for the ClusterNtpServersModify structure.
type ClusterNtpServersModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNtpServersModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNtpServersModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewClusterNtpServersModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNtpServersModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNtpServersModifyOK creates a ClusterNtpServersModifyOK with default headers values
func NewClusterNtpServersModifyOK() *ClusterNtpServersModifyOK {
	return &ClusterNtpServersModifyOK{}
}

/*
ClusterNtpServersModifyOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNtpServersModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cluster ntp servers modify o k response has a 2xx status code
func (o *ClusterNtpServersModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ntp servers modify o k response has a 3xx status code
func (o *ClusterNtpServersModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ntp servers modify o k response has a 4xx status code
func (o *ClusterNtpServersModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ntp servers modify o k response has a 5xx status code
func (o *ClusterNtpServersModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ntp servers modify o k response a status code equal to that given
func (o *ClusterNtpServersModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster ntp servers modify o k response
func (o *ClusterNtpServersModifyOK) Code() int {
	return 200
}

func (o *ClusterNtpServersModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers/{server}][%d] clusterNtpServersModifyOK %s", 200, payload)
}

func (o *ClusterNtpServersModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers/{server}][%d] clusterNtpServersModifyOK %s", 200, payload)
}

func (o *ClusterNtpServersModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterNtpServersModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNtpServersModifyAccepted creates a ClusterNtpServersModifyAccepted with default headers values
func NewClusterNtpServersModifyAccepted() *ClusterNtpServersModifyAccepted {
	return &ClusterNtpServersModifyAccepted{}
}

/*
ClusterNtpServersModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ClusterNtpServersModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cluster ntp servers modify accepted response has a 2xx status code
func (o *ClusterNtpServersModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ntp servers modify accepted response has a 3xx status code
func (o *ClusterNtpServersModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ntp servers modify accepted response has a 4xx status code
func (o *ClusterNtpServersModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ntp servers modify accepted response has a 5xx status code
func (o *ClusterNtpServersModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ntp servers modify accepted response a status code equal to that given
func (o *ClusterNtpServersModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cluster ntp servers modify accepted response
func (o *ClusterNtpServersModifyAccepted) Code() int {
	return 202
}

func (o *ClusterNtpServersModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers/{server}][%d] clusterNtpServersModifyAccepted %s", 202, payload)
}

func (o *ClusterNtpServersModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers/{server}][%d] clusterNtpServersModifyAccepted %s", 202, payload)
}

func (o *ClusterNtpServersModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterNtpServersModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNtpServersModifyDefault creates a ClusterNtpServersModifyDefault with default headers values
func NewClusterNtpServersModifyDefault(code int) *ClusterNtpServersModifyDefault {
	return &ClusterNtpServersModifyDefault{
		_statusCode: code,
	}
}

/*
	ClusterNtpServersModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2097163 | NTP server address was invalid. |
| 2097164 | NTP server address was invalid. |
| 2097165 | Could not resolve NTP server hostname. |
| 2097166 | NTP server address query returned no valid IP addresses. |
| 2097167 | Failed to connect to NTP server. |
| 2097169 | NTP server provided was not synchronized. |
| 2097174 | NTP server provided had too high of root distance. |
| 2097177 | NTP server provided had an invalid stratum. |
| 2097181 | NTP server address was invalid. |
| 2097182 | NTP server address was invalid. |
| 2097183 | NTP symmetric key authentication cannot be used for a node not in a cluster. |
| 2097185 | NTP key authentication failed for the provided key. |
| 2097188 | An invalid key identifier was provided. Identifiers must be in the range from 1 to 65535. |
| 2097193 | An unknown key was provided. |
| 2097194 | The field \"authentication_enabled\" cannot be false when the field NTP key is given. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ClusterNtpServersModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster ntp servers modify default response has a 2xx status code
func (o *ClusterNtpServersModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster ntp servers modify default response has a 3xx status code
func (o *ClusterNtpServersModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster ntp servers modify default response has a 4xx status code
func (o *ClusterNtpServersModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster ntp servers modify default response has a 5xx status code
func (o *ClusterNtpServersModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster ntp servers modify default response a status code equal to that given
func (o *ClusterNtpServersModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster ntp servers modify default response
func (o *ClusterNtpServersModifyDefault) Code() int {
	return o._statusCode
}

func (o *ClusterNtpServersModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers/{server}][%d] cluster_ntp_servers_modify default %s", o._statusCode, payload)
}

func (o *ClusterNtpServersModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/ntp/servers/{server}][%d] cluster_ntp_servers_modify default %s", o._statusCode, payload)
}

func (o *ClusterNtpServersModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNtpServersModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
