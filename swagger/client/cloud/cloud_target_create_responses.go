// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// CloudTargetCreateReader is a Reader for the CloudTargetCreate structure.
type CloudTargetCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudTargetCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCloudTargetCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCloudTargetCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCloudTargetCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudTargetCreateCreated creates a CloudTargetCreateCreated with default headers values
func NewCloudTargetCreateCreated() *CloudTargetCreateCreated {
	return &CloudTargetCreateCreated{}
}

/*
CloudTargetCreateCreated describes a response with status code 201, with default header values.

Created
*/
type CloudTargetCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cloud target create created response has a 2xx status code
func (o *CloudTargetCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud target create created response has a 3xx status code
func (o *CloudTargetCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud target create created response has a 4xx status code
func (o *CloudTargetCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud target create created response has a 5xx status code
func (o *CloudTargetCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud target create created response a status code equal to that given
func (o *CloudTargetCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the cloud target create created response
func (o *CloudTargetCreateCreated) Code() int {
	return 201
}

func (o *CloudTargetCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloud/targets][%d] cloudTargetCreateCreated %s", 201, payload)
}

func (o *CloudTargetCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloud/targets][%d] cloudTargetCreateCreated %s", 201, payload)
}

func (o *CloudTargetCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *CloudTargetCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudTargetCreateAccepted creates a CloudTargetCreateAccepted with default headers values
func NewCloudTargetCreateAccepted() *CloudTargetCreateAccepted {
	return &CloudTargetCreateAccepted{}
}

/*
CloudTargetCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type CloudTargetCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cloud target create accepted response has a 2xx status code
func (o *CloudTargetCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud target create accepted response has a 3xx status code
func (o *CloudTargetCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud target create accepted response has a 4xx status code
func (o *CloudTargetCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud target create accepted response has a 5xx status code
func (o *CloudTargetCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud target create accepted response a status code equal to that given
func (o *CloudTargetCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cloud target create accepted response
func (o *CloudTargetCreateAccepted) Code() int {
	return 202
}

func (o *CloudTargetCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloud/targets][%d] cloudTargetCreateAccepted %s", 202, payload)
}

func (o *CloudTargetCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloud/targets][%d] cloudTargetCreateAccepted %s", 202, payload)
}

func (o *CloudTargetCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *CloudTargetCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudTargetCreateDefault creates a CloudTargetCreateDefault with default headers values
func NewCloudTargetCreateDefault(code int) *CloudTargetCreateDefault {
	return &CloudTargetCreateDefault{
		_statusCode: code,
	}
}

/*
	CloudTargetCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 786436 | VLDB is not running. |
| 786908 | Capability check failed. |
| 787015 | \"encryption-key-id\" is not expected when encrypt is false. |
| 787016 | An object store configuration with the same combination of server and container name already exists. |
| 787020 | Failed to check whether the FabricPool capability exists. |
| 787021 | The FabricPool capability is not enabled. |
| 787030 | Object store configuration creation is not supported in this cluster. |
| 787036 | Server name is invalid. A valid server name must be a fully qualified domain name. |
| 787037 | Cannot verify availability of the object store. |
| 787038 | Object store provider type requires a FabricPool license. |
| 787039 | Failed to retrieve FabricPool capacity license information. |
| 787044 | Cannot create object store configuration as the number of object stores in the cluster has reached maximum allowed limit. |
| 787054 | An object store configuration with the same combination of server, azure account and container name already exists. |
| 787064 | Object store server provider type does not match object store provider type. Use the provider type that matches the server. |
| 787065 | Certificate validation must be enabled for the object store provider. |
| 787066 | Certificate validation cannot be specified when the \"-is-ssl-enabled\" parameter is false. |
| 787068 | Disabling certificate validation requires an effective cluster version of ONTAP 9.4.0 or later. |
| 787071 | Object store configuration creation requires an effective cluster version of ONTAP 9.4.0 or later. |
| 787082 | Creating an object store configuration requires an effective cluster version of ONTAP 9.5.0 or later. |
| 787089 | The object store provider supports the virtual hosted-style, and the specified \"-server\" contains the container name.  The container specified in the \"-server\" parameter must be the same as the name of the container specified in the \"-container\" parameter. |
| 787133 | Could not retrieve temporary credentials from the CAP server. |
| 787134 | Could not retrieve temporary credentials from the CAP server due to an unexpected response. |
| 787136 | Specifying \"CAP\" as the \"-auth-type\" requires an effective cluster version of ONTAP 9.5.0 or later. |
| 787148 | The clock on node is behind by more than the maximum of 5 minutes. |
| 787149 | The clock on node is ahead by more than the maximum of 5 minutes. |
| 787158 | An object store configuration with the same combination of server and container name already exists. |
| 787159 | An object store configuration with the same name already exists. |
| 787179 | One or more clusters in this MetroCluster configuration do not have an effective cluster version of ONTAP 9.7.0 or later. |
| 787184 | Using HTTP proxies with FabricPool requires an effective cluster version of ONTAP 9.7.0 or later. |
| 787185 | There is no HTTP proxy for IPspace. Refer to the \"vserver http-proxy\" man page for details. |
| 787188 | Object store configuration belongs to another cluster and cannot be modified from the local cluster, unless the cluster is in switchover mode. |
| 787189 | Object store configuration name must not have the \"-mc\" suffix when the configuration is created for a local cluster. To create an object store configuration which belongs to another cluster, the cluster must be in switchover mode and \"-cluster\" must be specified. |
| 787209 | Object store is not accessible from the partner cluster in a MetroCluster configuration. |
| 787216 | Cannot perform object store configuration operations on a cluster that is waiting for switchback. |
| 787222 | Object store connectivity check failed on partner cluster in MetroCluster configuration. Wait a few minutes, and try the operation again. |
| 787223 | Specifying \"GCP_SA\" as the \"-auth-type\" requires an effective cluster version of ONTAP 9.7.0 or later. |
| 787227 | Specifying \"Azure_MSI\" as the \"-auth-type\" requires an effective cluster version of ONTAP 9.7 or later. |
| 787228 | SSL is required for this configuration. |
| 787229 | Cannot perform operation as URL style is not supported with object store provider type. |
| 787233 | The hash key for enabling this FabricPool feature is not present on the cluster. |
| 787234 | The hash key provided for the node to enable this FabricPool feature is not valid. |
| 787254 | The parameter is not supported on this system. |
| 787257 | The parameter \"-encryption-context\" is only applicable for AWS object store provider. |
| 787262 | The \"create-container\" option is applicable only for an SGWS or ONTAP_S3 object store provider. |
| 787301 | ONTAP S3 Bucket is of type NAS and is not supported as a target for FabricPool. |
| 787302 | Cannot use HTTP port with \"-is-ssl-enabled\" set to true. |
| 787303 | Cannot use HTTPS port with \"-is-ssl-enabled\" set to false. |
| 787350 | Creating an object store configuration with a Managed Service Identity (MSI) token is only supported on Azure NetApp files. |
| 787351 | Internal Error. Invalid authentication type. |
| 787352 | Creating an object store configuration with a Managed Service Identity (MSI) token requires an effective cluster version of ONTAP 9.16.1 or later. |
| 787353 | Creating an object store configuration with a Managed Service Identity (MSI) token is not supported for this owner. |
| 787354 | The specified properties are mutually exclusive. |
| 787355 | _azure_msi_ `authentication-type` is only supported with _Azure_Cloud_ `provider-type`. |
| 45940761 | Hostname cannot be resolved. Check the spelling of the hostname and check the system DNS availability using the \"vserver services name-service dns check\" command. |
| 45940778 | Bucket already exists. |
| 139591795 | Object store configuration for S3 SnapMirror representing ONTAP S3 object store provider only supports \\\"*\\\" as the container name. |
| 139591796 | Object store configuration name for S3 SnapMirror representing ONTAP S3 object store provider must be of the format \\\"vserver:<destination_vserver_uuid>:<source_vserver_uuid>\\\". |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type CloudTargetCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cloud target create default response has a 2xx status code
func (o *CloudTargetCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cloud target create default response has a 3xx status code
func (o *CloudTargetCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cloud target create default response has a 4xx status code
func (o *CloudTargetCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cloud target create default response has a 5xx status code
func (o *CloudTargetCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cloud target create default response a status code equal to that given
func (o *CloudTargetCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cloud target create default response
func (o *CloudTargetCreateDefault) Code() int {
	return o._statusCode
}

func (o *CloudTargetCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloud/targets][%d] cloud_target_create default %s", o._statusCode, payload)
}

func (o *CloudTargetCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cloud/targets][%d] cloud_target_create default %s", o._statusCode, payload)
}

func (o *CloudTargetCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CloudTargetCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
