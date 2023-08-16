/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.410.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// CreditBankStatementUploadItem An object containing information about the bank statement upload Item.
type CreditBankStatementUploadItem struct {
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	BankStatements []CreditBankStatementUploadObject `json:"bank_statements"`
	Status NullablePayrollItemStatus `json:"status"`
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DDTHH:mm:ssZ) indicating the last time that the Item was updated.
	UpdatedAt NullableTime `json:"updated_at"`
}

// NewCreditBankStatementUploadItem instantiates a new CreditBankStatementUploadItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankStatementUploadItem(itemId string, bankStatements []CreditBankStatementUploadObject, status NullablePayrollItemStatus, updatedAt NullableTime) *CreditBankStatementUploadItem {
	this := CreditBankStatementUploadItem{}
	this.ItemId = itemId
	this.BankStatements = bankStatements
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewCreditBankStatementUploadItemWithDefaults instantiates a new CreditBankStatementUploadItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankStatementUploadItemWithDefaults() *CreditBankStatementUploadItem {
	this := CreditBankStatementUploadItem{}
	return &this
}

// GetItemId returns the ItemId field value
func (o *CreditBankStatementUploadItem) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *CreditBankStatementUploadItem) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *CreditBankStatementUploadItem) SetItemId(v string) {
	o.ItemId = v
}

// GetBankStatements returns the BankStatements field value
func (o *CreditBankStatementUploadItem) GetBankStatements() []CreditBankStatementUploadObject {
	if o == nil {
		var ret []CreditBankStatementUploadObject
		return ret
	}

	return o.BankStatements
}

// GetBankStatementsOk returns a tuple with the BankStatements field value
// and a boolean to check if the value has been set.
func (o *CreditBankStatementUploadItem) GetBankStatementsOk() (*[]CreditBankStatementUploadObject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankStatements, true
}

// SetBankStatements sets field value
func (o *CreditBankStatementUploadItem) SetBankStatements(v []CreditBankStatementUploadObject) {
	o.BankStatements = v
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for PayrollItemStatus will be returned
func (o *CreditBankStatementUploadItem) GetStatus() PayrollItemStatus {
	if o == nil || o.Status.Get() == nil {
		var ret PayrollItemStatus
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankStatementUploadItem) GetStatusOk() (*PayrollItemStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *CreditBankStatementUploadItem) SetStatus(v PayrollItemStatus) {
	o.Status.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CreditBankStatementUploadItem) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankStatementUploadItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *CreditBankStatementUploadItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

func (o CreditBankStatementUploadItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["bank_statements"] = o.BankStatements
	}
	if true {
		toSerialize["status"] = o.Status.Get()
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankStatementUploadItem struct {
	value *CreditBankStatementUploadItem
	isSet bool
}

func (v NullableCreditBankStatementUploadItem) Get() *CreditBankStatementUploadItem {
	return v.value
}

func (v *NullableCreditBankStatementUploadItem) Set(val *CreditBankStatementUploadItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankStatementUploadItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankStatementUploadItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankStatementUploadItem(val *CreditBankStatementUploadItem) *NullableCreditBankStatementUploadItem {
	return &NullableCreditBankStatementUploadItem{value: val, isSet: true}
}

func (v NullableCreditBankStatementUploadItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankStatementUploadItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


