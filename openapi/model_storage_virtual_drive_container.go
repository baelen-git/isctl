/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-12-08T20:53:20Z.
 *
 * API version: 1.0.9-2908
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// StorageVirtualDriveContainer A Virtual Disk Drive Container.
type StorageVirtualDriveContainer struct {
	EquipmentBase `yaml:"EquipmentBase,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The identifier for this container.
	ContainerId         *int64                               `json:"ContainerId,omitempty" yaml:"ContainerId,omitempty"`
	EquipmentChassis    *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty" yaml:"EquipmentChassis,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty" yaml:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
	// An array of relationships to storageVirtualDrive resources.
	VirtualDrive []StorageVirtualDriveRelationship `json:"VirtualDrive,omitempty" yaml:"VirtualDrive,omitempty"`
}

// NewStorageVirtualDriveContainer instantiates a new StorageVirtualDriveContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveContainer(classId string, objectType string) *StorageVirtualDriveContainer {
	this := StorageVirtualDriveContainer{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageVirtualDriveContainerWithDefaults instantiates a new StorageVirtualDriveContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveContainerWithDefaults() *StorageVirtualDriveContainer {
	this := StorageVirtualDriveContainer{}
	var classId string = "storage.VirtualDriveContainer"
	this.ClassId = classId
	var objectType string = "storage.VirtualDriveContainer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDriveContainer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainer) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDriveContainer) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDriveContainer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainer) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDriveContainer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *StorageVirtualDriveContainer) GetContainerId() int64 {
	if o == nil || o.ContainerId == nil {
		var ret int64
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainer) GetContainerIdOk() (*int64, bool) {
	if o == nil || o.ContainerId == nil {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *StorageVirtualDriveContainer) HasContainerId() bool {
	if o != nil && o.ContainerId != nil {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given int64 and assigns it to the ContainerId field.
func (o *StorageVirtualDriveContainer) SetContainerId(v int64) {
	o.ContainerId = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *StorageVirtualDriveContainer) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainer) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *StorageVirtualDriveContainer) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *StorageVirtualDriveContainer) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageVirtualDriveContainer) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainer) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageVirtualDriveContainer) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageVirtualDriveContainer) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageVirtualDriveContainer) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveContainer) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageVirtualDriveContainer) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageVirtualDriveContainer) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageVirtualDriveContainer) GetVirtualDrive() []StorageVirtualDriveRelationship {
	if o == nil {
		var ret []StorageVirtualDriveRelationship
		return ret
	}
	return o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageVirtualDriveContainer) GetVirtualDriveOk() (*[]StorageVirtualDriveRelationship, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return &o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveContainer) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given []StorageVirtualDriveRelationship and assigns it to the VirtualDrive field.
func (o *StorageVirtualDriveContainer) SetVirtualDrive(v []StorageVirtualDriveRelationship) {
	o.VirtualDrive = v
}

func (o StorageVirtualDriveContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ContainerId != nil {
		toSerialize["ContainerId"] = o.ContainerId
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}
	return json.Marshal(toSerialize)
}

type NullableStorageVirtualDriveContainer struct {
	value *StorageVirtualDriveContainer
	isSet bool
}

func (v NullableStorageVirtualDriveContainer) Get() *StorageVirtualDriveContainer {
	return v.value
}

func (v *NullableStorageVirtualDriveContainer) Set(val *StorageVirtualDriveContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveContainer(val *StorageVirtualDriveContainer) *NullableStorageVirtualDriveContainer {
	return &NullableStorageVirtualDriveContainer{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
