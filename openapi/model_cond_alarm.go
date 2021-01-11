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
	"time"
)

// CondAlarm A state-full entity representing a found problem. Alarms can be reported by the managed system itself or can be determined by Intersight.
type CondAlarm struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Alarm acknowledgment state. Default value is None. * `None` - The Enum value None represents that the alarm is not acknowledged and is included as part of health status and overall alarm count. * `Acknowledge` - The Enum value Acknowledge represents that the alarm is acknowledged by user. The alarm will be ignored from the health status and overall alarm count.
	Acknowledge *string `json:"Acknowledge,omitempty" yaml:"Acknowledge,omitempty"`
	// User who acknowledged the alarm.
	AcknowledgeBy *string `json:"AcknowledgeBy,omitempty" yaml:"AcknowledgeBy,omitempty"`
	// Time at which the alarm was acknowledged by the user.
	AcknowledgeTime *time.Time `json:"AcknowledgeTime,omitempty" yaml:"AcknowledgeTime,omitempty"`
	// Display name of the affected object on which the alarm is raised. The name uniquely identifies an alarm target such that it can be distinguished from similar objects managed by Intersight.
	AffectedMoDisplayName *string `json:"AffectedMoDisplayName,omitempty" yaml:"AffectedMoDisplayName,omitempty"`
	// MoId of the affected object from the managed system's point of view.
	AffectedMoId *string `json:"AffectedMoId,omitempty" yaml:"AffectedMoId,omitempty"`
	// Managed system affected object type. For example Adaptor, FI, CIMC.
	AffectedMoType *string `json:"AffectedMoType,omitempty" yaml:"AffectedMoType,omitempty"`
	// A unique key for an alarm instance, consists of a combination of a unique system name and msAffectedObject.
	AffectedObject *string `json:"AffectedObject,omitempty" yaml:"AffectedObject,omitempty"`
	// Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault.
	AncestorMoId *string `json:"AncestorMoId,omitempty" yaml:"AncestorMoId,omitempty"`
	// Parent MO type of the fault from managed system. For example, Blade for adaptor fault.
	AncestorMoType *string `json:"AncestorMoType,omitempty" yaml:"AncestorMoType,omitempty"`
	// A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code.
	Code *string `json:"Code,omitempty" yaml:"Code,omitempty"`
	// The time the alarm was created.
	CreationTime *time.Time `json:"CreationTime,omitempty" yaml:"CreationTime,omitempty"`
	// A longer description of the alarm than the name. The description contains details of the component reporting the issue.
	Description *string `json:"Description,omitempty" yaml:"Description,omitempty"`
	// The time the alarm last had a change in severity.
	LastTransitionTime *time.Time `json:"LastTransitionTime,omitempty" yaml:"LastTransitionTime,omitempty"`
	// A unique key for the alarm from the managed system's point of view. For example, in the case of UCS, this is the fault's dn.
	MsAffectedObject *string `json:"MsAffectedObject,omitempty" yaml:"MsAffectedObject,omitempty"`
	// Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// The original severity when the alarm was first created. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
	OrigSeverity *string `json:"OrigSeverity,omitempty" yaml:"OrigSeverity,omitempty"`
	// The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
	Severity         *string                              `json:"Severity,omitempty" yaml:"Severity,omitempty"`
	AffectedMo       *MoBaseMoRelationship                `json:"AffectedMo,omitempty" yaml:"AffectedMo,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewCondAlarm instantiates a new CondAlarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondAlarm(classId string, objectType string) *CondAlarm {
	this := CondAlarm{}
	this.ClassId = classId
	this.ObjectType = objectType
	var acknowledge string = "None"
	this.Acknowledge = &acknowledge
	var origSeverity string = "None"
	this.OrigSeverity = &origSeverity
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// NewCondAlarmWithDefaults instantiates a new CondAlarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondAlarmWithDefaults() *CondAlarm {
	this := CondAlarm{}
	var classId string = "cond.Alarm"
	this.ClassId = classId
	var objectType string = "cond.Alarm"
	this.ObjectType = objectType
	var acknowledge string = "None"
	this.Acknowledge = &acknowledge
	var origSeverity string = "None"
	this.OrigSeverity = &origSeverity
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// GetClassId returns the ClassId field value
func (o *CondAlarm) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CondAlarm) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CondAlarm) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CondAlarm) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAcknowledge returns the Acknowledge field value if set, zero value otherwise.
func (o *CondAlarm) GetAcknowledge() string {
	if o == nil || o.Acknowledge == nil {
		var ret string
		return ret
	}
	return *o.Acknowledge
}

// GetAcknowledgeOk returns a tuple with the Acknowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetAcknowledgeOk() (*string, bool) {
	if o == nil || o.Acknowledge == nil {
		return nil, false
	}
	return o.Acknowledge, true
}

// HasAcknowledge returns a boolean if a field has been set.
func (o *CondAlarm) HasAcknowledge() bool {
	if o != nil && o.Acknowledge != nil {
		return true
	}

	return false
}

// SetAcknowledge gets a reference to the given string and assigns it to the Acknowledge field.
func (o *CondAlarm) SetAcknowledge(v string) {
	o.Acknowledge = &v
}

// GetAcknowledgeBy returns the AcknowledgeBy field value if set, zero value otherwise.
func (o *CondAlarm) GetAcknowledgeBy() string {
	if o == nil || o.AcknowledgeBy == nil {
		var ret string
		return ret
	}
	return *o.AcknowledgeBy
}

// GetAcknowledgeByOk returns a tuple with the AcknowledgeBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetAcknowledgeByOk() (*string, bool) {
	if o == nil || o.AcknowledgeBy == nil {
		return nil, false
	}
	return o.AcknowledgeBy, true
}

// HasAcknowledgeBy returns a boolean if a field has been set.
func (o *CondAlarm) HasAcknowledgeBy() bool {
	if o != nil && o.AcknowledgeBy != nil {
		return true
	}

	return false
}

// SetAcknowledgeBy gets a reference to the given string and assigns it to the AcknowledgeBy field.
func (o *CondAlarm) SetAcknowledgeBy(v string) {
	o.AcknowledgeBy = &v
}

// GetAcknowledgeTime returns the AcknowledgeTime field value if set, zero value otherwise.
func (o *CondAlarm) GetAcknowledgeTime() time.Time {
	if o == nil || o.AcknowledgeTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AcknowledgeTime
}

// GetAcknowledgeTimeOk returns a tuple with the AcknowledgeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetAcknowledgeTimeOk() (*time.Time, bool) {
	if o == nil || o.AcknowledgeTime == nil {
		return nil, false
	}
	return o.AcknowledgeTime, true
}

// HasAcknowledgeTime returns a boolean if a field has been set.
func (o *CondAlarm) HasAcknowledgeTime() bool {
	if o != nil && o.AcknowledgeTime != nil {
		return true
	}

	return false
}

// SetAcknowledgeTime gets a reference to the given time.Time and assigns it to the AcknowledgeTime field.
func (o *CondAlarm) SetAcknowledgeTime(v time.Time) {
	o.AcknowledgeTime = &v
}

// GetAffectedMoDisplayName returns the AffectedMoDisplayName field value if set, zero value otherwise.
func (o *CondAlarm) GetAffectedMoDisplayName() string {
	if o == nil || o.AffectedMoDisplayName == nil {
		var ret string
		return ret
	}
	return *o.AffectedMoDisplayName
}

// GetAffectedMoDisplayNameOk returns a tuple with the AffectedMoDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetAffectedMoDisplayNameOk() (*string, bool) {
	if o == nil || o.AffectedMoDisplayName == nil {
		return nil, false
	}
	return o.AffectedMoDisplayName, true
}

// HasAffectedMoDisplayName returns a boolean if a field has been set.
func (o *CondAlarm) HasAffectedMoDisplayName() bool {
	if o != nil && o.AffectedMoDisplayName != nil {
		return true
	}

	return false
}

// SetAffectedMoDisplayName gets a reference to the given string and assigns it to the AffectedMoDisplayName field.
func (o *CondAlarm) SetAffectedMoDisplayName(v string) {
	o.AffectedMoDisplayName = &v
}

// GetAffectedMoId returns the AffectedMoId field value if set, zero value otherwise.
func (o *CondAlarm) GetAffectedMoId() string {
	if o == nil || o.AffectedMoId == nil {
		var ret string
		return ret
	}
	return *o.AffectedMoId
}

// GetAffectedMoIdOk returns a tuple with the AffectedMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetAffectedMoIdOk() (*string, bool) {
	if o == nil || o.AffectedMoId == nil {
		return nil, false
	}
	return o.AffectedMoId, true
}

// HasAffectedMoId returns a boolean if a field has been set.
func (o *CondAlarm) HasAffectedMoId() bool {
	if o != nil && o.AffectedMoId != nil {
		return true
	}

	return false
}

// SetAffectedMoId gets a reference to the given string and assigns it to the AffectedMoId field.
func (o *CondAlarm) SetAffectedMoId(v string) {
	o.AffectedMoId = &v
}

// GetAffectedMoType returns the AffectedMoType field value if set, zero value otherwise.
func (o *CondAlarm) GetAffectedMoType() string {
	if o == nil || o.AffectedMoType == nil {
		var ret string
		return ret
	}
	return *o.AffectedMoType
}

// GetAffectedMoTypeOk returns a tuple with the AffectedMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetAffectedMoTypeOk() (*string, bool) {
	if o == nil || o.AffectedMoType == nil {
		return nil, false
	}
	return o.AffectedMoType, true
}

// HasAffectedMoType returns a boolean if a field has been set.
func (o *CondAlarm) HasAffectedMoType() bool {
	if o != nil && o.AffectedMoType != nil {
		return true
	}

	return false
}

// SetAffectedMoType gets a reference to the given string and assigns it to the AffectedMoType field.
func (o *CondAlarm) SetAffectedMoType(v string) {
	o.AffectedMoType = &v
}

// GetAffectedObject returns the AffectedObject field value if set, zero value otherwise.
func (o *CondAlarm) GetAffectedObject() string {
	if o == nil || o.AffectedObject == nil {
		var ret string
		return ret
	}
	return *o.AffectedObject
}

// GetAffectedObjectOk returns a tuple with the AffectedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetAffectedObjectOk() (*string, bool) {
	if o == nil || o.AffectedObject == nil {
		return nil, false
	}
	return o.AffectedObject, true
}

// HasAffectedObject returns a boolean if a field has been set.
func (o *CondAlarm) HasAffectedObject() bool {
	if o != nil && o.AffectedObject != nil {
		return true
	}

	return false
}

// SetAffectedObject gets a reference to the given string and assigns it to the AffectedObject field.
func (o *CondAlarm) SetAffectedObject(v string) {
	o.AffectedObject = &v
}

// GetAncestorMoId returns the AncestorMoId field value if set, zero value otherwise.
func (o *CondAlarm) GetAncestorMoId() string {
	if o == nil || o.AncestorMoId == nil {
		var ret string
		return ret
	}
	return *o.AncestorMoId
}

// GetAncestorMoIdOk returns a tuple with the AncestorMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetAncestorMoIdOk() (*string, bool) {
	if o == nil || o.AncestorMoId == nil {
		return nil, false
	}
	return o.AncestorMoId, true
}

// HasAncestorMoId returns a boolean if a field has been set.
func (o *CondAlarm) HasAncestorMoId() bool {
	if o != nil && o.AncestorMoId != nil {
		return true
	}

	return false
}

// SetAncestorMoId gets a reference to the given string and assigns it to the AncestorMoId field.
func (o *CondAlarm) SetAncestorMoId(v string) {
	o.AncestorMoId = &v
}

// GetAncestorMoType returns the AncestorMoType field value if set, zero value otherwise.
func (o *CondAlarm) GetAncestorMoType() string {
	if o == nil || o.AncestorMoType == nil {
		var ret string
		return ret
	}
	return *o.AncestorMoType
}

// GetAncestorMoTypeOk returns a tuple with the AncestorMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetAncestorMoTypeOk() (*string, bool) {
	if o == nil || o.AncestorMoType == nil {
		return nil, false
	}
	return o.AncestorMoType, true
}

// HasAncestorMoType returns a boolean if a field has been set.
func (o *CondAlarm) HasAncestorMoType() bool {
	if o != nil && o.AncestorMoType != nil {
		return true
	}

	return false
}

// SetAncestorMoType gets a reference to the given string and assigns it to the AncestorMoType field.
func (o *CondAlarm) SetAncestorMoType(v string) {
	o.AncestorMoType = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CondAlarm) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CondAlarm) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CondAlarm) SetCode(v string) {
	o.Code = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *CondAlarm) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *CondAlarm) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *CondAlarm) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CondAlarm) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CondAlarm) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CondAlarm) SetDescription(v string) {
	o.Description = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *CondAlarm) GetLastTransitionTime() time.Time {
	if o == nil || o.LastTransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *CondAlarm) HasLastTransitionTime() bool {
	if o != nil && o.LastTransitionTime != nil {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given time.Time and assigns it to the LastTransitionTime field.
func (o *CondAlarm) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime = &v
}

// GetMsAffectedObject returns the MsAffectedObject field value if set, zero value otherwise.
func (o *CondAlarm) GetMsAffectedObject() string {
	if o == nil || o.MsAffectedObject == nil {
		var ret string
		return ret
	}
	return *o.MsAffectedObject
}

// GetMsAffectedObjectOk returns a tuple with the MsAffectedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetMsAffectedObjectOk() (*string, bool) {
	if o == nil || o.MsAffectedObject == nil {
		return nil, false
	}
	return o.MsAffectedObject, true
}

// HasMsAffectedObject returns a boolean if a field has been set.
func (o *CondAlarm) HasMsAffectedObject() bool {
	if o != nil && o.MsAffectedObject != nil {
		return true
	}

	return false
}

// SetMsAffectedObject gets a reference to the given string and assigns it to the MsAffectedObject field.
func (o *CondAlarm) SetMsAffectedObject(v string) {
	o.MsAffectedObject = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CondAlarm) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CondAlarm) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CondAlarm) SetName(v string) {
	o.Name = &v
}

// GetOrigSeverity returns the OrigSeverity field value if set, zero value otherwise.
func (o *CondAlarm) GetOrigSeverity() string {
	if o == nil || o.OrigSeverity == nil {
		var ret string
		return ret
	}
	return *o.OrigSeverity
}

// GetOrigSeverityOk returns a tuple with the OrigSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetOrigSeverityOk() (*string, bool) {
	if o == nil || o.OrigSeverity == nil {
		return nil, false
	}
	return o.OrigSeverity, true
}

// HasOrigSeverity returns a boolean if a field has been set.
func (o *CondAlarm) HasOrigSeverity() bool {
	if o != nil && o.OrigSeverity != nil {
		return true
	}

	return false
}

// SetOrigSeverity gets a reference to the given string and assigns it to the OrigSeverity field.
func (o *CondAlarm) SetOrigSeverity(v string) {
	o.OrigSeverity = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *CondAlarm) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *CondAlarm) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *CondAlarm) SetSeverity(v string) {
	o.Severity = &v
}

// GetAffectedMo returns the AffectedMo field value if set, zero value otherwise.
func (o *CondAlarm) GetAffectedMo() MoBaseMoRelationship {
	if o == nil || o.AffectedMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AffectedMo
}

// GetAffectedMoOk returns a tuple with the AffectedMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetAffectedMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AffectedMo == nil {
		return nil, false
	}
	return o.AffectedMo, true
}

// HasAffectedMo returns a boolean if a field has been set.
func (o *CondAlarm) HasAffectedMo() bool {
	if o != nil && o.AffectedMo != nil {
		return true
	}

	return false
}

// SetAffectedMo gets a reference to the given MoBaseMoRelationship and assigns it to the AffectedMo field.
func (o *CondAlarm) SetAffectedMo(v MoBaseMoRelationship) {
	o.AffectedMo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *CondAlarm) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarm) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *CondAlarm) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *CondAlarm) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o CondAlarm) MarshalJSON() ([]byte, error) {
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
	if o.Acknowledge != nil {
		toSerialize["Acknowledge"] = o.Acknowledge
	}
	if o.AcknowledgeBy != nil {
		toSerialize["AcknowledgeBy"] = o.AcknowledgeBy
	}
	if o.AcknowledgeTime != nil {
		toSerialize["AcknowledgeTime"] = o.AcknowledgeTime
	}
	if o.AffectedMoDisplayName != nil {
		toSerialize["AffectedMoDisplayName"] = o.AffectedMoDisplayName
	}
	if o.AffectedMoId != nil {
		toSerialize["AffectedMoId"] = o.AffectedMoId
	}
	if o.AffectedMoType != nil {
		toSerialize["AffectedMoType"] = o.AffectedMoType
	}
	if o.AffectedObject != nil {
		toSerialize["AffectedObject"] = o.AffectedObject
	}
	if o.AncestorMoId != nil {
		toSerialize["AncestorMoId"] = o.AncestorMoId
	}
	if o.AncestorMoType != nil {
		toSerialize["AncestorMoType"] = o.AncestorMoType
	}
	if o.Code != nil {
		toSerialize["Code"] = o.Code
	}
	if o.CreationTime != nil {
		toSerialize["CreationTime"] = o.CreationTime
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.LastTransitionTime != nil {
		toSerialize["LastTransitionTime"] = o.LastTransitionTime
	}
	if o.MsAffectedObject != nil {
		toSerialize["MsAffectedObject"] = o.MsAffectedObject
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OrigSeverity != nil {
		toSerialize["OrigSeverity"] = o.OrigSeverity
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.AffectedMo != nil {
		toSerialize["AffectedMo"] = o.AffectedMo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableCondAlarm struct {
	value *CondAlarm
	isSet bool
}

func (v NullableCondAlarm) Get() *CondAlarm {
	return v.value
}

func (v *NullableCondAlarm) Set(val *CondAlarm) {
	v.value = val
	v.isSet = true
}

func (v NullableCondAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableCondAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondAlarm(val *CondAlarm) *NullableCondAlarm {
	return &NullableCondAlarm{value: val, isSet: true}
}

func (v NullableCondAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
