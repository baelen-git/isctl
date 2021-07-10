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

// NiaapiNibFileDownloaderAllOf Definition of the list of properties defined in 'niaapi.NibFileDownloader', excluding properties defined in parent classes.
type NiaapiNibFileDownloaderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Filename of this Metadata package file, folder will be handled by api.
	FileName *string `json:"FileName,omitempty"`
	// The presigned URL from server to download this file.
	PresignedUrl         *string `json:"PresignedUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiNibFileDownloaderAllOf NiaapiNibFileDownloaderAllOf

// NewNiaapiNibFileDownloaderAllOf instantiates a new NiaapiNibFileDownloaderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiNibFileDownloaderAllOf(classId string, objectType string) *NiaapiNibFileDownloaderAllOf {
	this := NiaapiNibFileDownloaderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiNibFileDownloaderAllOfWithDefaults instantiates a new NiaapiNibFileDownloaderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiNibFileDownloaderAllOfWithDefaults() *NiaapiNibFileDownloaderAllOf {
	this := NiaapiNibFileDownloaderAllOf{}
	var classId string = "niaapi.NibFileDownloader"
	this.ClassId = classId
	var objectType string = "niaapi.NibFileDownloader"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiNibFileDownloaderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiNibFileDownloaderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiNibFileDownloaderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiNibFileDownloaderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiNibFileDownloaderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiNibFileDownloaderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *NiaapiNibFileDownloaderAllOf) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNibFileDownloaderAllOf) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *NiaapiNibFileDownloaderAllOf) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *NiaapiNibFileDownloaderAllOf) SetFileName(v string) {
	o.FileName = &v
}

// GetPresignedUrl returns the PresignedUrl field value if set, zero value otherwise.
func (o *NiaapiNibFileDownloaderAllOf) GetPresignedUrl() string {
	if o == nil || o.PresignedUrl == nil {
		var ret string
		return ret
	}
	return *o.PresignedUrl
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNibFileDownloaderAllOf) GetPresignedUrlOk() (*string, bool) {
	if o == nil || o.PresignedUrl == nil {
		return nil, false
	}
	return o.PresignedUrl, true
}

// HasPresignedUrl returns a boolean if a field has been set.
func (o *NiaapiNibFileDownloaderAllOf) HasPresignedUrl() bool {
	if o != nil && o.PresignedUrl != nil {
		return true
	}

	return false
}

// SetPresignedUrl gets a reference to the given string and assigns it to the PresignedUrl field.
func (o *NiaapiNibFileDownloaderAllOf) SetPresignedUrl(v string) {
	o.PresignedUrl = &v
}

func (o NiaapiNibFileDownloaderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.PresignedUrl != nil {
		toSerialize["PresignedUrl"] = o.PresignedUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiNibFileDownloaderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiNibFileDownloaderAllOf := _NiaapiNibFileDownloaderAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiNibFileDownloaderAllOf); err == nil {
		*o = NiaapiNibFileDownloaderAllOf(varNiaapiNibFileDownloaderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileName")
		delete(additionalProperties, "PresignedUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiNibFileDownloaderAllOf struct {
	value *NiaapiNibFileDownloaderAllOf
	isSet bool
}

func (v NullableNiaapiNibFileDownloaderAllOf) Get() *NiaapiNibFileDownloaderAllOf {
	return v.value
}

func (v *NullableNiaapiNibFileDownloaderAllOf) Set(val *NiaapiNibFileDownloaderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiNibFileDownloaderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiNibFileDownloaderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiNibFileDownloaderAllOf(val *NiaapiNibFileDownloaderAllOf) *NullableNiaapiNibFileDownloaderAllOf {
	return &NullableNiaapiNibFileDownloaderAllOf{value: val, isSet: true}
}

func (v NullableNiaapiNibFileDownloaderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiNibFileDownloaderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
