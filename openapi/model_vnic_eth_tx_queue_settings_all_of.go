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

// VnicEthTxQueueSettingsAllOf Definition of the list of properties defined in 'vnic.EthTxQueueSettings', excluding properties defined in parent classes.
type VnicEthTxQueueSettingsAllOf struct {
	// The number of queue resources to allocate.
	Count *int64 `json:"Count,omitempty" yaml:"Count,omitempty"`
	// The number of descriptors in each queue.
	RingSize *int64 `json:"RingSize,omitempty" yaml:"RingSize,omitempty"`
}

// NewVnicEthTxQueueSettingsAllOf instantiates a new VnicEthTxQueueSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthTxQueueSettingsAllOf() *VnicEthTxQueueSettingsAllOf {
	this := VnicEthTxQueueSettingsAllOf{}
	return &this
}

// NewVnicEthTxQueueSettingsAllOfWithDefaults instantiates a new VnicEthTxQueueSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthTxQueueSettingsAllOfWithDefaults() *VnicEthTxQueueSettingsAllOf {
	this := VnicEthTxQueueSettingsAllOf{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VnicEthTxQueueSettingsAllOf) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthTxQueueSettingsAllOf) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VnicEthTxQueueSettingsAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *VnicEthTxQueueSettingsAllOf) SetCount(v int64) {
	o.Count = &v
}

// GetRingSize returns the RingSize field value if set, zero value otherwise.
func (o *VnicEthTxQueueSettingsAllOf) GetRingSize() int64 {
	if o == nil || o.RingSize == nil {
		var ret int64
		return ret
	}
	return *o.RingSize
}

// GetRingSizeOk returns a tuple with the RingSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthTxQueueSettingsAllOf) GetRingSizeOk() (*int64, bool) {
	if o == nil || o.RingSize == nil {
		return nil, false
	}
	return o.RingSize, true
}

// HasRingSize returns a boolean if a field has been set.
func (o *VnicEthTxQueueSettingsAllOf) HasRingSize() bool {
	if o != nil && o.RingSize != nil {
		return true
	}

	return false
}

// SetRingSize gets a reference to the given int64 and assigns it to the RingSize field.
func (o *VnicEthTxQueueSettingsAllOf) SetRingSize(v int64) {
	o.RingSize = &v
}

func (o VnicEthTxQueueSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.RingSize != nil {
		toSerialize["RingSize"] = o.RingSize
	}
	return json.Marshal(toSerialize)
}

type NullableVnicEthTxQueueSettingsAllOf struct {
	value *VnicEthTxQueueSettingsAllOf
	isSet bool
}

func (v NullableVnicEthTxQueueSettingsAllOf) Get() *VnicEthTxQueueSettingsAllOf {
	return v.value
}

func (v *NullableVnicEthTxQueueSettingsAllOf) Set(val *VnicEthTxQueueSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthTxQueueSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthTxQueueSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthTxQueueSettingsAllOf(val *VnicEthTxQueueSettingsAllOf) *NullableVnicEthTxQueueSettingsAllOf {
	return &NullableVnicEthTxQueueSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicEthTxQueueSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthTxQueueSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
