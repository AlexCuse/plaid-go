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

// PartnerEndCustomerOAuthInstitution The OAuth registration information for an institution.
type PartnerEndCustomerOAuthInstitution struct {
	Name *string `json:"name,omitempty"`
	InstitutionId *string `json:"institution_id,omitempty"`
	Environments *PartnerEndCustomerOAuthInstitutionEnvironments `json:"environments,omitempty"`
	// The date on which the end customer's application was approved by the institution, or an empty string if their application has not yet been approved.
	ProductionEnablementDate NullableString `json:"production_enablement_date,omitempty"`
	// The date on which non-OAuth Item adds will no longer be supported for this institution, or an empty string if no such date has been set by the institution.
	ClassicDisablementDate NullableString `json:"classic_disablement_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerEndCustomerOAuthInstitution PartnerEndCustomerOAuthInstitution

// NewPartnerEndCustomerOAuthInstitution instantiates a new PartnerEndCustomerOAuthInstitution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomerOAuthInstitution() *PartnerEndCustomerOAuthInstitution {
	this := PartnerEndCustomerOAuthInstitution{}
	return &this
}

// NewPartnerEndCustomerOAuthInstitutionWithDefaults instantiates a new PartnerEndCustomerOAuthInstitution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerOAuthInstitutionWithDefaults() *PartnerEndCustomerOAuthInstitution {
	this := PartnerEndCustomerOAuthInstitution{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartnerEndCustomerOAuthInstitution) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthInstitution) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartnerEndCustomerOAuthInstitution) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartnerEndCustomerOAuthInstitution) SetName(v string) {
	o.Name = &v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *PartnerEndCustomerOAuthInstitution) GetInstitutionId() string {
	if o == nil || o.InstitutionId == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthInstitution) GetInstitutionIdOk() (*string, bool) {
	if o == nil || o.InstitutionId == nil {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *PartnerEndCustomerOAuthInstitution) HasInstitutionId() bool {
	if o != nil && o.InstitutionId != nil {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *PartnerEndCustomerOAuthInstitution) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *PartnerEndCustomerOAuthInstitution) GetEnvironments() PartnerEndCustomerOAuthInstitutionEnvironments {
	if o == nil || o.Environments == nil {
		var ret PartnerEndCustomerOAuthInstitutionEnvironments
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerOAuthInstitution) GetEnvironmentsOk() (*PartnerEndCustomerOAuthInstitutionEnvironments, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *PartnerEndCustomerOAuthInstitution) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given PartnerEndCustomerOAuthInstitutionEnvironments and assigns it to the Environments field.
func (o *PartnerEndCustomerOAuthInstitution) SetEnvironments(v PartnerEndCustomerOAuthInstitutionEnvironments) {
	o.Environments = &v
}

// GetProductionEnablementDate returns the ProductionEnablementDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerEndCustomerOAuthInstitution) GetProductionEnablementDate() string {
	if o == nil || o.ProductionEnablementDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProductionEnablementDate.Get()
}

// GetProductionEnablementDateOk returns a tuple with the ProductionEnablementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerEndCustomerOAuthInstitution) GetProductionEnablementDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProductionEnablementDate.Get(), o.ProductionEnablementDate.IsSet()
}

// HasProductionEnablementDate returns a boolean if a field has been set.
func (o *PartnerEndCustomerOAuthInstitution) HasProductionEnablementDate() bool {
	if o != nil && o.ProductionEnablementDate.IsSet() {
		return true
	}

	return false
}

// SetProductionEnablementDate gets a reference to the given NullableString and assigns it to the ProductionEnablementDate field.
func (o *PartnerEndCustomerOAuthInstitution) SetProductionEnablementDate(v string) {
	o.ProductionEnablementDate.Set(&v)
}
// SetProductionEnablementDateNil sets the value for ProductionEnablementDate to be an explicit nil
func (o *PartnerEndCustomerOAuthInstitution) SetProductionEnablementDateNil() {
	o.ProductionEnablementDate.Set(nil)
}

// UnsetProductionEnablementDate ensures that no value is present for ProductionEnablementDate, not even an explicit nil
func (o *PartnerEndCustomerOAuthInstitution) UnsetProductionEnablementDate() {
	o.ProductionEnablementDate.Unset()
}

// GetClassicDisablementDate returns the ClassicDisablementDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerEndCustomerOAuthInstitution) GetClassicDisablementDate() string {
	if o == nil || o.ClassicDisablementDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClassicDisablementDate.Get()
}

// GetClassicDisablementDateOk returns a tuple with the ClassicDisablementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerEndCustomerOAuthInstitution) GetClassicDisablementDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClassicDisablementDate.Get(), o.ClassicDisablementDate.IsSet()
}

// HasClassicDisablementDate returns a boolean if a field has been set.
func (o *PartnerEndCustomerOAuthInstitution) HasClassicDisablementDate() bool {
	if o != nil && o.ClassicDisablementDate.IsSet() {
		return true
	}

	return false
}

// SetClassicDisablementDate gets a reference to the given NullableString and assigns it to the ClassicDisablementDate field.
func (o *PartnerEndCustomerOAuthInstitution) SetClassicDisablementDate(v string) {
	o.ClassicDisablementDate.Set(&v)
}
// SetClassicDisablementDateNil sets the value for ClassicDisablementDate to be an explicit nil
func (o *PartnerEndCustomerOAuthInstitution) SetClassicDisablementDateNil() {
	o.ClassicDisablementDate.Set(nil)
}

// UnsetClassicDisablementDate ensures that no value is present for ClassicDisablementDate, not even an explicit nil
func (o *PartnerEndCustomerOAuthInstitution) UnsetClassicDisablementDate() {
	o.ClassicDisablementDate.Unset()
}

func (o PartnerEndCustomerOAuthInstitution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.InstitutionId != nil {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	if o.ProductionEnablementDate.IsSet() {
		toSerialize["production_enablement_date"] = o.ProductionEnablementDate.Get()
	}
	if o.ClassicDisablementDate.IsSet() {
		toSerialize["classic_disablement_date"] = o.ClassicDisablementDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerEndCustomerOAuthInstitution) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerEndCustomerOAuthInstitution := _PartnerEndCustomerOAuthInstitution{}

	if err = json.Unmarshal(bytes, &varPartnerEndCustomerOAuthInstitution); err == nil {
		*o = PartnerEndCustomerOAuthInstitution(varPartnerEndCustomerOAuthInstitution)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "institution_id")
		delete(additionalProperties, "environments")
		delete(additionalProperties, "production_enablement_date")
		delete(additionalProperties, "classic_disablement_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerEndCustomerOAuthInstitution struct {
	value *PartnerEndCustomerOAuthInstitution
	isSet bool
}

func (v NullablePartnerEndCustomerOAuthInstitution) Get() *PartnerEndCustomerOAuthInstitution {
	return v.value
}

func (v *NullablePartnerEndCustomerOAuthInstitution) Set(val *PartnerEndCustomerOAuthInstitution) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerOAuthInstitution) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerOAuthInstitution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerOAuthInstitution(val *PartnerEndCustomerOAuthInstitution) *NullablePartnerEndCustomerOAuthInstitution {
	return &NullablePartnerEndCustomerOAuthInstitution{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerOAuthInstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerOAuthInstitution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


