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

// IppoolPoolMember PoolMember represents a single IPv4 address that is part of a pool.
type IppoolPoolMember struct {
	PoolAbstractPoolMember `yaml:"PoolAbstractPoolMember,inline"`
	// IPv4 Address of this pool member.
	IpV4Address      *string                        `json:"IpV4Address,omitempty" yaml:"IpV4Address,omitempty"`
	AssignedToEntity *MoBaseMoRelationship          `json:"AssignedToEntity,omitempty" yaml:"AssignedToEntity,omitempty"`
	IpV4Block        *IppoolShadowBlockRelationship `json:"IpV4Block,omitempty" yaml:"IpV4Block,omitempty"`
	Peer             *IppoolIpLeaseRelationship     `json:"Peer,omitempty" yaml:"Peer,omitempty"`
	Pool             *IppoolShadowPoolRelationship  `json:"Pool,omitempty" yaml:"Pool,omitempty"`
}

// NewIppoolPoolMember instantiates a new IppoolPoolMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolPoolMember() *IppoolPoolMember {
	this := IppoolPoolMember{}
	return &this
}

// NewIppoolPoolMemberWithDefaults instantiates a new IppoolPoolMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolPoolMemberWithDefaults() *IppoolPoolMember {
	this := IppoolPoolMember{}
	return &this
}

// GetIpV4Address returns the IpV4Address field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetIpV4Address() string {
	if o == nil || o.IpV4Address == nil {
		var ret string
		return ret
	}
	return *o.IpV4Address
}

// GetIpV4AddressOk returns a tuple with the IpV4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetIpV4AddressOk() (*string, bool) {
	if o == nil || o.IpV4Address == nil {
		return nil, false
	}
	return o.IpV4Address, true
}

// HasIpV4Address returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasIpV4Address() bool {
	if o != nil && o.IpV4Address != nil {
		return true
	}

	return false
}

// SetIpV4Address gets a reference to the given string and assigns it to the IpV4Address field.
func (o *IppoolPoolMember) SetIpV4Address(v string) {
	o.IpV4Address = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *IppoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetIpV4Block returns the IpV4Block field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetIpV4Block() IppoolShadowBlockRelationship {
	if o == nil || o.IpV4Block == nil {
		var ret IppoolShadowBlockRelationship
		return ret
	}
	return *o.IpV4Block
}

// GetIpV4BlockOk returns a tuple with the IpV4Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetIpV4BlockOk() (*IppoolShadowBlockRelationship, bool) {
	if o == nil || o.IpV4Block == nil {
		return nil, false
	}
	return o.IpV4Block, true
}

// HasIpV4Block returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasIpV4Block() bool {
	if o != nil && o.IpV4Block != nil {
		return true
	}

	return false
}

// SetIpV4Block gets a reference to the given IppoolShadowBlockRelationship and assigns it to the IpV4Block field.
func (o *IppoolPoolMember) SetIpV4Block(v IppoolShadowBlockRelationship) {
	o.IpV4Block = &v
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetPeer() IppoolIpLeaseRelationship {
	if o == nil || o.Peer == nil {
		var ret IppoolIpLeaseRelationship
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetPeerOk() (*IppoolIpLeaseRelationship, bool) {
	if o == nil || o.Peer == nil {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasPeer() bool {
	if o != nil && o.Peer != nil {
		return true
	}

	return false
}

// SetPeer gets a reference to the given IppoolIpLeaseRelationship and assigns it to the Peer field.
func (o *IppoolPoolMember) SetPeer(v IppoolIpLeaseRelationship) {
	o.Peer = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetPool() IppoolShadowPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolShadowPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetPoolOk() (*IppoolShadowPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolShadowPoolRelationship and assigns it to the Pool field.
func (o *IppoolPoolMember) SetPool(v IppoolShadowPoolRelationship) {
	o.Pool = &v
}

func (o IppoolPoolMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPoolMember, errPoolAbstractPoolMember := json.Marshal(o.PoolAbstractPoolMember)
	if errPoolAbstractPoolMember != nil {
		return []byte{}, errPoolAbstractPoolMember
	}
	errPoolAbstractPoolMember = json.Unmarshal([]byte(serializedPoolAbstractPoolMember), &toSerialize)
	if errPoolAbstractPoolMember != nil {
		return []byte{}, errPoolAbstractPoolMember
	}
	if o.IpV4Address != nil {
		toSerialize["IpV4Address"] = o.IpV4Address
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.IpV4Block != nil {
		toSerialize["IpV4Block"] = o.IpV4Block
	}
	if o.Peer != nil {
		toSerialize["Peer"] = o.Peer
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	return json.Marshal(toSerialize)
}

type NullableIppoolPoolMember struct {
	value *IppoolPoolMember
	isSet bool
}

func (v NullableIppoolPoolMember) Get() *IppoolPoolMember {
	return v.value
}

func (v *NullableIppoolPoolMember) Set(val *IppoolPoolMember) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolPoolMember) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolPoolMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolPoolMember(val *IppoolPoolMember) *NullableIppoolPoolMember {
	return &NullableIppoolPoolMember{value: val, isSet: true}
}

func (v NullableIppoolPoolMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolPoolMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
