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

// IamAccountExperienceAllOf Definition of the list of properties defined in 'iam.AccountExperience', excluding properties defined in parent classes.
type IamAccountExperienceAllOf struct {
	Features *[]IamFeatureDefinition `json:"Features,omitempty" yaml:"Features,omitempty"`
	Account  *IamAccountRelationship `json:"Account,omitempty" yaml:"Account,omitempty"`
}

// NewIamAccountExperienceAllOf instantiates a new IamAccountExperienceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamAccountExperienceAllOf() *IamAccountExperienceAllOf {
	this := IamAccountExperienceAllOf{}
	return &this
}

// NewIamAccountExperienceAllOfWithDefaults instantiates a new IamAccountExperienceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamAccountExperienceAllOfWithDefaults() *IamAccountExperienceAllOf {
	this := IamAccountExperienceAllOf{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *IamAccountExperienceAllOf) GetFeatures() []IamFeatureDefinition {
	if o == nil || o.Features == nil {
		var ret []IamFeatureDefinition
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccountExperienceAllOf) GetFeaturesOk() (*[]IamFeatureDefinition, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *IamAccountExperienceAllOf) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []IamFeatureDefinition and assigns it to the Features field.
func (o *IamAccountExperienceAllOf) SetFeatures(v []IamFeatureDefinition) {
	o.Features = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamAccountExperienceAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAccountExperienceAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamAccountExperienceAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamAccountExperienceAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o IamAccountExperienceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Features != nil {
		toSerialize["Features"] = o.Features
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableIamAccountExperienceAllOf struct {
	value *IamAccountExperienceAllOf
	isSet bool
}

func (v NullableIamAccountExperienceAllOf) Get() *IamAccountExperienceAllOf {
	return v.value
}

func (v *NullableIamAccountExperienceAllOf) Set(val *IamAccountExperienceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamAccountExperienceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamAccountExperienceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamAccountExperienceAllOf(val *IamAccountExperienceAllOf) *NullableIamAccountExperienceAllOf {
	return &NullableIamAccountExperienceAllOf{value: val, isSet: true}
}

func (v NullableIamAccountExperienceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamAccountExperienceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
