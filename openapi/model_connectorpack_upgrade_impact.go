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

// ConnectorpackUpgradeImpact Used to determine the list of connector packs to be installed on a target UCS Director in its next upgrade cycle. Accepts the moid of the target UcsdInfo as part of the filter query. Given below is a sample url :- https://{{target}}/api/v1/connectorpack/UpgradeImpacts?$filter= ( UcsdInfo.Moid eq <<MoId>> ).
type ConnectorpackUpgradeImpact struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                             `json:"ObjectType" yaml:"ObjectType"`
	ConnectorPack []ConnectorpackConnectorPackUpdate `json:"ConnectorPack,omitempty" yaml:"ConnectorPack,omitempty"`
	// States whether the UCS Director is eligible for an upgrade. Set to true if connector packs are available for upgrade, else set to false.
	IsEligibleForUpgrade *bool `json:"IsEligibleForUpgrade,omitempty" yaml:"IsEligibleForUpgrade,omitempty"`
	// States whether all the requisite updates have been downloaded to the target UCS Director. Set to true if all connector packs required to upgrade UCS Director to the next iteration have been downloaded, else set to false.
	IsUpdateDownloaded *bool                     `json:"IsUpdateDownloaded,omitempty" yaml:"IsUpdateDownloaded,omitempty"`
	UcsdInfo           *IaasUcsdInfoRelationship `json:"UcsdInfo,omitempty" yaml:"UcsdInfo,omitempty"`
}

// NewConnectorpackUpgradeImpact instantiates a new ConnectorpackUpgradeImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorpackUpgradeImpact(classId string, objectType string) *ConnectorpackUpgradeImpact {
	this := ConnectorpackUpgradeImpact{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorpackUpgradeImpactWithDefaults instantiates a new ConnectorpackUpgradeImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorpackUpgradeImpactWithDefaults() *ConnectorpackUpgradeImpact {
	this := ConnectorpackUpgradeImpact{}
	var classId string = "connectorpack.UpgradeImpact"
	this.ClassId = classId
	var objectType string = "connectorpack.UpgradeImpact"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorpackUpgradeImpact) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorpackUpgradeImpact) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorpackUpgradeImpact) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorpackUpgradeImpact) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorpackUpgradeImpact) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorpackUpgradeImpact) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectorPack returns the ConnectorPack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorpackUpgradeImpact) GetConnectorPack() []ConnectorpackConnectorPackUpdate {
	if o == nil {
		var ret []ConnectorpackConnectorPackUpdate
		return ret
	}
	return o.ConnectorPack
}

// GetConnectorPackOk returns a tuple with the ConnectorPack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorpackUpgradeImpact) GetConnectorPackOk() (*[]ConnectorpackConnectorPackUpdate, bool) {
	if o == nil || o.ConnectorPack == nil {
		return nil, false
	}
	return &o.ConnectorPack, true
}

// HasConnectorPack returns a boolean if a field has been set.
func (o *ConnectorpackUpgradeImpact) HasConnectorPack() bool {
	if o != nil && o.ConnectorPack != nil {
		return true
	}

	return false
}

// SetConnectorPack gets a reference to the given []ConnectorpackConnectorPackUpdate and assigns it to the ConnectorPack field.
func (o *ConnectorpackUpgradeImpact) SetConnectorPack(v []ConnectorpackConnectorPackUpdate) {
	o.ConnectorPack = v
}

// GetIsEligibleForUpgrade returns the IsEligibleForUpgrade field value if set, zero value otherwise.
func (o *ConnectorpackUpgradeImpact) GetIsEligibleForUpgrade() bool {
	if o == nil || o.IsEligibleForUpgrade == nil {
		var ret bool
		return ret
	}
	return *o.IsEligibleForUpgrade
}

// GetIsEligibleForUpgradeOk returns a tuple with the IsEligibleForUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackUpgradeImpact) GetIsEligibleForUpgradeOk() (*bool, bool) {
	if o == nil || o.IsEligibleForUpgrade == nil {
		return nil, false
	}
	return o.IsEligibleForUpgrade, true
}

// HasIsEligibleForUpgrade returns a boolean if a field has been set.
func (o *ConnectorpackUpgradeImpact) HasIsEligibleForUpgrade() bool {
	if o != nil && o.IsEligibleForUpgrade != nil {
		return true
	}

	return false
}

// SetIsEligibleForUpgrade gets a reference to the given bool and assigns it to the IsEligibleForUpgrade field.
func (o *ConnectorpackUpgradeImpact) SetIsEligibleForUpgrade(v bool) {
	o.IsEligibleForUpgrade = &v
}

// GetIsUpdateDownloaded returns the IsUpdateDownloaded field value if set, zero value otherwise.
func (o *ConnectorpackUpgradeImpact) GetIsUpdateDownloaded() bool {
	if o == nil || o.IsUpdateDownloaded == nil {
		var ret bool
		return ret
	}
	return *o.IsUpdateDownloaded
}

// GetIsUpdateDownloadedOk returns a tuple with the IsUpdateDownloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackUpgradeImpact) GetIsUpdateDownloadedOk() (*bool, bool) {
	if o == nil || o.IsUpdateDownloaded == nil {
		return nil, false
	}
	return o.IsUpdateDownloaded, true
}

// HasIsUpdateDownloaded returns a boolean if a field has been set.
func (o *ConnectorpackUpgradeImpact) HasIsUpdateDownloaded() bool {
	if o != nil && o.IsUpdateDownloaded != nil {
		return true
	}

	return false
}

// SetIsUpdateDownloaded gets a reference to the given bool and assigns it to the IsUpdateDownloaded field.
func (o *ConnectorpackUpgradeImpact) SetIsUpdateDownloaded(v bool) {
	o.IsUpdateDownloaded = &v
}

// GetUcsdInfo returns the UcsdInfo field value if set, zero value otherwise.
func (o *ConnectorpackUpgradeImpact) GetUcsdInfo() IaasUcsdInfoRelationship {
	if o == nil || o.UcsdInfo == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.UcsdInfo
}

// GetUcsdInfoOk returns a tuple with the UcsdInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackUpgradeImpact) GetUcsdInfoOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.UcsdInfo == nil {
		return nil, false
	}
	return o.UcsdInfo, true
}

// HasUcsdInfo returns a boolean if a field has been set.
func (o *ConnectorpackUpgradeImpact) HasUcsdInfo() bool {
	if o != nil && o.UcsdInfo != nil {
		return true
	}

	return false
}

// SetUcsdInfo gets a reference to the given IaasUcsdInfoRelationship and assigns it to the UcsdInfo field.
func (o *ConnectorpackUpgradeImpact) SetUcsdInfo(v IaasUcsdInfoRelationship) {
	o.UcsdInfo = &v
}

func (o ConnectorpackUpgradeImpact) MarshalJSON() ([]byte, error) {
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
	if o.ConnectorPack != nil {
		toSerialize["ConnectorPack"] = o.ConnectorPack
	}
	if o.IsEligibleForUpgrade != nil {
		toSerialize["IsEligibleForUpgrade"] = o.IsEligibleForUpgrade
	}
	if o.IsUpdateDownloaded != nil {
		toSerialize["IsUpdateDownloaded"] = o.IsUpdateDownloaded
	}
	if o.UcsdInfo != nil {
		toSerialize["UcsdInfo"] = o.UcsdInfo
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorpackUpgradeImpact struct {
	value *ConnectorpackUpgradeImpact
	isSet bool
}

func (v NullableConnectorpackUpgradeImpact) Get() *ConnectorpackUpgradeImpact {
	return v.value
}

func (v *NullableConnectorpackUpgradeImpact) Set(val *ConnectorpackUpgradeImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorpackUpgradeImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorpackUpgradeImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorpackUpgradeImpact(val *ConnectorpackUpgradeImpact) *NullableConnectorpackUpgradeImpact {
	return &NullableConnectorpackUpgradeImpact{value: val, isSet: true}
}

func (v NullableConnectorpackUpgradeImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorpackUpgradeImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
