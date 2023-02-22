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

// WalletTransactionExecuteResponse WalletTransactionExecuteResponse defines the response schema for `/wallet/transaction/execute`
type WalletTransactionExecuteResponse struct {
	// A unique ID identifying the transaction
	TransactionId string `json:"transaction_id"`
	Status WalletTransactionStatus `json:"status"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransactionExecuteResponse WalletTransactionExecuteResponse

// NewWalletTransactionExecuteResponse instantiates a new WalletTransactionExecuteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionExecuteResponse(transactionId string, status WalletTransactionStatus, requestId string) *WalletTransactionExecuteResponse {
	this := WalletTransactionExecuteResponse{}
	this.TransactionId = transactionId
	this.Status = status
	this.RequestId = requestId
	return &this
}

// NewWalletTransactionExecuteResponseWithDefaults instantiates a new WalletTransactionExecuteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionExecuteResponseWithDefaults() *WalletTransactionExecuteResponse {
	this := WalletTransactionExecuteResponse{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *WalletTransactionExecuteResponse) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionExecuteResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *WalletTransactionExecuteResponse) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetStatus returns the Status field value
func (o *WalletTransactionExecuteResponse) GetStatus() WalletTransactionStatus {
	if o == nil {
		var ret WalletTransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionExecuteResponse) GetStatusOk() (*WalletTransactionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WalletTransactionExecuteResponse) SetStatus(v WalletTransactionStatus) {
	o.Status = v
}

// GetRequestId returns the RequestId field value
func (o *WalletTransactionExecuteResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionExecuteResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WalletTransactionExecuteResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WalletTransactionExecuteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletTransactionExecuteResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransactionExecuteResponse := _WalletTransactionExecuteResponse{}

	if err = json.Unmarshal(bytes, &varWalletTransactionExecuteResponse); err == nil {
		*o = WalletTransactionExecuteResponse(varWalletTransactionExecuteResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransactionExecuteResponse struct {
	value *WalletTransactionExecuteResponse
	isSet bool
}

func (v NullableWalletTransactionExecuteResponse) Get() *WalletTransactionExecuteResponse {
	return v.value
}

func (v *NullableWalletTransactionExecuteResponse) Set(val *WalletTransactionExecuteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionExecuteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionExecuteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionExecuteResponse(val *WalletTransactionExecuteResponse) *NullableWalletTransactionExecuteResponse {
	return &NullableWalletTransactionExecuteResponse{value: val, isSet: true}
}

func (v NullableWalletTransactionExecuteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionExecuteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


