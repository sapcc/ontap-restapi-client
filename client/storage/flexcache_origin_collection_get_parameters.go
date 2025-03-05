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
	"github.com/go-openapi/swag"
)

// NewFlexcacheOriginCollectionGetParams creates a new FlexcacheOriginCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFlexcacheOriginCollectionGetParams() *FlexcacheOriginCollectionGetParams {
	return &FlexcacheOriginCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFlexcacheOriginCollectionGetParamsWithTimeout creates a new FlexcacheOriginCollectionGetParams object
// with the ability to set a timeout on a request.
func NewFlexcacheOriginCollectionGetParamsWithTimeout(timeout time.Duration) *FlexcacheOriginCollectionGetParams {
	return &FlexcacheOriginCollectionGetParams{
		timeout: timeout,
	}
}

// NewFlexcacheOriginCollectionGetParamsWithContext creates a new FlexcacheOriginCollectionGetParams object
// with the ability to set a context for a request.
func NewFlexcacheOriginCollectionGetParamsWithContext(ctx context.Context) *FlexcacheOriginCollectionGetParams {
	return &FlexcacheOriginCollectionGetParams{
		Context: ctx,
	}
}

// NewFlexcacheOriginCollectionGetParamsWithHTTPClient creates a new FlexcacheOriginCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFlexcacheOriginCollectionGetParamsWithHTTPClient(client *http.Client) *FlexcacheOriginCollectionGetParams {
	return &FlexcacheOriginCollectionGetParams{
		HTTPClient: client,
	}
}

/*
FlexcacheOriginCollectionGetParams contains all the parameters to send to the API endpoint

	for the flexcache origin collection get operation.

	Typically these are written to a http.Request.
*/
type FlexcacheOriginCollectionGetParams struct {

	/* BlockLevelInvalidation.

	   Filter by block_level_invalidation
	*/
	BlockLevelInvalidation *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* FlexcachesClusterName.

	   Filter by flexcaches.cluster.name
	*/
	FlexcachesClusterName *string

	/* FlexcachesClusterUUID.

	   Filter by flexcaches.cluster.uuid
	*/
	FlexcachesClusterUUID *string

	/* FlexcachesCreateTime.

	   Filter by flexcaches.create_time
	*/
	FlexcachesCreateTime *string

	/* FlexcachesIPAddress.

	   Filter by flexcaches.ip_address
	*/
	FlexcachesIPAddress *string

	/* FlexcachesSize.

	   Filter by flexcaches.size
	*/
	FlexcachesSize *int64

	/* FlexcachesState.

	   Filter by flexcaches.state
	*/
	FlexcachesState *string

	/* FlexcachesSvmName.

	   Filter by flexcaches.svm.name
	*/
	FlexcachesSvmName *string

	/* FlexcachesSvmUUID.

	   Filter by flexcaches.svm.uuid
	*/
	FlexcachesSvmUUID *string

	/* FlexcachesVolumeName.

	   Filter by flexcaches.volume.name
	*/
	FlexcachesVolumeName *string

	/* FlexcachesVolumeUUID.

	   Filter by flexcaches.volume.uuid
	*/
	FlexcachesVolumeUUID *string

	/* GlobalFileLockingEnabled.

	   Filter by global_file_locking_enabled
	*/
	GlobalFileLockingEnabled *bool

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	Name *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the flexcache origin collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlexcacheOriginCollectionGetParams) WithDefaults() *FlexcacheOriginCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the flexcache origin collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlexcacheOriginCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := FlexcacheOriginCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithTimeout(timeout time.Duration) *FlexcacheOriginCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithContext(ctx context.Context) *FlexcacheOriginCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithHTTPClient(client *http.Client) *FlexcacheOriginCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlockLevelInvalidation adds the blockLevelInvalidation to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithBlockLevelInvalidation(blockLevelInvalidation *bool) *FlexcacheOriginCollectionGetParams {
	o.SetBlockLevelInvalidation(blockLevelInvalidation)
	return o
}

// SetBlockLevelInvalidation adds the blockLevelInvalidation to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetBlockLevelInvalidation(blockLevelInvalidation *bool) {
	o.BlockLevelInvalidation = blockLevelInvalidation
}

// WithFields adds the fields to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFields(fields []string) *FlexcacheOriginCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFlexcachesClusterName adds the flexcachesClusterName to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFlexcachesClusterName(flexcachesClusterName *string) *FlexcacheOriginCollectionGetParams {
	o.SetFlexcachesClusterName(flexcachesClusterName)
	return o
}

// SetFlexcachesClusterName adds the flexcachesClusterName to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFlexcachesClusterName(flexcachesClusterName *string) {
	o.FlexcachesClusterName = flexcachesClusterName
}

// WithFlexcachesClusterUUID adds the flexcachesClusterUUID to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFlexcachesClusterUUID(flexcachesClusterUUID *string) *FlexcacheOriginCollectionGetParams {
	o.SetFlexcachesClusterUUID(flexcachesClusterUUID)
	return o
}

// SetFlexcachesClusterUUID adds the flexcachesClusterUuid to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFlexcachesClusterUUID(flexcachesClusterUUID *string) {
	o.FlexcachesClusterUUID = flexcachesClusterUUID
}

// WithFlexcachesCreateTime adds the flexcachesCreateTime to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFlexcachesCreateTime(flexcachesCreateTime *string) *FlexcacheOriginCollectionGetParams {
	o.SetFlexcachesCreateTime(flexcachesCreateTime)
	return o
}

// SetFlexcachesCreateTime adds the flexcachesCreateTime to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFlexcachesCreateTime(flexcachesCreateTime *string) {
	o.FlexcachesCreateTime = flexcachesCreateTime
}

// WithFlexcachesIPAddress adds the flexcachesIPAddress to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFlexcachesIPAddress(flexcachesIPAddress *string) *FlexcacheOriginCollectionGetParams {
	o.SetFlexcachesIPAddress(flexcachesIPAddress)
	return o
}

// SetFlexcachesIPAddress adds the flexcachesIpAddress to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFlexcachesIPAddress(flexcachesIPAddress *string) {
	o.FlexcachesIPAddress = flexcachesIPAddress
}

// WithFlexcachesSize adds the flexcachesSize to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFlexcachesSize(flexcachesSize *int64) *FlexcacheOriginCollectionGetParams {
	o.SetFlexcachesSize(flexcachesSize)
	return o
}

// SetFlexcachesSize adds the flexcachesSize to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFlexcachesSize(flexcachesSize *int64) {
	o.FlexcachesSize = flexcachesSize
}

// WithFlexcachesState adds the flexcachesState to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFlexcachesState(flexcachesState *string) *FlexcacheOriginCollectionGetParams {
	o.SetFlexcachesState(flexcachesState)
	return o
}

// SetFlexcachesState adds the flexcachesState to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFlexcachesState(flexcachesState *string) {
	o.FlexcachesState = flexcachesState
}

// WithFlexcachesSvmName adds the flexcachesSvmName to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFlexcachesSvmName(flexcachesSvmName *string) *FlexcacheOriginCollectionGetParams {
	o.SetFlexcachesSvmName(flexcachesSvmName)
	return o
}

// SetFlexcachesSvmName adds the flexcachesSvmName to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFlexcachesSvmName(flexcachesSvmName *string) {
	o.FlexcachesSvmName = flexcachesSvmName
}

// WithFlexcachesSvmUUID adds the flexcachesSvmUUID to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFlexcachesSvmUUID(flexcachesSvmUUID *string) *FlexcacheOriginCollectionGetParams {
	o.SetFlexcachesSvmUUID(flexcachesSvmUUID)
	return o
}

// SetFlexcachesSvmUUID adds the flexcachesSvmUuid to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFlexcachesSvmUUID(flexcachesSvmUUID *string) {
	o.FlexcachesSvmUUID = flexcachesSvmUUID
}

// WithFlexcachesVolumeName adds the flexcachesVolumeName to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFlexcachesVolumeName(flexcachesVolumeName *string) *FlexcacheOriginCollectionGetParams {
	o.SetFlexcachesVolumeName(flexcachesVolumeName)
	return o
}

// SetFlexcachesVolumeName adds the flexcachesVolumeName to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFlexcachesVolumeName(flexcachesVolumeName *string) {
	o.FlexcachesVolumeName = flexcachesVolumeName
}

// WithFlexcachesVolumeUUID adds the flexcachesVolumeUUID to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithFlexcachesVolumeUUID(flexcachesVolumeUUID *string) *FlexcacheOriginCollectionGetParams {
	o.SetFlexcachesVolumeUUID(flexcachesVolumeUUID)
	return o
}

// SetFlexcachesVolumeUUID adds the flexcachesVolumeUuid to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetFlexcachesVolumeUUID(flexcachesVolumeUUID *string) {
	o.FlexcachesVolumeUUID = flexcachesVolumeUUID
}

// WithGlobalFileLockingEnabled adds the globalFileLockingEnabled to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithGlobalFileLockingEnabled(globalFileLockingEnabled *bool) *FlexcacheOriginCollectionGetParams {
	o.SetGlobalFileLockingEnabled(globalFileLockingEnabled)
	return o
}

// SetGlobalFileLockingEnabled adds the globalFileLockingEnabled to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetGlobalFileLockingEnabled(globalFileLockingEnabled *bool) {
	o.GlobalFileLockingEnabled = globalFileLockingEnabled
}

// WithMaxRecords adds the maxRecords to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithMaxRecords(maxRecords *int64) *FlexcacheOriginCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithName adds the name to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithName(name *string) *FlexcacheOriginCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithOrderBy(orderBy []string) *FlexcacheOriginCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithReturnRecords(returnRecords *bool) *FlexcacheOriginCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *FlexcacheOriginCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithSvmName(svmName *string) *FlexcacheOriginCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithSvmUUID(svmUUID *string) *FlexcacheOriginCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) WithUUID(uuid *string) *FlexcacheOriginCollectionGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the flexcache origin collection get params
func (o *FlexcacheOriginCollectionGetParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *FlexcacheOriginCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BlockLevelInvalidation != nil {

		// query param block_level_invalidation
		var qrBlockLevelInvalidation bool

		if o.BlockLevelInvalidation != nil {
			qrBlockLevelInvalidation = *o.BlockLevelInvalidation
		}
		qBlockLevelInvalidation := swag.FormatBool(qrBlockLevelInvalidation)
		if qBlockLevelInvalidation != "" {

			if err := r.SetQueryParam("block_level_invalidation", qBlockLevelInvalidation); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.FlexcachesClusterName != nil {

		// query param flexcaches.cluster.name
		var qrFlexcachesClusterName string

		if o.FlexcachesClusterName != nil {
			qrFlexcachesClusterName = *o.FlexcachesClusterName
		}
		qFlexcachesClusterName := qrFlexcachesClusterName
		if qFlexcachesClusterName != "" {

			if err := r.SetQueryParam("flexcaches.cluster.name", qFlexcachesClusterName); err != nil {
				return err
			}
		}
	}

	if o.FlexcachesClusterUUID != nil {

		// query param flexcaches.cluster.uuid
		var qrFlexcachesClusterUUID string

		if o.FlexcachesClusterUUID != nil {
			qrFlexcachesClusterUUID = *o.FlexcachesClusterUUID
		}
		qFlexcachesClusterUUID := qrFlexcachesClusterUUID
		if qFlexcachesClusterUUID != "" {

			if err := r.SetQueryParam("flexcaches.cluster.uuid", qFlexcachesClusterUUID); err != nil {
				return err
			}
		}
	}

	if o.FlexcachesCreateTime != nil {

		// query param flexcaches.create_time
		var qrFlexcachesCreateTime string

		if o.FlexcachesCreateTime != nil {
			qrFlexcachesCreateTime = *o.FlexcachesCreateTime
		}
		qFlexcachesCreateTime := qrFlexcachesCreateTime
		if qFlexcachesCreateTime != "" {

			if err := r.SetQueryParam("flexcaches.create_time", qFlexcachesCreateTime); err != nil {
				return err
			}
		}
	}

	if o.FlexcachesIPAddress != nil {

		// query param flexcaches.ip_address
		var qrFlexcachesIPAddress string

		if o.FlexcachesIPAddress != nil {
			qrFlexcachesIPAddress = *o.FlexcachesIPAddress
		}
		qFlexcachesIPAddress := qrFlexcachesIPAddress
		if qFlexcachesIPAddress != "" {

			if err := r.SetQueryParam("flexcaches.ip_address", qFlexcachesIPAddress); err != nil {
				return err
			}
		}
	}

	if o.FlexcachesSize != nil {

		// query param flexcaches.size
		var qrFlexcachesSize int64

		if o.FlexcachesSize != nil {
			qrFlexcachesSize = *o.FlexcachesSize
		}
		qFlexcachesSize := swag.FormatInt64(qrFlexcachesSize)
		if qFlexcachesSize != "" {

			if err := r.SetQueryParam("flexcaches.size", qFlexcachesSize); err != nil {
				return err
			}
		}
	}

	if o.FlexcachesState != nil {

		// query param flexcaches.state
		var qrFlexcachesState string

		if o.FlexcachesState != nil {
			qrFlexcachesState = *o.FlexcachesState
		}
		qFlexcachesState := qrFlexcachesState
		if qFlexcachesState != "" {

			if err := r.SetQueryParam("flexcaches.state", qFlexcachesState); err != nil {
				return err
			}
		}
	}

	if o.FlexcachesSvmName != nil {

		// query param flexcaches.svm.name
		var qrFlexcachesSvmName string

		if o.FlexcachesSvmName != nil {
			qrFlexcachesSvmName = *o.FlexcachesSvmName
		}
		qFlexcachesSvmName := qrFlexcachesSvmName
		if qFlexcachesSvmName != "" {

			if err := r.SetQueryParam("flexcaches.svm.name", qFlexcachesSvmName); err != nil {
				return err
			}
		}
	}

	if o.FlexcachesSvmUUID != nil {

		// query param flexcaches.svm.uuid
		var qrFlexcachesSvmUUID string

		if o.FlexcachesSvmUUID != nil {
			qrFlexcachesSvmUUID = *o.FlexcachesSvmUUID
		}
		qFlexcachesSvmUUID := qrFlexcachesSvmUUID
		if qFlexcachesSvmUUID != "" {

			if err := r.SetQueryParam("flexcaches.svm.uuid", qFlexcachesSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.FlexcachesVolumeName != nil {

		// query param flexcaches.volume.name
		var qrFlexcachesVolumeName string

		if o.FlexcachesVolumeName != nil {
			qrFlexcachesVolumeName = *o.FlexcachesVolumeName
		}
		qFlexcachesVolumeName := qrFlexcachesVolumeName
		if qFlexcachesVolumeName != "" {

			if err := r.SetQueryParam("flexcaches.volume.name", qFlexcachesVolumeName); err != nil {
				return err
			}
		}
	}

	if o.FlexcachesVolumeUUID != nil {

		// query param flexcaches.volume.uuid
		var qrFlexcachesVolumeUUID string

		if o.FlexcachesVolumeUUID != nil {
			qrFlexcachesVolumeUUID = *o.FlexcachesVolumeUUID
		}
		qFlexcachesVolumeUUID := qrFlexcachesVolumeUUID
		if qFlexcachesVolumeUUID != "" {

			if err := r.SetQueryParam("flexcaches.volume.uuid", qFlexcachesVolumeUUID); err != nil {
				return err
			}
		}
	}

	if o.GlobalFileLockingEnabled != nil {

		// query param global_file_locking_enabled
		var qrGlobalFileLockingEnabled bool

		if o.GlobalFileLockingEnabled != nil {
			qrGlobalFileLockingEnabled = *o.GlobalFileLockingEnabled
		}
		qGlobalFileLockingEnabled := swag.FormatBool(qrGlobalFileLockingEnabled)
		if qGlobalFileLockingEnabled != "" {

			if err := r.SetQueryParam("global_file_locking_enabled", qGlobalFileLockingEnabled); err != nil {
				return err
			}
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
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

	if o.SvmName != nil {

		// query param svm.name
		var qrSvmName string

		if o.SvmName != nil {
			qrSvmName = *o.SvmName
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SvmUUID != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SvmUUID != nil {
			qrSvmUUID = *o.SvmUUID
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.UUID != nil {

		// query param uuid
		var qrUUID string

		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFlexcacheOriginCollectionGet binds the parameter fields
func (o *FlexcacheOriginCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamFlexcacheOriginCollectionGet binds the parameter order_by
func (o *FlexcacheOriginCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
