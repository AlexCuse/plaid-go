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

// ApplicationGetRequest ApplicationGetRequest defines the schema for `/application/get`
type ApplicationGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId string `json:"client_id"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret string `json:"secret"`
	// This field will map to the application ID that is returned from /item/applications/list, or provided to the institution in an oauth redirect.
	ApplicationId string `json:"application_id"`
}

// NewApplicationGetRequest instantiates a new ApplicationGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationGetRequest(clientId string, secret string, applicationId string) *ApplicationGetRequest {
	this := ApplicationGetRequest{}
	this.ClientId = clientId
	this.Secret = secret
	this.ApplicationId = applicationId
	return &this
}

// NewApplicationGetRequestWithDefaults instantiates a new ApplicationGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationGetRequestWithDefaults() *ApplicationGetRequest {
	this := ApplicationGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *ApplicationGetRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ApplicationGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ApplicationGetRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetSecret returns the Secret field value
func (o *ApplicationGetRequest) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *ApplicationGetRequest) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *ApplicationGetRequest) SetSecret(v string) {
	o.Secret = v
}

// GetApplicationId returns the ApplicationId field value
func (o *ApplicationGetRequest) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *ApplicationGetRequest) GetApplicationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *ApplicationGetRequest) SetApplicationId(v string) {
	o.ApplicationId = v
}

func (o ApplicationGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["application_id"] = o.ApplicationId
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationGetRequest struct {
	value *ApplicationGetRequest
	isSet bool
}

func (v NullableApplicationGetRequest) Get() *ApplicationGetRequest {
	return v.value
}

func (v *NullableApplicationGetRequest) Set(val *ApplicationGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationGetRequest(val *ApplicationGetRequest) *NullableApplicationGetRequest {
	return &NullableApplicationGetRequest{value: val, isSet: true}
}

func (v NullableApplicationGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


