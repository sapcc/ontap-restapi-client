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

// ExportRuleModifyReader is a Reader for the ExportRuleModify structure.
type ExportRuleModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRuleModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportRuleModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportRuleModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportRuleModifyOK creates a ExportRuleModifyOK with default headers values
func NewExportRuleModifyOK() *ExportRuleModifyOK {
	return &ExportRuleModifyOK{}
}

/*
ExportRuleModifyOK describes a response with status code 200, with default header values.

OK
*/
type ExportRuleModifyOK struct {
}

// IsSuccess returns true when this export rule modify o k response has a 2xx status code
func (o *ExportRuleModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export rule modify o k response has a 3xx status code
func (o *ExportRuleModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export rule modify o k response has a 4xx status code
func (o *ExportRuleModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export rule modify o k response has a 5xx status code
func (o *ExportRuleModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export rule modify o k response a status code equal to that given
func (o *ExportRuleModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export rule modify o k response
func (o *ExportRuleModifyOK) Code() int {
	return 200
}

func (o *ExportRuleModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/export-policies/{policy.id}/rules/{index}][%d] exportRuleModifyOK", 200)
}

func (o *ExportRuleModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/export-policies/{policy.id}/rules/{index}][%d] exportRuleModifyOK", 200)
}

func (o *ExportRuleModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportRuleModifyDefault creates a ExportRuleModifyDefault with default headers values
func NewExportRuleModifyDefault(code int) *ExportRuleModifyDefault {
	return &ExportRuleModifyDefault{
		_statusCode: code,
	}
}

/*
	ExportRuleModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262196     | Field 'svm.name' is not supported in the body of PATCH request |
| 262197     | The value provided is invalid for the field |
| 262203     | Field 'svm.uuid' is not supported in the body of PATCH request |
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
| 1704050    | Invalid clientmatch: clientmatch list contains a duplicate string. Duplicate strings in a clientmatch list are not supported |
| 1704054    | Invalid clientmatch: invalid characters in netgroup name. Valid characters for a netgroup name are 0-9, A-Z, a-z, ".", "_" and "-" |
| 1704064    | Clientmatch host name too long |
| 1704065    | Clientmatch domain name too long |
| 3277000    | Upgrade all nodes to ONTAP 9.0.0 or above to use krb5p as a security flavor in export-policy rules |
| 3277083    | User ID is not valid. Enter a value for User ID from 0 to 4294967295 |
| 3277149    | The "Anon" field cannot be an empty string |
| 6691623    | User is not authorized |
*/
type ExportRuleModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this export rule modify default response has a 2xx status code
func (o *ExportRuleModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export rule modify default response has a 3xx status code
func (o *ExportRuleModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export rule modify default response has a 4xx status code
func (o *ExportRuleModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export rule modify default response has a 5xx status code
func (o *ExportRuleModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export rule modify default response a status code equal to that given
func (o *ExportRuleModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export rule modify default response
func (o *ExportRuleModifyDefault) Code() int {
	return o._statusCode
}

func (o *ExportRuleModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/export-policies/{policy.id}/rules/{index}][%d] export_rule_modify default %s", o._statusCode, payload)
}

func (o *ExportRuleModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/export-policies/{policy.id}/rules/{index}][%d] export_rule_modify default %s", o._statusCode, payload)
}

func (o *ExportRuleModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportRuleModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
