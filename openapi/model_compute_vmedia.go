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

// ComputeVmedia Inventory of Virtual Media configuration and images uploaded.
type ComputeVmedia struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// State of the Virtual Media service on the server.
	Enabled *bool `json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
	// If enabled, allows encryption of all Virtual Media communications.
	Encryption *bool `json:"Encryption,omitempty" yaml:"Encryption,omitempty"`
	// If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host.
	LowPowerUsb         *bool                        `json:"LowPowerUsb,omitempty" yaml:"LowPowerUsb,omitempty"`
	ComputePhysicalUnit *ComputePhysicalRelationship `json:"ComputePhysicalUnit,omitempty" yaml:"ComputePhysicalUnit,omitempty"`
}

// NewComputeVmedia instantiates a new ComputeVmedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeVmedia(classId string, objectType string) *ComputeVmedia {
	this := ComputeVmedia{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var lowPowerUsb bool = true
	this.LowPowerUsb = &lowPowerUsb
	return &this
}

// NewComputeVmediaWithDefaults instantiates a new ComputeVmedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeVmediaWithDefaults() *ComputeVmedia {
	this := ComputeVmedia{}
	var classId string = "compute.Vmedia"
	this.ClassId = classId
	var objectType string = "compute.Vmedia"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var lowPowerUsb bool = true
	this.LowPowerUsb = &lowPowerUsb
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeVmedia) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeVmedia) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeVmedia) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeVmedia) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ComputeVmedia) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ComputeVmedia) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ComputeVmedia) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *ComputeVmedia) GetEncryption() bool {
	if o == nil || o.Encryption == nil {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetEncryptionOk() (*bool, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *ComputeVmedia) HasEncryption() bool {
	if o != nil && o.Encryption != nil {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *ComputeVmedia) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetLowPowerUsb returns the LowPowerUsb field value if set, zero value otherwise.
func (o *ComputeVmedia) GetLowPowerUsb() bool {
	if o == nil || o.LowPowerUsb == nil {
		var ret bool
		return ret
	}
	return *o.LowPowerUsb
}

// GetLowPowerUsbOk returns a tuple with the LowPowerUsb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetLowPowerUsbOk() (*bool, bool) {
	if o == nil || o.LowPowerUsb == nil {
		return nil, false
	}
	return o.LowPowerUsb, true
}

// HasLowPowerUsb returns a boolean if a field has been set.
func (o *ComputeVmedia) HasLowPowerUsb() bool {
	if o != nil && o.LowPowerUsb != nil {
		return true
	}

	return false
}

// SetLowPowerUsb gets a reference to the given bool and assigns it to the LowPowerUsb field.
func (o *ComputeVmedia) SetLowPowerUsb(v bool) {
	o.LowPowerUsb = &v
}

// GetComputePhysicalUnit returns the ComputePhysicalUnit field value if set, zero value otherwise.
func (o *ComputeVmedia) GetComputePhysicalUnit() ComputePhysicalRelationship {
	if o == nil || o.ComputePhysicalUnit == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.ComputePhysicalUnit
}

// GetComputePhysicalUnitOk returns a tuple with the ComputePhysicalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmedia) GetComputePhysicalUnitOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.ComputePhysicalUnit == nil {
		return nil, false
	}
	return o.ComputePhysicalUnit, true
}

// HasComputePhysicalUnit returns a boolean if a field has been set.
func (o *ComputeVmedia) HasComputePhysicalUnit() bool {
	if o != nil && o.ComputePhysicalUnit != nil {
		return true
	}

	return false
}

// SetComputePhysicalUnit gets a reference to the given ComputePhysicalRelationship and assigns it to the ComputePhysicalUnit field.
func (o *ComputeVmedia) SetComputePhysicalUnit(v ComputePhysicalRelationship) {
	o.ComputePhysicalUnit = &v
}

func (o ComputeVmedia) MarshalJSON() ([]byte, error) {
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
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Encryption != nil {
		toSerialize["Encryption"] = o.Encryption
	}
	if o.LowPowerUsb != nil {
		toSerialize["LowPowerUsb"] = o.LowPowerUsb
	}
	if o.ComputePhysicalUnit != nil {
		toSerialize["ComputePhysicalUnit"] = o.ComputePhysicalUnit
	}
	return json.Marshal(toSerialize)
}

type NullableComputeVmedia struct {
	value *ComputeVmedia
	isSet bool
}

func (v NullableComputeVmedia) Get() *ComputeVmedia {
	return v.value
}

func (v *NullableComputeVmedia) Set(val *ComputeVmedia) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeVmedia) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeVmedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeVmedia(val *ComputeVmedia) *NullableComputeVmedia {
	return &NullableComputeVmedia{value: val, isSet: true}
}

func (v NullableComputeVmedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeVmedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
