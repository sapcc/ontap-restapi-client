// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// LunMapReportingNodeCreateReader is a Reader for the LunMapReportingNodeCreate structure.
type LunMapReportingNodeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunMapReportingNodeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLunMapReportingNodeCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunMapReportingNodeCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunMapReportingNodeCreateCreated creates a LunMapReportingNodeCreateCreated with default headers values
func NewLunMapReportingNodeCreateCreated() *LunMapReportingNodeCreateCreated {
	return &LunMapReportingNodeCreateCreated{}
}

/*
LunMapReportingNodeCreateCreated describes a response with status code 201, with default header values.

Created
*/
type LunMapReportingNodeCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.LunMapReportingNodeResponse
}

// IsSuccess returns true when this lun map reporting node create created response has a 2xx status code
func (o *LunMapReportingNodeCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun map reporting node create created response has a 3xx status code
func (o *LunMapReportingNodeCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun map reporting node create created response has a 4xx status code
func (o *LunMapReportingNodeCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun map reporting node create created response has a 5xx status code
func (o *LunMapReportingNodeCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this lun map reporting node create created response a status code equal to that given
func (o *LunMapReportingNodeCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the lun map reporting node create created response
func (o *LunMapReportingNodeCreateCreated) Code() int {
	return 201
}

func (o *LunMapReportingNodeCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}/reporting-nodes][%d] lunMapReportingNodeCreateCreated %s", 201, payload)
}

func (o *LunMapReportingNodeCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}/reporting-nodes][%d] lunMapReportingNodeCreateCreated %s", 201, payload)
}

func (o *LunMapReportingNodeCreateCreated) GetPayload() *models.LunMapReportingNodeResponse {
	return o.Payload
}

func (o *LunMapReportingNodeCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.LunMapReportingNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunMapReportingNodeCreateDefault creates a LunMapReportingNodeCreateDefault with default headers values
func NewLunMapReportingNodeCreateDefault(code int) *LunMapReportingNodeCreateDefault {
	return &LunMapReportingNodeCreateDefault{
		_statusCode: code,
	}
}

/*
	LunMapReportingNodeCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The specified LUN does not exist or is not accessible to the caller. |
| 5374878 | The specified initiator group does not exist, is not accessible to the caller, or is not in the same SVM as the specified LUN. |
| 5374920 | The specified cluster node does not exist. |
| 5374921 | The specified cluster node name and UUID do not refer to the same cluster node. |
| 5374922 | The specified LUN map does not exist. |
| 5374923 | A cluster node `uuid` or `name` must be specified to add a reporting node. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LunMapReportingNodeCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this lun map reporting node create default response has a 2xx status code
func (o *LunMapReportingNodeCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun map reporting node create default response has a 3xx status code
func (o *LunMapReportingNodeCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun map reporting node create default response has a 4xx status code
func (o *LunMapReportingNodeCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun map reporting node create default response has a 5xx status code
func (o *LunMapReportingNodeCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun map reporting node create default response a status code equal to that given
func (o *LunMapReportingNodeCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lun map reporting node create default response
func (o *LunMapReportingNodeCreateDefault) Code() int {
	return o._statusCode
}

func (o *LunMapReportingNodeCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}/reporting-nodes][%d] lun_map_reporting_node_create default %s", o._statusCode, payload)
}

func (o *LunMapReportingNodeCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}/reporting-nodes][%d] lun_map_reporting_node_create default %s", o._statusCode, payload)
}

func (o *LunMapReportingNodeCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunMapReportingNodeCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
