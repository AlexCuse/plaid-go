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

// CreditAuditCopyTokenUpdateRequest CreditAuditCopyTokenUpdateRequest defines the request schema for `/credit/audit_copy_token/update`
type CreditAuditCopyTokenUpdateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The `audit_copy_token` you would like to update.
	AuditCopyToken string `json:"audit_copy_token"`
	// Array of tokens which the specified Audit Copy Token will be updated with. The types of token supported are asset report token and employment report token. There can be at most 1 of each token type in the array.
	ReportTokens []string `json:"report_tokens"`
}

// NewCreditAuditCopyTokenUpdateRequest instantiates a new CreditAuditCopyTokenUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditAuditCopyTokenUpdateRequest(auditCopyToken string, reportTokens []string) *CreditAuditCopyTokenUpdateRequest {
	this := CreditAuditCopyTokenUpdateRequest{}
	this.AuditCopyToken = auditCopyToken
	this.ReportTokens = reportTokens
	return &this
}

// NewCreditAuditCopyTokenUpdateRequestWithDefaults instantiates a new CreditAuditCopyTokenUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditAuditCopyTokenUpdateRequestWithDefaults() *CreditAuditCopyTokenUpdateRequest {
	this := CreditAuditCopyTokenUpdateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditAuditCopyTokenUpdateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenUpdateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditAuditCopyTokenUpdateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditAuditCopyTokenUpdateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditAuditCopyTokenUpdateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenUpdateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditAuditCopyTokenUpdateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditAuditCopyTokenUpdateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAuditCopyToken returns the AuditCopyToken field value
func (o *CreditAuditCopyTokenUpdateRequest) GetAuditCopyToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditCopyToken
}

// GetAuditCopyTokenOk returns a tuple with the AuditCopyToken field value
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenUpdateRequest) GetAuditCopyTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditCopyToken, true
}

// SetAuditCopyToken sets field value
func (o *CreditAuditCopyTokenUpdateRequest) SetAuditCopyToken(v string) {
	o.AuditCopyToken = v
}

// GetReportTokens returns the ReportTokens field value
func (o *CreditAuditCopyTokenUpdateRequest) GetReportTokens() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReportTokens
}

// GetReportTokensOk returns a tuple with the ReportTokens field value
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenUpdateRequest) GetReportTokensOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReportTokens, true
}

// SetReportTokens sets field value
func (o *CreditAuditCopyTokenUpdateRequest) SetReportTokens(v []string) {
	o.ReportTokens = v
}

func (o CreditAuditCopyTokenUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["audit_copy_token"] = o.AuditCopyToken
	}
	if true {
		toSerialize["report_tokens"] = o.ReportTokens
	}
	return json.Marshal(toSerialize)
}

type NullableCreditAuditCopyTokenUpdateRequest struct {
	value *CreditAuditCopyTokenUpdateRequest
	isSet bool
}

func (v NullableCreditAuditCopyTokenUpdateRequest) Get() *CreditAuditCopyTokenUpdateRequest {
	return v.value
}

func (v *NullableCreditAuditCopyTokenUpdateRequest) Set(val *CreditAuditCopyTokenUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditAuditCopyTokenUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditAuditCopyTokenUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditAuditCopyTokenUpdateRequest(val *CreditAuditCopyTokenUpdateRequest) *NullableCreditAuditCopyTokenUpdateRequest {
	return &NullableCreditAuditCopyTokenUpdateRequest{value: val, isSet: true}
}

func (v NullableCreditAuditCopyTokenUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditAuditCopyTokenUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


