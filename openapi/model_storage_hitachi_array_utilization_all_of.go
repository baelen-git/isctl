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

// StorageHitachiArrayUtilizationAllOf Definition of the list of properties defined in 'storage.HitachiArrayUtilization', excluding properties defined in parent classes.
type StorageHitachiArrayUtilizationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array's flash modules.
	DataReduction *float32 `json:"DataReduction,omitempty"`
	// Total provisioned storage capacity in Hitachi Virtual Storage, represented in bytes.
	Provisioned          *int64 `json:"Provisioned,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiArrayUtilizationAllOf StorageHitachiArrayUtilizationAllOf

// NewStorageHitachiArrayUtilizationAllOf instantiates a new StorageHitachiArrayUtilizationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiArrayUtilizationAllOf(classId string, objectType string) *StorageHitachiArrayUtilizationAllOf {
	this := StorageHitachiArrayUtilizationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiArrayUtilizationAllOfWithDefaults instantiates a new StorageHitachiArrayUtilizationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiArrayUtilizationAllOfWithDefaults() *StorageHitachiArrayUtilizationAllOf {
	this := StorageHitachiArrayUtilizationAllOf{}
	var classId string = "storage.HitachiArrayUtilization"
	this.ClassId = classId
	var objectType string = "storage.HitachiArrayUtilization"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiArrayUtilizationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiArrayUtilizationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiArrayUtilizationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiArrayUtilizationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiArrayUtilizationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiArrayUtilizationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataReduction returns the DataReduction field value if set, zero value otherwise.
func (o *StorageHitachiArrayUtilizationAllOf) GetDataReduction() float32 {
	if o == nil || o.DataReduction == nil {
		var ret float32
		return ret
	}
	return *o.DataReduction
}

// GetDataReductionOk returns a tuple with the DataReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArrayUtilizationAllOf) GetDataReductionOk() (*float32, bool) {
	if o == nil || o.DataReduction == nil {
		return nil, false
	}
	return o.DataReduction, true
}

// HasDataReduction returns a boolean if a field has been set.
func (o *StorageHitachiArrayUtilizationAllOf) HasDataReduction() bool {
	if o != nil && o.DataReduction != nil {
		return true
	}

	return false
}

// SetDataReduction gets a reference to the given float32 and assigns it to the DataReduction field.
func (o *StorageHitachiArrayUtilizationAllOf) SetDataReduction(v float32) {
	o.DataReduction = &v
}

// GetProvisioned returns the Provisioned field value if set, zero value otherwise.
func (o *StorageHitachiArrayUtilizationAllOf) GetProvisioned() int64 {
	if o == nil || o.Provisioned == nil {
		var ret int64
		return ret
	}
	return *o.Provisioned
}

// GetProvisionedOk returns a tuple with the Provisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArrayUtilizationAllOf) GetProvisionedOk() (*int64, bool) {
	if o == nil || o.Provisioned == nil {
		return nil, false
	}
	return o.Provisioned, true
}

// HasProvisioned returns a boolean if a field has been set.
func (o *StorageHitachiArrayUtilizationAllOf) HasProvisioned() bool {
	if o != nil && o.Provisioned != nil {
		return true
	}

	return false
}

// SetProvisioned gets a reference to the given int64 and assigns it to the Provisioned field.
func (o *StorageHitachiArrayUtilizationAllOf) SetProvisioned(v int64) {
	o.Provisioned = &v
}

func (o StorageHitachiArrayUtilizationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DataReduction != nil {
		toSerialize["DataReduction"] = o.DataReduction
	}
	if o.Provisioned != nil {
		toSerialize["Provisioned"] = o.Provisioned
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiArrayUtilizationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageHitachiArrayUtilizationAllOf := _StorageHitachiArrayUtilizationAllOf{}

	if err = json.Unmarshal(bytes, &varStorageHitachiArrayUtilizationAllOf); err == nil {
		*o = StorageHitachiArrayUtilizationAllOf(varStorageHitachiArrayUtilizationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DataReduction")
		delete(additionalProperties, "Provisioned")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiArrayUtilizationAllOf struct {
	value *StorageHitachiArrayUtilizationAllOf
	isSet bool
}

func (v NullableStorageHitachiArrayUtilizationAllOf) Get() *StorageHitachiArrayUtilizationAllOf {
	return v.value
}

func (v *NullableStorageHitachiArrayUtilizationAllOf) Set(val *StorageHitachiArrayUtilizationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiArrayUtilizationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiArrayUtilizationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiArrayUtilizationAllOf(val *StorageHitachiArrayUtilizationAllOf) *NullableStorageHitachiArrayUtilizationAllOf {
	return &NullableStorageHitachiArrayUtilizationAllOf{value: val, isSet: true}
}

func (v NullableStorageHitachiArrayUtilizationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiArrayUtilizationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
