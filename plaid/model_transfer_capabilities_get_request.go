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

// TransferCapabilitiesGetRequest Defines the request schema for `/transfer/capabilities/get`
type TransferCapabilitiesGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The Plaid `access_token` for the account that will be debited or credited. Required if not using `payment_profile_token`.
	AccessToken *string `json:"access_token,omitempty"`
	// The Plaid `account_id` corresponding to the end-user account that will be debited or credited. Required when creating a transfer using an `access_token`.
	AccountId *string `json:"account_id,omitempty"`
	// A payment profile token associated with the Payment Profile data that is being requested.
	PaymentProfileToken *string `json:"payment_profile_token,omitempty"`
}

// NewTransferCapabilitiesGetRequest instantiates a new TransferCapabilitiesGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCapabilitiesGetRequest() *TransferCapabilitiesGetRequest {
	this := TransferCapabilitiesGetRequest{}
	return &this
}

// NewTransferCapabilitiesGetRequestWithDefaults instantiates a new TransferCapabilitiesGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCapabilitiesGetRequestWithDefaults() *TransferCapabilitiesGetRequest {
	this := TransferCapabilitiesGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferCapabilitiesGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCapabilitiesGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferCapabilitiesGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferCapabilitiesGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferCapabilitiesGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCapabilitiesGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferCapabilitiesGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferCapabilitiesGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TransferCapabilitiesGetRequest) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCapabilitiesGetRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TransferCapabilitiesGetRequest) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TransferCapabilitiesGetRequest) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TransferCapabilitiesGetRequest) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCapabilitiesGetRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferCapabilitiesGetRequest) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *TransferCapabilitiesGetRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetPaymentProfileToken returns the PaymentProfileToken field value if set, zero value otherwise.
func (o *TransferCapabilitiesGetRequest) GetPaymentProfileToken() string {
	if o == nil || o.PaymentProfileToken == nil {
		var ret string
		return ret
	}
	return *o.PaymentProfileToken
}

// GetPaymentProfileTokenOk returns a tuple with the PaymentProfileToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCapabilitiesGetRequest) GetPaymentProfileTokenOk() (*string, bool) {
	if o == nil || o.PaymentProfileToken == nil {
		return nil, false
	}
	return o.PaymentProfileToken, true
}

// HasPaymentProfileToken returns a boolean if a field has been set.
func (o *TransferCapabilitiesGetRequest) HasPaymentProfileToken() bool {
	if o != nil && o.PaymentProfileToken != nil {
		return true
	}

	return false
}

// SetPaymentProfileToken gets a reference to the given string and assigns it to the PaymentProfileToken field.
func (o *TransferCapabilitiesGetRequest) SetPaymentProfileToken(v string) {
	o.PaymentProfileToken = &v
}

func (o TransferCapabilitiesGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.PaymentProfileToken != nil {
		toSerialize["payment_profile_token"] = o.PaymentProfileToken
	}
	return json.Marshal(toSerialize)
}

type NullableTransferCapabilitiesGetRequest struct {
	value *TransferCapabilitiesGetRequest
	isSet bool
}

func (v NullableTransferCapabilitiesGetRequest) Get() *TransferCapabilitiesGetRequest {
	return v.value
}

func (v *NullableTransferCapabilitiesGetRequest) Set(val *TransferCapabilitiesGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCapabilitiesGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCapabilitiesGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCapabilitiesGetRequest(val *TransferCapabilitiesGetRequest) *NullableTransferCapabilitiesGetRequest {
	return &NullableTransferCapabilitiesGetRequest{value: val, isSet: true}
}

func (v NullableTransferCapabilitiesGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCapabilitiesGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

