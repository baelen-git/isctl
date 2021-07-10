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

// ComputeIpAddressAllOf Definition of the list of properties defined in 'compute.IpAddress', excluding properties defined in parent classes.
type ComputeIpAddressAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP Address to be used for KVM.
	Address *string `json:"Address,omitempty"`
	// Category of the Kvm IP Address. * `Equipment` - Ip Address assigned to an equipment. * `ServiceProfile` - Ip Address assigned to a Service Profile.
	Category *string `json:"Category,omitempty"`
	// Default gateway property of KVM IP Address.
	DefaultGateway *string `json:"DefaultGateway,omitempty"`
	// The distinguished name for this managed object.
	Dn *string `json:"Dn,omitempty"`
	// HTTP port of an IP Address.
	HttpPort *int64 `json:"HttpPort,omitempty"`
	// Secured HTTP port of an IP Address.
	HttpsPort *int64 `json:"HttpsPort,omitempty"`
	// Port number on which the KVM is running and used for connecting to KVM console.
	KvmPort *int64 `json:"KvmPort,omitempty"`
	// VLAN Identifier of Inband IP Address.
	KvmVlan *int64 `json:"KvmVlan,omitempty"`
	// Name to identify the KVM IP Address. * `Outband` - The user assigned Out of band IP Address. * `Inband` - The user assigned Inband IP Address.
	Name *string `json:"Name,omitempty"`
	// Subnet detail of a KVM IP Address.
	Subnet *string `json:"Subnet,omitempty"`
	// Type of the KVM IP Address. * `MgmtInterface` - Ip Address of a Management Interface. * `VnicIpV4StaticAddr` - Static Ipv4 Address of a Virtual Network Interface. * `VnicIpV4PooledAddr` - Ipv4 Address of a Virtual Network Interface from an address pool. * `VnicIpV4MgmtPooledAddr` - Ipv4 Address of a Virtual Network Interface from a Managed address pool. * `VnicIpV6StaticAddr` - Static Ipv6 Address of a Virtual Network Interface. * `VnicIpV6MgmtPooledAddr` - Ipv6 Address of a Virtual Network Interface from a Managed address pool. * `VnicIpV4ProfDerivedAddr` - Server Profile derived Ipv4 Address of a Virtual Network Interface. * `MgmtIpV6ProfDerivedAddr` - Server Profile derived Ipv6 Address used for accessing server management services.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeIpAddressAllOf ComputeIpAddressAllOf

// NewComputeIpAddressAllOf instantiates a new ComputeIpAddressAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeIpAddressAllOf(classId string, objectType string) *ComputeIpAddressAllOf {
	this := ComputeIpAddressAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var category string = "Equipment"
	this.Category = &category
	var httpPort int64 = 80
	this.HttpPort = &httpPort
	var httpsPort int64 = 443
	this.HttpsPort = &httpsPort
	var kvmPort int64 = 2068
	this.KvmPort = &kvmPort
	var name string = "Outband"
	this.Name = &name
	var type_ string = "MgmtInterface"
	this.Type = &type_
	return &this
}

// NewComputeIpAddressAllOfWithDefaults instantiates a new ComputeIpAddressAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeIpAddressAllOfWithDefaults() *ComputeIpAddressAllOf {
	this := ComputeIpAddressAllOf{}
	var classId string = "compute.IpAddress"
	this.ClassId = classId
	var objectType string = "compute.IpAddress"
	this.ObjectType = objectType
	var category string = "Equipment"
	this.Category = &category
	var httpPort int64 = 80
	this.HttpPort = &httpPort
	var httpsPort int64 = 443
	this.HttpsPort = &httpsPort
	var kvmPort int64 = 2068
	this.KvmPort = &kvmPort
	var name string = "Outband"
	this.Name = &name
	var type_ string = "MgmtInterface"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeIpAddressAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeIpAddressAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeIpAddressAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeIpAddressAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ComputeIpAddressAllOf) SetAddress(v string) {
	o.Address = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ComputeIpAddressAllOf) SetCategory(v string) {
	o.Category = &v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetDefaultGateway() string {
	if o == nil || o.DefaultGateway == nil {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || o.DefaultGateway == nil {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasDefaultGateway() bool {
	if o != nil && o.DefaultGateway != nil {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *ComputeIpAddressAllOf) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *ComputeIpAddressAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetHttpPort returns the HttpPort field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetHttpPort() int64 {
	if o == nil || o.HttpPort == nil {
		var ret int64
		return ret
	}
	return *o.HttpPort
}

// GetHttpPortOk returns a tuple with the HttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetHttpPortOk() (*int64, bool) {
	if o == nil || o.HttpPort == nil {
		return nil, false
	}
	return o.HttpPort, true
}

// HasHttpPort returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasHttpPort() bool {
	if o != nil && o.HttpPort != nil {
		return true
	}

	return false
}

// SetHttpPort gets a reference to the given int64 and assigns it to the HttpPort field.
func (o *ComputeIpAddressAllOf) SetHttpPort(v int64) {
	o.HttpPort = &v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetHttpsPort() int64 {
	if o == nil || o.HttpsPort == nil {
		var ret int64
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetHttpsPortOk() (*int64, bool) {
	if o == nil || o.HttpsPort == nil {
		return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasHttpsPort() bool {
	if o != nil && o.HttpsPort != nil {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given int64 and assigns it to the HttpsPort field.
func (o *ComputeIpAddressAllOf) SetHttpsPort(v int64) {
	o.HttpsPort = &v
}

// GetKvmPort returns the KvmPort field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetKvmPort() int64 {
	if o == nil || o.KvmPort == nil {
		var ret int64
		return ret
	}
	return *o.KvmPort
}

// GetKvmPortOk returns a tuple with the KvmPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetKvmPortOk() (*int64, bool) {
	if o == nil || o.KvmPort == nil {
		return nil, false
	}
	return o.KvmPort, true
}

// HasKvmPort returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasKvmPort() bool {
	if o != nil && o.KvmPort != nil {
		return true
	}

	return false
}

// SetKvmPort gets a reference to the given int64 and assigns it to the KvmPort field.
func (o *ComputeIpAddressAllOf) SetKvmPort(v int64) {
	o.KvmPort = &v
}

// GetKvmVlan returns the KvmVlan field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetKvmVlan() int64 {
	if o == nil || o.KvmVlan == nil {
		var ret int64
		return ret
	}
	return *o.KvmVlan
}

// GetKvmVlanOk returns a tuple with the KvmVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetKvmVlanOk() (*int64, bool) {
	if o == nil || o.KvmVlan == nil {
		return nil, false
	}
	return o.KvmVlan, true
}

// HasKvmVlan returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasKvmVlan() bool {
	if o != nil && o.KvmVlan != nil {
		return true
	}

	return false
}

// SetKvmVlan gets a reference to the given int64 and assigns it to the KvmVlan field.
func (o *ComputeIpAddressAllOf) SetKvmVlan(v int64) {
	o.KvmVlan = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeIpAddressAllOf) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *ComputeIpAddressAllOf) SetSubnet(v string) {
	o.Subnet = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ComputeIpAddressAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeIpAddressAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ComputeIpAddressAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ComputeIpAddressAllOf) SetType(v string) {
	o.Type = &v
}

func (o ComputeIpAddressAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.DefaultGateway != nil {
		toSerialize["DefaultGateway"] = o.DefaultGateway
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.HttpPort != nil {
		toSerialize["HttpPort"] = o.HttpPort
	}
	if o.HttpsPort != nil {
		toSerialize["HttpsPort"] = o.HttpsPort
	}
	if o.KvmPort != nil {
		toSerialize["KvmPort"] = o.KvmPort
	}
	if o.KvmVlan != nil {
		toSerialize["KvmVlan"] = o.KvmVlan
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Subnet != nil {
		toSerialize["Subnet"] = o.Subnet
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeIpAddressAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeIpAddressAllOf := _ComputeIpAddressAllOf{}

	if err = json.Unmarshal(bytes, &varComputeIpAddressAllOf); err == nil {
		*o = ComputeIpAddressAllOf(varComputeIpAddressAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "DefaultGateway")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "HttpPort")
		delete(additionalProperties, "HttpsPort")
		delete(additionalProperties, "KvmPort")
		delete(additionalProperties, "KvmVlan")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Subnet")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeIpAddressAllOf struct {
	value *ComputeIpAddressAllOf
	isSet bool
}

func (v NullableComputeIpAddressAllOf) Get() *ComputeIpAddressAllOf {
	return v.value
}

func (v *NullableComputeIpAddressAllOf) Set(val *ComputeIpAddressAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeIpAddressAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeIpAddressAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeIpAddressAllOf(val *ComputeIpAddressAllOf) *NullableComputeIpAddressAllOf {
	return &NullableComputeIpAddressAllOf{value: val, isSet: true}
}

func (v NullableComputeIpAddressAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeIpAddressAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
