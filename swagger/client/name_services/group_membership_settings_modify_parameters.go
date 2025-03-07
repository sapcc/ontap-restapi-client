// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewGroupMembershipSettingsModifyParams creates a new GroupMembershipSettingsModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupMembershipSettingsModifyParams() *GroupMembershipSettingsModifyParams {
	return &GroupMembershipSettingsModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupMembershipSettingsModifyParamsWithTimeout creates a new GroupMembershipSettingsModifyParams object
// with the ability to set a timeout on a request.
func NewGroupMembershipSettingsModifyParamsWithTimeout(timeout time.Duration) *GroupMembershipSettingsModifyParams {
	return &GroupMembershipSettingsModifyParams{
		timeout: timeout,
	}
}

// NewGroupMembershipSettingsModifyParamsWithContext creates a new GroupMembershipSettingsModifyParams object
// with the ability to set a context for a request.
func NewGroupMembershipSettingsModifyParamsWithContext(ctx context.Context) *GroupMembershipSettingsModifyParams {
	return &GroupMembershipSettingsModifyParams{
		Context: ctx,
	}
}

// NewGroupMembershipSettingsModifyParamsWithHTTPClient creates a new GroupMembershipSettingsModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupMembershipSettingsModifyParamsWithHTTPClient(client *http.Client) *GroupMembershipSettingsModifyParams {
	return &GroupMembershipSettingsModifyParams{
		HTTPClient: client,
	}
}

/*
GroupMembershipSettingsModifyParams contains all the parameters to send to the API endpoint

	for the group membership settings modify operation.

	Typically these are written to a http.Request.
*/
type GroupMembershipSettingsModifyParams struct {

	/* Info.

	   Info specification.
	*/
	Info *models.GroupMembershipSettings

	/* SvmUUID.

	   SVM UUID.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group membership settings modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupMembershipSettingsModifyParams) WithDefaults() *GroupMembershipSettingsModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group membership settings modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupMembershipSettingsModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group membership settings modify params
func (o *GroupMembershipSettingsModifyParams) WithTimeout(timeout time.Duration) *GroupMembershipSettingsModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group membership settings modify params
func (o *GroupMembershipSettingsModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group membership settings modify params
func (o *GroupMembershipSettingsModifyParams) WithContext(ctx context.Context) *GroupMembershipSettingsModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group membership settings modify params
func (o *GroupMembershipSettingsModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group membership settings modify params
func (o *GroupMembershipSettingsModifyParams) WithHTTPClient(client *http.Client) *GroupMembershipSettingsModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group membership settings modify params
func (o *GroupMembershipSettingsModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the group membership settings modify params
func (o *GroupMembershipSettingsModifyParams) WithInfo(info *models.GroupMembershipSettings) *GroupMembershipSettingsModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the group membership settings modify params
func (o *GroupMembershipSettingsModifyParams) SetInfo(info *models.GroupMembershipSettings) {
	o.Info = info
}

// WithSvmUUID adds the svmUUID to the group membership settings modify params
func (o *GroupMembershipSettingsModifyParams) WithSvmUUID(svmUUID string) *GroupMembershipSettingsModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the group membership settings modify params
func (o *GroupMembershipSettingsModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GroupMembershipSettingsModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
