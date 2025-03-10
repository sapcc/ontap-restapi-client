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
	"github.com/go-openapi/swag"

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewMediatorDeleteParams creates a new MediatorDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMediatorDeleteParams() *MediatorDeleteParams {
	return &MediatorDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMediatorDeleteParamsWithTimeout creates a new MediatorDeleteParams object
// with the ability to set a timeout on a request.
func NewMediatorDeleteParamsWithTimeout(timeout time.Duration) *MediatorDeleteParams {
	return &MediatorDeleteParams{
		timeout: timeout,
	}
}

// NewMediatorDeleteParamsWithContext creates a new MediatorDeleteParams object
// with the ability to set a context for a request.
func NewMediatorDeleteParamsWithContext(ctx context.Context) *MediatorDeleteParams {
	return &MediatorDeleteParams{
		Context: ctx,
	}
}

// NewMediatorDeleteParamsWithHTTPClient creates a new MediatorDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewMediatorDeleteParamsWithHTTPClient(client *http.Client) *MediatorDeleteParams {
	return &MediatorDeleteParams{
		HTTPClient: client,
	}
}

/*
MediatorDeleteParams contains all the parameters to send to the API endpoint

	for the mediator delete operation.

	Typically these are written to a http.Request.
*/
type MediatorDeleteParams struct {

	/* Info.

	   Mediator information for the delete operation
	*/
	Info *models.Mediator

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	// UUID.
	//
	// Format: uuid
	UUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mediator delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediatorDeleteParams) WithDefaults() *MediatorDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mediator delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediatorDeleteParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := MediatorDeleteParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the mediator delete params
func (o *MediatorDeleteParams) WithTimeout(timeout time.Duration) *MediatorDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mediator delete params
func (o *MediatorDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mediator delete params
func (o *MediatorDeleteParams) WithContext(ctx context.Context) *MediatorDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mediator delete params
func (o *MediatorDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mediator delete params
func (o *MediatorDeleteParams) WithHTTPClient(client *http.Client) *MediatorDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mediator delete params
func (o *MediatorDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the mediator delete params
func (o *MediatorDeleteParams) WithInfo(info *models.Mediator) *MediatorDeleteParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the mediator delete params
func (o *MediatorDeleteParams) SetInfo(info *models.Mediator) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the mediator delete params
func (o *MediatorDeleteParams) WithReturnTimeout(returnTimeout *int64) *MediatorDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the mediator delete params
func (o *MediatorDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the mediator delete params
func (o *MediatorDeleteParams) WithUUID(uuid strfmt.UUID) *MediatorDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the mediator delete params
func (o *MediatorDeleteParams) SetUUID(uuid strfmt.UUID) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *MediatorDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
