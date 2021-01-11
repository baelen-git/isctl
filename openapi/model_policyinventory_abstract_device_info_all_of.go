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

// PolicyinventoryAbstractDeviceInfoAllOf Definition of the list of properties defined in 'policyinventory.AbstractDeviceInfo', excluding properties defined in parent classes.
type PolicyinventoryAbstractDeviceInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Configuration state of server profile config context.
	ConfigState *string `json:"ConfigState,omitempty" yaml:"ConfigState,omitempty"`
	// Control action of server profile config context.
	ControlAction *string `json:"ControlAction,omitempty" yaml:"ControlAction,omitempty"`
	// Error state of server profile config context.
	ErrorState *string                  `json:"ErrorState,omitempty" yaml:"ErrorState,omitempty"`
	JobInfo    []PolicyinventoryJobInfo `json:"JobInfo,omitempty" yaml:"JobInfo,omitempty"`
	// Operational state of server profile config context.
	OperState *string `json:"OperState,omitempty" yaml:"OperState,omitempty"`
	// Server profile MO ID of the server.
	ProfileMoId      *string                              `json:"ProfileMoId,omitempty" yaml:"ProfileMoId,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewPolicyinventoryAbstractDeviceInfoAllOf instantiates a new PolicyinventoryAbstractDeviceInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyinventoryAbstractDeviceInfoAllOf(classId string, objectType string) *PolicyinventoryAbstractDeviceInfoAllOf {
	this := PolicyinventoryAbstractDeviceInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyinventoryAbstractDeviceInfoAllOfWithDefaults instantiates a new PolicyinventoryAbstractDeviceInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyinventoryAbstractDeviceInfoAllOfWithDefaults() *PolicyinventoryAbstractDeviceInfoAllOf {
	this := PolicyinventoryAbstractDeviceInfoAllOf{}
	var classId string = "inventory.DeviceInfo"
	this.ClassId = classId
	var objectType string = "inventory.DeviceInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetControlAction returns the ControlAction field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetControlAction() string {
	if o == nil || o.ControlAction == nil {
		var ret string
		return ret
	}
	return *o.ControlAction
}

// GetControlActionOk returns a tuple with the ControlAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetControlActionOk() (*string, bool) {
	if o == nil || o.ControlAction == nil {
		return nil, false
	}
	return o.ControlAction, true
}

// HasControlAction returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasControlAction() bool {
	if o != nil && o.ControlAction != nil {
		return true
	}

	return false
}

// SetControlAction gets a reference to the given string and assigns it to the ControlAction field.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetControlAction(v string) {
	o.ControlAction = &v
}

// GetErrorState returns the ErrorState field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetErrorState() string {
	if o == nil || o.ErrorState == nil {
		var ret string
		return ret
	}
	return *o.ErrorState
}

// GetErrorStateOk returns a tuple with the ErrorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetErrorStateOk() (*string, bool) {
	if o == nil || o.ErrorState == nil {
		return nil, false
	}
	return o.ErrorState, true
}

// HasErrorState returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasErrorState() bool {
	if o != nil && o.ErrorState != nil {
		return true
	}

	return false
}

// SetErrorState gets a reference to the given string and assigns it to the ErrorState field.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetErrorState(v string) {
	o.ErrorState = &v
}

// GetJobInfo returns the JobInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetJobInfo() []PolicyinventoryJobInfo {
	if o == nil {
		var ret []PolicyinventoryJobInfo
		return ret
	}
	return o.JobInfo
}

// GetJobInfoOk returns a tuple with the JobInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetJobInfoOk() (*[]PolicyinventoryJobInfo, bool) {
	if o == nil || o.JobInfo == nil {
		return nil, false
	}
	return &o.JobInfo, true
}

// HasJobInfo returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasJobInfo() bool {
	if o != nil && o.JobInfo != nil {
		return true
	}

	return false
}

// SetJobInfo gets a reference to the given []PolicyinventoryJobInfo and assigns it to the JobInfo field.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetJobInfo(v []PolicyinventoryJobInfo) {
	o.JobInfo = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetProfileMoId returns the ProfileMoId field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetProfileMoId() string {
	if o == nil || o.ProfileMoId == nil {
		var ret string
		return ret
	}
	return *o.ProfileMoId
}

// GetProfileMoIdOk returns a tuple with the ProfileMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetProfileMoIdOk() (*string, bool) {
	if o == nil || o.ProfileMoId == nil {
		return nil, false
	}
	return o.ProfileMoId, true
}

// HasProfileMoId returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasProfileMoId() bool {
	if o != nil && o.ProfileMoId != nil {
		return true
	}

	return false
}

// SetProfileMoId gets a reference to the given string and assigns it to the ProfileMoId field.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetProfileMoId(v string) {
	o.ProfileMoId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o PolicyinventoryAbstractDeviceInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.ControlAction != nil {
		toSerialize["ControlAction"] = o.ControlAction
	}
	if o.ErrorState != nil {
		toSerialize["ErrorState"] = o.ErrorState
	}
	if o.JobInfo != nil {
		toSerialize["JobInfo"] = o.JobInfo
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.ProfileMoId != nil {
		toSerialize["ProfileMoId"] = o.ProfileMoId
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyinventoryAbstractDeviceInfoAllOf struct {
	value *PolicyinventoryAbstractDeviceInfoAllOf
	isSet bool
}

func (v NullablePolicyinventoryAbstractDeviceInfoAllOf) Get() *PolicyinventoryAbstractDeviceInfoAllOf {
	return v.value
}

func (v *NullablePolicyinventoryAbstractDeviceInfoAllOf) Set(val *PolicyinventoryAbstractDeviceInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyinventoryAbstractDeviceInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyinventoryAbstractDeviceInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyinventoryAbstractDeviceInfoAllOf(val *PolicyinventoryAbstractDeviceInfoAllOf) *NullablePolicyinventoryAbstractDeviceInfoAllOf {
	return &NullablePolicyinventoryAbstractDeviceInfoAllOf{value: val, isSet: true}
}

func (v NullablePolicyinventoryAbstractDeviceInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyinventoryAbstractDeviceInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
