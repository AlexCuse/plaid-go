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

// CreditSessionsGetResponse CreditSessionsGetResponse defines the response schema for `/credit/sessions/get`
type CreditSessionsGetResponse struct {
	// A list of Link sessions for the user. Sessions will be sorted in reverse chronological order.
	Sessions *[]CreditSession `json:"sessions,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditSessionsGetResponse CreditSessionsGetResponse

// NewCreditSessionsGetResponse instantiates a new CreditSessionsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditSessionsGetResponse(requestId string) *CreditSessionsGetResponse {
	this := CreditSessionsGetResponse{}
	this.RequestId = requestId
	return &this
}

// NewCreditSessionsGetResponseWithDefaults instantiates a new CreditSessionsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditSessionsGetResponseWithDefaults() *CreditSessionsGetResponse {
	this := CreditSessionsGetResponse{}
	return &this
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *CreditSessionsGetResponse) GetSessions() []CreditSession {
	if o == nil || o.Sessions == nil {
		var ret []CreditSession
		return ret
	}
	return *o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditSessionsGetResponse) GetSessionsOk() (*[]CreditSession, bool) {
	if o == nil || o.Sessions == nil {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *CreditSessionsGetResponse) HasSessions() bool {
	if o != nil && o.Sessions != nil {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []CreditSession and assigns it to the Sessions field.
func (o *CreditSessionsGetResponse) SetSessions(v []CreditSession) {
	o.Sessions = &v
}

// GetRequestId returns the RequestId field value
func (o *CreditSessionsGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditSessionsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditSessionsGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditSessionsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sessions != nil {
		toSerialize["sessions"] = o.Sessions
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditSessionsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditSessionsGetResponse := _CreditSessionsGetResponse{}

	if err = json.Unmarshal(bytes, &varCreditSessionsGetResponse); err == nil {
		*o = CreditSessionsGetResponse(varCreditSessionsGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sessions")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditSessionsGetResponse struct {
	value *CreditSessionsGetResponse
	isSet bool
}

func (v NullableCreditSessionsGetResponse) Get() *CreditSessionsGetResponse {
	return v.value
}

func (v *NullableCreditSessionsGetResponse) Set(val *CreditSessionsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditSessionsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditSessionsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditSessionsGetResponse(val *CreditSessionsGetResponse) *NullableCreditSessionsGetResponse {
	return &NullableCreditSessionsGetResponse{value: val, isSet: true}
}

func (v NullableCreditSessionsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditSessionsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


