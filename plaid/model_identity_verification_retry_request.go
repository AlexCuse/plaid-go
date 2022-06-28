/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.131.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IdentityVerificationRetryRequest Request input for retrying an identity verification attempt
type IdentityVerificationRetryRequest struct {
	// An identifier to help you connect this object to your internal systems. For example, your database ID corresponding to this object.
	ClientUserId string `json:"client_user_id"`
	// ID of the associated Identity Verification template.
	TemplateId string `json:"template_id"`
	Strategy Strategy `json:"strategy"`
	Steps NullableIdentityVerificationRetryRequestStepsObject `json:"steps,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
}

// NewIdentityVerificationRetryRequest instantiates a new IdentityVerificationRetryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationRetryRequest(clientUserId string, templateId string, strategy Strategy) *IdentityVerificationRetryRequest {
	this := IdentityVerificationRetryRequest{}
	this.ClientUserId = clientUserId
	this.TemplateId = templateId
	this.Strategy = strategy
	return &this
}

// NewIdentityVerificationRetryRequestWithDefaults instantiates a new IdentityVerificationRetryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationRetryRequestWithDefaults() *IdentityVerificationRetryRequest {
	this := IdentityVerificationRetryRequest{}
	return &this
}

// GetClientUserId returns the ClientUserId field value
func (o *IdentityVerificationRetryRequest) GetClientUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRetryRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientUserId, true
}

// SetClientUserId sets field value
func (o *IdentityVerificationRetryRequest) SetClientUserId(v string) {
	o.ClientUserId = v
}

// GetTemplateId returns the TemplateId field value
func (o *IdentityVerificationRetryRequest) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRetryRequest) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *IdentityVerificationRetryRequest) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetStrategy returns the Strategy field value
func (o *IdentityVerificationRetryRequest) GetStrategy() Strategy {
	if o == nil {
		var ret Strategy
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRetryRequest) GetStrategyOk() (*Strategy, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *IdentityVerificationRetryRequest) SetStrategy(v Strategy) {
	o.Strategy = v
}

// GetSteps returns the Steps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationRetryRequest) GetSteps() IdentityVerificationRetryRequestStepsObject {
	if o == nil || o.Steps.Get() == nil {
		var ret IdentityVerificationRetryRequestStepsObject
		return ret
	}
	return *o.Steps.Get()
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationRetryRequest) GetStepsOk() (*IdentityVerificationRetryRequestStepsObject, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Steps.Get(), o.Steps.IsSet()
}

// HasSteps returns a boolean if a field has been set.
func (o *IdentityVerificationRetryRequest) HasSteps() bool {
	if o != nil && o.Steps.IsSet() {
		return true
	}

	return false
}

// SetSteps gets a reference to the given NullableIdentityVerificationRetryRequestStepsObject and assigns it to the Steps field.
func (o *IdentityVerificationRetryRequest) SetSteps(v IdentityVerificationRetryRequestStepsObject) {
	o.Steps.Set(&v)
}
// SetStepsNil sets the value for Steps to be an explicit nil
func (o *IdentityVerificationRetryRequest) SetStepsNil() {
	o.Steps.Set(nil)
}

// UnsetSteps ensures that no value is present for Steps, not even an explicit nil
func (o *IdentityVerificationRetryRequest) UnsetSteps() {
	o.Steps.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IdentityVerificationRetryRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRetryRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IdentityVerificationRetryRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IdentityVerificationRetryRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IdentityVerificationRetryRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRetryRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IdentityVerificationRetryRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IdentityVerificationRetryRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o IdentityVerificationRetryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if true {
		toSerialize["template_id"] = o.TemplateId
	}
	if true {
		toSerialize["strategy"] = o.Strategy
	}
	if o.Steps.IsSet() {
		toSerialize["steps"] = o.Steps.Get()
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityVerificationRetryRequest struct {
	value *IdentityVerificationRetryRequest
	isSet bool
}

func (v NullableIdentityVerificationRetryRequest) Get() *IdentityVerificationRetryRequest {
	return v.value
}

func (v *NullableIdentityVerificationRetryRequest) Set(val *IdentityVerificationRetryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationRetryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationRetryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationRetryRequest(val *IdentityVerificationRetryRequest) *NullableIdentityVerificationRetryRequest {
	return &NullableIdentityVerificationRetryRequest{value: val, isSet: true}
}

func (v NullableIdentityVerificationRetryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationRetryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


