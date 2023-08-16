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

// FDXHateoasLink REST application constraint (Hypermedia As The Engine Of Application State)
type FDXHateoasLink struct {
	// URL to invoke the action on the resource
	Href string `json:"href"`
	Action *FDXHateoasLinkAction `json:"action,omitempty"`
	// Relation of this link to its containing entity, as defined by and with many example relation values at [IETF RFC5988](https://datatracker.ietf.org/doc/html/rfc5988)
	Rel *string `json:"rel,omitempty"`
	// Content-types that can be used in the Accept header
	Types *[]FDXContentTypes `json:"types,omitempty"`
}

// NewFDXHateoasLink instantiates a new FDXHateoasLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFDXHateoasLink(href string) *FDXHateoasLink {
	this := FDXHateoasLink{}
	this.Href = href
	return &this
}

// NewFDXHateoasLinkWithDefaults instantiates a new FDXHateoasLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFDXHateoasLinkWithDefaults() *FDXHateoasLink {
	this := FDXHateoasLink{}
	return &this
}

// GetHref returns the Href field value
func (o *FDXHateoasLink) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *FDXHateoasLink) GetHrefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *FDXHateoasLink) SetHref(v string) {
	o.Href = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *FDXHateoasLink) GetAction() FDXHateoasLinkAction {
	if o == nil || o.Action == nil {
		var ret FDXHateoasLinkAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXHateoasLink) GetActionOk() (*FDXHateoasLinkAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *FDXHateoasLink) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given FDXHateoasLinkAction and assigns it to the Action field.
func (o *FDXHateoasLink) SetAction(v FDXHateoasLinkAction) {
	o.Action = &v
}

// GetRel returns the Rel field value if set, zero value otherwise.
func (o *FDXHateoasLink) GetRel() string {
	if o == nil || o.Rel == nil {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXHateoasLink) GetRelOk() (*string, bool) {
	if o == nil || o.Rel == nil {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *FDXHateoasLink) HasRel() bool {
	if o != nil && o.Rel != nil {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *FDXHateoasLink) SetRel(v string) {
	o.Rel = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *FDXHateoasLink) GetTypes() []FDXContentTypes {
	if o == nil || o.Types == nil {
		var ret []FDXContentTypes
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXHateoasLink) GetTypesOk() (*[]FDXContentTypes, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *FDXHateoasLink) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []FDXContentTypes and assigns it to the Types field.
func (o *FDXHateoasLink) SetTypes(v []FDXContentTypes) {
	o.Types = &v
}

func (o FDXHateoasLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["href"] = o.Href
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Rel != nil {
		toSerialize["rel"] = o.Rel
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullableFDXHateoasLink struct {
	value *FDXHateoasLink
	isSet bool
}

func (v NullableFDXHateoasLink) Get() *FDXHateoasLink {
	return v.value
}

func (v *NullableFDXHateoasLink) Set(val *FDXHateoasLink) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXHateoasLink) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXHateoasLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXHateoasLink(val *FDXHateoasLink) *NullableFDXHateoasLink {
	return &NullableFDXHateoasLink{value: val, isSet: true}
}

func (v NullableFDXHateoasLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXHateoasLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

