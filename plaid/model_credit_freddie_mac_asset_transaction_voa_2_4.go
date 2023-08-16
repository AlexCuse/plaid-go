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

// CreditFreddieMacAssetTransactionVOA24 An object representing...
type CreditFreddieMacAssetTransactionVOA24 struct {
	ASSET_TRANSACTION_DETAIL AssetTransactionDetail `json:"ASSET_TRANSACTION_DETAIL"`
	// Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
	ASSET_TRANSACTION_DESCRIPTION []AssetTransactionDescription `json:"ASSET_TRANSACTION_DESCRIPTION"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacAssetTransactionVOA24 CreditFreddieMacAssetTransactionVOA24

// NewCreditFreddieMacAssetTransactionVOA24 instantiates a new CreditFreddieMacAssetTransactionVOA24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacAssetTransactionVOA24(aSSETTRANSACTIONDETAIL AssetTransactionDetail, aSSETTRANSACTIONDESCRIPTION []AssetTransactionDescription) *CreditFreddieMacAssetTransactionVOA24 {
	this := CreditFreddieMacAssetTransactionVOA24{}
	this.ASSET_TRANSACTION_DETAIL = aSSETTRANSACTIONDETAIL
	this.ASSET_TRANSACTION_DESCRIPTION = aSSETTRANSACTIONDESCRIPTION
	return &this
}

// NewCreditFreddieMacAssetTransactionVOA24WithDefaults instantiates a new CreditFreddieMacAssetTransactionVOA24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacAssetTransactionVOA24WithDefaults() *CreditFreddieMacAssetTransactionVOA24 {
	this := CreditFreddieMacAssetTransactionVOA24{}
	return &this
}

// GetASSET_TRANSACTION_DETAIL returns the ASSET_TRANSACTION_DETAIL field value
func (o *CreditFreddieMacAssetTransactionVOA24) GetASSET_TRANSACTION_DETAIL() AssetTransactionDetail {
	if o == nil {
		var ret AssetTransactionDetail
		return ret
	}

	return o.ASSET_TRANSACTION_DETAIL
}

// GetASSET_TRANSACTION_DETAILOk returns a tuple with the ASSET_TRANSACTION_DETAIL field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetTransactionVOA24) GetASSET_TRANSACTION_DETAILOk() (*AssetTransactionDetail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTION_DETAIL, true
}

// SetASSET_TRANSACTION_DETAIL sets field value
func (o *CreditFreddieMacAssetTransactionVOA24) SetASSET_TRANSACTION_DETAIL(v AssetTransactionDetail) {
	o.ASSET_TRANSACTION_DETAIL = v
}

// GetASSET_TRANSACTION_DESCRIPTION returns the ASSET_TRANSACTION_DESCRIPTION field value
func (o *CreditFreddieMacAssetTransactionVOA24) GetASSET_TRANSACTION_DESCRIPTION() []AssetTransactionDescription {
	if o == nil {
		var ret []AssetTransactionDescription
		return ret
	}

	return o.ASSET_TRANSACTION_DESCRIPTION
}

// GetASSET_TRANSACTION_DESCRIPTIONOk returns a tuple with the ASSET_TRANSACTION_DESCRIPTION field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetTransactionVOA24) GetASSET_TRANSACTION_DESCRIPTIONOk() (*[]AssetTransactionDescription, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTION_DESCRIPTION, true
}

// SetASSET_TRANSACTION_DESCRIPTION sets field value
func (o *CreditFreddieMacAssetTransactionVOA24) SetASSET_TRANSACTION_DESCRIPTION(v []AssetTransactionDescription) {
	o.ASSET_TRANSACTION_DESCRIPTION = v
}

func (o CreditFreddieMacAssetTransactionVOA24) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSET_TRANSACTION_DETAIL"] = o.ASSET_TRANSACTION_DETAIL
	}
	if true {
		toSerialize["ASSET_TRANSACTION_DESCRIPTION"] = o.ASSET_TRANSACTION_DESCRIPTION
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacAssetTransactionVOA24) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacAssetTransactionVOA24 := _CreditFreddieMacAssetTransactionVOA24{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacAssetTransactionVOA24); err == nil {
		*o = CreditFreddieMacAssetTransactionVOA24(varCreditFreddieMacAssetTransactionVOA24)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSET_TRANSACTION_DETAIL")
		delete(additionalProperties, "ASSET_TRANSACTION_DESCRIPTION")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacAssetTransactionVOA24 struct {
	value *CreditFreddieMacAssetTransactionVOA24
	isSet bool
}

func (v NullableCreditFreddieMacAssetTransactionVOA24) Get() *CreditFreddieMacAssetTransactionVOA24 {
	return v.value
}

func (v *NullableCreditFreddieMacAssetTransactionVOA24) Set(val *CreditFreddieMacAssetTransactionVOA24) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacAssetTransactionVOA24) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacAssetTransactionVOA24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacAssetTransactionVOA24(val *CreditFreddieMacAssetTransactionVOA24) *NullableCreditFreddieMacAssetTransactionVOA24 {
	return &NullableCreditFreddieMacAssetTransactionVOA24{value: val, isSet: true}
}

func (v NullableCreditFreddieMacAssetTransactionVOA24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacAssetTransactionVOA24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


