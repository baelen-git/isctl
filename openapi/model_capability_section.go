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

// CapabilitySection Refers to a section in the capability catalog. A capability catalog is divided into sections which are managed independently.
type CapabilitySection struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// Administrative action to initialize/load the catalog section from a particular source. * `None` - Nil value to indicate that no action is required. * `LoadLocal` - Load the catalog file packaged with the service. * `LoadIntersightRepository` - Load a catalog file hosted in the Intersight Repository.
	Action *string `json:"Action,omitempty" yaml:"Action,omitempty"`
	// The catalog name reference.
	CatalogName *string `json:"CatalogName,omitempty" yaml:"CatalogName,omitempty"`
	// A unique name for the section inside a catalog.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// The configured source for this section of the catalog. * `Local` - The catalog file is packaged with the service. * `IntersightRepository` - The catalog file is hosted in the Intersight Repository.
	Source *string `json:"Source,omitempty" yaml:"Source,omitempty"`
	// Version of the section inside a catalog.
	Version *string                        `json:"Version,omitempty" yaml:"Version,omitempty"`
	Catalog *CapabilityCatalogRelationship `json:"Catalog,omitempty" yaml:"Catalog,omitempty"`
}

// NewCapabilitySection instantiates a new CapabilitySection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySection() *CapabilitySection {
	this := CapabilitySection{}
	var action string = "None"
	this.Action = &action
	var source string = "Local"
	this.Source = &source
	return &this
}

// NewCapabilitySectionWithDefaults instantiates a new CapabilitySection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySectionWithDefaults() *CapabilitySection {
	this := CapabilitySection{}
	var action string = "None"
	this.Action = &action
	var source string = "Local"
	this.Source = &source
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CapabilitySection) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySection) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CapabilitySection) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *CapabilitySection) SetAction(v string) {
	o.Action = &v
}

// GetCatalogName returns the CatalogName field value if set, zero value otherwise.
func (o *CapabilitySection) GetCatalogName() string {
	if o == nil || o.CatalogName == nil {
		var ret string
		return ret
	}
	return *o.CatalogName
}

// GetCatalogNameOk returns a tuple with the CatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySection) GetCatalogNameOk() (*string, bool) {
	if o == nil || o.CatalogName == nil {
		return nil, false
	}
	return o.CatalogName, true
}

// HasCatalogName returns a boolean if a field has been set.
func (o *CapabilitySection) HasCatalogName() bool {
	if o != nil && o.CatalogName != nil {
		return true
	}

	return false
}

// SetCatalogName gets a reference to the given string and assigns it to the CatalogName field.
func (o *CapabilitySection) SetCatalogName(v string) {
	o.CatalogName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CapabilitySection) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySection) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CapabilitySection) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CapabilitySection) SetName(v string) {
	o.Name = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CapabilitySection) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySection) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CapabilitySection) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CapabilitySection) SetSource(v string) {
	o.Source = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CapabilitySection) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySection) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CapabilitySection) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CapabilitySection) SetVersion(v string) {
	o.Version = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *CapabilitySection) GetCatalog() CapabilityCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret CapabilityCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySection) GetCatalogOk() (*CapabilityCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *CapabilitySection) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given CapabilityCatalogRelationship and assigns it to the Catalog field.
func (o *CapabilitySection) SetCatalog(v CapabilityCatalogRelationship) {
	o.Catalog = &v
}

func (o CapabilitySection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.CatalogName != nil {
		toSerialize["CatalogName"] = o.CatalogName
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	return json.Marshal(toSerialize)
}

type NullableCapabilitySection struct {
	value *CapabilitySection
	isSet bool
}

func (v NullableCapabilitySection) Get() *CapabilitySection {
	return v.value
}

func (v *NullableCapabilitySection) Set(val *CapabilitySection) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySection) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySection(val *CapabilitySection) *NullableCapabilitySection {
	return &NullableCapabilitySection{value: val, isSet: true}
}

func (v NullableCapabilitySection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
