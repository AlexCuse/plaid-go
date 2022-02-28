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

// NewAccountsAvailableWebhook Fired when Plaid detects a new account for Items created or updated with [Account Select v2](https://plaid.com/docs/link/customization/#account-select). Upon receiving this webhook, you can prompt your users to share new accounts with you through [Account Select v2 update mode](https://plaid.com/docs/link/update-mode/#using-update-mode-to-request-new-accounts).
type NewAccountsAvailableWebhook struct {
	// `ITEM`
	WebhookType *string `json:"webhook_type,omitempty"`
	// `NEW_ACCOUNTS_AVAILABLE`
	WebhookCode *string `json:"webhook_code,omitempty"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId *string `json:"item_id,omitempty"`
	Error *PlaidError `json:"error,omitempty"`
}

// NewNewAccountsAvailableWebhook instantiates a new NewAccountsAvailableWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewAccountsAvailableWebhook() *NewAccountsAvailableWebhook {
	this := NewAccountsAvailableWebhook{}
	return &this
}

// NewNewAccountsAvailableWebhookWithDefaults instantiates a new NewAccountsAvailableWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewAccountsAvailableWebhookWithDefaults() *NewAccountsAvailableWebhook {
	this := NewAccountsAvailableWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value if set, zero value otherwise.
func (o *NewAccountsAvailableWebhook) GetWebhookType() string {
	if o == nil || o.WebhookType == nil {
		var ret string
		return ret
	}
	return *o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountsAvailableWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil || o.WebhookType == nil {
		return nil, false
	}
	return o.WebhookType, true
}

// HasWebhookType returns a boolean if a field has been set.
func (o *NewAccountsAvailableWebhook) HasWebhookType() bool {
	if o != nil && o.WebhookType != nil {
		return true
	}

	return false
}

// SetWebhookType gets a reference to the given string and assigns it to the WebhookType field.
func (o *NewAccountsAvailableWebhook) SetWebhookType(v string) {
	o.WebhookType = &v
}

// GetWebhookCode returns the WebhookCode field value if set, zero value otherwise.
func (o *NewAccountsAvailableWebhook) GetWebhookCode() string {
	if o == nil || o.WebhookCode == nil {
		var ret string
		return ret
	}
	return *o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountsAvailableWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil || o.WebhookCode == nil {
		return nil, false
	}
	return o.WebhookCode, true
}

// HasWebhookCode returns a boolean if a field has been set.
func (o *NewAccountsAvailableWebhook) HasWebhookCode() bool {
	if o != nil && o.WebhookCode != nil {
		return true
	}

	return false
}

// SetWebhookCode gets a reference to the given string and assigns it to the WebhookCode field.
func (o *NewAccountsAvailableWebhook) SetWebhookCode(v string) {
	o.WebhookCode = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *NewAccountsAvailableWebhook) GetItemId() string {
	if o == nil || o.ItemId == nil {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountsAvailableWebhook) GetItemIdOk() (*string, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *NewAccountsAvailableWebhook) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *NewAccountsAvailableWebhook) SetItemId(v string) {
	o.ItemId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *NewAccountsAvailableWebhook) GetError() PlaidError {
	if o == nil || o.Error == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountsAvailableWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *NewAccountsAvailableWebhook) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given PlaidError and assigns it to the Error field.
func (o *NewAccountsAvailableWebhook) SetError(v PlaidError) {
	o.Error = &v
}

func (o NewAccountsAvailableWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WebhookType != nil {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if o.WebhookCode != nil {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if o.ItemId != nil {
		toSerialize["item_id"] = o.ItemId
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableNewAccountsAvailableWebhook struct {
	value *NewAccountsAvailableWebhook
	isSet bool
}

func (v NullableNewAccountsAvailableWebhook) Get() *NewAccountsAvailableWebhook {
	return v.value
}

func (v *NullableNewAccountsAvailableWebhook) Set(val *NewAccountsAvailableWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAccountsAvailableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAccountsAvailableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAccountsAvailableWebhook(val *NewAccountsAvailableWebhook) *NullableNewAccountsAvailableWebhook {
	return &NullableNewAccountsAvailableWebhook{value: val, isSet: true}
}

func (v NullableNewAccountsAvailableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAccountsAvailableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


