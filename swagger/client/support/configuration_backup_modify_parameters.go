// Code generated by go-swagger; DO NOT EDIT.

package support

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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewConfigurationBackupModifyParams creates a new ConfigurationBackupModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfigurationBackupModifyParams() *ConfigurationBackupModifyParams {
	return &ConfigurationBackupModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfigurationBackupModifyParamsWithTimeout creates a new ConfigurationBackupModifyParams object
// with the ability to set a timeout on a request.
func NewConfigurationBackupModifyParamsWithTimeout(timeout time.Duration) *ConfigurationBackupModifyParams {
	return &ConfigurationBackupModifyParams{
		timeout: timeout,
	}
}

// NewConfigurationBackupModifyParamsWithContext creates a new ConfigurationBackupModifyParams object
// with the ability to set a context for a request.
func NewConfigurationBackupModifyParamsWithContext(ctx context.Context) *ConfigurationBackupModifyParams {
	return &ConfigurationBackupModifyParams{
		Context: ctx,
	}
}

// NewConfigurationBackupModifyParamsWithHTTPClient creates a new ConfigurationBackupModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfigurationBackupModifyParamsWithHTTPClient(client *http.Client) *ConfigurationBackupModifyParams {
	return &ConfigurationBackupModifyParams{
		HTTPClient: client,
	}
}

/*
ConfigurationBackupModifyParams contains all the parameters to send to the API endpoint

	for the configuration backup modify operation.

	Typically these are written to a http.Request.
*/
type ConfigurationBackupModifyParams struct {

	/* Info.

	   Cluster configuration backup information
	*/
	Info *models.ConfigurationBackup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the configuration backup modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigurationBackupModifyParams) WithDefaults() *ConfigurationBackupModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the configuration backup modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigurationBackupModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the configuration backup modify params
func (o *ConfigurationBackupModifyParams) WithTimeout(timeout time.Duration) *ConfigurationBackupModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the configuration backup modify params
func (o *ConfigurationBackupModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the configuration backup modify params
func (o *ConfigurationBackupModifyParams) WithContext(ctx context.Context) *ConfigurationBackupModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the configuration backup modify params
func (o *ConfigurationBackupModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the configuration backup modify params
func (o *ConfigurationBackupModifyParams) WithHTTPClient(client *http.Client) *ConfigurationBackupModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the configuration backup modify params
func (o *ConfigurationBackupModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the configuration backup modify params
func (o *ConfigurationBackupModifyParams) WithInfo(info *models.ConfigurationBackup) *ConfigurationBackupModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the configuration backup modify params
func (o *ConfigurationBackupModifyParams) SetInfo(info *models.ConfigurationBackup) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigurationBackupModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
