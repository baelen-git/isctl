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

// EquipmentRackEnclosureSlot Rack Server Slot in a RackEnclosure.
type EquipmentRackEnclosureSlot struct {
	EquipmentBase `yaml:"EquipmentBase,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Server ID which is part of Rack Enclosure Slot.
	RackId *int64 `json:"RackId,omitempty" yaml:"RackId,omitempty"`
	// Server DN which is part of Rack Enclosure Slot.
	RackUnitDn             *string                              `json:"RackUnitDn,omitempty" yaml:"RackUnitDn,omitempty"`
	EquipmentRackEnclosure *EquipmentRackEnclosureRelationship  `json:"EquipmentRackEnclosure,omitempty" yaml:"EquipmentRackEnclosure,omitempty"`
	InventoryDeviceInfo    *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty" yaml:"InventoryDeviceInfo,omitempty"`
	RackUnit               *ComputeRackUnitRelationship         `json:"RackUnit,omitempty" yaml:"RackUnit,omitempty"`
	RegisteredDevice       *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewEquipmentRackEnclosureSlot instantiates a new EquipmentRackEnclosureSlot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentRackEnclosureSlot(classId string, objectType string) *EquipmentRackEnclosureSlot {
	this := EquipmentRackEnclosureSlot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentRackEnclosureSlotWithDefaults instantiates a new EquipmentRackEnclosureSlot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentRackEnclosureSlotWithDefaults() *EquipmentRackEnclosureSlot {
	this := EquipmentRackEnclosureSlot{}
	var classId string = "equipment.RackEnclosureSlot"
	this.ClassId = classId
	var objectType string = "equipment.RackEnclosureSlot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentRackEnclosureSlot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureSlot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentRackEnclosureSlot) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentRackEnclosureSlot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureSlot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentRackEnclosureSlot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRackId returns the RackId field value if set, zero value otherwise.
func (o *EquipmentRackEnclosureSlot) GetRackId() int64 {
	if o == nil || o.RackId == nil {
		var ret int64
		return ret
	}
	return *o.RackId
}

// GetRackIdOk returns a tuple with the RackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureSlot) GetRackIdOk() (*int64, bool) {
	if o == nil || o.RackId == nil {
		return nil, false
	}
	return o.RackId, true
}

// HasRackId returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureSlot) HasRackId() bool {
	if o != nil && o.RackId != nil {
		return true
	}

	return false
}

// SetRackId gets a reference to the given int64 and assigns it to the RackId field.
func (o *EquipmentRackEnclosureSlot) SetRackId(v int64) {
	o.RackId = &v
}

// GetRackUnitDn returns the RackUnitDn field value if set, zero value otherwise.
func (o *EquipmentRackEnclosureSlot) GetRackUnitDn() string {
	if o == nil || o.RackUnitDn == nil {
		var ret string
		return ret
	}
	return *o.RackUnitDn
}

// GetRackUnitDnOk returns a tuple with the RackUnitDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureSlot) GetRackUnitDnOk() (*string, bool) {
	if o == nil || o.RackUnitDn == nil {
		return nil, false
	}
	return o.RackUnitDn, true
}

// HasRackUnitDn returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureSlot) HasRackUnitDn() bool {
	if o != nil && o.RackUnitDn != nil {
		return true
	}

	return false
}

// SetRackUnitDn gets a reference to the given string and assigns it to the RackUnitDn field.
func (o *EquipmentRackEnclosureSlot) SetRackUnitDn(v string) {
	o.RackUnitDn = &v
}

// GetEquipmentRackEnclosure returns the EquipmentRackEnclosure field value if set, zero value otherwise.
func (o *EquipmentRackEnclosureSlot) GetEquipmentRackEnclosure() EquipmentRackEnclosureRelationship {
	if o == nil || o.EquipmentRackEnclosure == nil {
		var ret EquipmentRackEnclosureRelationship
		return ret
	}
	return *o.EquipmentRackEnclosure
}

// GetEquipmentRackEnclosureOk returns a tuple with the EquipmentRackEnclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureSlot) GetEquipmentRackEnclosureOk() (*EquipmentRackEnclosureRelationship, bool) {
	if o == nil || o.EquipmentRackEnclosure == nil {
		return nil, false
	}
	return o.EquipmentRackEnclosure, true
}

// HasEquipmentRackEnclosure returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureSlot) HasEquipmentRackEnclosure() bool {
	if o != nil && o.EquipmentRackEnclosure != nil {
		return true
	}

	return false
}

// SetEquipmentRackEnclosure gets a reference to the given EquipmentRackEnclosureRelationship and assigns it to the EquipmentRackEnclosure field.
func (o *EquipmentRackEnclosureSlot) SetEquipmentRackEnclosure(v EquipmentRackEnclosureRelationship) {
	o.EquipmentRackEnclosure = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentRackEnclosureSlot) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureSlot) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureSlot) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentRackEnclosureSlot) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRackUnit returns the RackUnit field value if set, zero value otherwise.
func (o *EquipmentRackEnclosureSlot) GetRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.RackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.RackUnit
}

// GetRackUnitOk returns a tuple with the RackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureSlot) GetRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.RackUnit == nil {
		return nil, false
	}
	return o.RackUnit, true
}

// HasRackUnit returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureSlot) HasRackUnit() bool {
	if o != nil && o.RackUnit != nil {
		return true
	}

	return false
}

// SetRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the RackUnit field.
func (o *EquipmentRackEnclosureSlot) SetRackUnit(v ComputeRackUnitRelationship) {
	o.RackUnit = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentRackEnclosureSlot) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureSlot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureSlot) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentRackEnclosureSlot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentRackEnclosureSlot) MarshalJSON() ([]byte, error) {
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
	if o.RackId != nil {
		toSerialize["RackId"] = o.RackId
	}
	if o.RackUnitDn != nil {
		toSerialize["RackUnitDn"] = o.RackUnitDn
	}
	if o.EquipmentRackEnclosure != nil {
		toSerialize["EquipmentRackEnclosure"] = o.EquipmentRackEnclosure
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RackUnit != nil {
		toSerialize["RackUnit"] = o.RackUnit
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableEquipmentRackEnclosureSlot struct {
	value *EquipmentRackEnclosureSlot
	isSet bool
}

func (v NullableEquipmentRackEnclosureSlot) Get() *EquipmentRackEnclosureSlot {
	return v.value
}

func (v *NullableEquipmentRackEnclosureSlot) Set(val *EquipmentRackEnclosureSlot) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentRackEnclosureSlot) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentRackEnclosureSlot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentRackEnclosureSlot(val *EquipmentRackEnclosureSlot) *NullableEquipmentRackEnclosureSlot {
	return &NullableEquipmentRackEnclosureSlot{value: val, isSet: true}
}

func (v NullableEquipmentRackEnclosureSlot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentRackEnclosureSlot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
