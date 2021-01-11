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

// DeviceconnectorPolicyAllOf Definition of the list of properties defined in 'deviceconnector.Policy', excluding properties defined in parent classes.
type DeviceconnectorPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Enables configuration lockout on the endpoint.
	LockoutEnabled *bool                                 `json:"LockoutEnabled,omitempty" yaml:"LockoutEnabled,omitempty"`
	Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty" yaml:"Profiles,omitempty"`
}

// NewDeviceconnectorPolicyAllOf instantiates a new DeviceconnectorPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceconnectorPolicyAllOf(classId string, objectType string) *DeviceconnectorPolicyAllOf {
	this := DeviceconnectorPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var lockoutEnabled bool = true
	this.LockoutEnabled = &lockoutEnabled
	return &this
}

// NewDeviceconnectorPolicyAllOfWithDefaults instantiates a new DeviceconnectorPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceconnectorPolicyAllOfWithDefaults() *DeviceconnectorPolicyAllOf {
	this := DeviceconnectorPolicyAllOf{}
	var classId string = "deviceconnector.Policy"
	this.ClassId = classId
	var objectType string = "deviceconnector.Policy"
	this.ObjectType = objectType
	var lockoutEnabled bool = true
	this.LockoutEnabled = &lockoutEnabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *DeviceconnectorPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *DeviceconnectorPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *DeviceconnectorPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *DeviceconnectorPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DeviceconnectorPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DeviceconnectorPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLockoutEnabled returns the LockoutEnabled field value if set, zero value otherwise.
func (o *DeviceconnectorPolicyAllOf) GetLockoutEnabled() bool {
	if o == nil || o.LockoutEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LockoutEnabled
}

// GetLockoutEnabledOk returns a tuple with the LockoutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceconnectorPolicyAllOf) GetLockoutEnabledOk() (*bool, bool) {
	if o == nil || o.LockoutEnabled == nil {
		return nil, false
	}
	return o.LockoutEnabled, true
}

// HasLockoutEnabled returns a boolean if a field has been set.
func (o *DeviceconnectorPolicyAllOf) HasLockoutEnabled() bool {
	if o != nil && o.LockoutEnabled != nil {
		return true
	}

	return false
}

// SetLockoutEnabled gets a reference to the given bool and assigns it to the LockoutEnabled field.
func (o *DeviceconnectorPolicyAllOf) SetLockoutEnabled(v bool) {
	o.LockoutEnabled = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *DeviceconnectorPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceconnectorPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *DeviceconnectorPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *DeviceconnectorPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceconnectorPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceconnectorPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *DeviceconnectorPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *DeviceconnectorPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o DeviceconnectorPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LockoutEnabled != nil {
		toSerialize["LockoutEnabled"] = o.LockoutEnabled
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceconnectorPolicyAllOf struct {
	value *DeviceconnectorPolicyAllOf
	isSet bool
}

func (v NullableDeviceconnectorPolicyAllOf) Get() *DeviceconnectorPolicyAllOf {
	return v.value
}

func (v *NullableDeviceconnectorPolicyAllOf) Set(val *DeviceconnectorPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceconnectorPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceconnectorPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceconnectorPolicyAllOf(val *DeviceconnectorPolicyAllOf) *NullableDeviceconnectorPolicyAllOf {
	return &NullableDeviceconnectorPolicyAllOf{value: val, isSet: true}
}

func (v NullableDeviceconnectorPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceconnectorPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
