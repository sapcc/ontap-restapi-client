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

// NewNetworkIPInterfacesCreateParams creates a new NetworkIPInterfacesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPInterfacesCreateParams() *NetworkIPInterfacesCreateParams {
	return &NetworkIPInterfacesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPInterfacesCreateParamsWithTimeout creates a new NetworkIPInterfacesCreateParams object
// with the ability to set a timeout on a request.
func NewNetworkIPInterfacesCreateParamsWithTimeout(timeout time.Duration) *NetworkIPInterfacesCreateParams {
	return &NetworkIPInterfacesCreateParams{
		timeout: timeout,
	}
}

// NewNetworkIPInterfacesCreateParamsWithContext creates a new NetworkIPInterfacesCreateParams object
// with the ability to set a context for a request.
func NewNetworkIPInterfacesCreateParamsWithContext(ctx context.Context) *NetworkIPInterfacesCreateParams {
	return &NetworkIPInterfacesCreateParams{
		Context: ctx,
	}
}

// NewNetworkIPInterfacesCreateParamsWithHTTPClient creates a new NetworkIPInterfacesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPInterfacesCreateParamsWithHTTPClient(client *http.Client) *NetworkIPInterfacesCreateParams {
	return &NetworkIPInterfacesCreateParams{
		HTTPClient: client,
	}
}

/*
NetworkIPInterfacesCreateParams contains all the parameters to send to the API endpoint

	for the network ip interfaces create operation.

	Typically these are written to a http.Request.
*/
type NetworkIPInterfacesCreateParams struct {

	/* Info.

	   IP interface parameters
	*/
	Info *models.IPInterface

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip interfaces create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPInterfacesCreateParams) WithDefaults() *NetworkIPInterfacesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip interfaces create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPInterfacesCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := NetworkIPInterfacesCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the network ip interfaces create params
func (o *NetworkIPInterfacesCreateParams) WithTimeout(timeout time.Duration) *NetworkIPInterfacesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip interfaces create params
func (o *NetworkIPInterfacesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip interfaces create params
func (o *NetworkIPInterfacesCreateParams) WithContext(ctx context.Context) *NetworkIPInterfacesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip interfaces create params
func (o *NetworkIPInterfacesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip interfaces create params
func (o *NetworkIPInterfacesCreateParams) WithHTTPClient(client *http.Client) *NetworkIPInterfacesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip interfaces create params
func (o *NetworkIPInterfacesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the network ip interfaces create params
func (o *NetworkIPInterfacesCreateParams) WithInfo(info *models.IPInterface) *NetworkIPInterfacesCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the network ip interfaces create params
func (o *NetworkIPInterfacesCreateParams) SetInfo(info *models.IPInterface) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the network ip interfaces create params
func (o *NetworkIPInterfacesCreateParams) WithReturnRecords(returnRecords *bool) *NetworkIPInterfacesCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the network ip interfaces create params
func (o *NetworkIPInterfacesCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPInterfacesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
