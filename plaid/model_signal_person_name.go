/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.131.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SignalPersonName The user's legal name
type SignalPersonName struct {
	// The user's name prefix (e.g. \"Mr.\")
	Prefix NullableString `json:"prefix,omitempty"`
	// The user's given name. If the user has a one-word name, it should be provided in this field.
	GivenName NullableString `json:"given_name,omitempty"`
	// The user's middle name
	MiddleName NullableString `json:"middle_name,omitempty"`
	// The user's family name / surname
	FamilyName NullableString `json:"family_name,omitempty"`
	// The user's name suffix (e.g. \"II\")
	Suffix NullableString `json:"suffix,omitempty"`
}

// NewSignalPersonName instantiates a new SignalPersonName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalPersonName() *SignalPersonName {
	this := SignalPersonName{}
	return &this
}

// NewSignalPersonNameWithDefaults instantiates a new SignalPersonName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalPersonNameWithDefaults() *SignalPersonName {
	this := SignalPersonName{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalPersonName) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalPersonName) GetPrefixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *SignalPersonName) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given NullableString and assigns it to the Prefix field.
func (o *SignalPersonName) SetPrefix(v string) {
	o.Prefix.Set(&v)
}
// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *SignalPersonName) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *SignalPersonName) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetGivenName returns the GivenName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalPersonName) GetGivenName() string {
	if o == nil || o.GivenName.Get() == nil {
		var ret string
		return ret
	}
	return *o.GivenName.Get()
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalPersonName) GetGivenNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GivenName.Get(), o.GivenName.IsSet()
}

// HasGivenName returns a boolean if a field has been set.
func (o *SignalPersonName) HasGivenName() bool {
	if o != nil && o.GivenName.IsSet() {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given NullableString and assigns it to the GivenName field.
func (o *SignalPersonName) SetGivenName(v string) {
	o.GivenName.Set(&v)
}
// SetGivenNameNil sets the value for GivenName to be an explicit nil
func (o *SignalPersonName) SetGivenNameNil() {
	o.GivenName.Set(nil)
}

// UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
func (o *SignalPersonName) UnsetGivenName() {
	o.GivenName.Unset()
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalPersonName) GetMiddleName() string {
	if o == nil || o.MiddleName.Get() == nil {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalPersonName) GetMiddleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *SignalPersonName) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *SignalPersonName) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *SignalPersonName) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *SignalPersonName) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalPersonName) GetFamilyName() string {
	if o == nil || o.FamilyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FamilyName.Get()
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalPersonName) GetFamilyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FamilyName.Get(), o.FamilyName.IsSet()
}

// HasFamilyName returns a boolean if a field has been set.
func (o *SignalPersonName) HasFamilyName() bool {
	if o != nil && o.FamilyName.IsSet() {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given NullableString and assigns it to the FamilyName field.
func (o *SignalPersonName) SetFamilyName(v string) {
	o.FamilyName.Set(&v)
}
// SetFamilyNameNil sets the value for FamilyName to be an explicit nil
func (o *SignalPersonName) SetFamilyNameNil() {
	o.FamilyName.Set(nil)
}

// UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
func (o *SignalPersonName) UnsetFamilyName() {
	o.FamilyName.Unset()
}

// GetSuffix returns the Suffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalPersonName) GetSuffix() string {
	if o == nil || o.Suffix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Suffix.Get()
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalPersonName) GetSuffixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Suffix.Get(), o.Suffix.IsSet()
}

// HasSuffix returns a boolean if a field has been set.
func (o *SignalPersonName) HasSuffix() bool {
	if o != nil && o.Suffix.IsSet() {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given NullableString and assigns it to the Suffix field.
func (o *SignalPersonName) SetSuffix(v string) {
	o.Suffix.Set(&v)
}
// SetSuffixNil sets the value for Suffix to be an explicit nil
func (o *SignalPersonName) SetSuffixNil() {
	o.Suffix.Set(nil)
}

// UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
func (o *SignalPersonName) UnsetSuffix() {
	o.Suffix.Unset()
}

func (o SignalPersonName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	if o.GivenName.IsSet() {
		toSerialize["given_name"] = o.GivenName.Get()
	}
	if o.MiddleName.IsSet() {
		toSerialize["middle_name"] = o.MiddleName.Get()
	}
	if o.FamilyName.IsSet() {
		toSerialize["family_name"] = o.FamilyName.Get()
	}
	if o.Suffix.IsSet() {
		toSerialize["suffix"] = o.Suffix.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSignalPersonName struct {
	value *SignalPersonName
	isSet bool
}

func (v NullableSignalPersonName) Get() *SignalPersonName {
	return v.value
}

func (v *NullableSignalPersonName) Set(val *SignalPersonName) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalPersonName) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalPersonName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalPersonName(val *SignalPersonName) *NullableSignalPersonName {
	return &NullableSignalPersonName{value: val, isSet: true}
}

func (v NullableSignalPersonName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalPersonName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


