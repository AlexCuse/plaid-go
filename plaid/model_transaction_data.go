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

// TransactionData Information about the matched direct deposit transaction used to verify a user's payroll information.
type TransactionData struct {
	// The description of the transaction.
	Description string `json:"description"`
	// The amount of the transaction.
	Amount float64 `json:"amount"`
	// The date of the transaction, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (\"yyyy-mm-dd\").
	Date string `json:"date"`
	// A unique identifier for the end user's account.
	AccountId string `json:"account_id"`
	// A unique identifier for the transaction.
	TransactionId string `json:"transaction_id"`
	AdditionalProperties map[string]interface{}
}

type _TransactionData TransactionData

// NewTransactionData instantiates a new TransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionData(description string, amount float64, date string, accountId string, transactionId string) *TransactionData {
	this := TransactionData{}
	this.Description = description
	this.Amount = amount
	this.Date = date
	this.AccountId = accountId
	this.TransactionId = transactionId
	return &this
}

// NewTransactionDataWithDefaults instantiates a new TransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDataWithDefaults() *TransactionData {
	this := TransactionData{}
	return &this
}

// GetDescription returns the Description field value
func (o *TransactionData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransactionData) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransactionData) SetDescription(v string) {
	o.Description = v
}

// GetAmount returns the Amount field value
func (o *TransactionData) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionData) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionData) SetAmount(v float64) {
	o.Amount = v
}

// GetDate returns the Date field value
func (o *TransactionData) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *TransactionData) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *TransactionData) SetDate(v string) {
	o.Date = v
}

// GetAccountId returns the AccountId field value
func (o *TransactionData) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransactionData) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransactionData) SetAccountId(v string) {
	o.AccountId = v
}

// GetTransactionId returns the TransactionId field value
func (o *TransactionData) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionData) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *TransactionData) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o TransactionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionData) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionData := _TransactionData{}

	if err = json.Unmarshal(bytes, &varTransactionData); err == nil {
		*o = TransactionData(varTransactionData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "date")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "transaction_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionData struct {
	value *TransactionData
	isSet bool
}

func (v NullableTransactionData) Get() *TransactionData {
	return v.value
}

func (v *NullableTransactionData) Set(val *TransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionData(val *TransactionData) *NullableTransactionData {
	return &NullableTransactionData{value: val, isSet: true}
}

func (v NullableTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


