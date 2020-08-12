/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// SdcardPartition This adds and configures operating system partitions.
type SdcardPartition struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// This specifies the type of the partition. Allowed values are OS, Utility. * `OS` - This partition contains virtual drives where user can install operating system. * `Utility` - This partition contains virtual drives for utilities such as SCU, HUU, Drivers and Diagnostics.
	Type          *string               `json:"Type,omitempty" yaml:"Type,omitempty"`
	VirtualDrives *[]SdcardVirtualDrive `json:"VirtualDrives,omitempty" yaml:"VirtualDrives,omitempty"`
}

// NewSdcardPartition instantiates a new SdcardPartition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdcardPartition() *SdcardPartition {
	this := SdcardPartition{}
	var type_ string = "OS"
	this.Type = &type_
	return &this
}

// NewSdcardPartitionWithDefaults instantiates a new SdcardPartition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdcardPartitionWithDefaults() *SdcardPartition {
	this := SdcardPartition{}
	var type_ string = "OS"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SdcardPartition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdcardPartition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SdcardPartition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SdcardPartition) SetType(v string) {
	o.Type = &v
}

// GetVirtualDrives returns the VirtualDrives field value if set, zero value otherwise.
func (o *SdcardPartition) GetVirtualDrives() []SdcardVirtualDrive {
	if o == nil || o.VirtualDrives == nil {
		var ret []SdcardVirtualDrive
		return ret
	}
	return *o.VirtualDrives
}

// GetVirtualDrivesOk returns a tuple with the VirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdcardPartition) GetVirtualDrivesOk() (*[]SdcardVirtualDrive, bool) {
	if o == nil || o.VirtualDrives == nil {
		return nil, false
	}
	return o.VirtualDrives, true
}

// HasVirtualDrives returns a boolean if a field has been set.
func (o *SdcardPartition) HasVirtualDrives() bool {
	if o != nil && o.VirtualDrives != nil {
		return true
	}

	return false
}

// SetVirtualDrives gets a reference to the given []SdcardVirtualDrive and assigns it to the VirtualDrives field.
func (o *SdcardPartition) SetVirtualDrives(v []SdcardVirtualDrive) {
	o.VirtualDrives = &v
}

func (o SdcardPartition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VirtualDrives != nil {
		toSerialize["VirtualDrives"] = o.VirtualDrives
	}
	return json.Marshal(toSerialize)
}

type NullableSdcardPartition struct {
	value *SdcardPartition
	isSet bool
}

func (v NullableSdcardPartition) Get() *SdcardPartition {
	return v.value
}

func (v *NullableSdcardPartition) Set(val *SdcardPartition) {
	v.value = val
	v.isSet = true
}

func (v NullableSdcardPartition) IsSet() bool {
	return v.isSet
}

func (v *NullableSdcardPartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdcardPartition(val *SdcardPartition) *NullableSdcardPartition {
	return &NullableSdcardPartition{value: val, isSet: true}
}

func (v NullableSdcardPartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdcardPartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
