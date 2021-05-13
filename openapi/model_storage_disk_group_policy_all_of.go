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

// StorageDiskGroupPolicyAllOf Definition of the list of properties defined in 'storage.DiskGroupPolicy', excluding properties defined in parent classes.
type StorageDiskGroupPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType         string             `json:"ObjectType" yaml:"ObjectType"`
	DedicatedHotSpares []StorageLocalDisk `json:"DedicatedHotSpares,omitempty" yaml:"DedicatedHotSpares,omitempty"`
	// The supported RAID level for the disk group. * `Raid0` - RAID 0 Stripe Raid Level. * `Raid1` - RAID 1 Mirror Raid Level. * `Raid5` - RAID 5 Mirror Raid Level. * `Raid6` - RAID 6 Mirror Raid Level. * `Raid10` - RAID 10 Mirror Raid Level. * `Raid50` - RAID 50 Mirror Raid Level. * `Raid60` - RAID 60 Mirror Raid Level.
	RaidLevel  *string            `json:"RaidLevel,omitempty" yaml:"RaidLevel,omitempty"`
	SpanGroups []StorageSpanGroup `json:"SpanGroups,omitempty" yaml:"SpanGroups,omitempty"`
	// Selected disks in the disk group in JBOD state will be set to Unconfigured Good state before they are used in virtual drive creation.
	UseJbods     *bool                                 `json:"UseJbods,omitempty" yaml:"UseJbods,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
	// An array of relationships to storageStoragePolicy resources.
	StoragePolicies []StorageStoragePolicyRelationship `json:"StoragePolicies,omitempty" yaml:"StoragePolicies,omitempty"`
}

// NewStorageDiskGroupPolicyAllOf instantiates a new StorageDiskGroupPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDiskGroupPolicyAllOf(classId string, objectType string) *StorageDiskGroupPolicyAllOf {
	this := StorageDiskGroupPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var raidLevel string = "Raid0"
	this.RaidLevel = &raidLevel
	return &this
}

// NewStorageDiskGroupPolicyAllOfWithDefaults instantiates a new StorageDiskGroupPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDiskGroupPolicyAllOfWithDefaults() *StorageDiskGroupPolicyAllOf {
	this := StorageDiskGroupPolicyAllOf{}
	var classId string = "storage.DiskGroupPolicy"
	this.ClassId = classId
	var objectType string = "storage.DiskGroupPolicy"
	this.ObjectType = objectType
	var raidLevel string = "Raid0"
	this.RaidLevel = &raidLevel
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageDiskGroupPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageDiskGroupPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageDiskGroupPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageDiskGroupPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDedicatedHotSpares returns the DedicatedHotSpares field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDiskGroupPolicyAllOf) GetDedicatedHotSpares() []StorageLocalDisk {
	if o == nil {
		var ret []StorageLocalDisk
		return ret
	}
	return o.DedicatedHotSpares
}

// GetDedicatedHotSparesOk returns a tuple with the DedicatedHotSpares field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDiskGroupPolicyAllOf) GetDedicatedHotSparesOk() (*[]StorageLocalDisk, bool) {
	if o == nil || o.DedicatedHotSpares == nil {
		return nil, false
	}
	return &o.DedicatedHotSpares, true
}

// HasDedicatedHotSpares returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicyAllOf) HasDedicatedHotSpares() bool {
	if o != nil && o.DedicatedHotSpares != nil {
		return true
	}

	return false
}

// SetDedicatedHotSpares gets a reference to the given []StorageLocalDisk and assigns it to the DedicatedHotSpares field.
func (o *StorageDiskGroupPolicyAllOf) SetDedicatedHotSpares(v []StorageLocalDisk) {
	o.DedicatedHotSpares = v
}

// GetRaidLevel returns the RaidLevel field value if set, zero value otherwise.
func (o *StorageDiskGroupPolicyAllOf) GetRaidLevel() string {
	if o == nil || o.RaidLevel == nil {
		var ret string
		return ret
	}
	return *o.RaidLevel
}

// GetRaidLevelOk returns a tuple with the RaidLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicyAllOf) GetRaidLevelOk() (*string, bool) {
	if o == nil || o.RaidLevel == nil {
		return nil, false
	}
	return o.RaidLevel, true
}

// HasRaidLevel returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicyAllOf) HasRaidLevel() bool {
	if o != nil && o.RaidLevel != nil {
		return true
	}

	return false
}

// SetRaidLevel gets a reference to the given string and assigns it to the RaidLevel field.
func (o *StorageDiskGroupPolicyAllOf) SetRaidLevel(v string) {
	o.RaidLevel = &v
}

// GetSpanGroups returns the SpanGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDiskGroupPolicyAllOf) GetSpanGroups() []StorageSpanGroup {
	if o == nil {
		var ret []StorageSpanGroup
		return ret
	}
	return o.SpanGroups
}

// GetSpanGroupsOk returns a tuple with the SpanGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDiskGroupPolicyAllOf) GetSpanGroupsOk() (*[]StorageSpanGroup, bool) {
	if o == nil || o.SpanGroups == nil {
		return nil, false
	}
	return &o.SpanGroups, true
}

// HasSpanGroups returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicyAllOf) HasSpanGroups() bool {
	if o != nil && o.SpanGroups != nil {
		return true
	}

	return false
}

// SetSpanGroups gets a reference to the given []StorageSpanGroup and assigns it to the SpanGroups field.
func (o *StorageDiskGroupPolicyAllOf) SetSpanGroups(v []StorageSpanGroup) {
	o.SpanGroups = v
}

// GetUseJbods returns the UseJbods field value if set, zero value otherwise.
func (o *StorageDiskGroupPolicyAllOf) GetUseJbods() bool {
	if o == nil || o.UseJbods == nil {
		var ret bool
		return ret
	}
	return *o.UseJbods
}

// GetUseJbodsOk returns a tuple with the UseJbods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicyAllOf) GetUseJbodsOk() (*bool, bool) {
	if o == nil || o.UseJbods == nil {
		return nil, false
	}
	return o.UseJbods, true
}

// HasUseJbods returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicyAllOf) HasUseJbods() bool {
	if o != nil && o.UseJbods != nil {
		return true
	}

	return false
}

// SetUseJbods gets a reference to the given bool and assigns it to the UseJbods field.
func (o *StorageDiskGroupPolicyAllOf) SetUseJbods(v bool) {
	o.UseJbods = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *StorageDiskGroupPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroupPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *StorageDiskGroupPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetStoragePolicies returns the StoragePolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDiskGroupPolicyAllOf) GetStoragePolicies() []StorageStoragePolicyRelationship {
	if o == nil {
		var ret []StorageStoragePolicyRelationship
		return ret
	}
	return o.StoragePolicies
}

// GetStoragePoliciesOk returns a tuple with the StoragePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDiskGroupPolicyAllOf) GetStoragePoliciesOk() (*[]StorageStoragePolicyRelationship, bool) {
	if o == nil || o.StoragePolicies == nil {
		return nil, false
	}
	return &o.StoragePolicies, true
}

// HasStoragePolicies returns a boolean if a field has been set.
func (o *StorageDiskGroupPolicyAllOf) HasStoragePolicies() bool {
	if o != nil && o.StoragePolicies != nil {
		return true
	}

	return false
}

// SetStoragePolicies gets a reference to the given []StorageStoragePolicyRelationship and assigns it to the StoragePolicies field.
func (o *StorageDiskGroupPolicyAllOf) SetStoragePolicies(v []StorageStoragePolicyRelationship) {
	o.StoragePolicies = v
}

func (o StorageDiskGroupPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DedicatedHotSpares != nil {
		toSerialize["DedicatedHotSpares"] = o.DedicatedHotSpares
	}
	if o.RaidLevel != nil {
		toSerialize["RaidLevel"] = o.RaidLevel
	}
	if o.SpanGroups != nil {
		toSerialize["SpanGroups"] = o.SpanGroups
	}
	if o.UseJbods != nil {
		toSerialize["UseJbods"] = o.UseJbods
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.StoragePolicies != nil {
		toSerialize["StoragePolicies"] = o.StoragePolicies
	}
	return json.Marshal(toSerialize)
}

type NullableStorageDiskGroupPolicyAllOf struct {
	value *StorageDiskGroupPolicyAllOf
	isSet bool
}

func (v NullableStorageDiskGroupPolicyAllOf) Get() *StorageDiskGroupPolicyAllOf {
	return v.value
}

func (v *NullableStorageDiskGroupPolicyAllOf) Set(val *StorageDiskGroupPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDiskGroupPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDiskGroupPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDiskGroupPolicyAllOf(val *StorageDiskGroupPolicyAllOf) *NullableStorageDiskGroupPolicyAllOf {
	return &NullableStorageDiskGroupPolicyAllOf{value: val, isSet: true}
}

func (v NullableStorageDiskGroupPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDiskGroupPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
