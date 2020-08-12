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

// FabricPortIdentifier Type to represent switch ports.
type FabricPortIdentifier struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// Breakout port Identifier of the Switch Interface.
	AggregatePortId *int64 `json:"AggregatePortId,omitempty" yaml:"AggregatePortId,omitempty"`
	// Port Identifier of the Switch/FEX/Chassis Interface.
	PortId *int64 `json:"PortId,omitempty" yaml:"PortId,omitempty"`
	// Slot Identifier of the Switch/FEX/Chassis Interface.
	SlotId *int64 `json:"SlotId,omitempty" yaml:"SlotId,omitempty"`
}

// NewFabricPortIdentifier instantiates a new FabricPortIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPortIdentifier() *FabricPortIdentifier {
	this := FabricPortIdentifier{}
	return &this
}

// NewFabricPortIdentifierWithDefaults instantiates a new FabricPortIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPortIdentifierWithDefaults() *FabricPortIdentifier {
	this := FabricPortIdentifier{}
	return &this
}

// GetAggregatePortId returns the AggregatePortId field value if set, zero value otherwise.
func (o *FabricPortIdentifier) GetAggregatePortId() int64 {
	if o == nil || o.AggregatePortId == nil {
		var ret int64
		return ret
	}
	return *o.AggregatePortId
}

// GetAggregatePortIdOk returns a tuple with the AggregatePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortIdentifier) GetAggregatePortIdOk() (*int64, bool) {
	if o == nil || o.AggregatePortId == nil {
		return nil, false
	}
	return o.AggregatePortId, true
}

// HasAggregatePortId returns a boolean if a field has been set.
func (o *FabricPortIdentifier) HasAggregatePortId() bool {
	if o != nil && o.AggregatePortId != nil {
		return true
	}

	return false
}

// SetAggregatePortId gets a reference to the given int64 and assigns it to the AggregatePortId field.
func (o *FabricPortIdentifier) SetAggregatePortId(v int64) {
	o.AggregatePortId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *FabricPortIdentifier) GetPortId() int64 {
	if o == nil || o.PortId == nil {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortIdentifier) GetPortIdOk() (*int64, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *FabricPortIdentifier) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *FabricPortIdentifier) SetPortId(v int64) {
	o.PortId = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *FabricPortIdentifier) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortIdentifier) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *FabricPortIdentifier) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *FabricPortIdentifier) SetSlotId(v int64) {
	o.SlotId = &v
}

func (o FabricPortIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.AggregatePortId != nil {
		toSerialize["AggregatePortId"] = o.AggregatePortId
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	return json.Marshal(toSerialize)
}

type NullableFabricPortIdentifier struct {
	value *FabricPortIdentifier
	isSet bool
}

func (v NullableFabricPortIdentifier) Get() *FabricPortIdentifier {
	return v.value
}

func (v *NullableFabricPortIdentifier) Set(val *FabricPortIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPortIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPortIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPortIdentifier(val *FabricPortIdentifier) *NullableFabricPortIdentifier {
	return &NullableFabricPortIdentifier{value: val, isSet: true}
}

func (v NullableFabricPortIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPortIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
