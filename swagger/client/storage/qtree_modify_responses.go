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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// QtreeModifyReader is a Reader for the QtreeModify structure.
type QtreeModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QtreeModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQtreeModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewQtreeModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQtreeModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQtreeModifyOK creates a QtreeModifyOK with default headers values
func NewQtreeModifyOK() *QtreeModifyOK {
	return &QtreeModifyOK{}
}

/*
QtreeModifyOK describes a response with status code 200, with default header values.

OK
*/
type QtreeModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this qtree modify o k response has a 2xx status code
func (o *QtreeModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qtree modify o k response has a 3xx status code
func (o *QtreeModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qtree modify o k response has a 4xx status code
func (o *QtreeModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this qtree modify o k response has a 5xx status code
func (o *QtreeModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this qtree modify o k response a status code equal to that given
func (o *QtreeModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the qtree modify o k response
func (o *QtreeModifyOK) Code() int {
	return 200
}

func (o *QtreeModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qtrees/{volume.uuid}/{id}][%d] qtreeModifyOK %s", 200, payload)
}

func (o *QtreeModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qtrees/{volume.uuid}/{id}][%d] qtreeModifyOK %s", 200, payload)
}

func (o *QtreeModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QtreeModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQtreeModifyAccepted creates a QtreeModifyAccepted with default headers values
func NewQtreeModifyAccepted() *QtreeModifyAccepted {
	return &QtreeModifyAccepted{}
}

/*
QtreeModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QtreeModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this qtree modify accepted response has a 2xx status code
func (o *QtreeModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qtree modify accepted response has a 3xx status code
func (o *QtreeModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qtree modify accepted response has a 4xx status code
func (o *QtreeModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this qtree modify accepted response has a 5xx status code
func (o *QtreeModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this qtree modify accepted response a status code equal to that given
func (o *QtreeModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the qtree modify accepted response
func (o *QtreeModifyAccepted) Code() int {
	return 202
}

func (o *QtreeModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qtrees/{volume.uuid}/{id}][%d] qtreeModifyAccepted %s", 202, payload)
}

func (o *QtreeModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qtrees/{volume.uuid}/{id}][%d] qtreeModifyAccepted %s", 202, payload)
}

func (o *QtreeModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QtreeModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQtreeModifyDefault creates a QtreeModifyDefault with default headers values
func NewQtreeModifyDefault(code int) *QtreeModifyDefault {
	return &QtreeModifyDefault{
		_statusCode: code,
	}
}

/*
	QtreeModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262196 | The request contains a field which cannot be set in this operation. |
| 262278 | Required field is missing in the request. |
| 917505 | Vserver not found. |
| 917525 | The specified volume does not exist in Vserver. |
| 918235 | A volume with UUID was not found. |
| 5242887 | Failed to modify qtree. |
| 5242897 | This operation is not permitted on read-only volume. |
| 5242898 | This operation is only permitted on a data Vserver. |
| 5242902 | Missing inputs. |
| 5242915 | Failed to assign qtree export policy to qtree. |
| 5242927 | Unable to find qtree. |
| 5242945 | Failed to modify qtree. |
| 5242951 | Export policy supplied does not belong to the specified export policy ID. |
| 5242954 | Failed to get the qtree from volume. |
| 5242955 | The UUID of the volume is required. |
| 5242956 | Failed to obtain a qtree with ID. |
| 5242957 | Failed to delete the qtree. |
| 5242958 | Failed to rename the qtree with ID in the volume and SVM. |
| 5242959 | Successfully renamed qtree but the modify operation failed. |
| 5242965 | Invalid qtree path. The volume name component of the qtree path, must be the same as the volume specified with the parameter. |
| 5242967 | UNIX user or group ID must be 32-bit unsigned integer. |
| 5242971 | Qtree was renamed. However, the path modification failed. |
| 5242972 | Cannot rename qtree as that name already exists on a volume in the Vserver. |
| 5242973 | Cannot rename qtree to name with path concurrently on volume in Vserver, unless non-root qtrees in enabled on the volume. |
| 5242974 | Moved qtree. However, other properties were not modified. |
| 5242975 | Renamed qtree and moved the qtree. However, other properties were not modified. |
| 6622064 | Security-style NTFS is not supported on a SnapMirror active sync relationship volume. |
| 8454348 | QoS on qtrees is not supported because not all nodes in the cluster can support it. |
| 9437324 | The security style unified is not supported. |
| 23724050 | Failed to resolve user or group name. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type QtreeModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this qtree modify default response has a 2xx status code
func (o *QtreeModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qtree modify default response has a 3xx status code
func (o *QtreeModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qtree modify default response has a 4xx status code
func (o *QtreeModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qtree modify default response has a 5xx status code
func (o *QtreeModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qtree modify default response a status code equal to that given
func (o *QtreeModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the qtree modify default response
func (o *QtreeModifyDefault) Code() int {
	return o._statusCode
}

func (o *QtreeModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qtrees/{volume.uuid}/{id}][%d] qtree_modify default %s", o._statusCode, payload)
}

func (o *QtreeModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/qtrees/{volume.uuid}/{id}][%d] qtree_modify default %s", o._statusCode, payload)
}

func (o *QtreeModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QtreeModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
