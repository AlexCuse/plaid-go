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

// ProductAccess The product access being requested. Used to or disallow product access across all accounts. If unset, defaults to all products allowed.
type ProductAccess struct {
	// Allow access to statements. Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	Statements NullableBool `json:"statements,omitempty"`
	// Allow access to the Identity product (name, email, phone, address). Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	Identity NullableBool `json:"identity,omitempty"`
	// Allow access to account number details. Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	Auth NullableBool `json:"auth,omitempty"`
	// Allow access to transaction details. Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	Transactions NullableBool `json:"transactions,omitempty"`
	// Allow access to \"accounts_details_transactions\". Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	AccountsDetailsTransactions NullableBool `json:"accounts_details_transactions,omitempty"`
	// Allow access to \"accounts_routing_number\". Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	AccountsRoutingNumber NullableBool `json:"accounts_routing_number,omitempty"`
	// Allow access to \"accounts_statements\". Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	AccountsStatements NullableBool `json:"accounts_statements,omitempty"`
	// Allow access to \"accounts_tax_statements\". Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	AccountsTaxStatements NullableBool `json:"accounts_tax_statements,omitempty"`
	// Allow access to \"customers_profiles\". Only used by certain partners. If relevant to the partner and unset, defaults to `true`.
	CustomersProfiles NullableBool `json:"customers_profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductAccess ProductAccess

// NewProductAccess instantiates a new ProductAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAccess() *ProductAccess {
	this := ProductAccess{}
	var statements bool = true
	this.Statements = *NewNullableBool(&statements)
	var identity bool = true
	this.Identity = *NewNullableBool(&identity)
	var auth bool = true
	this.Auth = *NewNullableBool(&auth)
	var transactions bool = true
	this.Transactions = *NewNullableBool(&transactions)
	var accountsDetailsTransactions bool = true
	this.AccountsDetailsTransactions = *NewNullableBool(&accountsDetailsTransactions)
	var accountsRoutingNumber bool = true
	this.AccountsRoutingNumber = *NewNullableBool(&accountsRoutingNumber)
	var accountsStatements bool = true
	this.AccountsStatements = *NewNullableBool(&accountsStatements)
	var accountsTaxStatements bool = true
	this.AccountsTaxStatements = *NewNullableBool(&accountsTaxStatements)
	var customersProfiles bool = true
	this.CustomersProfiles = *NewNullableBool(&customersProfiles)
	return &this
}

// NewProductAccessWithDefaults instantiates a new ProductAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAccessWithDefaults() *ProductAccess {
	this := ProductAccess{}
	var statements bool = true
	this.Statements = *NewNullableBool(&statements)
	var identity bool = true
	this.Identity = *NewNullableBool(&identity)
	var auth bool = true
	this.Auth = *NewNullableBool(&auth)
	var transactions bool = true
	this.Transactions = *NewNullableBool(&transactions)
	var accountsDetailsTransactions bool = true
	this.AccountsDetailsTransactions = *NewNullableBool(&accountsDetailsTransactions)
	var accountsRoutingNumber bool = true
	this.AccountsRoutingNumber = *NewNullableBool(&accountsRoutingNumber)
	var accountsStatements bool = true
	this.AccountsStatements = *NewNullableBool(&accountsStatements)
	var accountsTaxStatements bool = true
	this.AccountsTaxStatements = *NewNullableBool(&accountsTaxStatements)
	var customersProfiles bool = true
	this.CustomersProfiles = *NewNullableBool(&customersProfiles)
	return &this
}

// GetStatements returns the Statements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetStatements() bool {
	if o == nil || o.Statements.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Statements.Get()
}

// GetStatementsOk returns a tuple with the Statements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetStatementsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Statements.Get(), o.Statements.IsSet()
}

// HasStatements returns a boolean if a field has been set.
func (o *ProductAccess) HasStatements() bool {
	if o != nil && o.Statements.IsSet() {
		return true
	}

	return false
}

// SetStatements gets a reference to the given NullableBool and assigns it to the Statements field.
func (o *ProductAccess) SetStatements(v bool) {
	o.Statements.Set(&v)
}
// SetStatementsNil sets the value for Statements to be an explicit nil
func (o *ProductAccess) SetStatementsNil() {
	o.Statements.Set(nil)
}

// UnsetStatements ensures that no value is present for Statements, not even an explicit nil
func (o *ProductAccess) UnsetStatements() {
	o.Statements.Unset()
}

// GetIdentity returns the Identity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetIdentity() bool {
	if o == nil || o.Identity.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetIdentityOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// HasIdentity returns a boolean if a field has been set.
func (o *ProductAccess) HasIdentity() bool {
	if o != nil && o.Identity.IsSet() {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NullableBool and assigns it to the Identity field.
func (o *ProductAccess) SetIdentity(v bool) {
	o.Identity.Set(&v)
}
// SetIdentityNil sets the value for Identity to be an explicit nil
func (o *ProductAccess) SetIdentityNil() {
	o.Identity.Set(nil)
}

// UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
func (o *ProductAccess) UnsetIdentity() {
	o.Identity.Unset()
}

// GetAuth returns the Auth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetAuth() bool {
	if o == nil || o.Auth.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Auth.Get()
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetAuthOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Auth.Get(), o.Auth.IsSet()
}

// HasAuth returns a boolean if a field has been set.
func (o *ProductAccess) HasAuth() bool {
	if o != nil && o.Auth.IsSet() {
		return true
	}

	return false
}

// SetAuth gets a reference to the given NullableBool and assigns it to the Auth field.
func (o *ProductAccess) SetAuth(v bool) {
	o.Auth.Set(&v)
}
// SetAuthNil sets the value for Auth to be an explicit nil
func (o *ProductAccess) SetAuthNil() {
	o.Auth.Set(nil)
}

// UnsetAuth ensures that no value is present for Auth, not even an explicit nil
func (o *ProductAccess) UnsetAuth() {
	o.Auth.Unset()
}

// GetTransactions returns the Transactions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetTransactions() bool {
	if o == nil || o.Transactions.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Transactions.Get()
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetTransactionsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Transactions.Get(), o.Transactions.IsSet()
}

// HasTransactions returns a boolean if a field has been set.
func (o *ProductAccess) HasTransactions() bool {
	if o != nil && o.Transactions.IsSet() {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given NullableBool and assigns it to the Transactions field.
func (o *ProductAccess) SetTransactions(v bool) {
	o.Transactions.Set(&v)
}
// SetTransactionsNil sets the value for Transactions to be an explicit nil
func (o *ProductAccess) SetTransactionsNil() {
	o.Transactions.Set(nil)
}

// UnsetTransactions ensures that no value is present for Transactions, not even an explicit nil
func (o *ProductAccess) UnsetTransactions() {
	o.Transactions.Unset()
}

// GetAccountsDetailsTransactions returns the AccountsDetailsTransactions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetAccountsDetailsTransactions() bool {
	if o == nil || o.AccountsDetailsTransactions.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AccountsDetailsTransactions.Get()
}

// GetAccountsDetailsTransactionsOk returns a tuple with the AccountsDetailsTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetAccountsDetailsTransactionsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountsDetailsTransactions.Get(), o.AccountsDetailsTransactions.IsSet()
}

// HasAccountsDetailsTransactions returns a boolean if a field has been set.
func (o *ProductAccess) HasAccountsDetailsTransactions() bool {
	if o != nil && o.AccountsDetailsTransactions.IsSet() {
		return true
	}

	return false
}

// SetAccountsDetailsTransactions gets a reference to the given NullableBool and assigns it to the AccountsDetailsTransactions field.
func (o *ProductAccess) SetAccountsDetailsTransactions(v bool) {
	o.AccountsDetailsTransactions.Set(&v)
}
// SetAccountsDetailsTransactionsNil sets the value for AccountsDetailsTransactions to be an explicit nil
func (o *ProductAccess) SetAccountsDetailsTransactionsNil() {
	o.AccountsDetailsTransactions.Set(nil)
}

// UnsetAccountsDetailsTransactions ensures that no value is present for AccountsDetailsTransactions, not even an explicit nil
func (o *ProductAccess) UnsetAccountsDetailsTransactions() {
	o.AccountsDetailsTransactions.Unset()
}

// GetAccountsRoutingNumber returns the AccountsRoutingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetAccountsRoutingNumber() bool {
	if o == nil || o.AccountsRoutingNumber.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AccountsRoutingNumber.Get()
}

// GetAccountsRoutingNumberOk returns a tuple with the AccountsRoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetAccountsRoutingNumberOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountsRoutingNumber.Get(), o.AccountsRoutingNumber.IsSet()
}

// HasAccountsRoutingNumber returns a boolean if a field has been set.
func (o *ProductAccess) HasAccountsRoutingNumber() bool {
	if o != nil && o.AccountsRoutingNumber.IsSet() {
		return true
	}

	return false
}

// SetAccountsRoutingNumber gets a reference to the given NullableBool and assigns it to the AccountsRoutingNumber field.
func (o *ProductAccess) SetAccountsRoutingNumber(v bool) {
	o.AccountsRoutingNumber.Set(&v)
}
// SetAccountsRoutingNumberNil sets the value for AccountsRoutingNumber to be an explicit nil
func (o *ProductAccess) SetAccountsRoutingNumberNil() {
	o.AccountsRoutingNumber.Set(nil)
}

// UnsetAccountsRoutingNumber ensures that no value is present for AccountsRoutingNumber, not even an explicit nil
func (o *ProductAccess) UnsetAccountsRoutingNumber() {
	o.AccountsRoutingNumber.Unset()
}

// GetAccountsStatements returns the AccountsStatements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetAccountsStatements() bool {
	if o == nil || o.AccountsStatements.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AccountsStatements.Get()
}

// GetAccountsStatementsOk returns a tuple with the AccountsStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetAccountsStatementsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountsStatements.Get(), o.AccountsStatements.IsSet()
}

// HasAccountsStatements returns a boolean if a field has been set.
func (o *ProductAccess) HasAccountsStatements() bool {
	if o != nil && o.AccountsStatements.IsSet() {
		return true
	}

	return false
}

// SetAccountsStatements gets a reference to the given NullableBool and assigns it to the AccountsStatements field.
func (o *ProductAccess) SetAccountsStatements(v bool) {
	o.AccountsStatements.Set(&v)
}
// SetAccountsStatementsNil sets the value for AccountsStatements to be an explicit nil
func (o *ProductAccess) SetAccountsStatementsNil() {
	o.AccountsStatements.Set(nil)
}

// UnsetAccountsStatements ensures that no value is present for AccountsStatements, not even an explicit nil
func (o *ProductAccess) UnsetAccountsStatements() {
	o.AccountsStatements.Unset()
}

// GetAccountsTaxStatements returns the AccountsTaxStatements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetAccountsTaxStatements() bool {
	if o == nil || o.AccountsTaxStatements.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AccountsTaxStatements.Get()
}

// GetAccountsTaxStatementsOk returns a tuple with the AccountsTaxStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetAccountsTaxStatementsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountsTaxStatements.Get(), o.AccountsTaxStatements.IsSet()
}

// HasAccountsTaxStatements returns a boolean if a field has been set.
func (o *ProductAccess) HasAccountsTaxStatements() bool {
	if o != nil && o.AccountsTaxStatements.IsSet() {
		return true
	}

	return false
}

// SetAccountsTaxStatements gets a reference to the given NullableBool and assigns it to the AccountsTaxStatements field.
func (o *ProductAccess) SetAccountsTaxStatements(v bool) {
	o.AccountsTaxStatements.Set(&v)
}
// SetAccountsTaxStatementsNil sets the value for AccountsTaxStatements to be an explicit nil
func (o *ProductAccess) SetAccountsTaxStatementsNil() {
	o.AccountsTaxStatements.Set(nil)
}

// UnsetAccountsTaxStatements ensures that no value is present for AccountsTaxStatements, not even an explicit nil
func (o *ProductAccess) UnsetAccountsTaxStatements() {
	o.AccountsTaxStatements.Unset()
}

// GetCustomersProfiles returns the CustomersProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAccess) GetCustomersProfiles() bool {
	if o == nil || o.CustomersProfiles.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CustomersProfiles.Get()
}

// GetCustomersProfilesOk returns a tuple with the CustomersProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAccess) GetCustomersProfilesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomersProfiles.Get(), o.CustomersProfiles.IsSet()
}

// HasCustomersProfiles returns a boolean if a field has been set.
func (o *ProductAccess) HasCustomersProfiles() bool {
	if o != nil && o.CustomersProfiles.IsSet() {
		return true
	}

	return false
}

// SetCustomersProfiles gets a reference to the given NullableBool and assigns it to the CustomersProfiles field.
func (o *ProductAccess) SetCustomersProfiles(v bool) {
	o.CustomersProfiles.Set(&v)
}
// SetCustomersProfilesNil sets the value for CustomersProfiles to be an explicit nil
func (o *ProductAccess) SetCustomersProfilesNil() {
	o.CustomersProfiles.Set(nil)
}

// UnsetCustomersProfiles ensures that no value is present for CustomersProfiles, not even an explicit nil
func (o *ProductAccess) UnsetCustomersProfiles() {
	o.CustomersProfiles.Unset()
}

func (o ProductAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Statements.IsSet() {
		toSerialize["statements"] = o.Statements.Get()
	}
	if o.Identity.IsSet() {
		toSerialize["identity"] = o.Identity.Get()
	}
	if o.Auth.IsSet() {
		toSerialize["auth"] = o.Auth.Get()
	}
	if o.Transactions.IsSet() {
		toSerialize["transactions"] = o.Transactions.Get()
	}
	if o.AccountsDetailsTransactions.IsSet() {
		toSerialize["accounts_details_transactions"] = o.AccountsDetailsTransactions.Get()
	}
	if o.AccountsRoutingNumber.IsSet() {
		toSerialize["accounts_routing_number"] = o.AccountsRoutingNumber.Get()
	}
	if o.AccountsStatements.IsSet() {
		toSerialize["accounts_statements"] = o.AccountsStatements.Get()
	}
	if o.AccountsTaxStatements.IsSet() {
		toSerialize["accounts_tax_statements"] = o.AccountsTaxStatements.Get()
	}
	if o.CustomersProfiles.IsSet() {
		toSerialize["customers_profiles"] = o.CustomersProfiles.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProductAccess) UnmarshalJSON(bytes []byte) (err error) {
	varProductAccess := _ProductAccess{}

	if err = json.Unmarshal(bytes, &varProductAccess); err == nil {
		*o = ProductAccess(varProductAccess)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "statements")
		delete(additionalProperties, "identity")
		delete(additionalProperties, "auth")
		delete(additionalProperties, "transactions")
		delete(additionalProperties, "accounts_details_transactions")
		delete(additionalProperties, "accounts_routing_number")
		delete(additionalProperties, "accounts_statements")
		delete(additionalProperties, "accounts_tax_statements")
		delete(additionalProperties, "customers_profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductAccess struct {
	value *ProductAccess
	isSet bool
}

func (v NullableProductAccess) Get() *ProductAccess {
	return v.value
}

func (v *NullableProductAccess) Set(val *ProductAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAccess(val *ProductAccess) *NullableProductAccess {
	return &NullableProductAccess{value: val, isSet: true}
}

func (v NullableProductAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


