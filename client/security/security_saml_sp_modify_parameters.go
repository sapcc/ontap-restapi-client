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

	"github.com/sapcc/ontap-restapi/models"
)

// NewSecuritySamlSpModifyParams creates a new SecuritySamlSpModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecuritySamlSpModifyParams() *SecuritySamlSpModifyParams {
	return &SecuritySamlSpModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecuritySamlSpModifyParamsWithTimeout creates a new SecuritySamlSpModifyParams object
// with the ability to set a timeout on a request.
func NewSecuritySamlSpModifyParamsWithTimeout(timeout time.Duration) *SecuritySamlSpModifyParams {
	return &SecuritySamlSpModifyParams{
		timeout: timeout,
	}
}

// NewSecuritySamlSpModifyParamsWithContext creates a new SecuritySamlSpModifyParams object
// with the ability to set a context for a request.
func NewSecuritySamlSpModifyParamsWithContext(ctx context.Context) *SecuritySamlSpModifyParams {
	return &SecuritySamlSpModifyParams{
		Context: ctx,
	}
}

// NewSecuritySamlSpModifyParamsWithHTTPClient creates a new SecuritySamlSpModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecuritySamlSpModifyParamsWithHTTPClient(client *http.Client) *SecuritySamlSpModifyParams {
	return &SecuritySamlSpModifyParams{
		HTTPClient: client,
	}
}

/*
SecuritySamlSpModifyParams contains all the parameters to send to the API endpoint

	for the security saml sp modify operation.

	Typically these are written to a http.Request.
*/
type SecuritySamlSpModifyParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.SecuritySamlSp

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security saml sp modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecuritySamlSpModifyParams) WithDefaults() *SecuritySamlSpModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security saml sp modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecuritySamlSpModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security saml sp modify params
func (o *SecuritySamlSpModifyParams) WithTimeout(timeout time.Duration) *SecuritySamlSpModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security saml sp modify params
func (o *SecuritySamlSpModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security saml sp modify params
func (o *SecuritySamlSpModifyParams) WithContext(ctx context.Context) *SecuritySamlSpModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security saml sp modify params
func (o *SecuritySamlSpModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security saml sp modify params
func (o *SecuritySamlSpModifyParams) WithHTTPClient(client *http.Client) *SecuritySamlSpModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security saml sp modify params
func (o *SecuritySamlSpModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the security saml sp modify params
func (o *SecuritySamlSpModifyParams) WithInfo(info *models.SecuritySamlSp) *SecuritySamlSpModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security saml sp modify params
func (o *SecuritySamlSpModifyParams) SetInfo(info *models.SecuritySamlSp) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *SecuritySamlSpModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
