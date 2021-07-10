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

// ConfigMoRefAllOf Definition of the list of properties defined in 'config.MoRef', excluding properties defined in parent classes.
type ConfigMoRefAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Moid represents the MoId of the object.
	Moid                 *string `json:"Moid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigMoRefAllOf ConfigMoRefAllOf

// NewConfigMoRefAllOf instantiates a new ConfigMoRefAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigMoRefAllOf(classId string, objectType string) *ConfigMoRefAllOf {
	this := ConfigMoRefAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConfigMoRefAllOfWithDefaults instantiates a new ConfigMoRefAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigMoRefAllOfWithDefaults() *ConfigMoRefAllOf {
	this := ConfigMoRefAllOf{}
	var classId string = "config.MoRef"
	this.ClassId = classId
	var objectType string = "config.MoRef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConfigMoRefAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConfigMoRefAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConfigMoRefAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConfigMoRefAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConfigMoRefAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConfigMoRefAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMoid returns the Moid field value if set, zero value otherwise.
func (o *ConfigMoRefAllOf) GetMoid() string {
	if o == nil || o.Moid == nil {
		var ret string
		return ret
	}
	return *o.Moid
}

// GetMoidOk returns a tuple with the Moid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMoRefAllOf) GetMoidOk() (*string, bool) {
	if o == nil || o.Moid == nil {
		return nil, false
	}
	return o.Moid, true
}

// HasMoid returns a boolean if a field has been set.
func (o *ConfigMoRefAllOf) HasMoid() bool {
	if o != nil && o.Moid != nil {
		return true
	}

	return false
}

// SetMoid gets a reference to the given string and assigns it to the Moid field.
func (o *ConfigMoRefAllOf) SetMoid(v string) {
	o.Moid = &v
}

func (o ConfigMoRefAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Moid != nil {
		toSerialize["Moid"] = o.Moid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConfigMoRefAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConfigMoRefAllOf := _ConfigMoRefAllOf{}

	if err = json.Unmarshal(bytes, &varConfigMoRefAllOf); err == nil {
		*o = ConfigMoRefAllOf(varConfigMoRefAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Moid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigMoRefAllOf struct {
	value *ConfigMoRefAllOf
	isSet bool
}

func (v NullableConfigMoRefAllOf) Get() *ConfigMoRefAllOf {
	return v.value
}

func (v *NullableConfigMoRefAllOf) Set(val *ConfigMoRefAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigMoRefAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigMoRefAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigMoRefAllOf(val *ConfigMoRefAllOf) *NullableConfigMoRefAllOf {
	return &NullableConfigMoRefAllOf{value: val, isSet: true}
}

func (v NullableConfigMoRefAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigMoRefAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
