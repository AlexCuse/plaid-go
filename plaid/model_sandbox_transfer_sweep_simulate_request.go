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

// SandboxTransferSweepSimulateRequest Defines the request schema for `/sandbox/transfer/sweep/simulate`
type SandboxTransferSweepSimulateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Plaid’s unique identifier for a test clock. If provided, the sweep to be simulated is created on the day of the `virtual_time` on the `test_clock`. If the date of `virtual_time` is on weekend or a federal holiday, the next available banking day is used.
	TestClockId NullableString `json:"test_clock_id,omitempty"`
}

// NewSandboxTransferSweepSimulateRequest instantiates a new SandboxTransferSweepSimulateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferSweepSimulateRequest() *SandboxTransferSweepSimulateRequest {
	this := SandboxTransferSweepSimulateRequest{}
	return &this
}

// NewSandboxTransferSweepSimulateRequestWithDefaults instantiates a new SandboxTransferSweepSimulateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferSweepSimulateRequestWithDefaults() *SandboxTransferSweepSimulateRequest {
	this := SandboxTransferSweepSimulateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxTransferSweepSimulateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferSweepSimulateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxTransferSweepSimulateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxTransferSweepSimulateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxTransferSweepSimulateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferSweepSimulateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxTransferSweepSimulateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxTransferSweepSimulateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetTestClockId returns the TestClockId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxTransferSweepSimulateRequest) GetTestClockId() string {
	if o == nil || o.TestClockId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TestClockId.Get()
}

// GetTestClockIdOk returns a tuple with the TestClockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxTransferSweepSimulateRequest) GetTestClockIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TestClockId.Get(), o.TestClockId.IsSet()
}

// HasTestClockId returns a boolean if a field has been set.
func (o *SandboxTransferSweepSimulateRequest) HasTestClockId() bool {
	if o != nil && o.TestClockId.IsSet() {
		return true
	}

	return false
}

// SetTestClockId gets a reference to the given NullableString and assigns it to the TestClockId field.
func (o *SandboxTransferSweepSimulateRequest) SetTestClockId(v string) {
	o.TestClockId.Set(&v)
}
// SetTestClockIdNil sets the value for TestClockId to be an explicit nil
func (o *SandboxTransferSweepSimulateRequest) SetTestClockIdNil() {
	o.TestClockId.Set(nil)
}

// UnsetTestClockId ensures that no value is present for TestClockId, not even an explicit nil
func (o *SandboxTransferSweepSimulateRequest) UnsetTestClockId() {
	o.TestClockId.Unset()
}

func (o SandboxTransferSweepSimulateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.TestClockId.IsSet() {
		toSerialize["test_clock_id"] = o.TestClockId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxTransferSweepSimulateRequest struct {
	value *SandboxTransferSweepSimulateRequest
	isSet bool
}

func (v NullableSandboxTransferSweepSimulateRequest) Get() *SandboxTransferSweepSimulateRequest {
	return v.value
}

func (v *NullableSandboxTransferSweepSimulateRequest) Set(val *SandboxTransferSweepSimulateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferSweepSimulateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferSweepSimulateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferSweepSimulateRequest(val *SandboxTransferSweepSimulateRequest) *NullableSandboxTransferSweepSimulateRequest {
	return &NullableSandboxTransferSweepSimulateRequest{value: val, isSet: true}
}

func (v NullableSandboxTransferSweepSimulateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferSweepSimulateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


