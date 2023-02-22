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

// DepositSwitchCreateResponse DepositSwitchCreateResponse defines the response schema for `/deposit_switch/create`
type DepositSwitchCreateResponse struct {
	// ID of the deposit switch. This ID is persisted throughout the lifetime of the deposit switch.
	DepositSwitchId string `json:"deposit_switch_id"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _DepositSwitchCreateResponse DepositSwitchCreateResponse

// NewDepositSwitchCreateResponse instantiates a new DepositSwitchCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchCreateResponse(depositSwitchId string, requestId string) *DepositSwitchCreateResponse {
	this := DepositSwitchCreateResponse{}
	this.DepositSwitchId = depositSwitchId
	this.RequestId = requestId
	return &this
}

// NewDepositSwitchCreateResponseWithDefaults instantiates a new DepositSwitchCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchCreateResponseWithDefaults() *DepositSwitchCreateResponse {
	this := DepositSwitchCreateResponse{}
	return &this
}

// GetDepositSwitchId returns the DepositSwitchId field value
func (o *DepositSwitchCreateResponse) GetDepositSwitchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositSwitchId
}

// GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchCreateResponse) GetDepositSwitchIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DepositSwitchId, true
}

// SetDepositSwitchId sets field value
func (o *DepositSwitchCreateResponse) SetDepositSwitchId(v string) {
	o.DepositSwitchId = v
}

// GetRequestId returns the RequestId field value
func (o *DepositSwitchCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *DepositSwitchCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o DepositSwitchCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deposit_switch_id"] = o.DepositSwitchId
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DepositSwitchCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDepositSwitchCreateResponse := _DepositSwitchCreateResponse{}

	if err = json.Unmarshal(bytes, &varDepositSwitchCreateResponse); err == nil {
		*o = DepositSwitchCreateResponse(varDepositSwitchCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "deposit_switch_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositSwitchCreateResponse struct {
	value *DepositSwitchCreateResponse
	isSet bool
}

func (v NullableDepositSwitchCreateResponse) Get() *DepositSwitchCreateResponse {
	return v.value
}

func (v *NullableDepositSwitchCreateResponse) Set(val *DepositSwitchCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchCreateResponse(val *DepositSwitchCreateResponse) *NullableDepositSwitchCreateResponse {
	return &NullableDepositSwitchCreateResponse{value: val, isSet: true}
}

func (v NullableDepositSwitchCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


