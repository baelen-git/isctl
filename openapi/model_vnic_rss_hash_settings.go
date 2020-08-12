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

// VnicRssHashSettings The RSS Hash parameters help the traffic distribution across the Receive Queues based on the IP address (IPv4 or IPv6) and TCP Port numbers. These options help ensure that a single flow is directed to the same receive queue ensuring in-order delivery.
type VnicRssHashSettings struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// When enabled, the IPv4 address is used for traffic distribution.
	Ipv4Hash *bool `json:"Ipv4Hash,omitempty" yaml:"Ipv4Hash,omitempty"`
	// When enabled, the IPv6 extensions are used for traffic distribution.
	Ipv6ExtHash *bool `json:"Ipv6ExtHash,omitempty" yaml:"Ipv6ExtHash,omitempty"`
	// When enabled, the IPv6 address is used for traffic distribution.
	Ipv6Hash *bool `json:"Ipv6Hash,omitempty" yaml:"Ipv6Hash,omitempty"`
	// When enabled, both the IPv4 address and TCP port number are used for traffic distribution.
	TcpIpv4Hash *bool `json:"TcpIpv4Hash,omitempty" yaml:"TcpIpv4Hash,omitempty"`
	// When enabled, both the IPv6 extensions and TCP port number are used for traffic distribution.
	TcpIpv6ExtHash *bool `json:"TcpIpv6ExtHash,omitempty" yaml:"TcpIpv6ExtHash,omitempty"`
	// When enabled, both the IPv6 address and TCP port number are used for traffic distribution.
	TcpIpv6Hash *bool `json:"TcpIpv6Hash,omitempty" yaml:"TcpIpv6Hash,omitempty"`
	// When enabled, both the IPv4 address and UDP port number are used for traffic distribution.
	UdpIpv4Hash *bool `json:"UdpIpv4Hash,omitempty" yaml:"UdpIpv4Hash,omitempty"`
	// When enabled, both the IPv6 address and UDP port number are used for traffic distribution.
	UdpIpv6Hash *bool `json:"UdpIpv6Hash,omitempty" yaml:"UdpIpv6Hash,omitempty"`
}

// NewVnicRssHashSettings instantiates a new VnicRssHashSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicRssHashSettings() *VnicRssHashSettings {
	this := VnicRssHashSettings{}
	return &this
}

// NewVnicRssHashSettingsWithDefaults instantiates a new VnicRssHashSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicRssHashSettingsWithDefaults() *VnicRssHashSettings {
	this := VnicRssHashSettings{}
	return &this
}

// GetIpv4Hash returns the Ipv4Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettings) GetIpv4Hash() bool {
	if o == nil || o.Ipv4Hash == nil {
		var ret bool
		return ret
	}
	return *o.Ipv4Hash
}

// GetIpv4HashOk returns a tuple with the Ipv4Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettings) GetIpv4HashOk() (*bool, bool) {
	if o == nil || o.Ipv4Hash == nil {
		return nil, false
	}
	return o.Ipv4Hash, true
}

// HasIpv4Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettings) HasIpv4Hash() bool {
	if o != nil && o.Ipv4Hash != nil {
		return true
	}

	return false
}

// SetIpv4Hash gets a reference to the given bool and assigns it to the Ipv4Hash field.
func (o *VnicRssHashSettings) SetIpv4Hash(v bool) {
	o.Ipv4Hash = &v
}

// GetIpv6ExtHash returns the Ipv6ExtHash field value if set, zero value otherwise.
func (o *VnicRssHashSettings) GetIpv6ExtHash() bool {
	if o == nil || o.Ipv6ExtHash == nil {
		var ret bool
		return ret
	}
	return *o.Ipv6ExtHash
}

// GetIpv6ExtHashOk returns a tuple with the Ipv6ExtHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettings) GetIpv6ExtHashOk() (*bool, bool) {
	if o == nil || o.Ipv6ExtHash == nil {
		return nil, false
	}
	return o.Ipv6ExtHash, true
}

// HasIpv6ExtHash returns a boolean if a field has been set.
func (o *VnicRssHashSettings) HasIpv6ExtHash() bool {
	if o != nil && o.Ipv6ExtHash != nil {
		return true
	}

	return false
}

// SetIpv6ExtHash gets a reference to the given bool and assigns it to the Ipv6ExtHash field.
func (o *VnicRssHashSettings) SetIpv6ExtHash(v bool) {
	o.Ipv6ExtHash = &v
}

// GetIpv6Hash returns the Ipv6Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettings) GetIpv6Hash() bool {
	if o == nil || o.Ipv6Hash == nil {
		var ret bool
		return ret
	}
	return *o.Ipv6Hash
}

// GetIpv6HashOk returns a tuple with the Ipv6Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettings) GetIpv6HashOk() (*bool, bool) {
	if o == nil || o.Ipv6Hash == nil {
		return nil, false
	}
	return o.Ipv6Hash, true
}

// HasIpv6Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettings) HasIpv6Hash() bool {
	if o != nil && o.Ipv6Hash != nil {
		return true
	}

	return false
}

// SetIpv6Hash gets a reference to the given bool and assigns it to the Ipv6Hash field.
func (o *VnicRssHashSettings) SetIpv6Hash(v bool) {
	o.Ipv6Hash = &v
}

// GetTcpIpv4Hash returns the TcpIpv4Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettings) GetTcpIpv4Hash() bool {
	if o == nil || o.TcpIpv4Hash == nil {
		var ret bool
		return ret
	}
	return *o.TcpIpv4Hash
}

// GetTcpIpv4HashOk returns a tuple with the TcpIpv4Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettings) GetTcpIpv4HashOk() (*bool, bool) {
	if o == nil || o.TcpIpv4Hash == nil {
		return nil, false
	}
	return o.TcpIpv4Hash, true
}

// HasTcpIpv4Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettings) HasTcpIpv4Hash() bool {
	if o != nil && o.TcpIpv4Hash != nil {
		return true
	}

	return false
}

// SetTcpIpv4Hash gets a reference to the given bool and assigns it to the TcpIpv4Hash field.
func (o *VnicRssHashSettings) SetTcpIpv4Hash(v bool) {
	o.TcpIpv4Hash = &v
}

// GetTcpIpv6ExtHash returns the TcpIpv6ExtHash field value if set, zero value otherwise.
func (o *VnicRssHashSettings) GetTcpIpv6ExtHash() bool {
	if o == nil || o.TcpIpv6ExtHash == nil {
		var ret bool
		return ret
	}
	return *o.TcpIpv6ExtHash
}

// GetTcpIpv6ExtHashOk returns a tuple with the TcpIpv6ExtHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettings) GetTcpIpv6ExtHashOk() (*bool, bool) {
	if o == nil || o.TcpIpv6ExtHash == nil {
		return nil, false
	}
	return o.TcpIpv6ExtHash, true
}

// HasTcpIpv6ExtHash returns a boolean if a field has been set.
func (o *VnicRssHashSettings) HasTcpIpv6ExtHash() bool {
	if o != nil && o.TcpIpv6ExtHash != nil {
		return true
	}

	return false
}

// SetTcpIpv6ExtHash gets a reference to the given bool and assigns it to the TcpIpv6ExtHash field.
func (o *VnicRssHashSettings) SetTcpIpv6ExtHash(v bool) {
	o.TcpIpv6ExtHash = &v
}

// GetTcpIpv6Hash returns the TcpIpv6Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettings) GetTcpIpv6Hash() bool {
	if o == nil || o.TcpIpv6Hash == nil {
		var ret bool
		return ret
	}
	return *o.TcpIpv6Hash
}

// GetTcpIpv6HashOk returns a tuple with the TcpIpv6Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettings) GetTcpIpv6HashOk() (*bool, bool) {
	if o == nil || o.TcpIpv6Hash == nil {
		return nil, false
	}
	return o.TcpIpv6Hash, true
}

// HasTcpIpv6Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettings) HasTcpIpv6Hash() bool {
	if o != nil && o.TcpIpv6Hash != nil {
		return true
	}

	return false
}

// SetTcpIpv6Hash gets a reference to the given bool and assigns it to the TcpIpv6Hash field.
func (o *VnicRssHashSettings) SetTcpIpv6Hash(v bool) {
	o.TcpIpv6Hash = &v
}

// GetUdpIpv4Hash returns the UdpIpv4Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettings) GetUdpIpv4Hash() bool {
	if o == nil || o.UdpIpv4Hash == nil {
		var ret bool
		return ret
	}
	return *o.UdpIpv4Hash
}

// GetUdpIpv4HashOk returns a tuple with the UdpIpv4Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettings) GetUdpIpv4HashOk() (*bool, bool) {
	if o == nil || o.UdpIpv4Hash == nil {
		return nil, false
	}
	return o.UdpIpv4Hash, true
}

// HasUdpIpv4Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettings) HasUdpIpv4Hash() bool {
	if o != nil && o.UdpIpv4Hash != nil {
		return true
	}

	return false
}

// SetUdpIpv4Hash gets a reference to the given bool and assigns it to the UdpIpv4Hash field.
func (o *VnicRssHashSettings) SetUdpIpv4Hash(v bool) {
	o.UdpIpv4Hash = &v
}

// GetUdpIpv6Hash returns the UdpIpv6Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettings) GetUdpIpv6Hash() bool {
	if o == nil || o.UdpIpv6Hash == nil {
		var ret bool
		return ret
	}
	return *o.UdpIpv6Hash
}

// GetUdpIpv6HashOk returns a tuple with the UdpIpv6Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettings) GetUdpIpv6HashOk() (*bool, bool) {
	if o == nil || o.UdpIpv6Hash == nil {
		return nil, false
	}
	return o.UdpIpv6Hash, true
}

// HasUdpIpv6Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettings) HasUdpIpv6Hash() bool {
	if o != nil && o.UdpIpv6Hash != nil {
		return true
	}

	return false
}

// SetUdpIpv6Hash gets a reference to the given bool and assigns it to the UdpIpv6Hash field.
func (o *VnicRssHashSettings) SetUdpIpv6Hash(v bool) {
	o.UdpIpv6Hash = &v
}

func (o VnicRssHashSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Ipv4Hash != nil {
		toSerialize["Ipv4Hash"] = o.Ipv4Hash
	}
	if o.Ipv6ExtHash != nil {
		toSerialize["Ipv6ExtHash"] = o.Ipv6ExtHash
	}
	if o.Ipv6Hash != nil {
		toSerialize["Ipv6Hash"] = o.Ipv6Hash
	}
	if o.TcpIpv4Hash != nil {
		toSerialize["TcpIpv4Hash"] = o.TcpIpv4Hash
	}
	if o.TcpIpv6ExtHash != nil {
		toSerialize["TcpIpv6ExtHash"] = o.TcpIpv6ExtHash
	}
	if o.TcpIpv6Hash != nil {
		toSerialize["TcpIpv6Hash"] = o.TcpIpv6Hash
	}
	if o.UdpIpv4Hash != nil {
		toSerialize["UdpIpv4Hash"] = o.UdpIpv4Hash
	}
	if o.UdpIpv6Hash != nil {
		toSerialize["UdpIpv6Hash"] = o.UdpIpv6Hash
	}
	return json.Marshal(toSerialize)
}

type NullableVnicRssHashSettings struct {
	value *VnicRssHashSettings
	isSet bool
}

func (v NullableVnicRssHashSettings) Get() *VnicRssHashSettings {
	return v.value
}

func (v *NullableVnicRssHashSettings) Set(val *VnicRssHashSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicRssHashSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicRssHashSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicRssHashSettings(val *VnicRssHashSettings) *NullableVnicRssHashSettings {
	return &NullableVnicRssHashSettings{value: val, isSet: true}
}

func (v NullableVnicRssHashSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicRssHashSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
