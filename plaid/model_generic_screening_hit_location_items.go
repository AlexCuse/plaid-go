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

// GenericScreeningHitLocationItems Analyzed location information for the associated hit
type GenericScreeningHitLocationItems struct {
	Analysis *MatchSummary `json:"analysis,omitempty"`
	Data *WatchlistScreeningHitLocations `json:"data,omitempty"`
}

// NewGenericScreeningHitLocationItems instantiates a new GenericScreeningHitLocationItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericScreeningHitLocationItems() *GenericScreeningHitLocationItems {
	this := GenericScreeningHitLocationItems{}
	return &this
}

// NewGenericScreeningHitLocationItemsWithDefaults instantiates a new GenericScreeningHitLocationItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericScreeningHitLocationItemsWithDefaults() *GenericScreeningHitLocationItems {
	this := GenericScreeningHitLocationItems{}
	return &this
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *GenericScreeningHitLocationItems) GetAnalysis() MatchSummary {
	if o == nil || o.Analysis == nil {
		var ret MatchSummary
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericScreeningHitLocationItems) GetAnalysisOk() (*MatchSummary, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *GenericScreeningHitLocationItems) HasAnalysis() bool {
	if o != nil && o.Analysis != nil {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given MatchSummary and assigns it to the Analysis field.
func (o *GenericScreeningHitLocationItems) SetAnalysis(v MatchSummary) {
	o.Analysis = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GenericScreeningHitLocationItems) GetData() WatchlistScreeningHitLocations {
	if o == nil || o.Data == nil {
		var ret WatchlistScreeningHitLocations
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericScreeningHitLocationItems) GetDataOk() (*WatchlistScreeningHitLocations, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GenericScreeningHitLocationItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given WatchlistScreeningHitLocations and assigns it to the Data field.
func (o *GenericScreeningHitLocationItems) SetData(v WatchlistScreeningHitLocations) {
	o.Data = &v
}

func (o GenericScreeningHitLocationItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGenericScreeningHitLocationItems struct {
	value *GenericScreeningHitLocationItems
	isSet bool
}

func (v NullableGenericScreeningHitLocationItems) Get() *GenericScreeningHitLocationItems {
	return v.value
}

func (v *NullableGenericScreeningHitLocationItems) Set(val *GenericScreeningHitLocationItems) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericScreeningHitLocationItems) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericScreeningHitLocationItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericScreeningHitLocationItems(val *GenericScreeningHitLocationItems) *NullableGenericScreeningHitLocationItems {
	return &NullableGenericScreeningHitLocationItems{value: val, isSet: true}
}

func (v NullableGenericScreeningHitLocationItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericScreeningHitLocationItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


