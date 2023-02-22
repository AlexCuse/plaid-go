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

// WatchlistScreeningIndividualListResponse Paginated list of individual watchlist screenings.
type WatchlistScreeningIndividualListResponse struct {
	// List of individual watchlist screenings
	WatchlistScreenings []WatchlistScreeningIndividual `json:"watchlist_screenings"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningIndividualListResponse WatchlistScreeningIndividualListResponse

// NewWatchlistScreeningIndividualListResponse instantiates a new WatchlistScreeningIndividualListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividualListResponse(watchlistScreenings []WatchlistScreeningIndividual, nextCursor NullableString, requestId string) *WatchlistScreeningIndividualListResponse {
	this := WatchlistScreeningIndividualListResponse{}
	this.WatchlistScreenings = watchlistScreenings
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningIndividualListResponseWithDefaults instantiates a new WatchlistScreeningIndividualListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualListResponseWithDefaults() *WatchlistScreeningIndividualListResponse {
	this := WatchlistScreeningIndividualListResponse{}
	return &this
}

// GetWatchlistScreenings returns the WatchlistScreenings field value
func (o *WatchlistScreeningIndividualListResponse) GetWatchlistScreenings() []WatchlistScreeningIndividual {
	if o == nil {
		var ret []WatchlistScreeningIndividual
		return ret
	}

	return o.WatchlistScreenings
}

// GetWatchlistScreeningsOk returns a tuple with the WatchlistScreenings field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualListResponse) GetWatchlistScreeningsOk() (*[]WatchlistScreeningIndividual, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistScreenings, true
}

// SetWatchlistScreenings sets field value
func (o *WatchlistScreeningIndividualListResponse) SetWatchlistScreenings(v []WatchlistScreeningIndividual) {
	o.WatchlistScreenings = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningIndividualListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividualListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *WatchlistScreeningIndividualListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningIndividualListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningIndividualListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningIndividualListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["watchlist_screenings"] = o.WatchlistScreenings
	}
	if true {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningIndividualListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningIndividualListResponse := _WatchlistScreeningIndividualListResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningIndividualListResponse); err == nil {
		*o = WatchlistScreeningIndividualListResponse(varWatchlistScreeningIndividualListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "watchlist_screenings")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningIndividualListResponse struct {
	value *WatchlistScreeningIndividualListResponse
	isSet bool
}

func (v NullableWatchlistScreeningIndividualListResponse) Get() *WatchlistScreeningIndividualListResponse {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualListResponse) Set(val *WatchlistScreeningIndividualListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualListResponse(val *WatchlistScreeningIndividualListResponse) *NullableWatchlistScreeningIndividualListResponse {
	return &NullableWatchlistScreeningIndividualListResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


