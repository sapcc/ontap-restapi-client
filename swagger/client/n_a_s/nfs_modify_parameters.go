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

// NewNfsModifyParams creates a new NfsModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNfsModifyParams() *NfsModifyParams {
	return &NfsModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNfsModifyParamsWithTimeout creates a new NfsModifyParams object
// with the ability to set a timeout on a request.
func NewNfsModifyParamsWithTimeout(timeout time.Duration) *NfsModifyParams {
	return &NfsModifyParams{
		timeout: timeout,
	}
}

// NewNfsModifyParamsWithContext creates a new NfsModifyParams object
// with the ability to set a context for a request.
func NewNfsModifyParamsWithContext(ctx context.Context) *NfsModifyParams {
	return &NfsModifyParams{
		Context: ctx,
	}
}

// NewNfsModifyParamsWithHTTPClient creates a new NfsModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewNfsModifyParamsWithHTTPClient(client *http.Client) *NfsModifyParams {
	return &NfsModifyParams{
		HTTPClient: client,
	}
}

/*
NfsModifyParams contains all the parameters to send to the API endpoint

	for the nfs modify operation.

	Typically these are written to a http.Request.
*/
type NfsModifyParams struct {

	/* Info.

	   Info Specification
	*/
	Info *models.NfsService

	// SvmUUID.
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nfs modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsModifyParams) WithDefaults() *NfsModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nfs modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nfs modify params
func (o *NfsModifyParams) WithTimeout(timeout time.Duration) *NfsModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nfs modify params
func (o *NfsModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nfs modify params
func (o *NfsModifyParams) WithContext(ctx context.Context) *NfsModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nfs modify params
func (o *NfsModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nfs modify params
func (o *NfsModifyParams) WithHTTPClient(client *http.Client) *NfsModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nfs modify params
func (o *NfsModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the nfs modify params
func (o *NfsModifyParams) WithInfo(info *models.NfsService) *NfsModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the nfs modify params
func (o *NfsModifyParams) SetInfo(info *models.NfsService) {
	o.Info = info
}

// WithSvmUUID adds the svmUUID to the nfs modify params
func (o *NfsModifyParams) WithSvmUUID(svmUUID string) *NfsModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the nfs modify params
func (o *NfsModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NfsModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
