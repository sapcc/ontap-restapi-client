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

// NewS3BucketModifyParams creates a new S3BucketModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3BucketModifyParams() *S3BucketModifyParams {
	return &S3BucketModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3BucketModifyParamsWithTimeout creates a new S3BucketModifyParams object
// with the ability to set a timeout on a request.
func NewS3BucketModifyParamsWithTimeout(timeout time.Duration) *S3BucketModifyParams {
	return &S3BucketModifyParams{
		timeout: timeout,
	}
}

// NewS3BucketModifyParamsWithContext creates a new S3BucketModifyParams object
// with the ability to set a context for a request.
func NewS3BucketModifyParamsWithContext(ctx context.Context) *S3BucketModifyParams {
	return &S3BucketModifyParams{
		Context: ctx,
	}
}

// NewS3BucketModifyParamsWithHTTPClient creates a new S3BucketModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3BucketModifyParamsWithHTTPClient(client *http.Client) *S3BucketModifyParams {
	return &S3BucketModifyParams{
		HTTPClient: client,
	}
}

/*
S3BucketModifyParams contains all the parameters to send to the API endpoint

	for the s3 bucket modify operation.

	Typically these are written to a http.Request.
*/
type S3BucketModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.S3Bucket

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	/* UUID.

	   The unique identifier of the bucket.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 bucket modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketModifyParams) WithDefaults() *S3BucketModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 bucket modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := S3BucketModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 bucket modify params
func (o *S3BucketModifyParams) WithTimeout(timeout time.Duration) *S3BucketModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 bucket modify params
func (o *S3BucketModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 bucket modify params
func (o *S3BucketModifyParams) WithContext(ctx context.Context) *S3BucketModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 bucket modify params
func (o *S3BucketModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 bucket modify params
func (o *S3BucketModifyParams) WithHTTPClient(client *http.Client) *S3BucketModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 bucket modify params
func (o *S3BucketModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the s3 bucket modify params
func (o *S3BucketModifyParams) WithInfo(info *models.S3Bucket) *S3BucketModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 bucket modify params
func (o *S3BucketModifyParams) SetInfo(info *models.S3Bucket) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the s3 bucket modify params
func (o *S3BucketModifyParams) WithReturnTimeout(returnTimeout *int64) *S3BucketModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the s3 bucket modify params
func (o *S3BucketModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmUUID adds the svmUUID to the s3 bucket modify params
func (o *S3BucketModifyParams) WithSvmUUID(svmUUID string) *S3BucketModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 bucket modify params
func (o *S3BucketModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the s3 bucket modify params
func (o *S3BucketModifyParams) WithUUID(uuid string) *S3BucketModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the s3 bucket modify params
func (o *S3BucketModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *S3BucketModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
