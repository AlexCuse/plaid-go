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

// CreditBankIncomeCause An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.
type CreditBankIncomeCause struct {
	ErrorType CreditBankIncomeErrorType `json:"error_type"`
	// We use standard HTTP response codes for success and failure notifications, and our errors are further classified by `error_type`. In general, 200 HTTP codes correspond to success, 40X codes are for developer- or user-related failures, and 50X codes are for Plaid-related issues. Error fields will be `null` if no error has occurred.
	ErrorCode string `json:"error_code"`
	// A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.
	ErrorMessage string `json:"error_message"`
	// A user-friendly representation of the error code. null if the error is not related to user action. This may change over time and is not safe for programmatic use.
	DisplayMessage string `json:"display_message"`
	// The `item_id` of the Item associated with this warning.
	ItemId string `json:"item_id"`
}

// NewCreditBankIncomeCause instantiates a new CreditBankIncomeCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeCause(errorType CreditBankIncomeErrorType, errorCode string, errorMessage string, displayMessage string, itemId string) *CreditBankIncomeCause {
	this := CreditBankIncomeCause{}
	this.ErrorType = errorType
	this.ErrorCode = errorCode
	this.ErrorMessage = errorMessage
	this.DisplayMessage = displayMessage
	this.ItemId = itemId
	return &this
}

// NewCreditBankIncomeCauseWithDefaults instantiates a new CreditBankIncomeCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeCauseWithDefaults() *CreditBankIncomeCause {
	this := CreditBankIncomeCause{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *CreditBankIncomeCause) GetErrorType() CreditBankIncomeErrorType {
	if o == nil {
		var ret CreditBankIncomeErrorType
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeCause) GetErrorTypeOk() (*CreditBankIncomeErrorType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *CreditBankIncomeCause) SetErrorType(v CreditBankIncomeErrorType) {
	o.ErrorType = v
}

// GetErrorCode returns the ErrorCode field value
func (o *CreditBankIncomeCause) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeCause) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *CreditBankIncomeCause) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value
func (o *CreditBankIncomeCause) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeCause) GetErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *CreditBankIncomeCause) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

// GetDisplayMessage returns the DisplayMessage field value
func (o *CreditBankIncomeCause) GetDisplayMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayMessage
}

// GetDisplayMessageOk returns a tuple with the DisplayMessage field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeCause) GetDisplayMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayMessage, true
}

// SetDisplayMessage sets field value
func (o *CreditBankIncomeCause) SetDisplayMessage(v string) {
	o.DisplayMessage = v
}

// GetItemId returns the ItemId field value
func (o *CreditBankIncomeCause) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeCause) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *CreditBankIncomeCause) SetItemId(v string) {
	o.ItemId = v
}

func (o CreditBankIncomeCause) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error_type"] = o.ErrorType
	}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if true {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if true {
		toSerialize["display_message"] = o.DisplayMessage
	}
	if true {
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

