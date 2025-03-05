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

// SvmPeerPermissionCollectionGetReader is a Reader for the SvmPeerPermissionCollectionGet structure.
type SvmPeerPermissionCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmPeerPermissionCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmPeerPermissionCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmPeerPermissionCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmPeerPermissionCollectionGetOK creates a SvmPeerPermissionCollectionGetOK with default headers values
func NewSvmPeerPermissionCollectionGetOK() *SvmPeerPermissionCollectionGetOK {
	return &SvmPeerPermissionCollectionGetOK{}
}

/*
SvmPeerPermissionCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SvmPeerPermissionCollectionGetOK struct {
	Payload *models.SvmPeerPermissionResponse
}

// IsSuccess returns true when this svm peer permission collection get o k response has a 2xx status code
func (o *SvmPeerPermissionCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer permission collection get o k response has a 3xx status code
func (o *SvmPeerPermissionCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer permission collection get o k response has a 4xx status code
func (o *SvmPeerPermissionCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer permission collection get o k response has a 5xx status code
func (o *SvmPeerPermissionCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer permission collection get o k response a status code equal to that given
func (o *SvmPeerPermissionCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm peer permission collection get o k response
func (o *SvmPeerPermissionCollectionGetOK) Code() int {
	return 200
}

func (o *SvmPeerPermissionCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peer-permissions][%d] svmPeerPermissionCollectionGetOK %s", 200, payload)
}

func (o *SvmPeerPermissionCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peer-permissions][%d] svmPeerPermissionCollectionGetOK %s", 200, payload)
}

func (o *SvmPeerPermissionCollectionGetOK) GetPayload() *models.SvmPeerPermissionResponse {
	return o.Payload
}

func (o *SvmPeerPermissionCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SvmPeerPermissionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmPeerPermissionCollectionGetDefault creates a SvmPeerPermissionCollectionGetDefault with default headers values
func NewSvmPeerPermissionCollectionGetDefault(code int) *SvmPeerPermissionCollectionGetDefault {
	return &SvmPeerPermissionCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	SvmPeerPermissionCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 26345574    | Failed to find the SVM or volume name with UUID. |
```
<br/>
*/
type SvmPeerPermissionCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm peer permission collection get default response has a 2xx status code
func (o *SvmPeerPermissionCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm peer permission collection get default response has a 3xx status code
func (o *SvmPeerPermissionCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm peer permission collection get default response has a 4xx status code
func (o *SvmPeerPermissionCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm peer permission collection get default response has a 5xx status code
func (o *SvmPeerPermissionCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm peer permission collection get default response a status code equal to that given
func (o *SvmPeerPermissionCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm peer permission collection get default response
func (o *SvmPeerPermissionCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *SvmPeerPermissionCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peer-permissions][%d] svm_peer_permission_collection_get default %s", o._statusCode, payload)
}

func (o *SvmPeerPermissionCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peer-permissions][%d] svm_peer_permission_collection_get default %s", o._statusCode, payload)
}

func (o *SvmPeerPermissionCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmPeerPermissionCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
