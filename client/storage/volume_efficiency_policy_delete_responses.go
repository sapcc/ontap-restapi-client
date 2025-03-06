// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// VolumeEfficiencyPolicyDeleteReader is a Reader for the VolumeEfficiencyPolicyDelete structure.
type VolumeEfficiencyPolicyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeEfficiencyPolicyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeEfficiencyPolicyDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeEfficiencyPolicyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeEfficiencyPolicyDeleteOK creates a VolumeEfficiencyPolicyDeleteOK with default headers values
func NewVolumeEfficiencyPolicyDeleteOK() *VolumeEfficiencyPolicyDeleteOK {
	return &VolumeEfficiencyPolicyDeleteOK{}
}

/*
VolumeEfficiencyPolicyDeleteOK describes a response with status code 200, with default header values.

OK
*/
type VolumeEfficiencyPolicyDeleteOK struct {
}

// IsSuccess returns true when this volume efficiency policy delete o k response has a 2xx status code
func (o *VolumeEfficiencyPolicyDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume efficiency policy delete o k response has a 3xx status code
func (o *VolumeEfficiencyPolicyDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume efficiency policy delete o k response has a 4xx status code
func (o *VolumeEfficiencyPolicyDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume efficiency policy delete o k response has a 5xx status code
func (o *VolumeEfficiencyPolicyDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume efficiency policy delete o k response a status code equal to that given
func (o *VolumeEfficiencyPolicyDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volume efficiency policy delete o k response
func (o *VolumeEfficiencyPolicyDeleteOK) Code() int {
	return 200
}

func (o *VolumeEfficiencyPolicyDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/volume-efficiency-policies/{uuid}][%d] volumeEfficiencyPolicyDeleteOK", 200)
}

func (o *VolumeEfficiencyPolicyDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /storage/volume-efficiency-policies/{uuid}][%d] volumeEfficiencyPolicyDeleteOK", 200)
}

func (o *VolumeEfficiencyPolicyDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVolumeEfficiencyPolicyDeleteDefault creates a VolumeEfficiencyPolicyDeleteDefault with default headers values
func NewVolumeEfficiencyPolicyDeleteDefault(code int) *VolumeEfficiencyPolicyDeleteDefault {
	return &VolumeEfficiencyPolicyDeleteDefault{
		_statusCode: code,
	}
}

/*
	VolumeEfficiencyPolicyDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Code

| Error Code | Description |
| ---------  | ----------- |
|  918702    | The specified operation on the volume efficiency policies endpoint is not supported on this platform. |
|  6881346   | The policy was not deleted because the policy is in use by at least one volume. |
|  6881431   | The specified policy is a predefined policy and cannot be deleted. |
*/
type VolumeEfficiencyPolicyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this volume efficiency policy delete default response has a 2xx status code
func (o *VolumeEfficiencyPolicyDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume efficiency policy delete default response has a 3xx status code
func (o *VolumeEfficiencyPolicyDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume efficiency policy delete default response has a 4xx status code
func (o *VolumeEfficiencyPolicyDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume efficiency policy delete default response has a 5xx status code
func (o *VolumeEfficiencyPolicyDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume efficiency policy delete default response a status code equal to that given
func (o *VolumeEfficiencyPolicyDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume efficiency policy delete default response
func (o *VolumeEfficiencyPolicyDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VolumeEfficiencyPolicyDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volume-efficiency-policies/{uuid}][%d] volume_efficiency_policy_delete default %s", o._statusCode, payload)
}

func (o *VolumeEfficiencyPolicyDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/volume-efficiency-policies/{uuid}][%d] volume_efficiency_policy_delete default %s", o._statusCode, payload)
}

func (o *VolumeEfficiencyPolicyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeEfficiencyPolicyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
