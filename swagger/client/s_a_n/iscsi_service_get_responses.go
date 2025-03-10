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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// IscsiServiceGetReader is a Reader for the IscsiServiceGet structure.
type IscsiServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiServiceGetOK creates a IscsiServiceGetOK with default headers values
func NewIscsiServiceGetOK() *IscsiServiceGetOK {
	return &IscsiServiceGetOK{}
}

/*
IscsiServiceGetOK describes a response with status code 200, with default header values.

OK
*/
type IscsiServiceGetOK struct {
	Payload *models.IscsiService
}

// IsSuccess returns true when this iscsi service get o k response has a 2xx status code
func (o *IscsiServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iscsi service get o k response has a 3xx status code
func (o *IscsiServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iscsi service get o k response has a 4xx status code
func (o *IscsiServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iscsi service get o k response has a 5xx status code
func (o *IscsiServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iscsi service get o k response a status code equal to that given
func (o *IscsiServiceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iscsi service get o k response
func (o *IscsiServiceGetOK) Code() int {
	return 200
}

func (o *IscsiServiceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services/{svm.uuid}][%d] iscsiServiceGetOK %s", 200, payload)
}

func (o *IscsiServiceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services/{svm.uuid}][%d] iscsiServiceGetOK %s", 200, payload)
}

func (o *IscsiServiceGetOK) GetPayload() *models.IscsiService {
	return o.Payload
}

func (o *IscsiServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IscsiService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIscsiServiceGetDefault creates a IscsiServiceGetDefault with default headers values
func NewIscsiServiceGetDefault(code int) *IscsiServiceGetDefault {
	return &IscsiServiceGetDefault{
		_statusCode: code,
	}
}

/*
	IscsiServiceGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | An SVM with the specified UUID does not exist. |
| 5374078 | The SVM does not have an iSCSI service. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IscsiServiceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this iscsi service get default response has a 2xx status code
func (o *IscsiServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iscsi service get default response has a 3xx status code
func (o *IscsiServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iscsi service get default response has a 4xx status code
func (o *IscsiServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iscsi service get default response has a 5xx status code
func (o *IscsiServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iscsi service get default response a status code equal to that given
func (o *IscsiServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iscsi service get default response
func (o *IscsiServiceGetDefault) Code() int {
	return o._statusCode
}

func (o *IscsiServiceGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services/{svm.uuid}][%d] iscsi_service_get default %s", o._statusCode, payload)
}

func (o *IscsiServiceGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/san/iscsi/services/{svm.uuid}][%d] iscsi_service_get default %s", o._statusCode, payload)
}

func (o *IscsiServiceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
