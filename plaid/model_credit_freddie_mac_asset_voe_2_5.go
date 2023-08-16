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

// CreditFreddieMacAssetVOE25 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacAssetVOE25 struct {
	ASSET_DETAIL CreditFreddieMacAssetDetailVOE25 `json:"ASSET_DETAIL"`
	ASSET_OWNERS AssetOwners `json:"ASSET_OWNERS"`
	ASSET_HOLDER AssetHolder `json:"ASSET_HOLDER"`
	ASSET_TRANSACTIONS CreditFreddieMacAssetTransactionsVOE25 `json:"ASSET_TRANSACTIONS"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacAssetVOE25 CreditFreddieMacAssetVOE25

// NewCreditFreddieMacAssetVOE25 instantiates a new CreditFreddieMacAssetVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacAssetVOE25(aSSETDETAIL CreditFreddieMacAssetDetailVOE25, aSSETOWNERS AssetOwners, aSSETHOLDER AssetHolder, aSSETTRANSACTIONS CreditFreddieMacAssetTransactionsVOE25) *CreditFreddieMacAssetVOE25 {
	this := CreditFreddieMacAssetVOE25{}
	this.ASSET_DETAIL = aSSETDETAIL
	this.ASSET_OWNERS = aSSETOWNERS
	this.ASSET_HOLDER = aSSETHOLDER
	this.ASSET_TRANSACTIONS = aSSETTRANSACTIONS
	return &this
}

// NewCreditFreddieMacAssetVOE25WithDefaults instantiates a new CreditFreddieMacAssetVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacAssetVOE25WithDefaults() *CreditFreddieMacAssetVOE25 {
	this := CreditFreddieMacAssetVOE25{}
	return &this
}

// GetASSET_DETAIL returns the ASSET_DETAIL field value
func (o *CreditFreddieMacAssetVOE25) GetASSET_DETAIL() CreditFreddieMacAssetDetailVOE25 {
	if o == nil {
		var ret CreditFreddieMacAssetDetailVOE25
		return ret
	}

	return o.ASSET_DETAIL
}

// GetASSET_DETAILOk returns a tuple with the ASSET_DETAIL field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetVOE25) GetASSET_DETAILOk() (*CreditFreddieMacAssetDetailVOE25, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_DETAIL, true
}

// SetASSET_DETAIL sets field value
func (o *CreditFreddieMacAssetVOE25) SetASSET_DETAIL(v CreditFreddieMacAssetDetailVOE25) {
	o.ASSET_DETAIL = v
}

// GetASSET_OWNERS returns the ASSET_OWNERS field value
func (o *CreditFreddieMacAssetVOE25) GetASSET_OWNERS() AssetOwners {
	if o == nil {
		var ret AssetOwners
		return ret
	}

	return o.ASSET_OWNERS
}

// GetASSET_OWNERSOk returns a tuple with the ASSET_OWNERS field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetVOE25) GetASSET_OWNERSOk() (*AssetOwners, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_OWNERS, true
}

// SetASSET_OWNERS sets field value
func (o *CreditFreddieMacAssetVOE25) SetASSET_OWNERS(v AssetOwners) {
	o.ASSET_OWNERS = v
}

// GetASSET_HOLDER returns the ASSET_HOLDER field value
func (o *CreditFreddieMacAssetVOE25) GetASSET_HOLDER() AssetHolder {
	if o == nil {
		var ret AssetHolder
		return ret
	}

	return o.ASSET_HOLDER
}

// GetASSET_HOLDEROk returns a tuple with the ASSET_HOLDER field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetVOE25) GetASSET_HOLDEROk() (*AssetHolder, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_HOLDER, true
}

// SetASSET_HOLDER sets field value
func (o *CreditFreddieMacAssetVOE25) SetASSET_HOLDER(v AssetHolder) {
	o.ASSET_HOLDER = v
}

// GetASSET_TRANSACTIONS returns the ASSET_TRANSACTIONS field value
func (o *CreditFreddieMacAssetVOE25) GetASSET_TRANSACTIONS() CreditFreddieMacAssetTransactionsVOE25 {
	if o == nil {
		var ret CreditFreddieMacAssetTransactionsVOE25
		return ret
	}

	return o.ASSET_TRANSACTIONS
}

// GetASSET_TRANSACTIONSOk returns a tuple with the ASSET_TRANSACTIONS field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetVOE25) GetASSET_TRANSACTIONSOk() (*CreditFreddieMacAssetTransactionsVOE25, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET_TRANSACTIONS, true
}

// SetASSET_TRANSACTIONS sets field value
func (o *CreditFreddieMacAssetVOE25) SetASSET_TRANSACTIONS(v CreditFreddieMacAssetTransactionsVOE25) {
	o.ASSET_TRANSACTIONS = v
}

func (o CreditFreddieMacAssetVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSET_DETAIL"] = o.ASSET_DETAIL
	}
	if true {
		toSerialize["ASSET_OWNERS"] = o.ASSET_OWNERS
	}
	if true {
		toSerialize["ASSET_HOLDER"] = o.ASSET_HOLDER
	}
	if true {
		toSerialize["ASSET_TRANSACTIONS"] = o.ASSET_TRANSACTIONS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacAssetVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacAssetVOE25 := _CreditFreddieMacAssetVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacAssetVOE25); err == nil {
		*o = CreditFreddieMacAssetVOE25(varCreditFreddieMacAssetVOE25)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSET_DETAIL")
		delete(additionalProperties, "ASSET_OWNERS")
		delete(additionalProperties, "ASSET_HOLDER")
		delete(additionalProperties, "ASSET_TRANSACTIONS")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacAssetVOE25 struct {
	value *CreditFreddieMacAssetVOE25
	isSet bool
}

func (v NullableCreditFreddieMacAssetVOE25) Get() *CreditFreddieMacAssetVOE25 {
	return v.value
}

func (v *NullableCreditFreddieMacAssetVOE25) Set(val *CreditFreddieMacAssetVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacAssetVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacAssetVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacAssetVOE25(val *CreditFreddieMacAssetVOE25) *NullableCreditFreddieMacAssetVOE25 {
	return &NullableCreditFreddieMacAssetVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieMacAssetVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacAssetVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


