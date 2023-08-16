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

// TransferQuestionnaireCreateResponse Defines the response schema for `/transfer/questionnaire/create`
type TransferQuestionnaireCreateResponse struct {
	// Plaid-hosted onboarding URL that you will redirect the end customer to.
	OnboardingUrl string `json:"onboarding_url"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferQuestionnaireCreateResponse TransferQuestionnaireCreateResponse

// NewTransferQuestionnaireCreateResponse instantiates a new TransferQuestionnaireCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferQuestionnaireCreateResponse(onboardingUrl string, requestId string) *TransferQuestionnaireCreateResponse {
	this := TransferQuestionnaireCreateResponse{}
	this.OnboardingUrl = onboardingUrl
	this.RequestId = requestId
	return &this
}

// NewTransferQuestionnaireCreateResponseWithDefaults instantiates a new TransferQuestionnaireCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferQuestionnaireCreateResponseWithDefaults() *TransferQuestionnaireCreateResponse {
	this := TransferQuestionnaireCreateResponse{}
	return &this
}

// GetOnboardingUrl returns the OnboardingUrl field value
func (o *TransferQuestionnaireCreateResponse) GetOnboardingUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OnboardingUrl
}

// GetOnboardingUrlOk returns a tuple with the OnboardingUrl field value
// and a boolean to check if the value has been set.
func (o *TransferQuestionnaireCreateResponse) GetOnboardingUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OnboardingUrl, true
}

// SetOnboardingUrl sets field value
func (o *TransferQuestionnaireCreateResponse) SetOnboardingUrl(v string) {
	o.OnboardingUrl = v
}

// GetRequestId returns the RequestId field value
func (o *TransferQuestionnaireCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferQuestionnaireCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferQuestionnaireCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferQuestionnaireCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["onboarding_url"] = o.OnboardingUrl
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferQuestionnaireCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferQuestionnaireCreateResponse := _TransferQuestionnaireCreateResponse{}

	if err = json.Unmarshal(bytes, &varTransferQuestionnaireCreateResponse); err == nil {
		*o = TransferQuestionnaireCreateResponse(varTransferQuestionnaireCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "onboarding_url")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferQuestionnaireCreateResponse struct {
	value *TransferQuestionnaireCreateResponse
	isSet bool
}

func (v NullableTransferQuestionnaireCreateResponse) Get() *TransferQuestionnaireCreateResponse {
	return v.value
}

func (v *NullableTransferQuestionnaireCreateResponse) Set(val *TransferQuestionnaireCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferQuestionnaireCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferQuestionnaireCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferQuestionnaireCreateResponse(val *TransferQuestionnaireCreateResponse) *NullableTransferQuestionnaireCreateResponse {
	return &NullableTransferQuestionnaireCreateResponse{value: val, isSet: true}
}

func (v NullableTransferQuestionnaireCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferQuestionnaireCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


