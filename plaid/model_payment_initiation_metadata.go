/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.131.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentInitiationMetadata Metadata that captures what specific payment configurations an institution supports when making Payment Initiation requests.
type PaymentInitiationMetadata struct {
	// Indicates whether the institution supports payments from a different country.
	SupportsInternationalPayments bool `json:"supports_international_payments"`
	// Indicates whether the institution supports SEPA Instant payments.
	SupportsSepaInstant bool `json:"supports_sepa_instant"`
	// A mapping of currency to maximum payment amount (denominated in the smallest unit of currency) supported by the institution.  Example: `{\"GBP\": \"10000\"}` 
	MaximumPaymentAmount map[string]string `json:"maximum_payment_amount"`
	// Indicates whether the institution supports returning refund details when initiating a payment.
	SupportsRefundDetails bool `json:"supports_refund_details"`
	StandingOrderMetadata NullablePaymentInitiationStandingOrderMetadata `json:"standing_order_metadata"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationMetadata PaymentInitiationMetadata

// NewPaymentInitiationMetadata instantiates a new PaymentInitiationMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationMetadata(supportsInternationalPayments bool, supportsSepaInstant bool, maximumPaymentAmount map[string]string, supportsRefundDetails bool, standingOrderMetadata NullablePaymentInitiationStandingOrderMetadata) *PaymentInitiationMetadata {
	this := PaymentInitiationMetadata{}
	this.SupportsInternationalPayments = supportsInternationalPayments
	this.SupportsSepaInstant = supportsSepaInstant
	this.MaximumPaymentAmount = maximumPaymentAmount
	this.SupportsRefundDetails = supportsRefundDetails
	this.StandingOrderMetadata = standingOrderMetadata
	return &this
}

// NewPaymentInitiationMetadataWithDefaults instantiates a new PaymentInitiationMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationMetadataWithDefaults() *PaymentInitiationMetadata {
	this := PaymentInitiationMetadata{}
	return &this
}

// GetSupportsInternationalPayments returns the SupportsInternationalPayments field value
func (o *PaymentInitiationMetadata) GetSupportsInternationalPayments() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportsInternationalPayments
}

// GetSupportsInternationalPaymentsOk returns a tuple with the SupportsInternationalPayments field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationMetadata) GetSupportsInternationalPaymentsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupportsInternationalPayments, true
}

// SetSupportsInternationalPayments sets field value
func (o *PaymentInitiationMetadata) SetSupportsInternationalPayments(v bool) {
	o.SupportsInternationalPayments = v
}

// GetSupportsSepaInstant returns the SupportsSepaInstant field value
func (o *PaymentInitiationMetadata) GetSupportsSepaInstant() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportsSepaInstant
}

// GetSupportsSepaInstantOk returns a tuple with the SupportsSepaInstant field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationMetadata) GetSupportsSepaInstantOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupportsSepaInstant, true
}

// SetSupportsSepaInstant sets field value
func (o *PaymentInitiationMetadata) SetSupportsSepaInstant(v bool) {
	o.SupportsSepaInstant = v
}

// GetMaximumPaymentAmount returns the MaximumPaymentAmount field value
func (o *PaymentInitiationMetadata) GetMaximumPaymentAmount() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.MaximumPaymentAmount
}

// GetMaximumPaymentAmountOk returns a tuple with the MaximumPaymentAmount field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationMetadata) GetMaximumPaymentAmountOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaximumPaymentAmount, true
}

// SetMaximumPaymentAmount sets field value
func (o *PaymentInitiationMetadata) SetMaximumPaymentAmount(v map[string]string) {
	o.MaximumPaymentAmount = v
}

// GetSupportsRefundDetails returns the SupportsRefundDetails field value
func (o *PaymentInitiationMetadata) GetSupportsRefundDetails() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportsRefundDetails
}

// GetSupportsRefundDetailsOk returns a tuple with the SupportsRefundDetails field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationMetadata) GetSupportsRefundDetailsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupportsRefundDetails, true
}

// SetSupportsRefundDetails sets field value
func (o *PaymentInitiationMetadata) SetSupportsRefundDetails(v bool) {
	o.SupportsRefundDetails = v
}

// GetStandingOrderMetadata returns the StandingOrderMetadata field value
// If the value is explicit nil, the zero value for PaymentInitiationStandingOrderMetadata will be returned
func (o *PaymentInitiationMetadata) GetStandingOrderMetadata() PaymentInitiationStandingOrderMetadata {
	if o == nil || o.StandingOrderMetadata.Get() == nil {
		var ret PaymentInitiationStandingOrderMetadata
		return ret
	}

	return *o.StandingOrderMetadata.Get()
}

// GetStandingOrderMetadataOk returns a tuple with the StandingOrderMetadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationMetadata) GetStandingOrderMetadataOk() (*PaymentInitiationStandingOrderMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StandingOrderMetadata.Get(), o.StandingOrderMetadata.IsSet()
}

// SetStandingOrderMetadata sets field value
func (o *PaymentInitiationMetadata) SetStandingOrderMetadata(v PaymentInitiationStandingOrderMetadata) {
	o.StandingOrderMetadata.Set(&v)
}

func (o PaymentInitiationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supports_international_payments"] = o.SupportsInternationalPayments
	}
	if true {
		toSerialize["supports_sepa_instant"] = o.SupportsSepaInstant
	}
	if true {
		toSerialize["maximum_payment_amount"] = o.MaximumPaymentAmount
	}
	if true {
		toSerialize["supports_refund_details"] = o.SupportsRefundDetails
	}
	if true {
		toSerialize["standing_order_metadata"] = o.StandingOrderMetadata.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationMetadata := _PaymentInitiationMetadata{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationMetadata); err == nil {
		*o = PaymentInitiationMetadata(varPaymentInitiationMetadata)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "supports_international_payments")
		delete(additionalProperties, "supports_sepa_instant")
		delete(additionalProperties, "maximum_payment_amount")
		delete(additionalProperties, "supports_refund_details")
		delete(additionalProperties, "standing_order_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationMetadata struct {
	value *PaymentInitiationMetadata
	isSet bool
}

func (v NullablePaymentInitiationMetadata) Get() *PaymentInitiationMetadata {
	return v.value
}

func (v *NullablePaymentInitiationMetadata) Set(val *PaymentInitiationMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationMetadata(val *PaymentInitiationMetadata) *NullablePaymentInitiationMetadata {
	return &NullablePaymentInitiationMetadata{value: val, isSet: true}
}

func (v NullablePaymentInitiationMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


