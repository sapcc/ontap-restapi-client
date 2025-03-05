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

	"github.com/sapcc/ontap-restapi/models"
)

// NewLdapSchemaModifyParams creates a new LdapSchemaModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLdapSchemaModifyParams() *LdapSchemaModifyParams {
	return &LdapSchemaModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLdapSchemaModifyParamsWithTimeout creates a new LdapSchemaModifyParams object
// with the ability to set a timeout on a request.
func NewLdapSchemaModifyParamsWithTimeout(timeout time.Duration) *LdapSchemaModifyParams {
	return &LdapSchemaModifyParams{
		timeout: timeout,
	}
}

// NewLdapSchemaModifyParamsWithContext creates a new LdapSchemaModifyParams object
// with the ability to set a context for a request.
func NewLdapSchemaModifyParamsWithContext(ctx context.Context) *LdapSchemaModifyParams {
	return &LdapSchemaModifyParams{
		Context: ctx,
	}
}

// NewLdapSchemaModifyParamsWithHTTPClient creates a new LdapSchemaModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewLdapSchemaModifyParamsWithHTTPClient(client *http.Client) *LdapSchemaModifyParams {
	return &LdapSchemaModifyParams{
		HTTPClient: client,
	}
}

/*
LdapSchemaModifyParams contains all the parameters to send to the API endpoint

	for the ldap schema modify operation.

	Typically these are written to a http.Request.
*/
type LdapSchemaModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.LdapSchema

	/* Name.

	   LDAP schema name.
	*/
	Name string

	/* OwnerUUID.

	   UUID of the owner to which this object belongs.
	*/
	OwnerUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ldap schema modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapSchemaModifyParams) WithDefaults() *LdapSchemaModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ldap schema modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapSchemaModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ldap schema modify params
func (o *LdapSchemaModifyParams) WithTimeout(timeout time.Duration) *LdapSchemaModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ldap schema modify params
func (o *LdapSchemaModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ldap schema modify params
func (o *LdapSchemaModifyParams) WithContext(ctx context.Context) *LdapSchemaModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ldap schema modify params
func (o *LdapSchemaModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ldap schema modify params
func (o *LdapSchemaModifyParams) WithHTTPClient(client *http.Client) *LdapSchemaModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ldap schema modify params
func (o *LdapSchemaModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ldap schema modify params
func (o *LdapSchemaModifyParams) WithInfo(info *models.LdapSchema) *LdapSchemaModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ldap schema modify params
func (o *LdapSchemaModifyParams) SetInfo(info *models.LdapSchema) {
	o.Info = info
}

// WithName adds the name to the ldap schema modify params
func (o *LdapSchemaModifyParams) WithName(name string) *LdapSchemaModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ldap schema modify params
func (o *LdapSchemaModifyParams) SetName(name string) {
	o.Name = name
}

// WithOwnerUUID adds the ownerUUID to the ldap schema modify params
func (o *LdapSchemaModifyParams) WithOwnerUUID(ownerUUID string) *LdapSchemaModifyParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the ldap schema modify params
func (o *LdapSchemaModifyParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LdapSchemaModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
