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

// NewActiveDirectoryPreferredDcCreateParams creates a new ActiveDirectoryPreferredDcCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActiveDirectoryPreferredDcCreateParams() *ActiveDirectoryPreferredDcCreateParams {
	return &ActiveDirectoryPreferredDcCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActiveDirectoryPreferredDcCreateParamsWithTimeout creates a new ActiveDirectoryPreferredDcCreateParams object
// with the ability to set a timeout on a request.
func NewActiveDirectoryPreferredDcCreateParamsWithTimeout(timeout time.Duration) *ActiveDirectoryPreferredDcCreateParams {
	return &ActiveDirectoryPreferredDcCreateParams{
		timeout: timeout,
	}
}

// NewActiveDirectoryPreferredDcCreateParamsWithContext creates a new ActiveDirectoryPreferredDcCreateParams object
// with the ability to set a context for a request.
func NewActiveDirectoryPreferredDcCreateParamsWithContext(ctx context.Context) *ActiveDirectoryPreferredDcCreateParams {
	return &ActiveDirectoryPreferredDcCreateParams{
		Context: ctx,
	}
}

// NewActiveDirectoryPreferredDcCreateParamsWithHTTPClient creates a new ActiveDirectoryPreferredDcCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewActiveDirectoryPreferredDcCreateParamsWithHTTPClient(client *http.Client) *ActiveDirectoryPreferredDcCreateParams {
	return &ActiveDirectoryPreferredDcCreateParams{
		HTTPClient: client,
	}
}

/*
ActiveDirectoryPreferredDcCreateParams contains all the parameters to send to the API endpoint

	for the active directory preferred dc create operation.

	Typically these are written to a http.Request.
*/
type ActiveDirectoryPreferredDcCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.ActiveDirectoryPreferredDc

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* SkipConfigValidation.

	   Skip the validation of the specified preferred DC configuration.
	*/
	SkipConfigValidation *bool

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the active directory preferred dc create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActiveDirectoryPreferredDcCreateParams) WithDefaults() *ActiveDirectoryPreferredDcCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the active directory preferred dc create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActiveDirectoryPreferredDcCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		skipConfigValidationDefault = bool(false)
	)

	val := ActiveDirectoryPreferredDcCreateParams{
		ReturnRecords:        &returnRecordsDefault,
		SkipConfigValidation: &skipConfigValidationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) WithTimeout(timeout time.Duration) *ActiveDirectoryPreferredDcCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) WithContext(ctx context.Context) *ActiveDirectoryPreferredDcCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) WithHTTPClient(client *http.Client) *ActiveDirectoryPreferredDcCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) WithInfo(info *models.ActiveDirectoryPreferredDc) *ActiveDirectoryPreferredDcCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) SetInfo(info *models.ActiveDirectoryPreferredDc) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) WithReturnRecords(returnRecords *bool) *ActiveDirectoryPreferredDcCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithSkipConfigValidation adds the skipConfigValidation to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) WithSkipConfigValidation(skipConfigValidation *bool) *ActiveDirectoryPreferredDcCreateParams {
	o.SetSkipConfigValidation(skipConfigValidation)
	return o
}

// SetSkipConfigValidation adds the skipConfigValidation to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) SetSkipConfigValidation(skipConfigValidation *bool) {
	o.SkipConfigValidation = skipConfigValidation
}

// WithSvmUUID adds the svmUUID to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) WithSvmUUID(svmUUID string) *ActiveDirectoryPreferredDcCreateParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the active directory preferred dc create params
func (o *ActiveDirectoryPreferredDcCreateParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ActiveDirectoryPreferredDcCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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

	if o.SkipConfigValidation != nil {

		// query param skip_config_validation
		var qrSkipConfigValidation bool

		if o.SkipConfigValidation != nil {
			qrSkipConfigValidation = *o.SkipConfigValidation
		}
		qSkipConfigValidation := swag.FormatBool(qrSkipConfigValidation)
		if qSkipConfigValidation != "" {

			if err := r.SetQueryParam("skip_config_validation", qSkipConfigValidation); err != nil {
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
