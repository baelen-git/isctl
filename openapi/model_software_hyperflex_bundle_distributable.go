/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-12-08T20:53:20Z.
 *
 * API version: 1.0.9-2908
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// SoftwareHyperflexBundleDistributable A HyperFlex image bundle distributed by Cisco for Private Appliance.
type SoftwareHyperflexBundleDistributable struct {
	FirmwareBaseDistributable `yaml:"FirmwareBaseDistributable,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                                 `json:"ObjectType" yaml:"ObjectType"`
	Catalog    *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty" yaml:"Catalog,omitempty"`
	// An array of relationships to softwareHyperflexDistributable resources.
	Images []SoftwareHyperflexDistributableRelationship `json:"Images,omitempty" yaml:"Images,omitempty"`
}

// NewSoftwareHyperflexBundleDistributable instantiates a new SoftwareHyperflexBundleDistributable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareHyperflexBundleDistributable(classId string, objectType string) *SoftwareHyperflexBundleDistributable {
	this := SoftwareHyperflexBundleDistributable{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwareHyperflexBundleDistributableWithDefaults instantiates a new SoftwareHyperflexBundleDistributable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareHyperflexBundleDistributableWithDefaults() *SoftwareHyperflexBundleDistributable {
	this := SoftwareHyperflexBundleDistributable{}
	var classId string = "software.HyperflexBundleDistributable"
	this.ClassId = classId
	var objectType string = "software.HyperflexBundleDistributable"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareHyperflexBundleDistributable) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareHyperflexBundleDistributable) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareHyperflexBundleDistributable) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareHyperflexBundleDistributable) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareHyperflexBundleDistributable) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareHyperflexBundleDistributable) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwareHyperflexBundleDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareHyperflexBundleDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareHyperflexBundleDistributable) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareHyperflexBundleDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwareHyperflexBundleDistributable) GetImages() []SoftwareHyperflexDistributableRelationship {
	if o == nil {
		var ret []SoftwareHyperflexDistributableRelationship
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwareHyperflexBundleDistributable) GetImagesOk() (*[]SoftwareHyperflexDistributableRelationship, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return &o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *SoftwareHyperflexBundleDistributable) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []SoftwareHyperflexDistributableRelationship and assigns it to the Images field.
func (o *SoftwareHyperflexBundleDistributable) SetImages(v []SoftwareHyperflexDistributableRelationship) {
	o.Images = v
}

func (o SoftwareHyperflexBundleDistributable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseDistributable, errFirmwareBaseDistributable := json.Marshal(o.FirmwareBaseDistributable)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	errFirmwareBaseDistributable = json.Unmarshal([]byte(serializedFirmwareBaseDistributable), &toSerialize)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.Images != nil {
		toSerialize["Images"] = o.Images
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwareHyperflexBundleDistributable struct {
	value *SoftwareHyperflexBundleDistributable
	isSet bool
}

func (v NullableSoftwareHyperflexBundleDistributable) Get() *SoftwareHyperflexBundleDistributable {
	return v.value
}

func (v *NullableSoftwareHyperflexBundleDistributable) Set(val *SoftwareHyperflexBundleDistributable) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareHyperflexBundleDistributable) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareHyperflexBundleDistributable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareHyperflexBundleDistributable(val *SoftwareHyperflexBundleDistributable) *NullableSoftwareHyperflexBundleDistributable {
	return &NullableSoftwareHyperflexBundleDistributable{value: val, isSet: true}
}

func (v NullableSoftwareHyperflexBundleDistributable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareHyperflexBundleDistributable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
