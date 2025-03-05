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
	"github.com/go-openapi/swag"
)

// NewCifsShareGetParams creates a new CifsShareGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsShareGetParams() *CifsShareGetParams {
	return &CifsShareGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsShareGetParamsWithTimeout creates a new CifsShareGetParams object
// with the ability to set a timeout on a request.
func NewCifsShareGetParamsWithTimeout(timeout time.Duration) *CifsShareGetParams {
	return &CifsShareGetParams{
		timeout: timeout,
	}
}

// NewCifsShareGetParamsWithContext creates a new CifsShareGetParams object
// with the ability to set a context for a request.
func NewCifsShareGetParamsWithContext(ctx context.Context) *CifsShareGetParams {
	return &CifsShareGetParams{
		Context: ctx,
	}
}

// NewCifsShareGetParamsWithHTTPClient creates a new CifsShareGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsShareGetParamsWithHTTPClient(client *http.Client) *CifsShareGetParams {
	return &CifsShareGetParams{
		HTTPClient: client,
	}
}

/*
CifsShareGetParams contains all the parameters to send to the API endpoint

	for the cifs share get operation.

	Typically these are written to a http.Request.
*/
type CifsShareGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Name.

	   Share Name
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

// WithDefaults hydrates default values in the cifs share get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsShareGetParams) WithDefaults() *CifsShareGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs share get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsShareGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cifs share get params
func (o *CifsShareGetParams) WithTimeout(timeout time.Duration) *CifsShareGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs share get params
func (o *CifsShareGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs share get params
func (o *CifsShareGetParams) WithContext(ctx context.Context) *CifsShareGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs share get params
func (o *CifsShareGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs share get params
func (o *CifsShareGetParams) WithHTTPClient(client *http.Client) *CifsShareGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs share get params
func (o *CifsShareGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the cifs share get params
func (o *CifsShareGetParams) WithFields(fields []string) *CifsShareGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cifs share get params
func (o *CifsShareGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithName adds the name to the cifs share get params
func (o *CifsShareGetParams) WithName(name string) *CifsShareGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the cifs share get params
func (o *CifsShareGetParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the cifs share get params
func (o *CifsShareGetParams) WithSvmUUID(svmUUID string) *CifsShareGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the cifs share get params
func (o *CifsShareGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsShareGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamCifsShareGet binds the parameter fields
func (o *CifsShareGetParams) bindParamFields(formats strfmt.Registry) []string {
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
