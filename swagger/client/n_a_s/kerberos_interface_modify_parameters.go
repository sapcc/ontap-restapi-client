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

// NewKerberosInterfaceModifyParams creates a new KerberosInterfaceModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKerberosInterfaceModifyParams() *KerberosInterfaceModifyParams {
	return &KerberosInterfaceModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKerberosInterfaceModifyParamsWithTimeout creates a new KerberosInterfaceModifyParams object
// with the ability to set a timeout on a request.
func NewKerberosInterfaceModifyParamsWithTimeout(timeout time.Duration) *KerberosInterfaceModifyParams {
	return &KerberosInterfaceModifyParams{
		timeout: timeout,
	}
}

// NewKerberosInterfaceModifyParamsWithContext creates a new KerberosInterfaceModifyParams object
// with the ability to set a context for a request.
func NewKerberosInterfaceModifyParamsWithContext(ctx context.Context) *KerberosInterfaceModifyParams {
	return &KerberosInterfaceModifyParams{
		Context: ctx,
	}
}

// NewKerberosInterfaceModifyParamsWithHTTPClient creates a new KerberosInterfaceModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewKerberosInterfaceModifyParamsWithHTTPClient(client *http.Client) *KerberosInterfaceModifyParams {
	return &KerberosInterfaceModifyParams{
		HTTPClient: client,
	}
}

/*
KerberosInterfaceModifyParams contains all the parameters to send to the API endpoint

	for the kerberos interface modify operation.

	Typically these are written to a http.Request.
*/
type KerberosInterfaceModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.KerberosInterface

	/* InterfaceUUID.

	   Network interface UUID
	*/
	InterfaceUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kerberos interface modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KerberosInterfaceModifyParams) WithDefaults() *KerberosInterfaceModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kerberos interface modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KerberosInterfaceModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kerberos interface modify params
func (o *KerberosInterfaceModifyParams) WithTimeout(timeout time.Duration) *KerberosInterfaceModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kerberos interface modify params
func (o *KerberosInterfaceModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kerberos interface modify params
func (o *KerberosInterfaceModifyParams) WithContext(ctx context.Context) *KerberosInterfaceModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kerberos interface modify params
func (o *KerberosInterfaceModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kerberos interface modify params
func (o *KerberosInterfaceModifyParams) WithHTTPClient(client *http.Client) *KerberosInterfaceModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kerberos interface modify params
func (o *KerberosInterfaceModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the kerberos interface modify params
func (o *KerberosInterfaceModifyParams) WithInfo(info *models.KerberosInterface) *KerberosInterfaceModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the kerberos interface modify params
func (o *KerberosInterfaceModifyParams) SetInfo(info *models.KerberosInterface) {
	o.Info = info
}

// WithInterfaceUUID adds the interfaceUUID to the kerberos interface modify params
func (o *KerberosInterfaceModifyParams) WithInterfaceUUID(interfaceUUID string) *KerberosInterfaceModifyParams {
	o.SetInterfaceUUID(interfaceUUID)
	return o
}

// SetInterfaceUUID adds the interfaceUuid to the kerberos interface modify params
func (o *KerberosInterfaceModifyParams) SetInterfaceUUID(interfaceUUID string) {
	o.InterfaceUUID = interfaceUUID
}

// WriteToRequest writes these params to a swagger request
func (o *KerberosInterfaceModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param interface.uuid
	if err := r.SetPathParam("interface.uuid", o.InterfaceUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
