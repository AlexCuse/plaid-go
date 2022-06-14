/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ConnectedApplication Describes the connected application for a particular end user.
type ConnectedApplication struct {
	// This field will map to the application ID that is returned from /item/applications/list, or provided to the institution in an oauth redirect.
	ApplicationId string `json:"application_id"`
	// The name of the application
	Name string `json:"name"`
	// A human-readable name of the application for display purposes
	DisplayName NullableString `json:"display_name,omitempty"`
	// A URL that links to the application logo image.
	LogoUrl NullableString `json:"logo_url,omitempty"`
	// The URL for the application's website
	ApplicationUrl NullableString `json:"application_url,omitempty"`
	// A string provided by the connected app stating why they use their respective enabled products.
	ReasonForAccess NullableString `json:"reason_for_access,omitempty"`
	// The date this application was linked in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC.
	CreatedAt string `json:"created_at"`
	Scopes NullableScopesNullable `json:"scopes,omitempty"`
}

// NewConnectedApplication instantiates a new ConnectedApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedApplication(applicationId string, name string, createdAt string) *ConnectedApplication {
	this := ConnectedApplication{}
	this.ApplicationId = applicationId
	this.Name = name
	this.CreatedAt = createdAt
	return &this
}

// NewConnectedApplicationWithDefaults instantiates a new ConnectedApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedApplicationWithDefaults() *ConnectedApplication {
	this := ConnectedApplication{}
	return &this
}

// GetApplicationId returns the ApplicationId field value
func (o *ConnectedApplication) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *ConnectedApplication) GetApplicationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *ConnectedApplication) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetName returns the Name field value
func (o *ConnectedApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectedApplication) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectedApplication) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectedApplication) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedApplication) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ConnectedApplication) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *ConnectedApplication) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *ConnectedApplication) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *ConnectedApplication) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectedApplication) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedApplication) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *ConnectedApplication) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *ConnectedApplication) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}
// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *ConnectedApplication) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *ConnectedApplication) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetApplicationUrl returns the ApplicationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectedApplication) GetApplicationUrl() string {
	if o == nil || o.ApplicationUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationUrl.Get()
}

// GetApplicationUrlOk returns a tuple with the ApplicationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedApplication) GetApplicationUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationUrl.Get(), o.ApplicationUrl.IsSet()
}

// HasApplicationUrl returns a boolean if a field has been set.
func (o *ConnectedApplication) HasApplicationUrl() bool {
	if o != nil && o.ApplicationUrl.IsSet() {
		return true
	}

	return false
}

// SetApplicationUrl gets a reference to the given NullableString and assigns it to the ApplicationUrl field.
func (o *ConnectedApplication) SetApplicationUrl(v string) {
	o.ApplicationUrl.Set(&v)
}
// SetApplicationUrlNil sets the value for ApplicationUrl to be an explicit nil
func (o *ConnectedApplication) SetApplicationUrlNil() {
	o.ApplicationUrl.Set(nil)
}

// UnsetApplicationUrl ensures that no value is present for ApplicationUrl, not even an explicit nil
func (o *ConnectedApplication) UnsetApplicationUrl() {
	o.ApplicationUrl.Unset()
}

// GetReasonForAccess returns the ReasonForAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectedApplication) GetReasonForAccess() string {
	if o == nil || o.ReasonForAccess.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReasonForAccess.Get()
}

// GetReasonForAccessOk returns a tuple with the ReasonForAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedApplication) GetReasonForAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReasonForAccess.Get(), o.ReasonForAccess.IsSet()
}

// HasReasonForAccess returns a boolean if a field has been set.
func (o *ConnectedApplication) HasReasonForAccess() bool {
	if o != nil && o.ReasonForAccess.IsSet() {
		return true
	}

	return false
}

// SetReasonForAccess gets a reference to the given NullableString and assigns it to the ReasonForAccess field.
func (o *ConnectedApplication) SetReasonForAccess(v string) {
	o.ReasonForAccess.Set(&v)
}
// SetReasonForAccessNil sets the value for ReasonForAccess to be an explicit nil
func (o *ConnectedApplication) SetReasonForAccessNil() {
	o.ReasonForAccess.Set(nil)
}

// UnsetReasonForAccess ensures that no value is present for ReasonForAccess, not even an explicit nil
func (o *ConnectedApplication) UnsetReasonForAccess() {
	o.ReasonForAccess.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ConnectedApplication) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ConnectedApplication) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ConnectedApplication) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectedApplication) GetScopes() ScopesNullable {
	if o == nil || o.Scopes.Get() == nil {
		var ret ScopesNullable
		return ret
	}
	return *o.Scopes.Get()
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedApplication) GetScopesOk() (*ScopesNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scopes.Get(), o.Scopes.IsSet()
}

// HasScopes returns a boolean if a field has been set.
func (o *ConnectedApplication) HasScopes() bool {
	if o != nil && o.Scopes.IsSet() {
		return true
	}

	return false
}

// SetScopes gets a reference to the given NullableScopesNullable and assigns it to the Scopes field.
func (o *ConnectedApplication) SetScopes(v ScopesNullable) {
	o.Scopes.Set(&v)
}
// SetScopesNil sets the value for Scopes to be an explicit nil
func (o *ConnectedApplication) SetScopesNil() {
	o.Scopes.Set(nil)
}

// UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
func (o *ConnectedApplication) UnsetScopes() {
	o.Scopes.Unset()
}

func (o ConnectedApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["application_id"] = o.ApplicationId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if o.LogoUrl.IsSet() {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if o.ApplicationUrl.IsSet() {
		toSerialize["application_url"] = o.ApplicationUrl.Get()
	}
	if o.ReasonForAccess.IsSet() {
		toSerialize["reason_for_access"] = o.ReasonForAccess.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Scopes.IsSet() {
		toSerialize["scopes"] = o.Scopes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableConnectedApplication struct {
	value *ConnectedApplication
	isSet bool
}

func (v NullableConnectedApplication) Get() *ConnectedApplication {
	return v.value
}

func (v *NullableConnectedApplication) Set(val *ConnectedApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedApplication(val *ConnectedApplication) *NullableConnectedApplication {
	return &NullableConnectedApplication{value: val, isSet: true}
}

func (v NullableConnectedApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


