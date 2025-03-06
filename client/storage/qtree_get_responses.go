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

// QtreeGetReader is a Reader for the QtreeGet structure.
type QtreeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QtreeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQtreeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQtreeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQtreeGetOK creates a QtreeGetOK with default headers values
func NewQtreeGetOK() *QtreeGetOK {
	return &QtreeGetOK{}
}

/*
QtreeGetOK describes a response with status code 200, with default header values.

OK
*/
type QtreeGetOK struct {
	Payload *models.Qtree
}

// IsSuccess returns true when this qtree get o k response has a 2xx status code
func (o *QtreeGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qtree get o k response has a 3xx status code
func (o *QtreeGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qtree get o k response has a 4xx status code
func (o *QtreeGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this qtree get o k response has a 5xx status code
func (o *QtreeGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this qtree get o k response a status code equal to that given
func (o *QtreeGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the qtree get o k response
func (o *QtreeGetOK) Code() int {
	return 200
}

func (o *QtreeGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qtrees/{volume.uuid}/{id}][%d] qtreeGetOK %s", 200, payload)
}

func (o *QtreeGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qtrees/{volume.uuid}/{id}][%d] qtreeGetOK %s", 200, payload)
}

func (o *QtreeGetOK) GetPayload() *models.Qtree {
	return o.Payload
}

func (o *QtreeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Qtree)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQtreeGetDefault creates a QtreeGetDefault with default headers values
func NewQtreeGetDefault(code int) *QtreeGetDefault {
	return &QtreeGetDefault{
		_statusCode: code,
	}
}

/*
	QtreeGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 918235 | A volume with UUID was not found. |
| 2621462 | The specified SVM does not exist. |
| 5242889 | Failed to get the qtree from volume. |
| 5242956 | Failed to obtain a qtree with ID. |
| 5242965 | Invalid qtree path. The volume name component of the qtree path, must be the same as the volume specified with the parameter. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type QtreeGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this qtree get default response has a 2xx status code
func (o *QtreeGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qtree get default response has a 3xx status code
func (o *QtreeGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qtree get default response has a 4xx status code
func (o *QtreeGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qtree get default response has a 5xx status code
func (o *QtreeGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qtree get default response a status code equal to that given
func (o *QtreeGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the qtree get default response
func (o *QtreeGetDefault) Code() int {
	return o._statusCode
}

func (o *QtreeGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qtrees/{volume.uuid}/{id}][%d] qtree_get default %s", o._statusCode, payload)
}

func (o *QtreeGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/qtrees/{volume.uuid}/{id}][%d] qtree_get default %s", o._statusCode, payload)
}

func (o *QtreeGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QtreeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
