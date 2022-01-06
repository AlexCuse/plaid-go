/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.61.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// InvestmentsTransactionsGetResponse InvestmentsTransactionsGetResponse defines the response schema for `/investments/transactions/get`
type InvestmentsTransactionsGetResponse struct {
	Item Item `json:"item"`
	// The accounts for which transaction history is being fetched.
	Accounts []AccountBase `json:"accounts"`
	// All securities for which there is a corresponding transaction being fetched.
	Securities []Security `json:"securities"`
	// The transactions being fetched
	InvestmentTransactions []InvestmentTransaction `json:"investment_transactions"`
	// The total number of transactions available within the date range specified. If `total_investment_transactions` is larger than the size of the `transactions` array, more transactions are available and can be fetched via manipulating the `offset` parameter.'
	TotalInvestmentTransactions int32 `json:"total_investment_transactions"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _InvestmentsTransactionsGetResponse InvestmentsTransactionsGetResponse

// NewInvestmentsTransactionsGetResponse instantiates a new InvestmentsTransactionsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsTransactionsGetResponse(item Item, accounts []AccountBase, securities []Security, investmentTransactions []InvestmentTransaction, totalInvestmentTransactions int32, requestId string) *InvestmentsTransactionsGetResponse {
	this := InvestmentsTransactionsGetResponse{}
	this.Item = item
	this.Accounts = accounts
	this.Securities = securities
	this.InvestmentTransactions = investmentTransactions
	this.TotalInvestmentTransactions = totalInvestmentTransactions
	this.RequestId = requestId
	return &this
}

// NewInvestmentsTransactionsGetResponseWithDefaults instantiates a new InvestmentsTransactionsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsTransactionsGetResponseWithDefaults() *InvestmentsTransactionsGetResponse {
	this := InvestmentsTransactionsGetResponse{}
	return &this
}

// GetItem returns the Item field value
func (o *InvestmentsTransactionsGetResponse) GetItem() Item {
	if o == nil {
		var ret Item
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetResponse) GetItemOk() (*Item, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *InvestmentsTransactionsGetResponse) SetItem(v Item) {
	o.Item = v
}

// GetAccounts returns the Accounts field value
func (o *InvestmentsTransactionsGetResponse) GetAccounts() []AccountBase {
	if o == nil {
		var ret []AccountBase
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetResponse) GetAccountsOk() (*[]AccountBase, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *InvestmentsTransactionsGetResponse) SetAccounts(v []AccountBase) {
	o.Accounts = v
}

// GetSecurities returns the Securities field value
func (o *InvestmentsTransactionsGetResponse) GetSecurities() []Security {
	if o == nil {
		var ret []Security
		return ret
	}

	return o.Securities
}

// GetSecuritiesOk returns a tuple with the Securities field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetResponse) GetSecuritiesOk() (*[]Security, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Securities, true
}

// SetSecurities sets field value
func (o *InvestmentsTransactionsGetResponse) SetSecurities(v []Security) {
	o.Securities = v
}

// GetInvestmentTransactions returns the InvestmentTransactions field value
func (o *InvestmentsTransactionsGetResponse) GetInvestmentTransactions() []InvestmentTransaction {
	if o == nil {
		var ret []InvestmentTransaction
		return ret
	}

	return o.InvestmentTransactions
}

// GetInvestmentTransactionsOk returns a tuple with the InvestmentTransactions field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetResponse) GetInvestmentTransactionsOk() (*[]InvestmentTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InvestmentTransactions, true
}

// SetInvestmentTransactions sets field value
func (o *InvestmentsTransactionsGetResponse) SetInvestmentTransactions(v []InvestmentTransaction) {
	o.InvestmentTransactions = v
}

// GetTotalInvestmentTransactions returns the TotalInvestmentTransactions field value
func (o *InvestmentsTransactionsGetResponse) GetTotalInvestmentTransactions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalInvestmentTransactions
}

// GetTotalInvestmentTransactionsOk returns a tuple with the TotalInvestmentTransactions field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetResponse) GetTotalInvestmentTransactionsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalInvestmentTransactions, true
}

// SetTotalInvestmentTransactions sets field value
func (o *InvestmentsTransactionsGetResponse) SetTotalInvestmentTransactions(v int32) {
	o.TotalInvestmentTransactions = v
}

// GetRequestId returns the RequestId field value
func (o *InvestmentsTransactionsGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *InvestmentsTransactionsGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o InvestmentsTransactionsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	if true {
		toSerialize["securities"] = o.Securities
	}
	if true {
		toSerialize["investment_transactions"] = o.InvestmentTransactions
	}
	if true {
		toSerialize["total_investment_transactions"] = o.TotalInvestmentTransactions
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InvestmentsTransactionsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInvestmentsTransactionsGetResponse := _InvestmentsTransactionsGetResponse{}

	if err = json.Unmarshal(bytes, &varInvestmentsTransactionsGetResponse); err == nil {
		*o = InvestmentsTransactionsGetResponse(varInvestmentsTransactionsGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item")
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "securities")
		delete(additionalProperties, "investment_transactions")
		delete(additionalProperties, "total_investment_transactions")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvestmentsTransactionsGetResponse struct {
	value *InvestmentsTransactionsGetResponse
	isSet bool
}

func (v NullableInvestmentsTransactionsGetResponse) Get() *InvestmentsTransactionsGetResponse {
	return v.value
}

func (v *NullableInvestmentsTransactionsGetResponse) Set(val *InvestmentsTransactionsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsTransactionsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsTransactionsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsTransactionsGetResponse(val *InvestmentsTransactionsGetResponse) *NullableInvestmentsTransactionsGetResponse {
	return &NullableInvestmentsTransactionsGetResponse{value: val, isSet: true}
}

func (v NullableInvestmentsTransactionsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsTransactionsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


