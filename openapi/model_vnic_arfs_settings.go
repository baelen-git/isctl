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

// VnicArfsSettings Settings for Accelerated Receive Flow Steering to reduce the network latency and increase CPU cache efficiency.
type VnicArfsSettings struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// Status of Accelerated Receive Flow Steering on the virtual ethernet interface.
	Enabled *bool `json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
}

// NewVnicArfsSettings instantiates a new VnicArfsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicArfsSettings() *VnicArfsSettings {
	this := VnicArfsSettings{}
	return &this
}

// NewVnicArfsSettingsWithDefaults instantiates a new VnicArfsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicArfsSettingsWithDefaults() *VnicArfsSettings {
	this := VnicArfsSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VnicArfsSettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicArfsSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VnicArfsSettings) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VnicArfsSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o VnicArfsSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableVnicArfsSettings struct {
	value *VnicArfsSettings
	isSet bool
}

func (v NullableVnicArfsSettings) Get() *VnicArfsSettings {
	return v.value
}

func (v *NullableVnicArfsSettings) Set(val *VnicArfsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicArfsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicArfsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicArfsSettings(val *VnicArfsSettings) *NullableVnicArfsSettings {
	return &NullableVnicArfsSettings{value: val, isSet: true}
}

func (v NullableVnicArfsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicArfsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
