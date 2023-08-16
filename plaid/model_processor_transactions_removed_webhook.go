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

// ProcessorTransactionsRemovedWebhook This webhook is only sent to [Plaid processor partners](https://plaid.com/docs/auth/partnerships/).  Fired when transaction(s) for an Item are deleted. The deleted transaction IDs are included in the webhook payload. Plaid will typically check for deleted transaction data several times a day.  This webhook is intended for use with `/processor/transactions/get`; if you are using the newer `/processor/transactions/sync` endpoint, this webhook will still be fired to maintain backwards compatibility, but it is recommended to listen for and respond to the `SYNC_UPDATES_AVAILABLE` webhook instead.
type ProcessorTransactionsRemovedWebhook struct {
	// `TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `TRANSACTIONS_REMOVED`
	WebhookCode string `json:"webhook_code"`
	Error NullablePlaidError `json:"error,omitempty"`
	// An array of `transaction_ids` corresponding to the removed transactions
	RemovedTransactions []string `json:"removed_transactions"`
	// The ID of the account.
	AccountId string `json:"account_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorTransactionsRemovedWebhook ProcessorTransactionsRemovedWebhook

// NewProcessorTransactionsRemovedWebhook instantiates a new ProcessorTransactionsRemovedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorTransactionsRemovedWebhook(webhookType string, webhookCode string, removedTransactions []string, accountId string, environment WebhookEnvironmentValues) *ProcessorTransactionsRemovedWebhook {
	this := ProcessorTransactionsRemovedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.RemovedTransactions = removedTransactions
	this.AccountId = accountId
	this.Environment = environment
	return &this
}

// NewProcessorTransactionsRemovedWebhookWithDefaults instantiates a new ProcessorTransactionsRemovedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorTransactionsRemovedWebhookWithDefaults() *ProcessorTransactionsRemovedWebhook {
	this := ProcessorTransactionsRemovedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *ProcessorTransactionsRemovedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsRemovedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *ProcessorTransactionsRemovedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *ProcessorTransactionsRemovedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsRemovedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *ProcessorTransactionsRemovedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorTransactionsRemovedWebhook) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorTransactionsRemovedWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ProcessorTransactionsRemovedWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullablePlaidError and assigns it to the Error field.
func (o *ProcessorTransactionsRemovedWebhook) SetError(v PlaidError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ProcessorTransactionsRemovedWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ProcessorTransactionsRemovedWebhook) UnsetError() {
	o.Error.Unset()
}

// GetRemovedTransactions returns the RemovedTransactions field value
func (o *ProcessorTransactionsRemovedWebhook) GetRemovedTransactions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RemovedTransactions
}

// GetRemovedTransactionsOk returns a tuple with the RemovedTransactions field value
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsRemovedWebhook) GetRemovedTransactionsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RemovedTransactions, true
}

// SetRemovedTransactions sets field value
func (o *ProcessorTransactionsRemovedWebhook) SetRemovedTransactions(v []string) {
	o.RemovedTransactions = v
}

// GetAccountId returns the AccountId field value
func (o *ProcessorTransactionsRemovedWebhook) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsRemovedWebhook) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ProcessorTransactionsRemovedWebhook) SetAccountId(v string) {
	o.AccountId = v
}

// GetEnvironment returns the Environment field value
func (o *ProcessorTransactionsRemovedWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsRemovedWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ProcessorTransactionsRemovedWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o ProcessorTransactionsRemovedWebhook) MarshalJSON() ([]byte, error) {
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
		toSerialize["removed_transactions"] = o.RemovedTransactions
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorTransactionsRemovedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorTransactionsRemovedWebhook := _ProcessorTransactionsRemovedWebhook{}

	if err = json.Unmarshal(bytes, &varProcessorTransactionsRemovedWebhook); err == nil {
		*o = ProcessorTransactionsRemovedWebhook(varProcessorTransactionsRemovedWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "error")
		delete(additionalProperties, "removed_transactions")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorTransactionsRemovedWebhook struct {
	value *ProcessorTransactionsRemovedWebhook
	isSet bool
}

func (v NullableProcessorTransactionsRemovedWebhook) Get() *ProcessorTransactionsRemovedWebhook {
	return v.value
}

func (v *NullableProcessorTransactionsRemovedWebhook) Set(val *ProcessorTransactionsRemovedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorTransactionsRemovedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorTransactionsRemovedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorTransactionsRemovedWebhook(val *ProcessorTransactionsRemovedWebhook) *NullableProcessorTransactionsRemovedWebhook {
	return &NullableProcessorTransactionsRemovedWebhook{value: val, isSet: true}
}

func (v NullableProcessorTransactionsRemovedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorTransactionsRemovedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


