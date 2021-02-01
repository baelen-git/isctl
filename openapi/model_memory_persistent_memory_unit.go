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

// MemoryPersistentMemoryUnit Persistent Memory Module on a server.
type MemoryPersistentMemoryUnit struct {
	MemoryAbstractUnit `yaml:"MemoryAbstractUnit,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// AppDirect capacity in GiB of the Persistent Memory Module on a server.
	AppDirectCapacity *string `json:"AppDirectCapacity,omitempty" yaml:"AppDirectCapacity,omitempty"`
	// Count status of the Persistent Memory Module on a server.
	CountStatus *string `json:"CountStatus,omitempty" yaml:"CountStatus,omitempty"`
	// Firmware version of the firware running on the Persistent Memory Module on a server.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" yaml:"FirmwareVersion,omitempty"`
	// Frozen status of the Persistent Memory Module on a server.
	FrozenStatus *string `json:"FrozenStatus,omitempty" yaml:"FrozenStatus,omitempty"`
	// Health state of the Persistent Memory Module on a server.
	HealthState *string `json:"HealthState,omitempty" yaml:"HealthState,omitempty"`
	// Lock status of the Persistent Memory Module on a server.
	LockStatus *string `json:"LockStatus,omitempty" yaml:"LockStatus,omitempty"`
	// Memory capacity in GiB of the Persistent Memory Module on a server.
	MemoryCapacity *string `json:"MemoryCapacity,omitempty" yaml:"MemoryCapacity,omitempty"`
	// ID of the Persistent Memory Module on a server.
	MemoryId *int64 `json:"MemoryId,omitempty" yaml:"MemoryId,omitempty"`
	// Persistent Memory capacity in GiB of the Persistent Memory Module on a server.
	PersistentMemoryCapacity *string `json:"PersistentMemoryCapacity,omitempty" yaml:"PersistentMemoryCapacity,omitempty"`
	// Reserved capacity in GiB of the Persistent Memory Module on a server.
	ReservedCapacity *string `json:"ReservedCapacity,omitempty" yaml:"ReservedCapacity,omitempty"`
	// Security status of the Persistent Memory Module on a server.
	SecurityStatus *string `json:"SecurityStatus,omitempty" yaml:"SecurityStatus,omitempty"`
	// Socket ID of the Persistent Memory Module on a server.
	SocketId *string `json:"SocketId,omitempty" yaml:"SocketId,omitempty"`
	// Socket Memory ID of the Persistent Memory Module on a server.
	SocketMemoryId *string `json:"SocketMemoryId,omitempty" yaml:"SocketMemoryId,omitempty"`
	// Total capacity in GiB of the Persistent Memory Module on a server.
	TotalCapacity *string `json:"TotalCapacity,omitempty" yaml:"TotalCapacity,omitempty"`
	// UID of the Persistent Memory Module on a server.
	Uid                 *string                              `json:"Uid,omitempty" yaml:"Uid,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty" yaml:"InventoryDeviceInfo,omitempty"`
	MemoryArray         *MemoryArrayRelationship             `json:"MemoryArray,omitempty" yaml:"MemoryArray,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewMemoryPersistentMemoryUnit instantiates a new MemoryPersistentMemoryUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryUnit(classId string, objectType string) *MemoryPersistentMemoryUnit {
	this := MemoryPersistentMemoryUnit{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryPersistentMemoryUnitWithDefaults instantiates a new MemoryPersistentMemoryUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryUnitWithDefaults() *MemoryPersistentMemoryUnit {
	this := MemoryPersistentMemoryUnit{}
	var classId string = "memory.PersistentMemoryUnit"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryUnit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryUnit) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryUnit) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryUnit) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryUnit) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAppDirectCapacity returns the AppDirectCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetAppDirectCapacity() string {
	if o == nil || o.AppDirectCapacity == nil {
		var ret string
		return ret
	}
	return *o.AppDirectCapacity
}

// GetAppDirectCapacityOk returns a tuple with the AppDirectCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetAppDirectCapacityOk() (*string, bool) {
	if o == nil || o.AppDirectCapacity == nil {
		return nil, false
	}
	return o.AppDirectCapacity, true
}

// HasAppDirectCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasAppDirectCapacity() bool {
	if o != nil && o.AppDirectCapacity != nil {
		return true
	}

	return false
}

// SetAppDirectCapacity gets a reference to the given string and assigns it to the AppDirectCapacity field.
func (o *MemoryPersistentMemoryUnit) SetAppDirectCapacity(v string) {
	o.AppDirectCapacity = &v
}

// GetCountStatus returns the CountStatus field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetCountStatus() string {
	if o == nil || o.CountStatus == nil {
		var ret string
		return ret
	}
	return *o.CountStatus
}

// GetCountStatusOk returns a tuple with the CountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetCountStatusOk() (*string, bool) {
	if o == nil || o.CountStatus == nil {
		return nil, false
	}
	return o.CountStatus, true
}

// HasCountStatus returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasCountStatus() bool {
	if o != nil && o.CountStatus != nil {
		return true
	}

	return false
}

// SetCountStatus gets a reference to the given string and assigns it to the CountStatus field.
func (o *MemoryPersistentMemoryUnit) SetCountStatus(v string) {
	o.CountStatus = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *MemoryPersistentMemoryUnit) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetFrozenStatus returns the FrozenStatus field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetFrozenStatus() string {
	if o == nil || o.FrozenStatus == nil {
		var ret string
		return ret
	}
	return *o.FrozenStatus
}

// GetFrozenStatusOk returns a tuple with the FrozenStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetFrozenStatusOk() (*string, bool) {
	if o == nil || o.FrozenStatus == nil {
		return nil, false
	}
	return o.FrozenStatus, true
}

// HasFrozenStatus returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasFrozenStatus() bool {
	if o != nil && o.FrozenStatus != nil {
		return true
	}

	return false
}

// SetFrozenStatus gets a reference to the given string and assigns it to the FrozenStatus field.
func (o *MemoryPersistentMemoryUnit) SetFrozenStatus(v string) {
	o.FrozenStatus = &v
}

// GetHealthState returns the HealthState field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetHealthState() string {
	if o == nil || o.HealthState == nil {
		var ret string
		return ret
	}
	return *o.HealthState
}

// GetHealthStateOk returns a tuple with the HealthState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetHealthStateOk() (*string, bool) {
	if o == nil || o.HealthState == nil {
		return nil, false
	}
	return o.HealthState, true
}

// HasHealthState returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasHealthState() bool {
	if o != nil && o.HealthState != nil {
		return true
	}

	return false
}

// SetHealthState gets a reference to the given string and assigns it to the HealthState field.
func (o *MemoryPersistentMemoryUnit) SetHealthState(v string) {
	o.HealthState = &v
}

// GetLockStatus returns the LockStatus field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetLockStatus() string {
	if o == nil || o.LockStatus == nil {
		var ret string
		return ret
	}
	return *o.LockStatus
}

// GetLockStatusOk returns a tuple with the LockStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetLockStatusOk() (*string, bool) {
	if o == nil || o.LockStatus == nil {
		return nil, false
	}
	return o.LockStatus, true
}

// HasLockStatus returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasLockStatus() bool {
	if o != nil && o.LockStatus != nil {
		return true
	}

	return false
}

// SetLockStatus gets a reference to the given string and assigns it to the LockStatus field.
func (o *MemoryPersistentMemoryUnit) SetLockStatus(v string) {
	o.LockStatus = &v
}

// GetMemoryCapacity returns the MemoryCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetMemoryCapacity() string {
	if o == nil || o.MemoryCapacity == nil {
		var ret string
		return ret
	}
	return *o.MemoryCapacity
}

// GetMemoryCapacityOk returns a tuple with the MemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetMemoryCapacityOk() (*string, bool) {
	if o == nil || o.MemoryCapacity == nil {
		return nil, false
	}
	return o.MemoryCapacity, true
}

// HasMemoryCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasMemoryCapacity() bool {
	if o != nil && o.MemoryCapacity != nil {
		return true
	}

	return false
}

// SetMemoryCapacity gets a reference to the given string and assigns it to the MemoryCapacity field.
func (o *MemoryPersistentMemoryUnit) SetMemoryCapacity(v string) {
	o.MemoryCapacity = &v
}

// GetMemoryId returns the MemoryId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetMemoryId() int64 {
	if o == nil || o.MemoryId == nil {
		var ret int64
		return ret
	}
	return *o.MemoryId
}

// GetMemoryIdOk returns a tuple with the MemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetMemoryIdOk() (*int64, bool) {
	if o == nil || o.MemoryId == nil {
		return nil, false
	}
	return o.MemoryId, true
}

// HasMemoryId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasMemoryId() bool {
	if o != nil && o.MemoryId != nil {
		return true
	}

	return false
}

// SetMemoryId gets a reference to the given int64 and assigns it to the MemoryId field.
func (o *MemoryPersistentMemoryUnit) SetMemoryId(v int64) {
	o.MemoryId = &v
}

// GetPersistentMemoryCapacity returns the PersistentMemoryCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetPersistentMemoryCapacity() string {
	if o == nil || o.PersistentMemoryCapacity == nil {
		var ret string
		return ret
	}
	return *o.PersistentMemoryCapacity
}

// GetPersistentMemoryCapacityOk returns a tuple with the PersistentMemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetPersistentMemoryCapacityOk() (*string, bool) {
	if o == nil || o.PersistentMemoryCapacity == nil {
		return nil, false
	}
	return o.PersistentMemoryCapacity, true
}

// HasPersistentMemoryCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasPersistentMemoryCapacity() bool {
	if o != nil && o.PersistentMemoryCapacity != nil {
		return true
	}

	return false
}

// SetPersistentMemoryCapacity gets a reference to the given string and assigns it to the PersistentMemoryCapacity field.
func (o *MemoryPersistentMemoryUnit) SetPersistentMemoryCapacity(v string) {
	o.PersistentMemoryCapacity = &v
}

// GetReservedCapacity returns the ReservedCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetReservedCapacity() string {
	if o == nil || o.ReservedCapacity == nil {
		var ret string
		return ret
	}
	return *o.ReservedCapacity
}

// GetReservedCapacityOk returns a tuple with the ReservedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetReservedCapacityOk() (*string, bool) {
	if o == nil || o.ReservedCapacity == nil {
		return nil, false
	}
	return o.ReservedCapacity, true
}

// HasReservedCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasReservedCapacity() bool {
	if o != nil && o.ReservedCapacity != nil {
		return true
	}

	return false
}

// SetReservedCapacity gets a reference to the given string and assigns it to the ReservedCapacity field.
func (o *MemoryPersistentMemoryUnit) SetReservedCapacity(v string) {
	o.ReservedCapacity = &v
}

// GetSecurityStatus returns the SecurityStatus field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetSecurityStatus() string {
	if o == nil || o.SecurityStatus == nil {
		var ret string
		return ret
	}
	return *o.SecurityStatus
}

// GetSecurityStatusOk returns a tuple with the SecurityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetSecurityStatusOk() (*string, bool) {
	if o == nil || o.SecurityStatus == nil {
		return nil, false
	}
	return o.SecurityStatus, true
}

// HasSecurityStatus returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasSecurityStatus() bool {
	if o != nil && o.SecurityStatus != nil {
		return true
	}

	return false
}

// SetSecurityStatus gets a reference to the given string and assigns it to the SecurityStatus field.
func (o *MemoryPersistentMemoryUnit) SetSecurityStatus(v string) {
	o.SecurityStatus = &v
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetSocketId() string {
	if o == nil || o.SocketId == nil {
		var ret string
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetSocketIdOk() (*string, bool) {
	if o == nil || o.SocketId == nil {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasSocketId() bool {
	if o != nil && o.SocketId != nil {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given string and assigns it to the SocketId field.
func (o *MemoryPersistentMemoryUnit) SetSocketId(v string) {
	o.SocketId = &v
}

// GetSocketMemoryId returns the SocketMemoryId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetSocketMemoryId() string {
	if o == nil || o.SocketMemoryId == nil {
		var ret string
		return ret
	}
	return *o.SocketMemoryId
}

// GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetSocketMemoryIdOk() (*string, bool) {
	if o == nil || o.SocketMemoryId == nil {
		return nil, false
	}
	return o.SocketMemoryId, true
}

// HasSocketMemoryId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasSocketMemoryId() bool {
	if o != nil && o.SocketMemoryId != nil {
		return true
	}

	return false
}

// SetSocketMemoryId gets a reference to the given string and assigns it to the SocketMemoryId field.
func (o *MemoryPersistentMemoryUnit) SetSocketMemoryId(v string) {
	o.SocketMemoryId = &v
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetTotalCapacity() string {
	if o == nil || o.TotalCapacity == nil {
		var ret string
		return ret
	}
	return *o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetTotalCapacityOk() (*string, bool) {
	if o == nil || o.TotalCapacity == nil {
		return nil, false
	}
	return o.TotalCapacity, true
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasTotalCapacity() bool {
	if o != nil && o.TotalCapacity != nil {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given string and assigns it to the TotalCapacity field.
func (o *MemoryPersistentMemoryUnit) SetTotalCapacity(v string) {
	o.TotalCapacity = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *MemoryPersistentMemoryUnit) SetUid(v string) {
	o.Uid = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryPersistentMemoryUnit) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetMemoryArray returns the MemoryArray field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetMemoryArray() MemoryArrayRelationship {
	if o == nil || o.MemoryArray == nil {
		var ret MemoryArrayRelationship
		return ret
	}
	return *o.MemoryArray
}

// GetMemoryArrayOk returns a tuple with the MemoryArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetMemoryArrayOk() (*MemoryArrayRelationship, bool) {
	if o == nil || o.MemoryArray == nil {
		return nil, false
	}
	return o.MemoryArray, true
}

// HasMemoryArray returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasMemoryArray() bool {
	if o != nil && o.MemoryArray != nil {
		return true
	}

	return false
}

// SetMemoryArray gets a reference to the given MemoryArrayRelationship and assigns it to the MemoryArray field.
func (o *MemoryPersistentMemoryUnit) SetMemoryArray(v MemoryArrayRelationship) {
	o.MemoryArray = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryUnit) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryPersistentMemoryUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o MemoryPersistentMemoryUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMemoryAbstractUnit, errMemoryAbstractUnit := json.Marshal(o.MemoryAbstractUnit)
	if errMemoryAbstractUnit != nil {
		return []byte{}, errMemoryAbstractUnit
	}
	errMemoryAbstractUnit = json.Unmarshal([]byte(serializedMemoryAbstractUnit), &toSerialize)
	if errMemoryAbstractUnit != nil {
		return []byte{}, errMemoryAbstractUnit
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AppDirectCapacity != nil {
		toSerialize["AppDirectCapacity"] = o.AppDirectCapacity
	}
	if o.CountStatus != nil {
		toSerialize["CountStatus"] = o.CountStatus
	}
	if o.FirmwareVersion != nil {
		toSerialize["FirmwareVersion"] = o.FirmwareVersion
	}
	if o.FrozenStatus != nil {
		toSerialize["FrozenStatus"] = o.FrozenStatus
	}
	if o.HealthState != nil {
		toSerialize["HealthState"] = o.HealthState
	}
	if o.LockStatus != nil {
		toSerialize["LockStatus"] = o.LockStatus
	}
	if o.MemoryCapacity != nil {
		toSerialize["MemoryCapacity"] = o.MemoryCapacity
	}
	if o.MemoryId != nil {
		toSerialize["MemoryId"] = o.MemoryId
	}
	if o.PersistentMemoryCapacity != nil {
		toSerialize["PersistentMemoryCapacity"] = o.PersistentMemoryCapacity
	}
	if o.ReservedCapacity != nil {
		toSerialize["ReservedCapacity"] = o.ReservedCapacity
	}
	if o.SecurityStatus != nil {
		toSerialize["SecurityStatus"] = o.SecurityStatus
	}
	if o.SocketId != nil {
		toSerialize["SocketId"] = o.SocketId
	}
	if o.SocketMemoryId != nil {
		toSerialize["SocketMemoryId"] = o.SocketMemoryId
	}
	if o.TotalCapacity != nil {
		toSerialize["TotalCapacity"] = o.TotalCapacity
	}
	if o.Uid != nil {
		toSerialize["Uid"] = o.Uid
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.MemoryArray != nil {
		toSerialize["MemoryArray"] = o.MemoryArray
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableMemoryPersistentMemoryUnit struct {
	value *MemoryPersistentMemoryUnit
	isSet bool
}

func (v NullableMemoryPersistentMemoryUnit) Get() *MemoryPersistentMemoryUnit {
	return v.value
}

func (v *NullableMemoryPersistentMemoryUnit) Set(val *MemoryPersistentMemoryUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryUnit(val *MemoryPersistentMemoryUnit) *NullableMemoryPersistentMemoryUnit {
	return &NullableMemoryPersistentMemoryUnit{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
