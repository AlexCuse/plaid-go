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

// WatchlistScreeningSearchTerms Search terms for creating an individual watchlist screening
type WatchlistScreeningSearchTerms struct {
	// ID of the associated program.
	WatchlistProgramId string `json:"watchlist_program_id"`
	// The legal name of the individual being screened.
	LegalName string `json:"legal_name"`
	DateOfBirth NullableString `json:"date_of_birth"`
	DocumentNumber NullableString `json:"document_number"`
	Country NullableString `json:"country"`
	// The current version of the search terms. Starts at `1` and increments with each edit to `search_terms`.
	Version float32 `json:"version"`
}

// NewWatchlistScreeningSearchTerms instantiates a new WatchlistScreeningSearchTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningSearchTerms(watchlistProgramId string, legalName string, dateOfBirth NullableString, documentNumber NullableString, country NullableString, version float32) *WatchlistScreeningSearchTerms {
	this := WatchlistScreeningSearchTerms{}
	this.WatchlistProgramId = watchlistProgramId
	this.LegalName = legalName
	this.DateOfBirth = dateOfBirth
	this.DocumentNumber = documentNumber
	this.Country = country
	this.Version = version
	return &this
}

// NewWatchlistScreeningSearchTermsWithDefaults instantiates a new WatchlistScreeningSearchTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningSearchTermsWithDefaults() *WatchlistScreeningSearchTerms {
	this := WatchlistScreeningSearchTerms{}
	return &this
}

// GetWatchlistProgramId returns the WatchlistProgramId field value
func (o *WatchlistScreeningSearchTerms) GetWatchlistProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WatchlistProgramId
}

// GetWatchlistProgramIdOk returns a tuple with the WatchlistProgramId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningSearchTerms) GetWatchlistProgramIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistProgramId, true
}

// SetWatchlistProgramId sets field value
func (o *WatchlistScreeningSearchTerms) SetWatchlistProgramId(v string) {
	o.WatchlistProgramId = v
}

// GetLegalName returns the LegalName field value
func (o *WatchlistScreeningSearchTerms) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningSearchTerms) GetLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *WatchlistScreeningSearchTerms) SetLegalName(v string) {
	o.LegalName = v
}

// GetDateOfBirth returns the DateOfBirth field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningSearchTerms) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth.Get() == nil {
		var ret string
		return ret
	}

	return *o.DateOfBirth.Get()
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningSearchTerms) GetDateOfBirthOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateOfBirth.Get(), o.DateOfBirth.IsSet()
}

// SetDateOfBirth sets field value
func (o *WatchlistScreeningSearchTerms) SetDateOfBirth(v string) {
	o.DateOfBirth.Set(&v)
}

// GetDocumentNumber returns the DocumentNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningSearchTerms) GetDocumentNumber() string {
	if o == nil || o.DocumentNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentNumber.Get()
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningSearchTerms) GetDocumentNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentNumber.Get(), o.DocumentNumber.IsSet()
}

// SetDocumentNumber sets field value
func (o *WatchlistScreeningSearchTerms) SetDocumentNumber(v string) {
	o.DocumentNumber.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningSearchTerms) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningSearchTerms) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *WatchlistScreeningSearchTerms) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetVersion returns the Version field value
func (o *WatchlistScreeningSearchTerms) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningSearchTerms) GetVersionOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *WatchlistScreeningSearchTerms) SetVersion(v float32) {
	o.Version = v
}

func (o WatchlistScreeningSearchTerms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["watchlist_program_id"] = o.WatchlistProgramId
	}
	if true {
		toSerialize["legal_name"] = o.LegalName
	}
	if true {
		toSerialize["date_of_birth"] = o.DateOfBirth.Get()
	}
	if true {
		toSerialize["document_number"] = o.DocumentNumber.Get()
	}
	if true {
		toSerialize["country"] = o.Country.Get()
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistScreeningSearchTerms struct {
	value *WatchlistScreeningSearchTerms
	isSet bool
}

func (v NullableWatchlistScreeningSearchTerms) Get() *WatchlistScreeningSearchTerms {
	return v.value
}

func (v *NullableWatchlistScreeningSearchTerms) Set(val *WatchlistScreeningSearchTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningSearchTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningSearchTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningSearchTerms(val *WatchlistScreeningSearchTerms) *NullableWatchlistScreeningSearchTerms {
	return &NullableWatchlistScreeningSearchTerms{value: val, isSet: true}
}

func (v NullableWatchlistScreeningSearchTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningSearchTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


