// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewFileDirectorySecurityCreateParams creates a new FileDirectorySecurityCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileDirectorySecurityCreateParams() *FileDirectorySecurityCreateParams {
	return &FileDirectorySecurityCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileDirectorySecurityCreateParamsWithTimeout creates a new FileDirectorySecurityCreateParams object
// with the ability to set a timeout on a request.
func NewFileDirectorySecurityCreateParamsWithTimeout(timeout time.Duration) *FileDirectorySecurityCreateParams {
	return &FileDirectorySecurityCreateParams{
		timeout: timeout,
	}
}

// NewFileDirectorySecurityCreateParamsWithContext creates a new FileDirectorySecurityCreateParams object
// with the ability to set a context for a request.
func NewFileDirectorySecurityCreateParamsWithContext(ctx context.Context) *FileDirectorySecurityCreateParams {
	return &FileDirectorySecurityCreateParams{
		Context: ctx,
	}
}

// NewFileDirectorySecurityCreateParamsWithHTTPClient creates a new FileDirectorySecurityCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileDirectorySecurityCreateParamsWithHTTPClient(client *http.Client) *FileDirectorySecurityCreateParams {
	return &FileDirectorySecurityCreateParams{
		HTTPClient: client,
	}
}

/*
FileDirectorySecurityCreateParams contains all the parameters to send to the API endpoint

	for the file directory security create operation.

	Typically these are written to a http.Request.
*/
type FileDirectorySecurityCreateParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.FileDirectorySecurity

	/* Path.

	   target path
	*/
	Path string

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

// WithDefaults hydrates default values in the file directory security create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileDirectorySecurityCreateParams) WithDefaults() *FileDirectorySecurityCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file directory security create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileDirectorySecurityCreateParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := FileDirectorySecurityCreateParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the file directory security create params
func (o *FileDirectorySecurityCreateParams) WithTimeout(timeout time.Duration) *FileDirectorySecurityCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file directory security create params
func (o *FileDirectorySecurityCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file directory security create params
func (o *FileDirectorySecurityCreateParams) WithContext(ctx context.Context) *FileDirectorySecurityCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file directory security create params
func (o *FileDirectorySecurityCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file directory security create params
func (o *FileDirectorySecurityCreateParams) WithHTTPClient(client *http.Client) *FileDirectorySecurityCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file directory security create params
func (o *FileDirectorySecurityCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the file directory security create params
func (o *FileDirectorySecurityCreateParams) WithInfo(info *models.FileDirectorySecurity) *FileDirectorySecurityCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the file directory security create params
func (o *FileDirectorySecurityCreateParams) SetInfo(info *models.FileDirectorySecurity) {
	o.Info = info
}

// WithPath adds the path to the file directory security create params
func (o *FileDirectorySecurityCreateParams) WithPath(path string) *FileDirectorySecurityCreateParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the file directory security create params
func (o *FileDirectorySecurityCreateParams) SetPath(path string) {
	o.Path = path
}

// WithReturnTimeout adds the returnTimeout to the file directory security create params
func (o *FileDirectorySecurityCreateParams) WithReturnTimeout(returnTimeout *int64) *FileDirectorySecurityCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the file directory security create params
func (o *FileDirectorySecurityCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmUUID adds the svmUUID to the file directory security create params
func (o *FileDirectorySecurityCreateParams) WithSvmUUID(svmUUID string) *FileDirectorySecurityCreateParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the file directory security create params
func (o *FileDirectorySecurityCreateParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FileDirectorySecurityCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
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
