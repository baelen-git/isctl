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

// WorkflowBuildTaskMetaOwner Contains the list of dynamic workflow types that a microservice supports.
type WorkflowBuildTaskMetaOwner struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The microservice owner responsible for the tasks.
	Service       *string  `json:"Service,omitempty" yaml:"Service,omitempty"`
	WorkflowTypes []string `json:"WorkflowTypes,omitempty" yaml:"WorkflowTypes,omitempty"`
}

// NewWorkflowBuildTaskMetaOwner instantiates a new WorkflowBuildTaskMetaOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBuildTaskMetaOwner(classId string, objectType string) *WorkflowBuildTaskMetaOwner {
	this := WorkflowBuildTaskMetaOwner{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowBuildTaskMetaOwnerWithDefaults instantiates a new WorkflowBuildTaskMetaOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBuildTaskMetaOwnerWithDefaults() *WorkflowBuildTaskMetaOwner {
	this := WorkflowBuildTaskMetaOwner{}
	var classId string = "workflow.BuildTaskMetaOwner"
	this.ClassId = classId
	var objectType string = "workflow.BuildTaskMetaOwner"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowBuildTaskMetaOwner) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaOwner) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowBuildTaskMetaOwner) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowBuildTaskMetaOwner) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaOwner) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowBuildTaskMetaOwner) SetObjectType(v string) {
	o.ObjectType = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *WorkflowBuildTaskMetaOwner) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaOwner) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *WorkflowBuildTaskMetaOwner) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *WorkflowBuildTaskMetaOwner) SetService(v string) {
	o.Service = &v
}

// GetWorkflowTypes returns the WorkflowTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBuildTaskMetaOwner) GetWorkflowTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.WorkflowTypes
}

// GetWorkflowTypesOk returns a tuple with the WorkflowTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBuildTaskMetaOwner) GetWorkflowTypesOk() (*[]string, bool) {
	if o == nil || o.WorkflowTypes == nil {
		return nil, false
	}
	return &o.WorkflowTypes, true
}

// HasWorkflowTypes returns a boolean if a field has been set.
func (o *WorkflowBuildTaskMetaOwner) HasWorkflowTypes() bool {
	if o != nil && o.WorkflowTypes != nil {
		return true
	}

	return false
}

// SetWorkflowTypes gets a reference to the given []string and assigns it to the WorkflowTypes field.
func (o *WorkflowBuildTaskMetaOwner) SetWorkflowTypes(v []string) {
	o.WorkflowTypes = v
}

func (o WorkflowBuildTaskMetaOwner) MarshalJSON() ([]byte, error) {
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
	if o.Service != nil {
		toSerialize["Service"] = o.Service
	}
	if o.WorkflowTypes != nil {
		toSerialize["WorkflowTypes"] = o.WorkflowTypes
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowBuildTaskMetaOwner struct {
	value *WorkflowBuildTaskMetaOwner
	isSet bool
}

func (v NullableWorkflowBuildTaskMetaOwner) Get() *WorkflowBuildTaskMetaOwner {
	return v.value
}

func (v *NullableWorkflowBuildTaskMetaOwner) Set(val *WorkflowBuildTaskMetaOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBuildTaskMetaOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBuildTaskMetaOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBuildTaskMetaOwner(val *WorkflowBuildTaskMetaOwner) *NullableWorkflowBuildTaskMetaOwner {
	return &NullableWorkflowBuildTaskMetaOwner{value: val, isSet: true}
}

func (v NullableWorkflowBuildTaskMetaOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBuildTaskMetaOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
