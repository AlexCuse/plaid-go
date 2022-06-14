/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// EntityWatchlistScreening The entity screening object allows you to represent an entity in your system, update its profile, and search for it on various watchlists. Note: Rejected entity screenings will not receive new hits, regardless of entity program configuration.
type EntityWatchlistScreening struct {
	// ID of the associated entity screening.
	Id string `json:"id"`
	SearchTerms EntityWatchlistScreeningSearchTerms `json:"search_terms"`
	Assignee NullableString `json:"assignee"`
	Status WatchlistScreeningStatus `json:"status"`
	ClientUserId NullableString `json:"client_user_id"`
	AuditTrail WatchlistScreeningAuditTrail `json:"audit_trail"`
}

// NewEntityWatchlistScreening instantiates a new EntityWatchlistScreening object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityWatchlistScreening(id string, searchTerms EntityWatchlistScreeningSearchTerms, assignee NullableString, status WatchlistScreeningStatus, clientUserId NullableString, auditTrail WatchlistScreeningAuditTrail) *EntityWatchlistScreening {
	this := EntityWatchlistScreening{}
	this.Id = id
	this.SearchTerms = searchTerms
	this.Assignee = assignee
	this.Status = status
	this.ClientUserId = clientUserId
	this.AuditTrail = auditTrail
	return &this
}

// NewEntityWatchlistScreeningWithDefaults instantiates a new EntityWatchlistScreening object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityWatchlistScreeningWithDefaults() *EntityWatchlistScreening {
	this := EntityWatchlistScreening{}
	return &this
}

// GetId returns the Id field value
func (o *EntityWatchlistScreening) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreening) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntityWatchlistScreening) SetId(v string) {
	o.Id = v
}

// GetSearchTerms returns the SearchTerms field value
func (o *EntityWatchlistScreening) GetSearchTerms() EntityWatchlistScreeningSearchTerms {
	if o == nil {
		var ret EntityWatchlistScreeningSearchTerms
		return ret
	}

	return o.SearchTerms
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreening) GetSearchTermsOk() (*EntityWatchlistScreeningSearchTerms, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchTerms, true
}

// SetSearchTerms sets field value
func (o *EntityWatchlistScreening) SetSearchTerms(v EntityWatchlistScreeningSearchTerms) {
	o.SearchTerms = v
}

// GetAssignee returns the Assignee field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EntityWatchlistScreening) GetAssignee() string {
	if o == nil || o.Assignee.Get() == nil {
		var ret string
		return ret
	}

	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistScreening) GetAssigneeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// SetAssignee sets field value
func (o *EntityWatchlistScreening) SetAssignee(v string) {
	o.Assignee.Set(&v)
}

// GetStatus returns the Status field value
func (o *EntityWatchlistScreening) GetStatus() WatchlistScreeningStatus {
	if o == nil {
		var ret WatchlistScreeningStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreening) GetStatusOk() (*WatchlistScreeningStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EntityWatchlistScreening) SetStatus(v WatchlistScreeningStatus) {
	o.Status = v
}

// GetClientUserId returns the ClientUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EntityWatchlistScreening) GetClientUserId() string {
	if o == nil || o.ClientUserId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientUserId.Get()
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistScreening) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientUserId.Get(), o.ClientUserId.IsSet()
}

// SetClientUserId sets field value
func (o *EntityWatchlistScreening) SetClientUserId(v string) {
	o.ClientUserId.Set(&v)
}

// GetAuditTrail returns the AuditTrail field value
func (o *EntityWatchlistScreening) GetAuditTrail() WatchlistScreeningAuditTrail {
	if o == nil {
		var ret WatchlistScreeningAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreening) GetAuditTrailOk() (*WatchlistScreeningAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *EntityWatchlistScreening) SetAuditTrail(v WatchlistScreeningAuditTrail) {
	o.AuditTrail = v
}

func (o EntityWatchlistScreening) MarshalJSON() ([]byte, error) {
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

type NullableEntityWatchlistScreening struct {
	value *EntityWatchlistScreening
	isSet bool
}

func (v NullableEntityWatchlistScreening) Get() *EntityWatchlistScreening {
	return v.value
}

func (v *NullableEntityWatchlistScreening) Set(val *EntityWatchlistScreening) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityWatchlistScreening) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityWatchlistScreening) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityWatchlistScreening(val *EntityWatchlistScreening) *NullableEntityWatchlistScreening {
	return &NullableEntityWatchlistScreening{value: val, isSet: true}
}

func (v NullableEntityWatchlistScreening) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityWatchlistScreening) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


