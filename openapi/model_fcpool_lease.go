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

// FcpoolLease Lease represents a single WWN ID that is part of the universe, allocated either from a pool or through static assignment.
type FcpoolLease struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// Purpose of this WWN pool.
	PoolPurpose *string `json:"PoolPurpose,omitempty" yaml:"PoolPurpose,omitempty"`
	// WWN ID allocated for pool based allocation.
	WwnId            *string                       `json:"WwnId,omitempty" yaml:"WwnId,omitempty"`
	AssignedToEntity *MoBaseMoRelationship         `json:"AssignedToEntity,omitempty" yaml:"AssignedToEntity,omitempty"`
	Pool             *FcpoolPoolRelationship       `json:"Pool,omitempty" yaml:"Pool,omitempty"`
	PoolMember       *FcpoolPoolMemberRelationship `json:"PoolMember,omitempty" yaml:"PoolMember,omitempty"`
	Universe         *FcpoolUniverseRelationship   `json:"Universe,omitempty" yaml:"Universe,omitempty"`
}

// NewFcpoolLease instantiates a new FcpoolLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolLease() *FcpoolLease {
	this := FcpoolLease{}
	return &this
}

// NewFcpoolLeaseWithDefaults instantiates a new FcpoolLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolLeaseWithDefaults() *FcpoolLease {
	this := FcpoolLease{}
	return &this
}

// GetPoolPurpose returns the PoolPurpose field value if set, zero value otherwise.
func (o *FcpoolLease) GetPoolPurpose() string {
	if o == nil || o.PoolPurpose == nil {
		var ret string
		return ret
	}
	return *o.PoolPurpose
}

// GetPoolPurposeOk returns a tuple with the PoolPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetPoolPurposeOk() (*string, bool) {
	if o == nil || o.PoolPurpose == nil {
		return nil, false
	}
	return o.PoolPurpose, true
}

// HasPoolPurpose returns a boolean if a field has been set.
func (o *FcpoolLease) HasPoolPurpose() bool {
	if o != nil && o.PoolPurpose != nil {
		return true
	}

	return false
}

// SetPoolPurpose gets a reference to the given string and assigns it to the PoolPurpose field.
func (o *FcpoolLease) SetPoolPurpose(v string) {
	o.PoolPurpose = &v
}

// GetWwnId returns the WwnId field value if set, zero value otherwise.
func (o *FcpoolLease) GetWwnId() string {
	if o == nil || o.WwnId == nil {
		var ret string
		return ret
	}
	return *o.WwnId
}

// GetWwnIdOk returns a tuple with the WwnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetWwnIdOk() (*string, bool) {
	if o == nil || o.WwnId == nil {
		return nil, false
	}
	return o.WwnId, true
}

// HasWwnId returns a boolean if a field has been set.
func (o *FcpoolLease) HasWwnId() bool {
	if o != nil && o.WwnId != nil {
		return true
	}

	return false
}

// SetWwnId gets a reference to the given string and assigns it to the WwnId field.
func (o *FcpoolLease) SetWwnId(v string) {
	o.WwnId = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *FcpoolLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *FcpoolLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *FcpoolLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *FcpoolLease) GetPool() FcpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *FcpoolLease) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given FcpoolPoolRelationship and assigns it to the Pool field.
func (o *FcpoolLease) SetPool(v FcpoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *FcpoolLease) GetPoolMember() FcpoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret FcpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetPoolMemberOk() (*FcpoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *FcpoolLease) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given FcpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *FcpoolLease) SetPoolMember(v FcpoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *FcpoolLease) GetUniverse() FcpoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret FcpoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolLease) GetUniverseOk() (*FcpoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *FcpoolLease) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given FcpoolUniverseRelationship and assigns it to the Universe field.
func (o *FcpoolLease) SetUniverse(v FcpoolUniverseRelationship) {
	o.Universe = &v
}

func (o FcpoolLease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.PoolPurpose != nil {
		toSerialize["PoolPurpose"] = o.PoolPurpose
	}
	if o.WwnId != nil {
		toSerialize["WwnId"] = o.WwnId
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}
	return json.Marshal(toSerialize)
}

type NullableFcpoolLease struct {
	value *FcpoolLease
	isSet bool
}

func (v NullableFcpoolLease) Get() *FcpoolLease {
	return v.value
}

func (v *NullableFcpoolLease) Set(val *FcpoolLease) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolLease) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolLease(val *FcpoolLease) *NullableFcpoolLease {
	return &NullableFcpoolLease{value: val, isSet: true}
}

func (v NullableFcpoolLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
