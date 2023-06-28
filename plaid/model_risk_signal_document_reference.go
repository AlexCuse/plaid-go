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

// RiskSignalDocumentReference Object containing metadata for the document
type RiskSignalDocumentReference struct {
	// An identifier of the document referenced by the document metadata.
	DocumentId NullableString `json:"document_id,omitempty"`
	// The name of the document
	DocumentName *string `json:"document_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskSignalDocumentReference RiskSignalDocumentReference

// NewRiskSignalDocumentReference instantiates a new RiskSignalDocumentReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskSignalDocumentReference() *RiskSignalDocumentReference {
	this := RiskSignalDocumentReference{}
	return &this
}

// NewRiskSignalDocumentReferenceWithDefaults instantiates a new RiskSignalDocumentReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskSignalDocumentReferenceWithDefaults() *RiskSignalDocumentReference {
	this := RiskSignalDocumentReference{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskSignalDocumentReference) GetDocumentId() string {
	if o == nil || o.DocumentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DocumentId.Get()
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskSignalDocumentReference) GetDocumentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentId.Get(), o.DocumentId.IsSet()
}

// HasDocumentId returns a boolean if a field has been set.
func (o *RiskSignalDocumentReference) HasDocumentId() bool {
	if o != nil && o.DocumentId.IsSet() {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given NullableString and assigns it to the DocumentId field.
func (o *RiskSignalDocumentReference) SetDocumentId(v string) {
	o.DocumentId.Set(&v)
}
// SetDocumentIdNil sets the value for DocumentId to be an explicit nil
func (o *RiskSignalDocumentReference) SetDocumentIdNil() {
	o.DocumentId.Set(nil)
}

// UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil
func (o *RiskSignalDocumentReference) UnsetDocumentId() {
	o.DocumentId.Unset()
}

// GetDocumentName returns the DocumentName field value if set, zero value otherwise.
func (o *RiskSignalDocumentReference) GetDocumentName() string {
	if o == nil || o.DocumentName == nil {
		var ret string
		return ret
	}
	return *o.DocumentName
}

// GetDocumentNameOk returns a tuple with the DocumentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskSignalDocumentReference) GetDocumentNameOk() (*string, bool) {
	if o == nil || o.DocumentName == nil {
		return nil, false
	}
	return o.DocumentName, true
}

// HasDocumentName returns a boolean if a field has been set.
func (o *RiskSignalDocumentReference) HasDocumentName() bool {
	if o != nil && o.DocumentName != nil {
		return true
	}

	return false
}

// SetDocumentName gets a reference to the given string and assigns it to the DocumentName field.
func (o *RiskSignalDocumentReference) SetDocumentName(v string) {
	o.DocumentName = &v
}

func (o RiskSignalDocumentReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId.IsSet() {
		toSerialize["document_id"] = o.DocumentId.Get()
	}
	if o.DocumentName != nil {
		toSerialize["document_name"] = o.DocumentName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskSignalDocumentReference) UnmarshalJSON(bytes []byte) (err error) {
	varRiskSignalDocumentReference := _RiskSignalDocumentReference{}

	if err = json.Unmarshal(bytes, &varRiskSignalDocumentReference); err == nil {
		*o = RiskSignalDocumentReference(varRiskSignalDocumentReference)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "document_id")
		delete(additionalProperties, "document_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskSignalDocumentReference struct {
	value *RiskSignalDocumentReference
	isSet bool
}

func (v NullableRiskSignalDocumentReference) Get() *RiskSignalDocumentReference {
	return v.value
}

func (v *NullableRiskSignalDocumentReference) Set(val *RiskSignalDocumentReference) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSignalDocumentReference) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSignalDocumentReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSignalDocumentReference(val *RiskSignalDocumentReference) *NullableRiskSignalDocumentReference {
	return &NullableRiskSignalDocumentReference{value: val, isSet: true}
}

func (v NullableRiskSignalDocumentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSignalDocumentReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


