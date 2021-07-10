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

// HyperflexBaseClusterAllOf Definition of the list of properties defined in 'hyperflex.BaseCluster', excluding properties defined in parent classes.
type HyperflexBaseClusterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType   string                        `json:"ObjectType"`
	AlarmSummary NullableHyperflexAlarmSummary `json:"AlarmSummary,omitempty"`
	// The number of days remaining before the cluster's storage utilization reaches the recommended capacity limit of 76%. Default value is math.MaxInt32 to indicate that the capacity runway is \"Unknown\" for a cluster that is not connected or with not sufficient data.
	CapacityRunway *int64 `json:"CapacityRunway,omitempty"`
	// The name of this HyperFlex cluster.
	ClusterName *string `json:"ClusterName,omitempty"`
	// This can be a Storage or Compute cluster. A storage cluster contains storage nodes that are used to persist data. A compute cluster contains compute nodes that are used for executing business logic. * `Storage` - Cluster of storage nodes used to persist data. * `Compute` - Cluster of compute nodes used to execute business logic. * `Unknown` - This cluster type is Unknown. Expect Compute or Storage as valid values.
	ClusterPurpose *string `json:"ClusterPurpose,omitempty"`
	// The number of compute nodes that belong to this cluster.
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitempty"`
	// The number of converged nodes that belong to this cluster.
	ConvergedNodeCount *int64 `json:"ConvergedNodeCount,omitempty"`
	// The deployment type of the HyperFlex cluster. The cluster can have one of the following configurations: 1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site. 2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites. 3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes. If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined, the deployment type is set as 'NA' (not available). * `NA` - The deployment type of the HyperFlex cluster is not available. * `Datacenter` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site. * `Stretched Cluster` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites. * `Edge` - The deployment type of a HyperFlex cluster consisting of 2 or more standalone nodes.
	DeploymentType *string `json:"DeploymentType,omitempty"`
	// The type of the drives used for storage in this cluster. * `NA` - The drive type of the HyperFlex cluster is not available. * `All-Flash` - Indicates that this HyperFlex cluster contains flash drives only. * `Hybrid` - Indicates that this HyperFlex cluster contains both flash and hard disk drives.
	DriveType *string `json:"DriveType,omitempty"`
	// The HyperFlex Data or Application Platform version of this cluster.
	HxVersion *string `json:"HxVersion,omitempty"`
	// The version of hypervisor running on this cluster.
	HypervisorVersion *string `json:"HypervisorVersion,omitempty"`
	// The storage capacity in this cluster.
	StorageCapacity *int64 `json:"StorageCapacity,omitempty"`
	// The number of storage nodes that belong to this cluster.
	StorageNodeCount *int64 `json:"StorageNodeCount,omitempty"`
	// The storage utilization is computed based on total capacity and current capacity utilization.
	StorageUtilization *float32 `json:"StorageUtilization,omitempty"`
	// The storage utilization percentage is computed based on total capacity and current capacity utilization.
	UtilizationPercentage *float32 `json:"UtilizationPercentage,omitempty"`
	// The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data.
	UtilizationTrendPercentage *float32                           `json:"UtilizationTrendPercentage,omitempty"`
	AssociatedProfile          *PolicyAbstractProfileRelationship `json:"AssociatedProfile,omitempty"`
	// An array of relationships to hyperflexBaseCluster resources.
	ChildClusters        []HyperflexBaseClusterRelationship `json:"ChildClusters,omitempty"`
	ParentCluster        *HyperflexBaseClusterRelationship  `json:"ParentCluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexBaseClusterAllOf HyperflexBaseClusterAllOf

// NewHyperflexBaseClusterAllOf instantiates a new HyperflexBaseClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexBaseClusterAllOf(classId string, objectType string) *HyperflexBaseClusterAllOf {
	this := HyperflexBaseClusterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var capacityRunway int64 = 2147483647
	this.CapacityRunway = &capacityRunway
	var clusterPurpose string = "Storage"
	this.ClusterPurpose = &clusterPurpose
	var deploymentType string = "NA"
	this.DeploymentType = &deploymentType
	var driveType string = "NA"
	this.DriveType = &driveType
	return &this
}

// NewHyperflexBaseClusterAllOfWithDefaults instantiates a new HyperflexBaseClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexBaseClusterAllOfWithDefaults() *HyperflexBaseClusterAllOf {
	this := HyperflexBaseClusterAllOf{}
	var capacityRunway int64 = 2147483647
	this.CapacityRunway = &capacityRunway
	var clusterPurpose string = "Storage"
	this.ClusterPurpose = &clusterPurpose
	var deploymentType string = "NA"
	this.DeploymentType = &deploymentType
	var driveType string = "NA"
	this.DriveType = &driveType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexBaseClusterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexBaseClusterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexBaseClusterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexBaseClusterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlarmSummary returns the AlarmSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexBaseClusterAllOf) GetAlarmSummary() HyperflexAlarmSummary {
	if o == nil || o.AlarmSummary.Get() == nil {
		var ret HyperflexAlarmSummary
		return ret
	}
	return *o.AlarmSummary.Get()
}

// GetAlarmSummaryOk returns a tuple with the AlarmSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexBaseClusterAllOf) GetAlarmSummaryOk() (*HyperflexAlarmSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlarmSummary.Get(), o.AlarmSummary.IsSet()
}

// HasAlarmSummary returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasAlarmSummary() bool {
	if o != nil && o.AlarmSummary.IsSet() {
		return true
	}

	return false
}

// SetAlarmSummary gets a reference to the given NullableHyperflexAlarmSummary and assigns it to the AlarmSummary field.
func (o *HyperflexBaseClusterAllOf) SetAlarmSummary(v HyperflexAlarmSummary) {
	o.AlarmSummary.Set(&v)
}

// SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil
func (o *HyperflexBaseClusterAllOf) SetAlarmSummaryNil() {
	o.AlarmSummary.Set(nil)
}

// UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
func (o *HyperflexBaseClusterAllOf) UnsetAlarmSummary() {
	o.AlarmSummary.Unset()
}

// GetCapacityRunway returns the CapacityRunway field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetCapacityRunway() int64 {
	if o == nil || o.CapacityRunway == nil {
		var ret int64
		return ret
	}
	return *o.CapacityRunway
}

// GetCapacityRunwayOk returns a tuple with the CapacityRunway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetCapacityRunwayOk() (*int64, bool) {
	if o == nil || o.CapacityRunway == nil {
		return nil, false
	}
	return o.CapacityRunway, true
}

// HasCapacityRunway returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasCapacityRunway() bool {
	if o != nil && o.CapacityRunway != nil {
		return true
	}

	return false
}

// SetCapacityRunway gets a reference to the given int64 and assigns it to the CapacityRunway field.
func (o *HyperflexBaseClusterAllOf) SetCapacityRunway(v int64) {
	o.CapacityRunway = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *HyperflexBaseClusterAllOf) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterPurpose returns the ClusterPurpose field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetClusterPurpose() string {
	if o == nil || o.ClusterPurpose == nil {
		var ret string
		return ret
	}
	return *o.ClusterPurpose
}

// GetClusterPurposeOk returns a tuple with the ClusterPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetClusterPurposeOk() (*string, bool) {
	if o == nil || o.ClusterPurpose == nil {
		return nil, false
	}
	return o.ClusterPurpose, true
}

// HasClusterPurpose returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasClusterPurpose() bool {
	if o != nil && o.ClusterPurpose != nil {
		return true
	}

	return false
}

// SetClusterPurpose gets a reference to the given string and assigns it to the ClusterPurpose field.
func (o *HyperflexBaseClusterAllOf) SetClusterPurpose(v string) {
	o.ClusterPurpose = &v
}

// GetComputeNodeCount returns the ComputeNodeCount field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetComputeNodeCount() int64 {
	if o == nil || o.ComputeNodeCount == nil {
		var ret int64
		return ret
	}
	return *o.ComputeNodeCount
}

// GetComputeNodeCountOk returns a tuple with the ComputeNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetComputeNodeCountOk() (*int64, bool) {
	if o == nil || o.ComputeNodeCount == nil {
		return nil, false
	}
	return o.ComputeNodeCount, true
}

// HasComputeNodeCount returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasComputeNodeCount() bool {
	if o != nil && o.ComputeNodeCount != nil {
		return true
	}

	return false
}

// SetComputeNodeCount gets a reference to the given int64 and assigns it to the ComputeNodeCount field.
func (o *HyperflexBaseClusterAllOf) SetComputeNodeCount(v int64) {
	o.ComputeNodeCount = &v
}

// GetConvergedNodeCount returns the ConvergedNodeCount field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetConvergedNodeCount() int64 {
	if o == nil || o.ConvergedNodeCount == nil {
		var ret int64
		return ret
	}
	return *o.ConvergedNodeCount
}

// GetConvergedNodeCountOk returns a tuple with the ConvergedNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetConvergedNodeCountOk() (*int64, bool) {
	if o == nil || o.ConvergedNodeCount == nil {
		return nil, false
	}
	return o.ConvergedNodeCount, true
}

// HasConvergedNodeCount returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasConvergedNodeCount() bool {
	if o != nil && o.ConvergedNodeCount != nil {
		return true
	}

	return false
}

// SetConvergedNodeCount gets a reference to the given int64 and assigns it to the ConvergedNodeCount field.
func (o *HyperflexBaseClusterAllOf) SetConvergedNodeCount(v int64) {
	o.ConvergedNodeCount = &v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasDeploymentType() bool {
	if o != nil && o.DeploymentType != nil {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *HyperflexBaseClusterAllOf) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetDriveType returns the DriveType field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetDriveType() string {
	if o == nil || o.DriveType == nil {
		var ret string
		return ret
	}
	return *o.DriveType
}

// GetDriveTypeOk returns a tuple with the DriveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetDriveTypeOk() (*string, bool) {
	if o == nil || o.DriveType == nil {
		return nil, false
	}
	return o.DriveType, true
}

// HasDriveType returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasDriveType() bool {
	if o != nil && o.DriveType != nil {
		return true
	}

	return false
}

// SetDriveType gets a reference to the given string and assigns it to the DriveType field.
func (o *HyperflexBaseClusterAllOf) SetDriveType(v string) {
	o.DriveType = &v
}

// GetHxVersion returns the HxVersion field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetHxVersion() string {
	if o == nil || o.HxVersion == nil {
		var ret string
		return ret
	}
	return *o.HxVersion
}

// GetHxVersionOk returns a tuple with the HxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetHxVersionOk() (*string, bool) {
	if o == nil || o.HxVersion == nil {
		return nil, false
	}
	return o.HxVersion, true
}

// HasHxVersion returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasHxVersion() bool {
	if o != nil && o.HxVersion != nil {
		return true
	}

	return false
}

// SetHxVersion gets a reference to the given string and assigns it to the HxVersion field.
func (o *HyperflexBaseClusterAllOf) SetHxVersion(v string) {
	o.HxVersion = &v
}

// GetHypervisorVersion returns the HypervisorVersion field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetHypervisorVersion() string {
	if o == nil || o.HypervisorVersion == nil {
		var ret string
		return ret
	}
	return *o.HypervisorVersion
}

// GetHypervisorVersionOk returns a tuple with the HypervisorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetHypervisorVersionOk() (*string, bool) {
	if o == nil || o.HypervisorVersion == nil {
		return nil, false
	}
	return o.HypervisorVersion, true
}

// HasHypervisorVersion returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasHypervisorVersion() bool {
	if o != nil && o.HypervisorVersion != nil {
		return true
	}

	return false
}

// SetHypervisorVersion gets a reference to the given string and assigns it to the HypervisorVersion field.
func (o *HyperflexBaseClusterAllOf) SetHypervisorVersion(v string) {
	o.HypervisorVersion = &v
}

// GetStorageCapacity returns the StorageCapacity field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetStorageCapacity() int64 {
	if o == nil || o.StorageCapacity == nil {
		var ret int64
		return ret
	}
	return *o.StorageCapacity
}

// GetStorageCapacityOk returns a tuple with the StorageCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetStorageCapacityOk() (*int64, bool) {
	if o == nil || o.StorageCapacity == nil {
		return nil, false
	}
	return o.StorageCapacity, true
}

// HasStorageCapacity returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasStorageCapacity() bool {
	if o != nil && o.StorageCapacity != nil {
		return true
	}

	return false
}

// SetStorageCapacity gets a reference to the given int64 and assigns it to the StorageCapacity field.
func (o *HyperflexBaseClusterAllOf) SetStorageCapacity(v int64) {
	o.StorageCapacity = &v
}

// GetStorageNodeCount returns the StorageNodeCount field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetStorageNodeCount() int64 {
	if o == nil || o.StorageNodeCount == nil {
		var ret int64
		return ret
	}
	return *o.StorageNodeCount
}

// GetStorageNodeCountOk returns a tuple with the StorageNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetStorageNodeCountOk() (*int64, bool) {
	if o == nil || o.StorageNodeCount == nil {
		return nil, false
	}
	return o.StorageNodeCount, true
}

// HasStorageNodeCount returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasStorageNodeCount() bool {
	if o != nil && o.StorageNodeCount != nil {
		return true
	}

	return false
}

// SetStorageNodeCount gets a reference to the given int64 and assigns it to the StorageNodeCount field.
func (o *HyperflexBaseClusterAllOf) SetStorageNodeCount(v int64) {
	o.StorageNodeCount = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetStorageUtilization() float32 {
	if o == nil || o.StorageUtilization == nil {
		var ret float32
		return ret
	}
	return *o.StorageUtilization
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetStorageUtilizationOk() (*float32, bool) {
	if o == nil || o.StorageUtilization == nil {
		return nil, false
	}
	return o.StorageUtilization, true
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization != nil {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given float32 and assigns it to the StorageUtilization field.
func (o *HyperflexBaseClusterAllOf) SetStorageUtilization(v float32) {
	o.StorageUtilization = &v
}

// GetUtilizationPercentage returns the UtilizationPercentage field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetUtilizationPercentage() float32 {
	if o == nil || o.UtilizationPercentage == nil {
		var ret float32
		return ret
	}
	return *o.UtilizationPercentage
}

// GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetUtilizationPercentageOk() (*float32, bool) {
	if o == nil || o.UtilizationPercentage == nil {
		return nil, false
	}
	return o.UtilizationPercentage, true
}

// HasUtilizationPercentage returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasUtilizationPercentage() bool {
	if o != nil && o.UtilizationPercentage != nil {
		return true
	}

	return false
}

// SetUtilizationPercentage gets a reference to the given float32 and assigns it to the UtilizationPercentage field.
func (o *HyperflexBaseClusterAllOf) SetUtilizationPercentage(v float32) {
	o.UtilizationPercentage = &v
}

// GetUtilizationTrendPercentage returns the UtilizationTrendPercentage field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetUtilizationTrendPercentage() float32 {
	if o == nil || o.UtilizationTrendPercentage == nil {
		var ret float32
		return ret
	}
	return *o.UtilizationTrendPercentage
}

// GetUtilizationTrendPercentageOk returns a tuple with the UtilizationTrendPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetUtilizationTrendPercentageOk() (*float32, bool) {
	if o == nil || o.UtilizationTrendPercentage == nil {
		return nil, false
	}
	return o.UtilizationTrendPercentage, true
}

// HasUtilizationTrendPercentage returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasUtilizationTrendPercentage() bool {
	if o != nil && o.UtilizationTrendPercentage != nil {
		return true
	}

	return false
}

// SetUtilizationTrendPercentage gets a reference to the given float32 and assigns it to the UtilizationTrendPercentage field.
func (o *HyperflexBaseClusterAllOf) SetUtilizationTrendPercentage(v float32) {
	o.UtilizationTrendPercentage = &v
}

// GetAssociatedProfile returns the AssociatedProfile field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetAssociatedProfile() PolicyAbstractProfileRelationship {
	if o == nil || o.AssociatedProfile == nil {
		var ret PolicyAbstractProfileRelationship
		return ret
	}
	return *o.AssociatedProfile
}

// GetAssociatedProfileOk returns a tuple with the AssociatedProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetAssociatedProfileOk() (*PolicyAbstractProfileRelationship, bool) {
	if o == nil || o.AssociatedProfile == nil {
		return nil, false
	}
	return o.AssociatedProfile, true
}

// HasAssociatedProfile returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasAssociatedProfile() bool {
	if o != nil && o.AssociatedProfile != nil {
		return true
	}

	return false
}

// SetAssociatedProfile gets a reference to the given PolicyAbstractProfileRelationship and assigns it to the AssociatedProfile field.
func (o *HyperflexBaseClusterAllOf) SetAssociatedProfile(v PolicyAbstractProfileRelationship) {
	o.AssociatedProfile = &v
}

// GetChildClusters returns the ChildClusters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexBaseClusterAllOf) GetChildClusters() []HyperflexBaseClusterRelationship {
	if o == nil {
		var ret []HyperflexBaseClusterRelationship
		return ret
	}
	return o.ChildClusters
}

// GetChildClustersOk returns a tuple with the ChildClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexBaseClusterAllOf) GetChildClustersOk() (*[]HyperflexBaseClusterRelationship, bool) {
	if o == nil || o.ChildClusters == nil {
		return nil, false
	}
	return &o.ChildClusters, true
}

// HasChildClusters returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasChildClusters() bool {
	if o != nil && o.ChildClusters != nil {
		return true
	}

	return false
}

// SetChildClusters gets a reference to the given []HyperflexBaseClusterRelationship and assigns it to the ChildClusters field.
func (o *HyperflexBaseClusterAllOf) SetChildClusters(v []HyperflexBaseClusterRelationship) {
	o.ChildClusters = v
}

// GetParentCluster returns the ParentCluster field value if set, zero value otherwise.
func (o *HyperflexBaseClusterAllOf) GetParentCluster() HyperflexBaseClusterRelationship {
	if o == nil || o.ParentCluster == nil {
		var ret HyperflexBaseClusterRelationship
		return ret
	}
	return *o.ParentCluster
}

// GetParentClusterOk returns a tuple with the ParentCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBaseClusterAllOf) GetParentClusterOk() (*HyperflexBaseClusterRelationship, bool) {
	if o == nil || o.ParentCluster == nil {
		return nil, false
	}
	return o.ParentCluster, true
}

// HasParentCluster returns a boolean if a field has been set.
func (o *HyperflexBaseClusterAllOf) HasParentCluster() bool {
	if o != nil && o.ParentCluster != nil {
		return true
	}

	return false
}

// SetParentCluster gets a reference to the given HyperflexBaseClusterRelationship and assigns it to the ParentCluster field.
func (o *HyperflexBaseClusterAllOf) SetParentCluster(v HyperflexBaseClusterRelationship) {
	o.ParentCluster = &v
}

func (o HyperflexBaseClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AlarmSummary.IsSet() {
		toSerialize["AlarmSummary"] = o.AlarmSummary.Get()
	}
	if o.CapacityRunway != nil {
		toSerialize["CapacityRunway"] = o.CapacityRunway
	}
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.ClusterPurpose != nil {
		toSerialize["ClusterPurpose"] = o.ClusterPurpose
	}
	if o.ComputeNodeCount != nil {
		toSerialize["ComputeNodeCount"] = o.ComputeNodeCount
	}
	if o.ConvergedNodeCount != nil {
		toSerialize["ConvergedNodeCount"] = o.ConvergedNodeCount
	}
	if o.DeploymentType != nil {
		toSerialize["DeploymentType"] = o.DeploymentType
	}
	if o.DriveType != nil {
		toSerialize["DriveType"] = o.DriveType
	}
	if o.HxVersion != nil {
		toSerialize["HxVersion"] = o.HxVersion
	}
	if o.HypervisorVersion != nil {
		toSerialize["HypervisorVersion"] = o.HypervisorVersion
	}
	if o.StorageCapacity != nil {
		toSerialize["StorageCapacity"] = o.StorageCapacity
	}
	if o.StorageNodeCount != nil {
		toSerialize["StorageNodeCount"] = o.StorageNodeCount
	}
	if o.StorageUtilization != nil {
		toSerialize["StorageUtilization"] = o.StorageUtilization
	}
	if o.UtilizationPercentage != nil {
		toSerialize["UtilizationPercentage"] = o.UtilizationPercentage
	}
	if o.UtilizationTrendPercentage != nil {
		toSerialize["UtilizationTrendPercentage"] = o.UtilizationTrendPercentage
	}
	if o.AssociatedProfile != nil {
		toSerialize["AssociatedProfile"] = o.AssociatedProfile
	}
	if o.ChildClusters != nil {
		toSerialize["ChildClusters"] = o.ChildClusters
	}
	if o.ParentCluster != nil {
		toSerialize["ParentCluster"] = o.ParentCluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexBaseClusterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexBaseClusterAllOf := _HyperflexBaseClusterAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexBaseClusterAllOf); err == nil {
		*o = HyperflexBaseClusterAllOf(varHyperflexBaseClusterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlarmSummary")
		delete(additionalProperties, "CapacityRunway")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "ClusterPurpose")
		delete(additionalProperties, "ComputeNodeCount")
		delete(additionalProperties, "ConvergedNodeCount")
		delete(additionalProperties, "DeploymentType")
		delete(additionalProperties, "DriveType")
		delete(additionalProperties, "HxVersion")
		delete(additionalProperties, "HypervisorVersion")
		delete(additionalProperties, "StorageCapacity")
		delete(additionalProperties, "StorageNodeCount")
		delete(additionalProperties, "StorageUtilization")
		delete(additionalProperties, "UtilizationPercentage")
		delete(additionalProperties, "UtilizationTrendPercentage")
		delete(additionalProperties, "AssociatedProfile")
		delete(additionalProperties, "ChildClusters")
		delete(additionalProperties, "ParentCluster")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexBaseClusterAllOf struct {
	value *HyperflexBaseClusterAllOf
	isSet bool
}

func (v NullableHyperflexBaseClusterAllOf) Get() *HyperflexBaseClusterAllOf {
	return v.value
}

func (v *NullableHyperflexBaseClusterAllOf) Set(val *HyperflexBaseClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexBaseClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexBaseClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexBaseClusterAllOf(val *HyperflexBaseClusterAllOf) *NullableHyperflexBaseClusterAllOf {
	return &NullableHyperflexBaseClusterAllOf{value: val, isSet: true}
}

func (v NullableHyperflexBaseClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexBaseClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
