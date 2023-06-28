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

// ExternalPaymentScheduleRequest The schedule that the payment will be executed on. If a schedule is provided, the payment is automatically set up as a standing order. If no schedule is specified, the payment will be executed only once.
type ExternalPaymentScheduleRequest struct {
	Interval PaymentScheduleInterval `json:"interval"`
	// The day of the interval on which to schedule the payment.  If the payment interval is weekly, `interval_execution_day` should be an integer from 1 (Monday) to 7 (Sunday).  If the payment interval is monthly, `interval_execution_day` should be an integer indicating which day of the month to make the payment on. Integers from 1 to 28 can be used to make a payment on that day of the month. Negative integers from -1 to -5 can be used to make a payment relative to the end of the month. To make a payment on the last day of the month, use -1; to make the payment on the second-to-last day, use -2, and so on.
	IntervalExecutionDay int32 `json:"interval_execution_day"`
	// A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will begin on the first `interval_execution_day` on or after the `start_date`.  If the first `interval_execution_day` on or after the start date is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so.
	StartDate string `json:"start_date"`
	// A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will end on the last `interval_execution_day` on or before the `end_date`. If the only `interval_execution_day` between the start date and the end date (inclusive) is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make a payment on that day, but it is not guaranteed to do so.
	EndDate NullableString `json:"end_date,omitempty"`
	// The start date sent to the bank after adjusting for holidays or weekends.  Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, this field will be `null`.
	AdjustedStartDate NullableString `json:"adjusted_start_date,omitempty"`
}

// NewExternalPaymentScheduleRequest instantiates a new ExternalPaymentScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentScheduleRequest(interval PaymentScheduleInterval, intervalExecutionDay int32, startDate string) *ExternalPaymentScheduleRequest {
	this := ExternalPaymentScheduleRequest{}
	this.Interval = interval
	this.IntervalExecutionDay = intervalExecutionDay
	this.StartDate = startDate
	return &this
}

// NewExternalPaymentScheduleRequestWithDefaults instantiates a new ExternalPaymentScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentScheduleRequestWithDefaults() *ExternalPaymentScheduleRequest {
	this := ExternalPaymentScheduleRequest{}
	return &this
}

// GetInterval returns the Interval field value
func (o *ExternalPaymentScheduleRequest) GetInterval() PaymentScheduleInterval {
	if o == nil {
		var ret PaymentScheduleInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentScheduleRequest) GetIntervalOk() (*PaymentScheduleInterval, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ExternalPaymentScheduleRequest) SetInterval(v PaymentScheduleInterval) {
	o.Interval = v
}

// GetIntervalExecutionDay returns the IntervalExecutionDay field value
func (o *ExternalPaymentScheduleRequest) GetIntervalExecutionDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntervalExecutionDay
}

// GetIntervalExecutionDayOk returns a tuple with the IntervalExecutionDay field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentScheduleRequest) GetIntervalExecutionDayOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IntervalExecutionDay, true
}

// SetIntervalExecutionDay sets field value
func (o *ExternalPaymentScheduleRequest) SetIntervalExecutionDay(v int32) {
	o.IntervalExecutionDay = v
}

// GetStartDate returns the StartDate field value
func (o *ExternalPaymentScheduleRequest) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentScheduleRequest) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *ExternalPaymentScheduleRequest) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalPaymentScheduleRequest) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentScheduleRequest) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *ExternalPaymentScheduleRequest) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *ExternalPaymentScheduleRequest) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *ExternalPaymentScheduleRequest) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *ExternalPaymentScheduleRequest) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetAdjustedStartDate returns the AdjustedStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalPaymentScheduleRequest) GetAdjustedStartDate() string {
	if o == nil || o.AdjustedStartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.AdjustedStartDate.Get()
}

// GetAdjustedStartDateOk returns a tuple with the AdjustedStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentScheduleRequest) GetAdjustedStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdjustedStartDate.Get(), o.AdjustedStartDate.IsSet()
}

// HasAdjustedStartDate returns a boolean if a field has been set.
func (o *ExternalPaymentScheduleRequest) HasAdjustedStartDate() bool {
	if o != nil && o.AdjustedStartDate.IsSet() {
		return true
	}

	return false
}

// SetAdjustedStartDate gets a reference to the given NullableString and assigns it to the AdjustedStartDate field.
func (o *ExternalPaymentScheduleRequest) SetAdjustedStartDate(v string) {
	o.AdjustedStartDate.Set(&v)
}
// SetAdjustedStartDateNil sets the value for AdjustedStartDate to be an explicit nil
func (o *ExternalPaymentScheduleRequest) SetAdjustedStartDateNil() {
	o.AdjustedStartDate.Set(nil)
}

// UnsetAdjustedStartDate ensures that no value is present for AdjustedStartDate, not even an explicit nil
func (o *ExternalPaymentScheduleRequest) UnsetAdjustedStartDate() {
	o.AdjustedStartDate.Unset()
}

func (o ExternalPaymentScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if true {
		toSerialize["interval_execution_day"] = o.IntervalExecutionDay
	}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.AdjustedStartDate.IsSet() {
		toSerialize["adjusted_start_date"] = o.AdjustedStartDate.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPaymentScheduleRequest struct {
	value *ExternalPaymentScheduleRequest
	isSet bool
}

func (v NullableExternalPaymentScheduleRequest) Get() *ExternalPaymentScheduleRequest {
	return v.value
}

func (v *NullableExternalPaymentScheduleRequest) Set(val *ExternalPaymentScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentScheduleRequest(val *ExternalPaymentScheduleRequest) *NullableExternalPaymentScheduleRequest {
	return &NullableExternalPaymentScheduleRequest{value: val, isSet: true}
}

func (v NullableExternalPaymentScheduleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


