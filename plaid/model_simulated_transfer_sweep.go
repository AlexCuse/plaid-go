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
	"time"
)

// SimulatedTransferSweep A sweep returned from the `/sandbox/transfer/sweep/simulate` endpoint. Can be null if there are no transfers to include in a sweep.
type SimulatedTransferSweep struct {
	// Identifier of the sweep.
	Id string `json:"id"`
	// The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited.
	FundingAccountId string `json:"funding_account_id"`
	// The datetime when the sweep occurred, in RFC 3339 format.
	Created time.Time `json:"created"`
	// Signed decimal amount of the sweep as it appears on your sweep account ledger (e.g. \"-10.00\")  If amount is not present, the sweep was net-settled to zero and outstanding debits and credits between the sweep account and Plaid are balanced.
	Amount string `json:"amount"`
	// The currency of the sweep, e.g. \"USD\".
	IsoCurrencyCode string `json:"iso_currency_code"`
	// The date when the sweep settled, in the YYYY-MM-DD format.
	Settled NullableString `json:"settled"`
	Status NullableSweepStatus `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimulatedTransferSweep SimulatedTransferSweep

// NewSimulatedTransferSweep instantiates a new SimulatedTransferSweep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulatedTransferSweep(id string, fundingAccountId string, created time.Time, amount string, isoCurrencyCode string, settled NullableString) *SimulatedTransferSweep {
	this := SimulatedTransferSweep{}
	this.Id = id
	this.FundingAccountId = fundingAccountId
	this.Created = created
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	this.Settled = settled
	return &this
}

// NewSimulatedTransferSweepWithDefaults instantiates a new SimulatedTransferSweep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulatedTransferSweepWithDefaults() *SimulatedTransferSweep {
	this := SimulatedTransferSweep{}
	return &this
}

// GetId returns the Id field value
func (o *SimulatedTransferSweep) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimulatedTransferSweep) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimulatedTransferSweep) SetId(v string) {
	o.Id = v
}

// GetFundingAccountId returns the FundingAccountId field value
func (o *SimulatedTransferSweep) GetFundingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingAccountId
}

// GetFundingAccountIdOk returns a tuple with the FundingAccountId field value
// and a boolean to check if the value has been set.
func (o *SimulatedTransferSweep) GetFundingAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FundingAccountId, true
}

// SetFundingAccountId sets field value
func (o *SimulatedTransferSweep) SetFundingAccountId(v string) {
	o.FundingAccountId = v
}

// GetCreated returns the Created field value
func (o *SimulatedTransferSweep) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *SimulatedTransferSweep) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *SimulatedTransferSweep) SetCreated(v time.Time) {
	o.Created = v
}

// GetAmount returns the Amount field value
func (o *SimulatedTransferSweep) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *SimulatedTransferSweep) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *SimulatedTransferSweep) SetAmount(v string) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *SimulatedTransferSweep) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *SimulatedTransferSweep) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *SimulatedTransferSweep) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetSettled returns the Settled field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SimulatedTransferSweep) GetSettled() string {
	if o == nil || o.Settled.Get() == nil {
		var ret string
		return ret
	}

	return *o.Settled.Get()
}

// GetSettledOk returns a tuple with the Settled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimulatedTransferSweep) GetSettledOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Settled.Get(), o.Settled.IsSet()
}

// SetSettled sets field value
func (o *SimulatedTransferSweep) SetSettled(v string) {
	o.Settled.Set(&v)
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimulatedTransferSweep) GetStatus() SweepStatus {
	if o == nil || o.Status.Get() == nil {
		var ret SweepStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimulatedTransferSweep) GetStatusOk() (*SweepStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *SimulatedTransferSweep) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableSweepStatus and assigns it to the Status field.
func (o *SimulatedTransferSweep) SetStatus(v SweepStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *SimulatedTransferSweep) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *SimulatedTransferSweep) UnsetStatus() {
	o.Status.Unset()
}

func (o SimulatedTransferSweep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["funding_account_id"] = o.FundingAccountId
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["settled"] = o.Settled.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SimulatedTransferSweep) UnmarshalJSON(bytes []byte) (err error) {
	varSimulatedTransferSweep := _SimulatedTransferSweep{}

	if err = json.Unmarshal(bytes, &varSimulatedTransferSweep); err == nil {
		*o = SimulatedTransferSweep(varSimulatedTransferSweep)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "funding_account_id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "settled")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimulatedTransferSweep struct {
	value *SimulatedTransferSweep
	isSet bool
}

func (v NullableSimulatedTransferSweep) Get() *SimulatedTransferSweep {
	return v.value
}

func (v *NullableSimulatedTransferSweep) Set(val *SimulatedTransferSweep) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulatedTransferSweep) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulatedTransferSweep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulatedTransferSweep(val *SimulatedTransferSweep) *NullableSimulatedTransferSweep {
	return &NullableSimulatedTransferSweep{value: val, isSet: true}
}

func (v NullableSimulatedTransferSweep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulatedTransferSweep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


