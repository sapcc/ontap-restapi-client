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

// NewFpolicyEngineModifyParams creates a new FpolicyEngineModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyEngineModifyParams() *FpolicyEngineModifyParams {
	return &FpolicyEngineModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyEngineModifyParamsWithTimeout creates a new FpolicyEngineModifyParams object
// with the ability to set a timeout on a request.
func NewFpolicyEngineModifyParamsWithTimeout(timeout time.Duration) *FpolicyEngineModifyParams {
	return &FpolicyEngineModifyParams{
		timeout: timeout,
	}
}

// NewFpolicyEngineModifyParamsWithContext creates a new FpolicyEngineModifyParams object
// with the ability to set a context for a request.
func NewFpolicyEngineModifyParamsWithContext(ctx context.Context) *FpolicyEngineModifyParams {
	return &FpolicyEngineModifyParams{
		Context: ctx,
	}
}

// NewFpolicyEngineModifyParamsWithHTTPClient creates a new FpolicyEngineModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyEngineModifyParamsWithHTTPClient(client *http.Client) *FpolicyEngineModifyParams {
	return &FpolicyEngineModifyParams{
		HTTPClient: client,
	}
}

/*
FpolicyEngineModifyParams contains all the parameters to send to the API endpoint

	for the fpolicy engine modify operation.

	Typically these are written to a http.Request.
*/
type FpolicyEngineModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.FpolicyEngine

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

// WithDefaults hydrates default values in the fpolicy engine modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyEngineModifyParams) WithDefaults() *FpolicyEngineModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy engine modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyEngineModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) WithTimeout(timeout time.Duration) *FpolicyEngineModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) WithContext(ctx context.Context) *FpolicyEngineModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) WithHTTPClient(client *http.Client) *FpolicyEngineModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) WithInfo(info *models.FpolicyEngine) *FpolicyEngineModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) SetInfo(info *models.FpolicyEngine) {
	o.Info = info
}

// WithName adds the name to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) WithName(name string) *FpolicyEngineModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) WithSvmUUID(svmUUID string) *FpolicyEngineModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the fpolicy engine modify params
func (o *FpolicyEngineModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyEngineModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
