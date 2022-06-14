/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditPayStubVerification An object containing details on the paystub's verification status. This object will only be populated if the [`income_verification.access_tokens`](/docs/api/tokens/#link-token-create-request-income-verification-access-tokens) parameter was provided during the `/link/token/create` call or if a problem was detected with the information supplied by the user; otherwise it will be `null`.
type CreditPayStubVerification struct {
	// Derived verification status.
	VerificationStatus NullableString `json:"verification_status"`
	VerificationAttributes []PayStubVerificationAttribute `json:"verification_attributes"`
	AdditionalProperties map[string]interface{}
}

type _CreditPayStubVerification CreditPayStubVerification

// NewCreditPayStubVerification instantiates a new CreditPayStubVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayStubVerification(verificationStatus NullableString, verificationAttributes []PayStubVerificationAttribute) *CreditPayStubVerification {
	this := CreditPayStubVerification{}
	this.VerificationStatus = verificationStatus
	this.VerificationAttributes = verificationAttributes
	return &this
}

// NewCreditPayStubVerificationWithDefaults instantiates a new CreditPayStubVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayStubVerificationWithDefaults() *CreditPayStubVerification {
	this := CreditPayStubVerification{}
	return &this
}

// GetVerificationStatus returns the VerificationStatus field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPayStubVerification) GetVerificationStatus() string {
	if o == nil || o.VerificationStatus.Get() == nil {
		var ret string
		return ret
	}

	return *o.VerificationStatus.Get()
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubVerification) GetVerificationStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VerificationStatus.Get(), o.VerificationStatus.IsSet()
}

// SetVerificationStatus sets field value
func (o *CreditPayStubVerification) SetVerificationStatus(v string) {
	o.VerificationStatus.Set(&v)
}

// GetVerificationAttributes returns the VerificationAttributes field value
func (o *CreditPayStubVerification) GetVerificationAttributes() []PayStubVerificationAttribute {
	if o == nil {
		var ret []PayStubVerificationAttribute
		return ret
	}

	return o.VerificationAttributes
}

// GetVerificationAttributesOk returns a tuple with the VerificationAttributes field value
// and a boolean to check if the value has been set.
func (o *CreditPayStubVerification) GetVerificationAttributesOk() (*[]PayStubVerificationAttribute, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationAttributes, true
}

// SetVerificationAttributes sets field value
func (o *CreditPayStubVerification) SetVerificationAttributes(v []PayStubVerificationAttribute) {
	o.VerificationAttributes = v
}

func (o CreditPayStubVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus.Get()
	}
	if true {
		toSerialize["verification_attributes"] = o.VerificationAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditPayStubVerification) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPayStubVerification := _CreditPayStubVerification{}

	if err = json.Unmarshal(bytes, &varCreditPayStubVerification); err == nil {
		*o = CreditPayStubVerification(varCreditPayStubVerification)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "verification_status")
		delete(additionalProperties, "verification_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPayStubVerification struct {
	value *CreditPayStubVerification
	isSet bool
}

func (v NullableCreditPayStubVerification) Get() *CreditPayStubVerification {
	return v.value
}

func (v *NullableCreditPayStubVerification) Set(val *CreditPayStubVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayStubVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayStubVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayStubVerification(val *CreditPayStubVerification) *NullableCreditPayStubVerification {
	return &NullableCreditPayStubVerification{value: val, isSet: true}
}

func (v NullableCreditPayStubVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayStubVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


