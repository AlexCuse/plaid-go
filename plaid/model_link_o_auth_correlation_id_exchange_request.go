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

// LinkOAuthCorrelationIdExchangeRequest LinkOAuthCorrelationIdExchangeRequest defines the request schema for `/link/oauth/correlation_id/exchange`
type LinkOAuthCorrelationIdExchangeRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A `link_correlation_id` from a received OAuth redirect URI callback
	LinkCorrelationId string `json:"link_correlation_id"`
}

// NewLinkOAuthCorrelationIdExchangeRequest instantiates a new LinkOAuthCorrelationIdExchangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkOAuthCorrelationIdExchangeRequest(linkCorrelationId string) *LinkOAuthCorrelationIdExchangeRequest {
	this := LinkOAuthCorrelationIdExchangeRequest{}
	this.LinkCorrelationId = linkCorrelationId
	return &this
}

// NewLinkOAuthCorrelationIdExchangeRequestWithDefaults instantiates a new LinkOAuthCorrelationIdExchangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkOAuthCorrelationIdExchangeRequestWithDefaults() *LinkOAuthCorrelationIdExchangeRequest {
	this := LinkOAuthCorrelationIdExchangeRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *LinkOAuthCorrelationIdExchangeRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkOAuthCorrelationIdExchangeRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *LinkOAuthCorrelationIdExchangeRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *LinkOAuthCorrelationIdExchangeRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *LinkOAuthCorrelationIdExchangeRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkOAuthCorrelationIdExchangeRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *LinkOAuthCorrelationIdExchangeRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *LinkOAuthCorrelationIdExchangeRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetLinkCorrelationId returns the LinkCorrelationId field value
func (o *LinkOAuthCorrelationIdExchangeRequest) GetLinkCorrelationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkCorrelationId
}

// GetLinkCorrelationIdOk returns a tuple with the LinkCorrelationId field value
// and a boolean to check if the value has been set.
func (o *LinkOAuthCorrelationIdExchangeRequest) GetLinkCorrelationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkCorrelationId, true
}

// SetLinkCorrelationId sets field value
func (o *LinkOAuthCorrelationIdExchangeRequest) SetLinkCorrelationId(v string) {
	o.LinkCorrelationId = v
}

func (o LinkOAuthCorrelationIdExchangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["link_correlation_id"] = o.LinkCorrelationId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkOAuthCorrelationIdExchangeRequest struct {
	value *LinkOAuthCorrelationIdExchangeRequest
	isSet bool
}

func (v NullableLinkOAuthCorrelationIdExchangeRequest) Get() *LinkOAuthCorrelationIdExchangeRequest {
	return v.value
}

func (v *NullableLinkOAuthCorrelationIdExchangeRequest) Set(val *LinkOAuthCorrelationIdExchangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkOAuthCorrelationIdExchangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkOAuthCorrelationIdExchangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkOAuthCorrelationIdExchangeRequest(val *LinkOAuthCorrelationIdExchangeRequest) *NullableLinkOAuthCorrelationIdExchangeRequest {
	return &NullableLinkOAuthCorrelationIdExchangeRequest{value: val, isSet: true}
}

func (v NullableLinkOAuthCorrelationIdExchangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkOAuthCorrelationIdExchangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


