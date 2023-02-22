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

// TransactionsGetRequest TransactionsGetRequest defines the request schema for `/transactions/get`
type TransactionsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	Options *TransactionsGetRequestOptions `json:"options,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The earliest date for which data should be returned. Dates should be formatted as YYYY-MM-DD.
	StartDate string `json:"start_date"`
	// The latest date for which data should be returned. Dates should be formatted as YYYY-MM-DD.
	EndDate string `json:"end_date"`
}

// NewTransactionsGetRequest instantiates a new TransactionsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsGetRequest(accessToken string, startDate string, endDate string) *TransactionsGetRequest {
	this := TransactionsGetRequest{}
	this.AccessToken = accessToken
	this.StartDate = startDate
	this.EndDate = endDate
	return &this
}

// NewTransactionsGetRequestWithDefaults instantiates a new TransactionsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsGetRequestWithDefaults() *TransactionsGetRequest {
	this := TransactionsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransactionsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransactionsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransactionsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TransactionsGetRequest) GetOptions() TransactionsGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret TransactionsGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsGetRequest) GetOptionsOk() (*TransactionsGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TransactionsGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given TransactionsGetRequestOptions and assigns it to the Options field.
func (o *TransactionsGetRequest) SetOptions(v TransactionsGetRequestOptions) {
	o.Options = &v
}

// GetAccessToken returns the AccessToken field value
func (o *TransactionsGetRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TransactionsGetRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TransactionsGetRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransactionsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransactionsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransactionsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetStartDate returns the StartDate field value
func (o *TransactionsGetRequest) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *TransactionsGetRequest) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *TransactionsGetRequest) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *TransactionsGetRequest) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *TransactionsGetRequest) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *TransactionsGetRequest) SetEndDate(v string) {
	o.EndDate = v
}

func (o TransactionsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if true {
		toSerialize["end_date"] = o.EndDate
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsGetRequest struct {
	value *TransactionsGetRequest
	isSet bool
}

func (v NullableTransactionsGetRequest) Get() *TransactionsGetRequest {
	return v.value
}

func (v *NullableTransactionsGetRequest) Set(val *TransactionsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsGetRequest(val *TransactionsGetRequest) *NullableTransactionsGetRequest {
	return &NullableTransactionsGetRequest{value: val, isSet: true}
}

func (v NullableTransactionsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


