/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.61.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// WalletGetResponse WalletGetResponse defines the response schema for `/wallet/get`
type WalletGetResponse struct {
	// A unique ID identifying the e-wallet
	WalletId string `json:"wallet_id"`
	Balance NullableWalletBalance `json:"balance"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WalletGetResponse WalletGetResponse

// NewWalletGetResponse instantiates a new WalletGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletGetResponse(walletId string, balance NullableWalletBalance, requestId string) *WalletGetResponse {
	this := WalletGetResponse{}
	this.WalletId = walletId
	this.Balance = balance
	this.RequestId = requestId
	return &this
}

// NewWalletGetResponseWithDefaults instantiates a new WalletGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletGetResponseWithDefaults() *WalletGetResponse {
	this := WalletGetResponse{}
	return &this
}

// GetWalletId returns the WalletId field value
func (o *WalletGetResponse) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *WalletGetResponse) GetWalletIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *WalletGetResponse) SetWalletId(v string) {
	o.WalletId = v
}

// GetBalance returns the Balance field value
// If the value is explicit nil, the zero value for WalletBalance will be returned
func (o *WalletGetResponse) GetBalance() WalletBalance {
	if o == nil || o.Balance.Get() == nil {
		var ret WalletBalance
		return ret
	}

	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletGetResponse) GetBalanceOk() (*WalletBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// SetBalance sets field value
func (o *WalletGetResponse) SetBalance(v WalletBalance) {
	o.Balance.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *WalletGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WalletGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WalletGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WalletGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wallet_id"] = o.WalletId
	}
	if true {
		toSerialize["balance"] = o.Balance.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWalletGetResponse := _WalletGetResponse{}

	if err = json.Unmarshal(bytes, &varWalletGetResponse); err == nil {
		*o = WalletGetResponse(varWalletGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "wallet_id")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletGetResponse struct {
	value *WalletGetResponse
	isSet bool
}

func (v NullableWalletGetResponse) Get() *WalletGetResponse {
	return v.value
}

func (v *NullableWalletGetResponse) Set(val *WalletGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletGetResponse(val *WalletGetResponse) *NullableWalletGetResponse {
	return &NullableWalletGetResponse{value: val, isSet: true}
}

func (v NullableWalletGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


