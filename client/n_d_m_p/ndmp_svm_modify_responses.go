// Code generated by go-swagger; DO NOT EDIT.

package n_d_m_p

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

// NdmpSvmModifyReader is a Reader for the NdmpSvmModify structure.
type NdmpSvmModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NdmpSvmModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNdmpSvmModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNdmpSvmModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNdmpSvmModifyOK creates a NdmpSvmModifyOK with default headers values
func NewNdmpSvmModifyOK() *NdmpSvmModifyOK {
	return &NdmpSvmModifyOK{}
}

/*
NdmpSvmModifyOK describes a response with status code 200, with default header values.

OK
*/
type NdmpSvmModifyOK struct {
	Payload *models.NdmpSvm
}

// IsSuccess returns true when this ndmp svm modify o k response has a 2xx status code
func (o *NdmpSvmModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ndmp svm modify o k response has a 3xx status code
func (o *NdmpSvmModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ndmp svm modify o k response has a 4xx status code
func (o *NdmpSvmModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ndmp svm modify o k response has a 5xx status code
func (o *NdmpSvmModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ndmp svm modify o k response a status code equal to that given
func (o *NdmpSvmModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ndmp svm modify o k response
func (o *NdmpSvmModifyOK) Code() int {
	return 200
}

func (o *NdmpSvmModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/ndmp/svms/{svm.uuid}][%d] ndmpSvmModifyOK %s", 200, payload)
}

func (o *NdmpSvmModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/ndmp/svms/{svm.uuid}][%d] ndmpSvmModifyOK %s", 200, payload)
}

func (o *NdmpSvmModifyOK) GetPayload() *models.NdmpSvm {
	return o.Payload
}

func (o *NdmpSvmModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NdmpSvm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNdmpSvmModifyDefault creates a NdmpSvmModifyDefault with default headers values
func NewNdmpSvmModifyDefault(code int) *NdmpSvmModifyDefault {
	return &NdmpSvmModifyDefault{
		_statusCode: code,
	}
}

/*
	NdmpSvmModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 65601536    | The operation is not supported because NDMP SVM-aware mode is disabled.|
| 65601551    | Authentication type \"plaintext_sso\" cannot be combined with other authentication types.|
*/
type NdmpSvmModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ndmp svm modify default response has a 2xx status code
func (o *NdmpSvmModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ndmp svm modify default response has a 3xx status code
func (o *NdmpSvmModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ndmp svm modify default response has a 4xx status code
func (o *NdmpSvmModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ndmp svm modify default response has a 5xx status code
func (o *NdmpSvmModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ndmp svm modify default response a status code equal to that given
func (o *NdmpSvmModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ndmp svm modify default response
func (o *NdmpSvmModifyDefault) Code() int {
	return o._statusCode
}

func (o *NdmpSvmModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/ndmp/svms/{svm.uuid}][%d] ndmp_svm_modify default %s", o._statusCode, payload)
}

func (o *NdmpSvmModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/ndmp/svms/{svm.uuid}][%d] ndmp_svm_modify default %s", o._statusCode, payload)
}

func (o *NdmpSvmModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NdmpSvmModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
