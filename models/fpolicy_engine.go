// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FpolicyEngine Defines how ONTAP makes and manages connections to external FPolicy servers.
//
// swagger:model fpolicy_engine
type FpolicyEngine struct {

	// buffer size
	BufferSize *FpolicyEngineInlineBufferSize `json:"buffer_size,omitempty"`

	// certificate
	Certificate *FpolicyEngineInlineCertificate `json:"certificate,omitempty"`

	// The format for the notification messages sent to the FPolicy servers.
	//   The possible values are:
	//     * xml  - Notifications sent to the FPolicy server will be formatted using the XML schema.
	//     * protobuf - Notifications sent to the FPolicy server will be formatted using Protobuf schema, which is a binary form.
	//
	// Enum: ["xml","protobuf"]
	Format *string `json:"format,omitempty"`

	// fpolicy engine inline primary servers
	// Example: ["10.132.145.20","10.140.101.109"]
	FpolicyEngineInlinePrimaryServers []*string `json:"primary_servers,omitempty"`

	// fpolicy engine inline secondary servers
	// Example: ["10.132.145.20","10.132.145.21"]
	FpolicyEngineInlineSecondaryServers []*string `json:"secondary_servers,omitempty"`

	// Specifies the ISO-8601 interval time for a storage appliance to send Keep Alive message to an FPolicy server. The allowed range is between 10 to 600 seconds.
	// Example: PT2M
	KeepAliveInterval *string `json:"keep_alive_interval,omitempty"`

	// Specifies the maximum number of outstanding requests for the FPolicy server. It is used to specify maximum outstanding requests that will be queued up for the FPolicy server. The value for this field must be between 1 and 10000.  The default values are 500, 1000 or 2000 for Low-end(<64 GB memory), Mid-end(>=64 GB memory) and High-end(>=128 GB memory) Platforms respectively.
	// Example: 500
	// Maximum: 10000
	// Minimum: 1
	MaxServerRequests *int64 `json:"max_server_requests,omitempty"`

	// Specifies the name to assign to the external server configuration.
	// Example: fp_ex_eng
	Name *string `json:"name,omitempty"`

	// Port number of the FPolicy server application.
	// Example: 9876
	Port *int64 `json:"port,omitempty"`

	// Specifies the ISO-8601 timeout duration for a screen request to be aborted by a storage appliance. The allowed range is between 0 to 200 seconds.
	// Example: PT40S
	RequestAbortTimeout *string `json:"request_abort_timeout,omitempty"`

	// Specifies the ISO-8601 timeout duration for a screen request to be processed by an FPolicy server. The allowed range is between 0 to 100 seconds.
	// Example: PT20S
	RequestCancelTimeout *string `json:"request_cancel_timeout,omitempty"`

	// resiliency
	Resiliency *FpolicyEngineInlineResiliency `json:"resiliency,omitempty"`

	// Specifies the ISO-8601 timeout duration in which a throttled FPolicy server must complete at least one screen request. If no request is processed within the timeout, connection to the FPolicy server is terminated. The allowed range is between 0 to 100 seconds.
	// Example: PT1M
	ServerProgressTimeout *string `json:"server_progress_timeout,omitempty"`

	// Specifies the SSL option for external communication with the FPolicy server. Possible values include the following:
	// * no_auth       When set to "no_auth", no authentication takes place.
	// * server_auth   When set to "server_auth", only the FPolicy server is authenticated by the SVM. With this option, before creating the FPolicy external engine, the administrator must install the public certificate of the certificate authority (CA) that signed the FPolicy server certificate.
	// * mutual_auth   When set to "mutual_auth", mutual authentication takes place between the SVM and the FPolicy server. This means authentication of the FPolicy server by the SVM along with authentication of the SVM by the FPolicy server. With this option, before creating the FPolicy external engine, the administrator must install the public certificate of the certificate authority (CA) that signed the FPolicy server certificate along with the public certificate and key file for authentication of the SVM.
	//
	// Enum: ["no_auth","server_auth","mutual_auth"]
	SslOption *string `json:"ssl_option,omitempty"`

	// Specifies the ISO-8601 interval time for a storage appliance to query a status request from an FPolicy server. The allowed range is between 0 to 50 seconds.
	// Example: PT10S
	StatusRequestInterval *string `json:"status_request_interval,omitempty"`

	// svm
	Svm *FpolicyEngineInlineSvm `json:"svm,omitempty"`

	// The notification mode determines what ONTAP does after sending notifications to FPolicy servers.
	//   The possible values are:
	//     * synchronous  - After sending a notification, wait for a response from the FPolicy server.
	//     * asynchronous - After sending a notification, file request processing continues.
	//
	// Enum: ["synchronous","asynchronous"]
	Type *string `json:"type,omitempty"`
}

// Validate validates this fpolicy engine
func (m *FpolicyEngine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBufferSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxServerRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResiliency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyEngine) validateBufferSize(formats strfmt.Registry) error {
	if swag.IsZero(m.BufferSize) { // not required
		return nil
	}

	if m.BufferSize != nil {
		if err := m.BufferSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buffer_size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("buffer_size")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyEngine) validateCertificate(formats strfmt.Registry) error {
	if swag.IsZero(m.Certificate) { // not required
		return nil
	}

	if m.Certificate != nil {
		if err := m.Certificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

var fpolicyEngineTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["xml","protobuf"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyEngineTypeFormatPropEnum = append(fpolicyEngineTypeFormatPropEnum, v)
	}
}

const (

	// FpolicyEngineFormatXML captures enum value "xml"
	FpolicyEngineFormatXML string = "xml"

	// FpolicyEngineFormatProtobuf captures enum value "protobuf"
	FpolicyEngineFormatProtobuf string = "protobuf"
)

// prop value enum
func (m *FpolicyEngine) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyEngineTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FpolicyEngine) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", *m.Format); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyEngine) validateMaxServerRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxServerRequests) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_server_requests", "body", *m.MaxServerRequests, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_server_requests", "body", *m.MaxServerRequests, 10000, false); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyEngine) validateResiliency(formats strfmt.Registry) error {
	if swag.IsZero(m.Resiliency) { // not required
		return nil
	}

	if m.Resiliency != nil {
		if err := m.Resiliency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resiliency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resiliency")
			}
			return err
		}
	}

	return nil
}

var fpolicyEngineTypeSslOptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["no_auth","server_auth","mutual_auth"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyEngineTypeSslOptionPropEnum = append(fpolicyEngineTypeSslOptionPropEnum, v)
	}
}

const (

	// FpolicyEngineSslOptionNoAuth captures enum value "no_auth"
	FpolicyEngineSslOptionNoAuth string = "no_auth"

	// FpolicyEngineSslOptionServerAuth captures enum value "server_auth"
	FpolicyEngineSslOptionServerAuth string = "server_auth"

	// FpolicyEngineSslOptionMutualAuth captures enum value "mutual_auth"
	FpolicyEngineSslOptionMutualAuth string = "mutual_auth"
)

// prop value enum
func (m *FpolicyEngine) validateSslOptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyEngineTypeSslOptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FpolicyEngine) validateSslOption(formats strfmt.Registry) error {
	if swag.IsZero(m.SslOption) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslOptionEnum("ssl_option", "body", *m.SslOption); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyEngine) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

var fpolicyEngineTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["synchronous","asynchronous"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyEngineTypeTypePropEnum = append(fpolicyEngineTypeTypePropEnum, v)
	}
}

const (

	// FpolicyEngineTypeSynchronous captures enum value "synchronous"
	FpolicyEngineTypeSynchronous string = "synchronous"

	// FpolicyEngineTypeAsynchronous captures enum value "asynchronous"
	FpolicyEngineTypeAsynchronous string = "asynchronous"
)

// prop value enum
func (m *FpolicyEngine) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyEngineTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FpolicyEngine) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fpolicy engine based on the context it is used
func (m *FpolicyEngine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBufferSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCertificate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResiliency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyEngine) contextValidateBufferSize(ctx context.Context, formats strfmt.Registry) error {

	if m.BufferSize != nil {

		if swag.IsZero(m.BufferSize) { // not required
			return nil
		}

		if err := m.BufferSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buffer_size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("buffer_size")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyEngine) contextValidateCertificate(ctx context.Context, formats strfmt.Registry) error {

	if m.Certificate != nil {

		if swag.IsZero(m.Certificate) { // not required
			return nil
		}

		if err := m.Certificate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyEngine) contextValidateResiliency(ctx context.Context, formats strfmt.Registry) error {

	if m.Resiliency != nil {

		if swag.IsZero(m.Resiliency) { // not required
			return nil
		}

		if err := m.Resiliency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resiliency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resiliency")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyEngine) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {

		if swag.IsZero(m.Svm) { // not required
			return nil
		}

		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEngine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEngine) UnmarshalBinary(b []byte) error {
	var res FpolicyEngine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyEngineInlineBufferSize Specifies the send and receive buffer size of the connected socket for the FPolicy server.
//
// swagger:model fpolicy_engine_inline_buffer_size
type FpolicyEngineInlineBufferSize struct {

	// Specifies the receive buffer size of the connected socket for the FPolicy server. Default value is 256KB.
	// Maximum: 7.89516e+06
	// Minimum: 0
	RecvBuffer *int64 `json:"recv_buffer,omitempty"`

	// Specifies the send buffer size of the connected socket for the FPolicy server. Default value 1MB.
	// Maximum: 7.89516e+06
	// Minimum: 0
	SendBuffer *int64 `json:"send_buffer,omitempty"`
}

// Validate validates this fpolicy engine inline buffer size
func (m *FpolicyEngineInlineBufferSize) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecvBuffer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSendBuffer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyEngineInlineBufferSize) validateRecvBuffer(formats strfmt.Registry) error {
	if swag.IsZero(m.RecvBuffer) { // not required
		return nil
	}

	if err := validate.MinimumInt("buffer_size"+"."+"recv_buffer", "body", *m.RecvBuffer, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("buffer_size"+"."+"recv_buffer", "body", *m.RecvBuffer, 7.89516e+06, false); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyEngineInlineBufferSize) validateSendBuffer(formats strfmt.Registry) error {
	if swag.IsZero(m.SendBuffer) { // not required
		return nil
	}

	if err := validate.MinimumInt("buffer_size"+"."+"send_buffer", "body", *m.SendBuffer, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("buffer_size"+"."+"send_buffer", "body", *m.SendBuffer, 7.89516e+06, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fpolicy engine inline buffer size based on context it is used
func (m *FpolicyEngineInlineBufferSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEngineInlineBufferSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEngineInlineBufferSize) UnmarshalBinary(b []byte) error {
	var res FpolicyEngineInlineBufferSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyEngineInlineCertificate Provides details about certificate used to authenticate the FPolicy server.
//
// swagger:model fpolicy_engine_inline_certificate
type FpolicyEngineInlineCertificate struct {

	// Specifies the certificate authority (CA) name of the certificate
	// used for authentication if SSL authentication between the SVM and the FPolicy
	// server is configured.
	//
	// Example: TASample1
	Ca *string `json:"ca,omitempty"`

	// Specifies the certificate name as a fully qualified domain
	// name (FQDN) or custom common name. The certificate is used if SSL authentication
	// between the SVM and the FPolicy server is configured.
	//
	// Example: Sample1-FPolicy-Client
	Name *string `json:"name,omitempty"`

	// Specifies the serial number of the certificate used for
	// authentication if SSL authentication between the SVM and the FPolicy
	// server is configured.
	//
	// Example: 8DDE112A114D1FBC
	SerialNumber *string `json:"serial_number,omitempty"`
}

// Validate validates this fpolicy engine inline certificate
func (m *FpolicyEngineInlineCertificate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fpolicy engine inline certificate based on context it is used
func (m *FpolicyEngineInlineCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEngineInlineCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEngineInlineCertificate) UnmarshalBinary(b []byte) error {
	var res FpolicyEngineInlineCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyEngineInlineResiliency If all primary and secondary servers are down, or if no response is received from the FPolicy servers, file access events are stored inside the storage controller under the specified resiliency-directory-path.
//
// swagger:model fpolicy_engine_inline_resiliency
type FpolicyEngineInlineResiliency struct {

	// Specifies the directory path under the SVM namespace,
	// where notifications are stored in the files whenever a network outage happens.
	//
	// Example: /dir1
	DirectoryPath *string `json:"directory_path,omitempty"`

	// Specifies whether the resiliency feature is enabled or not.
	// Default is false.
	//
	Enabled *bool `json:"enabled,omitempty"`

	// Specifies the ISO-8601 duration, for which the notifications are written
	// to files inside the storage controller during a network outage. The value for
	// this field must be between 0 and 600 seconds. Default is 180 seconds.
	//
	// Example: PT3M
	RetentionDuration *string `json:"retention_duration,omitempty"`
}

// Validate validates this fpolicy engine inline resiliency
func (m *FpolicyEngineInlineResiliency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fpolicy engine inline resiliency based on context it is used
func (m *FpolicyEngineInlineResiliency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEngineInlineResiliency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEngineInlineResiliency) UnmarshalBinary(b []byte) error {
	var res FpolicyEngineInlineResiliency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyEngineInlineSvm fpolicy engine inline svm
//
// swagger:model fpolicy_engine_inline_svm
type FpolicyEngineInlineSvm struct {

	// SVM UUID
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this fpolicy engine inline svm
func (m *FpolicyEngineInlineSvm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fpolicy engine inline svm based on the context it is used
func (m *FpolicyEngineInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyEngineInlineSvm) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "svm"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEngineInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEngineInlineSvm) UnmarshalBinary(b []byte) error {
	var res FpolicyEngineInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
