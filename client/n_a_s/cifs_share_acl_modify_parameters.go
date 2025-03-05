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

	"github.com/sapcc/ontap-restapi/models"
)

// NewCifsShareACLModifyParams creates a new CifsShareACLModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsShareACLModifyParams() *CifsShareACLModifyParams {
	return &CifsShareACLModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsShareACLModifyParamsWithTimeout creates a new CifsShareACLModifyParams object
// with the ability to set a timeout on a request.
func NewCifsShareACLModifyParamsWithTimeout(timeout time.Duration) *CifsShareACLModifyParams {
	return &CifsShareACLModifyParams{
		timeout: timeout,
	}
}

// NewCifsShareACLModifyParamsWithContext creates a new CifsShareACLModifyParams object
// with the ability to set a context for a request.
func NewCifsShareACLModifyParamsWithContext(ctx context.Context) *CifsShareACLModifyParams {
	return &CifsShareACLModifyParams{
		Context: ctx,
	}
}

// NewCifsShareACLModifyParamsWithHTTPClient creates a new CifsShareACLModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsShareACLModifyParamsWithHTTPClient(client *http.Client) *CifsShareACLModifyParams {
	return &CifsShareACLModifyParams{
		HTTPClient: client,
	}
}

/*
CifsShareACLModifyParams contains all the parameters to send to the API endpoint

	for the cifs share acl modify operation.

	Typically these are written to a http.Request.
*/
type CifsShareACLModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.CifsShareACL

	/* Share.

	   Share name
	*/
	Share string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	/* Type.

	   User or group type
	*/
	Type string

	/* UserOrGroup.

	   User or group name
	*/
	UserOrGroup string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs share acl modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsShareACLModifyParams) WithDefaults() *CifsShareACLModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs share acl modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsShareACLModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cifs share acl modify params
func (o *CifsShareACLModifyParams) WithTimeout(timeout time.Duration) *CifsShareACLModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs share acl modify params
func (o *CifsShareACLModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs share acl modify params
func (o *CifsShareACLModifyParams) WithContext(ctx context.Context) *CifsShareACLModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs share acl modify params
func (o *CifsShareACLModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs share acl modify params
func (o *CifsShareACLModifyParams) WithHTTPClient(client *http.Client) *CifsShareACLModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs share acl modify params
func (o *CifsShareACLModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cifs share acl modify params
func (o *CifsShareACLModifyParams) WithInfo(info *models.CifsShareACL) *CifsShareACLModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cifs share acl modify params
func (o *CifsShareACLModifyParams) SetInfo(info *models.CifsShareACL) {
	o.Info = info
}

// WithShare adds the share to the cifs share acl modify params
func (o *CifsShareACLModifyParams) WithShare(share string) *CifsShareACLModifyParams {
	o.SetShare(share)
	return o
}

// SetShare adds the share to the cifs share acl modify params
func (o *CifsShareACLModifyParams) SetShare(share string) {
	o.Share = share
}

// WithSvmUUID adds the svmUUID to the cifs share acl modify params
func (o *CifsShareACLModifyParams) WithSvmUUID(svmUUID string) *CifsShareACLModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the cifs share acl modify params
func (o *CifsShareACLModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithType adds the typeVar to the cifs share acl modify params
func (o *CifsShareACLModifyParams) WithType(typeVar string) *CifsShareACLModifyParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the cifs share acl modify params
func (o *CifsShareACLModifyParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WithUserOrGroup adds the userOrGroup to the cifs share acl modify params
func (o *CifsShareACLModifyParams) WithUserOrGroup(userOrGroup string) *CifsShareACLModifyParams {
	o.SetUserOrGroup(userOrGroup)
	return o
}

// SetUserOrGroup adds the userOrGroup to the cifs share acl modify params
func (o *CifsShareACLModifyParams) SetUserOrGroup(userOrGroup string) {
	o.UserOrGroup = userOrGroup
}

// WriteToRequest writes these params to a swagger request
func (o *CifsShareACLModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param share
	if err := r.SetPathParam("share", o.Share); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	// path param user_or_group
	if err := r.SetPathParam("user_or_group", o.UserOrGroup); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
