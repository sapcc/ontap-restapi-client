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

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// NewS3PolicyModifyParams creates a new S3PolicyModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3PolicyModifyParams() *S3PolicyModifyParams {
	return &S3PolicyModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3PolicyModifyParamsWithTimeout creates a new S3PolicyModifyParams object
// with the ability to set a timeout on a request.
func NewS3PolicyModifyParamsWithTimeout(timeout time.Duration) *S3PolicyModifyParams {
	return &S3PolicyModifyParams{
		timeout: timeout,
	}
}

// NewS3PolicyModifyParamsWithContext creates a new S3PolicyModifyParams object
// with the ability to set a context for a request.
func NewS3PolicyModifyParamsWithContext(ctx context.Context) *S3PolicyModifyParams {
	return &S3PolicyModifyParams{
		Context: ctx,
	}
}

// NewS3PolicyModifyParamsWithHTTPClient creates a new S3PolicyModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3PolicyModifyParamsWithHTTPClient(client *http.Client) *S3PolicyModifyParams {
	return &S3PolicyModifyParams{
		HTTPClient: client,
	}
}

/*
S3PolicyModifyParams contains all the parameters to send to the API endpoint

	for the s3 policy modify operation.

	Typically these are written to a http.Request.
*/
type S3PolicyModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.S3Policy

	/* Name.

	   Policy name
	*/
	Name string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 policy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3PolicyModifyParams) WithDefaults() *S3PolicyModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 policy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3PolicyModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 policy modify params
func (o *S3PolicyModifyParams) WithTimeout(timeout time.Duration) *S3PolicyModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 policy modify params
func (o *S3PolicyModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 policy modify params
func (o *S3PolicyModifyParams) WithContext(ctx context.Context) *S3PolicyModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 policy modify params
func (o *S3PolicyModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 policy modify params
func (o *S3PolicyModifyParams) WithHTTPClient(client *http.Client) *S3PolicyModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 policy modify params
func (o *S3PolicyModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the s3 policy modify params
func (o *S3PolicyModifyParams) WithInfo(info *models.S3Policy) *S3PolicyModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 policy modify params
func (o *S3PolicyModifyParams) SetInfo(info *models.S3Policy) {
	o.Info = info
}

// WithName adds the name to the s3 policy modify params
func (o *S3PolicyModifyParams) WithName(name string) *S3PolicyModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the s3 policy modify params
func (o *S3PolicyModifyParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the s3 policy modify params
func (o *S3PolicyModifyParams) WithSvmUUID(svmUUID string) *S3PolicyModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 policy modify params
func (o *S3PolicyModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3PolicyModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
