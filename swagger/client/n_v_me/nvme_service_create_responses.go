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

// NvmeServiceCreateReader is a Reader for the NvmeServiceCreate structure.
type NvmeServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNvmeServiceCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeServiceCreateCreated creates a NvmeServiceCreateCreated with default headers values
func NewNvmeServiceCreateCreated() *NvmeServiceCreateCreated {
	return &NvmeServiceCreateCreated{}
}

/*
NvmeServiceCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NvmeServiceCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.NvmeServiceResponse
}

// IsSuccess returns true when this nvme service create created response has a 2xx status code
func (o *NvmeServiceCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme service create created response has a 3xx status code
func (o *NvmeServiceCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme service create created response has a 4xx status code
func (o *NvmeServiceCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme service create created response has a 5xx status code
func (o *NvmeServiceCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme service create created response a status code equal to that given
func (o *NvmeServiceCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the nvme service create created response
func (o *NvmeServiceCreateCreated) Code() int {
	return 201
}

func (o *NvmeServiceCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nvme/services][%d] nvmeServiceCreateCreated %s", 201, payload)
}

func (o *NvmeServiceCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nvme/services][%d] nvmeServiceCreateCreated %s", 201, payload)
}

func (o *NvmeServiceCreateCreated) GetPayload() *models.NvmeServiceResponse {
	return o.Payload
}

func (o *NvmeServiceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.NvmeServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeServiceCreateDefault creates a NvmeServiceCreateDefault with default headers values
func NewNvmeServiceCreateDefault(code int) *NvmeServiceCreateDefault {
	return &NvmeServiceCreateDefault{
		_statusCode: code,
	}
}

/*
	NvmeServiceCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1115127 | The cluster lacks a valid NVMe license. |
| 2621462 | The supplied SVM does not exist. |
| 2621507 | NVMe is not allowed for the specified SVM. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5374893 | The SVM is stopped. The SVM must be running to create an NVMe service. |
| 72089650 | An NVMe service already exists for the specified SVM. |
| 72089900 | An NVMe service cannot be creating in an SVM that is configured for a SAN protocol. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeServiceCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme service create default response has a 2xx status code
func (o *NvmeServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme service create default response has a 3xx status code
func (o *NvmeServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme service create default response has a 4xx status code
func (o *NvmeServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme service create default response has a 5xx status code
func (o *NvmeServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme service create default response a status code equal to that given
func (o *NvmeServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme service create default response
func (o *NvmeServiceCreateDefault) Code() int {
	return o._statusCode
}

func (o *NvmeServiceCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nvme/services][%d] nvme_service_create default %s", o._statusCode, payload)
}

func (o *NvmeServiceCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nvme/services][%d] nvme_service_create default %s", o._statusCode, payload)
}

func (o *NvmeServiceCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
