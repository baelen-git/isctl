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

// MemoryPersistentMemoryNamespaceAllOf Definition of the list of properties defined in 'memory.PersistentMemoryNamespace', excluding properties defined in parent classes.
type MemoryPersistentMemoryNamespaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Capacity in GiB of the Persistent Memory Namespace.
	Capacity *string `json:"Capacity,omitempty" yaml:"Capacity,omitempty"`
	// Health state of the Persistent Memory Namespace.
	HealthState *string `json:"HealthState,omitempty" yaml:"HealthState,omitempty"`
	// Label version of the Persistent Memory Namespace.
	LabelVersion *string `json:"LabelVersion,omitempty" yaml:"LabelVersion,omitempty"`
	// Mode of the Persistent Memory Namespace.
	Mode *string `json:"Mode,omitempty" yaml:"Mode,omitempty"`
	// Name of the Persistent Memory Namespace.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// UUID of the Persistent Memory Namespace.
	Uuid                         *string                                   `json:"Uuid,omitempty" yaml:"Uuid,omitempty"`
	InventoryDeviceInfo          *InventoryDeviceInfoRelationship          `json:"InventoryDeviceInfo,omitempty" yaml:"InventoryDeviceInfo,omitempty"`
	MemoryPersistentMemoryRegion *MemoryPersistentMemoryRegionRelationship `json:"MemoryPersistentMemoryRegion,omitempty" yaml:"MemoryPersistentMemoryRegion,omitempty"`
	RegisteredDevice             *AssetDeviceRegistrationRelationship      `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewMemoryPersistentMemoryNamespaceAllOf instantiates a new MemoryPersistentMemoryNamespaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryNamespaceAllOf(classId string, objectType string) *MemoryPersistentMemoryNamespaceAllOf {
	this := MemoryPersistentMemoryNamespaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryPersistentMemoryNamespaceAllOfWithDefaults instantiates a new MemoryPersistentMemoryNamespaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryNamespaceAllOfWithDefaults() *MemoryPersistentMemoryNamespaceAllOf {
	this := MemoryPersistentMemoryNamespaceAllOf{}
	var classId string = "memory.PersistentMemoryNamespace"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryNamespace"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryNamespaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryNamespaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryNamespaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryNamespaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *MemoryPersistentMemoryNamespaceAllOf) SetCapacity(v string) {
	o.Capacity = &v
}

// GetHealthState returns the HealthState field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetHealthState() string {
	if o == nil || o.HealthState == nil {
		var ret string
		return ret
	}
	return *o.HealthState
}

// GetHealthStateOk returns a tuple with the HealthState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetHealthStateOk() (*string, bool) {
	if o == nil || o.HealthState == nil {
		return nil, false
	}
	return o.HealthState, true
}

// HasHealthState returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) HasHealthState() bool {
	if o != nil && o.HealthState != nil {
		return true
	}

	return false
}

// SetHealthState gets a reference to the given string and assigns it to the HealthState field.
func (o *MemoryPersistentMemoryNamespaceAllOf) SetHealthState(v string) {
	o.HealthState = &v
}

// GetLabelVersion returns the LabelVersion field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetLabelVersion() string {
	if o == nil || o.LabelVersion == nil {
		var ret string
		return ret
	}
	return *o.LabelVersion
}

// GetLabelVersionOk returns a tuple with the LabelVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetLabelVersionOk() (*string, bool) {
	if o == nil || o.LabelVersion == nil {
		return nil, false
	}
	return o.LabelVersion, true
}

// HasLabelVersion returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) HasLabelVersion() bool {
	if o != nil && o.LabelVersion != nil {
		return true
	}

	return false
}

// SetLabelVersion gets a reference to the given string and assigns it to the LabelVersion field.
func (o *MemoryPersistentMemoryNamespaceAllOf) SetLabelVersion(v string) {
	o.LabelVersion = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *MemoryPersistentMemoryNamespaceAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MemoryPersistentMemoryNamespaceAllOf) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *MemoryPersistentMemoryNamespaceAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryPersistentMemoryNamespaceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetMemoryPersistentMemoryRegion returns the MemoryPersistentMemoryRegion field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetMemoryPersistentMemoryRegion() MemoryPersistentMemoryRegionRelationship {
	if o == nil || o.MemoryPersistentMemoryRegion == nil {
		var ret MemoryPersistentMemoryRegionRelationship
		return ret
	}
	return *o.MemoryPersistentMemoryRegion
}

// GetMemoryPersistentMemoryRegionOk returns a tuple with the MemoryPersistentMemoryRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetMemoryPersistentMemoryRegionOk() (*MemoryPersistentMemoryRegionRelationship, bool) {
	if o == nil || o.MemoryPersistentMemoryRegion == nil {
		return nil, false
	}
	return o.MemoryPersistentMemoryRegion, true
}

// HasMemoryPersistentMemoryRegion returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) HasMemoryPersistentMemoryRegion() bool {
	if o != nil && o.MemoryPersistentMemoryRegion != nil {
		return true
	}

	return false
}

// SetMemoryPersistentMemoryRegion gets a reference to the given MemoryPersistentMemoryRegionRelationship and assigns it to the MemoryPersistentMemoryRegion field.
func (o *MemoryPersistentMemoryNamespaceAllOf) SetMemoryPersistentMemoryRegion(v MemoryPersistentMemoryRegionRelationship) {
	o.MemoryPersistentMemoryRegion = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryPersistentMemoryNamespaceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o MemoryPersistentMemoryNamespaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.HealthState != nil {
		toSerialize["HealthState"] = o.HealthState
	}
	if o.LabelVersion != nil {
		toSerialize["LabelVersion"] = o.LabelVersion
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.MemoryPersistentMemoryRegion != nil {
		toSerialize["MemoryPersistentMemoryRegion"] = o.MemoryPersistentMemoryRegion
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableMemoryPersistentMemoryNamespaceAllOf struct {
	value *MemoryPersistentMemoryNamespaceAllOf
	isSet bool
}

func (v NullableMemoryPersistentMemoryNamespaceAllOf) Get() *MemoryPersistentMemoryNamespaceAllOf {
	return v.value
}

func (v *NullableMemoryPersistentMemoryNamespaceAllOf) Set(val *MemoryPersistentMemoryNamespaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryNamespaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryNamespaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryNamespaceAllOf(val *MemoryPersistentMemoryNamespaceAllOf) *NullableMemoryPersistentMemoryNamespaceAllOf {
	return &NullableMemoryPersistentMemoryNamespaceAllOf{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryNamespaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryNamespaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
