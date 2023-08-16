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

// WalletListRequest WalletListRequest defines the request schema for `/wallet/list`
type WalletListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	IsoCurrencyCode *WalletISOCurrencyCode `json:"iso_currency_code,omitempty"`
	// A base64 value representing the latest e-wallet that has already been requested. Set this to `next_cursor` received from the previous `/wallet/list` request. If provided, the response will only contain e-wallets created before that e-wallet. If omitted, the response will contain e-wallets starting from the most recent, and in descending order.
	Cursor *string `json:"cursor,omitempty"`
	// The number of e-wallets to fetch
	Count *int32 `json:"count,omitempty"`
}

// NewWalletListRequest instantiates a new WalletListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletListRequest() *WalletListRequest {
	this := WalletListRequest{}
	var count int32 = 10
	this.Count = &count
	return &this
}

// NewWalletListRequestWithDefaults instantiates a new WalletListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletListRequestWithDefaults() *WalletListRequest {
	this := WalletListRequest{}
	var count int32 = 10
	this.Count = &count
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WalletListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WalletListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WalletListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WalletListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WalletListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WalletListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise.
func (o *WalletListRequest) GetIsoCurrencyCode() WalletISOCurrencyCode {
	if o == nil || o.IsoCurrencyCode == nil {
		var ret WalletISOCurrencyCode
		return ret
	}
	return *o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletListRequest) GetIsoCurrencyCodeOk() (*WalletISOCurrencyCode, bool) {
	if o == nil || o.IsoCurrencyCode == nil {
		return nil, false
	}
	return o.IsoCurrencyCode, true
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *WalletListRequest) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode != nil {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given WalletISOCurrencyCode and assigns it to the IsoCurrencyCode field.
func (o *WalletListRequest) SetIsoCurrencyCode(v WalletISOCurrencyCode) {
	o.IsoCurrencyCode = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *WalletListRequest) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletListRequest) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *WalletListRequest) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *WalletListRequest) SetCursor(v string) {
	o.Cursor = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *WalletListRequest) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletListRequest) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *WalletListRequest) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *WalletListRequest) SetCount(v int32) {
	o.Count = &v
}

func (o WalletListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.IsoCurrencyCode != nil {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableWalletListRequest struct {
	value *WalletListRequest
	isSet bool
}

func (v NullableWalletListRequest) Get() *WalletListRequest {
	return v.value
}

func (v *NullableWalletListRequest) Set(val *WalletListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletListRequest(val *WalletListRequest) *NullableWalletListRequest {
	return &NullableWalletListRequest{value: val, isSet: true}
}

func (v NullableWalletListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


