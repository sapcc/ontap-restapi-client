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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NdmpSvmGetReader is a Reader for the NdmpSvmGet structure.
type NdmpSvmGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NdmpSvmGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNdmpSvmGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNdmpSvmGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNdmpSvmGetOK creates a NdmpSvmGetOK with default headers values
func NewNdmpSvmGetOK() *NdmpSvmGetOK {
	return &NdmpSvmGetOK{}
}

/*
NdmpSvmGetOK describes a response with status code 200, with default header values.

OK
*/
type NdmpSvmGetOK struct {
	Payload *models.NdmpSvm
}

// IsSuccess returns true when this ndmp svm get o k response has a 2xx status code
func (o *NdmpSvmGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ndmp svm get o k response has a 3xx status code
func (o *NdmpSvmGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ndmp svm get o k response has a 4xx status code
func (o *NdmpSvmGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ndmp svm get o k response has a 5xx status code
func (o *NdmpSvmGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ndmp svm get o k response a status code equal to that given
func (o *NdmpSvmGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ndmp svm get o k response
func (o *NdmpSvmGetOK) Code() int {
	return 200
}

func (o *NdmpSvmGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/svms/{svm.uuid}][%d] ndmpSvmGetOK %s", 200, payload)
}

func (o *NdmpSvmGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/svms/{svm.uuid}][%d] ndmpSvmGetOK %s", 200, payload)
}

func (o *NdmpSvmGetOK) GetPayload() *models.NdmpSvm {
	return o.Payload
}

func (o *NdmpSvmGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NdmpSvm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNdmpSvmGetDefault creates a NdmpSvmGetDefault with default headers values
func NewNdmpSvmGetDefault(code int) *NdmpSvmGetDefault {
	return &NdmpSvmGetDefault{
		_statusCode: code,
	}
}

/*
	NdmpSvmGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 2           | The UUID provided is an invalid value for field \"svm.uuid\".|
| 262197      | The value provided is invalid for field \"fields\".|
| 65601536    | The operation is not supported because NDMP SVM-aware mode is disabled.|
*/
type NdmpSvmGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ndmp svm get default response has a 2xx status code
func (o *NdmpSvmGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ndmp svm get default response has a 3xx status code
func (o *NdmpSvmGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ndmp svm get default response has a 4xx status code
func (o *NdmpSvmGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ndmp svm get default response has a 5xx status code
func (o *NdmpSvmGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ndmp svm get default response a status code equal to that given
func (o *NdmpSvmGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ndmp svm get default response
func (o *NdmpSvmGetDefault) Code() int {
	return o._statusCode
}

func (o *NdmpSvmGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/svms/{svm.uuid}][%d] ndmp_svm_get default %s", o._statusCode, payload)
}

func (o *NdmpSvmGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/svms/{svm.uuid}][%d] ndmp_svm_get default %s", o._statusCode, payload)
}

func (o *NdmpSvmGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NdmpSvmGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
