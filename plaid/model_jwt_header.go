/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.77.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// JWTHeader A JWT Header, used for webhook validation
type JWTHeader struct {
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _JWTHeader JWTHeader

// NewJWTHeader instantiates a new JWTHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJWTHeader(id string) *JWTHeader {
	this := JWTHeader{}
	this.Id = id
	return &this
}

// NewJWTHeaderWithDefaults instantiates a new JWTHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWTHeaderWithDefaults() *JWTHeader {
	this := JWTHeader{}
	return &this
}

// GetId returns the Id field value
func (o *JWTHeader) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JWTHeader) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JWTHeader) SetId(v string) {
	o.Id = v
}

func (o JWTHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *JWTHeader) UnmarshalJSON(bytes []byte) (err error) {
	varJWTHeader := _JWTHeader{}

	if err = json.Unmarshal(bytes, &varJWTHeader); err == nil {
		*o = JWTHeader(varJWTHeader)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJWTHeader struct {
	value *JWTHeader
	isSet bool
}

func (v NullableJWTHeader) Get() *JWTHeader {
	return v.value
}

func (v *NullableJWTHeader) Set(val *JWTHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableJWTHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableJWTHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJWTHeader(val *JWTHeader) *NullableJWTHeader {
	return &NullableJWTHeader{value: val, isSet: true}
}

func (v NullableJWTHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJWTHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


