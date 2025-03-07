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

// NewUnixGroupCreateParams creates a new UnixGroupCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixGroupCreateParams() *UnixGroupCreateParams {
	return &UnixGroupCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixGroupCreateParamsWithTimeout creates a new UnixGroupCreateParams object
// with the ability to set a timeout on a request.
func NewUnixGroupCreateParamsWithTimeout(timeout time.Duration) *UnixGroupCreateParams {
	return &UnixGroupCreateParams{
		timeout: timeout,
	}
}

// NewUnixGroupCreateParamsWithContext creates a new UnixGroupCreateParams object
// with the ability to set a context for a request.
func NewUnixGroupCreateParamsWithContext(ctx context.Context) *UnixGroupCreateParams {
	return &UnixGroupCreateParams{
		Context: ctx,
	}
}

// NewUnixGroupCreateParamsWithHTTPClient creates a new UnixGroupCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixGroupCreateParamsWithHTTPClient(client *http.Client) *UnixGroupCreateParams {
	return &UnixGroupCreateParams{
		HTTPClient: client,
	}
}

/*
UnixGroupCreateParams contains all the parameters to send to the API endpoint

	for the unix group create operation.

	Typically these are written to a http.Request.
*/
type UnixGroupCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.UnixGroup

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unix group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupCreateParams) WithDefaults() *UnixGroupCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := UnixGroupCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the unix group create params
func (o *UnixGroupCreateParams) WithTimeout(timeout time.Duration) *UnixGroupCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix group create params
func (o *UnixGroupCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix group create params
func (o *UnixGroupCreateParams) WithContext(ctx context.Context) *UnixGroupCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix group create params
func (o *UnixGroupCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix group create params
func (o *UnixGroupCreateParams) WithHTTPClient(client *http.Client) *UnixGroupCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix group create params
func (o *UnixGroupCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the unix group create params
func (o *UnixGroupCreateParams) WithInfo(info *models.UnixGroup) *UnixGroupCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the unix group create params
func (o *UnixGroupCreateParams) SetInfo(info *models.UnixGroup) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the unix group create params
func (o *UnixGroupCreateParams) WithReturnRecords(returnRecords *bool) *UnixGroupCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the unix group create params
func (o *UnixGroupCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *UnixGroupCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
