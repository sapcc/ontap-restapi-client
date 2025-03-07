// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NvmeServiceModifyReader is a Reader for the NvmeServiceModify structure.
type NvmeServiceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeServiceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeServiceModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeServiceModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeServiceModifyOK creates a NvmeServiceModifyOK with default headers values
func NewNvmeServiceModifyOK() *NvmeServiceModifyOK {
	return &NvmeServiceModifyOK{}
}

/*
NvmeServiceModifyOK describes a response with status code 200, with default header values.

OK
*/
type NvmeServiceModifyOK struct {
}

// IsSuccess returns true when this nvme service modify o k response has a 2xx status code
func (o *NvmeServiceModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme service modify o k response has a 3xx status code
func (o *NvmeServiceModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme service modify o k response has a 4xx status code
func (o *NvmeServiceModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme service modify o k response has a 5xx status code
func (o *NvmeServiceModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme service modify o k response a status code equal to that given
func (o *NvmeServiceModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme service modify o k response
func (o *NvmeServiceModifyOK) Code() int {
	return 200
}

func (o *NvmeServiceModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/nvme/services/{svm.uuid}][%d] nvmeServiceModifyOK", 200)
}

func (o *NvmeServiceModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/nvme/services/{svm.uuid}][%d] nvmeServiceModifyOK", 200)
}

func (o *NvmeServiceModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNvmeServiceModifyDefault creates a NvmeServiceModifyDefault with default headers values
func NewNvmeServiceModifyDefault(code int) *NvmeServiceModifyDefault {
	return &NvmeServiceModifyDefault{
		_statusCode: code,
	}
}

/*
	NvmeServiceModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1115127 | The cluster lacks a valid NVMe license. |
| 2621462 | The supplied SVM does not exist. |
| 5374893 | The SVM is stopped. The SVM must be running to create an NVMe service. |
| 72089651 | The supplied SVM does not have an NVMe service. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeServiceModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme service modify default response has a 2xx status code
func (o *NvmeServiceModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme service modify default response has a 3xx status code
func (o *NvmeServiceModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme service modify default response has a 4xx status code
func (o *NvmeServiceModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme service modify default response has a 5xx status code
func (o *NvmeServiceModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme service modify default response a status code equal to that given
func (o *NvmeServiceModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme service modify default response
func (o *NvmeServiceModifyDefault) Code() int {
	return o._statusCode
}

func (o *NvmeServiceModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nvme/services/{svm.uuid}][%d] nvme_service_modify default %s", o._statusCode, payload)
}

func (o *NvmeServiceModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nvme/services/{svm.uuid}][%d] nvme_service_modify default %s", o._statusCode, payload)
}

func (o *NvmeServiceModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeServiceModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
