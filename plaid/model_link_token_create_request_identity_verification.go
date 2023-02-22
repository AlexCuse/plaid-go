/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LinkTokenCreateRequestIdentityVerification Specifies option for initializing Link for use with the Identity Verification product.
type LinkTokenCreateRequestIdentityVerification struct {
	// ID of the associated Identity Verification template.
	TemplateId string `json:"template_id"`
	Consent *bool `json:"consent,omitempty"`
	// A flag specifying whether the end user has already agreed to a privacy policy specifying that their data will be shared with Plaid for verification purposes.  If `gave_consent` is set to `true`, the `accept_tos` step will be marked as `skipped` and the end user's session will start at the next step requirement.
	GaveConsent *bool `json:"gave_consent,omitempty"`
}

// NewLinkTokenCreateRequestIdentityVerification instantiates a new LinkTokenCreateRequestIdentityVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestIdentityVerification(templateId string) *LinkTokenCreateRequestIdentityVerification {
	this := LinkTokenCreateRequestIdentityVerification{}
	this.TemplateId = templateId
	var gaveConsent bool = false
	this.GaveConsent = &gaveConsent
	return &this
}

// NewLinkTokenCreateRequestIdentityVerificationWithDefaults instantiates a new LinkTokenCreateRequestIdentityVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestIdentityVerificationWithDefaults() *LinkTokenCreateRequestIdentityVerification {
	this := LinkTokenCreateRequestIdentityVerification{}
	var gaveConsent bool = false
	this.GaveConsent = &gaveConsent
	return &this
}

// GetTemplateId returns the TemplateId field value
func (o *LinkTokenCreateRequestIdentityVerification) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIdentityVerification) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *LinkTokenCreateRequestIdentityVerification) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIdentityVerification) GetConsent() bool {
	if o == nil || o.Consent == nil {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIdentityVerification) GetConsentOk() (*bool, bool) {
	if o == nil || o.Consent == nil {
		return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIdentityVerification) HasConsent() bool {
	if o != nil && o.Consent != nil {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *LinkTokenCreateRequestIdentityVerification) SetConsent(v bool) {
	o.Consent = &v
}

// GetGaveConsent returns the GaveConsent field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIdentityVerification) GetGaveConsent() bool {
	if o == nil || o.GaveConsent == nil {
		var ret bool
		return ret
	}
	return *o.GaveConsent
}

// GetGaveConsentOk returns a tuple with the GaveConsent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIdentityVerification) GetGaveConsentOk() (*bool, bool) {
	if o == nil || o.GaveConsent == nil {
		return nil, false
	}
	return o.GaveConsent, true
}

// HasGaveConsent returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIdentityVerification) HasGaveConsent() bool {
	if o != nil && o.GaveConsent != nil {
		return true
	}

	return false
}

// SetGaveConsent gets a reference to the given bool and assigns it to the GaveConsent field.
func (o *LinkTokenCreateRequestIdentityVerification) SetGaveConsent(v bool) {
	o.GaveConsent = &v
}

func (o LinkTokenCreateRequestIdentityVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["template_id"] = o.TemplateId
	}
	if o.Consent != nil {
		toSerialize["consent"] = o.Consent
	}
	if o.GaveConsent != nil {
		toSerialize["gave_consent"] = o.GaveConsent
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestIdentityVerification struct {
	value *LinkTokenCreateRequestIdentityVerification
	isSet bool
}

func (v NullableLinkTokenCreateRequestIdentityVerification) Get() *LinkTokenCreateRequestIdentityVerification {
	return v.value
}

func (v *NullableLinkTokenCreateRequestIdentityVerification) Set(val *LinkTokenCreateRequestIdentityVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestIdentityVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestIdentityVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestIdentityVerification(val *LinkTokenCreateRequestIdentityVerification) *NullableLinkTokenCreateRequestIdentityVerification {
	return &NullableLinkTokenCreateRequestIdentityVerification{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestIdentityVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestIdentityVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


