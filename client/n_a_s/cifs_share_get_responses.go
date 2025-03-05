// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// CifsShareGetReader is a Reader for the CifsShareGet structure.
type CifsShareGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsShareGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareGetOK creates a CifsShareGetOK with default headers values
func NewCifsShareGetOK() *CifsShareGetOK {
	return &CifsShareGetOK{}
}

/*
CifsShareGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsShareGetOK struct {
	Payload *models.CifsShare
}

// IsSuccess returns true when this cifs share get o k response has a 2xx status code
func (o *CifsShareGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs share get o k response has a 3xx status code
func (o *CifsShareGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs share get o k response has a 4xx status code
func (o *CifsShareGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs share get o k response has a 5xx status code
func (o *CifsShareGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs share get o k response a status code equal to that given
func (o *CifsShareGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs share get o k response
func (o *CifsShareGetOK) Code() int {
	return 200
}

func (o *CifsShareGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifsShareGetOK %s", 200, payload)
}

func (o *CifsShareGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifsShareGetOK %s", 200, payload)
}

func (o *CifsShareGetOK) GetPayload() *models.CifsShare {
	return o.Payload
}

func (o *CifsShareGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsShare)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsShareGetDefault creates a CifsShareGetDefault with default headers values
func NewCifsShareGetDefault(code int) *CifsShareGetDefault {
	return &CifsShareGetDefault{
		_statusCode: code,
	}
}

/*
CifsShareGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsShareGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs share get default response has a 2xx status code
func (o *CifsShareGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs share get default response has a 3xx status code
func (o *CifsShareGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs share get default response has a 4xx status code
func (o *CifsShareGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs share get default response has a 5xx status code
func (o *CifsShareGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs share get default response a status code equal to that given
func (o *CifsShareGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs share get default response
func (o *CifsShareGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifs_share_get default %s", o._statusCode, payload)
}

func (o *CifsShareGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/shares/{svm.uuid}/{name}][%d] cifs_share_get default %s", o._statusCode, payload)
}

func (o *CifsShareGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
