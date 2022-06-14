/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// RecaptchaRequiredError The request was flagged by Plaid's fraud system, and requires additional verification to ensure they are not a bot.
type RecaptchaRequiredError struct {
	// RECAPTCHA_ERROR
	ErrorType string `json:"error_type"`
	// RECAPTCHA_REQUIRED
	ErrorCode string `json:"error_code"`
	DisplayMessage string `json:"display_message"`
	// 400
	HttpCode string `json:"http_code"`
	// Your user will be prompted to solve a Google reCAPTCHA challenge in the Link Recaptcha pane. If they solve the challenge successfully, the user's request is resubmitted and they are directed to the next Item creation step.
	LinkUserExperience string `json:"link_user_experience"`
	// Plaid's fraud system detects abusive traffic and considers a variety of parameters throughout Item creation requests. When a request is considered risky or possibly fraudulent, Link presents a reCAPTCHA for the user to solve.
	CommonCauses string `json:"common_causes"`
	// Link will automatically guide your user through reCAPTCHA verification. As a general rule, we recommend instrumenting basic fraud monitoring to detect and protect your website from spam and abuse.  If your user cannot verify their session, please submit a Support ticket with the following identifiers: `link_session_id` or `request_id`
	TroubleshootingSteps string `json:"troubleshooting_steps"`
	AdditionalProperties map[string]interface{}
}

type _RecaptchaRequiredError RecaptchaRequiredError

// NewRecaptchaRequiredError instantiates a new RecaptchaRequiredError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecaptchaRequiredError(errorType string, errorCode string, displayMessage string, httpCode string, linkUserExperience string, commonCauses string, troubleshootingSteps string) *RecaptchaRequiredError {
	this := RecaptchaRequiredError{}
	this.ErrorType = errorType
	this.ErrorCode = errorCode
	this.DisplayMessage = displayMessage
	this.HttpCode = httpCode
	this.LinkUserExperience = linkUserExperience
	this.CommonCauses = commonCauses
	this.TroubleshootingSteps = troubleshootingSteps
	return &this
}

// NewRecaptchaRequiredErrorWithDefaults instantiates a new RecaptchaRequiredError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecaptchaRequiredErrorWithDefaults() *RecaptchaRequiredError {
	this := RecaptchaRequiredError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *RecaptchaRequiredError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *RecaptchaRequiredError) GetErrorTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *RecaptchaRequiredError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorCode returns the ErrorCode field value
func (o *RecaptchaRequiredError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *RecaptchaRequiredError) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *RecaptchaRequiredError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetDisplayMessage returns the DisplayMessage field value
func (o *RecaptchaRequiredError) GetDisplayMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayMessage
}

// GetDisplayMessageOk returns a tuple with the DisplayMessage field value
// and a boolean to check if the value has been set.
func (o *RecaptchaRequiredError) GetDisplayMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayMessage, true
}

// SetDisplayMessage sets field value
func (o *RecaptchaRequiredError) SetDisplayMessage(v string) {
	o.DisplayMessage = v
}

// GetHttpCode returns the HttpCode field value
func (o *RecaptchaRequiredError) GetHttpCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpCode
}

// GetHttpCodeOk returns a tuple with the HttpCode field value
// and a boolean to check if the value has been set.
func (o *RecaptchaRequiredError) GetHttpCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HttpCode, true
}

// SetHttpCode sets field value
func (o *RecaptchaRequiredError) SetHttpCode(v string) {
	o.HttpCode = v
}

// GetLinkUserExperience returns the LinkUserExperience field value
func (o *RecaptchaRequiredError) GetLinkUserExperience() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkUserExperience
}

// GetLinkUserExperienceOk returns a tuple with the LinkUserExperience field value
// and a boolean to check if the value has been set.
func (o *RecaptchaRequiredError) GetLinkUserExperienceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkUserExperience, true
}

// SetLinkUserExperience sets field value
func (o *RecaptchaRequiredError) SetLinkUserExperience(v string) {
	o.LinkUserExperience = v
}

// GetCommonCauses returns the CommonCauses field value
func (o *RecaptchaRequiredError) GetCommonCauses() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommonCauses
}

// GetCommonCausesOk returns a tuple with the CommonCauses field value
// and a boolean to check if the value has been set.
func (o *RecaptchaRequiredError) GetCommonCausesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommonCauses, true
}

// SetCommonCauses sets field value
func (o *RecaptchaRequiredError) SetCommonCauses(v string) {
	o.CommonCauses = v
}

// GetTroubleshootingSteps returns the TroubleshootingSteps field value
func (o *RecaptchaRequiredError) GetTroubleshootingSteps() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TroubleshootingSteps
}

// GetTroubleshootingStepsOk returns a tuple with the TroubleshootingSteps field value
// and a boolean to check if the value has been set.
func (o *RecaptchaRequiredError) GetTroubleshootingStepsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TroubleshootingSteps, true
}

// SetTroubleshootingSteps sets field value
func (o *RecaptchaRequiredError) SetTroubleshootingSteps(v string) {
	o.TroubleshootingSteps = v
}

func (o RecaptchaRequiredError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error_type"] = o.ErrorType
	}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if true {
		toSerialize["display_message"] = o.DisplayMessage
	}
	if true {
		toSerialize["http_code"] = o.HttpCode
	}
	if true {
		toSerialize["link_user_experience"] = o.LinkUserExperience
	}
	if true {
		toSerialize["common_causes"] = o.CommonCauses
	}
	if true {
		toSerialize["troubleshooting_steps"] = o.TroubleshootingSteps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecaptchaRequiredError) UnmarshalJSON(bytes []byte) (err error) {
	varRecaptchaRequiredError := _RecaptchaRequiredError{}

	if err = json.Unmarshal(bytes, &varRecaptchaRequiredError); err == nil {
		*o = RecaptchaRequiredError(varRecaptchaRequiredError)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error_type")
		delete(additionalProperties, "error_code")
		delete(additionalProperties, "display_message")
		delete(additionalProperties, "http_code")
		delete(additionalProperties, "link_user_experience")
		delete(additionalProperties, "common_causes")
		delete(additionalProperties, "troubleshooting_steps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecaptchaRequiredError struct {
	value *RecaptchaRequiredError
	isSet bool
}

func (v NullableRecaptchaRequiredError) Get() *RecaptchaRequiredError {
	return v.value
}

func (v *NullableRecaptchaRequiredError) Set(val *RecaptchaRequiredError) {
	v.value = val
	v.isSet = true
}

func (v NullableRecaptchaRequiredError) IsSet() bool {
	return v.isSet
}

func (v *NullableRecaptchaRequiredError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecaptchaRequiredError(val *RecaptchaRequiredError) *NullableRecaptchaRequiredError {
	return &NullableRecaptchaRequiredError{value: val, isSet: true}
}

func (v NullableRecaptchaRequiredError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecaptchaRequiredError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


