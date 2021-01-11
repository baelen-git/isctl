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

// HyperflexClusterNetworkPolicy A policy specifying VLANs for management, VM migration, and VM traffic.
type HyperflexClusterNetworkPolicy struct {
	PolicyAbstractPolicy `yaml:"PolicyAbstractPolicy,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Enable or disable jumbo frames.
	JumboFrame     *bool                               `json:"JumboFrame,omitempty" yaml:"JumboFrame,omitempty"`
	KvmIpRange     NullableHyperflexIpAddrRange        `json:"KvmIpRange,omitempty" yaml:"KvmIpRange,omitempty"`
	MacPrefixRange NullableHyperflexMacAddrPrefixRange `json:"MacPrefixRange,omitempty" yaml:"MacPrefixRange,omitempty"`
	MgmtVlan       NullableHyperflexNamedVlan          `json:"MgmtVlan,omitempty" yaml:"MgmtVlan,omitempty"`
	// Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be '1G' or '10G+'. Use '10G+' for link speeds of 10G or above. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be 'default' only. * `default` - Current default value set on the hardware platform. * `1G` - A link speed of 1 gigabit per second. * `10G` - A link speed of 10 gigabits per second or above.
	UplinkSpeed     *string                    `json:"UplinkSpeed,omitempty" yaml:"UplinkSpeed,omitempty"`
	VmMigrationVlan NullableHyperflexNamedVlan `json:"VmMigrationVlan,omitempty" yaml:"VmMigrationVlan,omitempty"`
	VmNetworkVlans  []HyperflexNamedVlan       `json:"VmNetworkVlans,omitempty" yaml:"VmNetworkVlans,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty" yaml:"ClusterProfiles,omitempty"`
	Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
}

// NewHyperflexClusterNetworkPolicy instantiates a new HyperflexClusterNetworkPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterNetworkPolicy(classId string, objectType string) *HyperflexClusterNetworkPolicy {
	this := HyperflexClusterNetworkPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var uplinkSpeed string = "default"
	this.UplinkSpeed = &uplinkSpeed
	return &this
}

// NewHyperflexClusterNetworkPolicyWithDefaults instantiates a new HyperflexClusterNetworkPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterNetworkPolicyWithDefaults() *HyperflexClusterNetworkPolicy {
	this := HyperflexClusterNetworkPolicy{}
	var classId string = "hyperflex.ClusterNetworkPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterNetworkPolicy"
	this.ObjectType = objectType
	var uplinkSpeed string = "default"
	this.UplinkSpeed = &uplinkSpeed
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterNetworkPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterNetworkPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterNetworkPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterNetworkPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterNetworkPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterNetworkPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetJumboFrame returns the JumboFrame field value if set, zero value otherwise.
func (o *HyperflexClusterNetworkPolicy) GetJumboFrame() bool {
	if o == nil || o.JumboFrame == nil {
		var ret bool
		return ret
	}
	return *o.JumboFrame
}

// GetJumboFrameOk returns a tuple with the JumboFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterNetworkPolicy) GetJumboFrameOk() (*bool, bool) {
	if o == nil || o.JumboFrame == nil {
		return nil, false
	}
	return o.JumboFrame, true
}

// HasJumboFrame returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasJumboFrame() bool {
	if o != nil && o.JumboFrame != nil {
		return true
	}

	return false
}

// SetJumboFrame gets a reference to the given bool and assigns it to the JumboFrame field.
func (o *HyperflexClusterNetworkPolicy) SetJumboFrame(v bool) {
	o.JumboFrame = &v
}

// GetKvmIpRange returns the KvmIpRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetKvmIpRange() HyperflexIpAddrRange {
	if o == nil || o.KvmIpRange.Get() == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.KvmIpRange.Get()
}

// GetKvmIpRangeOk returns a tuple with the KvmIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetKvmIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.KvmIpRange.Get(), o.KvmIpRange.IsSet()
}

// HasKvmIpRange returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasKvmIpRange() bool {
	if o != nil && o.KvmIpRange.IsSet() {
		return true
	}

	return false
}

// SetKvmIpRange gets a reference to the given NullableHyperflexIpAddrRange and assigns it to the KvmIpRange field.
func (o *HyperflexClusterNetworkPolicy) SetKvmIpRange(v HyperflexIpAddrRange) {
	o.KvmIpRange.Set(&v)
}

// SetKvmIpRangeNil sets the value for KvmIpRange to be an explicit nil
func (o *HyperflexClusterNetworkPolicy) SetKvmIpRangeNil() {
	o.KvmIpRange.Set(nil)
}

// UnsetKvmIpRange ensures that no value is present for KvmIpRange, not even an explicit nil
func (o *HyperflexClusterNetworkPolicy) UnsetKvmIpRange() {
	o.KvmIpRange.Unset()
}

// GetMacPrefixRange returns the MacPrefixRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetMacPrefixRange() HyperflexMacAddrPrefixRange {
	if o == nil || o.MacPrefixRange.Get() == nil {
		var ret HyperflexMacAddrPrefixRange
		return ret
	}
	return *o.MacPrefixRange.Get()
}

// GetMacPrefixRangeOk returns a tuple with the MacPrefixRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetMacPrefixRangeOk() (*HyperflexMacAddrPrefixRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacPrefixRange.Get(), o.MacPrefixRange.IsSet()
}

// HasMacPrefixRange returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasMacPrefixRange() bool {
	if o != nil && o.MacPrefixRange.IsSet() {
		return true
	}

	return false
}

// SetMacPrefixRange gets a reference to the given NullableHyperflexMacAddrPrefixRange and assigns it to the MacPrefixRange field.
func (o *HyperflexClusterNetworkPolicy) SetMacPrefixRange(v HyperflexMacAddrPrefixRange) {
	o.MacPrefixRange.Set(&v)
}

// SetMacPrefixRangeNil sets the value for MacPrefixRange to be an explicit nil
func (o *HyperflexClusterNetworkPolicy) SetMacPrefixRangeNil() {
	o.MacPrefixRange.Set(nil)
}

// UnsetMacPrefixRange ensures that no value is present for MacPrefixRange, not even an explicit nil
func (o *HyperflexClusterNetworkPolicy) UnsetMacPrefixRange() {
	o.MacPrefixRange.Unset()
}

// GetMgmtVlan returns the MgmtVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetMgmtVlan() HyperflexNamedVlan {
	if o == nil || o.MgmtVlan.Get() == nil {
		var ret HyperflexNamedVlan
		return ret
	}
	return *o.MgmtVlan.Get()
}

// GetMgmtVlanOk returns a tuple with the MgmtVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetMgmtVlanOk() (*HyperflexNamedVlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.MgmtVlan.Get(), o.MgmtVlan.IsSet()
}

// HasMgmtVlan returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasMgmtVlan() bool {
	if o != nil && o.MgmtVlan.IsSet() {
		return true
	}

	return false
}

// SetMgmtVlan gets a reference to the given NullableHyperflexNamedVlan and assigns it to the MgmtVlan field.
func (o *HyperflexClusterNetworkPolicy) SetMgmtVlan(v HyperflexNamedVlan) {
	o.MgmtVlan.Set(&v)
}

// SetMgmtVlanNil sets the value for MgmtVlan to be an explicit nil
func (o *HyperflexClusterNetworkPolicy) SetMgmtVlanNil() {
	o.MgmtVlan.Set(nil)
}

// UnsetMgmtVlan ensures that no value is present for MgmtVlan, not even an explicit nil
func (o *HyperflexClusterNetworkPolicy) UnsetMgmtVlan() {
	o.MgmtVlan.Unset()
}

// GetUplinkSpeed returns the UplinkSpeed field value if set, zero value otherwise.
func (o *HyperflexClusterNetworkPolicy) GetUplinkSpeed() string {
	if o == nil || o.UplinkSpeed == nil {
		var ret string
		return ret
	}
	return *o.UplinkSpeed
}

// GetUplinkSpeedOk returns a tuple with the UplinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterNetworkPolicy) GetUplinkSpeedOk() (*string, bool) {
	if o == nil || o.UplinkSpeed == nil {
		return nil, false
	}
	return o.UplinkSpeed, true
}

// HasUplinkSpeed returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasUplinkSpeed() bool {
	if o != nil && o.UplinkSpeed != nil {
		return true
	}

	return false
}

// SetUplinkSpeed gets a reference to the given string and assigns it to the UplinkSpeed field.
func (o *HyperflexClusterNetworkPolicy) SetUplinkSpeed(v string) {
	o.UplinkSpeed = &v
}

// GetVmMigrationVlan returns the VmMigrationVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetVmMigrationVlan() HyperflexNamedVlan {
	if o == nil || o.VmMigrationVlan.Get() == nil {
		var ret HyperflexNamedVlan
		return ret
	}
	return *o.VmMigrationVlan.Get()
}

// GetVmMigrationVlanOk returns a tuple with the VmMigrationVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetVmMigrationVlanOk() (*HyperflexNamedVlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmMigrationVlan.Get(), o.VmMigrationVlan.IsSet()
}

// HasVmMigrationVlan returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasVmMigrationVlan() bool {
	if o != nil && o.VmMigrationVlan.IsSet() {
		return true
	}

	return false
}

// SetVmMigrationVlan gets a reference to the given NullableHyperflexNamedVlan and assigns it to the VmMigrationVlan field.
func (o *HyperflexClusterNetworkPolicy) SetVmMigrationVlan(v HyperflexNamedVlan) {
	o.VmMigrationVlan.Set(&v)
}

// SetVmMigrationVlanNil sets the value for VmMigrationVlan to be an explicit nil
func (o *HyperflexClusterNetworkPolicy) SetVmMigrationVlanNil() {
	o.VmMigrationVlan.Set(nil)
}

// UnsetVmMigrationVlan ensures that no value is present for VmMigrationVlan, not even an explicit nil
func (o *HyperflexClusterNetworkPolicy) UnsetVmMigrationVlan() {
	o.VmMigrationVlan.Unset()
}

// GetVmNetworkVlans returns the VmNetworkVlans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetVmNetworkVlans() []HyperflexNamedVlan {
	if o == nil {
		var ret []HyperflexNamedVlan
		return ret
	}
	return o.VmNetworkVlans
}

// GetVmNetworkVlansOk returns a tuple with the VmNetworkVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetVmNetworkVlansOk() (*[]HyperflexNamedVlan, bool) {
	if o == nil || o.VmNetworkVlans == nil {
		return nil, false
	}
	return &o.VmNetworkVlans, true
}

// HasVmNetworkVlans returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasVmNetworkVlans() bool {
	if o != nil && o.VmNetworkVlans != nil {
		return true
	}

	return false
}

// SetVmNetworkVlans gets a reference to the given []HyperflexNamedVlan and assigns it to the VmNetworkVlans field.
func (o *HyperflexClusterNetworkPolicy) SetVmNetworkVlans(v []HyperflexNamedVlan) {
	o.VmNetworkVlans = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterNetworkPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterNetworkPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexClusterNetworkPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexClusterNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexClusterNetworkPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexClusterNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexClusterNetworkPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.JumboFrame != nil {
		toSerialize["JumboFrame"] = o.JumboFrame
	}
	if o.KvmIpRange.IsSet() {
		toSerialize["KvmIpRange"] = o.KvmIpRange.Get()
	}
	if o.MacPrefixRange.IsSet() {
		toSerialize["MacPrefixRange"] = o.MacPrefixRange.Get()
	}
	if o.MgmtVlan.IsSet() {
		toSerialize["MgmtVlan"] = o.MgmtVlan.Get()
	}
	if o.UplinkSpeed != nil {
		toSerialize["UplinkSpeed"] = o.UplinkSpeed
	}
	if o.VmMigrationVlan.IsSet() {
		toSerialize["VmMigrationVlan"] = o.VmMigrationVlan.Get()
	}
	if o.VmNetworkVlans != nil {
		toSerialize["VmNetworkVlans"] = o.VmNetworkVlans
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexClusterNetworkPolicy struct {
	value *HyperflexClusterNetworkPolicy
	isSet bool
}

func (v NullableHyperflexClusterNetworkPolicy) Get() *HyperflexClusterNetworkPolicy {
	return v.value
}

func (v *NullableHyperflexClusterNetworkPolicy) Set(val *HyperflexClusterNetworkPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterNetworkPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterNetworkPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterNetworkPolicy(val *HyperflexClusterNetworkPolicy) *NullableHyperflexClusterNetworkPolicy {
	return &NullableHyperflexClusterNetworkPolicy{value: val, isSet: true}
}

func (v NullableHyperflexClusterNetworkPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterNetworkPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
