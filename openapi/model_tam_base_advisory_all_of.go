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

// TamBaseAdvisoryAllOf Definition of the list of properties defined in 'tam.BaseAdvisory', excluding properties defined in parent classes.
type TamBaseAdvisoryAllOf struct {
	// Brief description of the advisory details.
	Description *string `json:"Description,omitempty" yaml:"Description,omitempty"`
	// A user defined name for the Intersight Advisory.
	Name     *string      `json:"Name,omitempty" yaml:"Name,omitempty"`
	Severity *TamSeverity `json:"Severity,omitempty" yaml:"Severity,omitempty"`
	// Current state of the advisory. * `ready` - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created. * `evaluating` - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation.
	State *string `json:"State,omitempty" yaml:"State,omitempty"`
}

// NewTamBaseAdvisoryAllOf instantiates a new TamBaseAdvisoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamBaseAdvisoryAllOf() *TamBaseAdvisoryAllOf {
	this := TamBaseAdvisoryAllOf{}
	var state string = "ready"
	this.State = &state
	return &this
}

// NewTamBaseAdvisoryAllOfWithDefaults instantiates a new TamBaseAdvisoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamBaseAdvisoryAllOfWithDefaults() *TamBaseAdvisoryAllOf {
	this := TamBaseAdvisoryAllOf{}
	var state string = "ready"
	this.State = &state
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TamBaseAdvisoryAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisoryAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TamBaseAdvisoryAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TamBaseAdvisoryAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TamBaseAdvisoryAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisoryAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TamBaseAdvisoryAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TamBaseAdvisoryAllOf) SetName(v string) {
	o.Name = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *TamBaseAdvisoryAllOf) GetSeverity() TamSeverity {
	if o == nil || o.Severity == nil {
		var ret TamSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisoryAllOf) GetSeverityOk() (*TamSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *TamBaseAdvisoryAllOf) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given TamSeverity and assigns it to the Severity field.
func (o *TamBaseAdvisoryAllOf) SetSeverity(v TamSeverity) {
	o.Severity = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TamBaseAdvisoryAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamBaseAdvisoryAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TamBaseAdvisoryAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TamBaseAdvisoryAllOf) SetState(v string) {
	o.State = &v
}

func (o TamBaseAdvisoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableTamBaseAdvisoryAllOf struct {
	value *TamBaseAdvisoryAllOf
	isSet bool
}

func (v NullableTamBaseAdvisoryAllOf) Get() *TamBaseAdvisoryAllOf {
	return v.value
}

func (v *NullableTamBaseAdvisoryAllOf) Set(val *TamBaseAdvisoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamBaseAdvisoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamBaseAdvisoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamBaseAdvisoryAllOf(val *TamBaseAdvisoryAllOf) *NullableTamBaseAdvisoryAllOf {
	return &NullableTamBaseAdvisoryAllOf{value: val, isSet: true}
}

func (v NullableTamBaseAdvisoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamBaseAdvisoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}