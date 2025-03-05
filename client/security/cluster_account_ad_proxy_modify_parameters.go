// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewClusterAccountAdProxyModifyParams creates a new ClusterAccountAdProxyModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterAccountAdProxyModifyParams() *ClusterAccountAdProxyModifyParams {
	return &ClusterAccountAdProxyModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterAccountAdProxyModifyParamsWithTimeout creates a new ClusterAccountAdProxyModifyParams object
// with the ability to set a timeout on a request.
func NewClusterAccountAdProxyModifyParamsWithTimeout(timeout time.Duration) *ClusterAccountAdProxyModifyParams {
	return &ClusterAccountAdProxyModifyParams{
		timeout: timeout,
	}
}

// NewClusterAccountAdProxyModifyParamsWithContext creates a new ClusterAccountAdProxyModifyParams object
// with the ability to set a context for a request.
func NewClusterAccountAdProxyModifyParamsWithContext(ctx context.Context) *ClusterAccountAdProxyModifyParams {
	return &ClusterAccountAdProxyModifyParams{
		Context: ctx,
	}
}

// NewClusterAccountAdProxyModifyParamsWithHTTPClient creates a new ClusterAccountAdProxyModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterAccountAdProxyModifyParamsWithHTTPClient(client *http.Client) *ClusterAccountAdProxyModifyParams {
	return &ClusterAccountAdProxyModifyParams{
		HTTPClient: client,
	}
}

/*
ClusterAccountAdProxyModifyParams contains all the parameters to send to the API endpoint

	for the cluster account ad proxy modify operation.

	Typically these are written to a http.Request.
*/
type ClusterAccountAdProxyModifyParams struct {

	/* Info.

	   The data SVM that tunnels the Active Directory authentication requests.
	*/
	Info *models.ClusterAdProxy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster account ad proxy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterAccountAdProxyModifyParams) WithDefaults() *ClusterAccountAdProxyModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster account ad proxy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterAccountAdProxyModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster account ad proxy modify params
func (o *ClusterAccountAdProxyModifyParams) WithTimeout(timeout time.Duration) *ClusterAccountAdProxyModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster account ad proxy modify params
func (o *ClusterAccountAdProxyModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster account ad proxy modify params
func (o *ClusterAccountAdProxyModifyParams) WithContext(ctx context.Context) *ClusterAccountAdProxyModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster account ad proxy modify params
func (o *ClusterAccountAdProxyModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster account ad proxy modify params
func (o *ClusterAccountAdProxyModifyParams) WithHTTPClient(client *http.Client) *ClusterAccountAdProxyModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster account ad proxy modify params
func (o *ClusterAccountAdProxyModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cluster account ad proxy modify params
func (o *ClusterAccountAdProxyModifyParams) WithInfo(info *models.ClusterAdProxy) *ClusterAccountAdProxyModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cluster account ad proxy modify params
func (o *ClusterAccountAdProxyModifyParams) SetInfo(info *models.ClusterAdProxy) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterAccountAdProxyModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
