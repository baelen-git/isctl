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

// HyperflexCiscoHypervisorManager A hypervisor manager to manage Cisco HyperFlex compute clusters and is available per user account.
type HyperflexCiscoHypervisorManager struct {
	VirtualizationBaseHypervisorManager `yaml:"VirtualizationBaseHypervisorManager,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                  `json:"ObjectType" yaml:"ObjectType"`
	Account    *IamAccountRelationship `json:"Account,omitempty" yaml:"Account,omitempty"`
}

// NewHyperflexCiscoHypervisorManager instantiates a new HyperflexCiscoHypervisorManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexCiscoHypervisorManager(classId string, objectType string) *HyperflexCiscoHypervisorManager {
	this := HyperflexCiscoHypervisorManager{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexCiscoHypervisorManagerWithDefaults instantiates a new HyperflexCiscoHypervisorManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexCiscoHypervisorManagerWithDefaults() *HyperflexCiscoHypervisorManager {
	this := HyperflexCiscoHypervisorManager{}
	var classId string = "hyperflex.CiscoHypervisorManager"
	this.ClassId = classId
	var objectType string = "hyperflex.CiscoHypervisorManager"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexCiscoHypervisorManager) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexCiscoHypervisorManager) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexCiscoHypervisorManager) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexCiscoHypervisorManager) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexCiscoHypervisorManager) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexCiscoHypervisorManager) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *HyperflexCiscoHypervisorManager) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCiscoHypervisorManager) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *HyperflexCiscoHypervisorManager) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *HyperflexCiscoHypervisorManager) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o HyperflexCiscoHypervisorManager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseHypervisorManager, errVirtualizationBaseHypervisorManager := json.Marshal(o.VirtualizationBaseHypervisorManager)
	if errVirtualizationBaseHypervisorManager != nil {
		return []byte{}, errVirtualizationBaseHypervisorManager
	}
	errVirtualizationBaseHypervisorManager = json.Unmarshal([]byte(serializedVirtualizationBaseHypervisorManager), &toSerialize)
	if errVirtualizationBaseHypervisorManager != nil {
		return []byte{}, errVirtualizationBaseHypervisorManager
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexCiscoHypervisorManager struct {
	value *HyperflexCiscoHypervisorManager
	isSet bool
}

func (v NullableHyperflexCiscoHypervisorManager) Get() *HyperflexCiscoHypervisorManager {
	return v.value
}

func (v *NullableHyperflexCiscoHypervisorManager) Set(val *HyperflexCiscoHypervisorManager) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexCiscoHypervisorManager) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexCiscoHypervisorManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexCiscoHypervisorManager(val *HyperflexCiscoHypervisorManager) *NullableHyperflexCiscoHypervisorManager {
	return &NullableHyperflexCiscoHypervisorManager{value: val, isSet: true}
}

func (v NullableHyperflexCiscoHypervisorManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexCiscoHypervisorManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
