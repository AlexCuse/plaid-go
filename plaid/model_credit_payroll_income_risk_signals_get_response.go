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

// CreditPayrollIncomeRiskSignalsGetResponse CreditPayrollIncomeRiskSignalsGetRequest defines the response schema for `/credit/payroll_income/risk_signals/get`
type CreditPayrollIncomeRiskSignalsGetResponse struct {
	// Array of payroll items.
	Items []PayrollRiskSignalsItem `json:"items"`
	Error NullablePlaidError `json:"error,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditPayrollIncomeRiskSignalsGetResponse CreditPayrollIncomeRiskSignalsGetResponse

// NewCreditPayrollIncomeRiskSignalsGetResponse instantiates a new CreditPayrollIncomeRiskSignalsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayrollIncomeRiskSignalsGetResponse(items []PayrollRiskSignalsItem, requestId string) *CreditPayrollIncomeRiskSignalsGetResponse {
	this := CreditPayrollIncomeRiskSignalsGetResponse{}
	this.Items = items
	this.RequestId = requestId
	return &this
}

// NewCreditPayrollIncomeRiskSignalsGetResponseWithDefaults instantiates a new CreditPayrollIncomeRiskSignalsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayrollIncomeRiskSignalsGetResponseWithDefaults() *CreditPayrollIncomeRiskSignalsGetResponse {
	this := CreditPayrollIncomeRiskSignalsGetResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *CreditPayrollIncomeRiskSignalsGetResponse) GetItems() []PayrollRiskSignalsItem {
	if o == nil {
		var ret []PayrollRiskSignalsItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRiskSignalsGetResponse) GetItemsOk() (*[]PayrollRiskSignalsItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CreditPayrollIncomeRiskSignalsGetResponse) SetItems(v []PayrollRiskSignalsItem) {
	o.Items = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditPayrollIncomeRiskSignalsGetResponse) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayrollIncomeRiskSignalsGetResponse) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *CreditPayrollIncomeRiskSignalsGetResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullablePlaidError and assigns it to the Error field.
func (o *CreditPayrollIncomeRiskSignalsGetResponse) SetError(v PlaidError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *CreditPayrollIncomeRiskSignalsGetResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *CreditPayrollIncomeRiskSignalsGetResponse) UnsetError() {
	o.Error.Unset()
}

// GetRequestId returns the RequestId field value
func (o *CreditPayrollIncomeRiskSignalsGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRiskSignalsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditPayrollIncomeRiskSignalsGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditPayrollIncomeRiskSignalsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditPayrollIncomeRiskSignalsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPayrollIncomeRiskSignalsGetResponse := _CreditPayrollIncomeRiskSignalsGetResponse{}

	if err = json.Unmarshal(bytes, &varCreditPayrollIncomeRiskSignalsGetResponse); err == nil {
		*o = CreditPayrollIncomeRiskSignalsGetResponse(varCreditPayrollIncomeRiskSignalsGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "error")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPayrollIncomeRiskSignalsGetResponse struct {
	value *CreditPayrollIncomeRiskSignalsGetResponse
	isSet bool
}

func (v NullableCreditPayrollIncomeRiskSignalsGetResponse) Get() *CreditPayrollIncomeRiskSignalsGetResponse {
	return v.value
}

func (v *NullableCreditPayrollIncomeRiskSignalsGetResponse) Set(val *CreditPayrollIncomeRiskSignalsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayrollIncomeRiskSignalsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayrollIncomeRiskSignalsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayrollIncomeRiskSignalsGetResponse(val *CreditPayrollIncomeRiskSignalsGetResponse) *NullableCreditPayrollIncomeRiskSignalsGetResponse {
	return &NullableCreditPayrollIncomeRiskSignalsGetResponse{value: val, isSet: true}
}

func (v NullableCreditPayrollIncomeRiskSignalsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayrollIncomeRiskSignalsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


