/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.413.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// BaseReportLongestGapInsights Largest number of days between sequential transactions per calendar month
type BaseReportLongestGapInsights struct {
	// The start date of this time period. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	StartDate *string `json:"start_date,omitempty"`
	// The end date of this time period. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	EndDate *string `json:"end_date,omitempty"`
	// Largest number of days between sequential transactions for this time period.
	Days *int32 `json:"days,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseReportLongestGapInsights BaseReportLongestGapInsights

// NewBaseReportLongestGapInsights instantiates a new BaseReportLongestGapInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseReportLongestGapInsights() *BaseReportLongestGapInsights {
	this := BaseReportLongestGapInsights{}
	return &this
}

// NewBaseReportLongestGapInsightsWithDefaults instantiates a new BaseReportLongestGapInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseReportLongestGapInsightsWithDefaults() *BaseReportLongestGapInsights {
	this := BaseReportLongestGapInsights{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BaseReportLongestGapInsights) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportLongestGapInsights) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BaseReportLongestGapInsights) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *BaseReportLongestGapInsights) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BaseReportLongestGapInsights) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportLongestGapInsights) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BaseReportLongestGapInsights) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *BaseReportLongestGapInsights) SetEndDate(v string) {
	o.EndDate = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *BaseReportLongestGapInsights) GetDays() int32 {
	if o == nil || o.Days == nil {
		var ret int32
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportLongestGapInsights) GetDaysOk() (*int32, bool) {
	if o == nil || o.Days == nil {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *BaseReportLongestGapInsights) HasDays() bool {
	if o != nil && o.Days != nil {
		return true
	}

	return false
}

// SetDays gets a reference to the given int32 and assigns it to the Days field.
func (o *BaseReportLongestGapInsights) SetDays(v int32) {
	o.Days = &v
}

func (o BaseReportLongestGapInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.Days != nil {
		toSerialize["days"] = o.Days
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseReportLongestGapInsights) UnmarshalJSON(bytes []byte) (err error) {
	varBaseReportLongestGapInsights := _BaseReportLongestGapInsights{}

	if err = json.Unmarshal(bytes, &varBaseReportLongestGapInsights); err == nil {
		*o = BaseReportLongestGapInsights(varBaseReportLongestGapInsights)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "days")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseReportLongestGapInsights struct {
	value *BaseReportLongestGapInsights
	isSet bool
}

func (v NullableBaseReportLongestGapInsights) Get() *BaseReportLongestGapInsights {
	return v.value
}

func (v *NullableBaseReportLongestGapInsights) Set(val *BaseReportLongestGapInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseReportLongestGapInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseReportLongestGapInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseReportLongestGapInsights(val *BaseReportLongestGapInsights) *NullableBaseReportLongestGapInsights {
	return &NullableBaseReportLongestGapInsights{value: val, isSet: true}
}

func (v NullableBaseReportLongestGapInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseReportLongestGapInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


