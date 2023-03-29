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

// WarehouseContact struct for WarehouseContact
type WarehouseContact struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Note *string `json:"note,omitempty"`
	TypeId *int32 `json:"typeId,omitempty"`
	OnStockAvailability *int32 `json:"onStockAvailability,omitempty"`
	OutOfStockAvailability *int32 `json:"outOfStockAvailability,omitempty"`
	SplitByShippingProfile *bool `json:"splitByShippingProfile,omitempty"`
	StorageLocationType *string `json:"storageLocationType,omitempty"`
	StorageLocationZone *int32 `json:"storageLocationZone,omitempty"`
	RepairWarehouseId *int32 `json:"repairWarehouseId,omitempty"`
	IsInventoryModeActive *bool `json:"isInventoryModeActive,omitempty"`
	LogisticsType *string `json:"logisticsType,omitempty"`
	AveragePriceSource *string `json:"average_price_source,omitempty"`
	ReorderLevelDynamic *string `json:"reorder_level_dynamic,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	AllocationReferrerIds *[]string `json:"allocationReferrerIds,omitempty"`
	Address *WarehouseContactAddress `json:"address,omitempty"`
}

// NewWarehouseContact instantiates a new WarehouseContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseContact() *WarehouseContact {
	this := WarehouseContact{}
	return &this
}

// NewWarehouseContactWithDefaults instantiates a new WarehouseContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseContactWithDefaults() *WarehouseContact {
	this := WarehouseContact{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WarehouseContact) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WarehouseContact) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WarehouseContact) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WarehouseContact) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WarehouseContact) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WarehouseContact) SetName(v string) {
	o.Name = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *WarehouseContact) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *WarehouseContact) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *WarehouseContact) SetNote(v string) {
	o.Note = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *WarehouseContact) GetTypeId() int32 {
	if o == nil || o.TypeId == nil {
		var ret int32
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetTypeIdOk() (*int32, bool) {
	if o == nil || o.TypeId == nil {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *WarehouseContact) HasTypeId() bool {
	if o != nil && o.TypeId != nil {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given int32 and assigns it to the TypeId field.
func (o *WarehouseContact) SetTypeId(v int32) {
	o.TypeId = &v
}

// GetOnStockAvailability returns the OnStockAvailability field value if set, zero value otherwise.
func (o *WarehouseContact) GetOnStockAvailability() int32 {
	if o == nil || o.OnStockAvailability == nil {
		var ret int32
		return ret
	}
	return *o.OnStockAvailability
}

// GetOnStockAvailabilityOk returns a tuple with the OnStockAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetOnStockAvailabilityOk() (*int32, bool) {
	if o == nil || o.OnStockAvailability == nil {
		return nil, false
	}
	return o.OnStockAvailability, true
}

// HasOnStockAvailability returns a boolean if a field has been set.
func (o *WarehouseContact) HasOnStockAvailability() bool {
	if o != nil && o.OnStockAvailability != nil {
		return true
	}

	return false
}

// SetOnStockAvailability gets a reference to the given int32 and assigns it to the OnStockAvailability field.
func (o *WarehouseContact) SetOnStockAvailability(v int32) {
	o.OnStockAvailability = &v
}

// GetOutOfStockAvailability returns the OutOfStockAvailability field value if set, zero value otherwise.
func (o *WarehouseContact) GetOutOfStockAvailability() int32 {
	if o == nil || o.OutOfStockAvailability == nil {
		var ret int32
		return ret
	}
	return *o.OutOfStockAvailability
}

// GetOutOfStockAvailabilityOk returns a tuple with the OutOfStockAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetOutOfStockAvailabilityOk() (*int32, bool) {
	if o == nil || o.OutOfStockAvailability == nil {
		return nil, false
	}
	return o.OutOfStockAvailability, true
}

// HasOutOfStockAvailability returns a boolean if a field has been set.
func (o *WarehouseContact) HasOutOfStockAvailability() bool {
	if o != nil && o.OutOfStockAvailability != nil {
		return true
	}

	return false
}

// SetOutOfStockAvailability gets a reference to the given int32 and assigns it to the OutOfStockAvailability field.
func (o *WarehouseContact) SetOutOfStockAvailability(v int32) {
	o.OutOfStockAvailability = &v
}

// GetSplitByShippingProfile returns the SplitByShippingProfile field value if set, zero value otherwise.
func (o *WarehouseContact) GetSplitByShippingProfile() bool {
	if o == nil || o.SplitByShippingProfile == nil {
		var ret bool
		return ret
	}
	return *o.SplitByShippingProfile
}

// GetSplitByShippingProfileOk returns a tuple with the SplitByShippingProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetSplitByShippingProfileOk() (*bool, bool) {
	if o == nil || o.SplitByShippingProfile == nil {
		return nil, false
	}
	return o.SplitByShippingProfile, true
}

// HasSplitByShippingProfile returns a boolean if a field has been set.
func (o *WarehouseContact) HasSplitByShippingProfile() bool {
	if o != nil && o.SplitByShippingProfile != nil {
		return true
	}

	return false
}

// SetSplitByShippingProfile gets a reference to the given bool and assigns it to the SplitByShippingProfile field.
func (o *WarehouseContact) SetSplitByShippingProfile(v bool) {
	o.SplitByShippingProfile = &v
}

// GetStorageLocationType returns the StorageLocationType field value if set, zero value otherwise.
func (o *WarehouseContact) GetStorageLocationType() string {
	if o == nil || o.StorageLocationType == nil {
		var ret string
		return ret
	}
	return *o.StorageLocationType
}

// GetStorageLocationTypeOk returns a tuple with the StorageLocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetStorageLocationTypeOk() (*string, bool) {
	if o == nil || o.StorageLocationType == nil {
		return nil, false
	}
	return o.StorageLocationType, true
}

// HasStorageLocationType returns a boolean if a field has been set.
func (o *WarehouseContact) HasStorageLocationType() bool {
	if o != nil && o.StorageLocationType != nil {
		return true
	}

	return false
}

// SetStorageLocationType gets a reference to the given string and assigns it to the StorageLocationType field.
func (o *WarehouseContact) SetStorageLocationType(v string) {
	o.StorageLocationType = &v
}

// GetStorageLocationZone returns the StorageLocationZone field value if set, zero value otherwise.
func (o *WarehouseContact) GetStorageLocationZone() int32 {
	if o == nil || o.StorageLocationZone == nil {
		var ret int32
		return ret
	}
	return *o.StorageLocationZone
}

// GetStorageLocationZoneOk returns a tuple with the StorageLocationZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetStorageLocationZoneOk() (*int32, bool) {
	if o == nil || o.StorageLocationZone == nil {
		return nil, false
	}
	return o.StorageLocationZone, true
}

// HasStorageLocationZone returns a boolean if a field has been set.
func (o *WarehouseContact) HasStorageLocationZone() bool {
	if o != nil && o.StorageLocationZone != nil {
		return true
	}

	return false
}

// SetStorageLocationZone gets a reference to the given int32 and assigns it to the StorageLocationZone field.
func (o *WarehouseContact) SetStorageLocationZone(v int32) {
	o.StorageLocationZone = &v
}

// GetRepairWarehouseId returns the RepairWarehouseId field value if set, zero value otherwise.
func (o *WarehouseContact) GetRepairWarehouseId() int32 {
	if o == nil || o.RepairWarehouseId == nil {
		var ret int32
		return ret
	}
	return *o.RepairWarehouseId
}

// GetRepairWarehouseIdOk returns a tuple with the RepairWarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetRepairWarehouseIdOk() (*int32, bool) {
	if o == nil || o.RepairWarehouseId == nil {
		return nil, false
	}
	return o.RepairWarehouseId, true
}

// HasRepairWarehouseId returns a boolean if a field has been set.
func (o *WarehouseContact) HasRepairWarehouseId() bool {
	if o != nil && o.RepairWarehouseId != nil {
		return true
	}

	return false
}

// SetRepairWarehouseId gets a reference to the given int32 and assigns it to the RepairWarehouseId field.
func (o *WarehouseContact) SetRepairWarehouseId(v int32) {
	o.RepairWarehouseId = &v
}

// GetIsInventoryModeActive returns the IsInventoryModeActive field value if set, zero value otherwise.
func (o *WarehouseContact) GetIsInventoryModeActive() bool {
	if o == nil || o.IsInventoryModeActive == nil {
		var ret bool
		return ret
	}
	return *o.IsInventoryModeActive
}

// GetIsInventoryModeActiveOk returns a tuple with the IsInventoryModeActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetIsInventoryModeActiveOk() (*bool, bool) {
	if o == nil || o.IsInventoryModeActive == nil {
		return nil, false
	}
	return o.IsInventoryModeActive, true
}

// HasIsInventoryModeActive returns a boolean if a field has been set.
func (o *WarehouseContact) HasIsInventoryModeActive() bool {
	if o != nil && o.IsInventoryModeActive != nil {
		return true
	}

	return false
}

// SetIsInventoryModeActive gets a reference to the given bool and assigns it to the IsInventoryModeActive field.
func (o *WarehouseContact) SetIsInventoryModeActive(v bool) {
	o.IsInventoryModeActive = &v
}

// GetLogisticsType returns the LogisticsType field value if set, zero value otherwise.
func (o *WarehouseContact) GetLogisticsType() string {
	if o == nil || o.LogisticsType == nil {
		var ret string
		return ret
	}
	return *o.LogisticsType
}

// GetLogisticsTypeOk returns a tuple with the LogisticsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetLogisticsTypeOk() (*string, bool) {
	if o == nil || o.LogisticsType == nil {
		return nil, false
	}
	return o.LogisticsType, true
}

// HasLogisticsType returns a boolean if a field has been set.
func (o *WarehouseContact) HasLogisticsType() bool {
	if o != nil && o.LogisticsType != nil {
		return true
	}

	return false
}

// SetLogisticsType gets a reference to the given string and assigns it to the LogisticsType field.
func (o *WarehouseContact) SetLogisticsType(v string) {
	o.LogisticsType = &v
}

// GetAveragePriceSource returns the AveragePriceSource field value if set, zero value otherwise.
func (o *WarehouseContact) GetAveragePriceSource() string {
	if o == nil || o.AveragePriceSource == nil {
		var ret string
		return ret
	}
	return *o.AveragePriceSource
}

// GetAveragePriceSourceOk returns a tuple with the AveragePriceSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetAveragePriceSourceOk() (*string, bool) {
	if o == nil || o.AveragePriceSource == nil {
		return nil, false
	}
	return o.AveragePriceSource, true
}

// HasAveragePriceSource returns a boolean if a field has been set.
func (o *WarehouseContact) HasAveragePriceSource() bool {
	if o != nil && o.AveragePriceSource != nil {
		return true
	}

	return false
}

// SetAveragePriceSource gets a reference to the given string and assigns it to the AveragePriceSource field.
func (o *WarehouseContact) SetAveragePriceSource(v string) {
	o.AveragePriceSource = &v
}

// GetReorderLevelDynamic returns the ReorderLevelDynamic field value if set, zero value otherwise.
func (o *WarehouseContact) GetReorderLevelDynamic() string {
	if o == nil || o.ReorderLevelDynamic == nil {
		var ret string
		return ret
	}
	return *o.ReorderLevelDynamic
}

// GetReorderLevelDynamicOk returns a tuple with the ReorderLevelDynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetReorderLevelDynamicOk() (*string, bool) {
	if o == nil || o.ReorderLevelDynamic == nil {
		return nil, false
	}
	return o.ReorderLevelDynamic, true
}

// HasReorderLevelDynamic returns a boolean if a field has been set.
func (o *WarehouseContact) HasReorderLevelDynamic() bool {
	if o != nil && o.ReorderLevelDynamic != nil {
		return true
	}

	return false
}

// SetReorderLevelDynamic gets a reference to the given string and assigns it to the ReorderLevelDynamic field.
func (o *WarehouseContact) SetReorderLevelDynamic(v string) {
	o.ReorderLevelDynamic = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WarehouseContact) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WarehouseContact) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WarehouseContact) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WarehouseContact) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WarehouseContact) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WarehouseContact) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAllocationReferrerIds returns the AllocationReferrerIds field value if set, zero value otherwise.
func (o *WarehouseContact) GetAllocationReferrerIds() []string {
	if o == nil || o.AllocationReferrerIds == nil {
		var ret []string
		return ret
	}
	return *o.AllocationReferrerIds
}

// GetAllocationReferrerIdsOk returns a tuple with the AllocationReferrerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetAllocationReferrerIdsOk() (*[]string, bool) {
	if o == nil || o.AllocationReferrerIds == nil {
		return nil, false
	}
	return o.AllocationReferrerIds, true
}

// HasAllocationReferrerIds returns a boolean if a field has been set.
func (o *WarehouseContact) HasAllocationReferrerIds() bool {
	if o != nil && o.AllocationReferrerIds != nil {
		return true
	}

	return false
}

// SetAllocationReferrerIds gets a reference to the given []string and assigns it to the AllocationReferrerIds field.
func (o *WarehouseContact) SetAllocationReferrerIds(v []string) {
	o.AllocationReferrerIds = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *WarehouseContact) GetAddress() WarehouseContactAddress {
	if o == nil || o.Address == nil {
		var ret WarehouseContactAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContact) GetAddressOk() (*WarehouseContactAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *WarehouseContact) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given WarehouseContactAddress and assigns it to the Address field.
func (o *WarehouseContact) SetAddress(v WarehouseContactAddress) {
	o.Address = &v
}

func (o WarehouseContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if o.TypeId != nil {
		toSerialize["typeId"] = o.TypeId
	}
	if o.OnStockAvailability != nil {
		toSerialize["onStockAvailability"] = o.OnStockAvailability
	}
	if o.OutOfStockAvailability != nil {
		toSerialize["outOfStockAvailability"] = o.OutOfStockAvailability
	}
	if o.SplitByShippingProfile != nil {
		toSerialize["splitByShippingProfile"] = o.SplitByShippingProfile
	}
	if o.StorageLocationType != nil {
		toSerialize["storageLocationType"] = o.StorageLocationType
	}
	if o.StorageLocationZone != nil {
		toSerialize["storageLocationZone"] = o.StorageLocationZone
	}
	if o.RepairWarehouseId != nil {
		toSerialize["repairWarehouseId"] = o.RepairWarehouseId
	}
	if o.IsInventoryModeActive != nil {
		toSerialize["isInventoryModeActive"] = o.IsInventoryModeActive
	}
	if o.LogisticsType != nil {
		toSerialize["logisticsType"] = o.LogisticsType
	}
	if o.AveragePriceSource != nil {
		toSerialize["average_price_source"] = o.AveragePriceSource
	}
	if o.ReorderLevelDynamic != nil {
		toSerialize["reorder_level_dynamic"] = o.ReorderLevelDynamic
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.AllocationReferrerIds != nil {
		toSerialize["allocationReferrerIds"] = o.AllocationReferrerIds
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableWarehouseContact struct {
	value *WarehouseContact
	isSet bool
}

func (v NullableWarehouseContact) Get() *WarehouseContact {
	return v.value
}

func (v *NullableWarehouseContact) Set(val *WarehouseContact) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseContact) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseContact(val *WarehouseContact) *NullableWarehouseContact {
	return &NullableWarehouseContact{value: val, isSet: true}
}

func (v NullableWarehouseContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

