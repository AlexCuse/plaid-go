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

// TransactionsRulesListResponse TransactionsRulesListResponse defines the response schema for `/beta/transactions/rules/v1/list`
type TransactionsRulesListResponse struct {
	// A list of the Item's transaction rules
	Rules []TransactionsCategoryRule `json:"rules"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransactionsRulesListResponse TransactionsRulesListResponse

// NewTransactionsRulesListResponse instantiates a new TransactionsRulesListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRulesListResponse(rules []TransactionsCategoryRule, requestId string) *TransactionsRulesListResponse {
	this := TransactionsRulesListResponse{}
	this.Rules = rules
	this.RequestId = requestId
	return &this
}

// NewTransactionsRulesListResponseWithDefaults instantiates a new TransactionsRulesListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRulesListResponseWithDefaults() *TransactionsRulesListResponse {
	this := TransactionsRulesListResponse{}
	return &this
}

// GetRules returns the Rules field value
func (o *TransactionsRulesListResponse) GetRules() []TransactionsCategoryRule {
	if o == nil {
		var ret []TransactionsCategoryRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesListResponse) GetRulesOk() (*[]TransactionsCategoryRule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value
func (o *TransactionsRulesListResponse) SetRules(v []TransactionsCategoryRule) {
	o.Rules = v
}

// GetRequestId returns the RequestId field value
func (o *TransactionsRulesListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransactionsRulesListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransactionsRulesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionsRulesListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionsRulesListResponse := _TransactionsRulesListResponse{}

	if err = json.Unmarshal(bytes, &varTransactionsRulesListResponse); err == nil {
		*o = TransactionsRulesListResponse(varTransactionsRulesListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "rules")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionsRulesListResponse struct {
	value *TransactionsRulesListResponse
	isSet bool
}

func (v NullableTransactionsRulesListResponse) Get() *TransactionsRulesListResponse {
	return v.value
}

func (v *NullableTransactionsRulesListResponse) Set(val *TransactionsRulesListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRulesListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRulesListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRulesListResponse(val *TransactionsRulesListResponse) *NullableTransactionsRulesListResponse {
	return &NullableTransactionsRulesListResponse{value: val, isSet: true}
}

func (v NullableTransactionsRulesListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRulesListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


