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

// WatchlistScreeningEntityCreateRequest Request input for creating an entity screening review
type WatchlistScreeningEntityCreateRequest struct {
	SearchTerms EntityWatchlistSearchTerms `json:"search_terms"`
	// A unique ID that identifies the end user in your system. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the Link Token Create `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.
	ClientUserId *string `json:"client_user_id,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
}

// NewWatchlistScreeningEntityCreateRequest instantiates a new WatchlistScreeningEntityCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningEntityCreateRequest(searchTerms EntityWatchlistSearchTerms) *WatchlistScreeningEntityCreateRequest {
	this := WatchlistScreeningEntityCreateRequest{}
	this.SearchTerms = searchTerms
	return &this
}

// NewWatchlistScreeningEntityCreateRequestWithDefaults instantiates a new WatchlistScreeningEntityCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningEntityCreateRequestWithDefaults() *WatchlistScreeningEntityCreateRequest {
	this := WatchlistScreeningEntityCreateRequest{}
	return &this
}

// GetSearchTerms returns the SearchTerms field value
func (o *WatchlistScreeningEntityCreateRequest) GetSearchTerms() EntityWatchlistSearchTerms {
	if o == nil {
		var ret EntityWatchlistSearchTerms
		return ret
	}

	return o.SearchTerms
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityCreateRequest) GetSearchTermsOk() (*EntityWatchlistSearchTerms, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchTerms, true
}

// SetSearchTerms sets field value
func (o *WatchlistScreeningEntityCreateRequest) SetSearchTerms(v EntityWatchlistSearchTerms) {
	o.SearchTerms = v
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise.
func (o *WatchlistScreeningEntityCreateRequest) GetClientUserId() string {
	if o == nil || o.ClientUserId == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityCreateRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil || o.ClientUserId == nil {
		return nil, false
	}
	return o.ClientUserId, true
}

// HasClientUserId returns a boolean if a field has been set.
func (o *WatchlistScreeningEntityCreateRequest) HasClientUserId() bool {
	if o != nil && o.ClientUserId != nil {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given string and assigns it to the ClientUserId field.
func (o *WatchlistScreeningEntityCreateRequest) SetClientUserId(v string) {
	o.ClientUserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WatchlistScreeningEntityCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WatchlistScreeningEntityCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WatchlistScreeningEntityCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WatchlistScreeningEntityCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WatchlistScreeningEntityCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WatchlistScreeningEntityCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o WatchlistScreeningEntityCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["search_terms"] = o.SearchTerms
	}
	if o.ClientUserId != nil {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistScreeningEntityCreateRequest struct {
	value *WatchlistScreeningEntityCreateRequest
	isSet bool
}

func (v NullableWatchlistScreeningEntityCreateRequest) Get() *WatchlistScreeningEntityCreateRequest {
	return v.value
}

func (v *NullableWatchlistScreeningEntityCreateRequest) Set(val *WatchlistScreeningEntityCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningEntityCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningEntityCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningEntityCreateRequest(val *WatchlistScreeningEntityCreateRequest) *NullableWatchlistScreeningEntityCreateRequest {
	return &NullableWatchlistScreeningEntityCreateRequest{value: val, isSet: true}
}

func (v NullableWatchlistScreeningEntityCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningEntityCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


