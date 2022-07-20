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

// AssetReportRelayRefreshRequest AssetReportRelayRefreshRequest defines the request schema for `/asset_report/relay/refresh`
type AssetReportRelayRefreshRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	AssetRelayToken string `json:"asset_relay_token"`
	// The URL registered to receive webhooks when the Asset Report of a Relay Token has been refreshed.
	Webhook NullableString `json:"webhook,omitempty"`
}

// NewAssetReportRelayRefreshRequest instantiates a new AssetReportRelayRefreshRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportRelayRefreshRequest(assetRelayToken string) *AssetReportRelayRefreshRequest {
	this := AssetReportRelayRefreshRequest{}
	this.AssetRelayToken = assetRelayToken
	return &this
}

// NewAssetReportRelayRefreshRequestWithDefaults instantiates a new AssetReportRelayRefreshRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportRelayRefreshRequestWithDefaults() *AssetReportRelayRefreshRequest {
	this := AssetReportRelayRefreshRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportRelayRefreshRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRelayRefreshRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportRelayRefreshRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportRelayRefreshRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportRelayRefreshRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRelayRefreshRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportRelayRefreshRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportRelayRefreshRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAssetRelayToken returns the AssetRelayToken field value
func (o *AssetReportRelayRefreshRequest) GetAssetRelayToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetRelayToken
}

// GetAssetRelayTokenOk returns a tuple with the AssetRelayToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportRelayRefreshRequest) GetAssetRelayTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetRelayToken, true
}

// SetAssetRelayToken sets field value
func (o *AssetReportRelayRefreshRequest) SetAssetRelayToken(v string) {
	o.AssetRelayToken = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportRelayRefreshRequest) GetWebhook() string {
	if o == nil || o.Webhook.Get() == nil {
		var ret string
		return ret
	}
	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportRelayRefreshRequest) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// HasWebhook returns a boolean if a field has been set.
func (o *AssetReportRelayRefreshRequest) HasWebhook() bool {
	if o != nil && o.Webhook.IsSet() {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NullableString and assigns it to the Webhook field.
func (o *AssetReportRelayRefreshRequest) SetWebhook(v string) {
	o.Webhook.Set(&v)
}
// SetWebhookNil sets the value for Webhook to be an explicit nil
func (o *AssetReportRelayRefreshRequest) SetWebhookNil() {
	o.Webhook.Set(nil)
}

// UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
func (o *AssetReportRelayRefreshRequest) UnsetWebhook() {
	o.Webhook.Unset()
}

func (o AssetReportRelayRefreshRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["asset_relay_token"] = o.AssetRelayToken
	}
	if o.Webhook.IsSet() {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportRelayRefreshRequest struct {
	value *AssetReportRelayRefreshRequest
	isSet bool
}

func (v NullableAssetReportRelayRefreshRequest) Get() *AssetReportRelayRefreshRequest {
	return v.value
}

func (v *NullableAssetReportRelayRefreshRequest) Set(val *AssetReportRelayRefreshRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportRelayRefreshRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportRelayRefreshRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportRelayRefreshRequest(val *AssetReportRelayRefreshRequest) *NullableAssetReportRelayRefreshRequest {
	return &NullableAssetReportRelayRefreshRequest{value: val, isSet: true}
}

func (v NullableAssetReportRelayRefreshRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportRelayRefreshRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


