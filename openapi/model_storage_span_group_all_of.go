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

// StorageSpanGroupAllOf Definition of the list of properties defined in 'storage.SpanGroup', excluding properties defined in parent classes.
type StorageSpanGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string             `json:"ObjectType" yaml:"ObjectType"`
	Disks      []StorageLocalDisk `json:"Disks,omitempty" yaml:"Disks,omitempty"`
}

// NewStorageSpanGroupAllOf instantiates a new StorageSpanGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSpanGroupAllOf(classId string, objectType string) *StorageSpanGroupAllOf {
	this := StorageSpanGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageSpanGroupAllOfWithDefaults instantiates a new StorageSpanGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSpanGroupAllOfWithDefaults() *StorageSpanGroupAllOf {
	this := StorageSpanGroupAllOf{}
	var classId string = "storage.SpanGroup"
	this.ClassId = classId
	var objectType string = "storage.SpanGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageSpanGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageSpanGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageSpanGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageSpanGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageSpanGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageSpanGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisks returns the Disks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageSpanGroupAllOf) GetDisks() []StorageLocalDisk {
	if o == nil {
		var ret []StorageLocalDisk
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageSpanGroupAllOf) GetDisksOk() (*[]StorageLocalDisk, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return &o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *StorageSpanGroupAllOf) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []StorageLocalDisk and assigns it to the Disks field.
func (o *StorageSpanGroupAllOf) SetDisks(v []StorageLocalDisk) {
	o.Disks = v
}

func (o StorageSpanGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Disks != nil {
		toSerialize["Disks"] = o.Disks
	}
	return json.Marshal(toSerialize)
}

type NullableStorageSpanGroupAllOf struct {
	value *StorageSpanGroupAllOf
	isSet bool
}

func (v NullableStorageSpanGroupAllOf) Get() *StorageSpanGroupAllOf {
	return v.value
}

func (v *NullableStorageSpanGroupAllOf) Set(val *StorageSpanGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSpanGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSpanGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSpanGroupAllOf(val *StorageSpanGroupAllOf) *NullableStorageSpanGroupAllOf {
	return &NullableStorageSpanGroupAllOf{value: val, isSet: true}
}

func (v NullableStorageSpanGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSpanGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
