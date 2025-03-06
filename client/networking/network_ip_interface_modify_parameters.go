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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NewNetworkIPInterfaceModifyParams creates a new NetworkIPInterfaceModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPInterfaceModifyParams() *NetworkIPInterfaceModifyParams {
	return &NetworkIPInterfaceModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPInterfaceModifyParamsWithTimeout creates a new NetworkIPInterfaceModifyParams object
// with the ability to set a timeout on a request.
func NewNetworkIPInterfaceModifyParamsWithTimeout(timeout time.Duration) *NetworkIPInterfaceModifyParams {
	return &NetworkIPInterfaceModifyParams{
		timeout: timeout,
	}
}

// NewNetworkIPInterfaceModifyParamsWithContext creates a new NetworkIPInterfaceModifyParams object
// with the ability to set a context for a request.
func NewNetworkIPInterfaceModifyParamsWithContext(ctx context.Context) *NetworkIPInterfaceModifyParams {
	return &NetworkIPInterfaceModifyParams{
		Context: ctx,
	}
}

// NewNetworkIPInterfaceModifyParamsWithHTTPClient creates a new NetworkIPInterfaceModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPInterfaceModifyParamsWithHTTPClient(client *http.Client) *NetworkIPInterfaceModifyParams {
	return &NetworkIPInterfaceModifyParams{
		HTTPClient: client,
	}
}

/*
NetworkIPInterfaceModifyParams contains all the parameters to send to the API endpoint

	for the network ip interface modify operation.

	Typically these are written to a http.Request.
*/
type NetworkIPInterfaceModifyParams struct {

	// Info.
	Info *models.IPInterface

	/* UUID.

	   IP interface UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip interface modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPInterfaceModifyParams) WithDefaults() *NetworkIPInterfaceModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip interface modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPInterfaceModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network ip interface modify params
func (o *NetworkIPInterfaceModifyParams) WithTimeout(timeout time.Duration) *NetworkIPInterfaceModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip interface modify params
func (o *NetworkIPInterfaceModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip interface modify params
func (o *NetworkIPInterfaceModifyParams) WithContext(ctx context.Context) *NetworkIPInterfaceModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip interface modify params
func (o *NetworkIPInterfaceModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip interface modify params
func (o *NetworkIPInterfaceModifyParams) WithHTTPClient(client *http.Client) *NetworkIPInterfaceModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip interface modify params
func (o *NetworkIPInterfaceModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the network ip interface modify params
func (o *NetworkIPInterfaceModifyParams) WithInfo(info *models.IPInterface) *NetworkIPInterfaceModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the network ip interface modify params
func (o *NetworkIPInterfaceModifyParams) SetInfo(info *models.IPInterface) {
	o.Info = info
}

// WithUUID adds the uuid to the network ip interface modify params
func (o *NetworkIPInterfaceModifyParams) WithUUID(uuid string) *NetworkIPInterfaceModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the network ip interface modify params
func (o *NetworkIPInterfaceModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPInterfaceModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
