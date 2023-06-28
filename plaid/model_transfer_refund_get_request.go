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

// TransferRefundGetRequest Defines the request schema for `/transfer/refund/get`
type TransferRefundGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Plaid’s unique identifier for a refund.
	RefundId string `json:"refund_id"`
}

// NewTransferRefundGetRequest instantiates a new TransferRefundGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRefundGetRequest(refundId string) *TransferRefundGetRequest {
	this := TransferRefundGetRequest{}
	this.RefundId = refundId
	return &this
}

// NewTransferRefundGetRequestWithDefaults instantiates a new TransferRefundGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRefundGetRequestWithDefaults() *TransferRefundGetRequest {
	this := TransferRefundGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferRefundGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRefundGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferRefundGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferRefundGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferRefundGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRefundGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferRefundGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferRefundGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetRefundId returns the RefundId field value
func (o *TransferRefundGetRequest) GetRefundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value
// and a boolean to check if the value has been set.
func (o *TransferRefundGetRequest) GetRefundIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RefundId, true
}

// SetRefundId sets field value
func (o *TransferRefundGetRequest) SetRefundId(v string) {
	o.RefundId = v
}

func (o TransferRefundGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["refund_id"] = o.RefundId
	}
	return json.Marshal(toSerialize)
}

type NullableTransferRefundGetRequest struct {
	value *TransferRefundGetRequest
	isSet bool
}

func (v NullableTransferRefundGetRequest) Get() *TransferRefundGetRequest {
	return v.value
}

func (v *NullableTransferRefundGetRequest) Set(val *TransferRefundGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRefundGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRefundGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRefundGetRequest(val *TransferRefundGetRequest) *NullableTransferRefundGetRequest {
	return &NullableTransferRefundGetRequest{value: val, isSet: true}
}

func (v NullableTransferRefundGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRefundGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


