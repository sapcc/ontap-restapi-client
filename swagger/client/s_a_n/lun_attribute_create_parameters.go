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

// NewLunAttributeCreateParams creates a new LunAttributeCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunAttributeCreateParams() *LunAttributeCreateParams {
	return &LunAttributeCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunAttributeCreateParamsWithTimeout creates a new LunAttributeCreateParams object
// with the ability to set a timeout on a request.
func NewLunAttributeCreateParamsWithTimeout(timeout time.Duration) *LunAttributeCreateParams {
	return &LunAttributeCreateParams{
		timeout: timeout,
	}
}

// NewLunAttributeCreateParamsWithContext creates a new LunAttributeCreateParams object
// with the ability to set a context for a request.
func NewLunAttributeCreateParamsWithContext(ctx context.Context) *LunAttributeCreateParams {
	return &LunAttributeCreateParams{
		Context: ctx,
	}
}

// NewLunAttributeCreateParamsWithHTTPClient creates a new LunAttributeCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunAttributeCreateParamsWithHTTPClient(client *http.Client) *LunAttributeCreateParams {
	return &LunAttributeCreateParams{
		HTTPClient: client,
	}
}

/*
LunAttributeCreateParams contains all the parameters to send to the API endpoint

	for the lun attribute create operation.

	Typically these are written to a http.Request.
*/
type LunAttributeCreateParams struct {

	/* Info.

	   The attribute name and value to add to the LUN.

	*/
	Info *models.LunAttribute

	/* LunUUID.

	   The unique identifier of the LUN.

	*/
	LunUUID string

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lun attribute create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunAttributeCreateParams) WithDefaults() *LunAttributeCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun attribute create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunAttributeCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := LunAttributeCreateParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the lun attribute create params
func (o *LunAttributeCreateParams) WithTimeout(timeout time.Duration) *LunAttributeCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun attribute create params
func (o *LunAttributeCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun attribute create params
func (o *LunAttributeCreateParams) WithContext(ctx context.Context) *LunAttributeCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun attribute create params
func (o *LunAttributeCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun attribute create params
func (o *LunAttributeCreateParams) WithHTTPClient(client *http.Client) *LunAttributeCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun attribute create params
func (o *LunAttributeCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the lun attribute create params
func (o *LunAttributeCreateParams) WithInfo(info *models.LunAttribute) *LunAttributeCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the lun attribute create params
func (o *LunAttributeCreateParams) SetInfo(info *models.LunAttribute) {
	o.Info = info
}

// WithLunUUID adds the lunUUID to the lun attribute create params
func (o *LunAttributeCreateParams) WithLunUUID(lunUUID string) *LunAttributeCreateParams {
	o.SetLunUUID(lunUUID)
	return o
}

// SetLunUUID adds the lunUuid to the lun attribute create params
func (o *LunAttributeCreateParams) SetLunUUID(lunUUID string) {
	o.LunUUID = lunUUID
}

// WithReturnRecords adds the returnRecords to the lun attribute create params
func (o *LunAttributeCreateParams) WithReturnRecords(returnRecords *bool) *LunAttributeCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the lun attribute create params
func (o *LunAttributeCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the lun attribute create params
func (o *LunAttributeCreateParams) WithReturnTimeout(returnTimeout *int64) *LunAttributeCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the lun attribute create params
func (o *LunAttributeCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *LunAttributeCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param lun.uuid
	if err := r.SetPathParam("lun.uuid", o.LunUUID); err != nil {
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
