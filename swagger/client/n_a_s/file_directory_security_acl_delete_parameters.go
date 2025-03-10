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

// NewFileDirectorySecurityACLDeleteParams creates a new FileDirectorySecurityACLDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileDirectorySecurityACLDeleteParams() *FileDirectorySecurityACLDeleteParams {
	return &FileDirectorySecurityACLDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileDirectorySecurityACLDeleteParamsWithTimeout creates a new FileDirectorySecurityACLDeleteParams object
// with the ability to set a timeout on a request.
func NewFileDirectorySecurityACLDeleteParamsWithTimeout(timeout time.Duration) *FileDirectorySecurityACLDeleteParams {
	return &FileDirectorySecurityACLDeleteParams{
		timeout: timeout,
	}
}

// NewFileDirectorySecurityACLDeleteParamsWithContext creates a new FileDirectorySecurityACLDeleteParams object
// with the ability to set a context for a request.
func NewFileDirectorySecurityACLDeleteParamsWithContext(ctx context.Context) *FileDirectorySecurityACLDeleteParams {
	return &FileDirectorySecurityACLDeleteParams{
		Context: ctx,
	}
}

// NewFileDirectorySecurityACLDeleteParamsWithHTTPClient creates a new FileDirectorySecurityACLDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileDirectorySecurityACLDeleteParamsWithHTTPClient(client *http.Client) *FileDirectorySecurityACLDeleteParams {
	return &FileDirectorySecurityACLDeleteParams{
		HTTPClient: client,
	}
}

/*
FileDirectorySecurityACLDeleteParams contains all the parameters to send to the API endpoint

	for the file directory security acl delete operation.

	Typically these are written to a http.Request.
*/
type FileDirectorySecurityACLDeleteParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.ACLDelete

	/* Path.

	   path
	*/
	Path string

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

	/* User.

	   User Name
	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the file directory security acl delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileDirectorySecurityACLDeleteParams) WithDefaults() *FileDirectorySecurityACLDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file directory security acl delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileDirectorySecurityACLDeleteParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := FileDirectorySecurityACLDeleteParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) WithTimeout(timeout time.Duration) *FileDirectorySecurityACLDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) WithContext(ctx context.Context) *FileDirectorySecurityACLDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) WithHTTPClient(client *http.Client) *FileDirectorySecurityACLDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) WithInfo(info *models.ACLDelete) *FileDirectorySecurityACLDeleteParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) SetInfo(info *models.ACLDelete) {
	o.Info = info
}

// WithPath adds the path to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) WithPath(path string) *FileDirectorySecurityACLDeleteParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) SetPath(path string) {
	o.Path = path
}

// WithReturnRecords adds the returnRecords to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) WithReturnRecords(returnRecords *bool) *FileDirectorySecurityACLDeleteParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) WithReturnTimeout(returnTimeout *int64) *FileDirectorySecurityACLDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmUUID adds the svmUUID to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) WithSvmUUID(svmUUID string) *FileDirectorySecurityACLDeleteParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithUser adds the user to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) WithUser(user string) *FileDirectorySecurityACLDeleteParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the file directory security acl delete params
func (o *FileDirectorySecurityACLDeleteParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *FileDirectorySecurityACLDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
