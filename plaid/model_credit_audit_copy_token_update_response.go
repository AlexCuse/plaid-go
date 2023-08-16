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

// CreditAuditCopyTokenUpdateResponse CreditAuditCopyTokenUpdateResponse defines the response schema for `/credit/audit_copy_token/update`
type CreditAuditCopyTokenUpdateResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	// `true` if the Audit Copy Token was successfully updated.
	Updated bool `json:"updated"`
	AdditionalProperties map[string]interface{}
}

type _CreditAuditCopyTokenUpdateResponse CreditAuditCopyTokenUpdateResponse

// NewCreditAuditCopyTokenUpdateResponse instantiates a new CreditAuditCopyTokenUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditAuditCopyTokenUpdateResponse(requestId string, updated bool) *CreditAuditCopyTokenUpdateResponse {
	this := CreditAuditCopyTokenUpdateResponse{}
	this.RequestId = requestId
	this.Updated = updated
	return &this
}

// NewCreditAuditCopyTokenUpdateResponseWithDefaults instantiates a new CreditAuditCopyTokenUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditAuditCopyTokenUpdateResponseWithDefaults() *CreditAuditCopyTokenUpdateResponse {
	this := CreditAuditCopyTokenUpdateResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *CreditAuditCopyTokenUpdateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenUpdateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditAuditCopyTokenUpdateResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetUpdated returns the Updated field value
func (o *CreditAuditCopyTokenUpdateResponse) GetUpdated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenUpdateResponse) GetUpdatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *CreditAuditCopyTokenUpdateResponse) SetUpdated(v bool) {
	o.Updated = v
}

func (o CreditAuditCopyTokenUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["updated"] = o.Updated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditAuditCopyTokenUpdateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditAuditCopyTokenUpdateResponse := _CreditAuditCopyTokenUpdateResponse{}

	if err = json.Unmarshal(bytes, &varCreditAuditCopyTokenUpdateResponse); err == nil {
		*o = CreditAuditCopyTokenUpdateResponse(varCreditAuditCopyTokenUpdateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditAuditCopyTokenUpdateResponse struct {
	value *CreditAuditCopyTokenUpdateResponse
	isSet bool
}

func (v NullableCreditAuditCopyTokenUpdateResponse) Get() *CreditAuditCopyTokenUpdateResponse {
	return v.value
}

func (v *NullableCreditAuditCopyTokenUpdateResponse) Set(val *CreditAuditCopyTokenUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditAuditCopyTokenUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditAuditCopyTokenUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditAuditCopyTokenUpdateResponse(val *CreditAuditCopyTokenUpdateResponse) *NullableCreditAuditCopyTokenUpdateResponse {
	return &NullableCreditAuditCopyTokenUpdateResponse{value: val, isSet: true}
}

func (v NullableCreditAuditCopyTokenUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditAuditCopyTokenUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


