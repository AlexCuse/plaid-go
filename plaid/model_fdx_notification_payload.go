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

// FDXNotificationPayload Custom key-value pairs payload for a notification
type FDXNotificationPayload struct {
	// ID for the origination entity related to the notification
	Id *string `json:"id,omitempty"`
	IdType *FDXNotificationPayloadIdType `json:"idType,omitempty"`
	CustomFields *[]FDXFiAttribute `json:"customFields,omitempty"`
}

// NewFDXNotificationPayload instantiates a new FDXNotificationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFDXNotificationPayload() *FDXNotificationPayload {
	this := FDXNotificationPayload{}
	return &this
}

// NewFDXNotificationPayloadWithDefaults instantiates a new FDXNotificationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFDXNotificationPayloadWithDefaults() *FDXNotificationPayload {
	this := FDXNotificationPayload{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FDXNotificationPayload) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXNotificationPayload) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FDXNotificationPayload) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FDXNotificationPayload) SetId(v string) {
	o.Id = &v
}

// GetIdType returns the IdType field value if set, zero value otherwise.
func (o *FDXNotificationPayload) GetIdType() FDXNotificationPayloadIdType {
	if o == nil || o.IdType == nil {
		var ret FDXNotificationPayloadIdType
		return ret
	}
	return *o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXNotificationPayload) GetIdTypeOk() (*FDXNotificationPayloadIdType, bool) {
	if o == nil || o.IdType == nil {
		return nil, false
	}
	return o.IdType, true
}

// HasIdType returns a boolean if a field has been set.
func (o *FDXNotificationPayload) HasIdType() bool {
	if o != nil && o.IdType != nil {
		return true
	}

	return false
}

// SetIdType gets a reference to the given FDXNotificationPayloadIdType and assigns it to the IdType field.
func (o *FDXNotificationPayload) SetIdType(v FDXNotificationPayloadIdType) {
	o.IdType = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *FDXNotificationPayload) GetCustomFields() []FDXFiAttribute {
	if o == nil || o.CustomFields == nil {
		var ret []FDXFiAttribute
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXNotificationPayload) GetCustomFieldsOk() (*[]FDXFiAttribute, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *FDXNotificationPayload) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []FDXFiAttribute and assigns it to the CustomFields field.
func (o *FDXNotificationPayload) SetCustomFields(v []FDXFiAttribute) {
	o.CustomFields = &v
}

func (o FDXNotificationPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdType != nil {
		toSerialize["idType"] = o.IdType
	}
	if o.CustomFields != nil {
		toSerialize["customFields"] = o.CustomFields
	}
	return json.Marshal(toSerialize)
}

type NullableFDXNotificationPayload struct {
	value *FDXNotificationPayload
	isSet bool
}

func (v NullableFDXNotificationPayload) Get() *FDXNotificationPayload {
	return v.value
}

func (v *NullableFDXNotificationPayload) Set(val *FDXNotificationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXNotificationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXNotificationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXNotificationPayload(val *FDXNotificationPayload) *NullableFDXNotificationPayload {
	return &NullableFDXNotificationPayload{value: val, isSet: true}
}

func (v NullableFDXNotificationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXNotificationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


