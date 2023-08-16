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

// LinkSessionFinishedWebhook Contains the state of a completed link session, along with the public token if available.
type LinkSessionFinishedWebhook struct {
	// `LINK`
	WebhookType string `json:"webhook_type"`
	// `SESSION_FINISHED`
	WebhookCode string `json:"webhook_code"`
	// The final status of the link session. Will always be \"SUCCESS\" or \"EXITED\".
	Status string `json:"status"`
	// The identifier for the link session.
	LinkSessionId string `json:"link_session_id"`
	// The link token used to create the link session.
	LinkToken string `json:"link_token"`
	// The public token generated by the link session.
	PublicToken *string `json:"public_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkSessionFinishedWebhook LinkSessionFinishedWebhook

// NewLinkSessionFinishedWebhook instantiates a new LinkSessionFinishedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkSessionFinishedWebhook(webhookType string, webhookCode string, status string, linkSessionId string, linkToken string) *LinkSessionFinishedWebhook {
	this := LinkSessionFinishedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.Status = status
	this.LinkSessionId = linkSessionId
	this.LinkToken = linkToken
	return &this
}

// NewLinkSessionFinishedWebhookWithDefaults instantiates a new LinkSessionFinishedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkSessionFinishedWebhookWithDefaults() *LinkSessionFinishedWebhook {
	this := LinkSessionFinishedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *LinkSessionFinishedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *LinkSessionFinishedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *LinkSessionFinishedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *LinkSessionFinishedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *LinkSessionFinishedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *LinkSessionFinishedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetStatus returns the Status field value
func (o *LinkSessionFinishedWebhook) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LinkSessionFinishedWebhook) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LinkSessionFinishedWebhook) SetStatus(v string) {
	o.Status = v
}

// GetLinkSessionId returns the LinkSessionId field value
func (o *LinkSessionFinishedWebhook) GetLinkSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkSessionId
}

// GetLinkSessionIdOk returns a tuple with the LinkSessionId field value
// and a boolean to check if the value has been set.
func (o *LinkSessionFinishedWebhook) GetLinkSessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkSessionId, true
}

// SetLinkSessionId sets field value
func (o *LinkSessionFinishedWebhook) SetLinkSessionId(v string) {
	o.LinkSessionId = v
}

// GetLinkToken returns the LinkToken field value
func (o *LinkSessionFinishedWebhook) GetLinkToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkToken
}

// GetLinkTokenOk returns a tuple with the LinkToken field value
// and a boolean to check if the value has been set.
func (o *LinkSessionFinishedWebhook) GetLinkTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkToken, true
}

// SetLinkToken sets field value
func (o *LinkSessionFinishedWebhook) SetLinkToken(v string) {
	o.LinkToken = v
}

// GetPublicToken returns the PublicToken field value if set, zero value otherwise.
func (o *LinkSessionFinishedWebhook) GetPublicToken() string {
	if o == nil || o.PublicToken == nil {
		var ret string
		return ret
	}
	return *o.PublicToken
}

// GetPublicTokenOk returns a tuple with the PublicToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkSessionFinishedWebhook) GetPublicTokenOk() (*string, bool) {
	if o == nil || o.PublicToken == nil {
		return nil, false
	}
	return o.PublicToken, true
}

// HasPublicToken returns a boolean if a field has been set.
func (o *LinkSessionFinishedWebhook) HasPublicToken() bool {
	if o != nil && o.PublicToken != nil {
		return true
	}

	return false
}

// SetPublicToken gets a reference to the given string and assigns it to the PublicToken field.
func (o *LinkSessionFinishedWebhook) SetPublicToken(v string) {
	o.PublicToken = &v
}

func (o LinkSessionFinishedWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["link_session_id"] = o.LinkSessionId
	}
	if true {
		toSerialize["link_token"] = o.LinkToken
	}
	if o.PublicToken != nil {
		toSerialize["public_token"] = o.PublicToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkSessionFinishedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varLinkSessionFinishedWebhook := _LinkSessionFinishedWebhook{}

	if err = json.Unmarshal(bytes, &varLinkSessionFinishedWebhook); err == nil {
		*o = LinkSessionFinishedWebhook(varLinkSessionFinishedWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "status")
		delete(additionalProperties, "link_session_id")
		delete(additionalProperties, "link_token")
		delete(additionalProperties, "public_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkSessionFinishedWebhook struct {
	value *LinkSessionFinishedWebhook
	isSet bool
}

func (v NullableLinkSessionFinishedWebhook) Get() *LinkSessionFinishedWebhook {
	return v.value
}

func (v *NullableLinkSessionFinishedWebhook) Set(val *LinkSessionFinishedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkSessionFinishedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkSessionFinishedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkSessionFinishedWebhook(val *LinkSessionFinishedWebhook) *NullableLinkSessionFinishedWebhook {
	return &NullableLinkSessionFinishedWebhook{value: val, isSet: true}
}

func (v NullableLinkSessionFinishedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkSessionFinishedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


