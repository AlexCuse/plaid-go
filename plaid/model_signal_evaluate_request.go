/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.77.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SignalEvaluateRequest SignalEvaluateRequest defines the request schema for `/signal/evaluate`
type SignalEvaluateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// The `account_id` of the account whose verification status is to be modified
	AccountId string `json:"account_id"`
	// The unique ID that you would like to use to refer to this transaction. For your convenience mapping your internal data, you could use your internal ID/identifier for this transaction. The max length for this field is 36 characters.
	ClientTransactionId string `json:"client_transaction_id"`
	// The transaction amount, in USD (e.g. `102.05`)
	Amount float64 `json:"amount"`
	// `true` if the end user is present while initiating the ACH transfer and the endpoint is being called; `false` otherwise (for example, when the ACH transfer is scheduled and the end user is not present, or you call this endpoint after the ACH transfer but before submitting the Nacha file for ACH processing).
	UserPresent NullableBool `json:"user_present,omitempty"`
	// A unique ID that identifies the end user in your system. This ID is used to correlate requests by a user with multiple Items. The max length for this field is 36 characters.
	ClientUserId *string `json:"client_user_id,omitempty"`
	User *SignalUser `json:"user,omitempty"`
	Device *SignalDevice `json:"device,omitempty"`
}

// NewSignalEvaluateRequest instantiates a new SignalEvaluateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalEvaluateRequest(accessToken string, accountId string, clientTransactionId string, amount float64) *SignalEvaluateRequest {
	this := SignalEvaluateRequest{}
	this.AccessToken = accessToken
	this.AccountId = accountId
	this.ClientTransactionId = clientTransactionId
	this.Amount = amount
	return &this
}

// NewSignalEvaluateRequestWithDefaults instantiates a new SignalEvaluateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalEvaluateRequestWithDefaults() *SignalEvaluateRequest {
	this := SignalEvaluateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SignalEvaluateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalEvaluateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SignalEvaluateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SignalEvaluateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SignalEvaluateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalEvaluateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SignalEvaluateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SignalEvaluateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *SignalEvaluateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *SignalEvaluateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *SignalEvaluateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetAccountId returns the AccountId field value
func (o *SignalEvaluateRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *SignalEvaluateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *SignalEvaluateRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetClientTransactionId returns the ClientTransactionId field value
func (o *SignalEvaluateRequest) GetClientTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientTransactionId
}

// GetClientTransactionIdOk returns a tuple with the ClientTransactionId field value
// and a boolean to check if the value has been set.
func (o *SignalEvaluateRequest) GetClientTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientTransactionId, true
}

// SetClientTransactionId sets field value
func (o *SignalEvaluateRequest) SetClientTransactionId(v string) {
	o.ClientTransactionId = v
}

// GetAmount returns the Amount field value
func (o *SignalEvaluateRequest) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *SignalEvaluateRequest) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *SignalEvaluateRequest) SetAmount(v float64) {
	o.Amount = v
}

// GetUserPresent returns the UserPresent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalEvaluateRequest) GetUserPresent() bool {
	if o == nil || o.UserPresent.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UserPresent.Get()
}

// GetUserPresentOk returns a tuple with the UserPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalEvaluateRequest) GetUserPresentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPresent.Get(), o.UserPresent.IsSet()
}

// HasUserPresent returns a boolean if a field has been set.
func (o *SignalEvaluateRequest) HasUserPresent() bool {
	if o != nil && o.UserPresent.IsSet() {
		return true
	}

	return false
}

// SetUserPresent gets a reference to the given NullableBool and assigns it to the UserPresent field.
func (o *SignalEvaluateRequest) SetUserPresent(v bool) {
	o.UserPresent.Set(&v)
}
// SetUserPresentNil sets the value for UserPresent to be an explicit nil
func (o *SignalEvaluateRequest) SetUserPresentNil() {
	o.UserPresent.Set(nil)
}

// UnsetUserPresent ensures that no value is present for UserPresent, not even an explicit nil
func (o *SignalEvaluateRequest) UnsetUserPresent() {
	o.UserPresent.Unset()
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise.
func (o *SignalEvaluateRequest) GetClientUserId() string {
	if o == nil || o.ClientUserId == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalEvaluateRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil || o.ClientUserId == nil {
		return nil, false
	}
	return o.ClientUserId, true
}

// HasClientUserId returns a boolean if a field has been set.
func (o *SignalEvaluateRequest) HasClientUserId() bool {
	if o != nil && o.ClientUserId != nil {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given string and assigns it to the ClientUserId field.
func (o *SignalEvaluateRequest) SetClientUserId(v string) {
	o.ClientUserId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SignalEvaluateRequest) GetUser() SignalUser {
	if o == nil || o.User == nil {
		var ret SignalUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalEvaluateRequest) GetUserOk() (*SignalUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SignalEvaluateRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given SignalUser and assigns it to the User field.
func (o *SignalEvaluateRequest) SetUser(v SignalUser) {
	o.User = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *SignalEvaluateRequest) GetDevice() SignalDevice {
	if o == nil || o.Device == nil {
		var ret SignalDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalEvaluateRequest) GetDeviceOk() (*SignalDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *SignalEvaluateRequest) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given SignalDevice and assigns it to the Device field.
func (o *SignalEvaluateRequest) SetDevice(v SignalDevice) {
	o.Device = &v
}

func (o SignalEvaluateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["client_transaction_id"] = o.ClientTransactionId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.UserPresent.IsSet() {
		toSerialize["user_present"] = o.UserPresent.Get()
	}
	if o.ClientUserId != nil {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableSignalEvaluateRequest struct {
	value *SignalEvaluateRequest
	isSet bool
}

func (v NullableSignalEvaluateRequest) Get() *SignalEvaluateRequest {
	return v.value
}

func (v *NullableSignalEvaluateRequest) Set(val *SignalEvaluateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalEvaluateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalEvaluateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalEvaluateRequest(val *SignalEvaluateRequest) *NullableSignalEvaluateRequest {
	return &NullableSignalEvaluateRequest{value: val, isSet: true}
}

func (v NullableSignalEvaluateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalEvaluateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


