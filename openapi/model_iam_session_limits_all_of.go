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

// IamSessionLimitsAllOf Definition of the list of properties defined in 'iam.SessionLimits', excluding properties defined in parent classes.
type IamSessionLimitsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The idle timeout interval for the web session in seconds. When a session is not refreshed for this duration, the session is marked as idle and removed. The minimum value is 300 seconds and the maximum value is 18000 seconds (5 hours). The system default value is 1800 seconds.
	IdleTimeOut *int64 `json:"IdleTimeOut,omitempty" yaml:"IdleTimeOut,omitempty"`
	// The maximum number of sessions allowed in an account. The default value is 128.
	MaximumLimit *int64 `json:"MaximumLimit,omitempty" yaml:"MaximumLimit,omitempty"`
	// The maximum number of sessions allowed per user. Default value is 32.
	PerUserLimit *int64 `json:"PerUserLimit,omitempty" yaml:"PerUserLimit,omitempty"`
	// The session expiry duration in seconds. The minimum value is 350 seconds and the maximum value is 31536000 seconds (1 year). The system default value is 57600 seconds.
	SessionTimeOut *int64                     `json:"SessionTimeOut,omitempty" yaml:"SessionTimeOut,omitempty"`
	Account        *IamAccountRelationship    `json:"Account,omitempty" yaml:"Account,omitempty"`
	Permission     *IamPermissionRelationship `json:"Permission,omitempty" yaml:"Permission,omitempty"`
}

// NewIamSessionLimitsAllOf instantiates a new IamSessionLimitsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSessionLimitsAllOf(classId string, objectType string) *IamSessionLimitsAllOf {
	this := IamSessionLimitsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var idleTimeOut int64 = 1800
	this.IdleTimeOut = &idleTimeOut
	var maximumLimit int64 = 128
	this.MaximumLimit = &maximumLimit
	var perUserLimit int64 = 32
	this.PerUserLimit = &perUserLimit
	var sessionTimeOut int64 = 57600
	this.SessionTimeOut = &sessionTimeOut
	return &this
}

// NewIamSessionLimitsAllOfWithDefaults instantiates a new IamSessionLimitsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSessionLimitsAllOfWithDefaults() *IamSessionLimitsAllOf {
	this := IamSessionLimitsAllOf{}
	var classId string = "iam.SessionLimits"
	this.ClassId = classId
	var objectType string = "iam.SessionLimits"
	this.ObjectType = objectType
	var idleTimeOut int64 = 1800
	this.IdleTimeOut = &idleTimeOut
	var maximumLimit int64 = 128
	this.MaximumLimit = &maximumLimit
	var perUserLimit int64 = 32
	this.PerUserLimit = &perUserLimit
	var sessionTimeOut int64 = 57600
	this.SessionTimeOut = &sessionTimeOut
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamSessionLimitsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamSessionLimitsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamSessionLimitsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamSessionLimitsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamSessionLimitsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamSessionLimitsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdleTimeOut returns the IdleTimeOut field value if set, zero value otherwise.
func (o *IamSessionLimitsAllOf) GetIdleTimeOut() int64 {
	if o == nil || o.IdleTimeOut == nil {
		var ret int64
		return ret
	}
	return *o.IdleTimeOut
}

// GetIdleTimeOutOk returns a tuple with the IdleTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimitsAllOf) GetIdleTimeOutOk() (*int64, bool) {
	if o == nil || o.IdleTimeOut == nil {
		return nil, false
	}
	return o.IdleTimeOut, true
}

// HasIdleTimeOut returns a boolean if a field has been set.
func (o *IamSessionLimitsAllOf) HasIdleTimeOut() bool {
	if o != nil && o.IdleTimeOut != nil {
		return true
	}

	return false
}

// SetIdleTimeOut gets a reference to the given int64 and assigns it to the IdleTimeOut field.
func (o *IamSessionLimitsAllOf) SetIdleTimeOut(v int64) {
	o.IdleTimeOut = &v
}

// GetMaximumLimit returns the MaximumLimit field value if set, zero value otherwise.
func (o *IamSessionLimitsAllOf) GetMaximumLimit() int64 {
	if o == nil || o.MaximumLimit == nil {
		var ret int64
		return ret
	}
	return *o.MaximumLimit
}

// GetMaximumLimitOk returns a tuple with the MaximumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimitsAllOf) GetMaximumLimitOk() (*int64, bool) {
	if o == nil || o.MaximumLimit == nil {
		return nil, false
	}
	return o.MaximumLimit, true
}

// HasMaximumLimit returns a boolean if a field has been set.
func (o *IamSessionLimitsAllOf) HasMaximumLimit() bool {
	if o != nil && o.MaximumLimit != nil {
		return true
	}

	return false
}

// SetMaximumLimit gets a reference to the given int64 and assigns it to the MaximumLimit field.
func (o *IamSessionLimitsAllOf) SetMaximumLimit(v int64) {
	o.MaximumLimit = &v
}

// GetPerUserLimit returns the PerUserLimit field value if set, zero value otherwise.
func (o *IamSessionLimitsAllOf) GetPerUserLimit() int64 {
	if o == nil || o.PerUserLimit == nil {
		var ret int64
		return ret
	}
	return *o.PerUserLimit
}

// GetPerUserLimitOk returns a tuple with the PerUserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimitsAllOf) GetPerUserLimitOk() (*int64, bool) {
	if o == nil || o.PerUserLimit == nil {
		return nil, false
	}
	return o.PerUserLimit, true
}

// HasPerUserLimit returns a boolean if a field has been set.
func (o *IamSessionLimitsAllOf) HasPerUserLimit() bool {
	if o != nil && o.PerUserLimit != nil {
		return true
	}

	return false
}

// SetPerUserLimit gets a reference to the given int64 and assigns it to the PerUserLimit field.
func (o *IamSessionLimitsAllOf) SetPerUserLimit(v int64) {
	o.PerUserLimit = &v
}

// GetSessionTimeOut returns the SessionTimeOut field value if set, zero value otherwise.
func (o *IamSessionLimitsAllOf) GetSessionTimeOut() int64 {
	if o == nil || o.SessionTimeOut == nil {
		var ret int64
		return ret
	}
	return *o.SessionTimeOut
}

// GetSessionTimeOutOk returns a tuple with the SessionTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimitsAllOf) GetSessionTimeOutOk() (*int64, bool) {
	if o == nil || o.SessionTimeOut == nil {
		return nil, false
	}
	return o.SessionTimeOut, true
}

// HasSessionTimeOut returns a boolean if a field has been set.
func (o *IamSessionLimitsAllOf) HasSessionTimeOut() bool {
	if o != nil && o.SessionTimeOut != nil {
		return true
	}

	return false
}

// SetSessionTimeOut gets a reference to the given int64 and assigns it to the SessionTimeOut field.
func (o *IamSessionLimitsAllOf) SetSessionTimeOut(v int64) {
	o.SessionTimeOut = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamSessionLimitsAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimitsAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamSessionLimitsAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamSessionLimitsAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamSessionLimitsAllOf) GetPermission() IamPermissionRelationship {
	if o == nil || o.Permission == nil {
		var ret IamPermissionRelationship
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimitsAllOf) GetPermissionOk() (*IamPermissionRelationship, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamSessionLimitsAllOf) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given IamPermissionRelationship and assigns it to the Permission field.
func (o *IamSessionLimitsAllOf) SetPermission(v IamPermissionRelationship) {
	o.Permission = &v
}

func (o IamSessionLimitsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IdleTimeOut != nil {
		toSerialize["IdleTimeOut"] = o.IdleTimeOut
	}
	if o.MaximumLimit != nil {
		toSerialize["MaximumLimit"] = o.MaximumLimit
	}
	if o.PerUserLimit != nil {
		toSerialize["PerUserLimit"] = o.PerUserLimit
	}
	if o.SessionTimeOut != nil {
		toSerialize["SessionTimeOut"] = o.SessionTimeOut
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	return json.Marshal(toSerialize)
}

type NullableIamSessionLimitsAllOf struct {
	value *IamSessionLimitsAllOf
	isSet bool
}

func (v NullableIamSessionLimitsAllOf) Get() *IamSessionLimitsAllOf {
	return v.value
}

func (v *NullableIamSessionLimitsAllOf) Set(val *IamSessionLimitsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSessionLimitsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSessionLimitsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSessionLimitsAllOf(val *IamSessionLimitsAllOf) *NullableIamSessionLimitsAllOf {
	return &NullableIamSessionLimitsAllOf{value: val, isSet: true}
}

func (v NullableIamSessionLimitsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSessionLimitsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
