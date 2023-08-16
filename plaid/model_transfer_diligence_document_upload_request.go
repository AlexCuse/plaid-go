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
	"os"
)

// TransferDiligenceDocumentUploadRequest Defines the request schema for `/transfer/diligence/document/upload`
type TransferDiligenceDocumentUploadRequest struct {
	// The Client ID of the originator whose document that you want to upload.
	OriginatorClientId string `json:"originator_client_id"`
	// A file to upload.
	File *os.File `json:"file"`
	Purpose TransferDocumentPurpose `json:"purpose"`
}

// NewTransferDiligenceDocumentUploadRequest instantiates a new TransferDiligenceDocumentUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferDiligenceDocumentUploadRequest(originatorClientId string, file *os.File, purpose TransferDocumentPurpose) *TransferDiligenceDocumentUploadRequest {
	this := TransferDiligenceDocumentUploadRequest{}
	this.OriginatorClientId = originatorClientId
	this.File = file
	this.Purpose = purpose
	return &this
}

// NewTransferDiligenceDocumentUploadRequestWithDefaults instantiates a new TransferDiligenceDocumentUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferDiligenceDocumentUploadRequestWithDefaults() *TransferDiligenceDocumentUploadRequest {
	this := TransferDiligenceDocumentUploadRequest{}
	return &this
}

// GetOriginatorClientId returns the OriginatorClientId field value
func (o *TransferDiligenceDocumentUploadRequest) GetOriginatorClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatorClientId
}

// GetOriginatorClientIdOk returns a tuple with the OriginatorClientId field value
// and a boolean to check if the value has been set.
func (o *TransferDiligenceDocumentUploadRequest) GetOriginatorClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginatorClientId, true
}

// SetOriginatorClientId sets field value
func (o *TransferDiligenceDocumentUploadRequest) SetOriginatorClientId(v string) {
	o.OriginatorClientId = v
}

// GetFile returns the File field value
func (o *TransferDiligenceDocumentUploadRequest) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *TransferDiligenceDocumentUploadRequest) GetFileOk() (**os.File, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *TransferDiligenceDocumentUploadRequest) SetFile(v *os.File) {
	o.File = v
}

// GetPurpose returns the Purpose field value
func (o *TransferDiligenceDocumentUploadRequest) GetPurpose() TransferDocumentPurpose {
	if o == nil {
		var ret TransferDocumentPurpose
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *TransferDiligenceDocumentUploadRequest) GetPurposeOk() (*TransferDocumentPurpose, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *TransferDiligenceDocumentUploadRequest) SetPurpose(v TransferDocumentPurpose) {
	o.Purpose = v
}

func (o TransferDiligenceDocumentUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["originator_client_id"] = o.OriginatorClientId
	}
	if true {
		toSerialize["file"] = o.File
	}
	if true {
		toSerialize["purpose"] = o.Purpose
	}
	return json.Marshal(toSerialize)
}

type NullableTransferDiligenceDocumentUploadRequest struct {
	value *TransferDiligenceDocumentUploadRequest
	isSet bool
}

func (v NullableTransferDiligenceDocumentUploadRequest) Get() *TransferDiligenceDocumentUploadRequest {
	return v.value
}

func (v *NullableTransferDiligenceDocumentUploadRequest) Set(val *TransferDiligenceDocumentUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDiligenceDocumentUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDiligenceDocumentUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDiligenceDocumentUploadRequest(val *TransferDiligenceDocumentUploadRequest) *NullableTransferDiligenceDocumentUploadRequest {
	return &NullableTransferDiligenceDocumentUploadRequest{value: val, isSet: true}
}

func (v NullableTransferDiligenceDocumentUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDiligenceDocumentUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


