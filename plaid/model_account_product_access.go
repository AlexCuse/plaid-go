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

// AccountProductAccess Allow the application to access specific products on this account
type AccountProductAccess struct {
	// Allow the application to access account data. Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	AccountData NullableBool `json:"account_data,omitempty"`
	// Allow the application to access bank statements. Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	Statements NullableBool `json:"statements,omitempty"`
	// Allow the application to access tax documents. Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	TaxDocuments NullableBool `json:"tax_documents,omitempty"`
}

// NewAccountProductAccess instantiates a new AccountProductAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountProductAccess() *AccountProductAccess {
	this := AccountProductAccess{}
	var accountData bool = true
	this.AccountData = *NewNullableBool(&accountData)
	var statements bool = true
	this.Statements = *NewNullableBool(&statements)
	var taxDocuments bool = true
	this.TaxDocuments = *NewNullableBool(&taxDocuments)
	return &this
}

// NewAccountProductAccessWithDefaults instantiates a new AccountProductAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountProductAccessWithDefaults() *AccountProductAccess {
	this := AccountProductAccess{}
	var accountData bool = true
	this.AccountData = *NewNullableBool(&accountData)
	var statements bool = true
	this.Statements = *NewNullableBool(&statements)
	var taxDocuments bool = true
	this.TaxDocuments = *NewNullableBool(&taxDocuments)
	return &this
}

// GetAccountData returns the AccountData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountProductAccess) GetAccountData() bool {
	if o == nil || o.AccountData.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AccountData.Get()
}

// GetAccountDataOk returns a tuple with the AccountData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountProductAccess) GetAccountDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountData.Get(), o.AccountData.IsSet()
}

// HasAccountData returns a boolean if a field has been set.
func (o *AccountProductAccess) HasAccountData() bool {
	if o != nil && o.AccountData.IsSet() {
		return true
	}

	return false
}

// SetAccountData gets a reference to the given NullableBool and assigns it to the AccountData field.
func (o *AccountProductAccess) SetAccountData(v bool) {
	o.AccountData.Set(&v)
}
// SetAccountDataNil sets the value for AccountData to be an explicit nil
func (o *AccountProductAccess) SetAccountDataNil() {
	o.AccountData.Set(nil)
}

// UnsetAccountData ensures that no value is present for AccountData, not even an explicit nil
func (o *AccountProductAccess) UnsetAccountData() {
	o.AccountData.Unset()
}

// GetStatements returns the Statements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountProductAccess) GetStatements() bool {
	if o == nil || o.Statements.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Statements.Get()
}

// GetStatementsOk returns a tuple with the Statements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountProductAccess) GetStatementsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Statements.Get(), o.Statements.IsSet()
}

// HasStatements returns a boolean if a field has been set.
func (o *AccountProductAccess) HasStatements() bool {
	if o != nil && o.Statements.IsSet() {
		return true
	}

	return false
}

// SetStatements gets a reference to the given NullableBool and assigns it to the Statements field.
func (o *AccountProductAccess) SetStatements(v bool) {
	o.Statements.Set(&v)
}
// SetStatementsNil sets the value for Statements to be an explicit nil
func (o *AccountProductAccess) SetStatementsNil() {
	o.Statements.Set(nil)
}

// UnsetStatements ensures that no value is present for Statements, not even an explicit nil
func (o *AccountProductAccess) UnsetStatements() {
	o.Statements.Unset()
}

// GetTaxDocuments returns the TaxDocuments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountProductAccess) GetTaxDocuments() bool {
	if o == nil || o.TaxDocuments.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TaxDocuments.Get()
}

// GetTaxDocumentsOk returns a tuple with the TaxDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountProductAccess) GetTaxDocumentsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxDocuments.Get(), o.TaxDocuments.IsSet()
}

// HasTaxDocuments returns a boolean if a field has been set.
func (o *AccountProductAccess) HasTaxDocuments() bool {
	if o != nil && o.TaxDocuments.IsSet() {
		return true
	}

	return false
}

// SetTaxDocuments gets a reference to the given NullableBool and assigns it to the TaxDocuments field.
func (o *AccountProductAccess) SetTaxDocuments(v bool) {
	o.TaxDocuments.Set(&v)
}
// SetTaxDocumentsNil sets the value for TaxDocuments to be an explicit nil
func (o *AccountProductAccess) SetTaxDocumentsNil() {
	o.TaxDocuments.Set(nil)
}

// UnsetTaxDocuments ensures that no value is present for TaxDocuments, not even an explicit nil
func (o *AccountProductAccess) UnsetTaxDocuments() {
	o.TaxDocuments.Unset()
}

func (o AccountProductAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountData.IsSet() {
		toSerialize["account_data"] = o.AccountData.Get()
	}
	if o.Statements.IsSet() {
		toSerialize["statements"] = o.Statements.Get()
	}
	if o.TaxDocuments.IsSet() {
		toSerialize["tax_documents"] = o.TaxDocuments.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAccountProductAccess struct {
	value *AccountProductAccess
	isSet bool
}

func (v NullableAccountProductAccess) Get() *AccountProductAccess {
	return v.value
}

func (v *NullableAccountProductAccess) Set(val *AccountProductAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountProductAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountProductAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountProductAccess(val *AccountProductAccess) *NullableAccountProductAccess {
	return &NullableAccountProductAccess{value: val, isSet: true}
}

func (v NullableAccountProductAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountProductAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


