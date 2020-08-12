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

// IaasUcsdMessagesAllOf Definition of the list of properties defined in 'iaas.UcsdMessages', excluding properties defined in parent classes.
type IaasUcsdMessagesAllOf struct {
	// Last checked time of the alerts.
	StatusId         *string                              `json:"StatusId,omitempty" yaml:"StatusId,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewIaasUcsdMessagesAllOf instantiates a new IaasUcsdMessagesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasUcsdMessagesAllOf() *IaasUcsdMessagesAllOf {
	this := IaasUcsdMessagesAllOf{}
	return &this
}

// NewIaasUcsdMessagesAllOfWithDefaults instantiates a new IaasUcsdMessagesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasUcsdMessagesAllOfWithDefaults() *IaasUcsdMessagesAllOf {
	this := IaasUcsdMessagesAllOf{}
	return &this
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *IaasUcsdMessagesAllOf) GetStatusId() string {
	if o == nil || o.StatusId == nil {
		var ret string
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdMessagesAllOf) GetStatusIdOk() (*string, bool) {
	if o == nil || o.StatusId == nil {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *IaasUcsdMessagesAllOf) HasStatusId() bool {
	if o != nil && o.StatusId != nil {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given string and assigns it to the StatusId field.
func (o *IaasUcsdMessagesAllOf) SetStatusId(v string) {
	o.StatusId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *IaasUcsdMessagesAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdMessagesAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *IaasUcsdMessagesAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *IaasUcsdMessagesAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o IaasUcsdMessagesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StatusId != nil {
		toSerialize["StatusId"] = o.StatusId
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableIaasUcsdMessagesAllOf struct {
	value *IaasUcsdMessagesAllOf
	isSet bool
}

func (v NullableIaasUcsdMessagesAllOf) Get() *IaasUcsdMessagesAllOf {
	return v.value
}

func (v *NullableIaasUcsdMessagesAllOf) Set(val *IaasUcsdMessagesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasUcsdMessagesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasUcsdMessagesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasUcsdMessagesAllOf(val *IaasUcsdMessagesAllOf) *NullableIaasUcsdMessagesAllOf {
	return &NullableIaasUcsdMessagesAllOf{value: val, isSet: true}
}

func (v NullableIaasUcsdMessagesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasUcsdMessagesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
