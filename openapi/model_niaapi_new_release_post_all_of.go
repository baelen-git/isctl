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
	"time"
)

// NiaapiNewReleasePostAllOf Definition of the list of properties defined in 'niaapi.NewReleasePost', excluding properties defined in parent classes.
type NiaapiNewReleasePostAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The date when this new release notice is posted.
	PostDate   *time.Time                     `json:"PostDate,omitempty" yaml:"PostDate,omitempty"`
	PostDetail NullableNiaapiNewReleaseDetail `json:"PostDetail,omitempty" yaml:"PostDetail,omitempty"`
	// The document type of this post.
	PostType *string `json:"PostType,omitempty" yaml:"PostType,omitempty"`
	// Identificator of this inbox post.
	Postid *string `json:"Postid,omitempty" yaml:"Postid,omitempty"`
	// Revision number of this notice.
	Revision *string `json:"Revision,omitempty" yaml:"Revision,omitempty"`
}

// NewNiaapiNewReleasePostAllOf instantiates a new NiaapiNewReleasePostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiNewReleasePostAllOf(classId string, objectType string) *NiaapiNewReleasePostAllOf {
	this := NiaapiNewReleasePostAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiNewReleasePostAllOfWithDefaults instantiates a new NiaapiNewReleasePostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiNewReleasePostAllOfWithDefaults() *NiaapiNewReleasePostAllOf {
	this := NiaapiNewReleasePostAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiNewReleasePostAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePostAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiNewReleasePostAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiNewReleasePostAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePostAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiNewReleasePostAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPostDate returns the PostDate field value if set, zero value otherwise.
func (o *NiaapiNewReleasePostAllOf) GetPostDate() time.Time {
	if o == nil || o.PostDate == nil {
		var ret time.Time
		return ret
	}
	return *o.PostDate
}

// GetPostDateOk returns a tuple with the PostDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePostAllOf) GetPostDateOk() (*time.Time, bool) {
	if o == nil || o.PostDate == nil {
		return nil, false
	}
	return o.PostDate, true
}

// HasPostDate returns a boolean if a field has been set.
func (o *NiaapiNewReleasePostAllOf) HasPostDate() bool {
	if o != nil && o.PostDate != nil {
		return true
	}

	return false
}

// SetPostDate gets a reference to the given time.Time and assigns it to the PostDate field.
func (o *NiaapiNewReleasePostAllOf) SetPostDate(v time.Time) {
	o.PostDate = &v
}

// GetPostDetail returns the PostDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiNewReleasePostAllOf) GetPostDetail() NiaapiNewReleaseDetail {
	if o == nil || o.PostDetail.Get() == nil {
		var ret NiaapiNewReleaseDetail
		return ret
	}
	return *o.PostDetail.Get()
}

// GetPostDetailOk returns a tuple with the PostDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiNewReleasePostAllOf) GetPostDetailOk() (*NiaapiNewReleaseDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostDetail.Get(), o.PostDetail.IsSet()
}

// HasPostDetail returns a boolean if a field has been set.
func (o *NiaapiNewReleasePostAllOf) HasPostDetail() bool {
	if o != nil && o.PostDetail.IsSet() {
		return true
	}

	return false
}

// SetPostDetail gets a reference to the given NullableNiaapiNewReleaseDetail and assigns it to the PostDetail field.
func (o *NiaapiNewReleasePostAllOf) SetPostDetail(v NiaapiNewReleaseDetail) {
	o.PostDetail.Set(&v)
}

// SetPostDetailNil sets the value for PostDetail to be an explicit nil
func (o *NiaapiNewReleasePostAllOf) SetPostDetailNil() {
	o.PostDetail.Set(nil)
}

// UnsetPostDetail ensures that no value is present for PostDetail, not even an explicit nil
func (o *NiaapiNewReleasePostAllOf) UnsetPostDetail() {
	o.PostDetail.Unset()
}

// GetPostType returns the PostType field value if set, zero value otherwise.
func (o *NiaapiNewReleasePostAllOf) GetPostType() string {
	if o == nil || o.PostType == nil {
		var ret string
		return ret
	}
	return *o.PostType
}

// GetPostTypeOk returns a tuple with the PostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePostAllOf) GetPostTypeOk() (*string, bool) {
	if o == nil || o.PostType == nil {
		return nil, false
	}
	return o.PostType, true
}

// HasPostType returns a boolean if a field has been set.
func (o *NiaapiNewReleasePostAllOf) HasPostType() bool {
	if o != nil && o.PostType != nil {
		return true
	}

	return false
}

// SetPostType gets a reference to the given string and assigns it to the PostType field.
func (o *NiaapiNewReleasePostAllOf) SetPostType(v string) {
	o.PostType = &v
}

// GetPostid returns the Postid field value if set, zero value otherwise.
func (o *NiaapiNewReleasePostAllOf) GetPostid() string {
	if o == nil || o.Postid == nil {
		var ret string
		return ret
	}
	return *o.Postid
}

// GetPostidOk returns a tuple with the Postid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePostAllOf) GetPostidOk() (*string, bool) {
	if o == nil || o.Postid == nil {
		return nil, false
	}
	return o.Postid, true
}

// HasPostid returns a boolean if a field has been set.
func (o *NiaapiNewReleasePostAllOf) HasPostid() bool {
	if o != nil && o.Postid != nil {
		return true
	}

	return false
}

// SetPostid gets a reference to the given string and assigns it to the Postid field.
func (o *NiaapiNewReleasePostAllOf) SetPostid(v string) {
	o.Postid = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *NiaapiNewReleasePostAllOf) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePostAllOf) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *NiaapiNewReleasePostAllOf) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *NiaapiNewReleasePostAllOf) SetRevision(v string) {
	o.Revision = &v
}

func (o NiaapiNewReleasePostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PostDate != nil {
		toSerialize["PostDate"] = o.PostDate
	}
	if o.PostDetail.IsSet() {
		toSerialize["PostDetail"] = o.PostDetail.Get()
	}
	if o.PostType != nil {
		toSerialize["PostType"] = o.PostType
	}
	if o.Postid != nil {
		toSerialize["Postid"] = o.Postid
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}
	return json.Marshal(toSerialize)
}

type NullableNiaapiNewReleasePostAllOf struct {
	value *NiaapiNewReleasePostAllOf
	isSet bool
}

func (v NullableNiaapiNewReleasePostAllOf) Get() *NiaapiNewReleasePostAllOf {
	return v.value
}

func (v *NullableNiaapiNewReleasePostAllOf) Set(val *NiaapiNewReleasePostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiNewReleasePostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiNewReleasePostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiNewReleasePostAllOf(val *NiaapiNewReleasePostAllOf) *NullableNiaapiNewReleasePostAllOf {
	return &NullableNiaapiNewReleasePostAllOf{value: val, isSet: true}
}

func (v NullableNiaapiNewReleasePostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiNewReleasePostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
