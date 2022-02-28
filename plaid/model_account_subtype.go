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
	"fmt"
)

// AccountSubtype See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.
type AccountSubtype string

// List of AccountSubtype
const (
	ACCOUNTSUBTYPE__401A AccountSubtype = "401a"
	ACCOUNTSUBTYPE__401K AccountSubtype = "401k"
	ACCOUNTSUBTYPE__403_B AccountSubtype = "403B"
	ACCOUNTSUBTYPE__457B AccountSubtype = "457b"
	ACCOUNTSUBTYPE__529 AccountSubtype = "529"
	ACCOUNTSUBTYPE_BROKERAGE AccountSubtype = "brokerage"
	ACCOUNTSUBTYPE_CASH_ISA AccountSubtype = "cash isa"
	ACCOUNTSUBTYPE_EDUCATION_SAVINGS_ACCOUNT AccountSubtype = "education savings account"
	ACCOUNTSUBTYPE_EBT AccountSubtype = "ebt"
	ACCOUNTSUBTYPE_FIXED_ANNUITY AccountSubtype = "fixed annuity"
	ACCOUNTSUBTYPE_GIC AccountSubtype = "gic"
	ACCOUNTSUBTYPE_HEALTH_REIMBURSEMENT_ARRANGEMENT AccountSubtype = "health reimbursement arrangement"
	ACCOUNTSUBTYPE_HSA AccountSubtype = "hsa"
	ACCOUNTSUBTYPE_ISA AccountSubtype = "isa"
	ACCOUNTSUBTYPE_IRA AccountSubtype = "ira"
	ACCOUNTSUBTYPE_LIF AccountSubtype = "lif"
	ACCOUNTSUBTYPE_LIFE_INSURANCE AccountSubtype = "life insurance"
	ACCOUNTSUBTYPE_LIRA AccountSubtype = "lira"
	ACCOUNTSUBTYPE_LRIF AccountSubtype = "lrif"
	ACCOUNTSUBTYPE_LRSP AccountSubtype = "lrsp"
	ACCOUNTSUBTYPE_NON_TAXABLE_BROKERAGE_ACCOUNT AccountSubtype = "non-taxable brokerage account"
	ACCOUNTSUBTYPE_OTHER AccountSubtype = "other"
	ACCOUNTSUBTYPE_OTHER_INSURANCE AccountSubtype = "other insurance"
	ACCOUNTSUBTYPE_OTHER_ANNUITY AccountSubtype = "other annuity"
	ACCOUNTSUBTYPE_PRIF AccountSubtype = "prif"
	ACCOUNTSUBTYPE_RDSP AccountSubtype = "rdsp"
	ACCOUNTSUBTYPE_RESP AccountSubtype = "resp"
	ACCOUNTSUBTYPE_RLIF AccountSubtype = "rlif"
	ACCOUNTSUBTYPE_RRIF AccountSubtype = "rrif"
	ACCOUNTSUBTYPE_PENSION AccountSubtype = "pension"
	ACCOUNTSUBTYPE_PROFIT_SHARING_PLAN AccountSubtype = "profit sharing plan"
	ACCOUNTSUBTYPE_RETIREMENT AccountSubtype = "retirement"
	ACCOUNTSUBTYPE_ROTH AccountSubtype = "roth"
	ACCOUNTSUBTYPE_ROTH_401K AccountSubtype = "roth 401k"
	ACCOUNTSUBTYPE_RRSP AccountSubtype = "rrsp"
	ACCOUNTSUBTYPE_SEP_IRA AccountSubtype = "sep ira"
	ACCOUNTSUBTYPE_SIMPLE_IRA AccountSubtype = "simple ira"
	ACCOUNTSUBTYPE_SIPP AccountSubtype = "sipp"
	ACCOUNTSUBTYPE_STOCK_PLAN AccountSubtype = "stock plan"
	ACCOUNTSUBTYPE_THRIFT_SAVINGS_PLAN AccountSubtype = "thrift savings plan"
	ACCOUNTSUBTYPE_TFSA AccountSubtype = "tfsa"
	ACCOUNTSUBTYPE_TRUST AccountSubtype = "trust"
	ACCOUNTSUBTYPE_UGMA AccountSubtype = "ugma"
	ACCOUNTSUBTYPE_UTMA AccountSubtype = "utma"
	ACCOUNTSUBTYPE_VARIABLE_ANNUITY AccountSubtype = "variable annuity"
	ACCOUNTSUBTYPE_CREDIT_CARD AccountSubtype = "credit card"
	ACCOUNTSUBTYPE_PAYPAL AccountSubtype = "paypal"
	ACCOUNTSUBTYPE_CD AccountSubtype = "cd"
	ACCOUNTSUBTYPE_CHECKING AccountSubtype = "checking"
	ACCOUNTSUBTYPE_SAVINGS AccountSubtype = "savings"
	ACCOUNTSUBTYPE_MONEY_MARKET AccountSubtype = "money market"
	ACCOUNTSUBTYPE_PREPAID AccountSubtype = "prepaid"
	ACCOUNTSUBTYPE_AUTO AccountSubtype = "auto"
	ACCOUNTSUBTYPE_BUSINESS AccountSubtype = "business"
	ACCOUNTSUBTYPE_COMMERCIAL AccountSubtype = "commercial"
	ACCOUNTSUBTYPE_CONSTRUCTION AccountSubtype = "construction"
	ACCOUNTSUBTYPE_CONSUMER AccountSubtype = "consumer"
	ACCOUNTSUBTYPE_HOME_EQUITY AccountSubtype = "home equity"
	ACCOUNTSUBTYPE_LOAN AccountSubtype = "loan"
	ACCOUNTSUBTYPE_MORTGAGE AccountSubtype = "mortgage"
	ACCOUNTSUBTYPE_OVERDRAFT AccountSubtype = "overdraft"
	ACCOUNTSUBTYPE_LINE_OF_CREDIT AccountSubtype = "line of credit"
	ACCOUNTSUBTYPE_STUDENT AccountSubtype = "student"
	ACCOUNTSUBTYPE_CASH_MANAGEMENT AccountSubtype = "cash management"
	ACCOUNTSUBTYPE_KEOGH AccountSubtype = "keogh"
	ACCOUNTSUBTYPE_MUTUAL_FUND AccountSubtype = "mutual fund"
	ACCOUNTSUBTYPE_RECURRING AccountSubtype = "recurring"
	ACCOUNTSUBTYPE_REWARDS AccountSubtype = "rewards"
	ACCOUNTSUBTYPE_SAFE_DEPOSIT AccountSubtype = "safe deposit"
	ACCOUNTSUBTYPE_SARSEP AccountSubtype = "sarsep"
	ACCOUNTSUBTYPE_PAYROLL AccountSubtype = "payroll"
	ACCOUNTSUBTYPE_NULL AccountSubtype = "null"
)

var allowedAccountSubtypeEnumValues = []AccountSubtype{
	"401a",
	"401k",
	"403B",
	"457b",
	"529",
	"brokerage",
	"cash isa",
	"education savings account",
	"ebt",
	"fixed annuity",
	"gic",
	"health reimbursement arrangement",
	"hsa",
	"isa",
	"ira",
	"lif",
	"life insurance",
	"lira",
	"lrif",
	"lrsp",
	"non-taxable brokerage account",
	"other",
	"other insurance",
	"other annuity",
	"prif",
	"rdsp",
	"resp",
	"rlif",
	"rrif",
	"pension",
	"profit sharing plan",
	"retirement",
	"roth",
	"roth 401k",
	"rrsp",
	"sep ira",
	"simple ira",
	"sipp",
	"stock plan",
	"thrift savings plan",
	"tfsa",
	"trust",
	"ugma",
	"utma",
	"variable annuity",
	"credit card",
	"paypal",
	"cd",
	"checking",
	"savings",
	"money market",
	"prepaid",
	"auto",
	"business",
	"commercial",
	"construction",
	"consumer",
	"home equity",
	"loan",
	"mortgage",
	"overdraft",
	"line of credit",
	"student",
	"cash management",
	"keogh",
	"mutual fund",
	"recurring",
	"rewards",
	"safe deposit",
	"sarsep",
	"payroll",
	"null",
}

func (v *AccountSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountSubtype(value)
	for _, existing := range allowedAccountSubtypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountSubtype", value)
}

// NewAccountSubtypeFromValue returns a pointer to a valid AccountSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountSubtypeFromValue(v string) (*AccountSubtype, error) {
	ev := AccountSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountSubtype: valid values are %v", v, allowedAccountSubtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountSubtype) IsValid() bool {
	for _, existing := range allowedAccountSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountSubtype value
func (v AccountSubtype) Ptr() *AccountSubtype {
	return &v
}

type NullableAccountSubtype struct {
	value *AccountSubtype
	isSet bool
}

func (v NullableAccountSubtype) Get() *AccountSubtype {
	return v.value
}

func (v *NullableAccountSubtype) Set(val *AccountSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSubtype(val *AccountSubtype) *NullableAccountSubtype {
	return &NullableAccountSubtype{value: val, isSet: true}
}

func (v NullableAccountSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

