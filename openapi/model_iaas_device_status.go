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

// IaasDeviceStatus List of infra accounts managed by UCSD.
type IaasDeviceStatus struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD.
	AccountName *string `json:"AccountName,omitempty" yaml:"AccountName,omitempty"`
	// The UCSD Infra Account type.
	AccountType *string `json:"AccountType,omitempty" yaml:"AccountType,omitempty"`
	// The UCSD Infra Account category.
	Category *string `json:"Category,omitempty" yaml:"Category,omitempty"`
	// Describes if the device is claimed in Intersight or not. * `Unknown` - If UCS Director managed account claim status information is unknown. * `Yes` - If UCS Director managed account is claimed in Intersight. * `No` - If UCS Director managed account is not claimed in Intersight. * `Not Applicable` - If UCS Director managed account is not capable of providing claim status information.
	ClaimStatus *string `json:"ClaimStatus,omitempty" yaml:"ClaimStatus,omitempty"`
	// Describes about the connection status between the UCSD and the actual end device.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" yaml:"ConnectionStatus,omitempty"`
	// Describes about the device model.
	DeviceModel *string `json:"DeviceModel,omitempty" yaml:"DeviceModel,omitempty"`
	// Describes about the device vendor/manufacturer of the device.
	DeviceVendor *string `json:"DeviceVendor,omitempty" yaml:"DeviceVendor,omitempty"`
	// Describes about the current firmware version running on the device.
	DeviceVersion *string `json:"DeviceVersion,omitempty" yaml:"DeviceVersion,omitempty"`
	// The IPAddress of the device.
	IpAddress *string `json:"IpAddress,omitempty" yaml:"IpAddress,omitempty"`
	// Describes about the pod to which this device belongs to in UCSD.
	Pod *string `json:"Pod,omitempty" yaml:"Pod,omitempty"`
	// Describes about the podType of Pod to which this device belongs to in UCSD.
	PodType *string                   `json:"PodType,omitempty" yaml:"PodType,omitempty"`
	Guid    *IaasUcsdInfoRelationship `json:"Guid,omitempty" yaml:"Guid,omitempty"`
}

// NewIaasDeviceStatus instantiates a new IaasDeviceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasDeviceStatus(classId string, objectType string) *IaasDeviceStatus {
	this := IaasDeviceStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var claimStatus string = "Unknown"
	this.ClaimStatus = &claimStatus
	return &this
}

// NewIaasDeviceStatusWithDefaults instantiates a new IaasDeviceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasDeviceStatusWithDefaults() *IaasDeviceStatus {
	this := IaasDeviceStatus{}
	var classId string = "iaas.DeviceStatus"
	this.ClassId = classId
	var objectType string = "iaas.DeviceStatus"
	this.ObjectType = objectType
	var claimStatus string = "Unknown"
	this.ClaimStatus = &claimStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasDeviceStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasDeviceStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasDeviceStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasDeviceStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *IaasDeviceStatus) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetAccountTypeOk() (*string, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *IaasDeviceStatus) SetAccountType(v string) {
	o.AccountType = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *IaasDeviceStatus) SetCategory(v string) {
	o.Category = &v
}

// GetClaimStatus returns the ClaimStatus field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetClaimStatus() string {
	if o == nil || o.ClaimStatus == nil {
		var ret string
		return ret
	}
	return *o.ClaimStatus
}

// GetClaimStatusOk returns a tuple with the ClaimStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetClaimStatusOk() (*string, bool) {
	if o == nil || o.ClaimStatus == nil {
		return nil, false
	}
	return o.ClaimStatus, true
}

// HasClaimStatus returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasClaimStatus() bool {
	if o != nil && o.ClaimStatus != nil {
		return true
	}

	return false
}

// SetClaimStatus gets a reference to the given string and assigns it to the ClaimStatus field.
func (o *IaasDeviceStatus) SetClaimStatus(v string) {
	o.ClaimStatus = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *IaasDeviceStatus) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetDeviceModel() string {
	if o == nil || o.DeviceModel == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetDeviceModelOk() (*string, bool) {
	if o == nil || o.DeviceModel == nil {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasDeviceModel() bool {
	if o != nil && o.DeviceModel != nil {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *IaasDeviceStatus) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetDeviceVendor returns the DeviceVendor field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetDeviceVendor() string {
	if o == nil || o.DeviceVendor == nil {
		var ret string
		return ret
	}
	return *o.DeviceVendor
}

// GetDeviceVendorOk returns a tuple with the DeviceVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetDeviceVendorOk() (*string, bool) {
	if o == nil || o.DeviceVendor == nil {
		return nil, false
	}
	return o.DeviceVendor, true
}

// HasDeviceVendor returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasDeviceVendor() bool {
	if o != nil && o.DeviceVendor != nil {
		return true
	}

	return false
}

// SetDeviceVendor gets a reference to the given string and assigns it to the DeviceVendor field.
func (o *IaasDeviceStatus) SetDeviceVendor(v string) {
	o.DeviceVendor = &v
}

// GetDeviceVersion returns the DeviceVersion field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetDeviceVersion() string {
	if o == nil || o.DeviceVersion == nil {
		var ret string
		return ret
	}
	return *o.DeviceVersion
}

// GetDeviceVersionOk returns a tuple with the DeviceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetDeviceVersionOk() (*string, bool) {
	if o == nil || o.DeviceVersion == nil {
		return nil, false
	}
	return o.DeviceVersion, true
}

// HasDeviceVersion returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasDeviceVersion() bool {
	if o != nil && o.DeviceVersion != nil {
		return true
	}

	return false
}

// SetDeviceVersion gets a reference to the given string and assigns it to the DeviceVersion field.
func (o *IaasDeviceStatus) SetDeviceVersion(v string) {
	o.DeviceVersion = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *IaasDeviceStatus) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetPod() string {
	if o == nil || o.Pod == nil {
		var ret string
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetPodOk() (*string, bool) {
	if o == nil || o.Pod == nil {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasPod() bool {
	if o != nil && o.Pod != nil {
		return true
	}

	return false
}

// SetPod gets a reference to the given string and assigns it to the Pod field.
func (o *IaasDeviceStatus) SetPod(v string) {
	o.Pod = &v
}

// GetPodType returns the PodType field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetPodType() string {
	if o == nil || o.PodType == nil {
		var ret string
		return ret
	}
	return *o.PodType
}

// GetPodTypeOk returns a tuple with the PodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetPodTypeOk() (*string, bool) {
	if o == nil || o.PodType == nil {
		return nil, false
	}
	return o.PodType, true
}

// HasPodType returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasPodType() bool {
	if o != nil && o.PodType != nil {
		return true
	}

	return false
}

// SetPodType gets a reference to the given string and assigns it to the PodType field.
func (o *IaasDeviceStatus) SetPodType(v string) {
	o.PodType = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasDeviceStatus) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || o.Guid == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasDeviceStatus) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasDeviceStatus) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given IaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasDeviceStatus) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid = &v
}

func (o IaasDeviceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccountName != nil {
		toSerialize["AccountName"] = o.AccountName
	}
	if o.AccountType != nil {
		toSerialize["AccountType"] = o.AccountType
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.ClaimStatus != nil {
		toSerialize["ClaimStatus"] = o.ClaimStatus
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.DeviceModel != nil {
		toSerialize["DeviceModel"] = o.DeviceModel
	}
	if o.DeviceVendor != nil {
		toSerialize["DeviceVendor"] = o.DeviceVendor
	}
	if o.DeviceVersion != nil {
		toSerialize["DeviceVersion"] = o.DeviceVersion
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Pod != nil {
		toSerialize["Pod"] = o.Pod
	}
	if o.PodType != nil {
		toSerialize["PodType"] = o.PodType
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}
	return json.Marshal(toSerialize)
}

type NullableIaasDeviceStatus struct {
	value *IaasDeviceStatus
	isSet bool
}

func (v NullableIaasDeviceStatus) Get() *IaasDeviceStatus {
	return v.value
}

func (v *NullableIaasDeviceStatus) Set(val *IaasDeviceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasDeviceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasDeviceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasDeviceStatus(val *IaasDeviceStatus) *NullableIaasDeviceStatus {
	return &NullableIaasDeviceStatus{value: val, isSet: true}
}

func (v NullableIaasDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasDeviceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
