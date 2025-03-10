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

// NfsClientsCacheModifyReader is a Reader for the NfsClientsCacheModify structure.
type NfsClientsCacheModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsClientsCacheModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsClientsCacheModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsClientsCacheModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsClientsCacheModifyOK creates a NfsClientsCacheModifyOK with default headers values
func NewNfsClientsCacheModifyOK() *NfsClientsCacheModifyOK {
	return &NfsClientsCacheModifyOK{}
}

/*
NfsClientsCacheModifyOK describes a response with status code 200, with default header values.

OK
*/
type NfsClientsCacheModifyOK struct {
}

// IsSuccess returns true when this nfs clients cache modify o k response has a 2xx status code
func (o *NfsClientsCacheModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nfs clients cache modify o k response has a 3xx status code
func (o *NfsClientsCacheModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nfs clients cache modify o k response has a 4xx status code
func (o *NfsClientsCacheModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nfs clients cache modify o k response has a 5xx status code
func (o *NfsClientsCacheModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nfs clients cache modify o k response a status code equal to that given
func (o *NfsClientsCacheModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nfs clients cache modify o k response
func (o *NfsClientsCacheModifyOK) Code() int {
	return 200
}

func (o *NfsClientsCacheModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/connected-client-settings][%d] nfsClientsCacheModifyOK", 200)
}

func (o *NfsClientsCacheModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/connected-client-settings][%d] nfsClientsCacheModifyOK", 200)
}

func (o *NfsClientsCacheModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNfsClientsCacheModifyDefault creates a NfsClientsCacheModifyDefault with default headers values
func NewNfsClientsCacheModifyDefault(code int) *NfsClientsCacheModifyDefault {
	return &NfsClientsCacheModifyDefault{
		_statusCode: code,
	}
}

/*
	NfsClientsCacheModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 3277172    | The value for \"client_retention_interval\" must be between 12 and 168 hours or its equivalent in days, minutes or seconds. |
| 3277170    | The \"client_retention_interval\" must be set to a value which is a multiple of 12 hours or to its equivalent in days, minutes or seconds.|
*/
type NfsClientsCacheModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nfs clients cache modify default response has a 2xx status code
func (o *NfsClientsCacheModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nfs clients cache modify default response has a 3xx status code
func (o *NfsClientsCacheModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nfs clients cache modify default response has a 4xx status code
func (o *NfsClientsCacheModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nfs clients cache modify default response has a 5xx status code
func (o *NfsClientsCacheModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nfs clients cache modify default response a status code equal to that given
func (o *NfsClientsCacheModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nfs clients cache modify default response
func (o *NfsClientsCacheModifyDefault) Code() int {
	return o._statusCode
}

func (o *NfsClientsCacheModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/connected-client-settings][%d] nfs_clients_cache_modify default %s", o._statusCode, payload)
}

func (o *NfsClientsCacheModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/connected-client-settings][%d] nfs_clients_cache_modify default %s", o._statusCode, payload)
}

func (o *NfsClientsCacheModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsClientsCacheModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
