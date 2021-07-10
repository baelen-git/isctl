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

// HyperflexClusterBackupPolicyAllOf Definition of the list of properties defined in 'hyperflex.ClusterBackupPolicy', excluding properties defined in parent classes.
type HyperflexClusterBackupPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Backup datastore name prefix used during the auto creation of the datastore. All VMs created in this datastore will be automatically backed up.
	BackupDataStoreName *string `json:"BackupDataStoreName,omitempty"`
	// Replication data store size in backupDataStoreSizeUnit.
	BackupDataStoreSize *int64 `json:"BackupDataStoreSize,omitempty"`
	// Replication data store size.
	BackupDataStoreSizeUnit *string `json:"BackupDataStoreSizeUnit,omitempty"`
	// Replication cluster pairing name prefix.
	ReplicationPairNamePrefix *string                              `json:"ReplicationPairNamePrefix,omitempty"`
	ReplicationSchedule       NullableHyperflexReplicationSchedule `json:"ReplicationSchedule,omitempty"`
	// Number of snapshots that will be retained as part of the Multi Point in Time support.
	SnapshotRetentionCount *int64                        `json:"SnapshotRetentionCount,omitempty"`
	BackupTarget           *HyperflexClusterRelationship `json:"BackupTarget,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterBackupPolicyAllOf HyperflexClusterBackupPolicyAllOf

// NewHyperflexClusterBackupPolicyAllOf instantiates a new HyperflexClusterBackupPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterBackupPolicyAllOf(classId string, objectType string) *HyperflexClusterBackupPolicyAllOf {
	this := HyperflexClusterBackupPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var backupDataStoreName string = "backup-source-ds"
	this.BackupDataStoreName = &backupDataStoreName
	var backupDataStoreSize int64 = 2
	this.BackupDataStoreSize = &backupDataStoreSize
	var backupDataStoreSizeUnit string = "TB"
	this.BackupDataStoreSizeUnit = &backupDataStoreSizeUnit
	var replicationPairNamePrefix string = "backup"
	this.ReplicationPairNamePrefix = &replicationPairNamePrefix
	var snapshotRetentionCount int64 = 4
	this.SnapshotRetentionCount = &snapshotRetentionCount
	return &this
}

// NewHyperflexClusterBackupPolicyAllOfWithDefaults instantiates a new HyperflexClusterBackupPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterBackupPolicyAllOfWithDefaults() *HyperflexClusterBackupPolicyAllOf {
	this := HyperflexClusterBackupPolicyAllOf{}
	var classId string = "hyperflex.ClusterBackupPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterBackupPolicy"
	this.ObjectType = objectType
	var backupDataStoreName string = "backup-source-ds"
	this.BackupDataStoreName = &backupDataStoreName
	var backupDataStoreSize int64 = 2
	this.BackupDataStoreSize = &backupDataStoreSize
	var backupDataStoreSizeUnit string = "TB"
	this.BackupDataStoreSizeUnit = &backupDataStoreSizeUnit
	var replicationPairNamePrefix string = "backup"
	this.ReplicationPairNamePrefix = &replicationPairNamePrefix
	var snapshotRetentionCount int64 = 4
	this.SnapshotRetentionCount = &snapshotRetentionCount
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterBackupPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterBackupPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterBackupPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterBackupPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupDataStoreName returns the BackupDataStoreName field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreName() string {
	if o == nil || o.BackupDataStoreName == nil {
		var ret string
		return ret
	}
	return *o.BackupDataStoreName
}

// GetBackupDataStoreNameOk returns a tuple with the BackupDataStoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreNameOk() (*string, bool) {
	if o == nil || o.BackupDataStoreName == nil {
		return nil, false
	}
	return o.BackupDataStoreName, true
}

// HasBackupDataStoreName returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyAllOf) HasBackupDataStoreName() bool {
	if o != nil && o.BackupDataStoreName != nil {
		return true
	}

	return false
}

// SetBackupDataStoreName gets a reference to the given string and assigns it to the BackupDataStoreName field.
func (o *HyperflexClusterBackupPolicyAllOf) SetBackupDataStoreName(v string) {
	o.BackupDataStoreName = &v
}

// GetBackupDataStoreSize returns the BackupDataStoreSize field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreSize() int64 {
	if o == nil || o.BackupDataStoreSize == nil {
		var ret int64
		return ret
	}
	return *o.BackupDataStoreSize
}

// GetBackupDataStoreSizeOk returns a tuple with the BackupDataStoreSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreSizeOk() (*int64, bool) {
	if o == nil || o.BackupDataStoreSize == nil {
		return nil, false
	}
	return o.BackupDataStoreSize, true
}

// HasBackupDataStoreSize returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyAllOf) HasBackupDataStoreSize() bool {
	if o != nil && o.BackupDataStoreSize != nil {
		return true
	}

	return false
}

// SetBackupDataStoreSize gets a reference to the given int64 and assigns it to the BackupDataStoreSize field.
func (o *HyperflexClusterBackupPolicyAllOf) SetBackupDataStoreSize(v int64) {
	o.BackupDataStoreSize = &v
}

// GetBackupDataStoreSizeUnit returns the BackupDataStoreSizeUnit field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreSizeUnit() string {
	if o == nil || o.BackupDataStoreSizeUnit == nil {
		var ret string
		return ret
	}
	return *o.BackupDataStoreSizeUnit
}

// GetBackupDataStoreSizeUnitOk returns a tuple with the BackupDataStoreSizeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyAllOf) GetBackupDataStoreSizeUnitOk() (*string, bool) {
	if o == nil || o.BackupDataStoreSizeUnit == nil {
		return nil, false
	}
	return o.BackupDataStoreSizeUnit, true
}

// HasBackupDataStoreSizeUnit returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyAllOf) HasBackupDataStoreSizeUnit() bool {
	if o != nil && o.BackupDataStoreSizeUnit != nil {
		return true
	}

	return false
}

// SetBackupDataStoreSizeUnit gets a reference to the given string and assigns it to the BackupDataStoreSizeUnit field.
func (o *HyperflexClusterBackupPolicyAllOf) SetBackupDataStoreSizeUnit(v string) {
	o.BackupDataStoreSizeUnit = &v
}

// GetReplicationPairNamePrefix returns the ReplicationPairNamePrefix field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyAllOf) GetReplicationPairNamePrefix() string {
	if o == nil || o.ReplicationPairNamePrefix == nil {
		var ret string
		return ret
	}
	return *o.ReplicationPairNamePrefix
}

// GetReplicationPairNamePrefixOk returns a tuple with the ReplicationPairNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyAllOf) GetReplicationPairNamePrefixOk() (*string, bool) {
	if o == nil || o.ReplicationPairNamePrefix == nil {
		return nil, false
	}
	return o.ReplicationPairNamePrefix, true
}

// HasReplicationPairNamePrefix returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyAllOf) HasReplicationPairNamePrefix() bool {
	if o != nil && o.ReplicationPairNamePrefix != nil {
		return true
	}

	return false
}

// SetReplicationPairNamePrefix gets a reference to the given string and assigns it to the ReplicationPairNamePrefix field.
func (o *HyperflexClusterBackupPolicyAllOf) SetReplicationPairNamePrefix(v string) {
	o.ReplicationPairNamePrefix = &v
}

// GetReplicationSchedule returns the ReplicationSchedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterBackupPolicyAllOf) GetReplicationSchedule() HyperflexReplicationSchedule {
	if o == nil || o.ReplicationSchedule.Get() == nil {
		var ret HyperflexReplicationSchedule
		return ret
	}
	return *o.ReplicationSchedule.Get()
}

// GetReplicationScheduleOk returns a tuple with the ReplicationSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterBackupPolicyAllOf) GetReplicationScheduleOk() (*HyperflexReplicationSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicationSchedule.Get(), o.ReplicationSchedule.IsSet()
}

// HasReplicationSchedule returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyAllOf) HasReplicationSchedule() bool {
	if o != nil && o.ReplicationSchedule.IsSet() {
		return true
	}

	return false
}

// SetReplicationSchedule gets a reference to the given NullableHyperflexReplicationSchedule and assigns it to the ReplicationSchedule field.
func (o *HyperflexClusterBackupPolicyAllOf) SetReplicationSchedule(v HyperflexReplicationSchedule) {
	o.ReplicationSchedule.Set(&v)
}

// SetReplicationScheduleNil sets the value for ReplicationSchedule to be an explicit nil
func (o *HyperflexClusterBackupPolicyAllOf) SetReplicationScheduleNil() {
	o.ReplicationSchedule.Set(nil)
}

// UnsetReplicationSchedule ensures that no value is present for ReplicationSchedule, not even an explicit nil
func (o *HyperflexClusterBackupPolicyAllOf) UnsetReplicationSchedule() {
	o.ReplicationSchedule.Unset()
}

// GetSnapshotRetentionCount returns the SnapshotRetentionCount field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyAllOf) GetSnapshotRetentionCount() int64 {
	if o == nil || o.SnapshotRetentionCount == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotRetentionCount
}

// GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyAllOf) GetSnapshotRetentionCountOk() (*int64, bool) {
	if o == nil || o.SnapshotRetentionCount == nil {
		return nil, false
	}
	return o.SnapshotRetentionCount, true
}

// HasSnapshotRetentionCount returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyAllOf) HasSnapshotRetentionCount() bool {
	if o != nil && o.SnapshotRetentionCount != nil {
		return true
	}

	return false
}

// SetSnapshotRetentionCount gets a reference to the given int64 and assigns it to the SnapshotRetentionCount field.
func (o *HyperflexClusterBackupPolicyAllOf) SetSnapshotRetentionCount(v int64) {
	o.SnapshotRetentionCount = &v
}

// GetBackupTarget returns the BackupTarget field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyAllOf) GetBackupTarget() HyperflexClusterRelationship {
	if o == nil || o.BackupTarget == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.BackupTarget
}

// GetBackupTargetOk returns a tuple with the BackupTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyAllOf) GetBackupTargetOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.BackupTarget == nil {
		return nil, false
	}
	return o.BackupTarget, true
}

// HasBackupTarget returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyAllOf) HasBackupTarget() bool {
	if o != nil && o.BackupTarget != nil {
		return true
	}

	return false
}

// SetBackupTarget gets a reference to the given HyperflexClusterRelationship and assigns it to the BackupTarget field.
func (o *HyperflexClusterBackupPolicyAllOf) SetBackupTarget(v HyperflexClusterRelationship) {
	o.BackupTarget = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterBackupPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterBackupPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexClusterBackupPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexClusterBackupPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterBackupPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexClusterBackupPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexClusterBackupPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexClusterBackupPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BackupDataStoreName != nil {
		toSerialize["BackupDataStoreName"] = o.BackupDataStoreName
	}
	if o.BackupDataStoreSize != nil {
		toSerialize["BackupDataStoreSize"] = o.BackupDataStoreSize
	}
	if o.BackupDataStoreSizeUnit != nil {
		toSerialize["BackupDataStoreSizeUnit"] = o.BackupDataStoreSizeUnit
	}
	if o.ReplicationPairNamePrefix != nil {
		toSerialize["ReplicationPairNamePrefix"] = o.ReplicationPairNamePrefix
	}
	if o.ReplicationSchedule.IsSet() {
		toSerialize["ReplicationSchedule"] = o.ReplicationSchedule.Get()
	}
	if o.SnapshotRetentionCount != nil {
		toSerialize["SnapshotRetentionCount"] = o.SnapshotRetentionCount
	}
	if o.BackupTarget != nil {
		toSerialize["BackupTarget"] = o.BackupTarget
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterBackupPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexClusterBackupPolicyAllOf := _HyperflexClusterBackupPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexClusterBackupPolicyAllOf); err == nil {
		*o = HyperflexClusterBackupPolicyAllOf(varHyperflexClusterBackupPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupDataStoreName")
		delete(additionalProperties, "BackupDataStoreSize")
		delete(additionalProperties, "BackupDataStoreSizeUnit")
		delete(additionalProperties, "ReplicationPairNamePrefix")
		delete(additionalProperties, "ReplicationSchedule")
		delete(additionalProperties, "SnapshotRetentionCount")
		delete(additionalProperties, "BackupTarget")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexClusterBackupPolicyAllOf struct {
	value *HyperflexClusterBackupPolicyAllOf
	isSet bool
}

func (v NullableHyperflexClusterBackupPolicyAllOf) Get() *HyperflexClusterBackupPolicyAllOf {
	return v.value
}

func (v *NullableHyperflexClusterBackupPolicyAllOf) Set(val *HyperflexClusterBackupPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterBackupPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterBackupPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterBackupPolicyAllOf(val *HyperflexClusterBackupPolicyAllOf) *NullableHyperflexClusterBackupPolicyAllOf {
	return &NullableHyperflexClusterBackupPolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexClusterBackupPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterBackupPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
