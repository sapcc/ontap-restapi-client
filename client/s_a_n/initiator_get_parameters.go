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
)

// NewInitiatorGetParams creates a new InitiatorGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitiatorGetParams() *InitiatorGetParams {
	return &InitiatorGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitiatorGetParamsWithTimeout creates a new InitiatorGetParams object
// with the ability to set a timeout on a request.
func NewInitiatorGetParamsWithTimeout(timeout time.Duration) *InitiatorGetParams {
	return &InitiatorGetParams{
		timeout: timeout,
	}
}

// NewInitiatorGetParamsWithContext creates a new InitiatorGetParams object
// with the ability to set a context for a request.
func NewInitiatorGetParamsWithContext(ctx context.Context) *InitiatorGetParams {
	return &InitiatorGetParams{
		Context: ctx,
	}
}

// NewInitiatorGetParamsWithHTTPClient creates a new InitiatorGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitiatorGetParamsWithHTTPClient(client *http.Client) *InitiatorGetParams {
	return &InitiatorGetParams{
		HTTPClient: client,
	}
}

/*
InitiatorGetParams contains all the parameters to send to the API endpoint

	for the initiator get operation.

	Typically these are written to a http.Request.
*/
type InitiatorGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Name.

	   The name of the initiator.

	*/
	Name string

	/* SvmUUID.

	   The unique identifier of the SVM for which the initiator properties are configured.

	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initiator get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiatorGetParams) WithDefaults() *InitiatorGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initiator get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiatorGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initiator get params
func (o *InitiatorGetParams) WithTimeout(timeout time.Duration) *InitiatorGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initiator get params
func (o *InitiatorGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initiator get params
func (o *InitiatorGetParams) WithContext(ctx context.Context) *InitiatorGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initiator get params
func (o *InitiatorGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initiator get params
func (o *InitiatorGetParams) WithHTTPClient(client *http.Client) *InitiatorGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initiator get params
func (o *InitiatorGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the initiator get params
func (o *InitiatorGetParams) WithFields(fields []string) *InitiatorGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the initiator get params
func (o *InitiatorGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithName adds the name to the initiator get params
func (o *InitiatorGetParams) WithName(name string) *InitiatorGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the initiator get params
func (o *InitiatorGetParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the initiator get params
func (o *InitiatorGetParams) WithSvmUUID(svmUUID string) *InitiatorGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the initiator get params
func (o *InitiatorGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *InitiatorGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamInitiatorGet binds the parameter fields
func (o *InitiatorGetParams) bindParamFields(formats strfmt.Registry) []string {
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
