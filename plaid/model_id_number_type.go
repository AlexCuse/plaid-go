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

// IDNumberType A globally unique and human readable ID type, specific to the country and document category. For more context on this field, see [Hybrid Input Validation](https://plaid.com/docs/identity-verification/hybrid-input-validation).
type IDNumberType string

var _ = fmt.Printf

// List of IDNumberType
const (
	IDNUMBERTYPE_AR_DNI IDNumberType = "ar_dni"
	IDNUMBERTYPE_AU_DRIVERS_LICENSE IDNumberType = "au_drivers_license"
	IDNUMBERTYPE_AU_PASSPORT IDNumberType = "au_passport"
	IDNUMBERTYPE_BR_CPF IDNumberType = "br_cpf"
	IDNUMBERTYPE_CA_SIN IDNumberType = "ca_sin"
	IDNUMBERTYPE_CL_RUN IDNumberType = "cl_run"
	IDNUMBERTYPE_CN_RESIDENT_CARD IDNumberType = "cn_resident_card"
	IDNUMBERTYPE_CO_NIT IDNumberType = "co_nit"
	IDNUMBERTYPE_DK_CPR IDNumberType = "dk_cpr"
	IDNUMBERTYPE_EG_NATIONAL_ID IDNumberType = "eg_national_id"
	IDNUMBERTYPE_ES_DNI IDNumberType = "es_dni"
	IDNUMBERTYPE_ES_NIE IDNumberType = "es_nie"
	IDNUMBERTYPE_HK_HKID IDNumberType = "hk_hkid"
	IDNUMBERTYPE_IN_PAN IDNumberType = "in_pan"
	IDNUMBERTYPE_IT_CF IDNumberType = "it_cf"
	IDNUMBERTYPE_JO_CIVIL_ID IDNumberType = "jo_civil_id"
	IDNUMBERTYPE_JP_MY_NUMBER IDNumberType = "jp_my_number"
	IDNUMBERTYPE_KE_HUDUMA_NAMBA IDNumberType = "ke_huduma_namba"
	IDNUMBERTYPE_KW_CIVIL_ID IDNumberType = "kw_civil_id"
	IDNUMBERTYPE_MX_CURP IDNumberType = "mx_curp"
	IDNUMBERTYPE_MX_RFC IDNumberType = "mx_rfc"
	IDNUMBERTYPE_MY_NRIC IDNumberType = "my_nric"
	IDNUMBERTYPE_NG_NIN IDNumberType = "ng_nin"
	IDNUMBERTYPE_NZ_DRIVERS_LICENSE IDNumberType = "nz_drivers_license"
	IDNUMBERTYPE_OM_CIVIL_ID IDNumberType = "om_civil_id"
	IDNUMBERTYPE_PH_PSN IDNumberType = "ph_psn"
	IDNUMBERTYPE_PL_PESEL IDNumberType = "pl_pesel"
	IDNUMBERTYPE_RO_CNP IDNumberType = "ro_cnp"
	IDNUMBERTYPE_SA_NATIONAL_ID IDNumberType = "sa_national_id"
	IDNUMBERTYPE_SE_PIN IDNumberType = "se_pin"
	IDNUMBERTYPE_SG_NRIC IDNumberType = "sg_nric"
	IDNUMBERTYPE_TR_TC_KIMLIK IDNumberType = "tr_tc_kimlik"
	IDNUMBERTYPE_US_SSN IDNumberType = "us_ssn"
	IDNUMBERTYPE_US_SSN_LAST_4 IDNumberType = "us_ssn_last_4"
	IDNUMBERTYPE_ZA_SMART_ID IDNumberType = "za_smart_id"
)

var allowedIDNumberTypeEnumValues = []IDNumberType{
	"ar_dni",
	"au_drivers_license",
	"au_passport",
	"br_cpf",
	"ca_sin",
	"cl_run",
	"cn_resident_card",
	"co_nit",
	"dk_cpr",
	"eg_national_id",
	"es_dni",
	"es_nie",
	"hk_hkid",
	"in_pan",
	"it_cf",
	"jo_civil_id",
	"jp_my_number",
	"ke_huduma_namba",
	"kw_civil_id",
	"mx_curp",
	"mx_rfc",
	"my_nric",
	"ng_nin",
	"nz_drivers_license",
	"om_civil_id",
	"ph_psn",
	"pl_pesel",
	"ro_cnp",
	"sa_national_id",
	"se_pin",
	"sg_nric",
	"tr_tc_kimlik",
	"us_ssn",
	"us_ssn_last_4",
	"za_smart_id",
}

func (v *IDNumberType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := IDNumberType(value)


	*v = enumTypeValue
	return nil
}

// NewIDNumberTypeFromValue returns a pointer to a valid IDNumberType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIDNumberTypeFromValue(v string) (*IDNumberType, error) {
	ev := IDNumberType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IDNumberType) IsValid() bool {
	for _, existing := range allowedIDNumberTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IDNumberType value
func (v IDNumberType) Ptr() *IDNumberType {
	return &v
}

type NullableIDNumberType struct {
	value *IDNumberType
	isSet bool
}

func (v NullableIDNumberType) Get() *IDNumberType {
	return v.value
}

func (v *NullableIDNumberType) Set(val *IDNumberType) {
	v.value = val
	v.isSet = true
}

func (v NullableIDNumberType) IsSet() bool {
	return v.isSet
}

func (v *NullableIDNumberType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIDNumberType(val *IDNumberType) *NullableIDNumberType {
	return &NullableIDNumberType{value: val, isSet: true}
}

func (v NullableIDNumberType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIDNumberType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

