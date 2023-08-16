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

// BankTransferFailure The failure reason if the type of this transfer is `\"failed\"` or `\"reversed\"`. Null value otherwise.
type BankTransferFailure struct {
	// The ACH return code, e.g. `R01`.  A return code will be provided if and only if the transfer status is `reversed`. For a full listing of ACH return codes, see [Bank Transfers errors](https://plaid.com/docs/errors/bank-transfers/#ach-return-codes).
	AchReturnCode NullableString `json:"ach_return_code,omitempty"`
	// A human-readable description of the reason for the failure or reversal.
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferFailure BankTransferFailure

// NewBankTransferFailure instantiates a new BankTransferFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferFailure() *BankTransferFailure {
	this := BankTransferFailure{}
	return &this
}

// NewBankTransferFailureWithDefaults instantiates a new BankTransferFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferFailureWithDefaults() *BankTransferFailure {
	this := BankTransferFailure{}
	return &this
}

// GetAchReturnCode returns the AchReturnCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferFailure) GetAchReturnCode() string {
	if o == nil || o.AchReturnCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.AchReturnCode.Get()
}

// GetAchReturnCodeOk returns a tuple with the AchReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferFailure) GetAchReturnCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AchReturnCode.Get(), o.AchReturnCode.IsSet()
}

// HasAchReturnCode returns a boolean if a field has been set.
func (o *BankTransferFailure) HasAchReturnCode() bool {
	if o != nil && o.AchReturnCode.IsSet() {
		return true
	}

	return false
}

// SetAchReturnCode gets a reference to the given NullableString and assigns it to the AchReturnCode field.
func (o *BankTransferFailure) SetAchReturnCode(v string) {
	o.AchReturnCode.Set(&v)
}
// SetAchReturnCodeNil sets the value for AchReturnCode to be an explicit nil
func (o *BankTransferFailure) SetAchReturnCodeNil() {
	o.AchReturnCode.Set(nil)
}

// UnsetAchReturnCode ensures that no value is present for AchReturnCode, not even an explicit nil
func (o *BankTransferFailure) UnsetAchReturnCode() {
	o.AchReturnCode.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BankTransferFailure) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferFailure) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BankTransferFailure) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BankTransferFailure) SetDescription(v string) {
	o.Description = &v
}

func (o BankTransferFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AchReturnCode.IsSet() {
		toSerialize["ach_return_code"] = o.AchReturnCode.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferFailure) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferFailure := _BankTransferFailure{}

	if err = json.Unmarshal(bytes, &varBankTransferFailure); err == nil {
		*o = BankTransferFailure(varBankTransferFailure)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ach_return_code")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferFailure struct {
	value *BankTransferFailure
	isSet bool
}

func (v NullableBankTransferFailure) Get() *BankTransferFailure {
	return v.value
}

func (v *NullableBankTransferFailure) Set(val *BankTransferFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferFailure(val *BankTransferFailure) *NullableBankTransferFailure {
	return &NullableBankTransferFailure{value: val, isSet: true}
}

func (v NullableBankTransferFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


