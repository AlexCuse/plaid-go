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

// CreditFreddieMacAssetTransactionVOE25 An object representing...
type CreditFreddieMacAssetTransactionVOE25 struct {
	ASSET_TRANSACTION_DETAIL CreditFreddieMacAssetTransactionDetailVOE25 `json:"ASSET_TRANSACTION_DETAIL"`
	// Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
	ASSET_TRANSACTION_DESCRIPTION []AssetTransactionDescription `json:"ASSET_TRANSACTION_DESCRIPTION"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacAssetTransactionVOE25 CreditFreddieMacAssetTransactionVOE25

// NewCreditFreddieMacAssetTransactionVOE25 instantiates a new CreditFreddieMacAssetTransactionVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacAssetTransactionVOE25(aSSETTRANSACTIONDETAIL CreditFreddieMacAssetTransactionDetailVOE25, aSSETTRANSACTIONDESCRIPTION []AssetTransactionDescription) *CreditFreddieMacAssetTransactionVOE25 {
	this := CreditFreddieMacAssetTransactionVOE25{}
	this.ASSET_TRANSACTION_DETAIL = aSSETTRANSACTIONDETAIL
	this.ASSET_TRANSACTION_DESCRIPTION = aSSETTRANSACTIONDESCRIPTION
	return &this
}

// NewCreditFreddieMacAssetTransactionVOE25WithDefaults instantiates a new CreditFreddieMacAssetTransactionVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacAssetTransactionVOE25WithDefaults() *CreditFreddieMacAssetTransactionVOE25 {
	this := CreditFreddieMacAssetTransactionVOE25{}
	return &this
}

// GetASSET_TRANSACTION_DETAIL returns the ASSET_TRANSACTION_DETAIL field value
func (o *CreditFreddieMacAssetTransactionVOE25) GetASSET_TRANSACTION_DETAIL() CreditFreddieMacAssetTransactionDetailVOE25 {
	if o == nil {
		var ret CreditFreddieMacAssetTransactionDetailVOE25
		return ret
	}

	return o.ASSET_TRANSACTION_DETAIL
}

// GetASSET_TRANSACTION_DETAILOk returns a tuple with the ASSET_TRANSACTION_DETAIL field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetTransactionVOE25) GetASSET_TRANSACTION_DETAILOk() (*CreditFreddieMacAssetTransactionDetailVOE25, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTION_DETAIL, true
}

// SetASSET_TRANSACTION_DETAIL sets field value
func (o *CreditFreddieMacAssetTransactionVOE25) SetASSET_TRANSACTION_DETAIL(v CreditFreddieMacAssetTransactionDetailVOE25) {
	o.ASSET_TRANSACTION_DETAIL = v
}

// GetASSET_TRANSACTION_DESCRIPTION returns the ASSET_TRANSACTION_DESCRIPTION field value
func (o *CreditFreddieMacAssetTransactionVOE25) GetASSET_TRANSACTION_DESCRIPTION() []AssetTransactionDescription {
	if o == nil {
		var ret []AssetTransactionDescription
		return ret
	}

	return o.ASSET_TRANSACTION_DESCRIPTION
}

// GetASSET_TRANSACTION_DESCRIPTIONOk returns a tuple with the ASSET_TRANSACTION_DESCRIPTION field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetTransactionVOE25) GetASSET_TRANSACTION_DESCRIPTIONOk() (*[]AssetTransactionDescription, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTION_DESCRIPTION, true
}

// SetASSET_TRANSACTION_DESCRIPTION sets field value
func (o *CreditFreddieMacAssetTransactionVOE25) SetASSET_TRANSACTION_DESCRIPTION(v []AssetTransactionDescription) {
	o.ASSET_TRANSACTION_DESCRIPTION = v
}

func (o CreditFreddieMacAssetTransactionVOE25) MarshalJSON() ([]byte, error) {
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

func (o *CreditFreddieMacAssetTransactionVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacAssetTransactionVOE25 := _CreditFreddieMacAssetTransactionVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacAssetTransactionVOE25); err == nil {
		*o = CreditFreddieMacAssetTransactionVOE25(varCreditFreddieMacAssetTransactionVOE25)
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

type NullableCreditFreddieMacAssetTransactionVOE25 struct {
	value *CreditFreddieMacAssetTransactionVOE25
	isSet bool
}

func (v NullableCreditFreddieMacAssetTransactionVOE25) Get() *CreditFreddieMacAssetTransactionVOE25 {
	return v.value
}

func (v *NullableCreditFreddieMacAssetTransactionVOE25) Set(val *CreditFreddieMacAssetTransactionVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacAssetTransactionVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacAssetTransactionVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacAssetTransactionVOE25(val *CreditFreddieMacAssetTransactionVOE25) *NullableCreditFreddieMacAssetTransactionVOE25 {
	return &NullableCreditFreddieMacAssetTransactionVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieMacAssetTransactionVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacAssetTransactionVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


