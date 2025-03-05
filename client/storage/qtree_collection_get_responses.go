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

	"github.com/sapcc/ontap-restapi/models"
)

// QtreeCollectionGetReader is a Reader for the QtreeCollectionGet structure.
type QtreeCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QtreeCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQtreeCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQtreeCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQtreeCollectionGetOK creates a QtreeCollectionGetOK with default headers values
func NewQtreeCollectionGetOK() *QtreeCollectionGetOK {
	return &QtreeCollectionGetOK{}
}

/*
QtreeCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type QtreeCollectionGetOK struct {
	Payload *models.QtreeResponse
}

// IsSuccess returns true when this qtree collection get o k response has a 2xx status code
func (o *QtreeCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qtree collection get o k response has a 3xx status code
func (o *QtreeCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qtree collection get o k response has a 4xx status code
func (o *QtreeCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this qtree collection get o k response has a 5xx status code
func (o *QtreeCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this qtree collection get o k response a status code equal to that given
func (o *QtreeCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the qtree collection get o k response
func (o *QtreeCollectionGetOK) Code() int {
	return 200
}

func (o *QtreeCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qtrees][%d] qtreeCollectionGetOK %s", 200, payload)
}

func (o *QtreeCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qtrees][%d] qtreeCollectionGetOK %s", 200, payload)
}

func (o *QtreeCollectionGetOK) GetPayload() *models.QtreeResponse {
	return o.Payload
}

func (o *QtreeCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QtreeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQtreeCollectionGetDefault creates a QtreeCollectionGetDefault with default headers values
func NewQtreeCollectionGetDefault(code int) *QtreeCollectionGetDefault {
	return &QtreeCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	QtreeCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 918235 | A volume with UUID was not found. |
| 2621462 | The specified SVM does not exist. |
| 5242889 | Failed to get the qtree from volume. |
| 5242956 | Failed to obtain qtree. |
| 5242965 | Invalid qtree path. The volume name component of the qtree path, must be the same as the volume specified with the parameter. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type QtreeCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this qtree collection get default response has a 2xx status code
func (o *QtreeCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qtree collection get default response has a 3xx status code
func (o *QtreeCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qtree collection get default response has a 4xx status code
func (o *QtreeCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qtree collection get default response has a 5xx status code
func (o *QtreeCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qtree collection get default response a status code equal to that given
func (o *QtreeCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the qtree collection get default response
func (o *QtreeCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *QtreeCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qtrees][%d] qtree_collection_get default %s", o._statusCode, payload)
}

func (o *QtreeCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qtrees][%d] qtree_collection_get default %s", o._statusCode, payload)
}

func (o *QtreeCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QtreeCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
