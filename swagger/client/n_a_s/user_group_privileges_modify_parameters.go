// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewUserGroupPrivilegesModifyParams creates a new UserGroupPrivilegesModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserGroupPrivilegesModifyParams() *UserGroupPrivilegesModifyParams {
	return &UserGroupPrivilegesModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserGroupPrivilegesModifyParamsWithTimeout creates a new UserGroupPrivilegesModifyParams object
// with the ability to set a timeout on a request.
func NewUserGroupPrivilegesModifyParamsWithTimeout(timeout time.Duration) *UserGroupPrivilegesModifyParams {
	return &UserGroupPrivilegesModifyParams{
		timeout: timeout,
	}
}

// NewUserGroupPrivilegesModifyParamsWithContext creates a new UserGroupPrivilegesModifyParams object
// with the ability to set a context for a request.
func NewUserGroupPrivilegesModifyParamsWithContext(ctx context.Context) *UserGroupPrivilegesModifyParams {
	return &UserGroupPrivilegesModifyParams{
		Context: ctx,
	}
}

// NewUserGroupPrivilegesModifyParamsWithHTTPClient creates a new UserGroupPrivilegesModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserGroupPrivilegesModifyParamsWithHTTPClient(client *http.Client) *UserGroupPrivilegesModifyParams {
	return &UserGroupPrivilegesModifyParams{
		HTTPClient: client,
	}
}

/*
UserGroupPrivilegesModifyParams contains all the parameters to send to the API endpoint

	for the user group privileges modify operation.

	Typically these are written to a http.Request.
*/
type UserGroupPrivilegesModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.UserGroupPrivileges

	/* Name.

	   Local or Active Directory user or group name.
	*/
	Name string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user group privileges modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupPrivilegesModifyParams) WithDefaults() *UserGroupPrivilegesModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user group privileges modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupPrivilegesModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) WithTimeout(timeout time.Duration) *UserGroupPrivilegesModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) WithContext(ctx context.Context) *UserGroupPrivilegesModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) WithHTTPClient(client *http.Client) *UserGroupPrivilegesModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) WithInfo(info *models.UserGroupPrivileges) *UserGroupPrivilegesModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) SetInfo(info *models.UserGroupPrivileges) {
	o.Info = info
}

// WithName adds the name to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) WithName(name string) *UserGroupPrivilegesModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) WithSvmUUID(svmUUID string) *UserGroupPrivilegesModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the user group privileges modify params
func (o *UserGroupPrivilegesModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UserGroupPrivilegesModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
