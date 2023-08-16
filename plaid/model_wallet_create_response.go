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

// WalletCreateResponse WalletCreateResponse defines the response schema for `/wallet/create`
type WalletCreateResponse struct {
	// A unique ID identifying the e-wallet
	WalletId string `json:"wallet_id"`
	Balance WalletBalance `json:"balance"`
	Numbers WalletNumbers `json:"numbers"`
	// The ID of the recipient that corresponds to the e-wallet account numbers
	RecipientId *string `json:"recipient_id,omitempty"`
	Status WalletStatus `json:"status"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WalletCreateResponse WalletCreateResponse

// NewWalletCreateResponse instantiates a new WalletCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletCreateResponse(walletId string, balance WalletBalance, numbers WalletNumbers, status WalletStatus, requestId string) *WalletCreateResponse {
	this := WalletCreateResponse{}
	this.WalletId = walletId
	this.Balance = balance
	this.Numbers = numbers
	this.Status = status
	this.RequestId = requestId
	return &this
}

// NewWalletCreateResponseWithDefaults instantiates a new WalletCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletCreateResponseWithDefaults() *WalletCreateResponse {
	this := WalletCreateResponse{}
	return &this
}

// GetWalletId returns the WalletId field value
func (o *WalletCreateResponse) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *WalletCreateResponse) GetWalletIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *WalletCreateResponse) SetWalletId(v string) {
	o.WalletId = v
}

// GetBalance returns the Balance field value
func (o *WalletCreateResponse) GetBalance() WalletBalance {
	if o == nil {
		var ret WalletBalance
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *WalletCreateResponse) GetBalanceOk() (*WalletBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *WalletCreateResponse) SetBalance(v WalletBalance) {
	o.Balance = v
}

// GetNumbers returns the Numbers field value
func (o *WalletCreateResponse) GetNumbers() WalletNumbers {
	if o == nil {
		var ret WalletNumbers
		return ret
	}

	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value
// and a boolean to check if the value has been set.
func (o *WalletCreateResponse) GetNumbersOk() (*WalletNumbers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Numbers, true
}

// SetNumbers sets field value
func (o *WalletCreateResponse) SetNumbers(v WalletNumbers) {
	o.Numbers = v
}

// GetRecipientId returns the RecipientId field value if set, zero value otherwise.
func (o *WalletCreateResponse) GetRecipientId() string {
	if o == nil || o.RecipientId == nil {
		var ret string
		return ret
	}
	return *o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreateResponse) GetRecipientIdOk() (*string, bool) {
	if o == nil || o.RecipientId == nil {
		return nil, false
	}
	return o.RecipientId, true
}

// HasRecipientId returns a boolean if a field has been set.
func (o *WalletCreateResponse) HasRecipientId() bool {
	if o != nil && o.RecipientId != nil {
		return true
	}

	return false
}

// SetRecipientId gets a reference to the given string and assigns it to the RecipientId field.
func (o *WalletCreateResponse) SetRecipientId(v string) {
	o.RecipientId = &v
}

// GetStatus returns the Status field value
func (o *WalletCreateResponse) GetStatus() WalletStatus {
	if o == nil {
		var ret WalletStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WalletCreateResponse) GetStatusOk() (*WalletStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WalletCreateResponse) SetStatus(v WalletStatus) {
	o.Status = v
}

// GetRequestId returns the RequestId field value
func (o *WalletCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WalletCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WalletCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WalletCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wallet_id"] = o.WalletId
	}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["numbers"] = o.Numbers
	}
	if o.RecipientId != nil {
		toSerialize["recipient_id"] = o.RecipientId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWalletCreateResponse := _WalletCreateResponse{}

	if err = json.Unmarshal(bytes, &varWalletCreateResponse); err == nil {
		*o = WalletCreateResponse(varWalletCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "wallet_id")
		delete(additionalProperties, "balance")
		delete(additionalProperties, "numbers")
		delete(additionalProperties, "recipient_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletCreateResponse struct {
	value *WalletCreateResponse
	isSet bool
}

func (v NullableWalletCreateResponse) Get() *WalletCreateResponse {
	return v.value
}

func (v *NullableWalletCreateResponse) Set(val *WalletCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletCreateResponse(val *WalletCreateResponse) *NullableWalletCreateResponse {
	return &NullableWalletCreateResponse{value: val, isSet: true}
}

func (v NullableWalletCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


