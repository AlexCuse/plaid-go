/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.413.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IncomeVerificationRiskSignalsStatusWebhook Fired when risk signals have been processed for documents uploaded via Document Income. It will typically take a minute or two for this webhook to fire after the end user has uploaded their documents in the Document Income flow. Once this webhook has fired, `/credit/payroll_income/risk_signals/get` may then be called to determine whether the documents were successfully processed and to retrieve risk data.
type IncomeVerificationRiskSignalsStatusWebhook struct {
	// `\"INCOME\"`
	WebhookType string `json:"webhook_type"`
	// `INCOME_VERIFICATION_RISK_SIGNALS`
	WebhookCode string `json:"webhook_code"`
	// The Item ID associated with the verification.
	ItemId string `json:"item_id"`
	// The Plaid `user_id` of the User associated with this webhook, warning, or error.
	UserId *string `json:"user_id,omitempty"`
	// `RISK_SIGNALS_PROCESSING_COMPLETE`: The income verification fraud detection processing has completed. If the user uploaded multiple documents, this webhook will fire when all documents have finished processing. Call the `/credit/payroll_income/risk_signals/get` endpoint to get all risk signal data.
	RiskSignalsStatus *string `json:"risk_signals_status,omitempty"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _IncomeVerificationRiskSignalsStatusWebhook IncomeVerificationRiskSignalsStatusWebhook

// NewIncomeVerificationRiskSignalsStatusWebhook instantiates a new IncomeVerificationRiskSignalsStatusWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationRiskSignalsStatusWebhook(webhookType string, webhookCode string, itemId string, environment WebhookEnvironmentValues) *IncomeVerificationRiskSignalsStatusWebhook {
	this := IncomeVerificationRiskSignalsStatusWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.Environment = environment
	return &this
}

// NewIncomeVerificationRiskSignalsStatusWebhookWithDefaults instantiates a new IncomeVerificationRiskSignalsStatusWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationRiskSignalsStatusWebhookWithDefaults() *IncomeVerificationRiskSignalsStatusWebhook {
	this := IncomeVerificationRiskSignalsStatusWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *IncomeVerificationRiskSignalsStatusWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *IncomeVerificationRiskSignalsStatusWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *IncomeVerificationRiskSignalsStatusWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IncomeVerificationRiskSignalsStatusWebhook) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IncomeVerificationRiskSignalsStatusWebhook) SetUserId(v string) {
	o.UserId = &v
}

// GetRiskSignalsStatus returns the RiskSignalsStatus field value if set, zero value otherwise.
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetRiskSignalsStatus() string {
	if o == nil || o.RiskSignalsStatus == nil {
		var ret string
		return ret
	}
	return *o.RiskSignalsStatus
}

// GetRiskSignalsStatusOk returns a tuple with the RiskSignalsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetRiskSignalsStatusOk() (*string, bool) {
	if o == nil || o.RiskSignalsStatus == nil {
		return nil, false
	}
	return o.RiskSignalsStatus, true
}

// HasRiskSignalsStatus returns a boolean if a field has been set.
func (o *IncomeVerificationRiskSignalsStatusWebhook) HasRiskSignalsStatus() bool {
	if o != nil && o.RiskSignalsStatus != nil {
		return true
	}

	return false
}

// SetRiskSignalsStatus gets a reference to the given string and assigns it to the RiskSignalsStatus field.
func (o *IncomeVerificationRiskSignalsStatusWebhook) SetRiskSignalsStatus(v string) {
	o.RiskSignalsStatus = &v
}

// GetEnvironment returns the Environment field value
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationRiskSignalsStatusWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *IncomeVerificationRiskSignalsStatusWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o IncomeVerificationRiskSignalsStatusWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.RiskSignalsStatus != nil {
		toSerialize["risk_signals_status"] = o.RiskSignalsStatus
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeVerificationRiskSignalsStatusWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeVerificationRiskSignalsStatusWebhook := _IncomeVerificationRiskSignalsStatusWebhook{}

	if err = json.Unmarshal(bytes, &varIncomeVerificationRiskSignalsStatusWebhook); err == nil {
		*o = IncomeVerificationRiskSignalsStatusWebhook(varIncomeVerificationRiskSignalsStatusWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "risk_signals_status")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeVerificationRiskSignalsStatusWebhook struct {
	value *IncomeVerificationRiskSignalsStatusWebhook
	isSet bool
}

func (v NullableIncomeVerificationRiskSignalsStatusWebhook) Get() *IncomeVerificationRiskSignalsStatusWebhook {
	return v.value
}

func (v *NullableIncomeVerificationRiskSignalsStatusWebhook) Set(val *IncomeVerificationRiskSignalsStatusWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationRiskSignalsStatusWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationRiskSignalsStatusWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationRiskSignalsStatusWebhook(val *IncomeVerificationRiskSignalsStatusWebhook) *NullableIncomeVerificationRiskSignalsStatusWebhook {
	return &NullableIncomeVerificationRiskSignalsStatusWebhook{value: val, isSet: true}
}

func (v NullableIncomeVerificationRiskSignalsStatusWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationRiskSignalsStatusWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

