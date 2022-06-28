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

// EntityWatchlistScreeningReview A review submitted by a team member for an entity watchlist screening. A review can be either a comment on the current screening state, actions taken against hits attached to the watchlist screening, or both.
type EntityWatchlistScreeningReview struct {
	// ID of the associated entity review.
	Id string `json:"id"`
	// Hits marked as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.
	ConfirmedHits []string `json:"confirmed_hits"`
	// Hits marked as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.
	DismissedHits []string `json:"dismissed_hits"`
	// A comment submitted by a team member as part of reviewing a watchlist screening.
	Comment NullableString `json:"comment"`
	AuditTrail WatchlistScreeningAuditTrail `json:"audit_trail"`
}

// NewEntityWatchlistScreeningReview instantiates a new EntityWatchlistScreeningReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityWatchlistScreeningReview(id string, confirmedHits []string, dismissedHits []string, comment NullableString, auditTrail WatchlistScreeningAuditTrail) *EntityWatchlistScreeningReview {
	this := EntityWatchlistScreeningReview{}
	this.Id = id
	this.ConfirmedHits = confirmedHits
	this.DismissedHits = dismissedHits
	this.Comment = comment
	this.AuditTrail = auditTrail
	return &this
}

// NewEntityWatchlistScreeningReviewWithDefaults instantiates a new EntityWatchlistScreeningReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityWatchlistScreeningReviewWithDefaults() *EntityWatchlistScreeningReview {
	this := EntityWatchlistScreeningReview{}
	return &this
}

// GetId returns the Id field value
func (o *EntityWatchlistScreeningReview) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreeningReview) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntityWatchlistScreeningReview) SetId(v string) {
	o.Id = v
}

// GetConfirmedHits returns the ConfirmedHits field value
func (o *EntityWatchlistScreeningReview) GetConfirmedHits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfirmedHits
}

// GetConfirmedHitsOk returns a tuple with the ConfirmedHits field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreeningReview) GetConfirmedHitsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmedHits, true
}

// SetConfirmedHits sets field value
func (o *EntityWatchlistScreeningReview) SetConfirmedHits(v []string) {
	o.ConfirmedHits = v
}

// GetDismissedHits returns the DismissedHits field value
func (o *EntityWatchlistScreeningReview) GetDismissedHits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DismissedHits
}

// GetDismissedHitsOk returns a tuple with the DismissedHits field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreeningReview) GetDismissedHitsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DismissedHits, true
}

// SetDismissedHits sets field value
func (o *EntityWatchlistScreeningReview) SetDismissedHits(v []string) {
	o.DismissedHits = v
}

// GetComment returns the Comment field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EntityWatchlistScreeningReview) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}

	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistScreeningReview) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// SetComment sets field value
func (o *EntityWatchlistScreeningReview) SetComment(v string) {
	o.Comment.Set(&v)
}

// GetAuditTrail returns the AuditTrail field value
func (o *EntityWatchlistScreeningReview) GetAuditTrail() WatchlistScreeningAuditTrail {
	if o == nil {
		var ret WatchlistScreeningAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistScreeningReview) GetAuditTrailOk() (*WatchlistScreeningAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *EntityWatchlistScreeningReview) SetAuditTrail(v WatchlistScreeningAuditTrail) {
	o.AuditTrail = v
}

func (o EntityWatchlistScreeningReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["confirmed_hits"] = o.ConfirmedHits
	}
	if true {
		toSerialize["dismissed_hits"] = o.DismissedHits
	}
	if true {
		toSerialize["comment"] = o.Comment.Get()
	}
	if true {
		toSerialize["audit_trail"] = o.AuditTrail
	}
	return json.Marshal(toSerialize)
}

type NullableEntityWatchlistScreeningReview struct {
	value *EntityWatchlistScreeningReview
	isSet bool
}

func (v NullableEntityWatchlistScreeningReview) Get() *EntityWatchlistScreeningReview {
	return v.value
}

func (v *NullableEntityWatchlistScreeningReview) Set(val *EntityWatchlistScreeningReview) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityWatchlistScreeningReview) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityWatchlistScreeningReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityWatchlistScreeningReview(val *EntityWatchlistScreeningReview) *NullableEntityWatchlistScreeningReview {
	return &NullableEntityWatchlistScreeningReview{value: val, isSet: true}
}

func (v NullableEntityWatchlistScreeningReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityWatchlistScreeningReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


