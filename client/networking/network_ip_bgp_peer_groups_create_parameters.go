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
	"github.com/go-openapi/swag"

	"github.com/sapcc/ontap-restapi-client/models"
)

// NewNetworkIPBgpPeerGroupsCreateParams creates a new NetworkIPBgpPeerGroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPBgpPeerGroupsCreateParams() *NetworkIPBgpPeerGroupsCreateParams {
	return &NetworkIPBgpPeerGroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPBgpPeerGroupsCreateParamsWithTimeout creates a new NetworkIPBgpPeerGroupsCreateParams object
// with the ability to set a timeout on a request.
func NewNetworkIPBgpPeerGroupsCreateParamsWithTimeout(timeout time.Duration) *NetworkIPBgpPeerGroupsCreateParams {
	return &NetworkIPBgpPeerGroupsCreateParams{
		timeout: timeout,
	}
}

// NewNetworkIPBgpPeerGroupsCreateParamsWithContext creates a new NetworkIPBgpPeerGroupsCreateParams object
// with the ability to set a context for a request.
func NewNetworkIPBgpPeerGroupsCreateParamsWithContext(ctx context.Context) *NetworkIPBgpPeerGroupsCreateParams {
	return &NetworkIPBgpPeerGroupsCreateParams{
		Context: ctx,
	}
}

// NewNetworkIPBgpPeerGroupsCreateParamsWithHTTPClient creates a new NetworkIPBgpPeerGroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPBgpPeerGroupsCreateParamsWithHTTPClient(client *http.Client) *NetworkIPBgpPeerGroupsCreateParams {
	return &NetworkIPBgpPeerGroupsCreateParams{
		HTTPClient: client,
	}
}

/*
NetworkIPBgpPeerGroupsCreateParams contains all the parameters to send to the API endpoint

	for the network ip bgp peer groups create operation.

	Typically these are written to a http.Request.
*/
type NetworkIPBgpPeerGroupsCreateParams struct {

	/* Info.

	   Peer group parameters
	*/
	Info *models.BgpPeerGroup

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip bgp peer groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPBgpPeerGroupsCreateParams) WithDefaults() *NetworkIPBgpPeerGroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip bgp peer groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPBgpPeerGroupsCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := NetworkIPBgpPeerGroupsCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the network ip bgp peer groups create params
func (o *NetworkIPBgpPeerGroupsCreateParams) WithTimeout(timeout time.Duration) *NetworkIPBgpPeerGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip bgp peer groups create params
func (o *NetworkIPBgpPeerGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip bgp peer groups create params
func (o *NetworkIPBgpPeerGroupsCreateParams) WithContext(ctx context.Context) *NetworkIPBgpPeerGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip bgp peer groups create params
func (o *NetworkIPBgpPeerGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip bgp peer groups create params
func (o *NetworkIPBgpPeerGroupsCreateParams) WithHTTPClient(client *http.Client) *NetworkIPBgpPeerGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip bgp peer groups create params
func (o *NetworkIPBgpPeerGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the network ip bgp peer groups create params
func (o *NetworkIPBgpPeerGroupsCreateParams) WithInfo(info *models.BgpPeerGroup) *NetworkIPBgpPeerGroupsCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the network ip bgp peer groups create params
func (o *NetworkIPBgpPeerGroupsCreateParams) SetInfo(info *models.BgpPeerGroup) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the network ip bgp peer groups create params
func (o *NetworkIPBgpPeerGroupsCreateParams) WithReturnRecords(returnRecords *bool) *NetworkIPBgpPeerGroupsCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the network ip bgp peer groups create params
func (o *NetworkIPBgpPeerGroupsCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPBgpPeerGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
