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

// CreditRelayCreateRequest CreditRelayCreateRequest defines the request schema for `/credit/relay/create`
type CreditRelayCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// List of report tokens, with at most one token of each report type. Currently only Asset Report token is supported.
	ReportTokens []ReportToken `json:"report_tokens"`
	// The `secondary_client_id` is the client id of the third party with whom you would like to share the Relay Token.
	SecondaryClientId string `json:"secondary_client_id"`
	// URL to which Plaid will send webhooks when the Secondary Client successfully retrieves an Asset Report by calling `/credit/relay/get`.
	Webhook NullableString `json:"webhook,omitempty"`
}

// NewCreditRelayCreateRequest instantiates a new CreditRelayCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditRelayCreateRequest(reportTokens []ReportToken, secondaryClientId string) *CreditRelayCreateRequest {
	this := CreditRelayCreateRequest{}
	this.ReportTokens = reportTokens
	this.SecondaryClientId = secondaryClientId
	return &this
}

// NewCreditRelayCreateRequestWithDefaults instantiates a new CreditRelayCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditRelayCreateRequestWithDefaults() *CreditRelayCreateRequest {
	this := CreditRelayCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditRelayCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditRelayCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditRelayCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditRelayCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditRelayCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditRelayCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditRelayCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditRelayCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetReportTokens returns the ReportTokens field value
func (o *CreditRelayCreateRequest) GetReportTokens() []ReportToken {
	if o == nil {
		var ret []ReportToken
		return ret
	}

	return o.ReportTokens
}

// GetReportTokensOk returns a tuple with the ReportTokens field value
// and a boolean to check if the value has been set.
func (o *CreditRelayCreateRequest) GetReportTokensOk() (*[]ReportToken, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReportTokens, true
}

// SetReportTokens sets field value
func (o *CreditRelayCreateRequest) SetReportTokens(v []ReportToken) {
	o.ReportTokens = v
}

// GetSecondaryClientId returns the SecondaryClientId field value
func (o *CreditRelayCreateRequest) GetSecondaryClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondaryClientId
}

// GetSecondaryClientIdOk returns a tuple with the SecondaryClientId field value
// and a boolean to check if the value has been set.
func (o *CreditRelayCreateRequest) GetSecondaryClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecondaryClientId, true
}

// SetSecondaryClientId sets field value
func (o *CreditRelayCreateRequest) SetSecondaryClientId(v string) {
	o.SecondaryClientId = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreditRelayCreateRequest) GetWebhook() string {
	if o == nil || o.Webhook.Get() == nil {
		var ret string
		return ret
	}
	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditRelayCreateRequest) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// HasWebhook returns a boolean if a field has been set.
func (o *CreditRelayCreateRequest) HasWebhook() bool {
	if o != nil && o.Webhook.IsSet() {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NullableString and assigns it to the Webhook field.
func (o *CreditRelayCreateRequest) SetWebhook(v string) {
	o.Webhook.Set(&v)
}
// SetWebhookNil sets the value for Webhook to be an explicit nil
func (o *CreditRelayCreateRequest) SetWebhookNil() {
	o.Webhook.Set(nil)
}

// UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
func (o *CreditRelayCreateRequest) UnsetWebhook() {
	o.Webhook.Unset()
}

func (o CreditRelayCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["report_tokens"] = o.ReportTokens
	}
	if true {
		toSerialize["secondary_client_id"] = o.SecondaryClientId
	}
	if o.Webhook.IsSet() {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreditRelayCreateRequest struct {
	value *CreditRelayCreateRequest
	isSet bool
}

func (v NullableCreditRelayCreateRequest) Get() *CreditRelayCreateRequest {
	return v.value
}

func (v *NullableCreditRelayCreateRequest) Set(val *CreditRelayCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditRelayCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditRelayCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditRelayCreateRequest(val *CreditRelayCreateRequest) *NullableCreditRelayCreateRequest {
	return &NullableCreditRelayCreateRequest{value: val, isSet: true}
}

func (v NullableCreditRelayCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditRelayCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


