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

// CreditBankIncomeHistoricalSummary The end user's monthly summary for the income source(s).
type CreditBankIncomeHistoricalSummary struct {
	// Total amount of earnings for the income source(s) of the user for the month in the summary.
	TotalAmount *float32 `json:"total_amount,omitempty"`
	// The ISO 4217 currency code of the amount or balance.
	IsoCurrencyCode NullableString `json:"iso_currency_code,omitempty"`
	// The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code,omitempty"`
	// The start date of the period covered in this monthly summary. This date will be the first day of the month, unless the month being covered is a partial month because it is the first month included in the summary and the date range being requested does not begin with the first day of the month. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	StartDate *string `json:"start_date,omitempty"`
	// The end date of the period included in this monthly summary. This date will be the last day of the month, unless the month being covered is a partial month because it is the last month included in the summary and the date range being requested does not end with the last day of the month. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	EndDate *string `json:"end_date,omitempty"`
	Transactions *[]CreditBankIncomeTransaction `json:"transactions,omitempty"`
}

// NewCreditBankIncomeHistoricalSummary instantiates a new CreditBankIncomeHistoricalSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeHistoricalSummary() *CreditBankIncomeHistoricalSummary {
	this := CreditBankIncomeHistoricalSummary{}
	return &this
}

// NewCreditBankIncomeHistoricalSummaryWithDefaults instantiates a new CreditBankIncomeHistoricalSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeHistoricalSummaryWithDefaults() *CreditBankIncomeHistoricalSummary {
	this := CreditBankIncomeHistoricalSummary{}
	return &this
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *CreditBankIncomeHistoricalSummary) GetTotalAmount() float32 {
	if o == nil || o.TotalAmount == nil {
		var ret float32
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeHistoricalSummary) GetTotalAmountOk() (*float32, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *CreditBankIncomeHistoricalSummary) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given float32 and assigns it to the TotalAmount field.
func (o *CreditBankIncomeHistoricalSummary) SetTotalAmount(v float32) {
	o.TotalAmount = &v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditBankIncomeHistoricalSummary) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankIncomeHistoricalSummary) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *CreditBankIncomeHistoricalSummary) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given NullableString and assigns it to the IsoCurrencyCode field.
func (o *CreditBankIncomeHistoricalSummary) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}
// SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil
func (o *CreditBankIncomeHistoricalSummary) SetIsoCurrencyCodeNil() {
	o.IsoCurrencyCode.Set(nil)
}

// UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
func (o *CreditBankIncomeHistoricalSummary) UnsetIsoCurrencyCode() {
	o.IsoCurrencyCode.Unset()
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditBankIncomeHistoricalSummary) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankIncomeHistoricalSummary) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// HasUnofficialCurrencyCode returns a boolean if a field has been set.
func (o *CreditBankIncomeHistoricalSummary) HasUnofficialCurrencyCode() bool {
	if o != nil && o.UnofficialCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetUnofficialCurrencyCode gets a reference to the given NullableString and assigns it to the UnofficialCurrencyCode field.
func (o *CreditBankIncomeHistoricalSummary) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}
// SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil
func (o *CreditBankIncomeHistoricalSummary) SetUnofficialCurrencyCodeNil() {
	o.UnofficialCurrencyCode.Set(nil)
}

// UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
func (o *CreditBankIncomeHistoricalSummary) UnsetUnofficialCurrencyCode() {
	o.UnofficialCurrencyCode.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreditBankIncomeHistoricalSummary) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeHistoricalSummary) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreditBankIncomeHistoricalSummary) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CreditBankIncomeHistoricalSummary) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CreditBankIncomeHistoricalSummary) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeHistoricalSummary) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CreditBankIncomeHistoricalSummary) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *CreditBankIncomeHistoricalSummary) SetEndDate(v string) {
	o.EndDate = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *CreditBankIncomeHistoricalSummary) GetTransactions() []CreditBankIncomeTransaction {
	if o == nil || o.Transactions == nil {
		var ret []CreditBankIncomeTransaction
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeHistoricalSummary) GetTransactionsOk() (*[]CreditBankIncomeTransaction, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *CreditBankIncomeHistoricalSummary) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []CreditBankIncomeTransaction and assigns it to the Transactions field.
func (o *CreditBankIncomeHistoricalSummary) SetTransactions(v []CreditBankIncomeTransaction) {
	o.Transactions = &v
}

func (o CreditBankIncomeHistoricalSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalAmount != nil {
		toSerialize["total_amount"] = o.TotalAmount
	}
	if o.IsoCurrencyCode.IsSet() {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if o.UnofficialCurrencyCode.IsSet() {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankIncomeHistoricalSummary struct {
	value *CreditBankIncomeHistoricalSummary
	isSet bool
}

func (v NullableCreditBankIncomeHistoricalSummary) Get() *CreditBankIncomeHistoricalSummary {
	return v.value
}

func (v *NullableCreditBankIncomeHistoricalSummary) Set(val *CreditBankIncomeHistoricalSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeHistoricalSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeHistoricalSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeHistoricalSummary(val *CreditBankIncomeHistoricalSummary) *NullableCreditBankIncomeHistoricalSummary {
	return &NullableCreditBankIncomeHistoricalSummary{value: val, isSet: true}
}

func (v NullableCreditBankIncomeHistoricalSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeHistoricalSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


