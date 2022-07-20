/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransactionsSyncRequestOptions An optional object to be used with the request. If specified, `options` must not be `null`.
type TransactionsSyncRequestOptions struct {
	// Include the raw unparsed transaction description from the financial institution. This field is disabled by default. If you need this information in addition to the parsed data provided, contact your Plaid Account Manager.
	IncludeOriginalDescription NullableBool `json:"include_original_description,omitempty"`
	// Include the [`personal_finance_category`](https://plaid.com/docs/api/products/transactions/#transactions-sync-response-added-personal-finance-category) object in the response.  See the [`taxonomy csv file`](https://plaid.com/documents/transactions-personal-finance-category-taxonomy.csv) for a full list of personal finance categories.  We’re introducing Category Rules - a new beta endpoint that will enable you to change the `personal_finance_category` for a transaction based on your users’ needs. When rules are set, the selected category will override the Plaid provided category. To learn more, send a note to transactions-feedback@plaid.com.
	IncludePersonalFinanceCategory *bool `json:"include_personal_finance_category,omitempty"`
}

// NewTransactionsSyncRequestOptions instantiates a new TransactionsSyncRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsSyncRequestOptions() *TransactionsSyncRequestOptions {
	this := TransactionsSyncRequestOptions{}
	var includeOriginalDescription bool = false
	this.IncludeOriginalDescription = *NewNullableBool(&includeOriginalDescription)
	var includePersonalFinanceCategory bool = false
	this.IncludePersonalFinanceCategory = &includePersonalFinanceCategory
	return &this
}

// NewTransactionsSyncRequestOptionsWithDefaults instantiates a new TransactionsSyncRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsSyncRequestOptionsWithDefaults() *TransactionsSyncRequestOptions {
	this := TransactionsSyncRequestOptions{}
	var includeOriginalDescription bool = false
	this.IncludeOriginalDescription = *NewNullableBool(&includeOriginalDescription)
	var includePersonalFinanceCategory bool = false
	this.IncludePersonalFinanceCategory = &includePersonalFinanceCategory
	return &this
}

// GetIncludeOriginalDescription returns the IncludeOriginalDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionsSyncRequestOptions) GetIncludeOriginalDescription() bool {
	if o == nil || o.IncludeOriginalDescription.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludeOriginalDescription.Get()
}

// GetIncludeOriginalDescriptionOk returns a tuple with the IncludeOriginalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionsSyncRequestOptions) GetIncludeOriginalDescriptionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncludeOriginalDescription.Get(), o.IncludeOriginalDescription.IsSet()
}

// HasIncludeOriginalDescription returns a boolean if a field has been set.
func (o *TransactionsSyncRequestOptions) HasIncludeOriginalDescription() bool {
	if o != nil && o.IncludeOriginalDescription.IsSet() {
		return true
	}

	return false
}

// SetIncludeOriginalDescription gets a reference to the given NullableBool and assigns it to the IncludeOriginalDescription field.
func (o *TransactionsSyncRequestOptions) SetIncludeOriginalDescription(v bool) {
	o.IncludeOriginalDescription.Set(&v)
}
// SetIncludeOriginalDescriptionNil sets the value for IncludeOriginalDescription to be an explicit nil
func (o *TransactionsSyncRequestOptions) SetIncludeOriginalDescriptionNil() {
	o.IncludeOriginalDescription.Set(nil)
}

// UnsetIncludeOriginalDescription ensures that no value is present for IncludeOriginalDescription, not even an explicit nil
func (o *TransactionsSyncRequestOptions) UnsetIncludeOriginalDescription() {
	o.IncludeOriginalDescription.Unset()
}

// GetIncludePersonalFinanceCategory returns the IncludePersonalFinanceCategory field value if set, zero value otherwise.
func (o *TransactionsSyncRequestOptions) GetIncludePersonalFinanceCategory() bool {
	if o == nil || o.IncludePersonalFinanceCategory == nil {
		var ret bool
		return ret
	}
	return *o.IncludePersonalFinanceCategory
}

// GetIncludePersonalFinanceCategoryOk returns a tuple with the IncludePersonalFinanceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsSyncRequestOptions) GetIncludePersonalFinanceCategoryOk() (*bool, bool) {
	if o == nil || o.IncludePersonalFinanceCategory == nil {
		return nil, false
	}
	return o.IncludePersonalFinanceCategory, true
}

// HasIncludePersonalFinanceCategory returns a boolean if a field has been set.
func (o *TransactionsSyncRequestOptions) HasIncludePersonalFinanceCategory() bool {
	if o != nil && o.IncludePersonalFinanceCategory != nil {
		return true
	}

	return false
}

// SetIncludePersonalFinanceCategory gets a reference to the given bool and assigns it to the IncludePersonalFinanceCategory field.
func (o *TransactionsSyncRequestOptions) SetIncludePersonalFinanceCategory(v bool) {
	o.IncludePersonalFinanceCategory = &v
}

func (o TransactionsSyncRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeOriginalDescription.IsSet() {
		toSerialize["include_original_description"] = o.IncludeOriginalDescription.Get()
	}
	if o.IncludePersonalFinanceCategory != nil {
		toSerialize["include_personal_finance_category"] = o.IncludePersonalFinanceCategory
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsSyncRequestOptions struct {
	value *TransactionsSyncRequestOptions
	isSet bool
}

func (v NullableTransactionsSyncRequestOptions) Get() *TransactionsSyncRequestOptions {
	return v.value
}

func (v *NullableTransactionsSyncRequestOptions) Set(val *TransactionsSyncRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsSyncRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsSyncRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsSyncRequestOptions(val *TransactionsSyncRequestOptions) *NullableTransactionsSyncRequestOptions {
	return &NullableTransactionsSyncRequestOptions{value: val, isSet: true}
}

func (v NullableTransactionsSyncRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsSyncRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


