// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NisModifyReader is a Reader for the NisModify structure.
type NisModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NisModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNisModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNisModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNisModifyOK creates a NisModifyOK with default headers values
func NewNisModifyOK() *NisModifyOK {
	return &NisModifyOK{}
}

/*
NisModifyOK describes a response with status code 200, with default header values.

OK
*/
type NisModifyOK struct {
}

// IsSuccess returns true when this nis modify o k response has a 2xx status code
func (o *NisModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nis modify o k response has a 3xx status code
func (o *NisModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nis modify o k response has a 4xx status code
func (o *NisModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nis modify o k response has a 5xx status code
func (o *NisModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nis modify o k response a status code equal to that given
func (o *NisModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nis modify o k response
func (o *NisModifyOK) Code() int {
	return 200
}

func (o *NisModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/nis/{svm.uuid}][%d] nisModifyOK", 200)
}

func (o *NisModifyOK) String() string {
	return fmt.Sprintf("[PATCH /name-services/nis/{svm.uuid}][%d] nisModifyOK", 200)
}

func (o *NisModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNisModifyDefault creates a NisModifyDefault with default headers values
func NewNisModifyDefault(code int) *NisModifyDefault {
	return &NisModifyDefault{
		_statusCode: code,
	}
}

/*
	NisModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1966253    | IPv6 is not enabled in the cluster |
| 2621488    | Invalid SVM context |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name |
| 3276964    | NIS domain name or NIS server domain is too long. The maximum supported for domain name is 64 characters and the maximum supported for NIS server domain is 255 characters |
| 3276933    | A maximum of 10 NIS servers can be configured per SVM |
| 23724109   | DNS resolution failed for one or more specified servers  |
| 23724111   | Empty NIS servers cannot be specified.  |
| 23724112   | DNS resolution failed due to an internal error. Contact technical support if this issue persists |
| 23724132   | DNS resolution failed for all the specified servers  |
| 23724130   | Cannot use an IPv6 name server address because there are no IPv6 LIFs |
*/
type NisModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nis modify default response has a 2xx status code
func (o *NisModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nis modify default response has a 3xx status code
func (o *NisModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nis modify default response has a 4xx status code
func (o *NisModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nis modify default response has a 5xx status code
func (o *NisModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nis modify default response a status code equal to that given
func (o *NisModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nis modify default response
func (o *NisModifyDefault) Code() int {
	return o._statusCode
}

func (o *NisModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/nis/{svm.uuid}][%d] nis_modify default %s", o._statusCode, payload)
}

func (o *NisModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/nis/{svm.uuid}][%d] nis_modify default %s", o._statusCode, payload)
}

func (o *NisModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NisModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
