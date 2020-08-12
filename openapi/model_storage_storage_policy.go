/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// StorageStoragePolicy The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The encryption of drives can be enabled through this policy using remote keys from a KMIP server.
type StorageStoragePolicy struct {
	PolicyAbstractPolicy `yaml:"PolicyAbstractPolicy,inline"`
	GlobalHotSpares      *[]StorageLocalDisk `json:"GlobalHotSpares,omitempty" yaml:"GlobalHotSpares,omitempty"`
	// Retains the virtual drives defined in policy if they exist already. If this flag is false, the existing virtual drives are removed and created again based on virtual drives in the policy.
	RetainPolicyVirtualDrives *bool `json:"RetainPolicyVirtualDrives,omitempty" yaml:"RetainPolicyVirtualDrives,omitempty"`
	// Unused Disks State is used to specify the state, unconfigured good or jbod, in which the disks that are not used in this policy should be moved. * `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group. * `Jbod` - JBOD state where the disks start showing up to host os.
	UnusedDisksState *string                      `json:"UnusedDisksState,omitempty" yaml:"UnusedDisksState,omitempty"`
	VirtualDrives    *[]StorageVirtualDriveConfig `json:"VirtualDrives,omitempty" yaml:"VirtualDrives,omitempty"`
	// An array of relationships to storageDiskGroupPolicy resources.
	DiskGroupPolicies []StorageDiskGroupPolicyRelationship  `json:"DiskGroupPolicies,omitempty" yaml:"DiskGroupPolicies,omitempty"`
	Organization      *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty" yaml:"Profiles,omitempty"`
}

// NewStorageStoragePolicy instantiates a new StorageStoragePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageStoragePolicy() *StorageStoragePolicy {
	this := StorageStoragePolicy{}
	var unusedDisksState string = "UnconfiguredGood"
	this.UnusedDisksState = &unusedDisksState
	return &this
}

// NewStorageStoragePolicyWithDefaults instantiates a new StorageStoragePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageStoragePolicyWithDefaults() *StorageStoragePolicy {
	this := StorageStoragePolicy{}
	var unusedDisksState string = "UnconfiguredGood"
	this.UnusedDisksState = &unusedDisksState
	return &this
}

// GetGlobalHotSpares returns the GlobalHotSpares field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetGlobalHotSpares() []StorageLocalDisk {
	if o == nil || o.GlobalHotSpares == nil {
		var ret []StorageLocalDisk
		return ret
	}
	return *o.GlobalHotSpares
}

// GetGlobalHotSparesOk returns a tuple with the GlobalHotSpares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetGlobalHotSparesOk() (*[]StorageLocalDisk, bool) {
	if o == nil || o.GlobalHotSpares == nil {
		return nil, false
	}
	return o.GlobalHotSpares, true
}

// HasGlobalHotSpares returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasGlobalHotSpares() bool {
	if o != nil && o.GlobalHotSpares != nil {
		return true
	}

	return false
}

// SetGlobalHotSpares gets a reference to the given []StorageLocalDisk and assigns it to the GlobalHotSpares field.
func (o *StorageStoragePolicy) SetGlobalHotSpares(v []StorageLocalDisk) {
	o.GlobalHotSpares = &v
}

// GetRetainPolicyVirtualDrives returns the RetainPolicyVirtualDrives field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetRetainPolicyVirtualDrives() bool {
	if o == nil || o.RetainPolicyVirtualDrives == nil {
		var ret bool
		return ret
	}
	return *o.RetainPolicyVirtualDrives
}

// GetRetainPolicyVirtualDrivesOk returns a tuple with the RetainPolicyVirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetRetainPolicyVirtualDrivesOk() (*bool, bool) {
	if o == nil || o.RetainPolicyVirtualDrives == nil {
		return nil, false
	}
	return o.RetainPolicyVirtualDrives, true
}

// HasRetainPolicyVirtualDrives returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasRetainPolicyVirtualDrives() bool {
	if o != nil && o.RetainPolicyVirtualDrives != nil {
		return true
	}

	return false
}

// SetRetainPolicyVirtualDrives gets a reference to the given bool and assigns it to the RetainPolicyVirtualDrives field.
func (o *StorageStoragePolicy) SetRetainPolicyVirtualDrives(v bool) {
	o.RetainPolicyVirtualDrives = &v
}

// GetUnusedDisksState returns the UnusedDisksState field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetUnusedDisksState() string {
	if o == nil || o.UnusedDisksState == nil {
		var ret string
		return ret
	}
	return *o.UnusedDisksState
}

// GetUnusedDisksStateOk returns a tuple with the UnusedDisksState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetUnusedDisksStateOk() (*string, bool) {
	if o == nil || o.UnusedDisksState == nil {
		return nil, false
	}
	return o.UnusedDisksState, true
}

// HasUnusedDisksState returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasUnusedDisksState() bool {
	if o != nil && o.UnusedDisksState != nil {
		return true
	}

	return false
}

// SetUnusedDisksState gets a reference to the given string and assigns it to the UnusedDisksState field.
func (o *StorageStoragePolicy) SetUnusedDisksState(v string) {
	o.UnusedDisksState = &v
}

// GetVirtualDrives returns the VirtualDrives field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetVirtualDrives() []StorageVirtualDriveConfig {
	if o == nil || o.VirtualDrives == nil {
		var ret []StorageVirtualDriveConfig
		return ret
	}
	return *o.VirtualDrives
}

// GetVirtualDrivesOk returns a tuple with the VirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetVirtualDrivesOk() (*[]StorageVirtualDriveConfig, bool) {
	if o == nil || o.VirtualDrives == nil {
		return nil, false
	}
	return o.VirtualDrives, true
}

// HasVirtualDrives returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasVirtualDrives() bool {
	if o != nil && o.VirtualDrives != nil {
		return true
	}

	return false
}

// SetVirtualDrives gets a reference to the given []StorageVirtualDriveConfig and assigns it to the VirtualDrives field.
func (o *StorageStoragePolicy) SetVirtualDrives(v []StorageVirtualDriveConfig) {
	o.VirtualDrives = &v
}

// GetDiskGroupPolicies returns the DiskGroupPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicy) GetDiskGroupPolicies() []StorageDiskGroupPolicyRelationship {
	if o == nil {
		var ret []StorageDiskGroupPolicyRelationship
		return ret
	}
	return o.DiskGroupPolicies
}

// GetDiskGroupPoliciesOk returns a tuple with the DiskGroupPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicy) GetDiskGroupPoliciesOk() (*[]StorageDiskGroupPolicyRelationship, bool) {
	if o == nil || o.DiskGroupPolicies == nil {
		return nil, false
	}
	return &o.DiskGroupPolicies, true
}

// HasDiskGroupPolicies returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasDiskGroupPolicies() bool {
	if o != nil && o.DiskGroupPolicies != nil {
		return true
	}

	return false
}

// SetDiskGroupPolicies gets a reference to the given []StorageDiskGroupPolicyRelationship and assigns it to the DiskGroupPolicies field.
func (o *StorageStoragePolicy) SetDiskGroupPolicies(v []StorageDiskGroupPolicyRelationship) {
	o.DiskGroupPolicies = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *StorageStoragePolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *StorageStoragePolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o StorageStoragePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
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
	return json.Marshal(toSerialize)
}

type NullableStorageStoragePolicy struct {
	value *StorageStoragePolicy
	isSet bool
}

func (v NullableStorageStoragePolicy) Get() *StorageStoragePolicy {
	return v.value
}

func (v *NullableStorageStoragePolicy) Set(val *StorageStoragePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageStoragePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageStoragePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageStoragePolicy(val *StorageStoragePolicy) *NullableStorageStoragePolicy {
	return &NullableStorageStoragePolicy{value: val, isSet: true}
}

func (v NullableStorageStoragePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageStoragePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
