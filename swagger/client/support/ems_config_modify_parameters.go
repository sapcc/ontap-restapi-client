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

// NewEmsConfigModifyParams creates a new EmsConfigModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsConfigModifyParams() *EmsConfigModifyParams {
	return &EmsConfigModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsConfigModifyParamsWithTimeout creates a new EmsConfigModifyParams object
// with the ability to set a timeout on a request.
func NewEmsConfigModifyParamsWithTimeout(timeout time.Duration) *EmsConfigModifyParams {
	return &EmsConfigModifyParams{
		timeout: timeout,
	}
}

// NewEmsConfigModifyParamsWithContext creates a new EmsConfigModifyParams object
// with the ability to set a context for a request.
func NewEmsConfigModifyParamsWithContext(ctx context.Context) *EmsConfigModifyParams {
	return &EmsConfigModifyParams{
		Context: ctx,
	}
}

// NewEmsConfigModifyParamsWithHTTPClient creates a new EmsConfigModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsConfigModifyParamsWithHTTPClient(client *http.Client) *EmsConfigModifyParams {
	return &EmsConfigModifyParams{
		HTTPClient: client,
	}
}

/*
EmsConfigModifyParams contains all the parameters to send to the API endpoint

	for the ems config modify operation.

	Typically these are written to a http.Request.
*/
type EmsConfigModifyParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.EmsConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems config modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsConfigModifyParams) WithDefaults() *EmsConfigModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems config modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsConfigModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ems config modify params
func (o *EmsConfigModifyParams) WithTimeout(timeout time.Duration) *EmsConfigModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems config modify params
func (o *EmsConfigModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems config modify params
func (o *EmsConfigModifyParams) WithContext(ctx context.Context) *EmsConfigModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems config modify params
func (o *EmsConfigModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems config modify params
func (o *EmsConfigModifyParams) WithHTTPClient(client *http.Client) *EmsConfigModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems config modify params
func (o *EmsConfigModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ems config modify params
func (o *EmsConfigModifyParams) WithInfo(info *models.EmsConfig) *EmsConfigModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ems config modify params
func (o *EmsConfigModifyParams) SetInfo(info *models.EmsConfig) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *EmsConfigModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
