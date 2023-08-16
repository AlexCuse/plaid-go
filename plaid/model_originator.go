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

// Originator Originator and their status.
type Originator struct {
	// Originator’s client ID.
	ClientId string `json:"client_id"`
	TransferDiligenceStatus TransferDiligenceStatus `json:"transfer_diligence_status"`
	AdditionalProperties map[string]interface{}
}

type _Originator Originator

// NewOriginator instantiates a new Originator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginator(clientId string, transferDiligenceStatus TransferDiligenceStatus) *Originator {
	this := Originator{}
	this.ClientId = clientId
	this.TransferDiligenceStatus = transferDiligenceStatus
	return &this
}

// NewOriginatorWithDefaults instantiates a new Originator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginatorWithDefaults() *Originator {
	this := Originator{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *Originator) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *Originator) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *Originator) SetClientId(v string) {
	o.ClientId = v
}

// GetTransferDiligenceStatus returns the TransferDiligenceStatus field value
func (o *Originator) GetTransferDiligenceStatus() TransferDiligenceStatus {
	if o == nil {
		var ret TransferDiligenceStatus
		return ret
	}

	return o.TransferDiligenceStatus
}

// GetTransferDiligenceStatusOk returns a tuple with the TransferDiligenceStatus field value
// and a boolean to check if the value has been set.
func (o *Originator) GetTransferDiligenceStatusOk() (*TransferDiligenceStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferDiligenceStatus, true
}

// SetTransferDiligenceStatus sets field value
func (o *Originator) SetTransferDiligenceStatus(v TransferDiligenceStatus) {
	o.TransferDiligenceStatus = v
}

func (o Originator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["transfer_diligence_status"] = o.TransferDiligenceStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Originator) UnmarshalJSON(bytes []byte) (err error) {
	varOriginator := _Originator{}

	if err = json.Unmarshal(bytes, &varOriginator); err == nil {
		*o = Originator(varOriginator)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "transfer_diligence_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOriginator struct {
	value *Originator
	isSet bool
}

func (v NullableOriginator) Get() *Originator {
	return v.value
}

func (v *NullableOriginator) Set(val *Originator) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginator) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginator(val *Originator) *NullableOriginator {
	return &NullableOriginator{value: val, isSet: true}
}

func (v NullableOriginator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

