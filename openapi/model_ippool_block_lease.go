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

// IppoolBlockLease BlockLease represents an IP address that is allocated from a pool to a specific entity like server profile.
type IppoolBlockLease struct {
	PoolAbstractBlockLease `yaml:"PoolAbstractBlockLease,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Type of the IP address requested. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
	IpType           *string               `json:"IpType,omitempty" yaml:"IpType,omitempty"`
	AssignedToEntity *MoBaseMoRelationship `json:"AssignedToEntity,omitempty" yaml:"AssignedToEntity,omitempty"`
	// An array of relationships to ippoolIpLease resources.
	IpLeases []IppoolIpLeaseRelationship `json:"IpLeases,omitempty" yaml:"IpLeases,omitempty"`
	Pool     *IppoolPoolRelationship     `json:"Pool,omitempty" yaml:"Pool,omitempty"`
	Universe *IppoolUniverseRelationship `json:"Universe,omitempty" yaml:"Universe,omitempty"`
	Vrf      *VrfVrfRelationship         `json:"Vrf,omitempty" yaml:"Vrf,omitempty"`
}

// NewIppoolBlockLease instantiates a new IppoolBlockLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolBlockLease(classId string, objectType string) *IppoolBlockLease {
	this := IppoolBlockLease{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// NewIppoolBlockLeaseWithDefaults instantiates a new IppoolBlockLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolBlockLeaseWithDefaults() *IppoolBlockLease {
	this := IppoolBlockLease{}
	var classId string = "ippool.BlockLease"
	this.ClassId = classId
	var objectType string = "ippool.BlockLease"
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolBlockLease) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolBlockLease) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolBlockLease) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IppoolBlockLease) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolBlockLease) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolBlockLease) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *IppoolBlockLease) GetIpType() string {
	if o == nil || o.IpType == nil {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolBlockLease) GetIpTypeOk() (*string, bool) {
	if o == nil || o.IpType == nil {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *IppoolBlockLease) HasIpType() bool {
	if o != nil && o.IpType != nil {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *IppoolBlockLease) SetIpType(v string) {
	o.IpType = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *IppoolBlockLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolBlockLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *IppoolBlockLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *IppoolBlockLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetIpLeases returns the IpLeases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolBlockLease) GetIpLeases() []IppoolIpLeaseRelationship {
	if o == nil {
		var ret []IppoolIpLeaseRelationship
		return ret
	}
	return o.IpLeases
}

// GetIpLeasesOk returns a tuple with the IpLeases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolBlockLease) GetIpLeasesOk() (*[]IppoolIpLeaseRelationship, bool) {
	if o == nil || o.IpLeases == nil {
		return nil, false
	}
	return &o.IpLeases, true
}

// HasIpLeases returns a boolean if a field has been set.
func (o *IppoolBlockLease) HasIpLeases() bool {
	if o != nil && o.IpLeases != nil {
		return true
	}

	return false
}

// SetIpLeases gets a reference to the given []IppoolIpLeaseRelationship and assigns it to the IpLeases field.
func (o *IppoolBlockLease) SetIpLeases(v []IppoolIpLeaseRelationship) {
	o.IpLeases = v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolBlockLease) GetPool() IppoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolBlockLease) GetPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolBlockLease) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolPoolRelationship and assigns it to the Pool field.
func (o *IppoolBlockLease) SetPool(v IppoolPoolRelationship) {
	o.Pool = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *IppoolBlockLease) GetUniverse() IppoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret IppoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolBlockLease) GetUniverseOk() (*IppoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *IppoolBlockLease) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given IppoolUniverseRelationship and assigns it to the Universe field.
func (o *IppoolBlockLease) SetUniverse(v IppoolUniverseRelationship) {
	o.Universe = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *IppoolBlockLease) GetVrf() VrfVrfRelationship {
	if o == nil || o.Vrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolBlockLease) GetVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *IppoolBlockLease) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given VrfVrfRelationship and assigns it to the Vrf field.
func (o *IppoolBlockLease) SetVrf(v VrfVrfRelationship) {
	o.Vrf = &v
}

func (o IppoolBlockLease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractBlockLease, errPoolAbstractBlockLease := json.Marshal(o.PoolAbstractBlockLease)
	if errPoolAbstractBlockLease != nil {
		return []byte{}, errPoolAbstractBlockLease
	}
	errPoolAbstractBlockLease = json.Unmarshal([]byte(serializedPoolAbstractBlockLease), &toSerialize)
	if errPoolAbstractBlockLease != nil {
		return []byte{}, errPoolAbstractBlockLease
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpType != nil {
		toSerialize["IpType"] = o.IpType
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.IpLeases != nil {
		toSerialize["IpLeases"] = o.IpLeases
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}
	if o.Vrf != nil {
		toSerialize["Vrf"] = o.Vrf
	}
	return json.Marshal(toSerialize)
}

type NullableIppoolBlockLease struct {
	value *IppoolBlockLease
	isSet bool
}

func (v NullableIppoolBlockLease) Get() *IppoolBlockLease {
	return v.value
}

func (v *NullableIppoolBlockLease) Set(val *IppoolBlockLease) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolBlockLease) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolBlockLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolBlockLease(val *IppoolBlockLease) *NullableIppoolBlockLease {
	return &NullableIppoolBlockLease{value: val, isSet: true}
}

func (v NullableIppoolBlockLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolBlockLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
