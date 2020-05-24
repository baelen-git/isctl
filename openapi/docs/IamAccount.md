# IamAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Intersight account. By default, name is same as the MoID of the account. | [optional] 
**Status** | Pointer to **string** | Status of the account. To activate the Intersight account, claim a device to the account. | [optional] [readonly] 
**AppRegistrations** | Pointer to [**[]IamAppRegistrationRelationship**](iam.AppRegistration.Relationship.md) | An array of relationships to iamAppRegistration resources. | [optional] [readonly] 
**DomainGroups** | Pointer to [**[]IamDomainGroupRelationship**](iam.DomainGroup.Relationship.md) | An array of relationships to iamDomainGroup resources. | [optional] [readonly] 
**EndPointRoles** | Pointer to [**[]IamEndPointRoleRelationship**](iam.EndPointRole.Relationship.md) | An array of relationships to iamEndPointRole resources. | [optional] [readonly] 
**Idpreferences** | Pointer to [**[]IamIdpReferenceRelationship**](iam.IdpReference.Relationship.md) | An array of relationships to iamIdpReference resources. | [optional] [readonly] 
**Idps** | Pointer to [**[]IamIdpRelationship**](iam.Idp.Relationship.md) | An array of relationships to iamIdp resources. | [optional] [readonly] 
**Permissions** | Pointer to [**[]IamPermissionRelationship**](iam.Permission.Relationship.md) | An array of relationships to iamPermission resources. | [optional] [readonly] 
**PrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](iam.PrivilegeSet.Relationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] [readonly] 
**Privileges** | Pointer to [**[]IamPrivilegeRelationship**](iam.Privilege.Relationship.md) | An array of relationships to iamPrivilege resources. | [optional] [readonly] 
**ResourceLimits** | Pointer to [**IamResourceLimitsRelationship**](iam.ResourceLimits.Relationship.md) |  | [optional] 
**Roles** | Pointer to [**[]IamRoleRelationship**](iam.Role.Relationship.md) | An array of relationships to iamRole resources. | [optional] [readonly] 
**SecurityHolder** | Pointer to [**IamSecurityHolderRelationship**](iam.SecurityHolder.Relationship.md) |  | [optional] 
**SessionLimits** | Pointer to [**IamSessionLimitsRelationship**](iam.SessionLimits.Relationship.md) |  | [optional] 

## Methods

### NewIamAccount

`func NewIamAccount() *IamAccount`

NewIamAccount instantiates a new IamAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAccountWithDefaults

`func NewIamAccountWithDefaults() *IamAccount`

NewIamAccountWithDefaults instantiates a new IamAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *IamAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAppRegistrations

`func (o *IamAccount) GetAppRegistrations() []IamAppRegistrationRelationship`

GetAppRegistrations returns the AppRegistrations field if non-nil, zero value otherwise.

### GetAppRegistrationsOk

`func (o *IamAccount) GetAppRegistrationsOk() (*[]IamAppRegistrationRelationship, bool)`

GetAppRegistrationsOk returns a tuple with the AppRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRegistrations

`func (o *IamAccount) SetAppRegistrations(v []IamAppRegistrationRelationship)`

SetAppRegistrations sets AppRegistrations field to given value.

### HasAppRegistrations

`func (o *IamAccount) HasAppRegistrations() bool`

HasAppRegistrations returns a boolean if a field has been set.

### GetDomainGroups

`func (o *IamAccount) GetDomainGroups() []IamDomainGroupRelationship`

GetDomainGroups returns the DomainGroups field if non-nil, zero value otherwise.

### GetDomainGroupsOk

`func (o *IamAccount) GetDomainGroupsOk() (*[]IamDomainGroupRelationship, bool)`

GetDomainGroupsOk returns a tuple with the DomainGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroups

`func (o *IamAccount) SetDomainGroups(v []IamDomainGroupRelationship)`

SetDomainGroups sets DomainGroups field to given value.

### HasDomainGroups

`func (o *IamAccount) HasDomainGroups() bool`

HasDomainGroups returns a boolean if a field has been set.

### GetEndPointRoles

`func (o *IamAccount) GetEndPointRoles() []IamEndPointRoleRelationship`

GetEndPointRoles returns the EndPointRoles field if non-nil, zero value otherwise.

### GetEndPointRolesOk

`func (o *IamAccount) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool)`

GetEndPointRolesOk returns a tuple with the EndPointRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointRoles

`func (o *IamAccount) SetEndPointRoles(v []IamEndPointRoleRelationship)`

SetEndPointRoles sets EndPointRoles field to given value.

### HasEndPointRoles

`func (o *IamAccount) HasEndPointRoles() bool`

HasEndPointRoles returns a boolean if a field has been set.

### GetIdpreferences

`func (o *IamAccount) GetIdpreferences() []IamIdpReferenceRelationship`

GetIdpreferences returns the Idpreferences field if non-nil, zero value otherwise.

### GetIdpreferencesOk

`func (o *IamAccount) GetIdpreferencesOk() (*[]IamIdpReferenceRelationship, bool)`

GetIdpreferencesOk returns a tuple with the Idpreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpreferences

`func (o *IamAccount) SetIdpreferences(v []IamIdpReferenceRelationship)`

SetIdpreferences sets Idpreferences field to given value.

### HasIdpreferences

`func (o *IamAccount) HasIdpreferences() bool`

HasIdpreferences returns a boolean if a field has been set.

### GetIdps

`func (o *IamAccount) GetIdps() []IamIdpRelationship`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *IamAccount) GetIdpsOk() (*[]IamIdpRelationship, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *IamAccount) SetIdps(v []IamIdpRelationship)`

SetIdps sets Idps field to given value.

### HasIdps

`func (o *IamAccount) HasIdps() bool`

HasIdps returns a boolean if a field has been set.

### GetPermissions

`func (o *IamAccount) GetPermissions() []IamPermissionRelationship`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamAccount) GetPermissionsOk() (*[]IamPermissionRelationship, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamAccount) SetPermissions(v []IamPermissionRelationship)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamAccount) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPrivilegeSets

`func (o *IamAccount) GetPrivilegeSets() []IamPrivilegeSetRelationship`

GetPrivilegeSets returns the PrivilegeSets field if non-nil, zero value otherwise.

### GetPrivilegeSetsOk

`func (o *IamAccount) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSets

`func (o *IamAccount) SetPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetPrivilegeSets sets PrivilegeSets field to given value.

### HasPrivilegeSets

`func (o *IamAccount) HasPrivilegeSets() bool`

HasPrivilegeSets returns a boolean if a field has been set.

### GetPrivileges

`func (o *IamAccount) GetPrivileges() []IamPrivilegeRelationship`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *IamAccount) GetPrivilegesOk() (*[]IamPrivilegeRelationship, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *IamAccount) SetPrivileges(v []IamPrivilegeRelationship)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *IamAccount) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetResourceLimits

`func (o *IamAccount) GetResourceLimits() IamResourceLimitsRelationship`

GetResourceLimits returns the ResourceLimits field if non-nil, zero value otherwise.

### GetResourceLimitsOk

`func (o *IamAccount) GetResourceLimitsOk() (*IamResourceLimitsRelationship, bool)`

GetResourceLimitsOk returns a tuple with the ResourceLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLimits

`func (o *IamAccount) SetResourceLimits(v IamResourceLimitsRelationship)`

SetResourceLimits sets ResourceLimits field to given value.

### HasResourceLimits

`func (o *IamAccount) HasResourceLimits() bool`

HasResourceLimits returns a boolean if a field has been set.

### GetRoles

`func (o *IamAccount) GetRoles() []IamRoleRelationship`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamAccount) GetRolesOk() (*[]IamRoleRelationship, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamAccount) SetRoles(v []IamRoleRelationship)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamAccount) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSecurityHolder

`func (o *IamAccount) GetSecurityHolder() IamSecurityHolderRelationship`

GetSecurityHolder returns the SecurityHolder field if non-nil, zero value otherwise.

### GetSecurityHolderOk

`func (o *IamAccount) GetSecurityHolderOk() (*IamSecurityHolderRelationship, bool)`

GetSecurityHolderOk returns a tuple with the SecurityHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityHolder

`func (o *IamAccount) SetSecurityHolder(v IamSecurityHolderRelationship)`

SetSecurityHolder sets SecurityHolder field to given value.

### HasSecurityHolder

`func (o *IamAccount) HasSecurityHolder() bool`

HasSecurityHolder returns a boolean if a field has been set.

### GetSessionLimits

`func (o *IamAccount) GetSessionLimits() IamSessionLimitsRelationship`

GetSessionLimits returns the SessionLimits field if non-nil, zero value otherwise.

### GetSessionLimitsOk

`func (o *IamAccount) GetSessionLimitsOk() (*IamSessionLimitsRelationship, bool)`

GetSessionLimitsOk returns a tuple with the SessionLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLimits

`func (o *IamAccount) SetSessionLimits(v IamSessionLimitsRelationship)`

SetSessionLimits sets SessionLimits field to given value.

### HasSessionLimits

`func (o *IamAccount) HasSessionLimits() bool`

HasSessionLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


