/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditFreddieMacAssetTransactionsVOE25 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacAssetTransactionsVOE25 struct {
	ASSET_TRANSACTION []CreditFreddieMacAssetTransactionVOE25 `json:"ASSET_TRANSACTION"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacAssetTransactionsVOE25 CreditFreddieMacAssetTransactionsVOE25

// NewCreditFreddieMacAssetTransactionsVOE25 instantiates a new CreditFreddieMacAssetTransactionsVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacAssetTransactionsVOE25(aSSETTRANSACTION []CreditFreddieMacAssetTransactionVOE25) *CreditFreddieMacAssetTransactionsVOE25 {
	this := CreditFreddieMacAssetTransactionsVOE25{}
	this.ASSET_TRANSACTION = aSSETTRANSACTION
	return &this
}

// NewCreditFreddieMacAssetTransactionsVOE25WithDefaults instantiates a new CreditFreddieMacAssetTransactionsVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacAssetTransactionsVOE25WithDefaults() *CreditFreddieMacAssetTransactionsVOE25 {
	this := CreditFreddieMacAssetTransactionsVOE25{}
	return &this
}

// GetASSET_TRANSACTION returns the ASSET_TRANSACTION field value
func (o *CreditFreddieMacAssetTransactionsVOE25) GetASSET_TRANSACTION() []CreditFreddieMacAssetTransactionVOE25 {
	if o == nil {
		var ret []CreditFreddieMacAssetTransactionVOE25
		return ret
	}

	return o.ASSET_TRANSACTION
}

// GetASSET_TRANSACTIONOk returns a tuple with the ASSET_TRANSACTION field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetTransactionsVOE25) GetASSET_TRANSACTIONOk() (*[]CreditFreddieMacAssetTransactionVOE25, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTION, true
}

// SetASSET_TRANSACTION sets field value
func (o *CreditFreddieMacAssetTransactionsVOE25) SetASSET_TRANSACTION(v []CreditFreddieMacAssetTransactionVOE25) {
	o.ASSET_TRANSACTION = v
}

func (o CreditFreddieMacAssetTransactionsVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSET_TRANSACTION"] = o.ASSET_TRANSACTION
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacAssetTransactionsVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacAssetTransactionsVOE25 := _CreditFreddieMacAssetTransactionsVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacAssetTransactionsVOE25); err == nil {
		*o = CreditFreddieMacAssetTransactionsVOE25(varCreditFreddieMacAssetTransactionsVOE25)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSET_TRANSACTION")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacAssetTransactionsVOE25 struct {
	value *CreditFreddieMacAssetTransactionsVOE25
	isSet bool
}

func (v NullableCreditFreddieMacAssetTransactionsVOE25) Get() *CreditFreddieMacAssetTransactionsVOE25 {
	return v.value
}

func (v *NullableCreditFreddieMacAssetTransactionsVOE25) Set(val *CreditFreddieMacAssetTransactionsVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacAssetTransactionsVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacAssetTransactionsVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacAssetTransactionsVOE25(val *CreditFreddieMacAssetTransactionsVOE25) *NullableCreditFreddieMacAssetTransactionsVOE25 {
	return &NullableCreditFreddieMacAssetTransactionsVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieMacAssetTransactionsVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacAssetTransactionsVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


