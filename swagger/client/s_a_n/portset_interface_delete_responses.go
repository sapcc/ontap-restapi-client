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

// PortsetInterfaceDeleteReader is a Reader for the PortsetInterfaceDelete structure.
type PortsetInterfaceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortsetInterfaceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPortsetInterfaceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPortsetInterfaceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPortsetInterfaceDeleteOK creates a PortsetInterfaceDeleteOK with default headers values
func NewPortsetInterfaceDeleteOK() *PortsetInterfaceDeleteOK {
	return &PortsetInterfaceDeleteOK{}
}

/*
PortsetInterfaceDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PortsetInterfaceDeleteOK struct {
}

// IsSuccess returns true when this portset interface delete o k response has a 2xx status code
func (o *PortsetInterfaceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this portset interface delete o k response has a 3xx status code
func (o *PortsetInterfaceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this portset interface delete o k response has a 4xx status code
func (o *PortsetInterfaceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this portset interface delete o k response has a 5xx status code
func (o *PortsetInterfaceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this portset interface delete o k response a status code equal to that given
func (o *PortsetInterfaceDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the portset interface delete o k response
func (o *PortsetInterfaceDeleteOK) Code() int {
	return 200
}

func (o *PortsetInterfaceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/portsets/{portset.uuid}/interfaces/{uuid}][%d] portsetInterfaceDeleteOK", 200)
}

func (o *PortsetInterfaceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/portsets/{portset.uuid}/interfaces/{uuid}][%d] portsetInterfaceDeleteOK", 200)
}

func (o *PortsetInterfaceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPortsetInterfaceDeleteDefault creates a PortsetInterfaceDeleteDefault with default headers values
func NewPortsetInterfaceDeleteDefault(code int) *PortsetInterfaceDeleteDefault {
	return &PortsetInterfaceDeleteDefault{
		_statusCode: code,
	}
}

/*
	PortsetInterfaceDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374906 | A specified network interface was not found. |
| 5374908 | The portset specified in the URI does not exist. |
| 5374916 | The specified network interface is not in the portset. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type PortsetInterfaceDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this portset interface delete default response has a 2xx status code
func (o *PortsetInterfaceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this portset interface delete default response has a 3xx status code
func (o *PortsetInterfaceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this portset interface delete default response has a 4xx status code
func (o *PortsetInterfaceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this portset interface delete default response has a 5xx status code
func (o *PortsetInterfaceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this portset interface delete default response a status code equal to that given
func (o *PortsetInterfaceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the portset interface delete default response
func (o *PortsetInterfaceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *PortsetInterfaceDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/portsets/{portset.uuid}/interfaces/{uuid}][%d] portset_interface_delete default %s", o._statusCode, payload)
}

func (o *PortsetInterfaceDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/portsets/{portset.uuid}/interfaces/{uuid}][%d] portset_interface_delete default %s", o._statusCode, payload)
}

func (o *PortsetInterfaceDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PortsetInterfaceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
