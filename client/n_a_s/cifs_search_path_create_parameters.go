// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NewCifsSearchPathCreateParams creates a new CifsSearchPathCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsSearchPathCreateParams() *CifsSearchPathCreateParams {
	return &CifsSearchPathCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsSearchPathCreateParamsWithTimeout creates a new CifsSearchPathCreateParams object
// with the ability to set a timeout on a request.
func NewCifsSearchPathCreateParamsWithTimeout(timeout time.Duration) *CifsSearchPathCreateParams {
	return &CifsSearchPathCreateParams{
		timeout: timeout,
	}
}

// NewCifsSearchPathCreateParamsWithContext creates a new CifsSearchPathCreateParams object
// with the ability to set a context for a request.
func NewCifsSearchPathCreateParamsWithContext(ctx context.Context) *CifsSearchPathCreateParams {
	return &CifsSearchPathCreateParams{
		Context: ctx,
	}
}

// NewCifsSearchPathCreateParamsWithHTTPClient creates a new CifsSearchPathCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsSearchPathCreateParamsWithHTTPClient(client *http.Client) *CifsSearchPathCreateParams {
	return &CifsSearchPathCreateParams{
		HTTPClient: client,
	}
}

/*
CifsSearchPathCreateParams contains all the parameters to send to the API endpoint

	for the cifs search path create operation.

	Typically these are written to a http.Request.
*/
type CifsSearchPathCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.CifsSearchPath

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs search path create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSearchPathCreateParams) WithDefaults() *CifsSearchPathCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs search path create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSearchPathCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := CifsSearchPathCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs search path create params
func (o *CifsSearchPathCreateParams) WithTimeout(timeout time.Duration) *CifsSearchPathCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs search path create params
func (o *CifsSearchPathCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs search path create params
func (o *CifsSearchPathCreateParams) WithContext(ctx context.Context) *CifsSearchPathCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs search path create params
func (o *CifsSearchPathCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs search path create params
func (o *CifsSearchPathCreateParams) WithHTTPClient(client *http.Client) *CifsSearchPathCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs search path create params
func (o *CifsSearchPathCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cifs search path create params
func (o *CifsSearchPathCreateParams) WithInfo(info *models.CifsSearchPath) *CifsSearchPathCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cifs search path create params
func (o *CifsSearchPathCreateParams) SetInfo(info *models.CifsSearchPath) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the cifs search path create params
func (o *CifsSearchPathCreateParams) WithReturnRecords(returnRecords *bool) *CifsSearchPathCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the cifs search path create params
func (o *CifsSearchPathCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *CifsSearchPathCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
