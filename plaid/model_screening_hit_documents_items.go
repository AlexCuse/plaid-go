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

// ScreeningHitDocumentsItems Analyzed document information for the associated hit
type ScreeningHitDocumentsItems struct {
	Analysis *MatchSummary `json:"analysis,omitempty"`
	Data *WatchlistScreeningDocument `json:"data,omitempty"`
}

// NewScreeningHitDocumentsItems instantiates a new ScreeningHitDocumentsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreeningHitDocumentsItems() *ScreeningHitDocumentsItems {
	this := ScreeningHitDocumentsItems{}
	return &this
}

// NewScreeningHitDocumentsItemsWithDefaults instantiates a new ScreeningHitDocumentsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreeningHitDocumentsItemsWithDefaults() *ScreeningHitDocumentsItems {
	this := ScreeningHitDocumentsItems{}
	return &this
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *ScreeningHitDocumentsItems) GetAnalysis() MatchSummary {
	if o == nil || o.Analysis == nil {
		var ret MatchSummary
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitDocumentsItems) GetAnalysisOk() (*MatchSummary, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *ScreeningHitDocumentsItems) HasAnalysis() bool {
	if o != nil && o.Analysis != nil {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given MatchSummary and assigns it to the Analysis field.
func (o *ScreeningHitDocumentsItems) SetAnalysis(v MatchSummary) {
	o.Analysis = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ScreeningHitDocumentsItems) GetData() WatchlistScreeningDocument {
	if o == nil || o.Data == nil {
		var ret WatchlistScreeningDocument
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitDocumentsItems) GetDataOk() (*WatchlistScreeningDocument, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ScreeningHitDocumentsItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given WatchlistScreeningDocument and assigns it to the Data field.
func (o *ScreeningHitDocumentsItems) SetData(v WatchlistScreeningDocument) {
	o.Data = &v
}

func (o ScreeningHitDocumentsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableScreeningHitDocumentsItems struct {
	value *ScreeningHitDocumentsItems
	isSet bool
}

func (v NullableScreeningHitDocumentsItems) Get() *ScreeningHitDocumentsItems {
	return v.value
}

func (v *NullableScreeningHitDocumentsItems) Set(val *ScreeningHitDocumentsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableScreeningHitDocumentsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableScreeningHitDocumentsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreeningHitDocumentsItems(val *ScreeningHitDocumentsItems) *NullableScreeningHitDocumentsItems {
	return &NullableScreeningHitDocumentsItems{value: val, isSet: true}
}

func (v NullableScreeningHitDocumentsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreeningHitDocumentsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


