/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// OsTemplateFileAllOf Definition of the list of properties defined in 'os.TemplateFile', excluding properties defined in parent classes.
type OsTemplateFileAllOf struct {
	// The name of the OS Template File that user uploads for unattended installation.
	Name         *string   `json:"Name,omitempty" yaml:"Name,omitempty"`
	Placeholders *[]string `json:"Placeholders,omitempty" yaml:"Placeholders,omitempty"`
	// The content of the entire template file is stored as value. The content can either be a static file content or a template content. The template is expected to conform to the golang template syntax.  The placeholders, if any, would be populated and the values provided would be  used to populate this template.
	TemplateContent *string `json:"TemplateContent,omitempty" yaml:"TemplateContent,omitempty"`
}

// NewOsTemplateFileAllOf instantiates a new OsTemplateFileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsTemplateFileAllOf() *OsTemplateFileAllOf {
	this := OsTemplateFileAllOf{}
	return &this
}

// NewOsTemplateFileAllOfWithDefaults instantiates a new OsTemplateFileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsTemplateFileAllOfWithDefaults() *OsTemplateFileAllOf {
	this := OsTemplateFileAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsTemplateFileAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTemplateFileAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsTemplateFileAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsTemplateFileAllOf) SetName(v string) {
	o.Name = &v
}

// GetPlaceholders returns the Placeholders field value if set, zero value otherwise.
func (o *OsTemplateFileAllOf) GetPlaceholders() []string {
	if o == nil || o.Placeholders == nil {
		var ret []string
		return ret
	}
	return *o.Placeholders
}

// GetPlaceholdersOk returns a tuple with the Placeholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTemplateFileAllOf) GetPlaceholdersOk() (*[]string, bool) {
	if o == nil || o.Placeholders == nil {
		return nil, false
	}
	return o.Placeholders, true
}

// HasPlaceholders returns a boolean if a field has been set.
func (o *OsTemplateFileAllOf) HasPlaceholders() bool {
	if o != nil && o.Placeholders != nil {
		return true
	}

	return false
}

// SetPlaceholders gets a reference to the given []string and assigns it to the Placeholders field.
func (o *OsTemplateFileAllOf) SetPlaceholders(v []string) {
	o.Placeholders = &v
}

// GetTemplateContent returns the TemplateContent field value if set, zero value otherwise.
func (o *OsTemplateFileAllOf) GetTemplateContent() string {
	if o == nil || o.TemplateContent == nil {
		var ret string
		return ret
	}
	return *o.TemplateContent
}

// GetTemplateContentOk returns a tuple with the TemplateContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTemplateFileAllOf) GetTemplateContentOk() (*string, bool) {
	if o == nil || o.TemplateContent == nil {
		return nil, false
	}
	return o.TemplateContent, true
}

// HasTemplateContent returns a boolean if a field has been set.
func (o *OsTemplateFileAllOf) HasTemplateContent() bool {
	if o != nil && o.TemplateContent != nil {
		return true
	}

	return false
}

// SetTemplateContent gets a reference to the given string and assigns it to the TemplateContent field.
func (o *OsTemplateFileAllOf) SetTemplateContent(v string) {
	o.TemplateContent = &v
}

func (o OsTemplateFileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Placeholders != nil {
		toSerialize["Placeholders"] = o.Placeholders
	}
	if o.TemplateContent != nil {
		toSerialize["TemplateContent"] = o.TemplateContent
	}
	return json.Marshal(toSerialize)
}

type NullableOsTemplateFileAllOf struct {
	value *OsTemplateFileAllOf
	isSet bool
}

func (v NullableOsTemplateFileAllOf) Get() *OsTemplateFileAllOf {
	return v.value
}

func (v *NullableOsTemplateFileAllOf) Set(val *OsTemplateFileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsTemplateFileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsTemplateFileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsTemplateFileAllOf(val *OsTemplateFileAllOf) *NullableOsTemplateFileAllOf {
	return &NullableOsTemplateFileAllOf{value: val, isSet: true}
}

func (v NullableOsTemplateFileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsTemplateFileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
