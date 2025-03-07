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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewIPSubnetCreateParams creates a new IPSubnetCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIPSubnetCreateParams() *IPSubnetCreateParams {
	return &IPSubnetCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIPSubnetCreateParamsWithTimeout creates a new IPSubnetCreateParams object
// with the ability to set a timeout on a request.
func NewIPSubnetCreateParamsWithTimeout(timeout time.Duration) *IPSubnetCreateParams {
	return &IPSubnetCreateParams{
		timeout: timeout,
	}
}

// NewIPSubnetCreateParamsWithContext creates a new IPSubnetCreateParams object
// with the ability to set a context for a request.
func NewIPSubnetCreateParamsWithContext(ctx context.Context) *IPSubnetCreateParams {
	return &IPSubnetCreateParams{
		Context: ctx,
	}
}

// NewIPSubnetCreateParamsWithHTTPClient creates a new IPSubnetCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIPSubnetCreateParamsWithHTTPClient(client *http.Client) *IPSubnetCreateParams {
	return &IPSubnetCreateParams{
		HTTPClient: client,
	}
}

/*
IPSubnetCreateParams contains all the parameters to send to the API endpoint

	for the ip subnet create operation.

	Typically these are written to a http.Request.
*/
type IPSubnetCreateParams struct {

	/* Info.

	   IP subnet parameters
	*/
	Info *models.IPSubnet

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ip subnet create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPSubnetCreateParams) WithDefaults() *IPSubnetCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ip subnet create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPSubnetCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := IPSubnetCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ip subnet create params
func (o *IPSubnetCreateParams) WithTimeout(timeout time.Duration) *IPSubnetCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ip subnet create params
func (o *IPSubnetCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ip subnet create params
func (o *IPSubnetCreateParams) WithContext(ctx context.Context) *IPSubnetCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ip subnet create params
func (o *IPSubnetCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ip subnet create params
func (o *IPSubnetCreateParams) WithHTTPClient(client *http.Client) *IPSubnetCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ip subnet create params
func (o *IPSubnetCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ip subnet create params
func (o *IPSubnetCreateParams) WithInfo(info *models.IPSubnet) *IPSubnetCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ip subnet create params
func (o *IPSubnetCreateParams) SetInfo(info *models.IPSubnet) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the ip subnet create params
func (o *IPSubnetCreateParams) WithReturnRecords(returnRecords *bool) *IPSubnetCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ip subnet create params
func (o *IPSubnetCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *IPSubnetCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
