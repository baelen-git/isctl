/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// WorkflowFileTemplateOpAllOf Definition of the list of properties defined in 'workflow.FileTemplateOp', excluding properties defined in parent classes.
type WorkflowFileTemplateOpAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Path of the template file on the connected device.
	TemplateFilePath *string `json:"TemplateFilePath,omitempty"`
	// Input values to render text output file from template file.
	TemplateValues       interface{} `json:"TemplateValues,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowFileTemplateOpAllOf WorkflowFileTemplateOpAllOf

// NewWorkflowFileTemplateOpAllOf instantiates a new WorkflowFileTemplateOpAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowFileTemplateOpAllOf(classId string, objectType string) *WorkflowFileTemplateOpAllOf {
	this := WorkflowFileTemplateOpAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowFileTemplateOpAllOfWithDefaults instantiates a new WorkflowFileTemplateOpAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowFileTemplateOpAllOfWithDefaults() *WorkflowFileTemplateOpAllOf {
	this := WorkflowFileTemplateOpAllOf{}
	var classId string = "workflow.FileTemplateOp"
	this.ClassId = classId
	var objectType string = "workflow.FileTemplateOp"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowFileTemplateOpAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowFileTemplateOpAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowFileTemplateOpAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowFileTemplateOpAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowFileTemplateOpAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowFileTemplateOpAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTemplateFilePath returns the TemplateFilePath field value if set, zero value otherwise.
func (o *WorkflowFileTemplateOpAllOf) GetTemplateFilePath() string {
	if o == nil || o.TemplateFilePath == nil {
		var ret string
		return ret
	}
	return *o.TemplateFilePath
}

// GetTemplateFilePathOk returns a tuple with the TemplateFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileTemplateOpAllOf) GetTemplateFilePathOk() (*string, bool) {
	if o == nil || o.TemplateFilePath == nil {
		return nil, false
	}
	return o.TemplateFilePath, true
}

// HasTemplateFilePath returns a boolean if a field has been set.
func (o *WorkflowFileTemplateOpAllOf) HasTemplateFilePath() bool {
	if o != nil && o.TemplateFilePath != nil {
		return true
	}

	return false
}

// SetTemplateFilePath gets a reference to the given string and assigns it to the TemplateFilePath field.
func (o *WorkflowFileTemplateOpAllOf) SetTemplateFilePath(v string) {
	o.TemplateFilePath = &v
}

// GetTemplateValues returns the TemplateValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowFileTemplateOpAllOf) GetTemplateValues() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TemplateValues
}

// GetTemplateValuesOk returns a tuple with the TemplateValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowFileTemplateOpAllOf) GetTemplateValuesOk() (*interface{}, bool) {
	if o == nil || o.TemplateValues == nil {
		return nil, false
	}
	return &o.TemplateValues, true
}

// HasTemplateValues returns a boolean if a field has been set.
func (o *WorkflowFileTemplateOpAllOf) HasTemplateValues() bool {
	if o != nil && o.TemplateValues != nil {
		return true
	}

	return false
}

// SetTemplateValues gets a reference to the given interface{} and assigns it to the TemplateValues field.
func (o *WorkflowFileTemplateOpAllOf) SetTemplateValues(v interface{}) {
	o.TemplateValues = v
}

func (o WorkflowFileTemplateOpAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TemplateFilePath != nil {
		toSerialize["TemplateFilePath"] = o.TemplateFilePath
	}
	if o.TemplateValues != nil {
		toSerialize["TemplateValues"] = o.TemplateValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowFileTemplateOpAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowFileTemplateOpAllOf := _WorkflowFileTemplateOpAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowFileTemplateOpAllOf); err == nil {
		*o = WorkflowFileTemplateOpAllOf(varWorkflowFileTemplateOpAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TemplateFilePath")
		delete(additionalProperties, "TemplateValues")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowFileTemplateOpAllOf struct {
	value *WorkflowFileTemplateOpAllOf
	isSet bool
}

func (v NullableWorkflowFileTemplateOpAllOf) Get() *WorkflowFileTemplateOpAllOf {
	return v.value
}

func (v *NullableWorkflowFileTemplateOpAllOf) Set(val *WorkflowFileTemplateOpAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowFileTemplateOpAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowFileTemplateOpAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowFileTemplateOpAllOf(val *WorkflowFileTemplateOpAllOf) *NullableWorkflowFileTemplateOpAllOf {
	return &NullableWorkflowFileTemplateOpAllOf{value: val, isSet: true}
}

func (v NullableWorkflowFileTemplateOpAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowFileTemplateOpAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
