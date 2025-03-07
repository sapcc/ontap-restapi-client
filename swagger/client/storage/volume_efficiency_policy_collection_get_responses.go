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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// VolumeEfficiencyPolicyCollectionGetReader is a Reader for the VolumeEfficiencyPolicyCollectionGet structure.
type VolumeEfficiencyPolicyCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeEfficiencyPolicyCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeEfficiencyPolicyCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeEfficiencyPolicyCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeEfficiencyPolicyCollectionGetOK creates a VolumeEfficiencyPolicyCollectionGetOK with default headers values
func NewVolumeEfficiencyPolicyCollectionGetOK() *VolumeEfficiencyPolicyCollectionGetOK {
	return &VolumeEfficiencyPolicyCollectionGetOK{}
}

/*
VolumeEfficiencyPolicyCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type VolumeEfficiencyPolicyCollectionGetOK struct {
	Payload *models.VolumeEfficiencyPolicyResponse
}

// IsSuccess returns true when this volume efficiency policy collection get o k response has a 2xx status code
func (o *VolumeEfficiencyPolicyCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume efficiency policy collection get o k response has a 3xx status code
func (o *VolumeEfficiencyPolicyCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume efficiency policy collection get o k response has a 4xx status code
func (o *VolumeEfficiencyPolicyCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume efficiency policy collection get o k response has a 5xx status code
func (o *VolumeEfficiencyPolicyCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume efficiency policy collection get o k response a status code equal to that given
func (o *VolumeEfficiencyPolicyCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volume efficiency policy collection get o k response
func (o *VolumeEfficiencyPolicyCollectionGetOK) Code() int {
	return 200
}

func (o *VolumeEfficiencyPolicyCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volume-efficiency-policies][%d] volumeEfficiencyPolicyCollectionGetOK %s", 200, payload)
}

func (o *VolumeEfficiencyPolicyCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volume-efficiency-policies][%d] volumeEfficiencyPolicyCollectionGetOK %s", 200, payload)
}

func (o *VolumeEfficiencyPolicyCollectionGetOK) GetPayload() *models.VolumeEfficiencyPolicyResponse {
	return o.Payload
}

func (o *VolumeEfficiencyPolicyCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeEfficiencyPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeEfficiencyPolicyCollectionGetDefault creates a VolumeEfficiencyPolicyCollectionGetDefault with default headers values
func NewVolumeEfficiencyPolicyCollectionGetDefault(code int) *VolumeEfficiencyPolicyCollectionGetDefault {
	return &VolumeEfficiencyPolicyCollectionGetDefault{
		_statusCode: code,
	}
}

/*
VolumeEfficiencyPolicyCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type VolumeEfficiencyPolicyCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this volume efficiency policy collection get default response has a 2xx status code
func (o *VolumeEfficiencyPolicyCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume efficiency policy collection get default response has a 3xx status code
func (o *VolumeEfficiencyPolicyCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume efficiency policy collection get default response has a 4xx status code
func (o *VolumeEfficiencyPolicyCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume efficiency policy collection get default response has a 5xx status code
func (o *VolumeEfficiencyPolicyCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume efficiency policy collection get default response a status code equal to that given
func (o *VolumeEfficiencyPolicyCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume efficiency policy collection get default response
func (o *VolumeEfficiencyPolicyCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *VolumeEfficiencyPolicyCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volume-efficiency-policies][%d] volume_efficiency_policy_collection_get default %s", o._statusCode, payload)
}

func (o *VolumeEfficiencyPolicyCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/volume-efficiency-policies][%d] volume_efficiency_policy_collection_get default %s", o._statusCode, payload)
}

func (o *VolumeEfficiencyPolicyCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeEfficiencyPolicyCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
