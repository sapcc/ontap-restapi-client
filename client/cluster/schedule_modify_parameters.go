// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewScheduleModifyParams creates a new ScheduleModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScheduleModifyParams() *ScheduleModifyParams {
	return &ScheduleModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScheduleModifyParamsWithTimeout creates a new ScheduleModifyParams object
// with the ability to set a timeout on a request.
func NewScheduleModifyParamsWithTimeout(timeout time.Duration) *ScheduleModifyParams {
	return &ScheduleModifyParams{
		timeout: timeout,
	}
}

// NewScheduleModifyParamsWithContext creates a new ScheduleModifyParams object
// with the ability to set a context for a request.
func NewScheduleModifyParamsWithContext(ctx context.Context) *ScheduleModifyParams {
	return &ScheduleModifyParams{
		Context: ctx,
	}
}

// NewScheduleModifyParamsWithHTTPClient creates a new ScheduleModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewScheduleModifyParamsWithHTTPClient(client *http.Client) *ScheduleModifyParams {
	return &ScheduleModifyParams{
		HTTPClient: client,
	}
}

/*
ScheduleModifyParams contains all the parameters to send to the API endpoint

	for the schedule modify operation.

	Typically these are written to a http.Request.
*/
type ScheduleModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Schedule

	/* UUID.

	   Schedule UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schedule modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleModifyParams) WithDefaults() *ScheduleModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schedule modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schedule modify params
func (o *ScheduleModifyParams) WithTimeout(timeout time.Duration) *ScheduleModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedule modify params
func (o *ScheduleModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedule modify params
func (o *ScheduleModifyParams) WithContext(ctx context.Context) *ScheduleModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedule modify params
func (o *ScheduleModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedule modify params
func (o *ScheduleModifyParams) WithHTTPClient(client *http.Client) *ScheduleModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedule modify params
func (o *ScheduleModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the schedule modify params
func (o *ScheduleModifyParams) WithInfo(info *models.Schedule) *ScheduleModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the schedule modify params
func (o *ScheduleModifyParams) SetInfo(info *models.Schedule) {
	o.Info = info
}

// WithUUID adds the uuid to the schedule modify params
func (o *ScheduleModifyParams) WithUUID(uuid string) *ScheduleModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the schedule modify params
func (o *ScheduleModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduleModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
