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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NewSSHModifyParams creates a new SSHModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSSHModifyParams() *SSHModifyParams {
	return &SSHModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSSHModifyParamsWithTimeout creates a new SSHModifyParams object
// with the ability to set a timeout on a request.
func NewSSHModifyParamsWithTimeout(timeout time.Duration) *SSHModifyParams {
	return &SSHModifyParams{
		timeout: timeout,
	}
}

// NewSSHModifyParamsWithContext creates a new SSHModifyParams object
// with the ability to set a context for a request.
func NewSSHModifyParamsWithContext(ctx context.Context) *SSHModifyParams {
	return &SSHModifyParams{
		Context: ctx,
	}
}

// NewSSHModifyParamsWithHTTPClient creates a new SSHModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSSHModifyParamsWithHTTPClient(client *http.Client) *SSHModifyParams {
	return &SSHModifyParams{
		HTTPClient: client,
	}
}

/*
SSHModifyParams contains all the parameters to send to the API endpoint

	for the ssh modify operation.

	Typically these are written to a http.Request.
*/
type SSHModifyParams struct {

	/* Info.

	   SSH server algorithms and limits configuration details.
	*/
	Info *models.ClusterSSHServer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ssh modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSHModifyParams) WithDefaults() *SSHModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ssh modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSHModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ssh modify params
func (o *SSHModifyParams) WithTimeout(timeout time.Duration) *SSHModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ssh modify params
func (o *SSHModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ssh modify params
func (o *SSHModifyParams) WithContext(ctx context.Context) *SSHModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ssh modify params
func (o *SSHModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ssh modify params
func (o *SSHModifyParams) WithHTTPClient(client *http.Client) *SSHModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ssh modify params
func (o *SSHModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ssh modify params
func (o *SSHModifyParams) WithInfo(info *models.ClusterSSHServer) *SSHModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ssh modify params
func (o *SSHModifyParams) SetInfo(info *models.ClusterSSHServer) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *SSHModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
