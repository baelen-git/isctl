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

// IamUserPreferenceAllOf Definition of the list of properties defined in 'iam.UserPreference', excluding properties defined in parent classes.
type IamUserPreferenceAllOf struct {
	// UI preferences of the user.
	Preference interface{} `json:"Preference,omitempty" yaml:"Preference,omitempty"`
	// Unique id of the user used by the identity provider to store the user.
	UserUniqueIdentifier *string                      `json:"UserUniqueIdentifier,omitempty" yaml:"UserUniqueIdentifier,omitempty"`
	Idp                  *IamIdpRelationship          `json:"Idp,omitempty" yaml:"Idp,omitempty"`
	IdpReference         *IamIdpReferenceRelationship `json:"IdpReference,omitempty" yaml:"IdpReference,omitempty"`
}

// NewIamUserPreferenceAllOf instantiates a new IamUserPreferenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUserPreferenceAllOf() *IamUserPreferenceAllOf {
	this := IamUserPreferenceAllOf{}
	return &this
}

// NewIamUserPreferenceAllOfWithDefaults instantiates a new IamUserPreferenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserPreferenceAllOfWithDefaults() *IamUserPreferenceAllOf {
	this := IamUserPreferenceAllOf{}
	return &this
}

// GetPreference returns the Preference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserPreferenceAllOf) GetPreference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserPreferenceAllOf) GetPreferenceOk() (*interface{}, bool) {
	if o == nil || o.Preference == nil {
		return nil, false
	}
	return &o.Preference, true
}

// HasPreference returns a boolean if a field has been set.
func (o *IamUserPreferenceAllOf) HasPreference() bool {
	if o != nil && o.Preference != nil {
		return true
	}

	return false
}

// SetPreference gets a reference to the given interface{} and assigns it to the Preference field.
func (o *IamUserPreferenceAllOf) SetPreference(v interface{}) {
	o.Preference = v
}

// GetUserUniqueIdentifier returns the UserUniqueIdentifier field value if set, zero value otherwise.
func (o *IamUserPreferenceAllOf) GetUserUniqueIdentifier() string {
	if o == nil || o.UserUniqueIdentifier == nil {
		var ret string
		return ret
	}
	return *o.UserUniqueIdentifier
}

// GetUserUniqueIdentifierOk returns a tuple with the UserUniqueIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserPreferenceAllOf) GetUserUniqueIdentifierOk() (*string, bool) {
	if o == nil || o.UserUniqueIdentifier == nil {
		return nil, false
	}
	return o.UserUniqueIdentifier, true
}

// HasUserUniqueIdentifier returns a boolean if a field has been set.
func (o *IamUserPreferenceAllOf) HasUserUniqueIdentifier() bool {
	if o != nil && o.UserUniqueIdentifier != nil {
		return true
	}

	return false
}

// SetUserUniqueIdentifier gets a reference to the given string and assigns it to the UserUniqueIdentifier field.
func (o *IamUserPreferenceAllOf) SetUserUniqueIdentifier(v string) {
	o.UserUniqueIdentifier = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IamUserPreferenceAllOf) GetIdp() IamIdpRelationship {
	if o == nil || o.Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserPreferenceAllOf) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IamUserPreferenceAllOf) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IamIdpRelationship and assigns it to the Idp field.
func (o *IamUserPreferenceAllOf) SetIdp(v IamIdpRelationship) {
	o.Idp = &v
}

// GetIdpReference returns the IdpReference field value if set, zero value otherwise.
func (o *IamUserPreferenceAllOf) GetIdpReference() IamIdpReferenceRelationship {
	if o == nil || o.IdpReference == nil {
		var ret IamIdpReferenceRelationship
		return ret
	}
	return *o.IdpReference
}

// GetIdpReferenceOk returns a tuple with the IdpReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserPreferenceAllOf) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool) {
	if o == nil || o.IdpReference == nil {
		return nil, false
	}
	return o.IdpReference, true
}

// HasIdpReference returns a boolean if a field has been set.
func (o *IamUserPreferenceAllOf) HasIdpReference() bool {
	if o != nil && o.IdpReference != nil {
		return true
	}

	return false
}

// SetIdpReference gets a reference to the given IamIdpReferenceRelationship and assigns it to the IdpReference field.
func (o *IamUserPreferenceAllOf) SetIdpReference(v IamIdpReferenceRelationship) {
	o.IdpReference = &v
}

func (o IamUserPreferenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Preference != nil {
		toSerialize["Preference"] = o.Preference
	}
	if o.UserUniqueIdentifier != nil {
		toSerialize["UserUniqueIdentifier"] = o.UserUniqueIdentifier
	}
	if o.Idp != nil {
		toSerialize["Idp"] = o.Idp
	}
	if o.IdpReference != nil {
		toSerialize["IdpReference"] = o.IdpReference
	}
	return json.Marshal(toSerialize)
}

type NullableIamUserPreferenceAllOf struct {
	value *IamUserPreferenceAllOf
	isSet bool
}

func (v NullableIamUserPreferenceAllOf) Get() *IamUserPreferenceAllOf {
	return v.value
}

func (v *NullableIamUserPreferenceAllOf) Set(val *IamUserPreferenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamUserPreferenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamUserPreferenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamUserPreferenceAllOf(val *IamUserPreferenceAllOf) *NullableIamUserPreferenceAllOf {
	return &NullableIamUserPreferenceAllOf{value: val, isSet: true}
}

func (v NullableIamUserPreferenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamUserPreferenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
