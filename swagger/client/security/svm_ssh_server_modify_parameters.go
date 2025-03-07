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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewSvmSSHServerModifyParams creates a new SvmSSHServerModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmSSHServerModifyParams() *SvmSSHServerModifyParams {
	return &SvmSSHServerModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmSSHServerModifyParamsWithTimeout creates a new SvmSSHServerModifyParams object
// with the ability to set a timeout on a request.
func NewSvmSSHServerModifyParamsWithTimeout(timeout time.Duration) *SvmSSHServerModifyParams {
	return &SvmSSHServerModifyParams{
		timeout: timeout,
	}
}

// NewSvmSSHServerModifyParamsWithContext creates a new SvmSSHServerModifyParams object
// with the ability to set a context for a request.
func NewSvmSSHServerModifyParamsWithContext(ctx context.Context) *SvmSSHServerModifyParams {
	return &SvmSSHServerModifyParams{
		Context: ctx,
	}
}

// NewSvmSSHServerModifyParamsWithHTTPClient creates a new SvmSSHServerModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmSSHServerModifyParamsWithHTTPClient(client *http.Client) *SvmSSHServerModifyParams {
	return &SvmSSHServerModifyParams{
		HTTPClient: client,
	}
}

/*
SvmSSHServerModifyParams contains all the parameters to send to the API endpoint

	for the svm ssh server modify operation.

	Typically these are written to a http.Request.
*/
type SvmSSHServerModifyParams struct {

	/* Info.

	   SSH server algorithm configuration details.
	*/
	Info *models.SvmSSHServer

	/* SvmUUID.

	   SVM UUID
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm ssh server modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmSSHServerModifyParams) WithDefaults() *SvmSSHServerModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm ssh server modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmSSHServerModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the svm ssh server modify params
func (o *SvmSSHServerModifyParams) WithTimeout(timeout time.Duration) *SvmSSHServerModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm ssh server modify params
func (o *SvmSSHServerModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm ssh server modify params
func (o *SvmSSHServerModifyParams) WithContext(ctx context.Context) *SvmSSHServerModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm ssh server modify params
func (o *SvmSSHServerModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm ssh server modify params
func (o *SvmSSHServerModifyParams) WithHTTPClient(client *http.Client) *SvmSSHServerModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm ssh server modify params
func (o *SvmSSHServerModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the svm ssh server modify params
func (o *SvmSSHServerModifyParams) WithInfo(info *models.SvmSSHServer) *SvmSSHServerModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the svm ssh server modify params
func (o *SvmSSHServerModifyParams) SetInfo(info *models.SvmSSHServer) {
	o.Info = info
}

// WithSvmUUID adds the svmUUID to the svm ssh server modify params
func (o *SvmSSHServerModifyParams) WithSvmUUID(svmUUID string) *SvmSSHServerModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the svm ssh server modify params
func (o *SvmSSHServerModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SvmSSHServerModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
