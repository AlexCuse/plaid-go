/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IdentityVerificationTemplateReference The resource ID and version number of the template configuring the behavior of a given identity verification.
type IdentityVerificationTemplateReference struct {
	// ID of the associated Identity Verification template.
	Id string `json:"id"`
	// Version of the associated Identity Verification template.
	Version float32 `json:"version"`
}

// NewIdentityVerificationTemplateReference instantiates a new IdentityVerificationTemplateReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationTemplateReference(id string, version float32) *IdentityVerificationTemplateReference {
	this := IdentityVerificationTemplateReference{}
	this.Id = id
	this.Version = version
	return &this
}

// NewIdentityVerificationTemplateReferenceWithDefaults instantiates a new IdentityVerificationTemplateReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationTemplateReferenceWithDefaults() *IdentityVerificationTemplateReference {
	this := IdentityVerificationTemplateReference{}
	return &this
}

// GetId returns the Id field value
func (o *IdentityVerificationTemplateReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationTemplateReference) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdentityVerificationTemplateReference) SetId(v string) {
	o.Id = v
}

// GetVersion returns the Version field value
func (o *IdentityVerificationTemplateReference) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationTemplateReference) GetVersionOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *IdentityVerificationTemplateReference) SetVersion(v float32) {
	o.Version = v
}

func (o IdentityVerificationTemplateReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityVerificationTemplateReference struct {
	value *IdentityVerificationTemplateReference
	isSet bool
}

func (v NullableIdentityVerificationTemplateReference) Get() *IdentityVerificationTemplateReference {
	return v.value
}

func (v *NullableIdentityVerificationTemplateReference) Set(val *IdentityVerificationTemplateReference) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationTemplateReference) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationTemplateReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationTemplateReference(val *IdentityVerificationTemplateReference) *NullableIdentityVerificationTemplateReference {
	return &NullableIdentityVerificationTemplateReference{value: val, isSet: true}
}

func (v NullableIdentityVerificationTemplateReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationTemplateReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


