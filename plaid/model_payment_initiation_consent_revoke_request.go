/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.97.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentInitiationConsentRevokeRequest PaymentInitiationConsentRevokeRequest defines the request schema for `/payment_initiation/consent/revoke`
type PaymentInitiationConsentRevokeRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The consent ID.
	ConsentId string `json:"consent_id"`
}

// NewPaymentInitiationConsentRevokeRequest instantiates a new PaymentInitiationConsentRevokeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationConsentRevokeRequest(consentId string) *PaymentInitiationConsentRevokeRequest {
	this := PaymentInitiationConsentRevokeRequest{}
	this.ConsentId = consentId
	return &this
}

// NewPaymentInitiationConsentRevokeRequestWithDefaults instantiates a new PaymentInitiationConsentRevokeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationConsentRevokeRequestWithDefaults() *PaymentInitiationConsentRevokeRequest {
	this := PaymentInitiationConsentRevokeRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationConsentRevokeRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentRevokeRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationConsentRevokeRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationConsentRevokeRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationConsentRevokeRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentRevokeRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationConsentRevokeRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationConsentRevokeRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetConsentId returns the ConsentId field value
func (o *PaymentInitiationConsentRevokeRequest) GetConsentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentId
}

// GetConsentIdOk returns a tuple with the ConsentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentRevokeRequest) GetConsentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsentId, true
}

// SetConsentId sets field value
func (o *PaymentInitiationConsentRevokeRequest) SetConsentId(v string) {
	o.ConsentId = v
}

func (o PaymentInitiationConsentRevokeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["consent_id"] = o.ConsentId
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationConsentRevokeRequest struct {
	value *PaymentInitiationConsentRevokeRequest
	isSet bool
}

func (v NullablePaymentInitiationConsentRevokeRequest) Get() *PaymentInitiationConsentRevokeRequest {
	return v.value
}

func (v *NullablePaymentInitiationConsentRevokeRequest) Set(val *PaymentInitiationConsentRevokeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsentRevokeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsentRevokeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsentRevokeRequest(val *PaymentInitiationConsentRevokeRequest) *NullablePaymentInitiationConsentRevokeRequest {
	return &NullablePaymentInitiationConsentRevokeRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsentRevokeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsentRevokeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


