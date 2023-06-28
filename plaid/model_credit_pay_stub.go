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

// CreditPayStub An object representing an end user's pay stub.
type CreditPayStub struct {
	Deductions CreditPayStubDeductions `json:"deductions"`
	// An identifier of the document referenced by the document metadata.
	DocumentId NullableString `json:"document_id"`
	DocumentMetadata CreditDocumentMetadata `json:"document_metadata"`
	Earnings CreditPayStubEarnings `json:"earnings"`
	Employee CreditPayStubEmployee `json:"employee"`
	Employer CreditPayStubEmployer `json:"employer"`
	NetPay CreditPayStubNetPay `json:"net_pay"`
	PayPeriodDetails PayStubPayPeriodDetails `json:"pay_period_details"`
	AdditionalProperties map[string]interface{}
}

type _CreditPayStub CreditPayStub

// NewCreditPayStub instantiates a new CreditPayStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayStub(deductions CreditPayStubDeductions, documentId NullableString, documentMetadata CreditDocumentMetadata, earnings CreditPayStubEarnings, employee CreditPayStubEmployee, employer CreditPayStubEmployer, netPay CreditPayStubNetPay, payPeriodDetails PayStubPayPeriodDetails) *CreditPayStub {
	this := CreditPayStub{}
	this.Deductions = deductions
	this.DocumentId = documentId
	this.DocumentMetadata = documentMetadata
	this.Earnings = earnings
	this.Employee = employee
	this.Employer = employer
	this.NetPay = netPay
	this.PayPeriodDetails = payPeriodDetails
	return &this
}

// NewCreditPayStubWithDefaults instantiates a new CreditPayStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayStubWithDefaults() *CreditPayStub {
	this := CreditPayStub{}
	return &this
}

// GetDeductions returns the Deductions field value
func (o *CreditPayStub) GetDeductions() CreditPayStubDeductions {
	if o == nil {
		var ret CreditPayStubDeductions
		return ret
	}

	return o.Deductions
}

// GetDeductionsOk returns a tuple with the Deductions field value
// and a boolean to check if the value has been set.
func (o *CreditPayStub) GetDeductionsOk() (*CreditPayStubDeductions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Deductions, true
}

// SetDeductions sets field value
func (o *CreditPayStub) SetDeductions(v CreditPayStubDeductions) {
	o.Deductions = v
}

// GetDocumentId returns the DocumentId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPayStub) GetDocumentId() string {
	if o == nil || o.DocumentId.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentId.Get()
}

// GetDocumentIdOk returns a tuple with the DocumentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStub) GetDocumentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentId.Get(), o.DocumentId.IsSet()
}

// SetDocumentId sets field value
func (o *CreditPayStub) SetDocumentId(v string) {
	o.DocumentId.Set(&v)
}

// GetDocumentMetadata returns the DocumentMetadata field value
func (o *CreditPayStub) GetDocumentMetadata() CreditDocumentMetadata {
	if o == nil {
		var ret CreditDocumentMetadata
		return ret
	}

	return o.DocumentMetadata
}

// GetDocumentMetadataOk returns a tuple with the DocumentMetadata field value
// and a boolean to check if the value has been set.
func (o *CreditPayStub) GetDocumentMetadataOk() (*CreditDocumentMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentMetadata, true
}

// SetDocumentMetadata sets field value
func (o *CreditPayStub) SetDocumentMetadata(v CreditDocumentMetadata) {
	o.DocumentMetadata = v
}

// GetEarnings returns the Earnings field value
func (o *CreditPayStub) GetEarnings() CreditPayStubEarnings {
	if o == nil {
		var ret CreditPayStubEarnings
		return ret
	}

	return o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value
// and a boolean to check if the value has been set.
func (o *CreditPayStub) GetEarningsOk() (*CreditPayStubEarnings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Earnings, true
}

// SetEarnings sets field value
func (o *CreditPayStub) SetEarnings(v CreditPayStubEarnings) {
	o.Earnings = v
}

// GetEmployee returns the Employee field value
func (o *CreditPayStub) GetEmployee() CreditPayStubEmployee {
	if o == nil {
		var ret CreditPayStubEmployee
		return ret
	}

	return o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value
// and a boolean to check if the value has been set.
func (o *CreditPayStub) GetEmployeeOk() (*CreditPayStubEmployee, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employee, true
}

// SetEmployee sets field value
func (o *CreditPayStub) SetEmployee(v CreditPayStubEmployee) {
	o.Employee = v
}

// GetEmployer returns the Employer field value
func (o *CreditPayStub) GetEmployer() CreditPayStubEmployer {
	if o == nil {
		var ret CreditPayStubEmployer
		return ret
	}

	return o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
func (o *CreditPayStub) GetEmployerOk() (*CreditPayStubEmployer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employer, true
}

// SetEmployer sets field value
func (o *CreditPayStub) SetEmployer(v CreditPayStubEmployer) {
	o.Employer = v
}

// GetNetPay returns the NetPay field value
func (o *CreditPayStub) GetNetPay() CreditPayStubNetPay {
	if o == nil {
		var ret CreditPayStubNetPay
		return ret
	}

	return o.NetPay
}

// GetNetPayOk returns a tuple with the NetPay field value
// and a boolean to check if the value has been set.
func (o *CreditPayStub) GetNetPayOk() (*CreditPayStubNetPay, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetPay, true
}

// SetNetPay sets field value
func (o *CreditPayStub) SetNetPay(v CreditPayStubNetPay) {
	o.NetPay = v
}

// GetPayPeriodDetails returns the PayPeriodDetails field value
func (o *CreditPayStub) GetPayPeriodDetails() PayStubPayPeriodDetails {
	if o == nil {
		var ret PayStubPayPeriodDetails
		return ret
	}

	return o.PayPeriodDetails
}

// GetPayPeriodDetailsOk returns a tuple with the PayPeriodDetails field value
// and a boolean to check if the value has been set.
func (o *CreditPayStub) GetPayPeriodDetailsOk() (*PayStubPayPeriodDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PayPeriodDetails, true
}

// SetPayPeriodDetails sets field value
func (o *CreditPayStub) SetPayPeriodDetails(v PayStubPayPeriodDetails) {
	o.PayPeriodDetails = v
}

func (o CreditPayStub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deductions"] = o.Deductions
	}
	if true {
		toSerialize["document_id"] = o.DocumentId.Get()
	}
	if true {
		toSerialize["document_metadata"] = o.DocumentMetadata
	}
	if true {
		toSerialize["earnings"] = o.Earnings
	}
	if true {
		toSerialize["employee"] = o.Employee
	}
	if true {
		toSerialize["employer"] = o.Employer
	}
	if true {
		toSerialize["net_pay"] = o.NetPay
	}
	if true {
		toSerialize["pay_period_details"] = o.PayPeriodDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditPayStub) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPayStub := _CreditPayStub{}

	if err = json.Unmarshal(bytes, &varCreditPayStub); err == nil {
		*o = CreditPayStub(varCreditPayStub)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "deductions")
		delete(additionalProperties, "document_id")
		delete(additionalProperties, "document_metadata")
		delete(additionalProperties, "earnings")
		delete(additionalProperties, "employee")
		delete(additionalProperties, "employer")
		delete(additionalProperties, "net_pay")
		delete(additionalProperties, "pay_period_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPayStub struct {
	value *CreditPayStub
	isSet bool
}

func (v NullableCreditPayStub) Get() *CreditPayStub {
	return v.value
}

func (v *NullableCreditPayStub) Set(val *CreditPayStub) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayStub) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayStub(val *CreditPayStub) *NullableCreditPayStub {
	return &NullableCreditPayStub{value: val, isSet: true}
}

func (v NullableCreditPayStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


