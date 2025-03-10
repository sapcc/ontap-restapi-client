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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// SvmPeerInstanceGetReader is a Reader for the SvmPeerInstanceGet structure.
type SvmPeerInstanceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmPeerInstanceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmPeerInstanceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmPeerInstanceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmPeerInstanceGetOK creates a SvmPeerInstanceGetOK with default headers values
func NewSvmPeerInstanceGetOK() *SvmPeerInstanceGetOK {
	return &SvmPeerInstanceGetOK{}
}

/*
SvmPeerInstanceGetOK describes a response with status code 200, with default header values.

OK
*/
type SvmPeerInstanceGetOK struct {
	Payload *models.SvmPeer
}

// IsSuccess returns true when this svm peer instance get o k response has a 2xx status code
func (o *SvmPeerInstanceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer instance get o k response has a 3xx status code
func (o *SvmPeerInstanceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer instance get o k response has a 4xx status code
func (o *SvmPeerInstanceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer instance get o k response has a 5xx status code
func (o *SvmPeerInstanceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer instance get o k response a status code equal to that given
func (o *SvmPeerInstanceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm peer instance get o k response
func (o *SvmPeerInstanceGetOK) Code() int {
	return 200
}

func (o *SvmPeerInstanceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peers/{uuid}][%d] svmPeerInstanceGetOK %s", 200, payload)
}

func (o *SvmPeerInstanceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peers/{uuid}][%d] svmPeerInstanceGetOK %s", 200, payload)
}

func (o *SvmPeerInstanceGetOK) GetPayload() *models.SvmPeer {
	return o.Payload
}

func (o *SvmPeerInstanceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SvmPeer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmPeerInstanceGetDefault creates a SvmPeerInstanceGetDefault with default headers values
func NewSvmPeerInstanceGetDefault(code int) *SvmPeerInstanceGetDefault {
	return &SvmPeerInstanceGetDefault{
		_statusCode: code,
	}
}

/*
	SvmPeerInstanceGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 26345578    | Internal error. Unable to retrieve local or peer SVM name. |
| 9896086     | Peer SVM name conflicts with one of the following: a peer SVM in an existing SVM peer relationship, a local SVM, or an IPSpace. Use the \"name\" property to uniquely specify the peer SVM alias name. |
```
<br/>
*/
type SvmPeerInstanceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm peer instance get default response has a 2xx status code
func (o *SvmPeerInstanceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm peer instance get default response has a 3xx status code
func (o *SvmPeerInstanceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm peer instance get default response has a 4xx status code
func (o *SvmPeerInstanceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm peer instance get default response has a 5xx status code
func (o *SvmPeerInstanceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm peer instance get default response a status code equal to that given
func (o *SvmPeerInstanceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm peer instance get default response
func (o *SvmPeerInstanceGetDefault) Code() int {
	return o._statusCode
}

func (o *SvmPeerInstanceGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peers/{uuid}][%d] svm_peer_instance_get default %s", o._statusCode, payload)
}

func (o *SvmPeerInstanceGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /svm/peers/{uuid}][%d] svm_peer_instance_get default %s", o._statusCode, payload)
}

func (o *SvmPeerInstanceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmPeerInstanceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
