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

	"github.com/sapcc/ontap-restapi/models"
)

// NewVscanOnAccessModifyParams creates a new VscanOnAccessModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanOnAccessModifyParams() *VscanOnAccessModifyParams {
	return &VscanOnAccessModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanOnAccessModifyParamsWithTimeout creates a new VscanOnAccessModifyParams object
// with the ability to set a timeout on a request.
func NewVscanOnAccessModifyParamsWithTimeout(timeout time.Duration) *VscanOnAccessModifyParams {
	return &VscanOnAccessModifyParams{
		timeout: timeout,
	}
}

// NewVscanOnAccessModifyParamsWithContext creates a new VscanOnAccessModifyParams object
// with the ability to set a context for a request.
func NewVscanOnAccessModifyParamsWithContext(ctx context.Context) *VscanOnAccessModifyParams {
	return &VscanOnAccessModifyParams{
		Context: ctx,
	}
}

// NewVscanOnAccessModifyParamsWithHTTPClient creates a new VscanOnAccessModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanOnAccessModifyParamsWithHTTPClient(client *http.Client) *VscanOnAccessModifyParams {
	return &VscanOnAccessModifyParams{
		HTTPClient: client,
	}
}

/*
VscanOnAccessModifyParams contains all the parameters to send to the API endpoint

	for the vscan on access modify operation.

	Typically these are written to a http.Request.
*/
type VscanOnAccessModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.VscanOnAccess

	// Name.
	Name string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan on access modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnAccessModifyParams) WithDefaults() *VscanOnAccessModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan on access modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnAccessModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vscan on access modify params
func (o *VscanOnAccessModifyParams) WithTimeout(timeout time.Duration) *VscanOnAccessModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan on access modify params
func (o *VscanOnAccessModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan on access modify params
func (o *VscanOnAccessModifyParams) WithContext(ctx context.Context) *VscanOnAccessModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan on access modify params
func (o *VscanOnAccessModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan on access modify params
func (o *VscanOnAccessModifyParams) WithHTTPClient(client *http.Client) *VscanOnAccessModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan on access modify params
func (o *VscanOnAccessModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the vscan on access modify params
func (o *VscanOnAccessModifyParams) WithInfo(info *models.VscanOnAccess) *VscanOnAccessModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the vscan on access modify params
func (o *VscanOnAccessModifyParams) SetInfo(info *models.VscanOnAccess) {
	o.Info = info
}

// WithName adds the name to the vscan on access modify params
func (o *VscanOnAccessModifyParams) WithName(name string) *VscanOnAccessModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the vscan on access modify params
func (o *VscanOnAccessModifyParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the vscan on access modify params
func (o *VscanOnAccessModifyParams) WithSvmUUID(svmUUID string) *VscanOnAccessModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vscan on access modify params
func (o *VscanOnAccessModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VscanOnAccessModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
