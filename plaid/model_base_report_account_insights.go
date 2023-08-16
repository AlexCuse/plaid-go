/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.410.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// BaseReportAccountInsights Calculated insights derived from transaction-level data.
type BaseReportAccountInsights struct {
	// Date of the earliest transaction in the base report for the account.
	OldestTransactionDate *string `json:"oldest_transaction_date,omitempty"`
	// Date of the most recent transaction in the base report for the account.
	MostRecentTransactionDate *string `json:"most_recent_transaction_date,omitempty"`
	// Number of days days available in the base report for the account.
	DaysAvailable *int32 `json:"days_available,omitempty"`
	// Average number of days between sequential transactions
	AverageDaysBetweenTransactions *float32 `json:"average_days_between_transactions,omitempty"`
	// Longest gap between sequential transactions
	LongestGapBetweenTransactions *[]BaseReportLongestGapInsights `json:"longest_gap_between_transactions,omitempty"`
	// The number of debits into the account. This field will be null for non-depository accounts.
	NumberOfInflows []BaseReportNumberFlowInsights `json:"number_of_inflows,omitempty"`
	// Average amount of debit transactions into account. This field will be null for non-depository accounts. This field only takes into account USD transactions from the account.
	AverageInflowAmount []BaseReportAverageFlowInsights `json:"average_inflow_amount,omitempty"`
	// The number of credit into the account. This field will be null for non-depository accounts.
	NumberOfOutflows []BaseReportNumberFlowInsights `json:"number_of_outflows,omitempty"`
	// Average amount of credit transactions into account. This field will be null for non-depository accounts. This field only takes into account USD transactions from the account.
	AverageOutflowAmount []BaseReportAverageFlowInsights `json:"average_outflow_amount,omitempty"`
	// Number of days with no transactions
	NumberOfDaysNoTransactions *int32 `json:"number_of_days_no_transactions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseReportAccountInsights BaseReportAccountInsights

// NewBaseReportAccountInsights instantiates a new BaseReportAccountInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseReportAccountInsights() *BaseReportAccountInsights {
	this := BaseReportAccountInsights{}
	return &this
}

// NewBaseReportAccountInsightsWithDefaults instantiates a new BaseReportAccountInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseReportAccountInsightsWithDefaults() *BaseReportAccountInsights {
	this := BaseReportAccountInsights{}
	return &this
}

// GetOldestTransactionDate returns the OldestTransactionDate field value if set, zero value otherwise.
func (o *BaseReportAccountInsights) GetOldestTransactionDate() string {
	if o == nil || o.OldestTransactionDate == nil {
		var ret string
		return ret
	}
	return *o.OldestTransactionDate
}

// GetOldestTransactionDateOk returns a tuple with the OldestTransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportAccountInsights) GetOldestTransactionDateOk() (*string, bool) {
	if o == nil || o.OldestTransactionDate == nil {
		return nil, false
	}
	return o.OldestTransactionDate, true
}

// HasOldestTransactionDate returns a boolean if a field has been set.
func (o *BaseReportAccountInsights) HasOldestTransactionDate() bool {
	if o != nil && o.OldestTransactionDate != nil {
		return true
	}

	return false
}

// SetOldestTransactionDate gets a reference to the given string and assigns it to the OldestTransactionDate field.
func (o *BaseReportAccountInsights) SetOldestTransactionDate(v string) {
	o.OldestTransactionDate = &v
}

// GetMostRecentTransactionDate returns the MostRecentTransactionDate field value if set, zero value otherwise.
func (o *BaseReportAccountInsights) GetMostRecentTransactionDate() string {
	if o == nil || o.MostRecentTransactionDate == nil {
		var ret string
		return ret
	}
	return *o.MostRecentTransactionDate
}

// GetMostRecentTransactionDateOk returns a tuple with the MostRecentTransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportAccountInsights) GetMostRecentTransactionDateOk() (*string, bool) {
	if o == nil || o.MostRecentTransactionDate == nil {
		return nil, false
	}
	return o.MostRecentTransactionDate, true
}

// HasMostRecentTransactionDate returns a boolean if a field has been set.
func (o *BaseReportAccountInsights) HasMostRecentTransactionDate() bool {
	if o != nil && o.MostRecentTransactionDate != nil {
		return true
	}

	return false
}

// SetMostRecentTransactionDate gets a reference to the given string and assigns it to the MostRecentTransactionDate field.
func (o *BaseReportAccountInsights) SetMostRecentTransactionDate(v string) {
	o.MostRecentTransactionDate = &v
}

// GetDaysAvailable returns the DaysAvailable field value if set, zero value otherwise.
func (o *BaseReportAccountInsights) GetDaysAvailable() int32 {
	if o == nil || o.DaysAvailable == nil {
		var ret int32
		return ret
	}
	return *o.DaysAvailable
}

// GetDaysAvailableOk returns a tuple with the DaysAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportAccountInsights) GetDaysAvailableOk() (*int32, bool) {
	if o == nil || o.DaysAvailable == nil {
		return nil, false
	}
	return o.DaysAvailable, true
}

// HasDaysAvailable returns a boolean if a field has been set.
func (o *BaseReportAccountInsights) HasDaysAvailable() bool {
	if o != nil && o.DaysAvailable != nil {
		return true
	}

	return false
}

// SetDaysAvailable gets a reference to the given int32 and assigns it to the DaysAvailable field.
func (o *BaseReportAccountInsights) SetDaysAvailable(v int32) {
	o.DaysAvailable = &v
}

// GetAverageDaysBetweenTransactions returns the AverageDaysBetweenTransactions field value if set, zero value otherwise.
func (o *BaseReportAccountInsights) GetAverageDaysBetweenTransactions() float32 {
	if o == nil || o.AverageDaysBetweenTransactions == nil {
		var ret float32
		return ret
	}
	return *o.AverageDaysBetweenTransactions
}

// GetAverageDaysBetweenTransactionsOk returns a tuple with the AverageDaysBetweenTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportAccountInsights) GetAverageDaysBetweenTransactionsOk() (*float32, bool) {
	if o == nil || o.AverageDaysBetweenTransactions == nil {
		return nil, false
	}
	return o.AverageDaysBetweenTransactions, true
}

// HasAverageDaysBetweenTransactions returns a boolean if a field has been set.
func (o *BaseReportAccountInsights) HasAverageDaysBetweenTransactions() bool {
	if o != nil && o.AverageDaysBetweenTransactions != nil {
		return true
	}

	return false
}

// SetAverageDaysBetweenTransactions gets a reference to the given float32 and assigns it to the AverageDaysBetweenTransactions field.
func (o *BaseReportAccountInsights) SetAverageDaysBetweenTransactions(v float32) {
	o.AverageDaysBetweenTransactions = &v
}

// GetLongestGapBetweenTransactions returns the LongestGapBetweenTransactions field value if set, zero value otherwise.
func (o *BaseReportAccountInsights) GetLongestGapBetweenTransactions() []BaseReportLongestGapInsights {
	if o == nil || o.LongestGapBetweenTransactions == nil {
		var ret []BaseReportLongestGapInsights
		return ret
	}
	return *o.LongestGapBetweenTransactions
}

// GetLongestGapBetweenTransactionsOk returns a tuple with the LongestGapBetweenTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportAccountInsights) GetLongestGapBetweenTransactionsOk() (*[]BaseReportLongestGapInsights, bool) {
	if o == nil || o.LongestGapBetweenTransactions == nil {
		return nil, false
	}
	return o.LongestGapBetweenTransactions, true
}

// HasLongestGapBetweenTransactions returns a boolean if a field has been set.
func (o *BaseReportAccountInsights) HasLongestGapBetweenTransactions() bool {
	if o != nil && o.LongestGapBetweenTransactions != nil {
		return true
	}

	return false
}

// SetLongestGapBetweenTransactions gets a reference to the given []BaseReportLongestGapInsights and assigns it to the LongestGapBetweenTransactions field.
func (o *BaseReportAccountInsights) SetLongestGapBetweenTransactions(v []BaseReportLongestGapInsights) {
	o.LongestGapBetweenTransactions = &v
}

// GetNumberOfInflows returns the NumberOfInflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseReportAccountInsights) GetNumberOfInflows() []BaseReportNumberFlowInsights {
	if o == nil  {
		var ret []BaseReportNumberFlowInsights
		return ret
	}
	return o.NumberOfInflows
}

// GetNumberOfInflowsOk returns a tuple with the NumberOfInflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseReportAccountInsights) GetNumberOfInflowsOk() (*[]BaseReportNumberFlowInsights, bool) {
	if o == nil || o.NumberOfInflows == nil {
		return nil, false
	}
	return &o.NumberOfInflows, true
}

// HasNumberOfInflows returns a boolean if a field has been set.
func (o *BaseReportAccountInsights) HasNumberOfInflows() bool {
	if o != nil && o.NumberOfInflows != nil {
		return true
	}

	return false
}

// SetNumberOfInflows gets a reference to the given []BaseReportNumberFlowInsights and assigns it to the NumberOfInflows field.
func (o *BaseReportAccountInsights) SetNumberOfInflows(v []BaseReportNumberFlowInsights) {
	o.NumberOfInflows = v
}

// GetAverageInflowAmount returns the AverageInflowAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseReportAccountInsights) GetAverageInflowAmount() []BaseReportAverageFlowInsights {
	if o == nil  {
		var ret []BaseReportAverageFlowInsights
		return ret
	}
	return o.AverageInflowAmount
}

// GetAverageInflowAmountOk returns a tuple with the AverageInflowAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseReportAccountInsights) GetAverageInflowAmountOk() (*[]BaseReportAverageFlowInsights, bool) {
	if o == nil || o.AverageInflowAmount == nil {
		return nil, false
	}
	return &o.AverageInflowAmount, true
}

// HasAverageInflowAmount returns a boolean if a field has been set.
func (o *BaseReportAccountInsights) HasAverageInflowAmount() bool {
	if o != nil && o.AverageInflowAmount != nil {
		return true
	}

	return false
}

// SetAverageInflowAmount gets a reference to the given []BaseReportAverageFlowInsights and assigns it to the AverageInflowAmount field.
func (o *BaseReportAccountInsights) SetAverageInflowAmount(v []BaseReportAverageFlowInsights) {
	o.AverageInflowAmount = v
}

// GetNumberOfOutflows returns the NumberOfOutflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseReportAccountInsights) GetNumberOfOutflows() []BaseReportNumberFlowInsights {
	if o == nil  {
		var ret []BaseReportNumberFlowInsights
		return ret
	}
	return o.NumberOfOutflows
}

// GetNumberOfOutflowsOk returns a tuple with the NumberOfOutflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseReportAccountInsights) GetNumberOfOutflowsOk() (*[]BaseReportNumberFlowInsights, bool) {
	if o == nil || o.NumberOfOutflows == nil {
		return nil, false
	}
	return &o.NumberOfOutflows, true
}

// HasNumberOfOutflows returns a boolean if a field has been set.
func (o *BaseReportAccountInsights) HasNumberOfOutflows() bool {
	if o != nil && o.NumberOfOutflows != nil {
		return true
	}

	return false
}

// SetNumberOfOutflows gets a reference to the given []BaseReportNumberFlowInsights and assigns it to the NumberOfOutflows field.
func (o *BaseReportAccountInsights) SetNumberOfOutflows(v []BaseReportNumberFlowInsights) {
	o.NumberOfOutflows = v
}

// GetAverageOutflowAmount returns the AverageOutflowAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseReportAccountInsights) GetAverageOutflowAmount() []BaseReportAverageFlowInsights {
	if o == nil  {
		var ret []BaseReportAverageFlowInsights
		return ret
	}
	return o.AverageOutflowAmount
}

// GetAverageOutflowAmountOk returns a tuple with the AverageOutflowAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseReportAccountInsights) GetAverageOutflowAmountOk() (*[]BaseReportAverageFlowInsights, bool) {
	if o == nil || o.AverageOutflowAmount == nil {
		return nil, false
	}
	return &o.AverageOutflowAmount, true
}

// HasAverageOutflowAmount returns a boolean if a field has been set.
func (o *BaseReportAccountInsights) HasAverageOutflowAmount() bool {
	if o != nil && o.AverageOutflowAmount != nil {
		return true
	}

	return false
}

// SetAverageOutflowAmount gets a reference to the given []BaseReportAverageFlowInsights and assigns it to the AverageOutflowAmount field.
func (o *BaseReportAccountInsights) SetAverageOutflowAmount(v []BaseReportAverageFlowInsights) {
	o.AverageOutflowAmount = v
}

// GetNumberOfDaysNoTransactions returns the NumberOfDaysNoTransactions field value if set, zero value otherwise.
func (o *BaseReportAccountInsights) GetNumberOfDaysNoTransactions() int32 {
	if o == nil || o.NumberOfDaysNoTransactions == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfDaysNoTransactions
}

// GetNumberOfDaysNoTransactionsOk returns a tuple with the NumberOfDaysNoTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportAccountInsights) GetNumberOfDaysNoTransactionsOk() (*int32, bool) {
	if o == nil || o.NumberOfDaysNoTransactions == nil {
		return nil, false
	}
	return o.NumberOfDaysNoTransactions, true
}

// HasNumberOfDaysNoTransactions returns a boolean if a field has been set.
func (o *BaseReportAccountInsights) HasNumberOfDaysNoTransactions() bool {
	if o != nil && o.NumberOfDaysNoTransactions != nil {
		return true
	}

	return false
}

// SetNumberOfDaysNoTransactions gets a reference to the given int32 and assigns it to the NumberOfDaysNoTransactions field.
func (o *BaseReportAccountInsights) SetNumberOfDaysNoTransactions(v int32) {
	o.NumberOfDaysNoTransactions = &v
}

func (o BaseReportAccountInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OldestTransactionDate != nil {
		toSerialize["oldest_transaction_date"] = o.OldestTransactionDate
	}
	if o.MostRecentTransactionDate != nil {
		toSerialize["most_recent_transaction_date"] = o.MostRecentTransactionDate
	}
	if o.DaysAvailable != nil {
		toSerialize["days_available"] = o.DaysAvailable
	}
	if o.AverageDaysBetweenTransactions != nil {
		toSerialize["average_days_between_transactions"] = o.AverageDaysBetweenTransactions
	}
	if o.LongestGapBetweenTransactions != nil {
		toSerialize["longest_gap_between_transactions"] = o.LongestGapBetweenTransactions
	}
	if o.NumberOfInflows != nil {
		toSerialize["number_of_inflows"] = o.NumberOfInflows
	}
	if o.AverageInflowAmount != nil {
		toSerialize["average_inflow_amount"] = o.AverageInflowAmount
	}
	if o.NumberOfOutflows != nil {
		toSerialize["number_of_outflows"] = o.NumberOfOutflows
	}
	if o.AverageOutflowAmount != nil {
		toSerialize["average_outflow_amount"] = o.AverageOutflowAmount
	}
	if o.NumberOfDaysNoTransactions != nil {
		toSerialize["number_of_days_no_transactions"] = o.NumberOfDaysNoTransactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseReportAccountInsights) UnmarshalJSON(bytes []byte) (err error) {
	varBaseReportAccountInsights := _BaseReportAccountInsights{}

	if err = json.Unmarshal(bytes, &varBaseReportAccountInsights); err == nil {
		*o = BaseReportAccountInsights(varBaseReportAccountInsights)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "oldest_transaction_date")
		delete(additionalProperties, "most_recent_transaction_date")
		delete(additionalProperties, "days_available")
		delete(additionalProperties, "average_days_between_transactions")
		delete(additionalProperties, "longest_gap_between_transactions")
		delete(additionalProperties, "number_of_inflows")
		delete(additionalProperties, "average_inflow_amount")
		delete(additionalProperties, "number_of_outflows")
		delete(additionalProperties, "average_outflow_amount")
		delete(additionalProperties, "number_of_days_no_transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseReportAccountInsights struct {
	value *BaseReportAccountInsights
	isSet bool
}

func (v NullableBaseReportAccountInsights) Get() *BaseReportAccountInsights {
	return v.value
}

func (v *NullableBaseReportAccountInsights) Set(val *BaseReportAccountInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseReportAccountInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseReportAccountInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseReportAccountInsights(val *BaseReportAccountInsights) *NullableBaseReportAccountInsights {
	return &NullableBaseReportAccountInsights{value: val, isSet: true}
}

func (v NullableBaseReportAccountInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseReportAccountInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


