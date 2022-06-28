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

// RecipientBACS An object containing a BACS account number and sort code. If an IBAN is not provided or if you need to accept domestic GBP-denominated payments, BACS data is required.
type RecipientBACS struct {
	// The account number of the account. Maximum of 10 characters.
	Account *string `json:"account,omitempty"`
	// The 6-character sort code of the account.
	SortCode *string `json:"sort_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecipientBACS RecipientBACS

// NewRecipientBACS instantiates a new RecipientBACS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientBACS() *RecipientBACS {
	this := RecipientBACS{}
	return &this
}

// NewRecipientBACSWithDefaults instantiates a new RecipientBACS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientBACSWithDefaults() *RecipientBACS {
	this := RecipientBACS{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *RecipientBACS) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientBACS) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *RecipientBACS) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *RecipientBACS) SetAccount(v string) {
	o.Account = &v
}

// GetSortCode returns the SortCode field value if set, zero value otherwise.
func (o *RecipientBACS) GetSortCode() string {
	if o == nil || o.SortCode == nil {
		var ret string
		return ret
	}
	return *o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientBACS) GetSortCodeOk() (*string, bool) {
	if o == nil || o.SortCode == nil {
		return nil, false
	}
	return o.SortCode, true
}

// HasSortCode returns a boolean if a field has been set.
func (o *RecipientBACS) HasSortCode() bool {
	if o != nil && o.SortCode != nil {
		return true
	}

	return false
}

// SetSortCode gets a reference to the given string and assigns it to the SortCode field.
func (o *RecipientBACS) SetSortCode(v string) {
	o.SortCode = &v
}

func (o RecipientBACS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.SortCode != nil {
		toSerialize["sort_code"] = o.SortCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecipientBACS) UnmarshalJSON(bytes []byte) (err error) {
	varRecipientBACS := _RecipientBACS{}

	if err = json.Unmarshal(bytes, &varRecipientBACS); err == nil {
		*o = RecipientBACS(varRecipientBACS)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		delete(additionalProperties, "sort_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecipientBACS struct {
	value *RecipientBACS
	isSet bool
}

func (v NullableRecipientBACS) Get() *RecipientBACS {
	return v.value
}

func (v *NullableRecipientBACS) Set(val *RecipientBACS) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientBACS) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientBACS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientBACS(val *RecipientBACS) *NullableRecipientBACS {
	return &NullableRecipientBACS{value: val, isSet: true}
}

func (v NullableRecipientBACS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientBACS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


