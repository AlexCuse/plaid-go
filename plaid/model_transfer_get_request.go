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

// TransferGetRequest Defines the request schema for `/transfer/get`
type TransferGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Plaid’s unique identifier for a transfer.
	TransferId string `json:"transfer_id"`
	// The Plaid client ID of the transfer originator. Should only be present if `client_id` is a third-party sender (TPS).
	OriginatorClientId NullableString `json:"originator_client_id,omitempty"`
}

// NewTransferGetRequest instantiates a new TransferGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferGetRequest(transferId string) *TransferGetRequest {
	this := TransferGetRequest{}
	this.TransferId = transferId
	return &this
}

// NewTransferGetRequestWithDefaults instantiates a new TransferGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferGetRequestWithDefaults() *TransferGetRequest {
	this := TransferGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetTransferId returns the TransferId field value
func (o *TransferGetRequest) GetTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
func (o *TransferGetRequest) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferId, true
}

// SetTransferId sets field value
func (o *TransferGetRequest) SetTransferId(v string) {
	o.TransferId = v
}

// GetOriginatorClientId returns the OriginatorClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferGetRequest) GetOriginatorClientId() string {
	if o == nil || o.OriginatorClientId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginatorClientId.Get()
}

// GetOriginatorClientIdOk returns a tuple with the OriginatorClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferGetRequest) GetOriginatorClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginatorClientId.Get(), o.OriginatorClientId.IsSet()
}

// HasOriginatorClientId returns a boolean if a field has been set.
func (o *TransferGetRequest) HasOriginatorClientId() bool {
	if o != nil && o.OriginatorClientId.IsSet() {
		return true
	}

	return false
}

// SetOriginatorClientId gets a reference to the given NullableString and assigns it to the OriginatorClientId field.
func (o *TransferGetRequest) SetOriginatorClientId(v string) {
	o.OriginatorClientId.Set(&v)
}
// SetOriginatorClientIdNil sets the value for OriginatorClientId to be an explicit nil
func (o *TransferGetRequest) SetOriginatorClientIdNil() {
	o.OriginatorClientId.Set(nil)
}

// UnsetOriginatorClientId ensures that no value is present for OriginatorClientId, not even an explicit nil
func (o *TransferGetRequest) UnsetOriginatorClientId() {
	o.OriginatorClientId.Unset()
}

func (o TransferGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["transfer_id"] = o.TransferId
	}
	if o.OriginatorClientId.IsSet() {
		toSerialize["originator_client_id"] = o.OriginatorClientId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransferGetRequest struct {
	value *TransferGetRequest
	isSet bool
}

func (v NullableTransferGetRequest) Get() *TransferGetRequest {
	return v.value
}

func (v *NullableTransferGetRequest) Set(val *TransferGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferGetRequest(val *TransferGetRequest) *NullableTransferGetRequest {
	return &NullableTransferGetRequest{value: val, isSet: true}
}

func (v NullableTransferGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


