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

// SoftwarerepositoryImportResultAllOf Definition of the list of properties defined in 'softwarerepository.ImportResult', excluding properties defined in parent classes.
type SoftwarerepositoryImportResultAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The reason for the failure of an import operation, if applicable.
	ErrorMessage *string `json:"ErrorMessage,omitempty"`
	// The progress percentage for the import operation.
	Progress             *int64 `json:"Progress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryImportResultAllOf SoftwarerepositoryImportResultAllOf

// NewSoftwarerepositoryImportResultAllOf instantiates a new SoftwarerepositoryImportResultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryImportResultAllOf(classId string, objectType string) *SoftwarerepositoryImportResultAllOf {
	this := SoftwarerepositoryImportResultAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryImportResultAllOfWithDefaults instantiates a new SoftwarerepositoryImportResultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryImportResultAllOfWithDefaults() *SoftwarerepositoryImportResultAllOf {
	this := SoftwarerepositoryImportResultAllOf{}
	var classId string = "softwarerepository.ImportResult"
	this.ClassId = classId
	var objectType string = "softwarerepository.ImportResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryImportResultAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryImportResultAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryImportResultAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryImportResultAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryImportResultAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryImportResultAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SoftwarerepositoryImportResultAllOf) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryImportResultAllOf) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SoftwarerepositoryImportResultAllOf) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SoftwarerepositoryImportResultAllOf) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *SoftwarerepositoryImportResultAllOf) GetProgress() int64 {
	if o == nil || o.Progress == nil {
		var ret int64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryImportResultAllOf) GetProgressOk() (*int64, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *SoftwarerepositoryImportResultAllOf) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int64 and assigns it to the Progress field.
func (o *SoftwarerepositoryImportResultAllOf) SetProgress(v int64) {
	o.Progress = &v
}

func (o SoftwarerepositoryImportResultAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ErrorMessage != nil {
		toSerialize["ErrorMessage"] = o.ErrorMessage
	}
	if o.Progress != nil {
		toSerialize["Progress"] = o.Progress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryImportResultAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryImportResultAllOf := _SoftwarerepositoryImportResultAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryImportResultAllOf); err == nil {
		*o = SoftwarerepositoryImportResultAllOf(varSoftwarerepositoryImportResultAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorMessage")
		delete(additionalProperties, "Progress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryImportResultAllOf struct {
	value *SoftwarerepositoryImportResultAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryImportResultAllOf) Get() *SoftwarerepositoryImportResultAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryImportResultAllOf) Set(val *SoftwarerepositoryImportResultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryImportResultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryImportResultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryImportResultAllOf(val *SoftwarerepositoryImportResultAllOf) *NullableSoftwarerepositoryImportResultAllOf {
	return &NullableSoftwarerepositoryImportResultAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryImportResultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryImportResultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
