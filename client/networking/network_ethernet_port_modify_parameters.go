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

	"github.com/sapcc/ontap-restapi/models"
)

// NewNetworkEthernetPortModifyParams creates a new NetworkEthernetPortModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkEthernetPortModifyParams() *NetworkEthernetPortModifyParams {
	return &NetworkEthernetPortModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkEthernetPortModifyParamsWithTimeout creates a new NetworkEthernetPortModifyParams object
// with the ability to set a timeout on a request.
func NewNetworkEthernetPortModifyParamsWithTimeout(timeout time.Duration) *NetworkEthernetPortModifyParams {
	return &NetworkEthernetPortModifyParams{
		timeout: timeout,
	}
}

// NewNetworkEthernetPortModifyParamsWithContext creates a new NetworkEthernetPortModifyParams object
// with the ability to set a context for a request.
func NewNetworkEthernetPortModifyParamsWithContext(ctx context.Context) *NetworkEthernetPortModifyParams {
	return &NetworkEthernetPortModifyParams{
		Context: ctx,
	}
}

// NewNetworkEthernetPortModifyParamsWithHTTPClient creates a new NetworkEthernetPortModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkEthernetPortModifyParamsWithHTTPClient(client *http.Client) *NetworkEthernetPortModifyParams {
	return &NetworkEthernetPortModifyParams{
		HTTPClient: client,
	}
}

/*
NetworkEthernetPortModifyParams contains all the parameters to send to the API endpoint

	for the network ethernet port modify operation.

	Typically these are written to a http.Request.
*/
type NetworkEthernetPortModifyParams struct {

	// Info.
	Info *models.Port

	/* UUID.

	   Port UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ethernet port modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkEthernetPortModifyParams) WithDefaults() *NetworkEthernetPortModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ethernet port modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkEthernetPortModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network ethernet port modify params
func (o *NetworkEthernetPortModifyParams) WithTimeout(timeout time.Duration) *NetworkEthernetPortModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ethernet port modify params
func (o *NetworkEthernetPortModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ethernet port modify params
func (o *NetworkEthernetPortModifyParams) WithContext(ctx context.Context) *NetworkEthernetPortModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ethernet port modify params
func (o *NetworkEthernetPortModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ethernet port modify params
func (o *NetworkEthernetPortModifyParams) WithHTTPClient(client *http.Client) *NetworkEthernetPortModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ethernet port modify params
func (o *NetworkEthernetPortModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the network ethernet port modify params
func (o *NetworkEthernetPortModifyParams) WithInfo(info *models.Port) *NetworkEthernetPortModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the network ethernet port modify params
func (o *NetworkEthernetPortModifyParams) SetInfo(info *models.Port) {
	o.Info = info
}

// WithUUID adds the uuid to the network ethernet port modify params
func (o *NetworkEthernetPortModifyParams) WithUUID(uuid string) *NetworkEthernetPortModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the network ethernet port modify params
func (o *NetworkEthernetPortModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkEthernetPortModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
