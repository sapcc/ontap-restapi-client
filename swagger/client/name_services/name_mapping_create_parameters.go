// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewNameMappingCreateParams creates a new NameMappingCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNameMappingCreateParams() *NameMappingCreateParams {
	return &NameMappingCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNameMappingCreateParamsWithTimeout creates a new NameMappingCreateParams object
// with the ability to set a timeout on a request.
func NewNameMappingCreateParamsWithTimeout(timeout time.Duration) *NameMappingCreateParams {
	return &NameMappingCreateParams{
		timeout: timeout,
	}
}

// NewNameMappingCreateParamsWithContext creates a new NameMappingCreateParams object
// with the ability to set a context for a request.
func NewNameMappingCreateParamsWithContext(ctx context.Context) *NameMappingCreateParams {
	return &NameMappingCreateParams{
		Context: ctx,
	}
}

// NewNameMappingCreateParamsWithHTTPClient creates a new NameMappingCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNameMappingCreateParamsWithHTTPClient(client *http.Client) *NameMappingCreateParams {
	return &NameMappingCreateParams{
		HTTPClient: client,
	}
}

/*
NameMappingCreateParams contains all the parameters to send to the API endpoint

	for the name mapping create operation.

	Typically these are written to a http.Request.
*/
type NameMappingCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.NameMapping

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the name mapping create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameMappingCreateParams) WithDefaults() *NameMappingCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the name mapping create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameMappingCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := NameMappingCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the name mapping create params
func (o *NameMappingCreateParams) WithTimeout(timeout time.Duration) *NameMappingCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the name mapping create params
func (o *NameMappingCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the name mapping create params
func (o *NameMappingCreateParams) WithContext(ctx context.Context) *NameMappingCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the name mapping create params
func (o *NameMappingCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the name mapping create params
func (o *NameMappingCreateParams) WithHTTPClient(client *http.Client) *NameMappingCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the name mapping create params
func (o *NameMappingCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the name mapping create params
func (o *NameMappingCreateParams) WithInfo(info *models.NameMapping) *NameMappingCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the name mapping create params
func (o *NameMappingCreateParams) SetInfo(info *models.NameMapping) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the name mapping create params
func (o *NameMappingCreateParams) WithReturnRecords(returnRecords *bool) *NameMappingCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the name mapping create params
func (o *NameMappingCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *NameMappingCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
