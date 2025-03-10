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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewS3GroupCreateParams creates a new S3GroupCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3GroupCreateParams() *S3GroupCreateParams {
	return &S3GroupCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3GroupCreateParamsWithTimeout creates a new S3GroupCreateParams object
// with the ability to set a timeout on a request.
func NewS3GroupCreateParamsWithTimeout(timeout time.Duration) *S3GroupCreateParams {
	return &S3GroupCreateParams{
		timeout: timeout,
	}
}

// NewS3GroupCreateParamsWithContext creates a new S3GroupCreateParams object
// with the ability to set a context for a request.
func NewS3GroupCreateParamsWithContext(ctx context.Context) *S3GroupCreateParams {
	return &S3GroupCreateParams{
		Context: ctx,
	}
}

// NewS3GroupCreateParamsWithHTTPClient creates a new S3GroupCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3GroupCreateParamsWithHTTPClient(client *http.Client) *S3GroupCreateParams {
	return &S3GroupCreateParams{
		HTTPClient: client,
	}
}

/*
S3GroupCreateParams contains all the parameters to send to the API endpoint

	for the s3 group create operation.

	Typically these are written to a http.Request.
*/
type S3GroupCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.S3Group

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3GroupCreateParams) WithDefaults() *S3GroupCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3GroupCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := S3GroupCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 group create params
func (o *S3GroupCreateParams) WithTimeout(timeout time.Duration) *S3GroupCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 group create params
func (o *S3GroupCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 group create params
func (o *S3GroupCreateParams) WithContext(ctx context.Context) *S3GroupCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 group create params
func (o *S3GroupCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 group create params
func (o *S3GroupCreateParams) WithHTTPClient(client *http.Client) *S3GroupCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 group create params
func (o *S3GroupCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the s3 group create params
func (o *S3GroupCreateParams) WithInfo(info *models.S3Group) *S3GroupCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 group create params
func (o *S3GroupCreateParams) SetInfo(info *models.S3Group) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the s3 group create params
func (o *S3GroupCreateParams) WithReturnRecords(returnRecords *bool) *S3GroupCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the s3 group create params
func (o *S3GroupCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithSvmUUID adds the svmUUID to the s3 group create params
func (o *S3GroupCreateParams) WithSvmUUID(svmUUID string) *S3GroupCreateParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 group create params
func (o *S3GroupCreateParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3GroupCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
