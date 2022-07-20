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

// ItemImportRequestUserAuth Object of user ID and auth token pair, permitting Plaid to aggregate a user’s accounts
type ItemImportRequestUserAuth struct {
	// Opaque user identifier
	UserId string `json:"user_id"`
	// Authorization token Plaid will use to aggregate this user’s accounts
	AuthToken string `json:"auth_token"`
}

// NewItemImportRequestUserAuth instantiates a new ItemImportRequestUserAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemImportRequestUserAuth(userId string, authToken string) *ItemImportRequestUserAuth {
	this := ItemImportRequestUserAuth{}
	this.UserId = userId
	this.AuthToken = authToken
	return &this
}

// NewItemImportRequestUserAuthWithDefaults instantiates a new ItemImportRequestUserAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemImportRequestUserAuthWithDefaults() *ItemImportRequestUserAuth {
	this := ItemImportRequestUserAuth{}
	return &this
}

// GetUserId returns the UserId field value
func (o *ItemImportRequestUserAuth) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ItemImportRequestUserAuth) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ItemImportRequestUserAuth) SetUserId(v string) {
	o.UserId = v
}

// GetAuthToken returns the AuthToken field value
func (o *ItemImportRequestUserAuth) GetAuthToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value
// and a boolean to check if the value has been set.
func (o *ItemImportRequestUserAuth) GetAuthTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthToken, true
}

// SetAuthToken sets field value
func (o *ItemImportRequestUserAuth) SetAuthToken(v string) {
	o.AuthToken = v
}

func (o ItemImportRequestUserAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["auth_token"] = o.AuthToken
	}
	return json.Marshal(toSerialize)
}

type NullableItemImportRequestUserAuth struct {
	value *ItemImportRequestUserAuth
	isSet bool
}

func (v NullableItemImportRequestUserAuth) Get() *ItemImportRequestUserAuth {
	return v.value
}

func (v *NullableItemImportRequestUserAuth) Set(val *ItemImportRequestUserAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableItemImportRequestUserAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableItemImportRequestUserAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemImportRequestUserAuth(val *ItemImportRequestUserAuth) *NullableItemImportRequestUserAuth {
	return &NullableItemImportRequestUserAuth{value: val, isSet: true}
}

func (v NullableItemImportRequestUserAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemImportRequestUserAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


