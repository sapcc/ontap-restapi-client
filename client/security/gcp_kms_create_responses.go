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

// GcpKmsCreateReader is a Reader for the GcpKmsCreate structure.
type GcpKmsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GcpKmsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGcpKmsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGcpKmsCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGcpKmsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGcpKmsCreateCreated creates a GcpKmsCreateCreated with default headers values
func NewGcpKmsCreateCreated() *GcpKmsCreateCreated {
	return &GcpKmsCreateCreated{}
}

/*
GcpKmsCreateCreated describes a response with status code 201, with default header values.

Created
*/
type GcpKmsCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.GcpKmsResponse
}

// IsSuccess returns true when this gcp kms create created response has a 2xx status code
func (o *GcpKmsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gcp kms create created response has a 3xx status code
func (o *GcpKmsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gcp kms create created response has a 4xx status code
func (o *GcpKmsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this gcp kms create created response has a 5xx status code
func (o *GcpKmsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this gcp kms create created response a status code equal to that given
func (o *GcpKmsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the gcp kms create created response
func (o *GcpKmsCreateCreated) Code() int {
	return 201
}

func (o *GcpKmsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms][%d] gcpKmsCreateCreated %s", 201, payload)
}

func (o *GcpKmsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms][%d] gcpKmsCreateCreated %s", 201, payload)
}

func (o *GcpKmsCreateCreated) GetPayload() *models.GcpKmsResponse {
	return o.Payload
}

func (o *GcpKmsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.GcpKmsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGcpKmsCreateAccepted creates a GcpKmsCreateAccepted with default headers values
func NewGcpKmsCreateAccepted() *GcpKmsCreateAccepted {
	return &GcpKmsCreateAccepted{}
}

/*
GcpKmsCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type GcpKmsCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this gcp kms create accepted response has a 2xx status code
func (o *GcpKmsCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gcp kms create accepted response has a 3xx status code
func (o *GcpKmsCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gcp kms create accepted response has a 4xx status code
func (o *GcpKmsCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this gcp kms create accepted response has a 5xx status code
func (o *GcpKmsCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this gcp kms create accepted response a status code equal to that given
func (o *GcpKmsCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the gcp kms create accepted response
func (o *GcpKmsCreateAccepted) Code() int {
	return 202
}

func (o *GcpKmsCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms][%d] gcpKmsCreateAccepted %s", 202, payload)
}

func (o *GcpKmsCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms][%d] gcpKmsCreateAccepted %s", 202, payload)
}

func (o *GcpKmsCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *GcpKmsCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGcpKmsCreateDefault creates a GcpKmsCreateDefault with default headers values
func NewGcpKmsCreateDefault(code int) *GcpKmsCreateDefault {
	return &GcpKmsCreateDefault{
		_statusCode: code,
	}
}

/*
	GcpKmsCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537703 | The Google Cloud Key Management Service is not supported for the admin Vserver. |
| 65537704 | The Google Cloud Key Management Service is not supported in MetroCluster configurations. |
| 65537706 | Internal error. Failed to the encrypt the application credentials. |
| 65537713 | Internal Error. Failed to store the application credentials. |
| 65537719 | Failed to enable the Google Cloud Key Management Service for SVM <svm-name> because invalid application credentials were provided. |
| 65537720 | Failed to configure Google Cloud Key Management Service for SVM <svm-name> because a key manager has already been configured for this SVM. Use the REST API GET method \"/api/security/gcp-kms\" to view all of the configured key managers. |
| 65537740 | The privileged account must be an email address or an empty string. |
| 65537749 | The application credentials field must be in valid JSON format. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type GcpKmsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this gcp kms create default response has a 2xx status code
func (o *GcpKmsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this gcp kms create default response has a 3xx status code
func (o *GcpKmsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this gcp kms create default response has a 4xx status code
func (o *GcpKmsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this gcp kms create default response has a 5xx status code
func (o *GcpKmsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this gcp kms create default response a status code equal to that given
func (o *GcpKmsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the gcp kms create default response
func (o *GcpKmsCreateDefault) Code() int {
	return o._statusCode
}

func (o *GcpKmsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms][%d] gcp_kms_create default %s", o._statusCode, payload)
}

func (o *GcpKmsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/gcp-kms][%d] gcp_kms_create default %s", o._statusCode, payload)
}

func (o *GcpKmsCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GcpKmsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
