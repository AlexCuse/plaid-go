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
	"time"
)

// CreditBankIncome The report of the Bank Income data for an end user.
type CreditBankIncome struct {
	// The unique identifier associated with the Bank Income Report.
	BankIncomeId *string `json:"bank_income_id,omitempty"`
	// The time when the Bank Income Report was generated.
	GeneratedTime *time.Time `json:"generated_time,omitempty"`
	// The number of days requested by the customer for the Bank Income Report.
	DaysRequested *int32 `json:"days_requested,omitempty"`
	// The list of Items in the report along with the associated metadata about the Item.
	Items *[]CreditBankIncomeItem `json:"items,omitempty"`
	BankIncomeSummary *CreditBankIncomeSummary `json:"bank_income_summary,omitempty"`
	// If data from the Bank Income report was unable to be retrieved, the warnings will contain information about the error that caused the data to be incomplete.
	Warnings *[]CreditBankIncomeWarning `json:"warnings,omitempty"`
}

// NewCreditBankIncome instantiates a new CreditBankIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncome() *CreditBankIncome {
	this := CreditBankIncome{}
	return &this
}

// NewCreditBankIncomeWithDefaults instantiates a new CreditBankIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeWithDefaults() *CreditBankIncome {
	this := CreditBankIncome{}
	return &this
}

// GetBankIncomeId returns the BankIncomeId field value if set, zero value otherwise.
func (o *CreditBankIncome) GetBankIncomeId() string {
	if o == nil || o.BankIncomeId == nil {
		var ret string
		return ret
	}
	return *o.BankIncomeId
}

// GetBankIncomeIdOk returns a tuple with the BankIncomeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncome) GetBankIncomeIdOk() (*string, bool) {
	if o == nil || o.BankIncomeId == nil {
		return nil, false
	}
	return o.BankIncomeId, true
}

// HasBankIncomeId returns a boolean if a field has been set.
func (o *CreditBankIncome) HasBankIncomeId() bool {
	if o != nil && o.BankIncomeId != nil {
		return true
	}

	return false
}

// SetBankIncomeId gets a reference to the given string and assigns it to the BankIncomeId field.
func (o *CreditBankIncome) SetBankIncomeId(v string) {
	o.BankIncomeId = &v
}

// GetGeneratedTime returns the GeneratedTime field value if set, zero value otherwise.
func (o *CreditBankIncome) GetGeneratedTime() time.Time {
	if o == nil || o.GeneratedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.GeneratedTime
}

// GetGeneratedTimeOk returns a tuple with the GeneratedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncome) GetGeneratedTimeOk() (*time.Time, bool) {
	if o == nil || o.GeneratedTime == nil {
		return nil, false
	}
	return o.GeneratedTime, true
}

// HasGeneratedTime returns a boolean if a field has been set.
func (o *CreditBankIncome) HasGeneratedTime() bool {
	if o != nil && o.GeneratedTime != nil {
		return true
	}

	return false
}

// SetGeneratedTime gets a reference to the given time.Time and assigns it to the GeneratedTime field.
func (o *CreditBankIncome) SetGeneratedTime(v time.Time) {
	o.GeneratedTime = &v
}

// GetDaysRequested returns the DaysRequested field value if set, zero value otherwise.
func (o *CreditBankIncome) GetDaysRequested() int32 {
	if o == nil || o.DaysRequested == nil {
		var ret int32
		return ret
	}
	return *o.DaysRequested
}

// GetDaysRequestedOk returns a tuple with the DaysRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncome) GetDaysRequestedOk() (*int32, bool) {
	if o == nil || o.DaysRequested == nil {
		return nil, false
	}
	return o.DaysRequested, true
}

// HasDaysRequested returns a boolean if a field has been set.
func (o *CreditBankIncome) HasDaysRequested() bool {
	if o != nil && o.DaysRequested != nil {
		return true
	}

	return false
}

// SetDaysRequested gets a reference to the given int32 and assigns it to the DaysRequested field.
func (o *CreditBankIncome) SetDaysRequested(v int32) {
	o.DaysRequested = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *CreditBankIncome) GetItems() []CreditBankIncomeItem {
	if o == nil || o.Items == nil {
		var ret []CreditBankIncomeItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncome) GetItemsOk() (*[]CreditBankIncomeItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *CreditBankIncome) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []CreditBankIncomeItem and assigns it to the Items field.
func (o *CreditBankIncome) SetItems(v []CreditBankIncomeItem) {
	o.Items = &v
}

// GetBankIncomeSummary returns the BankIncomeSummary field value if set, zero value otherwise.
func (o *CreditBankIncome) GetBankIncomeSummary() CreditBankIncomeSummary {
	if o == nil || o.BankIncomeSummary == nil {
		var ret CreditBankIncomeSummary
		return ret
	}
	return *o.BankIncomeSummary
}

// GetBankIncomeSummaryOk returns a tuple with the BankIncomeSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncome) GetBankIncomeSummaryOk() (*CreditBankIncomeSummary, bool) {
	if o == nil || o.BankIncomeSummary == nil {
		return nil, false
	}
	return o.BankIncomeSummary, true
}

// HasBankIncomeSummary returns a boolean if a field has been set.
func (o *CreditBankIncome) HasBankIncomeSummary() bool {
	if o != nil && o.BankIncomeSummary != nil {
		return true
	}

	return false
}

// SetBankIncomeSummary gets a reference to the given CreditBankIncomeSummary and assigns it to the BankIncomeSummary field.
func (o *CreditBankIncome) SetBankIncomeSummary(v CreditBankIncomeSummary) {
	o.BankIncomeSummary = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CreditBankIncome) GetWarnings() []CreditBankIncomeWarning {
	if o == nil || o.Warnings == nil {
		var ret []CreditBankIncomeWarning
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncome) GetWarningsOk() (*[]CreditBankIncomeWarning, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CreditBankIncome) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []CreditBankIncomeWarning and assigns it to the Warnings field.
func (o *CreditBankIncome) SetWarnings(v []CreditBankIncomeWarning) {
	o.Warnings = &v
}

func (o CreditBankIncome) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BankIncomeId != nil {
		toSerialize["bank_income_id"] = o.BankIncomeId
	}
	if o.GeneratedTime != nil {
		toSerialize["generated_time"] = o.GeneratedTime
	}
	if o.DaysRequested != nil {
		toSerialize["days_requested"] = o.DaysRequested
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.BankIncomeSummary != nil {
		toSerialize["bank_income_summary"] = o.BankIncomeSummary
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankIncome struct {
	value *CreditBankIncome
	isSet bool
}

func (v NullableCreditBankIncome) Get() *CreditBankIncome {
	return v.value
}

func (v *NullableCreditBankIncome) Set(val *CreditBankIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncome(val *CreditBankIncome) *NullableCreditBankIncome {
	return &NullableCreditBankIncome{value: val, isSet: true}
}

func (v NullableCreditBankIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


