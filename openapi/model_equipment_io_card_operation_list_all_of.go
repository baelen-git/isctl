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

// EquipmentIoCardOperationListAllOf struct for EquipmentIoCardOperationListAllOf
type EquipmentIoCardOperationListAllOf struct {
	// The total number of 'equipment.IoCardOperation' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'equipment.IoCardOperation' resources matching the request.
	Results              []EquipmentIoCardOperation `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIoCardOperationListAllOf EquipmentIoCardOperationListAllOf

// NewEquipmentIoCardOperationListAllOf instantiates a new EquipmentIoCardOperationListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCardOperationListAllOf() *EquipmentIoCardOperationListAllOf {
	this := EquipmentIoCardOperationListAllOf{}
	return &this
}

// NewEquipmentIoCardOperationListAllOfWithDefaults instantiates a new EquipmentIoCardOperationListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardOperationListAllOfWithDefaults() *EquipmentIoCardOperationListAllOf {
	this := EquipmentIoCardOperationListAllOf{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *EquipmentIoCardOperationListAllOf) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationListAllOf) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *EquipmentIoCardOperationListAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *EquipmentIoCardOperationListAllOf) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardOperationListAllOf) GetResults() []EquipmentIoCardOperation {
	if o == nil {
		var ret []EquipmentIoCardOperation
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardOperationListAllOf) GetResultsOk() (*[]EquipmentIoCardOperation, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EquipmentIoCardOperationListAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EquipmentIoCardOperation and assigns it to the Results field.
func (o *EquipmentIoCardOperationListAllOf) SetResults(v []EquipmentIoCardOperation) {
	o.Results = v
}

func (o EquipmentIoCardOperationListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentIoCardOperationListAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentIoCardOperationListAllOf := _EquipmentIoCardOperationListAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentIoCardOperationListAllOf); err == nil {
		*o = EquipmentIoCardOperationListAllOf(varEquipmentIoCardOperationListAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentIoCardOperationListAllOf struct {
	value *EquipmentIoCardOperationListAllOf
	isSet bool
}

func (v NullableEquipmentIoCardOperationListAllOf) Get() *EquipmentIoCardOperationListAllOf {
	return v.value
}

func (v *NullableEquipmentIoCardOperationListAllOf) Set(val *EquipmentIoCardOperationListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardOperationListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardOperationListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardOperationListAllOf(val *EquipmentIoCardOperationListAllOf) *NullableEquipmentIoCardOperationListAllOf {
	return &NullableEquipmentIoCardOperationListAllOf{value: val, isSet: true}
}

func (v NullableEquipmentIoCardOperationListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardOperationListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
