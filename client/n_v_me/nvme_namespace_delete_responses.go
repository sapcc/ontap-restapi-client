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

	"github.com/sapcc/ontap-restapi/models"
)

// NvmeNamespaceDeleteReader is a Reader for the NvmeNamespaceDelete structure.
type NvmeNamespaceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeNamespaceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeNamespaceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewNvmeNamespaceDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeNamespaceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeNamespaceDeleteOK creates a NvmeNamespaceDeleteOK with default headers values
func NewNvmeNamespaceDeleteOK() *NvmeNamespaceDeleteOK {
	return &NvmeNamespaceDeleteOK{}
}

/*
NvmeNamespaceDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NvmeNamespaceDeleteOK struct {
}

// IsSuccess returns true when this nvme namespace delete o k response has a 2xx status code
func (o *NvmeNamespaceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme namespace delete o k response has a 3xx status code
func (o *NvmeNamespaceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme namespace delete o k response has a 4xx status code
func (o *NvmeNamespaceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme namespace delete o k response has a 5xx status code
func (o *NvmeNamespaceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme namespace delete o k response a status code equal to that given
func (o *NvmeNamespaceDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nvme namespace delete o k response
func (o *NvmeNamespaceDeleteOK) Code() int {
	return 200
}

func (o *NvmeNamespaceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/namespaces/{uuid}][%d] nvmeNamespaceDeleteOK", 200)
}

func (o *NvmeNamespaceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /storage/namespaces/{uuid}][%d] nvmeNamespaceDeleteOK", 200)
}

func (o *NvmeNamespaceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNvmeNamespaceDeleteAccepted creates a NvmeNamespaceDeleteAccepted with default headers values
func NewNvmeNamespaceDeleteAccepted() *NvmeNamespaceDeleteAccepted {
	return &NvmeNamespaceDeleteAccepted{}
}

/*
NvmeNamespaceDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type NvmeNamespaceDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this nvme namespace delete accepted response has a 2xx status code
func (o *NvmeNamespaceDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme namespace delete accepted response has a 3xx status code
func (o *NvmeNamespaceDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme namespace delete accepted response has a 4xx status code
func (o *NvmeNamespaceDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme namespace delete accepted response has a 5xx status code
func (o *NvmeNamespaceDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme namespace delete accepted response a status code equal to that given
func (o *NvmeNamespaceDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the nvme namespace delete accepted response
func (o *NvmeNamespaceDeleteAccepted) Code() int {
	return 202
}

func (o *NvmeNamespaceDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/namespaces/{uuid}][%d] nvmeNamespaceDeleteAccepted %s", 202, payload)
}

func (o *NvmeNamespaceDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/namespaces/{uuid}][%d] nvmeNamespaceDeleteAccepted %s", 202, payload)
}

func (o *NvmeNamespaceDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *NvmeNamespaceDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeNamespaceDeleteDefault creates a NvmeNamespaceDeleteDefault with default headers values
func NewNvmeNamespaceDeleteDefault(code int) *NvmeNamespaceDeleteDefault {
	return &NvmeNamespaceDeleteDefault{
		_statusCode: code,
	}
}

/*
	NvmeNamespaceDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 4 | The specified namespace was not found. |
| 72089796 | The namespace must be unmapped before deletion. |
| 72090016 | The namespace's aggregate is offline. The aggregate must be online to modify or remove the namespace. |
| 72090017 | The namespace's volume is offline. The volume must be online to modify or remove the namespace. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NvmeNamespaceDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nvme namespace delete default response has a 2xx status code
func (o *NvmeNamespaceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme namespace delete default response has a 3xx status code
func (o *NvmeNamespaceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme namespace delete default response has a 4xx status code
func (o *NvmeNamespaceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme namespace delete default response has a 5xx status code
func (o *NvmeNamespaceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme namespace delete default response a status code equal to that given
func (o *NvmeNamespaceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nvme namespace delete default response
func (o *NvmeNamespaceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NvmeNamespaceDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/namespaces/{uuid}][%d] nvme_namespace_delete default %s", o._statusCode, payload)
}

func (o *NvmeNamespaceDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/namespaces/{uuid}][%d] nvme_namespace_delete default %s", o._statusCode, payload)
}

func (o *NvmeNamespaceDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeNamespaceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
