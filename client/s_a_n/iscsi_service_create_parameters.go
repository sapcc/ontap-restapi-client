// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewIscsiServiceCreateParams creates a new IscsiServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIscsiServiceCreateParams() *IscsiServiceCreateParams {
	return &IscsiServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIscsiServiceCreateParamsWithTimeout creates a new IscsiServiceCreateParams object
// with the ability to set a timeout on a request.
func NewIscsiServiceCreateParamsWithTimeout(timeout time.Duration) *IscsiServiceCreateParams {
	return &IscsiServiceCreateParams{
		timeout: timeout,
	}
}

// NewIscsiServiceCreateParamsWithContext creates a new IscsiServiceCreateParams object
// with the ability to set a context for a request.
func NewIscsiServiceCreateParamsWithContext(ctx context.Context) *IscsiServiceCreateParams {
	return &IscsiServiceCreateParams{
		Context: ctx,
	}
}

// NewIscsiServiceCreateParamsWithHTTPClient creates a new IscsiServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIscsiServiceCreateParamsWithHTTPClient(client *http.Client) *IscsiServiceCreateParams {
	return &IscsiServiceCreateParams{
		HTTPClient: client,
	}
}

/*
IscsiServiceCreateParams contains all the parameters to send to the API endpoint

	for the iscsi service create operation.

	Typically these are written to a http.Request.
*/
type IscsiServiceCreateParams struct {

	/* Info.

	   The property values for the new iSCSI service.

	*/
	Info *models.IscsiService

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iscsi service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiServiceCreateParams) WithDefaults() *IscsiServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iscsi service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiServiceCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := IscsiServiceCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the iscsi service create params
func (o *IscsiServiceCreateParams) WithTimeout(timeout time.Duration) *IscsiServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iscsi service create params
func (o *IscsiServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iscsi service create params
func (o *IscsiServiceCreateParams) WithContext(ctx context.Context) *IscsiServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iscsi service create params
func (o *IscsiServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iscsi service create params
func (o *IscsiServiceCreateParams) WithHTTPClient(client *http.Client) *IscsiServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iscsi service create params
func (o *IscsiServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the iscsi service create params
func (o *IscsiServiceCreateParams) WithInfo(info *models.IscsiService) *IscsiServiceCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the iscsi service create params
func (o *IscsiServiceCreateParams) SetInfo(info *models.IscsiService) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the iscsi service create params
func (o *IscsiServiceCreateParams) WithReturnRecords(returnRecords *bool) *IscsiServiceCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the iscsi service create params
func (o *IscsiServiceCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *IscsiServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
