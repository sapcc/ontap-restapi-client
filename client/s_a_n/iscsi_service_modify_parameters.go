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

// NewIscsiServiceModifyParams creates a new IscsiServiceModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIscsiServiceModifyParams() *IscsiServiceModifyParams {
	return &IscsiServiceModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIscsiServiceModifyParamsWithTimeout creates a new IscsiServiceModifyParams object
// with the ability to set a timeout on a request.
func NewIscsiServiceModifyParamsWithTimeout(timeout time.Duration) *IscsiServiceModifyParams {
	return &IscsiServiceModifyParams{
		timeout: timeout,
	}
}

// NewIscsiServiceModifyParamsWithContext creates a new IscsiServiceModifyParams object
// with the ability to set a context for a request.
func NewIscsiServiceModifyParamsWithContext(ctx context.Context) *IscsiServiceModifyParams {
	return &IscsiServiceModifyParams{
		Context: ctx,
	}
}

// NewIscsiServiceModifyParamsWithHTTPClient creates a new IscsiServiceModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewIscsiServiceModifyParamsWithHTTPClient(client *http.Client) *IscsiServiceModifyParams {
	return &IscsiServiceModifyParams{
		HTTPClient: client,
	}
}

/*
IscsiServiceModifyParams contains all the parameters to send to the API endpoint

	for the iscsi service modify operation.

	Typically these are written to a http.Request.
*/
type IscsiServiceModifyParams struct {

	/* Info.

	   The new property values for the iSCSI service.

	*/
	Info *models.IscsiService

	/* SvmUUID.

	   The unique identifier of the SVM for which to update the iSCSI service.

	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iscsi service modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiServiceModifyParams) WithDefaults() *IscsiServiceModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iscsi service modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiServiceModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the iscsi service modify params
func (o *IscsiServiceModifyParams) WithTimeout(timeout time.Duration) *IscsiServiceModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iscsi service modify params
func (o *IscsiServiceModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iscsi service modify params
func (o *IscsiServiceModifyParams) WithContext(ctx context.Context) *IscsiServiceModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iscsi service modify params
func (o *IscsiServiceModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iscsi service modify params
func (o *IscsiServiceModifyParams) WithHTTPClient(client *http.Client) *IscsiServiceModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iscsi service modify params
func (o *IscsiServiceModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the iscsi service modify params
func (o *IscsiServiceModifyParams) WithInfo(info *models.IscsiService) *IscsiServiceModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the iscsi service modify params
func (o *IscsiServiceModifyParams) SetInfo(info *models.IscsiService) {
	o.Info = info
}

// WithSvmUUID adds the svmUUID to the iscsi service modify params
func (o *IscsiServiceModifyParams) WithSvmUUID(svmUUID string) *IscsiServiceModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the iscsi service modify params
func (o *IscsiServiceModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *IscsiServiceModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
