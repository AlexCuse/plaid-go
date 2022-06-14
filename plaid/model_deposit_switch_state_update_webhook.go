/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// DepositSwitchStateUpdateWebhook Fired when the status of a deposit switch request has changed.
type DepositSwitchStateUpdateWebhook struct {
	// `\"DEPOSIT_SWITCH\"`
	WebhookType *string `json:"webhook_type,omitempty"`
	// `\"SWITCH_STATE_UPDATE\"`
	WebhookCode *string `json:"webhook_code,omitempty"`
	//  The state, or status, of the deposit switch.  `initialized`: The deposit switch has been initialized with the user entering the information required to submit the deposit switch request.  `processing`: The deposit switch request has been submitted and is being processed.  `completed`: The user's employer has fulfilled and completed the deposit switch request.  `error`: There was an error processing the deposit switch request.  For more information, see the [Deposit Switch API reference](/docs/deposit-switch/reference#deposit_switchget).
	State *string `json:"state,omitempty"`
	// The ID of the deposit switch.
	DepositSwitchId *string `json:"deposit_switch_id,omitempty"`
}

// NewDepositSwitchStateUpdateWebhook instantiates a new DepositSwitchStateUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchStateUpdateWebhook() *DepositSwitchStateUpdateWebhook {
	this := DepositSwitchStateUpdateWebhook{}
	return &this
}

// NewDepositSwitchStateUpdateWebhookWithDefaults instantiates a new DepositSwitchStateUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchStateUpdateWebhookWithDefaults() *DepositSwitchStateUpdateWebhook {
	this := DepositSwitchStateUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value if set, zero value otherwise.
func (o *DepositSwitchStateUpdateWebhook) GetWebhookType() string {
	if o == nil || o.WebhookType == nil {
		var ret string
		return ret
	}
	return *o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchStateUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil || o.WebhookType == nil {
		return nil, false
	}
	return o.WebhookType, true
}

// HasWebhookType returns a boolean if a field has been set.
func (o *DepositSwitchStateUpdateWebhook) HasWebhookType() bool {
	if o != nil && o.WebhookType != nil {
		return true
	}

	return false
}

// SetWebhookType gets a reference to the given string and assigns it to the WebhookType field.
func (o *DepositSwitchStateUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = &v
}

// GetWebhookCode returns the WebhookCode field value if set, zero value otherwise.
func (o *DepositSwitchStateUpdateWebhook) GetWebhookCode() string {
	if o == nil || o.WebhookCode == nil {
		var ret string
		return ret
	}
	return *o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchStateUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil || o.WebhookCode == nil {
		return nil, false
	}
	return o.WebhookCode, true
}

// HasWebhookCode returns a boolean if a field has been set.
func (o *DepositSwitchStateUpdateWebhook) HasWebhookCode() bool {
	if o != nil && o.WebhookCode != nil {
		return true
	}

	return false
}

// SetWebhookCode gets a reference to the given string and assigns it to the WebhookCode field.
func (o *DepositSwitchStateUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DepositSwitchStateUpdateWebhook) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchStateUpdateWebhook) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DepositSwitchStateUpdateWebhook) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DepositSwitchStateUpdateWebhook) SetState(v string) {
	o.State = &v
}

// GetDepositSwitchId returns the DepositSwitchId field value if set, zero value otherwise.
func (o *DepositSwitchStateUpdateWebhook) GetDepositSwitchId() string {
	if o == nil || o.DepositSwitchId == nil {
		var ret string
		return ret
	}
	return *o.DepositSwitchId
}

// GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchStateUpdateWebhook) GetDepositSwitchIdOk() (*string, bool) {
	if o == nil || o.DepositSwitchId == nil {
		return nil, false
	}
	return o.DepositSwitchId, true
}

// HasDepositSwitchId returns a boolean if a field has been set.
func (o *DepositSwitchStateUpdateWebhook) HasDepositSwitchId() bool {
	if o != nil && o.DepositSwitchId != nil {
		return true
	}

	return false
}

// SetDepositSwitchId gets a reference to the given string and assigns it to the DepositSwitchId field.
func (o *DepositSwitchStateUpdateWebhook) SetDepositSwitchId(v string) {
	o.DepositSwitchId = &v
}

func (o DepositSwitchStateUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WebhookType != nil {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if o.WebhookCode != nil {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.DepositSwitchId != nil {
		toSerialize["deposit_switch_id"] = o.DepositSwitchId
	}
	return json.Marshal(toSerialize)
}

type NullableDepositSwitchStateUpdateWebhook struct {
	value *DepositSwitchStateUpdateWebhook
	isSet bool
}

func (v NullableDepositSwitchStateUpdateWebhook) Get() *DepositSwitchStateUpdateWebhook {
	return v.value
}

func (v *NullableDepositSwitchStateUpdateWebhook) Set(val *DepositSwitchStateUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchStateUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchStateUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchStateUpdateWebhook(val *DepositSwitchStateUpdateWebhook) *NullableDepositSwitchStateUpdateWebhook {
	return &NullableDepositSwitchStateUpdateWebhook{value: val, isSet: true}
}

func (v NullableDepositSwitchStateUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchStateUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


