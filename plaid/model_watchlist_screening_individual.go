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

// WatchlistScreeningIndividual The screening object allows you to represent a customer in your system, update their profile, and search for them on various watchlists. Note: Rejected customers will not receive new hits, regardless of program configuration.
type WatchlistScreeningIndividual struct {
	// ID of the associated screening.
	Id string `json:"id"`
	SearchTerms WatchlistScreeningSearchTerms `json:"search_terms"`
	Assignee NullableString `json:"assignee"`
	Status WatchlistScreeningStatus `json:"status"`
	ClientUserId NullableString `json:"client_user_id"`
	AuditTrail WatchlistScreeningAuditTrail `json:"audit_trail"`
}

// NewWatchlistScreeningIndividual instantiates a new WatchlistScreeningIndividual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividual(id string, searchTerms WatchlistScreeningSearchTerms, assignee NullableString, status WatchlistScreeningStatus, clientUserId NullableString, auditTrail WatchlistScreeningAuditTrail) *WatchlistScreeningIndividual {
	this := WatchlistScreeningIndividual{}
	this.Id = id
	this.SearchTerms = searchTerms
	this.Assignee = assignee
	this.Status = status
	this.ClientUserId = clientUserId
	this.AuditTrail = auditTrail
	return &this
}

// NewWatchlistScreeningIndividualWithDefaults instantiates a new WatchlistScreeningIndividual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualWithDefaults() *WatchlistScreeningIndividual {
	this := WatchlistScreeningIndividual{}
	return &this
}

// GetId returns the Id field value
func (o *WatchlistScreeningIndividual) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividual) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WatchlistScreeningIndividual) SetId(v string) {
	o.Id = v
}

// GetSearchTerms returns the SearchTerms field value
func (o *WatchlistScreeningIndividual) GetSearchTerms() WatchlistScreeningSearchTerms {
	if o == nil {
		var ret WatchlistScreeningSearchTerms
		return ret
	}

	return o.SearchTerms
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividual) GetSearchTermsOk() (*WatchlistScreeningSearchTerms, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchTerms, true
}

// SetSearchTerms sets field value
func (o *WatchlistScreeningIndividual) SetSearchTerms(v WatchlistScreeningSearchTerms) {
	o.SearchTerms = v
}

// GetAssignee returns the Assignee field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningIndividual) GetAssignee() string {
	if o == nil || o.Assignee.Get() == nil {
		var ret string
		return ret
	}

	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividual) GetAssigneeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// SetAssignee sets field value
func (o *WatchlistScreeningIndividual) SetAssignee(v string) {
	o.Assignee.Set(&v)
}

// GetStatus returns the Status field value
func (o *WatchlistScreeningIndividual) GetStatus() WatchlistScreeningStatus {
	if o == nil {
		var ret WatchlistScreeningStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividual) GetStatusOk() (*WatchlistScreeningStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WatchlistScreeningIndividual) SetStatus(v WatchlistScreeningStatus) {
	o.Status = v
}

// GetClientUserId returns the ClientUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningIndividual) GetClientUserId() string {
	if o == nil || o.ClientUserId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientUserId.Get()
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividual) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientUserId.Get(), o.ClientUserId.IsSet()
}

// SetClientUserId sets field value
func (o *WatchlistScreeningIndividual) SetClientUserId(v string) {
	o.ClientUserId.Set(&v)
}

// GetAuditTrail returns the AuditTrail field value
func (o *WatchlistScreeningIndividual) GetAuditTrail() WatchlistScreeningAuditTrail {
	if o == nil {
		var ret WatchlistScreeningAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividual) GetAuditTrailOk() (*WatchlistScreeningAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *WatchlistScreeningIndividual) SetAuditTrail(v WatchlistScreeningAuditTrail) {
	o.AuditTrail = v
}

func (o WatchlistScreeningIndividual) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["search_terms"] = o.SearchTerms
	}
	if true {
		toSerialize["assignee"] = o.Assignee.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId.Get()
	}
	if true {
		toSerialize["audit_trail"] = o.AuditTrail
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistScreeningIndividual struct {
	value *WatchlistScreeningIndividual
	isSet bool
}

func (v NullableWatchlistScreeningIndividual) Get() *WatchlistScreeningIndividual {
	return v.value
}

func (v *NullableWatchlistScreeningIndividual) Set(val *WatchlistScreeningIndividual) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividual) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividual(val *WatchlistScreeningIndividual) *NullableWatchlistScreeningIndividual {
	return &NullableWatchlistScreeningIndividual{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


