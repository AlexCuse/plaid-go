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
)

// PartnerCustomerOAuthInstitutionsGetResponse Response schema for `/partner/customer/oauth_institutions/get`.
type PartnerCustomerOAuthInstitutionsGetResponse struct {
	FlowdownStatus *PartnerEndCustomerFlowdownStatus `json:"flowdown_status,omitempty"`
	QuestionnaireStatus *PartnerEndCustomerQuestionnaireStatus `json:"questionnaire_status,omitempty"`
	// The OAuth institutions with which the end customer's application is being registered.
	Institutions *[]PartnerEndCustomerOAuthInstitution `json:"institutions,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId *string `json:"request_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerCustomerOAuthInstitutionsGetResponse PartnerCustomerOAuthInstitutionsGetResponse

// NewPartnerCustomerOAuthInstitutionsGetResponse instantiates a new PartnerCustomerOAuthInstitutionsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerCustomerOAuthInstitutionsGetResponse() *PartnerCustomerOAuthInstitutionsGetResponse {
	this := PartnerCustomerOAuthInstitutionsGetResponse{}
	return &this
}

// NewPartnerCustomerOAuthInstitutionsGetResponseWithDefaults instantiates a new PartnerCustomerOAuthInstitutionsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerCustomerOAuthInstitutionsGetResponseWithDefaults() *PartnerCustomerOAuthInstitutionsGetResponse {
	this := PartnerCustomerOAuthInstitutionsGetResponse{}
	return &this
}

// GetFlowdownStatus returns the FlowdownStatus field value if set, zero value otherwise.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) GetFlowdownStatus() PartnerEndCustomerFlowdownStatus {
	if o == nil || o.FlowdownStatus == nil {
		var ret PartnerEndCustomerFlowdownStatus
		return ret
	}
	return *o.FlowdownStatus
}

// GetFlowdownStatusOk returns a tuple with the FlowdownStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) GetFlowdownStatusOk() (*PartnerEndCustomerFlowdownStatus, bool) {
	if o == nil || o.FlowdownStatus == nil {
		return nil, false
	}
	return o.FlowdownStatus, true
}

// HasFlowdownStatus returns a boolean if a field has been set.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) HasFlowdownStatus() bool {
	if o != nil && o.FlowdownStatus != nil {
		return true
	}

	return false
}

// SetFlowdownStatus gets a reference to the given PartnerEndCustomerFlowdownStatus and assigns it to the FlowdownStatus field.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) SetFlowdownStatus(v PartnerEndCustomerFlowdownStatus) {
	o.FlowdownStatus = &v
}

// GetQuestionnaireStatus returns the QuestionnaireStatus field value if set, zero value otherwise.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) GetQuestionnaireStatus() PartnerEndCustomerQuestionnaireStatus {
	if o == nil || o.QuestionnaireStatus == nil {
		var ret PartnerEndCustomerQuestionnaireStatus
		return ret
	}
	return *o.QuestionnaireStatus
}

// GetQuestionnaireStatusOk returns a tuple with the QuestionnaireStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) GetQuestionnaireStatusOk() (*PartnerEndCustomerQuestionnaireStatus, bool) {
	if o == nil || o.QuestionnaireStatus == nil {
		return nil, false
	}
	return o.QuestionnaireStatus, true
}

// HasQuestionnaireStatus returns a boolean if a field has been set.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) HasQuestionnaireStatus() bool {
	if o != nil && o.QuestionnaireStatus != nil {
		return true
	}

	return false
}

// SetQuestionnaireStatus gets a reference to the given PartnerEndCustomerQuestionnaireStatus and assigns it to the QuestionnaireStatus field.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) SetQuestionnaireStatus(v PartnerEndCustomerQuestionnaireStatus) {
	o.QuestionnaireStatus = &v
}

// GetInstitutions returns the Institutions field value if set, zero value otherwise.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) GetInstitutions() []PartnerEndCustomerOAuthInstitution {
	if o == nil || o.Institutions == nil {
		var ret []PartnerEndCustomerOAuthInstitution
		return ret
	}
	return *o.Institutions
}

// GetInstitutionsOk returns a tuple with the Institutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) GetInstitutionsOk() (*[]PartnerEndCustomerOAuthInstitution, bool) {
	if o == nil || o.Institutions == nil {
		return nil, false
	}
	return o.Institutions, true
}

// HasInstitutions returns a boolean if a field has been set.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) HasInstitutions() bool {
	if o != nil && o.Institutions != nil {
		return true
	}

	return false
}

// SetInstitutions gets a reference to the given []PartnerEndCustomerOAuthInstitution and assigns it to the Institutions field.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) SetInstitutions(v []PartnerEndCustomerOAuthInstitution) {
	o.Institutions = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *PartnerCustomerOAuthInstitutionsGetResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o PartnerCustomerOAuthInstitutionsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlowdownStatus != nil {
		toSerialize["flowdown_status"] = o.FlowdownStatus
	}
	if o.QuestionnaireStatus != nil {
		toSerialize["questionnaire_status"] = o.QuestionnaireStatus
	}
	if o.Institutions != nil {
		toSerialize["institutions"] = o.Institutions
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerCustomerOAuthInstitutionsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerCustomerOAuthInstitutionsGetResponse := _PartnerCustomerOAuthInstitutionsGetResponse{}

	if err = json.Unmarshal(bytes, &varPartnerCustomerOAuthInstitutionsGetResponse); err == nil {
		*o = PartnerCustomerOAuthInstitutionsGetResponse(varPartnerCustomerOAuthInstitutionsGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "flowdown_status")
		delete(additionalProperties, "questionnaire_status")
		delete(additionalProperties, "institutions")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerCustomerOAuthInstitutionsGetResponse struct {
	value *PartnerCustomerOAuthInstitutionsGetResponse
	isSet bool
}

func (v NullablePartnerCustomerOAuthInstitutionsGetResponse) Get() *PartnerCustomerOAuthInstitutionsGetResponse {
	return v.value
}

func (v *NullablePartnerCustomerOAuthInstitutionsGetResponse) Set(val *PartnerCustomerOAuthInstitutionsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerCustomerOAuthInstitutionsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerCustomerOAuthInstitutionsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerCustomerOAuthInstitutionsGetResponse(val *PartnerCustomerOAuthInstitutionsGetResponse) *NullablePartnerCustomerOAuthInstitutionsGetResponse {
	return &NullablePartnerCustomerOAuthInstitutionsGetResponse{value: val, isSet: true}
}

func (v NullablePartnerCustomerOAuthInstitutionsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerCustomerOAuthInstitutionsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


