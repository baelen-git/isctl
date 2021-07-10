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

// TamPsirtSeverityAllOf Definition of the list of properties defined in 'tam.PsirtSeverity', excluding properties defined in parent classes.
type TamPsirtSeverityAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Severity level associated with the security advisory. * `critical` - < If applicable, it may expose users to critical failures and it needs to be addressed immediately. For e.g. a PSIRT advisory with a corresponding CVSS score of above 9.0. * `high` - < If applicable, it may expose the users to critical failure and it needs to be addressed immediately. For e.g. a PSIRT advisory with a corresponding CVSS score between 7.0-8.9. These may be the vulnerabilities that are more difficult to exploit than the ones deemed critical but once exploited, the will cause critical failures. * `medium` - < If applicable, it may expose the users to failure of certain functions. for e.g. a PSIRT advisory with a corresponding CVSS score between 4.0-6.9. These may be the vulnerabilities that are mitigated to a large extent but that may still be exploited in certain restricted cases. * `info` - < If applicable, it may have some minimal impact for certain functions in the user environment. For e.g. a PSIRT advisory with a corresponding CVSS score below 4.0.
	Level                *string `json:"Level,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamPsirtSeverityAllOf TamPsirtSeverityAllOf

// NewTamPsirtSeverityAllOf instantiates a new TamPsirtSeverityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamPsirtSeverityAllOf(classId string, objectType string) *TamPsirtSeverityAllOf {
	this := TamPsirtSeverityAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var level string = "critical"
	this.Level = &level
	return &this
}

// NewTamPsirtSeverityAllOfWithDefaults instantiates a new TamPsirtSeverityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamPsirtSeverityAllOfWithDefaults() *TamPsirtSeverityAllOf {
	this := TamPsirtSeverityAllOf{}
	var classId string = "tam.PsirtSeverity"
	this.ClassId = classId
	var objectType string = "tam.PsirtSeverity"
	this.ObjectType = objectType
	var level string = "critical"
	this.Level = &level
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamPsirtSeverityAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamPsirtSeverityAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamPsirtSeverityAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamPsirtSeverityAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamPsirtSeverityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamPsirtSeverityAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *TamPsirtSeverityAllOf) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamPsirtSeverityAllOf) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *TamPsirtSeverityAllOf) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *TamPsirtSeverityAllOf) SetLevel(v string) {
	o.Level = &v
}

func (o TamPsirtSeverityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Level != nil {
		toSerialize["Level"] = o.Level
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamPsirtSeverityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamPsirtSeverityAllOf := _TamPsirtSeverityAllOf{}

	if err = json.Unmarshal(bytes, &varTamPsirtSeverityAllOf); err == nil {
		*o = TamPsirtSeverityAllOf(varTamPsirtSeverityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Level")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamPsirtSeverityAllOf struct {
	value *TamPsirtSeverityAllOf
	isSet bool
}

func (v NullableTamPsirtSeverityAllOf) Get() *TamPsirtSeverityAllOf {
	return v.value
}

func (v *NullableTamPsirtSeverityAllOf) Set(val *TamPsirtSeverityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamPsirtSeverityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamPsirtSeverityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamPsirtSeverityAllOf(val *TamPsirtSeverityAllOf) *NullableTamPsirtSeverityAllOf {
	return &NullableTamPsirtSeverityAllOf{value: val, isSet: true}
}

func (v NullableTamPsirtSeverityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamPsirtSeverityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
