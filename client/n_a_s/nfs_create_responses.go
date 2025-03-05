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

	"github.com/sapcc/ontap-restapi/models"
)

// NfsCreateReader is a Reader for the NfsCreate structure.
type NfsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNfsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsCreateCreated creates a NfsCreateCreated with default headers values
func NewNfsCreateCreated() *NfsCreateCreated {
	return &NfsCreateCreated{}
}

/*
NfsCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NfsCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.NfsServiceResponse
}

// IsSuccess returns true when this nfs create created response has a 2xx status code
func (o *NfsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nfs create created response has a 3xx status code
func (o *NfsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nfs create created response has a 4xx status code
func (o *NfsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this nfs create created response has a 5xx status code
func (o *NfsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this nfs create created response a status code equal to that given
func (o *NfsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the nfs create created response
func (o *NfsCreateCreated) Code() int {
	return 201
}

func (o *NfsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/services][%d] nfsCreateCreated %s", 201, payload)
}

func (o *NfsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/services][%d] nfsCreateCreated %s", 201, payload)
}

func (o *NfsCreateCreated) GetPayload() *models.NfsServiceResponse {
	return o.Payload
}

func (o *NfsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.NfsServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNfsCreateDefault creates a NfsCreateDefault with default headers values
func NewNfsCreateDefault(code int) *NfsCreateDefault {
	return &NfsCreateDefault{
		_statusCode: code,
	}
}

/*
	NfsCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1534829    | The port numbers allowed are 635 (the default) and 1024 through 9999 |
| 2621516    | This operation is only supported on a data SVM |
| 2621570    | This operation is not permitted on an SVM that is configured as the destination for SVM DR.|
| 2621574    | This operation is not permitted on an SVM that is configured as the destination of a MetroCluster SVM relationship |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name |
| 3276852    | NFSv4.1 implementation name length for the SVM must be less than 256 bytes.|
| 3276900    | NFSv4.1 implementation name cannot be an empty string |
| 3276916    | Vserver is not running |
| 3276994    | Kerberos must be disabled on all LIFs on Vserver before adding or removing AES encryption. Disable Kerberos on the LIF and try again |
| 3277038    | Cannot enable \\\"showmount\\\" feature because it requires an effective cluster version of Data ONTAP 8.3.0 or later |
| 3277048    | The port numbers allowed are 635 (the default) and 1024 through 9999 |
| 3277049    | Cannot enable \\\"showmount\\\" feature on ID-Discard Vserver. Ensure that the Vserver is initialized and retry the command |
| 3277052    | NFSv4.x access to transitioned volumes in this Vserver could trigger conversion of non-Unicode directories to Unicode, which might impact data-serving performance. Before enabling NFSv4.x for this Vserver, refer to the Data and Configuration Transition Guide |
| 3277069    | Cannot disable TCP because the SnapDiff RPC server is in the \\\"on\\\" state |
| 3277085    | The port numbers allowed are 1024 through 9999 |
| 3277089    | Attempting to create an NFS server using 64-bits for NFSv3 FSIDs and File IDs on Vserver. Older client software might not work with 64-bit identifiers|
| 3277099    | Domain name contains invalid characters or it is too short. Allowed characters are: alphabetical characters (A-Za-z), numeric characters (0-9), minus sign (-), and the period (.). The first character must be alphabetical or numeric, last character must not be a minus sign or a period. Minimum supported length: 2 characters, maximum of 256 characters |
| 3277140    | Cannot set \"transport.tcp_max_transfer_size\" to a value other than multiples of 4096 |
| 2621507    | The NFS protocol is not allowed for the specified SVM.|
| 2621519    | SVM name is invalid. The SVM name must begin with a letter or an underscore. If the SVM is of type \"sync-source\", the maximum supported length is 41. Otherwise, the maximum supported length is 47.|
| 262196     | Field \"access_cache_config.harvest_timeout\" cannot be set in this operation.|
*/
type NfsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nfs create default response has a 2xx status code
func (o *NfsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nfs create default response has a 3xx status code
func (o *NfsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nfs create default response has a 4xx status code
func (o *NfsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nfs create default response has a 5xx status code
func (o *NfsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nfs create default response a status code equal to that given
func (o *NfsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nfs create default response
func (o *NfsCreateDefault) Code() int {
	return o._statusCode
}

func (o *NfsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/services][%d] nfs_create default %s", o._statusCode, payload)
}

func (o *NfsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/nfs/services][%d] nfs_create default %s", o._statusCode, payload)
}

func (o *NfsCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
