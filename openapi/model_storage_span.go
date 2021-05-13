/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// StorageSpan Group of disks to configure virtual drive.
type StorageSpan struct {
	InventoryBase `yaml:"InventoryBase,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string  `json:"ObjectType" yaml:"ObjectType"`
	Slots      []int64 `json:"Slots,omitempty" yaml:"Slots,omitempty"`
	// Unique identifier value of this span.
	SpanId    *int64                        `json:"SpanId,omitempty" yaml:"SpanId,omitempty"`
	DiskGroup *StorageDiskGroupRelationship `json:"DiskGroup,omitempty" yaml:"DiskGroup,omitempty"`
	// An array of relationships to storagePhysicalDisk resources.
	PhysicalDisks    []StoragePhysicalDiskRelationship    `json:"PhysicalDisks,omitempty" yaml:"PhysicalDisks,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewStorageSpan instantiates a new StorageSpan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSpan(classId string, objectType string) *StorageSpan {
	this := StorageSpan{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageSpanWithDefaults instantiates a new StorageSpan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSpanWithDefaults() *StorageSpan {
	this := StorageSpan{}
	var classId string = "storage.Span"
	this.ClassId = classId
	var objectType string = "storage.Span"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageSpan) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageSpan) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageSpan) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageSpan) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageSpan) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageSpan) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSlots returns the Slots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageSpan) GetSlots() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageSpan) GetSlotsOk() (*[]int64, bool) {
	if o == nil || o.Slots == nil {
		return nil, false
	}
	return &o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *StorageSpan) HasSlots() bool {
	if o != nil && o.Slots != nil {
		return true
	}

	return false
}

// SetSlots gets a reference to the given []int64 and assigns it to the Slots field.
func (o *StorageSpan) SetSlots(v []int64) {
	o.Slots = v
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *StorageSpan) GetSpanId() int64 {
	if o == nil || o.SpanId == nil {
		var ret int64
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpan) GetSpanIdOk() (*int64, bool) {
	if o == nil || o.SpanId == nil {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *StorageSpan) HasSpanId() bool {
	if o != nil && o.SpanId != nil {
		return true
	}

	return false
}

// SetSpanId gets a reference to the given int64 and assigns it to the SpanId field.
func (o *StorageSpan) SetSpanId(v int64) {
	o.SpanId = &v
}

// GetDiskGroup returns the DiskGroup field value if set, zero value otherwise.
func (o *StorageSpan) GetDiskGroup() StorageDiskGroupRelationship {
	if o == nil || o.DiskGroup == nil {
		var ret StorageDiskGroupRelationship
		return ret
	}
	return *o.DiskGroup
}

// GetDiskGroupOk returns a tuple with the DiskGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpan) GetDiskGroupOk() (*StorageDiskGroupRelationship, bool) {
	if o == nil || o.DiskGroup == nil {
		return nil, false
	}
	return o.DiskGroup, true
}

// HasDiskGroup returns a boolean if a field has been set.
func (o *StorageSpan) HasDiskGroup() bool {
	if o != nil && o.DiskGroup != nil {
		return true
	}

	return false
}

// SetDiskGroup gets a reference to the given StorageDiskGroupRelationship and assigns it to the DiskGroup field.
func (o *StorageSpan) SetDiskGroup(v StorageDiskGroupRelationship) {
	o.DiskGroup = &v
}

// GetPhysicalDisks returns the PhysicalDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageSpan) GetPhysicalDisks() []StoragePhysicalDiskRelationship {
	if o == nil {
		var ret []StoragePhysicalDiskRelationship
		return ret
	}
	return o.PhysicalDisks
}

// GetPhysicalDisksOk returns a tuple with the PhysicalDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageSpan) GetPhysicalDisksOk() (*[]StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.PhysicalDisks == nil {
		return nil, false
	}
	return &o.PhysicalDisks, true
}

// HasPhysicalDisks returns a boolean if a field has been set.
func (o *StorageSpan) HasPhysicalDisks() bool {
	if o != nil && o.PhysicalDisks != nil {
		return true
	}

	return false
}

// SetPhysicalDisks gets a reference to the given []StoragePhysicalDiskRelationship and assigns it to the PhysicalDisks field.
func (o *StorageSpan) SetPhysicalDisks(v []StoragePhysicalDiskRelationship) {
	o.PhysicalDisks = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageSpan) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpan) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageSpan) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageSpan) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageSpan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Slots != nil {
		toSerialize["Slots"] = o.Slots
	}
	if o.SpanId != nil {
		toSerialize["SpanId"] = o.SpanId
	}
	if o.DiskGroup != nil {
		toSerialize["DiskGroup"] = o.DiskGroup
	}
	if o.PhysicalDisks != nil {
		toSerialize["PhysicalDisks"] = o.PhysicalDisks
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableStorageSpan struct {
	value *StorageSpan
	isSet bool
}

func (v NullableStorageSpan) Get() *StorageSpan {
	return v.value
}

func (v *NullableStorageSpan) Set(val *StorageSpan) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSpan) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSpan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSpan(val *StorageSpan) *NullableStorageSpan {
	return &NullableStorageSpan{value: val, isSet: true}
}

func (v NullableStorageSpan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSpan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
