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

// PaymentInitiationPaymentGetRequest PaymentInitiationPaymentGetRequest defines the request schema for `/payment_initiation/payment/get`
type PaymentInitiationPaymentGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The `payment_id` returned from `/payment_initiation/payment/create`.
	PaymentId string `json:"payment_id"`
}

// NewPaymentInitiationPaymentGetRequest instantiates a new PaymentInitiationPaymentGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentGetRequest(paymentId string) *PaymentInitiationPaymentGetRequest {
	this := PaymentInitiationPaymentGetRequest{}
	this.PaymentId = paymentId
	return &this
}

// NewPaymentInitiationPaymentGetRequestWithDefaults instantiates a new PaymentInitiationPaymentGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentGetRequestWithDefaults() *PaymentInitiationPaymentGetRequest {
	this := PaymentInitiationPaymentGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationPaymentGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationPaymentGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentInitiationPaymentGetRequest) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentInitiationPaymentGetRequest) SetPaymentId(v string) {
	o.PaymentId = v
}

func (o PaymentInitiationPaymentGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["payment_id"] = o.PaymentId
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationPaymentGetRequest struct {
	value *PaymentInitiationPaymentGetRequest
	isSet bool
}

func (v NullablePaymentInitiationPaymentGetRequest) Get() *PaymentInitiationPaymentGetRequest {
	return v.value
}

func (v *NullablePaymentInitiationPaymentGetRequest) Set(val *PaymentInitiationPaymentGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentGetRequest(val *PaymentInitiationPaymentGetRequest) *NullablePaymentInitiationPaymentGetRequest {
	return &NullablePaymentInitiationPaymentGetRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


