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
	"github.com/go-openapi/swag"

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewAuditModifyParams creates a new AuditModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditModifyParams() *AuditModifyParams {
	return &AuditModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditModifyParamsWithTimeout creates a new AuditModifyParams object
// with the ability to set a timeout on a request.
func NewAuditModifyParamsWithTimeout(timeout time.Duration) *AuditModifyParams {
	return &AuditModifyParams{
		timeout: timeout,
	}
}

// NewAuditModifyParamsWithContext creates a new AuditModifyParams object
// with the ability to set a context for a request.
func NewAuditModifyParamsWithContext(ctx context.Context) *AuditModifyParams {
	return &AuditModifyParams{
		Context: ctx,
	}
}

// NewAuditModifyParamsWithHTTPClient creates a new AuditModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditModifyParamsWithHTTPClient(client *http.Client) *AuditModifyParams {
	return &AuditModifyParams{
		HTTPClient: client,
	}
}

/*
AuditModifyParams contains all the parameters to send to the API endpoint

	for the audit modify operation.

	Typically these are written to a http.Request.
*/
type AuditModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Audit

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the audit modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditModifyParams) WithDefaults() *AuditModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := AuditModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the audit modify params
func (o *AuditModifyParams) WithTimeout(timeout time.Duration) *AuditModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit modify params
func (o *AuditModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit modify params
func (o *AuditModifyParams) WithContext(ctx context.Context) *AuditModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit modify params
func (o *AuditModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit modify params
func (o *AuditModifyParams) WithHTTPClient(client *http.Client) *AuditModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit modify params
func (o *AuditModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the audit modify params
func (o *AuditModifyParams) WithInfo(info *models.Audit) *AuditModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the audit modify params
func (o *AuditModifyParams) SetInfo(info *models.Audit) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the audit modify params
func (o *AuditModifyParams) WithReturnTimeout(returnTimeout *int64) *AuditModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the audit modify params
func (o *AuditModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmUUID adds the svmUUID to the audit modify params
func (o *AuditModifyParams) WithSvmUUID(svmUUID string) *AuditModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the audit modify params
func (o *AuditModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *AuditModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
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
