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

// IncomeVerificationWebhookStatus struct for IncomeVerificationWebhookStatus
type IncomeVerificationWebhookStatus struct {
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _IncomeVerificationWebhookStatus IncomeVerificationWebhookStatus

// NewIncomeVerificationWebhookStatus instantiates a new IncomeVerificationWebhookStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationWebhookStatus(id string) *IncomeVerificationWebhookStatus {
	this := IncomeVerificationWebhookStatus{}
	this.Id = id
	return &this
}

// NewIncomeVerificationWebhookStatusWithDefaults instantiates a new IncomeVerificationWebhookStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationWebhookStatusWithDefaults() *IncomeVerificationWebhookStatus {
	this := IncomeVerificationWebhookStatus{}
	return &this
}

// GetId returns the Id field value
func (o *IncomeVerificationWebhookStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationWebhookStatus) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IncomeVerificationWebhookStatus) SetId(v string) {
	o.Id = v
}

func (o IncomeVerificationWebhookStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeVerificationWebhookStatus) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeVerificationWebhookStatus := _IncomeVerificationWebhookStatus{}

	if err = json.Unmarshal(bytes, &varIncomeVerificationWebhookStatus); err == nil {
		*o = IncomeVerificationWebhookStatus(varIncomeVerificationWebhookStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeVerificationWebhookStatus struct {
	value *IncomeVerificationWebhookStatus
	isSet bool
}

func (v NullableIncomeVerificationWebhookStatus) Get() *IncomeVerificationWebhookStatus {
	return v.value
}

func (v *NullableIncomeVerificationWebhookStatus) Set(val *IncomeVerificationWebhookStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationWebhookStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationWebhookStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationWebhookStatus(val *IncomeVerificationWebhookStatus) *NullableIncomeVerificationWebhookStatus {
	return &NullableIncomeVerificationWebhookStatus{value: val, isSet: true}
}

func (v NullableIncomeVerificationWebhookStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationWebhookStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


