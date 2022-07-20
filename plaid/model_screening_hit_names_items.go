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

// ScreeningHitNamesItems Analyzed name information for the associated hit
type ScreeningHitNamesItems struct {
	Analysis *MatchSummary `json:"analysis,omitempty"`
	Data *IndividualScreeningHitNames `json:"data,omitempty"`
}

// NewScreeningHitNamesItems instantiates a new ScreeningHitNamesItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreeningHitNamesItems() *ScreeningHitNamesItems {
	this := ScreeningHitNamesItems{}
	return &this
}

// NewScreeningHitNamesItemsWithDefaults instantiates a new ScreeningHitNamesItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreeningHitNamesItemsWithDefaults() *ScreeningHitNamesItems {
	this := ScreeningHitNamesItems{}
	return &this
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *ScreeningHitNamesItems) GetAnalysis() MatchSummary {
	if o == nil || o.Analysis == nil {
		var ret MatchSummary
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitNamesItems) GetAnalysisOk() (*MatchSummary, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *ScreeningHitNamesItems) HasAnalysis() bool {
	if o != nil && o.Analysis != nil {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given MatchSummary and assigns it to the Analysis field.
func (o *ScreeningHitNamesItems) SetAnalysis(v MatchSummary) {
	o.Analysis = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ScreeningHitNamesItems) GetData() IndividualScreeningHitNames {
	if o == nil || o.Data == nil {
		var ret IndividualScreeningHitNames
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitNamesItems) GetDataOk() (*IndividualScreeningHitNames, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ScreeningHitNamesItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given IndividualScreeningHitNames and assigns it to the Data field.
func (o *ScreeningHitNamesItems) SetData(v IndividualScreeningHitNames) {
	o.Data = &v
}

func (o ScreeningHitNamesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableScreeningHitNamesItems struct {
	value *ScreeningHitNamesItems
	isSet bool
}

func (v NullableScreeningHitNamesItems) Get() *ScreeningHitNamesItems {
	return v.value
}

func (v *NullableScreeningHitNamesItems) Set(val *ScreeningHitNamesItems) {
	v.value = val
	v.isSet = true
}

func (v NullableScreeningHitNamesItems) IsSet() bool {
	return v.isSet
}

func (v *NullableScreeningHitNamesItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreeningHitNamesItems(val *ScreeningHitNamesItems) *NullableScreeningHitNamesItems {
	return &NullableScreeningHitNamesItems{value: val, isSet: true}
}

func (v NullableScreeningHitNamesItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreeningHitNamesItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


