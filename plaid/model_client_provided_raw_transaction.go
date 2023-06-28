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

// ClientProvidedRawTransaction A client-provided transaction for Plaid to enhance.
type ClientProvidedRawTransaction struct {
	// A unique ID for the transaction used to help you tie data back to your systems.
	Id string `json:"id"`
	// The raw description of the transaction.
	Description string `json:"description"`
	// The value of the transaction with direction. (NOTE: this will affect enrichment results, so directions are important):.   Negative (-) for credits (e.g., incoming transfers, refunds)   Positive (+) for debits (e.g., purchases, fees, outgoing transfers)
	Amount float64 `json:"amount"`
	// The ISO-4217 currency code of the transaction e.g. USD.
	IsoCurrencyCode string `json:"iso_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _ClientProvidedRawTransaction ClientProvidedRawTransaction

// NewClientProvidedRawTransaction instantiates a new ClientProvidedRawTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProvidedRawTransaction(id string, description string, amount float64, isoCurrencyCode string) *ClientProvidedRawTransaction {
	this := ClientProvidedRawTransaction{}
	this.Id = id
	this.Description = description
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	return &this
}

// NewClientProvidedRawTransactionWithDefaults instantiates a new ClientProvidedRawTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProvidedRawTransactionWithDefaults() *ClientProvidedRawTransaction {
	this := ClientProvidedRawTransaction{}
	return &this
}

// GetId returns the Id field value
func (o *ClientProvidedRawTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedRawTransaction) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClientProvidedRawTransaction) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *ClientProvidedRawTransaction) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedRawTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ClientProvidedRawTransaction) SetDescription(v string) {
	o.Description = v
}

// GetAmount returns the Amount field value
func (o *ClientProvidedRawTransaction) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedRawTransaction) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ClientProvidedRawTransaction) SetAmount(v float64) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *ClientProvidedRawTransaction) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedRawTransaction) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *ClientProvidedRawTransaction) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

func (o ClientProvidedRawTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ClientProvidedRawTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varClientProvidedRawTransaction := _ClientProvidedRawTransaction{}

	if err = json.Unmarshal(bytes, &varClientProvidedRawTransaction); err == nil {
		*o = ClientProvidedRawTransaction(varClientProvidedRawTransaction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "iso_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientProvidedRawTransaction struct {
	value *ClientProvidedRawTransaction
	isSet bool
}

func (v NullableClientProvidedRawTransaction) Get() *ClientProvidedRawTransaction {
	return v.value
}

func (v *NullableClientProvidedRawTransaction) Set(val *ClientProvidedRawTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProvidedRawTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProvidedRawTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProvidedRawTransaction(val *ClientProvidedRawTransaction) *NullableClientProvidedRawTransaction {
	return &NullableClientProvidedRawTransaction{value: val, isSet: true}
}

func (v NullableClientProvidedRawTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProvidedRawTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


