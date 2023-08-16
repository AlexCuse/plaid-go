/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.413.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditSessionBankEmploymentResult The details of a bank employment verification in Link.
type CreditSessionBankEmploymentResult struct {
	Status *CreditSessionBankEmploymentStatus `json:"status,omitempty"`
	// The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.
	ItemId *string `json:"item_id,omitempty"`
	// The Plaid Institution ID associated with the Item.
	InstitutionId *string `json:"institution_id,omitempty"`
}

// NewCreditSessionBankEmploymentResult instantiates a new CreditSessionBankEmploymentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditSessionBankEmploymentResult() *CreditSessionBankEmploymentResult {
	this := CreditSessionBankEmploymentResult{}
	return &this
}

// NewCreditSessionBankEmploymentResultWithDefaults instantiates a new CreditSessionBankEmploymentResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditSessionBankEmploymentResultWithDefaults() *CreditSessionBankEmploymentResult {
	this := CreditSessionBankEmploymentResult{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreditSessionBankEmploymentResult) GetStatus() CreditSessionBankEmploymentStatus {
	if o == nil || o.Status == nil {
		var ret CreditSessionBankEmploymentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSessionBankEmploymentResult) GetStatusOk() (*CreditSessionBankEmploymentStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreditSessionBankEmploymentResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CreditSessionBankEmploymentStatus and assigns it to the Status field.
func (o *CreditSessionBankEmploymentResult) SetStatus(v CreditSessionBankEmploymentStatus) {
	o.Status = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *CreditSessionBankEmploymentResult) GetItemId() string {
	if o == nil || o.ItemId == nil {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSessionBankEmploymentResult) GetItemIdOk() (*string, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *CreditSessionBankEmploymentResult) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *CreditSessionBankEmploymentResult) SetItemId(v string) {
	o.ItemId = &v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *CreditSessionBankEmploymentResult) GetInstitutionId() string {
	if o == nil || o.InstitutionId == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSessionBankEmploymentResult) GetInstitutionIdOk() (*string, bool) {
	if o == nil || o.InstitutionId == nil {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *CreditSessionBankEmploymentResult) HasInstitutionId() bool {
	if o != nil && o.InstitutionId != nil {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *CreditSessionBankEmploymentResult) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

func (o CreditSessionBankEmploymentResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ItemId != nil {
		toSerialize["item_id"] = o.ItemId
	}
	if o.InstitutionId != nil {
		toSerialize["institution_id"] = o.InstitutionId
	}
	return json.Marshal(toSerialize)
}

type NullableCreditSessionBankEmploymentResult struct {
	value *CreditSessionBankEmploymentResult
	isSet bool
}

func (v NullableCreditSessionBankEmploymentResult) Get() *CreditSessionBankEmploymentResult {
	return v.value
}

func (v *NullableCreditSessionBankEmploymentResult) Set(val *CreditSessionBankEmploymentResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditSessionBankEmploymentResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditSessionBankEmploymentResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditSessionBankEmploymentResult(val *CreditSessionBankEmploymentResult) *NullableCreditSessionBankEmploymentResult {
	return &NullableCreditSessionBankEmploymentResult{value: val, isSet: true}
}

func (v NullableCreditSessionBankEmploymentResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditSessionBankEmploymentResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

