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
	"time"
)

// DashboardUser Account information associated with a team member with access to the Plaid dashboard.
type DashboardUser struct {
	// ID of the associated user.
	Id string `json:"id"`
	// An ISO8601 formatted timestamp.
	CreatedAt time.Time `json:"created_at"`
	// A valid email address.
	EmailAddress string `json:"email_address"`
	Status DashboardUserStatus `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _DashboardUser DashboardUser

// NewDashboardUser instantiates a new DashboardUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardUser(id string, createdAt time.Time, emailAddress string, status DashboardUserStatus) *DashboardUser {
	this := DashboardUser{}
	this.Id = id
	this.CreatedAt = createdAt
	this.EmailAddress = emailAddress
	this.Status = status
	return &this
}

// NewDashboardUserWithDefaults instantiates a new DashboardUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardUserWithDefaults() *DashboardUser {
	this := DashboardUser{}
	return &this
}

// GetId returns the Id field value
func (o *DashboardUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DashboardUser) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DashboardUser) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DashboardUser) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DashboardUser) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DashboardUser) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *DashboardUser) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *DashboardUser) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *DashboardUser) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetStatus returns the Status field value
func (o *DashboardUser) GetStatus() DashboardUserStatus {
	if o == nil {
		var ret DashboardUserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DashboardUser) GetStatusOk() (*DashboardUserStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DashboardUser) SetStatus(v DashboardUserStatus) {
	o.Status = v
}

func (o DashboardUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["email_address"] = o.EmailAddress
	}
	if true {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DashboardUser) UnmarshalJSON(bytes []byte) (err error) {
	varDashboardUser := _DashboardUser{}

	if err = json.Unmarshal(bytes, &varDashboardUser); err == nil {
		*o = DashboardUser(varDashboardUser)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "email_address")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDashboardUser struct {
	value *DashboardUser
	isSet bool
}

func (v NullableDashboardUser) Get() *DashboardUser {
	return v.value
}

func (v *NullableDashboardUser) Set(val *DashboardUser) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardUser) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardUser(val *DashboardUser) *NullableDashboardUser {
	return &NullableDashboardUser{value: val, isSet: true}
}

func (v NullableDashboardUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


