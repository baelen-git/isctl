/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-11T18:30:19Z.
 *
 * API version: 1.0.9-3252
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// MoDocumentCount A document returned in a response body when the HTTP GET request specifies the '$count' query parameter.
type MoDocumentCount struct {
	MoBaseResponse `yaml:"MoBaseResponse,inline"`
	// The total number of resources matching the query filter, accross all pages.
	Count *int32 `json:"Count,omitempty" yaml:"Count,omitempty"`
}

// NewMoDocumentCount instantiates a new MoDocumentCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoDocumentCount() *MoDocumentCount {
	this := MoDocumentCount{}
	return &this
}

// NewMoDocumentCountWithDefaults instantiates a new MoDocumentCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoDocumentCountWithDefaults() *MoDocumentCount {
	this := MoDocumentCount{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *MoDocumentCount) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoDocumentCount) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *MoDocumentCount) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *MoDocumentCount) SetCount(v int32) {
	o.Count = &v
}

func (o MoDocumentCount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseResponse, errMoBaseResponse := json.Marshal(o.MoBaseResponse)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	errMoBaseResponse = json.Unmarshal([]byte(serializedMoBaseResponse), &toSerialize)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableMoDocumentCount struct {
	value *MoDocumentCount
	isSet bool
}

func (v NullableMoDocumentCount) Get() *MoDocumentCount {
	return v.value
}

func (v *NullableMoDocumentCount) Set(val *MoDocumentCount) {
	v.value = val
	v.isSet = true
}

func (v NullableMoDocumentCount) IsSet() bool {
	return v.isSet
}

func (v *NullableMoDocumentCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoDocumentCount(val *MoDocumentCount) *NullableMoDocumentCount {
	return &NullableMoDocumentCount{value: val, isSet: true}
}

func (v NullableMoDocumentCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoDocumentCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
