// Code generated by go-swagger; DO NOT EDIT.

package n_d_m_p

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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewNdmpNodeModifyParams creates a new NdmpNodeModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNdmpNodeModifyParams() *NdmpNodeModifyParams {
	return &NdmpNodeModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNdmpNodeModifyParamsWithTimeout creates a new NdmpNodeModifyParams object
// with the ability to set a timeout on a request.
func NewNdmpNodeModifyParamsWithTimeout(timeout time.Duration) *NdmpNodeModifyParams {
	return &NdmpNodeModifyParams{
		timeout: timeout,
	}
}

// NewNdmpNodeModifyParamsWithContext creates a new NdmpNodeModifyParams object
// with the ability to set a context for a request.
func NewNdmpNodeModifyParamsWithContext(ctx context.Context) *NdmpNodeModifyParams {
	return &NdmpNodeModifyParams{
		Context: ctx,
	}
}

// NewNdmpNodeModifyParamsWithHTTPClient creates a new NdmpNodeModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewNdmpNodeModifyParamsWithHTTPClient(client *http.Client) *NdmpNodeModifyParams {
	return &NdmpNodeModifyParams{
		HTTPClient: client,
	}
}

/*
NdmpNodeModifyParams contains all the parameters to send to the API endpoint

	for the ndmp node modify operation.

	Typically these are written to a http.Request.
*/
type NdmpNodeModifyParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.NdmpNode

	/* NodeUUID.

	   Node UUID
	*/
	NodeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ndmp node modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpNodeModifyParams) WithDefaults() *NdmpNodeModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ndmp node modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpNodeModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ndmp node modify params
func (o *NdmpNodeModifyParams) WithTimeout(timeout time.Duration) *NdmpNodeModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ndmp node modify params
func (o *NdmpNodeModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ndmp node modify params
func (o *NdmpNodeModifyParams) WithContext(ctx context.Context) *NdmpNodeModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ndmp node modify params
func (o *NdmpNodeModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ndmp node modify params
func (o *NdmpNodeModifyParams) WithHTTPClient(client *http.Client) *NdmpNodeModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ndmp node modify params
func (o *NdmpNodeModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ndmp node modify params
func (o *NdmpNodeModifyParams) WithInfo(info *models.NdmpNode) *NdmpNodeModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ndmp node modify params
func (o *NdmpNodeModifyParams) SetInfo(info *models.NdmpNode) {
	o.Info = info
}

// WithNodeUUID adds the nodeUUID to the ndmp node modify params
func (o *NdmpNodeModifyParams) WithNodeUUID(nodeUUID string) *NdmpNodeModifyParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the ndmp node modify params
func (o *NdmpNodeModifyParams) SetNodeUUID(nodeUUID string) {
	o.NodeUUID = nodeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NdmpNodeModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
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
