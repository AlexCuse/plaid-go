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

// LoanIdentifier The information used to identify this loan by various parties to the transaction or other organizations.
type LoanIdentifier struct {
	// The value of the identifier for the specified type.
	LoanIdentifier NullableString `json:"LoanIdentifier"`
	LoanIdentifierType NullableLoanIdentifierType `json:"LoanIdentifierType"`
	AdditionalProperties map[string]interface{}
}

type _LoanIdentifier LoanIdentifier

// NewLoanIdentifier instantiates a new LoanIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoanIdentifier(loanIdentifier NullableString, loanIdentifierType NullableLoanIdentifierType) *LoanIdentifier {
	this := LoanIdentifier{}
	this.LoanIdentifier = loanIdentifier
	this.LoanIdentifierType = loanIdentifierType
	return &this
}

// NewLoanIdentifierWithDefaults instantiates a new LoanIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoanIdentifierWithDefaults() *LoanIdentifier {
	this := LoanIdentifier{}
	return &this
}

// GetLoanIdentifier returns the LoanIdentifier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LoanIdentifier) GetLoanIdentifier() string {
	if o == nil || o.LoanIdentifier.Get() == nil {
		var ret string
		return ret
	}

	return *o.LoanIdentifier.Get()
}

// GetLoanIdentifierOk returns a tuple with the LoanIdentifier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoanIdentifier) GetLoanIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoanIdentifier.Get(), o.LoanIdentifier.IsSet()
}

// SetLoanIdentifier sets field value
func (o *LoanIdentifier) SetLoanIdentifier(v string) {
	o.LoanIdentifier.Set(&v)
}

// GetLoanIdentifierType returns the LoanIdentifierType field value
// If the value is explicit nil, the zero value for LoanIdentifierType will be returned
func (o *LoanIdentifier) GetLoanIdentifierType() LoanIdentifierType {
	if o == nil || o.LoanIdentifierType.Get() == nil {
		var ret LoanIdentifierType
		return ret
	}

	return *o.LoanIdentifierType.Get()
}

// GetLoanIdentifierTypeOk returns a tuple with the LoanIdentifierType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoanIdentifier) GetLoanIdentifierTypeOk() (*LoanIdentifierType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoanIdentifierType.Get(), o.LoanIdentifierType.IsSet()
}

// SetLoanIdentifierType sets field value
func (o *LoanIdentifier) SetLoanIdentifierType(v LoanIdentifierType) {
	o.LoanIdentifierType.Set(&v)
}

func (o LoanIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LoanIdentifier"] = o.LoanIdentifier.Get()
	}
	if true {
		toSerialize["LoanIdentifierType"] = o.LoanIdentifierType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LoanIdentifier) UnmarshalJSON(bytes []byte) (err error) {
	varLoanIdentifier := _LoanIdentifier{}

	if err = json.Unmarshal(bytes, &varLoanIdentifier); err == nil {
		*o = LoanIdentifier(varLoanIdentifier)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LoanIdentifier")
		delete(additionalProperties, "LoanIdentifierType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoanIdentifier struct {
	value *LoanIdentifier
	isSet bool
}

func (v NullableLoanIdentifier) Get() *LoanIdentifier {
	return v.value
}

func (v *NullableLoanIdentifier) Set(val *LoanIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableLoanIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableLoanIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoanIdentifier(val *LoanIdentifier) *NullableLoanIdentifier {
	return &NullableLoanIdentifier{value: val, isSet: true}
}

func (v NullableLoanIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoanIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


