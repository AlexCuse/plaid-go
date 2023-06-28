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

// CreditFreddieMacLoanIdentifiersVOA24 Collection of current and previous identifiers for this loan.
type CreditFreddieMacLoanIdentifiersVOA24 struct {
	LOAN_IDENTIFIER []LoanIdentifier `json:"LOAN_IDENTIFIER"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacLoanIdentifiersVOA24 CreditFreddieMacLoanIdentifiersVOA24

// NewCreditFreddieMacLoanIdentifiersVOA24 instantiates a new CreditFreddieMacLoanIdentifiersVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacLoanIdentifiersVOA24(lOANIDENTIFIER []LoanIdentifier) *CreditFreddieMacLoanIdentifiersVOA24 {
	this := CreditFreddieMacLoanIdentifiersVOA24{}
	this.LOAN_IDENTIFIER = lOANIDENTIFIER
	return &this
}

// NewCreditFreddieMacLoanIdentifiersVOA24WithDefaults instantiates a new CreditFreddieMacLoanIdentifiersVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacLoanIdentifiersVOA24WithDefaults() *CreditFreddieMacLoanIdentifiersVOA24 {
	this := CreditFreddieMacLoanIdentifiersVOA24{}
	return &this
}

// GetLOAN_IDENTIFIER returns the LOAN_IDENTIFIER field value
func (o *CreditFreddieMacLoanIdentifiersVOA24) GetLOAN_IDENTIFIER() []LoanIdentifier {
	if o == nil {
		var ret []LoanIdentifier
		return ret
	}

	return o.LOAN_IDENTIFIER
}

// GetLOAN_IDENTIFIEROk returns a tuple with the LOAN_IDENTIFIER field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacLoanIdentifiersVOA24) GetLOAN_IDENTIFIEROk() (*[]LoanIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LOAN_IDENTIFIER, true
}

// SetLOAN_IDENTIFIER sets field value
func (o *CreditFreddieMacLoanIdentifiersVOA24) SetLOAN_IDENTIFIER(v []LoanIdentifier) {
	o.LOAN_IDENTIFIER = v
}

func (o CreditFreddieMacLoanIdentifiersVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LOAN_IDENTIFIER"] = o.LOAN_IDENTIFIER
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacLoanIdentifiersVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacLoanIdentifiersVOA24 := _CreditFreddieMacLoanIdentifiersVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacLoanIdentifiersVOA24); err == nil {
		*o = CreditFreddieMacLoanIdentifiersVOA24(varCreditFreddieMacLoanIdentifiersVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LOAN_IDENTIFIER")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacLoanIdentifiersVOA24 struct {
	value *CreditFreddieMacLoanIdentifiersVOA24
	isSet bool
}

func (v NullableCreditFreddieMacLoanIdentifiersVOA24) Get() *CreditFreddieMacLoanIdentifiersVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacLoanIdentifiersVOA24) Set(val *CreditFreddieMacLoanIdentifiersVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacLoanIdentifiersVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacLoanIdentifiersVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacLoanIdentifiersVOA24(val *CreditFreddieMacLoanIdentifiersVOA24) *NullableCreditFreddieMacLoanIdentifiersVOA24 {
	return &NullableCreditFreddieMacLoanIdentifiersVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacLoanIdentifiersVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacLoanIdentifiersVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


