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

// NewSecurityOauth2GlobalModifyParams creates a new SecurityOauth2GlobalModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityOauth2GlobalModifyParams() *SecurityOauth2GlobalModifyParams {
	return &SecurityOauth2GlobalModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityOauth2GlobalModifyParamsWithTimeout creates a new SecurityOauth2GlobalModifyParams object
// with the ability to set a timeout on a request.
func NewSecurityOauth2GlobalModifyParamsWithTimeout(timeout time.Duration) *SecurityOauth2GlobalModifyParams {
	return &SecurityOauth2GlobalModifyParams{
		timeout: timeout,
	}
}

// NewSecurityOauth2GlobalModifyParamsWithContext creates a new SecurityOauth2GlobalModifyParams object
// with the ability to set a context for a request.
func NewSecurityOauth2GlobalModifyParamsWithContext(ctx context.Context) *SecurityOauth2GlobalModifyParams {
	return &SecurityOauth2GlobalModifyParams{
		Context: ctx,
	}
}

// NewSecurityOauth2GlobalModifyParamsWithHTTPClient creates a new SecurityOauth2GlobalModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityOauth2GlobalModifyParamsWithHTTPClient(client *http.Client) *SecurityOauth2GlobalModifyParams {
	return &SecurityOauth2GlobalModifyParams{
		HTTPClient: client,
	}
}

/*
SecurityOauth2GlobalModifyParams contains all the parameters to send to the API endpoint

	for the security oauth2 global modify operation.

	Typically these are written to a http.Request.
*/
type SecurityOauth2GlobalModifyParams struct {

	/* Info.

	   Status specification
	*/
	Info *models.SecurityOauth2Global

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security oauth2 global modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityOauth2GlobalModifyParams) WithDefaults() *SecurityOauth2GlobalModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security oauth2 global modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityOauth2GlobalModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security oauth2 global modify params
func (o *SecurityOauth2GlobalModifyParams) WithTimeout(timeout time.Duration) *SecurityOauth2GlobalModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security oauth2 global modify params
func (o *SecurityOauth2GlobalModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security oauth2 global modify params
func (o *SecurityOauth2GlobalModifyParams) WithContext(ctx context.Context) *SecurityOauth2GlobalModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security oauth2 global modify params
func (o *SecurityOauth2GlobalModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security oauth2 global modify params
func (o *SecurityOauth2GlobalModifyParams) WithHTTPClient(client *http.Client) *SecurityOauth2GlobalModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security oauth2 global modify params
func (o *SecurityOauth2GlobalModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the security oauth2 global modify params
func (o *SecurityOauth2GlobalModifyParams) WithInfo(info *models.SecurityOauth2Global) *SecurityOauth2GlobalModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security oauth2 global modify params
func (o *SecurityOauth2GlobalModifyParams) SetInfo(info *models.SecurityOauth2Global) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityOauth2GlobalModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
