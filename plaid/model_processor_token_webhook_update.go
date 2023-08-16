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

// ProcessorTokenWebhookUpdate This webhook is only sent to [Plaid processor partners](https://plaid.com/docs/auth/partnerships/).  Fired when a processor updates the webhook URL for a processor token via `/processor/token/webhook/update`.
type ProcessorTokenWebhookUpdate struct {
	// `PROCESSOR_TOKEN`
	WebhookType string `json:"webhook_type"`
	// `WEBHOOK_UPDATE_ACKNOWLEDGED`
	WebhookCode string `json:"webhook_code"`
	Error NullablePlaidError `json:"error,omitempty"`
	// The ID of the account.
	AccountId string `json:"account_id"`
	// The new webhook URL.
	NewWebhookUrl string `json:"new_webhook_url"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorTokenWebhookUpdate ProcessorTokenWebhookUpdate

// NewProcessorTokenWebhookUpdate instantiates a new ProcessorTokenWebhookUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorTokenWebhookUpdate(webhookType string, webhookCode string, accountId string, newWebhookUrl string, environment WebhookEnvironmentValues) *ProcessorTokenWebhookUpdate {
	this := ProcessorTokenWebhookUpdate{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.AccountId = accountId
	this.NewWebhookUrl = newWebhookUrl
	this.Environment = environment
	return &this
}

// NewProcessorTokenWebhookUpdateWithDefaults instantiates a new ProcessorTokenWebhookUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorTokenWebhookUpdateWithDefaults() *ProcessorTokenWebhookUpdate {
	this := ProcessorTokenWebhookUpdate{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *ProcessorTokenWebhookUpdate) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenWebhookUpdate) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *ProcessorTokenWebhookUpdate) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *ProcessorTokenWebhookUpdate) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenWebhookUpdate) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *ProcessorTokenWebhookUpdate) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorTokenWebhookUpdate) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorTokenWebhookUpdate) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ProcessorTokenWebhookUpdate) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullablePlaidError and assigns it to the Error field.
func (o *ProcessorTokenWebhookUpdate) SetError(v PlaidError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ProcessorTokenWebhookUpdate) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ProcessorTokenWebhookUpdate) UnsetError() {
	o.Error.Unset()
}

// GetAccountId returns the AccountId field value
func (o *ProcessorTokenWebhookUpdate) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenWebhookUpdate) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ProcessorTokenWebhookUpdate) SetAccountId(v string) {
	o.AccountId = v
}

// GetNewWebhookUrl returns the NewWebhookUrl field value
func (o *ProcessorTokenWebhookUpdate) GetNewWebhookUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewWebhookUrl
}

// GetNewWebhookUrlOk returns a tuple with the NewWebhookUrl field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenWebhookUpdate) GetNewWebhookUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewWebhookUrl, true
}

// SetNewWebhookUrl sets field value
func (o *ProcessorTokenWebhookUpdate) SetNewWebhookUrl(v string) {
	o.NewWebhookUrl = v
}

// GetEnvironment returns the Environment field value
func (o *ProcessorTokenWebhookUpdate) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenWebhookUpdate) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ProcessorTokenWebhookUpdate) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o ProcessorTokenWebhookUpdate) MarshalJSON() ([]byte, error) {
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
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["new_webhook_url"] = o.NewWebhookUrl
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorTokenWebhookUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorTokenWebhookUpdate := _ProcessorTokenWebhookUpdate{}

	if err = json.Unmarshal(bytes, &varProcessorTokenWebhookUpdate); err == nil {
		*o = ProcessorTokenWebhookUpdate(varProcessorTokenWebhookUpdate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "error")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "new_webhook_url")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorTokenWebhookUpdate struct {
	value *ProcessorTokenWebhookUpdate
	isSet bool
}

func (v NullableProcessorTokenWebhookUpdate) Get() *ProcessorTokenWebhookUpdate {
	return v.value
}

func (v *NullableProcessorTokenWebhookUpdate) Set(val *ProcessorTokenWebhookUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorTokenWebhookUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorTokenWebhookUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorTokenWebhookUpdate(val *ProcessorTokenWebhookUpdate) *NullableProcessorTokenWebhookUpdate {
	return &NullableProcessorTokenWebhookUpdate{value: val, isSet: true}
}

func (v NullableProcessorTokenWebhookUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorTokenWebhookUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


