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

// IgroupNestedDeleteReader is a Reader for the IgroupNestedDelete structure.
type IgroupNestedDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IgroupNestedDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIgroupNestedDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIgroupNestedDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIgroupNestedDeleteOK creates a IgroupNestedDeleteOK with default headers values
func NewIgroupNestedDeleteOK() *IgroupNestedDeleteOK {
	return &IgroupNestedDeleteOK{}
}

/*
IgroupNestedDeleteOK describes a response with status code 200, with default header values.

OK
*/
type IgroupNestedDeleteOK struct {
}

// IsSuccess returns true when this igroup nested delete o k response has a 2xx status code
func (o *IgroupNestedDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this igroup nested delete o k response has a 3xx status code
func (o *IgroupNestedDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this igroup nested delete o k response has a 4xx status code
func (o *IgroupNestedDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this igroup nested delete o k response has a 5xx status code
func (o *IgroupNestedDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this igroup nested delete o k response a status code equal to that given
func (o *IgroupNestedDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the igroup nested delete o k response
func (o *IgroupNestedDeleteOK) Code() int {
	return 200
}

func (o *IgroupNestedDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/igroups/{igroup.uuid}/igroups/{uuid}][%d] igroupNestedDeleteOK", 200)
}

func (o *IgroupNestedDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/igroups/{igroup.uuid}/igroups/{uuid}][%d] igroupNestedDeleteOK", 200)
}

func (o *IgroupNestedDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIgroupNestedDeleteDefault creates a IgroupNestedDeleteDefault with default headers values
func NewIgroupNestedDeleteDefault(code int) *IgroupNestedDeleteDefault {
	return &IgroupNestedDeleteDefault{
		_statusCode: code,
	}
}

/*
	IgroupNestedDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1254213 | The initiator group is mapped to one or more LUNs and `allow_delete_while_mapped` has not been specified. |
| 5374738 | The child initiator group is not owned by the parent initiator group. |
| 5374743 | LUN maps exist for a parent initiator group. |
| 5374852 | The initiator group specified in the URI does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IgroupNestedDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this igroup nested delete default response has a 2xx status code
func (o *IgroupNestedDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this igroup nested delete default response has a 3xx status code
func (o *IgroupNestedDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this igroup nested delete default response has a 4xx status code
func (o *IgroupNestedDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this igroup nested delete default response has a 5xx status code
func (o *IgroupNestedDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this igroup nested delete default response a status code equal to that given
func (o *IgroupNestedDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the igroup nested delete default response
func (o *IgroupNestedDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IgroupNestedDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/igroups/{igroup.uuid}/igroups/{uuid}][%d] igroup_nested_delete default %s", o._statusCode, payload)
}

func (o *IgroupNestedDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/igroups/{igroup.uuid}/igroups/{uuid}][%d] igroup_nested_delete default %s", o._statusCode, payload)
}

func (o *IgroupNestedDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IgroupNestedDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
