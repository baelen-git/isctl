/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StorageStoragePolicyAllOf Definition of the list of properties defined in 'storage.StoragePolicy', excluding properties defined in parent classes.
type StorageStoragePolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string             `json:"ObjectType"`
	GlobalHotSpares []StorageLocalDisk `json:"GlobalHotSpares,omitempty"`
	// Retains the virtual drives defined in policy if they exist already. If this flag is false, the existing virtual drives are removed and created again based on virtual drives in the policy.
	RetainPolicyVirtualDrives *bool `json:"RetainPolicyVirtualDrives,omitempty"`
	// Unused Disks State is used to specify the state, unconfigured good or jbod, in which the disks that are not used in this policy should be moved. * `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group. * `Jbod` - JBOD state where the disks start showing up to host os.
	UnusedDisksState *string                     `json:"UnusedDisksState,omitempty"`
	VirtualDrives    []StorageVirtualDriveConfig `json:"VirtualDrives,omitempty"`
	// An array of relationships to storageDiskGroupPolicy resources.
	DiskGroupPolicies []StorageDiskGroupPolicyRelationship  `json:"DiskGroupPolicies,omitempty"`
	Organization      *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageStoragePolicyAllOf StorageStoragePolicyAllOf

// NewStorageStoragePolicyAllOf instantiates a new StorageStoragePolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageStoragePolicyAllOf(classId string, objectType string) *StorageStoragePolicyAllOf {
	this := StorageStoragePolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var retainPolicyVirtualDrives bool = true
	this.RetainPolicyVirtualDrives = &retainPolicyVirtualDrives
	var unusedDisksState string = "UnconfiguredGood"
	this.UnusedDisksState = &unusedDisksState
	return &this
}

// NewStorageStoragePolicyAllOfWithDefaults instantiates a new StorageStoragePolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageStoragePolicyAllOfWithDefaults() *StorageStoragePolicyAllOf {
	this := StorageStoragePolicyAllOf{}
	var classId string = "storage.StoragePolicy"
	this.ClassId = classId
	var objectType string = "storage.StoragePolicy"
	this.ObjectType = objectType
	var retainPolicyVirtualDrives bool = true
	this.RetainPolicyVirtualDrives = &retainPolicyVirtualDrives
	var unusedDisksState string = "UnconfiguredGood"
	this.UnusedDisksState = &unusedDisksState
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageStoragePolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageStoragePolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageStoragePolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageStoragePolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGlobalHotSpares returns the GlobalHotSpares field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicyAllOf) GetGlobalHotSpares() []StorageLocalDisk {
	if o == nil {
		var ret []StorageLocalDisk
		return ret
	}
	return o.GlobalHotSpares
}

// GetGlobalHotSparesOk returns a tuple with the GlobalHotSpares field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicyAllOf) GetGlobalHotSparesOk() (*[]StorageLocalDisk, bool) {
	if o == nil || o.GlobalHotSpares == nil {
		return nil, false
	}
	return &o.GlobalHotSpares, true
}

// HasGlobalHotSpares returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasGlobalHotSpares() bool {
	if o != nil && o.GlobalHotSpares != nil {
		return true
	}

	return false
}

// SetGlobalHotSpares gets a reference to the given []StorageLocalDisk and assigns it to the GlobalHotSpares field.
func (o *StorageStoragePolicyAllOf) SetGlobalHotSpares(v []StorageLocalDisk) {
	o.GlobalHotSpares = v
}

// GetRetainPolicyVirtualDrives returns the RetainPolicyVirtualDrives field value if set, zero value otherwise.
func (o *StorageStoragePolicyAllOf) GetRetainPolicyVirtualDrives() bool {
	if o == nil || o.RetainPolicyVirtualDrives == nil {
		var ret bool
		return ret
	}
	return *o.RetainPolicyVirtualDrives
}

// GetRetainPolicyVirtualDrivesOk returns a tuple with the RetainPolicyVirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetRetainPolicyVirtualDrivesOk() (*bool, bool) {
	if o == nil || o.RetainPolicyVirtualDrives == nil {
		return nil, false
	}
	return o.RetainPolicyVirtualDrives, true
}

// HasRetainPolicyVirtualDrives returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasRetainPolicyVirtualDrives() bool {
	if o != nil && o.RetainPolicyVirtualDrives != nil {
		return true
	}

	return false
}

// SetRetainPolicyVirtualDrives gets a reference to the given bool and assigns it to the RetainPolicyVirtualDrives field.
func (o *StorageStoragePolicyAllOf) SetRetainPolicyVirtualDrives(v bool) {
	o.RetainPolicyVirtualDrives = &v
}

// GetUnusedDisksState returns the UnusedDisksState field value if set, zero value otherwise.
func (o *StorageStoragePolicyAllOf) GetUnusedDisksState() string {
	if o == nil || o.UnusedDisksState == nil {
		var ret string
		return ret
	}
	return *o.UnusedDisksState
}

// GetUnusedDisksStateOk returns a tuple with the UnusedDisksState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetUnusedDisksStateOk() (*string, bool) {
	if o == nil || o.UnusedDisksState == nil {
		return nil, false
	}
	return o.UnusedDisksState, true
}

// HasUnusedDisksState returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasUnusedDisksState() bool {
	if o != nil && o.UnusedDisksState != nil {
		return true
	}

	return false
}

// SetUnusedDisksState gets a reference to the given string and assigns it to the UnusedDisksState field.
func (o *StorageStoragePolicyAllOf) SetUnusedDisksState(v string) {
	o.UnusedDisksState = &v
}

// GetVirtualDrives returns the VirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicyAllOf) GetVirtualDrives() []StorageVirtualDriveConfig {
	if o == nil {
		var ret []StorageVirtualDriveConfig
		return ret
	}
	return o.VirtualDrives
}

// GetVirtualDrivesOk returns a tuple with the VirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicyAllOf) GetVirtualDrivesOk() (*[]StorageVirtualDriveConfig, bool) {
	if o == nil || o.VirtualDrives == nil {
		return nil, false
	}
	return &o.VirtualDrives, true
}

// HasVirtualDrives returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasVirtualDrives() bool {
	if o != nil && o.VirtualDrives != nil {
		return true
	}

	return false
}

// SetVirtualDrives gets a reference to the given []StorageVirtualDriveConfig and assigns it to the VirtualDrives field.
func (o *StorageStoragePolicyAllOf) SetVirtualDrives(v []StorageVirtualDriveConfig) {
	o.VirtualDrives = v
}

// GetDiskGroupPolicies returns the DiskGroupPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicyAllOf) GetDiskGroupPolicies() []StorageDiskGroupPolicyRelationship {
	if o == nil {
		var ret []StorageDiskGroupPolicyRelationship
		return ret
	}
	return o.DiskGroupPolicies
}

// GetDiskGroupPoliciesOk returns a tuple with the DiskGroupPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicyAllOf) GetDiskGroupPoliciesOk() (*[]StorageDiskGroupPolicyRelationship, bool) {
	if o == nil || o.DiskGroupPolicies == nil {
		return nil, false
	}
	return &o.DiskGroupPolicies, true
}

// HasDiskGroupPolicies returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasDiskGroupPolicies() bool {
	if o != nil && o.DiskGroupPolicies != nil {
		return true
	}

	return false
}

// SetDiskGroupPolicies gets a reference to the given []StorageDiskGroupPolicyRelationship and assigns it to the DiskGroupPolicies field.
func (o *StorageStoragePolicyAllOf) SetDiskGroupPolicies(v []StorageDiskGroupPolicyRelationship) {
	o.DiskGroupPolicies = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *StorageStoragePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *StorageStoragePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *StorageStoragePolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *StorageStoragePolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o StorageStoragePolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.GlobalHotSpares != nil {
		toSerialize["GlobalHotSpares"] = o.GlobalHotSpares
	}
	if o.RetainPolicyVirtualDrives != nil {
		toSerialize["RetainPolicyVirtualDrives"] = o.RetainPolicyVirtualDrives
	}
	if o.UnusedDisksState != nil {
		toSerialize["UnusedDisksState"] = o.UnusedDisksState
	}
	if o.VirtualDrives != nil {
		toSerialize["VirtualDrives"] = o.VirtualDrives
	}
	if o.DiskGroupPolicies != nil {
		toSerialize["DiskGroupPolicies"] = o.DiskGroupPolicies
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageStoragePolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageStoragePolicyAllOf := _StorageStoragePolicyAllOf{}

	if err = json.Unmarshal(bytes, &varStorageStoragePolicyAllOf); err == nil {
		*o = StorageStoragePolicyAllOf(varStorageStoragePolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "GlobalHotSpares")
		delete(additionalProperties, "RetainPolicyVirtualDrives")
		delete(additionalProperties, "UnusedDisksState")
		delete(additionalProperties, "VirtualDrives")
		delete(additionalProperties, "DiskGroupPolicies")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageStoragePolicyAllOf struct {
	value *StorageStoragePolicyAllOf
	isSet bool
}

func (v NullableStorageStoragePolicyAllOf) Get() *StorageStoragePolicyAllOf {
	return v.value
}

func (v *NullableStorageStoragePolicyAllOf) Set(val *StorageStoragePolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageStoragePolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageStoragePolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageStoragePolicyAllOf(val *StorageStoragePolicyAllOf) *NullableStorageStoragePolicyAllOf {
	return &NullableStorageStoragePolicyAllOf{value: val, isSet: true}
}

func (v NullableStorageStoragePolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageStoragePolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
