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

// HyperflexVmBackupInfo Virtual Machine backup information.
type HyperflexVmBackupInfo struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Description of the backup status of this VmBackupInfo. * `InitializingProtection` - Protection has started, but not completed. * `Protected` - Protection has completed successfully. * `ExceedsInterval` - Protection has not completed successfully in over two times the backup interval.
	BackupStatus               *string                                 `json:"BackupStatus,omitempty" yaml:"BackupStatus,omitempty"`
	ClusterEntityReference     NullableHyperflexEntityReference        `json:"ClusterEntityReference,omitempty" yaml:"ClusterEntityReference,omitempty"`
	ClusterIdProtectionInfoMap []HyperflexMapClusterIdToProtectionInfo `json:"ClusterIdProtectionInfoMap,omitempty" yaml:"ClusterIdProtectionInfoMap,omitempty"`
	Error                      NullableHyperflexErrorStack             `json:"Error,omitempty" yaml:"Error,omitempty"`
	// Description of the protection status of this VmBackupInfo. * `PREPARE_FAILOVER_STARTED` - The protection status is prepare failover started. * `PREPARE_FAILOVER_FAILED` - The protection status is prepare failover failed. * `PREPARE_FAILOVER_COMPLETED` - The protection status is prepaire failover completed. * `FAILOVER_STARTED` - The protection status is failover started. * `FAILOVER_FAILED` - The protection status is failover failed. * `FAILOVER_COMPLETED` - The protection status is failover completed. * `PREPARE_REVERSEPROTECT_STARTED` - The protection status is prepare reverse protect started. * `PREPARE_REVERSEPROTECT_FAILED` - The protection status is prepare reverse protect failed. * `PREPARE_REVERSEPROTECT_COMPLETED` - The protection status is prepair reverse protect completed. * `REVERSEPROTECT_STARTED` - The protection status is reverse protect started. * `REVERSEPROTECT_FAILED` - The protection status is reverse protect failed. * `ACTIVE` - The protection status is active. * `CREATION_IN_PROGRESS` - The protection status is failover in progress. * `CREATION_FAILED` - The protection status is creation failed. * `LOCAL_RESTORE_STARTED` - The protection status is local restore started. * `LOCAL_RESTORE_FAILED` - The protection status is local restore failed.
	ProtectionStatus  *string                                          `json:"ProtectionStatus,omitempty" yaml:"ProtectionStatus,omitempty"`
	Schedule          []HyperflexReplicationClusterReferenceToSchedule `json:"Schedule,omitempty" yaml:"Schedule,omitempty"`
	VmEntityReference NullableHyperflexEntityReference                 `json:"VmEntityReference,omitempty" yaml:"VmEntityReference,omitempty"`
	VmInfo            NullableHyperflexVirtualMachine                  `json:"VmInfo,omitempty" yaml:"VmInfo,omitempty"`
	SrcBackupCluster  *HyperflexBackupClusterRelationship              `json:"SrcBackupCluster,omitempty" yaml:"SrcBackupCluster,omitempty"`
	SrcCluster        *HyperflexClusterRelationship                    `json:"SrcCluster,omitempty" yaml:"SrcCluster,omitempty"`
	TgtCluster        *HyperflexClusterRelationship                    `json:"TgtCluster,omitempty" yaml:"TgtCluster,omitempty"`
}

// NewHyperflexVmBackupInfo instantiates a new HyperflexVmBackupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVmBackupInfo(classId string, objectType string) *HyperflexVmBackupInfo {
	this := HyperflexVmBackupInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var backupStatus string = "InitializingProtection"
	this.BackupStatus = &backupStatus
	var protectionStatus string = "PREPARE_FAILOVER_STARTED"
	this.ProtectionStatus = &protectionStatus
	return &this
}

// NewHyperflexVmBackupInfoWithDefaults instantiates a new HyperflexVmBackupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVmBackupInfoWithDefaults() *HyperflexVmBackupInfo {
	this := HyperflexVmBackupInfo{}
	var classId string = "hyperflex.VmBackupInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.VmBackupInfo"
	this.ObjectType = objectType
	var backupStatus string = "InitializingProtection"
	this.BackupStatus = &backupStatus
	var protectionStatus string = "PREPARE_FAILOVER_STARTED"
	this.ProtectionStatus = &protectionStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVmBackupInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmBackupInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVmBackupInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVmBackupInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmBackupInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVmBackupInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupStatus returns the BackupStatus field value if set, zero value otherwise.
func (o *HyperflexVmBackupInfo) GetBackupStatus() string {
	if o == nil || o.BackupStatus == nil {
		var ret string
		return ret
	}
	return *o.BackupStatus
}

// GetBackupStatusOk returns a tuple with the BackupStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmBackupInfo) GetBackupStatusOk() (*string, bool) {
	if o == nil || o.BackupStatus == nil {
		return nil, false
	}
	return o.BackupStatus, true
}

// HasBackupStatus returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasBackupStatus() bool {
	if o != nil && o.BackupStatus != nil {
		return true
	}

	return false
}

// SetBackupStatus gets a reference to the given string and assigns it to the BackupStatus field.
func (o *HyperflexVmBackupInfo) SetBackupStatus(v string) {
	o.BackupStatus = &v
}

// GetClusterEntityReference returns the ClusterEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmBackupInfo) GetClusterEntityReference() HyperflexEntityReference {
	if o == nil || o.ClusterEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.ClusterEntityReference.Get()
}

// GetClusterEntityReferenceOk returns a tuple with the ClusterEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmBackupInfo) GetClusterEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterEntityReference.Get(), o.ClusterEntityReference.IsSet()
}

// HasClusterEntityReference returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasClusterEntityReference() bool {
	if o != nil && o.ClusterEntityReference.IsSet() {
		return true
	}

	return false
}

// SetClusterEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the ClusterEntityReference field.
func (o *HyperflexVmBackupInfo) SetClusterEntityReference(v HyperflexEntityReference) {
	o.ClusterEntityReference.Set(&v)
}

// SetClusterEntityReferenceNil sets the value for ClusterEntityReference to be an explicit nil
func (o *HyperflexVmBackupInfo) SetClusterEntityReferenceNil() {
	o.ClusterEntityReference.Set(nil)
}

// UnsetClusterEntityReference ensures that no value is present for ClusterEntityReference, not even an explicit nil
func (o *HyperflexVmBackupInfo) UnsetClusterEntityReference() {
	o.ClusterEntityReference.Unset()
}

// GetClusterIdProtectionInfoMap returns the ClusterIdProtectionInfoMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmBackupInfo) GetClusterIdProtectionInfoMap() []HyperflexMapClusterIdToProtectionInfo {
	if o == nil {
		var ret []HyperflexMapClusterIdToProtectionInfo
		return ret
	}
	return o.ClusterIdProtectionInfoMap
}

// GetClusterIdProtectionInfoMapOk returns a tuple with the ClusterIdProtectionInfoMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmBackupInfo) GetClusterIdProtectionInfoMapOk() (*[]HyperflexMapClusterIdToProtectionInfo, bool) {
	if o == nil || o.ClusterIdProtectionInfoMap == nil {
		return nil, false
	}
	return &o.ClusterIdProtectionInfoMap, true
}

// HasClusterIdProtectionInfoMap returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasClusterIdProtectionInfoMap() bool {
	if o != nil && o.ClusterIdProtectionInfoMap != nil {
		return true
	}

	return false
}

// SetClusterIdProtectionInfoMap gets a reference to the given []HyperflexMapClusterIdToProtectionInfo and assigns it to the ClusterIdProtectionInfoMap field.
func (o *HyperflexVmBackupInfo) SetClusterIdProtectionInfoMap(v []HyperflexMapClusterIdToProtectionInfo) {
	o.ClusterIdProtectionInfoMap = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmBackupInfo) GetError() HyperflexErrorStack {
	if o == nil || o.Error.Get() == nil {
		var ret HyperflexErrorStack
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmBackupInfo) GetErrorOk() (*HyperflexErrorStack, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableHyperflexErrorStack and assigns it to the Error field.
func (o *HyperflexVmBackupInfo) SetError(v HyperflexErrorStack) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *HyperflexVmBackupInfo) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *HyperflexVmBackupInfo) UnsetError() {
	o.Error.Unset()
}

// GetProtectionStatus returns the ProtectionStatus field value if set, zero value otherwise.
func (o *HyperflexVmBackupInfo) GetProtectionStatus() string {
	if o == nil || o.ProtectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ProtectionStatus
}

// GetProtectionStatusOk returns a tuple with the ProtectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmBackupInfo) GetProtectionStatusOk() (*string, bool) {
	if o == nil || o.ProtectionStatus == nil {
		return nil, false
	}
	return o.ProtectionStatus, true
}

// HasProtectionStatus returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasProtectionStatus() bool {
	if o != nil && o.ProtectionStatus != nil {
		return true
	}

	return false
}

// SetProtectionStatus gets a reference to the given string and assigns it to the ProtectionStatus field.
func (o *HyperflexVmBackupInfo) SetProtectionStatus(v string) {
	o.ProtectionStatus = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmBackupInfo) GetSchedule() []HyperflexReplicationClusterReferenceToSchedule {
	if o == nil {
		var ret []HyperflexReplicationClusterReferenceToSchedule
		return ret
	}
	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmBackupInfo) GetScheduleOk() (*[]HyperflexReplicationClusterReferenceToSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given []HyperflexReplicationClusterReferenceToSchedule and assigns it to the Schedule field.
func (o *HyperflexVmBackupInfo) SetSchedule(v []HyperflexReplicationClusterReferenceToSchedule) {
	o.Schedule = v
}

// GetVmEntityReference returns the VmEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmBackupInfo) GetVmEntityReference() HyperflexEntityReference {
	if o == nil || o.VmEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.VmEntityReference.Get()
}

// GetVmEntityReferenceOk returns a tuple with the VmEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmBackupInfo) GetVmEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmEntityReference.Get(), o.VmEntityReference.IsSet()
}

// HasVmEntityReference returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasVmEntityReference() bool {
	if o != nil && o.VmEntityReference.IsSet() {
		return true
	}

	return false
}

// SetVmEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the VmEntityReference field.
func (o *HyperflexVmBackupInfo) SetVmEntityReference(v HyperflexEntityReference) {
	o.VmEntityReference.Set(&v)
}

// SetVmEntityReferenceNil sets the value for VmEntityReference to be an explicit nil
func (o *HyperflexVmBackupInfo) SetVmEntityReferenceNil() {
	o.VmEntityReference.Set(nil)
}

// UnsetVmEntityReference ensures that no value is present for VmEntityReference, not even an explicit nil
func (o *HyperflexVmBackupInfo) UnsetVmEntityReference() {
	o.VmEntityReference.Unset()
}

// GetVmInfo returns the VmInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmBackupInfo) GetVmInfo() HyperflexVirtualMachine {
	if o == nil || o.VmInfo.Get() == nil {
		var ret HyperflexVirtualMachine
		return ret
	}
	return *o.VmInfo.Get()
}

// GetVmInfoOk returns a tuple with the VmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmBackupInfo) GetVmInfoOk() (*HyperflexVirtualMachine, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmInfo.Get(), o.VmInfo.IsSet()
}

// HasVmInfo returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasVmInfo() bool {
	if o != nil && o.VmInfo.IsSet() {
		return true
	}

	return false
}

// SetVmInfo gets a reference to the given NullableHyperflexVirtualMachine and assigns it to the VmInfo field.
func (o *HyperflexVmBackupInfo) SetVmInfo(v HyperflexVirtualMachine) {
	o.VmInfo.Set(&v)
}

// SetVmInfoNil sets the value for VmInfo to be an explicit nil
func (o *HyperflexVmBackupInfo) SetVmInfoNil() {
	o.VmInfo.Set(nil)
}

// UnsetVmInfo ensures that no value is present for VmInfo, not even an explicit nil
func (o *HyperflexVmBackupInfo) UnsetVmInfo() {
	o.VmInfo.Unset()
}

// GetSrcBackupCluster returns the SrcBackupCluster field value if set, zero value otherwise.
func (o *HyperflexVmBackupInfo) GetSrcBackupCluster() HyperflexBackupClusterRelationship {
	if o == nil || o.SrcBackupCluster == nil {
		var ret HyperflexBackupClusterRelationship
		return ret
	}
	return *o.SrcBackupCluster
}

// GetSrcBackupClusterOk returns a tuple with the SrcBackupCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmBackupInfo) GetSrcBackupClusterOk() (*HyperflexBackupClusterRelationship, bool) {
	if o == nil || o.SrcBackupCluster == nil {
		return nil, false
	}
	return o.SrcBackupCluster, true
}

// HasSrcBackupCluster returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasSrcBackupCluster() bool {
	if o != nil && o.SrcBackupCluster != nil {
		return true
	}

	return false
}

// SetSrcBackupCluster gets a reference to the given HyperflexBackupClusterRelationship and assigns it to the SrcBackupCluster field.
func (o *HyperflexVmBackupInfo) SetSrcBackupCluster(v HyperflexBackupClusterRelationship) {
	o.SrcBackupCluster = &v
}

// GetSrcCluster returns the SrcCluster field value if set, zero value otherwise.
func (o *HyperflexVmBackupInfo) GetSrcCluster() HyperflexClusterRelationship {
	if o == nil || o.SrcCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.SrcCluster
}

// GetSrcClusterOk returns a tuple with the SrcCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmBackupInfo) GetSrcClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.SrcCluster == nil {
		return nil, false
	}
	return o.SrcCluster, true
}

// HasSrcCluster returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasSrcCluster() bool {
	if o != nil && o.SrcCluster != nil {
		return true
	}

	return false
}

// SetSrcCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the SrcCluster field.
func (o *HyperflexVmBackupInfo) SetSrcCluster(v HyperflexClusterRelationship) {
	o.SrcCluster = &v
}

// GetTgtCluster returns the TgtCluster field value if set, zero value otherwise.
func (o *HyperflexVmBackupInfo) GetTgtCluster() HyperflexClusterRelationship {
	if o == nil || o.TgtCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.TgtCluster
}

// GetTgtClusterOk returns a tuple with the TgtCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmBackupInfo) GetTgtClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.TgtCluster == nil {
		return nil, false
	}
	return o.TgtCluster, true
}

// HasTgtCluster returns a boolean if a field has been set.
func (o *HyperflexVmBackupInfo) HasTgtCluster() bool {
	if o != nil && o.TgtCluster != nil {
		return true
	}

	return false
}

// SetTgtCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the TgtCluster field.
func (o *HyperflexVmBackupInfo) SetTgtCluster(v HyperflexClusterRelationship) {
	o.TgtCluster = &v
}

func (o HyperflexVmBackupInfo) MarshalJSON() ([]byte, error) {
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
	if o.BackupStatus != nil {
		toSerialize["BackupStatus"] = o.BackupStatus
	}
	if o.ClusterEntityReference.IsSet() {
		toSerialize["ClusterEntityReference"] = o.ClusterEntityReference.Get()
	}
	if o.ClusterIdProtectionInfoMap != nil {
		toSerialize["ClusterIdProtectionInfoMap"] = o.ClusterIdProtectionInfoMap
	}
	if o.Error.IsSet() {
		toSerialize["Error"] = o.Error.Get()
	}
	if o.ProtectionStatus != nil {
		toSerialize["ProtectionStatus"] = o.ProtectionStatus
	}
	if o.Schedule != nil {
		toSerialize["Schedule"] = o.Schedule
	}
	if o.VmEntityReference.IsSet() {
		toSerialize["VmEntityReference"] = o.VmEntityReference.Get()
	}
	if o.VmInfo.IsSet() {
		toSerialize["VmInfo"] = o.VmInfo.Get()
	}
	if o.SrcBackupCluster != nil {
		toSerialize["SrcBackupCluster"] = o.SrcBackupCluster
	}
	if o.SrcCluster != nil {
		toSerialize["SrcCluster"] = o.SrcCluster
	}
	if o.TgtCluster != nil {
		toSerialize["TgtCluster"] = o.TgtCluster
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexVmBackupInfo struct {
	value *HyperflexVmBackupInfo
	isSet bool
}

func (v NullableHyperflexVmBackupInfo) Get() *HyperflexVmBackupInfo {
	return v.value
}

func (v *NullableHyperflexVmBackupInfo) Set(val *HyperflexVmBackupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVmBackupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVmBackupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVmBackupInfo(val *HyperflexVmBackupInfo) *NullableHyperflexVmBackupInfo {
	return &NullableHyperflexVmBackupInfo{value: val, isSet: true}
}

func (v NullableHyperflexVmBackupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVmBackupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
