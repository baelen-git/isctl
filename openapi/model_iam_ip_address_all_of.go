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

// IamIpAddressAllOf Definition of the list of properties defined in 'iam.IpAddress', excluding properties defined in parent classes.
type IamIpAddressAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The Trusted IP range's address. IP address, CIDR range, and IP address range formats are supported. For example '12.13.14.15', '12.13.14.0/24', and '12.13.14.15-12.13.14.200'. Reserved IP ranges '127.0.0.1', '10.0.0.0/8', '172.16.0.0/12', and '192.168.0.0/16' are not allowed.
	Address *string `json:"Address,omitempty" yaml:"Address,omitempty"`
	// Description of Trusted IP address range.
	Description        *string                            `json:"Description,omitempty" yaml:"Description,omitempty"`
	IpAccessManagement *IamIpAccessManagementRelationship `json:"IpAccessManagement,omitempty" yaml:"IpAccessManagement,omitempty"`
}

// NewIamIpAddressAllOf instantiates a new IamIpAddressAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamIpAddressAllOf(classId string, objectType string) *IamIpAddressAllOf {
	this := IamIpAddressAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamIpAddressAllOfWithDefaults instantiates a new IamIpAddressAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamIpAddressAllOfWithDefaults() *IamIpAddressAllOf {
	this := IamIpAddressAllOf{}
	var classId string = "iam.IpAddress"
	this.ClassId = classId
	var objectType string = "iam.IpAddress"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamIpAddressAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamIpAddressAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamIpAddressAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamIpAddressAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamIpAddressAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamIpAddressAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *IamIpAddressAllOf) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAddressAllOf) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *IamIpAddressAllOf) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *IamIpAddressAllOf) SetAddress(v string) {
	o.Address = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamIpAddressAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAddressAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamIpAddressAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamIpAddressAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetIpAccessManagement returns the IpAccessManagement field value if set, zero value otherwise.
func (o *IamIpAddressAllOf) GetIpAccessManagement() IamIpAccessManagementRelationship {
	if o == nil || o.IpAccessManagement == nil {
		var ret IamIpAccessManagementRelationship
		return ret
	}
	return *o.IpAccessManagement
}

// GetIpAccessManagementOk returns a tuple with the IpAccessManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAddressAllOf) GetIpAccessManagementOk() (*IamIpAccessManagementRelationship, bool) {
	if o == nil || o.IpAccessManagement == nil {
		return nil, false
	}
	return o.IpAccessManagement, true
}

// HasIpAccessManagement returns a boolean if a field has been set.
func (o *IamIpAddressAllOf) HasIpAccessManagement() bool {
	if o != nil && o.IpAccessManagement != nil {
		return true
	}

	return false
}

// SetIpAccessManagement gets a reference to the given IamIpAccessManagementRelationship and assigns it to the IpAccessManagement field.
func (o *IamIpAddressAllOf) SetIpAccessManagement(v IamIpAccessManagementRelationship) {
	o.IpAccessManagement = &v
}

func (o IamIpAddressAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.IpAccessManagement != nil {
		toSerialize["IpAccessManagement"] = o.IpAccessManagement
	}
	return json.Marshal(toSerialize)
}

type NullableIamIpAddressAllOf struct {
	value *IamIpAddressAllOf
	isSet bool
}

func (v NullableIamIpAddressAllOf) Get() *IamIpAddressAllOf {
	return v.value
}

func (v *NullableIamIpAddressAllOf) Set(val *IamIpAddressAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamIpAddressAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamIpAddressAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamIpAddressAllOf(val *IamIpAddressAllOf) *NullableIamIpAddressAllOf {
	return &NullableIamIpAddressAllOf{value: val, isSet: true}
}

func (v NullableIamIpAddressAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamIpAddressAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
