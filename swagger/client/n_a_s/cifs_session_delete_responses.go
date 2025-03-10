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

// CifsSessionDeleteReader is a Reader for the CifsSessionDelete structure.
type CifsSessionDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSessionDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsSessionDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSessionDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSessionDeleteOK creates a CifsSessionDeleteOK with default headers values
func NewCifsSessionDeleteOK() *CifsSessionDeleteOK {
	return &CifsSessionDeleteOK{}
}

/*
CifsSessionDeleteOK describes a response with status code 200, with default header values.

OK
*/
type CifsSessionDeleteOK struct {
}

// IsSuccess returns true when this cifs session delete o k response has a 2xx status code
func (o *CifsSessionDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs session delete o k response has a 3xx status code
func (o *CifsSessionDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs session delete o k response has a 4xx status code
func (o *CifsSessionDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs session delete o k response has a 5xx status code
func (o *CifsSessionDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs session delete o k response a status code equal to that given
func (o *CifsSessionDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs session delete o k response
func (o *CifsSessionDeleteOK) Code() int {
	return 200
}

func (o *CifsSessionDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/sessions/{node.uuid}/{svm.uuid}/{identifier}/{connection_id}][%d] cifsSessionDeleteOK", 200)
}

func (o *CifsSessionDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/sessions/{node.uuid}/{svm.uuid}/{identifier}/{connection_id}][%d] cifsSessionDeleteOK", 200)
}

func (o *CifsSessionDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsSessionDeleteDefault creates a CifsSessionDeleteDefault with default headers values
func NewCifsSessionDeleteDefault(code int) *CifsSessionDeleteDefault {
	return &CifsSessionDeleteDefault{
		_statusCode: code,
	}
}

/*
	CifsSessionDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655503     | The SMB session delete does not allow a connection ID of zero (0). |
*/
type CifsSessionDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs session delete default response has a 2xx status code
func (o *CifsSessionDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs session delete default response has a 3xx status code
func (o *CifsSessionDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs session delete default response has a 4xx status code
func (o *CifsSessionDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs session delete default response has a 5xx status code
func (o *CifsSessionDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs session delete default response a status code equal to that given
func (o *CifsSessionDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs session delete default response
func (o *CifsSessionDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CifsSessionDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/sessions/{node.uuid}/{svm.uuid}/{identifier}/{connection_id}][%d] cifs_session_delete default %s", o._statusCode, payload)
}

func (o *CifsSessionDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/sessions/{node.uuid}/{svm.uuid}/{identifier}/{connection_id}][%d] cifs_session_delete default %s", o._statusCode, payload)
}

func (o *CifsSessionDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSessionDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
