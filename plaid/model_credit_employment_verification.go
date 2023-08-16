/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.413.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditEmploymentVerification The object containing proof of employment data for an individual.
type CreditEmploymentVerification struct {
	// ID of the payroll provider account.
	AccountId NullableString `json:"account_id"`
	// Current employment status.
	Status NullableString `json:"status"`
	// Start of employment in ISO 8601 format (YYYY-MM-DD).
	StartDate NullableString `json:"start_date"`
	// End of employment, if applicable. Provided in ISO 8601 format (YYY-MM-DD).
	EndDate NullableString `json:"end_date"`
	Employer CreditEmployerVerification `json:"employer"`
	// Current title of employee.
	Title NullableString `json:"title"`
	PlatformIds CreditPlatformIds `json:"platform_ids"`
	// The type of employment for the individual. `\"FULL_TIME\"`: A full-time employee. `\"PART_TIME\"`: A part-time employee. `\"CONTRACTOR\"`: An employee typically hired externally through a contracting group. `\"TEMPORARY\"`: A temporary employee. `\"OTHER\"`: The employee type is not one of the above defined types.
	EmployeeType NullableString `json:"employee_type"`
	// The date of the employee's most recent paystub in ISO 8601 format (YYYY-MM-DD).
	LastPaystubDate NullableString `json:"last_paystub_date"`
	AdditionalProperties map[string]interface{}
}

type _CreditEmploymentVerification CreditEmploymentVerification

// NewCreditEmploymentVerification instantiates a new CreditEmploymentVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditEmploymentVerification(accountId NullableString, status NullableString, startDate NullableString, endDate NullableString, employer CreditEmployerVerification, title NullableString, platformIds CreditPlatformIds, employeeType NullableString, lastPaystubDate NullableString) *CreditEmploymentVerification {
	this := CreditEmploymentVerification{}
	this.AccountId = accountId
	this.Status = status
	this.StartDate = startDate
	this.EndDate = endDate
	this.Employer = employer
	this.Title = title
	this.PlatformIds = platformIds
	this.EmployeeType = employeeType
	this.LastPaystubDate = lastPaystubDate
	return &this
}

// NewCreditEmploymentVerificationWithDefaults instantiates a new CreditEmploymentVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditEmploymentVerificationWithDefaults() *CreditEmploymentVerification {
	this := CreditEmploymentVerification{}
	return &this
}

// GetAccountId returns the AccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditEmploymentVerification) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditEmploymentVerification) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// SetAccountId sets field value
func (o *CreditEmploymentVerification) SetAccountId(v string) {
	o.AccountId.Set(&v)
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditEmploymentVerification) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditEmploymentVerification) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *CreditEmploymentVerification) SetStatus(v string) {
	o.Status.Set(&v)
}

// GetStartDate returns the StartDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditEmploymentVerification) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditEmploymentVerification) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// SetStartDate sets field value
func (o *CreditEmploymentVerification) SetStartDate(v string) {
	o.StartDate.Set(&v)
}

// GetEndDate returns the EndDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditEmploymentVerification) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditEmploymentVerification) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// SetEndDate sets field value
func (o *CreditEmploymentVerification) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// GetEmployer returns the Employer field value
func (o *CreditEmploymentVerification) GetEmployer() CreditEmployerVerification {
	if o == nil {
		var ret CreditEmployerVerification
		return ret
	}

	return o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
func (o *CreditEmploymentVerification) GetEmployerOk() (*CreditEmployerVerification, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employer, true
}

// SetEmployer sets field value
func (o *CreditEmploymentVerification) SetEmployer(v CreditEmployerVerification) {
	o.Employer = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditEmploymentVerification) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditEmploymentVerification) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreditEmploymentVerification) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetPlatformIds returns the PlatformIds field value
func (o *CreditEmploymentVerification) GetPlatformIds() CreditPlatformIds {
	if o == nil {
		var ret CreditPlatformIds
		return ret
	}

	return o.PlatformIds
}

// GetPlatformIdsOk returns a tuple with the PlatformIds field value
// and a boolean to check if the value has been set.
func (o *CreditEmploymentVerification) GetPlatformIdsOk() (*CreditPlatformIds, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlatformIds, true
}

// SetPlatformIds sets field value
func (o *CreditEmploymentVerification) SetPlatformIds(v CreditPlatformIds) {
	o.PlatformIds = v
}

// GetEmployeeType returns the EmployeeType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditEmploymentVerification) GetEmployeeType() string {
	if o == nil || o.EmployeeType.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmployeeType.Get()
}

// GetEmployeeTypeOk returns a tuple with the EmployeeType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditEmploymentVerification) GetEmployeeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeeType.Get(), o.EmployeeType.IsSet()
}

// SetEmployeeType sets field value
func (o *CreditEmploymentVerification) SetEmployeeType(v string) {
	o.EmployeeType.Set(&v)
}

// GetLastPaystubDate returns the LastPaystubDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditEmploymentVerification) GetLastPaystubDate() string {
	if o == nil || o.LastPaystubDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastPaystubDate.Get()
}

// GetLastPaystubDateOk returns a tuple with the LastPaystubDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditEmploymentVerification) GetLastPaystubDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastPaystubDate.Get(), o.LastPaystubDate.IsSet()
}

// SetLastPaystubDate sets field value
func (o *CreditEmploymentVerification) SetLastPaystubDate(v string) {
	o.LastPaystubDate.Set(&v)
}

func (o CreditEmploymentVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if true {
		toSerialize["status"] = o.Status.Get()
	}
	if true {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if true {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if true {
		toSerialize["employer"] = o.Employer
	}
	if true {
		toSerialize["title"] = o.Title.Get()
	}
	if true {
		toSerialize["platform_ids"] = o.PlatformIds
	}
	if true {
		toSerialize["employee_type"] = o.EmployeeType.Get()
	}
	if true {
		toSerialize["last_paystub_date"] = o.LastPaystubDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditEmploymentVerification) UnmarshalJSON(bytes []byte) (err error) {
	varCreditEmploymentVerification := _CreditEmploymentVerification{}

	if err = json.Unmarshal(bytes, &varCreditEmploymentVerification); err == nil {
		*o = CreditEmploymentVerification(varCreditEmploymentVerification)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "employer")
		delete(additionalProperties, "title")
		delete(additionalProperties, "platform_ids")
		delete(additionalProperties, "employee_type")
		delete(additionalProperties, "last_paystub_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditEmploymentVerification struct {
	value *CreditEmploymentVerification
	isSet bool
}

func (v NullableCreditEmploymentVerification) Get() *CreditEmploymentVerification {
	return v.value
}

func (v *NullableCreditEmploymentVerification) Set(val *CreditEmploymentVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditEmploymentVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditEmploymentVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditEmploymentVerification(val *CreditEmploymentVerification) *NullableCreditEmploymentVerification {
	return &NullableCreditEmploymentVerification{value: val, isSet: true}
}

func (v NullableCreditEmploymentVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditEmploymentVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


