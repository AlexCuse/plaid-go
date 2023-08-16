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

// TransferOriginatorCreateResponse Defines the response schema for `/transfer/originator/create`
type TransferOriginatorCreateResponse struct {
	// Client ID of the originator. This identifier will be used when creating transfers and should be stored associated with end user information.
	OriginatorClientId string `json:"originator_client_id"`
	// The company name of the end customer.
	CompanyName string `json:"company_name"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferOriginatorCreateResponse TransferOriginatorCreateResponse

// NewTransferOriginatorCreateResponse instantiates a new TransferOriginatorCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOriginatorCreateResponse(originatorClientId string, companyName string, requestId string) *TransferOriginatorCreateResponse {
	this := TransferOriginatorCreateResponse{}
	this.OriginatorClientId = originatorClientId
	this.CompanyName = companyName
	this.RequestId = requestId
	return &this
}

// NewTransferOriginatorCreateResponseWithDefaults instantiates a new TransferOriginatorCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOriginatorCreateResponseWithDefaults() *TransferOriginatorCreateResponse {
	this := TransferOriginatorCreateResponse{}
	return &this
}

// GetOriginatorClientId returns the OriginatorClientId field value
func (o *TransferOriginatorCreateResponse) GetOriginatorClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatorClientId
}

// GetOriginatorClientIdOk returns a tuple with the OriginatorClientId field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorCreateResponse) GetOriginatorClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginatorClientId, true
}

// SetOriginatorClientId sets field value
func (o *TransferOriginatorCreateResponse) SetOriginatorClientId(v string) {
	o.OriginatorClientId = v
}

// GetCompanyName returns the CompanyName field value
func (o *TransferOriginatorCreateResponse) GetCompanyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorCreateResponse) GetCompanyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyName, true
}

// SetCompanyName sets field value
func (o *TransferOriginatorCreateResponse) SetCompanyName(v string) {
	o.CompanyName = v
}

// GetRequestId returns the RequestId field value
func (o *TransferOriginatorCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferOriginatorCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferOriginatorCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["originator_client_id"] = o.OriginatorClientId
	}
	if true {
		toSerialize["company_name"] = o.CompanyName
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferOriginatorCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferOriginatorCreateResponse := _TransferOriginatorCreateResponse{}

	if err = json.Unmarshal(bytes, &varTransferOriginatorCreateResponse); err == nil {
		*o = TransferOriginatorCreateResponse(varTransferOriginatorCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "originator_client_id")
		delete(additionalProperties, "company_name")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferOriginatorCreateResponse struct {
	value *TransferOriginatorCreateResponse
	isSet bool
}

func (v NullableTransferOriginatorCreateResponse) Get() *TransferOriginatorCreateResponse {
	return v.value
}

func (v *NullableTransferOriginatorCreateResponse) Set(val *TransferOriginatorCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOriginatorCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOriginatorCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOriginatorCreateResponse(val *TransferOriginatorCreateResponse) *NullableTransferOriginatorCreateResponse {
	return &NullableTransferOriginatorCreateResponse{value: val, isSet: true}
}

func (v NullableTransferOriginatorCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOriginatorCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


