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

// NewIscsiCredentialsCreateParams creates a new IscsiCredentialsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIscsiCredentialsCreateParams() *IscsiCredentialsCreateParams {
	return &IscsiCredentialsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIscsiCredentialsCreateParamsWithTimeout creates a new IscsiCredentialsCreateParams object
// with the ability to set a timeout on a request.
func NewIscsiCredentialsCreateParamsWithTimeout(timeout time.Duration) *IscsiCredentialsCreateParams {
	return &IscsiCredentialsCreateParams{
		timeout: timeout,
	}
}

// NewIscsiCredentialsCreateParamsWithContext creates a new IscsiCredentialsCreateParams object
// with the ability to set a context for a request.
func NewIscsiCredentialsCreateParamsWithContext(ctx context.Context) *IscsiCredentialsCreateParams {
	return &IscsiCredentialsCreateParams{
		Context: ctx,
	}
}

// NewIscsiCredentialsCreateParamsWithHTTPClient creates a new IscsiCredentialsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIscsiCredentialsCreateParamsWithHTTPClient(client *http.Client) *IscsiCredentialsCreateParams {
	return &IscsiCredentialsCreateParams{
		HTTPClient: client,
	}
}

/*
IscsiCredentialsCreateParams contains all the parameters to send to the API endpoint

	for the iscsi credentials create operation.

	Typically these are written to a http.Request.
*/
type IscsiCredentialsCreateParams struct {

	/* Info.

	   The property values for the new iSCSI credentials object.

	*/
	Info *models.IscsiCredentials

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iscsi credentials create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiCredentialsCreateParams) WithDefaults() *IscsiCredentialsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iscsi credentials create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiCredentialsCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := IscsiCredentialsCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the iscsi credentials create params
func (o *IscsiCredentialsCreateParams) WithTimeout(timeout time.Duration) *IscsiCredentialsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iscsi credentials create params
func (o *IscsiCredentialsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iscsi credentials create params
func (o *IscsiCredentialsCreateParams) WithContext(ctx context.Context) *IscsiCredentialsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iscsi credentials create params
func (o *IscsiCredentialsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iscsi credentials create params
func (o *IscsiCredentialsCreateParams) WithHTTPClient(client *http.Client) *IscsiCredentialsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iscsi credentials create params
func (o *IscsiCredentialsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the iscsi credentials create params
func (o *IscsiCredentialsCreateParams) WithInfo(info *models.IscsiCredentials) *IscsiCredentialsCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the iscsi credentials create params
func (o *IscsiCredentialsCreateParams) SetInfo(info *models.IscsiCredentials) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the iscsi credentials create params
func (o *IscsiCredentialsCreateParams) WithReturnRecords(returnRecords *bool) *IscsiCredentialsCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the iscsi credentials create params
func (o *IscsiCredentialsCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *IscsiCredentialsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
