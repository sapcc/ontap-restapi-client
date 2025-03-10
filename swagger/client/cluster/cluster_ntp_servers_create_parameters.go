// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewClusterNtpServersCreateParams creates a new ClusterNtpServersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterNtpServersCreateParams() *ClusterNtpServersCreateParams {
	return &ClusterNtpServersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterNtpServersCreateParamsWithTimeout creates a new ClusterNtpServersCreateParams object
// with the ability to set a timeout on a request.
func NewClusterNtpServersCreateParamsWithTimeout(timeout time.Duration) *ClusterNtpServersCreateParams {
	return &ClusterNtpServersCreateParams{
		timeout: timeout,
	}
}

// NewClusterNtpServersCreateParamsWithContext creates a new ClusterNtpServersCreateParams object
// with the ability to set a context for a request.
func NewClusterNtpServersCreateParamsWithContext(ctx context.Context) *ClusterNtpServersCreateParams {
	return &ClusterNtpServersCreateParams{
		Context: ctx,
	}
}

// NewClusterNtpServersCreateParamsWithHTTPClient creates a new ClusterNtpServersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterNtpServersCreateParamsWithHTTPClient(client *http.Client) *ClusterNtpServersCreateParams {
	return &ClusterNtpServersCreateParams{
		HTTPClient: client,
	}
}

/*
ClusterNtpServersCreateParams contains all the parameters to send to the API endpoint

	for the cluster ntp servers create operation.

	Typically these are written to a http.Request.
*/
type ClusterNtpServersCreateParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.NtpServer

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

// WithDefaults hydrates default values in the cluster ntp servers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNtpServersCreateParams) WithDefaults() *ClusterNtpServersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster ntp servers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNtpServersCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := ClusterNtpServersCreateParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) WithTimeout(timeout time.Duration) *ClusterNtpServersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) WithContext(ctx context.Context) *ClusterNtpServersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) WithHTTPClient(client *http.Client) *ClusterNtpServersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) WithInfo(info *models.NtpServer) *ClusterNtpServersCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) SetInfo(info *models.NtpServer) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) WithReturnRecords(returnRecords *bool) *ClusterNtpServersCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) WithReturnTimeout(returnTimeout *int64) *ClusterNtpServersCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cluster ntp servers create params
func (o *ClusterNtpServersCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterNtpServersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
