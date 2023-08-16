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

// InvestmentsRefreshResponse InvestmentsRefreshResponse defines the response schema for `/investments/refresh`
type InvestmentsRefreshResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _InvestmentsRefreshResponse InvestmentsRefreshResponse

// NewInvestmentsRefreshResponse instantiates a new InvestmentsRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsRefreshResponse(requestId string) *InvestmentsRefreshResponse {
	this := InvestmentsRefreshResponse{}
	this.RequestId = requestId
	return &this
}

// NewInvestmentsRefreshResponseWithDefaults instantiates a new InvestmentsRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsRefreshResponseWithDefaults() *InvestmentsRefreshResponse {
	this := InvestmentsRefreshResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *InvestmentsRefreshResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *InvestmentsRefreshResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *InvestmentsRefreshResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o InvestmentsRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InvestmentsRefreshResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInvestmentsRefreshResponse := _InvestmentsRefreshResponse{}

	if err = json.Unmarshal(bytes, &varInvestmentsRefreshResponse); err == nil {
		*o = InvestmentsRefreshResponse(varInvestmentsRefreshResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvestmentsRefreshResponse struct {
	value *InvestmentsRefreshResponse
	isSet bool
}

func (v NullableInvestmentsRefreshResponse) Get() *InvestmentsRefreshResponse {
	return v.value
}

func (v *NullableInvestmentsRefreshResponse) Set(val *InvestmentsRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsRefreshResponse(val *InvestmentsRefreshResponse) *NullableInvestmentsRefreshResponse {
	return &NullableInvestmentsRefreshResponse{value: val, isSet: true}
}

func (v NullableInvestmentsRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


