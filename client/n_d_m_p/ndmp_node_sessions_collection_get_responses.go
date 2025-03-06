// Code generated by go-swagger; DO NOT EDIT.

package n_d_m_p

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

// NdmpNodeSessionsCollectionGetReader is a Reader for the NdmpNodeSessionsCollectionGet structure.
type NdmpNodeSessionsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NdmpNodeSessionsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNdmpNodeSessionsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNdmpNodeSessionsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNdmpNodeSessionsCollectionGetOK creates a NdmpNodeSessionsCollectionGetOK with default headers values
func NewNdmpNodeSessionsCollectionGetOK() *NdmpNodeSessionsCollectionGetOK {
	return &NdmpNodeSessionsCollectionGetOK{}
}

/*
NdmpNodeSessionsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NdmpNodeSessionsCollectionGetOK struct {
	Payload *models.NdmpSessionResponse
}

// IsSuccess returns true when this ndmp node sessions collection get o k response has a 2xx status code
func (o *NdmpNodeSessionsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ndmp node sessions collection get o k response has a 3xx status code
func (o *NdmpNodeSessionsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ndmp node sessions collection get o k response has a 4xx status code
func (o *NdmpNodeSessionsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ndmp node sessions collection get o k response has a 5xx status code
func (o *NdmpNodeSessionsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ndmp node sessions collection get o k response a status code equal to that given
func (o *NdmpNodeSessionsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ndmp node sessions collection get o k response
func (o *NdmpNodeSessionsCollectionGetOK) Code() int {
	return 200
}

func (o *NdmpNodeSessionsCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/sessions][%d] ndmpNodeSessionsCollectionGetOK %s", 200, payload)
}

func (o *NdmpNodeSessionsCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/sessions][%d] ndmpNodeSessionsCollectionGetOK %s", 200, payload)
}

func (o *NdmpNodeSessionsCollectionGetOK) GetPayload() *models.NdmpSessionResponse {
	return o.Payload
}

func (o *NdmpNodeSessionsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NdmpSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNdmpNodeSessionsCollectionGetDefault creates a NdmpNodeSessionsCollectionGetDefault with default headers values
func NewNdmpNodeSessionsCollectionGetDefault(code int) *NdmpNodeSessionsCollectionGetDefault {
	return &NdmpNodeSessionsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	NdmpNodeSessionsCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 68812805    | Failed to obtain the NDMP mode of the operation.|
*/
type NdmpNodeSessionsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ndmp node sessions collection get default response has a 2xx status code
func (o *NdmpNodeSessionsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ndmp node sessions collection get default response has a 3xx status code
func (o *NdmpNodeSessionsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ndmp node sessions collection get default response has a 4xx status code
func (o *NdmpNodeSessionsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ndmp node sessions collection get default response has a 5xx status code
func (o *NdmpNodeSessionsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ndmp node sessions collection get default response a status code equal to that given
func (o *NdmpNodeSessionsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ndmp node sessions collection get default response
func (o *NdmpNodeSessionsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *NdmpNodeSessionsCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/sessions][%d] ndmp_node_sessions_collection_get default %s", o._statusCode, payload)
}

func (o *NdmpNodeSessionsCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/ndmp/sessions][%d] ndmp_node_sessions_collection_get default %s", o._statusCode, payload)
}

func (o *NdmpNodeSessionsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NdmpNodeSessionsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
