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

// EmsFiltersRulesCreateReader is a Reader for the EmsFiltersRulesCreate structure.
type EmsFiltersRulesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsFiltersRulesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewEmsFiltersRulesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsFiltersRulesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsFiltersRulesCreateCreated creates a EmsFiltersRulesCreateCreated with default headers values
func NewEmsFiltersRulesCreateCreated() *EmsFiltersRulesCreateCreated {
	return &EmsFiltersRulesCreateCreated{}
}

/*
EmsFiltersRulesCreateCreated describes a response with status code 201, with default header values.

Created
*/
type EmsFiltersRulesCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.EmsFilterRuleResponse
}

// IsSuccess returns true when this ems filters rules create created response has a 2xx status code
func (o *EmsFiltersRulesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems filters rules create created response has a 3xx status code
func (o *EmsFiltersRulesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems filters rules create created response has a 4xx status code
func (o *EmsFiltersRulesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems filters rules create created response has a 5xx status code
func (o *EmsFiltersRulesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ems filters rules create created response a status code equal to that given
func (o *EmsFiltersRulesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ems filters rules create created response
func (o *EmsFiltersRulesCreateCreated) Code() int {
	return 201
}

func (o *EmsFiltersRulesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/filters/{name}/rules][%d] emsFiltersRulesCreateCreated %s", 201, payload)
}

func (o *EmsFiltersRulesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/filters/{name}/rules][%d] emsFiltersRulesCreateCreated %s", 201, payload)
}

func (o *EmsFiltersRulesCreateCreated) GetPayload() *models.EmsFilterRuleResponse {
	return o.Payload
}

func (o *EmsFiltersRulesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.EmsFilterRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsFiltersRulesCreateDefault creates a EmsFiltersRulesCreateDefault with default headers values
func NewEmsFiltersRulesCreateDefault(code int) *EmsFiltersRulesCreateDefault {
	return &EmsFiltersRulesCreateDefault{
		_statusCode: code,
	}
}

/*
	EmsFiltersRulesCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 983092 | The index of the rule provided is outside the allowed range for the filter provided |
| 983095 | The rule index provided is invalid for the filter provided |
| 983113 | Default filters cannot be modified or removed |
| 983115 | The maximum number of filter rules is reached |
| 983126 | A rule requires at least one name_pattern, severities, snmp_trap_types, or parameter pattern to be defined |
| 983127 | A property cannot contain a combination of the wildcard characters and other values |
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
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type EmsFiltersRulesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ems filters rules create default response has a 2xx status code
func (o *EmsFiltersRulesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems filters rules create default response has a 3xx status code
func (o *EmsFiltersRulesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems filters rules create default response has a 4xx status code
func (o *EmsFiltersRulesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems filters rules create default response has a 5xx status code
func (o *EmsFiltersRulesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems filters rules create default response a status code equal to that given
func (o *EmsFiltersRulesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ems filters rules create default response
func (o *EmsFiltersRulesCreateDefault) Code() int {
	return o._statusCode
}

func (o *EmsFiltersRulesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/filters/{name}/rules][%d] ems_filters_rules_create default %s", o._statusCode, payload)
}

func (o *EmsFiltersRulesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support/ems/filters/{name}/rules][%d] ems_filters_rules_create default %s", o._statusCode, payload)
}

func (o *EmsFiltersRulesCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsFiltersRulesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
