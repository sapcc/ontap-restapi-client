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

// NewLdapSchemaCreateParams creates a new LdapSchemaCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLdapSchemaCreateParams() *LdapSchemaCreateParams {
	return &LdapSchemaCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLdapSchemaCreateParamsWithTimeout creates a new LdapSchemaCreateParams object
// with the ability to set a timeout on a request.
func NewLdapSchemaCreateParamsWithTimeout(timeout time.Duration) *LdapSchemaCreateParams {
	return &LdapSchemaCreateParams{
		timeout: timeout,
	}
}

// NewLdapSchemaCreateParamsWithContext creates a new LdapSchemaCreateParams object
// with the ability to set a context for a request.
func NewLdapSchemaCreateParamsWithContext(ctx context.Context) *LdapSchemaCreateParams {
	return &LdapSchemaCreateParams{
		Context: ctx,
	}
}

// NewLdapSchemaCreateParamsWithHTTPClient creates a new LdapSchemaCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewLdapSchemaCreateParamsWithHTTPClient(client *http.Client) *LdapSchemaCreateParams {
	return &LdapSchemaCreateParams{
		HTTPClient: client,
	}
}

/*
LdapSchemaCreateParams contains all the parameters to send to the API endpoint

	for the ldap schema create operation.

	Typically these are written to a http.Request.
*/
type LdapSchemaCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.LdapSchema

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ldap schema create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapSchemaCreateParams) WithDefaults() *LdapSchemaCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ldap schema create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapSchemaCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := LdapSchemaCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ldap schema create params
func (o *LdapSchemaCreateParams) WithTimeout(timeout time.Duration) *LdapSchemaCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ldap schema create params
func (o *LdapSchemaCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ldap schema create params
func (o *LdapSchemaCreateParams) WithContext(ctx context.Context) *LdapSchemaCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ldap schema create params
func (o *LdapSchemaCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ldap schema create params
func (o *LdapSchemaCreateParams) WithHTTPClient(client *http.Client) *LdapSchemaCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ldap schema create params
func (o *LdapSchemaCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ldap schema create params
func (o *LdapSchemaCreateParams) WithInfo(info *models.LdapSchema) *LdapSchemaCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ldap schema create params
func (o *LdapSchemaCreateParams) SetInfo(info *models.LdapSchema) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the ldap schema create params
func (o *LdapSchemaCreateParams) WithReturnRecords(returnRecords *bool) *LdapSchemaCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ldap schema create params
func (o *LdapSchemaCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *LdapSchemaCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
