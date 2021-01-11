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

// WorkflowRollbackWorkflow Rollback workflow contains details about the workflow instance, tasks to be rollback along with the status and workflow instances.
type WorkflowRollbackWorkflow struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The action of the rollback workflow such as Create and Start. * `None` - If no action is set, then the default value is set to none for the action field. * `Create` - Create rollback workflow data for the execution of the rollback workflow. * `Start` - Start a new execution of the rollback workflow.
	Action *string `json:"Action,omitempty" yaml:"Action,omitempty"`
	// When set to true, if a task in the workflow fails, the rollback workflow continues to the subsequent task. When set to false, the rollback workflow execution halts if a task fails.
	ContinueOnTaskFailure *bool                          `json:"ContinueOnTaskFailure,omitempty" yaml:"ContinueOnTaskFailure,omitempty"`
	RollbackTasks         []WorkflowRollbackWorkflowTask `json:"RollbackTasks,omitempty" yaml:"RollbackTasks,omitempty"`
	SelectedTasks         []WorkflowRollbackWorkflowTask `json:"SelectedTasks,omitempty" yaml:"SelectedTasks,omitempty"`
	// Status of the rollback workflow instance (Created, Running, Completed, Failed). * `None` - If no status is set, then the default value is set none for the status field. * `Created` - Status of the rollback workflow when it identifies the eligible tasks for rollback. * `Running` - Status of the rollback workflow when it is in-progress. * `Completed` - Status of the rollback workflow after execution is successful. * `Failed` - Status of the rollback workflow after execution results in failure.
	Status          *string                           `json:"Status,omitempty" yaml:"Status,omitempty"`
	PrimaryWorkflow *WorkflowWorkflowInfoRelationship `json:"PrimaryWorkflow,omitempty" yaml:"PrimaryWorkflow,omitempty"`
	// An array of relationships to workflowWorkflowInfo resources.
	RollbackWorkflows []WorkflowWorkflowInfoRelationship `json:"RollbackWorkflows,omitempty" yaml:"RollbackWorkflows,omitempty"`
}

// NewWorkflowRollbackWorkflow instantiates a new WorkflowRollbackWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRollbackWorkflow(classId string, objectType string) *WorkflowRollbackWorkflow {
	this := WorkflowRollbackWorkflow{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var continueOnTaskFailure bool = true
	this.ContinueOnTaskFailure = &continueOnTaskFailure
	var status string = "None"
	this.Status = &status
	return &this
}

// NewWorkflowRollbackWorkflowWithDefaults instantiates a new WorkflowRollbackWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRollbackWorkflowWithDefaults() *WorkflowRollbackWorkflow {
	this := WorkflowRollbackWorkflow{}
	var classId string = "workflow.RollbackWorkflow"
	this.ClassId = classId
	var objectType string = "workflow.RollbackWorkflow"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var continueOnTaskFailure bool = true
	this.ContinueOnTaskFailure = &continueOnTaskFailure
	var status string = "None"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowRollbackWorkflow) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflow) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowRollbackWorkflow) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowRollbackWorkflow) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflow) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowRollbackWorkflow) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflow) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflow) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflow) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *WorkflowRollbackWorkflow) SetAction(v string) {
	o.Action = &v
}

// GetContinueOnTaskFailure returns the ContinueOnTaskFailure field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflow) GetContinueOnTaskFailure() bool {
	if o == nil || o.ContinueOnTaskFailure == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnTaskFailure
}

// GetContinueOnTaskFailureOk returns a tuple with the ContinueOnTaskFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflow) GetContinueOnTaskFailureOk() (*bool, bool) {
	if o == nil || o.ContinueOnTaskFailure == nil {
		return nil, false
	}
	return o.ContinueOnTaskFailure, true
}

// HasContinueOnTaskFailure returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflow) HasContinueOnTaskFailure() bool {
	if o != nil && o.ContinueOnTaskFailure != nil {
		return true
	}

	return false
}

// SetContinueOnTaskFailure gets a reference to the given bool and assigns it to the ContinueOnTaskFailure field.
func (o *WorkflowRollbackWorkflow) SetContinueOnTaskFailure(v bool) {
	o.ContinueOnTaskFailure = &v
}

// GetRollbackTasks returns the RollbackTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRollbackWorkflow) GetRollbackTasks() []WorkflowRollbackWorkflowTask {
	if o == nil {
		var ret []WorkflowRollbackWorkflowTask
		return ret
	}
	return o.RollbackTasks
}

// GetRollbackTasksOk returns a tuple with the RollbackTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRollbackWorkflow) GetRollbackTasksOk() (*[]WorkflowRollbackWorkflowTask, bool) {
	if o == nil || o.RollbackTasks == nil {
		return nil, false
	}
	return &o.RollbackTasks, true
}

// HasRollbackTasks returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflow) HasRollbackTasks() bool {
	if o != nil && o.RollbackTasks != nil {
		return true
	}

	return false
}

// SetRollbackTasks gets a reference to the given []WorkflowRollbackWorkflowTask and assigns it to the RollbackTasks field.
func (o *WorkflowRollbackWorkflow) SetRollbackTasks(v []WorkflowRollbackWorkflowTask) {
	o.RollbackTasks = v
}

// GetSelectedTasks returns the SelectedTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRollbackWorkflow) GetSelectedTasks() []WorkflowRollbackWorkflowTask {
	if o == nil {
		var ret []WorkflowRollbackWorkflowTask
		return ret
	}
	return o.SelectedTasks
}

// GetSelectedTasksOk returns a tuple with the SelectedTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRollbackWorkflow) GetSelectedTasksOk() (*[]WorkflowRollbackWorkflowTask, bool) {
	if o == nil || o.SelectedTasks == nil {
		return nil, false
	}
	return &o.SelectedTasks, true
}

// HasSelectedTasks returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflow) HasSelectedTasks() bool {
	if o != nil && o.SelectedTasks != nil {
		return true
	}

	return false
}

// SetSelectedTasks gets a reference to the given []WorkflowRollbackWorkflowTask and assigns it to the SelectedTasks field.
func (o *WorkflowRollbackWorkflow) SetSelectedTasks(v []WorkflowRollbackWorkflowTask) {
	o.SelectedTasks = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflow) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflow) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflow) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowRollbackWorkflow) SetStatus(v string) {
	o.Status = &v
}

// GetPrimaryWorkflow returns the PrimaryWorkflow field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflow) GetPrimaryWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.PrimaryWorkflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.PrimaryWorkflow
}

// GetPrimaryWorkflowOk returns a tuple with the PrimaryWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflow) GetPrimaryWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.PrimaryWorkflow == nil {
		return nil, false
	}
	return o.PrimaryWorkflow, true
}

// HasPrimaryWorkflow returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflow) HasPrimaryWorkflow() bool {
	if o != nil && o.PrimaryWorkflow != nil {
		return true
	}

	return false
}

// SetPrimaryWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the PrimaryWorkflow field.
func (o *WorkflowRollbackWorkflow) SetPrimaryWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.PrimaryWorkflow = &v
}

// GetRollbackWorkflows returns the RollbackWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRollbackWorkflow) GetRollbackWorkflows() []WorkflowWorkflowInfoRelationship {
	if o == nil {
		var ret []WorkflowWorkflowInfoRelationship
		return ret
	}
	return o.RollbackWorkflows
}

// GetRollbackWorkflowsOk returns a tuple with the RollbackWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRollbackWorkflow) GetRollbackWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RollbackWorkflows == nil {
		return nil, false
	}
	return &o.RollbackWorkflows, true
}

// HasRollbackWorkflows returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflow) HasRollbackWorkflows() bool {
	if o != nil && o.RollbackWorkflows != nil {
		return true
	}

	return false
}

// SetRollbackWorkflows gets a reference to the given []WorkflowWorkflowInfoRelationship and assigns it to the RollbackWorkflows field.
func (o *WorkflowRollbackWorkflow) SetRollbackWorkflows(v []WorkflowWorkflowInfoRelationship) {
	o.RollbackWorkflows = v
}

func (o WorkflowRollbackWorkflow) MarshalJSON() ([]byte, error) {
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
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.ContinueOnTaskFailure != nil {
		toSerialize["ContinueOnTaskFailure"] = o.ContinueOnTaskFailure
	}
	if o.RollbackTasks != nil {
		toSerialize["RollbackTasks"] = o.RollbackTasks
	}
	if o.SelectedTasks != nil {
		toSerialize["SelectedTasks"] = o.SelectedTasks
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.PrimaryWorkflow != nil {
		toSerialize["PrimaryWorkflow"] = o.PrimaryWorkflow
	}
	if o.RollbackWorkflows != nil {
		toSerialize["RollbackWorkflows"] = o.RollbackWorkflows
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRollbackWorkflow struct {
	value *WorkflowRollbackWorkflow
	isSet bool
}

func (v NullableWorkflowRollbackWorkflow) Get() *WorkflowRollbackWorkflow {
	return v.value
}

func (v *NullableWorkflowRollbackWorkflow) Set(val *WorkflowRollbackWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRollbackWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRollbackWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRollbackWorkflow(val *WorkflowRollbackWorkflow) *NullableWorkflowRollbackWorkflow {
	return &NullableWorkflowRollbackWorkflow{value: val, isSet: true}
}

func (v NullableWorkflowRollbackWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRollbackWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
