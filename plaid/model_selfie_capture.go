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

// SelfieCapture The image or video capture of a selfie. Only one of image or video URL will be populated per selfie.
type SelfieCapture struct {
	// Temporary URL for downloading an image selfie capture.
	ImageUrl NullableString `json:"image_url,omitempty"`
	// Temporary URL for downloading a video selfie capture.
	VideoUrl NullableString `json:"video_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfieCapture SelfieCapture

// NewSelfieCapture instantiates a new SelfieCapture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfieCapture() *SelfieCapture {
	this := SelfieCapture{}
	return &this
}

// NewSelfieCaptureWithDefaults instantiates a new SelfieCapture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfieCaptureWithDefaults() *SelfieCapture {
	this := SelfieCapture{}
	return &this
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelfieCapture) GetImageUrl() string {
	if o == nil || o.ImageUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelfieCapture) GetImageUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *SelfieCapture) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *SelfieCapture) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *SelfieCapture) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *SelfieCapture) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelfieCapture) GetVideoUrl() string {
	if o == nil || o.VideoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.VideoUrl.Get()
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelfieCapture) GetVideoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VideoUrl.Get(), o.VideoUrl.IsSet()
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *SelfieCapture) HasVideoUrl() bool {
	if o != nil && o.VideoUrl.IsSet() {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given NullableString and assigns it to the VideoUrl field.
func (o *SelfieCapture) SetVideoUrl(v string) {
	o.VideoUrl.Set(&v)
}
// SetVideoUrlNil sets the value for VideoUrl to be an explicit nil
func (o *SelfieCapture) SetVideoUrlNil() {
	o.VideoUrl.Set(nil)
}

// UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
func (o *SelfieCapture) UnsetVideoUrl() {
	o.VideoUrl.Unset()
}

func (o SelfieCapture) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	if o.VideoUrl.IsSet() {
		toSerialize["video_url"] = o.VideoUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SelfieCapture) UnmarshalJSON(bytes []byte) (err error) {
	varSelfieCapture := _SelfieCapture{}

	if err = json.Unmarshal(bytes, &varSelfieCapture); err == nil {
		*o = SelfieCapture(varSelfieCapture)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "video_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfieCapture struct {
	value *SelfieCapture
	isSet bool
}

func (v NullableSelfieCapture) Get() *SelfieCapture {
	return v.value
}

func (v *NullableSelfieCapture) Set(val *SelfieCapture) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfieCapture) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfieCapture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfieCapture(val *SelfieCapture) *NullableSelfieCapture {
	return &NullableSelfieCapture{value: val, isSet: true}
}

func (v NullableSelfieCapture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfieCapture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


