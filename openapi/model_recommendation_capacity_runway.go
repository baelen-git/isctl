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

// RecommendationCapacityRunway Entity representing the new capacity runway based on recommendations.
type RecommendationCapacityRunway struct {
	RecommendationBase `yaml:"RecommendationBase,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Additional capacity is the capacity which is needed more after exhausing all hardware on current cluster.
	AdditionalCapacity *int64 `json:"AdditionalCapacity,omitempty" yaml:"AdditionalCapacity,omitempty"`
	// Number of months in future for which recommendation is provided for.
	Period *int64 `json:"Period,omitempty" yaml:"Period,omitempty"`
	// This represents the new runway, that is the number of days remaining before the cluster's storage utilization reaches the recommended capacity limit after the recommended hardware is added.
	Runway *int64 `json:"Runway,omitempty" yaml:"Runway,omitempty"`
	// Total capacity of the cluster after the recommended hardware is added.
	TotalCapacity *int64 `json:"TotalCapacity,omitempty" yaml:"TotalCapacity,omitempty"`
	// Unit for the new capacity. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes.
	Unit             *string                       `json:"Unit,omitempty" yaml:"Unit,omitempty"`
	ForecastInstance *ForecastInstanceRelationship `json:"ForecastInstance,omitempty" yaml:"ForecastInstance,omitempty"`
	// An array of relationships to recommendationPhysicalItem resources.
	PhysicalItem     []RecommendationPhysicalItemRelationship `json:"PhysicalItem,omitempty" yaml:"PhysicalItem,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship     `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewRecommendationCapacityRunway instantiates a new RecommendationCapacityRunway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationCapacityRunway(classId string, objectType string) *RecommendationCapacityRunway {
	this := RecommendationCapacityRunway{}
	this.ClassId = classId
	this.ObjectType = objectType
	var unit string = "TB"
	this.Unit = &unit
	return &this
}

// NewRecommendationCapacityRunwayWithDefaults instantiates a new RecommendationCapacityRunway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationCapacityRunwayWithDefaults() *RecommendationCapacityRunway {
	this := RecommendationCapacityRunway{}
	var classId string = "recommendation.CapacityRunway"
	this.ClassId = classId
	var objectType string = "recommendation.CapacityRunway"
	this.ObjectType = objectType
	var unit string = "TB"
	this.Unit = &unit
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationCapacityRunway) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunway) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationCapacityRunway) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationCapacityRunway) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunway) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationCapacityRunway) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdditionalCapacity returns the AdditionalCapacity field value if set, zero value otherwise.
func (o *RecommendationCapacityRunway) GetAdditionalCapacity() int64 {
	if o == nil || o.AdditionalCapacity == nil {
		var ret int64
		return ret
	}
	return *o.AdditionalCapacity
}

// GetAdditionalCapacityOk returns a tuple with the AdditionalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunway) GetAdditionalCapacityOk() (*int64, bool) {
	if o == nil || o.AdditionalCapacity == nil {
		return nil, false
	}
	return o.AdditionalCapacity, true
}

// HasAdditionalCapacity returns a boolean if a field has been set.
func (o *RecommendationCapacityRunway) HasAdditionalCapacity() bool {
	if o != nil && o.AdditionalCapacity != nil {
		return true
	}

	return false
}

// SetAdditionalCapacity gets a reference to the given int64 and assigns it to the AdditionalCapacity field.
func (o *RecommendationCapacityRunway) SetAdditionalCapacity(v int64) {
	o.AdditionalCapacity = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *RecommendationCapacityRunway) GetPeriod() int64 {
	if o == nil || o.Period == nil {
		var ret int64
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunway) GetPeriodOk() (*int64, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *RecommendationCapacityRunway) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int64 and assigns it to the Period field.
func (o *RecommendationCapacityRunway) SetPeriod(v int64) {
	o.Period = &v
}

// GetRunway returns the Runway field value if set, zero value otherwise.
func (o *RecommendationCapacityRunway) GetRunway() int64 {
	if o == nil || o.Runway == nil {
		var ret int64
		return ret
	}
	return *o.Runway
}

// GetRunwayOk returns a tuple with the Runway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunway) GetRunwayOk() (*int64, bool) {
	if o == nil || o.Runway == nil {
		return nil, false
	}
	return o.Runway, true
}

// HasRunway returns a boolean if a field has been set.
func (o *RecommendationCapacityRunway) HasRunway() bool {
	if o != nil && o.Runway != nil {
		return true
	}

	return false
}

// SetRunway gets a reference to the given int64 and assigns it to the Runway field.
func (o *RecommendationCapacityRunway) SetRunway(v int64) {
	o.Runway = &v
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise.
func (o *RecommendationCapacityRunway) GetTotalCapacity() int64 {
	if o == nil || o.TotalCapacity == nil {
		var ret int64
		return ret
	}
	return *o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunway) GetTotalCapacityOk() (*int64, bool) {
	if o == nil || o.TotalCapacity == nil {
		return nil, false
	}
	return o.TotalCapacity, true
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *RecommendationCapacityRunway) HasTotalCapacity() bool {
	if o != nil && o.TotalCapacity != nil {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given int64 and assigns it to the TotalCapacity field.
func (o *RecommendationCapacityRunway) SetTotalCapacity(v int64) {
	o.TotalCapacity = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *RecommendationCapacityRunway) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunway) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *RecommendationCapacityRunway) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *RecommendationCapacityRunway) SetUnit(v string) {
	o.Unit = &v
}

// GetForecastInstance returns the ForecastInstance field value if set, zero value otherwise.
func (o *RecommendationCapacityRunway) GetForecastInstance() ForecastInstanceRelationship {
	if o == nil || o.ForecastInstance == nil {
		var ret ForecastInstanceRelationship
		return ret
	}
	return *o.ForecastInstance
}

// GetForecastInstanceOk returns a tuple with the ForecastInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunway) GetForecastInstanceOk() (*ForecastInstanceRelationship, bool) {
	if o == nil || o.ForecastInstance == nil {
		return nil, false
	}
	return o.ForecastInstance, true
}

// HasForecastInstance returns a boolean if a field has been set.
func (o *RecommendationCapacityRunway) HasForecastInstance() bool {
	if o != nil && o.ForecastInstance != nil {
		return true
	}

	return false
}

// SetForecastInstance gets a reference to the given ForecastInstanceRelationship and assigns it to the ForecastInstance field.
func (o *RecommendationCapacityRunway) SetForecastInstance(v ForecastInstanceRelationship) {
	o.ForecastInstance = &v
}

// GetPhysicalItem returns the PhysicalItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationCapacityRunway) GetPhysicalItem() []RecommendationPhysicalItemRelationship {
	if o == nil {
		var ret []RecommendationPhysicalItemRelationship
		return ret
	}
	return o.PhysicalItem
}

// GetPhysicalItemOk returns a tuple with the PhysicalItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationCapacityRunway) GetPhysicalItemOk() (*[]RecommendationPhysicalItemRelationship, bool) {
	if o == nil || o.PhysicalItem == nil {
		return nil, false
	}
	return &o.PhysicalItem, true
}

// HasPhysicalItem returns a boolean if a field has been set.
func (o *RecommendationCapacityRunway) HasPhysicalItem() bool {
	if o != nil && o.PhysicalItem != nil {
		return true
	}

	return false
}

// SetPhysicalItem gets a reference to the given []RecommendationPhysicalItemRelationship and assigns it to the PhysicalItem field.
func (o *RecommendationCapacityRunway) SetPhysicalItem(v []RecommendationPhysicalItemRelationship) {
	o.PhysicalItem = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *RecommendationCapacityRunway) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunway) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *RecommendationCapacityRunway) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *RecommendationCapacityRunway) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o RecommendationCapacityRunway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRecommendationBase, errRecommendationBase := json.Marshal(o.RecommendationBase)
	if errRecommendationBase != nil {
		return []byte{}, errRecommendationBase
	}
	errRecommendationBase = json.Unmarshal([]byte(serializedRecommendationBase), &toSerialize)
	if errRecommendationBase != nil {
		return []byte{}, errRecommendationBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdditionalCapacity != nil {
		toSerialize["AdditionalCapacity"] = o.AdditionalCapacity
	}
	if o.Period != nil {
		toSerialize["Period"] = o.Period
	}
	if o.Runway != nil {
		toSerialize["Runway"] = o.Runway
	}
	if o.TotalCapacity != nil {
		toSerialize["TotalCapacity"] = o.TotalCapacity
	}
	if o.Unit != nil {
		toSerialize["Unit"] = o.Unit
	}
	if o.ForecastInstance != nil {
		toSerialize["ForecastInstance"] = o.ForecastInstance
	}
	if o.PhysicalItem != nil {
		toSerialize["PhysicalItem"] = o.PhysicalItem
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableRecommendationCapacityRunway struct {
	value *RecommendationCapacityRunway
	isSet bool
}

func (v NullableRecommendationCapacityRunway) Get() *RecommendationCapacityRunway {
	return v.value
}

func (v *NullableRecommendationCapacityRunway) Set(val *RecommendationCapacityRunway) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationCapacityRunway) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationCapacityRunway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationCapacityRunway(val *RecommendationCapacityRunway) *NullableRecommendationCapacityRunway {
	return &NullableRecommendationCapacityRunway{value: val, isSet: true}
}

func (v NullableRecommendationCapacityRunway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationCapacityRunway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
