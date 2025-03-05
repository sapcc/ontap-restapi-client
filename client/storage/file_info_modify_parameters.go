// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/sapcc/ontap-restapi/models"
)

// NewFileInfoModifyParams creates a new FileInfoModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileInfoModifyParams() *FileInfoModifyParams {
	return &FileInfoModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileInfoModifyParamsWithTimeout creates a new FileInfoModifyParams object
// with the ability to set a timeout on a request.
func NewFileInfoModifyParamsWithTimeout(timeout time.Duration) *FileInfoModifyParams {
	return &FileInfoModifyParams{
		timeout: timeout,
	}
}

// NewFileInfoModifyParamsWithContext creates a new FileInfoModifyParams object
// with the ability to set a context for a request.
func NewFileInfoModifyParamsWithContext(ctx context.Context) *FileInfoModifyParams {
	return &FileInfoModifyParams{
		Context: ctx,
	}
}

// NewFileInfoModifyParamsWithHTTPClient creates a new FileInfoModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileInfoModifyParamsWithHTTPClient(client *http.Client) *FileInfoModifyParams {
	return &FileInfoModifyParams{
		HTTPClient: client,
	}
}

/*
FileInfoModifyParams contains all the parameters to send to the API endpoint

	for the file info modify operation.

	Typically these are written to a http.Request.
*/
type FileInfoModifyParams struct {

	/* ByteOffset.

	   Indicates the number of bytes into the file to begin writing. Use "-1" to append (default). Note that the byte-offset field is only supported for writing to a new or existing file, which requires specifying the Content-Type as 'multipart/form-data'.
	*/
	ByteOffset *int64

	/* Info.

	   Info specification
	*/
	Info *models.FileInfo

	/* Path.

	   Relative path of a file in the volume. The path field requires using "%2E" to represent "." and "%2F" to represent "/" for the path provided.
	*/
	Path string

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* StreamName.

	   Name of stream associated with the file to write data to.
	*/
	StreamName *string

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the file info modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileInfoModifyParams) WithDefaults() *FileInfoModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file info modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileInfoModifyParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := FileInfoModifyParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the file info modify params
func (o *FileInfoModifyParams) WithTimeout(timeout time.Duration) *FileInfoModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file info modify params
func (o *FileInfoModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file info modify params
func (o *FileInfoModifyParams) WithContext(ctx context.Context) *FileInfoModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file info modify params
func (o *FileInfoModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file info modify params
func (o *FileInfoModifyParams) WithHTTPClient(client *http.Client) *FileInfoModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file info modify params
func (o *FileInfoModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithByteOffset adds the byteOffset to the file info modify params
func (o *FileInfoModifyParams) WithByteOffset(byteOffset *int64) *FileInfoModifyParams {
	o.SetByteOffset(byteOffset)
	return o
}

// SetByteOffset adds the byteOffset to the file info modify params
func (o *FileInfoModifyParams) SetByteOffset(byteOffset *int64) {
	o.ByteOffset = byteOffset
}

// WithInfo adds the info to the file info modify params
func (o *FileInfoModifyParams) WithInfo(info *models.FileInfo) *FileInfoModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the file info modify params
func (o *FileInfoModifyParams) SetInfo(info *models.FileInfo) {
	o.Info = info
}

// WithPath adds the path to the file info modify params
func (o *FileInfoModifyParams) WithPath(path string) *FileInfoModifyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the file info modify params
func (o *FileInfoModifyParams) SetPath(path string) {
	o.Path = path
}

// WithReturnRecords adds the returnRecords to the file info modify params
func (o *FileInfoModifyParams) WithReturnRecords(returnRecords *bool) *FileInfoModifyParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the file info modify params
func (o *FileInfoModifyParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithStreamName adds the streamName to the file info modify params
func (o *FileInfoModifyParams) WithStreamName(streamName *string) *FileInfoModifyParams {
	o.SetStreamName(streamName)
	return o
}

// SetStreamName adds the streamName to the file info modify params
func (o *FileInfoModifyParams) SetStreamName(streamName *string) {
	o.StreamName = streamName
}

// WithVolumeUUID adds the volumeUUID to the file info modify params
func (o *FileInfoModifyParams) WithVolumeUUID(volumeUUID string) *FileInfoModifyParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the file info modify params
func (o *FileInfoModifyParams) SetVolumeUUID(volumeUUID string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FileInfoModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ByteOffset != nil {

		// query param byte_offset
		var qrByteOffset int64

		if o.ByteOffset != nil {
			qrByteOffset = *o.ByteOffset
		}
		qByteOffset := swag.FormatInt64(qrByteOffset)
		if qByteOffset != "" {

			if err := r.SetQueryParam("byte_offset", qByteOffset); err != nil {
				return err
			}
		}
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.StreamName != nil {

		// query param stream_name
		var qrStreamName string

		if o.StreamName != nil {
			qrStreamName = *o.StreamName
		}
		qStreamName := qrStreamName
		if qStreamName != "" {

			if err := r.SetQueryParam("stream_name", qStreamName); err != nil {
				return err
			}
		}
	}

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
