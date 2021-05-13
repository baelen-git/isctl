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

// VirtualizationVirtualDisk Depicts disk configuration used to be create a virtual disk on a hypervisor datastore.
type VirtualizationVirtualDisk struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Disk capacity to be provided with units example - 10Gi.
	Capacity *string `json:"Capacity,omitempty" yaml:"Capacity,omitempty"`
	// Flag to indicate whether the configuration is created from inventory object.
	Discovered *bool `json:"Discovered,omitempty" yaml:"Discovered,omitempty"`
	// File mode of the disk  example - Filesystem, Block. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
	Mode *string `json:"Mode,omitempty" yaml:"Mode,omitempty"`
	// Name of the storage disk. Name must be unique within a Datastore.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// Base64 encoded CA certificates of the https source to check against.
	SourceCerts *string `json:"SourceCerts,omitempty" yaml:"SourceCerts,omitempty"`
	// Source disk from which the content is copied.
	SourceDiskToClone *string `json:"SourceDiskToClone,omitempty" yaml:"SourceDiskToClone,omitempty"`
	// Image path used to import on the created disk.
	SourceFilePath   *string                                    `json:"SourceFilePath,omitempty" yaml:"SourceFilePath,omitempty"`
	Cluster          *VirtualizationBaseClusterRelationship     `json:"Cluster,omitempty" yaml:"Cluster,omitempty"`
	Inventory        *VirtualizationBaseVirtualDiskRelationship `json:"Inventory,omitempty" yaml:"Inventory,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship       `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
	WorkflowInfo     *WorkflowWorkflowInfoRelationship          `json:"WorkflowInfo,omitempty" yaml:"WorkflowInfo,omitempty"`
}

// NewVirtualizationVirtualDisk instantiates a new VirtualizationVirtualDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVirtualDisk(classId string, objectType string) *VirtualizationVirtualDisk {
	this := VirtualizationVirtualDisk{}
	this.ClassId = classId
	this.ObjectType = objectType
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// NewVirtualizationVirtualDiskWithDefaults instantiates a new VirtualizationVirtualDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVirtualDiskWithDefaults() *VirtualizationVirtualDisk {
	this := VirtualizationVirtualDisk{}
	var classId string = "virtualization.VirtualDisk"
	this.ClassId = classId
	var objectType string = "virtualization.VirtualDisk"
	this.ObjectType = objectType
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVirtualDisk) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVirtualDisk) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVirtualDisk) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVirtualDisk) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *VirtualizationVirtualDisk) SetCapacity(v string) {
	o.Capacity = &v
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetDiscovered() bool {
	if o == nil || o.Discovered == nil {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetDiscoveredOk() (*bool, bool) {
	if o == nil || o.Discovered == nil {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasDiscovered() bool {
	if o != nil && o.Discovered != nil {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *VirtualizationVirtualDisk) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VirtualizationVirtualDisk) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVirtualDisk) SetName(v string) {
	o.Name = &v
}

// GetSourceCerts returns the SourceCerts field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetSourceCerts() string {
	if o == nil || o.SourceCerts == nil {
		var ret string
		return ret
	}
	return *o.SourceCerts
}

// GetSourceCertsOk returns a tuple with the SourceCerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetSourceCertsOk() (*string, bool) {
	if o == nil || o.SourceCerts == nil {
		return nil, false
	}
	return o.SourceCerts, true
}

// HasSourceCerts returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasSourceCerts() bool {
	if o != nil && o.SourceCerts != nil {
		return true
	}

	return false
}

// SetSourceCerts gets a reference to the given string and assigns it to the SourceCerts field.
func (o *VirtualizationVirtualDisk) SetSourceCerts(v string) {
	o.SourceCerts = &v
}

// GetSourceDiskToClone returns the SourceDiskToClone field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetSourceDiskToClone() string {
	if o == nil || o.SourceDiskToClone == nil {
		var ret string
		return ret
	}
	return *o.SourceDiskToClone
}

// GetSourceDiskToCloneOk returns a tuple with the SourceDiskToClone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetSourceDiskToCloneOk() (*string, bool) {
	if o == nil || o.SourceDiskToClone == nil {
		return nil, false
	}
	return o.SourceDiskToClone, true
}

// HasSourceDiskToClone returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasSourceDiskToClone() bool {
	if o != nil && o.SourceDiskToClone != nil {
		return true
	}

	return false
}

// SetSourceDiskToClone gets a reference to the given string and assigns it to the SourceDiskToClone field.
func (o *VirtualizationVirtualDisk) SetSourceDiskToClone(v string) {
	o.SourceDiskToClone = &v
}

// GetSourceFilePath returns the SourceFilePath field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetSourceFilePath() string {
	if o == nil || o.SourceFilePath == nil {
		var ret string
		return ret
	}
	return *o.SourceFilePath
}

// GetSourceFilePathOk returns a tuple with the SourceFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetSourceFilePathOk() (*string, bool) {
	if o == nil || o.SourceFilePath == nil {
		return nil, false
	}
	return o.SourceFilePath, true
}

// HasSourceFilePath returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasSourceFilePath() bool {
	if o != nil && o.SourceFilePath != nil {
		return true
	}

	return false
}

// SetSourceFilePath gets a reference to the given string and assigns it to the SourceFilePath field.
func (o *VirtualizationVirtualDisk) SetSourceFilePath(v string) {
	o.SourceFilePath = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetCluster() VirtualizationBaseClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationBaseClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationBaseClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationVirtualDisk) SetCluster(v VirtualizationBaseClusterRelationship) {
	o.Cluster = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetInventory() VirtualizationBaseVirtualDiskRelationship {
	if o == nil || o.Inventory == nil {
		var ret VirtualizationBaseVirtualDiskRelationship
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetInventoryOk() (*VirtualizationBaseVirtualDiskRelationship, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given VirtualizationBaseVirtualDiskRelationship and assigns it to the Inventory field.
func (o *VirtualizationVirtualDisk) SetInventory(v VirtualizationBaseVirtualDiskRelationship) {
	o.Inventory = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationVirtualDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *VirtualizationVirtualDisk) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDisk) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *VirtualizationVirtualDisk) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *VirtualizationVirtualDisk) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o VirtualizationVirtualDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Discovered != nil {
		toSerialize["Discovered"] = o.Discovered
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SourceCerts != nil {
		toSerialize["SourceCerts"] = o.SourceCerts
	}
	if o.SourceDiskToClone != nil {
		toSerialize["SourceDiskToClone"] = o.SourceDiskToClone
	}
	if o.SourceFilePath != nil {
		toSerialize["SourceFilePath"] = o.SourceFilePath
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Inventory != nil {
		toSerialize["Inventory"] = o.Inventory
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationVirtualDisk struct {
	value *VirtualizationVirtualDisk
	isSet bool
}

func (v NullableVirtualizationVirtualDisk) Get() *VirtualizationVirtualDisk {
	return v.value
}

func (v *NullableVirtualizationVirtualDisk) Set(val *VirtualizationVirtualDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVirtualDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVirtualDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVirtualDisk(val *VirtualizationVirtualDisk) *NullableVirtualizationVirtualDisk {
	return &NullableVirtualizationVirtualDisk{value: val, isSet: true}
}

func (v NullableVirtualizationVirtualDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVirtualDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
