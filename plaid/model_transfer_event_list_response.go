/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferEventListResponse Defines the response schema for `/transfer/event/list`
type TransferEventListResponse struct {
	TransferEvents []TransferEvent `json:"transfer_events"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferEventListResponse TransferEventListResponse

// NewTransferEventListResponse instantiates a new TransferEventListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferEventListResponse(transferEvents []TransferEvent, requestId string) *TransferEventListResponse {
	this := TransferEventListResponse{}
	this.TransferEvents = transferEvents
	this.RequestId = requestId
	return &this
}

// NewTransferEventListResponseWithDefaults instantiates a new TransferEventListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferEventListResponseWithDefaults() *TransferEventListResponse {
	this := TransferEventListResponse{}
	return &this
}

// GetTransferEvents returns the TransferEvents field value
func (o *TransferEventListResponse) GetTransferEvents() []TransferEvent {
	if o == nil {
		var ret []TransferEvent
		return ret
	}

	return o.TransferEvents
}

// GetTransferEventsOk returns a tuple with the TransferEvents field value
// and a boolean to check if the value has been set.
func (o *TransferEventListResponse) GetTransferEventsOk() (*[]TransferEvent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferEvents, true
}

// SetTransferEvents sets field value
func (o *TransferEventListResponse) SetTransferEvents(v []TransferEvent) {
	o.TransferEvents = v
}

// GetRequestId returns the RequestId field value
func (o *TransferEventListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferEventListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferEventListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferEventListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transfer_events"] = o.TransferEvents
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferEventListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferEventListResponse := _TransferEventListResponse{}

	if err = json.Unmarshal(bytes, &varTransferEventListResponse); err == nil {
		*o = TransferEventListResponse(varTransferEventListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transfer_events")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferEventListResponse struct {
	value *TransferEventListResponse
	isSet bool
}

func (v NullableTransferEventListResponse) Get() *TransferEventListResponse {
	return v.value
}

func (v *NullableTransferEventListResponse) Set(val *TransferEventListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEventListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEventListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEventListResponse(val *TransferEventListResponse) *NullableTransferEventListResponse {
	return &NullableTransferEventListResponse{value: val, isSet: true}
}

func (v NullableTransferEventListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEventListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


