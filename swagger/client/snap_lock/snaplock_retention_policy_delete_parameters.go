// Code generated by go-swagger; DO NOT EDIT.

package snap_lock

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

// NewSnaplockRetentionPolicyDeleteParams creates a new SnaplockRetentionPolicyDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockRetentionPolicyDeleteParams() *SnaplockRetentionPolicyDeleteParams {
	return &SnaplockRetentionPolicyDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockRetentionPolicyDeleteParamsWithTimeout creates a new SnaplockRetentionPolicyDeleteParams object
// with the ability to set a timeout on a request.
func NewSnaplockRetentionPolicyDeleteParamsWithTimeout(timeout time.Duration) *SnaplockRetentionPolicyDeleteParams {
	return &SnaplockRetentionPolicyDeleteParams{
		timeout: timeout,
	}
}

// NewSnaplockRetentionPolicyDeleteParamsWithContext creates a new SnaplockRetentionPolicyDeleteParams object
// with the ability to set a context for a request.
func NewSnaplockRetentionPolicyDeleteParamsWithContext(ctx context.Context) *SnaplockRetentionPolicyDeleteParams {
	return &SnaplockRetentionPolicyDeleteParams{
		Context: ctx,
	}
}

// NewSnaplockRetentionPolicyDeleteParamsWithHTTPClient creates a new SnaplockRetentionPolicyDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockRetentionPolicyDeleteParamsWithHTTPClient(client *http.Client) *SnaplockRetentionPolicyDeleteParams {
	return &SnaplockRetentionPolicyDeleteParams{
		HTTPClient: client,
	}
}

/*
SnaplockRetentionPolicyDeleteParams contains all the parameters to send to the API endpoint

	for the snaplock retention policy delete operation.

	Typically these are written to a http.Request.
*/
type SnaplockRetentionPolicyDeleteParams struct {

	/* PolicyName.

	   Name of the retention policy
	*/
	PolicyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock retention policy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockRetentionPolicyDeleteParams) WithDefaults() *SnaplockRetentionPolicyDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock retention policy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockRetentionPolicyDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snaplock retention policy delete params
func (o *SnaplockRetentionPolicyDeleteParams) WithTimeout(timeout time.Duration) *SnaplockRetentionPolicyDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock retention policy delete params
func (o *SnaplockRetentionPolicyDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock retention policy delete params
func (o *SnaplockRetentionPolicyDeleteParams) WithContext(ctx context.Context) *SnaplockRetentionPolicyDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock retention policy delete params
func (o *SnaplockRetentionPolicyDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock retention policy delete params
func (o *SnaplockRetentionPolicyDeleteParams) WithHTTPClient(client *http.Client) *SnaplockRetentionPolicyDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock retention policy delete params
func (o *SnaplockRetentionPolicyDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicyName adds the policyName to the snaplock retention policy delete params
func (o *SnaplockRetentionPolicyDeleteParams) WithPolicyName(policyName string) *SnaplockRetentionPolicyDeleteParams {
	o.SetPolicyName(policyName)
	return o
}

// SetPolicyName adds the policyName to the snaplock retention policy delete params
func (o *SnaplockRetentionPolicyDeleteParams) SetPolicyName(policyName string) {
	o.PolicyName = policyName
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockRetentionPolicyDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param policy.name
	if err := r.SetPathParam("policy.name", o.PolicyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
