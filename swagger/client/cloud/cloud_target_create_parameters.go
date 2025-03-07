// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewCloudTargetCreateParams creates a new CloudTargetCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudTargetCreateParams() *CloudTargetCreateParams {
	return &CloudTargetCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudTargetCreateParamsWithTimeout creates a new CloudTargetCreateParams object
// with the ability to set a timeout on a request.
func NewCloudTargetCreateParamsWithTimeout(timeout time.Duration) *CloudTargetCreateParams {
	return &CloudTargetCreateParams{
		timeout: timeout,
	}
}

// NewCloudTargetCreateParamsWithContext creates a new CloudTargetCreateParams object
// with the ability to set a context for a request.
func NewCloudTargetCreateParamsWithContext(ctx context.Context) *CloudTargetCreateParams {
	return &CloudTargetCreateParams{
		Context: ctx,
	}
}

// NewCloudTargetCreateParamsWithHTTPClient creates a new CloudTargetCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudTargetCreateParamsWithHTTPClient(client *http.Client) *CloudTargetCreateParams {
	return &CloudTargetCreateParams{
		HTTPClient: client,
	}
}

/*
CloudTargetCreateParams contains all the parameters to send to the API endpoint

	for the cloud target create operation.

	Typically these are written to a http.Request.
*/
type CloudTargetCreateParams struct {

	/* CheckOnly.

	   Do not create the target configuration, only check that the POST request succeeds.
	*/
	CheckOnly *bool

	/* IgnoreWarnings.

	   Specifies whether or not warning codes should be ignored.
	*/
	IgnoreWarnings *bool

	/* Info.

	   Info specification
	*/
	Info *models.CloudTarget

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

// WithDefaults hydrates default values in the cloud target create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudTargetCreateParams) WithDefaults() *CloudTargetCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud target create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudTargetCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := CloudTargetCreateParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cloud target create params
func (o *CloudTargetCreateParams) WithTimeout(timeout time.Duration) *CloudTargetCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud target create params
func (o *CloudTargetCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud target create params
func (o *CloudTargetCreateParams) WithContext(ctx context.Context) *CloudTargetCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud target create params
func (o *CloudTargetCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud target create params
func (o *CloudTargetCreateParams) WithHTTPClient(client *http.Client) *CloudTargetCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud target create params
func (o *CloudTargetCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckOnly adds the checkOnly to the cloud target create params
func (o *CloudTargetCreateParams) WithCheckOnly(checkOnly *bool) *CloudTargetCreateParams {
	o.SetCheckOnly(checkOnly)
	return o
}

// SetCheckOnly adds the checkOnly to the cloud target create params
func (o *CloudTargetCreateParams) SetCheckOnly(checkOnly *bool) {
	o.CheckOnly = checkOnly
}

// WithIgnoreWarnings adds the ignoreWarnings to the cloud target create params
func (o *CloudTargetCreateParams) WithIgnoreWarnings(ignoreWarnings *bool) *CloudTargetCreateParams {
	o.SetIgnoreWarnings(ignoreWarnings)
	return o
}

// SetIgnoreWarnings adds the ignoreWarnings to the cloud target create params
func (o *CloudTargetCreateParams) SetIgnoreWarnings(ignoreWarnings *bool) {
	o.IgnoreWarnings = ignoreWarnings
}

// WithInfo adds the info to the cloud target create params
func (o *CloudTargetCreateParams) WithInfo(info *models.CloudTarget) *CloudTargetCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cloud target create params
func (o *CloudTargetCreateParams) SetInfo(info *models.CloudTarget) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the cloud target create params
func (o *CloudTargetCreateParams) WithReturnRecords(returnRecords *bool) *CloudTargetCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the cloud target create params
func (o *CloudTargetCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the cloud target create params
func (o *CloudTargetCreateParams) WithReturnTimeout(returnTimeout *int64) *CloudTargetCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cloud target create params
func (o *CloudTargetCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *CloudTargetCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CheckOnly != nil {

		// query param check_only
		var qrCheckOnly bool

		if o.CheckOnly != nil {
			qrCheckOnly = *o.CheckOnly
		}
		qCheckOnly := swag.FormatBool(qrCheckOnly)
		if qCheckOnly != "" {

			if err := r.SetQueryParam("check_only", qCheckOnly); err != nil {
				return err
			}
		}
	}

	if o.IgnoreWarnings != nil {

		// query param ignore_warnings
		var qrIgnoreWarnings bool

		if o.IgnoreWarnings != nil {
			qrIgnoreWarnings = *o.IgnoreWarnings
		}
		qIgnoreWarnings := swag.FormatBool(qrIgnoreWarnings)
		if qIgnoreWarnings != "" {

			if err := r.SetQueryParam("ignore_warnings", qIgnoreWarnings); err != nil {
				return err
			}
		}
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
