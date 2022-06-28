/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.131.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransactionsRulesRemoveResponse TransactionsRulesRemoveResponse defines the response schema for `/beta/transactions/rules/v1/remove`
type TransactionsRulesRemoveResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransactionsRulesRemoveResponse TransactionsRulesRemoveResponse

// NewTransactionsRulesRemoveResponse instantiates a new TransactionsRulesRemoveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRulesRemoveResponse(requestId string) *TransactionsRulesRemoveResponse {
	this := TransactionsRulesRemoveResponse{}
	this.RequestId = requestId
	return &this
}

// NewTransactionsRulesRemoveResponseWithDefaults instantiates a new TransactionsRulesRemoveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRulesRemoveResponseWithDefaults() *TransactionsRulesRemoveResponse {
	this := TransactionsRulesRemoveResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *TransactionsRulesRemoveResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesRemoveResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransactionsRulesRemoveResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransactionsRulesRemoveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionsRulesRemoveResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionsRulesRemoveResponse := _TransactionsRulesRemoveResponse{}

	if err = json.Unmarshal(bytes, &varTransactionsRulesRemoveResponse); err == nil {
		*o = TransactionsRulesRemoveResponse(varTransactionsRulesRemoveResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionsRulesRemoveResponse struct {
	value *TransactionsRulesRemoveResponse
	isSet bool
}

func (v NullableTransactionsRulesRemoveResponse) Get() *TransactionsRulesRemoveResponse {
	return v.value
}

func (v *NullableTransactionsRulesRemoveResponse) Set(val *TransactionsRulesRemoveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRulesRemoveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRulesRemoveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRulesRemoveResponse(val *TransactionsRulesRemoveResponse) *NullableTransactionsRulesRemoveResponse {
	return &NullableTransactionsRulesRemoveResponse{value: val, isSet: true}
}

func (v NullableTransactionsRulesRemoveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRulesRemoveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


