/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.97.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditBankIncomeCause An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.
type CreditBankIncomeCause struct {
	ErrorType *CreditBankIncomeErrorType `json:"error_type,omitempty"`
	// We use standard HTTP response codes for success and failure notifications, and our errors are further classified by `error_type`. In general, 200 HTTP codes correspond to success, 40X codes are for developer- or user-related failures, and 50X codes are for Plaid-related issues. Error fields will be `null` if no error has occurred.
	ErrorCode *string `json:"error_code,omitempty"`
	// A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.
	ErrorMessage *string `json:"error_message,omitempty"`
	// A user-friendly representation of the error code. null if the error is not related to user action. This may change over time and is not safe for programmatic use.
	DisplayMessage *string `json:"display_message,omitempty"`
	// The `item_id` of the Item associated with this warning.
	ItemId *string `json:"item_id,omitempty"`
}

// NewCreditBankIncomeCause instantiates a new CreditBankIncomeCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeCause() *CreditBankIncomeCause {
	this := CreditBankIncomeCause{}
	return &this
}

// NewCreditBankIncomeCauseWithDefaults instantiates a new CreditBankIncomeCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeCauseWithDefaults() *CreditBankIncomeCause {
	this := CreditBankIncomeCause{}
	return &this
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *CreditBankIncomeCause) GetErrorType() CreditBankIncomeErrorType {
	if o == nil || o.ErrorType == nil {
		var ret CreditBankIncomeErrorType
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeCause) GetErrorTypeOk() (*CreditBankIncomeErrorType, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *CreditBankIncomeCause) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given CreditBankIncomeErrorType and assigns it to the ErrorType field.
func (o *CreditBankIncomeCause) SetErrorType(v CreditBankIncomeErrorType) {
	o.ErrorType = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CreditBankIncomeCause) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeCause) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CreditBankIncomeCause) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CreditBankIncomeCause) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CreditBankIncomeCause) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeCause) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CreditBankIncomeCause) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CreditBankIncomeCause) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetDisplayMessage returns the DisplayMessage field value if set, zero value otherwise.
func (o *CreditBankIncomeCause) GetDisplayMessage() string {
	if o == nil || o.DisplayMessage == nil {
		var ret string
		return ret
	}
	return *o.DisplayMessage
}

// GetDisplayMessageOk returns a tuple with the DisplayMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeCause) GetDisplayMessageOk() (*string, bool) {
	if o == nil || o.DisplayMessage == nil {
		return nil, false
	}
	return o.DisplayMessage, true
}

// HasDisplayMessage returns a boolean if a field has been set.
func (o *CreditBankIncomeCause) HasDisplayMessage() bool {
	if o != nil && o.DisplayMessage != nil {
		return true
	}

	return false
}

// SetDisplayMessage gets a reference to the given string and assigns it to the DisplayMessage field.
func (o *CreditBankIncomeCause) SetDisplayMessage(v string) {
	o.DisplayMessage = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *CreditBankIncomeCause) GetItemId() string {
	if o == nil || o.ItemId == nil {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeCause) GetItemIdOk() (*string, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *CreditBankIncomeCause) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *CreditBankIncomeCause) SetItemId(v string) {
	o.ItemId = &v
}

func (o CreditBankIncomeCause) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorType != nil {
		toSerialize["error_type"] = o.ErrorType
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.DisplayMessage != nil {
		toSerialize["display_message"] = o.DisplayMessage
	}
	if o.ItemId != nil {
		toSerialize["item_id"] = o.ItemId
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankIncomeCause struct {
	value *CreditBankIncomeCause
	isSet bool
}

func (v NullableCreditBankIncomeCause) Get() *CreditBankIncomeCause {
	return v.value
}

func (v *NullableCreditBankIncomeCause) Set(val *CreditBankIncomeCause) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeCause) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeCause(val *CreditBankIncomeCause) *NullableCreditBankIncomeCause {
	return &NullableCreditBankIncomeCause{value: val, isSet: true}
}

func (v NullableCreditBankIncomeCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


