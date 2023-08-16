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

// CreditEmploymentGetRequest CreditEmploymentGetRequest defines the request schema for `/credit/employment/get`.
type CreditEmploymentGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken string `json:"user_token"`
}

// NewCreditEmploymentGetRequest instantiates a new CreditEmploymentGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditEmploymentGetRequest(userToken string) *CreditEmploymentGetRequest {
	this := CreditEmploymentGetRequest{}
	this.UserToken = userToken
	return &this
}

// NewCreditEmploymentGetRequestWithDefaults instantiates a new CreditEmploymentGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditEmploymentGetRequestWithDefaults() *CreditEmploymentGetRequest {
	this := CreditEmploymentGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditEmploymentGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditEmploymentGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditEmploymentGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditEmploymentGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditEmploymentGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditEmploymentGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditEmploymentGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditEmploymentGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value
func (o *CreditEmploymentGetRequest) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *CreditEmploymentGetRequest) GetUserTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *CreditEmploymentGetRequest) SetUserToken(v string) {
	o.UserToken = v
}

func (o CreditEmploymentGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["user_token"] = o.UserToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreditEmploymentGetRequest struct {
	value *CreditEmploymentGetRequest
	isSet bool
}

func (v NullableCreditEmploymentGetRequest) Get() *CreditEmploymentGetRequest {
	return v.value
}

func (v *NullableCreditEmploymentGetRequest) Set(val *CreditEmploymentGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditEmploymentGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditEmploymentGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditEmploymentGetRequest(val *CreditEmploymentGetRequest) *NullableCreditEmploymentGetRequest {
	return &NullableCreditEmploymentGetRequest{value: val, isSet: true}
}

func (v NullableCreditEmploymentGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditEmploymentGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

