/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.385.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// EntityWatchlistScreeningSearchTerms Search terms associated with an entity used for searching against watchlists
type EntityWatchlistScreeningSearchTerms struct {
	// ID of the associated entity program.
	EntityWatchlistProgramId string `json:"entity_watchlist_program_id"`
	// The name of the organization being screened.
	LegalName string `json:"legal_name"`
	// The numeric or alphanumeric identifier associated with this document.
	DocumentNumber NullableString `json:"document_number"`
	// A valid email address.
	EmailAddress NullableString `json:"email_address"`
	// Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.
	Country NullableString `json:"country"`
	// A phone number in E.164 format.
	PhoneNumber NullableString `json:"phone_number"`
	// An 'http' or 'https' URL (must begin with either of those).
	Url NullableString `json:"url"`
	// The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.
	Version int32 `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _EntityWatchlistScreeningSearchTerms EntityWatchlistScreeningSearchTerms

// NewEntityWatchlistScreeningSearchTerms instantiates a new EntityWatchlistScreeningSearchTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityWatchlistScreeningSearchTerms(entityWatchlistProgramId string, legalName string, documentNumber NullableString, emailAddress NullableString, country NullableString, phoneNumber NullableString, url NullableString, version int32) *EntityWatchlistScreeningSearchTerms {
	this := EntityWatchlistScreeningSearchTerms{}
	this.EntityWatchlistProgramId = entityWatchlistProgramId
	this.LegalName = legalName
	this.DocumentNumber = documentNumber
	this.EmailAddress = emailAddress
	this.Country = country
	this.PhoneNumber = phoneNumber
	this.Url = url
	this.Version = version
	return &this
}

// NewEntityWatchlistScreeningSearchTermsWithDefaults instantiates a new EntityWatchlistScreeningSearchTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityWatchlistScreeningSearchTermsWithDefaults() *EntityWatchlistScreeningSearchTerms {
	this := EntityWatchlistScreeningSearchTerms{}
	return &this
}

// GetEntityWatchlistProgramId returns the EntityWatchlistProgramId field value
func (o *EntityWatchlistScreeningSearchTerms) GetEntityWatchlistProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityWatchlistProgramId
}

// GetEntityWatchlistProgramIdOk returns a tuple with the EntityWatchlistProgramId field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreeningSearchTerms) GetEntityWatchlistProgramIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistProgramId, true
}

// SetEntityWatchlistProgramId sets field value
func (o *EntityWatchlistScreeningSearchTerms) SetEntityWatchlistProgramId(v string) {
	o.EntityWatchlistProgramId = v
}

// GetLegalName returns the LegalName field value
func (o *EntityWatchlistScreeningSearchTerms) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreeningSearchTerms) GetLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *EntityWatchlistScreeningSearchTerms) SetLegalName(v string) {
	o.LegalName = v
}

// GetDocumentNumber returns the DocumentNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EntityWatchlistScreeningSearchTerms) GetDocumentNumber() string {
	if o == nil || o.DocumentNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentNumber.Get()
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistScreeningSearchTerms) GetDocumentNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentNumber.Get(), o.DocumentNumber.IsSet()
}

// SetDocumentNumber sets field value
func (o *EntityWatchlistScreeningSearchTerms) SetDocumentNumber(v string) {
	o.DocumentNumber.Set(&v)
}

// GetEmailAddress returns the EmailAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EntityWatchlistScreeningSearchTerms) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistScreeningSearchTerms) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// SetEmailAddress sets field value
func (o *EntityWatchlistScreeningSearchTerms) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EntityWatchlistScreeningSearchTerms) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistScreeningSearchTerms) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *EntityWatchlistScreeningSearchTerms) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetPhoneNumber returns the PhoneNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EntityWatchlistScreeningSearchTerms) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistScreeningSearchTerms) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// SetPhoneNumber sets field value
func (o *EntityWatchlistScreeningSearchTerms) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EntityWatchlistScreeningSearchTerms) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistScreeningSearchTerms) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *EntityWatchlistScreeningSearchTerms) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetVersion returns the Version field value
func (o *EntityWatchlistScreeningSearchTerms) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreeningSearchTerms) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EntityWatchlistScreeningSearchTerms) SetVersion(v int32) {
	o.Version = v
}

func (o EntityWatchlistScreeningSearchTerms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_watchlist_program_id"] = o.EntityWatchlistProgramId
	}
	if true {
		toSerialize["legal_name"] = o.LegalName
	}
	if true {
		toSerialize["document_number"] = o.DocumentNumber.Get()
	}
	if true {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if true {
		toSerialize["country"] = o.Country.Get()
	}
	if true {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if true {
		toSerialize["url"] = o.Url.Get()
	}
	if true {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityWatchlistScreeningSearchTerms) UnmarshalJSON(bytes []byte) (err error) {
	varEntityWatchlistScreeningSearchTerms := _EntityWatchlistScreeningSearchTerms{}

	if err = json.Unmarshal(bytes, &varEntityWatchlistScreeningSearchTerms); err == nil {
		*o = EntityWatchlistScreeningSearchTerms(varEntityWatchlistScreeningSearchTerms)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_watchlist_program_id")
		delete(additionalProperties, "legal_name")
		delete(additionalProperties, "document_number")
		delete(additionalProperties, "email_address")
		delete(additionalProperties, "country")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "url")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityWatchlistScreeningSearchTerms struct {
	value *EntityWatchlistScreeningSearchTerms
	isSet bool
}

func (v NullableEntityWatchlistScreeningSearchTerms) Get() *EntityWatchlistScreeningSearchTerms {
	return v.value
}

func (v *NullableEntityWatchlistScreeningSearchTerms) Set(val *EntityWatchlistScreeningSearchTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityWatchlistScreeningSearchTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityWatchlistScreeningSearchTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityWatchlistScreeningSearchTerms(val *EntityWatchlistScreeningSearchTerms) *NullableEntityWatchlistScreeningSearchTerms {
	return &NullableEntityWatchlistScreeningSearchTerms{value: val, isSet: true}
}

func (v NullableEntityWatchlistScreeningSearchTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityWatchlistScreeningSearchTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


