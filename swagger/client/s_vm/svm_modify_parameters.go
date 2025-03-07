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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewSvmModifyParams creates a new SvmModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmModifyParams() *SvmModifyParams {
	return &SvmModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmModifyParamsWithTimeout creates a new SvmModifyParams object
// with the ability to set a timeout on a request.
func NewSvmModifyParamsWithTimeout(timeout time.Duration) *SvmModifyParams {
	return &SvmModifyParams{
		timeout: timeout,
	}
}

// NewSvmModifyParamsWithContext creates a new SvmModifyParams object
// with the ability to set a context for a request.
func NewSvmModifyParamsWithContext(ctx context.Context) *SvmModifyParams {
	return &SvmModifyParams{
		Context: ctx,
	}
}

// NewSvmModifyParamsWithHTTPClient creates a new SvmModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmModifyParamsWithHTTPClient(client *http.Client) *SvmModifyParams {
	return &SvmModifyParams{
		HTTPClient: client,
	}
}

/*
SvmModifyParams contains all the parameters to send to the API endpoint

	for the svm modify operation.

	Typically these are written to a http.Request.
*/
type SvmModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Svm

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

// WithDefaults hydrates default values in the svm modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmModifyParams) WithDefaults() *SvmModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := SvmModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm modify params
func (o *SvmModifyParams) WithTimeout(timeout time.Duration) *SvmModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm modify params
func (o *SvmModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm modify params
func (o *SvmModifyParams) WithContext(ctx context.Context) *SvmModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm modify params
func (o *SvmModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm modify params
func (o *SvmModifyParams) WithHTTPClient(client *http.Client) *SvmModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm modify params
func (o *SvmModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the svm modify params
func (o *SvmModifyParams) WithInfo(info *models.Svm) *SvmModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the svm modify params
func (o *SvmModifyParams) SetInfo(info *models.Svm) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the svm modify params
func (o *SvmModifyParams) WithReturnTimeout(returnTimeout *int64) *SvmModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the svm modify params
func (o *SvmModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the svm modify params
func (o *SvmModifyParams) WithUUID(uuid string) *SvmModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the svm modify params
func (o *SvmModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SvmModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
