// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewSwitchModifyParams creates a new SwitchModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSwitchModifyParams() *SwitchModifyParams {
	return &SwitchModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSwitchModifyParamsWithTimeout creates a new SwitchModifyParams object
// with the ability to set a timeout on a request.
func NewSwitchModifyParamsWithTimeout(timeout time.Duration) *SwitchModifyParams {
	return &SwitchModifyParams{
		timeout: timeout,
	}
}

// NewSwitchModifyParamsWithContext creates a new SwitchModifyParams object
// with the ability to set a context for a request.
func NewSwitchModifyParamsWithContext(ctx context.Context) *SwitchModifyParams {
	return &SwitchModifyParams{
		Context: ctx,
	}
}

// NewSwitchModifyParamsWithHTTPClient creates a new SwitchModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSwitchModifyParamsWithHTTPClient(client *http.Client) *SwitchModifyParams {
	return &SwitchModifyParams{
		HTTPClient: client,
	}
}

/*
SwitchModifyParams contains all the parameters to send to the API endpoint

	for the switch modify operation.

	Typically these are written to a http.Request.
*/
type SwitchModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Switch

	/* Name.

	   Switch Name
	*/
	Name string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the switch modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SwitchModifyParams) WithDefaults() *SwitchModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the switch modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SwitchModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := SwitchModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the switch modify params
func (o *SwitchModifyParams) WithTimeout(timeout time.Duration) *SwitchModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the switch modify params
func (o *SwitchModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the switch modify params
func (o *SwitchModifyParams) WithContext(ctx context.Context) *SwitchModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the switch modify params
func (o *SwitchModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the switch modify params
func (o *SwitchModifyParams) WithHTTPClient(client *http.Client) *SwitchModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the switch modify params
func (o *SwitchModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the switch modify params
func (o *SwitchModifyParams) WithInfo(info *models.Switch) *SwitchModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the switch modify params
func (o *SwitchModifyParams) SetInfo(info *models.Switch) {
	o.Info = info
}

// WithName adds the name to the switch modify params
func (o *SwitchModifyParams) WithName(name string) *SwitchModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the switch modify params
func (o *SwitchModifyParams) SetName(name string) {
	o.Name = name
}

// WithReturnTimeout adds the returnTimeout to the switch modify params
func (o *SwitchModifyParams) WithReturnTimeout(returnTimeout *int64) *SwitchModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the switch modify params
func (o *SwitchModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *SwitchModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
