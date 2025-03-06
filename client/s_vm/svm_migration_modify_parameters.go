// Code generated by go-swagger; DO NOT EDIT.

package s_vm

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

	"github.com/sapcc/ontap-restapi-client/models"
)

// NewSvmMigrationModifyParams creates a new SvmMigrationModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmMigrationModifyParams() *SvmMigrationModifyParams {
	return &SvmMigrationModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmMigrationModifyParamsWithTimeout creates a new SvmMigrationModifyParams object
// with the ability to set a timeout on a request.
func NewSvmMigrationModifyParamsWithTimeout(timeout time.Duration) *SvmMigrationModifyParams {
	return &SvmMigrationModifyParams{
		timeout: timeout,
	}
}

// NewSvmMigrationModifyParamsWithContext creates a new SvmMigrationModifyParams object
// with the ability to set a context for a request.
func NewSvmMigrationModifyParamsWithContext(ctx context.Context) *SvmMigrationModifyParams {
	return &SvmMigrationModifyParams{
		Context: ctx,
	}
}

// NewSvmMigrationModifyParamsWithHTTPClient creates a new SvmMigrationModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmMigrationModifyParamsWithHTTPClient(client *http.Client) *SvmMigrationModifyParams {
	return &SvmMigrationModifyParams{
		HTTPClient: client,
	}
}

/*
SvmMigrationModifyParams contains all the parameters to send to the API endpoint

	for the svm migration modify operation.

	Typically these are written to a http.Request.
*/
type SvmMigrationModifyParams struct {

	/* Action.

	     The pause action pauses the SVM migration. This action stops data transfer and configuration replication. This operation must be performed on the destination cluster.
	The resume action resumes an SVM migration from a paused or failed state. If the SVM migration is in a cleanup_failed state, volume placement is ignored. This operation must be performed on the destination cluster.
	The cutover action triggers the cutover of an SVM from the source cluster to the destination cluster.
	The source_clean up action performs a clean up of the SVM on the source cluster.

	*/
	Action *string

	/* AutoCutover.

	   Optional property that when set to true automatically performs cutover when the migration state reaches "ready for cutover".

	   Default: true
	*/
	AutoCutover *bool

	/* AutoSourceCleanup.

	   Optional property that when set to true automatically cleans up the SVM on the source cluster after the migration cutover.

	   Default: true
	*/
	AutoSourceCleanup *bool

	/* Info.

	   Info specification
	*/
	Info *models.SvmMigration

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   SVM migration UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm migration modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmMigrationModifyParams) WithDefaults() *SvmMigrationModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm migration modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmMigrationModifyParams) SetDefaults() {
	var (
		autoCutoverDefault = bool(true)

		autoSourceCleanupDefault = bool(true)

		returnTimeoutDefault = int64(0)
	)

	val := SvmMigrationModifyParams{
		AutoCutover:       &autoCutoverDefault,
		AutoSourceCleanup: &autoSourceCleanupDefault,
		ReturnTimeout:     &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm migration modify params
func (o *SvmMigrationModifyParams) WithTimeout(timeout time.Duration) *SvmMigrationModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm migration modify params
func (o *SvmMigrationModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm migration modify params
func (o *SvmMigrationModifyParams) WithContext(ctx context.Context) *SvmMigrationModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm migration modify params
func (o *SvmMigrationModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm migration modify params
func (o *SvmMigrationModifyParams) WithHTTPClient(client *http.Client) *SvmMigrationModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm migration modify params
func (o *SvmMigrationModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the svm migration modify params
func (o *SvmMigrationModifyParams) WithAction(action *string) *SvmMigrationModifyParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the svm migration modify params
func (o *SvmMigrationModifyParams) SetAction(action *string) {
	o.Action = action
}

// WithAutoCutover adds the autoCutover to the svm migration modify params
func (o *SvmMigrationModifyParams) WithAutoCutover(autoCutover *bool) *SvmMigrationModifyParams {
	o.SetAutoCutover(autoCutover)
	return o
}

// SetAutoCutover adds the autoCutover to the svm migration modify params
func (o *SvmMigrationModifyParams) SetAutoCutover(autoCutover *bool) {
	o.AutoCutover = autoCutover
}

// WithAutoSourceCleanup adds the autoSourceCleanup to the svm migration modify params
func (o *SvmMigrationModifyParams) WithAutoSourceCleanup(autoSourceCleanup *bool) *SvmMigrationModifyParams {
	o.SetAutoSourceCleanup(autoSourceCleanup)
	return o
}

// SetAutoSourceCleanup adds the autoSourceCleanup to the svm migration modify params
func (o *SvmMigrationModifyParams) SetAutoSourceCleanup(autoSourceCleanup *bool) {
	o.AutoSourceCleanup = autoSourceCleanup
}

// WithInfo adds the info to the svm migration modify params
func (o *SvmMigrationModifyParams) WithInfo(info *models.SvmMigration) *SvmMigrationModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the svm migration modify params
func (o *SvmMigrationModifyParams) SetInfo(info *models.SvmMigration) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the svm migration modify params
func (o *SvmMigrationModifyParams) WithReturnTimeout(returnTimeout *int64) *SvmMigrationModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the svm migration modify params
func (o *SvmMigrationModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the svm migration modify params
func (o *SvmMigrationModifyParams) WithUUID(uuid string) *SvmMigrationModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the svm migration modify params
func (o *SvmMigrationModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SvmMigrationModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// query param action
		var qrAction string

		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {

			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}
	}

	if o.AutoCutover != nil {

		// query param auto_cutover
		var qrAutoCutover bool

		if o.AutoCutover != nil {
			qrAutoCutover = *o.AutoCutover
		}
		qAutoCutover := swag.FormatBool(qrAutoCutover)
		if qAutoCutover != "" {

			if err := r.SetQueryParam("auto_cutover", qAutoCutover); err != nil {
				return err
			}
		}
	}

	if o.AutoSourceCleanup != nil {

		// query param auto_source_cleanup
		var qrAutoSourceCleanup bool

		if o.AutoSourceCleanup != nil {
			qrAutoSourceCleanup = *o.AutoSourceCleanup
		}
		qAutoSourceCleanup := swag.FormatBool(qrAutoSourceCleanup)
		if qAutoSourceCleanup != "" {

			if err := r.SetQueryParam("auto_source_cleanup", qAutoSourceCleanup); err != nil {
				return err
			}
		}
	}
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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
