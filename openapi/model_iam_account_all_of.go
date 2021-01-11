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

// IamAccountAllOf Definition of the list of properties defined in 'iam.Account', excluding properties defined in parent classes.
type IamAccountAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Name of the Intersight account. By default, name is same as the MoID of the account.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// Status of the account. To activate the Intersight account, claim a device to the account.
	Status *string `json:"Status,omitempty" yaml:"Status,omitempty"`
	// An array of relationships to iamAppRegistration resources.
	AppRegistrations []IamAppRegistrationRelationship `json:"AppRegistrations,omitempty" yaml:"AppRegistrations,omitempty"`
	// An array of relationships to iamDomainGroup resources.
	DomainGroups []IamDomainGroupRelationship `json:"DomainGroups,omitempty" yaml:"DomainGroups,omitempty"`
	// An array of relationships to iamEndPointRole resources.
	EndPointRoles []IamEndPointRoleRelationship `json:"EndPointRoles,omitempty" yaml:"EndPointRoles,omitempty"`
	// An array of relationships to iamIdpReference resources.
	Idpreferences []IamIdpReferenceRelationship `json:"Idpreferences,omitempty" yaml:"Idpreferences,omitempty"`
	// An array of relationships to iamIdp resources.
	Idps []IamIdpRelationship `json:"Idps,omitempty" yaml:"Idps,omitempty"`
	// An array of relationships to iamPermission resources.
	Permissions []IamPermissionRelationship `json:"Permissions,omitempty" yaml:"Permissions,omitempty"`
	// An array of relationships to iamPrivilegeSet resources.
	PrivilegeSets []IamPrivilegeSetRelationship `json:"PrivilegeSets,omitempty" yaml:"PrivilegeSets,omitempty"`
	// An array of relationships to iamPrivilege resources.
	Privileges     []IamPrivilegeRelationship     `json:"Privileges,omitempty" yaml:"Privileges,omitempty"`
	ResourceLimits *IamResourceLimitsRelationship `json:"ResourceLimits,omitempty" yaml:"ResourceLimits,omitempty"`
	// An array of relationships to iamRole resources.
	Roles          []IamRoleRelationship          `json:"Roles,omitempty" yaml:"Roles,omitempty"`
	SecurityHolder *IamSecurityHolderRelationship `json:"SecurityHolder,omitempty" yaml:"SecurityHolder,omitempty"`
	SessionLimits  *IamSessionLimitsRelationship  `json:"SessionLimits,omitempty" yaml:"SessionLimits,omitempty"`
}

// NewIamAccountAllOf instantiates a new IamAccountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamAccountAllOf(classId string, objectType string) *IamAccountAllOf {
	this := IamAccountAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamAccountAllOfWithDefaults instantiates a new IamAccountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamAccountAllOfWithDefaults() *IamAccountAllOf {
	this := IamAccountAllOf{}
	var classId string = "iam.Account"
	this.ClassId = classId
	var objectType string = "iam.Account"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamAccountAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamAccountAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamAccountAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamAccountAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamAccountAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamAccountAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamAccountAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccountAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamAccountAllOf) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IamAccountAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccountAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IamAccountAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetAppRegistrations returns the AppRegistrations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccountAllOf) GetAppRegistrations() []IamAppRegistrationRelationship {
	if o == nil {
		var ret []IamAppRegistrationRelationship
		return ret
	}
	return o.AppRegistrations
}

// GetAppRegistrationsOk returns a tuple with the AppRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccountAllOf) GetAppRegistrationsOk() (*[]IamAppRegistrationRelationship, bool) {
	if o == nil || o.AppRegistrations == nil {
		return nil, false
	}
	return &o.AppRegistrations, true
}

// HasAppRegistrations returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasAppRegistrations() bool {
	if o != nil && o.AppRegistrations != nil {
		return true
	}

	return false
}

// SetAppRegistrations gets a reference to the given []IamAppRegistrationRelationship and assigns it to the AppRegistrations field.
func (o *IamAccountAllOf) SetAppRegistrations(v []IamAppRegistrationRelationship) {
	o.AppRegistrations = v
}

// GetDomainGroups returns the DomainGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccountAllOf) GetDomainGroups() []IamDomainGroupRelationship {
	if o == nil {
		var ret []IamDomainGroupRelationship
		return ret
	}
	return o.DomainGroups
}

// GetDomainGroupsOk returns a tuple with the DomainGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccountAllOf) GetDomainGroupsOk() (*[]IamDomainGroupRelationship, bool) {
	if o == nil || o.DomainGroups == nil {
		return nil, false
	}
	return &o.DomainGroups, true
}

// HasDomainGroups returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasDomainGroups() bool {
	if o != nil && o.DomainGroups != nil {
		return true
	}

	return false
}

// SetDomainGroups gets a reference to the given []IamDomainGroupRelationship and assigns it to the DomainGroups field.
func (o *IamAccountAllOf) SetDomainGroups(v []IamDomainGroupRelationship) {
	o.DomainGroups = v
}

// GetEndPointRoles returns the EndPointRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccountAllOf) GetEndPointRoles() []IamEndPointRoleRelationship {
	if o == nil {
		var ret []IamEndPointRoleRelationship
		return ret
	}
	return o.EndPointRoles
}

// GetEndPointRolesOk returns a tuple with the EndPointRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccountAllOf) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool) {
	if o == nil || o.EndPointRoles == nil {
		return nil, false
	}
	return &o.EndPointRoles, true
}

// HasEndPointRoles returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasEndPointRoles() bool {
	if o != nil && o.EndPointRoles != nil {
		return true
	}

	return false
}

// SetEndPointRoles gets a reference to the given []IamEndPointRoleRelationship and assigns it to the EndPointRoles field.
func (o *IamAccountAllOf) SetEndPointRoles(v []IamEndPointRoleRelationship) {
	o.EndPointRoles = v
}

// GetIdpreferences returns the Idpreferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccountAllOf) GetIdpreferences() []IamIdpReferenceRelationship {
	if o == nil {
		var ret []IamIdpReferenceRelationship
		return ret
	}
	return o.Idpreferences
}

// GetIdpreferencesOk returns a tuple with the Idpreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccountAllOf) GetIdpreferencesOk() (*[]IamIdpReferenceRelationship, bool) {
	if o == nil || o.Idpreferences == nil {
		return nil, false
	}
	return &o.Idpreferences, true
}

// HasIdpreferences returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasIdpreferences() bool {
	if o != nil && o.Idpreferences != nil {
		return true
	}

	return false
}

// SetIdpreferences gets a reference to the given []IamIdpReferenceRelationship and assigns it to the Idpreferences field.
func (o *IamAccountAllOf) SetIdpreferences(v []IamIdpReferenceRelationship) {
	o.Idpreferences = v
}

// GetIdps returns the Idps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccountAllOf) GetIdps() []IamIdpRelationship {
	if o == nil {
		var ret []IamIdpRelationship
		return ret
	}
	return o.Idps
}

// GetIdpsOk returns a tuple with the Idps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccountAllOf) GetIdpsOk() (*[]IamIdpRelationship, bool) {
	if o == nil || o.Idps == nil {
		return nil, false
	}
	return &o.Idps, true
}

// HasIdps returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasIdps() bool {
	if o != nil && o.Idps != nil {
		return true
	}

	return false
}

// SetIdps gets a reference to the given []IamIdpRelationship and assigns it to the Idps field.
func (o *IamAccountAllOf) SetIdps(v []IamIdpRelationship) {
	o.Idps = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccountAllOf) GetPermissions() []IamPermissionRelationship {
	if o == nil {
		var ret []IamPermissionRelationship
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccountAllOf) GetPermissionsOk() (*[]IamPermissionRelationship, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []IamPermissionRelationship and assigns it to the Permissions field.
func (o *IamAccountAllOf) SetPermissions(v []IamPermissionRelationship) {
	o.Permissions = v
}

// GetPrivilegeSets returns the PrivilegeSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccountAllOf) GetPrivilegeSets() []IamPrivilegeSetRelationship {
	if o == nil {
		var ret []IamPrivilegeSetRelationship
		return ret
	}
	return o.PrivilegeSets
}

// GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccountAllOf) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool) {
	if o == nil || o.PrivilegeSets == nil {
		return nil, false
	}
	return &o.PrivilegeSets, true
}

// HasPrivilegeSets returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasPrivilegeSets() bool {
	if o != nil && o.PrivilegeSets != nil {
		return true
	}

	return false
}

// SetPrivilegeSets gets a reference to the given []IamPrivilegeSetRelationship and assigns it to the PrivilegeSets field.
func (o *IamAccountAllOf) SetPrivilegeSets(v []IamPrivilegeSetRelationship) {
	o.PrivilegeSets = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccountAllOf) GetPrivileges() []IamPrivilegeRelationship {
	if o == nil {
		var ret []IamPrivilegeRelationship
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccountAllOf) GetPrivilegesOk() (*[]IamPrivilegeRelationship, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []IamPrivilegeRelationship and assigns it to the Privileges field.
func (o *IamAccountAllOf) SetPrivileges(v []IamPrivilegeRelationship) {
	o.Privileges = v
}

// GetResourceLimits returns the ResourceLimits field value if set, zero value otherwise.
func (o *IamAccountAllOf) GetResourceLimits() IamResourceLimitsRelationship {
	if o == nil || o.ResourceLimits == nil {
		var ret IamResourceLimitsRelationship
		return ret
	}
	return *o.ResourceLimits
}

// GetResourceLimitsOk returns a tuple with the ResourceLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccountAllOf) GetResourceLimitsOk() (*IamResourceLimitsRelationship, bool) {
	if o == nil || o.ResourceLimits == nil {
		return nil, false
	}
	return o.ResourceLimits, true
}

// HasResourceLimits returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasResourceLimits() bool {
	if o != nil && o.ResourceLimits != nil {
		return true
	}

	return false
}

// SetResourceLimits gets a reference to the given IamResourceLimitsRelationship and assigns it to the ResourceLimits field.
func (o *IamAccountAllOf) SetResourceLimits(v IamResourceLimitsRelationship) {
	o.ResourceLimits = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAccountAllOf) GetRoles() []IamRoleRelationship {
	if o == nil {
		var ret []IamRoleRelationship
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAccountAllOf) GetRolesOk() (*[]IamRoleRelationship, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []IamRoleRelationship and assigns it to the Roles field.
func (o *IamAccountAllOf) SetRoles(v []IamRoleRelationship) {
	o.Roles = v
}

// GetSecurityHolder returns the SecurityHolder field value if set, zero value otherwise.
func (o *IamAccountAllOf) GetSecurityHolder() IamSecurityHolderRelationship {
	if o == nil || o.SecurityHolder == nil {
		var ret IamSecurityHolderRelationship
		return ret
	}
	return *o.SecurityHolder
}

// GetSecurityHolderOk returns a tuple with the SecurityHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccountAllOf) GetSecurityHolderOk() (*IamSecurityHolderRelationship, bool) {
	if o == nil || o.SecurityHolder == nil {
		return nil, false
	}
	return o.SecurityHolder, true
}

// HasSecurityHolder returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasSecurityHolder() bool {
	if o != nil && o.SecurityHolder != nil {
		return true
	}

	return false
}

// SetSecurityHolder gets a reference to the given IamSecurityHolderRelationship and assigns it to the SecurityHolder field.
func (o *IamAccountAllOf) SetSecurityHolder(v IamSecurityHolderRelationship) {
	o.SecurityHolder = &v
}

// GetSessionLimits returns the SessionLimits field value if set, zero value otherwise.
func (o *IamAccountAllOf) GetSessionLimits() IamSessionLimitsRelationship {
	if o == nil || o.SessionLimits == nil {
		var ret IamSessionLimitsRelationship
		return ret
	}
	return *o.SessionLimits
}

// GetSessionLimitsOk returns a tuple with the SessionLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccountAllOf) GetSessionLimitsOk() (*IamSessionLimitsRelationship, bool) {
	if o == nil || o.SessionLimits == nil {
		return nil, false
	}
	return o.SessionLimits, true
}

// HasSessionLimits returns a boolean if a field has been set.
func (o *IamAccountAllOf) HasSessionLimits() bool {
	if o != nil && o.SessionLimits != nil {
		return true
	}

	return false
}

// SetSessionLimits gets a reference to the given IamSessionLimitsRelationship and assigns it to the SessionLimits field.
func (o *IamAccountAllOf) SetSessionLimits(v IamSessionLimitsRelationship) {
	o.SessionLimits = &v
}

func (o IamAccountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.AppRegistrations != nil {
		toSerialize["AppRegistrations"] = o.AppRegistrations
	}
	if o.DomainGroups != nil {
		toSerialize["DomainGroups"] = o.DomainGroups
	}
	if o.EndPointRoles != nil {
		toSerialize["EndPointRoles"] = o.EndPointRoles
	}
	if o.Idpreferences != nil {
		toSerialize["Idpreferences"] = o.Idpreferences
	}
	if o.Idps != nil {
		toSerialize["Idps"] = o.Idps
	}
	if o.Permissions != nil {
		toSerialize["Permissions"] = o.Permissions
	}
	if o.PrivilegeSets != nil {
		toSerialize["PrivilegeSets"] = o.PrivilegeSets
	}
	if o.Privileges != nil {
		toSerialize["Privileges"] = o.Privileges
	}
	if o.ResourceLimits != nil {
		toSerialize["ResourceLimits"] = o.ResourceLimits
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}
	if o.SecurityHolder != nil {
		toSerialize["SecurityHolder"] = o.SecurityHolder
	}
	if o.SessionLimits != nil {
		toSerialize["SessionLimits"] = o.SessionLimits
	}
	return json.Marshal(toSerialize)
}

type NullableIamAccountAllOf struct {
	value *IamAccountAllOf
	isSet bool
}

func (v NullableIamAccountAllOf) Get() *IamAccountAllOf {
	return v.value
}

func (v *NullableIamAccountAllOf) Set(val *IamAccountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamAccountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamAccountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamAccountAllOf(val *IamAccountAllOf) *NullableIamAccountAllOf {
	return &NullableIamAccountAllOf{value: val, isSet: true}
}

func (v NullableIamAccountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamAccountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
