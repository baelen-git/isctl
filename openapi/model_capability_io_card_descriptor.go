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

// CapabilityIoCardDescriptor Descriptor that uniquely identifies an IO card module.
type CapabilityIoCardDescriptor struct {
	CapabilityHardwareDescriptor `yaml:"CapabilityHardwareDescriptor,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Number of hif ports per blade for the iocard module.
	NumHifPorts *int64 `json:"NumHifPorts,omitempty" yaml:"NumHifPorts,omitempty"`
	// Revision for the iocard module.
	Revision *string `json:"Revision,omitempty" yaml:"Revision,omitempty"`
}

// NewCapabilityIoCardDescriptor instantiates a new CapabilityIoCardDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityIoCardDescriptor(classId string, objectType string) *CapabilityIoCardDescriptor {
	this := CapabilityIoCardDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityIoCardDescriptorWithDefaults instantiates a new CapabilityIoCardDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityIoCardDescriptorWithDefaults() *CapabilityIoCardDescriptor {
	this := CapabilityIoCardDescriptor{}
	var classId string = "capability.IoCardDescriptor"
	this.ClassId = classId
	var objectType string = "capability.IoCardDescriptor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityIoCardDescriptor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityIoCardDescriptor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityIoCardDescriptor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityIoCardDescriptor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityIoCardDescriptor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityIoCardDescriptor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNumHifPorts returns the NumHifPorts field value if set, zero value otherwise.
func (o *CapabilityIoCardDescriptor) GetNumHifPorts() int64 {
	if o == nil || o.NumHifPorts == nil {
		var ret int64
		return ret
	}
	return *o.NumHifPorts
}

// GetNumHifPortsOk returns a tuple with the NumHifPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityIoCardDescriptor) GetNumHifPortsOk() (*int64, bool) {
	if o == nil || o.NumHifPorts == nil {
		return nil, false
	}
	return o.NumHifPorts, true
}

// HasNumHifPorts returns a boolean if a field has been set.
func (o *CapabilityIoCardDescriptor) HasNumHifPorts() bool {
	if o != nil && o.NumHifPorts != nil {
		return true
	}

	return false
}

// SetNumHifPorts gets a reference to the given int64 and assigns it to the NumHifPorts field.
func (o *CapabilityIoCardDescriptor) SetNumHifPorts(v int64) {
	o.NumHifPorts = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *CapabilityIoCardDescriptor) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityIoCardDescriptor) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *CapabilityIoCardDescriptor) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *CapabilityIoCardDescriptor) SetRevision(v string) {
	o.Revision = &v
}

func (o CapabilityIoCardDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityHardwareDescriptor, errCapabilityHardwareDescriptor := json.Marshal(o.CapabilityHardwareDescriptor)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	errCapabilityHardwareDescriptor = json.Unmarshal([]byte(serializedCapabilityHardwareDescriptor), &toSerialize)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NumHifPorts != nil {
		toSerialize["NumHifPorts"] = o.NumHifPorts
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}
	return json.Marshal(toSerialize)
}

type NullableCapabilityIoCardDescriptor struct {
	value *CapabilityIoCardDescriptor
	isSet bool
}

func (v NullableCapabilityIoCardDescriptor) Get() *CapabilityIoCardDescriptor {
	return v.value
}

func (v *NullableCapabilityIoCardDescriptor) Set(val *CapabilityIoCardDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityIoCardDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityIoCardDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityIoCardDescriptor(val *CapabilityIoCardDescriptor) *NullableCapabilityIoCardDescriptor {
	return &NullableCapabilityIoCardDescriptor{value: val, isSet: true}
}

func (v NullableCapabilityIoCardDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityIoCardDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
