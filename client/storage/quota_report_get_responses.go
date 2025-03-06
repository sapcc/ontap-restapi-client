// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// QuotaReportGetReader is a Reader for the QuotaReportGet structure.
type QuotaReportGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuotaReportGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuotaReportGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQuotaReportGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuotaReportGetOK creates a QuotaReportGetOK with default headers values
func NewQuotaReportGetOK() *QuotaReportGetOK {
	return &QuotaReportGetOK{}
}

/*
QuotaReportGetOK describes a response with status code 200, with default header values.

OK
*/
type QuotaReportGetOK struct {
	Payload *models.QuotaReport
}

// IsSuccess returns true when this quota report get o k response has a 2xx status code
func (o *QuotaReportGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this quota report get o k response has a 3xx status code
func (o *QuotaReportGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this quota report get o k response has a 4xx status code
func (o *QuotaReportGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this quota report get o k response has a 5xx status code
func (o *QuotaReportGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this quota report get o k response a status code equal to that given
func (o *QuotaReportGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the quota report get o k response
func (o *QuotaReportGetOK) Code() int {
	return 200
}

func (o *QuotaReportGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/quota/reports/{volume.uuid}/{index}][%d] quotaReportGetOK %s", 200, payload)
}

func (o *QuotaReportGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/quota/reports/{volume.uuid}/{index}][%d] quotaReportGetOK %s", 200, payload)
}

func (o *QuotaReportGetOK) GetPayload() *models.QuotaReport {
	return o.Payload
}

func (o *QuotaReportGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuotaReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotaReportGetDefault creates a QuotaReportGetDefault with default headers values
func NewQuotaReportGetDefault(code int) *QuotaReportGetDefault {
	return &QuotaReportGetDefault{
		_statusCode: code,
	}
}

/*
	QuotaReportGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 918235 | A volume with UUID was not found. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type QuotaReportGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this quota report get default response has a 2xx status code
func (o *QuotaReportGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this quota report get default response has a 3xx status code
func (o *QuotaReportGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this quota report get default response has a 4xx status code
func (o *QuotaReportGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this quota report get default response has a 5xx status code
func (o *QuotaReportGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this quota report get default response a status code equal to that given
func (o *QuotaReportGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the quota report get default response
func (o *QuotaReportGetDefault) Code() int {
	return o._statusCode
}

func (o *QuotaReportGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/quota/reports/{volume.uuid}/{index}][%d] quota_report_get default %s", o._statusCode, payload)
}

func (o *QuotaReportGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/quota/reports/{volume.uuid}/{index}][%d] quota_report_get default %s", o._statusCode, payload)
}

func (o *QuotaReportGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuotaReportGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
