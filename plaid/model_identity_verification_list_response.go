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

// IdentityVerificationListResponse Paginated list of Plaid sessions.
type IdentityVerificationListResponse struct {
	// List of Plaid sessions
	IdentityVerifications []IdentityVerification `json:"identity_verifications"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _IdentityVerificationListResponse IdentityVerificationListResponse

// NewIdentityVerificationListResponse instantiates a new IdentityVerificationListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationListResponse(identityVerifications []IdentityVerification, nextCursor NullableString, requestId string) *IdentityVerificationListResponse {
	this := IdentityVerificationListResponse{}
	this.IdentityVerifications = identityVerifications
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewIdentityVerificationListResponseWithDefaults instantiates a new IdentityVerificationListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationListResponseWithDefaults() *IdentityVerificationListResponse {
	this := IdentityVerificationListResponse{}
	return &this
}

// GetIdentityVerifications returns the IdentityVerifications field value
func (o *IdentityVerificationListResponse) GetIdentityVerifications() []IdentityVerification {
	if o == nil {
		var ret []IdentityVerification
		return ret
	}

	return o.IdentityVerifications
}

// GetIdentityVerificationsOk returns a tuple with the IdentityVerifications field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationListResponse) GetIdentityVerificationsOk() (*[]IdentityVerification, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdentityVerifications, true
}

// SetIdentityVerifications sets field value
func (o *IdentityVerificationListResponse) SetIdentityVerifications(v []IdentityVerification) {
	o.IdentityVerifications = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *IdentityVerificationListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *IdentityVerificationListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *IdentityVerificationListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o IdentityVerificationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity_verifications"] = o.IdentityVerifications
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

func (o *IdentityVerificationListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityVerificationListResponse := _IdentityVerificationListResponse{}

	if err = json.Unmarshal(bytes, &varIdentityVerificationListResponse); err == nil {
		*o = IdentityVerificationListResponse(varIdentityVerificationListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity_verifications")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityVerificationListResponse struct {
	value *IdentityVerificationListResponse
	isSet bool
}

func (v NullableIdentityVerificationListResponse) Get() *IdentityVerificationListResponse {
	return v.value
}

func (v *NullableIdentityVerificationListResponse) Set(val *IdentityVerificationListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationListResponse(val *IdentityVerificationListResponse) *NullableIdentityVerificationListResponse {
	return &NullableIdentityVerificationListResponse{value: val, isSet: true}
}

func (v NullableIdentityVerificationListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


