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
	"reflect"
	"strings"
)

// HyperflexSoftwareDistributionVersion A HyperFlex Software Distribution version.
type HyperflexSoftwareDistributionVersion struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The HyperFlex Software Distribution version.
	Version *string `json:"Version,omitempty"`
	// An array of relationships to hyperflexSoftwareDistributionComponent resources.
	DistributionComponents    []HyperflexSoftwareDistributionComponentRelationship `json:"DistributionComponents,omitempty"`
	SoftwareDistributionEntry *HyperflexSoftwareDistributionEntryRelationship      `json:"SoftwareDistributionEntry,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _HyperflexSoftwareDistributionVersion HyperflexSoftwareDistributionVersion

// NewHyperflexSoftwareDistributionVersion instantiates a new HyperflexSoftwareDistributionVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSoftwareDistributionVersion(classId string, objectType string) *HyperflexSoftwareDistributionVersion {
	this := HyperflexSoftwareDistributionVersion{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSoftwareDistributionVersionWithDefaults instantiates a new HyperflexSoftwareDistributionVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSoftwareDistributionVersionWithDefaults() *HyperflexSoftwareDistributionVersion {
	this := HyperflexSoftwareDistributionVersion{}
	var classId string = "hyperflex.SoftwareDistributionVersion"
	this.ClassId = classId
	var objectType string = "hyperflex.SoftwareDistributionVersion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSoftwareDistributionVersion) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionVersion) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSoftwareDistributionVersion) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSoftwareDistributionVersion) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionVersion) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSoftwareDistributionVersion) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionVersion) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionVersion) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionVersion) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexSoftwareDistributionVersion) SetVersion(v string) {
	o.Version = &v
}

// GetDistributionComponents returns the DistributionComponents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSoftwareDistributionVersion) GetDistributionComponents() []HyperflexSoftwareDistributionComponentRelationship {
	if o == nil {
		var ret []HyperflexSoftwareDistributionComponentRelationship
		return ret
	}
	return o.DistributionComponents
}

// GetDistributionComponentsOk returns a tuple with the DistributionComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSoftwareDistributionVersion) GetDistributionComponentsOk() (*[]HyperflexSoftwareDistributionComponentRelationship, bool) {
	if o == nil || o.DistributionComponents == nil {
		return nil, false
	}
	return &o.DistributionComponents, true
}

// HasDistributionComponents returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionVersion) HasDistributionComponents() bool {
	if o != nil && o.DistributionComponents != nil {
		return true
	}

	return false
}

// SetDistributionComponents gets a reference to the given []HyperflexSoftwareDistributionComponentRelationship and assigns it to the DistributionComponents field.
func (o *HyperflexSoftwareDistributionVersion) SetDistributionComponents(v []HyperflexSoftwareDistributionComponentRelationship) {
	o.DistributionComponents = v
}

// GetSoftwareDistributionEntry returns the SoftwareDistributionEntry field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionVersion) GetSoftwareDistributionEntry() HyperflexSoftwareDistributionEntryRelationship {
	if o == nil || o.SoftwareDistributionEntry == nil {
		var ret HyperflexSoftwareDistributionEntryRelationship
		return ret
	}
	return *o.SoftwareDistributionEntry
}

// GetSoftwareDistributionEntryOk returns a tuple with the SoftwareDistributionEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionVersion) GetSoftwareDistributionEntryOk() (*HyperflexSoftwareDistributionEntryRelationship, bool) {
	if o == nil || o.SoftwareDistributionEntry == nil {
		return nil, false
	}
	return o.SoftwareDistributionEntry, true
}

// HasSoftwareDistributionEntry returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionVersion) HasSoftwareDistributionEntry() bool {
	if o != nil && o.SoftwareDistributionEntry != nil {
		return true
	}

	return false
}

// SetSoftwareDistributionEntry gets a reference to the given HyperflexSoftwareDistributionEntryRelationship and assigns it to the SoftwareDistributionEntry field.
func (o *HyperflexSoftwareDistributionVersion) SetSoftwareDistributionEntry(v HyperflexSoftwareDistributionEntryRelationship) {
	o.SoftwareDistributionEntry = &v
}

func (o HyperflexSoftwareDistributionVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.DistributionComponents != nil {
		toSerialize["DistributionComponents"] = o.DistributionComponents
	}
	if o.SoftwareDistributionEntry != nil {
		toSerialize["SoftwareDistributionEntry"] = o.SoftwareDistributionEntry
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSoftwareDistributionVersion) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexSoftwareDistributionVersionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The HyperFlex Software Distribution version.
		Version *string `json:"Version,omitempty"`
		// An array of relationships to hyperflexSoftwareDistributionComponent resources.
		DistributionComponents    []HyperflexSoftwareDistributionComponentRelationship `json:"DistributionComponents,omitempty"`
		SoftwareDistributionEntry *HyperflexSoftwareDistributionEntryRelationship      `json:"SoftwareDistributionEntry,omitempty"`
	}

	varHyperflexSoftwareDistributionVersionWithoutEmbeddedStruct := HyperflexSoftwareDistributionVersionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexSoftwareDistributionVersionWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexSoftwareDistributionVersion := _HyperflexSoftwareDistributionVersion{}
		varHyperflexSoftwareDistributionVersion.ClassId = varHyperflexSoftwareDistributionVersionWithoutEmbeddedStruct.ClassId
		varHyperflexSoftwareDistributionVersion.ObjectType = varHyperflexSoftwareDistributionVersionWithoutEmbeddedStruct.ObjectType
		varHyperflexSoftwareDistributionVersion.Version = varHyperflexSoftwareDistributionVersionWithoutEmbeddedStruct.Version
		varHyperflexSoftwareDistributionVersion.DistributionComponents = varHyperflexSoftwareDistributionVersionWithoutEmbeddedStruct.DistributionComponents
		varHyperflexSoftwareDistributionVersion.SoftwareDistributionEntry = varHyperflexSoftwareDistributionVersionWithoutEmbeddedStruct.SoftwareDistributionEntry
		*o = HyperflexSoftwareDistributionVersion(varHyperflexSoftwareDistributionVersion)
	} else {
		return err
	}

	varHyperflexSoftwareDistributionVersion := _HyperflexSoftwareDistributionVersion{}

	err = json.Unmarshal(bytes, &varHyperflexSoftwareDistributionVersion)
	if err == nil {
		o.MoBaseMo = varHyperflexSoftwareDistributionVersion.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "DistributionComponents")
		delete(additionalProperties, "SoftwareDistributionEntry")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexSoftwareDistributionVersion struct {
	value *HyperflexSoftwareDistributionVersion
	isSet bool
}

func (v NullableHyperflexSoftwareDistributionVersion) Get() *HyperflexSoftwareDistributionVersion {
	return v.value
}

func (v *NullableHyperflexSoftwareDistributionVersion) Set(val *HyperflexSoftwareDistributionVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSoftwareDistributionVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSoftwareDistributionVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSoftwareDistributionVersion(val *HyperflexSoftwareDistributionVersion) *NullableHyperflexSoftwareDistributionVersion {
	return &NullableHyperflexSoftwareDistributionVersion{value: val, isSet: true}
}

func (v NullableHyperflexSoftwareDistributionVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSoftwareDistributionVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
