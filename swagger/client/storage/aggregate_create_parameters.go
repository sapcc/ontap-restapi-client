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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewAggregateCreateParams creates a new AggregateCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateCreateParams() *AggregateCreateParams {
	return &AggregateCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateCreateParamsWithTimeout creates a new AggregateCreateParams object
// with the ability to set a timeout on a request.
func NewAggregateCreateParamsWithTimeout(timeout time.Duration) *AggregateCreateParams {
	return &AggregateCreateParams{
		timeout: timeout,
	}
}

// NewAggregateCreateParamsWithContext creates a new AggregateCreateParams object
// with the ability to set a context for a request.
func NewAggregateCreateParamsWithContext(ctx context.Context) *AggregateCreateParams {
	return &AggregateCreateParams{
		Context: ctx,
	}
}

// NewAggregateCreateParamsWithHTTPClient creates a new AggregateCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateCreateParamsWithHTTPClient(client *http.Client) *AggregateCreateParams {
	return &AggregateCreateParams{
		HTTPClient: client,
	}
}

/*
AggregateCreateParams contains all the parameters to send to the API endpoint

	for the aggregate create operation.

	Typically these are written to a http.Request.
*/
type AggregateCreateParams struct {

	/* DiskSize.

	   If set, POST only selects disks of the specified size.
	*/
	DiskSize *int64

	/* Info.

	   Info specification
	*/
	Info *models.Aggregate

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateCreateParams) WithDefaults() *AggregateCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := AggregateCreateParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the aggregate create params
func (o *AggregateCreateParams) WithTimeout(timeout time.Duration) *AggregateCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate create params
func (o *AggregateCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate create params
func (o *AggregateCreateParams) WithContext(ctx context.Context) *AggregateCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate create params
func (o *AggregateCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate create params
func (o *AggregateCreateParams) WithHTTPClient(client *http.Client) *AggregateCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate create params
func (o *AggregateCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDiskSize adds the diskSize to the aggregate create params
func (o *AggregateCreateParams) WithDiskSize(diskSize *int64) *AggregateCreateParams {
	o.SetDiskSize(diskSize)
	return o
}

// SetDiskSize adds the diskSize to the aggregate create params
func (o *AggregateCreateParams) SetDiskSize(diskSize *int64) {
	o.DiskSize = diskSize
}

// WithInfo adds the info to the aggregate create params
func (o *AggregateCreateParams) WithInfo(info *models.Aggregate) *AggregateCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the aggregate create params
func (o *AggregateCreateParams) SetInfo(info *models.Aggregate) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the aggregate create params
func (o *AggregateCreateParams) WithReturnRecords(returnRecords *bool) *AggregateCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the aggregate create params
func (o *AggregateCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the aggregate create params
func (o *AggregateCreateParams) WithReturnTimeout(returnTimeout *int64) *AggregateCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the aggregate create params
func (o *AggregateCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DiskSize != nil {

		// query param disk_size
		var qrDiskSize int64

		if o.DiskSize != nil {
			qrDiskSize = *o.DiskSize
		}
		qDiskSize := swag.FormatInt64(qrDiskSize)
		if qDiskSize != "" {

			if err := r.SetQueryParam("disk_size", qDiskSize); err != nil {
				return err
			}
		}
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
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

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
