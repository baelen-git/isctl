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

// StorageBaseHostLun Generic storage host lun object. It exists only if the volume is associated to host initiator.
type StorageBaseHostLun struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint.
	Hlu *int64 `json:"Hlu,omitempty" yaml:"Hlu,omitempty"`
	// Name of the host associated with LUN.
	HostName *string `json:"HostName,omitempty" yaml:"HostName,omitempty"`
	// Name of the storage volume associated with LUN.
	VolumeName *string `json:"VolumeName,omitempty" yaml:"VolumeName,omitempty"`
}

// NewStorageBaseHostLun instantiates a new StorageBaseHostLun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseHostLun(classId string, objectType string) *StorageBaseHostLun {
	this := StorageBaseHostLun{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseHostLunWithDefaults instantiates a new StorageBaseHostLun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseHostLunWithDefaults() *StorageBaseHostLun {
	this := StorageBaseHostLun{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseHostLun) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseHostLun) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseHostLun) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseHostLun) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseHostLun) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseHostLun) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHlu returns the Hlu field value if set, zero value otherwise.
func (o *StorageBaseHostLun) GetHlu() int64 {
	if o == nil || o.Hlu == nil {
		var ret int64
		return ret
	}
	return *o.Hlu
}

// GetHluOk returns a tuple with the Hlu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseHostLun) GetHluOk() (*int64, bool) {
	if o == nil || o.Hlu == nil {
		return nil, false
	}
	return o.Hlu, true
}

// HasHlu returns a boolean if a field has been set.
func (o *StorageBaseHostLun) HasHlu() bool {
	if o != nil && o.Hlu != nil {
		return true
	}

	return false
}

// SetHlu gets a reference to the given int64 and assigns it to the Hlu field.
func (o *StorageBaseHostLun) SetHlu(v int64) {
	o.Hlu = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *StorageBaseHostLun) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseHostLun) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *StorageBaseHostLun) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *StorageBaseHostLun) SetHostName(v string) {
	o.HostName = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *StorageBaseHostLun) GetVolumeName() string {
	if o == nil || o.VolumeName == nil {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseHostLun) GetVolumeNameOk() (*string, bool) {
	if o == nil || o.VolumeName == nil {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *StorageBaseHostLun) HasVolumeName() bool {
	if o != nil && o.VolumeName != nil {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *StorageBaseHostLun) SetVolumeName(v string) {
	o.VolumeName = &v
}

func (o StorageBaseHostLun) MarshalJSON() ([]byte, error) {
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
	if o.Hlu != nil {
		toSerialize["Hlu"] = o.Hlu
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.VolumeName != nil {
		toSerialize["VolumeName"] = o.VolumeName
	}
	return json.Marshal(toSerialize)
}

type NullableStorageBaseHostLun struct {
	value *StorageBaseHostLun
	isSet bool
}

func (v NullableStorageBaseHostLun) Get() *StorageBaseHostLun {
	return v.value
}

func (v *NullableStorageBaseHostLun) Set(val *StorageBaseHostLun) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseHostLun) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseHostLun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseHostLun(val *StorageBaseHostLun) *NullableStorageBaseHostLun {
	return &NullableStorageBaseHostLun{value: val, isSet: true}
}

func (v NullableStorageBaseHostLun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseHostLun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
