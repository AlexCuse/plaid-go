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
	"fmt"
)

// AssetType A value from a MISMO prescribed list that specifies financial assets in a mortgage loan transaction. Assets may be either liquid or fixed and are associated with a corresponding asset amount.
type AssetType string

var _ = fmt.Printf

// List of AssetType
const (
	ASSETTYPE_CHECKING_ACCOUNT AssetType = "CheckingAccount"
	ASSETTYPE_SAVINGS_ACCOUNT AssetType = "SavingsAccount"
	ASSETTYPE_INVESTMENT AssetType = "Investment"
	ASSETTYPE_MONEY_MARKET_FUND AssetType = "MoneyMarketFund"
	ASSETTYPE_OTHER AssetType = "Other"
)

var allowedAssetTypeEnumValues = []AssetType{
	"CheckingAccount",
	"SavingsAccount",
	"Investment",
	"MoneyMarketFund",
	"Other",
}

func (v *AssetType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := AssetType(value)


	*v = enumTypeValue
	return nil
}

// NewAssetTypeFromValue returns a pointer to a valid AssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetTypeFromValue(v string) (*AssetType, error) {
	ev := AssetType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetType) IsValid() bool {
	for _, existing := range allowedAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssetType value
func (v AssetType) Ptr() *AssetType {
	return &v
}

type NullableAssetType struct {
	value *AssetType
	isSet bool
}

func (v NullableAssetType) Get() *AssetType {
	return v.value
}

func (v *NullableAssetType) Set(val *AssetType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetType(val *AssetType) *NullableAssetType {
	return &NullableAssetType{value: val, isSet: true}
}

func (v NullableAssetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

