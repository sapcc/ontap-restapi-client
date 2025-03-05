// Code generated by go-swagger; DO NOT EDIT.

package snap_lock

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

// NewSnaplockRetentionPolicyCollectionGetParams creates a new SnaplockRetentionPolicyCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockRetentionPolicyCollectionGetParams() *SnaplockRetentionPolicyCollectionGetParams {
	return &SnaplockRetentionPolicyCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockRetentionPolicyCollectionGetParamsWithTimeout creates a new SnaplockRetentionPolicyCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSnaplockRetentionPolicyCollectionGetParamsWithTimeout(timeout time.Duration) *SnaplockRetentionPolicyCollectionGetParams {
	return &SnaplockRetentionPolicyCollectionGetParams{
		timeout: timeout,
	}
}

// NewSnaplockRetentionPolicyCollectionGetParamsWithContext creates a new SnaplockRetentionPolicyCollectionGetParams object
// with the ability to set a context for a request.
func NewSnaplockRetentionPolicyCollectionGetParamsWithContext(ctx context.Context) *SnaplockRetentionPolicyCollectionGetParams {
	return &SnaplockRetentionPolicyCollectionGetParams{
		Context: ctx,
	}
}

// NewSnaplockRetentionPolicyCollectionGetParamsWithHTTPClient creates a new SnaplockRetentionPolicyCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockRetentionPolicyCollectionGetParamsWithHTTPClient(client *http.Client) *SnaplockRetentionPolicyCollectionGetParams {
	return &SnaplockRetentionPolicyCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SnaplockRetentionPolicyCollectionGetParams contains all the parameters to send to the API endpoint

	for the snaplock retention policy collection get operation.

	Typically these are written to a http.Request.
*/
type SnaplockRetentionPolicyCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock retention policy collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockRetentionPolicyCollectionGetParams) WithDefaults() *SnaplockRetentionPolicyCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock retention policy collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockRetentionPolicyCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SnaplockRetentionPolicyCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) WithTimeout(timeout time.Duration) *SnaplockRetentionPolicyCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) WithContext(ctx context.Context) *SnaplockRetentionPolicyCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) WithHTTPClient(client *http.Client) *SnaplockRetentionPolicyCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) WithFields(fields []string) *SnaplockRetentionPolicyCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) WithMaxRecords(maxRecords *int64) *SnaplockRetentionPolicyCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) WithOrderBy(orderBy []string) *SnaplockRetentionPolicyCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) WithReturnRecords(returnRecords *bool) *SnaplockRetentionPolicyCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SnaplockRetentionPolicyCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snaplock retention policy collection get params
func (o *SnaplockRetentionPolicyCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockRetentionPolicyCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
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

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnaplockRetentionPolicyCollectionGet binds the parameter fields
func (o *SnaplockRetentionPolicyCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSnaplockRetentionPolicyCollectionGet binds the parameter order_by
func (o *SnaplockRetentionPolicyCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
