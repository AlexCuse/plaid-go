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
)

// BankTransferSweepListResponse BankTransferSweepListResponse defines the response schema for `/bank_transfer/sweep/list`
type BankTransferSweepListResponse struct {
	Sweeps []BankTransferSweep `json:"sweeps"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferSweepListResponse BankTransferSweepListResponse

// NewBankTransferSweepListResponse instantiates a new BankTransferSweepListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferSweepListResponse(sweeps []BankTransferSweep, requestId string) *BankTransferSweepListResponse {
	this := BankTransferSweepListResponse{}
	this.Sweeps = sweeps
	this.RequestId = requestId
	return &this
}

// NewBankTransferSweepListResponseWithDefaults instantiates a new BankTransferSweepListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferSweepListResponseWithDefaults() *BankTransferSweepListResponse {
	this := BankTransferSweepListResponse{}
	return &this
}

// GetSweeps returns the Sweeps field value
func (o *BankTransferSweepListResponse) GetSweeps() []BankTransferSweep {
	if o == nil {
		var ret []BankTransferSweep
		return ret
	}

	return o.Sweeps
}

// GetSweepsOk returns a tuple with the Sweeps field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweepListResponse) GetSweepsOk() (*[]BankTransferSweep, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sweeps, true
}

// SetSweeps sets field value
func (o *BankTransferSweepListResponse) SetSweeps(v []BankTransferSweep) {
	o.Sweeps = v
}

// GetRequestId returns the RequestId field value
func (o *BankTransferSweepListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweepListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BankTransferSweepListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BankTransferSweepListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sweeps"] = o.Sweeps
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferSweepListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferSweepListResponse := _BankTransferSweepListResponse{}

	if err = json.Unmarshal(bytes, &varBankTransferSweepListResponse); err == nil {
		*o = BankTransferSweepListResponse(varBankTransferSweepListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sweeps")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferSweepListResponse struct {
	value *BankTransferSweepListResponse
	isSet bool
}

func (v NullableBankTransferSweepListResponse) Get() *BankTransferSweepListResponse {
	return v.value
}

func (v *NullableBankTransferSweepListResponse) Set(val *BankTransferSweepListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferSweepListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferSweepListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferSweepListResponse(val *BankTransferSweepListResponse) *NullableBankTransferSweepListResponse {
	return &NullableBankTransferSweepListResponse{value: val, isSet: true}
}

func (v NullableBankTransferSweepListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferSweepListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


