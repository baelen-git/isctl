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

// HyperflexExtFcStoragePolicyAllOf Definition of the list of properties defined in 'hyperflex.ExtFcStoragePolicy', excluding properties defined in parent classes.
type HyperflexExtFcStoragePolicyAllOf struct {
	// Enables or disables external FC storage configuration.
	AdminState      *bool                     `json:"AdminState,omitempty" yaml:"AdminState,omitempty"`
	ExtaTraffic     *HyperflexNamedVsan       `json:"ExtaTraffic,omitempty" yaml:"ExtaTraffic,omitempty"`
	ExtbTraffic     *HyperflexNamedVsan       `json:"ExtbTraffic,omitempty" yaml:"ExtbTraffic,omitempty"`
	WwxnPrefixRange *HyperflexWwxnPrefixRange `json:"WwxnPrefixRange,omitempty" yaml:"WwxnPrefixRange,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty" yaml:"ClusterProfiles,omitempty"`
	Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
}

// NewHyperflexExtFcStoragePolicyAllOf instantiates a new HyperflexExtFcStoragePolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexExtFcStoragePolicyAllOf() *HyperflexExtFcStoragePolicyAllOf {
	this := HyperflexExtFcStoragePolicyAllOf{}
	return &this
}

// NewHyperflexExtFcStoragePolicyAllOfWithDefaults instantiates a new HyperflexExtFcStoragePolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexExtFcStoragePolicyAllOfWithDefaults() *HyperflexExtFcStoragePolicyAllOf {
	this := HyperflexExtFcStoragePolicyAllOf{}
	return &this
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *HyperflexExtFcStoragePolicyAllOf) GetAdminState() bool {
	if o == nil || o.AdminState == nil {
		var ret bool
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) GetAdminStateOk() (*bool, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given bool and assigns it to the AdminState field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetAdminState(v bool) {
	o.AdminState = &v
}

// GetExtaTraffic returns the ExtaTraffic field value if set, zero value otherwise.
func (o *HyperflexExtFcStoragePolicyAllOf) GetExtaTraffic() HyperflexNamedVsan {
	if o == nil || o.ExtaTraffic == nil {
		var ret HyperflexNamedVsan
		return ret
	}
	return *o.ExtaTraffic
}

// GetExtaTrafficOk returns a tuple with the ExtaTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) GetExtaTrafficOk() (*HyperflexNamedVsan, bool) {
	if o == nil || o.ExtaTraffic == nil {
		return nil, false
	}
	return o.ExtaTraffic, true
}

// HasExtaTraffic returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasExtaTraffic() bool {
	if o != nil && o.ExtaTraffic != nil {
		return true
	}

	return false
}

// SetExtaTraffic gets a reference to the given HyperflexNamedVsan and assigns it to the ExtaTraffic field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetExtaTraffic(v HyperflexNamedVsan) {
	o.ExtaTraffic = &v
}

// GetExtbTraffic returns the ExtbTraffic field value if set, zero value otherwise.
func (o *HyperflexExtFcStoragePolicyAllOf) GetExtbTraffic() HyperflexNamedVsan {
	if o == nil || o.ExtbTraffic == nil {
		var ret HyperflexNamedVsan
		return ret
	}
	return *o.ExtbTraffic
}

// GetExtbTrafficOk returns a tuple with the ExtbTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) GetExtbTrafficOk() (*HyperflexNamedVsan, bool) {
	if o == nil || o.ExtbTraffic == nil {
		return nil, false
	}
	return o.ExtbTraffic, true
}

// HasExtbTraffic returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasExtbTraffic() bool {
	if o != nil && o.ExtbTraffic != nil {
		return true
	}

	return false
}

// SetExtbTraffic gets a reference to the given HyperflexNamedVsan and assigns it to the ExtbTraffic field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetExtbTraffic(v HyperflexNamedVsan) {
	o.ExtbTraffic = &v
}

// GetWwxnPrefixRange returns the WwxnPrefixRange field value if set, zero value otherwise.
func (o *HyperflexExtFcStoragePolicyAllOf) GetWwxnPrefixRange() HyperflexWwxnPrefixRange {
	if o == nil || o.WwxnPrefixRange == nil {
		var ret HyperflexWwxnPrefixRange
		return ret
	}
	return *o.WwxnPrefixRange
}

// GetWwxnPrefixRangeOk returns a tuple with the WwxnPrefixRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) GetWwxnPrefixRangeOk() (*HyperflexWwxnPrefixRange, bool) {
	if o == nil || o.WwxnPrefixRange == nil {
		return nil, false
	}
	return o.WwxnPrefixRange, true
}

// HasWwxnPrefixRange returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasWwxnPrefixRange() bool {
	if o != nil && o.WwxnPrefixRange != nil {
		return true
	}

	return false
}

// SetWwxnPrefixRange gets a reference to the given HyperflexWwxnPrefixRange and assigns it to the WwxnPrefixRange field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetWwxnPrefixRange(v HyperflexWwxnPrefixRange) {
	o.WwxnPrefixRange = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexExtFcStoragePolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexExtFcStoragePolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexExtFcStoragePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexExtFcStoragePolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.ExtaTraffic != nil {
		toSerialize["ExtaTraffic"] = o.ExtaTraffic
	}
	if o.ExtbTraffic != nil {
		toSerialize["ExtbTraffic"] = o.ExtbTraffic
	}
	if o.WwxnPrefixRange != nil {
		toSerialize["WwxnPrefixRange"] = o.WwxnPrefixRange
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexExtFcStoragePolicyAllOf struct {
	value *HyperflexExtFcStoragePolicyAllOf
	isSet bool
}

func (v NullableHyperflexExtFcStoragePolicyAllOf) Get() *HyperflexExtFcStoragePolicyAllOf {
	return v.value
}

func (v *NullableHyperflexExtFcStoragePolicyAllOf) Set(val *HyperflexExtFcStoragePolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexExtFcStoragePolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexExtFcStoragePolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexExtFcStoragePolicyAllOf(val *HyperflexExtFcStoragePolicyAllOf) *NullableHyperflexExtFcStoragePolicyAllOf {
	return &NullableHyperflexExtFcStoragePolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexExtFcStoragePolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexExtFcStoragePolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
