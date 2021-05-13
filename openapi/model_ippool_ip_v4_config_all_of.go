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

// IppoolIpV4ConfigAllOf Definition of the list of properties defined in 'ippool.IpV4Config', excluding properties defined in parent classes.
type IppoolIpV4ConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// IP address of the default IPv4 gateway.
	Gateway *string `json:"Gateway,omitempty" yaml:"Gateway,omitempty"`
	// A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address.
	Netmask *string `json:"Netmask,omitempty" yaml:"Netmask,omitempty"`
	// IP Address of the primary Domain Name System (DNS) server.
	PrimaryDns *string `json:"PrimaryDns,omitempty" yaml:"PrimaryDns,omitempty"`
	// IP Address of the secondary Domain Name System (DNS) server.
	SecondaryDns *string `json:"SecondaryDns,omitempty" yaml:"SecondaryDns,omitempty"`
}

// NewIppoolIpV4ConfigAllOf instantiates a new IppoolIpV4ConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolIpV4ConfigAllOf(classId string, objectType string) *IppoolIpV4ConfigAllOf {
	this := IppoolIpV4ConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIppoolIpV4ConfigAllOfWithDefaults instantiates a new IppoolIpV4ConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolIpV4ConfigAllOfWithDefaults() *IppoolIpV4ConfigAllOf {
	this := IppoolIpV4ConfigAllOf{}
	var classId string = "ippool.IpV4Config"
	this.ClassId = classId
	var objectType string = "ippool.IpV4Config"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolIpV4ConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolIpV4ConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolIpV4ConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IppoolIpV4ConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolIpV4ConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolIpV4ConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *IppoolIpV4ConfigAllOf) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4ConfigAllOf) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *IppoolIpV4ConfigAllOf) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *IppoolIpV4ConfigAllOf) SetGateway(v string) {
	o.Gateway = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *IppoolIpV4ConfigAllOf) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4ConfigAllOf) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *IppoolIpV4ConfigAllOf) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *IppoolIpV4ConfigAllOf) SetNetmask(v string) {
	o.Netmask = &v
}

// GetPrimaryDns returns the PrimaryDns field value if set, zero value otherwise.
func (o *IppoolIpV4ConfigAllOf) GetPrimaryDns() string {
	if o == nil || o.PrimaryDns == nil {
		var ret string
		return ret
	}
	return *o.PrimaryDns
}

// GetPrimaryDnsOk returns a tuple with the PrimaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4ConfigAllOf) GetPrimaryDnsOk() (*string, bool) {
	if o == nil || o.PrimaryDns == nil {
		return nil, false
	}
	return o.PrimaryDns, true
}

// HasPrimaryDns returns a boolean if a field has been set.
func (o *IppoolIpV4ConfigAllOf) HasPrimaryDns() bool {
	if o != nil && o.PrimaryDns != nil {
		return true
	}

	return false
}

// SetPrimaryDns gets a reference to the given string and assigns it to the PrimaryDns field.
func (o *IppoolIpV4ConfigAllOf) SetPrimaryDns(v string) {
	o.PrimaryDns = &v
}

// GetSecondaryDns returns the SecondaryDns field value if set, zero value otherwise.
func (o *IppoolIpV4ConfigAllOf) GetSecondaryDns() string {
	if o == nil || o.SecondaryDns == nil {
		var ret string
		return ret
	}
	return *o.SecondaryDns
}

// GetSecondaryDnsOk returns a tuple with the SecondaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4ConfigAllOf) GetSecondaryDnsOk() (*string, bool) {
	if o == nil || o.SecondaryDns == nil {
		return nil, false
	}
	return o.SecondaryDns, true
}

// HasSecondaryDns returns a boolean if a field has been set.
func (o *IppoolIpV4ConfigAllOf) HasSecondaryDns() bool {
	if o != nil && o.SecondaryDns != nil {
		return true
	}

	return false
}

// SetSecondaryDns gets a reference to the given string and assigns it to the SecondaryDns field.
func (o *IppoolIpV4ConfigAllOf) SetSecondaryDns(v string) {
	o.SecondaryDns = &v
}

func (o IppoolIpV4ConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.Netmask != nil {
		toSerialize["Netmask"] = o.Netmask
	}
	if o.PrimaryDns != nil {
		toSerialize["PrimaryDns"] = o.PrimaryDns
	}
	if o.SecondaryDns != nil {
		toSerialize["SecondaryDns"] = o.SecondaryDns
	}
	return json.Marshal(toSerialize)
}

type NullableIppoolIpV4ConfigAllOf struct {
	value *IppoolIpV4ConfigAllOf
	isSet bool
}

func (v NullableIppoolIpV4ConfigAllOf) Get() *IppoolIpV4ConfigAllOf {
	return v.value
}

func (v *NullableIppoolIpV4ConfigAllOf) Set(val *IppoolIpV4ConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolIpV4ConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolIpV4ConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolIpV4ConfigAllOf(val *IppoolIpV4ConfigAllOf) *NullableIppoolIpV4ConfigAllOf {
	return &NullableIppoolIpV4ConfigAllOf{value: val, isSet: true}
}

func (v NullableIppoolIpV4ConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolIpV4ConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
