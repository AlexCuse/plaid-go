/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.77.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// MortgageLiability Contains details about a mortgage account.
type MortgageLiability struct {
	// The ID of the account that this liability belongs to.
	AccountId string `json:"account_id"`
	// The account number of the loan.
	AccountNumber string `json:"account_number"`
	// The current outstanding amount charged for late payment.
	CurrentLateFee NullableFloat64 `json:"current_late_fee"`
	// Total amount held in escrow to pay taxes and insurance on behalf of the borrower.
	EscrowBalance NullableFloat64 `json:"escrow_balance"`
	// Indicates whether the borrower has private mortgage insurance in effect.
	HasPmi NullableBool `json:"has_pmi"`
	// Indicates whether the borrower will pay a penalty for early payoff of mortgage.
	HasPrepaymentPenalty NullableBool `json:"has_prepayment_penalty"`
	InterestRate MortgageInterestRate `json:"interest_rate"`
	// The amount of the last payment.
	LastPaymentAmount NullableFloat64 `json:"last_payment_amount"`
	// The date of the last payment. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	LastPaymentDate NullableString `json:"last_payment_date"`
	// Description of the type of loan, for example `conventional`, `fixed`, or `variable`. This field is provided directly from the loan servicer and does not have an enumerated set of possible values.
	LoanTypeDescription NullableString `json:"loan_type_description"`
	// Full duration of mortgage as at origination (e.g. `10 year`).
	LoanTerm NullableString `json:"loan_term"`
	// Original date on which mortgage is due in full. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	MaturityDate NullableString `json:"maturity_date"`
	// The amount of the next payment.
	NextMonthlyPayment NullableFloat64 `json:"next_monthly_payment"`
	// The due date for the next payment. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	NextPaymentDueDate NullableString `json:"next_payment_due_date"`
	// The date on which the loan was initially lent. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	OriginationDate NullableString `json:"origination_date"`
	// The original principal balance of the mortgage.
	OriginationPrincipalAmount NullableFloat64 `json:"origination_principal_amount"`
	// Amount of loan (principal + interest) past due for payment.
	PastDueAmount NullableFloat64 `json:"past_due_amount"`
	PropertyAddress MortgagePropertyAddress `json:"property_address"`
	// The year to date (YTD) interest paid.
	YtdInterestPaid NullableFloat64 `json:"ytd_interest_paid"`
	// The YTD principal paid.
	YtdPrincipalPaid NullableFloat64 `json:"ytd_principal_paid"`
	AdditionalProperties map[string]interface{}
}

type _MortgageLiability MortgageLiability

// NewMortgageLiability instantiates a new MortgageLiability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMortgageLiability(accountId string, accountNumber string, currentLateFee NullableFloat64, escrowBalance NullableFloat64, hasPmi NullableBool, hasPrepaymentPenalty NullableBool, interestRate MortgageInterestRate, lastPaymentAmount NullableFloat64, lastPaymentDate NullableString, loanTypeDescription NullableString, loanTerm NullableString, maturityDate NullableString, nextMonthlyPayment NullableFloat64, nextPaymentDueDate NullableString, originationDate NullableString, originationPrincipalAmount NullableFloat64, pastDueAmount NullableFloat64, propertyAddress MortgagePropertyAddress, ytdInterestPaid NullableFloat64, ytdPrincipalPaid NullableFloat64) *MortgageLiability {
	this := MortgageLiability{}
	this.AccountId = accountId
	this.AccountNumber = accountNumber
	this.CurrentLateFee = currentLateFee
	this.EscrowBalance = escrowBalance
	this.HasPmi = hasPmi
	this.HasPrepaymentPenalty = hasPrepaymentPenalty
	this.InterestRate = interestRate
	this.LastPaymentAmount = lastPaymentAmount
	this.LastPaymentDate = lastPaymentDate
	this.LoanTypeDescription = loanTypeDescription
	this.LoanTerm = loanTerm
	this.MaturityDate = maturityDate
	this.NextMonthlyPayment = nextMonthlyPayment
	this.NextPaymentDueDate = nextPaymentDueDate
	this.OriginationDate = originationDate
	this.OriginationPrincipalAmount = originationPrincipalAmount
	this.PastDueAmount = pastDueAmount
	this.PropertyAddress = propertyAddress
	this.YtdInterestPaid = ytdInterestPaid
	this.YtdPrincipalPaid = ytdPrincipalPaid
	return &this
}

// NewMortgageLiabilityWithDefaults instantiates a new MortgageLiability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMortgageLiabilityWithDefaults() *MortgageLiability {
	this := MortgageLiability{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *MortgageLiability) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *MortgageLiability) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *MortgageLiability) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccountNumber returns the AccountNumber field value
func (o *MortgageLiability) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *MortgageLiability) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *MortgageLiability) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetCurrentLateFee returns the CurrentLateFee field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *MortgageLiability) GetCurrentLateFee() float64 {
	if o == nil || o.CurrentLateFee.Get() == nil {
		var ret float64
		return ret
	}

	return *o.CurrentLateFee.Get()
}

// GetCurrentLateFeeOk returns a tuple with the CurrentLateFee field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetCurrentLateFeeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentLateFee.Get(), o.CurrentLateFee.IsSet()
}

// SetCurrentLateFee sets field value
func (o *MortgageLiability) SetCurrentLateFee(v float64) {
	o.CurrentLateFee.Set(&v)
}

// GetEscrowBalance returns the EscrowBalance field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *MortgageLiability) GetEscrowBalance() float64 {
	if o == nil || o.EscrowBalance.Get() == nil {
		var ret float64
		return ret
	}

	return *o.EscrowBalance.Get()
}

// GetEscrowBalanceOk returns a tuple with the EscrowBalance field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetEscrowBalanceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EscrowBalance.Get(), o.EscrowBalance.IsSet()
}

// SetEscrowBalance sets field value
func (o *MortgageLiability) SetEscrowBalance(v float64) {
	o.EscrowBalance.Set(&v)
}

// GetHasPmi returns the HasPmi field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *MortgageLiability) GetHasPmi() bool {
	if o == nil || o.HasPmi.Get() == nil {
		var ret bool
		return ret
	}

	return *o.HasPmi.Get()
}

// GetHasPmiOk returns a tuple with the HasPmi field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetHasPmiOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasPmi.Get(), o.HasPmi.IsSet()
}

// SetHasPmi sets field value
func (o *MortgageLiability) SetHasPmi(v bool) {
	o.HasPmi.Set(&v)
}

// GetHasPrepaymentPenalty returns the HasPrepaymentPenalty field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *MortgageLiability) GetHasPrepaymentPenalty() bool {
	if o == nil || o.HasPrepaymentPenalty.Get() == nil {
		var ret bool
		return ret
	}

	return *o.HasPrepaymentPenalty.Get()
}

// GetHasPrepaymentPenaltyOk returns a tuple with the HasPrepaymentPenalty field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetHasPrepaymentPenaltyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasPrepaymentPenalty.Get(), o.HasPrepaymentPenalty.IsSet()
}

// SetHasPrepaymentPenalty sets field value
func (o *MortgageLiability) SetHasPrepaymentPenalty(v bool) {
	o.HasPrepaymentPenalty.Set(&v)
}

// GetInterestRate returns the InterestRate field value
func (o *MortgageLiability) GetInterestRate() MortgageInterestRate {
	if o == nil {
		var ret MortgageInterestRate
		return ret
	}

	return o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value
// and a boolean to check if the value has been set.
func (o *MortgageLiability) GetInterestRateOk() (*MortgageInterestRate, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InterestRate, true
}

// SetInterestRate sets field value
func (o *MortgageLiability) SetInterestRate(v MortgageInterestRate) {
	o.InterestRate = v
}

// GetLastPaymentAmount returns the LastPaymentAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *MortgageLiability) GetLastPaymentAmount() float64 {
	if o == nil || o.LastPaymentAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.LastPaymentAmount.Get()
}

// GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetLastPaymentAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastPaymentAmount.Get(), o.LastPaymentAmount.IsSet()
}

// SetLastPaymentAmount sets field value
func (o *MortgageLiability) SetLastPaymentAmount(v float64) {
	o.LastPaymentAmount.Set(&v)
}

// GetLastPaymentDate returns the LastPaymentDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MortgageLiability) GetLastPaymentDate() string {
	if o == nil || o.LastPaymentDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastPaymentDate.Get()
}

// GetLastPaymentDateOk returns a tuple with the LastPaymentDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetLastPaymentDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastPaymentDate.Get(), o.LastPaymentDate.IsSet()
}

// SetLastPaymentDate sets field value
func (o *MortgageLiability) SetLastPaymentDate(v string) {
	o.LastPaymentDate.Set(&v)
}

// GetLoanTypeDescription returns the LoanTypeDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MortgageLiability) GetLoanTypeDescription() string {
	if o == nil || o.LoanTypeDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.LoanTypeDescription.Get()
}

// GetLoanTypeDescriptionOk returns a tuple with the LoanTypeDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetLoanTypeDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoanTypeDescription.Get(), o.LoanTypeDescription.IsSet()
}

// SetLoanTypeDescription sets field value
func (o *MortgageLiability) SetLoanTypeDescription(v string) {
	o.LoanTypeDescription.Set(&v)
}

// GetLoanTerm returns the LoanTerm field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MortgageLiability) GetLoanTerm() string {
	if o == nil || o.LoanTerm.Get() == nil {
		var ret string
		return ret
	}

	return *o.LoanTerm.Get()
}

// GetLoanTermOk returns a tuple with the LoanTerm field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetLoanTermOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoanTerm.Get(), o.LoanTerm.IsSet()
}

// SetLoanTerm sets field value
func (o *MortgageLiability) SetLoanTerm(v string) {
	o.LoanTerm.Set(&v)
}

// GetMaturityDate returns the MaturityDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MortgageLiability) GetMaturityDate() string {
	if o == nil || o.MaturityDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.MaturityDate.Get()
}

// GetMaturityDateOk returns a tuple with the MaturityDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetMaturityDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaturityDate.Get(), o.MaturityDate.IsSet()
}

// SetMaturityDate sets field value
func (o *MortgageLiability) SetMaturityDate(v string) {
	o.MaturityDate.Set(&v)
}

// GetNextMonthlyPayment returns the NextMonthlyPayment field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *MortgageLiability) GetNextMonthlyPayment() float64 {
	if o == nil || o.NextMonthlyPayment.Get() == nil {
		var ret float64
		return ret
	}

	return *o.NextMonthlyPayment.Get()
}

// GetNextMonthlyPaymentOk returns a tuple with the NextMonthlyPayment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetNextMonthlyPaymentOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextMonthlyPayment.Get(), o.NextMonthlyPayment.IsSet()
}

// SetNextMonthlyPayment sets field value
func (o *MortgageLiability) SetNextMonthlyPayment(v float64) {
	o.NextMonthlyPayment.Set(&v)
}

// GetNextPaymentDueDate returns the NextPaymentDueDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MortgageLiability) GetNextPaymentDueDate() string {
	if o == nil || o.NextPaymentDueDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextPaymentDueDate.Get()
}

// GetNextPaymentDueDateOk returns a tuple with the NextPaymentDueDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetNextPaymentDueDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextPaymentDueDate.Get(), o.NextPaymentDueDate.IsSet()
}

// SetNextPaymentDueDate sets field value
func (o *MortgageLiability) SetNextPaymentDueDate(v string) {
	o.NextPaymentDueDate.Set(&v)
}

// GetOriginationDate returns the OriginationDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MortgageLiability) GetOriginationDate() string {
	if o == nil || o.OriginationDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginationDate.Get()
}

// GetOriginationDateOk returns a tuple with the OriginationDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetOriginationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationDate.Get(), o.OriginationDate.IsSet()
}

// SetOriginationDate sets field value
func (o *MortgageLiability) SetOriginationDate(v string) {
	o.OriginationDate.Set(&v)
}

// GetOriginationPrincipalAmount returns the OriginationPrincipalAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *MortgageLiability) GetOriginationPrincipalAmount() float64 {
	if o == nil || o.OriginationPrincipalAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.OriginationPrincipalAmount.Get()
}

// GetOriginationPrincipalAmountOk returns a tuple with the OriginationPrincipalAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetOriginationPrincipalAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationPrincipalAmount.Get(), o.OriginationPrincipalAmount.IsSet()
}

// SetOriginationPrincipalAmount sets field value
func (o *MortgageLiability) SetOriginationPrincipalAmount(v float64) {
	o.OriginationPrincipalAmount.Set(&v)
}

// GetPastDueAmount returns the PastDueAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *MortgageLiability) GetPastDueAmount() float64 {
	if o == nil || o.PastDueAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.PastDueAmount.Get()
}

// GetPastDueAmountOk returns a tuple with the PastDueAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetPastDueAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PastDueAmount.Get(), o.PastDueAmount.IsSet()
}

// SetPastDueAmount sets field value
func (o *MortgageLiability) SetPastDueAmount(v float64) {
	o.PastDueAmount.Set(&v)
}

// GetPropertyAddress returns the PropertyAddress field value
func (o *MortgageLiability) GetPropertyAddress() MortgagePropertyAddress {
	if o == nil {
		var ret MortgagePropertyAddress
		return ret
	}

	return o.PropertyAddress
}

// GetPropertyAddressOk returns a tuple with the PropertyAddress field value
// and a boolean to check if the value has been set.
func (o *MortgageLiability) GetPropertyAddressOk() (*MortgagePropertyAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PropertyAddress, true
}

// SetPropertyAddress sets field value
func (o *MortgageLiability) SetPropertyAddress(v MortgagePropertyAddress) {
	o.PropertyAddress = v
}

// GetYtdInterestPaid returns the YtdInterestPaid field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *MortgageLiability) GetYtdInterestPaid() float64 {
	if o == nil || o.YtdInterestPaid.Get() == nil {
		var ret float64
		return ret
	}

	return *o.YtdInterestPaid.Get()
}

// GetYtdInterestPaidOk returns a tuple with the YtdInterestPaid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetYtdInterestPaidOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdInterestPaid.Get(), o.YtdInterestPaid.IsSet()
}

// SetYtdInterestPaid sets field value
func (o *MortgageLiability) SetYtdInterestPaid(v float64) {
	o.YtdInterestPaid.Set(&v)
}

// GetYtdPrincipalPaid returns the YtdPrincipalPaid field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *MortgageLiability) GetYtdPrincipalPaid() float64 {
	if o == nil || o.YtdPrincipalPaid.Get() == nil {
		var ret float64
		return ret
	}

	return *o.YtdPrincipalPaid.Get()
}

// GetYtdPrincipalPaidOk returns a tuple with the YtdPrincipalPaid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MortgageLiability) GetYtdPrincipalPaidOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdPrincipalPaid.Get(), o.YtdPrincipalPaid.IsSet()
}

// SetYtdPrincipalPaid sets field value
func (o *MortgageLiability) SetYtdPrincipalPaid(v float64) {
	o.YtdPrincipalPaid.Set(&v)
}

func (o MortgageLiability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account_number"] = o.AccountNumber
	}
	if true {
		toSerialize["current_late_fee"] = o.CurrentLateFee.Get()
	}
	if true {
		toSerialize["escrow_balance"] = o.EscrowBalance.Get()
	}
	if true {
		toSerialize["has_pmi"] = o.HasPmi.Get()
	}
	if true {
		toSerialize["has_prepayment_penalty"] = o.HasPrepaymentPenalty.Get()
	}
	if true {
		toSerialize["interest_rate"] = o.InterestRate
	}
	if true {
		toSerialize["last_payment_amount"] = o.LastPaymentAmount.Get()
	}
	if true {
		toSerialize["last_payment_date"] = o.LastPaymentDate.Get()
	}
	if true {
		toSerialize["loan_type_description"] = o.LoanTypeDescription.Get()
	}
	if true {
		toSerialize["loan_term"] = o.LoanTerm.Get()
	}
	if true {
		toSerialize["maturity_date"] = o.MaturityDate.Get()
	}
	if true {
		toSerialize["next_monthly_payment"] = o.NextMonthlyPayment.Get()
	}
	if true {
		toSerialize["next_payment_due_date"] = o.NextPaymentDueDate.Get()
	}
	if true {
		toSerialize["origination_date"] = o.OriginationDate.Get()
	}
	if true {
		toSerialize["origination_principal_amount"] = o.OriginationPrincipalAmount.Get()
	}
	if true {
		toSerialize["past_due_amount"] = o.PastDueAmount.Get()
	}
	if true {
		toSerialize["property_address"] = o.PropertyAddress
	}
	if true {
		toSerialize["ytd_interest_paid"] = o.YtdInterestPaid.Get()
	}
	if true {
		toSerialize["ytd_principal_paid"] = o.YtdPrincipalPaid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MortgageLiability) UnmarshalJSON(bytes []byte) (err error) {
	varMortgageLiability := _MortgageLiability{}

	if err = json.Unmarshal(bytes, &varMortgageLiability); err == nil {
		*o = MortgageLiability(varMortgageLiability)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "account_number")
		delete(additionalProperties, "current_late_fee")
		delete(additionalProperties, "escrow_balance")
		delete(additionalProperties, "has_pmi")
		delete(additionalProperties, "has_prepayment_penalty")
		delete(additionalProperties, "interest_rate")
		delete(additionalProperties, "last_payment_amount")
		delete(additionalProperties, "last_payment_date")
		delete(additionalProperties, "loan_type_description")
		delete(additionalProperties, "loan_term")
		delete(additionalProperties, "maturity_date")
		delete(additionalProperties, "next_monthly_payment")
		delete(additionalProperties, "next_payment_due_date")
		delete(additionalProperties, "origination_date")
		delete(additionalProperties, "origination_principal_amount")
		delete(additionalProperties, "past_due_amount")
		delete(additionalProperties, "property_address")
		delete(additionalProperties, "ytd_interest_paid")
		delete(additionalProperties, "ytd_principal_paid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMortgageLiability struct {
	value *MortgageLiability
	isSet bool
}

func (v NullableMortgageLiability) Get() *MortgageLiability {
	return v.value
}

func (v *NullableMortgageLiability) Set(val *MortgageLiability) {
	v.value = val
	v.isSet = true
}

func (v NullableMortgageLiability) IsSet() bool {
	return v.isSet
}

func (v *NullableMortgageLiability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMortgageLiability(val *MortgageLiability) *NullableMortgageLiability {
	return &NullableMortgageLiability{value: val, isSet: true}
}

func (v NullableMortgageLiability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMortgageLiability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


