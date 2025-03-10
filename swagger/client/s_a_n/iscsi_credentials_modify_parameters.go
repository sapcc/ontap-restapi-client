// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewIscsiCredentialsModifyParams creates a new IscsiCredentialsModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIscsiCredentialsModifyParams() *IscsiCredentialsModifyParams {
	return &IscsiCredentialsModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIscsiCredentialsModifyParamsWithTimeout creates a new IscsiCredentialsModifyParams object
// with the ability to set a timeout on a request.
func NewIscsiCredentialsModifyParamsWithTimeout(timeout time.Duration) *IscsiCredentialsModifyParams {
	return &IscsiCredentialsModifyParams{
		timeout: timeout,
	}
}

// NewIscsiCredentialsModifyParamsWithContext creates a new IscsiCredentialsModifyParams object
// with the ability to set a context for a request.
func NewIscsiCredentialsModifyParamsWithContext(ctx context.Context) *IscsiCredentialsModifyParams {
	return &IscsiCredentialsModifyParams{
		Context: ctx,
	}
}

// NewIscsiCredentialsModifyParamsWithHTTPClient creates a new IscsiCredentialsModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewIscsiCredentialsModifyParamsWithHTTPClient(client *http.Client) *IscsiCredentialsModifyParams {
	return &IscsiCredentialsModifyParams{
		HTTPClient: client,
	}
}

/*
IscsiCredentialsModifyParams contains all the parameters to send to the API endpoint

	for the iscsi credentials modify operation.

	Typically these are written to a http.Request.
*/
type IscsiCredentialsModifyParams struct {

	/* AddInitiatorAddresses.

	   If _true_, the initiator addresses in the body merge into the existing addresses in the iSCSI security object rather than replace the existing addresses.

	*/
	AddInitiatorAddresses *bool

	/* Info.

	   The new property values for the iSCSI credentials object.

	*/
	Info *models.IscsiCredentials

	/* Initiator.

	   The iSCSI initiator of the credentials object.

	*/
	Initiator string

	/* RemoveInitiatorAddresses.

	   If _true_, the initiator addresses in the body are removed from the existing addresses in the iSCSI security object rather than replace the existing addresses.

	*/
	RemoveInitiatorAddresses *bool

	/* SvmUUID.

	   The unique identifier of an SVM.

	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iscsi credentials modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiCredentialsModifyParams) WithDefaults() *IscsiCredentialsModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iscsi credentials modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiCredentialsModifyParams) SetDefaults() {
	var (
		addInitiatorAddressesDefault = bool(false)

		removeInitiatorAddressesDefault = bool(false)
	)

	val := IscsiCredentialsModifyParams{
		AddInitiatorAddresses:    &addInitiatorAddressesDefault,
		RemoveInitiatorAddresses: &removeInitiatorAddressesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) WithTimeout(timeout time.Duration) *IscsiCredentialsModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) WithContext(ctx context.Context) *IscsiCredentialsModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) WithHTTPClient(client *http.Client) *IscsiCredentialsModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddInitiatorAddresses adds the addInitiatorAddresses to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) WithAddInitiatorAddresses(addInitiatorAddresses *bool) *IscsiCredentialsModifyParams {
	o.SetAddInitiatorAddresses(addInitiatorAddresses)
	return o
}

// SetAddInitiatorAddresses adds the addInitiatorAddresses to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) SetAddInitiatorAddresses(addInitiatorAddresses *bool) {
	o.AddInitiatorAddresses = addInitiatorAddresses
}

// WithInfo adds the info to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) WithInfo(info *models.IscsiCredentials) *IscsiCredentialsModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) SetInfo(info *models.IscsiCredentials) {
	o.Info = info
}

// WithInitiator adds the initiator to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) WithInitiator(initiator string) *IscsiCredentialsModifyParams {
	o.SetInitiator(initiator)
	return o
}

// SetInitiator adds the initiator to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) SetInitiator(initiator string) {
	o.Initiator = initiator
}

// WithRemoveInitiatorAddresses adds the removeInitiatorAddresses to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) WithRemoveInitiatorAddresses(removeInitiatorAddresses *bool) *IscsiCredentialsModifyParams {
	o.SetRemoveInitiatorAddresses(removeInitiatorAddresses)
	return o
}

// SetRemoveInitiatorAddresses adds the removeInitiatorAddresses to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) SetRemoveInitiatorAddresses(removeInitiatorAddresses *bool) {
	o.RemoveInitiatorAddresses = removeInitiatorAddresses
}

// WithSvmUUID adds the svmUUID to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) WithSvmUUID(svmUUID string) *IscsiCredentialsModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the iscsi credentials modify params
func (o *IscsiCredentialsModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *IscsiCredentialsModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddInitiatorAddresses != nil {

		// query param add_initiator_addresses
		var qrAddInitiatorAddresses bool

		if o.AddInitiatorAddresses != nil {
			qrAddInitiatorAddresses = *o.AddInitiatorAddresses
		}
		qAddInitiatorAddresses := swag.FormatBool(qrAddInitiatorAddresses)
		if qAddInitiatorAddresses != "" {

			if err := r.SetQueryParam("add_initiator_addresses", qAddInitiatorAddresses); err != nil {
				return err
			}
		}
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param initiator
	if err := r.SetPathParam("initiator", o.Initiator); err != nil {
		return err
	}

	if o.RemoveInitiatorAddresses != nil {

		// query param remove_initiator_addresses
		var qrRemoveInitiatorAddresses bool

		if o.RemoveInitiatorAddresses != nil {
			qrRemoveInitiatorAddresses = *o.RemoveInitiatorAddresses
		}
		qRemoveInitiatorAddresses := swag.FormatBool(qrRemoveInitiatorAddresses)
		if qRemoveInitiatorAddresses != "" {

			if err := r.SetQueryParam("remove_initiator_addresses", qRemoveInitiatorAddresses); err != nil {
				return err
			}
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
