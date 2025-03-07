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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// SnapmirrorRelationshipModifyReader is a Reader for the SnapmirrorRelationshipModify structure.
type SnapmirrorRelationshipModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapmirrorRelationshipModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapmirrorRelationshipModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnapmirrorRelationshipModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapmirrorRelationshipModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapmirrorRelationshipModifyOK creates a SnapmirrorRelationshipModifyOK with default headers values
func NewSnapmirrorRelationshipModifyOK() *SnapmirrorRelationshipModifyOK {
	return &SnapmirrorRelationshipModifyOK{}
}

/*
SnapmirrorRelationshipModifyOK describes a response with status code 200, with default header values.

OK
*/
type SnapmirrorRelationshipModifyOK struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapmirror relationship modify o k response has a 2xx status code
func (o *SnapmirrorRelationshipModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror relationship modify o k response has a 3xx status code
func (o *SnapmirrorRelationshipModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror relationship modify o k response has a 4xx status code
func (o *SnapmirrorRelationshipModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror relationship modify o k response has a 5xx status code
func (o *SnapmirrorRelationshipModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror relationship modify o k response a status code equal to that given
func (o *SnapmirrorRelationshipModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snapmirror relationship modify o k response
func (o *SnapmirrorRelationshipModifyOK) Code() int {
	return 200
}

func (o *SnapmirrorRelationshipModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/relationships/{uuid}][%d] snapmirrorRelationshipModifyOK %s", 200, payload)
}

func (o *SnapmirrorRelationshipModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/relationships/{uuid}][%d] snapmirrorRelationshipModifyOK %s", 200, payload)
}

func (o *SnapmirrorRelationshipModifyOK) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorRelationshipModifyAccepted creates a SnapmirrorRelationshipModifyAccepted with default headers values
func NewSnapmirrorRelationshipModifyAccepted() *SnapmirrorRelationshipModifyAccepted {
	return &SnapmirrorRelationshipModifyAccepted{}
}

/*
SnapmirrorRelationshipModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapmirrorRelationshipModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapmirror relationship modify accepted response has a 2xx status code
func (o *SnapmirrorRelationshipModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror relationship modify accepted response has a 3xx status code
func (o *SnapmirrorRelationshipModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror relationship modify accepted response has a 4xx status code
func (o *SnapmirrorRelationshipModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror relationship modify accepted response has a 5xx status code
func (o *SnapmirrorRelationshipModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror relationship modify accepted response a status code equal to that given
func (o *SnapmirrorRelationshipModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snapmirror relationship modify accepted response
func (o *SnapmirrorRelationshipModifyAccepted) Code() int {
	return 202
}

func (o *SnapmirrorRelationshipModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/relationships/{uuid}][%d] snapmirrorRelationshipModifyAccepted %s", 202, payload)
}

func (o *SnapmirrorRelationshipModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/relationships/{uuid}][%d] snapmirrorRelationshipModifyAccepted %s", 202, payload)
}

func (o *SnapmirrorRelationshipModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorRelationshipModifyDefault creates a SnapmirrorRelationshipModifyDefault with default headers values
func NewSnapmirrorRelationshipModifyDefault(code int) *SnapmirrorRelationshipModifyDefault {
	return &SnapmirrorRelationshipModifyDefault{
		_statusCode: code,
	}
}

/*
	SnapmirrorRelationshipModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 13303825    | Could not retrieve information for the SnapMirror policy type |
| 13303817    | Unknown value for the SnapMirror state |
| 13303829    | Invalid state |
| 13303830    | Transient state |
| 13303831    | Invalid state for async SnapMirror relationship |
| 13303834    | Given input valid only for FlexGroup SnapMirror relationship |
| 13303835    | Given flag is valid only when PATCH state is broken_off |
| 13303836    | Given flag is valid only when PATCH state is snapmirrored or in_sync |
| 13303818    | Invalid state transition requested |
| 13303828    | Given state change is not possible for SVM SnapMirror relationship |
| 13303833    | Requested state change is not possible |
| 13303832    | SnapMirror relationship is already initialized |
| 13303824    | Quiescing the SnapMirror relationship has failed |
| 13303826    | Required environment variables are not set |
| 13303827    | Internal Error |
| 13303823    | Quiesce operation timed out |
| 13303821    | Invalid SnapMirror policy name/UUID |
| 13303819    | Could not retrieve SnapMirror policy information |
| 13303851    | Cannot modify attributes of SnapMirror restore relationship |
| 13303816    | Could not retrieve state or status values |
| 13303837    | Given flags are valid only if SnapMirror state change is requested |
| 6619546     | Destination must be a dp volume |
| 13303808    | Transition to broken_off state failed |
| 13303809    | Transition to paused state failed |
| 13303810    | Transition to snapmirrored state failed |
| 13303811    | Transition from paused state failed |
| 13303820    | SnapMirror policy, transfer_schedule, and throttle, if specified were successfully updated, state transition failed |
| 13303856    | SVM is not configured with any data protocol |
| 13303857    | SVM is not configured with any network interface |
| 13303858    | Internal error. Failed to check LIF and protocols details for SVM |
| 13303859    | Internal error. SVM Failover operation failed. SVM operational state is unavailable. |
| 13303865    | Modifying the specified SnapMirror policy is not supported. |
| 13303866    | Cannot use the specified policy to modify the policy of the relationship. |
| 13303867    | Modifying the policy of an async-mirror or a vault relationship is not supported. |
| 13303884    | LIF and protocols details are configured incorrectly for SVM. |
| 13303996    | The source and destination clusters both have a policy with the same name, but they have different properties. |
| 13304062    | Cannot reverse the direction of a SnapMirror DP relationship when the source cluster version is earlier than the destination cluster version. |
| 13304070    | Remote peer cluster requires the dp_rest_support capability to support reversing the direction of a DP relationship. |
| 13304071    | Failed to access capabilities on remote cluster. |
| 13304080    | Specified uuid and name do not match. |
| 13304081    | Modifying a property during the operation is not supported. |
| 13304082    | The specified properties are mutually exclusive. |
| 13304083    | The specified property is not supported because all nodes in the cluster are not capable of supporting the property. |
| 13304086    | Reversing the direction of a SnapMirror relationship associated with a policy containing the property create_snapshot_on_source set to false is not supported. |
| 6619715     | Modification of relationship is in progress. Retry the command after a few minutes. |
| 6619699     | Schedule not found. |
| 13304108    | Schedule not found in the Administrative SVM or the SVM for the relationship. |
| 13304111    | The SnapMirror active sync relationship consistency groups are nested. Expanding an SnapMirror active sync relationship with a pre-existing DP volume is only supported for flat consistency groups. |
| 6621458     | The destination Consistency Group is the source of a SnapMirror Synchronous (SM-S) relationship. Sources of SM-S relationships cannot be the destination of any other SnapMirror relationship. |
| 6621782     | A property of the policy is not valid for relationships between these endpoints. |
| 13304120    | Values specified for the source.path and destination.path properties do not match the relationship's source.path or destination.path properties.

| 13304003    | IntraCluster flip operation is not supported. |
| 6621125     | The policy is not valid for relationships with FlexGroup volume endpoints. Only policies without snapshot creation schedules are supported for these relationships. |
| 13304093    | The property specified is not supported for the specified relationships. |
| 6622077     | The expand operation has failed on the SnapMirror active sync relationship with specified destination path. |
| 6619720     | Relationship information has been updated and is being propagated. Wait a few minutes and try the operation again. |
*/
type SnapmirrorRelationshipModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snapmirror relationship modify default response has a 2xx status code
func (o *SnapmirrorRelationshipModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapmirror relationship modify default response has a 3xx status code
func (o *SnapmirrorRelationshipModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapmirror relationship modify default response has a 4xx status code
func (o *SnapmirrorRelationshipModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapmirror relationship modify default response has a 5xx status code
func (o *SnapmirrorRelationshipModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapmirror relationship modify default response a status code equal to that given
func (o *SnapmirrorRelationshipModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snapmirror relationship modify default response
func (o *SnapmirrorRelationshipModifyDefault) Code() int {
	return o._statusCode
}

func (o *SnapmirrorRelationshipModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/relationships/{uuid}][%d] snapmirror_relationship_modify default %s", o._statusCode, payload)
}

func (o *SnapmirrorRelationshipModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /snapmirror/relationships/{uuid}][%d] snapmirror_relationship_modify default %s", o._statusCode, payload)
}

func (o *SnapmirrorRelationshipModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
