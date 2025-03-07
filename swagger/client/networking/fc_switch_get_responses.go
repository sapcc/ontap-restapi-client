// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// FcSwitchGetReader is a Reader for the FcSwitchGet structure.
type FcSwitchGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcSwitchGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcSwitchGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcSwitchGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcSwitchGetOK creates a FcSwitchGetOK with default headers values
func NewFcSwitchGetOK() *FcSwitchGetOK {
	return &FcSwitchGetOK{}
}

/*
FcSwitchGetOK describes a response with status code 200, with default header values.

OK
*/
type FcSwitchGetOK struct {
	Payload *models.FcSwitch
}

// IsSuccess returns true when this fc switch get o k response has a 2xx status code
func (o *FcSwitchGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fc switch get o k response has a 3xx status code
func (o *FcSwitchGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fc switch get o k response has a 4xx status code
func (o *FcSwitchGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fc switch get o k response has a 5xx status code
func (o *FcSwitchGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fc switch get o k response a status code equal to that given
func (o *FcSwitchGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fc switch get o k response
func (o *FcSwitchGetOK) Code() int {
	return 200
}

func (o *FcSwitchGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/fabrics/{fabric.name}/switches/{wwn}][%d] fcSwitchGetOK %s", 200, payload)
}

func (o *FcSwitchGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/fabrics/{fabric.name}/switches/{wwn}][%d] fcSwitchGetOK %s", 200, payload)
}

func (o *FcSwitchGetOK) GetPayload() *models.FcSwitch {
	return o.Payload
}

func (o *FcSwitchGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcSwitch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFcSwitchGetDefault creates a FcSwitchGetDefault with default headers values
func NewFcSwitchGetDefault(code int) *FcSwitchGetDefault {
	return &FcSwitchGetDefault{
		_statusCode: code,
	}
}

/*
	FcSwitchGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5373982 | An invalid WWN was specified. The length is incorrect. |
| 5373983 | An invalid WWN was specified. The format is incorrect. |
| 5375053 | The Fibre Channel fabric specified by name in the request URI was not found in the FC fabric cache. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FcSwitchGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fc switch get default response has a 2xx status code
func (o *FcSwitchGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fc switch get default response has a 3xx status code
func (o *FcSwitchGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fc switch get default response has a 4xx status code
func (o *FcSwitchGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fc switch get default response has a 5xx status code
func (o *FcSwitchGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fc switch get default response a status code equal to that given
func (o *FcSwitchGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fc switch get default response
func (o *FcSwitchGetDefault) Code() int {
	return o._statusCode
}

func (o *FcSwitchGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/fabrics/{fabric.name}/switches/{wwn}][%d] fc_switch_get default %s", o._statusCode, payload)
}

func (o *FcSwitchGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/fc/fabrics/{fabric.name}/switches/{wwn}][%d] fc_switch_get default %s", o._statusCode, payload)
}

func (o *FcSwitchGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcSwitchGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
