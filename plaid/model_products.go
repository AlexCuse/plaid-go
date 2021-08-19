/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.21.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// Products A list of products that an institution can support. All Items must be initialized with at least one product. The Balance product is always available and does not need to be specified during initialization.
type Products string

// List of Products
const (
	PRODUCTS_ASSETS Products = "assets"
	PRODUCTS_AUTH Products = "auth"
	PRODUCTS_BALANCE Products = "balance"
	PRODUCTS_IDENTITY Products = "identity"
	PRODUCTS_INVESTMENTS Products = "investments"
	PRODUCTS_LIABILITIES Products = "liabilities"
	PRODUCTS_PAYMENT_INITIATION Products = "payment_initiation"
	PRODUCTS_TRANSACTIONS Products = "transactions"
	PRODUCTS_CREDIT_DETAILS Products = "credit_details"
	PRODUCTS_INCOME Products = "income"
	PRODUCTS_INCOME_VERIFICATION Products = "income_verification"
	PRODUCTS_DEPOSIT_SWITCH Products = "deposit_switch"
	PRODUCTS_STANDING_ORDERS Products = "standing_orders"
)

var allowedProductsEnumValues = []Products{
	"assets",
	"auth",
	"balance",
	"identity",
	"investments",
	"liabilities",
	"payment_initiation",
	"transactions",
	"credit_details",
	"income",
	"income_verification",
	"deposit_switch",
	"standing_orders",
}

func (v *Products) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Products(value)
	for _, existing := range allowedProductsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Products", value)
}

// NewProductsFromValue returns a pointer to a valid Products
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductsFromValue(v string) (*Products, error) {
	ev := Products(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Products: valid values are %v", v, allowedProductsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Products) IsValid() bool {
	for _, existing := range allowedProductsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Products value
func (v Products) Ptr() *Products {
	return &v
}

type NullableProducts struct {
	value *Products
	isSet bool
}

func (v NullableProducts) Get() *Products {
	return v.value
}

func (v *NullableProducts) Set(val *Products) {
	v.value = val
	v.isSet = true
}

func (v NullableProducts) IsSet() bool {
	return v.isSet
}

func (v *NullableProducts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProducts(val *Products) *NullableProducts {
	return &NullableProducts{value: val, isSet: true}
}

func (v NullableProducts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProducts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
