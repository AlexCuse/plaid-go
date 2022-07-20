/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ListWatchlistScreeningEntityHistoryRequest Request input for listing changes to entity watchlist screenings
type ListWatchlistScreeningEntityHistoryRequest struct {
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// ID of the associated entity screening.
	EntityWatchlistScreeningId string `json:"entity_watchlist_screening_id"`
	// An identifier that determines which page of results you receive.
	Cursor NullableString `json:"cursor,omitempty"`
}

// NewListWatchlistScreeningEntityHistoryRequest instantiates a new ListWatchlistScreeningEntityHistoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWatchlistScreeningEntityHistoryRequest(entityWatchlistScreeningId string) *ListWatchlistScreeningEntityHistoryRequest {
	this := ListWatchlistScreeningEntityHistoryRequest{}
	this.EntityWatchlistScreeningId = entityWatchlistScreeningId
	return &this
}

// NewListWatchlistScreeningEntityHistoryRequestWithDefaults instantiates a new ListWatchlistScreeningEntityHistoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWatchlistScreeningEntityHistoryRequestWithDefaults() *ListWatchlistScreeningEntityHistoryRequest {
	this := ListWatchlistScreeningEntityHistoryRequest{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ListWatchlistScreeningEntityHistoryRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWatchlistScreeningEntityHistoryRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ListWatchlistScreeningEntityHistoryRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ListWatchlistScreeningEntityHistoryRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ListWatchlistScreeningEntityHistoryRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWatchlistScreeningEntityHistoryRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ListWatchlistScreeningEntityHistoryRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ListWatchlistScreeningEntityHistoryRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetEntityWatchlistScreeningId returns the EntityWatchlistScreeningId field value
func (o *ListWatchlistScreeningEntityHistoryRequest) GetEntityWatchlistScreeningId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityWatchlistScreeningId
}

// GetEntityWatchlistScreeningIdOk returns a tuple with the EntityWatchlistScreeningId field value
// and a boolean to check if the value has been set.
func (o *ListWatchlistScreeningEntityHistoryRequest) GetEntityWatchlistScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistScreeningId, true
}

// SetEntityWatchlistScreeningId sets field value
func (o *ListWatchlistScreeningEntityHistoryRequest) SetEntityWatchlistScreeningId(v string) {
	o.EntityWatchlistScreeningId = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListWatchlistScreeningEntityHistoryRequest) GetCursor() string {
	if o == nil || o.Cursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListWatchlistScreeningEntityHistoryRequest) GetCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *ListWatchlistScreeningEntityHistoryRequest) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *ListWatchlistScreeningEntityHistoryRequest) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *ListWatchlistScreeningEntityHistoryRequest) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *ListWatchlistScreeningEntityHistoryRequest) UnsetCursor() {
	o.Cursor.Unset()
}

func (o ListWatchlistScreeningEntityHistoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["entity_watchlist_screening_id"] = o.EntityWatchlistScreeningId
	}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableListWatchlistScreeningEntityHistoryRequest struct {
	value *ListWatchlistScreeningEntityHistoryRequest
	isSet bool
}

func (v NullableListWatchlistScreeningEntityHistoryRequest) Get() *ListWatchlistScreeningEntityHistoryRequest {
	return v.value
}

func (v *NullableListWatchlistScreeningEntityHistoryRequest) Set(val *ListWatchlistScreeningEntityHistoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListWatchlistScreeningEntityHistoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListWatchlistScreeningEntityHistoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWatchlistScreeningEntityHistoryRequest(val *ListWatchlistScreeningEntityHistoryRequest) *NullableListWatchlistScreeningEntityHistoryRequest {
	return &NullableListWatchlistScreeningEntityHistoryRequest{value: val, isSet: true}
}

func (v NullableListWatchlistScreeningEntityHistoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWatchlistScreeningEntityHistoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


