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

// HyperflexHxapClusterAllOf Definition of the list of properties defined in 'hyperflex.HxapCluster', excluding properties defined in parent classes.
type HyperflexHxapClusterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// CPU oversubscription factor configured on the cluster.
	ConfiguredCpuOverSubFactor *float64                            `json:"ConfiguredCpuOverSubFactor,omitempty" yaml:"ConfiguredCpuOverSubFactor,omitempty"`
	CpuAllocation              NullableVirtualizationCpuAllocation `json:"CpuAllocation,omitempty" yaml:"CpuAllocation,omitempty"`
	// Current oversubscription factor of the cluster.
	CurrentCpuOverSubFactor *float64 `json:"CurrentCpuOverSubFactor,omitempty" yaml:"CurrentCpuOverSubFactor,omitempty"`
	// Datacenter to which the cluster belongs.
	DatacenterName *string `json:"DatacenterName,omitempty" yaml:"DatacenterName,omitempty"`
	// Reason for the failure when cluster is in failed state.
	FailureReason *string `json:"FailureReason,omitempty" yaml:"FailureReason,omitempty"`
	// Hypervisor version of HyperFlex compute cluster along with build number.
	HypervisorBuild *string `json:"HypervisorBuild,omitempty" yaml:"HypervisorBuild,omitempty"`
	// Management IP Address of the cluster.
	ManagementIpAddress *string                                `json:"ManagementIpAddress,omitempty" yaml:"ManagementIpAddress,omitempty"`
	MemoryAllocation    NullableVirtualizationMemoryAllocation `json:"MemoryAllocation,omitempty" yaml:"MemoryAllocation,omitempty"`
	HxCluster           *HyperflexClusterRelationship          `json:"HxCluster,omitempty" yaml:"HxCluster,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship   `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewHyperflexHxapClusterAllOf instantiates a new HyperflexHxapClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapClusterAllOf(classId string, objectType string) *HyperflexHxapClusterAllOf {
	this := HyperflexHxapClusterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxapClusterAllOfWithDefaults instantiates a new HyperflexHxapClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapClusterAllOfWithDefaults() *HyperflexHxapClusterAllOf {
	this := HyperflexHxapClusterAllOf{}
	var classId string = "hyperflex.HxapCluster"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapCluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapClusterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapClusterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapClusterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapClusterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapClusterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapClusterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfiguredCpuOverSubFactor returns the ConfiguredCpuOverSubFactor field value if set, zero value otherwise.
func (o *HyperflexHxapClusterAllOf) GetConfiguredCpuOverSubFactor() float64 {
	if o == nil || o.ConfiguredCpuOverSubFactor == nil {
		var ret float64
		return ret
	}
	return *o.ConfiguredCpuOverSubFactor
}

// GetConfiguredCpuOverSubFactorOk returns a tuple with the ConfiguredCpuOverSubFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapClusterAllOf) GetConfiguredCpuOverSubFactorOk() (*float64, bool) {
	if o == nil || o.ConfiguredCpuOverSubFactor == nil {
		return nil, false
	}
	return o.ConfiguredCpuOverSubFactor, true
}

// HasConfiguredCpuOverSubFactor returns a boolean if a field has been set.
func (o *HyperflexHxapClusterAllOf) HasConfiguredCpuOverSubFactor() bool {
	if o != nil && o.ConfiguredCpuOverSubFactor != nil {
		return true
	}

	return false
}

// SetConfiguredCpuOverSubFactor gets a reference to the given float64 and assigns it to the ConfiguredCpuOverSubFactor field.
func (o *HyperflexHxapClusterAllOf) SetConfiguredCpuOverSubFactor(v float64) {
	o.ConfiguredCpuOverSubFactor = &v
}

// GetCpuAllocation returns the CpuAllocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapClusterAllOf) GetCpuAllocation() VirtualizationCpuAllocation {
	if o == nil || o.CpuAllocation.Get() == nil {
		var ret VirtualizationCpuAllocation
		return ret
	}
	return *o.CpuAllocation.Get()
}

// GetCpuAllocationOk returns a tuple with the CpuAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapClusterAllOf) GetCpuAllocationOk() (*VirtualizationCpuAllocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuAllocation.Get(), o.CpuAllocation.IsSet()
}

// HasCpuAllocation returns a boolean if a field has been set.
func (o *HyperflexHxapClusterAllOf) HasCpuAllocation() bool {
	if o != nil && o.CpuAllocation.IsSet() {
		return true
	}

	return false
}

// SetCpuAllocation gets a reference to the given NullableVirtualizationCpuAllocation and assigns it to the CpuAllocation field.
func (o *HyperflexHxapClusterAllOf) SetCpuAllocation(v VirtualizationCpuAllocation) {
	o.CpuAllocation.Set(&v)
}

// SetCpuAllocationNil sets the value for CpuAllocation to be an explicit nil
func (o *HyperflexHxapClusterAllOf) SetCpuAllocationNil() {
	o.CpuAllocation.Set(nil)
}

// UnsetCpuAllocation ensures that no value is present for CpuAllocation, not even an explicit nil
func (o *HyperflexHxapClusterAllOf) UnsetCpuAllocation() {
	o.CpuAllocation.Unset()
}

// GetCurrentCpuOverSubFactor returns the CurrentCpuOverSubFactor field value if set, zero value otherwise.
func (o *HyperflexHxapClusterAllOf) GetCurrentCpuOverSubFactor() float64 {
	if o == nil || o.CurrentCpuOverSubFactor == nil {
		var ret float64
		return ret
	}
	return *o.CurrentCpuOverSubFactor
}

// GetCurrentCpuOverSubFactorOk returns a tuple with the CurrentCpuOverSubFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapClusterAllOf) GetCurrentCpuOverSubFactorOk() (*float64, bool) {
	if o == nil || o.CurrentCpuOverSubFactor == nil {
		return nil, false
	}
	return o.CurrentCpuOverSubFactor, true
}

// HasCurrentCpuOverSubFactor returns a boolean if a field has been set.
func (o *HyperflexHxapClusterAllOf) HasCurrentCpuOverSubFactor() bool {
	if o != nil && o.CurrentCpuOverSubFactor != nil {
		return true
	}

	return false
}

// SetCurrentCpuOverSubFactor gets a reference to the given float64 and assigns it to the CurrentCpuOverSubFactor field.
func (o *HyperflexHxapClusterAllOf) SetCurrentCpuOverSubFactor(v float64) {
	o.CurrentCpuOverSubFactor = &v
}

// GetDatacenterName returns the DatacenterName field value if set, zero value otherwise.
func (o *HyperflexHxapClusterAllOf) GetDatacenterName() string {
	if o == nil || o.DatacenterName == nil {
		var ret string
		return ret
	}
	return *o.DatacenterName
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapClusterAllOf) GetDatacenterNameOk() (*string, bool) {
	if o == nil || o.DatacenterName == nil {
		return nil, false
	}
	return o.DatacenterName, true
}

// HasDatacenterName returns a boolean if a field has been set.
func (o *HyperflexHxapClusterAllOf) HasDatacenterName() bool {
	if o != nil && o.DatacenterName != nil {
		return true
	}

	return false
}

// SetDatacenterName gets a reference to the given string and assigns it to the DatacenterName field.
func (o *HyperflexHxapClusterAllOf) SetDatacenterName(v string) {
	o.DatacenterName = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *HyperflexHxapClusterAllOf) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapClusterAllOf) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *HyperflexHxapClusterAllOf) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *HyperflexHxapClusterAllOf) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetHypervisorBuild returns the HypervisorBuild field value if set, zero value otherwise.
func (o *HyperflexHxapClusterAllOf) GetHypervisorBuild() string {
	if o == nil || o.HypervisorBuild == nil {
		var ret string
		return ret
	}
	return *o.HypervisorBuild
}

// GetHypervisorBuildOk returns a tuple with the HypervisorBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapClusterAllOf) GetHypervisorBuildOk() (*string, bool) {
	if o == nil || o.HypervisorBuild == nil {
		return nil, false
	}
	return o.HypervisorBuild, true
}

// HasHypervisorBuild returns a boolean if a field has been set.
func (o *HyperflexHxapClusterAllOf) HasHypervisorBuild() bool {
	if o != nil && o.HypervisorBuild != nil {
		return true
	}

	return false
}

// SetHypervisorBuild gets a reference to the given string and assigns it to the HypervisorBuild field.
func (o *HyperflexHxapClusterAllOf) SetHypervisorBuild(v string) {
	o.HypervisorBuild = &v
}

// GetManagementIpAddress returns the ManagementIpAddress field value if set, zero value otherwise.
func (o *HyperflexHxapClusterAllOf) GetManagementIpAddress() string {
	if o == nil || o.ManagementIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementIpAddress
}

// GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapClusterAllOf) GetManagementIpAddressOk() (*string, bool) {
	if o == nil || o.ManagementIpAddress == nil {
		return nil, false
	}
	return o.ManagementIpAddress, true
}

// HasManagementIpAddress returns a boolean if a field has been set.
func (o *HyperflexHxapClusterAllOf) HasManagementIpAddress() bool {
	if o != nil && o.ManagementIpAddress != nil {
		return true
	}

	return false
}

// SetManagementIpAddress gets a reference to the given string and assigns it to the ManagementIpAddress field.
func (o *HyperflexHxapClusterAllOf) SetManagementIpAddress(v string) {
	o.ManagementIpAddress = &v
}

// GetMemoryAllocation returns the MemoryAllocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapClusterAllOf) GetMemoryAllocation() VirtualizationMemoryAllocation {
	if o == nil || o.MemoryAllocation.Get() == nil {
		var ret VirtualizationMemoryAllocation
		return ret
	}
	return *o.MemoryAllocation.Get()
}

// GetMemoryAllocationOk returns a tuple with the MemoryAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapClusterAllOf) GetMemoryAllocationOk() (*VirtualizationMemoryAllocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryAllocation.Get(), o.MemoryAllocation.IsSet()
}

// HasMemoryAllocation returns a boolean if a field has been set.
func (o *HyperflexHxapClusterAllOf) HasMemoryAllocation() bool {
	if o != nil && o.MemoryAllocation.IsSet() {
		return true
	}

	return false
}

// SetMemoryAllocation gets a reference to the given NullableVirtualizationMemoryAllocation and assigns it to the MemoryAllocation field.
func (o *HyperflexHxapClusterAllOf) SetMemoryAllocation(v VirtualizationMemoryAllocation) {
	o.MemoryAllocation.Set(&v)
}

// SetMemoryAllocationNil sets the value for MemoryAllocation to be an explicit nil
func (o *HyperflexHxapClusterAllOf) SetMemoryAllocationNil() {
	o.MemoryAllocation.Set(nil)
}

// UnsetMemoryAllocation ensures that no value is present for MemoryAllocation, not even an explicit nil
func (o *HyperflexHxapClusterAllOf) UnsetMemoryAllocation() {
	o.MemoryAllocation.Unset()
}

// GetHxCluster returns the HxCluster field value if set, zero value otherwise.
func (o *HyperflexHxapClusterAllOf) GetHxCluster() HyperflexClusterRelationship {
	if o == nil || o.HxCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.HxCluster
}

// GetHxClusterOk returns a tuple with the HxCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapClusterAllOf) GetHxClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.HxCluster == nil {
		return nil, false
	}
	return o.HxCluster, true
}

// HasHxCluster returns a boolean if a field has been set.
func (o *HyperflexHxapClusterAllOf) HasHxCluster() bool {
	if o != nil && o.HxCluster != nil {
		return true
	}

	return false
}

// SetHxCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the HxCluster field.
func (o *HyperflexHxapClusterAllOf) SetHxCluster(v HyperflexClusterRelationship) {
	o.HxCluster = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *HyperflexHxapClusterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapClusterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexHxapClusterAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexHxapClusterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o HyperflexHxapClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfiguredCpuOverSubFactor != nil {
		toSerialize["ConfiguredCpuOverSubFactor"] = o.ConfiguredCpuOverSubFactor
	}
	if o.CpuAllocation.IsSet() {
		toSerialize["CpuAllocation"] = o.CpuAllocation.Get()
	}
	if o.CurrentCpuOverSubFactor != nil {
		toSerialize["CurrentCpuOverSubFactor"] = o.CurrentCpuOverSubFactor
	}
	if o.DatacenterName != nil {
		toSerialize["DatacenterName"] = o.DatacenterName
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.HypervisorBuild != nil {
		toSerialize["HypervisorBuild"] = o.HypervisorBuild
	}
	if o.ManagementIpAddress != nil {
		toSerialize["ManagementIpAddress"] = o.ManagementIpAddress
	}
	if o.MemoryAllocation.IsSet() {
		toSerialize["MemoryAllocation"] = o.MemoryAllocation.Get()
	}
	if o.HxCluster != nil {
		toSerialize["HxCluster"] = o.HxCluster
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexHxapClusterAllOf struct {
	value *HyperflexHxapClusterAllOf
	isSet bool
}

func (v NullableHyperflexHxapClusterAllOf) Get() *HyperflexHxapClusterAllOf {
	return v.value
}

func (v *NullableHyperflexHxapClusterAllOf) Set(val *HyperflexHxapClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapClusterAllOf(val *HyperflexHxapClusterAllOf) *NullableHyperflexHxapClusterAllOf {
	return &NullableHyperflexHxapClusterAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxapClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
