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

// ApplicationGetResponse ApplicationGetResponse defines the response schema for `/application/get`
type ApplicationGetResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	Application Application `json:"application"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationGetResponse ApplicationGetResponse

// NewApplicationGetResponse instantiates a new ApplicationGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationGetResponse(requestId string, application Application) *ApplicationGetResponse {
	this := ApplicationGetResponse{}
	this.RequestId = requestId
	this.Application = application
	return &this
}

// NewApplicationGetResponseWithDefaults instantiates a new ApplicationGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationGetResponseWithDefaults() *ApplicationGetResponse {
	this := ApplicationGetResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ApplicationGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ApplicationGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ApplicationGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetApplication returns the Application field value
func (o *ApplicationGetResponse) GetApplication() Application {
	if o == nil {
		var ret Application
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *ApplicationGetResponse) GetApplicationOk() (*Application, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *ApplicationGetResponse) SetApplication(v Application) {
	o.Application = v
}

func (o ApplicationGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["application"] = o.Application
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationGetResponse := _ApplicationGetResponse{}

	if err = json.Unmarshal(bytes, &varApplicationGetResponse); err == nil {
		*o = ApplicationGetResponse(varApplicationGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "application")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationGetResponse struct {
	value *ApplicationGetResponse
	isSet bool
}

func (v NullableApplicationGetResponse) Get() *ApplicationGetResponse {
	return v.value
}

func (v *NullableApplicationGetResponse) Set(val *ApplicationGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationGetResponse(val *ApplicationGetResponse) *NullableApplicationGetResponse {
	return &NullableApplicationGetResponse{value: val, isSet: true}
}

func (v NullableApplicationGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


