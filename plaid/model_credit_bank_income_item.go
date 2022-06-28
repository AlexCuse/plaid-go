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
	"time"
)

// CreditBankIncomeItem The details and metadata for an end user's Item.
type CreditBankIncomeItem struct {
	// The Item's accounts that have Bank Income data.
	BankIncomeAccounts *[]CreditBankIncomeAccount `json:"bank_income_accounts,omitempty"`
	// The income sources for this Item. Each entry in the array is a single income source.
	BankIncomeSources *[]CreditBankIncomeSource `json:"bank_income_sources,omitempty"`
	// The time when this Item's data was last retrieved from the financial institution.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// The unique identifier of the institution associated with the Item.
	InstitutionId *string `json:"institution_id,omitempty"`
	// The name of the institution associated with the Item.
	InstitutionName *string `json:"institution_name,omitempty"`
	// The unique identifier for the Item.
	ItemId *string `json:"item_id,omitempty"`
}

// NewCreditBankIncomeItem instantiates a new CreditBankIncomeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeItem() *CreditBankIncomeItem {
	this := CreditBankIncomeItem{}
	return &this
}

// NewCreditBankIncomeItemWithDefaults instantiates a new CreditBankIncomeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeItemWithDefaults() *CreditBankIncomeItem {
	this := CreditBankIncomeItem{}
	return &this
}

// GetBankIncomeAccounts returns the BankIncomeAccounts field value if set, zero value otherwise.
func (o *CreditBankIncomeItem) GetBankIncomeAccounts() []CreditBankIncomeAccount {
	if o == nil || o.BankIncomeAccounts == nil {
		var ret []CreditBankIncomeAccount
		return ret
	}
	return *o.BankIncomeAccounts
}

// GetBankIncomeAccountsOk returns a tuple with the BankIncomeAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeItem) GetBankIncomeAccountsOk() (*[]CreditBankIncomeAccount, bool) {
	if o == nil || o.BankIncomeAccounts == nil {
		return nil, false
	}
	return o.BankIncomeAccounts, true
}

// HasBankIncomeAccounts returns a boolean if a field has been set.
func (o *CreditBankIncomeItem) HasBankIncomeAccounts() bool {
	if o != nil && o.BankIncomeAccounts != nil {
		return true
	}

	return false
}

// SetBankIncomeAccounts gets a reference to the given []CreditBankIncomeAccount and assigns it to the BankIncomeAccounts field.
func (o *CreditBankIncomeItem) SetBankIncomeAccounts(v []CreditBankIncomeAccount) {
	o.BankIncomeAccounts = &v
}

// GetBankIncomeSources returns the BankIncomeSources field value if set, zero value otherwise.
func (o *CreditBankIncomeItem) GetBankIncomeSources() []CreditBankIncomeSource {
	if o == nil || o.BankIncomeSources == nil {
		var ret []CreditBankIncomeSource
		return ret
	}
	return *o.BankIncomeSources
}

// GetBankIncomeSourcesOk returns a tuple with the BankIncomeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeItem) GetBankIncomeSourcesOk() (*[]CreditBankIncomeSource, bool) {
	if o == nil || o.BankIncomeSources == nil {
		return nil, false
	}
	return o.BankIncomeSources, true
}

// HasBankIncomeSources returns a boolean if a field has been set.
func (o *CreditBankIncomeItem) HasBankIncomeSources() bool {
	if o != nil && o.BankIncomeSources != nil {
		return true
	}

	return false
}

// SetBankIncomeSources gets a reference to the given []CreditBankIncomeSource and assigns it to the BankIncomeSources field.
func (o *CreditBankIncomeItem) SetBankIncomeSources(v []CreditBankIncomeSource) {
	o.BankIncomeSources = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *CreditBankIncomeItem) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeItem) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *CreditBankIncomeItem) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *CreditBankIncomeItem) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *CreditBankIncomeItem) GetInstitutionId() string {
	if o == nil || o.InstitutionId == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeItem) GetInstitutionIdOk() (*string, bool) {
	if o == nil || o.InstitutionId == nil {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *CreditBankIncomeItem) HasInstitutionId() bool {
	if o != nil && o.InstitutionId != nil {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *CreditBankIncomeItem) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

// GetInstitutionName returns the InstitutionName field value if set, zero value otherwise.
func (o *CreditBankIncomeItem) GetInstitutionName() string {
	if o == nil || o.InstitutionName == nil {
		var ret string
		return ret
	}
	return *o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeItem) GetInstitutionNameOk() (*string, bool) {
	if o == nil || o.InstitutionName == nil {
		return nil, false
	}
	return o.InstitutionName, true
}

// HasInstitutionName returns a boolean if a field has been set.
func (o *CreditBankIncomeItem) HasInstitutionName() bool {
	if o != nil && o.InstitutionName != nil {
		return true
	}

	return false
}

// SetInstitutionName gets a reference to the given string and assigns it to the InstitutionName field.
func (o *CreditBankIncomeItem) SetInstitutionName(v string) {
	o.InstitutionName = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *CreditBankIncomeItem) GetItemId() string {
	if o == nil || o.ItemId == nil {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeItem) GetItemIdOk() (*string, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *CreditBankIncomeItem) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *CreditBankIncomeItem) SetItemId(v string) {
	o.ItemId = &v
}

func (o CreditBankIncomeItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BankIncomeAccounts != nil {
		toSerialize["bank_income_accounts"] = o.BankIncomeAccounts
	}
	if o.BankIncomeSources != nil {
		toSerialize["bank_income_sources"] = o.BankIncomeSources
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.InstitutionId != nil {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if o.InstitutionName != nil {
		toSerialize["institution_name"] = o.InstitutionName
	}
	if o.ItemId != nil {
		toSerialize["item_id"] = o.ItemId
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankIncomeItem struct {
	value *CreditBankIncomeItem
	isSet bool
}

func (v NullableCreditBankIncomeItem) Get() *CreditBankIncomeItem {
	return v.value
}

func (v *NullableCreditBankIncomeItem) Set(val *CreditBankIncomeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeItem(val *CreditBankIncomeItem) *NullableCreditBankIncomeItem {
	return &NullableCreditBankIncomeItem{value: val, isSet: true}
}

func (v NullableCreditBankIncomeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


