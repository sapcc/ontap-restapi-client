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

// NewS3BucketSvmCreateParams creates a new S3BucketSvmCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3BucketSvmCreateParams() *S3BucketSvmCreateParams {
	return &S3BucketSvmCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3BucketSvmCreateParamsWithTimeout creates a new S3BucketSvmCreateParams object
// with the ability to set a timeout on a request.
func NewS3BucketSvmCreateParamsWithTimeout(timeout time.Duration) *S3BucketSvmCreateParams {
	return &S3BucketSvmCreateParams{
		timeout: timeout,
	}
}

// NewS3BucketSvmCreateParamsWithContext creates a new S3BucketSvmCreateParams object
// with the ability to set a context for a request.
func NewS3BucketSvmCreateParamsWithContext(ctx context.Context) *S3BucketSvmCreateParams {
	return &S3BucketSvmCreateParams{
		Context: ctx,
	}
}

// NewS3BucketSvmCreateParamsWithHTTPClient creates a new S3BucketSvmCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3BucketSvmCreateParamsWithHTTPClient(client *http.Client) *S3BucketSvmCreateParams {
	return &S3BucketSvmCreateParams{
		HTTPClient: client,
	}
}

/*
S3BucketSvmCreateParams contains all the parameters to send to the API endpoint

	for the s3 bucket svm create operation.

	Typically these are written to a http.Request.
*/
type S3BucketSvmCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.S3BucketSvm

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 bucket svm create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketSvmCreateParams) WithDefaults() *S3BucketSvmCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 bucket svm create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketSvmCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := S3BucketSvmCreateParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) WithTimeout(timeout time.Duration) *S3BucketSvmCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) WithContext(ctx context.Context) *S3BucketSvmCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) WithHTTPClient(client *http.Client) *S3BucketSvmCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) WithInfo(info *models.S3BucketSvm) *S3BucketSvmCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) SetInfo(info *models.S3BucketSvm) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) WithReturnRecords(returnRecords *bool) *S3BucketSvmCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) WithReturnTimeout(returnTimeout *int64) *S3BucketSvmCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmUUID adds the svmUUID to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) WithSvmUUID(svmUUID string) *S3BucketSvmCreateParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 bucket svm create params
func (o *S3BucketSvmCreateParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3BucketSvmCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
