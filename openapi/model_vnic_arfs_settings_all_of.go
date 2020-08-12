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

// VnicArfsSettingsAllOf Definition of the list of properties defined in 'vnic.ArfsSettings', excluding properties defined in parent classes.
type VnicArfsSettingsAllOf struct {
	// Status of Accelerated Receive Flow Steering on the virtual ethernet interface.
	Enabled *bool `json:"Enabled,omitempty" yaml:"Enabled,omitempty"`
}

// NewVnicArfsSettingsAllOf instantiates a new VnicArfsSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicArfsSettingsAllOf() *VnicArfsSettingsAllOf {
	this := VnicArfsSettingsAllOf{}
	return &this
}

// NewVnicArfsSettingsAllOfWithDefaults instantiates a new VnicArfsSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicArfsSettingsAllOfWithDefaults() *VnicArfsSettingsAllOf {
	this := VnicArfsSettingsAllOf{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VnicArfsSettingsAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicArfsSettingsAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VnicArfsSettingsAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VnicArfsSettingsAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o VnicArfsSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableVnicArfsSettingsAllOf struct {
	value *VnicArfsSettingsAllOf
	isSet bool
}

func (v NullableVnicArfsSettingsAllOf) Get() *VnicArfsSettingsAllOf {
	return v.value
}

func (v *NullableVnicArfsSettingsAllOf) Set(val *VnicArfsSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicArfsSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicArfsSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicArfsSettingsAllOf(val *VnicArfsSettingsAllOf) *NullableVnicArfsSettingsAllOf {
	return &NullableVnicArfsSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicArfsSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicArfsSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
