/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.61.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SandboxIncomeFireWebhookResponse SandboxIncomeFireWebhookResponse defines the response schema for `/sandbox/income/fire_webhook`
type SandboxIncomeFireWebhookResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
}

// NewSandboxIncomeFireWebhookResponse instantiates a new SandboxIncomeFireWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxIncomeFireWebhookResponse(requestId string) *SandboxIncomeFireWebhookResponse {
	this := SandboxIncomeFireWebhookResponse{}
	this.RequestId = requestId
	return &this
}

// NewSandboxIncomeFireWebhookResponseWithDefaults instantiates a new SandboxIncomeFireWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxIncomeFireWebhookResponseWithDefaults() *SandboxIncomeFireWebhookResponse {
	this := SandboxIncomeFireWebhookResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SandboxIncomeFireWebhookResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxIncomeFireWebhookResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxIncomeFireWebhookResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxIncomeFireWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxIncomeFireWebhookResponse struct {
	value *SandboxIncomeFireWebhookResponse
	isSet bool
}

func (v NullableSandboxIncomeFireWebhookResponse) Get() *SandboxIncomeFireWebhookResponse {
	return v.value
}

func (v *NullableSandboxIncomeFireWebhookResponse) Set(val *SandboxIncomeFireWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxIncomeFireWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxIncomeFireWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxIncomeFireWebhookResponse(val *SandboxIncomeFireWebhookResponse) *NullableSandboxIncomeFireWebhookResponse {
	return &NullableSandboxIncomeFireWebhookResponse{value: val, isSet: true}
}

func (v NullableSandboxIncomeFireWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxIncomeFireWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


