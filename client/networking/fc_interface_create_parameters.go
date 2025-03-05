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

	"github.com/sapcc/ontap-restapi/models"
)

// NewFcInterfaceCreateParams creates a new FcInterfaceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFcInterfaceCreateParams() *FcInterfaceCreateParams {
	return &FcInterfaceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFcInterfaceCreateParamsWithTimeout creates a new FcInterfaceCreateParams object
// with the ability to set a timeout on a request.
func NewFcInterfaceCreateParamsWithTimeout(timeout time.Duration) *FcInterfaceCreateParams {
	return &FcInterfaceCreateParams{
		timeout: timeout,
	}
}

// NewFcInterfaceCreateParamsWithContext creates a new FcInterfaceCreateParams object
// with the ability to set a context for a request.
func NewFcInterfaceCreateParamsWithContext(ctx context.Context) *FcInterfaceCreateParams {
	return &FcInterfaceCreateParams{
		Context: ctx,
	}
}

// NewFcInterfaceCreateParamsWithHTTPClient creates a new FcInterfaceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFcInterfaceCreateParamsWithHTTPClient(client *http.Client) *FcInterfaceCreateParams {
	return &FcInterfaceCreateParams{
		HTTPClient: client,
	}
}

/*
FcInterfaceCreateParams contains all the parameters to send to the API endpoint

	for the fc interface create operation.

	Typically these are written to a http.Request.
*/
type FcInterfaceCreateParams struct {

	/* Info.

	   The property values for the new FC interface.

	*/
	Info *models.FcInterface

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fc interface create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcInterfaceCreateParams) WithDefaults() *FcInterfaceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fc interface create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcInterfaceCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := FcInterfaceCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fc interface create params
func (o *FcInterfaceCreateParams) WithTimeout(timeout time.Duration) *FcInterfaceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fc interface create params
func (o *FcInterfaceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fc interface create params
func (o *FcInterfaceCreateParams) WithContext(ctx context.Context) *FcInterfaceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fc interface create params
func (o *FcInterfaceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fc interface create params
func (o *FcInterfaceCreateParams) WithHTTPClient(client *http.Client) *FcInterfaceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fc interface create params
func (o *FcInterfaceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the fc interface create params
func (o *FcInterfaceCreateParams) WithInfo(info *models.FcInterface) *FcInterfaceCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the fc interface create params
func (o *FcInterfaceCreateParams) SetInfo(info *models.FcInterface) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the fc interface create params
func (o *FcInterfaceCreateParams) WithReturnRecords(returnRecords *bool) *FcInterfaceCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the fc interface create params
func (o *FcInterfaceCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *FcInterfaceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
