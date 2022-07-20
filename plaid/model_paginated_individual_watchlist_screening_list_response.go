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

// PaginatedIndividualWatchlistScreeningListResponse Paginated list of individual watchlist screenings.
type PaginatedIndividualWatchlistScreeningListResponse struct {
	// List of individual watchlist screenings
	WatchlistScreenings []WatchlistScreeningIndividual `json:"watchlist_screenings"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedIndividualWatchlistScreeningListResponse PaginatedIndividualWatchlistScreeningListResponse

// NewPaginatedIndividualWatchlistScreeningListResponse instantiates a new PaginatedIndividualWatchlistScreeningListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedIndividualWatchlistScreeningListResponse(watchlistScreenings []WatchlistScreeningIndividual, nextCursor NullableString, requestId string) *PaginatedIndividualWatchlistScreeningListResponse {
	this := PaginatedIndividualWatchlistScreeningListResponse{}
	this.WatchlistScreenings = watchlistScreenings
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewPaginatedIndividualWatchlistScreeningListResponseWithDefaults instantiates a new PaginatedIndividualWatchlistScreeningListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedIndividualWatchlistScreeningListResponseWithDefaults() *PaginatedIndividualWatchlistScreeningListResponse {
	this := PaginatedIndividualWatchlistScreeningListResponse{}
	return &this
}

// GetWatchlistScreenings returns the WatchlistScreenings field value
func (o *PaginatedIndividualWatchlistScreeningListResponse) GetWatchlistScreenings() []WatchlistScreeningIndividual {
	if o == nil {
		var ret []WatchlistScreeningIndividual
		return ret
	}

	return o.WatchlistScreenings
}

// GetWatchlistScreeningsOk returns a tuple with the WatchlistScreenings field value
// and a boolean to check if the value has been set.
func (o *PaginatedIndividualWatchlistScreeningListResponse) GetWatchlistScreeningsOk() (*[]WatchlistScreeningIndividual, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistScreenings, true
}

// SetWatchlistScreenings sets field value
func (o *PaginatedIndividualWatchlistScreeningListResponse) SetWatchlistScreenings(v []WatchlistScreeningIndividual) {
	o.WatchlistScreenings = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginatedIndividualWatchlistScreeningListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedIndividualWatchlistScreeningListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *PaginatedIndividualWatchlistScreeningListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *PaginatedIndividualWatchlistScreeningListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaginatedIndividualWatchlistScreeningListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaginatedIndividualWatchlistScreeningListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaginatedIndividualWatchlistScreeningListResponse) MarshalJSON() ([]byte, error) {
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

func (o *PaginatedIndividualWatchlistScreeningListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaginatedIndividualWatchlistScreeningListResponse := _PaginatedIndividualWatchlistScreeningListResponse{}

	if err = json.Unmarshal(bytes, &varPaginatedIndividualWatchlistScreeningListResponse); err == nil {
		*o = PaginatedIndividualWatchlistScreeningListResponse(varPaginatedIndividualWatchlistScreeningListResponse)
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

type NullablePaginatedIndividualWatchlistScreeningListResponse struct {
	value *PaginatedIndividualWatchlistScreeningListResponse
	isSet bool
}

func (v NullablePaginatedIndividualWatchlistScreeningListResponse) Get() *PaginatedIndividualWatchlistScreeningListResponse {
	return v.value
}

func (v *NullablePaginatedIndividualWatchlistScreeningListResponse) Set(val *PaginatedIndividualWatchlistScreeningListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedIndividualWatchlistScreeningListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedIndividualWatchlistScreeningListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedIndividualWatchlistScreeningListResponse(val *PaginatedIndividualWatchlistScreeningListResponse) *NullablePaginatedIndividualWatchlistScreeningListResponse {
	return &NullablePaginatedIndividualWatchlistScreeningListResponse{value: val, isSet: true}
}

func (v NullablePaginatedIndividualWatchlistScreeningListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedIndividualWatchlistScreeningListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


