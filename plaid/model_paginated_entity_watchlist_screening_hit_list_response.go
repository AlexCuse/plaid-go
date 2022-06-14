/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaginatedEntityWatchlistScreeningHitListResponse Paginated list of entity watchlist screening hits
type PaginatedEntityWatchlistScreeningHitListResponse struct {
	// List of entity watchlist screening hits
	EntityWatchlistScreeningHits []EntityWatchlistScreeningHit `json:"entity_watchlist_screening_hits"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedEntityWatchlistScreeningHitListResponse PaginatedEntityWatchlistScreeningHitListResponse

// NewPaginatedEntityWatchlistScreeningHitListResponse instantiates a new PaginatedEntityWatchlistScreeningHitListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEntityWatchlistScreeningHitListResponse(entityWatchlistScreeningHits []EntityWatchlistScreeningHit, nextCursor NullableString, requestId string) *PaginatedEntityWatchlistScreeningHitListResponse {
	this := PaginatedEntityWatchlistScreeningHitListResponse{}
	this.EntityWatchlistScreeningHits = entityWatchlistScreeningHits
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewPaginatedEntityWatchlistScreeningHitListResponseWithDefaults instantiates a new PaginatedEntityWatchlistScreeningHitListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEntityWatchlistScreeningHitListResponseWithDefaults() *PaginatedEntityWatchlistScreeningHitListResponse {
	this := PaginatedEntityWatchlistScreeningHitListResponse{}
	return &this
}

// GetEntityWatchlistScreeningHits returns the EntityWatchlistScreeningHits field value
func (o *PaginatedEntityWatchlistScreeningHitListResponse) GetEntityWatchlistScreeningHits() []EntityWatchlistScreeningHit {
	if o == nil {
		var ret []EntityWatchlistScreeningHit
		return ret
	}

	return o.EntityWatchlistScreeningHits
}

// GetEntityWatchlistScreeningHitsOk returns a tuple with the EntityWatchlistScreeningHits field value
// and a boolean to check if the value has been set.
func (o *PaginatedEntityWatchlistScreeningHitListResponse) GetEntityWatchlistScreeningHitsOk() (*[]EntityWatchlistScreeningHit, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistScreeningHits, true
}

// SetEntityWatchlistScreeningHits sets field value
func (o *PaginatedEntityWatchlistScreeningHitListResponse) SetEntityWatchlistScreeningHits(v []EntityWatchlistScreeningHit) {
	o.EntityWatchlistScreeningHits = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginatedEntityWatchlistScreeningHitListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedEntityWatchlistScreeningHitListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *PaginatedEntityWatchlistScreeningHitListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *PaginatedEntityWatchlistScreeningHitListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaginatedEntityWatchlistScreeningHitListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaginatedEntityWatchlistScreeningHitListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaginatedEntityWatchlistScreeningHitListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_watchlist_screening_hits"] = o.EntityWatchlistScreeningHits
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

func (o *PaginatedEntityWatchlistScreeningHitListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaginatedEntityWatchlistScreeningHitListResponse := _PaginatedEntityWatchlistScreeningHitListResponse{}

	if err = json.Unmarshal(bytes, &varPaginatedEntityWatchlistScreeningHitListResponse); err == nil {
		*o = PaginatedEntityWatchlistScreeningHitListResponse(varPaginatedEntityWatchlistScreeningHitListResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_watchlist_screening_hits")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedEntityWatchlistScreeningHitListResponse struct {
	value *PaginatedEntityWatchlistScreeningHitListResponse
	isSet bool
}

func (v NullablePaginatedEntityWatchlistScreeningHitListResponse) Get() *PaginatedEntityWatchlistScreeningHitListResponse {
	return v.value
}

func (v *NullablePaginatedEntityWatchlistScreeningHitListResponse) Set(val *PaginatedEntityWatchlistScreeningHitListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEntityWatchlistScreeningHitListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEntityWatchlistScreeningHitListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEntityWatchlistScreeningHitListResponse(val *PaginatedEntityWatchlistScreeningHitListResponse) *NullablePaginatedEntityWatchlistScreeningHitListResponse {
	return &NullablePaginatedEntityWatchlistScreeningHitListResponse{value: val, isSet: true}
}

func (v NullablePaginatedEntityWatchlistScreeningHitListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEntityWatchlistScreeningHitListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


