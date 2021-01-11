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

// StorageHitachiParityGroup A parity group in Hitachi storage array.
type StorageHitachiParityGroup struct {
	StorageBaseRaidGroup `yaml:"StorageBaseRaidGroup,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Speed (rpm) of the disk belonging to the parity group.
	DiskSpeed *string `json:"DiskSpeed,omitempty" yaml:"DiskSpeed,omitempty"`
	// Type of the disk belonging to the parity group.
	DiskType *string `json:"DiskType,omitempty" yaml:"DiskType,omitempty"`
	// Value of the accelerated compression of the parity group. true, Accelerated compression for the parity group is enabled. false, Accelerated compression for the parity group is disabled.
	IsAcceleratedCompressionEnabled *bool `json:"IsAcceleratedCompressionEnabled,omitempty" yaml:"IsAcceleratedCompressionEnabled,omitempty"`
	// Value of the copy back mode setting of the parity group. true, Copy back mode is enabled. false, Copy back mode is disabled.
	IsCopyBackModeEnabled *bool `json:"IsCopyBackModeEnabled,omitempty" yaml:"IsCopyBackModeEnabled,omitempty"`
	// Value of the encryption setting of the parity group. true, Encryption is enabled. false, Encryption is disabled.
	IsEncryptionEnabled *bool                                `json:"IsEncryptionEnabled,omitempty" yaml:"IsEncryptionEnabled,omitempty"`
	Array               *StorageHitachiArrayRelationship     `json:"Array,omitempty" yaml:"Array,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewStorageHitachiParityGroup instantiates a new StorageHitachiParityGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiParityGroup(classId string, objectType string) *StorageHitachiParityGroup {
	this := StorageHitachiParityGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiParityGroupWithDefaults instantiates a new StorageHitachiParityGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiParityGroupWithDefaults() *StorageHitachiParityGroup {
	this := StorageHitachiParityGroup{}
	var classId string = "storage.HitachiParityGroup"
	this.ClassId = classId
	var objectType string = "storage.HitachiParityGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiParityGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiParityGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiParityGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiParityGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDiskSpeed returns the DiskSpeed field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetDiskSpeed() string {
	if o == nil || o.DiskSpeed == nil {
		var ret string
		return ret
	}
	return *o.DiskSpeed
}

// GetDiskSpeedOk returns a tuple with the DiskSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetDiskSpeedOk() (*string, bool) {
	if o == nil || o.DiskSpeed == nil {
		return nil, false
	}
	return o.DiskSpeed, true
}

// HasDiskSpeed returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasDiskSpeed() bool {
	if o != nil && o.DiskSpeed != nil {
		return true
	}

	return false
}

// SetDiskSpeed gets a reference to the given string and assigns it to the DiskSpeed field.
func (o *StorageHitachiParityGroup) SetDiskSpeed(v string) {
	o.DiskSpeed = &v
}

// GetDiskType returns the DiskType field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetDiskType() string {
	if o == nil || o.DiskType == nil {
		var ret string
		return ret
	}
	return *o.DiskType
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetDiskTypeOk() (*string, bool) {
	if o == nil || o.DiskType == nil {
		return nil, false
	}
	return o.DiskType, true
}

// HasDiskType returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasDiskType() bool {
	if o != nil && o.DiskType != nil {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given string and assigns it to the DiskType field.
func (o *StorageHitachiParityGroup) SetDiskType(v string) {
	o.DiskType = &v
}

// GetIsAcceleratedCompressionEnabled returns the IsAcceleratedCompressionEnabled field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetIsAcceleratedCompressionEnabled() bool {
	if o == nil || o.IsAcceleratedCompressionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsAcceleratedCompressionEnabled
}

// GetIsAcceleratedCompressionEnabledOk returns a tuple with the IsAcceleratedCompressionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetIsAcceleratedCompressionEnabledOk() (*bool, bool) {
	if o == nil || o.IsAcceleratedCompressionEnabled == nil {
		return nil, false
	}
	return o.IsAcceleratedCompressionEnabled, true
}

// HasIsAcceleratedCompressionEnabled returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasIsAcceleratedCompressionEnabled() bool {
	if o != nil && o.IsAcceleratedCompressionEnabled != nil {
		return true
	}

	return false
}

// SetIsAcceleratedCompressionEnabled gets a reference to the given bool and assigns it to the IsAcceleratedCompressionEnabled field.
func (o *StorageHitachiParityGroup) SetIsAcceleratedCompressionEnabled(v bool) {
	o.IsAcceleratedCompressionEnabled = &v
}

// GetIsCopyBackModeEnabled returns the IsCopyBackModeEnabled field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetIsCopyBackModeEnabled() bool {
	if o == nil || o.IsCopyBackModeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCopyBackModeEnabled
}

// GetIsCopyBackModeEnabledOk returns a tuple with the IsCopyBackModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetIsCopyBackModeEnabledOk() (*bool, bool) {
	if o == nil || o.IsCopyBackModeEnabled == nil {
		return nil, false
	}
	return o.IsCopyBackModeEnabled, true
}

// HasIsCopyBackModeEnabled returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasIsCopyBackModeEnabled() bool {
	if o != nil && o.IsCopyBackModeEnabled != nil {
		return true
	}

	return false
}

// SetIsCopyBackModeEnabled gets a reference to the given bool and assigns it to the IsCopyBackModeEnabled field.
func (o *StorageHitachiParityGroup) SetIsCopyBackModeEnabled(v bool) {
	o.IsCopyBackModeEnabled = &v
}

// GetIsEncryptionEnabled returns the IsEncryptionEnabled field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetIsEncryptionEnabled() bool {
	if o == nil || o.IsEncryptionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEncryptionEnabled
}

// GetIsEncryptionEnabledOk returns a tuple with the IsEncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetIsEncryptionEnabledOk() (*bool, bool) {
	if o == nil || o.IsEncryptionEnabled == nil {
		return nil, false
	}
	return o.IsEncryptionEnabled, true
}

// HasIsEncryptionEnabled returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasIsEncryptionEnabled() bool {
	if o != nil && o.IsEncryptionEnabled != nil {
		return true
	}

	return false
}

// SetIsEncryptionEnabled gets a reference to the given bool and assigns it to the IsEncryptionEnabled field.
func (o *StorageHitachiParityGroup) SetIsEncryptionEnabled(v bool) {
	o.IsEncryptionEnabled = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiParityGroup) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiParityGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiParityGroup) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiParityGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiParityGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseRaidGroup, errStorageBaseRaidGroup := json.Marshal(o.StorageBaseRaidGroup)
	if errStorageBaseRaidGroup != nil {
		return []byte{}, errStorageBaseRaidGroup
	}
	errStorageBaseRaidGroup = json.Unmarshal([]byte(serializedStorageBaseRaidGroup), &toSerialize)
	if errStorageBaseRaidGroup != nil {
		return []byte{}, errStorageBaseRaidGroup
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DiskSpeed != nil {
		toSerialize["DiskSpeed"] = o.DiskSpeed
	}
	if o.DiskType != nil {
		toSerialize["DiskType"] = o.DiskType
	}
	if o.IsAcceleratedCompressionEnabled != nil {
		toSerialize["IsAcceleratedCompressionEnabled"] = o.IsAcceleratedCompressionEnabled
	}
	if o.IsCopyBackModeEnabled != nil {
		toSerialize["IsCopyBackModeEnabled"] = o.IsCopyBackModeEnabled
	}
	if o.IsEncryptionEnabled != nil {
		toSerialize["IsEncryptionEnabled"] = o.IsEncryptionEnabled
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableStorageHitachiParityGroup struct {
	value *StorageHitachiParityGroup
	isSet bool
}

func (v NullableStorageHitachiParityGroup) Get() *StorageHitachiParityGroup {
	return v.value
}

func (v *NullableStorageHitachiParityGroup) Set(val *StorageHitachiParityGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiParityGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiParityGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiParityGroup(val *StorageHitachiParityGroup) *NullableStorageHitachiParityGroup {
	return &NullableStorageHitachiParityGroup{value: val, isSet: true}
}

func (v NullableStorageHitachiParityGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiParityGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
