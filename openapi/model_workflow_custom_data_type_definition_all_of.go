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

// WorkflowCustomDataTypeDefinitionAllOf Definition of the list of properties defined in 'workflow.CustomDataTypeDefinition', excluding properties defined in parent classes.
type WorkflowCustomDataTypeDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// When true this data type definition is a collection of type definitions to represent composite data like JSON.
	CompositeType *bool `json:"CompositeType,omitempty" yaml:"CompositeType,omitempty"`
	// A human-friendly description of this custom data type indicating it's domain and usage.
	Description *string `json:"Description,omitempty" yaml:"Description,omitempty"`
	// A user friendly short name to identify the custom data type definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), or an underscore (_).
	Label *string `json:"Label,omitempty" yaml:"Label,omitempty"`
	// The name of custom data type definition. The valid name can contain lower case and upper case alphabetic characters, digits and special characters '-' and '_'.
	Name           *string                      `json:"Name,omitempty" yaml:"Name,omitempty"`
	ParameterSet   []WorkflowParameterSet       `json:"ParameterSet,omitempty" yaml:"ParameterSet,omitempty"`
	TypeDefinition []WorkflowBaseDataType       `json:"TypeDefinition,omitempty" yaml:"TypeDefinition,omitempty"`
	Catalog        *WorkflowCatalogRelationship `json:"Catalog,omitempty" yaml:"Catalog,omitempty"`
}

// NewWorkflowCustomDataTypeDefinitionAllOf instantiates a new WorkflowCustomDataTypeDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCustomDataTypeDefinitionAllOf(classId string, objectType string) *WorkflowCustomDataTypeDefinitionAllOf {
	this := WorkflowCustomDataTypeDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var compositeType bool = false
	this.CompositeType = &compositeType
	return &this
}

// NewWorkflowCustomDataTypeDefinitionAllOfWithDefaults instantiates a new WorkflowCustomDataTypeDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCustomDataTypeDefinitionAllOfWithDefaults() *WorkflowCustomDataTypeDefinitionAllOf {
	this := WorkflowCustomDataTypeDefinitionAllOf{}
	var classId string = "workflow.CustomDataTypeDefinition"
	this.ClassId = classId
	var objectType string = "workflow.CustomDataTypeDefinition"
	this.ObjectType = objectType
	var compositeType bool = false
	this.CompositeType = &compositeType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCustomDataTypeDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCustomDataTypeDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCompositeType returns the CompositeType field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetCompositeType() bool {
	if o == nil || o.CompositeType == nil {
		var ret bool
		return ret
	}
	return *o.CompositeType
}

// GetCompositeTypeOk returns a tuple with the CompositeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetCompositeTypeOk() (*bool, bool) {
	if o == nil || o.CompositeType == nil {
		return nil, false
	}
	return o.CompositeType, true
}

// HasCompositeType returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) HasCompositeType() bool {
	if o != nil && o.CompositeType != nil {
		return true
	}

	return false
}

// SetCompositeType gets a reference to the given bool and assigns it to the CompositeType field.
func (o *WorkflowCustomDataTypeDefinitionAllOf) SetCompositeType(v bool) {
	o.CompositeType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowCustomDataTypeDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowCustomDataTypeDefinitionAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowCustomDataTypeDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetParameterSet returns the ParameterSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetParameterSet() []WorkflowParameterSet {
	if o == nil {
		var ret []WorkflowParameterSet
		return ret
	}
	return o.ParameterSet
}

// GetParameterSetOk returns a tuple with the ParameterSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetParameterSetOk() (*[]WorkflowParameterSet, bool) {
	if o == nil || o.ParameterSet == nil {
		return nil, false
	}
	return &o.ParameterSet, true
}

// HasParameterSet returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) HasParameterSet() bool {
	if o != nil && o.ParameterSet != nil {
		return true
	}

	return false
}

// SetParameterSet gets a reference to the given []WorkflowParameterSet and assigns it to the ParameterSet field.
func (o *WorkflowCustomDataTypeDefinitionAllOf) SetParameterSet(v []WorkflowParameterSet) {
	o.ParameterSet = v
}

// GetTypeDefinition returns the TypeDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetTypeDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.TypeDefinition
}

// GetTypeDefinitionOk returns a tuple with the TypeDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetTypeDefinitionOk() (*[]WorkflowBaseDataType, bool) {
	if o == nil || o.TypeDefinition == nil {
		return nil, false
	}
	return &o.TypeDefinition, true
}

// HasTypeDefinition returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) HasTypeDefinition() bool {
	if o != nil && o.TypeDefinition != nil {
		return true
	}

	return false
}

// SetTypeDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the TypeDefinition field.
func (o *WorkflowCustomDataTypeDefinitionAllOf) SetTypeDefinition(v []WorkflowBaseDataType) {
	o.TypeDefinition = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinitionAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given WorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowCustomDataTypeDefinitionAllOf) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog = &v
}

func (o WorkflowCustomDataTypeDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CompositeType != nil {
		toSerialize["CompositeType"] = o.CompositeType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ParameterSet != nil {
		toSerialize["ParameterSet"] = o.ParameterSet
	}
	if o.TypeDefinition != nil {
		toSerialize["TypeDefinition"] = o.TypeDefinition
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowCustomDataTypeDefinitionAllOf struct {
	value *WorkflowCustomDataTypeDefinitionAllOf
	isSet bool
}

func (v NullableWorkflowCustomDataTypeDefinitionAllOf) Get() *WorkflowCustomDataTypeDefinitionAllOf {
	return v.value
}

func (v *NullableWorkflowCustomDataTypeDefinitionAllOf) Set(val *WorkflowCustomDataTypeDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCustomDataTypeDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCustomDataTypeDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCustomDataTypeDefinitionAllOf(val *WorkflowCustomDataTypeDefinitionAllOf) *NullableWorkflowCustomDataTypeDefinitionAllOf {
	return &NullableWorkflowCustomDataTypeDefinitionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowCustomDataTypeDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCustomDataTypeDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
