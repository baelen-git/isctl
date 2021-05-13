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

// FabricLinkControlPolicy A policy to configure the link settings for all the ports (including UDLD).
type FabricLinkControlPolicy struct {
	PolicyAbstractPolicy `yaml:"PolicyAbstractPolicy,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                                `json:"ObjectType" yaml:"ObjectType"`
	UdldSettings NullableFabricUdldSettings            `json:"UdldSettings,omitempty" yaml:"UdldSettings,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
}

// NewFabricLinkControlPolicy instantiates a new FabricLinkControlPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricLinkControlPolicy(classId string, objectType string) *FabricLinkControlPolicy {
	this := FabricLinkControlPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricLinkControlPolicyWithDefaults instantiates a new FabricLinkControlPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricLinkControlPolicyWithDefaults() *FabricLinkControlPolicy {
	this := FabricLinkControlPolicy{}
	var classId string = "fabric.LinkControlPolicy"
	this.ClassId = classId
	var objectType string = "fabric.LinkControlPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricLinkControlPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricLinkControlPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricLinkControlPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricLinkControlPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricLinkControlPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricLinkControlPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUdldSettings returns the UdldSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricLinkControlPolicy) GetUdldSettings() FabricUdldSettings {
	if o == nil || o.UdldSettings.Get() == nil {
		var ret FabricUdldSettings
		return ret
	}
	return *o.UdldSettings.Get()
}

// GetUdldSettingsOk returns a tuple with the UdldSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricLinkControlPolicy) GetUdldSettingsOk() (*FabricUdldSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.UdldSettings.Get(), o.UdldSettings.IsSet()
}

// HasUdldSettings returns a boolean if a field has been set.
func (o *FabricLinkControlPolicy) HasUdldSettings() bool {
	if o != nil && o.UdldSettings.IsSet() {
		return true
	}

	return false
}

// SetUdldSettings gets a reference to the given NullableFabricUdldSettings and assigns it to the UdldSettings field.
func (o *FabricLinkControlPolicy) SetUdldSettings(v FabricUdldSettings) {
	o.UdldSettings.Set(&v)
}

// SetUdldSettingsNil sets the value for UdldSettings to be an explicit nil
func (o *FabricLinkControlPolicy) SetUdldSettingsNil() {
	o.UdldSettings.Set(nil)
}

// UnsetUdldSettings ensures that no value is present for UdldSettings, not even an explicit nil
func (o *FabricLinkControlPolicy) UnsetUdldSettings() {
	o.UdldSettings.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricLinkControlPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLinkControlPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricLinkControlPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricLinkControlPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o FabricLinkControlPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.UdldSettings.IsSet() {
		toSerialize["UdldSettings"] = o.UdldSettings.Get()
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableFabricLinkControlPolicy struct {
	value *FabricLinkControlPolicy
	isSet bool
}

func (v NullableFabricLinkControlPolicy) Get() *FabricLinkControlPolicy {
	return v.value
}

func (v *NullableFabricLinkControlPolicy) Set(val *FabricLinkControlPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricLinkControlPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricLinkControlPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricLinkControlPolicy(val *FabricLinkControlPolicy) *NullableFabricLinkControlPolicy {
	return &NullableFabricLinkControlPolicy{value: val, isSet: true}
}

func (v NullableFabricLinkControlPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricLinkControlPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
