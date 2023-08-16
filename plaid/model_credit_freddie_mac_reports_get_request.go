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

// CreditFreddieMacReportsGetRequest CreditFreddieMacReportsGetRequest defines the request schema for `credit/asset_report/freddie_mac/get`
type CreditFreddieMacReportsGetRequest struct {
	// A token that can be shared with a third party auditor to allow them to obtain access to the Asset Report. This token should be stored securely.
	AuditCopyToken string `json:"audit_copy_token"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
}

// NewCreditFreddieMacReportsGetRequest instantiates a new CreditFreddieMacReportsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacReportsGetRequest(auditCopyToken string) *CreditFreddieMacReportsGetRequest {
	this := CreditFreddieMacReportsGetRequest{}
	this.AuditCopyToken = auditCopyToken
	return &this
}

// NewCreditFreddieMacReportsGetRequestWithDefaults instantiates a new CreditFreddieMacReportsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacReportsGetRequestWithDefaults() *CreditFreddieMacReportsGetRequest {
	this := CreditFreddieMacReportsGetRequest{}
	return &this
}

// GetAuditCopyToken returns the AuditCopyToken field value
func (o *CreditFreddieMacReportsGetRequest) GetAuditCopyToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditCopyToken
}

// GetAuditCopyTokenOk returns a tuple with the AuditCopyToken field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacReportsGetRequest) GetAuditCopyTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditCopyToken, true
}

// SetAuditCopyToken sets field value
func (o *CreditFreddieMacReportsGetRequest) SetAuditCopyToken(v string) {
	o.AuditCopyToken = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditFreddieMacReportsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacReportsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditFreddieMacReportsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditFreddieMacReportsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditFreddieMacReportsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacReportsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditFreddieMacReportsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditFreddieMacReportsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o CreditFreddieMacReportsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["audit_copy_token"] = o.AuditCopyToken
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableCreditFreddieMacReportsGetRequest struct {
	value *CreditFreddieMacReportsGetRequest
	isSet bool
}

func (v NullableCreditFreddieMacReportsGetRequest) Get() *CreditFreddieMacReportsGetRequest {
	return v.value
}

func (v *NullableCreditFreddieMacReportsGetRequest) Set(val *CreditFreddieMacReportsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacReportsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacReportsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacReportsGetRequest(val *CreditFreddieMacReportsGetRequest) *NullableCreditFreddieMacReportsGetRequest {
	return &NullableCreditFreddieMacReportsGetRequest{value: val, isSet: true}
}

func (v NullableCreditFreddieMacReportsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacReportsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


