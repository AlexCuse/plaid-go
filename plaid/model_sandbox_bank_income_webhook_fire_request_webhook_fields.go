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

// SandboxBankIncomeWebhookFireRequestWebhookFields Optional fields which will be populated in the simulated webhook
type SandboxBankIncomeWebhookFireRequestWebhookFields struct {
	// The user id to be returned in INCOME webhooks
	UserId string `json:"user_id"`
	BankIncomeRefreshCompleteResult *BankIncomeRefreshCompleteResult `json:"bank_income_refresh_complete_result,omitempty"`
}

// NewSandboxBankIncomeWebhookFireRequestWebhookFields instantiates a new SandboxBankIncomeWebhookFireRequestWebhookFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxBankIncomeWebhookFireRequestWebhookFields(userId string) *SandboxBankIncomeWebhookFireRequestWebhookFields {
	this := SandboxBankIncomeWebhookFireRequestWebhookFields{}
	this.UserId = userId
	return &this
}

// NewSandboxBankIncomeWebhookFireRequestWebhookFieldsWithDefaults instantiates a new SandboxBankIncomeWebhookFireRequestWebhookFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxBankIncomeWebhookFireRequestWebhookFieldsWithDefaults() *SandboxBankIncomeWebhookFireRequestWebhookFields {
	this := SandboxBankIncomeWebhookFireRequestWebhookFields{}
	return &this
}

// GetUserId returns the UserId field value
func (o *SandboxBankIncomeWebhookFireRequestWebhookFields) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SandboxBankIncomeWebhookFireRequestWebhookFields) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SandboxBankIncomeWebhookFireRequestWebhookFields) SetUserId(v string) {
	o.UserId = v
}

// GetBankIncomeRefreshCompleteResult returns the BankIncomeRefreshCompleteResult field value if set, zero value otherwise.
func (o *SandboxBankIncomeWebhookFireRequestWebhookFields) GetBankIncomeRefreshCompleteResult() BankIncomeRefreshCompleteResult {
	if o == nil || o.BankIncomeRefreshCompleteResult == nil {
		var ret BankIncomeRefreshCompleteResult
		return ret
	}
	return *o.BankIncomeRefreshCompleteResult
}

// GetBankIncomeRefreshCompleteResultOk returns a tuple with the BankIncomeRefreshCompleteResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxBankIncomeWebhookFireRequestWebhookFields) GetBankIncomeRefreshCompleteResultOk() (*BankIncomeRefreshCompleteResult, bool) {
	if o == nil || o.BankIncomeRefreshCompleteResult == nil {
		return nil, false
	}
	return o.BankIncomeRefreshCompleteResult, true
}

// HasBankIncomeRefreshCompleteResult returns a boolean if a field has been set.
func (o *SandboxBankIncomeWebhookFireRequestWebhookFields) HasBankIncomeRefreshCompleteResult() bool {
	if o != nil && o.BankIncomeRefreshCompleteResult != nil {
		return true
	}

	return false
}

// SetBankIncomeRefreshCompleteResult gets a reference to the given BankIncomeRefreshCompleteResult and assigns it to the BankIncomeRefreshCompleteResult field.
func (o *SandboxBankIncomeWebhookFireRequestWebhookFields) SetBankIncomeRefreshCompleteResult(v BankIncomeRefreshCompleteResult) {
	o.BankIncomeRefreshCompleteResult = &v
}

func (o SandboxBankIncomeWebhookFireRequestWebhookFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if o.BankIncomeRefreshCompleteResult != nil {
		toSerialize["bank_income_refresh_complete_result"] = o.BankIncomeRefreshCompleteResult
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxBankIncomeWebhookFireRequestWebhookFields struct {
	value *SandboxBankIncomeWebhookFireRequestWebhookFields
	isSet bool
}

func (v NullableSandboxBankIncomeWebhookFireRequestWebhookFields) Get() *SandboxBankIncomeWebhookFireRequestWebhookFields {
	return v.value
}

func (v *NullableSandboxBankIncomeWebhookFireRequestWebhookFields) Set(val *SandboxBankIncomeWebhookFireRequestWebhookFields) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxBankIncomeWebhookFireRequestWebhookFields) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxBankIncomeWebhookFireRequestWebhookFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxBankIncomeWebhookFireRequestWebhookFields(val *SandboxBankIncomeWebhookFireRequestWebhookFields) *NullableSandboxBankIncomeWebhookFireRequestWebhookFields {
	return &NullableSandboxBankIncomeWebhookFireRequestWebhookFields{value: val, isSet: true}
}

func (v NullableSandboxBankIncomeWebhookFireRequestWebhookFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxBankIncomeWebhookFireRequestWebhookFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


