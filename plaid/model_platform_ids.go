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

// PlatformIds An object containing a set of ids related to an employee
type PlatformIds struct {
	// The ID of an employee as given by their employer
	EmployeeId NullableString `json:"employee_id,omitempty"`
	// The ID of an employee as given by their payroll
	PayrollId NullableString `json:"payroll_id,omitempty"`
	// The ID of the position of the employee
	PositionId NullableString `json:"position_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlatformIds PlatformIds

// NewPlatformIds instantiates a new PlatformIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformIds() *PlatformIds {
	this := PlatformIds{}
	return &this
}

// NewPlatformIdsWithDefaults instantiates a new PlatformIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformIdsWithDefaults() *PlatformIds {
	this := PlatformIds{}
	return &this
}

// GetEmployeeId returns the EmployeeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformIds) GetEmployeeId() string {
	if o == nil || o.EmployeeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployeeId.Get()
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformIds) GetEmployeeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeeId.Get(), o.EmployeeId.IsSet()
}

// HasEmployeeId returns a boolean if a field has been set.
func (o *PlatformIds) HasEmployeeId() bool {
	if o != nil && o.EmployeeId.IsSet() {
		return true
	}

	return false
}

// SetEmployeeId gets a reference to the given NullableString and assigns it to the EmployeeId field.
func (o *PlatformIds) SetEmployeeId(v string) {
	o.EmployeeId.Set(&v)
}
// SetEmployeeIdNil sets the value for EmployeeId to be an explicit nil
func (o *PlatformIds) SetEmployeeIdNil() {
	o.EmployeeId.Set(nil)
}

// UnsetEmployeeId ensures that no value is present for EmployeeId, not even an explicit nil
func (o *PlatformIds) UnsetEmployeeId() {
	o.EmployeeId.Unset()
}

// GetPayrollId returns the PayrollId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformIds) GetPayrollId() string {
	if o == nil || o.PayrollId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PayrollId.Get()
}

// GetPayrollIdOk returns a tuple with the PayrollId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformIds) GetPayrollIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayrollId.Get(), o.PayrollId.IsSet()
}

// HasPayrollId returns a boolean if a field has been set.
func (o *PlatformIds) HasPayrollId() bool {
	if o != nil && o.PayrollId.IsSet() {
		return true
	}

	return false
}

// SetPayrollId gets a reference to the given NullableString and assigns it to the PayrollId field.
func (o *PlatformIds) SetPayrollId(v string) {
	o.PayrollId.Set(&v)
}
// SetPayrollIdNil sets the value for PayrollId to be an explicit nil
func (o *PlatformIds) SetPayrollIdNil() {
	o.PayrollId.Set(nil)
}

// UnsetPayrollId ensures that no value is present for PayrollId, not even an explicit nil
func (o *PlatformIds) UnsetPayrollId() {
	o.PayrollId.Unset()
}

// GetPositionId returns the PositionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformIds) GetPositionId() string {
	if o == nil || o.PositionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PositionId.Get()
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformIds) GetPositionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PositionId.Get(), o.PositionId.IsSet()
}

// HasPositionId returns a boolean if a field has been set.
func (o *PlatformIds) HasPositionId() bool {
	if o != nil && o.PositionId.IsSet() {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given NullableString and assigns it to the PositionId field.
func (o *PlatformIds) SetPositionId(v string) {
	o.PositionId.Set(&v)
}
// SetPositionIdNil sets the value for PositionId to be an explicit nil
func (o *PlatformIds) SetPositionIdNil() {
	o.PositionId.Set(nil)
}

// UnsetPositionId ensures that no value is present for PositionId, not even an explicit nil
func (o *PlatformIds) UnsetPositionId() {
	o.PositionId.Unset()
}

func (o PlatformIds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmployeeId.IsSet() {
		toSerialize["employee_id"] = o.EmployeeId.Get()
	}
	if o.PayrollId.IsSet() {
		toSerialize["payroll_id"] = o.PayrollId.Get()
	}
	if o.PositionId.IsSet() {
		toSerialize["position_id"] = o.PositionId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PlatformIds) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformIds := _PlatformIds{}

	if err = json.Unmarshal(bytes, &varPlatformIds); err == nil {
		*o = PlatformIds(varPlatformIds)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "employee_id")
		delete(additionalProperties, "payroll_id")
		delete(additionalProperties, "position_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlatformIds struct {
	value *PlatformIds
	isSet bool
}

func (v NullablePlatformIds) Get() *PlatformIds {
	return v.value
}

func (v *NullablePlatformIds) Set(val *PlatformIds) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformIds) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformIds(val *PlatformIds) *NullablePlatformIds {
	return &NullablePlatformIds{value: val, isSet: true}
}

func (v NullablePlatformIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


