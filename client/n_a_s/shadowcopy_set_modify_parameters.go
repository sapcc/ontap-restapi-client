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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NewShadowcopySetModifyParams creates a new ShadowcopySetModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShadowcopySetModifyParams() *ShadowcopySetModifyParams {
	return &ShadowcopySetModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShadowcopySetModifyParamsWithTimeout creates a new ShadowcopySetModifyParams object
// with the ability to set a timeout on a request.
func NewShadowcopySetModifyParamsWithTimeout(timeout time.Duration) *ShadowcopySetModifyParams {
	return &ShadowcopySetModifyParams{
		timeout: timeout,
	}
}

// NewShadowcopySetModifyParamsWithContext creates a new ShadowcopySetModifyParams object
// with the ability to set a context for a request.
func NewShadowcopySetModifyParamsWithContext(ctx context.Context) *ShadowcopySetModifyParams {
	return &ShadowcopySetModifyParams{
		Context: ctx,
	}
}

// NewShadowcopySetModifyParamsWithHTTPClient creates a new ShadowcopySetModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewShadowcopySetModifyParamsWithHTTPClient(client *http.Client) *ShadowcopySetModifyParams {
	return &ShadowcopySetModifyParams{
		HTTPClient: client,
	}
}

/*
ShadowcopySetModifyParams contains all the parameters to send to the API endpoint

	for the shadowcopy set modify operation.

	Typically these are written to a http.Request.
*/
type ShadowcopySetModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.ShadowcopySet

	/* UUID.

	   Storage shadowcopy set ID.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the shadowcopy set modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShadowcopySetModifyParams) WithDefaults() *ShadowcopySetModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the shadowcopy set modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShadowcopySetModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the shadowcopy set modify params
func (o *ShadowcopySetModifyParams) WithTimeout(timeout time.Duration) *ShadowcopySetModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shadowcopy set modify params
func (o *ShadowcopySetModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shadowcopy set modify params
func (o *ShadowcopySetModifyParams) WithContext(ctx context.Context) *ShadowcopySetModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shadowcopy set modify params
func (o *ShadowcopySetModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shadowcopy set modify params
func (o *ShadowcopySetModifyParams) WithHTTPClient(client *http.Client) *ShadowcopySetModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shadowcopy set modify params
func (o *ShadowcopySetModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the shadowcopy set modify params
func (o *ShadowcopySetModifyParams) WithInfo(info *models.ShadowcopySet) *ShadowcopySetModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the shadowcopy set modify params
func (o *ShadowcopySetModifyParams) SetInfo(info *models.ShadowcopySet) {
	o.Info = info
}

// WithUUID adds the uuid to the shadowcopy set modify params
func (o *ShadowcopySetModifyParams) WithUUID(uuid string) *ShadowcopySetModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the shadowcopy set modify params
func (o *ShadowcopySetModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ShadowcopySetModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
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
