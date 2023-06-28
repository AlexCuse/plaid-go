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

// CreditPayrollIncomeRefreshResponse CreditPayrollIncomeRefreshResponse defines the response schema for `/credit/payroll_income/refresh`
type CreditPayrollIncomeRefreshResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	// The verification refresh status. One of the following:  `\"USER_PRESENCE_REQUIRED\"` User presence is required to refresh an income verification. `\"SUCCESSFUL\"` The income verification refresh was successful. `\"NOT_FOUND\"` No new data was found after the income verification refresh.
	VerificationRefreshStatus string `json:"verification_refresh_status"`
	AdditionalProperties map[string]interface{}
}

type _CreditPayrollIncomeRefreshResponse CreditPayrollIncomeRefreshResponse

// NewCreditPayrollIncomeRefreshResponse instantiates a new CreditPayrollIncomeRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayrollIncomeRefreshResponse(requestId string, verificationRefreshStatus string) *CreditPayrollIncomeRefreshResponse {
	this := CreditPayrollIncomeRefreshResponse{}
	this.RequestId = requestId
	this.VerificationRefreshStatus = verificationRefreshStatus
	return &this
}

// NewCreditPayrollIncomeRefreshResponseWithDefaults instantiates a new CreditPayrollIncomeRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayrollIncomeRefreshResponseWithDefaults() *CreditPayrollIncomeRefreshResponse {
	this := CreditPayrollIncomeRefreshResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *CreditPayrollIncomeRefreshResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRefreshResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditPayrollIncomeRefreshResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetVerificationRefreshStatus returns the VerificationRefreshStatus field value
func (o *CreditPayrollIncomeRefreshResponse) GetVerificationRefreshStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationRefreshStatus
}

// GetVerificationRefreshStatusOk returns a tuple with the VerificationRefreshStatus field value
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRefreshResponse) GetVerificationRefreshStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationRefreshStatus, true
}

// SetVerificationRefreshStatus sets field value
func (o *CreditPayrollIncomeRefreshResponse) SetVerificationRefreshStatus(v string) {
	o.VerificationRefreshStatus = v
}

func (o CreditPayrollIncomeRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["verification_refresh_status"] = o.VerificationRefreshStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditPayrollIncomeRefreshResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPayrollIncomeRefreshResponse := _CreditPayrollIncomeRefreshResponse{}

	if err = json.Unmarshal(bytes, &varCreditPayrollIncomeRefreshResponse); err == nil {
		*o = CreditPayrollIncomeRefreshResponse(varCreditPayrollIncomeRefreshResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "verification_refresh_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPayrollIncomeRefreshResponse struct {
	value *CreditPayrollIncomeRefreshResponse
	isSet bool
}

func (v NullableCreditPayrollIncomeRefreshResponse) Get() *CreditPayrollIncomeRefreshResponse {
	return v.value
}

func (v *NullableCreditPayrollIncomeRefreshResponse) Set(val *CreditPayrollIncomeRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayrollIncomeRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayrollIncomeRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayrollIncomeRefreshResponse(val *CreditPayrollIncomeRefreshResponse) *NullableCreditPayrollIncomeRefreshResponse {
	return &NullableCreditPayrollIncomeRefreshResponse{value: val, isSet: true}
}

func (v NullableCreditPayrollIncomeRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayrollIncomeRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


