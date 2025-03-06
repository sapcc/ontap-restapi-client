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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NewMultiAdminVerifyApprovalGroupModifyParams creates a new MultiAdminVerifyApprovalGroupModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyApprovalGroupModifyParams() *MultiAdminVerifyApprovalGroupModifyParams {
	return &MultiAdminVerifyApprovalGroupModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyApprovalGroupModifyParamsWithTimeout creates a new MultiAdminVerifyApprovalGroupModifyParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyApprovalGroupModifyParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyApprovalGroupModifyParams {
	return &MultiAdminVerifyApprovalGroupModifyParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyApprovalGroupModifyParamsWithContext creates a new MultiAdminVerifyApprovalGroupModifyParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyApprovalGroupModifyParamsWithContext(ctx context.Context) *MultiAdminVerifyApprovalGroupModifyParams {
	return &MultiAdminVerifyApprovalGroupModifyParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyApprovalGroupModifyParamsWithHTTPClient creates a new MultiAdminVerifyApprovalGroupModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyApprovalGroupModifyParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyApprovalGroupModifyParams {
	return &MultiAdminVerifyApprovalGroupModifyParams{
		HTTPClient: client,
	}
}

/*
MultiAdminVerifyApprovalGroupModifyParams contains all the parameters to send to the API endpoint

	for the multi admin verify approval group modify operation.

	Typically these are written to a http.Request.
*/
type MultiAdminVerifyApprovalGroupModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.MultiAdminVerifyApprovalGroup

	// Name.
	Name string

	// OwnerUUID.
	OwnerUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify approval group modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyApprovalGroupModifyParams) WithDefaults() *MultiAdminVerifyApprovalGroupModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify approval group modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyApprovalGroupModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyApprovalGroupModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) WithContext(ctx context.Context) *MultiAdminVerifyApprovalGroupModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyApprovalGroupModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) WithInfo(info *models.MultiAdminVerifyApprovalGroup) *MultiAdminVerifyApprovalGroupModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) SetInfo(info *models.MultiAdminVerifyApprovalGroup) {
	o.Info = info
}

// WithName adds the name to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) WithName(name string) *MultiAdminVerifyApprovalGroupModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) SetName(name string) {
	o.Name = name
}

// WithOwnerUUID adds the ownerUUID to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) WithOwnerUUID(ownerUUID string) *MultiAdminVerifyApprovalGroupModifyParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the multi admin verify approval group modify params
func (o *MultiAdminVerifyApprovalGroupModifyParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyApprovalGroupModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
