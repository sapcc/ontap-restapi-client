// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// CloudTargetDeleteReader is a Reader for the CloudTargetDelete structure.
type CloudTargetDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudTargetDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudTargetDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCloudTargetDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCloudTargetDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudTargetDeleteOK creates a CloudTargetDeleteOK with default headers values
func NewCloudTargetDeleteOK() *CloudTargetDeleteOK {
	return &CloudTargetDeleteOK{}
}

/*
CloudTargetDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CloudTargetDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cloud target delete o k response has a 2xx status code
func (o *CloudTargetDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud target delete o k response has a 3xx status code
func (o *CloudTargetDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud target delete o k response has a 4xx status code
func (o *CloudTargetDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud target delete o k response has a 5xx status code
func (o *CloudTargetDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud target delete o k response a status code equal to that given
func (o *CloudTargetDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cloud target delete o k response
func (o *CloudTargetDeleteOK) Code() int {
	return 200
}

func (o *CloudTargetDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloud/targets/{uuid}][%d] cloudTargetDeleteOK %s", 200, payload)
}

func (o *CloudTargetDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloud/targets/{uuid}][%d] cloudTargetDeleteOK %s", 200, payload)
}

func (o *CloudTargetDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *CloudTargetDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudTargetDeleteAccepted creates a CloudTargetDeleteAccepted with default headers values
func NewCloudTargetDeleteAccepted() *CloudTargetDeleteAccepted {
	return &CloudTargetDeleteAccepted{}
}

/*
CloudTargetDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type CloudTargetDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cloud target delete accepted response has a 2xx status code
func (o *CloudTargetDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud target delete accepted response has a 3xx status code
func (o *CloudTargetDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud target delete accepted response has a 4xx status code
func (o *CloudTargetDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud target delete accepted response has a 5xx status code
func (o *CloudTargetDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud target delete accepted response a status code equal to that given
func (o *CloudTargetDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cloud target delete accepted response
func (o *CloudTargetDeleteAccepted) Code() int {
	return 202
}

func (o *CloudTargetDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloud/targets/{uuid}][%d] cloudTargetDeleteAccepted %s", 202, payload)
}

func (o *CloudTargetDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloud/targets/{uuid}][%d] cloudTargetDeleteAccepted %s", 202, payload)
}

func (o *CloudTargetDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *CloudTargetDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudTargetDeleteDefault creates a CloudTargetDeleteDefault with default headers values
func NewCloudTargetDeleteDefault(code int) *CloudTargetDeleteDefault {
	return &CloudTargetDeleteDefault{
		_statusCode: code,
	}
}

/*
	CloudTargetDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 787013 | Cannot delete object store configuration as the store is attached to one or more aggregates. |
| 787014 | Cannot delete object store configuration as there are objects in the store. |
| 787078 | Cannot delete object store configuration. |
| 787188 | Object store configuration belongs to another cluster and cannot be modified from the local cluster, unless the cluster is in switchover mode. |
| 787216 | Cannot perform object store configuration operations on a cluster that is waiting for switchback. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type CloudTargetDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cloud target delete default response has a 2xx status code
func (o *CloudTargetDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cloud target delete default response has a 3xx status code
func (o *CloudTargetDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cloud target delete default response has a 4xx status code
func (o *CloudTargetDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cloud target delete default response has a 5xx status code
func (o *CloudTargetDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cloud target delete default response a status code equal to that given
func (o *CloudTargetDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cloud target delete default response
func (o *CloudTargetDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CloudTargetDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloud/targets/{uuid}][%d] cloud_target_delete default %s", o._statusCode, payload)
}

func (o *CloudTargetDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cloud/targets/{uuid}][%d] cloud_target_delete default %s", o._statusCode, payload)
}

func (o *CloudTargetDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CloudTargetDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
