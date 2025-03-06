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

	"github.com/sapcc/ontap-restapi-client/models"
)

// SoftwarePackageCreateReader is a Reader for the SoftwarePackageCreate structure.
type SoftwarePackageCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwarePackageCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSoftwarePackageCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSoftwarePackageCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwarePackageCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwarePackageCreateCreated creates a SoftwarePackageCreateCreated with default headers values
func NewSoftwarePackageCreateCreated() *SoftwarePackageCreateCreated {
	return &SoftwarePackageCreateCreated{}
}

/*
SoftwarePackageCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SoftwarePackageCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this software package create created response has a 2xx status code
func (o *SoftwarePackageCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this software package create created response has a 3xx status code
func (o *SoftwarePackageCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this software package create created response has a 4xx status code
func (o *SoftwarePackageCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this software package create created response has a 5xx status code
func (o *SoftwarePackageCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this software package create created response a status code equal to that given
func (o *SoftwarePackageCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the software package create created response
func (o *SoftwarePackageCreateCreated) Code() int {
	return 201
}

func (o *SoftwarePackageCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/software/download][%d] softwarePackageCreateCreated %s", 201, payload)
}

func (o *SoftwarePackageCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/software/download][%d] softwarePackageCreateCreated %s", 201, payload)
}

func (o *SoftwarePackageCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SoftwarePackageCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwarePackageCreateAccepted creates a SoftwarePackageCreateAccepted with default headers values
func NewSoftwarePackageCreateAccepted() *SoftwarePackageCreateAccepted {
	return &SoftwarePackageCreateAccepted{}
}

/*
SoftwarePackageCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SoftwarePackageCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this software package create accepted response has a 2xx status code
func (o *SoftwarePackageCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this software package create accepted response has a 3xx status code
func (o *SoftwarePackageCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this software package create accepted response has a 4xx status code
func (o *SoftwarePackageCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this software package create accepted response has a 5xx status code
func (o *SoftwarePackageCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this software package create accepted response a status code equal to that given
func (o *SoftwarePackageCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the software package create accepted response
func (o *SoftwarePackageCreateAccepted) Code() int {
	return 202
}

func (o *SoftwarePackageCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/software/download][%d] softwarePackageCreateAccepted %s", 202, payload)
}

func (o *SoftwarePackageCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/software/download][%d] softwarePackageCreateAccepted %s", 202, payload)
}

func (o *SoftwarePackageCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SoftwarePackageCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwarePackageCreateDefault creates a SoftwarePackageCreateDefault with default headers values
func NewSoftwarePackageCreateDefault(code int) *SoftwarePackageCreateDefault {
	return &SoftwarePackageCreateDefault{
		_statusCode: code,
	}
}

/*
	SoftwarePackageCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10551327 | Package image with the same name already exists. |
| 10551357 | Cannot perform an update when a previous update is still in progress. |
| 10551359 | Internal error. Failed to process the package after download completed. Try uploading the file again or contact technical support for assistance. |
| 10551381 | Package download failed. |
| 10551382 | Package download is still running. |
| 10551384 | Package download has not started. |
| 10551391 | Internal error. Contact technical support for assistance. |
| 10551392 | Internal error. Contact technical support for assistance. |
| 10551496 | Firmware file already exists. |
| 10551797 | Internal error. Failed to check if file upload is enabled. Contact technical support for assistance. |
| 10551859 | Failed to set primary and secondary nodes to store new image. |
| 39387137 | Invalid URL syntax was provided. Retry with a valid URL. |
| 39387138 | Unsupported URL scheme provided. Retry with one of FILE://, FTP://, or HTTPS://. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SoftwarePackageCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this software package create default response has a 2xx status code
func (o *SoftwarePackageCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this software package create default response has a 3xx status code
func (o *SoftwarePackageCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this software package create default response has a 4xx status code
func (o *SoftwarePackageCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this software package create default response has a 5xx status code
func (o *SoftwarePackageCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this software package create default response a status code equal to that given
func (o *SoftwarePackageCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the software package create default response
func (o *SoftwarePackageCreateDefault) Code() int {
	return o._statusCode
}

func (o *SoftwarePackageCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/software/download][%d] software_package_create default %s", o._statusCode, payload)
}

func (o *SoftwarePackageCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cluster/software/download][%d] software_package_create default %s", o._statusCode, payload)
}

func (o *SoftwarePackageCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SoftwarePackageCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
