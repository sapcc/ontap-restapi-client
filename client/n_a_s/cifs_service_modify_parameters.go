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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NewCifsServiceModifyParams creates a new CifsServiceModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsServiceModifyParams() *CifsServiceModifyParams {
	return &CifsServiceModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsServiceModifyParamsWithTimeout creates a new CifsServiceModifyParams object
// with the ability to set a timeout on a request.
func NewCifsServiceModifyParamsWithTimeout(timeout time.Duration) *CifsServiceModifyParams {
	return &CifsServiceModifyParams{
		timeout: timeout,
	}
}

// NewCifsServiceModifyParamsWithContext creates a new CifsServiceModifyParams object
// with the ability to set a context for a request.
func NewCifsServiceModifyParamsWithContext(ctx context.Context) *CifsServiceModifyParams {
	return &CifsServiceModifyParams{
		Context: ctx,
	}
}

// NewCifsServiceModifyParamsWithHTTPClient creates a new CifsServiceModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsServiceModifyParamsWithHTTPClient(client *http.Client) *CifsServiceModifyParams {
	return &CifsServiceModifyParams{
		HTTPClient: client,
	}
}

/*
CifsServiceModifyParams contains all the parameters to send to the API endpoint

	for the cifs service modify operation.

	Typically these are written to a http.Request.
*/
type CifsServiceModifyParams struct {

	/* Force.

	   If this is set and a machine account with the same name as specified in 'cifs-server name' exists in the Active Directory, existing machine account will be overwritten and reused. The default value for this field is false.
	*/
	Force *bool

	/* Info.

	   Info specification
	*/
	Info *models.CifsService

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

// WithDefaults hydrates default values in the cifs service modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsServiceModifyParams) WithDefaults() *CifsServiceModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs service modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsServiceModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := CifsServiceModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs service modify params
func (o *CifsServiceModifyParams) WithTimeout(timeout time.Duration) *CifsServiceModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs service modify params
func (o *CifsServiceModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs service modify params
func (o *CifsServiceModifyParams) WithContext(ctx context.Context) *CifsServiceModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs service modify params
func (o *CifsServiceModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs service modify params
func (o *CifsServiceModifyParams) WithHTTPClient(client *http.Client) *CifsServiceModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs service modify params
func (o *CifsServiceModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the cifs service modify params
func (o *CifsServiceModifyParams) WithForce(force *bool) *CifsServiceModifyParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the cifs service modify params
func (o *CifsServiceModifyParams) SetForce(force *bool) {
	o.Force = force
}

// WithInfo adds the info to the cifs service modify params
func (o *CifsServiceModifyParams) WithInfo(info *models.CifsService) *CifsServiceModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cifs service modify params
func (o *CifsServiceModifyParams) SetInfo(info *models.CifsService) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the cifs service modify params
func (o *CifsServiceModifyParams) WithReturnTimeout(returnTimeout *int64) *CifsServiceModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cifs service modify params
func (o *CifsServiceModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmUUID adds the svmUUID to the cifs service modify params
func (o *CifsServiceModifyParams) WithSvmUUID(svmUUID string) *CifsServiceModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the cifs service modify params
func (o *CifsServiceModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsServiceModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}
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
