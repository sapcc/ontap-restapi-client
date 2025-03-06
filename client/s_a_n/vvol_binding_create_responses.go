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

	"github.com/sapcc/ontap-restapi-client/models"
)

// VvolBindingCreateReader is a Reader for the VvolBindingCreate structure.
type VvolBindingCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VvolBindingCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVvolBindingCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVvolBindingCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVvolBindingCreateCreated creates a VvolBindingCreateCreated with default headers values
func NewVvolBindingCreateCreated() *VvolBindingCreateCreated {
	return &VvolBindingCreateCreated{}
}

/*
VvolBindingCreateCreated describes a response with status code 201, with default header values.

Created
*/
type VvolBindingCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.VvolBindingResponse
}

// IsSuccess returns true when this vvol binding create created response has a 2xx status code
func (o *VvolBindingCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vvol binding create created response has a 3xx status code
func (o *VvolBindingCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vvol binding create created response has a 4xx status code
func (o *VvolBindingCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this vvol binding create created response has a 5xx status code
func (o *VvolBindingCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this vvol binding create created response a status code equal to that given
func (o *VvolBindingCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the vvol binding create created response
func (o *VvolBindingCreateCreated) Code() int {
	return 201
}

func (o *VvolBindingCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/vvol-bindings][%d] vvolBindingCreateCreated %s", 201, payload)
}

func (o *VvolBindingCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/vvol-bindings][%d] vvolBindingCreateCreated %s", 201, payload)
}

func (o *VvolBindingCreateCreated) GetPayload() *models.VvolBindingResponse {
	return o.Payload
}

func (o *VvolBindingCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.VvolBindingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVvolBindingCreateDefault creates a VvolBindingCreateDefault with default headers values
func NewVvolBindingCreateDefault(code int) *VvolBindingCreateDefault {
	return &VvolBindingCreateDefault{
		_statusCode: code,
	}
}

/*
	VvolBindingCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | The specified SVM does not exist or is not accessible to the caller. |
| 2621706 | Both the SVM UUID and SVM name were supplied, but don't refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5374238 | A LUN in a snapshot was specified. |
| 5374323 | The LUN specified as the protocol endpoint LUN is not of class `protocol_endpoint`. |
| 5374325 | The LUN specified as the vVol LUN is not of class `vvol`. |
| 5374874 | The UUID and name supplied for the protocol endpoint of Vvol LUN do not refer to the same LUN. Use to the `target` property of the error object to differentiate between the protocol endpoint LUN and the vVol LUN. |
| 5374875 | The protocol endpoint or vVol LUN was not found or is not accessible to the caller. Use to the `target` property of the error object to differentiate between the protocol endpoint LUN and the vVol LUN. |
| 5374876 | The protocol endpoint or vVol LUN was not found in the SVM. Use to the `target` property of the error object to differentiate between the protocol endpoint LUN and the vVol LUN. |
| 5374924 | No protocol endpoint LUN was supplied. |
| 5374925 | No vVol LUN was supplied. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type VvolBindingCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vvol binding create default response has a 2xx status code
func (o *VvolBindingCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vvol binding create default response has a 3xx status code
func (o *VvolBindingCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vvol binding create default response has a 4xx status code
func (o *VvolBindingCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vvol binding create default response has a 5xx status code
func (o *VvolBindingCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vvol binding create default response a status code equal to that given
func (o *VvolBindingCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vvol binding create default response
func (o *VvolBindingCreateDefault) Code() int {
	return o._statusCode
}

func (o *VvolBindingCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/vvol-bindings][%d] vvol_binding_create default %s", o._statusCode, payload)
}

func (o *VvolBindingCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/vvol-bindings][%d] vvol_binding_create default %s", o._statusCode, payload)
}

func (o *VvolBindingCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VvolBindingCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
