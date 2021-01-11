/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-12-08T20:53:20Z.
 *
 * API version: 1.0.9-2908
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// FirmwareComponentMeta Contains the details for each component in the HSU bundle catalog.
type FirmwareComponentMeta struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The name of the component in the compressed HSU bundle.
	ComponentLabel *string `json:"ComponentLabel,omitempty" yaml:"ComponentLabel,omitempty"`
	// The type of component image within the distributable. * `ALL` - This represents all the components. * `ALL,HDD` - This represents all the components plus the HDDs. * `None` - This represents none of the components. * `NXOS` - This represents NXOS components. * `IOM` - This represents IOM components. * `PSU` - This represents PSU components. * `CIMC` - This represents CIMC components. * `BIOS` - This represents BIOS components. * `PCIE` - This represents PCIE components. * `Drive` - This represents Storage components. * `DIMM` - This represents DIMM components. * `BoardController` - This represents Board Controller components. * `StorageController` - This represents Storage Controller components. * `HBA` - This represents HBA components. * `GPU` - This represents GPU components. * `SasExpander` - This represents SasExpander components. * `MSwitch` - This represents mSwitch components. * `CMC` - This represents CMC components.
	ComponentType *string `json:"ComponentType,omitempty" yaml:"ComponentType,omitempty"`
	// This shows the description of component image within the distributable.
	Description *string `json:"Description,omitempty" yaml:"Description,omitempty"`
	// The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle. * `None` - Indicates that the component did not receive a disruption request. * `HostReboot` - Indicates that the component received a host reboot request. * `EndpointReboot` - Indicates that the component received an end point reboot request. * `ManualPowerCycle` - Indicates that the component received a manual power cycle request. * `AutomaticPowerCycle` - Indicates that the component received an automatic power cycle request.
	Disruption *string `json:"Disruption,omitempty" yaml:"Disruption,omitempty"`
	// This shows the path of component image within the distributable.
	ImagePath *string `json:"ImagePath,omitempty" yaml:"ImagePath,omitempty"`
	// If set, the component can be updated through out-of-band management, else, is updated through host service utility boot.
	IsOobSupported *bool `json:"IsOobSupported,omitempty" yaml:"IsOobSupported,omitempty"`
	// The model of the component image in the distributable.
	Model            *string  `json:"Model,omitempty" yaml:"Model,omitempty"`
	OobManageability []string `json:"OobManageability,omitempty" yaml:"OobManageability,omitempty"`
	// The image version of components packaged in the distributable.
	PackedVersion *string `json:"PackedVersion,omitempty" yaml:"PackedVersion,omitempty"`
	// The redfish target for each component.
	RedfishUrl *string `json:"RedfishUrl,omitempty" yaml:"RedfishUrl,omitempty"`
	// The version of component image in the distributable.
	Vendor *string `json:"Vendor,omitempty" yaml:"Vendor,omitempty"`
}

// NewFirmwareComponentMeta instantiates a new FirmwareComponentMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareComponentMeta(classId string, objectType string) *FirmwareComponentMeta {
	this := FirmwareComponentMeta{}
	this.ClassId = classId
	this.ObjectType = objectType
	var componentType string = "ALL"
	this.ComponentType = &componentType
	var disruption string = "None"
	this.Disruption = &disruption
	return &this
}

// NewFirmwareComponentMetaWithDefaults instantiates a new FirmwareComponentMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareComponentMetaWithDefaults() *FirmwareComponentMeta {
	this := FirmwareComponentMeta{}
	var classId string = "firmware.ComponentMeta"
	this.ClassId = classId
	var objectType string = "firmware.ComponentMeta"
	this.ObjectType = objectType
	var componentType string = "ALL"
	this.ComponentType = &componentType
	var disruption string = "None"
	this.Disruption = &disruption
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareComponentMeta) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareComponentMeta) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareComponentMeta) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareComponentMeta) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComponentLabel returns the ComponentLabel field value if set, zero value otherwise.
func (o *FirmwareComponentMeta) GetComponentLabel() string {
	if o == nil || o.ComponentLabel == nil {
		var ret string
		return ret
	}
	return *o.ComponentLabel
}

// GetComponentLabelOk returns a tuple with the ComponentLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetComponentLabelOk() (*string, bool) {
	if o == nil || o.ComponentLabel == nil {
		return nil, false
	}
	return o.ComponentLabel, true
}

// HasComponentLabel returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasComponentLabel() bool {
	if o != nil && o.ComponentLabel != nil {
		return true
	}

	return false
}

// SetComponentLabel gets a reference to the given string and assigns it to the ComponentLabel field.
func (o *FirmwareComponentMeta) SetComponentLabel(v string) {
	o.ComponentLabel = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *FirmwareComponentMeta) GetComponentType() string {
	if o == nil || o.ComponentType == nil {
		var ret string
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetComponentTypeOk() (*string, bool) {
	if o == nil || o.ComponentType == nil {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasComponentType() bool {
	if o != nil && o.ComponentType != nil {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given string and assigns it to the ComponentType field.
func (o *FirmwareComponentMeta) SetComponentType(v string) {
	o.ComponentType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FirmwareComponentMeta) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FirmwareComponentMeta) SetDescription(v string) {
	o.Description = &v
}

// GetDisruption returns the Disruption field value if set, zero value otherwise.
func (o *FirmwareComponentMeta) GetDisruption() string {
	if o == nil || o.Disruption == nil {
		var ret string
		return ret
	}
	return *o.Disruption
}

// GetDisruptionOk returns a tuple with the Disruption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetDisruptionOk() (*string, bool) {
	if o == nil || o.Disruption == nil {
		return nil, false
	}
	return o.Disruption, true
}

// HasDisruption returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasDisruption() bool {
	if o != nil && o.Disruption != nil {
		return true
	}

	return false
}

// SetDisruption gets a reference to the given string and assigns it to the Disruption field.
func (o *FirmwareComponentMeta) SetDisruption(v string) {
	o.Disruption = &v
}

// GetImagePath returns the ImagePath field value if set, zero value otherwise.
func (o *FirmwareComponentMeta) GetImagePath() string {
	if o == nil || o.ImagePath == nil {
		var ret string
		return ret
	}
	return *o.ImagePath
}

// GetImagePathOk returns a tuple with the ImagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetImagePathOk() (*string, bool) {
	if o == nil || o.ImagePath == nil {
		return nil, false
	}
	return o.ImagePath, true
}

// HasImagePath returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasImagePath() bool {
	if o != nil && o.ImagePath != nil {
		return true
	}

	return false
}

// SetImagePath gets a reference to the given string and assigns it to the ImagePath field.
func (o *FirmwareComponentMeta) SetImagePath(v string) {
	o.ImagePath = &v
}

// GetIsOobSupported returns the IsOobSupported field value if set, zero value otherwise.
func (o *FirmwareComponentMeta) GetIsOobSupported() bool {
	if o == nil || o.IsOobSupported == nil {
		var ret bool
		return ret
	}
	return *o.IsOobSupported
}

// GetIsOobSupportedOk returns a tuple with the IsOobSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetIsOobSupportedOk() (*bool, bool) {
	if o == nil || o.IsOobSupported == nil {
		return nil, false
	}
	return o.IsOobSupported, true
}

// HasIsOobSupported returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasIsOobSupported() bool {
	if o != nil && o.IsOobSupported != nil {
		return true
	}

	return false
}

// SetIsOobSupported gets a reference to the given bool and assigns it to the IsOobSupported field.
func (o *FirmwareComponentMeta) SetIsOobSupported(v bool) {
	o.IsOobSupported = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *FirmwareComponentMeta) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *FirmwareComponentMeta) SetModel(v string) {
	o.Model = &v
}

// GetOobManageability returns the OobManageability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareComponentMeta) GetOobManageability() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OobManageability
}

// GetOobManageabilityOk returns a tuple with the OobManageability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareComponentMeta) GetOobManageabilityOk() (*[]string, bool) {
	if o == nil || o.OobManageability == nil {
		return nil, false
	}
	return &o.OobManageability, true
}

// HasOobManageability returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasOobManageability() bool {
	if o != nil && o.OobManageability != nil {
		return true
	}

	return false
}

// SetOobManageability gets a reference to the given []string and assigns it to the OobManageability field.
func (o *FirmwareComponentMeta) SetOobManageability(v []string) {
	o.OobManageability = v
}

// GetPackedVersion returns the PackedVersion field value if set, zero value otherwise.
func (o *FirmwareComponentMeta) GetPackedVersion() string {
	if o == nil || o.PackedVersion == nil {
		var ret string
		return ret
	}
	return *o.PackedVersion
}

// GetPackedVersionOk returns a tuple with the PackedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetPackedVersionOk() (*string, bool) {
	if o == nil || o.PackedVersion == nil {
		return nil, false
	}
	return o.PackedVersion, true
}

// HasPackedVersion returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasPackedVersion() bool {
	if o != nil && o.PackedVersion != nil {
		return true
	}

	return false
}

// SetPackedVersion gets a reference to the given string and assigns it to the PackedVersion field.
func (o *FirmwareComponentMeta) SetPackedVersion(v string) {
	o.PackedVersion = &v
}

// GetRedfishUrl returns the RedfishUrl field value if set, zero value otherwise.
func (o *FirmwareComponentMeta) GetRedfishUrl() string {
	if o == nil || o.RedfishUrl == nil {
		var ret string
		return ret
	}
	return *o.RedfishUrl
}

// GetRedfishUrlOk returns a tuple with the RedfishUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetRedfishUrlOk() (*string, bool) {
	if o == nil || o.RedfishUrl == nil {
		return nil, false
	}
	return o.RedfishUrl, true
}

// HasRedfishUrl returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasRedfishUrl() bool {
	if o != nil && o.RedfishUrl != nil {
		return true
	}

	return false
}

// SetRedfishUrl gets a reference to the given string and assigns it to the RedfishUrl field.
func (o *FirmwareComponentMeta) SetRedfishUrl(v string) {
	o.RedfishUrl = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *FirmwareComponentMeta) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMeta) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *FirmwareComponentMeta) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *FirmwareComponentMeta) SetVendor(v string) {
	o.Vendor = &v
}

func (o FirmwareComponentMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ComponentLabel != nil {
		toSerialize["ComponentLabel"] = o.ComponentLabel
	}
	if o.ComponentType != nil {
		toSerialize["ComponentType"] = o.ComponentType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Disruption != nil {
		toSerialize["Disruption"] = o.Disruption
	}
	if o.ImagePath != nil {
		toSerialize["ImagePath"] = o.ImagePath
	}
	if o.IsOobSupported != nil {
		toSerialize["IsOobSupported"] = o.IsOobSupported
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.OobManageability != nil {
		toSerialize["OobManageability"] = o.OobManageability
	}
	if o.PackedVersion != nil {
		toSerialize["PackedVersion"] = o.PackedVersion
	}
	if o.RedfishUrl != nil {
		toSerialize["RedfishUrl"] = o.RedfishUrl
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableFirmwareComponentMeta struct {
	value *FirmwareComponentMeta
	isSet bool
}

func (v NullableFirmwareComponentMeta) Get() *FirmwareComponentMeta {
	return v.value
}

func (v *NullableFirmwareComponentMeta) Set(val *FirmwareComponentMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareComponentMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareComponentMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareComponentMeta(val *FirmwareComponentMeta) *NullableFirmwareComponentMeta {
	return &NullableFirmwareComponentMeta{value: val, isSet: true}
}

func (v NullableFirmwareComponentMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareComponentMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
