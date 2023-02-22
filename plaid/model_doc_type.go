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
	"fmt"
)

// DocType The type of document.  `DOCUMENT_TYPE_PAYSTUB`: A paystub.  `DOCUMENT_TYPE_BANK_STATEMENT`: A bank statement.  `DOCUMENT_TYPE_US_TAX_W2`: A W-2 wage and tax statement provided by a US employer reflecting wages earned by the employee.  `DOCUMENT_TYPE_US_MILITARY_ERAS`: An electronic Retirement Account Statement (eRAS) issued by the US military.  `DOCUMENT_TYPE_US_MILITARY_LES`: A Leave and Earnings Statement (LES) issued by the US military.  `DOCUMENT_TYPE_US_MILITARY_CLES`: A Civilian Leave and Earnings Statement (CLES) issued by the US military.  `DOCUMENT_TYPE_GIG`: Used to indicate that the income is related to gig work. Does not necessarily correspond to a specific document type.  `DOCUMENT_TYPE_NONE`: Used to indicate that there is no underlying document for the data.  `DOCUMENT_TYPE_PLAID_GENERATED_PAYSTUB_PDF`: Used to indicate that the PDF for the paystub was generated by Plaid.  `UNKNOWN`: Document type could not be determined.
type DocType string

var _ = fmt.Printf

// List of DocType
const (
	DOCTYPE_UNKNOWN DocType = "UNKNOWN"
	DOCTYPE_DOCUMENT_TYPE_PAYSTUB DocType = "DOCUMENT_TYPE_PAYSTUB"
	DOCTYPE_DOCUMENT_TYPE_BANK_STATEMENT DocType = "DOCUMENT_TYPE_BANK_STATEMENT"
	DOCTYPE_DOCUMENT_TYPE_US_TAX_W2 DocType = "DOCUMENT_TYPE_US_TAX_W2"
	DOCTYPE_DOCUMENT_TYPE_US_MILITARY_ERAS DocType = "DOCUMENT_TYPE_US_MILITARY_ERAS"
	DOCTYPE_DOCUMENT_TYPE_US_MILITARY_LES DocType = "DOCUMENT_TYPE_US_MILITARY_LES"
	DOCTYPE_DOCUMENT_TYPE_US_MILITARY_CLES DocType = "DOCUMENT_TYPE_US_MILITARY_CLES"
	DOCTYPE_DOCUMENT_TYPE_GIG DocType = "DOCUMENT_TYPE_GIG"
	DOCTYPE_DOCUMENT_TYPE_NONE DocType = "DOCUMENT_TYPE_NONE"
	DOCTYPE_DOCUMENT_TYPE_US_TAX_1099_MISC DocType = "DOCUMENT_TYPE_US_TAX_1099_MISC"
	DOCTYPE_DOCUMENT_TYPE_US_TAX_1099_K DocType = "DOCUMENT_TYPE_US_TAX_1099_K"
	DOCTYPE_DOCUMENT_TYPE_PLAID_GENERATED_PAYSTUB_PDF DocType = "DOCUMENT_TYPE_PLAID_GENERATED_PAYSTUB_PDF"
)

var allowedDocTypeEnumValues = []DocType{
	"UNKNOWN",
	"DOCUMENT_TYPE_PAYSTUB",
	"DOCUMENT_TYPE_BANK_STATEMENT",
	"DOCUMENT_TYPE_US_TAX_W2",
	"DOCUMENT_TYPE_US_MILITARY_ERAS",
	"DOCUMENT_TYPE_US_MILITARY_LES",
	"DOCUMENT_TYPE_US_MILITARY_CLES",
	"DOCUMENT_TYPE_GIG",
	"DOCUMENT_TYPE_NONE",
	"DOCUMENT_TYPE_US_TAX_1099_MISC",
	"DOCUMENT_TYPE_US_TAX_1099_K",
	"DOCUMENT_TYPE_PLAID_GENERATED_PAYSTUB_PDF",
}

func (v *DocType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := DocType(value)


	*v = enumTypeValue
	return nil
}

// NewDocTypeFromValue returns a pointer to a valid DocType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocTypeFromValue(v string) (*DocType, error) {
	ev := DocType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocType) IsValid() bool {
	for _, existing := range allowedDocTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocType value
func (v DocType) Ptr() *DocType {
	return &v
}

type NullableDocType struct {
	value *DocType
	isSet bool
}

func (v NullableDocType) Get() *DocType {
	return v.value
}

func (v *NullableDocType) Set(val *DocType) {
	v.value = val
	v.isSet = true
}

func (v NullableDocType) IsSet() bool {
	return v.isSet
}

func (v *NullableDocType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocType(val *DocType) *NullableDocType {
	return &NullableDocType{value: val, isSet: true}
}

func (v NullableDocType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

