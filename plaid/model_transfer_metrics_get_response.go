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

// TransferMetricsGetResponse Defines the response schema for `/transfer/metrics/get`
type TransferMetricsGetResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	// Sum of dollar amount of debit transfers in last 24 hours (decimal string with two digits of precision e.g. \"10.00\").
	DailyDebitTransferVolume string `json:"daily_debit_transfer_volume"`
	// Sum of dollar amount of credit transfers in last 24 hours (decimal string with two digits of precision e.g. \"10.00\").
	DailyCreditTransferVolume string `json:"daily_credit_transfer_volume"`
	// Sum of dollar amount of credit and debit transfers in current calendar month (decimal string with two digits of precision e.g. \"10.00\").
	MonthlyTransferVolume string `json:"monthly_transfer_volume"`
	// Sum of dollar amount of debit transfers in current calendar month (decimal string with two digits of precision e.g. \"10.00\").
	MonthlyDebitTransferVolume string `json:"monthly_debit_transfer_volume"`
	// Sum of dollar amount of credit transfers in current calendar month (decimal string with two digits of precision e.g. \"10.00\").
	MonthlyCreditTransferVolume string `json:"monthly_credit_transfer_volume"`
	// The currency of the dollar amount, e.g. \"USD\".
	IsoCurrencyCode string `json:"iso_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _TransferMetricsGetResponse TransferMetricsGetResponse

// NewTransferMetricsGetResponse instantiates a new TransferMetricsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferMetricsGetResponse(requestId string, dailyDebitTransferVolume string, dailyCreditTransferVolume string, monthlyTransferVolume string, monthlyDebitTransferVolume string, monthlyCreditTransferVolume string, isoCurrencyCode string) *TransferMetricsGetResponse {
	this := TransferMetricsGetResponse{}
	this.RequestId = requestId
	this.DailyDebitTransferVolume = dailyDebitTransferVolume
	this.DailyCreditTransferVolume = dailyCreditTransferVolume
	this.MonthlyTransferVolume = monthlyTransferVolume
	this.MonthlyDebitTransferVolume = monthlyDebitTransferVolume
	this.MonthlyCreditTransferVolume = monthlyCreditTransferVolume
	this.IsoCurrencyCode = isoCurrencyCode
	return &this
}

// NewTransferMetricsGetResponseWithDefaults instantiates a new TransferMetricsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferMetricsGetResponseWithDefaults() *TransferMetricsGetResponse {
	this := TransferMetricsGetResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *TransferMetricsGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferMetricsGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetDailyDebitTransferVolume returns the DailyDebitTransferVolume field value
func (o *TransferMetricsGetResponse) GetDailyDebitTransferVolume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DailyDebitTransferVolume
}

// GetDailyDebitTransferVolumeOk returns a tuple with the DailyDebitTransferVolume field value
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetResponse) GetDailyDebitTransferVolumeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DailyDebitTransferVolume, true
}

// SetDailyDebitTransferVolume sets field value
func (o *TransferMetricsGetResponse) SetDailyDebitTransferVolume(v string) {
	o.DailyDebitTransferVolume = v
}

// GetDailyCreditTransferVolume returns the DailyCreditTransferVolume field value
func (o *TransferMetricsGetResponse) GetDailyCreditTransferVolume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DailyCreditTransferVolume
}

// GetDailyCreditTransferVolumeOk returns a tuple with the DailyCreditTransferVolume field value
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetResponse) GetDailyCreditTransferVolumeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DailyCreditTransferVolume, true
}

// SetDailyCreditTransferVolume sets field value
func (o *TransferMetricsGetResponse) SetDailyCreditTransferVolume(v string) {
	o.DailyCreditTransferVolume = v
}

// GetMonthlyTransferVolume returns the MonthlyTransferVolume field value
func (o *TransferMetricsGetResponse) GetMonthlyTransferVolume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonthlyTransferVolume
}

// GetMonthlyTransferVolumeOk returns a tuple with the MonthlyTransferVolume field value
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetResponse) GetMonthlyTransferVolumeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MonthlyTransferVolume, true
}

// SetMonthlyTransferVolume sets field value
func (o *TransferMetricsGetResponse) SetMonthlyTransferVolume(v string) {
	o.MonthlyTransferVolume = v
}

// GetMonthlyDebitTransferVolume returns the MonthlyDebitTransferVolume field value
func (o *TransferMetricsGetResponse) GetMonthlyDebitTransferVolume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonthlyDebitTransferVolume
}

// GetMonthlyDebitTransferVolumeOk returns a tuple with the MonthlyDebitTransferVolume field value
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetResponse) GetMonthlyDebitTransferVolumeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MonthlyDebitTransferVolume, true
}

// SetMonthlyDebitTransferVolume sets field value
func (o *TransferMetricsGetResponse) SetMonthlyDebitTransferVolume(v string) {
	o.MonthlyDebitTransferVolume = v
}

// GetMonthlyCreditTransferVolume returns the MonthlyCreditTransferVolume field value
func (o *TransferMetricsGetResponse) GetMonthlyCreditTransferVolume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonthlyCreditTransferVolume
}

// GetMonthlyCreditTransferVolumeOk returns a tuple with the MonthlyCreditTransferVolume field value
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetResponse) GetMonthlyCreditTransferVolumeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MonthlyCreditTransferVolume, true
}

// SetMonthlyCreditTransferVolume sets field value
func (o *TransferMetricsGetResponse) SetMonthlyCreditTransferVolume(v string) {
	o.MonthlyCreditTransferVolume = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *TransferMetricsGetResponse) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetResponse) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *TransferMetricsGetResponse) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

func (o TransferMetricsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["daily_debit_transfer_volume"] = o.DailyDebitTransferVolume
	}
	if true {
		toSerialize["daily_credit_transfer_volume"] = o.DailyCreditTransferVolume
	}
	if true {
		toSerialize["monthly_transfer_volume"] = o.MonthlyTransferVolume
	}
	if true {
		toSerialize["monthly_debit_transfer_volume"] = o.MonthlyDebitTransferVolume
	}
	if true {
		toSerialize["monthly_credit_transfer_volume"] = o.MonthlyCreditTransferVolume
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferMetricsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferMetricsGetResponse := _TransferMetricsGetResponse{}

	if err = json.Unmarshal(bytes, &varTransferMetricsGetResponse); err == nil {
		*o = TransferMetricsGetResponse(varTransferMetricsGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "daily_debit_transfer_volume")
		delete(additionalProperties, "daily_credit_transfer_volume")
		delete(additionalProperties, "monthly_transfer_volume")
		delete(additionalProperties, "monthly_debit_transfer_volume")
		delete(additionalProperties, "monthly_credit_transfer_volume")
		delete(additionalProperties, "iso_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferMetricsGetResponse struct {
	value *TransferMetricsGetResponse
	isSet bool
}

func (v NullableTransferMetricsGetResponse) Get() *TransferMetricsGetResponse {
	return v.value
}

func (v *NullableTransferMetricsGetResponse) Set(val *TransferMetricsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferMetricsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferMetricsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferMetricsGetResponse(val *TransferMetricsGetResponse) *NullableTransferMetricsGetResponse {
	return &NullableTransferMetricsGetResponse{value: val, isSet: true}
}

func (v NullableTransferMetricsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferMetricsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

