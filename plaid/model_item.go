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
	"time"
)

// Item Metadata about the Item.
type Item struct {
	// The Plaid Item ID. The `item_id` is always unique; linking the same account at the same institution twice will result in two Items with different `item_id` values. Like all Plaid identifiers, the `item_id` is case-sensitive.
	ItemId string `json:"item_id"`
	// The Plaid Institution ID associated with the Item. Field is `null` for Items created via Same Day Micro-deposits.
	InstitutionId NullableString `json:"institution_id,omitempty"`
	// The URL registered to receive webhooks for the Item.
	Webhook NullableString `json:"webhook"`
	Error NullableError `json:"error"`
	// A list of products available for the Item that have not yet been accessed. The contents of this array will be mutually exclusive with `billed_products`.
	AvailableProducts []Products `json:"available_products"`
	// A list of products that have been billed for the Item. The contents of this array will be mutually exclusive with `available_products`. Note - `billed_products` is populated in all environments but only requests in Production are billed. Also note that products that are billed on a pay-per-call basis rather than a pay-per-Item basis, such as `balance`, will not appear here. 
	BilledProducts []Products `json:"billed_products"`
	// A list of authorized products for the Item. 
	Products *[]Products `json:"products,omitempty"`
	// Beta: A list of products that have gone through consent collection for the Item. Only present for those enabled in the beta. 
	ConsentedProducts *[]Products `json:"consented_products,omitempty"`
	// The RFC 3339 timestamp after which the consent provided by the end user will expire. Upon consent expiration, the item will enter the `ITEM_LOGIN_REQUIRED` error state. To circumvent the `ITEM_LOGIN_REQUIRED` error and maintain continuous consent, the end user can reauthenticate via Link’s update mode in advance of the consent expiration time.  Note - This is only relevant for certain OAuth-based institutions. For all other institutions, this field will be null. 
	ConsentExpirationTime NullableTime `json:"consent_expiration_time"`
	// Indicates whether an Item requires user interaction to be updated, which can be the case for Items with some forms of two-factor authentication.  `background` - Item can be updated in the background  `user_present_required` - Item requires user interaction to be updated
	UpdateType string `json:"update_type"`
	AdditionalProperties map[string]interface{}
}

type _Item Item

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem(itemId string, webhook NullableString, error_ NullableError, availableProducts []Products, billedProducts []Products, consentExpirationTime NullableTime, updateType string) *Item {
	this := Item{}
	this.ItemId = itemId
	this.Webhook = webhook
	this.Error = error_
	this.AvailableProducts = availableProducts
	this.BilledProducts = billedProducts
	this.ConsentExpirationTime = consentExpirationTime
	this.UpdateType = updateType
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{}
	return &this
}

// GetItemId returns the ItemId field value
func (o *Item) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *Item) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *Item) SetItemId(v string) {
	o.ItemId = v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Item) GetInstitutionId() string {
	if o == nil || o.InstitutionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId.Get()
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Item) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstitutionId.Get(), o.InstitutionId.IsSet()
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *Item) HasInstitutionId() bool {
	if o != nil && o.InstitutionId.IsSet() {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given NullableString and assigns it to the InstitutionId field.
func (o *Item) SetInstitutionId(v string) {
	o.InstitutionId.Set(&v)
}
// SetInstitutionIdNil sets the value for InstitutionId to be an explicit nil
func (o *Item) SetInstitutionIdNil() {
	o.InstitutionId.Set(nil)
}

// UnsetInstitutionId ensures that no value is present for InstitutionId, not even an explicit nil
func (o *Item) UnsetInstitutionId() {
	o.InstitutionId.Unset()
}

// GetWebhook returns the Webhook field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Item) GetWebhook() string {
	if o == nil || o.Webhook.Get() == nil {
		var ret string
		return ret
	}

	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Item) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// SetWebhook sets field value
func (o *Item) SetWebhook(v string) {
	o.Webhook.Set(&v)
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for Error will be returned
func (o *Item) GetError() Error {
	if o == nil || o.Error.Get() == nil {
		var ret Error
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Item) GetErrorOk() (*Error, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *Item) SetError(v Error) {
	o.Error.Set(&v)
}

// GetAvailableProducts returns the AvailableProducts field value
func (o *Item) GetAvailableProducts() []Products {
	if o == nil {
		var ret []Products
		return ret
	}

	return o.AvailableProducts
}

// GetAvailableProductsOk returns a tuple with the AvailableProducts field value
// and a boolean to check if the value has been set.
func (o *Item) GetAvailableProductsOk() (*[]Products, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailableProducts, true
}

// SetAvailableProducts sets field value
func (o *Item) SetAvailableProducts(v []Products) {
	o.AvailableProducts = v
}

// GetBilledProducts returns the BilledProducts field value
func (o *Item) GetBilledProducts() []Products {
	if o == nil {
		var ret []Products
		return ret
	}

	return o.BilledProducts
}

// GetBilledProductsOk returns a tuple with the BilledProducts field value
// and a boolean to check if the value has been set.
func (o *Item) GetBilledProductsOk() (*[]Products, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BilledProducts, true
}

// SetBilledProducts sets field value
func (o *Item) SetBilledProducts(v []Products) {
	o.BilledProducts = v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *Item) GetProducts() []Products {
	if o == nil || o.Products == nil {
		var ret []Products
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetProductsOk() (*[]Products, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *Item) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []Products and assigns it to the Products field.
func (o *Item) SetProducts(v []Products) {
	o.Products = &v
}

// GetConsentedProducts returns the ConsentedProducts field value if set, zero value otherwise.
func (o *Item) GetConsentedProducts() []Products {
	if o == nil || o.ConsentedProducts == nil {
		var ret []Products
		return ret
	}
	return *o.ConsentedProducts
}

// GetConsentedProductsOk returns a tuple with the ConsentedProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetConsentedProductsOk() (*[]Products, bool) {
	if o == nil || o.ConsentedProducts == nil {
		return nil, false
	}
	return o.ConsentedProducts, true
}

// HasConsentedProducts returns a boolean if a field has been set.
func (o *Item) HasConsentedProducts() bool {
	if o != nil && o.ConsentedProducts != nil {
		return true
	}

	return false
}

// SetConsentedProducts gets a reference to the given []Products and assigns it to the ConsentedProducts field.
func (o *Item) SetConsentedProducts(v []Products) {
	o.ConsentedProducts = &v
}

// GetConsentExpirationTime returns the ConsentExpirationTime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Item) GetConsentExpirationTime() time.Time {
	if o == nil || o.ConsentExpirationTime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ConsentExpirationTime.Get()
}

// GetConsentExpirationTimeOk returns a tuple with the ConsentExpirationTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Item) GetConsentExpirationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConsentExpirationTime.Get(), o.ConsentExpirationTime.IsSet()
}

// SetConsentExpirationTime sets field value
func (o *Item) SetConsentExpirationTime(v time.Time) {
	o.ConsentExpirationTime.Set(&v)
}

// GetUpdateType returns the UpdateType field value
func (o *Item) GetUpdateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateType
}

// GetUpdateTypeOk returns a tuple with the UpdateType field value
// and a boolean to check if the value has been set.
func (o *Item) GetUpdateTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateType, true
}

// SetUpdateType sets field value
func (o *Item) SetUpdateType(v string) {
	o.UpdateType = v
}

func (o Item) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if o.InstitutionId.IsSet() {
		toSerialize["institution_id"] = o.InstitutionId.Get()
	}
	if true {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	if true {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["available_products"] = o.AvailableProducts
	}
	if true {
		toSerialize["billed_products"] = o.BilledProducts
	}
	if o.Products != nil {
		toSerialize["products"] = o.Products
	}
	if o.ConsentedProducts != nil {
		toSerialize["consented_products"] = o.ConsentedProducts
	}
	if true {
		toSerialize["consent_expiration_time"] = o.ConsentExpirationTime.Get()
	}
	if true {
		toSerialize["update_type"] = o.UpdateType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Item) UnmarshalJSON(bytes []byte) (err error) {
	varItem := _Item{}

	if err = json.Unmarshal(bytes, &varItem); err == nil {
		*o = Item(varItem)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "institution_id")
		delete(additionalProperties, "webhook")
		delete(additionalProperties, "error")
		delete(additionalProperties, "available_products")
		delete(additionalProperties, "billed_products")
		delete(additionalProperties, "products")
		delete(additionalProperties, "consented_products")
		delete(additionalProperties, "consent_expiration_time")
		delete(additionalProperties, "update_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItem struct {
	value *Item
	isSet bool
}

func (v NullableItem) Get() *Item {
	return v.value
}

func (v *NullableItem) Set(val *Item) {
	v.value = val
	v.isSet = true
}

func (v NullableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem(val *Item) *NullableItem {
	return &NullableItem{value: val, isSet: true}
}

func (v NullableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


