// Code generated by go-swagger; DO NOT EDIT.

package s_vm

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

// SvmPeerPermissionInstanceGetReader is a Reader for the SvmPeerPermissionInstanceGet structure.
type SvmPeerPermissionInstanceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmPeerPermissionInstanceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmPeerPermissionInstanceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmPeerPermissionInstanceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmPeerPermissionInstanceGetOK creates a SvmPeerPermissionInstanceGetOK with default headers values
func NewSvmPeerPermissionInstanceGetOK() *SvmPeerPermissionInstanceGetOK {
	return &SvmPeerPermissionInstanceGetOK{}
}

/*
SvmPeerPermissionInstanceGetOK describes a response with status code 200, with default header values.

OK
*/
type SvmPeerPermissionInstanceGetOK struct {
	Payload *models.SvmPeerPermission
}

// IsSuccess returns true when this svm peer permission instance get o k response has a 2xx status code
func (o *SvmPeerPermissionInstanceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer permission instance get o k response has a 3xx status code
func (o *SvmPeerPermissionInstanceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer permission instance get o k response has a 4xx status code
func (o *SvmPeerPermissionInstanceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer permission instance get o k response has a 5xx status code
func (o *SvmPeerPermissionInstanceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer permission instance get o k response a status code equal to that given
func (o *SvmPeerPermissionInstanceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm peer permission instance get o k response
func (o *SvmPeerPermissionInstanceGetOK) Code() int {
	return 200
}

func (o *SvmPeerPermissionInstanceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peer-permissions/{cluster_peer.uuid}/{svm.uuid}][%d] svmPeerPermissionInstanceGetOK %s", 200, payload)
}

func (o *SvmPeerPermissionInstanceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peer-permissions/{cluster_peer.uuid}/{svm.uuid}][%d] svmPeerPermissionInstanceGetOK %s", 200, payload)
}

func (o *SvmPeerPermissionInstanceGetOK) GetPayload() *models.SvmPeerPermission {
	return o.Payload
}

func (o *SvmPeerPermissionInstanceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SvmPeerPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmPeerPermissionInstanceGetDefault creates a SvmPeerPermissionInstanceGetDefault with default headers values
func NewSvmPeerPermissionInstanceGetDefault(code int) *SvmPeerPermissionInstanceGetDefault {
	return &SvmPeerPermissionInstanceGetDefault{
		_statusCode: code,
	}
}

/*
	SvmPeerPermissionInstanceGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 26345574    | Failed to find the SVM or volume name with UUID. |
```
<br/>
*/
type SvmPeerPermissionInstanceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm peer permission instance get default response has a 2xx status code
func (o *SvmPeerPermissionInstanceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm peer permission instance get default response has a 3xx status code
func (o *SvmPeerPermissionInstanceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm peer permission instance get default response has a 4xx status code
func (o *SvmPeerPermissionInstanceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm peer permission instance get default response has a 5xx status code
func (o *SvmPeerPermissionInstanceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm peer permission instance get default response a status code equal to that given
func (o *SvmPeerPermissionInstanceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm peer permission instance get default response
func (o *SvmPeerPermissionInstanceGetDefault) Code() int {
	return o._statusCode
}

func (o *SvmPeerPermissionInstanceGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peer-permissions/{cluster_peer.uuid}/{svm.uuid}][%d] svm_peer_permission_instance_get default %s", o._statusCode, payload)
}

func (o *SvmPeerPermissionInstanceGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peer-permissions/{cluster_peer.uuid}/{svm.uuid}][%d] svm_peer_permission_instance_get default %s", o._statusCode, payload)
}

func (o *SvmPeerPermissionInstanceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmPeerPermissionInstanceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
