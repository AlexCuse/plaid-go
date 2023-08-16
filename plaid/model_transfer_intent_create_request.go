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

// TransferIntentCreateRequest Defines the request schema for `/transfer/intent/create`
type TransferIntentCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The Plaid `account_id` corresponding to the end-user account that will be debited or credited.
	AccountId NullableString `json:"account_id,omitempty"`
	// The id of the funding account to use, available in the Plaid Dashboard. This determines which of your business checking accounts will be credited or debited. Defaults to the account configured during onboarding.
	FundingAccountId NullableString `json:"funding_account_id,omitempty"`
	Mode TransferIntentCreateMode `json:"mode"`
	Network *TransferIntentCreateNetwork `json:"network,omitempty"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	// A description for the underlying transfer. Maximum of 8 characters.
	Description string `json:"description"`
	AchClass *ACHClass `json:"ach_class,omitempty"`
	// Plaid’s unique identifier for the origination account for the intent. If not provided, the default account will be used.
	OriginationAccountId NullableString `json:"origination_account_id,omitempty"`
	User TransferUserInRequest `json:"user"`
	// The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: The JSON values must be Strings (no nested JSON objects allowed) Only ASCII characters may be used Maximum of 50 key/value pairs Maximum key length of 40 characters Maximum value length of 500 characters 
	Metadata map[string]string `json:"metadata,omitempty"`
	// The currency of the transfer amount, e.g. \"USD\"
	IsoCurrencyCode *string `json:"iso_currency_code,omitempty"`
	// When `true`, the transfer requires a `GUARANTEED` decision by Plaid to proceed (Guarantee customers only).
	RequireGuarantee NullableBool `json:"require_guarantee,omitempty"`
}

// NewTransferIntentCreateRequest instantiates a new TransferIntentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferIntentCreateRequest(mode TransferIntentCreateMode, amount string, description string, user TransferUserInRequest) *TransferIntentCreateRequest {
	this := TransferIntentCreateRequest{}
	this.Mode = mode
	var network TransferIntentCreateNetwork = TRANSFERINTENTCREATENETWORK_SAME_DAY_ACH
	this.Network = &network
	this.Amount = amount
	this.Description = description
	this.User = user
	var requireGuarantee bool = false
	this.RequireGuarantee = *NewNullableBool(&requireGuarantee)
	return &this
}

// NewTransferIntentCreateRequestWithDefaults instantiates a new TransferIntentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferIntentCreateRequestWithDefaults() *TransferIntentCreateRequest {
	this := TransferIntentCreateRequest{}
	var network TransferIntentCreateNetwork = TRANSFERINTENTCREATENETWORK_SAME_DAY_ACH
	this.Network = &network
	var requireGuarantee bool = false
	this.RequireGuarantee = *NewNullableBool(&requireGuarantee)
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferIntentCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferIntentCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferIntentCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferIntentCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentCreateRequest) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *TransferIntentCreateRequest) SetAccountId(v string) {
	o.AccountId.Set(&v)
}
// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *TransferIntentCreateRequest) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *TransferIntentCreateRequest) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetFundingAccountId returns the FundingAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentCreateRequest) GetFundingAccountId() string {
	if o == nil || o.FundingAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.FundingAccountId.Get()
}

// GetFundingAccountIdOk returns a tuple with the FundingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentCreateRequest) GetFundingAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FundingAccountId.Get(), o.FundingAccountId.IsSet()
}

// HasFundingAccountId returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasFundingAccountId() bool {
	if o != nil && o.FundingAccountId.IsSet() {
		return true
	}

	return false
}

// SetFundingAccountId gets a reference to the given NullableString and assigns it to the FundingAccountId field.
func (o *TransferIntentCreateRequest) SetFundingAccountId(v string) {
	o.FundingAccountId.Set(&v)
}
// SetFundingAccountIdNil sets the value for FundingAccountId to be an explicit nil
func (o *TransferIntentCreateRequest) SetFundingAccountIdNil() {
	o.FundingAccountId.Set(nil)
}

// UnsetFundingAccountId ensures that no value is present for FundingAccountId, not even an explicit nil
func (o *TransferIntentCreateRequest) UnsetFundingAccountId() {
	o.FundingAccountId.Unset()
}

// GetMode returns the Mode field value
func (o *TransferIntentCreateRequest) GetMode() TransferIntentCreateMode {
	if o == nil {
		var ret TransferIntentCreateMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetModeOk() (*TransferIntentCreateMode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *TransferIntentCreateRequest) SetMode(v TransferIntentCreateMode) {
	o.Mode = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *TransferIntentCreateRequest) GetNetwork() TransferIntentCreateNetwork {
	if o == nil || o.Network == nil {
		var ret TransferIntentCreateNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetNetworkOk() (*TransferIntentCreateNetwork, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given TransferIntentCreateNetwork and assigns it to the Network field.
func (o *TransferIntentCreateRequest) SetNetwork(v TransferIntentCreateNetwork) {
	o.Network = &v
}

// GetAmount returns the Amount field value
func (o *TransferIntentCreateRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferIntentCreateRequest) SetAmount(v string) {
	o.Amount = v
}

// GetDescription returns the Description field value
func (o *TransferIntentCreateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransferIntentCreateRequest) SetDescription(v string) {
	o.Description = v
}

// GetAchClass returns the AchClass field value if set, zero value otherwise.
func (o *TransferIntentCreateRequest) GetAchClass() ACHClass {
	if o == nil || o.AchClass == nil {
		var ret ACHClass
		return ret
	}
	return *o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetAchClassOk() (*ACHClass, bool) {
	if o == nil || o.AchClass == nil {
		return nil, false
	}
	return o.AchClass, true
}

// HasAchClass returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasAchClass() bool {
	if o != nil && o.AchClass != nil {
		return true
	}

	return false
}

// SetAchClass gets a reference to the given ACHClass and assigns it to the AchClass field.
func (o *TransferIntentCreateRequest) SetAchClass(v ACHClass) {
	o.AchClass = &v
}

// GetOriginationAccountId returns the OriginationAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentCreateRequest) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentCreateRequest) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// HasOriginationAccountId returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasOriginationAccountId() bool {
	if o != nil && o.OriginationAccountId.IsSet() {
		return true
	}

	return false
}

// SetOriginationAccountId gets a reference to the given NullableString and assigns it to the OriginationAccountId field.
func (o *TransferIntentCreateRequest) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}
// SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil
func (o *TransferIntentCreateRequest) SetOriginationAccountIdNil() {
	o.OriginationAccountId.Set(nil)
}

// UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
func (o *TransferIntentCreateRequest) UnsetOriginationAccountId() {
	o.OriginationAccountId.Unset()
}

// GetUser returns the User field value
func (o *TransferIntentCreateRequest) GetUser() TransferUserInRequest {
	if o == nil {
		var ret TransferUserInRequest
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetUserOk() (*TransferUserInRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferIntentCreateRequest) SetUser(v TransferUserInRequest) {
	o.User = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentCreateRequest) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentCreateRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *TransferIntentCreateRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise.
func (o *TransferIntentCreateRequest) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateRequest) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil || o.IsoCurrencyCode == nil {
		return nil, false
	}
	return o.IsoCurrencyCode, true
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode != nil {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given string and assigns it to the IsoCurrencyCode field.
func (o *TransferIntentCreateRequest) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = &v
}

// GetRequireGuarantee returns the RequireGuarantee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferIntentCreateRequest) GetRequireGuarantee() bool {
	if o == nil || o.RequireGuarantee.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RequireGuarantee.Get()
}

// GetRequireGuaranteeOk returns a tuple with the RequireGuarantee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferIntentCreateRequest) GetRequireGuaranteeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequireGuarantee.Get(), o.RequireGuarantee.IsSet()
}

// HasRequireGuarantee returns a boolean if a field has been set.
func (o *TransferIntentCreateRequest) HasRequireGuarantee() bool {
	if o != nil && o.RequireGuarantee.IsSet() {
		return true
	}

	return false
}

// SetRequireGuarantee gets a reference to the given NullableBool and assigns it to the RequireGuarantee field.
func (o *TransferIntentCreateRequest) SetRequireGuarantee(v bool) {
	o.RequireGuarantee.Set(&v)
}
// SetRequireGuaranteeNil sets the value for RequireGuarantee to be an explicit nil
func (o *TransferIntentCreateRequest) SetRequireGuaranteeNil() {
	o.RequireGuarantee.Set(nil)
}

// UnsetRequireGuarantee ensures that no value is present for RequireGuarantee, not even an explicit nil
func (o *TransferIntentCreateRequest) UnsetRequireGuarantee() {
	o.RequireGuarantee.Unset()
}

func (o TransferIntentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if o.FundingAccountId.IsSet() {
		toSerialize["funding_account_id"] = o.FundingAccountId.Get()
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.AchClass != nil {
		toSerialize["ach_class"] = o.AchClass
	}
	if o.OriginationAccountId.IsSet() {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.IsoCurrencyCode != nil {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if o.RequireGuarantee.IsSet() {
		toSerialize["require_guarantee"] = o.RequireGuarantee.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransferIntentCreateRequest struct {
	value *TransferIntentCreateRequest
	isSet bool
}

func (v NullableTransferIntentCreateRequest) Get() *TransferIntentCreateRequest {
	return v.value
}

func (v *NullableTransferIntentCreateRequest) Set(val *TransferIntentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentCreateRequest(val *TransferIntentCreateRequest) *NullableTransferIntentCreateRequest {
	return &NullableTransferIntentCreateRequest{value: val, isSet: true}
}

func (v NullableTransferIntentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


