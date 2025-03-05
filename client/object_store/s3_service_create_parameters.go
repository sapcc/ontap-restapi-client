// Code generated by go-swagger; DO NOT EDIT.

package object_store

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

	"github.com/sapcc/ontap-restapi/models"
)

// NewS3ServiceCreateParams creates a new S3ServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3ServiceCreateParams() *S3ServiceCreateParams {
	return &S3ServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3ServiceCreateParamsWithTimeout creates a new S3ServiceCreateParams object
// with the ability to set a timeout on a request.
func NewS3ServiceCreateParamsWithTimeout(timeout time.Duration) *S3ServiceCreateParams {
	return &S3ServiceCreateParams{
		timeout: timeout,
	}
}

// NewS3ServiceCreateParamsWithContext creates a new S3ServiceCreateParams object
// with the ability to set a context for a request.
func NewS3ServiceCreateParamsWithContext(ctx context.Context) *S3ServiceCreateParams {
	return &S3ServiceCreateParams{
		Context: ctx,
	}
}

// NewS3ServiceCreateParamsWithHTTPClient creates a new S3ServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3ServiceCreateParamsWithHTTPClient(client *http.Client) *S3ServiceCreateParams {
	return &S3ServiceCreateParams{
		HTTPClient: client,
	}
}

/*
S3ServiceCreateParams contains all the parameters to send to the API endpoint

	for the s3 service create operation.

	Typically these are written to a http.Request.
*/
type S3ServiceCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.S3Service

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3ServiceCreateParams) WithDefaults() *S3ServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3ServiceCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := S3ServiceCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 service create params
func (o *S3ServiceCreateParams) WithTimeout(timeout time.Duration) *S3ServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 service create params
func (o *S3ServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 service create params
func (o *S3ServiceCreateParams) WithContext(ctx context.Context) *S3ServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 service create params
func (o *S3ServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 service create params
func (o *S3ServiceCreateParams) WithHTTPClient(client *http.Client) *S3ServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 service create params
func (o *S3ServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the s3 service create params
func (o *S3ServiceCreateParams) WithInfo(info *models.S3Service) *S3ServiceCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 service create params
func (o *S3ServiceCreateParams) SetInfo(info *models.S3Service) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the s3 service create params
func (o *S3ServiceCreateParams) WithReturnRecords(returnRecords *bool) *S3ServiceCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the s3 service create params
func (o *S3ServiceCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *S3ServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
