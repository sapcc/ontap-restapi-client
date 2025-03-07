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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// CifsSessionCollectionGetReader is a Reader for the CifsSessionCollectionGet structure.
type CifsSessionCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSessionCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsSessionCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSessionCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSessionCollectionGetOK creates a CifsSessionCollectionGetOK with default headers values
func NewCifsSessionCollectionGetOK() *CifsSessionCollectionGetOK {
	return &CifsSessionCollectionGetOK{}
}

/*
CifsSessionCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsSessionCollectionGetOK struct {
	Payload *models.CifsSessionResponse
}

// IsSuccess returns true when this cifs session collection get o k response has a 2xx status code
func (o *CifsSessionCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs session collection get o k response has a 3xx status code
func (o *CifsSessionCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs session collection get o k response has a 4xx status code
func (o *CifsSessionCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs session collection get o k response has a 5xx status code
func (o *CifsSessionCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs session collection get o k response a status code equal to that given
func (o *CifsSessionCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs session collection get o k response
func (o *CifsSessionCollectionGetOK) Code() int {
	return 200
}

func (o *CifsSessionCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/sessions][%d] cifsSessionCollectionGetOK %s", 200, payload)
}

func (o *CifsSessionCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/sessions][%d] cifsSessionCollectionGetOK %s", 200, payload)
}

func (o *CifsSessionCollectionGetOK) GetPayload() *models.CifsSessionResponse {
	return o.Payload
}

func (o *CifsSessionCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsSessionCollectionGetDefault creates a CifsSessionCollectionGetDefault with default headers values
func NewCifsSessionCollectionGetDefault(code int) *CifsSessionCollectionGetDefault {
	return &CifsSessionCollectionGetDefault{
		_statusCode: code,
	}
}

/*
CifsSessionCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsSessionCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs session collection get default response has a 2xx status code
func (o *CifsSessionCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs session collection get default response has a 3xx status code
func (o *CifsSessionCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs session collection get default response has a 4xx status code
func (o *CifsSessionCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs session collection get default response has a 5xx status code
func (o *CifsSessionCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs session collection get default response a status code equal to that given
func (o *CifsSessionCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs session collection get default response
func (o *CifsSessionCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *CifsSessionCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/sessions][%d] cifs_session_collection_get default %s", o._statusCode, payload)
}

func (o *CifsSessionCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/sessions][%d] cifs_session_collection_get default %s", o._statusCode, payload)
}

func (o *CifsSessionCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSessionCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
