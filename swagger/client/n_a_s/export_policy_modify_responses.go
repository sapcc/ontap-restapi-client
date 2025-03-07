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

// ExportPolicyModifyReader is a Reader for the ExportPolicyModify structure.
type ExportPolicyModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportPolicyModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportPolicyModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportPolicyModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportPolicyModifyOK creates a ExportPolicyModifyOK with default headers values
func NewExportPolicyModifyOK() *ExportPolicyModifyOK {
	return &ExportPolicyModifyOK{}
}

/*
ExportPolicyModifyOK describes a response with status code 200, with default header values.

OK
*/
type ExportPolicyModifyOK struct {
}

// IsSuccess returns true when this export policy modify o k response has a 2xx status code
func (o *ExportPolicyModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export policy modify o k response has a 3xx status code
func (o *ExportPolicyModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export policy modify o k response has a 4xx status code
func (o *ExportPolicyModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export policy modify o k response has a 5xx status code
func (o *ExportPolicyModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export policy modify o k response a status code equal to that given
func (o *ExportPolicyModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export policy modify o k response
func (o *ExportPolicyModifyOK) Code() int {
	return 200
}

func (o *ExportPolicyModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/export-policies/{id}][%d] exportPolicyModifyOK", 200)
}

func (o *ExportPolicyModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/export-policies/{id}][%d] exportPolicyModifyOK", 200)
}

func (o *ExportPolicyModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportPolicyModifyDefault creates a ExportPolicyModifyDefault with default headers values
func NewExportPolicyModifyDefault(code int) *ExportPolicyModifyDefault {
	return &ExportPolicyModifyDefault{
		_statusCode: code,
	}
}

/*
	ExportPolicyModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1703950    | Failed to rename ruleset |
| 1703952    | Invalid ruleset name provided. No spaces are allowed in a ruleset name|
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
| 1704047    | The export policy name cannot be longer than 256 characters |
| 1704050    | Invalid clientmatch: clientmatch list contains a duplicate string. Duplicate strings in a clientmatch list are not supported |
| 1704054    | Invalid clientmatch: invalid characters in netgroup name. Valid characters for a netgroup name are 0-9, A-Z, a-z, ".", "_" and "-" |
| 1704064    | Clientmatch host name too long |
| 1704065    | Clientmatch domain name too long |
| 3277000    | Upgrade all nodes to ONTAP 9.0.0 or above to use krb5p as a security flavor in export-policy rules |
| 3277083    | User ID is not valid. Enter a value for User ID from 0 to 4294967295 |
| 3277149    | The "Anon" field cannot be an empty string |
*/
type ExportPolicyModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this export policy modify default response has a 2xx status code
func (o *ExportPolicyModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export policy modify default response has a 3xx status code
func (o *ExportPolicyModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export policy modify default response has a 4xx status code
func (o *ExportPolicyModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export policy modify default response has a 5xx status code
func (o *ExportPolicyModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export policy modify default response a status code equal to that given
func (o *ExportPolicyModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export policy modify default response
func (o *ExportPolicyModifyDefault) Code() int {
	return o._statusCode
}

func (o *ExportPolicyModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/export-policies/{id}][%d] export_policy_modify default %s", o._statusCode, payload)
}

func (o *ExportPolicyModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/export-policies/{id}][%d] export_policy_modify default %s", o._statusCode, payload)
}

func (o *ExportPolicyModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportPolicyModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
