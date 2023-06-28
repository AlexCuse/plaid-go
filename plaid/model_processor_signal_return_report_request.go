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
	"time"
)

// ProcessorSignalReturnReportRequest ProcessorSignalReturnReportRequest defines the request schema for `/processor/signal/return/report`
type ProcessorSignalReturnReportRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The processor token obtained from the Plaid integration partner. Processor tokens are in the format: `processor-<environment>-<identifier>`
	ProcessorToken string `json:"processor_token"`
	// Must be the same as the `client_transaction_id` supplied when calling `/processor/signal/evaluate`
	ClientTransactionId string `json:"client_transaction_id"`
	// Must be a valid ACH return code (e.g. \"R01\")  If formatted incorrectly, this will result in an [`INVALID_FIELD`](/docs/errors/invalid-request/#invalid_field) error.
	ReturnCode string `json:"return_code"`
	// Date and time when you receive the returns from your payment processors, in ISO 8601 format (`YYYY-MM-DDTHH:mm:ssZ`).
	ReturnedAt NullableTime `json:"returned_at,omitempty"`
}

// NewProcessorSignalReturnReportRequest instantiates a new ProcessorSignalReturnReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorSignalReturnReportRequest(processorToken string, clientTransactionId string, returnCode string) *ProcessorSignalReturnReportRequest {
	this := ProcessorSignalReturnReportRequest{}
	this.ProcessorToken = processorToken
	this.ClientTransactionId = clientTransactionId
	this.ReturnCode = returnCode
	return &this
}

// NewProcessorSignalReturnReportRequestWithDefaults instantiates a new ProcessorSignalReturnReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorSignalReturnReportRequestWithDefaults() *ProcessorSignalReturnReportRequest {
	this := ProcessorSignalReturnReportRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorSignalReturnReportRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorSignalReturnReportRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorSignalReturnReportRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorSignalReturnReportRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorSignalReturnReportRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorSignalReturnReportRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorSignalReturnReportRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorSignalReturnReportRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetProcessorToken returns the ProcessorToken field value
func (o *ProcessorSignalReturnReportRequest) GetProcessorToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessorToken
}

// GetProcessorTokenOk returns a tuple with the ProcessorToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorSignalReturnReportRequest) GetProcessorTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProcessorToken, true
}

// SetProcessorToken sets field value
func (o *ProcessorSignalReturnReportRequest) SetProcessorToken(v string) {
	o.ProcessorToken = v
}

// GetClientTransactionId returns the ClientTransactionId field value
func (o *ProcessorSignalReturnReportRequest) GetClientTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientTransactionId
}

// GetClientTransactionIdOk returns a tuple with the ClientTransactionId field value
// and a boolean to check if the value has been set.
func (o *ProcessorSignalReturnReportRequest) GetClientTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientTransactionId, true
}

// SetClientTransactionId sets field value
func (o *ProcessorSignalReturnReportRequest) SetClientTransactionId(v string) {
	o.ClientTransactionId = v
}

// GetReturnCode returns the ReturnCode field value
func (o *ProcessorSignalReturnReportRequest) GetReturnCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value
// and a boolean to check if the value has been set.
func (o *ProcessorSignalReturnReportRequest) GetReturnCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReturnCode, true
}

// SetReturnCode sets field value
func (o *ProcessorSignalReturnReportRequest) SetReturnCode(v string) {
	o.ReturnCode = v
}

// GetReturnedAt returns the ReturnedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorSignalReturnReportRequest) GetReturnedAt() time.Time {
	if o == nil || o.ReturnedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ReturnedAt.Get()
}

// GetReturnedAtOk returns a tuple with the ReturnedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorSignalReturnReportRequest) GetReturnedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReturnedAt.Get(), o.ReturnedAt.IsSet()
}

// HasReturnedAt returns a boolean if a field has been set.
func (o *ProcessorSignalReturnReportRequest) HasReturnedAt() bool {
	if o != nil && o.ReturnedAt.IsSet() {
		return true
	}

	return false
}

// SetReturnedAt gets a reference to the given NullableTime and assigns it to the ReturnedAt field.
func (o *ProcessorSignalReturnReportRequest) SetReturnedAt(v time.Time) {
	o.ReturnedAt.Set(&v)
}
// SetReturnedAtNil sets the value for ReturnedAt to be an explicit nil
func (o *ProcessorSignalReturnReportRequest) SetReturnedAtNil() {
	o.ReturnedAt.Set(nil)
}

// UnsetReturnedAt ensures that no value is present for ReturnedAt, not even an explicit nil
func (o *ProcessorSignalReturnReportRequest) UnsetReturnedAt() {
	o.ReturnedAt.Unset()
}

func (o ProcessorSignalReturnReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["processor_token"] = o.ProcessorToken
	}
	if true {
		toSerialize["client_transaction_id"] = o.ClientTransactionId
	}
	if true {
		toSerialize["return_code"] = o.ReturnCode
	}
	if o.ReturnedAt.IsSet() {
		toSerialize["returned_at"] = o.ReturnedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorSignalReturnReportRequest struct {
	value *ProcessorSignalReturnReportRequest
	isSet bool
}

func (v NullableProcessorSignalReturnReportRequest) Get() *ProcessorSignalReturnReportRequest {
	return v.value
}

func (v *NullableProcessorSignalReturnReportRequest) Set(val *ProcessorSignalReturnReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorSignalReturnReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorSignalReturnReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorSignalReturnReportRequest(val *ProcessorSignalReturnReportRequest) *NullableProcessorSignalReturnReportRequest {
	return &NullableProcessorSignalReturnReportRequest{value: val, isSet: true}
}

func (v NullableProcessorSignalReturnReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorSignalReturnReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


