// Code generated by go-swagger; DO NOT EDIT.

package s_vm

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

// NewSvmPeerPermissionModifyParams creates a new SvmPeerPermissionModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmPeerPermissionModifyParams() *SvmPeerPermissionModifyParams {
	return &SvmPeerPermissionModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmPeerPermissionModifyParamsWithTimeout creates a new SvmPeerPermissionModifyParams object
// with the ability to set a timeout on a request.
func NewSvmPeerPermissionModifyParamsWithTimeout(timeout time.Duration) *SvmPeerPermissionModifyParams {
	return &SvmPeerPermissionModifyParams{
		timeout: timeout,
	}
}

// NewSvmPeerPermissionModifyParamsWithContext creates a new SvmPeerPermissionModifyParams object
// with the ability to set a context for a request.
func NewSvmPeerPermissionModifyParamsWithContext(ctx context.Context) *SvmPeerPermissionModifyParams {
	return &SvmPeerPermissionModifyParams{
		Context: ctx,
	}
}

// NewSvmPeerPermissionModifyParamsWithHTTPClient creates a new SvmPeerPermissionModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmPeerPermissionModifyParamsWithHTTPClient(client *http.Client) *SvmPeerPermissionModifyParams {
	return &SvmPeerPermissionModifyParams{
		HTTPClient: client,
	}
}

/*
SvmPeerPermissionModifyParams contains all the parameters to send to the API endpoint

	for the svm peer permission modify operation.

	Typically these are written to a http.Request.
*/
type SvmPeerPermissionModifyParams struct {

	/* ClusterPeerUUID.

	   Peer cluster UUID
	*/
	ClusterPeerUUID string

	/* Info.

	   Info specification
	*/
	Info *models.SvmPeerPermission

	/* SvmUUID.

	   SVM UUID
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm peer permission modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerPermissionModifyParams) WithDefaults() *SvmPeerPermissionModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm peer permission modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerPermissionModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) WithTimeout(timeout time.Duration) *SvmPeerPermissionModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) WithContext(ctx context.Context) *SvmPeerPermissionModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) WithHTTPClient(client *http.Client) *SvmPeerPermissionModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterPeerUUID adds the clusterPeerUUID to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) WithClusterPeerUUID(clusterPeerUUID string) *SvmPeerPermissionModifyParams {
	o.SetClusterPeerUUID(clusterPeerUUID)
	return o
}

// SetClusterPeerUUID adds the clusterPeerUuid to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) SetClusterPeerUUID(clusterPeerUUID string) {
	o.ClusterPeerUUID = clusterPeerUUID
}

// WithInfo adds the info to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) WithInfo(info *models.SvmPeerPermission) *SvmPeerPermissionModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) SetInfo(info *models.SvmPeerPermission) {
	o.Info = info
}

// WithSvmUUID adds the svmUUID to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) WithSvmUUID(svmUUID string) *SvmPeerPermissionModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the svm peer permission modify params
func (o *SvmPeerPermissionModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SvmPeerPermissionModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_peer.uuid
	if err := r.SetPathParam("cluster_peer.uuid", o.ClusterPeerUUID); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
