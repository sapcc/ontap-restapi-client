// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// IPServicePolicyModifyReader is a Reader for the IPServicePolicyModify structure.
type IPServicePolicyModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPServicePolicyModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPServicePolicyModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIPServicePolicyModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIPServicePolicyModifyOK creates a IPServicePolicyModifyOK with default headers values
func NewIPServicePolicyModifyOK() *IPServicePolicyModifyOK {
	return &IPServicePolicyModifyOK{}
}

/*
IPServicePolicyModifyOK describes a response with status code 200, with default header values.

OK
*/
type IPServicePolicyModifyOK struct {
}

// IsSuccess returns true when this ip service policy modify o k response has a 2xx status code
func (o *IPServicePolicyModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ip service policy modify o k response has a 3xx status code
func (o *IPServicePolicyModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ip service policy modify o k response has a 4xx status code
func (o *IPServicePolicyModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ip service policy modify o k response has a 5xx status code
func (o *IPServicePolicyModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ip service policy modify o k response a status code equal to that given
func (o *IPServicePolicyModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ip service policy modify o k response
func (o *IPServicePolicyModifyOK) Code() int {
	return 200
}

func (o *IPServicePolicyModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /network/ip/service-policies/{uuid}][%d] ipServicePolicyModifyOK", 200)
}

func (o *IPServicePolicyModifyOK) String() string {
	return fmt.Sprintf("[PATCH /network/ip/service-policies/{uuid}][%d] ipServicePolicyModifyOK", 200)
}

func (o *IPServicePolicyModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIPServicePolicyModifyDefault creates a IPServicePolicyModifyDefault with default headers values
func NewIPServicePolicyModifyDefault(code int) *IPServicePolicyModifyDefault {
	return &IPServicePolicyModifyDefault{
		_statusCode: code,
	}
}

/*
	IPServicePolicyModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1376669 | Port must reside in the same IPspace as the interface's SVM. |
| 2621740 | An unexpected error when trying to determine whether the target Vserver was locked or not on this cluster. |
| 53281911 | Modifying the service policy is disallowed because policies on this SVM are managed by the system |
| 53281929 | Service policies cannot combine block and file services. |
| 53281930 | Service policies maintained by the system cannot be renamed. |
| 53281931 | Service policy names cannot start with "default-". |
| 53281932 | Service cannot be added because the service does not exist for the specified SVM or IPspace. |
| 53281933 | A Cluster-scoped service cannot be added to a SVM-scoped service policy. |
| 53281934 | An SVM-scoped service cannot be added to a Cluster-scoped service policy. |
| 53281952 | The service policy on an SVM cannot be updated to include a block service. Use built-in service policies for SAN services. |
| 53281953 | The service policy on an SVM cannot be updated to include a new service. |
| 53281960 | Service cannot be removed from the service policy because it is used by one or more interfaces. |
| 53281961 | Service cannot be removed from the service policy because it is used by one or more interfaces. |
| 53281963 | Service cannot be removed from the service policy because it is used by one or more interfaces. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IPServicePolicyModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ip service policy modify default response has a 2xx status code
func (o *IPServicePolicyModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ip service policy modify default response has a 3xx status code
func (o *IPServicePolicyModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ip service policy modify default response has a 4xx status code
func (o *IPServicePolicyModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ip service policy modify default response has a 5xx status code
func (o *IPServicePolicyModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ip service policy modify default response a status code equal to that given
func (o *IPServicePolicyModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ip service policy modify default response
func (o *IPServicePolicyModifyDefault) Code() int {
	return o._statusCode
}

func (o *IPServicePolicyModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ip/service-policies/{uuid}][%d] ip_service_policy_modify default %s", o._statusCode, payload)
}

func (o *IPServicePolicyModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ip/service-policies/{uuid}][%d] ip_service_policy_modify default %s", o._statusCode, payload)
}

func (o *IPServicePolicyModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IPServicePolicyModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
