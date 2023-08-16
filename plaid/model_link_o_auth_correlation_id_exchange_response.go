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

// LinkOAuthCorrelationIdExchangeResponse LinkOAuthCorrelationIdExchangeResponse defines the response schema for `/link/oauth/correlation_id/exchange`
type LinkOAuthCorrelationIdExchangeResponse struct {
	// The `link_token` associated to the given `link_correlation_id`, which can be used to re-initialize Link.
	LinkToken string `json:"link_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _LinkOAuthCorrelationIdExchangeResponse LinkOAuthCorrelationIdExchangeResponse

// NewLinkOAuthCorrelationIdExchangeResponse instantiates a new LinkOAuthCorrelationIdExchangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkOAuthCorrelationIdExchangeResponse(linkToken string, requestId string) *LinkOAuthCorrelationIdExchangeResponse {
	this := LinkOAuthCorrelationIdExchangeResponse{}
	this.LinkToken = linkToken
	this.RequestId = requestId
	return &this
}

// NewLinkOAuthCorrelationIdExchangeResponseWithDefaults instantiates a new LinkOAuthCorrelationIdExchangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkOAuthCorrelationIdExchangeResponseWithDefaults() *LinkOAuthCorrelationIdExchangeResponse {
	this := LinkOAuthCorrelationIdExchangeResponse{}
	return &this
}

// GetLinkToken returns the LinkToken field value
func (o *LinkOAuthCorrelationIdExchangeResponse) GetLinkToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkToken
}

// GetLinkTokenOk returns a tuple with the LinkToken field value
// and a boolean to check if the value has been set.
func (o *LinkOAuthCorrelationIdExchangeResponse) GetLinkTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkToken, true
}

// SetLinkToken sets field value
func (o *LinkOAuthCorrelationIdExchangeResponse) SetLinkToken(v string) {
	o.LinkToken = v
}

// GetRequestId returns the RequestId field value
func (o *LinkOAuthCorrelationIdExchangeResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *LinkOAuthCorrelationIdExchangeResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *LinkOAuthCorrelationIdExchangeResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o LinkOAuthCorrelationIdExchangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["link_token"] = o.LinkToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkOAuthCorrelationIdExchangeResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLinkOAuthCorrelationIdExchangeResponse := _LinkOAuthCorrelationIdExchangeResponse{}

	if err = json.Unmarshal(bytes, &varLinkOAuthCorrelationIdExchangeResponse); err == nil {
		*o = LinkOAuthCorrelationIdExchangeResponse(varLinkOAuthCorrelationIdExchangeResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "link_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkOAuthCorrelationIdExchangeResponse struct {
	value *LinkOAuthCorrelationIdExchangeResponse
	isSet bool
}

func (v NullableLinkOAuthCorrelationIdExchangeResponse) Get() *LinkOAuthCorrelationIdExchangeResponse {
	return v.value
}

func (v *NullableLinkOAuthCorrelationIdExchangeResponse) Set(val *LinkOAuthCorrelationIdExchangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkOAuthCorrelationIdExchangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkOAuthCorrelationIdExchangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkOAuthCorrelationIdExchangeResponse(val *LinkOAuthCorrelationIdExchangeResponse) *NullableLinkOAuthCorrelationIdExchangeResponse {
	return &NullableLinkOAuthCorrelationIdExchangeResponse{value: val, isSet: true}
}

func (v NullableLinkOAuthCorrelationIdExchangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkOAuthCorrelationIdExchangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


