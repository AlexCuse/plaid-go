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

// ItemImportRequestOptions An optional object to configure `/item/import` request.
type ItemImportRequestOptions struct {
	// Specifies a webhook URL to associate with an Item. Plaid fires a webhook if credentials fail. 
	Webhook *string `json:"webhook,omitempty"`
}

// NewItemImportRequestOptions instantiates a new ItemImportRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemImportRequestOptions() *ItemImportRequestOptions {
	this := ItemImportRequestOptions{}
	return &this
}

// NewItemImportRequestOptionsWithDefaults instantiates a new ItemImportRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemImportRequestOptionsWithDefaults() *ItemImportRequestOptions {
	this := ItemImportRequestOptions{}
	return &this
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *ItemImportRequestOptions) GetWebhook() string {
	if o == nil || o.Webhook == nil {
		var ret string
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemImportRequestOptions) GetWebhookOk() (*string, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *ItemImportRequestOptions) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given string and assigns it to the Webhook field.
func (o *ItemImportRequestOptions) SetWebhook(v string) {
	o.Webhook = &v
}

func (o ItemImportRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableItemImportRequestOptions struct {
	value *ItemImportRequestOptions
	isSet bool
}

func (v NullableItemImportRequestOptions) Get() *ItemImportRequestOptions {
	return v.value
}

func (v *NullableItemImportRequestOptions) Set(val *ItemImportRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableItemImportRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableItemImportRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemImportRequestOptions(val *ItemImportRequestOptions) *NullableItemImportRequestOptions {
	return &NullableItemImportRequestOptions{value: val, isSet: true}
}

func (v NullableItemImportRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemImportRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


