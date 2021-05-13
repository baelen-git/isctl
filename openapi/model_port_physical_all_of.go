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

// PortPhysicalAllOf Definition of the list of properties defined in 'port.Physical', excluding properties defined in parent classes.
type PortPhysicalAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Operational state of this port (enabled/disabled).
	OperState *string `json:"OperState,omitempty" yaml:"OperState,omitempty"`
	// Reason for this port's Operational state.
	OperStateQual *string `json:"OperStateQual,omitempty" yaml:"OperStateQual,omitempty"`
	// Switch physical port identifier.
	PortId *int64 `json:"PortId,omitempty" yaml:"PortId,omitempty"`
	// The role assigned to this port.
	Role *string `json:"Role,omitempty" yaml:"Role,omitempty"`
	// Switch expansion slot module identifier.
	SlotId *int64 `json:"SlotId,omitempty" yaml:"SlotId,omitempty"`
	// Switch Identifier that is local to a cluster.
	SwitchId *string `json:"SwitchId,omitempty" yaml:"SwitchId,omitempty"`
}

// NewPortPhysicalAllOf instantiates a new PortPhysicalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortPhysicalAllOf(classId string, objectType string) *PortPhysicalAllOf {
	this := PortPhysicalAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPortPhysicalAllOfWithDefaults instantiates a new PortPhysicalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortPhysicalAllOfWithDefaults() *PortPhysicalAllOf {
	this := PortPhysicalAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PortPhysicalAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PortPhysicalAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PortPhysicalAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PortPhysicalAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PortPhysicalAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PortPhysicalAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *PortPhysicalAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysicalAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *PortPhysicalAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *PortPhysicalAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetOperStateQual returns the OperStateQual field value if set, zero value otherwise.
func (o *PortPhysicalAllOf) GetOperStateQual() string {
	if o == nil || o.OperStateQual == nil {
		var ret string
		return ret
	}
	return *o.OperStateQual
}

// GetOperStateQualOk returns a tuple with the OperStateQual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysicalAllOf) GetOperStateQualOk() (*string, bool) {
	if o == nil || o.OperStateQual == nil {
		return nil, false
	}
	return o.OperStateQual, true
}

// HasOperStateQual returns a boolean if a field has been set.
func (o *PortPhysicalAllOf) HasOperStateQual() bool {
	if o != nil && o.OperStateQual != nil {
		return true
	}

	return false
}

// SetOperStateQual gets a reference to the given string and assigns it to the OperStateQual field.
func (o *PortPhysicalAllOf) SetOperStateQual(v string) {
	o.OperStateQual = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *PortPhysicalAllOf) GetPortId() int64 {
	if o == nil || o.PortId == nil {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysicalAllOf) GetPortIdOk() (*int64, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *PortPhysicalAllOf) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *PortPhysicalAllOf) SetPortId(v int64) {
	o.PortId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PortPhysicalAllOf) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysicalAllOf) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PortPhysicalAllOf) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *PortPhysicalAllOf) SetRole(v string) {
	o.Role = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *PortPhysicalAllOf) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysicalAllOf) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *PortPhysicalAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *PortPhysicalAllOf) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *PortPhysicalAllOf) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortPhysicalAllOf) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *PortPhysicalAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *PortPhysicalAllOf) SetSwitchId(v string) {
	o.SwitchId = &v
}

func (o PortPhysicalAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.OperStateQual != nil {
		toSerialize["OperStateQual"] = o.OperStateQual
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	return json.Marshal(toSerialize)
}

type NullablePortPhysicalAllOf struct {
	value *PortPhysicalAllOf
	isSet bool
}

func (v NullablePortPhysicalAllOf) Get() *PortPhysicalAllOf {
	return v.value
}

func (v *NullablePortPhysicalAllOf) Set(val *PortPhysicalAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePortPhysicalAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePortPhysicalAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortPhysicalAllOf(val *PortPhysicalAllOf) *NullablePortPhysicalAllOf {
	return &NullablePortPhysicalAllOf{value: val, isSet: true}
}

func (v NullablePortPhysicalAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortPhysicalAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
