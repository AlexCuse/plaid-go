/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransactionsRulesListRequest TransactionsRulesListRequest defines the request schema for `/beta/transactions/rules/v1/list`
type TransactionsRulesListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId string `json:"client_id"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret string `json:"secret"`
}

// NewTransactionsRulesListRequest instantiates a new TransactionsRulesListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRulesListRequest(clientId string, accessToken string, secret string) *TransactionsRulesListRequest {
	this := TransactionsRulesListRequest{}
	this.ClientId = clientId
	this.AccessToken = accessToken
	this.Secret = secret
	return &this
}

// NewTransactionsRulesListRequestWithDefaults instantiates a new TransactionsRulesListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRulesListRequestWithDefaults() *TransactionsRulesListRequest {
	this := TransactionsRulesListRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *TransactionsRulesListRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesListRequest) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *TransactionsRulesListRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetAccessToken returns the AccessToken field value
func (o *TransactionsRulesListRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesListRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TransactionsRulesListRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetSecret returns the Secret field value
func (o *TransactionsRulesListRequest) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesListRequest) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *TransactionsRulesListRequest) SetSecret(v string) {
	o.Secret = v
}

func (o TransactionsRulesListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsRulesListRequest struct {
	value *TransactionsRulesListRequest
	isSet bool
}

func (v NullableTransactionsRulesListRequest) Get() *TransactionsRulesListRequest {
	return v.value
}

func (v *NullableTransactionsRulesListRequest) Set(val *TransactionsRulesListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRulesListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRulesListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRulesListRequest(val *TransactionsRulesListRequest) *NullableTransactionsRulesListRequest {
	return &NullableTransactionsRulesListRequest{value: val, isSet: true}
}

func (v NullableTransactionsRulesListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRulesListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


