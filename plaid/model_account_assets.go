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

// AccountAssets Asset information about an account
type AccountAssets struct {
	// Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  The `account_id` can also change if the `access_token` is deleted and the same credentials that were used to generate that `access_token` are used to generate a new `access_token` on a later date. In that case, the new `account_id` will be different from the old `account_id`.  If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  Like all Plaid identifiers, the `account_id` is case sensitive.
	AccountId string `json:"account_id"`
	Balances AccountBalance `json:"balances"`
	// The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.
	Mask NullableString `json:"mask"`
	// The name of the account, either assigned by the user or by the financial institution itself
	Name string `json:"name"`
	// The official name of the account as given by the financial institution
	OfficialName NullableString `json:"official_name"`
	Type AccountType `json:"type"`
	Subtype NullableAccountSubtype `json:"subtype"`
	// The current verification status of an Auth Item initiated through Automated or Manual micro-deposits.  Returned for Auth Items only.  `pending_automatic_verification`: The Item is pending automatic verification  `pending_manual_verification`: The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the two amounts.  `automatically_verified`: The Item has successfully been automatically verified   `manually_verified`: The Item has successfully been manually verified  `verification_expired`: Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.  `verification_failed`: The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.   
	VerificationStatus *string `json:"verification_status,omitempty"`
	// A unique and persistent identifier for accounts that can be used to trace multiple instances of the same account across different Items for depository accounts. This is currently an opt-in field and only supported for Chase Items.
	PersistentAccountId *string `json:"persistent_account_id,omitempty"`
	// The duration of transaction history available for this Item, typically defined as the time since the date of the earliest transaction in that account. Only returned by Assets endpoints.
	DaysAvailable float32 `json:"days_available"`
	// Transaction history associated with the account. Only returned by Assets endpoints. Transaction history returned by endpoints such as `/transactions/get` or `/investments/transactions/get` will be returned in the top-level `transactions` field instead.
	Transactions []AssetReportTransaction `json:"transactions"`
	// Data returned by the financial institution about the account owner or owners. Only returned by Identity or Assets endpoints. For business accounts, the name reported may be either the name of the individual or the name of the business, depending on the institution. Multiple owners on a single account will be represented in the same `owner` object, not in multiple owner objects within the array. In API versions 2018-05-22 and earlier, the `owners` object is not returned, and instead identity information is returned in the top level `identity` object. For more details, see [Plaid API versioning](https://plaid.com/docs/api/versioning/#version-2019-05-29)
	Owners []Owner `json:"owners"`
	OwnershipType NullableOwnershipType `json:"ownership_type,omitempty"`
	// Calculated data about the historical balances on the account. Only returned by Assets endpoints and currently not supported by `brokerage` or `investment` accounts.
	HistoricalBalances []HistoricalBalance `json:"historical_balances"`
}

// NewAccountAssets instantiates a new AccountAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAssets(accountId string, balances AccountBalance, mask NullableString, name string, officialName NullableString, type_ AccountType, subtype NullableAccountSubtype, daysAvailable float32, transactions []AssetReportTransaction, owners []Owner, historicalBalances []HistoricalBalance) *AccountAssets {
	this := AccountAssets{}
	this.AccountId = accountId
	this.Balances = balances
	this.Mask = mask
	this.Name = name
	this.OfficialName = officialName
	this.Type = type_
	this.Subtype = subtype
	this.DaysAvailable = daysAvailable
	this.Transactions = transactions
	this.Owners = owners
	this.HistoricalBalances = historicalBalances
	return &this
}

// NewAccountAssetsWithDefaults instantiates a new AccountAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAssetsWithDefaults() *AccountAssets {
	this := AccountAssets{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AccountAssets) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountAssets) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountAssets) SetAccountId(v string) {
	o.AccountId = v
}

// GetBalances returns the Balances field value
func (o *AccountAssets) GetBalances() AccountBalance {
	if o == nil {
		var ret AccountBalance
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *AccountAssets) GetBalancesOk() (*AccountBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *AccountAssets) SetBalances(v AccountBalance) {
	o.Balances = v
}

// GetMask returns the Mask field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountAssets) GetMask() string {
	if o == nil || o.Mask.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountAssets) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// SetMask sets field value
func (o *AccountAssets) SetMask(v string) {
	o.Mask.Set(&v)
}

// GetName returns the Name field value
func (o *AccountAssets) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountAssets) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountAssets) SetName(v string) {
	o.Name = v
}

// GetOfficialName returns the OfficialName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountAssets) GetOfficialName() string {
	if o == nil || o.OfficialName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OfficialName.Get()
}

// GetOfficialNameOk returns a tuple with the OfficialName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountAssets) GetOfficialNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OfficialName.Get(), o.OfficialName.IsSet()
}

// SetOfficialName sets field value
func (o *AccountAssets) SetOfficialName(v string) {
	o.OfficialName.Set(&v)
}

// GetType returns the Type field value
func (o *AccountAssets) GetType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountAssets) GetTypeOk() (*AccountType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountAssets) SetType(v AccountType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value
// If the value is explicit nil, the zero value for AccountSubtype will be returned
func (o *AccountAssets) GetSubtype() AccountSubtype {
	if o == nil || o.Subtype.Get() == nil {
		var ret AccountSubtype
		return ret
	}

	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountAssets) GetSubtypeOk() (*AccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// SetSubtype sets field value
func (o *AccountAssets) SetSubtype(v AccountSubtype) {
	o.Subtype.Set(&v)
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *AccountAssets) GetVerificationStatus() string {
	if o == nil || o.VerificationStatus == nil {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAssets) GetVerificationStatusOk() (*string, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *AccountAssets) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *AccountAssets) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

// GetPersistentAccountId returns the PersistentAccountId field value if set, zero value otherwise.
func (o *AccountAssets) GetPersistentAccountId() string {
	if o == nil || o.PersistentAccountId == nil {
		var ret string
		return ret
	}
	return *o.PersistentAccountId
}

// GetPersistentAccountIdOk returns a tuple with the PersistentAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAssets) GetPersistentAccountIdOk() (*string, bool) {
	if o == nil || o.PersistentAccountId == nil {
		return nil, false
	}
	return o.PersistentAccountId, true
}

// HasPersistentAccountId returns a boolean if a field has been set.
func (o *AccountAssets) HasPersistentAccountId() bool {
	if o != nil && o.PersistentAccountId != nil {
		return true
	}

	return false
}

// SetPersistentAccountId gets a reference to the given string and assigns it to the PersistentAccountId field.
func (o *AccountAssets) SetPersistentAccountId(v string) {
	o.PersistentAccountId = &v
}

// GetDaysAvailable returns the DaysAvailable field value
func (o *AccountAssets) GetDaysAvailable() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DaysAvailable
}

// GetDaysAvailableOk returns a tuple with the DaysAvailable field value
// and a boolean to check if the value has been set.
func (o *AccountAssets) GetDaysAvailableOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DaysAvailable, true
}

// SetDaysAvailable sets field value
func (o *AccountAssets) SetDaysAvailable(v float32) {
	o.DaysAvailable = v
}

// GetTransactions returns the Transactions field value
func (o *AccountAssets) GetTransactions() []AssetReportTransaction {
	if o == nil {
		var ret []AssetReportTransaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *AccountAssets) GetTransactionsOk() (*[]AssetReportTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *AccountAssets) SetTransactions(v []AssetReportTransaction) {
	o.Transactions = v
}

// GetOwners returns the Owners field value
func (o *AccountAssets) GetOwners() []Owner {
	if o == nil {
		var ret []Owner
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *AccountAssets) GetOwnersOk() (*[]Owner, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value
func (o *AccountAssets) SetOwners(v []Owner) {
	o.Owners = v
}

// GetOwnershipType returns the OwnershipType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountAssets) GetOwnershipType() OwnershipType {
	if o == nil || o.OwnershipType.Get() == nil {
		var ret OwnershipType
		return ret
	}
	return *o.OwnershipType.Get()
}

// GetOwnershipTypeOk returns a tuple with the OwnershipType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountAssets) GetOwnershipTypeOk() (*OwnershipType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OwnershipType.Get(), o.OwnershipType.IsSet()
}

// HasOwnershipType returns a boolean if a field has been set.
func (o *AccountAssets) HasOwnershipType() bool {
	if o != nil && o.OwnershipType.IsSet() {
		return true
	}

	return false
}

// SetOwnershipType gets a reference to the given NullableOwnershipType and assigns it to the OwnershipType field.
func (o *AccountAssets) SetOwnershipType(v OwnershipType) {
	o.OwnershipType.Set(&v)
}
// SetOwnershipTypeNil sets the value for OwnershipType to be an explicit nil
func (o *AccountAssets) SetOwnershipTypeNil() {
	o.OwnershipType.Set(nil)
}

// UnsetOwnershipType ensures that no value is present for OwnershipType, not even an explicit nil
func (o *AccountAssets) UnsetOwnershipType() {
	o.OwnershipType.Unset()
}

// GetHistoricalBalances returns the HistoricalBalances field value
func (o *AccountAssets) GetHistoricalBalances() []HistoricalBalance {
	if o == nil {
		var ret []HistoricalBalance
		return ret
	}

	return o.HistoricalBalances
}

// GetHistoricalBalancesOk returns a tuple with the HistoricalBalances field value
// and a boolean to check if the value has been set.
func (o *AccountAssets) GetHistoricalBalancesOk() (*[]HistoricalBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HistoricalBalances, true
}

// SetHistoricalBalances sets field value
func (o *AccountAssets) SetHistoricalBalances(v []HistoricalBalance) {
	o.HistoricalBalances = v
}

func (o AccountAssets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["balances"] = o.Balances
	}
	if true {
		toSerialize["mask"] = o.Mask.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["official_name"] = o.OfficialName.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	if o.VerificationStatus != nil {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if o.PersistentAccountId != nil {
		toSerialize["persistent_account_id"] = o.PersistentAccountId
	}
	if true {
		toSerialize["days_available"] = o.DaysAvailable
	}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	if true {
		toSerialize["owners"] = o.Owners
	}
	if o.OwnershipType.IsSet() {
		toSerialize["ownership_type"] = o.OwnershipType.Get()
	}
	if true {
		toSerialize["historical_balances"] = o.HistoricalBalances
	}
	return json.Marshal(toSerialize)
}

type NullableAccountAssets struct {
	value *AccountAssets
	isSet bool
}

func (v NullableAccountAssets) Get() *AccountAssets {
	return v.value
}

func (v *NullableAccountAssets) Set(val *AccountAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAssets(val *AccountAssets) *NullableAccountAssets {
	return &NullableAccountAssets{value: val, isSet: true}
}

func (v NullableAccountAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


