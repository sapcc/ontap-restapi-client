// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityOauth2GetParams creates a new SecurityOauth2GetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityOauth2GetParams() *SecurityOauth2GetParams {
	return &SecurityOauth2GetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityOauth2GetParamsWithTimeout creates a new SecurityOauth2GetParams object
// with the ability to set a timeout on a request.
func NewSecurityOauth2GetParamsWithTimeout(timeout time.Duration) *SecurityOauth2GetParams {
	return &SecurityOauth2GetParams{
		timeout: timeout,
	}
}

// NewSecurityOauth2GetParamsWithContext creates a new SecurityOauth2GetParams object
// with the ability to set a context for a request.
func NewSecurityOauth2GetParamsWithContext(ctx context.Context) *SecurityOauth2GetParams {
	return &SecurityOauth2GetParams{
		Context: ctx,
	}
}

// NewSecurityOauth2GetParamsWithHTTPClient creates a new SecurityOauth2GetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityOauth2GetParamsWithHTTPClient(client *http.Client) *SecurityOauth2GetParams {
	return &SecurityOauth2GetParams{
		HTTPClient: client,
	}
}

/*
SecurityOauth2GetParams contains all the parameters to send to the API endpoint

	for the security oauth2 get operation.

	Typically these are written to a http.Request.
*/
type SecurityOauth2GetParams struct {

	/* Application.

	   Filter by application
	*/
	Application *string

	/* Audience.

	   Filter by audience
	*/
	Audience *string

	/* ClientID.

	   Filter by client_id
	*/
	ClientID *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* HashedClientSecret.

	   Filter by hashed_client_secret
	*/
	HashedClientSecret *string

	/* IntrospectionEndpointURI.

	   Filter by introspection.endpoint_uri
	*/
	IntrospectionEndpointURI *string

	/* IntrospectionInterval.

	   Filter by introspection.interval
	*/
	IntrospectionInterval *string

	/* Issuer.

	   Filter by issuer
	*/
	Issuer *string

	/* JwksProviderURI.

	   Filter by jwks.provider_uri
	*/
	JwksProviderURI *string

	/* JwksRefreshInterval.

	   Filter by jwks.refresh_interval
	*/
	JwksRefreshInterval *string

	/* Name.

	   OAuth 2.0 configuration name.
	*/
	Name string

	/* OutgoingProxy.

	   Filter by outgoing_proxy
	*/
	OutgoingProxy *string

	/* Provider.

	   Filter by provider
	*/
	Provider *string

	/* RemoteUserClaim.

	   Filter by remote_user_claim
	*/
	RemoteUserClaim *string

	/* UseLocalRolesIfPresent.

	   Filter by use_local_roles_if_present
	*/
	UseLocalRolesIfPresent *bool

	/* UseMutualTLS.

	   Filter by use_mutual_tls
	*/
	UseMutualTLS *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security oauth2 get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityOauth2GetParams) WithDefaults() *SecurityOauth2GetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security oauth2 get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityOauth2GetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithTimeout(timeout time.Duration) *SecurityOauth2GetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithContext(ctx context.Context) *SecurityOauth2GetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithHTTPClient(client *http.Client) *SecurityOauth2GetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplication adds the application to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithApplication(application *string) *SecurityOauth2GetParams {
	o.SetApplication(application)
	return o
}

// SetApplication adds the application to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetApplication(application *string) {
	o.Application = application
}

// WithAudience adds the audience to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithAudience(audience *string) *SecurityOauth2GetParams {
	o.SetAudience(audience)
	return o
}

// SetAudience adds the audience to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetAudience(audience *string) {
	o.Audience = audience
}

// WithClientID adds the clientID to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithClientID(clientID *string) *SecurityOauth2GetParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithFields adds the fields to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithFields(fields []string) *SecurityOauth2GetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithHashedClientSecret adds the hashedClientSecret to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithHashedClientSecret(hashedClientSecret *string) *SecurityOauth2GetParams {
	o.SetHashedClientSecret(hashedClientSecret)
	return o
}

// SetHashedClientSecret adds the hashedClientSecret to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetHashedClientSecret(hashedClientSecret *string) {
	o.HashedClientSecret = hashedClientSecret
}

// WithIntrospectionEndpointURI adds the introspectionEndpointURI to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithIntrospectionEndpointURI(introspectionEndpointURI *string) *SecurityOauth2GetParams {
	o.SetIntrospectionEndpointURI(introspectionEndpointURI)
	return o
}

// SetIntrospectionEndpointURI adds the introspectionEndpointUri to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetIntrospectionEndpointURI(introspectionEndpointURI *string) {
	o.IntrospectionEndpointURI = introspectionEndpointURI
}

// WithIntrospectionInterval adds the introspectionInterval to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithIntrospectionInterval(introspectionInterval *string) *SecurityOauth2GetParams {
	o.SetIntrospectionInterval(introspectionInterval)
	return o
}

// SetIntrospectionInterval adds the introspectionInterval to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetIntrospectionInterval(introspectionInterval *string) {
	o.IntrospectionInterval = introspectionInterval
}

// WithIssuer adds the issuer to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithIssuer(issuer *string) *SecurityOauth2GetParams {
	o.SetIssuer(issuer)
	return o
}

// SetIssuer adds the issuer to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetIssuer(issuer *string) {
	o.Issuer = issuer
}

// WithJwksProviderURI adds the jwksProviderURI to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithJwksProviderURI(jwksProviderURI *string) *SecurityOauth2GetParams {
	o.SetJwksProviderURI(jwksProviderURI)
	return o
}

// SetJwksProviderURI adds the jwksProviderUri to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetJwksProviderURI(jwksProviderURI *string) {
	o.JwksProviderURI = jwksProviderURI
}

// WithJwksRefreshInterval adds the jwksRefreshInterval to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithJwksRefreshInterval(jwksRefreshInterval *string) *SecurityOauth2GetParams {
	o.SetJwksRefreshInterval(jwksRefreshInterval)
	return o
}

// SetJwksRefreshInterval adds the jwksRefreshInterval to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetJwksRefreshInterval(jwksRefreshInterval *string) {
	o.JwksRefreshInterval = jwksRefreshInterval
}

// WithName adds the name to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithName(name string) *SecurityOauth2GetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetName(name string) {
	o.Name = name
}

// WithOutgoingProxy adds the outgoingProxy to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithOutgoingProxy(outgoingProxy *string) *SecurityOauth2GetParams {
	o.SetOutgoingProxy(outgoingProxy)
	return o
}

// SetOutgoingProxy adds the outgoingProxy to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetOutgoingProxy(outgoingProxy *string) {
	o.OutgoingProxy = outgoingProxy
}

// WithProvider adds the provider to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithProvider(provider *string) *SecurityOauth2GetParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithRemoteUserClaim adds the remoteUserClaim to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithRemoteUserClaim(remoteUserClaim *string) *SecurityOauth2GetParams {
	o.SetRemoteUserClaim(remoteUserClaim)
	return o
}

// SetRemoteUserClaim adds the remoteUserClaim to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetRemoteUserClaim(remoteUserClaim *string) {
	o.RemoteUserClaim = remoteUserClaim
}

// WithUseLocalRolesIfPresent adds the useLocalRolesIfPresent to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithUseLocalRolesIfPresent(useLocalRolesIfPresent *bool) *SecurityOauth2GetParams {
	o.SetUseLocalRolesIfPresent(useLocalRolesIfPresent)
	return o
}

// SetUseLocalRolesIfPresent adds the useLocalRolesIfPresent to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetUseLocalRolesIfPresent(useLocalRolesIfPresent *bool) {
	o.UseLocalRolesIfPresent = useLocalRolesIfPresent
}

// WithUseMutualTLS adds the useMutualTLS to the security oauth2 get params
func (o *SecurityOauth2GetParams) WithUseMutualTLS(useMutualTLS *string) *SecurityOauth2GetParams {
	o.SetUseMutualTLS(useMutualTLS)
	return o
}

// SetUseMutualTLS adds the useMutualTls to the security oauth2 get params
func (o *SecurityOauth2GetParams) SetUseMutualTLS(useMutualTLS *string) {
	o.UseMutualTLS = useMutualTLS
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityOauth2GetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Application != nil {

		// query param application
		var qrApplication string

		if o.Application != nil {
			qrApplication = *o.Application
		}
		qApplication := qrApplication
		if qApplication != "" {

			if err := r.SetQueryParam("application", qApplication); err != nil {
				return err
			}
		}
	}

	if o.Audience != nil {

		// query param audience
		var qrAudience string

		if o.Audience != nil {
			qrAudience = *o.Audience
		}
		qAudience := qrAudience
		if qAudience != "" {

			if err := r.SetQueryParam("audience", qAudience); err != nil {
				return err
			}
		}
	}

	if o.ClientID != nil {

		// query param client_id
		var qrClientID string

		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := qrClientID
		if qClientID != "" {

			if err := r.SetQueryParam("client_id", qClientID); err != nil {
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

	if o.HashedClientSecret != nil {

		// query param hashed_client_secret
		var qrHashedClientSecret string

		if o.HashedClientSecret != nil {
			qrHashedClientSecret = *o.HashedClientSecret
		}
		qHashedClientSecret := qrHashedClientSecret
		if qHashedClientSecret != "" {

			if err := r.SetQueryParam("hashed_client_secret", qHashedClientSecret); err != nil {
				return err
			}
		}
	}

	if o.IntrospectionEndpointURI != nil {

		// query param introspection.endpoint_uri
		var qrIntrospectionEndpointURI string

		if o.IntrospectionEndpointURI != nil {
			qrIntrospectionEndpointURI = *o.IntrospectionEndpointURI
		}
		qIntrospectionEndpointURI := qrIntrospectionEndpointURI
		if qIntrospectionEndpointURI != "" {

			if err := r.SetQueryParam("introspection.endpoint_uri", qIntrospectionEndpointURI); err != nil {
				return err
			}
		}
	}

	if o.IntrospectionInterval != nil {

		// query param introspection.interval
		var qrIntrospectionInterval string

		if o.IntrospectionInterval != nil {
			qrIntrospectionInterval = *o.IntrospectionInterval
		}
		qIntrospectionInterval := qrIntrospectionInterval
		if qIntrospectionInterval != "" {

			if err := r.SetQueryParam("introspection.interval", qIntrospectionInterval); err != nil {
				return err
			}
		}
	}

	if o.Issuer != nil {

		// query param issuer
		var qrIssuer string

		if o.Issuer != nil {
			qrIssuer = *o.Issuer
		}
		qIssuer := qrIssuer
		if qIssuer != "" {

			if err := r.SetQueryParam("issuer", qIssuer); err != nil {
				return err
			}
		}
	}

	if o.JwksProviderURI != nil {

		// query param jwks.provider_uri
		var qrJwksProviderURI string

		if o.JwksProviderURI != nil {
			qrJwksProviderURI = *o.JwksProviderURI
		}
		qJwksProviderURI := qrJwksProviderURI
		if qJwksProviderURI != "" {

			if err := r.SetQueryParam("jwks.provider_uri", qJwksProviderURI); err != nil {
				return err
			}
		}
	}

	if o.JwksRefreshInterval != nil {

		// query param jwks.refresh_interval
		var qrJwksRefreshInterval string

		if o.JwksRefreshInterval != nil {
			qrJwksRefreshInterval = *o.JwksRefreshInterval
		}
		qJwksRefreshInterval := qrJwksRefreshInterval
		if qJwksRefreshInterval != "" {

			if err := r.SetQueryParam("jwks.refresh_interval", qJwksRefreshInterval); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.OutgoingProxy != nil {

		// query param outgoing_proxy
		var qrOutgoingProxy string

		if o.OutgoingProxy != nil {
			qrOutgoingProxy = *o.OutgoingProxy
		}
		qOutgoingProxy := qrOutgoingProxy
		if qOutgoingProxy != "" {

			if err := r.SetQueryParam("outgoing_proxy", qOutgoingProxy); err != nil {
				return err
			}
		}
	}

	if o.Provider != nil {

		// query param provider
		var qrProvider string

		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {

			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}
	}

	if o.RemoteUserClaim != nil {

		// query param remote_user_claim
		var qrRemoteUserClaim string

		if o.RemoteUserClaim != nil {
			qrRemoteUserClaim = *o.RemoteUserClaim
		}
		qRemoteUserClaim := qrRemoteUserClaim
		if qRemoteUserClaim != "" {

			if err := r.SetQueryParam("remote_user_claim", qRemoteUserClaim); err != nil {
				return err
			}
		}
	}

	if o.UseLocalRolesIfPresent != nil {

		// query param use_local_roles_if_present
		var qrUseLocalRolesIfPresent bool

		if o.UseLocalRolesIfPresent != nil {
			qrUseLocalRolesIfPresent = *o.UseLocalRolesIfPresent
		}
		qUseLocalRolesIfPresent := swag.FormatBool(qrUseLocalRolesIfPresent)
		if qUseLocalRolesIfPresent != "" {

			if err := r.SetQueryParam("use_local_roles_if_present", qUseLocalRolesIfPresent); err != nil {
				return err
			}
		}
	}

	if o.UseMutualTLS != nil {

		// query param use_mutual_tls
		var qrUseMutualTLS string

		if o.UseMutualTLS != nil {
			qrUseMutualTLS = *o.UseMutualTLS
		}
		qUseMutualTLS := qrUseMutualTLS
		if qUseMutualTLS != "" {

			if err := r.SetQueryParam("use_mutual_tls", qUseMutualTLS); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityOauth2Get binds the parameter fields
func (o *SecurityOauth2GetParams) bindParamFields(formats strfmt.Registry) []string {
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
