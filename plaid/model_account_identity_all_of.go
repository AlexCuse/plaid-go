/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AccountIdentityAllOf struct for AccountIdentityAllOf
type AccountIdentityAllOf struct {
	// Data returned by the financial institution about the account owner or owners. Only returned by Identity or Assets endpoints. For business accounts, the name reported may be either the name of the individual or the name of the business, depending on the institution. Multiple owners on a single account will be represented in the same `owner` object, not in multiple owner objects within the array. In API versions 2018-05-22 and earlier, the `owners` object is not returned, and instead identity information is returned in the top level `identity` object. For more details, see [Plaid API versioning](https://plaid.com/docs/api/versioning/#version-2019-05-29)
	Owners []Owner `json:"owners"`
}

// NewAccountIdentityAllOf instantiates a new AccountIdentityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountIdentityAllOf(owners []Owner) *AccountIdentityAllOf {
	this := AccountIdentityAllOf{}
	this.Owners = owners
	return &this
}

// NewAccountIdentityAllOfWithDefaults instantiates a new AccountIdentityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountIdentityAllOfWithDefaults() *AccountIdentityAllOf {
	this := AccountIdentityAllOf{}
	return &this
}

// GetOwners returns the Owners field value
func (o *AccountIdentityAllOf) GetOwners() []Owner {
	if o == nil {
		var ret []Owner
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *AccountIdentityAllOf) GetOwnersOk() (*[]Owner, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value
func (o *AccountIdentityAllOf) SetOwners(v []Owner) {
	o.Owners = v
}

func (o AccountIdentityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["owners"] = o.Owners
	}
	return json.Marshal(toSerialize)
}

type NullableAccountIdentityAllOf struct {
	value *AccountIdentityAllOf
	isSet bool
}

func (v NullableAccountIdentityAllOf) Get() *AccountIdentityAllOf {
	return v.value
}

func (v *NullableAccountIdentityAllOf) Set(val *AccountIdentityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountIdentityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountIdentityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountIdentityAllOf(val *AccountIdentityAllOf) *NullableAccountIdentityAllOf {
	return &NullableAccountIdentityAllOf{value: val, isSet: true}
}

func (v NullableAccountIdentityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountIdentityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


