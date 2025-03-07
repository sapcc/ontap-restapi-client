// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewAntiRansomwareModifyParams creates a new AntiRansomwareModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAntiRansomwareModifyParams() *AntiRansomwareModifyParams {
	return &AntiRansomwareModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAntiRansomwareModifyParamsWithTimeout creates a new AntiRansomwareModifyParams object
// with the ability to set a timeout on a request.
func NewAntiRansomwareModifyParamsWithTimeout(timeout time.Duration) *AntiRansomwareModifyParams {
	return &AntiRansomwareModifyParams{
		timeout: timeout,
	}
}

// NewAntiRansomwareModifyParamsWithContext creates a new AntiRansomwareModifyParams object
// with the ability to set a context for a request.
func NewAntiRansomwareModifyParamsWithContext(ctx context.Context) *AntiRansomwareModifyParams {
	return &AntiRansomwareModifyParams{
		Context: ctx,
	}
}

// NewAntiRansomwareModifyParamsWithHTTPClient creates a new AntiRansomwareModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAntiRansomwareModifyParamsWithHTTPClient(client *http.Client) *AntiRansomwareModifyParams {
	return &AntiRansomwareModifyParams{
		HTTPClient: client,
	}
}

/*
AntiRansomwareModifyParams contains all the parameters to send to the API endpoint

	for the anti ransomware modify operation.

	Typically these are written to a http.Request.
*/
type AntiRansomwareModifyParams struct {

	// Info.
	Info *models.AntiRansomware

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the anti ransomware modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AntiRansomwareModifyParams) WithDefaults() *AntiRansomwareModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the anti ransomware modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AntiRansomwareModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := AntiRansomwareModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the anti ransomware modify params
func (o *AntiRansomwareModifyParams) WithTimeout(timeout time.Duration) *AntiRansomwareModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the anti ransomware modify params
func (o *AntiRansomwareModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the anti ransomware modify params
func (o *AntiRansomwareModifyParams) WithContext(ctx context.Context) *AntiRansomwareModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the anti ransomware modify params
func (o *AntiRansomwareModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the anti ransomware modify params
func (o *AntiRansomwareModifyParams) WithHTTPClient(client *http.Client) *AntiRansomwareModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the anti ransomware modify params
func (o *AntiRansomwareModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the anti ransomware modify params
func (o *AntiRansomwareModifyParams) WithInfo(info *models.AntiRansomware) *AntiRansomwareModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the anti ransomware modify params
func (o *AntiRansomwareModifyParams) SetInfo(info *models.AntiRansomware) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the anti ransomware modify params
func (o *AntiRansomwareModifyParams) WithReturnTimeout(returnTimeout *int64) *AntiRansomwareModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the anti ransomware modify params
func (o *AntiRansomwareModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *AntiRansomwareModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
