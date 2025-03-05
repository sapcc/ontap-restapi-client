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

	"github.com/sapcc/ontap-restapi/models"
)

// AwsKmsDeleteReader is a Reader for the AwsKmsDelete structure.
type AwsKmsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsKmsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAwsKmsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAwsKmsDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAwsKmsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAwsKmsDeleteOK creates a AwsKmsDeleteOK with default headers values
func NewAwsKmsDeleteOK() *AwsKmsDeleteOK {
	return &AwsKmsDeleteOK{}
}

/*
AwsKmsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type AwsKmsDeleteOK struct {
}

// IsSuccess returns true when this aws kms delete o k response has a 2xx status code
func (o *AwsKmsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aws kms delete o k response has a 3xx status code
func (o *AwsKmsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws kms delete o k response has a 4xx status code
func (o *AwsKmsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws kms delete o k response has a 5xx status code
func (o *AwsKmsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aws kms delete o k response a status code equal to that given
func (o *AwsKmsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aws kms delete o k response
func (o *AwsKmsDeleteOK) Code() int {
	return 200
}

func (o *AwsKmsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/aws-kms/{uuid}][%d] awsKmsDeleteOK", 200)
}

func (o *AwsKmsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/aws-kms/{uuid}][%d] awsKmsDeleteOK", 200)
}

func (o *AwsKmsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAwsKmsDeleteAccepted creates a AwsKmsDeleteAccepted with default headers values
func NewAwsKmsDeleteAccepted() *AwsKmsDeleteAccepted {
	return &AwsKmsDeleteAccepted{}
}

/*
AwsKmsDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AwsKmsDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this aws kms delete accepted response has a 2xx status code
func (o *AwsKmsDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aws kms delete accepted response has a 3xx status code
func (o *AwsKmsDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws kms delete accepted response has a 4xx status code
func (o *AwsKmsDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws kms delete accepted response has a 5xx status code
func (o *AwsKmsDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this aws kms delete accepted response a status code equal to that given
func (o *AwsKmsDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the aws kms delete accepted response
func (o *AwsKmsDeleteAccepted) Code() int {
	return 202
}

func (o *AwsKmsDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/aws-kms/{uuid}][%d] awsKmsDeleteAccepted %s", 202, payload)
}

func (o *AwsKmsDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/aws-kms/{uuid}][%d] awsKmsDeleteAccepted %s", 202, payload)
}

func (o *AwsKmsDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *AwsKmsDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsKmsDeleteDefault creates a AwsKmsDeleteDefault with default headers values
func NewAwsKmsDeleteDefault(code int) *AwsKmsDeleteDefault {
	return &AwsKmsDeleteDefault{
		_statusCode: code,
	}
}

/*
	AwsKmsDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536817 | Internal error. Failed to determine if it is safe to disable key manager. |
| 65536827 | Internal error. Failed to determine if the given SVM has any encrypted volumes. |
| 65536834 | Internal error. Failed to get existing key-server details for the given SVM. |
| 65536883 | Internal error. Volume encryption key is missing for a volume. |
| 65536884 | Internal error. Volume encryption key is invalid for a volume. |
| 65537106 | Volume encryption keys (VEK) for one or more encrypted volumes are stored on the key manager configured for the given SVM. |
| 65537926 | Amazon Web Service Key Management Service is not configured for SVM. |
| 196608080 | One or more nodes in the cluster have the root volume encrypted using NVE (NetApp Volume Encryption). |
| 196608301 | Internal error. Failed to get encryption type. |
| 196608332 | NAE aggregates found in the cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AwsKmsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this aws kms delete default response has a 2xx status code
func (o *AwsKmsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aws kms delete default response has a 3xx status code
func (o *AwsKmsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aws kms delete default response has a 4xx status code
func (o *AwsKmsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aws kms delete default response has a 5xx status code
func (o *AwsKmsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aws kms delete default response a status code equal to that given
func (o *AwsKmsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the aws kms delete default response
func (o *AwsKmsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *AwsKmsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/aws-kms/{uuid}][%d] aws_kms_delete default %s", o._statusCode, payload)
}

func (o *AwsKmsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/aws-kms/{uuid}][%d] aws_kms_delete default %s", o._statusCode, payload)
}

func (o *AwsKmsDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AwsKmsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
