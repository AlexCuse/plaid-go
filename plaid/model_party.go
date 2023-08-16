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

// Party A collection of information about a single party to a transaction. Included direct participants like the borrower and seller as well as indirect participants such as the flood certificate provider.
type Party struct {
	INDIVIDUAL PartyIndividual `json:"INDIVIDUAL"`
	ROLES Roles `json:"ROLES"`
	TAXPAYER_IDENTIFIERS TaxpayerIdentifiers `json:"TAXPAYER_IDENTIFIERS"`
	AdditionalProperties map[string]interface{}
}

type _Party Party

// NewParty instantiates a new Party object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParty(iNDIVIDUAL PartyIndividual, rOLES Roles, tAXPAYERIDENTIFIERS TaxpayerIdentifiers) *Party {
	this := Party{}
	this.INDIVIDUAL = iNDIVIDUAL
	this.ROLES = rOLES
	this.TAXPAYER_IDENTIFIERS = tAXPAYERIDENTIFIERS
	return &this
}

// NewPartyWithDefaults instantiates a new Party object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyWithDefaults() *Party {
	this := Party{}
	return &this
}

// GetINDIVIDUAL returns the INDIVIDUAL field value
func (o *Party) GetINDIVIDUAL() PartyIndividual {
	if o == nil {
		var ret PartyIndividual
		return ret
	}

	return o.INDIVIDUAL
}

// GetINDIVIDUALOk returns a tuple with the INDIVIDUAL field value
// and a boolean to check if the value has been set.
func (o *Party) GetINDIVIDUALOk() (*PartyIndividual, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.INDIVIDUAL, true
}

// SetINDIVIDUAL sets field value
func (o *Party) SetINDIVIDUAL(v PartyIndividual) {
	o.INDIVIDUAL = v
}

// GetROLES returns the ROLES field value
func (o *Party) GetROLES() Roles {
	if o == nil {
		var ret Roles
		return ret
	}

	return o.ROLES
}

// GetROLESOk returns a tuple with the ROLES field value
// and a boolean to check if the value has been set.
func (o *Party) GetROLESOk() (*Roles, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ROLES, true
}

// SetROLES sets field value
func (o *Party) SetROLES(v Roles) {
	o.ROLES = v
}

// GetTAXPAYER_IDENTIFIERS returns the TAXPAYER_IDENTIFIERS field value
func (o *Party) GetTAXPAYER_IDENTIFIERS() TaxpayerIdentifiers {
	if o == nil {
		var ret TaxpayerIdentifiers
		return ret
	}

	return o.TAXPAYER_IDENTIFIERS
}

// GetTAXPAYER_IDENTIFIERSOk returns a tuple with the TAXPAYER_IDENTIFIERS field value
// and a boolean to check if the value has been set.
func (o *Party) GetTAXPAYER_IDENTIFIERSOk() (*TaxpayerIdentifiers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TAXPAYER_IDENTIFIERS, true
}

// SetTAXPAYER_IDENTIFIERS sets field value
func (o *Party) SetTAXPAYER_IDENTIFIERS(v TaxpayerIdentifiers) {
	o.TAXPAYER_IDENTIFIERS = v
}

func (o Party) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["INDIVIDUAL"] = o.INDIVIDUAL
	}
	if true {
		toSerialize["ROLES"] = o.ROLES
	}
	if true {
		toSerialize["TAXPAYER_IDENTIFIERS"] = o.TAXPAYER_IDENTIFIERS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Party) UnmarshalJSON(bytes []byte) (err error) {
	varParty := _Party{}

	if err = json.Unmarshal(bytes, &varParty); err == nil {
		*o = Party(varParty)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "INDIVIDUAL")
		delete(additionalProperties, "ROLES")
		delete(additionalProperties, "TAXPAYER_IDENTIFIERS")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableParty struct {
	value *Party
	isSet bool
}

func (v NullableParty) Get() *Party {
	return v.value
}

func (v *NullableParty) Set(val *Party) {
	v.value = val
	v.isSet = true
}

func (v NullableParty) IsSet() bool {
	return v.isSet
}

func (v *NullableParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParty(val *Party) *NullableParty {
	return &NullableParty{value: val, isSet: true}
}

func (v NullableParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


