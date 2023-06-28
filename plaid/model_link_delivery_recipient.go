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

// LinkDeliveryRecipient Metadata related to the recipient. If the information required to populate this field is not available, leave it blank.
type LinkDeliveryRecipient struct {
	// The list of communication methods to send the Hosted Link session URL to. If delivery is not required, leave this field blank.
	CommunicationMethods *[]LinkDeliveryCommunicationMethod `json:"communication_methods,omitempty"`
	// First name of the recipient. Will be used in the body of the email / text (if configured). If this information is not available, leave this field blank.
	FirstName *string `json:"first_name,omitempty"`
}

// NewLinkDeliveryRecipient instantiates a new LinkDeliveryRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDeliveryRecipient() *LinkDeliveryRecipient {
	this := LinkDeliveryRecipient{}
	return &this
}

// NewLinkDeliveryRecipientWithDefaults instantiates a new LinkDeliveryRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDeliveryRecipientWithDefaults() *LinkDeliveryRecipient {
	this := LinkDeliveryRecipient{}
	return &this
}

// GetCommunicationMethods returns the CommunicationMethods field value if set, zero value otherwise.
func (o *LinkDeliveryRecipient) GetCommunicationMethods() []LinkDeliveryCommunicationMethod {
	if o == nil || o.CommunicationMethods == nil {
		var ret []LinkDeliveryCommunicationMethod
		return ret
	}
	return *o.CommunicationMethods
}

// GetCommunicationMethodsOk returns a tuple with the CommunicationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryRecipient) GetCommunicationMethodsOk() (*[]LinkDeliveryCommunicationMethod, bool) {
	if o == nil || o.CommunicationMethods == nil {
		return nil, false
	}
	return o.CommunicationMethods, true
}

// HasCommunicationMethods returns a boolean if a field has been set.
func (o *LinkDeliveryRecipient) HasCommunicationMethods() bool {
	if o != nil && o.CommunicationMethods != nil {
		return true
	}

	return false
}

// SetCommunicationMethods gets a reference to the given []LinkDeliveryCommunicationMethod and assigns it to the CommunicationMethods field.
func (o *LinkDeliveryRecipient) SetCommunicationMethods(v []LinkDeliveryCommunicationMethod) {
	o.CommunicationMethods = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *LinkDeliveryRecipient) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryRecipient) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *LinkDeliveryRecipient) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *LinkDeliveryRecipient) SetFirstName(v string) {
	o.FirstName = &v
}

func (o LinkDeliveryRecipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommunicationMethods != nil {
		toSerialize["communication_methods"] = o.CommunicationMethods
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	return json.Marshal(toSerialize)
}

type NullableLinkDeliveryRecipient struct {
	value *LinkDeliveryRecipient
	isSet bool
}

func (v NullableLinkDeliveryRecipient) Get() *LinkDeliveryRecipient {
	return v.value
}

func (v *NullableLinkDeliveryRecipient) Set(val *LinkDeliveryRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryRecipient(val *LinkDeliveryRecipient) *NullableLinkDeliveryRecipient {
	return &NullableLinkDeliveryRecipient{value: val, isSet: true}
}

func (v NullableLinkDeliveryRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


