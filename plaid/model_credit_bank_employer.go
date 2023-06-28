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

// CreditBankEmployer Object containing employer data.
type CreditBankEmployer struct {
	// Name of the employer.
	Name string `json:"name"`
}

// NewCreditBankEmployer instantiates a new CreditBankEmployer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankEmployer(name string) *CreditBankEmployer {
	this := CreditBankEmployer{}
	this.Name = name
	return &this
}

// NewCreditBankEmployerWithDefaults instantiates a new CreditBankEmployer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankEmployerWithDefaults() *CreditBankEmployer {
	this := CreditBankEmployer{}
	return &this
}

// GetName returns the Name field value
func (o *CreditBankEmployer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmployer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreditBankEmployer) SetName(v string) {
	o.Name = v
}

func (o CreditBankEmployer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankEmployer struct {
	value *CreditBankEmployer
	isSet bool
}

func (v NullableCreditBankEmployer) Get() *CreditBankEmployer {
	return v.value
}

func (v *NullableCreditBankEmployer) Set(val *CreditBankEmployer) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankEmployer) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankEmployer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankEmployer(val *CreditBankEmployer) *NullableCreditBankEmployer {
	return &NullableCreditBankEmployer{value: val, isSet: true}
}

func (v NullableCreditBankEmployer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankEmployer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


