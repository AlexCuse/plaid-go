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

// KYCCheckDetails Additional information for the `kyc_check` step. This field will be `null` unless `steps.kyc_check` has reached a terminal state of either `success` or `failed`.
type KYCCheckDetails struct {
	// The outcome status for the associated Identity Verification attempt's `kyc_check` step. This field will always have the same value as `steps.kyc_check`.
	Status string `json:"status"`
	Address KYCCheckAddressSummary `json:"address"`
	Name KYCCheckNameSummary `json:"name"`
	DateOfBirth KYCCheckDateOfBirthSummary `json:"date_of_birth"`
	IdNumber KYCCheckIDNumberSummary `json:"id_number"`
	PhoneNumber KYCCheckPhoneSummary `json:"phone_number"`
	AdditionalProperties map[string]interface{}
}

type _KYCCheckDetails KYCCheckDetails

// NewKYCCheckDetails instantiates a new KYCCheckDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCCheckDetails(status string, address KYCCheckAddressSummary, name KYCCheckNameSummary, dateOfBirth KYCCheckDateOfBirthSummary, idNumber KYCCheckIDNumberSummary, phoneNumber KYCCheckPhoneSummary) *KYCCheckDetails {
	this := KYCCheckDetails{}
	this.Status = status
	this.Address = address
	this.Name = name
	this.DateOfBirth = dateOfBirth
	this.IdNumber = idNumber
	this.PhoneNumber = phoneNumber
	return &this
}

// NewKYCCheckDetailsWithDefaults instantiates a new KYCCheckDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCCheckDetailsWithDefaults() *KYCCheckDetails {
	this := KYCCheckDetails{}
	return &this
}

// GetStatus returns the Status field value
func (o *KYCCheckDetails) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *KYCCheckDetails) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *KYCCheckDetails) SetStatus(v string) {
	o.Status = v
}

// GetAddress returns the Address field value
func (o *KYCCheckDetails) GetAddress() KYCCheckAddressSummary {
	if o == nil {
		var ret KYCCheckAddressSummary
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *KYCCheckDetails) GetAddressOk() (*KYCCheckAddressSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *KYCCheckDetails) SetAddress(v KYCCheckAddressSummary) {
	o.Address = v
}

// GetName returns the Name field value
func (o *KYCCheckDetails) GetName() KYCCheckNameSummary {
	if o == nil {
		var ret KYCCheckNameSummary
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KYCCheckDetails) GetNameOk() (*KYCCheckNameSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KYCCheckDetails) SetName(v KYCCheckNameSummary) {
	o.Name = v
}

// GetDateOfBirth returns the DateOfBirth field value
func (o *KYCCheckDetails) GetDateOfBirth() KYCCheckDateOfBirthSummary {
	if o == nil {
		var ret KYCCheckDateOfBirthSummary
		return ret
	}

	return o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
func (o *KYCCheckDetails) GetDateOfBirthOk() (*KYCCheckDateOfBirthSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DateOfBirth, true
}

// SetDateOfBirth sets field value
func (o *KYCCheckDetails) SetDateOfBirth(v KYCCheckDateOfBirthSummary) {
	o.DateOfBirth = v
}

// GetIdNumber returns the IdNumber field value
func (o *KYCCheckDetails) GetIdNumber() KYCCheckIDNumberSummary {
	if o == nil {
		var ret KYCCheckIDNumberSummary
		return ret
	}

	return o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value
// and a boolean to check if the value has been set.
func (o *KYCCheckDetails) GetIdNumberOk() (*KYCCheckIDNumberSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdNumber, true
}

// SetIdNumber sets field value
func (o *KYCCheckDetails) SetIdNumber(v KYCCheckIDNumberSummary) {
	o.IdNumber = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *KYCCheckDetails) GetPhoneNumber() KYCCheckPhoneSummary {
	if o == nil {
		var ret KYCCheckPhoneSummary
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *KYCCheckDetails) GetPhoneNumberOk() (*KYCCheckPhoneSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *KYCCheckDetails) SetPhoneNumber(v KYCCheckPhoneSummary) {
	o.PhoneNumber = v
}

func (o KYCCheckDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}
	if true {
		toSerialize["id_number"] = o.IdNumber
	}
	if true {
		toSerialize["phone_number"] = o.PhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KYCCheckDetails) UnmarshalJSON(bytes []byte) (err error) {
	varKYCCheckDetails := _KYCCheckDetails{}

	if err = json.Unmarshal(bytes, &varKYCCheckDetails); err == nil {
		*o = KYCCheckDetails(varKYCCheckDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "date_of_birth")
		delete(additionalProperties, "id_number")
		delete(additionalProperties, "phone_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKYCCheckDetails struct {
	value *KYCCheckDetails
	isSet bool
}

func (v NullableKYCCheckDetails) Get() *KYCCheckDetails {
	return v.value
}

func (v *NullableKYCCheckDetails) Set(val *KYCCheckDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCCheckDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCCheckDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCCheckDetails(val *KYCCheckDetails) *NullableKYCCheckDetails {
	return &NullableKYCCheckDetails{value: val, isSet: true}
}

func (v NullableKYCCheckDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCCheckDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


