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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewEmsApplicationLogsCreateParams creates a new EmsApplicationLogsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsApplicationLogsCreateParams() *EmsApplicationLogsCreateParams {
	return &EmsApplicationLogsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsApplicationLogsCreateParamsWithTimeout creates a new EmsApplicationLogsCreateParams object
// with the ability to set a timeout on a request.
func NewEmsApplicationLogsCreateParamsWithTimeout(timeout time.Duration) *EmsApplicationLogsCreateParams {
	return &EmsApplicationLogsCreateParams{
		timeout: timeout,
	}
}

// NewEmsApplicationLogsCreateParamsWithContext creates a new EmsApplicationLogsCreateParams object
// with the ability to set a context for a request.
func NewEmsApplicationLogsCreateParamsWithContext(ctx context.Context) *EmsApplicationLogsCreateParams {
	return &EmsApplicationLogsCreateParams{
		Context: ctx,
	}
}

// NewEmsApplicationLogsCreateParamsWithHTTPClient creates a new EmsApplicationLogsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsApplicationLogsCreateParamsWithHTTPClient(client *http.Client) *EmsApplicationLogsCreateParams {
	return &EmsApplicationLogsCreateParams{
		HTTPClient: client,
	}
}

/*
EmsApplicationLogsCreateParams contains all the parameters to send to the API endpoint

	for the ems application logs create operation.

	Typically these are written to a http.Request.
*/
type EmsApplicationLogsCreateParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.EmsApplicationLog

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems application logs create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsApplicationLogsCreateParams) WithDefaults() *EmsApplicationLogsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems application logs create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsApplicationLogsCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := EmsApplicationLogsCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ems application logs create params
func (o *EmsApplicationLogsCreateParams) WithTimeout(timeout time.Duration) *EmsApplicationLogsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems application logs create params
func (o *EmsApplicationLogsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems application logs create params
func (o *EmsApplicationLogsCreateParams) WithContext(ctx context.Context) *EmsApplicationLogsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems application logs create params
func (o *EmsApplicationLogsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems application logs create params
func (o *EmsApplicationLogsCreateParams) WithHTTPClient(client *http.Client) *EmsApplicationLogsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems application logs create params
func (o *EmsApplicationLogsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ems application logs create params
func (o *EmsApplicationLogsCreateParams) WithInfo(info *models.EmsApplicationLog) *EmsApplicationLogsCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ems application logs create params
func (o *EmsApplicationLogsCreateParams) SetInfo(info *models.EmsApplicationLog) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the ems application logs create params
func (o *EmsApplicationLogsCreateParams) WithReturnRecords(returnRecords *bool) *EmsApplicationLogsCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ems application logs create params
func (o *EmsApplicationLogsCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *EmsApplicationLogsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
