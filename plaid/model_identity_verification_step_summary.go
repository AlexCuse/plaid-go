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

// IdentityVerificationStepSummary Each step will be one of the following values:   `active` - This step is the user's current step. They are either in the process of completing this step, or they recently closed their Identity Verification attempt while in the middle of this step. Only one step will be marked as `active` at any given point.  `success` - The Identity Verification attempt has completed this step.  `failed` - The user failed this step. This can either call the user to fail the session as a whole, or cause them to fallback to another step depending on how the Identity Verification template is configured. A failed step does not imply a failed session.  `waiting_for_prerequisite` - The user needs to complete another step first, before they progress to this step. This step may never run, depending on if the user fails an earlier step or if the step is only run as a fallback.  `not_applicable` - This step will not be run for this session.  `skipped` - The retry instructions that created this Identity Verification attempt specified that this step should be skipped.  `expired` - This step had not yet been completed when the Identity Verification attempt as a whole expired.  `canceled` - The Identity Verification attempt was canceled before the user completed this step.  `pending_review` - The Identity Verification attempt template was configured to perform a screening that had one or more hits needing review.  `manually_approved` - The step was manually overridden to pass by a team member in the dashboard.  `manually_rejected` - The step was manually overridden to fail by a team member in the dashboard.
type IdentityVerificationStepSummary struct {
	AcceptTos IdentityVerificationStepStatus `json:"accept_tos"`
	VerifySms IdentityVerificationStepStatus `json:"verify_sms"`
	KycCheck IdentityVerificationStepStatus `json:"kyc_check"`
	DocumentaryVerification IdentityVerificationStepStatus `json:"documentary_verification"`
	SelfieCheck IdentityVerificationStepStatus `json:"selfie_check"`
	WatchlistScreening IdentityVerificationStepStatus `json:"watchlist_screening"`
	RiskCheck IdentityVerificationStepStatus `json:"risk_check"`
	AdditionalProperties map[string]interface{}
}

type _IdentityVerificationStepSummary IdentityVerificationStepSummary

// NewIdentityVerificationStepSummary instantiates a new IdentityVerificationStepSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationStepSummary(acceptTos IdentityVerificationStepStatus, verifySms IdentityVerificationStepStatus, kycCheck IdentityVerificationStepStatus, documentaryVerification IdentityVerificationStepStatus, selfieCheck IdentityVerificationStepStatus, watchlistScreening IdentityVerificationStepStatus, riskCheck IdentityVerificationStepStatus) *IdentityVerificationStepSummary {
	this := IdentityVerificationStepSummary{}
	this.AcceptTos = acceptTos
	this.VerifySms = verifySms
	this.KycCheck = kycCheck
	this.DocumentaryVerification = documentaryVerification
	this.SelfieCheck = selfieCheck
	this.WatchlistScreening = watchlistScreening
	this.RiskCheck = riskCheck
	return &this
}

// NewIdentityVerificationStepSummaryWithDefaults instantiates a new IdentityVerificationStepSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationStepSummaryWithDefaults() *IdentityVerificationStepSummary {
	this := IdentityVerificationStepSummary{}
	return &this
}

// GetAcceptTos returns the AcceptTos field value
func (o *IdentityVerificationStepSummary) GetAcceptTos() IdentityVerificationStepStatus {
	if o == nil {
		var ret IdentityVerificationStepStatus
		return ret
	}

	return o.AcceptTos
}

// GetAcceptTosOk returns a tuple with the AcceptTos field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepSummary) GetAcceptTosOk() (*IdentityVerificationStepStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AcceptTos, true
}

// SetAcceptTos sets field value
func (o *IdentityVerificationStepSummary) SetAcceptTos(v IdentityVerificationStepStatus) {
	o.AcceptTos = v
}

// GetVerifySms returns the VerifySms field value
func (o *IdentityVerificationStepSummary) GetVerifySms() IdentityVerificationStepStatus {
	if o == nil {
		var ret IdentityVerificationStepStatus
		return ret
	}

	return o.VerifySms
}

// GetVerifySmsOk returns a tuple with the VerifySms field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepSummary) GetVerifySmsOk() (*IdentityVerificationStepStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerifySms, true
}

// SetVerifySms sets field value
func (o *IdentityVerificationStepSummary) SetVerifySms(v IdentityVerificationStepStatus) {
	o.VerifySms = v
}

// GetKycCheck returns the KycCheck field value
func (o *IdentityVerificationStepSummary) GetKycCheck() IdentityVerificationStepStatus {
	if o == nil {
		var ret IdentityVerificationStepStatus
		return ret
	}

	return o.KycCheck
}

// GetKycCheckOk returns a tuple with the KycCheck field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepSummary) GetKycCheckOk() (*IdentityVerificationStepStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KycCheck, true
}

// SetKycCheck sets field value
func (o *IdentityVerificationStepSummary) SetKycCheck(v IdentityVerificationStepStatus) {
	o.KycCheck = v
}

// GetDocumentaryVerification returns the DocumentaryVerification field value
func (o *IdentityVerificationStepSummary) GetDocumentaryVerification() IdentityVerificationStepStatus {
	if o == nil {
		var ret IdentityVerificationStepStatus
		return ret
	}

	return o.DocumentaryVerification
}

// GetDocumentaryVerificationOk returns a tuple with the DocumentaryVerification field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepSummary) GetDocumentaryVerificationOk() (*IdentityVerificationStepStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentaryVerification, true
}

// SetDocumentaryVerification sets field value
func (o *IdentityVerificationStepSummary) SetDocumentaryVerification(v IdentityVerificationStepStatus) {
	o.DocumentaryVerification = v
}

// GetSelfieCheck returns the SelfieCheck field value
func (o *IdentityVerificationStepSummary) GetSelfieCheck() IdentityVerificationStepStatus {
	if o == nil {
		var ret IdentityVerificationStepStatus
		return ret
	}

	return o.SelfieCheck
}

// GetSelfieCheckOk returns a tuple with the SelfieCheck field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepSummary) GetSelfieCheckOk() (*IdentityVerificationStepStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SelfieCheck, true
}

// SetSelfieCheck sets field value
func (o *IdentityVerificationStepSummary) SetSelfieCheck(v IdentityVerificationStepStatus) {
	o.SelfieCheck = v
}

// GetWatchlistScreening returns the WatchlistScreening field value
func (o *IdentityVerificationStepSummary) GetWatchlistScreening() IdentityVerificationStepStatus {
	if o == nil {
		var ret IdentityVerificationStepStatus
		return ret
	}

	return o.WatchlistScreening
}

// GetWatchlistScreeningOk returns a tuple with the WatchlistScreening field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepSummary) GetWatchlistScreeningOk() (*IdentityVerificationStepStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistScreening, true
}

// SetWatchlistScreening sets field value
func (o *IdentityVerificationStepSummary) SetWatchlistScreening(v IdentityVerificationStepStatus) {
	o.WatchlistScreening = v
}

// GetRiskCheck returns the RiskCheck field value
func (o *IdentityVerificationStepSummary) GetRiskCheck() IdentityVerificationStepStatus {
	if o == nil {
		var ret IdentityVerificationStepStatus
		return ret
	}

	return o.RiskCheck
}

// GetRiskCheckOk returns a tuple with the RiskCheck field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepSummary) GetRiskCheckOk() (*IdentityVerificationStepStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RiskCheck, true
}

// SetRiskCheck sets field value
func (o *IdentityVerificationStepSummary) SetRiskCheck(v IdentityVerificationStepStatus) {
	o.RiskCheck = v
}

func (o IdentityVerificationStepSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accept_tos"] = o.AcceptTos
	}
	if true {
		toSerialize["verify_sms"] = o.VerifySms
	}
	if true {
		toSerialize["kyc_check"] = o.KycCheck
	}
	if true {
		toSerialize["documentary_verification"] = o.DocumentaryVerification
	}
	if true {
		toSerialize["selfie_check"] = o.SelfieCheck
	}
	if true {
		toSerialize["watchlist_screening"] = o.WatchlistScreening
	}
	if true {
		toSerialize["risk_check"] = o.RiskCheck
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityVerificationStepSummary) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityVerificationStepSummary := _IdentityVerificationStepSummary{}

	if err = json.Unmarshal(bytes, &varIdentityVerificationStepSummary); err == nil {
		*o = IdentityVerificationStepSummary(varIdentityVerificationStepSummary)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accept_tos")
		delete(additionalProperties, "verify_sms")
		delete(additionalProperties, "kyc_check")
		delete(additionalProperties, "documentary_verification")
		delete(additionalProperties, "selfie_check")
		delete(additionalProperties, "watchlist_screening")
		delete(additionalProperties, "risk_check")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityVerificationStepSummary struct {
	value *IdentityVerificationStepSummary
	isSet bool
}

func (v NullableIdentityVerificationStepSummary) Get() *IdentityVerificationStepSummary {
	return v.value
}

func (v *NullableIdentityVerificationStepSummary) Set(val *IdentityVerificationStepSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationStepSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationStepSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationStepSummary(val *IdentityVerificationStepSummary) *NullableIdentityVerificationStepSummary {
	return &NullableIdentityVerificationStepSummary{value: val, isSet: true}
}

func (v NullableIdentityVerificationStepSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationStepSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


