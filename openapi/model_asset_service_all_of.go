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

// AssetServiceAllOf Definition of the list of properties defined in 'asset.Service', excluding properties defined in parent classes.
type AssetServiceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string                      `json:"ObjectType" yaml:"ObjectType"`
	Options    NullableAssetServiceOptions `json:"Options,omitempty" yaml:"Options,omitempty"`
}

// NewAssetServiceAllOf instantiates a new AssetServiceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetServiceAllOf(classId string, objectType string) *AssetServiceAllOf {
	this := AssetServiceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetServiceAllOfWithDefaults instantiates a new AssetServiceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetServiceAllOfWithDefaults() *AssetServiceAllOf {
	this := AssetServiceAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetServiceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetServiceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetServiceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetServiceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetServiceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetServiceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetServiceAllOf) GetOptions() AssetServiceOptions {
	if o == nil || o.Options.Get() == nil {
		var ret AssetServiceOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetServiceAllOf) GetOptionsOk() (*AssetServiceOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *AssetServiceAllOf) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableAssetServiceOptions and assigns it to the Options field.
func (o *AssetServiceAllOf) SetOptions(v AssetServiceOptions) {
	o.Options.Set(&v)
}

// SetOptionsNil sets the value for Options to be an explicit nil
func (o *AssetServiceAllOf) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *AssetServiceAllOf) UnsetOptions() {
	o.Options.Unset()
}

func (o AssetServiceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Options.IsSet() {
		toSerialize["Options"] = o.Options.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAssetServiceAllOf struct {
	value *AssetServiceAllOf
	isSet bool
}

func (v NullableAssetServiceAllOf) Get() *AssetServiceAllOf {
	return v.value
}

func (v *NullableAssetServiceAllOf) Set(val *AssetServiceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetServiceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetServiceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetServiceAllOf(val *AssetServiceAllOf) *NullableAssetServiceAllOf {
	return &NullableAssetServiceAllOf{value: val, isSet: true}
}

func (v NullableAssetServiceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetServiceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
