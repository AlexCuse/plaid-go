/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferUserInRequest The legal name and other information for the account holder.
type TransferUserInRequest struct {
	// The user's legal name.
	LegalName string `json:"legal_name"`
	// The user's phone number.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The user's email address.
	EmailAddress *string `json:"email_address,omitempty"`
	Address *TransferUserAddressInRequest `json:"address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferUserInRequest TransferUserInRequest

// NewTransferUserInRequest instantiates a new TransferUserInRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferUserInRequest(legalName string) *TransferUserInRequest {
	this := TransferUserInRequest{}
	this.LegalName = legalName
	return &this
}

// NewTransferUserInRequestWithDefaults instantiates a new TransferUserInRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferUserInRequestWithDefaults() *TransferUserInRequest {
	this := TransferUserInRequest{}
	return &this
}

// GetLegalName returns the LegalName field value
func (o *TransferUserInRequest) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *TransferUserInRequest) GetLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *TransferUserInRequest) SetLegalName(v string) {
	o.LegalName = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *TransferUserInRequest) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserInRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *TransferUserInRequest) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *TransferUserInRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *TransferUserInRequest) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserInRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *TransferUserInRequest) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *TransferUserInRequest) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TransferUserInRequest) GetAddress() TransferUserAddressInRequest {
	if o == nil || o.Address == nil {
		var ret TransferUserAddressInRequest
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferUserInRequest) GetAddressOk() (*TransferUserAddressInRequest, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TransferUserInRequest) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given TransferUserAddressInRequest and assigns it to the Address field.
func (o *TransferUserInRequest) SetAddress(v TransferUserAddressInRequest) {
	o.Address = &v
}

func (o TransferUserInRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["legal_name"] = o.LegalName
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.EmailAddress != nil {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferUserInRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTransferUserInRequest := _TransferUserInRequest{}

	if err = json.Unmarshal(bytes, &varTransferUserInRequest); err == nil {
		*o = TransferUserInRequest(varTransferUserInRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "legal_name")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "email_address")
		delete(additionalProperties, "address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferUserInRequest struct {
	value *TransferUserInRequest
	isSet bool
}

func (v NullableTransferUserInRequest) Get() *TransferUserInRequest {
	return v.value
}

func (v *NullableTransferUserInRequest) Set(val *TransferUserInRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferUserInRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferUserInRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferUserInRequest(val *TransferUserInRequest) *NullableTransferUserInRequest {
	return &NullableTransferUserInRequest{value: val, isSet: true}
}

func (v NullableTransferUserInRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferUserInRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


