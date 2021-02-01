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

// StorageNetAppVolumeSnapshot NetApp Volume Snapshot is a read-only image of a traditional or FlexVol volume, or an aggregate, that captures the state of the file system at a point in time.
type StorageNetAppVolumeSnapshot struct {
	StorageBaseSnapshot `yaml:"StorageBaseSnapshot,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// UUID of the volume snapshot.
	Uuid             *string                           `json:"Uuid,omitempty" yaml:"Uuid,omitempty"`
	Array            *StorageNetAppClusterRelationship `json:"Array,omitempty" yaml:"Array,omitempty"`
	StorageContainer *StorageNetAppVolumeRelationship  `json:"StorageContainer,omitempty" yaml:"StorageContainer,omitempty"`
}

// NewStorageNetAppVolumeSnapshot instantiates a new StorageNetAppVolumeSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppVolumeSnapshot(classId string, objectType string) *StorageNetAppVolumeSnapshot {
	this := StorageNetAppVolumeSnapshot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppVolumeSnapshotWithDefaults instantiates a new StorageNetAppVolumeSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppVolumeSnapshotWithDefaults() *StorageNetAppVolumeSnapshot {
	this := StorageNetAppVolumeSnapshot{}
	var classId string = "storage.NetAppVolumeSnapshot"
	this.ClassId = classId
	var objectType string = "storage.NetAppVolumeSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppVolumeSnapshot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeSnapshot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppVolumeSnapshot) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppVolumeSnapshot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeSnapshot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppVolumeSnapshot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppVolumeSnapshot) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeSnapshot) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppVolumeSnapshot) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppVolumeSnapshot) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppVolumeSnapshot) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeSnapshot) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppVolumeSnapshot) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppVolumeSnapshot) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *StorageNetAppVolumeSnapshot) GetStorageContainer() StorageNetAppVolumeRelationship {
	if o == nil || o.StorageContainer == nil {
		var ret StorageNetAppVolumeRelationship
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeSnapshot) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool) {
	if o == nil || o.StorageContainer == nil {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *StorageNetAppVolumeSnapshot) HasStorageContainer() bool {
	if o != nil && o.StorageContainer != nil {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given StorageNetAppVolumeRelationship and assigns it to the StorageContainer field.
func (o *StorageNetAppVolumeSnapshot) SetStorageContainer(v StorageNetAppVolumeRelationship) {
	o.StorageContainer = &v
}

func (o StorageNetAppVolumeSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseSnapshot, errStorageBaseSnapshot := json.Marshal(o.StorageBaseSnapshot)
	if errStorageBaseSnapshot != nil {
		return []byte{}, errStorageBaseSnapshot
	}
	errStorageBaseSnapshot = json.Unmarshal([]byte(serializedStorageBaseSnapshot), &toSerialize)
	if errStorageBaseSnapshot != nil {
		return []byte{}, errStorageBaseSnapshot
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.StorageContainer != nil {
		toSerialize["StorageContainer"] = o.StorageContainer
	}
	return json.Marshal(toSerialize)
}

type NullableStorageNetAppVolumeSnapshot struct {
	value *StorageNetAppVolumeSnapshot
	isSet bool
}

func (v NullableStorageNetAppVolumeSnapshot) Get() *StorageNetAppVolumeSnapshot {
	return v.value
}

func (v *NullableStorageNetAppVolumeSnapshot) Set(val *StorageNetAppVolumeSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppVolumeSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppVolumeSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppVolumeSnapshot(val *StorageNetAppVolumeSnapshot) *NullableStorageNetAppVolumeSnapshot {
	return &NullableStorageNetAppVolumeSnapshot{value: val, isSet: true}
}

func (v NullableStorageNetAppVolumeSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppVolumeSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
