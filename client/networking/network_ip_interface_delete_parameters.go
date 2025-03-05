// Code generated by go-swagger; DO NOT EDIT.

package networking

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
)

// NewNetworkIPInterfaceDeleteParams creates a new NetworkIPInterfaceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPInterfaceDeleteParams() *NetworkIPInterfaceDeleteParams {
	return &NetworkIPInterfaceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPInterfaceDeleteParamsWithTimeout creates a new NetworkIPInterfaceDeleteParams object
// with the ability to set a timeout on a request.
func NewNetworkIPInterfaceDeleteParamsWithTimeout(timeout time.Duration) *NetworkIPInterfaceDeleteParams {
	return &NetworkIPInterfaceDeleteParams{
		timeout: timeout,
	}
}

// NewNetworkIPInterfaceDeleteParamsWithContext creates a new NetworkIPInterfaceDeleteParams object
// with the ability to set a context for a request.
func NewNetworkIPInterfaceDeleteParamsWithContext(ctx context.Context) *NetworkIPInterfaceDeleteParams {
	return &NetworkIPInterfaceDeleteParams{
		Context: ctx,
	}
}

// NewNetworkIPInterfaceDeleteParamsWithHTTPClient creates a new NetworkIPInterfaceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPInterfaceDeleteParamsWithHTTPClient(client *http.Client) *NetworkIPInterfaceDeleteParams {
	return &NetworkIPInterfaceDeleteParams{
		HTTPClient: client,
	}
}

/*
NetworkIPInterfaceDeleteParams contains all the parameters to send to the API endpoint

	for the network ip interface delete operation.

	Typically these are written to a http.Request.
*/
type NetworkIPInterfaceDeleteParams struct {

	/* UUID.

	   IP interface UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip interface delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPInterfaceDeleteParams) WithDefaults() *NetworkIPInterfaceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip interface delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPInterfaceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network ip interface delete params
func (o *NetworkIPInterfaceDeleteParams) WithTimeout(timeout time.Duration) *NetworkIPInterfaceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip interface delete params
func (o *NetworkIPInterfaceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip interface delete params
func (o *NetworkIPInterfaceDeleteParams) WithContext(ctx context.Context) *NetworkIPInterfaceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip interface delete params
func (o *NetworkIPInterfaceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip interface delete params
func (o *NetworkIPInterfaceDeleteParams) WithHTTPClient(client *http.Client) *NetworkIPInterfaceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip interface delete params
func (o *NetworkIPInterfaceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the network ip interface delete params
func (o *NetworkIPInterfaceDeleteParams) WithUUID(uuid string) *NetworkIPInterfaceDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the network ip interface delete params
func (o *NetworkIPInterfaceDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPInterfaceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
