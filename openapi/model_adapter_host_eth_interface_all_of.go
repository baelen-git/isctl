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

// AdapterHostEthInterfaceAllOf Definition of the list of properties defined in 'adapter.HostEthInterface', excluding properties defined in parent classes.
type AdapterHostEthInterfaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin state of the Host Ethernet Interface.
	AdminState *string `json:"AdminState,omitempty"`
	// The Endpoint Config Dn of the Host Ethernet Interface.
	EpDn *string `json:"EpDn,omitempty"`
	// Unique Identifier for an Host Ethernet Interface within the adapter object.
	HostEthInterfaceId *int64 `json:"HostEthInterfaceId,omitempty"`
	// Type of External Ethernet Interface.
	InterfaceType *string `json:"InterfaceType,omitempty"`
	// Mac address of the Host Ethernet Interface.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Name of Host Ethernet Interface.
	Name       *string  `json:"Name,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	// Operability status of Host Ethernet Channel Interface.
	Operability *string `json:"Operability,omitempty"`
	// The factory default Mac address of the Host Ethernet Interface.
	OriginalMacAddress *string `json:"OriginalMacAddress,omitempty"`
	// The PCI address of the Host Ethernet Interface.
	PciAddr *string `json:"PciAddr,omitempty"`
	// The distinguished name of the peer endpoint connected to the Host Ethernet interface.
	PeerDn *string `json:"PeerDn,omitempty"`
	// Virtualization Preference of the Host Ethernet Interface indicating if virtualization is enabled or not.
	VirtualizationPreference *string `json:"VirtualizationPreference,omitempty"`
	// The Virtual Ethernet Interface DN connected to the Host Ethernet Interface.
	VnicDn               *string                              `json:"VnicDn,omitempty"`
	AdapterUnit          *AdapterUnitRelationship             `json:"AdapterUnit,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterHostEthInterfaceAllOf AdapterHostEthInterfaceAllOf

// NewAdapterHostEthInterfaceAllOf instantiates a new AdapterHostEthInterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterHostEthInterfaceAllOf(classId string, objectType string) *AdapterHostEthInterfaceAllOf {
	this := AdapterHostEthInterfaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAdapterHostEthInterfaceAllOfWithDefaults instantiates a new AdapterHostEthInterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterHostEthInterfaceAllOfWithDefaults() *AdapterHostEthInterfaceAllOf {
	this := AdapterHostEthInterfaceAllOf{}
	var classId string = "adapter.HostEthInterface"
	this.ClassId = classId
	var objectType string = "adapter.HostEthInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterHostEthInterfaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterHostEthInterfaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterHostEthInterfaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterHostEthInterfaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *AdapterHostEthInterfaceAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetEpDn returns the EpDn field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetEpDn() string {
	if o == nil || o.EpDn == nil {
		var ret string
		return ret
	}
	return *o.EpDn
}

// GetEpDnOk returns a tuple with the EpDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetEpDnOk() (*string, bool) {
	if o == nil || o.EpDn == nil {
		return nil, false
	}
	return o.EpDn, true
}

// HasEpDn returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasEpDn() bool {
	if o != nil && o.EpDn != nil {
		return true
	}

	return false
}

// SetEpDn gets a reference to the given string and assigns it to the EpDn field.
func (o *AdapterHostEthInterfaceAllOf) SetEpDn(v string) {
	o.EpDn = &v
}

// GetHostEthInterfaceId returns the HostEthInterfaceId field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetHostEthInterfaceId() int64 {
	if o == nil || o.HostEthInterfaceId == nil {
		var ret int64
		return ret
	}
	return *o.HostEthInterfaceId
}

// GetHostEthInterfaceIdOk returns a tuple with the HostEthInterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetHostEthInterfaceIdOk() (*int64, bool) {
	if o == nil || o.HostEthInterfaceId == nil {
		return nil, false
	}
	return o.HostEthInterfaceId, true
}

// HasHostEthInterfaceId returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasHostEthInterfaceId() bool {
	if o != nil && o.HostEthInterfaceId != nil {
		return true
	}

	return false
}

// SetHostEthInterfaceId gets a reference to the given int64 and assigns it to the HostEthInterfaceId field.
func (o *AdapterHostEthInterfaceAllOf) SetHostEthInterfaceId(v int64) {
	o.HostEthInterfaceId = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetInterfaceType() string {
	if o == nil || o.InterfaceType == nil {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || o.InterfaceType == nil {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasInterfaceType() bool {
	if o != nil && o.InterfaceType != nil {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *AdapterHostEthInterfaceAllOf) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *AdapterHostEthInterfaceAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdapterHostEthInterfaceAllOf) SetName(v string) {
	o.Name = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterHostEthInterfaceAllOf) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterHostEthInterfaceAllOf) GetOperReasonOk() (*[]string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return &o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *AdapterHostEthInterfaceAllOf) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *AdapterHostEthInterfaceAllOf) SetOperability(v string) {
	o.Operability = &v
}

// GetOriginalMacAddress returns the OriginalMacAddress field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetOriginalMacAddress() string {
	if o == nil || o.OriginalMacAddress == nil {
		var ret string
		return ret
	}
	return *o.OriginalMacAddress
}

// GetOriginalMacAddressOk returns a tuple with the OriginalMacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetOriginalMacAddressOk() (*string, bool) {
	if o == nil || o.OriginalMacAddress == nil {
		return nil, false
	}
	return o.OriginalMacAddress, true
}

// HasOriginalMacAddress returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasOriginalMacAddress() bool {
	if o != nil && o.OriginalMacAddress != nil {
		return true
	}

	return false
}

// SetOriginalMacAddress gets a reference to the given string and assigns it to the OriginalMacAddress field.
func (o *AdapterHostEthInterfaceAllOf) SetOriginalMacAddress(v string) {
	o.OriginalMacAddress = &v
}

// GetPciAddr returns the PciAddr field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetPciAddr() string {
	if o == nil || o.PciAddr == nil {
		var ret string
		return ret
	}
	return *o.PciAddr
}

// GetPciAddrOk returns a tuple with the PciAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetPciAddrOk() (*string, bool) {
	if o == nil || o.PciAddr == nil {
		return nil, false
	}
	return o.PciAddr, true
}

// HasPciAddr returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasPciAddr() bool {
	if o != nil && o.PciAddr != nil {
		return true
	}

	return false
}

// SetPciAddr gets a reference to the given string and assigns it to the PciAddr field.
func (o *AdapterHostEthInterfaceAllOf) SetPciAddr(v string) {
	o.PciAddr = &v
}

// GetPeerDn returns the PeerDn field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetPeerDn() string {
	if o == nil || o.PeerDn == nil {
		var ret string
		return ret
	}
	return *o.PeerDn
}

// GetPeerDnOk returns a tuple with the PeerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetPeerDnOk() (*string, bool) {
	if o == nil || o.PeerDn == nil {
		return nil, false
	}
	return o.PeerDn, true
}

// HasPeerDn returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasPeerDn() bool {
	if o != nil && o.PeerDn != nil {
		return true
	}

	return false
}

// SetPeerDn gets a reference to the given string and assigns it to the PeerDn field.
func (o *AdapterHostEthInterfaceAllOf) SetPeerDn(v string) {
	o.PeerDn = &v
}

// GetVirtualizationPreference returns the VirtualizationPreference field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetVirtualizationPreference() string {
	if o == nil || o.VirtualizationPreference == nil {
		var ret string
		return ret
	}
	return *o.VirtualizationPreference
}

// GetVirtualizationPreferenceOk returns a tuple with the VirtualizationPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetVirtualizationPreferenceOk() (*string, bool) {
	if o == nil || o.VirtualizationPreference == nil {
		return nil, false
	}
	return o.VirtualizationPreference, true
}

// HasVirtualizationPreference returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasVirtualizationPreference() bool {
	if o != nil && o.VirtualizationPreference != nil {
		return true
	}

	return false
}

// SetVirtualizationPreference gets a reference to the given string and assigns it to the VirtualizationPreference field.
func (o *AdapterHostEthInterfaceAllOf) SetVirtualizationPreference(v string) {
	o.VirtualizationPreference = &v
}

// GetVnicDn returns the VnicDn field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetVnicDn() string {
	if o == nil || o.VnicDn == nil {
		var ret string
		return ret
	}
	return *o.VnicDn
}

// GetVnicDnOk returns a tuple with the VnicDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetVnicDnOk() (*string, bool) {
	if o == nil || o.VnicDn == nil {
		return nil, false
	}
	return o.VnicDn, true
}

// HasVnicDn returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasVnicDn() bool {
	if o != nil && o.VnicDn != nil {
		return true
	}

	return false
}

// SetVnicDn gets a reference to the given string and assigns it to the VnicDn field.
func (o *AdapterHostEthInterfaceAllOf) SetVnicDn(v string) {
	o.VnicDn = &v
}

// GetAdapterUnit returns the AdapterUnit field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetAdapterUnit() AdapterUnitRelationship {
	if o == nil || o.AdapterUnit == nil {
		var ret AdapterUnitRelationship
		return ret
	}
	return *o.AdapterUnit
}

// GetAdapterUnitOk returns a tuple with the AdapterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetAdapterUnitOk() (*AdapterUnitRelationship, bool) {
	if o == nil || o.AdapterUnit == nil {
		return nil, false
	}
	return o.AdapterUnit, true
}

// HasAdapterUnit returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasAdapterUnit() bool {
	if o != nil && o.AdapterUnit != nil {
		return true
	}

	return false
}

// SetAdapterUnit gets a reference to the given AdapterUnitRelationship and assigns it to the AdapterUnit field.
func (o *AdapterHostEthInterfaceAllOf) SetAdapterUnit(v AdapterUnitRelationship) {
	o.AdapterUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *AdapterHostEthInterfaceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AdapterHostEthInterfaceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostEthInterfaceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AdapterHostEthInterfaceAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AdapterHostEthInterfaceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AdapterHostEthInterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.EpDn != nil {
		toSerialize["EpDn"] = o.EpDn
	}
	if o.HostEthInterfaceId != nil {
		toSerialize["HostEthInterfaceId"] = o.HostEthInterfaceId
	}
	if o.InterfaceType != nil {
		toSerialize["InterfaceType"] = o.InterfaceType
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.OriginalMacAddress != nil {
		toSerialize["OriginalMacAddress"] = o.OriginalMacAddress
	}
	if o.PciAddr != nil {
		toSerialize["PciAddr"] = o.PciAddr
	}
	if o.PeerDn != nil {
		toSerialize["PeerDn"] = o.PeerDn
	}
	if o.VirtualizationPreference != nil {
		toSerialize["VirtualizationPreference"] = o.VirtualizationPreference
	}
	if o.VnicDn != nil {
		toSerialize["VnicDn"] = o.VnicDn
	}
	if o.AdapterUnit != nil {
		toSerialize["AdapterUnit"] = o.AdapterUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterHostEthInterfaceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAdapterHostEthInterfaceAllOf := _AdapterHostEthInterfaceAllOf{}

	if err = json.Unmarshal(bytes, &varAdapterHostEthInterfaceAllOf); err == nil {
		*o = AdapterHostEthInterfaceAllOf(varAdapterHostEthInterfaceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "EpDn")
		delete(additionalProperties, "HostEthInterfaceId")
		delete(additionalProperties, "InterfaceType")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "OriginalMacAddress")
		delete(additionalProperties, "PciAddr")
		delete(additionalProperties, "PeerDn")
		delete(additionalProperties, "VirtualizationPreference")
		delete(additionalProperties, "VnicDn")
		delete(additionalProperties, "AdapterUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdapterHostEthInterfaceAllOf struct {
	value *AdapterHostEthInterfaceAllOf
	isSet bool
}

func (v NullableAdapterHostEthInterfaceAllOf) Get() *AdapterHostEthInterfaceAllOf {
	return v.value
}

func (v *NullableAdapterHostEthInterfaceAllOf) Set(val *AdapterHostEthInterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterHostEthInterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterHostEthInterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterHostEthInterfaceAllOf(val *AdapterHostEthInterfaceAllOf) *NullableAdapterHostEthInterfaceAllOf {
	return &NullableAdapterHostEthInterfaceAllOf{value: val, isSet: true}
}

func (v NullableAdapterHostEthInterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterHostEthInterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
