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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewNfsCreateParams creates a new NfsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNfsCreateParams() *NfsCreateParams {
	return &NfsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNfsCreateParamsWithTimeout creates a new NfsCreateParams object
// with the ability to set a timeout on a request.
func NewNfsCreateParamsWithTimeout(timeout time.Duration) *NfsCreateParams {
	return &NfsCreateParams{
		timeout: timeout,
	}
}

// NewNfsCreateParamsWithContext creates a new NfsCreateParams object
// with the ability to set a context for a request.
func NewNfsCreateParamsWithContext(ctx context.Context) *NfsCreateParams {
	return &NfsCreateParams{
		Context: ctx,
	}
}

// NewNfsCreateParamsWithHTTPClient creates a new NfsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNfsCreateParamsWithHTTPClient(client *http.Client) *NfsCreateParams {
	return &NfsCreateParams{
		HTTPClient: client,
	}
}

/*
NfsCreateParams contains all the parameters to send to the API endpoint

	for the nfs create operation.

	Typically these are written to a http.Request.
*/
type NfsCreateParams struct {

	/* Info.

	   Info Specification
	*/
	Info *models.NfsService

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nfs create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsCreateParams) WithDefaults() *NfsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nfs create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := NfsCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nfs create params
func (o *NfsCreateParams) WithTimeout(timeout time.Duration) *NfsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nfs create params
func (o *NfsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nfs create params
func (o *NfsCreateParams) WithContext(ctx context.Context) *NfsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nfs create params
func (o *NfsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nfs create params
func (o *NfsCreateParams) WithHTTPClient(client *http.Client) *NfsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nfs create params
func (o *NfsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the nfs create params
func (o *NfsCreateParams) WithInfo(info *models.NfsService) *NfsCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the nfs create params
func (o *NfsCreateParams) SetInfo(info *models.NfsService) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the nfs create params
func (o *NfsCreateParams) WithReturnRecords(returnRecords *bool) *NfsCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the nfs create params
func (o *NfsCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *NfsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
