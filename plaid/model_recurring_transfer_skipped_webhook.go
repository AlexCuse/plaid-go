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

// RecurringTransferSkippedWebhook Fired when Plaid is unable to originate a new ACH transaction of the recurring transfer on the planned date.
type RecurringTransferSkippedWebhook struct {
	// `TRANSFER`
	WebhookType string `json:"webhook_type"`
	// `RECURRING_TRANSFER_SKIPPED`
	WebhookCode string `json:"webhook_code"`
	// Plaid’s unique identifier for a recurring transfer.
	RecurringTransferId string `json:"recurring_transfer_id"`
	AuthorizationDecision TransferAuthorizationDecision `json:"authorization_decision"`
	AuthorizationDecisionRationaleCode *TransferAuthorizationDecisionRationaleCode `json:"authorization_decision_rationale_code,omitempty"`
	// The planned date on which Plaid is unable to originate a new ACH transaction of the recurring transfer. This will be of the form YYYY-MM-DD.
	SkippedOriginationDate string `json:"skipped_origination_date"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _RecurringTransferSkippedWebhook RecurringTransferSkippedWebhook

// NewRecurringTransferSkippedWebhook instantiates a new RecurringTransferSkippedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringTransferSkippedWebhook(webhookType string, webhookCode string, recurringTransferId string, authorizationDecision TransferAuthorizationDecision, skippedOriginationDate string, environment WebhookEnvironmentValues) *RecurringTransferSkippedWebhook {
	this := RecurringTransferSkippedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.RecurringTransferId = recurringTransferId
	this.AuthorizationDecision = authorizationDecision
	this.SkippedOriginationDate = skippedOriginationDate
	this.Environment = environment
	return &this
}

// NewRecurringTransferSkippedWebhookWithDefaults instantiates a new RecurringTransferSkippedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringTransferSkippedWebhookWithDefaults() *RecurringTransferSkippedWebhook {
	this := RecurringTransferSkippedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *RecurringTransferSkippedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *RecurringTransferSkippedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *RecurringTransferSkippedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *RecurringTransferSkippedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *RecurringTransferSkippedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *RecurringTransferSkippedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetRecurringTransferId returns the RecurringTransferId field value
func (o *RecurringTransferSkippedWebhook) GetRecurringTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecurringTransferId
}

// GetRecurringTransferIdOk returns a tuple with the RecurringTransferId field value
// and a boolean to check if the value has been set.
func (o *RecurringTransferSkippedWebhook) GetRecurringTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecurringTransferId, true
}

// SetRecurringTransferId sets field value
func (o *RecurringTransferSkippedWebhook) SetRecurringTransferId(v string) {
	o.RecurringTransferId = v
}

// GetAuthorizationDecision returns the AuthorizationDecision field value
func (o *RecurringTransferSkippedWebhook) GetAuthorizationDecision() TransferAuthorizationDecision {
	if o == nil {
		var ret TransferAuthorizationDecision
		return ret
	}

	return o.AuthorizationDecision
}

// GetAuthorizationDecisionOk returns a tuple with the AuthorizationDecision field value
// and a boolean to check if the value has been set.
func (o *RecurringTransferSkippedWebhook) GetAuthorizationDecisionOk() (*TransferAuthorizationDecision, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthorizationDecision, true
}

// SetAuthorizationDecision sets field value
func (o *RecurringTransferSkippedWebhook) SetAuthorizationDecision(v TransferAuthorizationDecision) {
	o.AuthorizationDecision = v
}

// GetAuthorizationDecisionRationaleCode returns the AuthorizationDecisionRationaleCode field value if set, zero value otherwise.
func (o *RecurringTransferSkippedWebhook) GetAuthorizationDecisionRationaleCode() TransferAuthorizationDecisionRationaleCode {
	if o == nil || o.AuthorizationDecisionRationaleCode == nil {
		var ret TransferAuthorizationDecisionRationaleCode
		return ret
	}
	return *o.AuthorizationDecisionRationaleCode
}

// GetAuthorizationDecisionRationaleCodeOk returns a tuple with the AuthorizationDecisionRationaleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringTransferSkippedWebhook) GetAuthorizationDecisionRationaleCodeOk() (*TransferAuthorizationDecisionRationaleCode, bool) {
	if o == nil || o.AuthorizationDecisionRationaleCode == nil {
		return nil, false
	}
	return o.AuthorizationDecisionRationaleCode, true
}

// HasAuthorizationDecisionRationaleCode returns a boolean if a field has been set.
func (o *RecurringTransferSkippedWebhook) HasAuthorizationDecisionRationaleCode() bool {
	if o != nil && o.AuthorizationDecisionRationaleCode != nil {
		return true
	}

	return false
}

// SetAuthorizationDecisionRationaleCode gets a reference to the given TransferAuthorizationDecisionRationaleCode and assigns it to the AuthorizationDecisionRationaleCode field.
func (o *RecurringTransferSkippedWebhook) SetAuthorizationDecisionRationaleCode(v TransferAuthorizationDecisionRationaleCode) {
	o.AuthorizationDecisionRationaleCode = &v
}

// GetSkippedOriginationDate returns the SkippedOriginationDate field value
func (o *RecurringTransferSkippedWebhook) GetSkippedOriginationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SkippedOriginationDate
}

// GetSkippedOriginationDateOk returns a tuple with the SkippedOriginationDate field value
// and a boolean to check if the value has been set.
func (o *RecurringTransferSkippedWebhook) GetSkippedOriginationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SkippedOriginationDate, true
}

// SetSkippedOriginationDate sets field value
func (o *RecurringTransferSkippedWebhook) SetSkippedOriginationDate(v string) {
	o.SkippedOriginationDate = v
}

// GetEnvironment returns the Environment field value
func (o *RecurringTransferSkippedWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *RecurringTransferSkippedWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *RecurringTransferSkippedWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o RecurringTransferSkippedWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["recurring_transfer_id"] = o.RecurringTransferId
	}
	if true {
		toSerialize["authorization_decision"] = o.AuthorizationDecision
	}
	if o.AuthorizationDecisionRationaleCode != nil {
		toSerialize["authorization_decision_rationale_code"] = o.AuthorizationDecisionRationaleCode
	}
	if true {
		toSerialize["skipped_origination_date"] = o.SkippedOriginationDate
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecurringTransferSkippedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varRecurringTransferSkippedWebhook := _RecurringTransferSkippedWebhook{}

	if err = json.Unmarshal(bytes, &varRecurringTransferSkippedWebhook); err == nil {
		*o = RecurringTransferSkippedWebhook(varRecurringTransferSkippedWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "recurring_transfer_id")
		delete(additionalProperties, "authorization_decision")
		delete(additionalProperties, "authorization_decision_rationale_code")
		delete(additionalProperties, "skipped_origination_date")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecurringTransferSkippedWebhook struct {
	value *RecurringTransferSkippedWebhook
	isSet bool
}

func (v NullableRecurringTransferSkippedWebhook) Get() *RecurringTransferSkippedWebhook {
	return v.value
}

func (v *NullableRecurringTransferSkippedWebhook) Set(val *RecurringTransferSkippedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringTransferSkippedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringTransferSkippedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringTransferSkippedWebhook(val *RecurringTransferSkippedWebhook) *NullableRecurringTransferSkippedWebhook {
	return &NullableRecurringTransferSkippedWebhook{value: val, isSet: true}
}

func (v NullableRecurringTransferSkippedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringTransferSkippedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


