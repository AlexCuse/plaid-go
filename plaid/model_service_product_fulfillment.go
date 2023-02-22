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

// ServiceProductFulfillment A collection of details related to a fulfillment service or product in terms of request, process and result.
type ServiceProductFulfillment struct {
	SERVICE_PRODUCT_FULFILLMENT_DETAIL ServiceProductFulfillmentDetail `json:"SERVICE_PRODUCT_FULFILLMENT_DETAIL"`
	AdditionalProperties map[string]interface{}
}

type _ServiceProductFulfillment ServiceProductFulfillment

// NewServiceProductFulfillment instantiates a new ServiceProductFulfillment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProductFulfillment(sERVICEPRODUCTFULFILLMENTDETAIL ServiceProductFulfillmentDetail) *ServiceProductFulfillment {
	this := ServiceProductFulfillment{}
	this.SERVICE_PRODUCT_FULFILLMENT_DETAIL = sERVICEPRODUCTFULFILLMENTDETAIL
	return &this
}

// NewServiceProductFulfillmentWithDefaults instantiates a new ServiceProductFulfillment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProductFulfillmentWithDefaults() *ServiceProductFulfillment {
	this := ServiceProductFulfillment{}
	return &this
}

// GetSERVICE_PRODUCT_FULFILLMENT_DETAIL returns the SERVICE_PRODUCT_FULFILLMENT_DETAIL field value
func (o *ServiceProductFulfillment) GetSERVICE_PRODUCT_FULFILLMENT_DETAIL() ServiceProductFulfillmentDetail {
	if o == nil {
		var ret ServiceProductFulfillmentDetail
		return ret
	}

	return o.SERVICE_PRODUCT_FULFILLMENT_DETAIL
}

// GetSERVICE_PRODUCT_FULFILLMENT_DETAILOk returns a tuple with the SERVICE_PRODUCT_FULFILLMENT_DETAIL field value
// and a boolean to check if the value has been set.
func (o *ServiceProductFulfillment) GetSERVICE_PRODUCT_FULFILLMENT_DETAILOk() (*ServiceProductFulfillmentDetail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SERVICE_PRODUCT_FULFILLMENT_DETAIL, true
}

// SetSERVICE_PRODUCT_FULFILLMENT_DETAIL sets field value
func (o *ServiceProductFulfillment) SetSERVICE_PRODUCT_FULFILLMENT_DETAIL(v ServiceProductFulfillmentDetail) {
	o.SERVICE_PRODUCT_FULFILLMENT_DETAIL = v
}

func (o ServiceProductFulfillment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["SERVICE_PRODUCT_FULFILLMENT_DETAIL"] = o.SERVICE_PRODUCT_FULFILLMENT_DETAIL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceProductFulfillment) UnmarshalJSON(bytes []byte) (err error) {
	varServiceProductFulfillment := _ServiceProductFulfillment{}

	if err = json.Unmarshal(bytes, &varServiceProductFulfillment); err == nil {
		*o = ServiceProductFulfillment(varServiceProductFulfillment)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "SERVICE_PRODUCT_FULFILLMENT_DETAIL")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceProductFulfillment struct {
	value *ServiceProductFulfillment
	isSet bool
}

func (v NullableServiceProductFulfillment) Get() *ServiceProductFulfillment {
	return v.value
}

func (v *NullableServiceProductFulfillment) Set(val *ServiceProductFulfillment) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProductFulfillment) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProductFulfillment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProductFulfillment(val *ServiceProductFulfillment) *NullableServiceProductFulfillment {
	return &NullableServiceProductFulfillment{value: val, isSet: true}
}

func (v NullableServiceProductFulfillment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProductFulfillment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


