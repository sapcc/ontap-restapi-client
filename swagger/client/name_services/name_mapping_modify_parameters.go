// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewNameMappingModifyParams creates a new NameMappingModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNameMappingModifyParams() *NameMappingModifyParams {
	return &NameMappingModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNameMappingModifyParamsWithTimeout creates a new NameMappingModifyParams object
// with the ability to set a timeout on a request.
func NewNameMappingModifyParamsWithTimeout(timeout time.Duration) *NameMappingModifyParams {
	return &NameMappingModifyParams{
		timeout: timeout,
	}
}

// NewNameMappingModifyParamsWithContext creates a new NameMappingModifyParams object
// with the ability to set a context for a request.
func NewNameMappingModifyParamsWithContext(ctx context.Context) *NameMappingModifyParams {
	return &NameMappingModifyParams{
		Context: ctx,
	}
}

// NewNameMappingModifyParamsWithHTTPClient creates a new NameMappingModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewNameMappingModifyParamsWithHTTPClient(client *http.Client) *NameMappingModifyParams {
	return &NameMappingModifyParams{
		HTTPClient: client,
	}
}

/*
NameMappingModifyParams contains all the parameters to send to the API endpoint

	for the name mapping modify operation.

	Typically these are written to a http.Request.
*/
type NameMappingModifyParams struct {

	/* Direction.

	   Direction
	*/
	Direction string

	/* Index.

	   Position of the entry in the list
	*/
	Index int64

	/* Info.

	   Info specification
	*/
	Info *models.NameMapping

	/* NewIndex.

	   New position of the Index after a swap is completed.
	*/
	NewIndex *int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the name mapping modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameMappingModifyParams) WithDefaults() *NameMappingModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the name mapping modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameMappingModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the name mapping modify params
func (o *NameMappingModifyParams) WithTimeout(timeout time.Duration) *NameMappingModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the name mapping modify params
func (o *NameMappingModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the name mapping modify params
func (o *NameMappingModifyParams) WithContext(ctx context.Context) *NameMappingModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the name mapping modify params
func (o *NameMappingModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the name mapping modify params
func (o *NameMappingModifyParams) WithHTTPClient(client *http.Client) *NameMappingModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the name mapping modify params
func (o *NameMappingModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirection adds the direction to the name mapping modify params
func (o *NameMappingModifyParams) WithDirection(direction string) *NameMappingModifyParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the name mapping modify params
func (o *NameMappingModifyParams) SetDirection(direction string) {
	o.Direction = direction
}

// WithIndex adds the index to the name mapping modify params
func (o *NameMappingModifyParams) WithIndex(index int64) *NameMappingModifyParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the name mapping modify params
func (o *NameMappingModifyParams) SetIndex(index int64) {
	o.Index = index
}

// WithInfo adds the info to the name mapping modify params
func (o *NameMappingModifyParams) WithInfo(info *models.NameMapping) *NameMappingModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the name mapping modify params
func (o *NameMappingModifyParams) SetInfo(info *models.NameMapping) {
	o.Info = info
}

// WithNewIndex adds the newIndex to the name mapping modify params
func (o *NameMappingModifyParams) WithNewIndex(newIndex *int64) *NameMappingModifyParams {
	o.SetNewIndex(newIndex)
	return o
}

// SetNewIndex adds the newIndex to the name mapping modify params
func (o *NameMappingModifyParams) SetNewIndex(newIndex *int64) {
	o.NewIndex = newIndex
}

// WithSvmUUID adds the svmUUID to the name mapping modify params
func (o *NameMappingModifyParams) WithSvmUUID(svmUUID string) *NameMappingModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the name mapping modify params
func (o *NameMappingModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NameMappingModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param direction
	if err := r.SetPathParam("direction", o.Direction); err != nil {
		return err
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.NewIndex != nil {

		// query param new_index
		var qrNewIndex int64

		if o.NewIndex != nil {
			qrNewIndex = *o.NewIndex
		}
		qNewIndex := swag.FormatInt64(qrNewIndex)
		if qNewIndex != "" {

			if err := r.SetQueryParam("new_index", qNewIndex); err != nil {
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
