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
)

// NewSecurityKeystoreDeleteParams creates a new SecurityKeystoreDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeystoreDeleteParams() *SecurityKeystoreDeleteParams {
	return &SecurityKeystoreDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeystoreDeleteParamsWithTimeout creates a new SecurityKeystoreDeleteParams object
// with the ability to set a timeout on a request.
func NewSecurityKeystoreDeleteParamsWithTimeout(timeout time.Duration) *SecurityKeystoreDeleteParams {
	return &SecurityKeystoreDeleteParams{
		timeout: timeout,
	}
}

// NewSecurityKeystoreDeleteParamsWithContext creates a new SecurityKeystoreDeleteParams object
// with the ability to set a context for a request.
func NewSecurityKeystoreDeleteParamsWithContext(ctx context.Context) *SecurityKeystoreDeleteParams {
	return &SecurityKeystoreDeleteParams{
		Context: ctx,
	}
}

// NewSecurityKeystoreDeleteParamsWithHTTPClient creates a new SecurityKeystoreDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeystoreDeleteParamsWithHTTPClient(client *http.Client) *SecurityKeystoreDeleteParams {
	return &SecurityKeystoreDeleteParams{
		HTTPClient: client,
	}
}

/*
SecurityKeystoreDeleteParams contains all the parameters to send to the API endpoint

	for the security keystore delete operation.

	Typically these are written to a http.Request.
*/
type SecurityKeystoreDeleteParams struct {

	/* UUID.

	   Keystore configuration UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security keystore delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeystoreDeleteParams) WithDefaults() *SecurityKeystoreDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security keystore delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeystoreDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security keystore delete params
func (o *SecurityKeystoreDeleteParams) WithTimeout(timeout time.Duration) *SecurityKeystoreDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security keystore delete params
func (o *SecurityKeystoreDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security keystore delete params
func (o *SecurityKeystoreDeleteParams) WithContext(ctx context.Context) *SecurityKeystoreDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security keystore delete params
func (o *SecurityKeystoreDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security keystore delete params
func (o *SecurityKeystoreDeleteParams) WithHTTPClient(client *http.Client) *SecurityKeystoreDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security keystore delete params
func (o *SecurityKeystoreDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the security keystore delete params
func (o *SecurityKeystoreDeleteParams) WithUUID(uuid string) *SecurityKeystoreDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the security keystore delete params
func (o *SecurityKeystoreDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeystoreDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
