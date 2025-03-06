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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NewMultiAdminVerifyApprovalGroupCreateParams creates a new MultiAdminVerifyApprovalGroupCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyApprovalGroupCreateParams() *MultiAdminVerifyApprovalGroupCreateParams {
	return &MultiAdminVerifyApprovalGroupCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyApprovalGroupCreateParamsWithTimeout creates a new MultiAdminVerifyApprovalGroupCreateParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyApprovalGroupCreateParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyApprovalGroupCreateParams {
	return &MultiAdminVerifyApprovalGroupCreateParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyApprovalGroupCreateParamsWithContext creates a new MultiAdminVerifyApprovalGroupCreateParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyApprovalGroupCreateParamsWithContext(ctx context.Context) *MultiAdminVerifyApprovalGroupCreateParams {
	return &MultiAdminVerifyApprovalGroupCreateParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyApprovalGroupCreateParamsWithHTTPClient creates a new MultiAdminVerifyApprovalGroupCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyApprovalGroupCreateParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyApprovalGroupCreateParams {
	return &MultiAdminVerifyApprovalGroupCreateParams{
		HTTPClient: client,
	}
}

/*
MultiAdminVerifyApprovalGroupCreateParams contains all the parameters to send to the API endpoint

	for the multi admin verify approval group create operation.

	Typically these are written to a http.Request.
*/
type MultiAdminVerifyApprovalGroupCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.MultiAdminVerifyApprovalGroup

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify approval group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyApprovalGroupCreateParams) WithDefaults() *MultiAdminVerifyApprovalGroupCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify approval group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyApprovalGroupCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := MultiAdminVerifyApprovalGroupCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the multi admin verify approval group create params
func (o *MultiAdminVerifyApprovalGroupCreateParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyApprovalGroupCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify approval group create params
func (o *MultiAdminVerifyApprovalGroupCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify approval group create params
func (o *MultiAdminVerifyApprovalGroupCreateParams) WithContext(ctx context.Context) *MultiAdminVerifyApprovalGroupCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify approval group create params
func (o *MultiAdminVerifyApprovalGroupCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify approval group create params
func (o *MultiAdminVerifyApprovalGroupCreateParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyApprovalGroupCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify approval group create params
func (o *MultiAdminVerifyApprovalGroupCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the multi admin verify approval group create params
func (o *MultiAdminVerifyApprovalGroupCreateParams) WithInfo(info *models.MultiAdminVerifyApprovalGroup) *MultiAdminVerifyApprovalGroupCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the multi admin verify approval group create params
func (o *MultiAdminVerifyApprovalGroupCreateParams) SetInfo(info *models.MultiAdminVerifyApprovalGroup) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the multi admin verify approval group create params
func (o *MultiAdminVerifyApprovalGroupCreateParams) WithReturnRecords(returnRecords *bool) *MultiAdminVerifyApprovalGroupCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the multi admin verify approval group create params
func (o *MultiAdminVerifyApprovalGroupCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyApprovalGroupCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
