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

// DiskCollectionGetReader is a Reader for the DiskCollectionGet structure.
type DiskCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiskCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiskCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDiskCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDiskCollectionGetOK creates a DiskCollectionGetOK with default headers values
func NewDiskCollectionGetOK() *DiskCollectionGetOK {
	return &DiskCollectionGetOK{}
}

/*
DiskCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type DiskCollectionGetOK struct {
	Payload *models.DiskResponse
}

// IsSuccess returns true when this disk collection get o k response has a 2xx status code
func (o *DiskCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disk collection get o k response has a 3xx status code
func (o *DiskCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disk collection get o k response has a 4xx status code
func (o *DiskCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this disk collection get o k response has a 5xx status code
func (o *DiskCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this disk collection get o k response a status code equal to that given
func (o *DiskCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the disk collection get o k response
func (o *DiskCollectionGetOK) Code() int {
	return 200
}

func (o *DiskCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/disks][%d] diskCollectionGetOK %s", 200, payload)
}

func (o *DiskCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/disks][%d] diskCollectionGetOK %s", 200, payload)
}

func (o *DiskCollectionGetOK) GetPayload() *models.DiskResponse {
	return o.Payload
}

func (o *DiskCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DiskResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiskCollectionGetDefault creates a DiskCollectionGetDefault with default headers values
func NewDiskCollectionGetDefault(code int) *DiskCollectionGetDefault {
	return &DiskCollectionGetDefault{
		_statusCode: code,
	}
}

/*
DiskCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type DiskCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this disk collection get default response has a 2xx status code
func (o *DiskCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this disk collection get default response has a 3xx status code
func (o *DiskCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this disk collection get default response has a 4xx status code
func (o *DiskCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this disk collection get default response has a 5xx status code
func (o *DiskCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this disk collection get default response a status code equal to that given
func (o *DiskCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the disk collection get default response
func (o *DiskCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *DiskCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/disks][%d] disk_collection_get default %s", o._statusCode, payload)
}

func (o *DiskCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/disks][%d] disk_collection_get default %s", o._statusCode, payload)
}

func (o *DiskCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DiskCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
