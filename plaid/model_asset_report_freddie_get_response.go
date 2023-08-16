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

// AssetReportFreddieGetResponse AssetReportFreddieGetResponse defines the response schema for `/asset_report/get`
type AssetReportFreddieGetResponse struct {
	DEAL AssetReportFreddie `json:"DEAL"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	// The Verification Of Assets (aka VOA or Freddie Mac Schema) schema version.
	SchemaVersion float32 `json:"SchemaVersion"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportFreddieGetResponse AssetReportFreddieGetResponse

// NewAssetReportFreddieGetResponse instantiates a new AssetReportFreddieGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportFreddieGetResponse(dEAL AssetReportFreddie, requestId string, schemaVersion float32) *AssetReportFreddieGetResponse {
	this := AssetReportFreddieGetResponse{}
	this.DEAL = dEAL
	this.RequestId = requestId
	this.SchemaVersion = schemaVersion
	return &this
}

// NewAssetReportFreddieGetResponseWithDefaults instantiates a new AssetReportFreddieGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportFreddieGetResponseWithDefaults() *AssetReportFreddieGetResponse {
	this := AssetReportFreddieGetResponse{}
	return &this
}

// GetDEAL returns the DEAL field value
func (o *AssetReportFreddieGetResponse) GetDEAL() AssetReportFreddie {
	if o == nil {
		var ret AssetReportFreddie
		return ret
	}

	return o.DEAL
}

// GetDEALOk returns a tuple with the DEAL field value
// and a boolean to check if the value has been set.
func (o *AssetReportFreddieGetResponse) GetDEALOk() (*AssetReportFreddie, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DEAL, true
}

// SetDEAL sets field value
func (o *AssetReportFreddieGetResponse) SetDEAL(v AssetReportFreddie) {
	o.DEAL = v
}

// GetRequestId returns the RequestId field value
func (o *AssetReportFreddieGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AssetReportFreddieGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AssetReportFreddieGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *AssetReportFreddieGetResponse) GetSchemaVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *AssetReportFreddieGetResponse) GetSchemaVersionOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *AssetReportFreddieGetResponse) SetSchemaVersion(v float32) {
	o.SchemaVersion = v
}

func (o AssetReportFreddieGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DEAL"] = o.DEAL
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["SchemaVersion"] = o.SchemaVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportFreddieGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportFreddieGetResponse := _AssetReportFreddieGetResponse{}

	if err = json.Unmarshal(bytes, &varAssetReportFreddieGetResponse); err == nil {
		*o = AssetReportFreddieGetResponse(varAssetReportFreddieGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DEAL")
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "SchemaVersion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportFreddieGetResponse struct {
	value *AssetReportFreddieGetResponse
	isSet bool
}

func (v NullableAssetReportFreddieGetResponse) Get() *AssetReportFreddieGetResponse {
	return v.value
}

func (v *NullableAssetReportFreddieGetResponse) Set(val *AssetReportFreddieGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportFreddieGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportFreddieGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportFreddieGetResponse(val *AssetReportFreddieGetResponse) *NullableAssetReportFreddieGetResponse {
	return &NullableAssetReportFreddieGetResponse{value: val, isSet: true}
}

func (v NullableAssetReportFreddieGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportFreddieGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


