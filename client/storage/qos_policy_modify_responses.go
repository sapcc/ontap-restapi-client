// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// QosPolicyModifyReader is a Reader for the QosPolicyModify structure.
type QosPolicyModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QosPolicyModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQosPolicyModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewQosPolicyModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQosPolicyModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQosPolicyModifyOK creates a QosPolicyModifyOK with default headers values
func NewQosPolicyModifyOK() *QosPolicyModifyOK {
	return &QosPolicyModifyOK{}
}

/*
QosPolicyModifyOK describes a response with status code 200, with default header values.

OK
*/
type QosPolicyModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this qos policy modify o k response has a 2xx status code
func (o *QosPolicyModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qos policy modify o k response has a 3xx status code
func (o *QosPolicyModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qos policy modify o k response has a 4xx status code
func (o *QosPolicyModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this qos policy modify o k response has a 5xx status code
func (o *QosPolicyModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this qos policy modify o k response a status code equal to that given
func (o *QosPolicyModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the qos policy modify o k response
func (o *QosPolicyModifyOK) Code() int {
	return 200
}

func (o *QosPolicyModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qos/policies/{uuid}][%d] qosPolicyModifyOK %s", 200, payload)
}

func (o *QosPolicyModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qos/policies/{uuid}][%d] qosPolicyModifyOK %s", 200, payload)
}

func (o *QosPolicyModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QosPolicyModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQosPolicyModifyAccepted creates a QosPolicyModifyAccepted with default headers values
func NewQosPolicyModifyAccepted() *QosPolicyModifyAccepted {
	return &QosPolicyModifyAccepted{}
}

/*
QosPolicyModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QosPolicyModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this qos policy modify accepted response has a 2xx status code
func (o *QosPolicyModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qos policy modify accepted response has a 3xx status code
func (o *QosPolicyModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qos policy modify accepted response has a 4xx status code
func (o *QosPolicyModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this qos policy modify accepted response has a 5xx status code
func (o *QosPolicyModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this qos policy modify accepted response a status code equal to that given
func (o *QosPolicyModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the qos policy modify accepted response
func (o *QosPolicyModifyAccepted) Code() int {
	return 202
}

func (o *QosPolicyModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qos/policies/{uuid}][%d] qosPolicyModifyAccepted %s", 202, payload)
}

func (o *QosPolicyModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qos/policies/{uuid}][%d] qosPolicyModifyAccepted %s", 202, payload)
}

func (o *QosPolicyModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QosPolicyModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQosPolicyModifyDefault creates a QosPolicyModifyDefault with default headers values
func NewQosPolicyModifyDefault(code int) *QosPolicyModifyDefault {
	return &QosPolicyModifyDefault{
		_statusCode: code,
	}
}

/*
	QosPolicyModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8454147 | The maximum limit for QoS policies has been reached. |
| 8454154 | The name specified for creating conflicts with an existing QoS policy name. |
| 8454194 | The minimum throughput value for the policy group must be less than or equal to the maximum throughput value. |
| 8454260 | Invalid value for maximum and minimum fields. Valid values for max_throughput_iops and max_throughput_mbps combination is for the ratio of max_throughput_mbps and max_throughput_iops to be within 1 to 4096. |
| 8454273 | Invalid value for an adaptive field. Value should be non-zero. |
| 8454277 | The name specified for creating an adaptive QoS policy conflicts with an existing fixed QoS policy name. |
| 8454278 | The name specified for creating a fixed QoS policy conflicts with an existing adaptive QoS policy name. |
| 8454286 | Modifications on these cluster scoped preset policies is prohibited. |
| 8454327 | The existing fixed QoS policy cannot be modified to an adaptive QoS policy. |
| 8454328 | The existing adaptive QoS policy cannot be modified to a fixed QoS policy. |
| 8454379 | The name specified for creating a fixed QoS policy already exists. |
| 8454380 | The name specified for creating an adaptive QoS policy already exists. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type QosPolicyModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this qos policy modify default response has a 2xx status code
func (o *QosPolicyModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qos policy modify default response has a 3xx status code
func (o *QosPolicyModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qos policy modify default response has a 4xx status code
func (o *QosPolicyModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qos policy modify default response has a 5xx status code
func (o *QosPolicyModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qos policy modify default response a status code equal to that given
func (o *QosPolicyModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the qos policy modify default response
func (o *QosPolicyModifyDefault) Code() int {
	return o._statusCode
}

func (o *QosPolicyModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qos/policies/{uuid}][%d] qos_policy_modify default %s", o._statusCode, payload)
}

func (o *QosPolicyModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qos/policies/{uuid}][%d] qos_policy_modify default %s", o._statusCode, payload)
}

func (o *QosPolicyModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QosPolicyModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
