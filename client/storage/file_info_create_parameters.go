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

// NewFileInfoCreateParams creates a new FileInfoCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileInfoCreateParams() *FileInfoCreateParams {
	return &FileInfoCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileInfoCreateParamsWithTimeout creates a new FileInfoCreateParams object
// with the ability to set a timeout on a request.
func NewFileInfoCreateParamsWithTimeout(timeout time.Duration) *FileInfoCreateParams {
	return &FileInfoCreateParams{
		timeout: timeout,
	}
}

// NewFileInfoCreateParamsWithContext creates a new FileInfoCreateParams object
// with the ability to set a context for a request.
func NewFileInfoCreateParamsWithContext(ctx context.Context) *FileInfoCreateParams {
	return &FileInfoCreateParams{
		Context: ctx,
	}
}

// NewFileInfoCreateParamsWithHTTPClient creates a new FileInfoCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileInfoCreateParamsWithHTTPClient(client *http.Client) *FileInfoCreateParams {
	return &FileInfoCreateParams{
		HTTPClient: client,
	}
}

/*
FileInfoCreateParams contains all the parameters to send to the API endpoint

	for the file info create operation.

	Typically these are written to a http.Request.
*/
type FileInfoCreateParams struct {

	/* ByteOffset.

	   Indicates the number of bytes into the file to begin writing. Use "-1" to append (default). Note that the byte-offset field is only supported for writing to a new or existing file, which requires specifying the Content-Type as 'multipart/form-data'.
	*/
	ByteOffset *int64

	/* Info.

	   Info specification
	*/
	Info *models.FileInfo

	/* Overwrite.

	   If false, and the file exists, the write will fail. Default is false.
	*/
	Overwrite *bool

	/* Path.

	   Relative path of a new file, directory or symlink. The path field requires using "%2E" to represent "." and "%2F" to represent "/" for the path provided.
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

// WithDefaults hydrates default values in the file info create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileInfoCreateParams) WithDefaults() *FileInfoCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file info create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileInfoCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := FileInfoCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the file info create params
func (o *FileInfoCreateParams) WithTimeout(timeout time.Duration) *FileInfoCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file info create params
func (o *FileInfoCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file info create params
func (o *FileInfoCreateParams) WithContext(ctx context.Context) *FileInfoCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file info create params
func (o *FileInfoCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file info create params
func (o *FileInfoCreateParams) WithHTTPClient(client *http.Client) *FileInfoCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file info create params
func (o *FileInfoCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithByteOffset adds the byteOffset to the file info create params
func (o *FileInfoCreateParams) WithByteOffset(byteOffset *int64) *FileInfoCreateParams {
	o.SetByteOffset(byteOffset)
	return o
}

// SetByteOffset adds the byteOffset to the file info create params
func (o *FileInfoCreateParams) SetByteOffset(byteOffset *int64) {
	o.ByteOffset = byteOffset
}

// WithInfo adds the info to the file info create params
func (o *FileInfoCreateParams) WithInfo(info *models.FileInfo) *FileInfoCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the file info create params
func (o *FileInfoCreateParams) SetInfo(info *models.FileInfo) {
	o.Info = info
}

// WithOverwrite adds the overwrite to the file info create params
func (o *FileInfoCreateParams) WithOverwrite(overwrite *bool) *FileInfoCreateParams {
	o.SetOverwrite(overwrite)
	return o
}

// SetOverwrite adds the overwrite to the file info create params
func (o *FileInfoCreateParams) SetOverwrite(overwrite *bool) {
	o.Overwrite = overwrite
}

// WithPath adds the path to the file info create params
func (o *FileInfoCreateParams) WithPath(path string) *FileInfoCreateParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the file info create params
func (o *FileInfoCreateParams) SetPath(path string) {
	o.Path = path
}

// WithReturnRecords adds the returnRecords to the file info create params
func (o *FileInfoCreateParams) WithReturnRecords(returnRecords *bool) *FileInfoCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the file info create params
func (o *FileInfoCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithStreamName adds the streamName to the file info create params
func (o *FileInfoCreateParams) WithStreamName(streamName *string) *FileInfoCreateParams {
	o.SetStreamName(streamName)
	return o
}

// SetStreamName adds the streamName to the file info create params
func (o *FileInfoCreateParams) SetStreamName(streamName *string) {
	o.StreamName = streamName
}

// WithVolumeUUID adds the volumeUUID to the file info create params
func (o *FileInfoCreateParams) WithVolumeUUID(volumeUUID string) *FileInfoCreateParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the file info create params
func (o *FileInfoCreateParams) SetVolumeUUID(volumeUUID string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FileInfoCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Overwrite != nil {

		// query param overwrite
		var qrOverwrite bool

		if o.Overwrite != nil {
			qrOverwrite = *o.Overwrite
		}
		qOverwrite := swag.FormatBool(qrOverwrite)
		if qOverwrite != "" {

			if err := r.SetQueryParam("overwrite", qOverwrite); err != nil {
				return err
			}
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
