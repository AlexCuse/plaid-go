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

// TransactionsSyncRequest TransactionsSyncRequest defines the request schema for `/transactions/sync`
type TransactionsSyncRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The cursor value represents the last update requested. Providing it will cause the response to only return changes after this update. If omitted, the entire history of updates will be returned, starting with the first-added transactions on the item. Note: The upper-bound length of this cursor is 256 characters of base64.
	Cursor *string `json:"cursor,omitempty"`
	// The number of transaction updates to fetch.
	Count *int32 `json:"count,omitempty"`
	Options *TransactionsSyncRequestOptions `json:"options,omitempty"`
}

// NewTransactionsSyncRequest instantiates a new TransactionsSyncRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsSyncRequest(accessToken string) *TransactionsSyncRequest {
	this := TransactionsSyncRequest{}
	this.AccessToken = accessToken
	var count int32 = 100
	this.Count = &count
	return &this
}

// NewTransactionsSyncRequestWithDefaults instantiates a new TransactionsSyncRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsSyncRequestWithDefaults() *TransactionsSyncRequest {
	this := TransactionsSyncRequest{}
	var count int32 = 100
	this.Count = &count
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransactionsSyncRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsSyncRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransactionsSyncRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransactionsSyncRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetAccessToken returns the AccessToken field value
func (o *TransactionsSyncRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TransactionsSyncRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TransactionsSyncRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransactionsSyncRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsSyncRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransactionsSyncRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransactionsSyncRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *TransactionsSyncRequest) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsSyncRequest) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *TransactionsSyncRequest) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *TransactionsSyncRequest) SetCursor(v string) {
	o.Cursor = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TransactionsSyncRequest) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsSyncRequest) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TransactionsSyncRequest) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *TransactionsSyncRequest) SetCount(v int32) {
	o.Count = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TransactionsSyncRequest) GetOptions() TransactionsSyncRequestOptions {
	if o == nil || o.Options == nil {
		var ret TransactionsSyncRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsSyncRequest) GetOptionsOk() (*TransactionsSyncRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TransactionsSyncRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given TransactionsSyncRequestOptions and assigns it to the Options field.
func (o *TransactionsSyncRequest) SetOptions(v TransactionsSyncRequestOptions) {
	o.Options = &v
}

func (o TransactionsSyncRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsSyncRequest struct {
	value *TransactionsSyncRequest
	isSet bool
}

func (v NullableTransactionsSyncRequest) Get() *TransactionsSyncRequest {
	return v.value
}

func (v *NullableTransactionsSyncRequest) Set(val *TransactionsSyncRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsSyncRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsSyncRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsSyncRequest(val *TransactionsSyncRequest) *NullableTransactionsSyncRequest {
	return &NullableTransactionsSyncRequest{value: val, isSet: true}
}

func (v NullableTransactionsSyncRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsSyncRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


