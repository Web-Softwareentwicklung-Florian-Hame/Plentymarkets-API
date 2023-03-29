/*
Plentymarkets-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OrderRelation struct for OrderRelation
type OrderRelation struct {
	OrderId *int32 `json:"orderId,omitempty"`
	ReferenceType *string `json:"referenceType,omitempty"`
	ReferenceId *int32 `json:"referenceId,omitempty"`
	Relation *float32 `json:"relation,omitempty"`
}

// NewOrderRelation instantiates a new OrderRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRelation() *OrderRelation {
	this := OrderRelation{}
	return &this
}

// NewOrderRelationWithDefaults instantiates a new OrderRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRelationWithDefaults() *OrderRelation {
	this := OrderRelation{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderRelation) GetOrderId() int32 {
	if o == nil || o.OrderId == nil {
		var ret int32
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRelation) GetOrderIdOk() (*int32, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderRelation) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int32 and assigns it to the OrderId field.
func (o *OrderRelation) SetOrderId(v int32) {
	o.OrderId = &v
}

// GetReferenceType returns the ReferenceType field value if set, zero value otherwise.
func (o *OrderRelation) GetReferenceType() string {
	if o == nil || o.ReferenceType == nil {
		var ret string
		return ret
	}
	return *o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRelation) GetReferenceTypeOk() (*string, bool) {
	if o == nil || o.ReferenceType == nil {
		return nil, false
	}
	return o.ReferenceType, true
}

// HasReferenceType returns a boolean if a field has been set.
func (o *OrderRelation) HasReferenceType() bool {
	if o != nil && o.ReferenceType != nil {
		return true
	}

	return false
}

// SetReferenceType gets a reference to the given string and assigns it to the ReferenceType field.
func (o *OrderRelation) SetReferenceType(v string) {
	o.ReferenceType = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *OrderRelation) GetReferenceId() int32 {
	if o == nil || o.ReferenceId == nil {
		var ret int32
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRelation) GetReferenceIdOk() (*int32, bool) {
	if o == nil || o.ReferenceId == nil {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *OrderRelation) HasReferenceId() bool {
	if o != nil && o.ReferenceId != nil {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given int32 and assigns it to the ReferenceId field.
func (o *OrderRelation) SetReferenceId(v int32) {
	o.ReferenceId = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *OrderRelation) GetRelation() float32 {
	if o == nil || o.Relation == nil {
		var ret float32
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRelation) GetRelationOk() (*float32, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *OrderRelation) HasRelation() bool {
	if o != nil && o.Relation != nil {
		return true
	}

	return false
}

// SetRelation gets a reference to the given float32 and assigns it to the Relation field.
func (o *OrderRelation) SetRelation(v float32) {
	o.Relation = &v
}

func (o OrderRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrderId != nil {
		toSerialize["orderId"] = o.OrderId
	}
	if o.ReferenceType != nil {
		toSerialize["referenceType"] = o.ReferenceType
	}
	if o.ReferenceId != nil {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
	}
	return json.Marshal(toSerialize)
}

type NullableOrderRelation struct {
	value *OrderRelation
	isSet bool
}

func (v NullableOrderRelation) Get() *OrderRelation {
	return v.value
}

func (v *NullableOrderRelation) Set(val *OrderRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRelation(val *OrderRelation) *NullableOrderRelation {
	return &NullableOrderRelation{value: val, isSet: true}
}

func (v NullableOrderRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

