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

// AssetReportUser The user object allows you to provide additional information about the user to be appended to the Asset Report. All fields are optional. The `first_name`, `last_name`, and `ssn` fields are required if you would like the Report to be eligible for Fannie Mae’s Day 1 Certainty™ program.
type AssetReportUser struct {
	// An identifier you determine and submit for the user.
	ClientUserId NullableString `json:"client_user_id,omitempty"`
	// The user's first name. Required for the Fannie Mae Day 1 Certainty™ program.
	FirstName NullableString `json:"first_name,omitempty"`
	// The user's middle name
	MiddleName NullableString `json:"middle_name,omitempty"`
	// The user's last name.  Required for the Fannie Mae Day 1 Certainty™ program.
	LastName NullableString `json:"last_name,omitempty"`
	// The user's Social Security Number. Required for the Fannie Mae Day 1 Certainty™ program.  Format: \"ddd-dd-dddd\"
	Ssn NullableString `json:"ssn,omitempty"`
	// The user's phone number, in E.164 format: +{countrycode}{number}. For example: \"+14151234567\". Phone numbers provided in other formats will be parsed on a best-effort basis.
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// The user's email address.
	Email NullableString `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportUser AssetReportUser

// NewAssetReportUser instantiates a new AssetReportUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportUser() *AssetReportUser {
	this := AssetReportUser{}
	return &this
}

// NewAssetReportUserWithDefaults instantiates a new AssetReportUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportUserWithDefaults() *AssetReportUser {
	this := AssetReportUser{}
	return &this
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportUser) GetClientUserId() string {
	if o == nil || o.ClientUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId.Get()
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportUser) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientUserId.Get(), o.ClientUserId.IsSet()
}

// HasClientUserId returns a boolean if a field has been set.
func (o *AssetReportUser) HasClientUserId() bool {
	if o != nil && o.ClientUserId.IsSet() {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given NullableString and assigns it to the ClientUserId field.
func (o *AssetReportUser) SetClientUserId(v string) {
	o.ClientUserId.Set(&v)
}
// SetClientUserIdNil sets the value for ClientUserId to be an explicit nil
func (o *AssetReportUser) SetClientUserIdNil() {
	o.ClientUserId.Set(nil)
}

// UnsetClientUserId ensures that no value is present for ClientUserId, not even an explicit nil
func (o *AssetReportUser) UnsetClientUserId() {
	o.ClientUserId.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportUser) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportUser) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *AssetReportUser) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *AssetReportUser) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *AssetReportUser) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *AssetReportUser) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportUser) GetMiddleName() string {
	if o == nil || o.MiddleName.Get() == nil {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportUser) GetMiddleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *AssetReportUser) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *AssetReportUser) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *AssetReportUser) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *AssetReportUser) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportUser) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportUser) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *AssetReportUser) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *AssetReportUser) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *AssetReportUser) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *AssetReportUser) UnsetLastName() {
	o.LastName.Unset()
}

// GetSsn returns the Ssn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportUser) GetSsn() string {
	if o == nil || o.Ssn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ssn.Get()
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportUser) GetSsnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ssn.Get(), o.Ssn.IsSet()
}

// HasSsn returns a boolean if a field has been set.
func (o *AssetReportUser) HasSsn() bool {
	if o != nil && o.Ssn.IsSet() {
		return true
	}

	return false
}

// SetSsn gets a reference to the given NullableString and assigns it to the Ssn field.
func (o *AssetReportUser) SetSsn(v string) {
	o.Ssn.Set(&v)
}
// SetSsnNil sets the value for Ssn to be an explicit nil
func (o *AssetReportUser) SetSsnNil() {
	o.Ssn.Set(nil)
}

// UnsetSsn ensures that no value is present for Ssn, not even an explicit nil
func (o *AssetReportUser) UnsetSsn() {
	o.Ssn.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportUser) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportUser) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AssetReportUser) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *AssetReportUser) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *AssetReportUser) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *AssetReportUser) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetReportUser) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetReportUser) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *AssetReportUser) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *AssetReportUser) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *AssetReportUser) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *AssetReportUser) UnsetEmail() {
	o.Email.Unset()
}

func (o AssetReportUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientUserId.IsSet() {
		toSerialize["client_user_id"] = o.ClientUserId.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.MiddleName.IsSet() {
		toSerialize["middle_name"] = o.MiddleName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.Ssn.IsSet() {
		toSerialize["ssn"] = o.Ssn.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportUser) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportUser := _AssetReportUser{}

	if err = json.Unmarshal(bytes, &varAssetReportUser); err == nil {
		*o = AssetReportUser(varAssetReportUser)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_user_id")
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "middle_name")
		delete(additionalProperties, "last_name")
		delete(additionalProperties, "ssn")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportUser struct {
	value *AssetReportUser
	isSet bool
}

func (v NullableAssetReportUser) Get() *AssetReportUser {
	return v.value
}

func (v *NullableAssetReportUser) Set(val *AssetReportUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportUser(val *AssetReportUser) *NullableAssetReportUser {
	return &NullableAssetReportUser{value: val, isSet: true}
}

func (v NullableAssetReportUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


