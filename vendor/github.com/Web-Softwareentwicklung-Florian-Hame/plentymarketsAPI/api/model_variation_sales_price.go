/*
Plentymarkets-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// VariationSalesPrice variation sales price model
type VariationSalesPrice struct {
	VariationId *int32 `json:"variationId,omitempty"`
	SalesPriceId *int32 `json:"salesPriceId,omitempty"`
	Price *float32 `json:"price,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewVariationSalesPrice instantiates a new VariationSalesPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariationSalesPrice() *VariationSalesPrice {
	this := VariationSalesPrice{}
	return &this
}

// NewVariationSalesPriceWithDefaults instantiates a new VariationSalesPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariationSalesPriceWithDefaults() *VariationSalesPrice {
	this := VariationSalesPrice{}
	return &this
}

// GetVariationId returns the VariationId field value if set, zero value otherwise.
func (o *VariationSalesPrice) GetVariationId() int32 {
	if o == nil || o.VariationId == nil {
		var ret int32
		return ret
	}
	return *o.VariationId
}

// GetVariationIdOk returns a tuple with the VariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationSalesPrice) GetVariationIdOk() (*int32, bool) {
	if o == nil || o.VariationId == nil {
		return nil, false
	}
	return o.VariationId, true
}

// HasVariationId returns a boolean if a field has been set.
func (o *VariationSalesPrice) HasVariationId() bool {
	if o != nil && o.VariationId != nil {
		return true
	}

	return false
}

// SetVariationId gets a reference to the given int32 and assigns it to the VariationId field.
func (o *VariationSalesPrice) SetVariationId(v int32) {
	o.VariationId = &v
}

// GetSalesPriceId returns the SalesPriceId field value if set, zero value otherwise.
func (o *VariationSalesPrice) GetSalesPriceId() int32 {
	if o == nil || o.SalesPriceId == nil {
		var ret int32
		return ret
	}
	return *o.SalesPriceId
}

// GetSalesPriceIdOk returns a tuple with the SalesPriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationSalesPrice) GetSalesPriceIdOk() (*int32, bool) {
	if o == nil || o.SalesPriceId == nil {
		return nil, false
	}
	return o.SalesPriceId, true
}

// HasSalesPriceId returns a boolean if a field has been set.
func (o *VariationSalesPrice) HasSalesPriceId() bool {
	if o != nil && o.SalesPriceId != nil {
		return true
	}

	return false
}

// SetSalesPriceId gets a reference to the given int32 and assigns it to the SalesPriceId field.
func (o *VariationSalesPrice) SetSalesPriceId(v int32) {
	o.SalesPriceId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *VariationSalesPrice) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationSalesPrice) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *VariationSalesPrice) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *VariationSalesPrice) SetPrice(v float32) {
	o.Price = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VariationSalesPrice) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationSalesPrice) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VariationSalesPrice) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VariationSalesPrice) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VariationSalesPrice) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationSalesPrice) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VariationSalesPrice) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VariationSalesPrice) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o VariationSalesPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VariationId != nil {
		toSerialize["variationId"] = o.VariationId
	}
	if o.SalesPriceId != nil {
		toSerialize["salesPriceId"] = o.SalesPriceId
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableVariationSalesPrice struct {
	value *VariationSalesPrice
	isSet bool
}

func (v NullableVariationSalesPrice) Get() *VariationSalesPrice {
	return v.value
}

func (v *NullableVariationSalesPrice) Set(val *VariationSalesPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableVariationSalesPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableVariationSalesPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariationSalesPrice(val *VariationSalesPrice) *NullableVariationSalesPrice {
	return &NullableVariationSalesPrice{value: val, isSet: true}
}

func (v NullableVariationSalesPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariationSalesPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

