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

// HyperflexVcenterConfigPolicyAllOf Definition of the list of properties defined in 'hyperflex.VcenterConfigPolicy', excluding properties defined in parent classes.
type HyperflexVcenterConfigPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The vCenter datacenter name.
	DataCenter *string `json:"DataCenter,omitempty" yaml:"DataCenter,omitempty"`
	// The vCenter server FQDN or IP.
	Hostname *string `json:"Hostname,omitempty" yaml:"Hostname,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty" yaml:"IsPasswordSet,omitempty"`
	// The password for authenticating with vCenter. Follow the corresponding password policy governed by vCenter.
	Password *string `json:"Password,omitempty" yaml:"Password,omitempty"`
	// Overrides the default vCenter Single Sign-On URL. Do not specify unless instructed by Cisco TAC.
	SsoUrl *string `json:"SsoUrl,omitempty" yaml:"SsoUrl,omitempty"`
	// The vCenter username (e.g. administrator@vsphere.local).
	Username *string `json:"Username,omitempty" yaml:"Username,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty" yaml:"ClusterProfiles,omitempty"`
	Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
}

// NewHyperflexVcenterConfigPolicyAllOf instantiates a new HyperflexVcenterConfigPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVcenterConfigPolicyAllOf(classId string, objectType string) *HyperflexVcenterConfigPolicyAllOf {
	this := HyperflexVcenterConfigPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// NewHyperflexVcenterConfigPolicyAllOfWithDefaults instantiates a new HyperflexVcenterConfigPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVcenterConfigPolicyAllOfWithDefaults() *HyperflexVcenterConfigPolicyAllOf {
	this := HyperflexVcenterConfigPolicyAllOf{}
	var classId string = "hyperflex.VcenterConfigPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.VcenterConfigPolicy"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVcenterConfigPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVcenterConfigPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVcenterConfigPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVcenterConfigPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataCenter returns the DataCenter field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicyAllOf) GetDataCenter() string {
	if o == nil || o.DataCenter == nil {
		var ret string
		return ret
	}
	return *o.DataCenter
}

// GetDataCenterOk returns a tuple with the DataCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) GetDataCenterOk() (*string, bool) {
	if o == nil || o.DataCenter == nil {
		return nil, false
	}
	return o.DataCenter, true
}

// HasDataCenter returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) HasDataCenter() bool {
	if o != nil && o.DataCenter != nil {
		return true
	}

	return false
}

// SetDataCenter gets a reference to the given string and assigns it to the DataCenter field.
func (o *HyperflexVcenterConfigPolicyAllOf) SetDataCenter(v string) {
	o.DataCenter = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicyAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *HyperflexVcenterConfigPolicyAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicyAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *HyperflexVcenterConfigPolicyAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicyAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *HyperflexVcenterConfigPolicyAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicyAllOf) GetSsoUrl() string {
	if o == nil || o.SsoUrl == nil {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) GetSsoUrlOk() (*string, bool) {
	if o == nil || o.SsoUrl == nil {
		return nil, false
	}
	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) HasSsoUrl() bool {
	if o != nil && o.SsoUrl != nil {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *HyperflexVcenterConfigPolicyAllOf) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicyAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *HyperflexVcenterConfigPolicyAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVcenterConfigPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVcenterConfigPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexVcenterConfigPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexVcenterConfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexVcenterConfigPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexVcenterConfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexVcenterConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DataCenter != nil {
		toSerialize["DataCenter"] = o.DataCenter
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.SsoUrl != nil {
		toSerialize["SsoUrl"] = o.SsoUrl
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexVcenterConfigPolicyAllOf struct {
	value *HyperflexVcenterConfigPolicyAllOf
	isSet bool
}

func (v NullableHyperflexVcenterConfigPolicyAllOf) Get() *HyperflexVcenterConfigPolicyAllOf {
	return v.value
}

func (v *NullableHyperflexVcenterConfigPolicyAllOf) Set(val *HyperflexVcenterConfigPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVcenterConfigPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVcenterConfigPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVcenterConfigPolicyAllOf(val *HyperflexVcenterConfigPolicyAllOf) *NullableHyperflexVcenterConfigPolicyAllOf {
	return &NullableHyperflexVcenterConfigPolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexVcenterConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVcenterConfigPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
