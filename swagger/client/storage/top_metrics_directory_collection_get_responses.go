// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// TopMetricsDirectoryCollectionGetReader is a Reader for the TopMetricsDirectoryCollectionGet structure.
type TopMetricsDirectoryCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TopMetricsDirectoryCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTopMetricsDirectoryCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTopMetricsDirectoryCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTopMetricsDirectoryCollectionGetOK creates a TopMetricsDirectoryCollectionGetOK with default headers values
func NewTopMetricsDirectoryCollectionGetOK() *TopMetricsDirectoryCollectionGetOK {
	return &TopMetricsDirectoryCollectionGetOK{}
}

/*
TopMetricsDirectoryCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type TopMetricsDirectoryCollectionGetOK struct {
	Payload *models.TopMetricsDirectoryResponse
}

// IsSuccess returns true when this top metrics directory collection get o k response has a 2xx status code
func (o *TopMetricsDirectoryCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this top metrics directory collection get o k response has a 3xx status code
func (o *TopMetricsDirectoryCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this top metrics directory collection get o k response has a 4xx status code
func (o *TopMetricsDirectoryCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this top metrics directory collection get o k response has a 5xx status code
func (o *TopMetricsDirectoryCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this top metrics directory collection get o k response a status code equal to that given
func (o *TopMetricsDirectoryCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the top metrics directory collection get o k response
func (o *TopMetricsDirectoryCollectionGetOK) Code() int {
	return 200
}

func (o *TopMetricsDirectoryCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/top-metrics/directories][%d] topMetricsDirectoryCollectionGetOK %s", 200, payload)
}

func (o *TopMetricsDirectoryCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/top-metrics/directories][%d] topMetricsDirectoryCollectionGetOK %s", 200, payload)
}

func (o *TopMetricsDirectoryCollectionGetOK) GetPayload() *models.TopMetricsDirectoryResponse {
	return o.Payload
}

func (o *TopMetricsDirectoryCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TopMetricsDirectoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTopMetricsDirectoryCollectionGetDefault creates a TopMetricsDirectoryCollectionGetDefault with default headers values
func NewTopMetricsDirectoryCollectionGetDefault(code int) *TopMetricsDirectoryCollectionGetDefault {
	return &TopMetricsDirectoryCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	TopMetricsDirectoryCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 124518405 | Volume activity tracking is not supported on volumes that contain LUNs. |
| 124518407 | Volume activity tracking is not supported on FlexCache volumes. |
| 124518408 | Volume activity tracking is not supported on audit staging volumes. |
| 124518409 | Volume activity tracking is not supported on object store server volumes. |
| 124518410 | Volume activity tracking is not supported on SnapMirror destination volumes. |
| 124518411 | Enabling or disabling volume activity tracking is not supported on individual FlexGroup constituents. |
| 124518412 | Volume activity tracking is not supported on SnapLock volumes. |
| 124518414 | Volume activity tracking is not supported on volumes that contain NVMe namespaces. |
| 124518415 | Failed to get the volume activity tracking report on volume volume.name in SVM svm.name. |
| 124518416 | Internal error. Volume activity tracking report timed out for volume volume.name in SVM svm.name. |
| 124518417 | Volume wildcard queries are not supported for activity tracking reports. |
| 124518418 | The activity tracking report for volume volume.name in SVM svm.name returned zero records. Check whether the volume has read/write traffic. Refer to the REST API documentation for more information. |
| 124518422 | Volume activity tracking is not supported on All SAN Array clusters. |
| 124519410 | The large directory report for volume volume.name in SVM svm.name is not available because the file system analytics database version doesn't support this report. |
| 124519411 | Volume activity tracking is not enabled on the volume. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type TopMetricsDirectoryCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this top metrics directory collection get default response has a 2xx status code
func (o *TopMetricsDirectoryCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this top metrics directory collection get default response has a 3xx status code
func (o *TopMetricsDirectoryCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this top metrics directory collection get default response has a 4xx status code
func (o *TopMetricsDirectoryCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this top metrics directory collection get default response has a 5xx status code
func (o *TopMetricsDirectoryCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this top metrics directory collection get default response a status code equal to that given
func (o *TopMetricsDirectoryCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the top metrics directory collection get default response
func (o *TopMetricsDirectoryCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *TopMetricsDirectoryCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/top-metrics/directories][%d] top_metrics_directory_collection_get default %s", o._statusCode, payload)
}

func (o *TopMetricsDirectoryCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volumes/{volume.uuid}/top-metrics/directories][%d] top_metrics_directory_collection_get default %s", o._statusCode, payload)
}

func (o *TopMetricsDirectoryCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TopMetricsDirectoryCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
