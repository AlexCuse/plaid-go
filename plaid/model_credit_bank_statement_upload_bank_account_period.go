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

// CreditBankStatementUploadBankAccountPeriod An object containing data on the overall period of the statement.
type CreditBankStatementUploadBankAccountPeriod struct {
	// The start date of the statement period in ISO 8601 format (YYYY-MM-DD).
	StartDate NullableString `json:"start_date"`
	// The end date of the statement period in ISO 8601 format (YYYY-MM-DD).
	EndDate NullableString `json:"end_date"`
	// The starting balance of the bank account for the period.
	StartingBalance NullableFloat32 `json:"starting_balance"`
	// The ending balance of the bank account for the period.
	EndingBalance NullableFloat32 `json:"ending_balance"`
}

// NewCreditBankStatementUploadBankAccountPeriod instantiates a new CreditBankStatementUploadBankAccountPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankStatementUploadBankAccountPeriod(startDate NullableString, endDate NullableString, startingBalance NullableFloat32, endingBalance NullableFloat32) *CreditBankStatementUploadBankAccountPeriod {
	this := CreditBankStatementUploadBankAccountPeriod{}
	this.StartDate = startDate
	this.EndDate = endDate
	this.StartingBalance = startingBalance
	this.EndingBalance = endingBalance
	return &this
}

// NewCreditBankStatementUploadBankAccountPeriodWithDefaults instantiates a new CreditBankStatementUploadBankAccountPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankStatementUploadBankAccountPeriodWithDefaults() *CreditBankStatementUploadBankAccountPeriod {
	this := CreditBankStatementUploadBankAccountPeriod{}
	return &this
}

// GetStartDate returns the StartDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditBankStatementUploadBankAccountPeriod) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankStatementUploadBankAccountPeriod) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// SetStartDate sets field value
func (o *CreditBankStatementUploadBankAccountPeriod) SetStartDate(v string) {
	o.StartDate.Set(&v)
}

// GetEndDate returns the EndDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditBankStatementUploadBankAccountPeriod) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankStatementUploadBankAccountPeriod) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// SetEndDate sets field value
func (o *CreditBankStatementUploadBankAccountPeriod) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// GetStartingBalance returns the StartingBalance field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *CreditBankStatementUploadBankAccountPeriod) GetStartingBalance() float32 {
	if o == nil || o.StartingBalance.Get() == nil {
		var ret float32
		return ret
	}

	return *o.StartingBalance.Get()
}

// GetStartingBalanceOk returns a tuple with the StartingBalance field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankStatementUploadBankAccountPeriod) GetStartingBalanceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartingBalance.Get(), o.StartingBalance.IsSet()
}

// SetStartingBalance sets field value
func (o *CreditBankStatementUploadBankAccountPeriod) SetStartingBalance(v float32) {
	o.StartingBalance.Set(&v)
}

// GetEndingBalance returns the EndingBalance field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *CreditBankStatementUploadBankAccountPeriod) GetEndingBalance() float32 {
	if o == nil || o.EndingBalance.Get() == nil {
		var ret float32
		return ret
	}

	return *o.EndingBalance.Get()
}

// GetEndingBalanceOk returns a tuple with the EndingBalance field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankStatementUploadBankAccountPeriod) GetEndingBalanceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndingBalance.Get(), o.EndingBalance.IsSet()
}

// SetEndingBalance sets field value
func (o *CreditBankStatementUploadBankAccountPeriod) SetEndingBalance(v float32) {
	o.EndingBalance.Set(&v)
}

func (o CreditBankStatementUploadBankAccountPeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if true {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if true {
		toSerialize["starting_balance"] = o.StartingBalance.Get()
	}
	if true {
		toSerialize["ending_balance"] = o.EndingBalance.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankStatementUploadBankAccountPeriod struct {
	value *CreditBankStatementUploadBankAccountPeriod
	isSet bool
}

func (v NullableCreditBankStatementUploadBankAccountPeriod) Get() *CreditBankStatementUploadBankAccountPeriod {
	return v.value
}

func (v *NullableCreditBankStatementUploadBankAccountPeriod) Set(val *CreditBankStatementUploadBankAccountPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankStatementUploadBankAccountPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankStatementUploadBankAccountPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankStatementUploadBankAccountPeriod(val *CreditBankStatementUploadBankAccountPeriod) *NullableCreditBankStatementUploadBankAccountPeriod {
	return &NullableCreditBankStatementUploadBankAccountPeriod{value: val, isSet: true}
}

func (v NullableCreditBankStatementUploadBankAccountPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankStatementUploadBankAccountPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


