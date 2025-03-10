// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewLunAttributeModifyParams creates a new LunAttributeModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunAttributeModifyParams() *LunAttributeModifyParams {
	return &LunAttributeModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunAttributeModifyParamsWithTimeout creates a new LunAttributeModifyParams object
// with the ability to set a timeout on a request.
func NewLunAttributeModifyParamsWithTimeout(timeout time.Duration) *LunAttributeModifyParams {
	return &LunAttributeModifyParams{
		timeout: timeout,
	}
}

// NewLunAttributeModifyParamsWithContext creates a new LunAttributeModifyParams object
// with the ability to set a context for a request.
func NewLunAttributeModifyParamsWithContext(ctx context.Context) *LunAttributeModifyParams {
	return &LunAttributeModifyParams{
		Context: ctx,
	}
}

// NewLunAttributeModifyParamsWithHTTPClient creates a new LunAttributeModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunAttributeModifyParamsWithHTTPClient(client *http.Client) *LunAttributeModifyParams {
	return &LunAttributeModifyParams{
		HTTPClient: client,
	}
}

/*
LunAttributeModifyParams contains all the parameters to send to the API endpoint

	for the lun attribute modify operation.

	Typically these are written to a http.Request.
*/
type LunAttributeModifyParams struct {

	/* Info.

	   The new property values for the LUN.

	*/
	Info *models.LunAttribute

	/* LunUUID.

	   The unique identifier of the LUN.

	*/
	LunUUID string

	/* Name.

	   The name of the attribute.

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lun attribute modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunAttributeModifyParams) WithDefaults() *LunAttributeModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun attribute modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunAttributeModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lun attribute modify params
func (o *LunAttributeModifyParams) WithTimeout(timeout time.Duration) *LunAttributeModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun attribute modify params
func (o *LunAttributeModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun attribute modify params
func (o *LunAttributeModifyParams) WithContext(ctx context.Context) *LunAttributeModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun attribute modify params
func (o *LunAttributeModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun attribute modify params
func (o *LunAttributeModifyParams) WithHTTPClient(client *http.Client) *LunAttributeModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun attribute modify params
func (o *LunAttributeModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the lun attribute modify params
func (o *LunAttributeModifyParams) WithInfo(info *models.LunAttribute) *LunAttributeModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the lun attribute modify params
func (o *LunAttributeModifyParams) SetInfo(info *models.LunAttribute) {
	o.Info = info
}

// WithLunUUID adds the lunUUID to the lun attribute modify params
func (o *LunAttributeModifyParams) WithLunUUID(lunUUID string) *LunAttributeModifyParams {
	o.SetLunUUID(lunUUID)
	return o
}

// SetLunUUID adds the lunUuid to the lun attribute modify params
func (o *LunAttributeModifyParams) SetLunUUID(lunUUID string) {
	o.LunUUID = lunUUID
}

// WithName adds the name to the lun attribute modify params
func (o *LunAttributeModifyParams) WithName(name string) *LunAttributeModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the lun attribute modify params
func (o *LunAttributeModifyParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LunAttributeModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param lun.uuid
	if err := r.SetPathParam("lun.uuid", o.LunUUID); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
