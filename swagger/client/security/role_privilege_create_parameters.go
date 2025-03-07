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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewRolePrivilegeCreateParams creates a new RolePrivilegeCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRolePrivilegeCreateParams() *RolePrivilegeCreateParams {
	return &RolePrivilegeCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRolePrivilegeCreateParamsWithTimeout creates a new RolePrivilegeCreateParams object
// with the ability to set a timeout on a request.
func NewRolePrivilegeCreateParamsWithTimeout(timeout time.Duration) *RolePrivilegeCreateParams {
	return &RolePrivilegeCreateParams{
		timeout: timeout,
	}
}

// NewRolePrivilegeCreateParamsWithContext creates a new RolePrivilegeCreateParams object
// with the ability to set a context for a request.
func NewRolePrivilegeCreateParamsWithContext(ctx context.Context) *RolePrivilegeCreateParams {
	return &RolePrivilegeCreateParams{
		Context: ctx,
	}
}

// NewRolePrivilegeCreateParamsWithHTTPClient creates a new RolePrivilegeCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewRolePrivilegeCreateParamsWithHTTPClient(client *http.Client) *RolePrivilegeCreateParams {
	return &RolePrivilegeCreateParams{
		HTTPClient: client,
	}
}

/*
RolePrivilegeCreateParams contains all the parameters to send to the API endpoint

	for the role privilege create operation.

	Typically these are written to a http.Request.
*/
type RolePrivilegeCreateParams struct {

	/* Info.

	   Role privilege specification
	*/
	Info *models.RolePrivilege

	/* Name.

	   Role name
	*/
	Name string

	/* OwnerUUID.

	   Role owner UUID
	*/
	OwnerUUID string

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the role privilege create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolePrivilegeCreateParams) WithDefaults() *RolePrivilegeCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the role privilege create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolePrivilegeCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := RolePrivilegeCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the role privilege create params
func (o *RolePrivilegeCreateParams) WithTimeout(timeout time.Duration) *RolePrivilegeCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the role privilege create params
func (o *RolePrivilegeCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the role privilege create params
func (o *RolePrivilegeCreateParams) WithContext(ctx context.Context) *RolePrivilegeCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the role privilege create params
func (o *RolePrivilegeCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the role privilege create params
func (o *RolePrivilegeCreateParams) WithHTTPClient(client *http.Client) *RolePrivilegeCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the role privilege create params
func (o *RolePrivilegeCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the role privilege create params
func (o *RolePrivilegeCreateParams) WithInfo(info *models.RolePrivilege) *RolePrivilegeCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the role privilege create params
func (o *RolePrivilegeCreateParams) SetInfo(info *models.RolePrivilege) {
	o.Info = info
}

// WithName adds the name to the role privilege create params
func (o *RolePrivilegeCreateParams) WithName(name string) *RolePrivilegeCreateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the role privilege create params
func (o *RolePrivilegeCreateParams) SetName(name string) {
	o.Name = name
}

// WithOwnerUUID adds the ownerUUID to the role privilege create params
func (o *RolePrivilegeCreateParams) WithOwnerUUID(ownerUUID string) *RolePrivilegeCreateParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the role privilege create params
func (o *RolePrivilegeCreateParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WithReturnRecords adds the returnRecords to the role privilege create params
func (o *RolePrivilegeCreateParams) WithReturnRecords(returnRecords *bool) *RolePrivilegeCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the role privilege create params
func (o *RolePrivilegeCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *RolePrivilegeCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
