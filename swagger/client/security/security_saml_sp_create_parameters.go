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

// NewSecuritySamlSpCreateParams creates a new SecuritySamlSpCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecuritySamlSpCreateParams() *SecuritySamlSpCreateParams {
	return &SecuritySamlSpCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecuritySamlSpCreateParamsWithTimeout creates a new SecuritySamlSpCreateParams object
// with the ability to set a timeout on a request.
func NewSecuritySamlSpCreateParamsWithTimeout(timeout time.Duration) *SecuritySamlSpCreateParams {
	return &SecuritySamlSpCreateParams{
		timeout: timeout,
	}
}

// NewSecuritySamlSpCreateParamsWithContext creates a new SecuritySamlSpCreateParams object
// with the ability to set a context for a request.
func NewSecuritySamlSpCreateParamsWithContext(ctx context.Context) *SecuritySamlSpCreateParams {
	return &SecuritySamlSpCreateParams{
		Context: ctx,
	}
}

// NewSecuritySamlSpCreateParamsWithHTTPClient creates a new SecuritySamlSpCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecuritySamlSpCreateParamsWithHTTPClient(client *http.Client) *SecuritySamlSpCreateParams {
	return &SecuritySamlSpCreateParams{
		HTTPClient: client,
	}
}

/*
SecuritySamlSpCreateParams contains all the parameters to send to the API endpoint

	for the security saml sp create operation.

	Typically these are written to a http.Request.
*/
type SecuritySamlSpCreateParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.SecuritySamlSp

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* VerifyMetadataServer.

	   Verify IdP metadata server identity.

	   Default: true
	*/
	VerifyMetadataServer *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security saml sp create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecuritySamlSpCreateParams) WithDefaults() *SecuritySamlSpCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security saml sp create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecuritySamlSpCreateParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)

		verifyMetadataServerDefault = bool(true)
	)

	val := SecuritySamlSpCreateParams{
		ReturnTimeout:        &returnTimeoutDefault,
		VerifyMetadataServer: &verifyMetadataServerDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security saml sp create params
func (o *SecuritySamlSpCreateParams) WithTimeout(timeout time.Duration) *SecuritySamlSpCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security saml sp create params
func (o *SecuritySamlSpCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security saml sp create params
func (o *SecuritySamlSpCreateParams) WithContext(ctx context.Context) *SecuritySamlSpCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security saml sp create params
func (o *SecuritySamlSpCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security saml sp create params
func (o *SecuritySamlSpCreateParams) WithHTTPClient(client *http.Client) *SecuritySamlSpCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security saml sp create params
func (o *SecuritySamlSpCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the security saml sp create params
func (o *SecuritySamlSpCreateParams) WithInfo(info *models.SecuritySamlSp) *SecuritySamlSpCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security saml sp create params
func (o *SecuritySamlSpCreateParams) SetInfo(info *models.SecuritySamlSp) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the security saml sp create params
func (o *SecuritySamlSpCreateParams) WithReturnTimeout(returnTimeout *int64) *SecuritySamlSpCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security saml sp create params
func (o *SecuritySamlSpCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithVerifyMetadataServer adds the verifyMetadataServer to the security saml sp create params
func (o *SecuritySamlSpCreateParams) WithVerifyMetadataServer(verifyMetadataServer *bool) *SecuritySamlSpCreateParams {
	o.SetVerifyMetadataServer(verifyMetadataServer)
	return o
}

// SetVerifyMetadataServer adds the verifyMetadataServer to the security saml sp create params
func (o *SecuritySamlSpCreateParams) SetVerifyMetadataServer(verifyMetadataServer *bool) {
	o.VerifyMetadataServer = verifyMetadataServer
}

// WriteToRequest writes these params to a swagger request
func (o *SecuritySamlSpCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.VerifyMetadataServer != nil {

		// query param verify_metadata_server
		var qrVerifyMetadataServer bool

		if o.VerifyMetadataServer != nil {
			qrVerifyMetadataServer = *o.VerifyMetadataServer
		}
		qVerifyMetadataServer := swag.FormatBool(qrVerifyMetadataServer)
		if qVerifyMetadataServer != "" {

			if err := r.SetQueryParam("verify_metadata_server", qVerifyMetadataServer); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
