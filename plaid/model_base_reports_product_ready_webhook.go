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

// BaseReportsProductReadyWebhook Fired when the Base Report has been generated and `/cra/base_report/get` is ready to be called.  If you attempt to retrieve a Base Report before this webhook has fired, you’ll receive a response with the HTTP status code 400 and a Plaid error code of `PRODUCT_NOT_READY`.
type BaseReportsProductReadyWebhook struct {
	// `BASE_REPORT`
	WebhookType string `json:"webhook_type"`
	// `PRODUCT_READY`
	WebhookCode string `json:"webhook_code"`
	// The `user_id` corresponding to the User ID the webhook has fired for.
	UserId *string `json:"user_id,omitempty"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _BaseReportsProductReadyWebhook BaseReportsProductReadyWebhook

// NewBaseReportsProductReadyWebhook instantiates a new BaseReportsProductReadyWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseReportsProductReadyWebhook(webhookType string, webhookCode string, environment WebhookEnvironmentValues) *BaseReportsProductReadyWebhook {
	this := BaseReportsProductReadyWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.Environment = environment
	return &this
}

// NewBaseReportsProductReadyWebhookWithDefaults instantiates a new BaseReportsProductReadyWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseReportsProductReadyWebhookWithDefaults() *BaseReportsProductReadyWebhook {
	this := BaseReportsProductReadyWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *BaseReportsProductReadyWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *BaseReportsProductReadyWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *BaseReportsProductReadyWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *BaseReportsProductReadyWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *BaseReportsProductReadyWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *BaseReportsProductReadyWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BaseReportsProductReadyWebhook) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportsProductReadyWebhook) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BaseReportsProductReadyWebhook) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BaseReportsProductReadyWebhook) SetUserId(v string) {
	o.UserId = &v
}

// GetEnvironment returns the Environment field value
func (o *BaseReportsProductReadyWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *BaseReportsProductReadyWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *BaseReportsProductReadyWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o BaseReportsProductReadyWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseReportsProductReadyWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varBaseReportsProductReadyWebhook := _BaseReportsProductReadyWebhook{}

	if err = json.Unmarshal(bytes, &varBaseReportsProductReadyWebhook); err == nil {
		*o = BaseReportsProductReadyWebhook(varBaseReportsProductReadyWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseReportsProductReadyWebhook struct {
	value *BaseReportsProductReadyWebhook
	isSet bool
}

func (v NullableBaseReportsProductReadyWebhook) Get() *BaseReportsProductReadyWebhook {
	return v.value
}

func (v *NullableBaseReportsProductReadyWebhook) Set(val *BaseReportsProductReadyWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseReportsProductReadyWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseReportsProductReadyWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseReportsProductReadyWebhook(val *BaseReportsProductReadyWebhook) *NullableBaseReportsProductReadyWebhook {
	return &NullableBaseReportsProductReadyWebhook{value: val, isSet: true}
}

func (v NullableBaseReportsProductReadyWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseReportsProductReadyWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


