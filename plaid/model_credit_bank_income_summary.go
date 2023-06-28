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

// CreditBankIncomeSummary Summary for bank income across all income sources and items (max history of 730 days).
type CreditBankIncomeSummary struct {
	// Total amount of earnings across all the income sources in the end user's Items for the days requested by the client. This may return an incorrect value if the summary includes income sources in multiple currencies. Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amounts) instead.
	TotalAmount *float32 `json:"total_amount,omitempty"`
	// The ISO 4217 currency code of the amount or balance. Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amounts) instead.
	IsoCurrencyCode NullableString `json:"iso_currency_code,omitempty"`
	// The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries. Please use [`total_amounts`](https://plaid.com/docs/api/products/income/#credit-bank_income-get-response-bank-income-bank-income-summary-total-amounts) instead.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code,omitempty"`
	// Total amount of earnings across all the income sources in the end user's Items for the days requested by the client. This can contain multiple amounts, with each amount denominated in one unique currency.
	TotalAmounts *[]CreditAmountWithCurrency `json:"total_amounts,omitempty"`
	// The earliest date within the days requested in which all income sources identified by Plaid appear in a user's account. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	StartDate *string `json:"start_date,omitempty"`
	// The latest date in which all income sources identified by Plaid appear in the user's account. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	EndDate *string `json:"end_date,omitempty"`
	// Number of income sources per end user.
	IncomeSourcesCount *int32 `json:"income_sources_count,omitempty"`
	// Number of income categories per end user.
	IncomeCategoriesCount *int32 `json:"income_categories_count,omitempty"`
	// Number of income transactions per end user.
	IncomeTransactionsCount *int32 `json:"income_transactions_count,omitempty"`
	HistoricalSummary *[]CreditBankIncomeHistoricalSummary `json:"historical_summary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreditBankIncomeSummary CreditBankIncomeSummary

// NewCreditBankIncomeSummary instantiates a new CreditBankIncomeSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeSummary() *CreditBankIncomeSummary {
	this := CreditBankIncomeSummary{}
	return &this
}

// NewCreditBankIncomeSummaryWithDefaults instantiates a new CreditBankIncomeSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeSummaryWithDefaults() *CreditBankIncomeSummary {
	this := CreditBankIncomeSummary{}
	return &this
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *CreditBankIncomeSummary) GetTotalAmount() float32 {
	if o == nil || o.TotalAmount == nil {
		var ret float32
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeSummary) GetTotalAmountOk() (*float32, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *CreditBankIncomeSummary) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given float32 and assigns it to the TotalAmount field.
func (o *CreditBankIncomeSummary) SetTotalAmount(v float32) {
	o.TotalAmount = &v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditBankIncomeSummary) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankIncomeSummary) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *CreditBankIncomeSummary) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given NullableString and assigns it to the IsoCurrencyCode field.
func (o *CreditBankIncomeSummary) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}
// SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil
func (o *CreditBankIncomeSummary) SetIsoCurrencyCodeNil() {
	o.IsoCurrencyCode.Set(nil)
}

// UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
func (o *CreditBankIncomeSummary) UnsetIsoCurrencyCode() {
	o.IsoCurrencyCode.Unset()
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditBankIncomeSummary) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankIncomeSummary) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// HasUnofficialCurrencyCode returns a boolean if a field has been set.
func (o *CreditBankIncomeSummary) HasUnofficialCurrencyCode() bool {
	if o != nil && o.UnofficialCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetUnofficialCurrencyCode gets a reference to the given NullableString and assigns it to the UnofficialCurrencyCode field.
func (o *CreditBankIncomeSummary) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}
// SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil
func (o *CreditBankIncomeSummary) SetUnofficialCurrencyCodeNil() {
	o.UnofficialCurrencyCode.Set(nil)
}

// UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
func (o *CreditBankIncomeSummary) UnsetUnofficialCurrencyCode() {
	o.UnofficialCurrencyCode.Unset()
}

// GetTotalAmounts returns the TotalAmounts field value if set, zero value otherwise.
func (o *CreditBankIncomeSummary) GetTotalAmounts() []CreditAmountWithCurrency {
	if o == nil || o.TotalAmounts == nil {
		var ret []CreditAmountWithCurrency
		return ret
	}
	return *o.TotalAmounts
}

// GetTotalAmountsOk returns a tuple with the TotalAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeSummary) GetTotalAmountsOk() (*[]CreditAmountWithCurrency, bool) {
	if o == nil || o.TotalAmounts == nil {
		return nil, false
	}
	return o.TotalAmounts, true
}

// HasTotalAmounts returns a boolean if a field has been set.
func (o *CreditBankIncomeSummary) HasTotalAmounts() bool {
	if o != nil && o.TotalAmounts != nil {
		return true
	}

	return false
}

// SetTotalAmounts gets a reference to the given []CreditAmountWithCurrency and assigns it to the TotalAmounts field.
func (o *CreditBankIncomeSummary) SetTotalAmounts(v []CreditAmountWithCurrency) {
	o.TotalAmounts = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreditBankIncomeSummary) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeSummary) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreditBankIncomeSummary) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CreditBankIncomeSummary) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CreditBankIncomeSummary) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeSummary) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CreditBankIncomeSummary) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *CreditBankIncomeSummary) SetEndDate(v string) {
	o.EndDate = &v
}

// GetIncomeSourcesCount returns the IncomeSourcesCount field value if set, zero value otherwise.
func (o *CreditBankIncomeSummary) GetIncomeSourcesCount() int32 {
	if o == nil || o.IncomeSourcesCount == nil {
		var ret int32
		return ret
	}
	return *o.IncomeSourcesCount
}

// GetIncomeSourcesCountOk returns a tuple with the IncomeSourcesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeSummary) GetIncomeSourcesCountOk() (*int32, bool) {
	if o == nil || o.IncomeSourcesCount == nil {
		return nil, false
	}
	return o.IncomeSourcesCount, true
}

// HasIncomeSourcesCount returns a boolean if a field has been set.
func (o *CreditBankIncomeSummary) HasIncomeSourcesCount() bool {
	if o != nil && o.IncomeSourcesCount != nil {
		return true
	}

	return false
}

// SetIncomeSourcesCount gets a reference to the given int32 and assigns it to the IncomeSourcesCount field.
func (o *CreditBankIncomeSummary) SetIncomeSourcesCount(v int32) {
	o.IncomeSourcesCount = &v
}

// GetIncomeCategoriesCount returns the IncomeCategoriesCount field value if set, zero value otherwise.
func (o *CreditBankIncomeSummary) GetIncomeCategoriesCount() int32 {
	if o == nil || o.IncomeCategoriesCount == nil {
		var ret int32
		return ret
	}
	return *o.IncomeCategoriesCount
}

// GetIncomeCategoriesCountOk returns a tuple with the IncomeCategoriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeSummary) GetIncomeCategoriesCountOk() (*int32, bool) {
	if o == nil || o.IncomeCategoriesCount == nil {
		return nil, false
	}
	return o.IncomeCategoriesCount, true
}

// HasIncomeCategoriesCount returns a boolean if a field has been set.
func (o *CreditBankIncomeSummary) HasIncomeCategoriesCount() bool {
	if o != nil && o.IncomeCategoriesCount != nil {
		return true
	}

	return false
}

// SetIncomeCategoriesCount gets a reference to the given int32 and assigns it to the IncomeCategoriesCount field.
func (o *CreditBankIncomeSummary) SetIncomeCategoriesCount(v int32) {
	o.IncomeCategoriesCount = &v
}

// GetIncomeTransactionsCount returns the IncomeTransactionsCount field value if set, zero value otherwise.
func (o *CreditBankIncomeSummary) GetIncomeTransactionsCount() int32 {
	if o == nil || o.IncomeTransactionsCount == nil {
		var ret int32
		return ret
	}
	return *o.IncomeTransactionsCount
}

// GetIncomeTransactionsCountOk returns a tuple with the IncomeTransactionsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeSummary) GetIncomeTransactionsCountOk() (*int32, bool) {
	if o == nil || o.IncomeTransactionsCount == nil {
		return nil, false
	}
	return o.IncomeTransactionsCount, true
}

// HasIncomeTransactionsCount returns a boolean if a field has been set.
func (o *CreditBankIncomeSummary) HasIncomeTransactionsCount() bool {
	if o != nil && o.IncomeTransactionsCount != nil {
		return true
	}

	return false
}

// SetIncomeTransactionsCount gets a reference to the given int32 and assigns it to the IncomeTransactionsCount field.
func (o *CreditBankIncomeSummary) SetIncomeTransactionsCount(v int32) {
	o.IncomeTransactionsCount = &v
}

// GetHistoricalSummary returns the HistoricalSummary field value if set, zero value otherwise.
func (o *CreditBankIncomeSummary) GetHistoricalSummary() []CreditBankIncomeHistoricalSummary {
	if o == nil || o.HistoricalSummary == nil {
		var ret []CreditBankIncomeHistoricalSummary
		return ret
	}
	return *o.HistoricalSummary
}

// GetHistoricalSummaryOk returns a tuple with the HistoricalSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeSummary) GetHistoricalSummaryOk() (*[]CreditBankIncomeHistoricalSummary, bool) {
	if o == nil || o.HistoricalSummary == nil {
		return nil, false
	}
	return o.HistoricalSummary, true
}

// HasHistoricalSummary returns a boolean if a field has been set.
func (o *CreditBankIncomeSummary) HasHistoricalSummary() bool {
	if o != nil && o.HistoricalSummary != nil {
		return true
	}

	return false
}

// SetHistoricalSummary gets a reference to the given []CreditBankIncomeHistoricalSummary and assigns it to the HistoricalSummary field.
func (o *CreditBankIncomeSummary) SetHistoricalSummary(v []CreditBankIncomeHistoricalSummary) {
	o.HistoricalSummary = &v
}

func (o CreditBankIncomeSummary) MarshalJSON() ([]byte, error) {
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
	if o.TotalAmounts != nil {
		toSerialize["total_amounts"] = o.TotalAmounts
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.IncomeSourcesCount != nil {
		toSerialize["income_sources_count"] = o.IncomeSourcesCount
	}
	if o.IncomeCategoriesCount != nil {
		toSerialize["income_categories_count"] = o.IncomeCategoriesCount
	}
	if o.IncomeTransactionsCount != nil {
		toSerialize["income_transactions_count"] = o.IncomeTransactionsCount
	}
	if o.HistoricalSummary != nil {
		toSerialize["historical_summary"] = o.HistoricalSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditBankIncomeSummary) UnmarshalJSON(bytes []byte) (err error) {
	varCreditBankIncomeSummary := _CreditBankIncomeSummary{}

	if err = json.Unmarshal(bytes, &varCreditBankIncomeSummary); err == nil {
		*o = CreditBankIncomeSummary(varCreditBankIncomeSummary)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "total_amount")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "total_amounts")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "income_sources_count")
		delete(additionalProperties, "income_categories_count")
		delete(additionalProperties, "income_transactions_count")
		delete(additionalProperties, "historical_summary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditBankIncomeSummary struct {
	value *CreditBankIncomeSummary
	isSet bool
}

func (v NullableCreditBankIncomeSummary) Get() *CreditBankIncomeSummary {
	return v.value
}

func (v *NullableCreditBankIncomeSummary) Set(val *CreditBankIncomeSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeSummary(val *CreditBankIncomeSummary) *NullableCreditBankIncomeSummary {
	return &NullableCreditBankIncomeSummary{value: val, isSet: true}
}

func (v NullableCreditBankIncomeSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


