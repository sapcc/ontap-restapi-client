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
	"github.com/go-openapi/swag"
)

// NewSecurityGroupGetParams creates a new SecurityGroupGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityGroupGetParams() *SecurityGroupGetParams {
	return &SecurityGroupGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityGroupGetParamsWithTimeout creates a new SecurityGroupGetParams object
// with the ability to set a timeout on a request.
func NewSecurityGroupGetParamsWithTimeout(timeout time.Duration) *SecurityGroupGetParams {
	return &SecurityGroupGetParams{
		timeout: timeout,
	}
}

// NewSecurityGroupGetParamsWithContext creates a new SecurityGroupGetParams object
// with the ability to set a context for a request.
func NewSecurityGroupGetParamsWithContext(ctx context.Context) *SecurityGroupGetParams {
	return &SecurityGroupGetParams{
		Context: ctx,
	}
}

// NewSecurityGroupGetParamsWithHTTPClient creates a new SecurityGroupGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityGroupGetParamsWithHTTPClient(client *http.Client) *SecurityGroupGetParams {
	return &SecurityGroupGetParams{
		HTTPClient: client,
	}
}

/*
SecurityGroupGetParams contains all the parameters to send to the API endpoint

	for the security group get operation.

	Typically these are written to a http.Request.
*/
type SecurityGroupGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Name.

	   Group name.
	*/
	Name string

	/* OwnerUUID.

	   Group owner. Used to identify a cluster or an SVM.
	*/
	OwnerUUID string

	/* Type.

	   Group type.
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security group get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityGroupGetParams) WithDefaults() *SecurityGroupGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security group get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityGroupGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security group get params
func (o *SecurityGroupGetParams) WithTimeout(timeout time.Duration) *SecurityGroupGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security group get params
func (o *SecurityGroupGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security group get params
func (o *SecurityGroupGetParams) WithContext(ctx context.Context) *SecurityGroupGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security group get params
func (o *SecurityGroupGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security group get params
func (o *SecurityGroupGetParams) WithHTTPClient(client *http.Client) *SecurityGroupGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security group get params
func (o *SecurityGroupGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the security group get params
func (o *SecurityGroupGetParams) WithFields(fields []string) *SecurityGroupGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the security group get params
func (o *SecurityGroupGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithName adds the name to the security group get params
func (o *SecurityGroupGetParams) WithName(name string) *SecurityGroupGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the security group get params
func (o *SecurityGroupGetParams) SetName(name string) {
	o.Name = name
}

// WithOwnerUUID adds the ownerUUID to the security group get params
func (o *SecurityGroupGetParams) WithOwnerUUID(ownerUUID string) *SecurityGroupGetParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the security group get params
func (o *SecurityGroupGetParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WithType adds the typeVar to the security group get params
func (o *SecurityGroupGetParams) WithType(typeVar string) *SecurityGroupGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the security group get params
func (o *SecurityGroupGetParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityGroupGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUID); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityGroupGet binds the parameter fields
func (o *SecurityGroupGetParams) bindParamFields(formats strfmt.Registry) []string {
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
