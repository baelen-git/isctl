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

// WorkflowInternalProperties Internal properties for a task definition which are not editable by the user.
type WorkflowInternalProperties struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// This field will hold the base task type like HttpBaseTask or RemoteAnsibleBaseTask.
	BaseTaskType *string                  `json:"BaseTaskType,omitempty" yaml:"BaseTaskType,omitempty"`
	Constraints  *WorkflowTaskConstraints `json:"Constraints,omitempty" yaml:"Constraints,omitempty"`
	// Denotes this is an internal task. Internal tasks will be hidden from the UI when executing a workflow.
	Internal *bool `json:"Internal,omitempty" yaml:"Internal,omitempty"`
	// The service that owns and is responsible for execution of the task.
	Owner *string `json:"Owner,omitempty" yaml:"Owner,omitempty"`
}

// NewWorkflowInternalProperties instantiates a new WorkflowInternalProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowInternalProperties() *WorkflowInternalProperties {
	this := WorkflowInternalProperties{}
	return &this
}

// NewWorkflowInternalPropertiesWithDefaults instantiates a new WorkflowInternalProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowInternalPropertiesWithDefaults() *WorkflowInternalProperties {
	this := WorkflowInternalProperties{}
	return &this
}

// GetBaseTaskType returns the BaseTaskType field value if set, zero value otherwise.
func (o *WorkflowInternalProperties) GetBaseTaskType() string {
	if o == nil || o.BaseTaskType == nil {
		var ret string
		return ret
	}
	return *o.BaseTaskType
}

// GetBaseTaskTypeOk returns a tuple with the BaseTaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInternalProperties) GetBaseTaskTypeOk() (*string, bool) {
	if o == nil || o.BaseTaskType == nil {
		return nil, false
	}
	return o.BaseTaskType, true
}

// HasBaseTaskType returns a boolean if a field has been set.
func (o *WorkflowInternalProperties) HasBaseTaskType() bool {
	if o != nil && o.BaseTaskType != nil {
		return true
	}

	return false
}

// SetBaseTaskType gets a reference to the given string and assigns it to the BaseTaskType field.
func (o *WorkflowInternalProperties) SetBaseTaskType(v string) {
	o.BaseTaskType = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *WorkflowInternalProperties) GetConstraints() WorkflowTaskConstraints {
	if o == nil || o.Constraints == nil {
		var ret WorkflowTaskConstraints
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInternalProperties) GetConstraintsOk() (*WorkflowTaskConstraints, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *WorkflowInternalProperties) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given WorkflowTaskConstraints and assigns it to the Constraints field.
func (o *WorkflowInternalProperties) SetConstraints(v WorkflowTaskConstraints) {
	o.Constraints = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *WorkflowInternalProperties) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInternalProperties) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *WorkflowInternalProperties) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *WorkflowInternalProperties) SetInternal(v bool) {
	o.Internal = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *WorkflowInternalProperties) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInternalProperties) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *WorkflowInternalProperties) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *WorkflowInternalProperties) SetOwner(v string) {
	o.Owner = &v
}

func (o WorkflowInternalProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.BaseTaskType != nil {
		toSerialize["BaseTaskType"] = o.BaseTaskType
	}
	if o.Constraints != nil {
		toSerialize["Constraints"] = o.Constraints
	}
	if o.Internal != nil {
		toSerialize["Internal"] = o.Internal
	}
	if o.Owner != nil {
		toSerialize["Owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowInternalProperties struct {
	value *WorkflowInternalProperties
	isSet bool
}

func (v NullableWorkflowInternalProperties) Get() *WorkflowInternalProperties {
	return v.value
}

func (v *NullableWorkflowInternalProperties) Set(val *WorkflowInternalProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowInternalProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowInternalProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowInternalProperties(val *WorkflowInternalProperties) *NullableWorkflowInternalProperties {
	return &NullableWorkflowInternalProperties{value: val, isSet: true}
}

func (v NullableWorkflowInternalProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowInternalProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
