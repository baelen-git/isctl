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
	"time"
)

// IamIpAccessManagement The access management based on IP address.
type IamIpAccessManagement struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Flag stores the state of IP address based access management. Access management is enabled when it's true.
	Enable *bool `json:"Enable,omitempty" yaml:"Enable,omitempty"`
	// The access to account gets locked out if wrong IP addresses are configured. Account Administrators have privilege to unblock the account. It stores the time when the account was last recovered from lock out.
	LastRecoveryTime *time.Time                     `json:"LastRecoveryTime,omitempty" yaml:"LastRecoveryTime,omitempty"`
	Holder           *IamSecurityHolderRelationship `json:"Holder,omitempty" yaml:"Holder,omitempty"`
	// An array of relationships to iamIpAddress resources.
	IpAddresses []IamIpAddressRelationship `json:"IpAddresses,omitempty" yaml:"IpAddresses,omitempty"`
}

// NewIamIpAccessManagement instantiates a new IamIpAccessManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamIpAccessManagement(classId string, objectType string) *IamIpAccessManagement {
	this := IamIpAccessManagement{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamIpAccessManagementWithDefaults instantiates a new IamIpAccessManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamIpAccessManagementWithDefaults() *IamIpAccessManagement {
	this := IamIpAccessManagement{}
	var classId string = "iam.IpAccessManagement"
	this.ClassId = classId
	var objectType string = "iam.IpAccessManagement"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamIpAccessManagement) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamIpAccessManagement) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamIpAccessManagement) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamIpAccessManagement) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamIpAccessManagement) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamIpAccessManagement) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *IamIpAccessManagement) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAccessManagement) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *IamIpAccessManagement) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *IamIpAccessManagement) SetEnable(v bool) {
	o.Enable = &v
}

// GetLastRecoveryTime returns the LastRecoveryTime field value if set, zero value otherwise.
func (o *IamIpAccessManagement) GetLastRecoveryTime() time.Time {
	if o == nil || o.LastRecoveryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRecoveryTime
}

// GetLastRecoveryTimeOk returns a tuple with the LastRecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAccessManagement) GetLastRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || o.LastRecoveryTime == nil {
		return nil, false
	}
	return o.LastRecoveryTime, true
}

// HasLastRecoveryTime returns a boolean if a field has been set.
func (o *IamIpAccessManagement) HasLastRecoveryTime() bool {
	if o != nil && o.LastRecoveryTime != nil {
		return true
	}

	return false
}

// SetLastRecoveryTime gets a reference to the given time.Time and assigns it to the LastRecoveryTime field.
func (o *IamIpAccessManagement) SetLastRecoveryTime(v time.Time) {
	o.LastRecoveryTime = &v
}

// GetHolder returns the Holder field value if set, zero value otherwise.
func (o *IamIpAccessManagement) GetHolder() IamSecurityHolderRelationship {
	if o == nil || o.Holder == nil {
		var ret IamSecurityHolderRelationship
		return ret
	}
	return *o.Holder
}

// GetHolderOk returns a tuple with the Holder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAccessManagement) GetHolderOk() (*IamSecurityHolderRelationship, bool) {
	if o == nil || o.Holder == nil {
		return nil, false
	}
	return o.Holder, true
}

// HasHolder returns a boolean if a field has been set.
func (o *IamIpAccessManagement) HasHolder() bool {
	if o != nil && o.Holder != nil {
		return true
	}

	return false
}

// SetHolder gets a reference to the given IamSecurityHolderRelationship and assigns it to the Holder field.
func (o *IamIpAccessManagement) SetHolder(v IamSecurityHolderRelationship) {
	o.Holder = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIpAccessManagement) GetIpAddresses() []IamIpAddressRelationship {
	if o == nil {
		var ret []IamIpAddressRelationship
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIpAccessManagement) GetIpAddressesOk() (*[]IamIpAddressRelationship, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *IamIpAccessManagement) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []IamIpAddressRelationship and assigns it to the IpAddresses field.
func (o *IamIpAccessManagement) SetIpAddresses(v []IamIpAddressRelationship) {
	o.IpAddresses = v
}

func (o IamIpAccessManagement) MarshalJSON() ([]byte, error) {
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
	if o.Enable != nil {
		toSerialize["Enable"] = o.Enable
	}
	if o.LastRecoveryTime != nil {
		toSerialize["LastRecoveryTime"] = o.LastRecoveryTime
	}
	if o.Holder != nil {
		toSerialize["Holder"] = o.Holder
	}
	if o.IpAddresses != nil {
		toSerialize["IpAddresses"] = o.IpAddresses
	}
	return json.Marshal(toSerialize)
}

type NullableIamIpAccessManagement struct {
	value *IamIpAccessManagement
	isSet bool
}

func (v NullableIamIpAccessManagement) Get() *IamIpAccessManagement {
	return v.value
}

func (v *NullableIamIpAccessManagement) Set(val *IamIpAccessManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableIamIpAccessManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableIamIpAccessManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamIpAccessManagement(val *IamIpAccessManagement) *NullableIamIpAccessManagement {
	return &NullableIamIpAccessManagement{value: val, isSet: true}
}

func (v NullableIamIpAccessManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamIpAccessManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
