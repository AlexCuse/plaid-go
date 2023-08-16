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

// CreditBankIncomeTransaction The transactions data for the end user's income source(s).
type CreditBankIncomeTransaction struct {
	// The settled value of the transaction, denominated in the transactions's currency as stated in `iso_currency_code` or `unofficial_currency_code`. Positive values when money moves out of the account; negative values when money moves in. For example, credit card purchases are positive; credit card payment, direct deposits, and refunds are negative.
	Amount *float32 `json:"amount,omitempty"`
	// For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted. Both dates are returned in an ISO 8601 format (YYYY-MM-DD).
	Date *string `json:"date,omitempty"`
	// The merchant name or transaction description.
	Name *string `json:"name,omitempty"`
	// The string returned by the financial institution to describe the transaction.
	OriginalDescription NullableString `json:"original_description,omitempty"`
	// When true, identifies the transaction as pending or unsettled. Pending transaction details (name, type, amount, category ID) may change before they are settled.
	Pending *bool `json:"pending,omitempty"`
	// The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.
	TransactionId *string `json:"transaction_id,omitempty"`
	// The check number of the transaction. This field is only populated for check transactions.
	CheckNumber NullableString `json:"check_number,omitempty"`
	// The ISO 4217 currency code of the amount or balance.
	IsoCurrencyCode NullableString `json:"iso_currency_code,omitempty"`
	// The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreditBankIncomeTransaction CreditBankIncomeTransaction

// NewCreditBankIncomeTransaction instantiates a new CreditBankIncomeTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeTransaction() *CreditBankIncomeTransaction {
	this := CreditBankIncomeTransaction{}
	return &this
}

// NewCreditBankIncomeTransactionWithDefaults instantiates a new CreditBankIncomeTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeTransactionWithDefaults() *CreditBankIncomeTransaction {
	this := CreditBankIncomeTransaction{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CreditBankIncomeTransaction) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeTransaction) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CreditBankIncomeTransaction) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *CreditBankIncomeTransaction) SetAmount(v float32) {
	o.Amount = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *CreditBankIncomeTransaction) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeTransaction) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *CreditBankIncomeTransaction) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *CreditBankIncomeTransaction) SetDate(v string) {
	o.Date = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreditBankIncomeTransaction) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeTransaction) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreditBankIncomeTransaction) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreditBankIncomeTransaction) SetName(v string) {
	o.Name = &v
}

// GetOriginalDescription returns the OriginalDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditBankIncomeTransaction) GetOriginalDescription() string {
	if o == nil || o.OriginalDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginalDescription.Get()
}

// GetOriginalDescriptionOk returns a tuple with the OriginalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankIncomeTransaction) GetOriginalDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalDescription.Get(), o.OriginalDescription.IsSet()
}

// HasOriginalDescription returns a boolean if a field has been set.
func (o *CreditBankIncomeTransaction) HasOriginalDescription() bool {
	if o != nil && o.OriginalDescription.IsSet() {
		return true
	}

	return false
}

// SetOriginalDescription gets a reference to the given NullableString and assigns it to the OriginalDescription field.
func (o *CreditBankIncomeTransaction) SetOriginalDescription(v string) {
	o.OriginalDescription.Set(&v)
}
// SetOriginalDescriptionNil sets the value for OriginalDescription to be an explicit nil
func (o *CreditBankIncomeTransaction) SetOriginalDescriptionNil() {
	o.OriginalDescription.Set(nil)
}

// UnsetOriginalDescription ensures that no value is present for OriginalDescription, not even an explicit nil
func (o *CreditBankIncomeTransaction) UnsetOriginalDescription() {
	o.OriginalDescription.Unset()
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *CreditBankIncomeTransaction) GetPending() bool {
	if o == nil || o.Pending == nil {
		var ret bool
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeTransaction) GetPendingOk() (*bool, bool) {
	if o == nil || o.Pending == nil {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *CreditBankIncomeTransaction) HasPending() bool {
	if o != nil && o.Pending != nil {
		return true
	}

	return false
}

// SetPending gets a reference to the given bool and assigns it to the Pending field.
func (o *CreditBankIncomeTransaction) SetPending(v bool) {
	o.Pending = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *CreditBankIncomeTransaction) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeTransaction) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *CreditBankIncomeTransaction) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *CreditBankIncomeTransaction) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetCheckNumber returns the CheckNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditBankIncomeTransaction) GetCheckNumber() string {
	if o == nil || o.CheckNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.CheckNumber.Get()
}

// GetCheckNumberOk returns a tuple with the CheckNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankIncomeTransaction) GetCheckNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CheckNumber.Get(), o.CheckNumber.IsSet()
}

// HasCheckNumber returns a boolean if a field has been set.
func (o *CreditBankIncomeTransaction) HasCheckNumber() bool {
	if o != nil && o.CheckNumber.IsSet() {
		return true
	}

	return false
}

// SetCheckNumber gets a reference to the given NullableString and assigns it to the CheckNumber field.
func (o *CreditBankIncomeTransaction) SetCheckNumber(v string) {
	o.CheckNumber.Set(&v)
}
// SetCheckNumberNil sets the value for CheckNumber to be an explicit nil
func (o *CreditBankIncomeTransaction) SetCheckNumberNil() {
	o.CheckNumber.Set(nil)
}

// UnsetCheckNumber ensures that no value is present for CheckNumber, not even an explicit nil
func (o *CreditBankIncomeTransaction) UnsetCheckNumber() {
	o.CheckNumber.Unset()
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditBankIncomeTransaction) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankIncomeTransaction) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *CreditBankIncomeTransaction) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given NullableString and assigns it to the IsoCurrencyCode field.
func (o *CreditBankIncomeTransaction) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}
// SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil
func (o *CreditBankIncomeTransaction) SetIsoCurrencyCodeNil() {
	o.IsoCurrencyCode.Set(nil)
}

// UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
func (o *CreditBankIncomeTransaction) UnsetIsoCurrencyCode() {
	o.IsoCurrencyCode.Unset()
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditBankIncomeTransaction) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankIncomeTransaction) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// HasUnofficialCurrencyCode returns a boolean if a field has been set.
func (o *CreditBankIncomeTransaction) HasUnofficialCurrencyCode() bool {
	if o != nil && o.UnofficialCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetUnofficialCurrencyCode gets a reference to the given NullableString and assigns it to the UnofficialCurrencyCode field.
func (o *CreditBankIncomeTransaction) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}
// SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil
func (o *CreditBankIncomeTransaction) SetUnofficialCurrencyCodeNil() {
	o.UnofficialCurrencyCode.Set(nil)
}

// UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
func (o *CreditBankIncomeTransaction) UnsetUnofficialCurrencyCode() {
	o.UnofficialCurrencyCode.Unset()
}

func (o CreditBankIncomeTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OriginalDescription.IsSet() {
		toSerialize["original_description"] = o.OriginalDescription.Get()
	}
	if o.Pending != nil {
		toSerialize["pending"] = o.Pending
	}
	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if o.CheckNumber.IsSet() {
		toSerialize["check_number"] = o.CheckNumber.Get()
	}
	if o.IsoCurrencyCode.IsSet() {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if o.UnofficialCurrencyCode.IsSet() {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditBankIncomeTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varCreditBankIncomeTransaction := _CreditBankIncomeTransaction{}

	if err = json.Unmarshal(bytes, &varCreditBankIncomeTransaction); err == nil {
		*o = CreditBankIncomeTransaction(varCreditBankIncomeTransaction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "date")
		delete(additionalProperties, "name")
		delete(additionalProperties, "original_description")
		delete(additionalProperties, "pending")
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "check_number")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditBankIncomeTransaction struct {
	value *CreditBankIncomeTransaction
	isSet bool
}

func (v NullableCreditBankIncomeTransaction) Get() *CreditBankIncomeTransaction {
	return v.value
}

func (v *NullableCreditBankIncomeTransaction) Set(val *CreditBankIncomeTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeTransaction(val *CreditBankIncomeTransaction) *NullableCreditBankIncomeTransaction {
	return &NullableCreditBankIncomeTransaction{value: val, isSet: true}
}

func (v NullableCreditBankIncomeTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


