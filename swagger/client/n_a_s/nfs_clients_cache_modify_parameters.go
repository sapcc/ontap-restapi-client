// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewNfsClientsCacheModifyParams creates a new NfsClientsCacheModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNfsClientsCacheModifyParams() *NfsClientsCacheModifyParams {
	return &NfsClientsCacheModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNfsClientsCacheModifyParamsWithTimeout creates a new NfsClientsCacheModifyParams object
// with the ability to set a timeout on a request.
func NewNfsClientsCacheModifyParamsWithTimeout(timeout time.Duration) *NfsClientsCacheModifyParams {
	return &NfsClientsCacheModifyParams{
		timeout: timeout,
	}
}

// NewNfsClientsCacheModifyParamsWithContext creates a new NfsClientsCacheModifyParams object
// with the ability to set a context for a request.
func NewNfsClientsCacheModifyParamsWithContext(ctx context.Context) *NfsClientsCacheModifyParams {
	return &NfsClientsCacheModifyParams{
		Context: ctx,
	}
}

// NewNfsClientsCacheModifyParamsWithHTTPClient creates a new NfsClientsCacheModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewNfsClientsCacheModifyParamsWithHTTPClient(client *http.Client) *NfsClientsCacheModifyParams {
	return &NfsClientsCacheModifyParams{
		HTTPClient: client,
	}
}

/*
NfsClientsCacheModifyParams contains all the parameters to send to the API endpoint

	for the nfs clients cache modify operation.

	Typically these are written to a http.Request.
*/
type NfsClientsCacheModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.NfsClientsCache

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nfs clients cache modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsClientsCacheModifyParams) WithDefaults() *NfsClientsCacheModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nfs clients cache modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsClientsCacheModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nfs clients cache modify params
func (o *NfsClientsCacheModifyParams) WithTimeout(timeout time.Duration) *NfsClientsCacheModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nfs clients cache modify params
func (o *NfsClientsCacheModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nfs clients cache modify params
func (o *NfsClientsCacheModifyParams) WithContext(ctx context.Context) *NfsClientsCacheModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nfs clients cache modify params
func (o *NfsClientsCacheModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nfs clients cache modify params
func (o *NfsClientsCacheModifyParams) WithHTTPClient(client *http.Client) *NfsClientsCacheModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nfs clients cache modify params
func (o *NfsClientsCacheModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the nfs clients cache modify params
func (o *NfsClientsCacheModifyParams) WithInfo(info *models.NfsClientsCache) *NfsClientsCacheModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the nfs clients cache modify params
func (o *NfsClientsCacheModifyParams) SetInfo(info *models.NfsClientsCache) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *NfsClientsCacheModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
