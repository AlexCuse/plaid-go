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

// InitialUpdateWebhook Fired when an Item's initial transaction pull is completed. Once this webhook has been fired, transaction data for the most recent 30 days can be fetched for the Item. If [Account Select v2](https://plaid.com/docs/link/customization/#account-select) is enabled, this webhook will also be fired if account selections for the Item are updated, with `new_transactions` set to the number of net new transactions pulled after the account selection update.
type InitialUpdateWebhook struct {
	// `TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `INITIAL_UPDATE`
	WebhookCode string `json:"webhook_code"`
	// The error code associated with the webhook.
	Error NullableString `json:"error,omitempty"`
	// The number of new, unfetched transactions available.
	NewTransactions float32 `json:"new_transactions"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	AdditionalProperties map[string]interface{}
}

type _InitialUpdateWebhook InitialUpdateWebhook

// NewInitialUpdateWebhook instantiates a new InitialUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitialUpdateWebhook(webhookType string, webhookCode string, newTransactions float32, itemId string) *InitialUpdateWebhook {
	this := InitialUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.NewTransactions = newTransactions
	this.ItemId = itemId
	return &this
}

// NewInitialUpdateWebhookWithDefaults instantiates a new InitialUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitialUpdateWebhookWithDefaults() *InitialUpdateWebhook {
	this := InitialUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *InitialUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *InitialUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *InitialUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *InitialUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *InitialUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *InitialUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InitialUpdateWebhook) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InitialUpdateWebhook) GetErrorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *InitialUpdateWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *InitialUpdateWebhook) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *InitialUpdateWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *InitialUpdateWebhook) UnsetError() {
	o.Error.Unset()
}

// GetNewTransactions returns the NewTransactions field value
func (o *InitialUpdateWebhook) GetNewTransactions() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NewTransactions
}

// GetNewTransactionsOk returns a tuple with the NewTransactions field value
// and a boolean to check if the value has been set.
func (o *InitialUpdateWebhook) GetNewTransactionsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewTransactions, true
}

// SetNewTransactions sets field value
func (o *InitialUpdateWebhook) SetNewTransactions(v float32) {
	o.NewTransactions = v
}

// GetItemId returns the ItemId field value
func (o *InitialUpdateWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *InitialUpdateWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *InitialUpdateWebhook) SetItemId(v string) {
	o.ItemId = v
}

func (o InitialUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["new_transactions"] = o.NewTransactions
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InitialUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varInitialUpdateWebhook := _InitialUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varInitialUpdateWebhook); err == nil {
		*o = InitialUpdateWebhook(varInitialUpdateWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "error")
		delete(additionalProperties, "new_transactions")
		delete(additionalProperties, "item_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInitialUpdateWebhook struct {
	value *InitialUpdateWebhook
	isSet bool
}

func (v NullableInitialUpdateWebhook) Get() *InitialUpdateWebhook {
	return v.value
}

func (v *NullableInitialUpdateWebhook) Set(val *InitialUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableInitialUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableInitialUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitialUpdateWebhook(val *InitialUpdateWebhook) *NullableInitialUpdateWebhook {
	return &NullableInitialUpdateWebhook{value: val, isSet: true}
}

func (v NullableInitialUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitialUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


