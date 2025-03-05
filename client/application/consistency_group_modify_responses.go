// Code generated by go-swagger; DO NOT EDIT.

package application

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

// ConsistencyGroupModifyReader is a Reader for the ConsistencyGroupModify structure.
type ConsistencyGroupModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsistencyGroupModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsistencyGroupModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewConsistencyGroupModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsistencyGroupModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsistencyGroupModifyOK creates a ConsistencyGroupModifyOK with default headers values
func NewConsistencyGroupModifyOK() *ConsistencyGroupModifyOK {
	return &ConsistencyGroupModifyOK{}
}

/*
ConsistencyGroupModifyOK describes a response with status code 200, with default header values.

OK
*/
type ConsistencyGroupModifyOK struct {
}

// IsSuccess returns true when this consistency group modify o k response has a 2xx status code
func (o *ConsistencyGroupModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group modify o k response has a 3xx status code
func (o *ConsistencyGroupModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group modify o k response has a 4xx status code
func (o *ConsistencyGroupModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group modify o k response has a 5xx status code
func (o *ConsistencyGroupModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group modify o k response a status code equal to that given
func (o *ConsistencyGroupModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the consistency group modify o k response
func (o *ConsistencyGroupModifyOK) Code() int {
	return 200
}

func (o *ConsistencyGroupModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /application/consistency-groups/{uuid}][%d] consistencyGroupModifyOK", 200)
}

func (o *ConsistencyGroupModifyOK) String() string {
	return fmt.Sprintf("[PATCH /application/consistency-groups/{uuid}][%d] consistencyGroupModifyOK", 200)
}

func (o *ConsistencyGroupModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsistencyGroupModifyAccepted creates a ConsistencyGroupModifyAccepted with default headers values
func NewConsistencyGroupModifyAccepted() *ConsistencyGroupModifyAccepted {
	return &ConsistencyGroupModifyAccepted{}
}

/*
ConsistencyGroupModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ConsistencyGroupModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this consistency group modify accepted response has a 2xx status code
func (o *ConsistencyGroupModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this consistency group modify accepted response has a 3xx status code
func (o *ConsistencyGroupModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this consistency group modify accepted response has a 4xx status code
func (o *ConsistencyGroupModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this consistency group modify accepted response has a 5xx status code
func (o *ConsistencyGroupModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this consistency group modify accepted response a status code equal to that given
func (o *ConsistencyGroupModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the consistency group modify accepted response
func (o *ConsistencyGroupModifyAccepted) Code() int {
	return 202
}

func (o *ConsistencyGroupModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /application/consistency-groups/{uuid}][%d] consistencyGroupModifyAccepted %s", 202, payload)
}

func (o *ConsistencyGroupModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /application/consistency-groups/{uuid}][%d] consistencyGroupModifyAccepted %s", 202, payload)
}

func (o *ConsistencyGroupModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ConsistencyGroupModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsistencyGroupModifyDefault creates a ConsistencyGroupModifyDefault with default headers values
func NewConsistencyGroupModifyDefault(code int) *ConsistencyGroupModifyDefault {
	return &ConsistencyGroupModifyDefault{
		_statusCode: code,
	}
}

/*
	ConsistencyGroupModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262285 | Consistency group does not support removing elements using a PATCH request. |
| 2621761 | Consistency groups with DP volumes are not supported on storage-limit enabled SVM. |
| 53411842 | Consistency group does not exist. |
| 53411843 | A consistency group with specified UUID was not found. |
| 53411844 | Specified consistency group was not found in the specified SVM. |
| 53411845 | The specified UUID and name refer to different consistency groups. |
| 53411846 | Either name or UUID must be provided. |
| 53411852 | A consistency group with the same identifier in the same scope exists. |
| 53411853 | Fields provided in the request conflict with each other. |
| 53411856 | Field provided is only supported when provisioning new objects. |
| 53411857 | LUNs that are not members of the application are not supported by this API. LUNs can be added to an application by adding the volume containing the LUNs to the application. |
| 53411860 | An object with the same identifier in the same scope exists. |
| 53411861 | Volume specified does not exist in provided volume array. |
| 53411862 | Modifying existing igroups is not supported using this API. |
| 53411864 | Request content insufficient to add an existing volume to an application. |
| 53411865 | Volumes contained in one consistency group cannot be added to a different consistency group. |
| 53411866 | LUNs are not supported on FlexGroup volumes. |
| 53411867 | LUN name is too long after appending a unique suffix. |
| 53411869 | Volume name is too long after appending a unique suffix. |
| 53411870 | When using the \"round_robin\" layout, the volume count must not be greater than the LUN count. |
| 53411942 | The application or component type of a consistency group that has an associated SnapMirror relationship cannot be changed. |
| 53411959 | Volumes with snapshot locking enabled cannot be added to a consistency group. |
| 53412027 | Failed to update the snapshot policy because the snapshot policies are not supported on the destination consistency group of SnapMirror active sync relationships. |
| 53412056 | The consistency group is not a FlexClone. |
| 53412057 | Consistency group split operation failed. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ConsistencyGroupModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this consistency group modify default response has a 2xx status code
func (o *ConsistencyGroupModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this consistency group modify default response has a 3xx status code
func (o *ConsistencyGroupModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this consistency group modify default response has a 4xx status code
func (o *ConsistencyGroupModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this consistency group modify default response has a 5xx status code
func (o *ConsistencyGroupModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this consistency group modify default response a status code equal to that given
func (o *ConsistencyGroupModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the consistency group modify default response
func (o *ConsistencyGroupModifyDefault) Code() int {
	return o._statusCode
}

func (o *ConsistencyGroupModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /application/consistency-groups/{uuid}][%d] consistency_group_modify default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /application/consistency-groups/{uuid}][%d] consistency_group_modify default %s", o._statusCode, payload)
}

func (o *ConsistencyGroupModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConsistencyGroupModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
