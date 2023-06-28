/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.385.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LinkDeliveryOptions Optional metadata related to the Hosted Link session
type LinkDeliveryOptions struct {
	Recipient *LinkDeliveryRecipient `json:"recipient,omitempty"`
}

// NewLinkDeliveryOptions instantiates a new LinkDeliveryOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDeliveryOptions() *LinkDeliveryOptions {
	this := LinkDeliveryOptions{}
	return &this
}

// NewLinkDeliveryOptionsWithDefaults instantiates a new LinkDeliveryOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDeliveryOptionsWithDefaults() *LinkDeliveryOptions {
	this := LinkDeliveryOptions{}
	return &this
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *LinkDeliveryOptions) GetRecipient() LinkDeliveryRecipient {
	if o == nil || o.Recipient == nil {
		var ret LinkDeliveryRecipient
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryOptions) GetRecipientOk() (*LinkDeliveryRecipient, bool) {
	if o == nil || o.Recipient == nil {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *LinkDeliveryOptions) HasRecipient() bool {
	if o != nil && o.Recipient != nil {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given LinkDeliveryRecipient and assigns it to the Recipient field.
func (o *LinkDeliveryOptions) SetRecipient(v LinkDeliveryRecipient) {
	o.Recipient = &v
}

func (o LinkDeliveryOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Recipient != nil {
		toSerialize["recipient"] = o.Recipient
	}
	return json.Marshal(toSerialize)
}

type NullableLinkDeliveryOptions struct {
	value *LinkDeliveryOptions
	isSet bool
}

func (v NullableLinkDeliveryOptions) Get() *LinkDeliveryOptions {
	return v.value
}

func (v *NullableLinkDeliveryOptions) Set(val *LinkDeliveryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryOptions(val *LinkDeliveryOptions) *NullableLinkDeliveryOptions {
	return &NullableLinkDeliveryOptions{value: val, isSet: true}
}

func (v NullableLinkDeliveryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


