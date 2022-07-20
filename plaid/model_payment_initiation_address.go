/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentInitiationAddress The optional address of the payment recipient.
type PaymentInitiationAddress struct {
	// An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.
	Street []string `json:"street"`
	// The city where the recipient is located. Maximum of 35 characters.
	City string `json:"city"`
	// The postal code where the recipient is located. Maximum of 16 characters.
	PostalCode string `json:"postal_code"`
	// The ISO 3166-1 alpha-2 country code where the recipient is located.
	Country string `json:"country"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationAddress PaymentInitiationAddress

// NewPaymentInitiationAddress instantiates a new PaymentInitiationAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationAddress(street []string, city string, postalCode string, country string) *PaymentInitiationAddress {
	this := PaymentInitiationAddress{}
	this.Street = street
	this.City = city
	this.PostalCode = postalCode
	this.Country = country
	return &this
}

// NewPaymentInitiationAddressWithDefaults instantiates a new PaymentInitiationAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationAddressWithDefaults() *PaymentInitiationAddress {
	this := PaymentInitiationAddress{}
	return &this
}

// GetStreet returns the Street field value
func (o *PaymentInitiationAddress) GetStreet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationAddress) GetStreetOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *PaymentInitiationAddress) SetStreet(v []string) {
	o.Street = v
}

// GetCity returns the City field value
func (o *PaymentInitiationAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationAddress) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *PaymentInitiationAddress) SetCity(v string) {
	o.City = v
}

// GetPostalCode returns the PostalCode field value
func (o *PaymentInitiationAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *PaymentInitiationAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetCountry returns the Country field value
func (o *PaymentInitiationAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationAddress) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *PaymentInitiationAddress) SetCountry(v string) {
	o.Country = v
}

func (o PaymentInitiationAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["street"] = o.Street
	}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode
	}
	if true {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationAddress) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationAddress := _PaymentInitiationAddress{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationAddress); err == nil {
		*o = PaymentInitiationAddress(varPaymentInitiationAddress)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "street")
		delete(additionalProperties, "city")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationAddress struct {
	value *PaymentInitiationAddress
	isSet bool
}

func (v NullablePaymentInitiationAddress) Get() *PaymentInitiationAddress {
	return v.value
}

func (v *NullablePaymentInitiationAddress) Set(val *PaymentInitiationAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationAddress(val *PaymentInitiationAddress) *NullablePaymentInitiationAddress {
	return &NullablePaymentInitiationAddress{value: val, isSet: true}
}

func (v NullablePaymentInitiationAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


