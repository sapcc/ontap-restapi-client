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

// StorageClusterGetReader is a Reader for the StorageClusterGet structure.
type StorageClusterGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageClusterGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageClusterGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageClusterGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageClusterGetOK creates a StorageClusterGetOK with default headers values
func NewStorageClusterGetOK() *StorageClusterGetOK {
	return &StorageClusterGetOK{}
}

/*
StorageClusterGetOK describes a response with status code 200, with default header values.

OK
*/
type StorageClusterGetOK struct {
	Payload *models.ClusterSpace
}

// IsSuccess returns true when this storage cluster get o k response has a 2xx status code
func (o *StorageClusterGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this storage cluster get o k response has a 3xx status code
func (o *StorageClusterGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage cluster get o k response has a 4xx status code
func (o *StorageClusterGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this storage cluster get o k response has a 5xx status code
func (o *StorageClusterGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this storage cluster get o k response a status code equal to that given
func (o *StorageClusterGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the storage cluster get o k response
func (o *StorageClusterGetOK) Code() int {
	return 200
}

func (o *StorageClusterGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/cluster][%d] storageClusterGetOK %s", 200, payload)
}

func (o *StorageClusterGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/cluster][%d] storageClusterGetOK %s", 200, payload)
}

func (o *StorageClusterGetOK) GetPayload() *models.ClusterSpace {
	return o.Payload
}

func (o *StorageClusterGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterSpace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageClusterGetDefault creates a StorageClusterGetDefault with default headers values
func NewStorageClusterGetDefault(code int) *StorageClusterGetDefault {
	return &StorageClusterGetDefault{
		_statusCode: code,
	}
}

/*
	StorageClusterGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 9240637 | This operation is not supported on this platform. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type StorageClusterGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this storage cluster get default response has a 2xx status code
func (o *StorageClusterGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this storage cluster get default response has a 3xx status code
func (o *StorageClusterGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this storage cluster get default response has a 4xx status code
func (o *StorageClusterGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this storage cluster get default response has a 5xx status code
func (o *StorageClusterGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this storage cluster get default response a status code equal to that given
func (o *StorageClusterGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the storage cluster get default response
func (o *StorageClusterGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageClusterGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/cluster][%d] storage_cluster_get default %s", o._statusCode, payload)
}

func (o *StorageClusterGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/cluster][%d] storage_cluster_get default %s", o._statusCode, payload)
}

func (o *StorageClusterGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StorageClusterGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
