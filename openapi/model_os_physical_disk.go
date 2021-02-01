/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-11T18:30:19Z.
 *
 * API version: 1.0.9-3252
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// OsPhysicalDisk Physical Disk target to do Operating System installation on.
type OsPhysicalDisk struct {
	OsInstallTarget `yaml:"OsInstallTarget,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The Physical Disk Name to be used as Install Target.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// Serial Number of the Physical Disk Target.
	SerialNumber *string `json:"SerialNumber,omitempty" yaml:"SerialNumber,omitempty"`
	// The SlotID of the Storage Controller associated to the physical disk.
	StorageControllerSlotId *string `json:"StorageControllerSlotId,omitempty" yaml:"StorageControllerSlotId,omitempty"`
}

// NewOsPhysicalDisk instantiates a new OsPhysicalDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsPhysicalDisk(classId string, objectType string) *OsPhysicalDisk {
	this := OsPhysicalDisk{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsPhysicalDiskWithDefaults instantiates a new OsPhysicalDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsPhysicalDiskWithDefaults() *OsPhysicalDisk {
	this := OsPhysicalDisk{}
	var classId string = "os.PhysicalDisk"
	this.ClassId = classId
	var objectType string = "os.PhysicalDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsPhysicalDisk) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsPhysicalDisk) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsPhysicalDisk) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsPhysicalDisk) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsPhysicalDisk) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsPhysicalDisk) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsPhysicalDisk) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDisk) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsPhysicalDisk) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsPhysicalDisk) SetName(v string) {
	o.Name = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *OsPhysicalDisk) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDisk) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *OsPhysicalDisk) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *OsPhysicalDisk) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStorageControllerSlotId returns the StorageControllerSlotId field value if set, zero value otherwise.
func (o *OsPhysicalDisk) GetStorageControllerSlotId() string {
	if o == nil || o.StorageControllerSlotId == nil {
		var ret string
		return ret
	}
	return *o.StorageControllerSlotId
}

// GetStorageControllerSlotIdOk returns a tuple with the StorageControllerSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDisk) GetStorageControllerSlotIdOk() (*string, bool) {
	if o == nil || o.StorageControllerSlotId == nil {
		return nil, false
	}
	return o.StorageControllerSlotId, true
}

// HasStorageControllerSlotId returns a boolean if a field has been set.
func (o *OsPhysicalDisk) HasStorageControllerSlotId() bool {
	if o != nil && o.StorageControllerSlotId != nil {
		return true
	}

	return false
}

// SetStorageControllerSlotId gets a reference to the given string and assigns it to the StorageControllerSlotId field.
func (o *OsPhysicalDisk) SetStorageControllerSlotId(v string) {
	o.StorageControllerSlotId = &v
}

func (o OsPhysicalDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOsInstallTarget, errOsInstallTarget := json.Marshal(o.OsInstallTarget)
	if errOsInstallTarget != nil {
		return []byte{}, errOsInstallTarget
	}
	errOsInstallTarget = json.Unmarshal([]byte(serializedOsInstallTarget), &toSerialize)
	if errOsInstallTarget != nil {
		return []byte{}, errOsInstallTarget
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.StorageControllerSlotId != nil {
		toSerialize["StorageControllerSlotId"] = o.StorageControllerSlotId
	}
	return json.Marshal(toSerialize)
}

type NullableOsPhysicalDisk struct {
	value *OsPhysicalDisk
	isSet bool
}

func (v NullableOsPhysicalDisk) Get() *OsPhysicalDisk {
	return v.value
}

func (v *NullableOsPhysicalDisk) Set(val *OsPhysicalDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableOsPhysicalDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableOsPhysicalDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsPhysicalDisk(val *OsPhysicalDisk) *NullableOsPhysicalDisk {
	return &NullableOsPhysicalDisk{value: val, isSet: true}
}

func (v NullableOsPhysicalDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsPhysicalDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
