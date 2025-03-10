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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// WwpnAliasDeleteReader is a Reader for the WwpnAliasDelete structure.
type WwpnAliasDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WwpnAliasDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWwpnAliasDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWwpnAliasDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWwpnAliasDeleteOK creates a WwpnAliasDeleteOK with default headers values
func NewWwpnAliasDeleteOK() *WwpnAliasDeleteOK {
	return &WwpnAliasDeleteOK{}
}

/*
WwpnAliasDeleteOK describes a response with status code 200, with default header values.

OK
*/
type WwpnAliasDeleteOK struct {
}

// IsSuccess returns true when this wwpn alias delete o k response has a 2xx status code
func (o *WwpnAliasDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wwpn alias delete o k response has a 3xx status code
func (o *WwpnAliasDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wwpn alias delete o k response has a 4xx status code
func (o *WwpnAliasDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wwpn alias delete o k response has a 5xx status code
func (o *WwpnAliasDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wwpn alias delete o k response a status code equal to that given
func (o *WwpnAliasDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wwpn alias delete o k response
func (o *WwpnAliasDeleteOK) Code() int {
	return 200
}

func (o *WwpnAliasDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /network/fc/wwpn-aliases/{svm.uuid}/{alias}][%d] wwpnAliasDeleteOK", 200)
}

func (o *WwpnAliasDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /network/fc/wwpn-aliases/{svm.uuid}/{alias}][%d] wwpnAliasDeleteOK", 200)
}

func (o *WwpnAliasDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWwpnAliasDeleteDefault creates a WwpnAliasDeleteDefault with default headers values
func NewWwpnAliasDeleteDefault(code int) *WwpnAliasDeleteDefault {
	return &WwpnAliasDeleteDefault{
		_statusCode: code,
	}
}

/*
	WwpnAliasDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | The specified SVM does not exist or cannot be accessed by this user. |
| 5374046 | The alias could not be found. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type WwpnAliasDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this wwpn alias delete default response has a 2xx status code
func (o *WwpnAliasDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wwpn alias delete default response has a 3xx status code
func (o *WwpnAliasDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wwpn alias delete default response has a 4xx status code
func (o *WwpnAliasDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wwpn alias delete default response has a 5xx status code
func (o *WwpnAliasDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wwpn alias delete default response a status code equal to that given
func (o *WwpnAliasDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the wwpn alias delete default response
func (o *WwpnAliasDeleteDefault) Code() int {
	return o._statusCode
}

func (o *WwpnAliasDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/fc/wwpn-aliases/{svm.uuid}/{alias}][%d] wwpn_alias_delete default %s", o._statusCode, payload)
}

func (o *WwpnAliasDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/fc/wwpn-aliases/{svm.uuid}/{alias}][%d] wwpn_alias_delete default %s", o._statusCode, payload)
}

func (o *WwpnAliasDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *WwpnAliasDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
