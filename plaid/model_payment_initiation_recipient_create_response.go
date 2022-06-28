/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.131.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentInitiationRecipientCreateResponse PaymentInitiationRecipientCreateResponse defines the response schema for `/payment_initation/recipient/create`
type PaymentInitiationRecipientCreateResponse struct {
	// A unique ID identifying the recipient
	RecipientId string `json:"recipient_id"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationRecipientCreateResponse PaymentInitiationRecipientCreateResponse

// NewPaymentInitiationRecipientCreateResponse instantiates a new PaymentInitiationRecipientCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationRecipientCreateResponse(recipientId string, requestId string) *PaymentInitiationRecipientCreateResponse {
	this := PaymentInitiationRecipientCreateResponse{}
	this.RecipientId = recipientId
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationRecipientCreateResponseWithDefaults instantiates a new PaymentInitiationRecipientCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationRecipientCreateResponseWithDefaults() *PaymentInitiationRecipientCreateResponse {
	this := PaymentInitiationRecipientCreateResponse{}
	return &this
}

// GetRecipientId returns the RecipientId field value
func (o *PaymentInitiationRecipientCreateResponse) GetRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientCreateResponse) GetRecipientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientId, true
}

// SetRecipientId sets field value
func (o *PaymentInitiationRecipientCreateResponse) SetRecipientId(v string) {
	o.RecipientId = v
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationRecipientCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationRecipientCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationRecipientCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipient_id"] = o.RecipientId
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationRecipientCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationRecipientCreateResponse := _PaymentInitiationRecipientCreateResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationRecipientCreateResponse); err == nil {
		*o = PaymentInitiationRecipientCreateResponse(varPaymentInitiationRecipientCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "recipient_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationRecipientCreateResponse struct {
	value *PaymentInitiationRecipientCreateResponse
	isSet bool
}

func (v NullablePaymentInitiationRecipientCreateResponse) Get() *PaymentInitiationRecipientCreateResponse {
	return v.value
}

func (v *NullablePaymentInitiationRecipientCreateResponse) Set(val *PaymentInitiationRecipientCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRecipientCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRecipientCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRecipientCreateResponse(val *PaymentInitiationRecipientCreateResponse) *NullablePaymentInitiationRecipientCreateResponse {
	return &NullablePaymentInitiationRecipientCreateResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationRecipientCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRecipientCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


