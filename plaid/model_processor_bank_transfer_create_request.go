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

// ProcessorBankTransferCreateRequest Defines the request schema for `/processor/bank_transfer/create`
type ProcessorBankTransferCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A random key provided by the client, per unique bank transfer. Maximum of 50 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. For example, if a request to create a bank transfer fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single bank transfer is created.
	IdempotencyKey string `json:"idempotency_key"`
	// The processor token obtained from the Plaid integration partner. Processor tokens are in the format: `processor-<environment>-<identifier>`
	ProcessorToken string `json:"processor_token"`
	Type BankTransferType `json:"type"`
	Network BankTransferNetwork `json:"network"`
	// The amount of the bank transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	// The currency of the transfer amount – should be set to \"USD\".
	IsoCurrencyCode string `json:"iso_currency_code"`
	// The transfer description. Maximum of 10 characters.
	Description string `json:"description"`
	AchClass *ACHClass `json:"ach_class,omitempty"`
	User BankTransferUser `json:"user"`
	// An arbitrary string provided by the client for storage with the bank transfer. May be up to 100 characters.
	CustomTag NullableString `json:"custom_tag,omitempty"`
	// The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: The JSON values must be Strings (no nested JSON objects allowed) Only ASCII characters may be used Maximum of 50 key/value pairs Maximum key length of 40 characters Maximum value length of 500 characters 
	Metadata map[string]string `json:"metadata,omitempty"`
	// Plaid’s unique identifier for the origination account for this transfer. If you have more than one origination account, this value must be specified.
	OriginationAccountId NullableString `json:"origination_account_id,omitempty"`
}

// NewProcessorBankTransferCreateRequest instantiates a new ProcessorBankTransferCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorBankTransferCreateRequest(idempotencyKey string, processorToken string, type_ BankTransferType, network BankTransferNetwork, amount string, isoCurrencyCode string, description string, user BankTransferUser) *ProcessorBankTransferCreateRequest {
	this := ProcessorBankTransferCreateRequest{}
	this.IdempotencyKey = idempotencyKey
	this.ProcessorToken = processorToken
	this.Type = type_
	this.Network = network
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	this.Description = description
	this.User = user
	return &this
}

// NewProcessorBankTransferCreateRequestWithDefaults instantiates a new ProcessorBankTransferCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorBankTransferCreateRequestWithDefaults() *ProcessorBankTransferCreateRequest {
	this := ProcessorBankTransferCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorBankTransferCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorBankTransferCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorBankTransferCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorBankTransferCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorBankTransferCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorBankTransferCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *ProcessorBankTransferCreateRequest) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *ProcessorBankTransferCreateRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetProcessorToken returns the ProcessorToken field value
func (o *ProcessorBankTransferCreateRequest) GetProcessorToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessorToken
}

// GetProcessorTokenOk returns a tuple with the ProcessorToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetProcessorTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProcessorToken, true
}

// SetProcessorToken sets field value
func (o *ProcessorBankTransferCreateRequest) SetProcessorToken(v string) {
	o.ProcessorToken = v
}

// GetType returns the Type field value
func (o *ProcessorBankTransferCreateRequest) GetType() BankTransferType {
	if o == nil {
		var ret BankTransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetTypeOk() (*BankTransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProcessorBankTransferCreateRequest) SetType(v BankTransferType) {
	o.Type = v
}

// GetNetwork returns the Network field value
func (o *ProcessorBankTransferCreateRequest) GetNetwork() BankTransferNetwork {
	if o == nil {
		var ret BankTransferNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetNetworkOk() (*BankTransferNetwork, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *ProcessorBankTransferCreateRequest) SetNetwork(v BankTransferNetwork) {
	o.Network = v
}

// GetAmount returns the Amount field value
func (o *ProcessorBankTransferCreateRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ProcessorBankTransferCreateRequest) SetAmount(v string) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *ProcessorBankTransferCreateRequest) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *ProcessorBankTransferCreateRequest) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetDescription returns the Description field value
func (o *ProcessorBankTransferCreateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ProcessorBankTransferCreateRequest) SetDescription(v string) {
	o.Description = v
}

// GetAchClass returns the AchClass field value if set, zero value otherwise.
func (o *ProcessorBankTransferCreateRequest) GetAchClass() ACHClass {
	if o == nil || o.AchClass == nil {
		var ret ACHClass
		return ret
	}
	return *o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetAchClassOk() (*ACHClass, bool) {
	if o == nil || o.AchClass == nil {
		return nil, false
	}
	return o.AchClass, true
}

// HasAchClass returns a boolean if a field has been set.
func (o *ProcessorBankTransferCreateRequest) HasAchClass() bool {
	if o != nil && o.AchClass != nil {
		return true
	}

	return false
}

// SetAchClass gets a reference to the given ACHClass and assigns it to the AchClass field.
func (o *ProcessorBankTransferCreateRequest) SetAchClass(v ACHClass) {
	o.AchClass = &v
}

// GetUser returns the User field value
func (o *ProcessorBankTransferCreateRequest) GetUser() BankTransferUser {
	if o == nil {
		var ret BankTransferUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ProcessorBankTransferCreateRequest) GetUserOk() (*BankTransferUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ProcessorBankTransferCreateRequest) SetUser(v BankTransferUser) {
	o.User = v
}

// GetCustomTag returns the CustomTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorBankTransferCreateRequest) GetCustomTag() string {
	if o == nil || o.CustomTag.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomTag.Get()
}

// GetCustomTagOk returns a tuple with the CustomTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorBankTransferCreateRequest) GetCustomTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomTag.Get(), o.CustomTag.IsSet()
}

// HasCustomTag returns a boolean if a field has been set.
func (o *ProcessorBankTransferCreateRequest) HasCustomTag() bool {
	if o != nil && o.CustomTag.IsSet() {
		return true
	}

	return false
}

// SetCustomTag gets a reference to the given NullableString and assigns it to the CustomTag field.
func (o *ProcessorBankTransferCreateRequest) SetCustomTag(v string) {
	o.CustomTag.Set(&v)
}
// SetCustomTagNil sets the value for CustomTag to be an explicit nil
func (o *ProcessorBankTransferCreateRequest) SetCustomTagNil() {
	o.CustomTag.Set(nil)
}

// UnsetCustomTag ensures that no value is present for CustomTag, not even an explicit nil
func (o *ProcessorBankTransferCreateRequest) UnsetCustomTag() {
	o.CustomTag.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorBankTransferCreateRequest) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorBankTransferCreateRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ProcessorBankTransferCreateRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ProcessorBankTransferCreateRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetOriginationAccountId returns the OriginationAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorBankTransferCreateRequest) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorBankTransferCreateRequest) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// HasOriginationAccountId returns a boolean if a field has been set.
func (o *ProcessorBankTransferCreateRequest) HasOriginationAccountId() bool {
	if o != nil && o.OriginationAccountId.IsSet() {
		return true
	}

	return false
}

// SetOriginationAccountId gets a reference to the given NullableString and assigns it to the OriginationAccountId field.
func (o *ProcessorBankTransferCreateRequest) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}
// SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil
func (o *ProcessorBankTransferCreateRequest) SetOriginationAccountIdNil() {
	o.OriginationAccountId.Set(nil)
}

// UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
func (o *ProcessorBankTransferCreateRequest) UnsetOriginationAccountId() {
	o.OriginationAccountId.Unset()
}

func (o ProcessorBankTransferCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if true {
		toSerialize["processor_token"] = o.ProcessorToken
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.AchClass != nil {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.CustomTag.IsSet() {
		toSerialize["custom_tag"] = o.CustomTag.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.OriginationAccountId.IsSet() {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorBankTransferCreateRequest struct {
	value *ProcessorBankTransferCreateRequest
	isSet bool
}

func (v NullableProcessorBankTransferCreateRequest) Get() *ProcessorBankTransferCreateRequest {
	return v.value
}

func (v *NullableProcessorBankTransferCreateRequest) Set(val *ProcessorBankTransferCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorBankTransferCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorBankTransferCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorBankTransferCreateRequest(val *ProcessorBankTransferCreateRequest) *NullableProcessorBankTransferCreateRequest {
	return &NullableProcessorBankTransferCreateRequest{value: val, isSet: true}
}

func (v NullableProcessorBankTransferCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorBankTransferCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


