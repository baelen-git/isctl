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

// SyslogPolicyAllOf Definition of the list of properties defined in 'syslog.Policy', excluding properties defined in parent classes.
type SyslogPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                                `json:"ObjectType"`
	LocalClients  []SyslogLocalClientBase               `json:"LocalClients,omitempty"`
	RemoteClients []SyslogRemoteClientBase              `json:"RemoteClients,omitempty"`
	Organization  *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SyslogPolicyAllOf SyslogPolicyAllOf

// NewSyslogPolicyAllOf instantiates a new SyslogPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogPolicyAllOf(classId string, objectType string) *SyslogPolicyAllOf {
	this := SyslogPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSyslogPolicyAllOfWithDefaults instantiates a new SyslogPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogPolicyAllOfWithDefaults() *SyslogPolicyAllOf {
	this := SyslogPolicyAllOf{}
	var classId string = "syslog.Policy"
	this.ClassId = classId
	var objectType string = "syslog.Policy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SyslogPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SyslogPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SyslogPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SyslogPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SyslogPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SyslogPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLocalClients returns the LocalClients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogPolicyAllOf) GetLocalClients() []SyslogLocalClientBase {
	if o == nil {
		var ret []SyslogLocalClientBase
		return ret
	}
	return o.LocalClients
}

// GetLocalClientsOk returns a tuple with the LocalClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogPolicyAllOf) GetLocalClientsOk() (*[]SyslogLocalClientBase, bool) {
	if o == nil || o.LocalClients == nil {
		return nil, false
	}
	return &o.LocalClients, true
}

// HasLocalClients returns a boolean if a field has been set.
func (o *SyslogPolicyAllOf) HasLocalClients() bool {
	if o != nil && o.LocalClients != nil {
		return true
	}

	return false
}

// SetLocalClients gets a reference to the given []SyslogLocalClientBase and assigns it to the LocalClients field.
func (o *SyslogPolicyAllOf) SetLocalClients(v []SyslogLocalClientBase) {
	o.LocalClients = v
}

// GetRemoteClients returns the RemoteClients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogPolicyAllOf) GetRemoteClients() []SyslogRemoteClientBase {
	if o == nil {
		var ret []SyslogRemoteClientBase
		return ret
	}
	return o.RemoteClients
}

// GetRemoteClientsOk returns a tuple with the RemoteClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogPolicyAllOf) GetRemoteClientsOk() (*[]SyslogRemoteClientBase, bool) {
	if o == nil || o.RemoteClients == nil {
		return nil, false
	}
	return &o.RemoteClients, true
}

// HasRemoteClients returns a boolean if a field has been set.
func (o *SyslogPolicyAllOf) HasRemoteClients() bool {
	if o != nil && o.RemoteClients != nil {
		return true
	}

	return false
}

// SetRemoteClients gets a reference to the given []SyslogRemoteClientBase and assigns it to the RemoteClients field.
func (o *SyslogPolicyAllOf) SetRemoteClients(v []SyslogRemoteClientBase) {
	o.RemoteClients = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SyslogPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SyslogPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SyslogPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SyslogPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *SyslogPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o SyslogPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LocalClients != nil {
		toSerialize["LocalClients"] = o.LocalClients
	}
	if o.RemoteClients != nil {
		toSerialize["RemoteClients"] = o.RemoteClients
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SyslogPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSyslogPolicyAllOf := _SyslogPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varSyslogPolicyAllOf); err == nil {
		*o = SyslogPolicyAllOf(varSyslogPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LocalClients")
		delete(additionalProperties, "RemoteClients")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSyslogPolicyAllOf struct {
	value *SyslogPolicyAllOf
	isSet bool
}

func (v NullableSyslogPolicyAllOf) Get() *SyslogPolicyAllOf {
	return v.value
}

func (v *NullableSyslogPolicyAllOf) Set(val *SyslogPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogPolicyAllOf(val *SyslogPolicyAllOf) *NullableSyslogPolicyAllOf {
	return &NullableSyslogPolicyAllOf{value: val, isSet: true}
}

func (v NullableSyslogPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
