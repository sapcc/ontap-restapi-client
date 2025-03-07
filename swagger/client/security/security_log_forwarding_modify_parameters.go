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

// NewSecurityLogForwardingModifyParams creates a new SecurityLogForwardingModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityLogForwardingModifyParams() *SecurityLogForwardingModifyParams {
	return &SecurityLogForwardingModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityLogForwardingModifyParamsWithTimeout creates a new SecurityLogForwardingModifyParams object
// with the ability to set a timeout on a request.
func NewSecurityLogForwardingModifyParamsWithTimeout(timeout time.Duration) *SecurityLogForwardingModifyParams {
	return &SecurityLogForwardingModifyParams{
		timeout: timeout,
	}
}

// NewSecurityLogForwardingModifyParamsWithContext creates a new SecurityLogForwardingModifyParams object
// with the ability to set a context for a request.
func NewSecurityLogForwardingModifyParamsWithContext(ctx context.Context) *SecurityLogForwardingModifyParams {
	return &SecurityLogForwardingModifyParams{
		Context: ctx,
	}
}

// NewSecurityLogForwardingModifyParamsWithHTTPClient creates a new SecurityLogForwardingModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityLogForwardingModifyParamsWithHTTPClient(client *http.Client) *SecurityLogForwardingModifyParams {
	return &SecurityLogForwardingModifyParams{
		HTTPClient: client,
	}
}

/*
SecurityLogForwardingModifyParams contains all the parameters to send to the API endpoint

	for the security log forwarding modify operation.

	Typically these are written to a http.Request.
*/
type SecurityLogForwardingModifyParams struct {

	/* Address.

	   IP address of remote syslog/splunk server.
	*/
	Address string

	/* Info.

	   Modify remote syslog/splunk server information.
	*/
	Info *models.SecurityAuditLogForward

	/* Port.

	   Port number of remote syslog/splunk server.
	*/
	Port int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security log forwarding modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityLogForwardingModifyParams) WithDefaults() *SecurityLogForwardingModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security log forwarding modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityLogForwardingModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) WithTimeout(timeout time.Duration) *SecurityLogForwardingModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) WithContext(ctx context.Context) *SecurityLogForwardingModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) WithHTTPClient(client *http.Client) *SecurityLogForwardingModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) WithAddress(address string) *SecurityLogForwardingModifyParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) SetAddress(address string) {
	o.Address = address
}

// WithInfo adds the info to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) WithInfo(info *models.SecurityAuditLogForward) *SecurityLogForwardingModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) SetInfo(info *models.SecurityAuditLogForward) {
	o.Info = info
}

// WithPort adds the port to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) WithPort(port int64) *SecurityLogForwardingModifyParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the security log forwarding modify params
func (o *SecurityLogForwardingModifyParams) SetPort(port int64) {
	o.Port = port
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityLogForwardingModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param port
	if err := r.SetPathParam("port", swag.FormatInt64(o.Port)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
