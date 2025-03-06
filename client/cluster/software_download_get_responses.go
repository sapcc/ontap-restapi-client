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

// SoftwareDownloadGetReader is a Reader for the SoftwareDownloadGet structure.
type SoftwareDownloadGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwareDownloadGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSoftwareDownloadGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwareDownloadGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwareDownloadGetOK creates a SoftwareDownloadGetOK with default headers values
func NewSoftwareDownloadGetOK() *SoftwareDownloadGetOK {
	return &SoftwareDownloadGetOK{}
}

/*
SoftwareDownloadGetOK describes a response with status code 200, with default header values.

OK
*/
type SoftwareDownloadGetOK struct {
	Payload *models.SoftwarePackageDownloadGet
}

// IsSuccess returns true when this software download get o k response has a 2xx status code
func (o *SoftwareDownloadGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this software download get o k response has a 3xx status code
func (o *SoftwareDownloadGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this software download get o k response has a 4xx status code
func (o *SoftwareDownloadGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this software download get o k response has a 5xx status code
func (o *SoftwareDownloadGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this software download get o k response a status code equal to that given
func (o *SoftwareDownloadGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the software download get o k response
func (o *SoftwareDownloadGetOK) Code() int {
	return 200
}

func (o *SoftwareDownloadGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/software/download][%d] softwareDownloadGetOK %s", 200, payload)
}

func (o *SoftwareDownloadGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/software/download][%d] softwareDownloadGetOK %s", 200, payload)
}

func (o *SoftwareDownloadGetOK) GetPayload() *models.SoftwarePackageDownloadGet {
	return o.Payload
}

func (o *SoftwareDownloadGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwarePackageDownloadGet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareDownloadGetDefault creates a SoftwareDownloadGetDefault with default headers values
func NewSoftwareDownloadGetDefault(code int) *SoftwareDownloadGetDefault {
	return &SoftwareDownloadGetDefault{
		_statusCode: code,
	}
}

/*
	SoftwareDownloadGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10551382 | Package download is still running. |
| 10551383 | Operation took longer than the maximum 1 hour time limit. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SoftwareDownloadGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this software download get default response has a 2xx status code
func (o *SoftwareDownloadGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this software download get default response has a 3xx status code
func (o *SoftwareDownloadGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this software download get default response has a 4xx status code
func (o *SoftwareDownloadGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this software download get default response has a 5xx status code
func (o *SoftwareDownloadGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this software download get default response a status code equal to that given
func (o *SoftwareDownloadGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the software download get default response
func (o *SoftwareDownloadGetDefault) Code() int {
	return o._statusCode
}

func (o *SoftwareDownloadGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/software/download][%d] software_download_get default %s", o._statusCode, payload)
}

func (o *SoftwareDownloadGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/software/download][%d] software_download_get default %s", o._statusCode, payload)
}

func (o *SoftwareDownloadGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SoftwareDownloadGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
