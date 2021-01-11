/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-12-08T20:53:20Z.
 *
 * API version: 1.0.9-2908
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// VirtualizationVmwareVmMemoryShareInfoAllOf Definition of the list of properties defined in 'virtualization.VmwareVmMemoryShareInfo', excluding properties defined in parent classes.
type VirtualizationVmwareVmMemoryShareInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Limit on the memory sharing imposed (in Mbytes).
	MemLimit *int64 `json:"MemLimit,omitempty" yaml:"MemLimit,omitempty"`
	// Limit on memory overhead imposed (in Mbytes).
	MemOverheadLimit *int64 `json:"MemOverheadLimit,omitempty" yaml:"MemOverheadLimit,omitempty"`
	// Similar to CPU reservations (Mbytes).
	MemReservation *int64 `json:"MemReservation,omitempty" yaml:"MemReservation,omitempty"`
	// Similar to CPU Shares but applicable to memory. There is no unit for this value. It is a relative measure based on the settings for other resource pools.
	MemShares *int64 `json:"MemShares,omitempty" yaml:"MemShares,omitempty"`
}

// NewVirtualizationVmwareVmMemoryShareInfoAllOf instantiates a new VirtualizationVmwareVmMemoryShareInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVmMemoryShareInfoAllOf(classId string, objectType string) *VirtualizationVmwareVmMemoryShareInfoAllOf {
	this := VirtualizationVmwareVmMemoryShareInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVmMemoryShareInfoAllOfWithDefaults instantiates a new VirtualizationVmwareVmMemoryShareInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVmMemoryShareInfoAllOfWithDefaults() *VirtualizationVmwareVmMemoryShareInfoAllOf {
	this := VirtualizationVmwareVmMemoryShareInfoAllOf{}
	var classId string = "virtualization.VmwareVmMemoryShareInfo"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVmMemoryShareInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMemLimit returns the MemLimit field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemLimit() int64 {
	if o == nil || o.MemLimit == nil {
		var ret int64
		return ret
	}
	return *o.MemLimit
}

// GetMemLimitOk returns a tuple with the MemLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemLimitOk() (*int64, bool) {
	if o == nil || o.MemLimit == nil {
		return nil, false
	}
	return o.MemLimit, true
}

// HasMemLimit returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) HasMemLimit() bool {
	if o != nil && o.MemLimit != nil {
		return true
	}

	return false
}

// SetMemLimit gets a reference to the given int64 and assigns it to the MemLimit field.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) SetMemLimit(v int64) {
	o.MemLimit = &v
}

// GetMemOverheadLimit returns the MemOverheadLimit field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemOverheadLimit() int64 {
	if o == nil || o.MemOverheadLimit == nil {
		var ret int64
		return ret
	}
	return *o.MemOverheadLimit
}

// GetMemOverheadLimitOk returns a tuple with the MemOverheadLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemOverheadLimitOk() (*int64, bool) {
	if o == nil || o.MemOverheadLimit == nil {
		return nil, false
	}
	return o.MemOverheadLimit, true
}

// HasMemOverheadLimit returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) HasMemOverheadLimit() bool {
	if o != nil && o.MemOverheadLimit != nil {
		return true
	}

	return false
}

// SetMemOverheadLimit gets a reference to the given int64 and assigns it to the MemOverheadLimit field.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) SetMemOverheadLimit(v int64) {
	o.MemOverheadLimit = &v
}

// GetMemReservation returns the MemReservation field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemReservation() int64 {
	if o == nil || o.MemReservation == nil {
		var ret int64
		return ret
	}
	return *o.MemReservation
}

// GetMemReservationOk returns a tuple with the MemReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemReservationOk() (*int64, bool) {
	if o == nil || o.MemReservation == nil {
		return nil, false
	}
	return o.MemReservation, true
}

// HasMemReservation returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) HasMemReservation() bool {
	if o != nil && o.MemReservation != nil {
		return true
	}

	return false
}

// SetMemReservation gets a reference to the given int64 and assigns it to the MemReservation field.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) SetMemReservation(v int64) {
	o.MemReservation = &v
}

// GetMemShares returns the MemShares field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemShares() int64 {
	if o == nil || o.MemShares == nil {
		var ret int64
		return ret
	}
	return *o.MemShares
}

// GetMemSharesOk returns a tuple with the MemShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) GetMemSharesOk() (*int64, bool) {
	if o == nil || o.MemShares == nil {
		return nil, false
	}
	return o.MemShares, true
}

// HasMemShares returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) HasMemShares() bool {
	if o != nil && o.MemShares != nil {
		return true
	}

	return false
}

// SetMemShares gets a reference to the given int64 and assigns it to the MemShares field.
func (o *VirtualizationVmwareVmMemoryShareInfoAllOf) SetMemShares(v int64) {
	o.MemShares = &v
}

func (o VirtualizationVmwareVmMemoryShareInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MemLimit != nil {
		toSerialize["MemLimit"] = o.MemLimit
	}
	if o.MemOverheadLimit != nil {
		toSerialize["MemOverheadLimit"] = o.MemOverheadLimit
	}
	if o.MemReservation != nil {
		toSerialize["MemReservation"] = o.MemReservation
	}
	if o.MemShares != nil {
		toSerialize["MemShares"] = o.MemShares
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationVmwareVmMemoryShareInfoAllOf struct {
	value *VirtualizationVmwareVmMemoryShareInfoAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareVmMemoryShareInfoAllOf) Get() *VirtualizationVmwareVmMemoryShareInfoAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareVmMemoryShareInfoAllOf) Set(val *VirtualizationVmwareVmMemoryShareInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVmMemoryShareInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVmMemoryShareInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVmMemoryShareInfoAllOf(val *VirtualizationVmwareVmMemoryShareInfoAllOf) *NullableVirtualizationVmwareVmMemoryShareInfoAllOf {
	return &NullableVirtualizationVmwareVmMemoryShareInfoAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVmMemoryShareInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVmMemoryShareInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
