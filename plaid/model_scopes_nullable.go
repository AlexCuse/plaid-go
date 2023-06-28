/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.385.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ScopesNullable The scopes object
type ScopesNullable struct {
	ProductAccess *ProductAccess `json:"product_access,omitempty"`
	Accounts *[]AccountAccess `json:"accounts,omitempty"`
	// Allow access to newly opened accounts as they are opened. If unset, defaults to `true`.
	NewAccounts NullableBool `json:"new_accounts,omitempty"`
}

// NewScopesNullable instantiates a new ScopesNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopesNullable() *ScopesNullable {
	this := ScopesNullable{}
	var newAccounts bool = true
	this.NewAccounts = *NewNullableBool(&newAccounts)
	return &this
}

// NewScopesNullableWithDefaults instantiates a new ScopesNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopesNullableWithDefaults() *ScopesNullable {
	this := ScopesNullable{}
	var newAccounts bool = true
	this.NewAccounts = *NewNullableBool(&newAccounts)
	return &this
}

// GetProductAccess returns the ProductAccess field value if set, zero value otherwise.
func (o *ScopesNullable) GetProductAccess() ProductAccess {
	if o == nil || o.ProductAccess == nil {
		var ret ProductAccess
		return ret
	}
	return *o.ProductAccess
}

// GetProductAccessOk returns a tuple with the ProductAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopesNullable) GetProductAccessOk() (*ProductAccess, bool) {
	if o == nil || o.ProductAccess == nil {
		return nil, false
	}
	return o.ProductAccess, true
}

// HasProductAccess returns a boolean if a field has been set.
func (o *ScopesNullable) HasProductAccess() bool {
	if o != nil && o.ProductAccess != nil {
		return true
	}

	return false
}

// SetProductAccess gets a reference to the given ProductAccess and assigns it to the ProductAccess field.
func (o *ScopesNullable) SetProductAccess(v ProductAccess) {
	o.ProductAccess = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *ScopesNullable) GetAccounts() []AccountAccess {
	if o == nil || o.Accounts == nil {
		var ret []AccountAccess
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopesNullable) GetAccountsOk() (*[]AccountAccess, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *ScopesNullable) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []AccountAccess and assigns it to the Accounts field.
func (o *ScopesNullable) SetAccounts(v []AccountAccess) {
	o.Accounts = &v
}

// GetNewAccounts returns the NewAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScopesNullable) GetNewAccounts() bool {
	if o == nil || o.NewAccounts.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NewAccounts.Get()
}

// GetNewAccountsOk returns a tuple with the NewAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScopesNullable) GetNewAccountsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewAccounts.Get(), o.NewAccounts.IsSet()
}

// HasNewAccounts returns a boolean if a field has been set.
func (o *ScopesNullable) HasNewAccounts() bool {
	if o != nil && o.NewAccounts.IsSet() {
		return true
	}

	return false
}

// SetNewAccounts gets a reference to the given NullableBool and assigns it to the NewAccounts field.
func (o *ScopesNullable) SetNewAccounts(v bool) {
	o.NewAccounts.Set(&v)
}
// SetNewAccountsNil sets the value for NewAccounts to be an explicit nil
func (o *ScopesNullable) SetNewAccountsNil() {
	o.NewAccounts.Set(nil)
}

// UnsetNewAccounts ensures that no value is present for NewAccounts, not even an explicit nil
func (o *ScopesNullable) UnsetNewAccounts() {
	o.NewAccounts.Unset()
}

func (o ScopesNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductAccess != nil {
		toSerialize["product_access"] = o.ProductAccess
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.NewAccounts.IsSet() {
		toSerialize["new_accounts"] = o.NewAccounts.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableScopesNullable struct {
	value *ScopesNullable
	isSet bool
}

func (v NullableScopesNullable) Get() *ScopesNullable {
	return v.value
}

func (v *NullableScopesNullable) Set(val *ScopesNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableScopesNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableScopesNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopesNullable(val *ScopesNullable) *NullableScopesNullable {
	return &NullableScopesNullable{value: val, isSet: true}
}

func (v NullableScopesNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopesNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


