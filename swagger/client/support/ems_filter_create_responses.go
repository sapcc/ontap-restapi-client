// Code generated by go-swagger; DO NOT EDIT.

package support

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

// EmsFilterCreateReader is a Reader for the EmsFilterCreate structure.
type EmsFilterCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsFilterCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEmsFilterCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsFilterCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsFilterCreateCreated creates a EmsFilterCreateCreated with default headers values
func NewEmsFilterCreateCreated() *EmsFilterCreateCreated {
	return &EmsFilterCreateCreated{}
}

/*
EmsFilterCreateCreated describes a response with status code 201, with default header values.

Created
*/
type EmsFilterCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.EmsFilterResponse
}

// IsSuccess returns true when this ems filter create created response has a 2xx status code
func (o *EmsFilterCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems filter create created response has a 3xx status code
func (o *EmsFilterCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems filter create created response has a 4xx status code
func (o *EmsFilterCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems filter create created response has a 5xx status code
func (o *EmsFilterCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ems filter create created response a status code equal to that given
func (o *EmsFilterCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ems filter create created response
func (o *EmsFilterCreateCreated) Code() int {
	return 201
}

func (o *EmsFilterCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/filters][%d] emsFilterCreateCreated %s", 201, payload)
}

func (o *EmsFilterCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/filters][%d] emsFilterCreateCreated %s", 201, payload)
}

func (o *EmsFilterCreateCreated) GetPayload() *models.EmsFilterResponse {
	return o.Payload
}

func (o *EmsFilterCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.EmsFilterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsFilterCreateDefault creates a EmsFilterCreateDefault with default headers values
func NewEmsFilterCreateDefault(code int) *EmsFilterCreateDefault {
	return &EmsFilterCreateDefault{
		_statusCode: code,
	}
}

/*
	EmsFilterCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 983088 | The filter name provided is empty |
| 983089 | The filter name provided cannot contain spaces |
| 983092 | The index of the rule provided is outside the allowed range for the filter provided |
| 983094 | The filter name provided is invalid. The filter name must contain between 2 and 64 characters and start and end with an alphanumeric symbol or (underscore). The allowed special characters are (underscore) and -(hyphen) |
| 983095 | The rule index provided is invalid for the filter provided |
| 983101 | No event is matched by the rule provided |
| 983113 | Default filters cannot be modified or removed |
| 983114 | The maximum number of filters is reached |
| 983115 | The maximum number of filter rules is reached |
| 983126 | A rule requires at least one name_pattern, severities, snmp_trap_types, or parameter pattern to be defined |
| 983127 | A property cannot contain a combination of the wildcard character and other values |
| 983128 | An invalid value is provided for the property 'snmp_trap_types' |
| 983146 | An invalid value is provided for the property 'severities' |
| 983147 | The severities provided are not supported |
| 983155 | The provided severities property does not match that of the name_pattern |
| 983156 | The provided snmp_trap_types property does not match that of the name_pattern |
| 983157 | The provided severities and snmp_trap_types properties do not match those of the name_pattern |
| 983158 | The name_pattern provided does not exist |
| 983195 | Empty field in parameter_criteria. Both name and value patterns must be specified |
| 983196 | name_pattern and value_pattern fields in parameter_criteria are empty |
| 983211 | Parameter criteria based filtering is not supported in this version of ONTAP |
| 983216 | The provided parameter criteria does not match that of the message name |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type EmsFilterCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems filter create default response has a 2xx status code
func (o *EmsFilterCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems filter create default response has a 3xx status code
func (o *EmsFilterCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems filter create default response has a 4xx status code
func (o *EmsFilterCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems filter create default response has a 5xx status code
func (o *EmsFilterCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems filter create default response a status code equal to that given
func (o *EmsFilterCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems filter create default response
func (o *EmsFilterCreateDefault) Code() int {
	return o._statusCode
}

func (o *EmsFilterCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/filters][%d] ems_filter_create default %s", o._statusCode, payload)
}

func (o *EmsFilterCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/filters][%d] ems_filter_create default %s", o._statusCode, payload)
}

func (o *EmsFilterCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsFilterCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
