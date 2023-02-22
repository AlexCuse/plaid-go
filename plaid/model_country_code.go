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

// CountryCode ISO-3166-1 alpha-2 country code standard.
type CountryCode string

var _ = fmt.Printf

// List of CountryCode
const (
	COUNTRYCODE_US CountryCode = "US"
	COUNTRYCODE_GB CountryCode = "GB"
	COUNTRYCODE_ES CountryCode = "ES"
	COUNTRYCODE_NL CountryCode = "NL"
	COUNTRYCODE_FR CountryCode = "FR"
	COUNTRYCODE_IE CountryCode = "IE"
	COUNTRYCODE_CA CountryCode = "CA"
	COUNTRYCODE_DE CountryCode = "DE"
	COUNTRYCODE_IT CountryCode = "IT"
	COUNTRYCODE_PL CountryCode = "PL"
	COUNTRYCODE_DK CountryCode = "DK"
	COUNTRYCODE_NO CountryCode = "NO"
	COUNTRYCODE_SE CountryCode = "SE"
	COUNTRYCODE_EE CountryCode = "EE"
	COUNTRYCODE_LT CountryCode = "LT"
	COUNTRYCODE_LV CountryCode = "LV"
)

var allowedCountryCodeEnumValues = []CountryCode{
	"US",
	"GB",
	"ES",
	"NL",
	"FR",
	"IE",
	"CA",
	"DE",
	"IT",
	"PL",
	"DK",
	"NO",
	"SE",
	"EE",
	"LT",
	"LV",
}

func (v *CountryCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CountryCode(value)


	*v = enumTypeValue
	return nil
}

// NewCountryCodeFromValue returns a pointer to a valid CountryCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCountryCodeFromValue(v string) (*CountryCode, error) {
	ev := CountryCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CountryCode) IsValid() bool {
	for _, existing := range allowedCountryCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CountryCode value
func (v CountryCode) Ptr() *CountryCode {
	return &v
}

type NullableCountryCode struct {
	value *CountryCode
	isSet bool
}

func (v NullableCountryCode) Get() *CountryCode {
	return v.value
}

func (v *NullableCountryCode) Set(val *CountryCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryCode(val *CountryCode) *NullableCountryCode {
	return &NullableCountryCode{value: val, isSet: true}
}

func (v NullableCountryCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

