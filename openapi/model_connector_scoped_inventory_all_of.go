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

// ConnectorScopedInventoryAllOf Definition of the list of properties defined in 'connector.ScopedInventory', excluding properties defined in parent classes.
type ConnectorScopedInventoryAllOf struct {
	// A property that uniquely identifies the object to be inventoried as a part of the scoped inventory.
	NamingProperty *string `json:"NamingProperty,omitempty" yaml:"NamingProperty,omitempty"`
	// Type of the object for which scoped inventory needs to be run.
	Type   *string   `json:"Type,omitempty" yaml:"Type,omitempty"`
	Values *[]string `json:"Values,omitempty" yaml:"Values,omitempty"`
}

// NewConnectorScopedInventoryAllOf instantiates a new ConnectorScopedInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorScopedInventoryAllOf() *ConnectorScopedInventoryAllOf {
	this := ConnectorScopedInventoryAllOf{}
	return &this
}

// NewConnectorScopedInventoryAllOfWithDefaults instantiates a new ConnectorScopedInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorScopedInventoryAllOfWithDefaults() *ConnectorScopedInventoryAllOf {
	this := ConnectorScopedInventoryAllOf{}
	return &this
}

// GetNamingProperty returns the NamingProperty field value if set, zero value otherwise.
func (o *ConnectorScopedInventoryAllOf) GetNamingProperty() string {
	if o == nil || o.NamingProperty == nil {
		var ret string
		return ret
	}
	return *o.NamingProperty
}

// GetNamingPropertyOk returns a tuple with the NamingProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorScopedInventoryAllOf) GetNamingPropertyOk() (*string, bool) {
	if o == nil || o.NamingProperty == nil {
		return nil, false
	}
	return o.NamingProperty, true
}

// HasNamingProperty returns a boolean if a field has been set.
func (o *ConnectorScopedInventoryAllOf) HasNamingProperty() bool {
	if o != nil && o.NamingProperty != nil {
		return true
	}

	return false
}

// SetNamingProperty gets a reference to the given string and assigns it to the NamingProperty field.
func (o *ConnectorScopedInventoryAllOf) SetNamingProperty(v string) {
	o.NamingProperty = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConnectorScopedInventoryAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorScopedInventoryAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConnectorScopedInventoryAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConnectorScopedInventoryAllOf) SetType(v string) {
	o.Type = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ConnectorScopedInventoryAllOf) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorScopedInventoryAllOf) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ConnectorScopedInventoryAllOf) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ConnectorScopedInventoryAllOf) SetValues(v []string) {
	o.Values = &v
}

func (o ConnectorScopedInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NamingProperty != nil {
		toSerialize["NamingProperty"] = o.NamingProperty
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Values != nil {
		toSerialize["Values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorScopedInventoryAllOf struct {
	value *ConnectorScopedInventoryAllOf
	isSet bool
}

func (v NullableConnectorScopedInventoryAllOf) Get() *ConnectorScopedInventoryAllOf {
	return v.value
}

func (v *NullableConnectorScopedInventoryAllOf) Set(val *ConnectorScopedInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorScopedInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorScopedInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorScopedInventoryAllOf(val *ConnectorScopedInventoryAllOf) *NullableConnectorScopedInventoryAllOf {
	return &NullableConnectorScopedInventoryAllOf{value: val, isSet: true}
}

func (v NullableConnectorScopedInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorScopedInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
