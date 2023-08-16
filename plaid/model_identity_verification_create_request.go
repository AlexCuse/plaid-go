/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.410.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IdentityVerificationCreateRequest Request schema for '/identity_verification/create'
type IdentityVerificationCreateRequest struct {
	// A unique ID that identifies the end user in your system. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.
	ClientUserId *string `json:"client_user_id,omitempty"`
	// A flag specifying whether you would like Plaid to expose a shareable URL for the verification being created.
	IsShareable bool `json:"is_shareable"`
	// ID of the associated Identity Verification template.
	TemplateId string `json:"template_id"`
	// A flag specifying whether the end user has already agreed to a privacy policy specifying that their data will be shared with Plaid for verification purposes.  If `gave_consent` is set to `true`, the `accept_tos` step will be marked as `skipped` and the end user's session will start at the next step requirement.
	GaveConsent bool `json:"gave_consent"`
	User NullableIdentityVerificationCreateRequestUser `json:"user,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// An optional flag specifying how you would like Plaid to handle attempts to create an Identity Verification when an Identity Verification already exists for the provided `client_user_id` and `template_id`. If idempotency is enabled, Plaid will return the existing Identity Verification. If idempotency is disabled, Plaid will reject the request with a `400 Bad Request` status code if an Identity Verification already exists for the supplied `client_user_id` and `template_id`.
	IsIdempotent NullableBool `json:"is_idempotent,omitempty"`
}

// NewIdentityVerificationCreateRequest instantiates a new IdentityVerificationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationCreateRequest(isShareable bool, templateId string, gaveConsent bool) *IdentityVerificationCreateRequest {
	this := IdentityVerificationCreateRequest{}
	this.IsShareable = isShareable
	this.TemplateId = templateId
	this.GaveConsent = gaveConsent
	return &this
}

// NewIdentityVerificationCreateRequestWithDefaults instantiates a new IdentityVerificationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationCreateRequestWithDefaults() *IdentityVerificationCreateRequest {
	this := IdentityVerificationCreateRequest{}
	var gaveConsent bool = false
	this.GaveConsent = gaveConsent
	return &this
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise.
func (o *IdentityVerificationCreateRequest) GetClientUserId() string {
	if o == nil || o.ClientUserId == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil || o.ClientUserId == nil {
		return nil, false
	}
	return o.ClientUserId, true
}

// HasClientUserId returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequest) HasClientUserId() bool {
	if o != nil && o.ClientUserId != nil {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given string and assigns it to the ClientUserId field.
func (o *IdentityVerificationCreateRequest) SetClientUserId(v string) {
	o.ClientUserId = &v
}

// GetIsShareable returns the IsShareable field value
func (o *IdentityVerificationCreateRequest) GetIsShareable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsShareable
}

// GetIsShareableOk returns a tuple with the IsShareable field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateRequest) GetIsShareableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsShareable, true
}

// SetIsShareable sets field value
func (o *IdentityVerificationCreateRequest) SetIsShareable(v bool) {
	o.IsShareable = v
}

// GetTemplateId returns the TemplateId field value
func (o *IdentityVerificationCreateRequest) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateRequest) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *IdentityVerificationCreateRequest) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetGaveConsent returns the GaveConsent field value
func (o *IdentityVerificationCreateRequest) GetGaveConsent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GaveConsent
}

// GetGaveConsentOk returns a tuple with the GaveConsent field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateRequest) GetGaveConsentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GaveConsent, true
}

// SetGaveConsent sets field value
func (o *IdentityVerificationCreateRequest) SetGaveConsent(v bool) {
	o.GaveConsent = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationCreateRequest) GetUser() IdentityVerificationCreateRequestUser {
	if o == nil || o.User.Get() == nil {
		var ret IdentityVerificationCreateRequestUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateRequest) GetUserOk() (*IdentityVerificationCreateRequestUser, bool) {
	if o == nil  {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableIdentityVerificationCreateRequestUser and assigns it to the User field.
func (o *IdentityVerificationCreateRequest) SetUser(v IdentityVerificationCreateRequestUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *IdentityVerificationCreateRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *IdentityVerificationCreateRequest) UnsetUser() {
	o.User.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IdentityVerificationCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IdentityVerificationCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IdentityVerificationCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IdentityVerificationCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIsIdempotent returns the IsIdempotent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationCreateRequest) GetIsIdempotent() bool {
	if o == nil || o.IsIdempotent.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsIdempotent.Get()
}

// GetIsIdempotentOk returns a tuple with the IsIdempotent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateRequest) GetIsIdempotentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsIdempotent.Get(), o.IsIdempotent.IsSet()
}

// HasIsIdempotent returns a boolean if a field has been set.
func (o *IdentityVerificationCreateRequest) HasIsIdempotent() bool {
	if o != nil && o.IsIdempotent.IsSet() {
		return true
	}

	return false
}

// SetIsIdempotent gets a reference to the given NullableBool and assigns it to the IsIdempotent field.
func (o *IdentityVerificationCreateRequest) SetIsIdempotent(v bool) {
	o.IsIdempotent.Set(&v)
}
// SetIsIdempotentNil sets the value for IsIdempotent to be an explicit nil
func (o *IdentityVerificationCreateRequest) SetIsIdempotentNil() {
	o.IsIdempotent.Set(nil)
}

// UnsetIsIdempotent ensures that no value is present for IsIdempotent, not even an explicit nil
func (o *IdentityVerificationCreateRequest) UnsetIsIdempotent() {
	o.IsIdempotent.Unset()
}

func (o IdentityVerificationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientUserId != nil {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if true {
		toSerialize["is_shareable"] = o.IsShareable
	}
	if true {
		toSerialize["template_id"] = o.TemplateId
	}
	if true {
		toSerialize["gave_consent"] = o.GaveConsent
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.IsIdempotent.IsSet() {
		toSerialize["is_idempotent"] = o.IsIdempotent.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityVerificationCreateRequest struct {
	value *IdentityVerificationCreateRequest
	isSet bool
}

func (v NullableIdentityVerificationCreateRequest) Get() *IdentityVerificationCreateRequest {
	return v.value
}

func (v *NullableIdentityVerificationCreateRequest) Set(val *IdentityVerificationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationCreateRequest(val *IdentityVerificationCreateRequest) *NullableIdentityVerificationCreateRequest {
	return &NullableIdentityVerificationCreateRequest{value: val, isSet: true}
}

func (v NullableIdentityVerificationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


