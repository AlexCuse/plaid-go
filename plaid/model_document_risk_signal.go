/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// DocumentRiskSignal Details about a certain reason as to why a document could potentially be fraudulent.
type DocumentRiskSignal struct {
	// The result from the risk signal check.
	Type NullableString `json:"type"`
	// The field which the risk signal was computed for
	Field NullableString `json:"field"`
	// A flag used to quickly identify if the signal indicates that this field is authentic or fraudulent
	HasFraudRisk NullableBool `json:"has_fraud_risk"`
	InstitutionMetadata NullableDocumentRiskSignalInstitutionMetadata `json:"institution_metadata"`
	// The expected value of the field, as seen on the document
	ExpectedValue NullableString `json:"expected_value"`
	// The derived value obtained in the risk signal calculation process for this field
	ActualValue NullableString `json:"actual_value"`
	// A human-readable explanation providing more detail into the particular risk signal
	SignalDescription NullableString `json:"signal_description"`
	AdditionalProperties map[string]interface{}
}

type _DocumentRiskSignal DocumentRiskSignal

// NewDocumentRiskSignal instantiates a new DocumentRiskSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentRiskSignal(type_ NullableString, field NullableString, hasFraudRisk NullableBool, institutionMetadata NullableDocumentRiskSignalInstitutionMetadata, expectedValue NullableString, actualValue NullableString, signalDescription NullableString) *DocumentRiskSignal {
	this := DocumentRiskSignal{}
	this.Type = type_
	this.Field = field
	this.HasFraudRisk = hasFraudRisk
	this.InstitutionMetadata = institutionMetadata
	this.ExpectedValue = expectedValue
	this.ActualValue = actualValue
	this.SignalDescription = signalDescription
	return &this
}

// NewDocumentRiskSignalWithDefaults instantiates a new DocumentRiskSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentRiskSignalWithDefaults() *DocumentRiskSignal {
	this := DocumentRiskSignal{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentRiskSignal) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *DocumentRiskSignal) SetType(v string) {
	o.Type.Set(&v)
}

// GetField returns the Field field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentRiskSignal) GetField() string {
	if o == nil || o.Field.Get() == nil {
		var ret string
		return ret
	}

	return *o.Field.Get()
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetFieldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Field.Get(), o.Field.IsSet()
}

// SetField sets field value
func (o *DocumentRiskSignal) SetField(v string) {
	o.Field.Set(&v)
}

// GetHasFraudRisk returns the HasFraudRisk field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *DocumentRiskSignal) GetHasFraudRisk() bool {
	if o == nil || o.HasFraudRisk.Get() == nil {
		var ret bool
		return ret
	}

	return *o.HasFraudRisk.Get()
}

// GetHasFraudRiskOk returns a tuple with the HasFraudRisk field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetHasFraudRiskOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasFraudRisk.Get(), o.HasFraudRisk.IsSet()
}

// SetHasFraudRisk sets field value
func (o *DocumentRiskSignal) SetHasFraudRisk(v bool) {
	o.HasFraudRisk.Set(&v)
}

// GetInstitutionMetadata returns the InstitutionMetadata field value
// If the value is explicit nil, the zero value for DocumentRiskSignalInstitutionMetadata will be returned
func (o *DocumentRiskSignal) GetInstitutionMetadata() DocumentRiskSignalInstitutionMetadata {
	if o == nil || o.InstitutionMetadata.Get() == nil {
		var ret DocumentRiskSignalInstitutionMetadata
		return ret
	}

	return *o.InstitutionMetadata.Get()
}

// GetInstitutionMetadataOk returns a tuple with the InstitutionMetadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetInstitutionMetadataOk() (*DocumentRiskSignalInstitutionMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstitutionMetadata.Get(), o.InstitutionMetadata.IsSet()
}

// SetInstitutionMetadata sets field value
func (o *DocumentRiskSignal) SetInstitutionMetadata(v DocumentRiskSignalInstitutionMetadata) {
	o.InstitutionMetadata.Set(&v)
}

// GetExpectedValue returns the ExpectedValue field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentRiskSignal) GetExpectedValue() string {
	if o == nil || o.ExpectedValue.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpectedValue.Get()
}

// GetExpectedValueOk returns a tuple with the ExpectedValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetExpectedValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpectedValue.Get(), o.ExpectedValue.IsSet()
}

// SetExpectedValue sets field value
func (o *DocumentRiskSignal) SetExpectedValue(v string) {
	o.ExpectedValue.Set(&v)
}

// GetActualValue returns the ActualValue field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentRiskSignal) GetActualValue() string {
	if o == nil || o.ActualValue.Get() == nil {
		var ret string
		return ret
	}

	return *o.ActualValue.Get()
}

// GetActualValueOk returns a tuple with the ActualValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetActualValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActualValue.Get(), o.ActualValue.IsSet()
}

// SetActualValue sets field value
func (o *DocumentRiskSignal) SetActualValue(v string) {
	o.ActualValue.Set(&v)
}

// GetSignalDescription returns the SignalDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentRiskSignal) GetSignalDescription() string {
	if o == nil || o.SignalDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.SignalDescription.Get()
}

// GetSignalDescriptionOk returns a tuple with the SignalDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentRiskSignal) GetSignalDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SignalDescription.Get(), o.SignalDescription.IsSet()
}

// SetSignalDescription sets field value
func (o *DocumentRiskSignal) SetSignalDescription(v string) {
	o.SignalDescription.Set(&v)
}

func (o DocumentRiskSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if true {
		toSerialize["field"] = o.Field.Get()
	}
	if true {
		toSerialize["has_fraud_risk"] = o.HasFraudRisk.Get()
	}
	if true {
		toSerialize["institution_metadata"] = o.InstitutionMetadata.Get()
	}
	if true {
		toSerialize["expected_value"] = o.ExpectedValue.Get()
	}
	if true {
		toSerialize["actual_value"] = o.ActualValue.Get()
	}
	if true {
		toSerialize["signal_description"] = o.SignalDescription.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DocumentRiskSignal) UnmarshalJSON(bytes []byte) (err error) {
	varDocumentRiskSignal := _DocumentRiskSignal{}

	if err = json.Unmarshal(bytes, &varDocumentRiskSignal); err == nil {
		*o = DocumentRiskSignal(varDocumentRiskSignal)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "field")
		delete(additionalProperties, "has_fraud_risk")
		delete(additionalProperties, "institution_metadata")
		delete(additionalProperties, "expected_value")
		delete(additionalProperties, "actual_value")
		delete(additionalProperties, "signal_description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentRiskSignal struct {
	value *DocumentRiskSignal
	isSet bool
}

func (v NullableDocumentRiskSignal) Get() *DocumentRiskSignal {
	return v.value
}

func (v *NullableDocumentRiskSignal) Set(val *DocumentRiskSignal) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentRiskSignal) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentRiskSignal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentRiskSignal(val *DocumentRiskSignal) *NullableDocumentRiskSignal {
	return &NullableDocumentRiskSignal{value: val, isSet: true}
}

func (v NullableDocumentRiskSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentRiskSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


