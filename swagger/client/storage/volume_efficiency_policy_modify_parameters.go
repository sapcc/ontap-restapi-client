// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewVolumeEfficiencyPolicyModifyParams creates a new VolumeEfficiencyPolicyModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeEfficiencyPolicyModifyParams() *VolumeEfficiencyPolicyModifyParams {
	return &VolumeEfficiencyPolicyModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeEfficiencyPolicyModifyParamsWithTimeout creates a new VolumeEfficiencyPolicyModifyParams object
// with the ability to set a timeout on a request.
func NewVolumeEfficiencyPolicyModifyParamsWithTimeout(timeout time.Duration) *VolumeEfficiencyPolicyModifyParams {
	return &VolumeEfficiencyPolicyModifyParams{
		timeout: timeout,
	}
}

// NewVolumeEfficiencyPolicyModifyParamsWithContext creates a new VolumeEfficiencyPolicyModifyParams object
// with the ability to set a context for a request.
func NewVolumeEfficiencyPolicyModifyParamsWithContext(ctx context.Context) *VolumeEfficiencyPolicyModifyParams {
	return &VolumeEfficiencyPolicyModifyParams{
		Context: ctx,
	}
}

// NewVolumeEfficiencyPolicyModifyParamsWithHTTPClient creates a new VolumeEfficiencyPolicyModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeEfficiencyPolicyModifyParamsWithHTTPClient(client *http.Client) *VolumeEfficiencyPolicyModifyParams {
	return &VolumeEfficiencyPolicyModifyParams{
		HTTPClient: client,
	}
}

/*
VolumeEfficiencyPolicyModifyParams contains all the parameters to send to the API endpoint

	for the volume efficiency policy modify operation.

	Typically these are written to a http.Request.
*/
type VolumeEfficiencyPolicyModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.VolumeEfficiencyPolicy

	/* UUID.

	   Volume efficiency policy UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume efficiency policy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeEfficiencyPolicyModifyParams) WithDefaults() *VolumeEfficiencyPolicyModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume efficiency policy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeEfficiencyPolicyModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume efficiency policy modify params
func (o *VolumeEfficiencyPolicyModifyParams) WithTimeout(timeout time.Duration) *VolumeEfficiencyPolicyModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume efficiency policy modify params
func (o *VolumeEfficiencyPolicyModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume efficiency policy modify params
func (o *VolumeEfficiencyPolicyModifyParams) WithContext(ctx context.Context) *VolumeEfficiencyPolicyModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume efficiency policy modify params
func (o *VolumeEfficiencyPolicyModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume efficiency policy modify params
func (o *VolumeEfficiencyPolicyModifyParams) WithHTTPClient(client *http.Client) *VolumeEfficiencyPolicyModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume efficiency policy modify params
func (o *VolumeEfficiencyPolicyModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the volume efficiency policy modify params
func (o *VolumeEfficiencyPolicyModifyParams) WithInfo(info *models.VolumeEfficiencyPolicy) *VolumeEfficiencyPolicyModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the volume efficiency policy modify params
func (o *VolumeEfficiencyPolicyModifyParams) SetInfo(info *models.VolumeEfficiencyPolicy) {
	o.Info = info
}

// WithUUID adds the uuid to the volume efficiency policy modify params
func (o *VolumeEfficiencyPolicyModifyParams) WithUUID(uuid string) *VolumeEfficiencyPolicyModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the volume efficiency policy modify params
func (o *VolumeEfficiencyPolicyModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeEfficiencyPolicyModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
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
