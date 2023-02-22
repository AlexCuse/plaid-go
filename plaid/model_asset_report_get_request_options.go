/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AssetReportGetRequestOptions An optional object to filter or add data to `/asset_report/get` results. If provided, must be non-`null`.
type AssetReportGetRequestOptions struct {
	// The maximum integer number of days of history to include in the Asset Report.
	DaysToInclude NullableInt32 `json:"days_to_include,omitempty"`
}

// NewAssetReportGetRequestOptions instantiates a new AssetReportGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportGetRequestOptions() *AssetReportGetRequestOptions {
	this := AssetReportGetRequestOptions{}
	return &this
}

// NewAssetReportGetRequestOptionsWithDefaults instantiates a new AssetReportGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportGetRequestOptionsWithDefaults() *AssetReportGetRequestOptions {
	this := AssetReportGetRequestOptions{}
	return &this
}

// GetDaysToInclude returns the DaysToInclude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportGetRequestOptions) GetDaysToInclude() int32 {
	if o == nil || o.DaysToInclude.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DaysToInclude.Get()
}

// GetDaysToIncludeOk returns a tuple with the DaysToInclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportGetRequestOptions) GetDaysToIncludeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DaysToInclude.Get(), o.DaysToInclude.IsSet()
}

// HasDaysToInclude returns a boolean if a field has been set.
func (o *AssetReportGetRequestOptions) HasDaysToInclude() bool {
	if o != nil && o.DaysToInclude.IsSet() {
		return true
	}

	return false
}

// SetDaysToInclude gets a reference to the given NullableInt32 and assigns it to the DaysToInclude field.
func (o *AssetReportGetRequestOptions) SetDaysToInclude(v int32) {
	o.DaysToInclude.Set(&v)
}
// SetDaysToIncludeNil sets the value for DaysToInclude to be an explicit nil
func (o *AssetReportGetRequestOptions) SetDaysToIncludeNil() {
	o.DaysToInclude.Set(nil)
}

// UnsetDaysToInclude ensures that no value is present for DaysToInclude, not even an explicit nil
func (o *AssetReportGetRequestOptions) UnsetDaysToInclude() {
	o.DaysToInclude.Unset()
}

func (o AssetReportGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DaysToInclude.IsSet() {
		toSerialize["days_to_include"] = o.DaysToInclude.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportGetRequestOptions struct {
	value *AssetReportGetRequestOptions
	isSet bool
}

func (v NullableAssetReportGetRequestOptions) Get() *AssetReportGetRequestOptions {
	return v.value
}

func (v *NullableAssetReportGetRequestOptions) Set(val *AssetReportGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportGetRequestOptions(val *AssetReportGetRequestOptions) *NullableAssetReportGetRequestOptions {
	return &NullableAssetReportGetRequestOptions{value: val, isSet: true}
}

func (v NullableAssetReportGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


