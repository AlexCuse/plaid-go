/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.97.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IncomeVerificationCreateRequestOptions Optional arguments for `/income/verification/create`
type IncomeVerificationCreateRequestOptions struct {
	// An array of access tokens corresponding to the Items that will be cross-referenced with the product data. Plaid will attempt to correlate transaction history from these Items with data from the user's paystub, such as date and amount. The `verification` status of the paystub as returned by `/income/verification/paystubs/get` will indicate if the verification status was successful, or, if not, why it failed. If the `transactions` product was not initialized for the Items during Link, it will be initialized after this Link session.
	AccessTokens *[]string `json:"access_tokens,omitempty"`
}

// NewIncomeVerificationCreateRequestOptions instantiates a new IncomeVerificationCreateRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationCreateRequestOptions() *IncomeVerificationCreateRequestOptions {
	this := IncomeVerificationCreateRequestOptions{}
	return &this
}

// NewIncomeVerificationCreateRequestOptionsWithDefaults instantiates a new IncomeVerificationCreateRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationCreateRequestOptionsWithDefaults() *IncomeVerificationCreateRequestOptions {
	this := IncomeVerificationCreateRequestOptions{}
	return &this
}

// GetAccessTokens returns the AccessTokens field value if set, zero value otherwise.
func (o *IncomeVerificationCreateRequestOptions) GetAccessTokens() []string {
	if o == nil || o.AccessTokens == nil {
		var ret []string
		return ret
	}
	return *o.AccessTokens
}

// GetAccessTokensOk returns a tuple with the AccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationCreateRequestOptions) GetAccessTokensOk() (*[]string, bool) {
	if o == nil || o.AccessTokens == nil {
		return nil, false
	}
	return o.AccessTokens, true
}

// HasAccessTokens returns a boolean if a field has been set.
func (o *IncomeVerificationCreateRequestOptions) HasAccessTokens() bool {
	if o != nil && o.AccessTokens != nil {
		return true
	}

	return false
}

// SetAccessTokens gets a reference to the given []string and assigns it to the AccessTokens field.
func (o *IncomeVerificationCreateRequestOptions) SetAccessTokens(v []string) {
	o.AccessTokens = &v
}

func (o IncomeVerificationCreateRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessTokens != nil {
		toSerialize["access_tokens"] = o.AccessTokens
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationCreateRequestOptions struct {
	value *IncomeVerificationCreateRequestOptions
	isSet bool
}

func (v NullableIncomeVerificationCreateRequestOptions) Get() *IncomeVerificationCreateRequestOptions {
	return v.value
}

func (v *NullableIncomeVerificationCreateRequestOptions) Set(val *IncomeVerificationCreateRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationCreateRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationCreateRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationCreateRequestOptions(val *IncomeVerificationCreateRequestOptions) *NullableIncomeVerificationCreateRequestOptions {
	return &NullableIncomeVerificationCreateRequestOptions{value: val, isSet: true}
}

func (v NullableIncomeVerificationCreateRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationCreateRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


