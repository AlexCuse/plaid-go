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
)

// TransactionsRecurringGetRequestOptions An optional object to be used with the request. If specified, `options` must not be `null`.
type TransactionsRecurringGetRequestOptions struct {
	// Include the [`personal_finance_category`](https://plaid.com/docs/api/products/transactions/#transactions-get-response-transactions-personal-finance-category) object for each transaction stream in the response.  All implementations are encouraged to set this field to `true` and to use the `personal_finance_category` field instead of `category`. Personal finance categories are the preferred categorization system for transactions, providing higher accuracy and more meaningful categories.  See the [`taxonomy csv file`](https://plaid.com/documents/transactions-personal-finance-category-taxonomy.csv) for a full list of personal finance categories.
	IncludePersonalFinanceCategory *bool `json:"include_personal_finance_category,omitempty"`
}

// NewTransactionsRecurringGetRequestOptions instantiates a new TransactionsRecurringGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRecurringGetRequestOptions() *TransactionsRecurringGetRequestOptions {
	this := TransactionsRecurringGetRequestOptions{}
	var includePersonalFinanceCategory bool = false
	this.IncludePersonalFinanceCategory = &includePersonalFinanceCategory
	return &this
}

// NewTransactionsRecurringGetRequestOptionsWithDefaults instantiates a new TransactionsRecurringGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRecurringGetRequestOptionsWithDefaults() *TransactionsRecurringGetRequestOptions {
	this := TransactionsRecurringGetRequestOptions{}
	var includePersonalFinanceCategory bool = false
	this.IncludePersonalFinanceCategory = &includePersonalFinanceCategory
	return &this
}

// GetIncludePersonalFinanceCategory returns the IncludePersonalFinanceCategory field value if set, zero value otherwise.
func (o *TransactionsRecurringGetRequestOptions) GetIncludePersonalFinanceCategory() bool {
	if o == nil || o.IncludePersonalFinanceCategory == nil {
		var ret bool
		return ret
	}
	return *o.IncludePersonalFinanceCategory
}

// GetIncludePersonalFinanceCategoryOk returns a tuple with the IncludePersonalFinanceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringGetRequestOptions) GetIncludePersonalFinanceCategoryOk() (*bool, bool) {
	if o == nil || o.IncludePersonalFinanceCategory == nil {
		return nil, false
	}
	return o.IncludePersonalFinanceCategory, true
}

// HasIncludePersonalFinanceCategory returns a boolean if a field has been set.
func (o *TransactionsRecurringGetRequestOptions) HasIncludePersonalFinanceCategory() bool {
	if o != nil && o.IncludePersonalFinanceCategory != nil {
		return true
	}

	return false
}

// SetIncludePersonalFinanceCategory gets a reference to the given bool and assigns it to the IncludePersonalFinanceCategory field.
func (o *TransactionsRecurringGetRequestOptions) SetIncludePersonalFinanceCategory(v bool) {
	o.IncludePersonalFinanceCategory = &v
}

func (o TransactionsRecurringGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludePersonalFinanceCategory != nil {
		toSerialize["include_personal_finance_category"] = o.IncludePersonalFinanceCategory
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsRecurringGetRequestOptions struct {
	value *TransactionsRecurringGetRequestOptions
	isSet bool
}

func (v NullableTransactionsRecurringGetRequestOptions) Get() *TransactionsRecurringGetRequestOptions {
	return v.value
}

func (v *NullableTransactionsRecurringGetRequestOptions) Set(val *TransactionsRecurringGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRecurringGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRecurringGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRecurringGetRequestOptions(val *TransactionsRecurringGetRequestOptions) *NullableTransactionsRecurringGetRequestOptions {
	return &NullableTransactionsRecurringGetRequestOptions{value: val, isSet: true}
}

func (v NullableTransactionsRecurringGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRecurringGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


