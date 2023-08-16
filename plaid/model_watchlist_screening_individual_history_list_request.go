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

// WatchlistScreeningIndividualHistoryListRequest Request input for listing changes to watchlist screenings for individuals
type WatchlistScreeningIndividualHistoryListRequest struct {
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// ID of the associated screening.
	WatchlistScreeningId string `json:"watchlist_screening_id"`
	// An identifier that determines which page of results you receive.
	Cursor NullableString `json:"cursor,omitempty"`
}

// NewWatchlistScreeningIndividualHistoryListRequest instantiates a new WatchlistScreeningIndividualHistoryListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividualHistoryListRequest(watchlistScreeningId string) *WatchlistScreeningIndividualHistoryListRequest {
	this := WatchlistScreeningIndividualHistoryListRequest{}
	this.WatchlistScreeningId = watchlistScreeningId
	return &this
}

// NewWatchlistScreeningIndividualHistoryListRequestWithDefaults instantiates a new WatchlistScreeningIndividualHistoryListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualHistoryListRequestWithDefaults() *WatchlistScreeningIndividualHistoryListRequest {
	this := WatchlistScreeningIndividualHistoryListRequest{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualHistoryListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualHistoryListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualHistoryListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WatchlistScreeningIndividualHistoryListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualHistoryListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualHistoryListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualHistoryListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WatchlistScreeningIndividualHistoryListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetWatchlistScreeningId returns the WatchlistScreeningId field value
func (o *WatchlistScreeningIndividualHistoryListRequest) GetWatchlistScreeningId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WatchlistScreeningId
}

// GetWatchlistScreeningIdOk returns a tuple with the WatchlistScreeningId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualHistoryListRequest) GetWatchlistScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistScreeningId, true
}

// SetWatchlistScreeningId sets field value
func (o *WatchlistScreeningIndividualHistoryListRequest) SetWatchlistScreeningId(v string) {
	o.WatchlistScreeningId = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatchlistScreeningIndividualHistoryListRequest) GetCursor() string {
	if o == nil || o.Cursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividualHistoryListRequest) GetCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualHistoryListRequest) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *WatchlistScreeningIndividualHistoryListRequest) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *WatchlistScreeningIndividualHistoryListRequest) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *WatchlistScreeningIndividualHistoryListRequest) UnsetCursor() {
	o.Cursor.Unset()
}

func (o WatchlistScreeningIndividualHistoryListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["watchlist_screening_id"] = o.WatchlistScreeningId
	}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistScreeningIndividualHistoryListRequest struct {
	value *WatchlistScreeningIndividualHistoryListRequest
	isSet bool
}

func (v NullableWatchlistScreeningIndividualHistoryListRequest) Get() *WatchlistScreeningIndividualHistoryListRequest {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualHistoryListRequest) Set(val *WatchlistScreeningIndividualHistoryListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualHistoryListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualHistoryListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualHistoryListRequest(val *WatchlistScreeningIndividualHistoryListRequest) *NullableWatchlistScreeningIndividualHistoryListRequest {
	return &NullableWatchlistScreeningIndividualHistoryListRequest{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualHistoryListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualHistoryListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


