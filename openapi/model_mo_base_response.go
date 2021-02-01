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

// MoBaseResponse The base response schema.
type MoBaseResponse struct {
	// A discriminator value to disambiguate the schema of a HTTP GET response body.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
}

// NewMoBaseResponse instantiates a new MoBaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoBaseResponse(objectType string) *MoBaseResponse {
	this := MoBaseResponse{}
	this.ObjectType = objectType
	return &this
}

// NewMoBaseResponseWithDefaults instantiates a new MoBaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoBaseResponseWithDefaults() *MoBaseResponse {
	this := MoBaseResponse{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *MoBaseResponse) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MoBaseResponse) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MoBaseResponse) SetObjectType(v string) {
	o.ObjectType = v
}

func (o MoBaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableMoBaseResponse struct {
	value *MoBaseResponse
	isSet bool
}

func (v NullableMoBaseResponse) Get() *MoBaseResponse {
	return v.value
}

func (v *NullableMoBaseResponse) Set(val *MoBaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMoBaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMoBaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoBaseResponse(val *MoBaseResponse) *NullableMoBaseResponse {
	return &NullableMoBaseResponse{value: val, isSet: true}
}

func (v NullableMoBaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoBaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
