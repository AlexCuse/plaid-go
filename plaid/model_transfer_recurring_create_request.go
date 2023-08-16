/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.413.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferRecurringCreateRequest Defines the request schema for `/transfer/recurring/create`
type TransferRecurringCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The Plaid `access_token` for the account that will be debited or credited. Required if not using `payment_profile_token`.
	AccessToken string `json:"access_token"`
	// A random key provided by the client, per unique recurring transfer. Maximum of 50 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. For example, if a request to create a recurring fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single recurring transfer is created.
	IdempotencyKey string `json:"idempotency_key"`
	// The Plaid `account_id` corresponding to the end-user account that will be debited or credited. Required when creating a transfer using an `access_token`.
	AccountId string `json:"account_id"`
	// The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited. Defaults to the account configured during onboarding.
	FundingAccountId NullableString `json:"funding_account_id,omitempty"`
	Type TransferType `json:"type"`
	Network TransferNetwork `json:"network"`
	AchClass *ACHClass `json:"ach_class,omitempty"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	// If the end user is initiating the specific transfer themselves via an interactive UI, this should be `true`; for automatic recurring payments where the end user is not actually initiating each individual transfer, it should be `false`.
	UserPresent NullableBool `json:"user_present"`
	// The currency of the transfer amount. The default value is \"USD\".
	IsoCurrencyCode *string `json:"iso_currency_code,omitempty"`
	// The description of the recurring transfer.
	Description string `json:"description"`
	// Plaid’s unique identifier for a test clock. This field may only be used when using `sandbox` environment. If provided, the created `recurring_transfer` is associated with the `test_clock`. New originations are automatically generated when the associated `test_clock` advances.
	TestClockId NullableString `json:"test_clock_id,omitempty"`
	Schedule TransferRecurringSchedule `json:"schedule"`
	User TransferUserInRequest `json:"user"`
	Device TransferDevice `json:"device"`
}

// NewTransferRecurringCreateRequest instantiates a new TransferRecurringCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRecurringCreateRequest(accessToken string, idempotencyKey string, accountId string, type_ TransferType, network TransferNetwork, amount string, userPresent NullableBool, description string, schedule TransferRecurringSchedule, user TransferUserInRequest, device TransferDevice) *TransferRecurringCreateRequest {
	this := TransferRecurringCreateRequest{}
	this.AccessToken = accessToken
	this.IdempotencyKey = idempotencyKey
	this.AccountId = accountId
	this.Type = type_
	this.Network = network
	this.Amount = amount
	this.UserPresent = userPresent
	this.Description = description
	this.Schedule = schedule
	this.User = user
	this.Device = device
	return &this
}

// NewTransferRecurringCreateRequestWithDefaults instantiates a new TransferRecurringCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRecurringCreateRequestWithDefaults() *TransferRecurringCreateRequest {
	this := TransferRecurringCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferRecurringCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferRecurringCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferRecurringCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferRecurringCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferRecurringCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferRecurringCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *TransferRecurringCreateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TransferRecurringCreateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *TransferRecurringCreateRequest) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *TransferRecurringCreateRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetAccountId returns the AccountId field value
func (o *TransferRecurringCreateRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransferRecurringCreateRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetFundingAccountId returns the FundingAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferRecurringCreateRequest) GetFundingAccountId() string {
	if o == nil || o.FundingAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.FundingAccountId.Get()
}

// GetFundingAccountIdOk returns a tuple with the FundingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferRecurringCreateRequest) GetFundingAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FundingAccountId.Get(), o.FundingAccountId.IsSet()
}

// HasFundingAccountId returns a boolean if a field has been set.
func (o *TransferRecurringCreateRequest) HasFundingAccountId() bool {
	if o != nil && o.FundingAccountId.IsSet() {
		return true
	}

	return false
}

// SetFundingAccountId gets a reference to the given NullableString and assigns it to the FundingAccountId field.
func (o *TransferRecurringCreateRequest) SetFundingAccountId(v string) {
	o.FundingAccountId.Set(&v)
}
// SetFundingAccountIdNil sets the value for FundingAccountId to be an explicit nil
func (o *TransferRecurringCreateRequest) SetFundingAccountIdNil() {
	o.FundingAccountId.Set(nil)
}

// UnsetFundingAccountId ensures that no value is present for FundingAccountId, not even an explicit nil
func (o *TransferRecurringCreateRequest) UnsetFundingAccountId() {
	o.FundingAccountId.Unset()
}

// GetType returns the Type field value
func (o *TransferRecurringCreateRequest) GetType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetTypeOk() (*TransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferRecurringCreateRequest) SetType(v TransferType) {
	o.Type = v
}

// GetNetwork returns the Network field value
func (o *TransferRecurringCreateRequest) GetNetwork() TransferNetwork {
	if o == nil {
		var ret TransferNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetNetworkOk() (*TransferNetwork, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TransferRecurringCreateRequest) SetNetwork(v TransferNetwork) {
	o.Network = v
}

// GetAchClass returns the AchClass field value if set, zero value otherwise.
func (o *TransferRecurringCreateRequest) GetAchClass() ACHClass {
	if o == nil || o.AchClass == nil {
		var ret ACHClass
		return ret
	}
	return *o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetAchClassOk() (*ACHClass, bool) {
	if o == nil || o.AchClass == nil {
		return nil, false
	}
	return o.AchClass, true
}

// HasAchClass returns a boolean if a field has been set.
func (o *TransferRecurringCreateRequest) HasAchClass() bool {
	if o != nil && o.AchClass != nil {
		return true
	}

	return false
}

// SetAchClass gets a reference to the given ACHClass and assigns it to the AchClass field.
func (o *TransferRecurringCreateRequest) SetAchClass(v ACHClass) {
	o.AchClass = &v
}

// GetAmount returns the Amount field value
func (o *TransferRecurringCreateRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferRecurringCreateRequest) SetAmount(v string) {
	o.Amount = v
}

// GetUserPresent returns the UserPresent field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *TransferRecurringCreateRequest) GetUserPresent() bool {
	if o == nil || o.UserPresent.Get() == nil {
		var ret bool
		return ret
	}

	return *o.UserPresent.Get()
}

// GetUserPresentOk returns a tuple with the UserPresent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferRecurringCreateRequest) GetUserPresentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPresent.Get(), o.UserPresent.IsSet()
}

// SetUserPresent sets field value
func (o *TransferRecurringCreateRequest) SetUserPresent(v bool) {
	o.UserPresent.Set(&v)
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise.
func (o *TransferRecurringCreateRequest) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil || o.IsoCurrencyCode == nil {
		return nil, false
	}
	return o.IsoCurrencyCode, true
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *TransferRecurringCreateRequest) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode != nil {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given string and assigns it to the IsoCurrencyCode field.
func (o *TransferRecurringCreateRequest) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = &v
}

// GetDescription returns the Description field value
func (o *TransferRecurringCreateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransferRecurringCreateRequest) SetDescription(v string) {
	o.Description = v
}

// GetTestClockId returns the TestClockId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferRecurringCreateRequest) GetTestClockId() string {
	if o == nil || o.TestClockId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TestClockId.Get()
}

// GetTestClockIdOk returns a tuple with the TestClockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferRecurringCreateRequest) GetTestClockIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TestClockId.Get(), o.TestClockId.IsSet()
}

// HasTestClockId returns a boolean if a field has been set.
func (o *TransferRecurringCreateRequest) HasTestClockId() bool {
	if o != nil && o.TestClockId.IsSet() {
		return true
	}

	return false
}

// SetTestClockId gets a reference to the given NullableString and assigns it to the TestClockId field.
func (o *TransferRecurringCreateRequest) SetTestClockId(v string) {
	o.TestClockId.Set(&v)
}
// SetTestClockIdNil sets the value for TestClockId to be an explicit nil
func (o *TransferRecurringCreateRequest) SetTestClockIdNil() {
	o.TestClockId.Set(nil)
}

// UnsetTestClockId ensures that no value is present for TestClockId, not even an explicit nil
func (o *TransferRecurringCreateRequest) UnsetTestClockId() {
	o.TestClockId.Unset()
}

// GetSchedule returns the Schedule field value
func (o *TransferRecurringCreateRequest) GetSchedule() TransferRecurringSchedule {
	if o == nil {
		var ret TransferRecurringSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetScheduleOk() (*TransferRecurringSchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *TransferRecurringCreateRequest) SetSchedule(v TransferRecurringSchedule) {
	o.Schedule = v
}

// GetUser returns the User field value
func (o *TransferRecurringCreateRequest) GetUser() TransferUserInRequest {
	if o == nil {
		var ret TransferUserInRequest
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetUserOk() (*TransferUserInRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferRecurringCreateRequest) SetUser(v TransferUserInRequest) {
	o.User = v
}

// GetDevice returns the Device field value
func (o *TransferRecurringCreateRequest) GetDevice() TransferDevice {
	if o == nil {
		var ret TransferDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *TransferRecurringCreateRequest) GetDeviceOk() (*TransferDevice, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *TransferRecurringCreateRequest) SetDevice(v TransferDevice) {
	o.Device = v
}

func (o TransferRecurringCreateRequest) MarshalJSON() ([]byte, error) {
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
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if o.FundingAccountId.IsSet() {
		toSerialize["funding_account_id"] = o.FundingAccountId.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if o.AchClass != nil {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["user_present"] = o.UserPresent.Get()
	}
	if o.IsoCurrencyCode != nil {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.TestClockId.IsSet() {
		toSerialize["test_clock_id"] = o.TestClockId.Get()
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableTransferRecurringCreateRequest struct {
	value *TransferRecurringCreateRequest
	isSet bool
}

func (v NullableTransferRecurringCreateRequest) Get() *TransferRecurringCreateRequest {
	return v.value
}

func (v *NullableTransferRecurringCreateRequest) Set(val *TransferRecurringCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRecurringCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRecurringCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRecurringCreateRequest(val *TransferRecurringCreateRequest) *NullableTransferRecurringCreateRequest {
	return &NullableTransferRecurringCreateRequest{value: val, isSet: true}
}

func (v NullableTransferRecurringCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRecurringCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


