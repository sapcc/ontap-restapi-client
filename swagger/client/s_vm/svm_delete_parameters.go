// Code generated by go-swagger; DO NOT EDIT.

package s_vm

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
)

// NewSvmDeleteParams creates a new SvmDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmDeleteParams() *SvmDeleteParams {
	return &SvmDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmDeleteParamsWithTimeout creates a new SvmDeleteParams object
// with the ability to set a timeout on a request.
func NewSvmDeleteParamsWithTimeout(timeout time.Duration) *SvmDeleteParams {
	return &SvmDeleteParams{
		timeout: timeout,
	}
}

// NewSvmDeleteParamsWithContext creates a new SvmDeleteParams object
// with the ability to set a context for a request.
func NewSvmDeleteParamsWithContext(ctx context.Context) *SvmDeleteParams {
	return &SvmDeleteParams{
		Context: ctx,
	}
}

// NewSvmDeleteParamsWithHTTPClient creates a new SvmDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmDeleteParamsWithHTTPClient(client *http.Client) *SvmDeleteParams {
	return &SvmDeleteParams{
		HTTPClient: client,
	}
}

/*
SvmDeleteParams contains all the parameters to send to the API endpoint

	for the svm delete operation.

	Typically these are written to a http.Request.
*/
type SvmDeleteParams struct {

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   Filter by UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmDeleteParams) WithDefaults() *SvmDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmDeleteParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := SvmDeleteParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm delete params
func (o *SvmDeleteParams) WithTimeout(timeout time.Duration) *SvmDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm delete params
func (o *SvmDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm delete params
func (o *SvmDeleteParams) WithContext(ctx context.Context) *SvmDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm delete params
func (o *SvmDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm delete params
func (o *SvmDeleteParams) WithHTTPClient(client *http.Client) *SvmDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm delete params
func (o *SvmDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnTimeout adds the returnTimeout to the svm delete params
func (o *SvmDeleteParams) WithReturnTimeout(returnTimeout *int64) *SvmDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the svm delete params
func (o *SvmDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the svm delete params
func (o *SvmDeleteParams) WithUUID(uuid string) *SvmDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the svm delete params
func (o *SvmDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SvmDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
