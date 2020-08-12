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

// FabricPcMemberAllOf Definition of the list of properties defined in 'fabric.PcMember', excluding properties defined in parent classes.
type FabricPcMemberAllOf struct {
	// Port Channel Identifier for the collection of ports.
	PcId       *int64                        `json:"PcId,omitempty" yaml:"PcId,omitempty"`
	PortPolicy *FabricPortPolicyRelationship `json:"PortPolicy,omitempty" yaml:"PortPolicy,omitempty"`
}

// NewFabricPcMemberAllOf instantiates a new FabricPcMemberAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPcMemberAllOf() *FabricPcMemberAllOf {
	this := FabricPcMemberAllOf{}
	return &this
}

// NewFabricPcMemberAllOfWithDefaults instantiates a new FabricPcMemberAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPcMemberAllOfWithDefaults() *FabricPcMemberAllOf {
	this := FabricPcMemberAllOf{}
	return &this
}

// GetPcId returns the PcId field value if set, zero value otherwise.
func (o *FabricPcMemberAllOf) GetPcId() int64 {
	if o == nil || o.PcId == nil {
		var ret int64
		return ret
	}
	return *o.PcId
}

// GetPcIdOk returns a tuple with the PcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcMemberAllOf) GetPcIdOk() (*int64, bool) {
	if o == nil || o.PcId == nil {
		return nil, false
	}
	return o.PcId, true
}

// HasPcId returns a boolean if a field has been set.
func (o *FabricPcMemberAllOf) HasPcId() bool {
	if o != nil && o.PcId != nil {
		return true
	}

	return false
}

// SetPcId gets a reference to the given int64 and assigns it to the PcId field.
func (o *FabricPcMemberAllOf) SetPcId(v int64) {
	o.PcId = &v
}

// GetPortPolicy returns the PortPolicy field value if set, zero value otherwise.
func (o *FabricPcMemberAllOf) GetPortPolicy() FabricPortPolicyRelationship {
	if o == nil || o.PortPolicy == nil {
		var ret FabricPortPolicyRelationship
		return ret
	}
	return *o.PortPolicy
}

// GetPortPolicyOk returns a tuple with the PortPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcMemberAllOf) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool) {
	if o == nil || o.PortPolicy == nil {
		return nil, false
	}
	return o.PortPolicy, true
}

// HasPortPolicy returns a boolean if a field has been set.
func (o *FabricPcMemberAllOf) HasPortPolicy() bool {
	if o != nil && o.PortPolicy != nil {
		return true
	}

	return false
}

// SetPortPolicy gets a reference to the given FabricPortPolicyRelationship and assigns it to the PortPolicy field.
func (o *FabricPcMemberAllOf) SetPortPolicy(v FabricPortPolicyRelationship) {
	o.PortPolicy = &v
}

func (o FabricPcMemberAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PcId != nil {
		toSerialize["PcId"] = o.PcId
	}
	if o.PortPolicy != nil {
		toSerialize["PortPolicy"] = o.PortPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableFabricPcMemberAllOf struct {
	value *FabricPcMemberAllOf
	isSet bool
}

func (v NullableFabricPcMemberAllOf) Get() *FabricPcMemberAllOf {
	return v.value
}

func (v *NullableFabricPcMemberAllOf) Set(val *FabricPcMemberAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPcMemberAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPcMemberAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPcMemberAllOf(val *FabricPcMemberAllOf) *NullableFabricPcMemberAllOf {
	return &NullableFabricPcMemberAllOf{value: val, isSet: true}
}

func (v NullableFabricPcMemberAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPcMemberAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
