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

// CreditAuditCopyTokenCreateResponse CreditAuditCopyTokenCreateResponse defines the response schema for `/credit/audit_copy_token/get`
type CreditAuditCopyTokenCreateResponse struct {
	// A token that can be shared with a third party auditor, which allows them to fetch the Asset Reports attached to the token. This token should be stored securely.
	AuditCopyToken string `json:"audit_copy_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditAuditCopyTokenCreateResponse CreditAuditCopyTokenCreateResponse

// NewCreditAuditCopyTokenCreateResponse instantiates a new CreditAuditCopyTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditAuditCopyTokenCreateResponse(auditCopyToken string, requestId string) *CreditAuditCopyTokenCreateResponse {
	this := CreditAuditCopyTokenCreateResponse{}
	this.AuditCopyToken = auditCopyToken
	this.RequestId = requestId
	return &this
}

// NewCreditAuditCopyTokenCreateResponseWithDefaults instantiates a new CreditAuditCopyTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditAuditCopyTokenCreateResponseWithDefaults() *CreditAuditCopyTokenCreateResponse {
	this := CreditAuditCopyTokenCreateResponse{}
	return &this
}

// GetAuditCopyToken returns the AuditCopyToken field value
func (o *CreditAuditCopyTokenCreateResponse) GetAuditCopyToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditCopyToken
}

// GetAuditCopyTokenOk returns a tuple with the AuditCopyToken field value
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenCreateResponse) GetAuditCopyTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditCopyToken, true
}

// SetAuditCopyToken sets field value
func (o *CreditAuditCopyTokenCreateResponse) SetAuditCopyToken(v string) {
	o.AuditCopyToken = v
}

// GetRequestId returns the RequestId field value
func (o *CreditAuditCopyTokenCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditAuditCopyTokenCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditAuditCopyTokenCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditAuditCopyTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["audit_copy_token"] = o.AuditCopyToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditAuditCopyTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditAuditCopyTokenCreateResponse := _CreditAuditCopyTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varCreditAuditCopyTokenCreateResponse); err == nil {
		*o = CreditAuditCopyTokenCreateResponse(varCreditAuditCopyTokenCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "audit_copy_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditAuditCopyTokenCreateResponse struct {
	value *CreditAuditCopyTokenCreateResponse
	isSet bool
}

func (v NullableCreditAuditCopyTokenCreateResponse) Get() *CreditAuditCopyTokenCreateResponse {
	return v.value
}

func (v *NullableCreditAuditCopyTokenCreateResponse) Set(val *CreditAuditCopyTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditAuditCopyTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditAuditCopyTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditAuditCopyTokenCreateResponse(val *CreditAuditCopyTokenCreateResponse) *NullableCreditAuditCopyTokenCreateResponse {
	return &NullableCreditAuditCopyTokenCreateResponse{value: val, isSet: true}
}

func (v NullableCreditAuditCopyTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditAuditCopyTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


