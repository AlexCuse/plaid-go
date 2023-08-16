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

// CreditCategory Information describing the intent of the transaction. Most relevant for credit use cases, but not limited to such use cases. Please reach out to your account manager or sales representative if you would like to receive this field.  See the [`taxonomy csv file`](https://plaid.com/documents/credit-category-taxonomy.csv) for a full list of credit categories.
type CreditCategory struct {
	// A high level category that communicates the broad category of the transaction.
	Primary string `json:"primary"`
	// A granular category conveying the transaction's intent. This field can also be used as a unique identifier for the category.
	Detailed string `json:"detailed"`
	AdditionalProperties map[string]interface{}
}

type _CreditCategory CreditCategory

// NewCreditCategory instantiates a new CreditCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCategory(primary string, detailed string) *CreditCategory {
	this := CreditCategory{}
	this.Primary = primary
	this.Detailed = detailed
	return &this
}

// NewCreditCategoryWithDefaults instantiates a new CreditCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCategoryWithDefaults() *CreditCategory {
	this := CreditCategory{}
	return &this
}

// GetPrimary returns the Primary field value
func (o *CreditCategory) GetPrimary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *CreditCategory) GetPrimaryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *CreditCategory) SetPrimary(v string) {
	o.Primary = v
}

// GetDetailed returns the Detailed field value
func (o *CreditCategory) GetDetailed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detailed
}

// GetDetailedOk returns a tuple with the Detailed field value
// and a boolean to check if the value has been set.
func (o *CreditCategory) GetDetailedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Detailed, true
}

// SetDetailed sets field value
func (o *CreditCategory) SetDetailed(v string) {
	o.Detailed = v
}

func (o CreditCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["primary"] = o.Primary
	}
	if true {
		toSerialize["detailed"] = o.Detailed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditCategory) UnmarshalJSON(bytes []byte) (err error) {
	varCreditCategory := _CreditCategory{}

	if err = json.Unmarshal(bytes, &varCreditCategory); err == nil {
		*o = CreditCategory(varCreditCategory)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "primary")
		delete(additionalProperties, "detailed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditCategory struct {
	value *CreditCategory
	isSet bool
}

func (v NullableCreditCategory) Get() *CreditCategory {
	return v.value
}

func (v *NullableCreditCategory) Set(val *CreditCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCategory(val *CreditCategory) *NullableCreditCategory {
	return &NullableCreditCategory{value: val, isSet: true}
}

func (v NullableCreditCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


