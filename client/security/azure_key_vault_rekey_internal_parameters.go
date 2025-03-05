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
)

// NewAzureKeyVaultRekeyInternalParams creates a new AzureKeyVaultRekeyInternalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAzureKeyVaultRekeyInternalParams() *AzureKeyVaultRekeyInternalParams {
	return &AzureKeyVaultRekeyInternalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAzureKeyVaultRekeyInternalParamsWithTimeout creates a new AzureKeyVaultRekeyInternalParams object
// with the ability to set a timeout on a request.
func NewAzureKeyVaultRekeyInternalParamsWithTimeout(timeout time.Duration) *AzureKeyVaultRekeyInternalParams {
	return &AzureKeyVaultRekeyInternalParams{
		timeout: timeout,
	}
}

// NewAzureKeyVaultRekeyInternalParamsWithContext creates a new AzureKeyVaultRekeyInternalParams object
// with the ability to set a context for a request.
func NewAzureKeyVaultRekeyInternalParamsWithContext(ctx context.Context) *AzureKeyVaultRekeyInternalParams {
	return &AzureKeyVaultRekeyInternalParams{
		Context: ctx,
	}
}

// NewAzureKeyVaultRekeyInternalParamsWithHTTPClient creates a new AzureKeyVaultRekeyInternalParams object
// with the ability to set a custom HTTPClient for a request.
func NewAzureKeyVaultRekeyInternalParamsWithHTTPClient(client *http.Client) *AzureKeyVaultRekeyInternalParams {
	return &AzureKeyVaultRekeyInternalParams{
		HTTPClient: client,
	}
}

/*
AzureKeyVaultRekeyInternalParams contains all the parameters to send to the API endpoint

	for the azure key vault rekey internal operation.

	Typically these are written to a http.Request.
*/
type AzureKeyVaultRekeyInternalParams struct {

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   UUID of the existing AKV configuration.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the azure key vault rekey internal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureKeyVaultRekeyInternalParams) WithDefaults() *AzureKeyVaultRekeyInternalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the azure key vault rekey internal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureKeyVaultRekeyInternalParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := AzureKeyVaultRekeyInternalParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) WithTimeout(timeout time.Duration) *AzureKeyVaultRekeyInternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) WithContext(ctx context.Context) *AzureKeyVaultRekeyInternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) WithHTTPClient(client *http.Client) *AzureKeyVaultRekeyInternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnRecords adds the returnRecords to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) WithReturnRecords(returnRecords *bool) *AzureKeyVaultRekeyInternalParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) WithReturnTimeout(returnTimeout *int64) *AzureKeyVaultRekeyInternalParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) WithUUID(uuid string) *AzureKeyVaultRekeyInternalParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the azure key vault rekey internal params
func (o *AzureKeyVaultRekeyInternalParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *AzureKeyVaultRekeyInternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
