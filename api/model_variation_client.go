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

// VariationClient variation client model
type VariationClient struct {
	VariationId *int32 `json:"variationId,omitempty"`
	PlentyId *int32 `json:"plentyId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewVariationClient instantiates a new VariationClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariationClient() *VariationClient {
	this := VariationClient{}
	return &this
}

// NewVariationClientWithDefaults instantiates a new VariationClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariationClientWithDefaults() *VariationClient {
	this := VariationClient{}
	return &this
}

// GetVariationId returns the VariationId field value if set, zero value otherwise.
func (o *VariationClient) GetVariationId() int32 {
	if o == nil || o.VariationId == nil {
		var ret int32
		return ret
	}
	return *o.VariationId
}

// GetVariationIdOk returns a tuple with the VariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationClient) GetVariationIdOk() (*int32, bool) {
	if o == nil || o.VariationId == nil {
		return nil, false
	}
	return o.VariationId, true
}

// HasVariationId returns a boolean if a field has been set.
func (o *VariationClient) HasVariationId() bool {
	if o != nil && o.VariationId != nil {
		return true
	}

	return false
}

// SetVariationId gets a reference to the given int32 and assigns it to the VariationId field.
func (o *VariationClient) SetVariationId(v int32) {
	o.VariationId = &v
}

// GetPlentyId returns the PlentyId field value if set, zero value otherwise.
func (o *VariationClient) GetPlentyId() int32 {
	if o == nil || o.PlentyId == nil {
		var ret int32
		return ret
	}
	return *o.PlentyId
}

// GetPlentyIdOk returns a tuple with the PlentyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationClient) GetPlentyIdOk() (*int32, bool) {
	if o == nil || o.PlentyId == nil {
		return nil, false
	}
	return o.PlentyId, true
}

// HasPlentyId returns a boolean if a field has been set.
func (o *VariationClient) HasPlentyId() bool {
	if o != nil && o.PlentyId != nil {
		return true
	}

	return false
}

// SetPlentyId gets a reference to the given int32 and assigns it to the PlentyId field.
func (o *VariationClient) SetPlentyId(v int32) {
	o.PlentyId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VariationClient) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationClient) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VariationClient) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VariationClient) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o VariationClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VariationId != nil {
		toSerialize["variationId"] = o.VariationId
	}
	if o.PlentyId != nil {
		toSerialize["plentyId"] = o.PlentyId
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableVariationClient struct {
	value *VariationClient
	isSet bool
}

func (v NullableVariationClient) Get() *VariationClient {
	return v.value
}

func (v *NullableVariationClient) Set(val *VariationClient) {
	v.value = val
	v.isSet = true
}

func (v NullableVariationClient) IsSet() bool {
	return v.isSet
}

func (v *NullableVariationClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariationClient(val *VariationClient) *NullableVariationClient {
	return &NullableVariationClient{value: val, isSet: true}
}

func (v NullableVariationClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariationClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

