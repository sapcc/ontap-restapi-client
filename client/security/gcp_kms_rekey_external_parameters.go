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

	"github.com/sapcc/ontap-restapi/models"
)

// NewGcpKmsRekeyExternalParams creates a new GcpKmsRekeyExternalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGcpKmsRekeyExternalParams() *GcpKmsRekeyExternalParams {
	return &GcpKmsRekeyExternalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGcpKmsRekeyExternalParamsWithTimeout creates a new GcpKmsRekeyExternalParams object
// with the ability to set a timeout on a request.
func NewGcpKmsRekeyExternalParamsWithTimeout(timeout time.Duration) *GcpKmsRekeyExternalParams {
	return &GcpKmsRekeyExternalParams{
		timeout: timeout,
	}
}

// NewGcpKmsRekeyExternalParamsWithContext creates a new GcpKmsRekeyExternalParams object
// with the ability to set a context for a request.
func NewGcpKmsRekeyExternalParamsWithContext(ctx context.Context) *GcpKmsRekeyExternalParams {
	return &GcpKmsRekeyExternalParams{
		Context: ctx,
	}
}

// NewGcpKmsRekeyExternalParamsWithHTTPClient creates a new GcpKmsRekeyExternalParams object
// with the ability to set a custom HTTPClient for a request.
func NewGcpKmsRekeyExternalParamsWithHTTPClient(client *http.Client) *GcpKmsRekeyExternalParams {
	return &GcpKmsRekeyExternalParams{
		HTTPClient: client,
	}
}

/*
GcpKmsRekeyExternalParams contains all the parameters to send to the API endpoint

	for the gcp kms rekey external operation.

	Typically these are written to a http.Request.
*/
type GcpKmsRekeyExternalParams struct {

	/* GcpKmsUUID.

	   UUID of the existing Google Cloud KMS configuration.
	*/
	GcpKmsUUID string

	/* Info.

	   Google Cloud KMS information.
	*/
	Info *models.GcpKmsKey

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gcp kms rekey external params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GcpKmsRekeyExternalParams) WithDefaults() *GcpKmsRekeyExternalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gcp kms rekey external params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GcpKmsRekeyExternalParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := GcpKmsRekeyExternalParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) WithTimeout(timeout time.Duration) *GcpKmsRekeyExternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) WithContext(ctx context.Context) *GcpKmsRekeyExternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) WithHTTPClient(client *http.Client) *GcpKmsRekeyExternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGcpKmsUUID adds the gcpKmsUUID to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) WithGcpKmsUUID(gcpKmsUUID string) *GcpKmsRekeyExternalParams {
	o.SetGcpKmsUUID(gcpKmsUUID)
	return o
}

// SetGcpKmsUUID adds the gcpKmsUuid to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) SetGcpKmsUUID(gcpKmsUUID string) {
	o.GcpKmsUUID = gcpKmsUUID
}

// WithInfo adds the info to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) WithInfo(info *models.GcpKmsKey) *GcpKmsRekeyExternalParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) SetInfo(info *models.GcpKmsKey) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) WithReturnRecords(returnRecords *bool) *GcpKmsRekeyExternalParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) WithReturnTimeout(returnTimeout *int64) *GcpKmsRekeyExternalParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the gcp kms rekey external params
func (o *GcpKmsRekeyExternalParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *GcpKmsRekeyExternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gcp_kms.uuid
	if err := r.SetPathParam("gcp_kms.uuid", o.GcpKmsUUID); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
