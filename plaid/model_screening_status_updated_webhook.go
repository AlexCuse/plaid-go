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

// ScreeningStatusUpdatedWebhook Fired when an individual screening status has changed, which can occur manually via the dashboard or during ongoing monitoring.
type ScreeningStatusUpdatedWebhook struct {
	// `SCREENING`
	WebhookType string `json:"webhook_type"`
	// `STATUS_UPDATED`
	WebhookCode string `json:"webhook_code"`
	// The ID of the associated screening.
	ScreeningId string `json:"screening_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _ScreeningStatusUpdatedWebhook ScreeningStatusUpdatedWebhook

// NewScreeningStatusUpdatedWebhook instantiates a new ScreeningStatusUpdatedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreeningStatusUpdatedWebhook(webhookType string, webhookCode string, screeningId string, environment WebhookEnvironmentValues) *ScreeningStatusUpdatedWebhook {
	this := ScreeningStatusUpdatedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ScreeningId = screeningId
	this.Environment = environment
	return &this
}

// NewScreeningStatusUpdatedWebhookWithDefaults instantiates a new ScreeningStatusUpdatedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreeningStatusUpdatedWebhookWithDefaults() *ScreeningStatusUpdatedWebhook {
	this := ScreeningStatusUpdatedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *ScreeningStatusUpdatedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *ScreeningStatusUpdatedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *ScreeningStatusUpdatedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *ScreeningStatusUpdatedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *ScreeningStatusUpdatedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *ScreeningStatusUpdatedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetScreeningId returns the ScreeningId field value
func (o *ScreeningStatusUpdatedWebhook) GetScreeningId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScreeningId
}

// GetScreeningIdOk returns a tuple with the ScreeningId field value
// and a boolean to check if the value has been set.
func (o *ScreeningStatusUpdatedWebhook) GetScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScreeningId, true
}

// SetScreeningId sets field value
func (o *ScreeningStatusUpdatedWebhook) SetScreeningId(v string) {
	o.ScreeningId = v
}

// GetEnvironment returns the Environment field value
func (o *ScreeningStatusUpdatedWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ScreeningStatusUpdatedWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ScreeningStatusUpdatedWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o ScreeningStatusUpdatedWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["screening_id"] = o.ScreeningId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ScreeningStatusUpdatedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varScreeningStatusUpdatedWebhook := _ScreeningStatusUpdatedWebhook{}

	if err = json.Unmarshal(bytes, &varScreeningStatusUpdatedWebhook); err == nil {
		*o = ScreeningStatusUpdatedWebhook(varScreeningStatusUpdatedWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "screening_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScreeningStatusUpdatedWebhook struct {
	value *ScreeningStatusUpdatedWebhook
	isSet bool
}

func (v NullableScreeningStatusUpdatedWebhook) Get() *ScreeningStatusUpdatedWebhook {
	return v.value
}

func (v *NullableScreeningStatusUpdatedWebhook) Set(val *ScreeningStatusUpdatedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableScreeningStatusUpdatedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableScreeningStatusUpdatedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreeningStatusUpdatedWebhook(val *ScreeningStatusUpdatedWebhook) *NullableScreeningStatusUpdatedWebhook {
	return &NullableScreeningStatusUpdatedWebhook{value: val, isSet: true}
}

func (v NullableScreeningStatusUpdatedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreeningStatusUpdatedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


