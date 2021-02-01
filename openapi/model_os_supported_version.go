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

// OsSupportedVersion The supported operating system version by SCU. The API is mainly for UI operation. In the software management page, operating system configuration will be created by providing the vendor and version. The version will be filtered based on vendor.
type OsSupportedVersion struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The OsInstall Supported Operating System Version Name.
	VersionName *string                               `json:"VersionName,omitempty" yaml:"VersionName,omitempty"`
	Vendor      *HclOperatingSystemVendorRelationship `json:"Vendor,omitempty" yaml:"Vendor,omitempty"`
	Version     *HclOperatingSystemRelationship       `json:"Version,omitempty" yaml:"Version,omitempty"`
}

// NewOsSupportedVersion instantiates a new OsSupportedVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsSupportedVersion(classId string, objectType string) *OsSupportedVersion {
	this := OsSupportedVersion{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsSupportedVersionWithDefaults instantiates a new OsSupportedVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsSupportedVersionWithDefaults() *OsSupportedVersion {
	this := OsSupportedVersion{}
	var classId string = "os.SupportedVersion"
	this.ClassId = classId
	var objectType string = "os.SupportedVersion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsSupportedVersion) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsSupportedVersion) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsSupportedVersion) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsSupportedVersion) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsSupportedVersion) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsSupportedVersion) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *OsSupportedVersion) GetVersionName() string {
	if o == nil || o.VersionName == nil {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsSupportedVersion) GetVersionNameOk() (*string, bool) {
	if o == nil || o.VersionName == nil {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *OsSupportedVersion) HasVersionName() bool {
	if o != nil && o.VersionName != nil {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *OsSupportedVersion) SetVersionName(v string) {
	o.VersionName = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *OsSupportedVersion) GetVendor() HclOperatingSystemVendorRelationship {
	if o == nil || o.Vendor == nil {
		var ret HclOperatingSystemVendorRelationship
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsSupportedVersion) GetVendorOk() (*HclOperatingSystemVendorRelationship, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *OsSupportedVersion) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given HclOperatingSystemVendorRelationship and assigns it to the Vendor field.
func (o *OsSupportedVersion) SetVendor(v HclOperatingSystemVendorRelationship) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OsSupportedVersion) GetVersion() HclOperatingSystemRelationship {
	if o == nil || o.Version == nil {
		var ret HclOperatingSystemRelationship
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsSupportedVersion) GetVersionOk() (*HclOperatingSystemRelationship, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OsSupportedVersion) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given HclOperatingSystemRelationship and assigns it to the Version field.
func (o *OsSupportedVersion) SetVersion(v HclOperatingSystemRelationship) {
	o.Version = &v
}

func (o OsSupportedVersion) MarshalJSON() ([]byte, error) {
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
	if o.VersionName != nil {
		toSerialize["VersionName"] = o.VersionName
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableOsSupportedVersion struct {
	value *OsSupportedVersion
	isSet bool
}

func (v NullableOsSupportedVersion) Get() *OsSupportedVersion {
	return v.value
}

func (v *NullableOsSupportedVersion) Set(val *OsSupportedVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableOsSupportedVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableOsSupportedVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsSupportedVersion(val *OsSupportedVersion) *NullableOsSupportedVersion {
	return &NullableOsSupportedVersion{value: val, isSet: true}
}

func (v NullableOsSupportedVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsSupportedVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
