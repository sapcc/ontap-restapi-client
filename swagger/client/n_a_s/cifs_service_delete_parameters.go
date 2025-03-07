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

// NewCifsServiceDeleteParams creates a new CifsServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsServiceDeleteParams() *CifsServiceDeleteParams {
	return &CifsServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsServiceDeleteParamsWithTimeout creates a new CifsServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewCifsServiceDeleteParamsWithTimeout(timeout time.Duration) *CifsServiceDeleteParams {
	return &CifsServiceDeleteParams{
		timeout: timeout,
	}
}

// NewCifsServiceDeleteParamsWithContext creates a new CifsServiceDeleteParams object
// with the ability to set a context for a request.
func NewCifsServiceDeleteParamsWithContext(ctx context.Context) *CifsServiceDeleteParams {
	return &CifsServiceDeleteParams{
		Context: ctx,
	}
}

// NewCifsServiceDeleteParamsWithHTTPClient creates a new CifsServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsServiceDeleteParamsWithHTTPClient(client *http.Client) *CifsServiceDeleteParams {
	return &CifsServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
CifsServiceDeleteParams contains all the parameters to send to the API endpoint

	for the cifs service delete operation.

	Typically these are written to a http.Request.
*/
type CifsServiceDeleteParams struct {

	/* Force.

	   When set, the local CIFS configuration is deleted irrespective of any communication errors. Default value for this field is false.
	*/
	Force *bool

	/* Info.

	   Info specification
	*/
	Info *models.CifsServiceDelete

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

// WithDefaults hydrates default values in the cifs service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsServiceDeleteParams) WithDefaults() *CifsServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsServiceDeleteParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := CifsServiceDeleteParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs service delete params
func (o *CifsServiceDeleteParams) WithTimeout(timeout time.Duration) *CifsServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs service delete params
func (o *CifsServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs service delete params
func (o *CifsServiceDeleteParams) WithContext(ctx context.Context) *CifsServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs service delete params
func (o *CifsServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs service delete params
func (o *CifsServiceDeleteParams) WithHTTPClient(client *http.Client) *CifsServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs service delete params
func (o *CifsServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the cifs service delete params
func (o *CifsServiceDeleteParams) WithForce(force *bool) *CifsServiceDeleteParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the cifs service delete params
func (o *CifsServiceDeleteParams) SetForce(force *bool) {
	o.Force = force
}

// WithInfo adds the info to the cifs service delete params
func (o *CifsServiceDeleteParams) WithInfo(info *models.CifsServiceDelete) *CifsServiceDeleteParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cifs service delete params
func (o *CifsServiceDeleteParams) SetInfo(info *models.CifsServiceDelete) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the cifs service delete params
func (o *CifsServiceDeleteParams) WithReturnTimeout(returnTimeout *int64) *CifsServiceDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cifs service delete params
func (o *CifsServiceDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmUUID adds the svmUUID to the cifs service delete params
func (o *CifsServiceDeleteParams) WithSvmUUID(svmUUID string) *CifsServiceDeleteParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the cifs service delete params
func (o *CifsServiceDeleteParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
