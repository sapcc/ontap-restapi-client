// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// IgroupCreateReader is a Reader for the IgroupCreate structure.
type IgroupCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IgroupCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIgroupCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIgroupCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIgroupCreateCreated creates a IgroupCreateCreated with default headers values
func NewIgroupCreateCreated() *IgroupCreateCreated {
	return &IgroupCreateCreated{}
}

/*
IgroupCreateCreated describes a response with status code 201, with default header values.

Created
*/
type IgroupCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.IgroupResponse
}

// IsSuccess returns true when this igroup create created response has a 2xx status code
func (o *IgroupCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this igroup create created response has a 3xx status code
func (o *IgroupCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this igroup create created response has a 4xx status code
func (o *IgroupCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this igroup create created response has a 5xx status code
func (o *IgroupCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this igroup create created response a status code equal to that given
func (o *IgroupCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the igroup create created response
func (o *IgroupCreateCreated) Code() int {
	return 201
}

func (o *IgroupCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/igroups][%d] igroupCreateCreated %s", 201, payload)
}

func (o *IgroupCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/igroups][%d] igroupCreateCreated %s", 201, payload)
}

func (o *IgroupCreateCreated) GetPayload() *models.IgroupResponse {
	return o.Payload
}

func (o *IgroupCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.IgroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIgroupCreateDefault creates a IgroupCreateDefault with default headers values
func NewIgroupCreateDefault(code int) *IgroupCreateDefault {
	return &IgroupCreateDefault{
		_statusCode: code,
	}
}

/*
	IgroupCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | The supplied SVM does not exist. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5373958 | An invalid initiator group name was supplied. |
| 5373966 | An initiator group cannot be created in an SVM that is configured for NVMe. |
| 5373969 | A supplied initiator name looks like an iSCSI IQN initiator, but the portions after the prefix are missing. |
| 5373971 | A supplied initiator name looks like an iSCSI IQN initiator, but the date portion is invalid. |
| 5373972 | A supplied initiator name looks like an iSCSI IQN initiator, but the naming authority portion is invalid. |
| 5373977 | A supplied initiator name looks like an iSCSI EUI initiator, but the length is invalid. |
| 5373978 | A supplied initiator name looks like an iSCSI EUI initiator, but the format is invalid. |
| 5373982 | An invalid WWN was specified. The length is incorrect. |
| 5373983 | An invalid WWN was specified. The format is incorrect. |
| 5373992 | A supplied initiator name was too long to be valid. |
| 5373993 | A supplied initiator name did not match any valid format. |
| 5374023 | An initiator group with the same name already exists. |
| 5374027 | An attempt was made to bind a portset with no member network interfaces to the initiator group. |
| 5374028 | An attempt was made to bind a portset with an incompatible protocol to the initiator group. |
| 5374038 | An invalid Fibre Channel WWPN was supplied. |
| 5374039 | An invalid iSCSI initiator name was supplied. |
| 5374040 | Initiators and child initiator groups were both supplied, but only one option is allowed. |
| 5374732 | An initiator is already in another initiator group with a conflicting operating system type. |
| 5374735 | An attempt was made to add a child igroup that would exceed the maximum allowable depth. |
| 5374737 | A supplied child initiator group already exists in another initiator group's hierarchy. |
| 5374739 | A supplied child initiator group has an operating system type that differs from the parent initiator group. |
| 5374740 | A supplied child initiator group has an protocol that differs from the parent initiator group. |
| 5374741 | A supplied child initiator group is already owned by a different child in the initiator group's hierarchy. |
| 5374742 | A supplied child initiator group contains an initiator that is already owned by another initiator group in the hierarchy. |
| 5374745 | Initiator group cannot be added as a child to itself. |
| 5374746 | The cluster is currently running in a mixed version and nested initiator groups cannot be created until the effective cluster version reaches 9.9.1. |
| 5374747 | The cluster is currently running in a mixed version and initiator group comments cannot be created until the effective cluster version reaches 9.9.1. |
| 5374758 | An error was reported by the peer cluster while creating a replicated initiator group. The specific error will be included as a nested error. |
| 5374878 | The supplied child initiator group does not exist. |
| 5374911 | The supplied portset does not exist. |
| 5374917 | Duplicated initiators have conflicting property values. |
| 5375055 | The `local_svm` property of an initiator proximity was not specified. |
| 5375056 | An SVM peering relationship that does not have the initiator group's SVM as the local SVM was specified. |
| 5375258 | The igroup is already replicated to a different peer SVM. |
| 5375261 | Setting initiator proximity is not supported for the SVM type. |
| 5376057 | Setting initiator proximity is not supported for the ONTAP version. |
| 5376059 | Setting initiator proximity to a peer that is either the destination of an SVM DR relationship or in a Metrocluster configuration is not supported. |
| 5376253 | Initiator group replication requires an effective cluster version of 9.15.1. |
| 5376255 | Initiator group replication requires the peer cluster to have an effective cluster version of 9.15.1. |
| 5376488 | An NVMe over Fabrics subsystem already exists with the requested name. |
| 6620376 | SVM peering information is unavailable. |
| 6620384 | The supplied SVMs are not peered. |
| 26345672 | The specified SVM peering relationship was not found. |
| 26345673 | An SVM peering relationship between the initiator group's SVM and specified peer SVM was not found. |
| 26345675 | An SVM peering relationship UUID and name were specified and they do not refer to the same SVM peering relationship. |
| 26345680 | Supplied SVM peer is on the local cluster. The operation requires a peer on a remote cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IgroupCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this igroup create default response has a 2xx status code
func (o *IgroupCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this igroup create default response has a 3xx status code
func (o *IgroupCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this igroup create default response has a 4xx status code
func (o *IgroupCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this igroup create default response has a 5xx status code
func (o *IgroupCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this igroup create default response a status code equal to that given
func (o *IgroupCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the igroup create default response
func (o *IgroupCreateDefault) Code() int {
	return o._statusCode
}

func (o *IgroupCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/igroups][%d] igroup_create default %s", o._statusCode, payload)
}

func (o *IgroupCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/san/igroups][%d] igroup_create default %s", o._statusCode, payload)
}

func (o *IgroupCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IgroupCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
