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

	"github.com/sapcc/ontap-restapi-client/models"
)

// ExportRuleClientsCreateReader is a Reader for the ExportRuleClientsCreate structure.
type ExportRuleClientsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRuleClientsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExportRuleClientsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportRuleClientsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportRuleClientsCreateCreated creates a ExportRuleClientsCreateCreated with default headers values
func NewExportRuleClientsCreateCreated() *ExportRuleClientsCreateCreated {
	return &ExportRuleClientsCreateCreated{}
}

/*
ExportRuleClientsCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ExportRuleClientsCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.ExportClientResponse
}

// IsSuccess returns true when this export rule clients create created response has a 2xx status code
func (o *ExportRuleClientsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export rule clients create created response has a 3xx status code
func (o *ExportRuleClientsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export rule clients create created response has a 4xx status code
func (o *ExportRuleClientsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this export rule clients create created response has a 5xx status code
func (o *ExportRuleClientsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this export rule clients create created response a status code equal to that given
func (o *ExportRuleClientsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the export rule clients create created response
func (o *ExportRuleClientsCreateCreated) Code() int {
	return 201
}

func (o *ExportRuleClientsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients][%d] exportRuleClientsCreateCreated %s", 201, payload)
}

func (o *ExportRuleClientsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients][%d] exportRuleClientsCreateCreated %s", 201, payload)
}

func (o *ExportRuleClientsCreateCreated) GetPayload() *models.ExportClientResponse {
	return o.Payload
}

func (o *ExportRuleClientsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.ExportClientResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportRuleClientsCreateDefault creates a ExportRuleClientsCreateDefault with default headers values
func NewExportRuleClientsCreateDefault(code int) *ExportRuleClientsCreateDefault {
	return &ExportRuleClientsCreateDefault{
		_statusCode: code,
	}
}

/*
	ExportRuleClientsCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262196     | Field 'svm.name' is not supported in the body of POST request |
| 1703954    | Export policy does not exist |
| 1704036    | Invalid clientmatch:  missing domain name |
| 1704037    | Invalid clientmatch:  missing network name |
| 1704038    | Invalid clientmatch:  missing netgroup name |
| 1704039    | Invalid clientmatch |
| 1704040    | Invalid clientmatch: address bytes masked out by netmask are non-zero |
| 1704041    | Invalid clientmatch: address bytes masked to zero by netmask |
| 1704042    | Invalid clientmatch: too many bits in netmask |
| 1704043    | Invalid clientmatch: invalid netmask |
| 1704044    | Invalid clientmatch: invalid characters in host name |
| 1704045    | Invalid clientmatch: invalid characters in domain name |
| 1704050    | Invalid clientmatch: the clientmatch list contains a duplicate string. Duplicate strings in a clientmatch list are not supported |
| 1704054    | Invalid clientmatch: invalid characters in netgroup name. Valid characters for a netgroup name are 0-9, A-Z, a-z, ".", "_" and "-" |
| 1704064    | Clientmatch host name too long |
| 1704065    | Clientmatch domain name too long |
| 6691623    | User is not authorized |
*/
type ExportRuleClientsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this export rule clients create default response has a 2xx status code
func (o *ExportRuleClientsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export rule clients create default response has a 3xx status code
func (o *ExportRuleClientsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export rule clients create default response has a 4xx status code
func (o *ExportRuleClientsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export rule clients create default response has a 5xx status code
func (o *ExportRuleClientsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export rule clients create default response a status code equal to that given
func (o *ExportRuleClientsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export rule clients create default response
func (o *ExportRuleClientsCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExportRuleClientsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients][%d] export_rule_clients_create default %s", o._statusCode, payload)
}

func (o *ExportRuleClientsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/export-policies/{policy.id}/rules/{index}/clients][%d] export_rule_clients_create default %s", o._statusCode, payload)
}

func (o *ExportRuleClientsCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportRuleClientsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
