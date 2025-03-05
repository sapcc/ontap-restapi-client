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
)

// NewMetroclusterNodeGetParams creates a new MetroclusterNodeGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetroclusterNodeGetParams() *MetroclusterNodeGetParams {
	return &MetroclusterNodeGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetroclusterNodeGetParamsWithTimeout creates a new MetroclusterNodeGetParams object
// with the ability to set a timeout on a request.
func NewMetroclusterNodeGetParamsWithTimeout(timeout time.Duration) *MetroclusterNodeGetParams {
	return &MetroclusterNodeGetParams{
		timeout: timeout,
	}
}

// NewMetroclusterNodeGetParamsWithContext creates a new MetroclusterNodeGetParams object
// with the ability to set a context for a request.
func NewMetroclusterNodeGetParamsWithContext(ctx context.Context) *MetroclusterNodeGetParams {
	return &MetroclusterNodeGetParams{
		Context: ctx,
	}
}

// NewMetroclusterNodeGetParamsWithHTTPClient creates a new MetroclusterNodeGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMetroclusterNodeGetParamsWithHTTPClient(client *http.Client) *MetroclusterNodeGetParams {
	return &MetroclusterNodeGetParams{
		HTTPClient: client,
	}
}

/*
MetroclusterNodeGetParams contains all the parameters to send to the API endpoint

	for the metrocluster node get operation.

	Typically these are written to a http.Request.
*/
type MetroclusterNodeGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* NodeUUID.

	   Node UUID
	*/
	NodeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the metrocluster node get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterNodeGetParams) WithDefaults() *MetroclusterNodeGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metrocluster node get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterNodeGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the metrocluster node get params
func (o *MetroclusterNodeGetParams) WithTimeout(timeout time.Duration) *MetroclusterNodeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metrocluster node get params
func (o *MetroclusterNodeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metrocluster node get params
func (o *MetroclusterNodeGetParams) WithContext(ctx context.Context) *MetroclusterNodeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metrocluster node get params
func (o *MetroclusterNodeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metrocluster node get params
func (o *MetroclusterNodeGetParams) WithHTTPClient(client *http.Client) *MetroclusterNodeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metrocluster node get params
func (o *MetroclusterNodeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the metrocluster node get params
func (o *MetroclusterNodeGetParams) WithFields(fields []string) *MetroclusterNodeGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the metrocluster node get params
func (o *MetroclusterNodeGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithNodeUUID adds the nodeUUID to the metrocluster node get params
func (o *MetroclusterNodeGetParams) WithNodeUUID(nodeUUID string) *MetroclusterNodeGetParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the metrocluster node get params
func (o *MetroclusterNodeGetParams) SetNodeUUID(nodeUUID string) {
	o.NodeUUID = nodeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *MetroclusterNodeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamMetroclusterNodeGet binds the parameter fields
func (o *MetroclusterNodeGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
