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

	"github.com/sapcc/ontap-restapi/models"
)

// PlexCollectionGetReader is a Reader for the PlexCollectionGet structure.
type PlexCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlexCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlexCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPlexCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPlexCollectionGetOK creates a PlexCollectionGetOK with default headers values
func NewPlexCollectionGetOK() *PlexCollectionGetOK {
	return &PlexCollectionGetOK{}
}

/*
PlexCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PlexCollectionGetOK struct {
	Payload *models.PlexResponse
}

// IsSuccess returns true when this plex collection get o k response has a 2xx status code
func (o *PlexCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plex collection get o k response has a 3xx status code
func (o *PlexCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plex collection get o k response has a 4xx status code
func (o *PlexCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plex collection get o k response has a 5xx status code
func (o *PlexCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plex collection get o k response a status code equal to that given
func (o *PlexCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plex collection get o k response
func (o *PlexCollectionGetOK) Code() int {
	return 200
}

func (o *PlexCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/plexes][%d] plexCollectionGetOK %s", 200, payload)
}

func (o *PlexCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/plexes][%d] plexCollectionGetOK %s", 200, payload)
}

func (o *PlexCollectionGetOK) GetPayload() *models.PlexResponse {
	return o.Payload
}

func (o *PlexCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlexResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlexCollectionGetDefault creates a PlexCollectionGetDefault with default headers values
func NewPlexCollectionGetDefault(code int) *PlexCollectionGetDefault {
	return &PlexCollectionGetDefault{
		_statusCode: code,
	}
}

/*
PlexCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type PlexCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this plex collection get default response has a 2xx status code
func (o *PlexCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plex collection get default response has a 3xx status code
func (o *PlexCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plex collection get default response has a 4xx status code
func (o *PlexCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plex collection get default response has a 5xx status code
func (o *PlexCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plex collection get default response a status code equal to that given
func (o *PlexCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plex collection get default response
func (o *PlexCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *PlexCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/plexes][%d] plex_collection_get default %s", o._statusCode, payload)
}

func (o *PlexCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/plexes][%d] plex_collection_get default %s", o._statusCode, payload)
}

func (o *PlexCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PlexCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
