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

// TransactionsRulesCreateResponse TransactionsRulesCreateResponse defines the response schema for `/beta/transactions/rules/v1/create`
type TransactionsRulesCreateResponse struct {
	Rule TransactionsCategoryRule `json:"rule"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransactionsRulesCreateResponse TransactionsRulesCreateResponse

// NewTransactionsRulesCreateResponse instantiates a new TransactionsRulesCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRulesCreateResponse(rule TransactionsCategoryRule, requestId string) *TransactionsRulesCreateResponse {
	this := TransactionsRulesCreateResponse{}
	this.Rule = rule
	this.RequestId = requestId
	return &this
}

// NewTransactionsRulesCreateResponseWithDefaults instantiates a new TransactionsRulesCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRulesCreateResponseWithDefaults() *TransactionsRulesCreateResponse {
	this := TransactionsRulesCreateResponse{}
	return &this
}

// GetRule returns the Rule field value
func (o *TransactionsRulesCreateResponse) GetRule() TransactionsCategoryRule {
	if o == nil {
		var ret TransactionsCategoryRule
		return ret
	}

	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesCreateResponse) GetRuleOk() (*TransactionsCategoryRule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value
func (o *TransactionsRulesCreateResponse) SetRule(v TransactionsCategoryRule) {
	o.Rule = v
}

// GetRequestId returns the RequestId field value
func (o *TransactionsRulesCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransactionsRulesCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransactionsRulesCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rule"] = o.Rule
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionsRulesCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionsRulesCreateResponse := _TransactionsRulesCreateResponse{}

	if err = json.Unmarshal(bytes, &varTransactionsRulesCreateResponse); err == nil {
		*o = TransactionsRulesCreateResponse(varTransactionsRulesCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "rule")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionsRulesCreateResponse struct {
	value *TransactionsRulesCreateResponse
	isSet bool
}

func (v NullableTransactionsRulesCreateResponse) Get() *TransactionsRulesCreateResponse {
	return v.value
}

func (v *NullableTransactionsRulesCreateResponse) Set(val *TransactionsRulesCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRulesCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRulesCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRulesCreateResponse(val *TransactionsRulesCreateResponse) *NullableTransactionsRulesCreateResponse {
	return &NullableTransactionsRulesCreateResponse{value: val, isSet: true}
}

func (v NullableTransactionsRulesCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRulesCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


