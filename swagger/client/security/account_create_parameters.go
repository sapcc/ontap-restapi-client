// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewAccountCreateParams creates a new AccountCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountCreateParams() *AccountCreateParams {
	return &AccountCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountCreateParamsWithTimeout creates a new AccountCreateParams object
// with the ability to set a timeout on a request.
func NewAccountCreateParamsWithTimeout(timeout time.Duration) *AccountCreateParams {
	return &AccountCreateParams{
		timeout: timeout,
	}
}

// NewAccountCreateParamsWithContext creates a new AccountCreateParams object
// with the ability to set a context for a request.
func NewAccountCreateParamsWithContext(ctx context.Context) *AccountCreateParams {
	return &AccountCreateParams{
		Context: ctx,
	}
}

// NewAccountCreateParamsWithHTTPClient creates a new AccountCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountCreateParamsWithHTTPClient(client *http.Client) *AccountCreateParams {
	return &AccountCreateParams{
		HTTPClient: client,
	}
}

/*
AccountCreateParams contains all the parameters to send to the API endpoint

	for the account create operation.

	Typically these are written to a http.Request.
*/
type AccountCreateParams struct {

	/* Info.

	   Details for the user account to be created.
	*/
	Info *models.Account

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountCreateParams) WithDefaults() *AccountCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := AccountCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the account create params
func (o *AccountCreateParams) WithTimeout(timeout time.Duration) *AccountCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account create params
func (o *AccountCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account create params
func (o *AccountCreateParams) WithContext(ctx context.Context) *AccountCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account create params
func (o *AccountCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account create params
func (o *AccountCreateParams) WithHTTPClient(client *http.Client) *AccountCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account create params
func (o *AccountCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the account create params
func (o *AccountCreateParams) WithInfo(info *models.Account) *AccountCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the account create params
func (o *AccountCreateParams) SetInfo(info *models.Account) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the account create params
func (o *AccountCreateParams) WithReturnRecords(returnRecords *bool) *AccountCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the account create params
func (o *AccountCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *AccountCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
