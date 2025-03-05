// Code generated by go-swagger; DO NOT EDIT.

package snap_mirror

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

// SnapmirrorRelationshipDeleteReader is a Reader for the SnapmirrorRelationshipDelete structure.
type SnapmirrorRelationshipDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapmirrorRelationshipDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapmirrorRelationshipDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnapmirrorRelationshipDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapmirrorRelationshipDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapmirrorRelationshipDeleteOK creates a SnapmirrorRelationshipDeleteOK with default headers values
func NewSnapmirrorRelationshipDeleteOK() *SnapmirrorRelationshipDeleteOK {
	return &SnapmirrorRelationshipDeleteOK{}
}

/*
SnapmirrorRelationshipDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SnapmirrorRelationshipDeleteOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapmirror relationship delete o k response has a 2xx status code
func (o *SnapmirrorRelationshipDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror relationship delete o k response has a 3xx status code
func (o *SnapmirrorRelationshipDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror relationship delete o k response has a 4xx status code
func (o *SnapmirrorRelationshipDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror relationship delete o k response has a 5xx status code
func (o *SnapmirrorRelationshipDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror relationship delete o k response a status code equal to that given
func (o *SnapmirrorRelationshipDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapmirror relationship delete o k response
func (o *SnapmirrorRelationshipDeleteOK) Code() int {
	return 200
}

func (o *SnapmirrorRelationshipDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapmirror/relationships/{uuid}][%d] snapmirrorRelationshipDeleteOK %s", 200, payload)
}

func (o *SnapmirrorRelationshipDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapmirror/relationships/{uuid}][%d] snapmirrorRelationshipDeleteOK %s", 200, payload)
}

func (o *SnapmirrorRelationshipDeleteOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorRelationshipDeleteAccepted creates a SnapmirrorRelationshipDeleteAccepted with default headers values
func NewSnapmirrorRelationshipDeleteAccepted() *SnapmirrorRelationshipDeleteAccepted {
	return &SnapmirrorRelationshipDeleteAccepted{}
}

/*
SnapmirrorRelationshipDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapmirrorRelationshipDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapmirror relationship delete accepted response has a 2xx status code
func (o *SnapmirrorRelationshipDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror relationship delete accepted response has a 3xx status code
func (o *SnapmirrorRelationshipDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror relationship delete accepted response has a 4xx status code
func (o *SnapmirrorRelationshipDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror relationship delete accepted response has a 5xx status code
func (o *SnapmirrorRelationshipDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror relationship delete accepted response a status code equal to that given
func (o *SnapmirrorRelationshipDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snapmirror relationship delete accepted response
func (o *SnapmirrorRelationshipDeleteAccepted) Code() int {
	return 202
}

func (o *SnapmirrorRelationshipDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapmirror/relationships/{uuid}][%d] snapmirrorRelationshipDeleteAccepted %s", 202, payload)
}

func (o *SnapmirrorRelationshipDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapmirror/relationships/{uuid}][%d] snapmirrorRelationshipDeleteAccepted %s", 202, payload)
}

func (o *SnapmirrorRelationshipDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorRelationshipDeleteDefault creates a SnapmirrorRelationshipDeleteDefault with default headers values
func NewSnapmirrorRelationshipDeleteDefault(code int) *SnapmirrorRelationshipDeleteDefault {
	return &SnapmirrorRelationshipDeleteDefault{
		_statusCode: code,
	}
}

/*
	SnapmirrorRelationshipDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 13303825    | Could not retrieve information for the SnapMirror policy type |
| 13303814    | Could not retrieve the source or destination SVM UUID |
| 13303815    | Could not retrieve information for the peer cluster |
| 13303822    | SnapMirror release has failed |
| 13303813    | SnapMirror release was successful but delete has failed |
| 13303854    | Cleanup of restore relationship failed |
| 13303855    | DELETE call on a restore relationship does not support the given flags |
| 13303865    | Deleting the specified SnapMirror policy is not supported. |
| 6619715     | Modification of relationship is in progress. Retry the command after a few minutes. |
*/
type SnapmirrorRelationshipDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapmirror relationship delete default response has a 2xx status code
func (o *SnapmirrorRelationshipDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapmirror relationship delete default response has a 3xx status code
func (o *SnapmirrorRelationshipDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapmirror relationship delete default response has a 4xx status code
func (o *SnapmirrorRelationshipDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapmirror relationship delete default response has a 5xx status code
func (o *SnapmirrorRelationshipDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapmirror relationship delete default response a status code equal to that given
func (o *SnapmirrorRelationshipDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapmirror relationship delete default response
func (o *SnapmirrorRelationshipDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SnapmirrorRelationshipDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapmirror/relationships/{uuid}][%d] snapmirror_relationship_delete default %s", o._statusCode, payload)
}

func (o *SnapmirrorRelationshipDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /snapmirror/relationships/{uuid}][%d] snapmirror_relationship_delete default %s", o._statusCode, payload)
}

func (o *SnapmirrorRelationshipDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
