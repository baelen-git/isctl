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

// ImcconnectorWebUiMessageAllOf Definition of the list of properties defined in 'imcconnector.WebUiMessage', excluding properties defined in parent classes.
type ImcconnectorWebUiMessageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The body content of the UI HTTP request to send to the BMC platform.
	WebUiRequest *string `json:"WebUiRequest,omitempty" yaml:"WebUiRequest,omitempty"`
}

// NewImcconnectorWebUiMessageAllOf instantiates a new ImcconnectorWebUiMessageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImcconnectorWebUiMessageAllOf(classId string, objectType string) *ImcconnectorWebUiMessageAllOf {
	this := ImcconnectorWebUiMessageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewImcconnectorWebUiMessageAllOfWithDefaults instantiates a new ImcconnectorWebUiMessageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImcconnectorWebUiMessageAllOfWithDefaults() *ImcconnectorWebUiMessageAllOf {
	this := ImcconnectorWebUiMessageAllOf{}
	var classId string = "imcconnector.WebUiMessage"
	this.ClassId = classId
	var objectType string = "imcconnector.WebUiMessage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ImcconnectorWebUiMessageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ImcconnectorWebUiMessageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ImcconnectorWebUiMessageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ImcconnectorWebUiMessageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ImcconnectorWebUiMessageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ImcconnectorWebUiMessageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetWebUiRequest returns the WebUiRequest field value if set, zero value otherwise.
func (o *ImcconnectorWebUiMessageAllOf) GetWebUiRequest() string {
	if o == nil || o.WebUiRequest == nil {
		var ret string
		return ret
	}
	return *o.WebUiRequest
}

// GetWebUiRequestOk returns a tuple with the WebUiRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImcconnectorWebUiMessageAllOf) GetWebUiRequestOk() (*string, bool) {
	if o == nil || o.WebUiRequest == nil {
		return nil, false
	}
	return o.WebUiRequest, true
}

// HasWebUiRequest returns a boolean if a field has been set.
func (o *ImcconnectorWebUiMessageAllOf) HasWebUiRequest() bool {
	if o != nil && o.WebUiRequest != nil {
		return true
	}

	return false
}

// SetWebUiRequest gets a reference to the given string and assigns it to the WebUiRequest field.
func (o *ImcconnectorWebUiMessageAllOf) SetWebUiRequest(v string) {
	o.WebUiRequest = &v
}

func (o ImcconnectorWebUiMessageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.WebUiRequest != nil {
		toSerialize["WebUiRequest"] = o.WebUiRequest
	}
	return json.Marshal(toSerialize)
}

type NullableImcconnectorWebUiMessageAllOf struct {
	value *ImcconnectorWebUiMessageAllOf
	isSet bool
}

func (v NullableImcconnectorWebUiMessageAllOf) Get() *ImcconnectorWebUiMessageAllOf {
	return v.value
}

func (v *NullableImcconnectorWebUiMessageAllOf) Set(val *ImcconnectorWebUiMessageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableImcconnectorWebUiMessageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableImcconnectorWebUiMessageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImcconnectorWebUiMessageAllOf(val *ImcconnectorWebUiMessageAllOf) *NullableImcconnectorWebUiMessageAllOf {
	return &NullableImcconnectorWebUiMessageAllOf{value: val, isSet: true}
}

func (v NullableImcconnectorWebUiMessageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImcconnectorWebUiMessageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
