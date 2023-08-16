/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.413.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ProcessorTokenCreateRequest ProcessorTokenCreateRequest defines the request schema for `/processor/token/create`
type ProcessorTokenCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// The `account_id` value obtained from the `onSuccess` callback in Link
	AccountId string `json:"account_id"`
	// The processor you are integrating with.
	Processor string `json:"processor"`
}

// NewProcessorTokenCreateRequest instantiates a new ProcessorTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorTokenCreateRequest(accessToken string, accountId string, processor string) *ProcessorTokenCreateRequest {
	this := ProcessorTokenCreateRequest{}
	this.AccessToken = accessToken
	this.AccountId = accountId
	this.Processor = processor
	return &this
}

// NewProcessorTokenCreateRequestWithDefaults instantiates a new ProcessorTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorTokenCreateRequestWithDefaults() *ProcessorTokenCreateRequest {
	this := ProcessorTokenCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorTokenCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorTokenCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorTokenCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorTokenCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorTokenCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorTokenCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorTokenCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorTokenCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *ProcessorTokenCreateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ProcessorTokenCreateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetAccountId returns the AccountId field value
func (o *ProcessorTokenCreateRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ProcessorTokenCreateRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetProcessor returns the Processor field value
func (o *ProcessorTokenCreateRequest) GetProcessor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Processor
}

// GetProcessorOk returns a tuple with the Processor field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenCreateRequest) GetProcessorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Processor, true
}

// SetProcessor sets field value
func (o *ProcessorTokenCreateRequest) SetProcessor(v string) {
	o.Processor = v
}

func (o ProcessorTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["processor"] = o.Processor
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorTokenCreateRequest struct {
	value *ProcessorTokenCreateRequest
	isSet bool
}

func (v NullableProcessorTokenCreateRequest) Get() *ProcessorTokenCreateRequest {
	return v.value
}

func (v *NullableProcessorTokenCreateRequest) Set(val *ProcessorTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorTokenCreateRequest(val *ProcessorTokenCreateRequest) *NullableProcessorTokenCreateRequest {
	return &NullableProcessorTokenCreateRequest{value: val, isSet: true}
}

func (v NullableProcessorTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


