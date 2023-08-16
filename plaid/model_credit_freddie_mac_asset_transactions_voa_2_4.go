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

// CreditFreddieMacAssetTransactionsVOA24 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacAssetTransactionsVOA24 struct {
	ASSET_TRANSACTION []CreditFreddieMacAssetTransactionVOA24 `json:"ASSET_TRANSACTION"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacAssetTransactionsVOA24 CreditFreddieMacAssetTransactionsVOA24

// NewCreditFreddieMacAssetTransactionsVOA24 instantiates a new CreditFreddieMacAssetTransactionsVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacAssetTransactionsVOA24(aSSETTRANSACTION []CreditFreddieMacAssetTransactionVOA24) *CreditFreddieMacAssetTransactionsVOA24 {
	this := CreditFreddieMacAssetTransactionsVOA24{}
	this.ASSET_TRANSACTION = aSSETTRANSACTION
	return &this
}

// NewCreditFreddieMacAssetTransactionsVOA24WithDefaults instantiates a new CreditFreddieMacAssetTransactionsVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacAssetTransactionsVOA24WithDefaults() *CreditFreddieMacAssetTransactionsVOA24 {
	this := CreditFreddieMacAssetTransactionsVOA24{}
	return &this
}

// GetASSET_TRANSACTION returns the ASSET_TRANSACTION field value
func (o *CreditFreddieMacAssetTransactionsVOA24) GetASSET_TRANSACTION() []CreditFreddieMacAssetTransactionVOA24 {
	if o == nil {
		var ret []CreditFreddieMacAssetTransactionVOA24
		return ret
	}

	return o.ASSET_TRANSACTION
}

// GetASSET_TRANSACTIONOk returns a tuple with the ASSET_TRANSACTION field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetTransactionsVOA24) GetASSET_TRANSACTIONOk() (*[]CreditFreddieMacAssetTransactionVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTION, true
}

// SetASSET_TRANSACTION sets field value
func (o *CreditFreddieMacAssetTransactionsVOA24) SetASSET_TRANSACTION(v []CreditFreddieMacAssetTransactionVOA24) {
	o.ASSET_TRANSACTION = v
}

func (o CreditFreddieMacAssetTransactionsVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSET_TRANSACTION"] = o.ASSET_TRANSACTION
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacAssetTransactionsVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacAssetTransactionsVOA24 := _CreditFreddieMacAssetTransactionsVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacAssetTransactionsVOA24); err == nil {
		*o = CreditFreddieMacAssetTransactionsVOA24(varCreditFreddieMacAssetTransactionsVOA24)
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

type NullableCreditFreddieMacAssetTransactionsVOA24 struct {
	value *CreditFreddieMacAssetTransactionsVOA24
	isSet bool
}

func (v NullableCreditFreddieMacAssetTransactionsVOA24) Get() *CreditFreddieMacAssetTransactionsVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacAssetTransactionsVOA24) Set(val *CreditFreddieMacAssetTransactionsVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacAssetTransactionsVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacAssetTransactionsVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacAssetTransactionsVOA24(val *CreditFreddieMacAssetTransactionsVOA24) *NullableCreditFreddieMacAssetTransactionsVOA24 {
	return &NullableCreditFreddieMacAssetTransactionsVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacAssetTransactionsVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacAssetTransactionsVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


