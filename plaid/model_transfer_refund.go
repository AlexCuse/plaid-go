/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.410.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// TransferRefund Represents a refund within the Transfers API.
type TransferRefund struct {
	// Plaid’s unique identifier for a refund.
	Id string `json:"id"`
	// The ID of the transfer to refund.
	TransferId string `json:"transfer_id"`
	// The amount of the refund (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	Status TransferRefundStatus `json:"status"`
	FailureReason NullableTransferRefundFailure `json:"failure_reason"`
	// The datetime when this refund was created. This will be of the form `2006-01-02T15:04:05Z`
	Created time.Time `json:"created"`
	AdditionalProperties map[string]interface{}
}

type _TransferRefund TransferRefund

// NewTransferRefund instantiates a new TransferRefund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRefund(id string, transferId string, amount string, status TransferRefundStatus, failureReason NullableTransferRefundFailure, created time.Time) *TransferRefund {
	this := TransferRefund{}
	this.Id = id
	this.TransferId = transferId
	this.Amount = amount
	this.Status = status
	this.FailureReason = failureReason
	this.Created = created
	return &this
}

// NewTransferRefundWithDefaults instantiates a new TransferRefund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRefundWithDefaults() *TransferRefund {
	this := TransferRefund{}
	return &this
}

// GetId returns the Id field value
func (o *TransferRefund) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferRefund) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferRefund) SetId(v string) {
	o.Id = v
}

// GetTransferId returns the TransferId field value
func (o *TransferRefund) GetTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
func (o *TransferRefund) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferId, true
}

// SetTransferId sets field value
func (o *TransferRefund) SetTransferId(v string) {
	o.TransferId = v
}

// GetAmount returns the Amount field value
func (o *TransferRefund) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferRefund) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferRefund) SetAmount(v string) {
	o.Amount = v
}

// GetStatus returns the Status field value
func (o *TransferRefund) GetStatus() TransferRefundStatus {
	if o == nil {
		var ret TransferRefundStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransferRefund) GetStatusOk() (*TransferRefundStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransferRefund) SetStatus(v TransferRefundStatus) {
	o.Status = v
}

// GetFailureReason returns the FailureReason field value
// If the value is explicit nil, the zero value for TransferRefundFailure will be returned
func (o *TransferRefund) GetFailureReason() TransferRefundFailure {
	if o == nil || o.FailureReason.Get() == nil {
		var ret TransferRefundFailure
		return ret
	}

	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferRefund) GetFailureReasonOk() (*TransferRefundFailure, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// SetFailureReason sets field value
func (o *TransferRefund) SetFailureReason(v TransferRefundFailure) {
	o.FailureReason.Set(&v)
}

// GetCreated returns the Created field value
func (o *TransferRefund) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TransferRefund) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TransferRefund) SetCreated(v time.Time) {
	o.Created = v
}

func (o TransferRefund) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["transfer_id"] = o.TransferId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}
	if true {
		toSerialize["created"] = o.Created
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferRefund) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRefund := _TransferRefund{}

	if err = json.Unmarshal(bytes, &varTransferRefund); err == nil {
		*o = TransferRefund(varTransferRefund)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "transfer_id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "status")
		delete(additionalProperties, "failure_reason")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferRefund struct {
	value *TransferRefund
	isSet bool
}

func (v NullableTransferRefund) Get() *TransferRefund {
	return v.value
}

func (v *NullableTransferRefund) Set(val *TransferRefund) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRefund) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRefund(val *TransferRefund) *NullableTransferRefund {
	return &NullableTransferRefund{value: val, isSet: true}
}

func (v NullableTransferRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


