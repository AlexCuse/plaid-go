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

// BankInitiatedReturnRisk The object contains a risk score and a risk tier that evaluate the transaction return risk because an account is overdrawn or because an ineligible account is used. Common return codes in this category include: \"R01\", \"R02\", \"R03\", \"R04\", \"R06\", \"R08\",  \"R09\", \"R13\", \"R16\", \"R17\", \"R20\", \"R23\". These returns have a turnaround time of 2 banking days.
type BankInitiatedReturnRisk struct {
	// A score from 1-99 that indicates the transaction return risk: a higher risk score suggests a higher return likelihood.
	Score int32 `json:"score"`
	// In the `bank_initiated_return_risk` object, there are eight risk tiers corresponding to the scores:   1: Predicted bank-initiated return incidence rate between 0.0% - 0.5%   2: Predicted bank-initiated return incidence rate between 0.5% - 1.5%   3: Predicted bank-initiated return incidence rate between 1.5% - 3%   4: Predicted bank-initiated return incidence rate between 3% - 5%   5: Predicted bank-initiated return incidence rate between 5% - 10%   6: Predicted bank-initiated return incidence rate between 10% - 15%   7: Predicted bank-initiated return incidence rate between 15% and 50%   8: Predicted bank-initiated return incidence rate greater than 50% 
	RiskTier int32 `json:"risk_tier"`
}

// NewBankInitiatedReturnRisk instantiates a new BankInitiatedReturnRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankInitiatedReturnRisk(score int32, riskTier int32) *BankInitiatedReturnRisk {
	this := BankInitiatedReturnRisk{}
	this.Score = score
	this.RiskTier = riskTier
	return &this
}

// NewBankInitiatedReturnRiskWithDefaults instantiates a new BankInitiatedReturnRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankInitiatedReturnRiskWithDefaults() *BankInitiatedReturnRisk {
	this := BankInitiatedReturnRisk{}
	return &this
}

// GetScore returns the Score field value
func (o *BankInitiatedReturnRisk) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *BankInitiatedReturnRisk) GetScoreOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *BankInitiatedReturnRisk) SetScore(v int32) {
	o.Score = v
}

// GetRiskTier returns the RiskTier field value
func (o *BankInitiatedReturnRisk) GetRiskTier() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RiskTier
}

// GetRiskTierOk returns a tuple with the RiskTier field value
// and a boolean to check if the value has been set.
func (o *BankInitiatedReturnRisk) GetRiskTierOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RiskTier, true
}

// SetRiskTier sets field value
func (o *BankInitiatedReturnRisk) SetRiskTier(v int32) {
	o.RiskTier = v
}

func (o BankInitiatedReturnRisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["score"] = o.Score
	}
	if true {
		toSerialize["risk_tier"] = o.RiskTier
	}
	return json.Marshal(toSerialize)
}

type NullableBankInitiatedReturnRisk struct {
	value *BankInitiatedReturnRisk
	isSet bool
}

func (v NullableBankInitiatedReturnRisk) Get() *BankInitiatedReturnRisk {
	return v.value
}

func (v *NullableBankInitiatedReturnRisk) Set(val *BankInitiatedReturnRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableBankInitiatedReturnRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableBankInitiatedReturnRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankInitiatedReturnRisk(val *BankInitiatedReturnRisk) *NullableBankInitiatedReturnRisk {
	return &NullableBankInitiatedReturnRisk{value: val, isSet: true}
}

func (v NullableBankInitiatedReturnRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankInitiatedReturnRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


