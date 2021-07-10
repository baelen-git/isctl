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
)

// HyperflexReplicationPlatDatastoreAllOf Definition of the list of properties defined in 'hyperflex.ReplicationPlatDatastore', excluding properties defined in parent classes.
type HyperflexReplicationPlatDatastoreAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                           `json:"ObjectType"`
	ClusterEr            NullableHyperflexEntityReference `json:"ClusterEr,omitempty"`
	DatastoreEr          NullableHyperflexEntityReference `json:"DatastoreEr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexReplicationPlatDatastoreAllOf HyperflexReplicationPlatDatastoreAllOf

// NewHyperflexReplicationPlatDatastoreAllOf instantiates a new HyperflexReplicationPlatDatastoreAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexReplicationPlatDatastoreAllOf(classId string, objectType string) *HyperflexReplicationPlatDatastoreAllOf {
	this := HyperflexReplicationPlatDatastoreAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexReplicationPlatDatastoreAllOfWithDefaults instantiates a new HyperflexReplicationPlatDatastoreAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexReplicationPlatDatastoreAllOfWithDefaults() *HyperflexReplicationPlatDatastoreAllOf {
	this := HyperflexReplicationPlatDatastoreAllOf{}
	var classId string = "hyperflex.ReplicationPlatDatastore"
	this.ClassId = classId
	var objectType string = "hyperflex.ReplicationPlatDatastore"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexReplicationPlatDatastoreAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPlatDatastoreAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexReplicationPlatDatastoreAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexReplicationPlatDatastoreAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPlatDatastoreAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexReplicationPlatDatastoreAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterEr returns the ClusterEr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPlatDatastoreAllOf) GetClusterEr() HyperflexEntityReference {
	if o == nil || o.ClusterEr.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.ClusterEr.Get()
}

// GetClusterErOk returns a tuple with the ClusterEr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPlatDatastoreAllOf) GetClusterErOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterEr.Get(), o.ClusterEr.IsSet()
}

// HasClusterEr returns a boolean if a field has been set.
func (o *HyperflexReplicationPlatDatastoreAllOf) HasClusterEr() bool {
	if o != nil && o.ClusterEr.IsSet() {
		return true
	}

	return false
}

// SetClusterEr gets a reference to the given NullableHyperflexEntityReference and assigns it to the ClusterEr field.
func (o *HyperflexReplicationPlatDatastoreAllOf) SetClusterEr(v HyperflexEntityReference) {
	o.ClusterEr.Set(&v)
}

// SetClusterErNil sets the value for ClusterEr to be an explicit nil
func (o *HyperflexReplicationPlatDatastoreAllOf) SetClusterErNil() {
	o.ClusterEr.Set(nil)
}

// UnsetClusterEr ensures that no value is present for ClusterEr, not even an explicit nil
func (o *HyperflexReplicationPlatDatastoreAllOf) UnsetClusterEr() {
	o.ClusterEr.Unset()
}

// GetDatastoreEr returns the DatastoreEr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPlatDatastoreAllOf) GetDatastoreEr() HyperflexEntityReference {
	if o == nil || o.DatastoreEr.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.DatastoreEr.Get()
}

// GetDatastoreErOk returns a tuple with the DatastoreEr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPlatDatastoreAllOf) GetDatastoreErOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatastoreEr.Get(), o.DatastoreEr.IsSet()
}

// HasDatastoreEr returns a boolean if a field has been set.
func (o *HyperflexReplicationPlatDatastoreAllOf) HasDatastoreEr() bool {
	if o != nil && o.DatastoreEr.IsSet() {
		return true
	}

	return false
}

// SetDatastoreEr gets a reference to the given NullableHyperflexEntityReference and assigns it to the DatastoreEr field.
func (o *HyperflexReplicationPlatDatastoreAllOf) SetDatastoreEr(v HyperflexEntityReference) {
	o.DatastoreEr.Set(&v)
}

// SetDatastoreErNil sets the value for DatastoreEr to be an explicit nil
func (o *HyperflexReplicationPlatDatastoreAllOf) SetDatastoreErNil() {
	o.DatastoreEr.Set(nil)
}

// UnsetDatastoreEr ensures that no value is present for DatastoreEr, not even an explicit nil
func (o *HyperflexReplicationPlatDatastoreAllOf) UnsetDatastoreEr() {
	o.DatastoreEr.Unset()
}

func (o HyperflexReplicationPlatDatastoreAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterEr.IsSet() {
		toSerialize["ClusterEr"] = o.ClusterEr.Get()
	}
	if o.DatastoreEr.IsSet() {
		toSerialize["DatastoreEr"] = o.DatastoreEr.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexReplicationPlatDatastoreAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexReplicationPlatDatastoreAllOf := _HyperflexReplicationPlatDatastoreAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexReplicationPlatDatastoreAllOf); err == nil {
		*o = HyperflexReplicationPlatDatastoreAllOf(varHyperflexReplicationPlatDatastoreAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterEr")
		delete(additionalProperties, "DatastoreEr")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexReplicationPlatDatastoreAllOf struct {
	value *HyperflexReplicationPlatDatastoreAllOf
	isSet bool
}

func (v NullableHyperflexReplicationPlatDatastoreAllOf) Get() *HyperflexReplicationPlatDatastoreAllOf {
	return v.value
}

func (v *NullableHyperflexReplicationPlatDatastoreAllOf) Set(val *HyperflexReplicationPlatDatastoreAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexReplicationPlatDatastoreAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexReplicationPlatDatastoreAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexReplicationPlatDatastoreAllOf(val *HyperflexReplicationPlatDatastoreAllOf) *NullableHyperflexReplicationPlatDatastoreAllOf {
	return &NullableHyperflexReplicationPlatDatastoreAllOf{value: val, isSet: true}
}

func (v NullableHyperflexReplicationPlatDatastoreAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexReplicationPlatDatastoreAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
