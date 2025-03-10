// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewEmsMessageCollectionGetParams creates a new EmsMessageCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsMessageCollectionGetParams() *EmsMessageCollectionGetParams {
	return &EmsMessageCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsMessageCollectionGetParamsWithTimeout creates a new EmsMessageCollectionGetParams object
// with the ability to set a timeout on a request.
func NewEmsMessageCollectionGetParamsWithTimeout(timeout time.Duration) *EmsMessageCollectionGetParams {
	return &EmsMessageCollectionGetParams{
		timeout: timeout,
	}
}

// NewEmsMessageCollectionGetParamsWithContext creates a new EmsMessageCollectionGetParams object
// with the ability to set a context for a request.
func NewEmsMessageCollectionGetParamsWithContext(ctx context.Context) *EmsMessageCollectionGetParams {
	return &EmsMessageCollectionGetParams{
		Context: ctx,
	}
}

// NewEmsMessageCollectionGetParamsWithHTTPClient creates a new EmsMessageCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsMessageCollectionGetParamsWithHTTPClient(client *http.Client) *EmsMessageCollectionGetParams {
	return &EmsMessageCollectionGetParams{
		HTTPClient: client,
	}
}

/*
EmsMessageCollectionGetParams contains all the parameters to send to the API endpoint

	for the ems message collection get operation.

	Typically these are written to a http.Request.
*/
type EmsMessageCollectionGetParams struct {

	/* CorrectiveAction.

	   Filter by corrective_action
	*/
	CorrectiveAction *string

	/* Deprecated.

	   Filter by deprecated
	*/
	Deprecated *bool

	/* Description.

	   Filter by description
	*/
	Description *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* FilterName.

	   The filter name that applies to the query.
	*/
	FilterName *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	Name *string

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

	/* Severity.

	   Filter by severity
	*/
	Severity *string

	/* SnmpTrapType.

	   Filter by snmp_trap_type
	*/
	SnmpTrapType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems message collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsMessageCollectionGetParams) WithDefaults() *EmsMessageCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems message collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsMessageCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := EmsMessageCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithTimeout(timeout time.Duration) *EmsMessageCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithContext(ctx context.Context) *EmsMessageCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithHTTPClient(client *http.Client) *EmsMessageCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCorrectiveAction adds the correctiveAction to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithCorrectiveAction(correctiveAction *string) *EmsMessageCollectionGetParams {
	o.SetCorrectiveAction(correctiveAction)
	return o
}

// SetCorrectiveAction adds the correctiveAction to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetCorrectiveAction(correctiveAction *string) {
	o.CorrectiveAction = correctiveAction
}

// WithDeprecated adds the deprecated to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithDeprecated(deprecated *bool) *EmsMessageCollectionGetParams {
	o.SetDeprecated(deprecated)
	return o
}

// SetDeprecated adds the deprecated to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetDeprecated(deprecated *bool) {
	o.Deprecated = deprecated
}

// WithDescription adds the description to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithDescription(description *string) *EmsMessageCollectionGetParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetDescription(description *string) {
	o.Description = description
}

// WithFields adds the fields to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithFields(fields []string) *EmsMessageCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFilterName adds the filterName to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithFilterName(filterName *string) *EmsMessageCollectionGetParams {
	o.SetFilterName(filterName)
	return o
}

// SetFilterName adds the filterName to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetFilterName(filterName *string) {
	o.FilterName = filterName
}

// WithMaxRecords adds the maxRecords to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithMaxRecords(maxRecords *int64) *EmsMessageCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithName adds the name to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithName(name *string) *EmsMessageCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithOrderBy(orderBy []string) *EmsMessageCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithReturnRecords(returnRecords *bool) *EmsMessageCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *EmsMessageCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSeverity adds the severity to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithSeverity(severity *string) *EmsMessageCollectionGetParams {
	o.SetSeverity(severity)
	return o
}

// SetSeverity adds the severity to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetSeverity(severity *string) {
	o.Severity = severity
}

// WithSnmpTrapType adds the snmpTrapType to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithSnmpTrapType(snmpTrapType *string) *EmsMessageCollectionGetParams {
	o.SetSnmpTrapType(snmpTrapType)
	return o
}

// SetSnmpTrapType adds the snmpTrapType to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetSnmpTrapType(snmpTrapType *string) {
	o.SnmpTrapType = snmpTrapType
}

// WriteToRequest writes these params to a swagger request
func (o *EmsMessageCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CorrectiveAction != nil {

		// query param corrective_action
		var qrCorrectiveAction string

		if o.CorrectiveAction != nil {
			qrCorrectiveAction = *o.CorrectiveAction
		}
		qCorrectiveAction := qrCorrectiveAction
		if qCorrectiveAction != "" {

			if err := r.SetQueryParam("corrective_action", qCorrectiveAction); err != nil {
				return err
			}
		}
	}

	if o.Deprecated != nil {

		// query param deprecated
		var qrDeprecated bool

		if o.Deprecated != nil {
			qrDeprecated = *o.Deprecated
		}
		qDeprecated := swag.FormatBool(qrDeprecated)
		if qDeprecated != "" {

			if err := r.SetQueryParam("deprecated", qDeprecated); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.FilterName != nil {

		// query param filter.name
		var qrFilterName string

		if o.FilterName != nil {
			qrFilterName = *o.FilterName
		}
		qFilterName := qrFilterName
		if qFilterName != "" {

			if err := r.SetQueryParam("filter.name", qFilterName); err != nil {
				return err
			}
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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.Severity != nil {

		// query param severity
		var qrSeverity string

		if o.Severity != nil {
			qrSeverity = *o.Severity
		}
		qSeverity := qrSeverity
		if qSeverity != "" {

			if err := r.SetQueryParam("severity", qSeverity); err != nil {
				return err
			}
		}
	}

	if o.SnmpTrapType != nil {

		// query param snmp_trap_type
		var qrSnmpTrapType string

		if o.SnmpTrapType != nil {
			qrSnmpTrapType = *o.SnmpTrapType
		}
		qSnmpTrapType := qrSnmpTrapType
		if qSnmpTrapType != "" {

			if err := r.SetQueryParam("snmp_trap_type", qSnmpTrapType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEmsMessageCollectionGet binds the parameter fields
func (o *EmsMessageCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamEmsMessageCollectionGet binds the parameter order_by
func (o *EmsMessageCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
