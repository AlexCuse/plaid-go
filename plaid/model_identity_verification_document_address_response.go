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

// IdentityVerificationDocumentAddressResponse The address extracted from the document. The address must at least contain the following fields to be a valid address: `street`, `city`, `country`. If any are missing or unable to be extracted, the address will be null.  `region`, and `postal_code` may be null based on the addressing system. For example:  Addresses from the United Kingdom will not include a region  Addresses from Hong Kong will not include postal code  Note: Optical Character Recognition (OCR) technology may sometimes extract incorrect data from a document.
type IdentityVerificationDocumentAddressResponse struct {
	// The full street address extracted from the document.
	Street string `json:"street"`
	// City extracted from the document.
	City string `json:"city"`
	// An ISO 3166-2 subdivision code extracted from the document. Related terms would be \"state\", \"province\", \"prefecture\", \"zone\", \"subdivision\", etc.
	Region NullableString `json:"region"`
	// The postal code extracted from the document. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.
	PostalCode NullableString `json:"postal_code"`
	// Valid, capitalized, two-letter ISO code representing the country extracted from the document. Must be in ISO 3166-1 alpha-2 form.
	Country string `json:"country"`
	AdditionalProperties map[string]interface{}
}

type _IdentityVerificationDocumentAddressResponse IdentityVerificationDocumentAddressResponse

// NewIdentityVerificationDocumentAddressResponse instantiates a new IdentityVerificationDocumentAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationDocumentAddressResponse(street string, city string, region NullableString, postalCode NullableString, country string) *IdentityVerificationDocumentAddressResponse {
	this := IdentityVerificationDocumentAddressResponse{}
	this.Street = street
	this.City = city
	this.Region = region
	this.PostalCode = postalCode
	this.Country = country
	return &this
}

// NewIdentityVerificationDocumentAddressResponseWithDefaults instantiates a new IdentityVerificationDocumentAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationDocumentAddressResponseWithDefaults() *IdentityVerificationDocumentAddressResponse {
	this := IdentityVerificationDocumentAddressResponse{}
	return &this
}

// GetStreet returns the Street field value
func (o *IdentityVerificationDocumentAddressResponse) GetStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationDocumentAddressResponse) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *IdentityVerificationDocumentAddressResponse) SetStreet(v string) {
	o.Street = v
}

// GetCity returns the City field value
func (o *IdentityVerificationDocumentAddressResponse) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationDocumentAddressResponse) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *IdentityVerificationDocumentAddressResponse) SetCity(v string) {
	o.City = v
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationDocumentAddressResponse) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationDocumentAddressResponse) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *IdentityVerificationDocumentAddressResponse) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationDocumentAddressResponse) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationDocumentAddressResponse) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *IdentityVerificationDocumentAddressResponse) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// GetCountry returns the Country field value
func (o *IdentityVerificationDocumentAddressResponse) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationDocumentAddressResponse) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *IdentityVerificationDocumentAddressResponse) SetCountry(v string) {
	o.Country = v
}

func (o IdentityVerificationDocumentAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["street"] = o.Street
	}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if true {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityVerificationDocumentAddressResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityVerificationDocumentAddressResponse := _IdentityVerificationDocumentAddressResponse{}

	if err = json.Unmarshal(bytes, &varIdentityVerificationDocumentAddressResponse); err == nil {
		*o = IdentityVerificationDocumentAddressResponse(varIdentityVerificationDocumentAddressResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "street")
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityVerificationDocumentAddressResponse struct {
	value *IdentityVerificationDocumentAddressResponse
	isSet bool
}

func (v NullableIdentityVerificationDocumentAddressResponse) Get() *IdentityVerificationDocumentAddressResponse {
	return v.value
}

func (v *NullableIdentityVerificationDocumentAddressResponse) Set(val *IdentityVerificationDocumentAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationDocumentAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationDocumentAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationDocumentAddressResponse(val *IdentityVerificationDocumentAddressResponse) *NullableIdentityVerificationDocumentAddressResponse {
	return &NullableIdentityVerificationDocumentAddressResponse{value: val, isSet: true}
}

func (v NullableIdentityVerificationDocumentAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationDocumentAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

