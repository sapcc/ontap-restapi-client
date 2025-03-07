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

// SvmPeerModifyReader is a Reader for the SvmPeerModify structure.
type SvmPeerModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmPeerModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmPeerModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSvmPeerModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmPeerModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmPeerModifyOK creates a SvmPeerModifyOK with default headers values
func NewSvmPeerModifyOK() *SvmPeerModifyOK {
	return &SvmPeerModifyOK{}
}

/*
SvmPeerModifyOK describes a response with status code 200, with default header values.

OK
*/
type SvmPeerModifyOK struct {
}

// IsSuccess returns true when this svm peer modify o k response has a 2xx status code
func (o *SvmPeerModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer modify o k response has a 3xx status code
func (o *SvmPeerModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer modify o k response has a 4xx status code
func (o *SvmPeerModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer modify o k response has a 5xx status code
func (o *SvmPeerModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer modify o k response a status code equal to that given
func (o *SvmPeerModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm peer modify o k response
func (o *SvmPeerModifyOK) Code() int {
	return 200
}

func (o *SvmPeerModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /svm/peers/{uuid}][%d] svmPeerModifyOK", 200)
}

func (o *SvmPeerModifyOK) String() string {
	return fmt.Sprintf("[PATCH /svm/peers/{uuid}][%d] svmPeerModifyOK", 200)
}

func (o *SvmPeerModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSvmPeerModifyAccepted creates a SvmPeerModifyAccepted with default headers values
func NewSvmPeerModifyAccepted() *SvmPeerModifyAccepted {
	return &SvmPeerModifyAccepted{}
}

/*
SvmPeerModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmPeerModifyAccepted struct {
}

// IsSuccess returns true when this svm peer modify accepted response has a 2xx status code
func (o *SvmPeerModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer modify accepted response has a 3xx status code
func (o *SvmPeerModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer modify accepted response has a 4xx status code
func (o *SvmPeerModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer modify accepted response has a 5xx status code
func (o *SvmPeerModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer modify accepted response a status code equal to that given
func (o *SvmPeerModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the svm peer modify accepted response
func (o *SvmPeerModifyAccepted) Code() int {
	return 202
}

func (o *SvmPeerModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /svm/peers/{uuid}][%d] svmPeerModifyAccepted", 202)
}

func (o *SvmPeerModifyAccepted) String() string {
	return fmt.Sprintf("[PATCH /svm/peers/{uuid}][%d] svmPeerModifyAccepted", 202)
}

func (o *SvmPeerModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSvmPeerModifyDefault creates a SvmPeerModifyDefault with default headers values
func NewSvmPeerModifyDefault(code int) *SvmPeerModifyDefault {
	return &SvmPeerModifyDefault{
		_statusCode: code,
	}
}

/*
	SvmPeerModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 13434889    | Internal error. Wait and retry. |
| 26345575    | The specified peer cluster name and peer cluster UUID do not match. |
| 26345576    | Given peer state is invalid. |
| 26345577    | One of the following is required: applications, state, or name. |
| 26345578    | Internal error. Unable to retrieve local or peer SVM name. |
| 26345579    | The specified field is invalid. |
| 26345581    | Peer cluster name could not be retrieved or validated. |
| 9896077     | The peer relationship is in use by FlexCache. View the FlexCache relationships, delete them and retry the operation. |
| 9896088     | System generated a name for the peer SVM because of a naming conflict. Use the name property to uniquely identify the peer SVM alias name. |
```
<br/>
*/
type SvmPeerModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm peer modify default response has a 2xx status code
func (o *SvmPeerModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm peer modify default response has a 3xx status code
func (o *SvmPeerModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm peer modify default response has a 4xx status code
func (o *SvmPeerModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm peer modify default response has a 5xx status code
func (o *SvmPeerModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm peer modify default response a status code equal to that given
func (o *SvmPeerModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm peer modify default response
func (o *SvmPeerModifyDefault) Code() int {
	return o._statusCode
}

func (o *SvmPeerModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /svm/peers/{uuid}][%d] svm_peer_modify default %s", o._statusCode, payload)
}

func (o *SvmPeerModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /svm/peers/{uuid}][%d] svm_peer_modify default %s", o._statusCode, payload)
}

func (o *SvmPeerModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmPeerModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
