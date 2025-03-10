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

// NewLicenseCreateParams creates a new LicenseCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLicenseCreateParams() *LicenseCreateParams {
	return &LicenseCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLicenseCreateParamsWithTimeout creates a new LicenseCreateParams object
// with the ability to set a timeout on a request.
func NewLicenseCreateParamsWithTimeout(timeout time.Duration) *LicenseCreateParams {
	return &LicenseCreateParams{
		timeout: timeout,
	}
}

// NewLicenseCreateParamsWithContext creates a new LicenseCreateParams object
// with the ability to set a context for a request.
func NewLicenseCreateParamsWithContext(ctx context.Context) *LicenseCreateParams {
	return &LicenseCreateParams{
		Context: ctx,
	}
}

// NewLicenseCreateParamsWithHTTPClient creates a new LicenseCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewLicenseCreateParamsWithHTTPClient(client *http.Client) *LicenseCreateParams {
	return &LicenseCreateParams{
		HTTPClient: client,
	}
}

/*
LicenseCreateParams contains all the parameters to send to the API endpoint

	for the license create operation.

	Typically these are written to a http.Request.
*/
type LicenseCreateParams struct {

	/* Info.

	   List of NLF or 26-character keys.
	*/
	Info *models.LicensePackage

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the license create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicenseCreateParams) WithDefaults() *LicenseCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the license create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicenseCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := LicenseCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the license create params
func (o *LicenseCreateParams) WithTimeout(timeout time.Duration) *LicenseCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the license create params
func (o *LicenseCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the license create params
func (o *LicenseCreateParams) WithContext(ctx context.Context) *LicenseCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the license create params
func (o *LicenseCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the license create params
func (o *LicenseCreateParams) WithHTTPClient(client *http.Client) *LicenseCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the license create params
func (o *LicenseCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the license create params
func (o *LicenseCreateParams) WithInfo(info *models.LicensePackage) *LicenseCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the license create params
func (o *LicenseCreateParams) SetInfo(info *models.LicensePackage) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the license create params
func (o *LicenseCreateParams) WithReturnRecords(returnRecords *bool) *LicenseCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the license create params
func (o *LicenseCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *LicenseCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
