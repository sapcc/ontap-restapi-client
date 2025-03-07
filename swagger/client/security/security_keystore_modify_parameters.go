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
	"github.com/go-openapi/swag"

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewSecurityKeystoreModifyParams creates a new SecurityKeystoreModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeystoreModifyParams() *SecurityKeystoreModifyParams {
	return &SecurityKeystoreModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeystoreModifyParamsWithTimeout creates a new SecurityKeystoreModifyParams object
// with the ability to set a timeout on a request.
func NewSecurityKeystoreModifyParamsWithTimeout(timeout time.Duration) *SecurityKeystoreModifyParams {
	return &SecurityKeystoreModifyParams{
		timeout: timeout,
	}
}

// NewSecurityKeystoreModifyParamsWithContext creates a new SecurityKeystoreModifyParams object
// with the ability to set a context for a request.
func NewSecurityKeystoreModifyParamsWithContext(ctx context.Context) *SecurityKeystoreModifyParams {
	return &SecurityKeystoreModifyParams{
		Context: ctx,
	}
}

// NewSecurityKeystoreModifyParamsWithHTTPClient creates a new SecurityKeystoreModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeystoreModifyParamsWithHTTPClient(client *http.Client) *SecurityKeystoreModifyParams {
	return &SecurityKeystoreModifyParams{
		HTTPClient: client,
	}
}

/*
SecurityKeystoreModifyParams contains all the parameters to send to the API endpoint

	for the security keystore modify operation.

	Typically these are written to a http.Request.
*/
type SecurityKeystoreModifyParams struct {

	/* Info.

	   Keystore information
	*/
	Info *models.SecurityKeystore

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   Keystore configuration UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security keystore modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeystoreModifyParams) WithDefaults() *SecurityKeystoreModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security keystore modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeystoreModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := SecurityKeystoreModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security keystore modify params
func (o *SecurityKeystoreModifyParams) WithTimeout(timeout time.Duration) *SecurityKeystoreModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security keystore modify params
func (o *SecurityKeystoreModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security keystore modify params
func (o *SecurityKeystoreModifyParams) WithContext(ctx context.Context) *SecurityKeystoreModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security keystore modify params
func (o *SecurityKeystoreModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security keystore modify params
func (o *SecurityKeystoreModifyParams) WithHTTPClient(client *http.Client) *SecurityKeystoreModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security keystore modify params
func (o *SecurityKeystoreModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the security keystore modify params
func (o *SecurityKeystoreModifyParams) WithInfo(info *models.SecurityKeystore) *SecurityKeystoreModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security keystore modify params
func (o *SecurityKeystoreModifyParams) SetInfo(info *models.SecurityKeystore) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the security keystore modify params
func (o *SecurityKeystoreModifyParams) WithReturnTimeout(returnTimeout *int64) *SecurityKeystoreModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security keystore modify params
func (o *SecurityKeystoreModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the security keystore modify params
func (o *SecurityKeystoreModifyParams) WithUUID(uuid string) *SecurityKeystoreModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the security keystore modify params
func (o *SecurityKeystoreModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeystoreModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
