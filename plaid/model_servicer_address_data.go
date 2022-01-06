/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.61.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ServicerAddressData The address of the student loan servicer. This is generally the remittance address to which payments should be sent.
type ServicerAddressData struct {
	// The full city name
	City NullableString `json:"city"`
	// The region or state Example: `\"NC\"`
	Region NullableString `json:"region"`
	// The full street address Example: `\"564 Main Street, APT 15\"`
	Street NullableString `json:"street"`
	// The postal code
	PostalCode NullableString `json:"postal_code"`
	// The ISO 3166-1 alpha-2 country code
	Country NullableString `json:"country"`
	AdditionalProperties map[string]interface{}
}

type _ServicerAddressData ServicerAddressData

// NewServicerAddressData instantiates a new ServicerAddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicerAddressData(city NullableString, region NullableString, street NullableString, postalCode NullableString, country NullableString) *ServicerAddressData {
	this := ServicerAddressData{}
	this.City = city
	this.Region = region
	this.Street = street
	this.PostalCode = postalCode
	this.Country = country
	return &this
}

// NewServicerAddressDataWithDefaults instantiates a new ServicerAddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicerAddressDataWithDefaults() *ServicerAddressData {
	this := ServicerAddressData{}
	return &this
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServicerAddressData) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicerAddressData) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *ServicerAddressData) SetCity(v string) {
	o.City.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServicerAddressData) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicerAddressData) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *ServicerAddressData) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetStreet returns the Street field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServicerAddressData) GetStreet() string {
	if o == nil || o.Street.Get() == nil {
		var ret string
		return ret
	}

	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicerAddressData) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// SetStreet sets field value
func (o *ServicerAddressData) SetStreet(v string) {
	o.Street.Set(&v)
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServicerAddressData) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicerAddressData) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *ServicerAddressData) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServicerAddressData) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicerAddressData) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *ServicerAddressData) SetCountry(v string) {
	o.Country.Set(&v)
}

func (o ServicerAddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["city"] = o.City.Get()
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["street"] = o.Street.Get()
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if true {
		toSerialize["country"] = o.Country.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicerAddressData) UnmarshalJSON(bytes []byte) (err error) {
	varServicerAddressData := _ServicerAddressData{}

	if err = json.Unmarshal(bytes, &varServicerAddressData); err == nil {
		*o = ServicerAddressData(varServicerAddressData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "street")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicerAddressData struct {
	value *ServicerAddressData
	isSet bool
}

func (v NullableServicerAddressData) Get() *ServicerAddressData {
	return v.value
}

func (v *NullableServicerAddressData) Set(val *ServicerAddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableServicerAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableServicerAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicerAddressData(val *ServicerAddressData) *NullableServicerAddressData {
	return &NullableServicerAddressData{value: val, isSet: true}
}

func (v NullableServicerAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicerAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


