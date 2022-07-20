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

// AssetReportRelayCreateRequest AssetReportRelayCreateRequest defines the request schema for `/asset_report/relay/create`
type AssetReportRelayCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A token that can be provided to endpoints such as `/asset_report/get` or `/asset_report/pdf/get` to fetch or update an Asset Report.
	AssetReportToken string `json:"asset_report_token"`
	// The `secondary_client_id` is the client id of the third party with whom you would like to share the Asset Report.
	SecondaryClientId string `json:"secondary_client_id"`
	// URL to which Plaid will send webhooks when the Secondary Client successfully retrieves an Asset Report by calling `asset_report/relay/get`.
	Webhook NullableString `json:"webhook,omitempty"`
}

// NewAssetReportRelayCreateRequest instantiates a new AssetReportRelayCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportRelayCreateRequest(assetReportToken string, secondaryClientId string) *AssetReportRelayCreateRequest {
	this := AssetReportRelayCreateRequest{}
	this.AssetReportToken = assetReportToken
	this.SecondaryClientId = secondaryClientId
	return &this
}

// NewAssetReportRelayCreateRequestWithDefaults instantiates a new AssetReportRelayCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportRelayCreateRequestWithDefaults() *AssetReportRelayCreateRequest {
	this := AssetReportRelayCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportRelayCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRelayCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportRelayCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportRelayCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportRelayCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRelayCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportRelayCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportRelayCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAssetReportToken returns the AssetReportToken field value
func (o *AssetReportRelayCreateRequest) GetAssetReportToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportToken
}

// GetAssetReportTokenOk returns a tuple with the AssetReportToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportRelayCreateRequest) GetAssetReportTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportToken, true
}

// SetAssetReportToken sets field value
func (o *AssetReportRelayCreateRequest) SetAssetReportToken(v string) {
	o.AssetReportToken = v
}

// GetSecondaryClientId returns the SecondaryClientId field value
func (o *AssetReportRelayCreateRequest) GetSecondaryClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondaryClientId
}

// GetSecondaryClientIdOk returns a tuple with the SecondaryClientId field value
// and a boolean to check if the value has been set.
func (o *AssetReportRelayCreateRequest) GetSecondaryClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecondaryClientId, true
}

// SetSecondaryClientId sets field value
func (o *AssetReportRelayCreateRequest) SetSecondaryClientId(v string) {
	o.SecondaryClientId = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportRelayCreateRequest) GetWebhook() string {
	if o == nil || o.Webhook.Get() == nil {
		var ret string
		return ret
	}
	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportRelayCreateRequest) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// HasWebhook returns a boolean if a field has been set.
func (o *AssetReportRelayCreateRequest) HasWebhook() bool {
	if o != nil && o.Webhook.IsSet() {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NullableString and assigns it to the Webhook field.
func (o *AssetReportRelayCreateRequest) SetWebhook(v string) {
	o.Webhook.Set(&v)
}
// SetWebhookNil sets the value for Webhook to be an explicit nil
func (o *AssetReportRelayCreateRequest) SetWebhookNil() {
	o.Webhook.Set(nil)
}

// UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
func (o *AssetReportRelayCreateRequest) UnsetWebhook() {
	o.Webhook.Unset()
}

func (o AssetReportRelayCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["asset_report_token"] = o.AssetReportToken
	}
	if true {
		toSerialize["secondary_client_id"] = o.SecondaryClientId
	}
	if o.Webhook.IsSet() {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportRelayCreateRequest struct {
	value *AssetReportRelayCreateRequest
	isSet bool
}

func (v NullableAssetReportRelayCreateRequest) Get() *AssetReportRelayCreateRequest {
	return v.value
}

func (v *NullableAssetReportRelayCreateRequest) Set(val *AssetReportRelayCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportRelayCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportRelayCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportRelayCreateRequest(val *AssetReportRelayCreateRequest) *NullableAssetReportRelayCreateRequest {
	return &NullableAssetReportRelayCreateRequest{value: val, isSet: true}
}

func (v NullableAssetReportRelayCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportRelayCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


