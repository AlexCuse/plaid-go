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

// EntityScreeningHitEmailsItems Analyzed emails for the associated hit
type EntityScreeningHitEmailsItems struct {
	Analysis *MatchSummary `json:"analysis,omitempty"`
	Data *EntityScreeningHitEmails `json:"data,omitempty"`
}

// NewEntityScreeningHitEmailsItems instantiates a new EntityScreeningHitEmailsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningHitEmailsItems() *EntityScreeningHitEmailsItems {
	this := EntityScreeningHitEmailsItems{}
	return &this
}

// NewEntityScreeningHitEmailsItemsWithDefaults instantiates a new EntityScreeningHitEmailsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningHitEmailsItemsWithDefaults() *EntityScreeningHitEmailsItems {
	this := EntityScreeningHitEmailsItems{}
	return &this
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *EntityScreeningHitEmailsItems) GetAnalysis() MatchSummary {
	if o == nil || o.Analysis == nil {
		var ret MatchSummary
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitEmailsItems) GetAnalysisOk() (*MatchSummary, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *EntityScreeningHitEmailsItems) HasAnalysis() bool {
	if o != nil && o.Analysis != nil {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given MatchSummary and assigns it to the Analysis field.
func (o *EntityScreeningHitEmailsItems) SetAnalysis(v MatchSummary) {
	o.Analysis = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EntityScreeningHitEmailsItems) GetData() EntityScreeningHitEmails {
	if o == nil || o.Data == nil {
		var ret EntityScreeningHitEmails
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitEmailsItems) GetDataOk() (*EntityScreeningHitEmails, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EntityScreeningHitEmailsItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given EntityScreeningHitEmails and assigns it to the Data field.
func (o *EntityScreeningHitEmailsItems) SetData(v EntityScreeningHitEmails) {
	o.Data = &v
}

func (o EntityScreeningHitEmailsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableEntityScreeningHitEmailsItems struct {
	value *EntityScreeningHitEmailsItems
	isSet bool
}

func (v NullableEntityScreeningHitEmailsItems) Get() *EntityScreeningHitEmailsItems {
	return v.value
}

func (v *NullableEntityScreeningHitEmailsItems) Set(val *EntityScreeningHitEmailsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningHitEmailsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningHitEmailsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningHitEmailsItems(val *EntityScreeningHitEmailsItems) *NullableEntityScreeningHitEmailsItems {
	return &NullableEntityScreeningHitEmailsItems{value: val, isSet: true}
}

func (v NullableEntityScreeningHitEmailsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningHitEmailsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


