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

// TechsupportmanagementApplianceParam Appliance parameter object for Tech Support requests.
type TechsupportmanagementApplianceParam struct {
	ConnectorPlatformParamBase `yaml:"ConnectorPlatformParamBase,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Specifies whether the techsupport request is from the cloud or by the appliance.
	IsApplianceRequest *bool `json:"IsApplianceRequest,omitempty" yaml:"IsApplianceRequest,omitempty"`
}

// NewTechsupportmanagementApplianceParam instantiates a new TechsupportmanagementApplianceParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementApplianceParam(classId string, objectType string) *TechsupportmanagementApplianceParam {
	this := TechsupportmanagementApplianceParam{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTechsupportmanagementApplianceParamWithDefaults instantiates a new TechsupportmanagementApplianceParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementApplianceParamWithDefaults() *TechsupportmanagementApplianceParam {
	this := TechsupportmanagementApplianceParam{}
	var classId string = "techsupportmanagement.ApplianceParam"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.ApplianceParam"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementApplianceParam) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementApplianceParam) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementApplianceParam) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementApplianceParam) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementApplianceParam) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementApplianceParam) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsApplianceRequest returns the IsApplianceRequest field value if set, zero value otherwise.
func (o *TechsupportmanagementApplianceParam) GetIsApplianceRequest() bool {
	if o == nil || o.IsApplianceRequest == nil {
		var ret bool
		return ret
	}
	return *o.IsApplianceRequest
}

// GetIsApplianceRequestOk returns a tuple with the IsApplianceRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementApplianceParam) GetIsApplianceRequestOk() (*bool, bool) {
	if o == nil || o.IsApplianceRequest == nil {
		return nil, false
	}
	return o.IsApplianceRequest, true
}

// HasIsApplianceRequest returns a boolean if a field has been set.
func (o *TechsupportmanagementApplianceParam) HasIsApplianceRequest() bool {
	if o != nil && o.IsApplianceRequest != nil {
		return true
	}

	return false
}

// SetIsApplianceRequest gets a reference to the given bool and assigns it to the IsApplianceRequest field.
func (o *TechsupportmanagementApplianceParam) SetIsApplianceRequest(v bool) {
	o.IsApplianceRequest = &v
}

func (o TechsupportmanagementApplianceParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorPlatformParamBase, errConnectorPlatformParamBase := json.Marshal(o.ConnectorPlatformParamBase)
	if errConnectorPlatformParamBase != nil {
		return []byte{}, errConnectorPlatformParamBase
	}
	errConnectorPlatformParamBase = json.Unmarshal([]byte(serializedConnectorPlatformParamBase), &toSerialize)
	if errConnectorPlatformParamBase != nil {
		return []byte{}, errConnectorPlatformParamBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsApplianceRequest != nil {
		toSerialize["IsApplianceRequest"] = o.IsApplianceRequest
	}
	return json.Marshal(toSerialize)
}

type NullableTechsupportmanagementApplianceParam struct {
	value *TechsupportmanagementApplianceParam
	isSet bool
}

func (v NullableTechsupportmanagementApplianceParam) Get() *TechsupportmanagementApplianceParam {
	return v.value
}

func (v *NullableTechsupportmanagementApplianceParam) Set(val *TechsupportmanagementApplianceParam) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementApplianceParam) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementApplianceParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementApplianceParam(val *TechsupportmanagementApplianceParam) *NullableTechsupportmanagementApplianceParam {
	return &NullableTechsupportmanagementApplianceParam{value: val, isSet: true}
}

func (v NullableTechsupportmanagementApplianceParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementApplianceParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
