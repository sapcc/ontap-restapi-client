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
	"github.com/go-openapi/swag"
)

// NewIgroupInitiatorDeleteParams creates a new IgroupInitiatorDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIgroupInitiatorDeleteParams() *IgroupInitiatorDeleteParams {
	return &IgroupInitiatorDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIgroupInitiatorDeleteParamsWithTimeout creates a new IgroupInitiatorDeleteParams object
// with the ability to set a timeout on a request.
func NewIgroupInitiatorDeleteParamsWithTimeout(timeout time.Duration) *IgroupInitiatorDeleteParams {
	return &IgroupInitiatorDeleteParams{
		timeout: timeout,
	}
}

// NewIgroupInitiatorDeleteParamsWithContext creates a new IgroupInitiatorDeleteParams object
// with the ability to set a context for a request.
func NewIgroupInitiatorDeleteParamsWithContext(ctx context.Context) *IgroupInitiatorDeleteParams {
	return &IgroupInitiatorDeleteParams{
		Context: ctx,
	}
}

// NewIgroupInitiatorDeleteParamsWithHTTPClient creates a new IgroupInitiatorDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIgroupInitiatorDeleteParamsWithHTTPClient(client *http.Client) *IgroupInitiatorDeleteParams {
	return &IgroupInitiatorDeleteParams{
		HTTPClient: client,
	}
}

/*
IgroupInitiatorDeleteParams contains all the parameters to send to the API endpoint

	for the igroup initiator delete operation.

	Typically these are written to a http.Request.
*/
type IgroupInitiatorDeleteParams struct {

	/* AllowDeleteWhileMapped.

	     Allows the deletion of an initiator from of a mapped initiator group.<br/>
	Deleting an initiator from a mapped initiator group makes the LUNs to which the initiator group is mapped no longer available to the initiator. This might cause a disruption in the availability of data.<br/>
	<b>This parameter should be used with caution.</b>

	*/
	AllowDeleteWhileMapped *bool

	/* IgroupUUID.

	   The unique identifier of the initiator group.

	*/
	IgroupUUID string

	/* Name.

	   The initiator name.

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the igroup initiator delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupInitiatorDeleteParams) WithDefaults() *IgroupInitiatorDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the igroup initiator delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupInitiatorDeleteParams) SetDefaults() {
	var (
		allowDeleteWhileMappedDefault = bool(false)
	)

	val := IgroupInitiatorDeleteParams{
		AllowDeleteWhileMapped: &allowDeleteWhileMappedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) WithTimeout(timeout time.Duration) *IgroupInitiatorDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) WithContext(ctx context.Context) *IgroupInitiatorDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) WithHTTPClient(client *http.Client) *IgroupInitiatorDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowDeleteWhileMapped adds the allowDeleteWhileMapped to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) WithAllowDeleteWhileMapped(allowDeleteWhileMapped *bool) *IgroupInitiatorDeleteParams {
	o.SetAllowDeleteWhileMapped(allowDeleteWhileMapped)
	return o
}

// SetAllowDeleteWhileMapped adds the allowDeleteWhileMapped to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) SetAllowDeleteWhileMapped(allowDeleteWhileMapped *bool) {
	o.AllowDeleteWhileMapped = allowDeleteWhileMapped
}

// WithIgroupUUID adds the igroupUUID to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) WithIgroupUUID(igroupUUID string) *IgroupInitiatorDeleteParams {
	o.SetIgroupUUID(igroupUUID)
	return o
}

// SetIgroupUUID adds the igroupUuid to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) SetIgroupUUID(igroupUUID string) {
	o.IgroupUUID = igroupUUID
}

// WithName adds the name to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) WithName(name string) *IgroupInitiatorDeleteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the igroup initiator delete params
func (o *IgroupInitiatorDeleteParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *IgroupInitiatorDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowDeleteWhileMapped != nil {

		// query param allow_delete_while_mapped
		var qrAllowDeleteWhileMapped bool

		if o.AllowDeleteWhileMapped != nil {
			qrAllowDeleteWhileMapped = *o.AllowDeleteWhileMapped
		}
		qAllowDeleteWhileMapped := swag.FormatBool(qrAllowDeleteWhileMapped)
		if qAllowDeleteWhileMapped != "" {

			if err := r.SetQueryParam("allow_delete_while_mapped", qAllowDeleteWhileMapped); err != nil {
				return err
			}
		}
	}

	// path param igroup.uuid
	if err := r.SetPathParam("igroup.uuid", o.IgroupUUID); err != nil {
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
