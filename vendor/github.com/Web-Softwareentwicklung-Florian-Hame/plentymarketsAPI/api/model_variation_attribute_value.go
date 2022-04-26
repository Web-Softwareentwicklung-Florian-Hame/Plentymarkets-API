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

// VariationAttributeValue variation attribute value model
type VariationAttributeValue struct {
	AttributeValueSetId *int32 `json:"attributeValueSetId,omitempty"`
	AttributeId *int32 `json:"attributeId,omitempty"`
	ValueId *int32 `json:"valueId,omitempty"`
	IsLinkableToImage *bool `json:"isLinkableToImage,omitempty"`
	Attribute *Attribute `json:"attribute,omitempty"`
	AttributeValue *AttributeValue `json:"attributeValue,omitempty"`
}

// NewVariationAttributeValue instantiates a new VariationAttributeValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariationAttributeValue() *VariationAttributeValue {
	this := VariationAttributeValue{}
	return &this
}

// NewVariationAttributeValueWithDefaults instantiates a new VariationAttributeValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariationAttributeValueWithDefaults() *VariationAttributeValue {
	this := VariationAttributeValue{}
	return &this
}

// GetAttributeValueSetId returns the AttributeValueSetId field value if set, zero value otherwise.
func (o *VariationAttributeValue) GetAttributeValueSetId() int32 {
	if o == nil || o.AttributeValueSetId == nil {
		var ret int32
		return ret
	}
	return *o.AttributeValueSetId
}

// GetAttributeValueSetIdOk returns a tuple with the AttributeValueSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationAttributeValue) GetAttributeValueSetIdOk() (*int32, bool) {
	if o == nil || o.AttributeValueSetId == nil {
		return nil, false
	}
	return o.AttributeValueSetId, true
}

// HasAttributeValueSetId returns a boolean if a field has been set.
func (o *VariationAttributeValue) HasAttributeValueSetId() bool {
	if o != nil && o.AttributeValueSetId != nil {
		return true
	}

	return false
}

// SetAttributeValueSetId gets a reference to the given int32 and assigns it to the AttributeValueSetId field.
func (o *VariationAttributeValue) SetAttributeValueSetId(v int32) {
	o.AttributeValueSetId = &v
}

// GetAttributeId returns the AttributeId field value if set, zero value otherwise.
func (o *VariationAttributeValue) GetAttributeId() int32 {
	if o == nil || o.AttributeId == nil {
		var ret int32
		return ret
	}
	return *o.AttributeId
}

// GetAttributeIdOk returns a tuple with the AttributeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationAttributeValue) GetAttributeIdOk() (*int32, bool) {
	if o == nil || o.AttributeId == nil {
		return nil, false
	}
	return o.AttributeId, true
}

// HasAttributeId returns a boolean if a field has been set.
func (o *VariationAttributeValue) HasAttributeId() bool {
	if o != nil && o.AttributeId != nil {
		return true
	}

	return false
}

// SetAttributeId gets a reference to the given int32 and assigns it to the AttributeId field.
func (o *VariationAttributeValue) SetAttributeId(v int32) {
	o.AttributeId = &v
}

// GetValueId returns the ValueId field value if set, zero value otherwise.
func (o *VariationAttributeValue) GetValueId() int32 {
	if o == nil || o.ValueId == nil {
		var ret int32
		return ret
	}
	return *o.ValueId
}

// GetValueIdOk returns a tuple with the ValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationAttributeValue) GetValueIdOk() (*int32, bool) {
	if o == nil || o.ValueId == nil {
		return nil, false
	}
	return o.ValueId, true
}

// HasValueId returns a boolean if a field has been set.
func (o *VariationAttributeValue) HasValueId() bool {
	if o != nil && o.ValueId != nil {
		return true
	}

	return false
}

// SetValueId gets a reference to the given int32 and assigns it to the ValueId field.
func (o *VariationAttributeValue) SetValueId(v int32) {
	o.ValueId = &v
}

// GetIsLinkableToImage returns the IsLinkableToImage field value if set, zero value otherwise.
func (o *VariationAttributeValue) GetIsLinkableToImage() bool {
	if o == nil || o.IsLinkableToImage == nil {
		var ret bool
		return ret
	}
	return *o.IsLinkableToImage
}

// GetIsLinkableToImageOk returns a tuple with the IsLinkableToImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationAttributeValue) GetIsLinkableToImageOk() (*bool, bool) {
	if o == nil || o.IsLinkableToImage == nil {
		return nil, false
	}
	return o.IsLinkableToImage, true
}

// HasIsLinkableToImage returns a boolean if a field has been set.
func (o *VariationAttributeValue) HasIsLinkableToImage() bool {
	if o != nil && o.IsLinkableToImage != nil {
		return true
	}

	return false
}

// SetIsLinkableToImage gets a reference to the given bool and assigns it to the IsLinkableToImage field.
func (o *VariationAttributeValue) SetIsLinkableToImage(v bool) {
	o.IsLinkableToImage = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *VariationAttributeValue) GetAttribute() Attribute {
	if o == nil || o.Attribute == nil {
		var ret Attribute
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationAttributeValue) GetAttributeOk() (*Attribute, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *VariationAttributeValue) HasAttribute() bool {
	if o != nil && o.Attribute != nil {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given Attribute and assigns it to the Attribute field.
func (o *VariationAttributeValue) SetAttribute(v Attribute) {
	o.Attribute = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *VariationAttributeValue) GetAttributeValue() AttributeValue {
	if o == nil || o.AttributeValue == nil {
		var ret AttributeValue
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationAttributeValue) GetAttributeValueOk() (*AttributeValue, bool) {
	if o == nil || o.AttributeValue == nil {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *VariationAttributeValue) HasAttributeValue() bool {
	if o != nil && o.AttributeValue != nil {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given AttributeValue and assigns it to the AttributeValue field.
func (o *VariationAttributeValue) SetAttributeValue(v AttributeValue) {
	o.AttributeValue = &v
}

func (o VariationAttributeValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttributeValueSetId != nil {
		toSerialize["attributeValueSetId"] = o.AttributeValueSetId
	}
	if o.AttributeId != nil {
		toSerialize["attributeId"] = o.AttributeId
	}
	if o.ValueId != nil {
		toSerialize["valueId"] = o.ValueId
	}
	if o.IsLinkableToImage != nil {
		toSerialize["isLinkableToImage"] = o.IsLinkableToImage
	}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.AttributeValue != nil {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	return json.Marshal(toSerialize)
}

type NullableVariationAttributeValue struct {
	value *VariationAttributeValue
	isSet bool
}

func (v NullableVariationAttributeValue) Get() *VariationAttributeValue {
	return v.value
}

func (v *NullableVariationAttributeValue) Set(val *VariationAttributeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableVariationAttributeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableVariationAttributeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariationAttributeValue(val *VariationAttributeValue) *NullableVariationAttributeValue {
	return &NullableVariationAttributeValue{value: val, isSet: true}
}

func (v NullableVariationAttributeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariationAttributeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

