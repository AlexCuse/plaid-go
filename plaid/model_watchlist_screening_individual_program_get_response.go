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
	"time"
)

// WatchlistScreeningIndividualProgramGetResponse A program that configures the active lists, search parameters, and other behavior for initial and ongoing screening of individuals.
type WatchlistScreeningIndividualProgramGetResponse struct {
	// ID of the associated program.
	Id string `json:"id"`
	// An ISO8601 formatted timestamp.
	CreatedAt time.Time `json:"created_at"`
	// Indicator specifying whether the program is enabled and will perform daily rescans.
	IsRescanningEnabled bool `json:"is_rescanning_enabled"`
	// Watchlists enabled for the associated program
	ListsEnabled []IndividualWatchlistCode `json:"lists_enabled"`
	// A name for the program to define its purpose. For example, \"High Risk Individuals\", \"US Cardholders\", or \"Applicants\".
	Name string `json:"name"`
	NameSensitivity ProgramNameSensitivity `json:"name_sensitivity"`
	AuditTrail WatchlistScreeningAuditTrail `json:"audit_trail"`
	// Archived programs are read-only and cannot screen new customers nor participate in ongoing monitoring.
	IsArchived bool `json:"is_archived"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningIndividualProgramGetResponse WatchlistScreeningIndividualProgramGetResponse

// NewWatchlistScreeningIndividualProgramGetResponse instantiates a new WatchlistScreeningIndividualProgramGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividualProgramGetResponse(id string, createdAt time.Time, isRescanningEnabled bool, listsEnabled []IndividualWatchlistCode, name string, nameSensitivity ProgramNameSensitivity, auditTrail WatchlistScreeningAuditTrail, isArchived bool, requestId string) *WatchlistScreeningIndividualProgramGetResponse {
	this := WatchlistScreeningIndividualProgramGetResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.IsRescanningEnabled = isRescanningEnabled
	this.ListsEnabled = listsEnabled
	this.Name = name
	this.NameSensitivity = nameSensitivity
	this.AuditTrail = auditTrail
	this.IsArchived = isArchived
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningIndividualProgramGetResponseWithDefaults instantiates a new WatchlistScreeningIndividualProgramGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualProgramGetResponseWithDefaults() *WatchlistScreeningIndividualProgramGetResponse {
	this := WatchlistScreeningIndividualProgramGetResponse{}
	return &this
}

// GetId returns the Id field value
func (o *WatchlistScreeningIndividualProgramGetResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualProgramGetResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WatchlistScreeningIndividualProgramGetResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WatchlistScreeningIndividualProgramGetResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualProgramGetResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WatchlistScreeningIndividualProgramGetResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetIsRescanningEnabled returns the IsRescanningEnabled field value
func (o *WatchlistScreeningIndividualProgramGetResponse) GetIsRescanningEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRescanningEnabled
}

// GetIsRescanningEnabledOk returns a tuple with the IsRescanningEnabled field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualProgramGetResponse) GetIsRescanningEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsRescanningEnabled, true
}

// SetIsRescanningEnabled sets field value
func (o *WatchlistScreeningIndividualProgramGetResponse) SetIsRescanningEnabled(v bool) {
	o.IsRescanningEnabled = v
}

// GetListsEnabled returns the ListsEnabled field value
func (o *WatchlistScreeningIndividualProgramGetResponse) GetListsEnabled() []IndividualWatchlistCode {
	if o == nil {
		var ret []IndividualWatchlistCode
		return ret
	}

	return o.ListsEnabled
}

// GetListsEnabledOk returns a tuple with the ListsEnabled field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualProgramGetResponse) GetListsEnabledOk() (*[]IndividualWatchlistCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListsEnabled, true
}

// SetListsEnabled sets field value
func (o *WatchlistScreeningIndividualProgramGetResponse) SetListsEnabled(v []IndividualWatchlistCode) {
	o.ListsEnabled = v
}

// GetName returns the Name field value
func (o *WatchlistScreeningIndividualProgramGetResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualProgramGetResponse) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WatchlistScreeningIndividualProgramGetResponse) SetName(v string) {
	o.Name = v
}

// GetNameSensitivity returns the NameSensitivity field value
func (o *WatchlistScreeningIndividualProgramGetResponse) GetNameSensitivity() ProgramNameSensitivity {
	if o == nil {
		var ret ProgramNameSensitivity
		return ret
	}

	return o.NameSensitivity
}

// GetNameSensitivityOk returns a tuple with the NameSensitivity field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualProgramGetResponse) GetNameSensitivityOk() (*ProgramNameSensitivity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NameSensitivity, true
}

// SetNameSensitivity sets field value
func (o *WatchlistScreeningIndividualProgramGetResponse) SetNameSensitivity(v ProgramNameSensitivity) {
	o.NameSensitivity = v
}

// GetAuditTrail returns the AuditTrail field value
func (o *WatchlistScreeningIndividualProgramGetResponse) GetAuditTrail() WatchlistScreeningAuditTrail {
	if o == nil {
		var ret WatchlistScreeningAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualProgramGetResponse) GetAuditTrailOk() (*WatchlistScreeningAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *WatchlistScreeningIndividualProgramGetResponse) SetAuditTrail(v WatchlistScreeningAuditTrail) {
	o.AuditTrail = v
}

// GetIsArchived returns the IsArchived field value
func (o *WatchlistScreeningIndividualProgramGetResponse) GetIsArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualProgramGetResponse) GetIsArchivedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsArchived, true
}

// SetIsArchived sets field value
func (o *WatchlistScreeningIndividualProgramGetResponse) SetIsArchived(v bool) {
	o.IsArchived = v
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningIndividualProgramGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualProgramGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningIndividualProgramGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningIndividualProgramGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["is_rescanning_enabled"] = o.IsRescanningEnabled
	}
	if true {
		toSerialize["lists_enabled"] = o.ListsEnabled
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["name_sensitivity"] = o.NameSensitivity
	}
	if true {
		toSerialize["audit_trail"] = o.AuditTrail
	}
	if true {
		toSerialize["is_archived"] = o.IsArchived
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningIndividualProgramGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningIndividualProgramGetResponse := _WatchlistScreeningIndividualProgramGetResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningIndividualProgramGetResponse); err == nil {
		*o = WatchlistScreeningIndividualProgramGetResponse(varWatchlistScreeningIndividualProgramGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "is_rescanning_enabled")
		delete(additionalProperties, "lists_enabled")
		delete(additionalProperties, "name")
		delete(additionalProperties, "name_sensitivity")
		delete(additionalProperties, "audit_trail")
		delete(additionalProperties, "is_archived")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningIndividualProgramGetResponse struct {
	value *WatchlistScreeningIndividualProgramGetResponse
	isSet bool
}

func (v NullableWatchlistScreeningIndividualProgramGetResponse) Get() *WatchlistScreeningIndividualProgramGetResponse {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualProgramGetResponse) Set(val *WatchlistScreeningIndividualProgramGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualProgramGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualProgramGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualProgramGetResponse(val *WatchlistScreeningIndividualProgramGetResponse) *NullableWatchlistScreeningIndividualProgramGetResponse {
	return &NullableWatchlistScreeningIndividualProgramGetResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualProgramGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualProgramGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


