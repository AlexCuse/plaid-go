/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.131.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// IncidentUpdate An update on the health incident
type IncidentUpdate struct {
	// The content of the update.
	Description *string `json:"description,omitempty"`
	// The status of the incident.
	Status *string `json:"status,omitempty"`
	// The date when the update was published, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `\"2020-10-30T15:26:48Z\"`.
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IncidentUpdate IncidentUpdate

// NewIncidentUpdate instantiates a new IncidentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentUpdate() *IncidentUpdate {
	this := IncidentUpdate{}
	return &this
}

// NewIncidentUpdateWithDefaults instantiates a new IncidentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentUpdateWithDefaults() *IncidentUpdate {
	this := IncidentUpdate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IncidentUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IncidentUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IncidentUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IncidentUpdate) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IncidentUpdate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IncidentUpdate) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *IncidentUpdate) GetUpdatedDate() time.Time {
	if o == nil || o.UpdatedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetUpdatedDateOk() (*time.Time, bool) {
	if o == nil || o.UpdatedDate == nil {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *IncidentUpdate) HasUpdatedDate() bool {
	if o != nil && o.UpdatedDate != nil {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given time.Time and assigns it to the UpdatedDate field.
func (o *IncidentUpdate) SetUpdatedDate(v time.Time) {
	o.UpdatedDate = &v
}

func (o IncidentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedDate != nil {
		toSerialize["updated_date"] = o.UpdatedDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncidentUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varIncidentUpdate := _IncidentUpdate{}

	if err = json.Unmarshal(bytes, &varIncidentUpdate); err == nil {
		*o = IncidentUpdate(varIncidentUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "status")
		delete(additionalProperties, "updated_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncidentUpdate struct {
	value *IncidentUpdate
	isSet bool
}

func (v NullableIncidentUpdate) Get() *IncidentUpdate {
	return v.value
}

func (v *NullableIncidentUpdate) Set(val *IncidentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentUpdate(val *IncidentUpdate) *NullableIncidentUpdate {
	return &NullableIncidentUpdate{value: val, isSet: true}
}

func (v NullableIncidentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


