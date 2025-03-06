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

	"github.com/sapcc/ontap-restapi-client/models"
)

// FpolicyPersistentStoreModifyReader is a Reader for the FpolicyPersistentStoreModify structure.
type FpolicyPersistentStoreModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyPersistentStoreModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyPersistentStoreModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyPersistentStoreModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyPersistentStoreModifyOK creates a FpolicyPersistentStoreModifyOK with default headers values
func NewFpolicyPersistentStoreModifyOK() *FpolicyPersistentStoreModifyOK {
	return &FpolicyPersistentStoreModifyOK{}
}

/*
FpolicyPersistentStoreModifyOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyPersistentStoreModifyOK struct {
}

// IsSuccess returns true when this fpolicy persistent store modify o k response has a 2xx status code
func (o *FpolicyPersistentStoreModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy persistent store modify o k response has a 3xx status code
func (o *FpolicyPersistentStoreModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy persistent store modify o k response has a 4xx status code
func (o *FpolicyPersistentStoreModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy persistent store modify o k response has a 5xx status code
func (o *FpolicyPersistentStoreModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy persistent store modify o k response a status code equal to that given
func (o *FpolicyPersistentStoreModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fpolicy persistent store modify o k response
func (o *FpolicyPersistentStoreModifyOK) Code() int {
	return 200
}

func (o *FpolicyPersistentStoreModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/persistent-stores/{name}][%d] fpolicyPersistentStoreModifyOK", 200)
}

func (o *FpolicyPersistentStoreModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/persistent-stores/{name}][%d] fpolicyPersistentStoreModifyOK", 200)
}

func (o *FpolicyPersistentStoreModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFpolicyPersistentStoreModifyDefault creates a FpolicyPersistentStoreModifyDefault with default headers values
func NewFpolicyPersistentStoreModifyDefault(code int) *FpolicyPersistentStoreModifyDefault {
	return &FpolicyPersistentStoreModifyDefault{
		_statusCode: code,
	}
}

/*
	FpolicyPersistentStoreModifyDefault describes a response with status code -1, with default header values.

	| Error Code | Description |

| ---------- | ----------- |
| 9765053    | Volume mentioned does not exist in the SVM |
| 9765056    | Specified Persistent Store name does not exist in the SVM |
| 9765069    | Volume is not a FlexVol volume. Only FlexVol volumes are supported for this operation. |
| 9765070    | Cannot update a Persistent Store volume when it is part of an enabled policy |
| 9765072    | Volume is not of type RW. Only volumes of type RW are supported. |
| 9765074    | Size is a required parameter for the creation of the Persistent Store. |
| 9765077    | The SVM is not associated with any aggregates. |
| 524849     | Attribute cannot be modified for the target volume because of the specified reason. |
| 917533     | The target volume is not eligible for modification due to certain constraints, such as the volume being offline or currently engaged in a volume move operation. |
| 917826     | Volume minimum autosize must be greater than zero. |
| 918647     | The designated size exceeds the maximum available capacity of the intended volume. |
| 917534     | The value specified for the field \\"-size\\" is too small. Update the field \\"-size\\" with the minimum size allowed and retry the operation. |
*/
type FpolicyPersistentStoreModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fpolicy persistent store modify default response has a 2xx status code
func (o *FpolicyPersistentStoreModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy persistent store modify default response has a 3xx status code
func (o *FpolicyPersistentStoreModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy persistent store modify default response has a 4xx status code
func (o *FpolicyPersistentStoreModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy persistent store modify default response has a 5xx status code
func (o *FpolicyPersistentStoreModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy persistent store modify default response a status code equal to that given
func (o *FpolicyPersistentStoreModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fpolicy persistent store modify default response
func (o *FpolicyPersistentStoreModifyDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyPersistentStoreModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/persistent-stores/{name}][%d] fpolicy_persistent_store_modify default %s", o._statusCode, payload)
}

func (o *FpolicyPersistentStoreModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/persistent-stores/{name}][%d] fpolicy_persistent_store_modify default %s", o._statusCode, payload)
}

func (o *FpolicyPersistentStoreModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyPersistentStoreModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
