/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.413.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// InvestmentFilter A filter to apply to `investment`-type accounts (or `brokerage`-type accounts for API versions 2018-05-22 and earlier).
type InvestmentFilter struct {
	// An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema). 
	AccountSubtypes []InvestmentAccountSubtype `json:"account_subtypes"`
	AdditionalProperties map[string]interface{}
}

type _InvestmentFilter InvestmentFilter

// NewInvestmentFilter instantiates a new InvestmentFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentFilter(accountSubtypes []InvestmentAccountSubtype) *InvestmentFilter {
	this := InvestmentFilter{}
	this.AccountSubtypes = accountSubtypes
	return &this
}

// NewInvestmentFilterWithDefaults instantiates a new InvestmentFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentFilterWithDefaults() *InvestmentFilter {
	this := InvestmentFilter{}
	return &this
}

// GetAccountSubtypes returns the AccountSubtypes field value
func (o *InvestmentFilter) GetAccountSubtypes() []InvestmentAccountSubtype {
	if o == nil {
		var ret []InvestmentAccountSubtype
		return ret
	}

	return o.AccountSubtypes
}

// GetAccountSubtypesOk returns a tuple with the AccountSubtypes field value
// and a boolean to check if the value has been set.
func (o *InvestmentFilter) GetAccountSubtypesOk() (*[]InvestmentAccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountSubtypes, true
}

// SetAccountSubtypes sets field value
func (o *InvestmentFilter) SetAccountSubtypes(v []InvestmentAccountSubtype) {
	o.AccountSubtypes = v
}

func (o InvestmentFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_subtypes"] = o.AccountSubtypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InvestmentFilter) UnmarshalJSON(bytes []byte) (err error) {
	varInvestmentFilter := _InvestmentFilter{}

	if err = json.Unmarshal(bytes, &varInvestmentFilter); err == nil {
		*o = InvestmentFilter(varInvestmentFilter)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_subtypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvestmentFilter struct {
	value *InvestmentFilter
	isSet bool
}

func (v NullableInvestmentFilter) Get() *InvestmentFilter {
	return v.value
}

func (v *NullableInvestmentFilter) Set(val *InvestmentFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentFilter(val *InvestmentFilter) *NullableInvestmentFilter {
	return &NullableInvestmentFilter{value: val, isSet: true}
}

func (v NullableInvestmentFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


