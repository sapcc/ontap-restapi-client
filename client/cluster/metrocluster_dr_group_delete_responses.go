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

// MetroclusterDrGroupDeleteReader is a Reader for the MetroclusterDrGroupDelete structure.
type MetroclusterDrGroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterDrGroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroclusterDrGroupDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewMetroclusterDrGroupDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterDrGroupDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterDrGroupDeleteOK creates a MetroclusterDrGroupDeleteOK with default headers values
func NewMetroclusterDrGroupDeleteOK() *MetroclusterDrGroupDeleteOK {
	return &MetroclusterDrGroupDeleteOK{}
}

/*
MetroclusterDrGroupDeleteOK describes a response with status code 200, with default header values.

OK
*/
type MetroclusterDrGroupDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this metrocluster dr group delete o k response has a 2xx status code
func (o *MetroclusterDrGroupDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster dr group delete o k response has a 3xx status code
func (o *MetroclusterDrGroupDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster dr group delete o k response has a 4xx status code
func (o *MetroclusterDrGroupDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster dr group delete o k response has a 5xx status code
func (o *MetroclusterDrGroupDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster dr group delete o k response a status code equal to that given
func (o *MetroclusterDrGroupDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the metrocluster dr group delete o k response
func (o *MetroclusterDrGroupDeleteOK) Code() int {
	return 200
}

func (o *MetroclusterDrGroupDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/metrocluster/dr-groups/{id}][%d] metroclusterDrGroupDeleteOK %s", 200, payload)
}

func (o *MetroclusterDrGroupDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/metrocluster/dr-groups/{id}][%d] metroclusterDrGroupDeleteOK %s", 200, payload)
}

func (o *MetroclusterDrGroupDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *MetroclusterDrGroupDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterDrGroupDeleteAccepted creates a MetroclusterDrGroupDeleteAccepted with default headers values
func NewMetroclusterDrGroupDeleteAccepted() *MetroclusterDrGroupDeleteAccepted {
	return &MetroclusterDrGroupDeleteAccepted{}
}

/*
MetroclusterDrGroupDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type MetroclusterDrGroupDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this metrocluster dr group delete accepted response has a 2xx status code
func (o *MetroclusterDrGroupDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster dr group delete accepted response has a 3xx status code
func (o *MetroclusterDrGroupDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster dr group delete accepted response has a 4xx status code
func (o *MetroclusterDrGroupDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster dr group delete accepted response has a 5xx status code
func (o *MetroclusterDrGroupDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster dr group delete accepted response a status code equal to that given
func (o *MetroclusterDrGroupDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the metrocluster dr group delete accepted response
func (o *MetroclusterDrGroupDeleteAccepted) Code() int {
	return 202
}

func (o *MetroclusterDrGroupDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/metrocluster/dr-groups/{id}][%d] metroclusterDrGroupDeleteAccepted %s", 202, payload)
}

func (o *MetroclusterDrGroupDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/metrocluster/dr-groups/{id}][%d] metroclusterDrGroupDeleteAccepted %s", 202, payload)
}

func (o *MetroclusterDrGroupDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *MetroclusterDrGroupDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterDrGroupDeleteDefault creates a MetroclusterDrGroupDeleteDefault with default headers values
func NewMetroclusterDrGroupDeleteDefault(code int) *MetroclusterDrGroupDeleteDefault {
	return &MetroclusterDrGroupDeleteDefault{
		_statusCode: code,
	}
}

/*
	MetroclusterDrGroupDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2425574 | Two or more DR groups must be configured to remove a DR group from the MetroCluster configuration. |
| 2425734 | An internal error occurred. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
| 2432833 | Operation is already running. |
| 2432859 | Unconfiguring MetroCluster DR Group |
| 2432860 | Unmirroring Aggregates |
| 2432861 | Unassigning Remote Disks |
| 2432862 | Disabling Cluster HA and Storage Failover HA |
| 2432863 | Disconnecting and deleting network connections |
| 2432864 | Unconfiguring and deleting the DR Group |
| 2432865 | Deleting MetroCluster DR Group |
| 2432866 | MetroCluster DR Group delete done |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type MetroclusterDrGroupDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this metrocluster dr group delete default response has a 2xx status code
func (o *MetroclusterDrGroupDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this metrocluster dr group delete default response has a 3xx status code
func (o *MetroclusterDrGroupDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this metrocluster dr group delete default response has a 4xx status code
func (o *MetroclusterDrGroupDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this metrocluster dr group delete default response has a 5xx status code
func (o *MetroclusterDrGroupDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this metrocluster dr group delete default response a status code equal to that given
func (o *MetroclusterDrGroupDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the metrocluster dr group delete default response
func (o *MetroclusterDrGroupDeleteDefault) Code() int {
	return o._statusCode
}

func (o *MetroclusterDrGroupDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/metrocluster/dr-groups/{id}][%d] metrocluster_dr_group_delete default %s", o._statusCode, payload)
}

func (o *MetroclusterDrGroupDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/metrocluster/dr-groups/{id}][%d] metrocluster_dr_group_delete default %s", o._statusCode, payload)
}

func (o *MetroclusterDrGroupDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetroclusterDrGroupDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
