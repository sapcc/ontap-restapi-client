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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewAzureKeyVaultRekeyExternalParams creates a new AzureKeyVaultRekeyExternalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAzureKeyVaultRekeyExternalParams() *AzureKeyVaultRekeyExternalParams {
	return &AzureKeyVaultRekeyExternalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAzureKeyVaultRekeyExternalParamsWithTimeout creates a new AzureKeyVaultRekeyExternalParams object
// with the ability to set a timeout on a request.
func NewAzureKeyVaultRekeyExternalParamsWithTimeout(timeout time.Duration) *AzureKeyVaultRekeyExternalParams {
	return &AzureKeyVaultRekeyExternalParams{
		timeout: timeout,
	}
}

// NewAzureKeyVaultRekeyExternalParamsWithContext creates a new AzureKeyVaultRekeyExternalParams object
// with the ability to set a context for a request.
func NewAzureKeyVaultRekeyExternalParamsWithContext(ctx context.Context) *AzureKeyVaultRekeyExternalParams {
	return &AzureKeyVaultRekeyExternalParams{
		Context: ctx,
	}
}

// NewAzureKeyVaultRekeyExternalParamsWithHTTPClient creates a new AzureKeyVaultRekeyExternalParams object
// with the ability to set a custom HTTPClient for a request.
func NewAzureKeyVaultRekeyExternalParamsWithHTTPClient(client *http.Client) *AzureKeyVaultRekeyExternalParams {
	return &AzureKeyVaultRekeyExternalParams{
		HTTPClient: client,
	}
}

/*
AzureKeyVaultRekeyExternalParams contains all the parameters to send to the API endpoint

	for the azure key vault rekey external operation.

	Typically these are written to a http.Request.
*/
type AzureKeyVaultRekeyExternalParams struct {

	/* AzureKeyVaultUUID.

	   UUID of the existing AKV configuration.
	*/
	AzureKeyVaultUUID string

	/* Info.

	   AKV information.
	*/
	Info *models.AzureKeyVaultKey

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

// WithDefaults hydrates default values in the azure key vault rekey external params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureKeyVaultRekeyExternalParams) WithDefaults() *AzureKeyVaultRekeyExternalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the azure key vault rekey external params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureKeyVaultRekeyExternalParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := AzureKeyVaultRekeyExternalParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) WithTimeout(timeout time.Duration) *AzureKeyVaultRekeyExternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) WithContext(ctx context.Context) *AzureKeyVaultRekeyExternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) WithHTTPClient(client *http.Client) *AzureKeyVaultRekeyExternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAzureKeyVaultUUID adds the azureKeyVaultUUID to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) WithAzureKeyVaultUUID(azureKeyVaultUUID string) *AzureKeyVaultRekeyExternalParams {
	o.SetAzureKeyVaultUUID(azureKeyVaultUUID)
	return o
}

// SetAzureKeyVaultUUID adds the azureKeyVaultUuid to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) SetAzureKeyVaultUUID(azureKeyVaultUUID string) {
	o.AzureKeyVaultUUID = azureKeyVaultUUID
}

// WithInfo adds the info to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) WithInfo(info *models.AzureKeyVaultKey) *AzureKeyVaultRekeyExternalParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) SetInfo(info *models.AzureKeyVaultKey) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) WithReturnRecords(returnRecords *bool) *AzureKeyVaultRekeyExternalParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) WithReturnTimeout(returnTimeout *int64) *AzureKeyVaultRekeyExternalParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the azure key vault rekey external params
func (o *AzureKeyVaultRekeyExternalParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *AzureKeyVaultRekeyExternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param azure_key_vault.uuid
	if err := r.SetPathParam("azure_key_vault.uuid", o.AzureKeyVaultUUID); err != nil {
		return err
	}
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
