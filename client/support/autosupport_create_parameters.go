// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewAutosupportCreateParams creates a new AutosupportCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAutosupportCreateParams() *AutosupportCreateParams {
	return &AutosupportCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAutosupportCreateParamsWithTimeout creates a new AutosupportCreateParams object
// with the ability to set a timeout on a request.
func NewAutosupportCreateParamsWithTimeout(timeout time.Duration) *AutosupportCreateParams {
	return &AutosupportCreateParams{
		timeout: timeout,
	}
}

// NewAutosupportCreateParamsWithContext creates a new AutosupportCreateParams object
// with the ability to set a context for a request.
func NewAutosupportCreateParamsWithContext(ctx context.Context) *AutosupportCreateParams {
	return &AutosupportCreateParams{
		Context: ctx,
	}
}

// NewAutosupportCreateParamsWithHTTPClient creates a new AutosupportCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAutosupportCreateParamsWithHTTPClient(client *http.Client) *AutosupportCreateParams {
	return &AutosupportCreateParams{
		HTTPClient: client,
	}
}

/*
AutosupportCreateParams contains all the parameters to send to the API endpoint

	for the autosupport create operation.

	Typically these are written to a http.Request.
*/
type AutosupportCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.AutosupportMessage

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the autosupport create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutosupportCreateParams) WithDefaults() *AutosupportCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the autosupport create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutosupportCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := AutosupportCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the autosupport create params
func (o *AutosupportCreateParams) WithTimeout(timeout time.Duration) *AutosupportCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the autosupport create params
func (o *AutosupportCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the autosupport create params
func (o *AutosupportCreateParams) WithContext(ctx context.Context) *AutosupportCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the autosupport create params
func (o *AutosupportCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the autosupport create params
func (o *AutosupportCreateParams) WithHTTPClient(client *http.Client) *AutosupportCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the autosupport create params
func (o *AutosupportCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the autosupport create params
func (o *AutosupportCreateParams) WithInfo(info *models.AutosupportMessage) *AutosupportCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the autosupport create params
func (o *AutosupportCreateParams) SetInfo(info *models.AutosupportMessage) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the autosupport create params
func (o *AutosupportCreateParams) WithReturnRecords(returnRecords *bool) *AutosupportCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the autosupport create params
func (o *AutosupportCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *AutosupportCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
