/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// WalletTransaction The transaction details
type WalletTransaction struct {
	// A unique ID identifying the transaction
	TransactionId string `json:"transaction_id"`
	// The EMI (E-Money Institution) wallet that this payment is associated with, if any. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.
	WalletId string `json:"wallet_id"`
	// A reference for the transaction
	Reference string `json:"reference"`
	// The type of the transaction. The supported transaction types that are returned are: `BANK_TRANSFER:` a transaction which credits an e-wallet through an external bank transfer.  `PAYOUT:` a transaction which debits an e-wallet by disbursing funds to a counterparty.  `PIS_PAY_IN:` a payment which credits an e-wallet through Plaid's Payment Initiation Services (PIS) APIs. For more information see the [Payment Initiation endpoints](https://plaid.com/docs/api/products/payment-initiation/).  `REFUND:` a transaction which debits an e-wallet by refunding a previously initiated payment made through Plaid's [PIS APIs](https://plaid.com/docs/api/products/payment-initiation/).  `FUNDS_SWEEP`: an automated transaction which debits funds from an e-wallet to a designated client-owned account.
	Type string `json:"type"`
	Amount WalletTransactionAmount `json:"amount"`
	Counterparty WalletTransactionCounterparty `json:"counterparty"`
	Status WalletTransactionStatus `json:"status"`
	// Timestamp when the transaction was created, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
	CreatedAt time.Time `json:"created_at"`
	// The date and time of the last time the `status` was updated, in IS0 8601 format
	LastStatusUpdate time.Time `json:"last_status_update"`
	// The payment id that this transaction is associated with, if any. This is present only for transaction types `PIS_PAY_IN` and `REFUND`.
	PaymentId NullableString `json:"payment_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransaction WalletTransaction

// NewWalletTransaction instantiates a new WalletTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransaction(transactionId string, walletId string, reference string, type_ string, amount WalletTransactionAmount, counterparty WalletTransactionCounterparty, status WalletTransactionStatus, createdAt time.Time, lastStatusUpdate time.Time) *WalletTransaction {
	this := WalletTransaction{}
	this.TransactionId = transactionId
	this.WalletId = walletId
	this.Reference = reference
	this.Type = type_
	this.Amount = amount
	this.Counterparty = counterparty
	this.Status = status
	this.CreatedAt = createdAt
	this.LastStatusUpdate = lastStatusUpdate
	return &this
}

// NewWalletTransactionWithDefaults instantiates a new WalletTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionWithDefaults() *WalletTransaction {
	this := WalletTransaction{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *WalletTransaction) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *WalletTransaction) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetWalletId returns the WalletId field value
func (o *WalletTransaction) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetWalletIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *WalletTransaction) SetWalletId(v string) {
	o.WalletId = v
}

// GetReference returns the Reference field value
func (o *WalletTransaction) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *WalletTransaction) SetReference(v string) {
	o.Reference = v
}

// GetType returns the Type field value
func (o *WalletTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WalletTransaction) SetType(v string) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *WalletTransaction) GetAmount() WalletTransactionAmount {
	if o == nil {
		var ret WalletTransactionAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetAmountOk() (*WalletTransactionAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *WalletTransaction) SetAmount(v WalletTransactionAmount) {
	o.Amount = v
}

// GetCounterparty returns the Counterparty field value
func (o *WalletTransaction) GetCounterparty() WalletTransactionCounterparty {
	if o == nil {
		var ret WalletTransactionCounterparty
		return ret
	}

	return o.Counterparty
}

// GetCounterpartyOk returns a tuple with the Counterparty field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetCounterpartyOk() (*WalletTransactionCounterparty, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Counterparty, true
}

// SetCounterparty sets field value
func (o *WalletTransaction) SetCounterparty(v WalletTransactionCounterparty) {
	o.Counterparty = v
}

// GetStatus returns the Status field value
func (o *WalletTransaction) GetStatus() WalletTransactionStatus {
	if o == nil {
		var ret WalletTransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetStatusOk() (*WalletTransactionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WalletTransaction) SetStatus(v WalletTransactionStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WalletTransaction) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WalletTransaction) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetLastStatusUpdate returns the LastStatusUpdate field value
func (o *WalletTransaction) GetLastStatusUpdate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastStatusUpdate
}

// GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetLastStatusUpdateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastStatusUpdate, true
}

// SetLastStatusUpdate sets field value
func (o *WalletTransaction) SetLastStatusUpdate(v time.Time) {
	o.LastStatusUpdate = v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletTransaction) GetPaymentId() string {
	if o == nil || o.PaymentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletTransaction) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// HasPaymentId returns a boolean if a field has been set.
func (o *WalletTransaction) HasPaymentId() bool {
	if o != nil && o.PaymentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given NullableString and assigns it to the PaymentId field.
func (o *WalletTransaction) SetPaymentId(v string) {
	o.PaymentId.Set(&v)
}
// SetPaymentIdNil sets the value for PaymentId to be an explicit nil
func (o *WalletTransaction) SetPaymentIdNil() {
	o.PaymentId.Set(nil)
}

// UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
func (o *WalletTransaction) UnsetPaymentId() {
	o.PaymentId.Unset()
}

func (o WalletTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if true {
		toSerialize["wallet_id"] = o.WalletId
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["counterparty"] = o.Counterparty
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["last_status_update"] = o.LastStatusUpdate
	}
	if o.PaymentId.IsSet() {
		toSerialize["payment_id"] = o.PaymentId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransaction := _WalletTransaction{}

	if err = json.Unmarshal(bytes, &varWalletTransaction); err == nil {
		*o = WalletTransaction(varWalletTransaction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "wallet_id")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "type")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "counterparty")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "last_status_update")
		delete(additionalProperties, "payment_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransaction struct {
	value *WalletTransaction
	isSet bool
}

func (v NullableWalletTransaction) Get() *WalletTransaction {
	return v.value
}

func (v *NullableWalletTransaction) Set(val *WalletTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransaction(val *WalletTransaction) *NullableWalletTransaction {
	return &NullableWalletTransaction{value: val, isSet: true}
}

func (v NullableWalletTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


