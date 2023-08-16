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
	"time"
)

// TransactionsRecurringGetResponse TransactionsRecurringGetResponse defines the response schema for `/transactions/recurring/get`
type TransactionsRecurringGetResponse struct {
	// An array of depository transaction streams.
	InflowStreams []TransactionStream `json:"inflow_streams"`
	// An array of expense transaction streams.
	OutflowStreams []TransactionStream `json:"outflow_streams"`
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the last time transaction streams for the given account were updated on
	UpdatedDatetime time.Time `json:"updated_datetime"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransactionsRecurringGetResponse TransactionsRecurringGetResponse

// NewTransactionsRecurringGetResponse instantiates a new TransactionsRecurringGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRecurringGetResponse(inflowStreams []TransactionStream, outflowStreams []TransactionStream, updatedDatetime time.Time, requestId string) *TransactionsRecurringGetResponse {
	this := TransactionsRecurringGetResponse{}
	this.InflowStreams = inflowStreams
	this.OutflowStreams = outflowStreams
	this.UpdatedDatetime = updatedDatetime
	this.RequestId = requestId
	return &this
}

// NewTransactionsRecurringGetResponseWithDefaults instantiates a new TransactionsRecurringGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRecurringGetResponseWithDefaults() *TransactionsRecurringGetResponse {
	this := TransactionsRecurringGetResponse{}
	return &this
}

// GetInflowStreams returns the InflowStreams field value
func (o *TransactionsRecurringGetResponse) GetInflowStreams() []TransactionStream {
	if o == nil {
		var ret []TransactionStream
		return ret
	}

	return o.InflowStreams
}

// GetInflowStreamsOk returns a tuple with the InflowStreams field value
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringGetResponse) GetInflowStreamsOk() (*[]TransactionStream, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InflowStreams, true
}

// SetInflowStreams sets field value
func (o *TransactionsRecurringGetResponse) SetInflowStreams(v []TransactionStream) {
	o.InflowStreams = v
}

// GetOutflowStreams returns the OutflowStreams field value
func (o *TransactionsRecurringGetResponse) GetOutflowStreams() []TransactionStream {
	if o == nil {
		var ret []TransactionStream
		return ret
	}

	return o.OutflowStreams
}

// GetOutflowStreamsOk returns a tuple with the OutflowStreams field value
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringGetResponse) GetOutflowStreamsOk() (*[]TransactionStream, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OutflowStreams, true
}

// SetOutflowStreams sets field value
func (o *TransactionsRecurringGetResponse) SetOutflowStreams(v []TransactionStream) {
	o.OutflowStreams = v
}

// GetUpdatedDatetime returns the UpdatedDatetime field value
func (o *TransactionsRecurringGetResponse) GetUpdatedDatetime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedDatetime
}

// GetUpdatedDatetimeOk returns a tuple with the UpdatedDatetime field value
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringGetResponse) GetUpdatedDatetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedDatetime, true
}

// SetUpdatedDatetime sets field value
func (o *TransactionsRecurringGetResponse) SetUpdatedDatetime(v time.Time) {
	o.UpdatedDatetime = v
}

// GetRequestId returns the RequestId field value
func (o *TransactionsRecurringGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransactionsRecurringGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransactionsRecurringGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inflow_streams"] = o.InflowStreams
	}
	if true {
		toSerialize["outflow_streams"] = o.OutflowStreams
	}
	if true {
		toSerialize["updated_datetime"] = o.UpdatedDatetime
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionsRecurringGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionsRecurringGetResponse := _TransactionsRecurringGetResponse{}

	if err = json.Unmarshal(bytes, &varTransactionsRecurringGetResponse); err == nil {
		*o = TransactionsRecurringGetResponse(varTransactionsRecurringGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "inflow_streams")
		delete(additionalProperties, "outflow_streams")
		delete(additionalProperties, "updated_datetime")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionsRecurringGetResponse struct {
	value *TransactionsRecurringGetResponse
	isSet bool
}

func (v NullableTransactionsRecurringGetResponse) Get() *TransactionsRecurringGetResponse {
	return v.value
}

func (v *NullableTransactionsRecurringGetResponse) Set(val *TransactionsRecurringGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRecurringGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRecurringGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRecurringGetResponse(val *TransactionsRecurringGetResponse) *NullableTransactionsRecurringGetResponse {
	return &NullableTransactionsRecurringGetResponse{value: val, isSet: true}
}

func (v NullableTransactionsRecurringGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRecurringGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


