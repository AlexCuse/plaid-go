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

// ItemActivityListResponse Describes a historical log of user consent events.
type ItemActivityListResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	// A list of activities.
	Activities []Activity `json:"activities"`
	// An array of objects containing timestamps for the last time each data type was accessed per application.
	LastDataAccessTimes []LastDataAccessTimes `json:"last_data_access_times"`
	// Cursor used for pagination.
	Cursor *string `json:"cursor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ItemActivityListResponse ItemActivityListResponse

// NewItemActivityListResponse instantiates a new ItemActivityListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemActivityListResponse(requestId string, activities []Activity, lastDataAccessTimes []LastDataAccessTimes) *ItemActivityListResponse {
	this := ItemActivityListResponse{}
	this.RequestId = requestId
	this.Activities = activities
	this.LastDataAccessTimes = lastDataAccessTimes
	return &this
}

// NewItemActivityListResponseWithDefaults instantiates a new ItemActivityListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemActivityListResponseWithDefaults() *ItemActivityListResponse {
	this := ItemActivityListResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ItemActivityListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ItemActivityListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ItemActivityListResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetActivities returns the Activities field value
func (o *ItemActivityListResponse) GetActivities() []Activity {
	if o == nil {
		var ret []Activity
		return ret
	}

	return o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field value
// and a boolean to check if the value has been set.
func (o *ItemActivityListResponse) GetActivitiesOk() (*[]Activity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Activities, true
}

// SetActivities sets field value
func (o *ItemActivityListResponse) SetActivities(v []Activity) {
	o.Activities = v
}

// GetLastDataAccessTimes returns the LastDataAccessTimes field value
func (o *ItemActivityListResponse) GetLastDataAccessTimes() []LastDataAccessTimes {
	if o == nil {
		var ret []LastDataAccessTimes
		return ret
	}

	return o.LastDataAccessTimes
}

// GetLastDataAccessTimesOk returns a tuple with the LastDataAccessTimes field value
// and a boolean to check if the value has been set.
func (o *ItemActivityListResponse) GetLastDataAccessTimesOk() (*[]LastDataAccessTimes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastDataAccessTimes, true
}

// SetLastDataAccessTimes sets field value
func (o *ItemActivityListResponse) SetLastDataAccessTimes(v []LastDataAccessTimes) {
	o.LastDataAccessTimes = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ItemActivityListResponse) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityListResponse) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ItemActivityListResponse) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ItemActivityListResponse) SetCursor(v string) {
	o.Cursor = &v
}

func (o ItemActivityListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["activities"] = o.Activities
	}
	if true {
		toSerialize["last_data_access_times"] = o.LastDataAccessTimes
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemActivityListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varItemActivityListResponse := _ItemActivityListResponse{}

	if err = json.Unmarshal(bytes, &varItemActivityListResponse); err == nil {
		*o = ItemActivityListResponse(varItemActivityListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "activities")
		delete(additionalProperties, "last_data_access_times")
		delete(additionalProperties, "cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemActivityListResponse struct {
	value *ItemActivityListResponse
	isSet bool
}

func (v NullableItemActivityListResponse) Get() *ItemActivityListResponse {
	return v.value
}

func (v *NullableItemActivityListResponse) Set(val *ItemActivityListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableItemActivityListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableItemActivityListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemActivityListResponse(val *ItemActivityListResponse) *NullableItemActivityListResponse {
	return &NullableItemActivityListResponse{value: val, isSet: true}
}

func (v NullableItemActivityListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemActivityListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


