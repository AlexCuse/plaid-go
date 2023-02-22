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

// ExternalPaymentScheduleGet The schedule that the payment will be executed on. If a schedule is provided, the payment is automatically set up as a standing order. If no schedule is specified, the payment will be executed only once.
type ExternalPaymentScheduleGet struct {
	Interval PaymentScheduleInterval `json:"interval"`
	// The day of the interval on which to schedule the payment.  If the payment interval is weekly, `interval_execution_day` should be an integer from 1 (Monday) to 7 (Sunday).  If the payment interval is monthly, `interval_execution_day` should be an integer indicating which day of the month to make the payment on. Integers from 1 to 28 can be used to make a payment on that day of the month. Negative integers from -1 to -5 can be used to make a payment relative to the end of the month. To make a payment on the last day of the month, use -1; to make the payment on the second-to-last day, use -2, and so on.
	IntervalExecutionDay int32 `json:"interval_execution_day"`
	// A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will begin on the first `interval_execution_day` on or after the `start_date`.  If the first `interval_execution_day` on or after the start date is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so.
	StartDate string `json:"start_date"`
	// A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will end on the last `interval_execution_day` on or before the `end_date`. If the only `interval_execution_day` between the start date and the end date (inclusive) is also the same day that `/payment_initiation/payment/create` was called, the bank *may* make a payment on that day, but it is not guaranteed to do so.
	EndDate NullableString `json:"end_date"`
	// The start date sent to the bank after adjusting for holidays or weekends.  Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, this field will be `null`.
	AdjustedStartDate NullableString `json:"adjusted_start_date"`
}

// NewExternalPaymentScheduleGet instantiates a new ExternalPaymentScheduleGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentScheduleGet(interval PaymentScheduleInterval, intervalExecutionDay int32, startDate string, endDate NullableString, adjustedStartDate NullableString) *ExternalPaymentScheduleGet {
	this := ExternalPaymentScheduleGet{}
	this.Interval = interval
	this.IntervalExecutionDay = intervalExecutionDay
	this.StartDate = startDate
	this.EndDate = endDate
	this.AdjustedStartDate = adjustedStartDate
	return &this
}

// NewExternalPaymentScheduleGetWithDefaults instantiates a new ExternalPaymentScheduleGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentScheduleGetWithDefaults() *ExternalPaymentScheduleGet {
	this := ExternalPaymentScheduleGet{}
	return &this
}

// GetInterval returns the Interval field value
func (o *ExternalPaymentScheduleGet) GetInterval() PaymentScheduleInterval {
	if o == nil {
		var ret PaymentScheduleInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentScheduleGet) GetIntervalOk() (*PaymentScheduleInterval, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ExternalPaymentScheduleGet) SetInterval(v PaymentScheduleInterval) {
	o.Interval = v
}

// GetIntervalExecutionDay returns the IntervalExecutionDay field value
func (o *ExternalPaymentScheduleGet) GetIntervalExecutionDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntervalExecutionDay
}

// GetIntervalExecutionDayOk returns a tuple with the IntervalExecutionDay field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentScheduleGet) GetIntervalExecutionDayOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IntervalExecutionDay, true
}

// SetIntervalExecutionDay sets field value
func (o *ExternalPaymentScheduleGet) SetIntervalExecutionDay(v int32) {
	o.IntervalExecutionDay = v
}

// GetStartDate returns the StartDate field value
func (o *ExternalPaymentScheduleGet) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentScheduleGet) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *ExternalPaymentScheduleGet) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalPaymentScheduleGet) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentScheduleGet) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// SetEndDate sets field value
func (o *ExternalPaymentScheduleGet) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// GetAdjustedStartDate returns the AdjustedStartDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalPaymentScheduleGet) GetAdjustedStartDate() string {
	if o == nil || o.AdjustedStartDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.AdjustedStartDate.Get()
}

// GetAdjustedStartDateOk returns a tuple with the AdjustedStartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentScheduleGet) GetAdjustedStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdjustedStartDate.Get(), o.AdjustedStartDate.IsSet()
}

// SetAdjustedStartDate sets field value
func (o *ExternalPaymentScheduleGet) SetAdjustedStartDate(v string) {
	o.AdjustedStartDate.Set(&v)
}

func (o ExternalPaymentScheduleGet) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if true {
		toSerialize["adjusted_start_date"] = o.AdjustedStartDate.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPaymentScheduleGet struct {
	value *ExternalPaymentScheduleGet
	isSet bool
}

func (v NullableExternalPaymentScheduleGet) Get() *ExternalPaymentScheduleGet {
	return v.value
}

func (v *NullableExternalPaymentScheduleGet) Set(val *ExternalPaymentScheduleGet) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentScheduleGet) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentScheduleGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentScheduleGet(val *ExternalPaymentScheduleGet) *NullableExternalPaymentScheduleGet {
	return &NullableExternalPaymentScheduleGet{value: val, isSet: true}
}

func (v NullableExternalPaymentScheduleGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentScheduleGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


