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

// InvestmentTransaction A transaction within an investment account.
type InvestmentTransaction struct {
	// The ID of the Investment transaction, unique across all Plaid transactions. Like all Plaid identifiers, the `investment_transaction_id` is case sensitive.
	InvestmentTransactionId string `json:"investment_transaction_id"`
	// A legacy field formerly used internally by Plaid to identify certain canceled transactions.
	CancelTransactionId NullableString `json:"cancel_transaction_id,omitempty"`
	// The `account_id` of the account against which this transaction posted.
	AccountId string `json:"account_id"`
	// The `security_id` to which this transaction is related.
	SecurityId NullableString `json:"security_id"`
	// The [ISO 8601](https://wikipedia.org/wiki/ISO_8601) posting date for the transaction.
	Date string `json:"date"`
	// The institution’s description of the transaction.
	Name string `json:"name"`
	// The number of units of the security involved in this transaction.
	Quantity float64 `json:"quantity"`
	// The complete value of the transaction. Positive values when cash is debited, e.g. purchases of stock; negative values when cash is credited, e.g. sales of stock. Treatment remains the same for cash-only movements unassociated with securities.
	Amount float64 `json:"amount"`
	// The price of the security at which this transaction occurred.
	Price float64 `json:"price"`
	// The combined value of all fees applied to this transaction
	Fees NullableFloat64 `json:"fees"`
	Type InvestmentTransactionType `json:"type"`
	Subtype InvestmentTransactionSubtype `json:"subtype"`
	// The ISO-4217 currency code of the transaction. Always `null` if `unofficial_currency_code` is non-`null`.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the holding. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _InvestmentTransaction InvestmentTransaction

// NewInvestmentTransaction instantiates a new InvestmentTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentTransaction(investmentTransactionId string, accountId string, securityId NullableString, date string, name string, quantity float64, amount float64, price float64, fees NullableFloat64, type_ InvestmentTransactionType, subtype InvestmentTransactionSubtype, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString) *InvestmentTransaction {
	this := InvestmentTransaction{}
	this.InvestmentTransactionId = investmentTransactionId
	this.AccountId = accountId
	this.SecurityId = securityId
	this.Date = date
	this.Name = name
	this.Quantity = quantity
	this.Amount = amount
	this.Price = price
	this.Fees = fees
	this.Type = type_
	this.Subtype = subtype
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	return &this
}

// NewInvestmentTransactionWithDefaults instantiates a new InvestmentTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentTransactionWithDefaults() *InvestmentTransaction {
	this := InvestmentTransaction{}
	return &this
}

// GetInvestmentTransactionId returns the InvestmentTransactionId field value
func (o *InvestmentTransaction) GetInvestmentTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvestmentTransactionId
}

// GetInvestmentTransactionIdOk returns a tuple with the InvestmentTransactionId field value
// and a boolean to check if the value has been set.
func (o *InvestmentTransaction) GetInvestmentTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InvestmentTransactionId, true
}

// SetInvestmentTransactionId sets field value
func (o *InvestmentTransaction) SetInvestmentTransactionId(v string) {
	o.InvestmentTransactionId = v
}

// GetCancelTransactionId returns the CancelTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvestmentTransaction) GetCancelTransactionId() string {
	if o == nil || o.CancelTransactionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CancelTransactionId.Get()
}

// GetCancelTransactionIdOk returns a tuple with the CancelTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvestmentTransaction) GetCancelTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CancelTransactionId.Get(), o.CancelTransactionId.IsSet()
}

// HasCancelTransactionId returns a boolean if a field has been set.
func (o *InvestmentTransaction) HasCancelTransactionId() bool {
	if o != nil && o.CancelTransactionId.IsSet() {
		return true
	}

	return false
}

// SetCancelTransactionId gets a reference to the given NullableString and assigns it to the CancelTransactionId field.
func (o *InvestmentTransaction) SetCancelTransactionId(v string) {
	o.CancelTransactionId.Set(&v)
}
// SetCancelTransactionIdNil sets the value for CancelTransactionId to be an explicit nil
func (o *InvestmentTransaction) SetCancelTransactionIdNil() {
	o.CancelTransactionId.Set(nil)
}

// UnsetCancelTransactionId ensures that no value is present for CancelTransactionId, not even an explicit nil
func (o *InvestmentTransaction) UnsetCancelTransactionId() {
	o.CancelTransactionId.Unset()
}

// GetAccountId returns the AccountId field value
func (o *InvestmentTransaction) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *InvestmentTransaction) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *InvestmentTransaction) SetAccountId(v string) {
	o.AccountId = v
}

// GetSecurityId returns the SecurityId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvestmentTransaction) GetSecurityId() string {
	if o == nil || o.SecurityId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SecurityId.Get()
}

// GetSecurityIdOk returns a tuple with the SecurityId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvestmentTransaction) GetSecurityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecurityId.Get(), o.SecurityId.IsSet()
}

// SetSecurityId sets field value
func (o *InvestmentTransaction) SetSecurityId(v string) {
	o.SecurityId.Set(&v)
}

// GetDate returns the Date field value
func (o *InvestmentTransaction) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *InvestmentTransaction) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *InvestmentTransaction) SetDate(v string) {
	o.Date = v
}

// GetName returns the Name field value
func (o *InvestmentTransaction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvestmentTransaction) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvestmentTransaction) SetName(v string) {
	o.Name = v
}

// GetQuantity returns the Quantity field value
func (o *InvestmentTransaction) GetQuantity() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *InvestmentTransaction) GetQuantityOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *InvestmentTransaction) SetQuantity(v float64) {
	o.Quantity = v
}

// GetAmount returns the Amount field value
func (o *InvestmentTransaction) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InvestmentTransaction) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InvestmentTransaction) SetAmount(v float64) {
	o.Amount = v
}

// GetPrice returns the Price field value
func (o *InvestmentTransaction) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *InvestmentTransaction) GetPriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *InvestmentTransaction) SetPrice(v float64) {
	o.Price = v
}

// GetFees returns the Fees field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *InvestmentTransaction) GetFees() float64 {
	if o == nil || o.Fees.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Fees.Get()
}

// GetFeesOk returns a tuple with the Fees field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvestmentTransaction) GetFeesOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Fees.Get(), o.Fees.IsSet()
}

// SetFees sets field value
func (o *InvestmentTransaction) SetFees(v float64) {
	o.Fees.Set(&v)
}

// GetType returns the Type field value
func (o *InvestmentTransaction) GetType() InvestmentTransactionType {
	if o == nil {
		var ret InvestmentTransactionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InvestmentTransaction) GetTypeOk() (*InvestmentTransactionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InvestmentTransaction) SetType(v InvestmentTransactionType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value
func (o *InvestmentTransaction) GetSubtype() InvestmentTransactionSubtype {
	if o == nil {
		var ret InvestmentTransactionSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *InvestmentTransaction) GetSubtypeOk() (*InvestmentTransactionSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *InvestmentTransaction) SetSubtype(v InvestmentTransactionSubtype) {
	o.Subtype = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvestmentTransaction) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvestmentTransaction) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *InvestmentTransaction) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvestmentTransaction) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvestmentTransaction) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *InvestmentTransaction) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

func (o InvestmentTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["investment_transaction_id"] = o.InvestmentTransactionId
	}
	if o.CancelTransactionId.IsSet() {
		toSerialize["cancel_transaction_id"] = o.CancelTransactionId.Get()
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["security_id"] = o.SecurityId.Get()
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["fees"] = o.Fees.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InvestmentTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varInvestmentTransaction := _InvestmentTransaction{}

	if err = json.Unmarshal(bytes, &varInvestmentTransaction); err == nil {
		*o = InvestmentTransaction(varInvestmentTransaction)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "investment_transaction_id")
		delete(additionalProperties, "cancel_transaction_id")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "security_id")
		delete(additionalProperties, "date")
		delete(additionalProperties, "name")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "price")
		delete(additionalProperties, "fees")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvestmentTransaction struct {
	value *InvestmentTransaction
	isSet bool
}

func (v NullableInvestmentTransaction) Get() *InvestmentTransaction {
	return v.value
}

func (v *NullableInvestmentTransaction) Set(val *InvestmentTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentTransaction(val *InvestmentTransaction) *NullableInvestmentTransaction {
	return &NullableInvestmentTransaction{value: val, isSet: true}
}

func (v NullableInvestmentTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


