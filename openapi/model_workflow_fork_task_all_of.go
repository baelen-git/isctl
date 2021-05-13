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

// WorkflowForkTaskAllOf Definition of the list of properties defined in 'workflow.ForkTask', excluding properties defined in parent classes.
type WorkflowForkTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string   `json:"ObjectType" yaml:"ObjectType"`
	ForkedTasks []string `json:"ForkedTasks,omitempty" yaml:"ForkedTasks,omitempty"`
	// Task name for the join control task that must follow a fork control task.
	JoinTask *string `json:"JoinTask,omitempty" yaml:"JoinTask,omitempty"`
}

// NewWorkflowForkTaskAllOf instantiates a new WorkflowForkTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowForkTaskAllOf(classId string, objectType string) *WorkflowForkTaskAllOf {
	this := WorkflowForkTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowForkTaskAllOfWithDefaults instantiates a new WorkflowForkTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowForkTaskAllOfWithDefaults() *WorkflowForkTaskAllOf {
	this := WorkflowForkTaskAllOf{}
	var classId string = "workflow.ForkTask"
	this.ClassId = classId
	var objectType string = "workflow.ForkTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowForkTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowForkTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowForkTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowForkTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowForkTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowForkTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetForkedTasks returns the ForkedTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowForkTaskAllOf) GetForkedTasks() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ForkedTasks
}

// GetForkedTasksOk returns a tuple with the ForkedTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowForkTaskAllOf) GetForkedTasksOk() (*[]string, bool) {
	if o == nil || o.ForkedTasks == nil {
		return nil, false
	}
	return &o.ForkedTasks, true
}

// HasForkedTasks returns a boolean if a field has been set.
func (o *WorkflowForkTaskAllOf) HasForkedTasks() bool {
	if o != nil && o.ForkedTasks != nil {
		return true
	}

	return false
}

// SetForkedTasks gets a reference to the given []string and assigns it to the ForkedTasks field.
func (o *WorkflowForkTaskAllOf) SetForkedTasks(v []string) {
	o.ForkedTasks = v
}

// GetJoinTask returns the JoinTask field value if set, zero value otherwise.
func (o *WorkflowForkTaskAllOf) GetJoinTask() string {
	if o == nil || o.JoinTask == nil {
		var ret string
		return ret
	}
	return *o.JoinTask
}

// GetJoinTaskOk returns a tuple with the JoinTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowForkTaskAllOf) GetJoinTaskOk() (*string, bool) {
	if o == nil || o.JoinTask == nil {
		return nil, false
	}
	return o.JoinTask, true
}

// HasJoinTask returns a boolean if a field has been set.
func (o *WorkflowForkTaskAllOf) HasJoinTask() bool {
	if o != nil && o.JoinTask != nil {
		return true
	}

	return false
}

// SetJoinTask gets a reference to the given string and assigns it to the JoinTask field.
func (o *WorkflowForkTaskAllOf) SetJoinTask(v string) {
	o.JoinTask = &v
}

func (o WorkflowForkTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ForkedTasks != nil {
		toSerialize["ForkedTasks"] = o.ForkedTasks
	}
	if o.JoinTask != nil {
		toSerialize["JoinTask"] = o.JoinTask
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowForkTaskAllOf struct {
	value *WorkflowForkTaskAllOf
	isSet bool
}

func (v NullableWorkflowForkTaskAllOf) Get() *WorkflowForkTaskAllOf {
	return v.value
}

func (v *NullableWorkflowForkTaskAllOf) Set(val *WorkflowForkTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowForkTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowForkTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowForkTaskAllOf(val *WorkflowForkTaskAllOf) *NullableWorkflowForkTaskAllOf {
	return &NullableWorkflowForkTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowForkTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowForkTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
