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

// CreditPayStubAddress Address on the pay stub.
type CreditPayStubAddress struct {
	// The full city name.
	City NullableString `json:"city"`
	// The ISO 3166-1 alpha-2 country code.
	Country NullableString `json:"country"`
	// The postal code of the address.
	PostalCode NullableString `json:"postal_code"`
	// The region or state. Example: `\"NC\"`
	Region NullableString `json:"region"`
	// The full street address.
	Street NullableString `json:"street"`
	AdditionalProperties map[string]interface{}
}

type _CreditPayStubAddress CreditPayStubAddress

// NewCreditPayStubAddress instantiates a new CreditPayStubAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayStubAddress(city NullableString, country NullableString, postalCode NullableString, region NullableString, street NullableString) *CreditPayStubAddress {
	this := CreditPayStubAddress{}
	this.City = city
	this.Country = country
	this.PostalCode = postalCode
	this.Region = region
	this.Street = street
	return &this
}

// NewCreditPayStubAddressWithDefaults instantiates a new CreditPayStubAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayStubAddressWithDefaults() *CreditPayStubAddress {
	this := CreditPayStubAddress{}
	return &this
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPayStubAddress) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubAddress) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *CreditPayStubAddress) SetCity(v string) {
	o.City.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPayStubAddress) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubAddress) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *CreditPayStubAddress) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPayStubAddress) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *CreditPayStubAddress) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPayStubAddress) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubAddress) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *CreditPayStubAddress) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetStreet returns the Street field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPayStubAddress) GetStreet() string {
	if o == nil || o.Street.Get() == nil {
		var ret string
		return ret
	}

	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubAddress) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// SetStreet sets field value
func (o *CreditPayStubAddress) SetStreet(v string) {
	o.Street.Set(&v)
}

func (o CreditPayStubAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["city"] = o.City.Get()
	}
	if true {
		toSerialize["country"] = o.Country.Get()
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["street"] = o.Street.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditPayStubAddress) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPayStubAddress := _CreditPayStubAddress{}

	if err = json.Unmarshal(bytes, &varCreditPayStubAddress); err == nil {
		*o = CreditPayStubAddress(varCreditPayStubAddress)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "country")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "region")
		delete(additionalProperties, "street")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPayStubAddress struct {
	value *CreditPayStubAddress
	isSet bool
}

func (v NullableCreditPayStubAddress) Get() *CreditPayStubAddress {
	return v.value
}

func (v *NullableCreditPayStubAddress) Set(val *CreditPayStubAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayStubAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayStubAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayStubAddress(val *CreditPayStubAddress) *NullableCreditPayStubAddress {
	return &NullableCreditPayStubAddress{value: val, isSet: true}
}

func (v NullableCreditPayStubAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayStubAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


