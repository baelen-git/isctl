/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HyperflexReplicationPeerInfo Information about the replication peer for this datastore.
type HyperflexReplicationPeerInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                                  `json:"ObjectType"`
	Datastores []HyperflexReplicationPlatDatastorePair `json:"Datastores,omitempty"`
	// Data Cluster IP for the replication peer.
	Dcip *string                          `json:"Dcip,omitempty"`
	Er   NullableHyperflexEntityReference `json:"Er,omitempty"`
	// Management Cluster IP for the replication peer.
	Mcip  *string                            `json:"Mcip,omitempty"`
	Ports []HyperflexPortTypeToPortNumberMap `json:"Ports,omitempty"`
	// Replication Cluster IP for the replication peer.
	ReplCip *string `json:"ReplCip,omitempty"`
	// Peer Cluster Status for the replication peer.
	Status *string `json:"Status,omitempty"`
	// Peer Cluster Status Details for the replication peer.
	StatusDetails        *string `json:"StatusDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexReplicationPeerInfo HyperflexReplicationPeerInfo

// NewHyperflexReplicationPeerInfo instantiates a new HyperflexReplicationPeerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexReplicationPeerInfo(classId string, objectType string) *HyperflexReplicationPeerInfo {
	this := HyperflexReplicationPeerInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexReplicationPeerInfoWithDefaults instantiates a new HyperflexReplicationPeerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexReplicationPeerInfoWithDefaults() *HyperflexReplicationPeerInfo {
	this := HyperflexReplicationPeerInfo{}
	var classId string = "hyperflex.ReplicationPeerInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.ReplicationPeerInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexReplicationPeerInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexReplicationPeerInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexReplicationPeerInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexReplicationPeerInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDatastores returns the Datastores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPeerInfo) GetDatastores() []HyperflexReplicationPlatDatastorePair {
	if o == nil {
		var ret []HyperflexReplicationPlatDatastorePair
		return ret
	}
	return o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPeerInfo) GetDatastoresOk() (*[]HyperflexReplicationPlatDatastorePair, bool) {
	if o == nil || o.Datastores == nil {
		return nil, false
	}
	return &o.Datastores, true
}

// HasDatastores returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfo) HasDatastores() bool {
	if o != nil && o.Datastores != nil {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []HyperflexReplicationPlatDatastorePair and assigns it to the Datastores field.
func (o *HyperflexReplicationPeerInfo) SetDatastores(v []HyperflexReplicationPlatDatastorePair) {
	o.Datastores = v
}

// GetDcip returns the Dcip field value if set, zero value otherwise.
func (o *HyperflexReplicationPeerInfo) GetDcip() string {
	if o == nil || o.Dcip == nil {
		var ret string
		return ret
	}
	return *o.Dcip
}

// GetDcipOk returns a tuple with the Dcip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfo) GetDcipOk() (*string, bool) {
	if o == nil || o.Dcip == nil {
		return nil, false
	}
	return o.Dcip, true
}

// HasDcip returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfo) HasDcip() bool {
	if o != nil && o.Dcip != nil {
		return true
	}

	return false
}

// SetDcip gets a reference to the given string and assigns it to the Dcip field.
func (o *HyperflexReplicationPeerInfo) SetDcip(v string) {
	o.Dcip = &v
}

// GetEr returns the Er field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPeerInfo) GetEr() HyperflexEntityReference {
	if o == nil || o.Er.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.Er.Get()
}

// GetErOk returns a tuple with the Er field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPeerInfo) GetErOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Er.Get(), o.Er.IsSet()
}

// HasEr returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfo) HasEr() bool {
	if o != nil && o.Er.IsSet() {
		return true
	}

	return false
}

// SetEr gets a reference to the given NullableHyperflexEntityReference and assigns it to the Er field.
func (o *HyperflexReplicationPeerInfo) SetEr(v HyperflexEntityReference) {
	o.Er.Set(&v)
}

// SetErNil sets the value for Er to be an explicit nil
func (o *HyperflexReplicationPeerInfo) SetErNil() {
	o.Er.Set(nil)
}

// UnsetEr ensures that no value is present for Er, not even an explicit nil
func (o *HyperflexReplicationPeerInfo) UnsetEr() {
	o.Er.Unset()
}

// GetMcip returns the Mcip field value if set, zero value otherwise.
func (o *HyperflexReplicationPeerInfo) GetMcip() string {
	if o == nil || o.Mcip == nil {
		var ret string
		return ret
	}
	return *o.Mcip
}

// GetMcipOk returns a tuple with the Mcip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfo) GetMcipOk() (*string, bool) {
	if o == nil || o.Mcip == nil {
		return nil, false
	}
	return o.Mcip, true
}

// HasMcip returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfo) HasMcip() bool {
	if o != nil && o.Mcip != nil {
		return true
	}

	return false
}

// SetMcip gets a reference to the given string and assigns it to the Mcip field.
func (o *HyperflexReplicationPeerInfo) SetMcip(v string) {
	o.Mcip = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPeerInfo) GetPorts() []HyperflexPortTypeToPortNumberMap {
	if o == nil {
		var ret []HyperflexPortTypeToPortNumberMap
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPeerInfo) GetPortsOk() (*[]HyperflexPortTypeToPortNumberMap, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return &o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfo) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []HyperflexPortTypeToPortNumberMap and assigns it to the Ports field.
func (o *HyperflexReplicationPeerInfo) SetPorts(v []HyperflexPortTypeToPortNumberMap) {
	o.Ports = v
}

// GetReplCip returns the ReplCip field value if set, zero value otherwise.
func (o *HyperflexReplicationPeerInfo) GetReplCip() string {
	if o == nil || o.ReplCip == nil {
		var ret string
		return ret
	}
	return *o.ReplCip
}

// GetReplCipOk returns a tuple with the ReplCip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfo) GetReplCipOk() (*string, bool) {
	if o == nil || o.ReplCip == nil {
		return nil, false
	}
	return o.ReplCip, true
}

// HasReplCip returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfo) HasReplCip() bool {
	if o != nil && o.ReplCip != nil {
		return true
	}

	return false
}

// SetReplCip gets a reference to the given string and assigns it to the ReplCip field.
func (o *HyperflexReplicationPeerInfo) SetReplCip(v string) {
	o.ReplCip = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexReplicationPeerInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexReplicationPeerInfo) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *HyperflexReplicationPeerInfo) GetStatusDetails() string {
	if o == nil || o.StatusDetails == nil {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPeerInfo) GetStatusDetailsOk() (*string, bool) {
	if o == nil || o.StatusDetails == nil {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *HyperflexReplicationPeerInfo) HasStatusDetails() bool {
	if o != nil && o.StatusDetails != nil {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *HyperflexReplicationPeerInfo) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

func (o HyperflexReplicationPeerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Datastores != nil {
		toSerialize["Datastores"] = o.Datastores
	}
	if o.Dcip != nil {
		toSerialize["Dcip"] = o.Dcip
	}
	if o.Er.IsSet() {
		toSerialize["Er"] = o.Er.Get()
	}
	if o.Mcip != nil {
		toSerialize["Mcip"] = o.Mcip
	}
	if o.Ports != nil {
		toSerialize["Ports"] = o.Ports
	}
	if o.ReplCip != nil {
		toSerialize["ReplCip"] = o.ReplCip
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusDetails != nil {
		toSerialize["StatusDetails"] = o.StatusDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexReplicationPeerInfo) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexReplicationPeerInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                  `json:"ObjectType"`
		Datastores []HyperflexReplicationPlatDatastorePair `json:"Datastores,omitempty"`
		// Data Cluster IP for the replication peer.
		Dcip *string                          `json:"Dcip,omitempty"`
		Er   NullableHyperflexEntityReference `json:"Er,omitempty"`
		// Management Cluster IP for the replication peer.
		Mcip  *string                            `json:"Mcip,omitempty"`
		Ports []HyperflexPortTypeToPortNumberMap `json:"Ports,omitempty"`
		// Replication Cluster IP for the replication peer.
		ReplCip *string `json:"ReplCip,omitempty"`
		// Peer Cluster Status for the replication peer.
		Status *string `json:"Status,omitempty"`
		// Peer Cluster Status Details for the replication peer.
		StatusDetails *string `json:"StatusDetails,omitempty"`
	}

	varHyperflexReplicationPeerInfoWithoutEmbeddedStruct := HyperflexReplicationPeerInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexReplicationPeerInfoWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexReplicationPeerInfo := _HyperflexReplicationPeerInfo{}
		varHyperflexReplicationPeerInfo.ClassId = varHyperflexReplicationPeerInfoWithoutEmbeddedStruct.ClassId
		varHyperflexReplicationPeerInfo.ObjectType = varHyperflexReplicationPeerInfoWithoutEmbeddedStruct.ObjectType
		varHyperflexReplicationPeerInfo.Datastores = varHyperflexReplicationPeerInfoWithoutEmbeddedStruct.Datastores
		varHyperflexReplicationPeerInfo.Dcip = varHyperflexReplicationPeerInfoWithoutEmbeddedStruct.Dcip
		varHyperflexReplicationPeerInfo.Er = varHyperflexReplicationPeerInfoWithoutEmbeddedStruct.Er
		varHyperflexReplicationPeerInfo.Mcip = varHyperflexReplicationPeerInfoWithoutEmbeddedStruct.Mcip
		varHyperflexReplicationPeerInfo.Ports = varHyperflexReplicationPeerInfoWithoutEmbeddedStruct.Ports
		varHyperflexReplicationPeerInfo.ReplCip = varHyperflexReplicationPeerInfoWithoutEmbeddedStruct.ReplCip
		varHyperflexReplicationPeerInfo.Status = varHyperflexReplicationPeerInfoWithoutEmbeddedStruct.Status
		varHyperflexReplicationPeerInfo.StatusDetails = varHyperflexReplicationPeerInfoWithoutEmbeddedStruct.StatusDetails
		*o = HyperflexReplicationPeerInfo(varHyperflexReplicationPeerInfo)
	} else {
		return err
	}

	varHyperflexReplicationPeerInfo := _HyperflexReplicationPeerInfo{}

	err = json.Unmarshal(bytes, &varHyperflexReplicationPeerInfo)
	if err == nil {
		o.MoBaseComplexType = varHyperflexReplicationPeerInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Datastores")
		delete(additionalProperties, "Dcip")
		delete(additionalProperties, "Er")
		delete(additionalProperties, "Mcip")
		delete(additionalProperties, "Ports")
		delete(additionalProperties, "ReplCip")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusDetails")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexReplicationPeerInfo struct {
	value *HyperflexReplicationPeerInfo
	isSet bool
}

func (v NullableHyperflexReplicationPeerInfo) Get() *HyperflexReplicationPeerInfo {
	return v.value
}

func (v *NullableHyperflexReplicationPeerInfo) Set(val *HyperflexReplicationPeerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexReplicationPeerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexReplicationPeerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexReplicationPeerInfo(val *HyperflexReplicationPeerInfo) *NullableHyperflexReplicationPeerInfo {
	return &NullableHyperflexReplicationPeerInfo{value: val, isSet: true}
}

func (v NullableHyperflexReplicationPeerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexReplicationPeerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
