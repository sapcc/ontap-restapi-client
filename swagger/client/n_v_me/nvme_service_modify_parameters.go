// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NewNvmeServiceModifyParams creates a new NvmeServiceModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeServiceModifyParams() *NvmeServiceModifyParams {
	return &NvmeServiceModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeServiceModifyParamsWithTimeout creates a new NvmeServiceModifyParams object
// with the ability to set a timeout on a request.
func NewNvmeServiceModifyParamsWithTimeout(timeout time.Duration) *NvmeServiceModifyParams {
	return &NvmeServiceModifyParams{
		timeout: timeout,
	}
}

// NewNvmeServiceModifyParamsWithContext creates a new NvmeServiceModifyParams object
// with the ability to set a context for a request.
func NewNvmeServiceModifyParamsWithContext(ctx context.Context) *NvmeServiceModifyParams {
	return &NvmeServiceModifyParams{
		Context: ctx,
	}
}

// NewNvmeServiceModifyParamsWithHTTPClient creates a new NvmeServiceModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeServiceModifyParamsWithHTTPClient(client *http.Client) *NvmeServiceModifyParams {
	return &NvmeServiceModifyParams{
		HTTPClient: client,
	}
}

/*
NvmeServiceModifyParams contains all the parameters to send to the API endpoint

	for the nvme service modify operation.

	Typically these are written to a http.Request.
*/
type NvmeServiceModifyParams struct {

	/* Info.

	   The new property values for the NVMe service.

	*/
	Info *models.NvmeService

	/* SvmUUID.

	   The unique identifier of the SVM whose NVMe service is to be updated.

	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme service modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeServiceModifyParams) WithDefaults() *NvmeServiceModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme service modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeServiceModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nvme service modify params
func (o *NvmeServiceModifyParams) WithTimeout(timeout time.Duration) *NvmeServiceModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme service modify params
func (o *NvmeServiceModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme service modify params
func (o *NvmeServiceModifyParams) WithContext(ctx context.Context) *NvmeServiceModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme service modify params
func (o *NvmeServiceModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme service modify params
func (o *NvmeServiceModifyParams) WithHTTPClient(client *http.Client) *NvmeServiceModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme service modify params
func (o *NvmeServiceModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the nvme service modify params
func (o *NvmeServiceModifyParams) WithInfo(info *models.NvmeService) *NvmeServiceModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the nvme service modify params
func (o *NvmeServiceModifyParams) SetInfo(info *models.NvmeService) {
	o.Info = info
}

// WithSvmUUID adds the svmUUID to the nvme service modify params
func (o *NvmeServiceModifyParams) WithSvmUUID(svmUUID string) *NvmeServiceModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the nvme service modify params
func (o *NvmeServiceModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeServiceModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
