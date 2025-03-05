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

	"github.com/sapcc/ontap-restapi/models"
)

// NewPortsetInterfaceCreateParams creates a new PortsetInterfaceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPortsetInterfaceCreateParams() *PortsetInterfaceCreateParams {
	return &PortsetInterfaceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPortsetInterfaceCreateParamsWithTimeout creates a new PortsetInterfaceCreateParams object
// with the ability to set a timeout on a request.
func NewPortsetInterfaceCreateParamsWithTimeout(timeout time.Duration) *PortsetInterfaceCreateParams {
	return &PortsetInterfaceCreateParams{
		timeout: timeout,
	}
}

// NewPortsetInterfaceCreateParamsWithContext creates a new PortsetInterfaceCreateParams object
// with the ability to set a context for a request.
func NewPortsetInterfaceCreateParamsWithContext(ctx context.Context) *PortsetInterfaceCreateParams {
	return &PortsetInterfaceCreateParams{
		Context: ctx,
	}
}

// NewPortsetInterfaceCreateParamsWithHTTPClient creates a new PortsetInterfaceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPortsetInterfaceCreateParamsWithHTTPClient(client *http.Client) *PortsetInterfaceCreateParams {
	return &PortsetInterfaceCreateParams{
		HTTPClient: client,
	}
}

/*
PortsetInterfaceCreateParams contains all the parameters to send to the API endpoint

	for the portset interface create operation.

	Typically these are written to a http.Request.
*/
type PortsetInterfaceCreateParams struct {

	/* Info.

	   The properties of the interface to add to the portset.

	*/
	Info *models.PortsetInterface

	/* PortsetUUID.

	   The unique identifier of the portset.

	*/
	PortsetUUID string

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the portset interface create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortsetInterfaceCreateParams) WithDefaults() *PortsetInterfaceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the portset interface create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortsetInterfaceCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := PortsetInterfaceCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the portset interface create params
func (o *PortsetInterfaceCreateParams) WithTimeout(timeout time.Duration) *PortsetInterfaceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the portset interface create params
func (o *PortsetInterfaceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the portset interface create params
func (o *PortsetInterfaceCreateParams) WithContext(ctx context.Context) *PortsetInterfaceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the portset interface create params
func (o *PortsetInterfaceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the portset interface create params
func (o *PortsetInterfaceCreateParams) WithHTTPClient(client *http.Client) *PortsetInterfaceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the portset interface create params
func (o *PortsetInterfaceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the portset interface create params
func (o *PortsetInterfaceCreateParams) WithInfo(info *models.PortsetInterface) *PortsetInterfaceCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the portset interface create params
func (o *PortsetInterfaceCreateParams) SetInfo(info *models.PortsetInterface) {
	o.Info = info
}

// WithPortsetUUID adds the portsetUUID to the portset interface create params
func (o *PortsetInterfaceCreateParams) WithPortsetUUID(portsetUUID string) *PortsetInterfaceCreateParams {
	o.SetPortsetUUID(portsetUUID)
	return o
}

// SetPortsetUUID adds the portsetUuid to the portset interface create params
func (o *PortsetInterfaceCreateParams) SetPortsetUUID(portsetUUID string) {
	o.PortsetUUID = portsetUUID
}

// WithReturnRecords adds the returnRecords to the portset interface create params
func (o *PortsetInterfaceCreateParams) WithReturnRecords(returnRecords *bool) *PortsetInterfaceCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the portset interface create params
func (o *PortsetInterfaceCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *PortsetInterfaceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param portset.uuid
	if err := r.SetPathParam("portset.uuid", o.PortsetUUID); err != nil {
		return err
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
