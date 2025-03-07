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

// NewIpsecCaCertificateCreateParams creates a new IpsecCaCertificateCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpsecCaCertificateCreateParams() *IpsecCaCertificateCreateParams {
	return &IpsecCaCertificateCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpsecCaCertificateCreateParamsWithTimeout creates a new IpsecCaCertificateCreateParams object
// with the ability to set a timeout on a request.
func NewIpsecCaCertificateCreateParamsWithTimeout(timeout time.Duration) *IpsecCaCertificateCreateParams {
	return &IpsecCaCertificateCreateParams{
		timeout: timeout,
	}
}

// NewIpsecCaCertificateCreateParamsWithContext creates a new IpsecCaCertificateCreateParams object
// with the ability to set a context for a request.
func NewIpsecCaCertificateCreateParamsWithContext(ctx context.Context) *IpsecCaCertificateCreateParams {
	return &IpsecCaCertificateCreateParams{
		Context: ctx,
	}
}

// NewIpsecCaCertificateCreateParamsWithHTTPClient creates a new IpsecCaCertificateCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpsecCaCertificateCreateParamsWithHTTPClient(client *http.Client) *IpsecCaCertificateCreateParams {
	return &IpsecCaCertificateCreateParams{
		HTTPClient: client,
	}
}

/*
IpsecCaCertificateCreateParams contains all the parameters to send to the API endpoint

	for the ipsec ca certificate create operation.

	Typically these are written to a http.Request.
*/
type IpsecCaCertificateCreateParams struct {

	/* Info.

	   IPsec CA certificate parameters.
	*/
	Info *models.IpsecCaCertificate

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipsec ca certificate create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpsecCaCertificateCreateParams) WithDefaults() *IpsecCaCertificateCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipsec ca certificate create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpsecCaCertificateCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := IpsecCaCertificateCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ipsec ca certificate create params
func (o *IpsecCaCertificateCreateParams) WithTimeout(timeout time.Duration) *IpsecCaCertificateCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipsec ca certificate create params
func (o *IpsecCaCertificateCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipsec ca certificate create params
func (o *IpsecCaCertificateCreateParams) WithContext(ctx context.Context) *IpsecCaCertificateCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipsec ca certificate create params
func (o *IpsecCaCertificateCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipsec ca certificate create params
func (o *IpsecCaCertificateCreateParams) WithHTTPClient(client *http.Client) *IpsecCaCertificateCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipsec ca certificate create params
func (o *IpsecCaCertificateCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ipsec ca certificate create params
func (o *IpsecCaCertificateCreateParams) WithInfo(info *models.IpsecCaCertificate) *IpsecCaCertificateCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ipsec ca certificate create params
func (o *IpsecCaCertificateCreateParams) SetInfo(info *models.IpsecCaCertificate) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the ipsec ca certificate create params
func (o *IpsecCaCertificateCreateParams) WithReturnRecords(returnRecords *bool) *IpsecCaCertificateCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ipsec ca certificate create params
func (o *IpsecCaCertificateCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *IpsecCaCertificateCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
