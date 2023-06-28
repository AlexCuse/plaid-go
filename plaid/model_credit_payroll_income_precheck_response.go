/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.385.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditPayrollIncomePrecheckResponse Defines the response schema for `/credit/payroll_income/precheck`.
type CreditPayrollIncomePrecheckResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	Confidence IncomeVerificationPrecheckConfidence `json:"confidence"`
	AdditionalProperties map[string]interface{}
}

type _CreditPayrollIncomePrecheckResponse CreditPayrollIncomePrecheckResponse

// NewCreditPayrollIncomePrecheckResponse instantiates a new CreditPayrollIncomePrecheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayrollIncomePrecheckResponse(requestId string, confidence IncomeVerificationPrecheckConfidence) *CreditPayrollIncomePrecheckResponse {
	this := CreditPayrollIncomePrecheckResponse{}
	this.RequestId = requestId
	this.Confidence = confidence
	return &this
}

// NewCreditPayrollIncomePrecheckResponseWithDefaults instantiates a new CreditPayrollIncomePrecheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayrollIncomePrecheckResponseWithDefaults() *CreditPayrollIncomePrecheckResponse {
	this := CreditPayrollIncomePrecheckResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *CreditPayrollIncomePrecheckResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomePrecheckResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditPayrollIncomePrecheckResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetConfidence returns the Confidence field value
func (o *CreditPayrollIncomePrecheckResponse) GetConfidence() IncomeVerificationPrecheckConfidence {
	if o == nil {
		var ret IncomeVerificationPrecheckConfidence
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomePrecheckResponse) GetConfidenceOk() (*IncomeVerificationPrecheckConfidence, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *CreditPayrollIncomePrecheckResponse) SetConfidence(v IncomeVerificationPrecheckConfidence) {
	o.Confidence = v
}

func (o CreditPayrollIncomePrecheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["confidence"] = o.Confidence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditPayrollIncomePrecheckResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPayrollIncomePrecheckResponse := _CreditPayrollIncomePrecheckResponse{}

	if err = json.Unmarshal(bytes, &varCreditPayrollIncomePrecheckResponse); err == nil {
		*o = CreditPayrollIncomePrecheckResponse(varCreditPayrollIncomePrecheckResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "confidence")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPayrollIncomePrecheckResponse struct {
	value *CreditPayrollIncomePrecheckResponse
	isSet bool
}

func (v NullableCreditPayrollIncomePrecheckResponse) Get() *CreditPayrollIncomePrecheckResponse {
	return v.value
}

func (v *NullableCreditPayrollIncomePrecheckResponse) Set(val *CreditPayrollIncomePrecheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayrollIncomePrecheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayrollIncomePrecheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayrollIncomePrecheckResponse(val *CreditPayrollIncomePrecheckResponse) *NullableCreditPayrollIncomePrecheckResponse {
	return &NullableCreditPayrollIncomePrecheckResponse{value: val, isSet: true}
}

func (v NullableCreditPayrollIncomePrecheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayrollIncomePrecheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


