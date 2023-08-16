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

// PartnerCustomerRemoveResponse Response schema for `/partner/customer/remove`.
type PartnerCustomerRemoveResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId *string `json:"request_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerCustomerRemoveResponse PartnerCustomerRemoveResponse

// NewPartnerCustomerRemoveResponse instantiates a new PartnerCustomerRemoveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerCustomerRemoveResponse() *PartnerCustomerRemoveResponse {
	this := PartnerCustomerRemoveResponse{}
	return &this
}

// NewPartnerCustomerRemoveResponseWithDefaults instantiates a new PartnerCustomerRemoveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerCustomerRemoveResponseWithDefaults() *PartnerCustomerRemoveResponse {
	this := PartnerCustomerRemoveResponse{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *PartnerCustomerRemoveResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerRemoveResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *PartnerCustomerRemoveResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *PartnerCustomerRemoveResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o PartnerCustomerRemoveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerCustomerRemoveResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerCustomerRemoveResponse := _PartnerCustomerRemoveResponse{}

	if err = json.Unmarshal(bytes, &varPartnerCustomerRemoveResponse); err == nil {
		*o = PartnerCustomerRemoveResponse(varPartnerCustomerRemoveResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerCustomerRemoveResponse struct {
	value *PartnerCustomerRemoveResponse
	isSet bool
}

func (v NullablePartnerCustomerRemoveResponse) Get() *PartnerCustomerRemoveResponse {
	return v.value
}

func (v *NullablePartnerCustomerRemoveResponse) Set(val *PartnerCustomerRemoveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerCustomerRemoveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerCustomerRemoveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerCustomerRemoveResponse(val *PartnerCustomerRemoveResponse) *NullablePartnerCustomerRemoveResponse {
	return &NullablePartnerCustomerRemoveResponse{value: val, isSet: true}
}

func (v NullablePartnerCustomerRemoveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerCustomerRemoveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


