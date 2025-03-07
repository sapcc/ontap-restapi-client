// Code generated by go-swagger; DO NOT EDIT.

package s_vm

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

// SvmCreateReader is a Reader for the SvmCreate structure.
type SvmCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSvmCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSvmCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmCreateCreated creates a SvmCreateCreated with default headers values
func NewSvmCreateCreated() *SvmCreateCreated {
	return &SvmCreateCreated{}
}

/*
SvmCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SvmCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this svm create created response has a 2xx status code
func (o *SvmCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm create created response has a 3xx status code
func (o *SvmCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm create created response has a 4xx status code
func (o *SvmCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm create created response has a 5xx status code
func (o *SvmCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this svm create created response a status code equal to that given
func (o *SvmCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the svm create created response
func (o *SvmCreateCreated) Code() int {
	return 201
}

func (o *SvmCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/svms][%d] svmCreateCreated %s", 201, payload)
}

func (o *SvmCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/svms][%d] svmCreateCreated %s", 201, payload)
}

func (o *SvmCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SvmCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSvmCreateAccepted creates a SvmCreateAccepted with default headers values
func NewSvmCreateAccepted() *SvmCreateAccepted {
	return &SvmCreateAccepted{}
}

/*
SvmCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this svm create accepted response has a 2xx status code
func (o *SvmCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm create accepted response has a 3xx status code
func (o *SvmCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm create accepted response has a 4xx status code
func (o *SvmCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm create accepted response has a 5xx status code
func (o *SvmCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this svm create accepted response a status code equal to that given
func (o *SvmCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the svm create accepted response
func (o *SvmCreateAccepted) Code() int {
	return 202
}

func (o *SvmCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/svms][%d] svmCreateAccepted %s", 202, payload)
}

func (o *SvmCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/svms][%d] svmCreateAccepted %s", 202, payload)
}

func (o *SvmCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SvmCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSvmCreateDefault creates a SvmCreateDefault with default headers values
func NewSvmCreateDefault(code int) *SvmCreateDefault {
	return &SvmCreateDefault{
		_statusCode: code,
	}
}

/*
	SvmCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 2621580     | Cannot specify options other than SVM name, comment and ipspace for a Vserver that is being configured as the destination for SVM DR. |
| 2621634     | \"sync-source\" SVM can only be created in a MetroCluster configuration. |
| 2621657     | \"sync-destination\" SVM can only be created by the system. |
| 13434884    | Cannot create an SVM because of incorrect fields. |
| 13434885    | Non-UTF8 language(s) not supported. |
| 13434888    | IPspace UUID and IPspace name mismatch. |
| 13434889    | Internal Error. Wait and retry. |
| 13434894    | Maximum allowed SVM jobs exceeded. Wait for the existing SVM jobs to complete and try again. |
| 13434908    | Invalid SVM name. The name is already in use by another SVM, IPSpace or cluster. |
| 13434909    | Internal Error. Failed to identify the aggregate to host SVM root volume. |
| 13434910    | Internal Error. Failed to allocate new SVM ID. |
| 13434911    | Invalid SVM name. Maximum supported length is 41 if SVM is of type \\\"sync-source\\\", otherwise 47. |
| 13434912    | Failed to find IPspace. |
| 13434913    | Internal error: Failed to check if an SVM create operation is in progress. Contact technical support for assistance. |
| 13434914    | Request to create the root volume of the SVM failed because there is not enough space in specified aggregate. |
| 13434915    | Failed to unlock the SVM because SVM create or delete job is in progress. Wait a few minutes, and then try the command again. |
| 13434916    | SVM is in the process of being created. Wait a few minutes, and then try the command again. |
| 13434917    | SVM creation successful. |
| 13434918    | IPspace name not provided for creating an SVM. |
| 458753      | Destination and gateway must belong to the same address family. |
| 13434935    | FCP, iSCSI and NVMe cannot be disabled or disallowed on this platform. |
| 23724038    | Invalid source for the provided ns-switch database. |
```
<br/>
*/
type SvmCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm create default response has a 2xx status code
func (o *SvmCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm create default response has a 3xx status code
func (o *SvmCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm create default response has a 4xx status code
func (o *SvmCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm create default response has a 5xx status code
func (o *SvmCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm create default response a status code equal to that given
func (o *SvmCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm create default response
func (o *SvmCreateDefault) Code() int {
	return o._statusCode
}

func (o *SvmCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/svms][%d] svm_create default %s", o._statusCode, payload)
}

func (o *SvmCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /svm/svms][%d] svm_create default %s", o._statusCode, payload)
}

func (o *SvmCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
