/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.77.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IdentityDefaultUpdateWebhook Fired when a change to identity data has been detected on an Item.
type IdentityDefaultUpdateWebhook struct {
	// `IDENTITY`
	WebhookType string `json:"webhook_type"`
	// `DEFAULT_UPDATE`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	// An object with keys of `account_id`'s that are mapped to their respective identity attributes that changed.  Example: `{ \"XMBvvyMGQ1UoLbKByoMqH3nXMj84ALSdE5B58\": [\"PHONES\"] }` 
	AccountIdsWithUpdatedIdentity map[string][]IdentityUpdateTypes `json:"account_ids_with_updated_identity"`
	Error PlaidError `json:"error"`
	AdditionalProperties map[string]interface{}
}

type _IdentityDefaultUpdateWebhook IdentityDefaultUpdateWebhook

// NewIdentityDefaultUpdateWebhook instantiates a new IdentityDefaultUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityDefaultUpdateWebhook(webhookType string, webhookCode string, itemId string, accountIdsWithUpdatedIdentity map[string][]IdentityUpdateTypes, error_ PlaidError) *IdentityDefaultUpdateWebhook {
	this := IdentityDefaultUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.AccountIdsWithUpdatedIdentity = accountIdsWithUpdatedIdentity
	this.Error = error_
	return &this
}

// NewIdentityDefaultUpdateWebhookWithDefaults instantiates a new IdentityDefaultUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityDefaultUpdateWebhookWithDefaults() *IdentityDefaultUpdateWebhook {
	this := IdentityDefaultUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *IdentityDefaultUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *IdentityDefaultUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *IdentityDefaultUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *IdentityDefaultUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *IdentityDefaultUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *IdentityDefaultUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *IdentityDefaultUpdateWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *IdentityDefaultUpdateWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *IdentityDefaultUpdateWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetAccountIdsWithUpdatedIdentity returns the AccountIdsWithUpdatedIdentity field value
func (o *IdentityDefaultUpdateWebhook) GetAccountIdsWithUpdatedIdentity() map[string][]IdentityUpdateTypes {
	if o == nil {
		var ret map[string][]IdentityUpdateTypes
		return ret
	}

	return o.AccountIdsWithUpdatedIdentity
}

// GetAccountIdsWithUpdatedIdentityOk returns a tuple with the AccountIdsWithUpdatedIdentity field value
// and a boolean to check if the value has been set.
func (o *IdentityDefaultUpdateWebhook) GetAccountIdsWithUpdatedIdentityOk() (*map[string][]IdentityUpdateTypes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountIdsWithUpdatedIdentity, true
}

// SetAccountIdsWithUpdatedIdentity sets field value
func (o *IdentityDefaultUpdateWebhook) SetAccountIdsWithUpdatedIdentity(v map[string][]IdentityUpdateTypes) {
	o.AccountIdsWithUpdatedIdentity = v
}

// GetError returns the Error field value
func (o *IdentityDefaultUpdateWebhook) GetError() PlaidError {
	if o == nil {
		var ret PlaidError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *IdentityDefaultUpdateWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *IdentityDefaultUpdateWebhook) SetError(v PlaidError) {
	o.Error = v
}

func (o IdentityDefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["account_ids_with_updated_identity"] = o.AccountIdsWithUpdatedIdentity
	}
	if true {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityDefaultUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityDefaultUpdateWebhook := _IdentityDefaultUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varIdentityDefaultUpdateWebhook); err == nil {
		*o = IdentityDefaultUpdateWebhook(varIdentityDefaultUpdateWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "account_ids_with_updated_identity")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityDefaultUpdateWebhook struct {
	value *IdentityDefaultUpdateWebhook
	isSet bool
}

func (v NullableIdentityDefaultUpdateWebhook) Get() *IdentityDefaultUpdateWebhook {
	return v.value
}

func (v *NullableIdentityDefaultUpdateWebhook) Set(val *IdentityDefaultUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityDefaultUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityDefaultUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityDefaultUpdateWebhook(val *IdentityDefaultUpdateWebhook) *NullableIdentityDefaultUpdateWebhook {
	return &NullableIdentityDefaultUpdateWebhook{value: val, isSet: true}
}

func (v NullableIdentityDefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityDefaultUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


