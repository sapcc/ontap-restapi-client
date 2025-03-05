// Code generated by go-swagger; DO NOT EDIT.

package snap_lock

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

// NewSnaplockLegalHoldBeginParams creates a new SnaplockLegalHoldBeginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLegalHoldBeginParams() *SnaplockLegalHoldBeginParams {
	return &SnaplockLegalHoldBeginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLegalHoldBeginParamsWithTimeout creates a new SnaplockLegalHoldBeginParams object
// with the ability to set a timeout on a request.
func NewSnaplockLegalHoldBeginParamsWithTimeout(timeout time.Duration) *SnaplockLegalHoldBeginParams {
	return &SnaplockLegalHoldBeginParams{
		timeout: timeout,
	}
}

// NewSnaplockLegalHoldBeginParamsWithContext creates a new SnaplockLegalHoldBeginParams object
// with the ability to set a context for a request.
func NewSnaplockLegalHoldBeginParamsWithContext(ctx context.Context) *SnaplockLegalHoldBeginParams {
	return &SnaplockLegalHoldBeginParams{
		Context: ctx,
	}
}

// NewSnaplockLegalHoldBeginParamsWithHTTPClient creates a new SnaplockLegalHoldBeginParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLegalHoldBeginParamsWithHTTPClient(client *http.Client) *SnaplockLegalHoldBeginParams {
	return &SnaplockLegalHoldBeginParams{
		HTTPClient: client,
	}
}

/*
SnaplockLegalHoldBeginParams contains all the parameters to send to the API endpoint

	for the snaplock legal hold begin operation.

	Typically these are written to a http.Request.
*/
type SnaplockLegalHoldBeginParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SnaplockLitigation

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock legal hold begin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldBeginParams) WithDefaults() *SnaplockLegalHoldBeginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock legal hold begin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldBeginParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := SnaplockLegalHoldBeginParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snaplock legal hold begin params
func (o *SnaplockLegalHoldBeginParams) WithTimeout(timeout time.Duration) *SnaplockLegalHoldBeginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock legal hold begin params
func (o *SnaplockLegalHoldBeginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock legal hold begin params
func (o *SnaplockLegalHoldBeginParams) WithContext(ctx context.Context) *SnaplockLegalHoldBeginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock legal hold begin params
func (o *SnaplockLegalHoldBeginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock legal hold begin params
func (o *SnaplockLegalHoldBeginParams) WithHTTPClient(client *http.Client) *SnaplockLegalHoldBeginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock legal hold begin params
func (o *SnaplockLegalHoldBeginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snaplock legal hold begin params
func (o *SnaplockLegalHoldBeginParams) WithInfo(info *models.SnaplockLitigation) *SnaplockLegalHoldBeginParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snaplock legal hold begin params
func (o *SnaplockLegalHoldBeginParams) SetInfo(info *models.SnaplockLitigation) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the snaplock legal hold begin params
func (o *SnaplockLegalHoldBeginParams) WithReturnRecords(returnRecords *bool) *SnaplockLegalHoldBeginParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snaplock legal hold begin params
func (o *SnaplockLegalHoldBeginParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLegalHoldBeginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
