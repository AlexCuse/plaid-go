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

// CreditFreddieVerificationOfEmploymentDealVOE25 An object representing a Verification of Employment report.
type CreditFreddieVerificationOfEmploymentDealVOE25 struct {
	LOANS CreditFreddieMacLoansVOA24 `json:"LOANS"`
	PARTIES CreditFreddieMacPartiesVOA24 `json:"PARTIES"`
	SERVICES CreditFreddieMacServicesVOE25 `json:"SERVICES"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieVerificationOfEmploymentDealVOE25 CreditFreddieVerificationOfEmploymentDealVOE25

// NewCreditFreddieVerificationOfEmploymentDealVOE25 instantiates a new CreditFreddieVerificationOfEmploymentDealVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieVerificationOfEmploymentDealVOE25(lOANS CreditFreddieMacLoansVOA24, pARTIES CreditFreddieMacPartiesVOA24, sERVICES CreditFreddieMacServicesVOE25) *CreditFreddieVerificationOfEmploymentDealVOE25 {
	this := CreditFreddieVerificationOfEmploymentDealVOE25{}
	this.LOANS = lOANS
	this.PARTIES = pARTIES
	this.SERVICES = sERVICES
	return &this
}

// NewCreditFreddieVerificationOfEmploymentDealVOE25WithDefaults instantiates a new CreditFreddieVerificationOfEmploymentDealVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieVerificationOfEmploymentDealVOE25WithDefaults() *CreditFreddieVerificationOfEmploymentDealVOE25 {
	this := CreditFreddieVerificationOfEmploymentDealVOE25{}
	return &this
}

// GetLOANS returns the LOANS field value
func (o *CreditFreddieVerificationOfEmploymentDealVOE25) GetLOANS() CreditFreddieMacLoansVOA24 {
	if o == nil {
		var ret CreditFreddieMacLoansVOA24
		return ret
	}

	return o.LOANS
}

// GetLOANSOk returns a tuple with the LOANS field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieVerificationOfEmploymentDealVOE25) GetLOANSOk() (*CreditFreddieMacLoansVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LOANS, true
}

// SetLOANS sets field value
func (o *CreditFreddieVerificationOfEmploymentDealVOE25) SetLOANS(v CreditFreddieMacLoansVOA24) {
	o.LOANS = v
}

// GetPARTIES returns the PARTIES field value
func (o *CreditFreddieVerificationOfEmploymentDealVOE25) GetPARTIES() CreditFreddieMacPartiesVOA24 {
	if o == nil {
		var ret CreditFreddieMacPartiesVOA24
		return ret
	}

	return o.PARTIES
}

// GetPARTIESOk returns a tuple with the PARTIES field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieVerificationOfEmploymentDealVOE25) GetPARTIESOk() (*CreditFreddieMacPartiesVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PARTIES, true
}

// SetPARTIES sets field value
func (o *CreditFreddieVerificationOfEmploymentDealVOE25) SetPARTIES(v CreditFreddieMacPartiesVOA24) {
	o.PARTIES = v
}

// GetSERVICES returns the SERVICES field value
func (o *CreditFreddieVerificationOfEmploymentDealVOE25) GetSERVICES() CreditFreddieMacServicesVOE25 {
	if o == nil {
		var ret CreditFreddieMacServicesVOE25
		return ret
	}

	return o.SERVICES
}

// GetSERVICESOk returns a tuple with the SERVICES field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieVerificationOfEmploymentDealVOE25) GetSERVICESOk() (*CreditFreddieMacServicesVOE25, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SERVICES, true
}

// SetSERVICES sets field value
func (o *CreditFreddieVerificationOfEmploymentDealVOE25) SetSERVICES(v CreditFreddieMacServicesVOE25) {
	o.SERVICES = v
}

func (o CreditFreddieVerificationOfEmploymentDealVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LOANS"] = o.LOANS
	}
	if true {
		toSerialize["PARTIES"] = o.PARTIES
	}
	if true {
		toSerialize["SERVICES"] = o.SERVICES
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieVerificationOfEmploymentDealVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieVerificationOfEmploymentDealVOE25 := _CreditFreddieVerificationOfEmploymentDealVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieVerificationOfEmploymentDealVOE25); err == nil {
		*o = CreditFreddieVerificationOfEmploymentDealVOE25(varCreditFreddieVerificationOfEmploymentDealVOE25)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LOANS")
		delete(additionalProperties, "PARTIES")
		delete(additionalProperties, "SERVICES")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieVerificationOfEmploymentDealVOE25 struct {
	value *CreditFreddieVerificationOfEmploymentDealVOE25
	isSet bool
}

func (v NullableCreditFreddieVerificationOfEmploymentDealVOE25) Get() *CreditFreddieVerificationOfEmploymentDealVOE25 {
	return v.value
}

func (v *NullableCreditFreddieVerificationOfEmploymentDealVOE25) Set(val *CreditFreddieVerificationOfEmploymentDealVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieVerificationOfEmploymentDealVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieVerificationOfEmploymentDealVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieVerificationOfEmploymentDealVOE25(val *CreditFreddieVerificationOfEmploymentDealVOE25) *NullableCreditFreddieVerificationOfEmploymentDealVOE25 {
	return &NullableCreditFreddieVerificationOfEmploymentDealVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieVerificationOfEmploymentDealVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieVerificationOfEmploymentDealVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


