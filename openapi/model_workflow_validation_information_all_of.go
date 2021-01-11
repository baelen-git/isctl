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

// WorkflowValidationInformationAllOf Definition of the list of properties defined in 'workflow.ValidationInformation', excluding properties defined in parent classes.
type WorkflowValidationInformationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default). * `NotValidated` - The state when workflow definition has not been validated. * `Valid` - The state when workflow definition is valid. * `Invalid` - The state when workflow definition is invalid.
	State           *string                   `json:"State,omitempty" yaml:"State,omitempty"`
	ValidationError []WorkflowValidationError `json:"ValidationError,omitempty" yaml:"ValidationError,omitempty"`
}

// NewWorkflowValidationInformationAllOf instantiates a new WorkflowValidationInformationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowValidationInformationAllOf(classId string, objectType string) *WorkflowValidationInformationAllOf {
	this := WorkflowValidationInformationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var state string = "NotValidated"
	this.State = &state
	return &this
}

// NewWorkflowValidationInformationAllOfWithDefaults instantiates a new WorkflowValidationInformationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowValidationInformationAllOfWithDefaults() *WorkflowValidationInformationAllOf {
	this := WorkflowValidationInformationAllOf{}
	var classId string = "workflow.ValidationInformation"
	this.ClassId = classId
	var objectType string = "workflow.ValidationInformation"
	this.ObjectType = objectType
	var state string = "NotValidated"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowValidationInformationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowValidationInformationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowValidationInformationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowValidationInformationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowValidationInformationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowValidationInformationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *WorkflowValidationInformationAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowValidationInformationAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *WorkflowValidationInformationAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *WorkflowValidationInformationAllOf) SetState(v string) {
	o.State = &v
}

// GetValidationError returns the ValidationError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowValidationInformationAllOf) GetValidationError() []WorkflowValidationError {
	if o == nil {
		var ret []WorkflowValidationError
		return ret
	}
	return o.ValidationError
}

// GetValidationErrorOk returns a tuple with the ValidationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowValidationInformationAllOf) GetValidationErrorOk() (*[]WorkflowValidationError, bool) {
	if o == nil || o.ValidationError == nil {
		return nil, false
	}
	return &o.ValidationError, true
}

// HasValidationError returns a boolean if a field has been set.
func (o *WorkflowValidationInformationAllOf) HasValidationError() bool {
	if o != nil && o.ValidationError != nil {
		return true
	}

	return false
}

// SetValidationError gets a reference to the given []WorkflowValidationError and assigns it to the ValidationError field.
func (o *WorkflowValidationInformationAllOf) SetValidationError(v []WorkflowValidationError) {
	o.ValidationError = v
}

func (o WorkflowValidationInformationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.ValidationError != nil {
		toSerialize["ValidationError"] = o.ValidationError
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowValidationInformationAllOf struct {
	value *WorkflowValidationInformationAllOf
	isSet bool
}

func (v NullableWorkflowValidationInformationAllOf) Get() *WorkflowValidationInformationAllOf {
	return v.value
}

func (v *NullableWorkflowValidationInformationAllOf) Set(val *WorkflowValidationInformationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowValidationInformationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowValidationInformationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowValidationInformationAllOf(val *WorkflowValidationInformationAllOf) *NullableWorkflowValidationInformationAllOf {
	return &NullableWorkflowValidationInformationAllOf{value: val, isSet: true}
}

func (v NullableWorkflowValidationInformationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowValidationInformationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
