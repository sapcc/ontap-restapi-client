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

// NewSnaplockRetentionOperationCollectionGetParams creates a new SnaplockRetentionOperationCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockRetentionOperationCollectionGetParams() *SnaplockRetentionOperationCollectionGetParams {
	return &SnaplockRetentionOperationCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockRetentionOperationCollectionGetParamsWithTimeout creates a new SnaplockRetentionOperationCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSnaplockRetentionOperationCollectionGetParamsWithTimeout(timeout time.Duration) *SnaplockRetentionOperationCollectionGetParams {
	return &SnaplockRetentionOperationCollectionGetParams{
		timeout: timeout,
	}
}

// NewSnaplockRetentionOperationCollectionGetParamsWithContext creates a new SnaplockRetentionOperationCollectionGetParams object
// with the ability to set a context for a request.
func NewSnaplockRetentionOperationCollectionGetParamsWithContext(ctx context.Context) *SnaplockRetentionOperationCollectionGetParams {
	return &SnaplockRetentionOperationCollectionGetParams{
		Context: ctx,
	}
}

// NewSnaplockRetentionOperationCollectionGetParamsWithHTTPClient creates a new SnaplockRetentionOperationCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockRetentionOperationCollectionGetParamsWithHTTPClient(client *http.Client) *SnaplockRetentionOperationCollectionGetParams {
	return &SnaplockRetentionOperationCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SnaplockRetentionOperationCollectionGetParams contains all the parameters to send to the API endpoint

	for the snaplock retention operation collection get operation.

	Typically these are written to a http.Request.
*/
type SnaplockRetentionOperationCollectionGetParams struct {

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

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock retention operation collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockRetentionOperationCollectionGetParams) WithDefaults() *SnaplockRetentionOperationCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock retention operation collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockRetentionOperationCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SnaplockRetentionOperationCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) WithTimeout(timeout time.Duration) *SnaplockRetentionOperationCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) WithContext(ctx context.Context) *SnaplockRetentionOperationCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) WithHTTPClient(client *http.Client) *SnaplockRetentionOperationCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) WithFields(fields []string) *SnaplockRetentionOperationCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) WithMaxRecords(maxRecords *int64) *SnaplockRetentionOperationCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) WithOrderBy(orderBy []string) *SnaplockRetentionOperationCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) WithReturnRecords(returnRecords *bool) *SnaplockRetentionOperationCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SnaplockRetentionOperationCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithVolumeUUID adds the volumeUUID to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) WithVolumeUUID(volumeUUID *string) *SnaplockRetentionOperationCollectionGetParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the snaplock retention operation collection get params
func (o *SnaplockRetentionOperationCollectionGetParams) SetVolumeUUID(volumeUUID *string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockRetentionOperationCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.VolumeUUID != nil {

		// query param volume.uuid
		var qrVolumeUUID string

		if o.VolumeUUID != nil {
			qrVolumeUUID = *o.VolumeUUID
		}
		qVolumeUUID := qrVolumeUUID
		if qVolumeUUID != "" {

			if err := r.SetQueryParam("volume.uuid", qVolumeUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnaplockRetentionOperationCollectionGet binds the parameter fields
func (o *SnaplockRetentionOperationCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSnaplockRetentionOperationCollectionGet binds the parameter order_by
func (o *SnaplockRetentionOperationCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
