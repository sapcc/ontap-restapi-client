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

// NewVscanOnDemandCreateParams creates a new VscanOnDemandCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanOnDemandCreateParams() *VscanOnDemandCreateParams {
	return &VscanOnDemandCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanOnDemandCreateParamsWithTimeout creates a new VscanOnDemandCreateParams object
// with the ability to set a timeout on a request.
func NewVscanOnDemandCreateParamsWithTimeout(timeout time.Duration) *VscanOnDemandCreateParams {
	return &VscanOnDemandCreateParams{
		timeout: timeout,
	}
}

// NewVscanOnDemandCreateParamsWithContext creates a new VscanOnDemandCreateParams object
// with the ability to set a context for a request.
func NewVscanOnDemandCreateParamsWithContext(ctx context.Context) *VscanOnDemandCreateParams {
	return &VscanOnDemandCreateParams{
		Context: ctx,
	}
}

// NewVscanOnDemandCreateParamsWithHTTPClient creates a new VscanOnDemandCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanOnDemandCreateParamsWithHTTPClient(client *http.Client) *VscanOnDemandCreateParams {
	return &VscanOnDemandCreateParams{
		HTTPClient: client,
	}
}

/*
VscanOnDemandCreateParams contains all the parameters to send to the API endpoint

	for the vscan on demand create operation.

	Typically these are written to a http.Request.
*/
type VscanOnDemandCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.VscanOnDemand

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan on demand create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnDemandCreateParams) WithDefaults() *VscanOnDemandCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan on demand create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnDemandCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := VscanOnDemandCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the vscan on demand create params
func (o *VscanOnDemandCreateParams) WithTimeout(timeout time.Duration) *VscanOnDemandCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan on demand create params
func (o *VscanOnDemandCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan on demand create params
func (o *VscanOnDemandCreateParams) WithContext(ctx context.Context) *VscanOnDemandCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan on demand create params
func (o *VscanOnDemandCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan on demand create params
func (o *VscanOnDemandCreateParams) WithHTTPClient(client *http.Client) *VscanOnDemandCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan on demand create params
func (o *VscanOnDemandCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the vscan on demand create params
func (o *VscanOnDemandCreateParams) WithInfo(info *models.VscanOnDemand) *VscanOnDemandCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the vscan on demand create params
func (o *VscanOnDemandCreateParams) SetInfo(info *models.VscanOnDemand) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the vscan on demand create params
func (o *VscanOnDemandCreateParams) WithReturnRecords(returnRecords *bool) *VscanOnDemandCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the vscan on demand create params
func (o *VscanOnDemandCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithSvmUUID adds the svmUUID to the vscan on demand create params
func (o *VscanOnDemandCreateParams) WithSvmUUID(svmUUID string) *VscanOnDemandCreateParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vscan on demand create params
func (o *VscanOnDemandCreateParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VscanOnDemandCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
