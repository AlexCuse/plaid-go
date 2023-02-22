/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LinkTokenCreateRequestTransfer Specifies options for initializing Link for use with the Transfer product.
type LinkTokenCreateRequestTransfer struct {
	// The `id` returned by the `/transfer/intent/create` endpoint.
	IntentId *string `json:"intent_id,omitempty"`
	// The `payment_profile_token` returned by the `/payment_profile/create` endpoint.
	PaymentProfileToken *string `json:"payment_profile_token,omitempty"`
}

// NewLinkTokenCreateRequestTransfer instantiates a new LinkTokenCreateRequestTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestTransfer() *LinkTokenCreateRequestTransfer {
	this := LinkTokenCreateRequestTransfer{}
	return &this
}

// NewLinkTokenCreateRequestTransferWithDefaults instantiates a new LinkTokenCreateRequestTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestTransferWithDefaults() *LinkTokenCreateRequestTransfer {
	this := LinkTokenCreateRequestTransfer{}
	return &this
}

// GetIntentId returns the IntentId field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestTransfer) GetIntentId() string {
	if o == nil || o.IntentId == nil {
		var ret string
		return ret
	}
	return *o.IntentId
}

// GetIntentIdOk returns a tuple with the IntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestTransfer) GetIntentIdOk() (*string, bool) {
	if o == nil || o.IntentId == nil {
		return nil, false
	}
	return o.IntentId, true
}

// HasIntentId returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestTransfer) HasIntentId() bool {
	if o != nil && o.IntentId != nil {
		return true
	}

	return false
}

// SetIntentId gets a reference to the given string and assigns it to the IntentId field.
func (o *LinkTokenCreateRequestTransfer) SetIntentId(v string) {
	o.IntentId = &v
}

// GetPaymentProfileToken returns the PaymentProfileToken field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestTransfer) GetPaymentProfileToken() string {
	if o == nil || o.PaymentProfileToken == nil {
		var ret string
		return ret
	}
	return *o.PaymentProfileToken
}

// GetPaymentProfileTokenOk returns a tuple with the PaymentProfileToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestTransfer) GetPaymentProfileTokenOk() (*string, bool) {
	if o == nil || o.PaymentProfileToken == nil {
		return nil, false
	}
	return o.PaymentProfileToken, true
}

// HasPaymentProfileToken returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestTransfer) HasPaymentProfileToken() bool {
	if o != nil && o.PaymentProfileToken != nil {
		return true
	}

	return false
}

// SetPaymentProfileToken gets a reference to the given string and assigns it to the PaymentProfileToken field.
func (o *LinkTokenCreateRequestTransfer) SetPaymentProfileToken(v string) {
	o.PaymentProfileToken = &v
}

func (o LinkTokenCreateRequestTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IntentId != nil {
		toSerialize["intent_id"] = o.IntentId
	}
	if o.PaymentProfileToken != nil {
		toSerialize["payment_profile_token"] = o.PaymentProfileToken
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestTransfer struct {
	value *LinkTokenCreateRequestTransfer
	isSet bool
}

func (v NullableLinkTokenCreateRequestTransfer) Get() *LinkTokenCreateRequestTransfer {
	return v.value
}

func (v *NullableLinkTokenCreateRequestTransfer) Set(val *LinkTokenCreateRequestTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestTransfer(val *LinkTokenCreateRequestTransfer) *NullableLinkTokenCreateRequestTransfer {
	return &NullableLinkTokenCreateRequestTransfer{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


