// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewFcpServiceModifyParams creates a new FcpServiceModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFcpServiceModifyParams() *FcpServiceModifyParams {
	return &FcpServiceModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFcpServiceModifyParamsWithTimeout creates a new FcpServiceModifyParams object
// with the ability to set a timeout on a request.
func NewFcpServiceModifyParamsWithTimeout(timeout time.Duration) *FcpServiceModifyParams {
	return &FcpServiceModifyParams{
		timeout: timeout,
	}
}

// NewFcpServiceModifyParamsWithContext creates a new FcpServiceModifyParams object
// with the ability to set a context for a request.
func NewFcpServiceModifyParamsWithContext(ctx context.Context) *FcpServiceModifyParams {
	return &FcpServiceModifyParams{
		Context: ctx,
	}
}

// NewFcpServiceModifyParamsWithHTTPClient creates a new FcpServiceModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewFcpServiceModifyParamsWithHTTPClient(client *http.Client) *FcpServiceModifyParams {
	return &FcpServiceModifyParams{
		HTTPClient: client,
	}
}

/*
FcpServiceModifyParams contains all the parameters to send to the API endpoint

	for the fcp service modify operation.

	Typically these are written to a http.Request.
*/
type FcpServiceModifyParams struct {

	/* Info.

	   The new property values for the FC Protocol service.

	*/
	Info *models.FcpService

	/* SvmUUID.

	   The unique identifier of the SVM whose FC Protocol service is to be updated.

	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fcp service modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcpServiceModifyParams) WithDefaults() *FcpServiceModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fcp service modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcpServiceModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fcp service modify params
func (o *FcpServiceModifyParams) WithTimeout(timeout time.Duration) *FcpServiceModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fcp service modify params
func (o *FcpServiceModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fcp service modify params
func (o *FcpServiceModifyParams) WithContext(ctx context.Context) *FcpServiceModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fcp service modify params
func (o *FcpServiceModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fcp service modify params
func (o *FcpServiceModifyParams) WithHTTPClient(client *http.Client) *FcpServiceModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fcp service modify params
func (o *FcpServiceModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the fcp service modify params
func (o *FcpServiceModifyParams) WithInfo(info *models.FcpService) *FcpServiceModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the fcp service modify params
func (o *FcpServiceModifyParams) SetInfo(info *models.FcpService) {
	o.Info = info
}

// WithSvmUUID adds the svmUUID to the fcp service modify params
func (o *FcpServiceModifyParams) WithSvmUUID(svmUUID string) *FcpServiceModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the fcp service modify params
func (o *FcpServiceModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FcpServiceModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
