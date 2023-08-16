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

// CreditRelayCreateResponse CreditRelayCreateResponse defines the response schema for `/credit/relay/create`
type CreditRelayCreateResponse struct {
	// A token that can be shared with a third party to allow them to access the Asset Report. This token should be stored securely.
	RelayToken string `json:"relay_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditRelayCreateResponse CreditRelayCreateResponse

// NewCreditRelayCreateResponse instantiates a new CreditRelayCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditRelayCreateResponse(relayToken string, requestId string) *CreditRelayCreateResponse {
	this := CreditRelayCreateResponse{}
	this.RelayToken = relayToken
	this.RequestId = requestId
	return &this
}

// NewCreditRelayCreateResponseWithDefaults instantiates a new CreditRelayCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditRelayCreateResponseWithDefaults() *CreditRelayCreateResponse {
	this := CreditRelayCreateResponse{}
	return &this
}

// GetRelayToken returns the RelayToken field value
func (o *CreditRelayCreateResponse) GetRelayToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelayToken
}

// GetRelayTokenOk returns a tuple with the RelayToken field value
// and a boolean to check if the value has been set.
func (o *CreditRelayCreateResponse) GetRelayTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RelayToken, true
}

// SetRelayToken sets field value
func (o *CreditRelayCreateResponse) SetRelayToken(v string) {
	o.RelayToken = v
}

// GetRequestId returns the RequestId field value
func (o *CreditRelayCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditRelayCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditRelayCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditRelayCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["relay_token"] = o.RelayToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditRelayCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditRelayCreateResponse := _CreditRelayCreateResponse{}

	if err = json.Unmarshal(bytes, &varCreditRelayCreateResponse); err == nil {
		*o = CreditRelayCreateResponse(varCreditRelayCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "relay_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditRelayCreateResponse struct {
	value *CreditRelayCreateResponse
	isSet bool
}

func (v NullableCreditRelayCreateResponse) Get() *CreditRelayCreateResponse {
	return v.value
}

func (v *NullableCreditRelayCreateResponse) Set(val *CreditRelayCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditRelayCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditRelayCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditRelayCreateResponse(val *CreditRelayCreateResponse) *NullableCreditRelayCreateResponse {
	return &NullableCreditRelayCreateResponse{value: val, isSet: true}
}

func (v NullableCreditRelayCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditRelayCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


