/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.21.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// Paystub An object representing data extracted from the end user's paystub.
type Paystub struct {
	// The account identifier for the account associated with this paystub.
	AccountId NullableString `json:"account_id,omitempty"`
	Employer PaystubEmployer `json:"employer"`
	Employee Employee `json:"employee"`
	PayPeriodDetails PayPeriodDetails `json:"pay_period_details"`
	IncomeBreakdown IncomeBreakdown `json:"income_breakdown"`
	YtdEarnings PaystubYTDDetails `json:"ytd_earnings"`
	AdditionalProperties map[string]interface{}
}

type _Paystub Paystub

// NewPaystub instantiates a new Paystub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystub(employer PaystubEmployer, employee Employee, payPeriodDetails PayPeriodDetails, incomeBreakdown IncomeBreakdown, ytdEarnings PaystubYTDDetails) *Paystub {
	this := Paystub{}
	this.Employer = employer
	this.Employee = employee
	this.PayPeriodDetails = payPeriodDetails
	this.IncomeBreakdown = incomeBreakdown
	this.YtdEarnings = ytdEarnings
	return &this
}

// NewPaystubWithDefaults instantiates a new Paystub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubWithDefaults() *Paystub {
	this := Paystub{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Paystub) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Paystub) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *Paystub) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *Paystub) SetAccountId(v string) {
	o.AccountId.Set(&v)
}
// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *Paystub) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *Paystub) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetEmployer returns the Employer field value
func (o *Paystub) GetEmployer() PaystubEmployer {
	if o == nil {
		var ret PaystubEmployer
		return ret
	}

	return o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetEmployerOk() (*PaystubEmployer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employer, true
}

// SetEmployer sets field value
func (o *Paystub) SetEmployer(v PaystubEmployer) {
	o.Employer = v
}

// GetEmployee returns the Employee field value
func (o *Paystub) GetEmployee() Employee {
	if o == nil {
		var ret Employee
		return ret
	}

	return o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetEmployeeOk() (*Employee, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employee, true
}

// SetEmployee sets field value
func (o *Paystub) SetEmployee(v Employee) {
	o.Employee = v
}

// GetPayPeriodDetails returns the PayPeriodDetails field value
func (o *Paystub) GetPayPeriodDetails() PayPeriodDetails {
	if o == nil {
		var ret PayPeriodDetails
		return ret
	}

	return o.PayPeriodDetails
}

// GetPayPeriodDetailsOk returns a tuple with the PayPeriodDetails field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetPayPeriodDetailsOk() (*PayPeriodDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PayPeriodDetails, true
}

// SetPayPeriodDetails sets field value
func (o *Paystub) SetPayPeriodDetails(v PayPeriodDetails) {
	o.PayPeriodDetails = v
}

// GetIncomeBreakdown returns the IncomeBreakdown field value
func (o *Paystub) GetIncomeBreakdown() IncomeBreakdown {
	if o == nil {
		var ret IncomeBreakdown
		return ret
	}

	return o.IncomeBreakdown
}

// GetIncomeBreakdownOk returns a tuple with the IncomeBreakdown field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetIncomeBreakdownOk() (*IncomeBreakdown, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IncomeBreakdown, true
}

// SetIncomeBreakdown sets field value
func (o *Paystub) SetIncomeBreakdown(v IncomeBreakdown) {
	o.IncomeBreakdown = v
}

// GetYtdEarnings returns the YtdEarnings field value
func (o *Paystub) GetYtdEarnings() PaystubYTDDetails {
	if o == nil {
		var ret PaystubYTDDetails
		return ret
	}

	return o.YtdEarnings
}

// GetYtdEarningsOk returns a tuple with the YtdEarnings field value
// and a boolean to check if the value has been set.
func (o *Paystub) GetYtdEarningsOk() (*PaystubYTDDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YtdEarnings, true
}

// SetYtdEarnings sets field value
func (o *Paystub) SetYtdEarnings(v PaystubYTDDetails) {
	o.YtdEarnings = v
}

func (o Paystub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if true {
		toSerialize["employer"] = o.Employer
	}
	if true {
		toSerialize["employee"] = o.Employee
	}
	if true {
		toSerialize["pay_period_details"] = o.PayPeriodDetails
	}
	if true {
		toSerialize["income_breakdown"] = o.IncomeBreakdown
	}
	if true {
		toSerialize["ytd_earnings"] = o.YtdEarnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Paystub) UnmarshalJSON(bytes []byte) (err error) {
	varPaystub := _Paystub{}

	if err = json.Unmarshal(bytes, &varPaystub); err == nil {
		*o = Paystub(varPaystub)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "employer")
		delete(additionalProperties, "employee")
		delete(additionalProperties, "pay_period_details")
		delete(additionalProperties, "income_breakdown")
		delete(additionalProperties, "ytd_earnings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaystub struct {
	value *Paystub
	isSet bool
}

func (v NullablePaystub) Get() *Paystub {
	return v.value
}

func (v *NullablePaystub) Set(val *Paystub) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystub) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystub(val *Paystub) *NullablePaystub {
	return &NullablePaystub{value: val, isSet: true}
}

func (v NullablePaystub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

