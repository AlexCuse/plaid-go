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

// TransferBalanceGetResponse Defines the response schema for `/transfer/balance/get`
type TransferBalanceGetResponse struct {
	Balance TransferBalance `json:"balance"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferBalanceGetResponse TransferBalanceGetResponse

// NewTransferBalanceGetResponse instantiates a new TransferBalanceGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferBalanceGetResponse(balance TransferBalance, requestId string) *TransferBalanceGetResponse {
	this := TransferBalanceGetResponse{}
	this.Balance = balance
	this.RequestId = requestId
	return &this
}

// NewTransferBalanceGetResponseWithDefaults instantiates a new TransferBalanceGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferBalanceGetResponseWithDefaults() *TransferBalanceGetResponse {
	this := TransferBalanceGetResponse{}
	return &this
}

// GetBalance returns the Balance field value
func (o *TransferBalanceGetResponse) GetBalance() TransferBalance {
	if o == nil {
		var ret TransferBalance
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *TransferBalanceGetResponse) GetBalanceOk() (*TransferBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *TransferBalanceGetResponse) SetBalance(v TransferBalance) {
	o.Balance = v
}

// GetRequestId returns the RequestId field value
func (o *TransferBalanceGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferBalanceGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferBalanceGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferBalanceGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferBalanceGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferBalanceGetResponse := _TransferBalanceGetResponse{}

	if err = json.Unmarshal(bytes, &varTransferBalanceGetResponse); err == nil {
		*o = TransferBalanceGetResponse(varTransferBalanceGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "balance")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferBalanceGetResponse struct {
	value *TransferBalanceGetResponse
	isSet bool
}

func (v NullableTransferBalanceGetResponse) Get() *TransferBalanceGetResponse {
	return v.value
}

func (v *NullableTransferBalanceGetResponse) Set(val *TransferBalanceGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferBalanceGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferBalanceGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferBalanceGetResponse(val *TransferBalanceGetResponse) *NullableTransferBalanceGetResponse {
	return &NullableTransferBalanceGetResponse{value: val, isSet: true}
}

func (v NullableTransferBalanceGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferBalanceGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


