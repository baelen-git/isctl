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

// CondHclStatus The HCL status of a managed object after we have validated the managed object components' firmware and drivers against the HCL.
type CondHclStatus struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The overall status for the components found in the HCL. This will provide the HCL validation status for all the components. It can be one of the following. \"Validated\" - all the components hardware/software profiles are listed in the HCL. \"Not-Listed\" - one or more components hardware/software profiles are not listed in the HCL \"Incomplete\" - the components are not evaluated as the server's software/hardware profiles are not listed in the HCL. \"Not-Evaluated\" - The components are not evaluated against the HCL because it is exempted. * `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL. * `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL. * `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL. * `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.
	ComponentStatus *string `json:"ComponentStatus,omitempty" yaml:"ComponentStatus,omitempty"`
	// The server model, processor and firmware are considered as part of the hardware profile for the server. This will provide the HCL validation status for the hardware profile. For the failure reason see the serverReason property. The status can be one of the following \"Validated\" - The server model, processor and firmware combination is listed in the HCL \"Not-Listed\" - The server model, processor and firmware combination is not listed in the HCL. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted. * `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL. * `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL. * `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL. * `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.
	HardwareStatus *string `json:"HardwareStatus,omitempty" yaml:"HardwareStatus,omitempty"`
	// The current CIMC version for the server normalized for querying HCL data. It is empty if we are missing this information.
	HclFirmwareVersion *string `json:"HclFirmwareVersion,omitempty" yaml:"HclFirmwareVersion,omitempty"`
	// The managed object's model to validate normalized for querying HCL data. It is empty if we are missing this information.
	HclModel *string `json:"HclModel,omitempty" yaml:"HclModel,omitempty"`
	// The OS Vendor for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information.
	HclOsVendor *string `json:"HclOsVendor,omitempty" yaml:"HclOsVendor,omitempty"`
	// The OS Version for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information.
	HclOsVersion *string `json:"HclOsVersion,omitempty" yaml:"HclOsVersion,omitempty"`
	// The managed object's processor to validate if applicable normalized for querying HCL data. It is empty if we are missing this information.
	HclProcessor *string `json:"HclProcessor,omitempty" yaml:"HclProcessor,omitempty"`
	// The current CIMC version for the server as received from inventory. It is empty if we are missing this information.
	InvFirmwareVersion *string `json:"InvFirmwareVersion,omitempty" yaml:"InvFirmwareVersion,omitempty"`
	// The managed object's model to validate as received from the inventory. It is empty if we are missing this information.
	InvModel *string `json:"InvModel,omitempty" yaml:"InvModel,omitempty"`
	// The OS Vendor for the managed object to validate as received from inventory. It is empty if we are missing this information.
	InvOsVendor *string `json:"InvOsVendor,omitempty" yaml:"InvOsVendor,omitempty"`
	// The OS Version for the managed object to validate as received from inventory. It is empty if we are missing this information.
	InvOsVersion *string `json:"InvOsVersion,omitempty" yaml:"InvOsVersion,omitempty"`
	// The managed object's processor to validate if applicable as received from inventory. It is empty if we are missing this information.
	InvProcessor *string `json:"InvProcessor,omitempty" yaml:"InvProcessor,omitempty"`
	// The reason for the HCL status. It will be one of the following \"Missing-Os-Info\" - we are missing os information in the inventory from the device connector \"Incompatible-Components\" - we have 1 or more components with \"Not-Validated\" status \"Compatible\" - all the components have \"Validated\" status. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted. * `Missing-Os-Info` - This means the HclStatus for the sever failed HCL validation because we have missing os information. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Incompatible-Components` - This means the HclStatus for the sever failed HCL validation because one or more of its components failed validation. To see why components failed check the related HclStatusDetails. * `Compatible` - This means the HclStatus for the sever has passed HCL validation for all of its related components. * `Not-Evaluated` - This means the HclStatus for the sever has not been evaluated because it is exempted.
	Reason *string `json:"Reason,omitempty" yaml:"Reason,omitempty"`
	// The reason generated by the server's HCL validation. For HCL the evaluation can be seen in three logical stages 1. Evaluate the server's hardware status 2. Evaluate the server's software status 3. Evaluate the server's components (each component has its own hardware/software evaluation) The evaluation does not proceed to the next stage until the previous stage is evaluated. Therefore there can be only one validation reason. \"Incompatible-Server\" - the server model is not listed in the HCL. \"Incompatible-Processor\" - the server model and processor combination is not listed in HCL. \"Incompatible-Firmware\" - the server model, processor and server firmware is not listed in HCL. \"Missing-Os-Info\" - the os vendor and version is not listed in HCL with the HW profile. \"Incompatible-Os-Info\" - the os vendor and version is not listed in HCL with the HW profile. \"Incompatible-Components\" - there is one or more components with \"Not-Validated\" status \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). \"Compatible\" - the server and all its components are validated. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted. * `Missing-Os-Driver-Info` - The validation failed becaue the given server has no OS driver information available in the inventory. Either install UCS Tools VIB on the host ESXi or use OS Discovery Tool scripts to provide proper OS information. * `Incompatible-Server` - The validation failed for this server because the server's model was not listed in the HCL. * `Incompatible-Processor` - The validation failed because the given processor was not listed for the given server model. * `Incompatible-Os-Info` - The validation failed because the given OS vendor or version was not listed in HCL for the server PID and processor combination. * `Incompatible-Firmware` - The validation failed because the given server firmware was not listed in the HCL for the given server PID, processor, OS vendor and version. * `Service-Unavailable` - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * `Service-Error` - The validation has failed because the HCL data service has returned a service error or unrecognized result. * `Not-Evaluated` - This means the HclStatus for the sever has not been evaluated because it is exempted. * `Incompatible-Components` - The validation has failed for this server because one or more components have \"Not-Listed\" status. * `Compatible` - The validation has passed for this server's model, processor, OS vendor and version.
	ServerReason *string `json:"ServerReason,omitempty" yaml:"ServerReason,omitempty"`
	// The OS vendor and version are considered part of the software profile for the server. This will provide the HCL validation status for the software profile. For the failure reason see the serverReason property. The status can be be one of the following \"Validated\" - The os vendor/version is listed in the HCL for the server model, processor and firmware \"Not-Listed\" - The os vendor/version is not listed in the HCL for the server model, processor and firmware \"Incomplete\" - The inventory is missing os vendor/version and HCL validation was not performed. \"Not-Evaluated\" - The server is not evaluated against the HCL because it is exempted. * `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL. * `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL. * `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL. * `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.
	SoftwareStatus *string `json:"SoftwareStatus,omitempty" yaml:"SoftwareStatus,omitempty"`
	// The HCL compatibility status of the managed object. The status can be one of the following \"Incomplete\" - there is no enough information to evaluate against the HCL data \"Validated\" - all components have been evaluated against the HCL and they all have \"Validated\" status \"Not-Listed\" - all components have been evaluated against the HCL and one or more have \"Not-Listed\" status. \"Not-Evaluated\" - server is not evaluated against the HCL because it is exempted. * `Incomplete` - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Not-Found` - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL. * `Not-Listed` - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL. * `Validated` - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL. * `Not-Evaluated` - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.
	Status *string `json:"Status,omitempty" yaml:"Status,omitempty"`
	// An array of relationships to condHclStatusDetail resources.
	Details          []CondHclStatusDetailRelationship    `json:"Details,omitempty" yaml:"Details,omitempty"`
	ManagedObject    *InventoryBaseRelationship           `json:"ManagedObject,omitempty" yaml:"ManagedObject,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewCondHclStatus instantiates a new CondHclStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondHclStatus(classId string, objectType string) *CondHclStatus {
	this := CondHclStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var componentStatus string = "Incomplete"
	this.ComponentStatus = &componentStatus
	var hardwareStatus string = "Incomplete"
	this.HardwareStatus = &hardwareStatus
	var reason string = "Missing-Os-Info"
	this.Reason = &reason
	var serverReason string = "Missing-Os-Driver-Info"
	this.ServerReason = &serverReason
	var softwareStatus string = "Incomplete"
	this.SoftwareStatus = &softwareStatus
	var status string = "Incomplete"
	this.Status = &status
	return &this
}

// NewCondHclStatusWithDefaults instantiates a new CondHclStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondHclStatusWithDefaults() *CondHclStatus {
	this := CondHclStatus{}
	var classId string = "cond.HclStatus"
	this.ClassId = classId
	var objectType string = "cond.HclStatus"
	this.ObjectType = objectType
	var componentStatus string = "Incomplete"
	this.ComponentStatus = &componentStatus
	var hardwareStatus string = "Incomplete"
	this.HardwareStatus = &hardwareStatus
	var reason string = "Missing-Os-Info"
	this.Reason = &reason
	var serverReason string = "Missing-Os-Driver-Info"
	this.ServerReason = &serverReason
	var softwareStatus string = "Incomplete"
	this.SoftwareStatus = &softwareStatus
	var status string = "Incomplete"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *CondHclStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CondHclStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CondHclStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CondHclStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComponentStatus returns the ComponentStatus field value if set, zero value otherwise.
func (o *CondHclStatus) GetComponentStatus() string {
	if o == nil || o.ComponentStatus == nil {
		var ret string
		return ret
	}
	return *o.ComponentStatus
}

// GetComponentStatusOk returns a tuple with the ComponentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetComponentStatusOk() (*string, bool) {
	if o == nil || o.ComponentStatus == nil {
		return nil, false
	}
	return o.ComponentStatus, true
}

// HasComponentStatus returns a boolean if a field has been set.
func (o *CondHclStatus) HasComponentStatus() bool {
	if o != nil && o.ComponentStatus != nil {
		return true
	}

	return false
}

// SetComponentStatus gets a reference to the given string and assigns it to the ComponentStatus field.
func (o *CondHclStatus) SetComponentStatus(v string) {
	o.ComponentStatus = &v
}

// GetHardwareStatus returns the HardwareStatus field value if set, zero value otherwise.
func (o *CondHclStatus) GetHardwareStatus() string {
	if o == nil || o.HardwareStatus == nil {
		var ret string
		return ret
	}
	return *o.HardwareStatus
}

// GetHardwareStatusOk returns a tuple with the HardwareStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetHardwareStatusOk() (*string, bool) {
	if o == nil || o.HardwareStatus == nil {
		return nil, false
	}
	return o.HardwareStatus, true
}

// HasHardwareStatus returns a boolean if a field has been set.
func (o *CondHclStatus) HasHardwareStatus() bool {
	if o != nil && o.HardwareStatus != nil {
		return true
	}

	return false
}

// SetHardwareStatus gets a reference to the given string and assigns it to the HardwareStatus field.
func (o *CondHclStatus) SetHardwareStatus(v string) {
	o.HardwareStatus = &v
}

// GetHclFirmwareVersion returns the HclFirmwareVersion field value if set, zero value otherwise.
func (o *CondHclStatus) GetHclFirmwareVersion() string {
	if o == nil || o.HclFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.HclFirmwareVersion
}

// GetHclFirmwareVersionOk returns a tuple with the HclFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetHclFirmwareVersionOk() (*string, bool) {
	if o == nil || o.HclFirmwareVersion == nil {
		return nil, false
	}
	return o.HclFirmwareVersion, true
}

// HasHclFirmwareVersion returns a boolean if a field has been set.
func (o *CondHclStatus) HasHclFirmwareVersion() bool {
	if o != nil && o.HclFirmwareVersion != nil {
		return true
	}

	return false
}

// SetHclFirmwareVersion gets a reference to the given string and assigns it to the HclFirmwareVersion field.
func (o *CondHclStatus) SetHclFirmwareVersion(v string) {
	o.HclFirmwareVersion = &v
}

// GetHclModel returns the HclModel field value if set, zero value otherwise.
func (o *CondHclStatus) GetHclModel() string {
	if o == nil || o.HclModel == nil {
		var ret string
		return ret
	}
	return *o.HclModel
}

// GetHclModelOk returns a tuple with the HclModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetHclModelOk() (*string, bool) {
	if o == nil || o.HclModel == nil {
		return nil, false
	}
	return o.HclModel, true
}

// HasHclModel returns a boolean if a field has been set.
func (o *CondHclStatus) HasHclModel() bool {
	if o != nil && o.HclModel != nil {
		return true
	}

	return false
}

// SetHclModel gets a reference to the given string and assigns it to the HclModel field.
func (o *CondHclStatus) SetHclModel(v string) {
	o.HclModel = &v
}

// GetHclOsVendor returns the HclOsVendor field value if set, zero value otherwise.
func (o *CondHclStatus) GetHclOsVendor() string {
	if o == nil || o.HclOsVendor == nil {
		var ret string
		return ret
	}
	return *o.HclOsVendor
}

// GetHclOsVendorOk returns a tuple with the HclOsVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetHclOsVendorOk() (*string, bool) {
	if o == nil || o.HclOsVendor == nil {
		return nil, false
	}
	return o.HclOsVendor, true
}

// HasHclOsVendor returns a boolean if a field has been set.
func (o *CondHclStatus) HasHclOsVendor() bool {
	if o != nil && o.HclOsVendor != nil {
		return true
	}

	return false
}

// SetHclOsVendor gets a reference to the given string and assigns it to the HclOsVendor field.
func (o *CondHclStatus) SetHclOsVendor(v string) {
	o.HclOsVendor = &v
}

// GetHclOsVersion returns the HclOsVersion field value if set, zero value otherwise.
func (o *CondHclStatus) GetHclOsVersion() string {
	if o == nil || o.HclOsVersion == nil {
		var ret string
		return ret
	}
	return *o.HclOsVersion
}

// GetHclOsVersionOk returns a tuple with the HclOsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetHclOsVersionOk() (*string, bool) {
	if o == nil || o.HclOsVersion == nil {
		return nil, false
	}
	return o.HclOsVersion, true
}

// HasHclOsVersion returns a boolean if a field has been set.
func (o *CondHclStatus) HasHclOsVersion() bool {
	if o != nil && o.HclOsVersion != nil {
		return true
	}

	return false
}

// SetHclOsVersion gets a reference to the given string and assigns it to the HclOsVersion field.
func (o *CondHclStatus) SetHclOsVersion(v string) {
	o.HclOsVersion = &v
}

// GetHclProcessor returns the HclProcessor field value if set, zero value otherwise.
func (o *CondHclStatus) GetHclProcessor() string {
	if o == nil || o.HclProcessor == nil {
		var ret string
		return ret
	}
	return *o.HclProcessor
}

// GetHclProcessorOk returns a tuple with the HclProcessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetHclProcessorOk() (*string, bool) {
	if o == nil || o.HclProcessor == nil {
		return nil, false
	}
	return o.HclProcessor, true
}

// HasHclProcessor returns a boolean if a field has been set.
func (o *CondHclStatus) HasHclProcessor() bool {
	if o != nil && o.HclProcessor != nil {
		return true
	}

	return false
}

// SetHclProcessor gets a reference to the given string and assigns it to the HclProcessor field.
func (o *CondHclStatus) SetHclProcessor(v string) {
	o.HclProcessor = &v
}

// GetInvFirmwareVersion returns the InvFirmwareVersion field value if set, zero value otherwise.
func (o *CondHclStatus) GetInvFirmwareVersion() string {
	if o == nil || o.InvFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.InvFirmwareVersion
}

// GetInvFirmwareVersionOk returns a tuple with the InvFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetInvFirmwareVersionOk() (*string, bool) {
	if o == nil || o.InvFirmwareVersion == nil {
		return nil, false
	}
	return o.InvFirmwareVersion, true
}

// HasInvFirmwareVersion returns a boolean if a field has been set.
func (o *CondHclStatus) HasInvFirmwareVersion() bool {
	if o != nil && o.InvFirmwareVersion != nil {
		return true
	}

	return false
}

// SetInvFirmwareVersion gets a reference to the given string and assigns it to the InvFirmwareVersion field.
func (o *CondHclStatus) SetInvFirmwareVersion(v string) {
	o.InvFirmwareVersion = &v
}

// GetInvModel returns the InvModel field value if set, zero value otherwise.
func (o *CondHclStatus) GetInvModel() string {
	if o == nil || o.InvModel == nil {
		var ret string
		return ret
	}
	return *o.InvModel
}

// GetInvModelOk returns a tuple with the InvModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetInvModelOk() (*string, bool) {
	if o == nil || o.InvModel == nil {
		return nil, false
	}
	return o.InvModel, true
}

// HasInvModel returns a boolean if a field has been set.
func (o *CondHclStatus) HasInvModel() bool {
	if o != nil && o.InvModel != nil {
		return true
	}

	return false
}

// SetInvModel gets a reference to the given string and assigns it to the InvModel field.
func (o *CondHclStatus) SetInvModel(v string) {
	o.InvModel = &v
}

// GetInvOsVendor returns the InvOsVendor field value if set, zero value otherwise.
func (o *CondHclStatus) GetInvOsVendor() string {
	if o == nil || o.InvOsVendor == nil {
		var ret string
		return ret
	}
	return *o.InvOsVendor
}

// GetInvOsVendorOk returns a tuple with the InvOsVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetInvOsVendorOk() (*string, bool) {
	if o == nil || o.InvOsVendor == nil {
		return nil, false
	}
	return o.InvOsVendor, true
}

// HasInvOsVendor returns a boolean if a field has been set.
func (o *CondHclStatus) HasInvOsVendor() bool {
	if o != nil && o.InvOsVendor != nil {
		return true
	}

	return false
}

// SetInvOsVendor gets a reference to the given string and assigns it to the InvOsVendor field.
func (o *CondHclStatus) SetInvOsVendor(v string) {
	o.InvOsVendor = &v
}

// GetInvOsVersion returns the InvOsVersion field value if set, zero value otherwise.
func (o *CondHclStatus) GetInvOsVersion() string {
	if o == nil || o.InvOsVersion == nil {
		var ret string
		return ret
	}
	return *o.InvOsVersion
}

// GetInvOsVersionOk returns a tuple with the InvOsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetInvOsVersionOk() (*string, bool) {
	if o == nil || o.InvOsVersion == nil {
		return nil, false
	}
	return o.InvOsVersion, true
}

// HasInvOsVersion returns a boolean if a field has been set.
func (o *CondHclStatus) HasInvOsVersion() bool {
	if o != nil && o.InvOsVersion != nil {
		return true
	}

	return false
}

// SetInvOsVersion gets a reference to the given string and assigns it to the InvOsVersion field.
func (o *CondHclStatus) SetInvOsVersion(v string) {
	o.InvOsVersion = &v
}

// GetInvProcessor returns the InvProcessor field value if set, zero value otherwise.
func (o *CondHclStatus) GetInvProcessor() string {
	if o == nil || o.InvProcessor == nil {
		var ret string
		return ret
	}
	return *o.InvProcessor
}

// GetInvProcessorOk returns a tuple with the InvProcessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetInvProcessorOk() (*string, bool) {
	if o == nil || o.InvProcessor == nil {
		return nil, false
	}
	return o.InvProcessor, true
}

// HasInvProcessor returns a boolean if a field has been set.
func (o *CondHclStatus) HasInvProcessor() bool {
	if o != nil && o.InvProcessor != nil {
		return true
	}

	return false
}

// SetInvProcessor gets a reference to the given string and assigns it to the InvProcessor field.
func (o *CondHclStatus) SetInvProcessor(v string) {
	o.InvProcessor = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CondHclStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CondHclStatus) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CondHclStatus) SetReason(v string) {
	o.Reason = &v
}

// GetServerReason returns the ServerReason field value if set, zero value otherwise.
func (o *CondHclStatus) GetServerReason() string {
	if o == nil || o.ServerReason == nil {
		var ret string
		return ret
	}
	return *o.ServerReason
}

// GetServerReasonOk returns a tuple with the ServerReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetServerReasonOk() (*string, bool) {
	if o == nil || o.ServerReason == nil {
		return nil, false
	}
	return o.ServerReason, true
}

// HasServerReason returns a boolean if a field has been set.
func (o *CondHclStatus) HasServerReason() bool {
	if o != nil && o.ServerReason != nil {
		return true
	}

	return false
}

// SetServerReason gets a reference to the given string and assigns it to the ServerReason field.
func (o *CondHclStatus) SetServerReason(v string) {
	o.ServerReason = &v
}

// GetSoftwareStatus returns the SoftwareStatus field value if set, zero value otherwise.
func (o *CondHclStatus) GetSoftwareStatus() string {
	if o == nil || o.SoftwareStatus == nil {
		var ret string
		return ret
	}
	return *o.SoftwareStatus
}

// GetSoftwareStatusOk returns a tuple with the SoftwareStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetSoftwareStatusOk() (*string, bool) {
	if o == nil || o.SoftwareStatus == nil {
		return nil, false
	}
	return o.SoftwareStatus, true
}

// HasSoftwareStatus returns a boolean if a field has been set.
func (o *CondHclStatus) HasSoftwareStatus() bool {
	if o != nil && o.SoftwareStatus != nil {
		return true
	}

	return false
}

// SetSoftwareStatus gets a reference to the given string and assigns it to the SoftwareStatus field.
func (o *CondHclStatus) SetSoftwareStatus(v string) {
	o.SoftwareStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CondHclStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CondHclStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CondHclStatus) SetStatus(v string) {
	o.Status = &v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondHclStatus) GetDetails() []CondHclStatusDetailRelationship {
	if o == nil {
		var ret []CondHclStatusDetailRelationship
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondHclStatus) GetDetailsOk() (*[]CondHclStatusDetailRelationship, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return &o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CondHclStatus) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []CondHclStatusDetailRelationship and assigns it to the Details field.
func (o *CondHclStatus) SetDetails(v []CondHclStatusDetailRelationship) {
	o.Details = v
}

// GetManagedObject returns the ManagedObject field value if set, zero value otherwise.
func (o *CondHclStatus) GetManagedObject() InventoryBaseRelationship {
	if o == nil || o.ManagedObject == nil {
		var ret InventoryBaseRelationship
		return ret
	}
	return *o.ManagedObject
}

// GetManagedObjectOk returns a tuple with the ManagedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetManagedObjectOk() (*InventoryBaseRelationship, bool) {
	if o == nil || o.ManagedObject == nil {
		return nil, false
	}
	return o.ManagedObject, true
}

// HasManagedObject returns a boolean if a field has been set.
func (o *CondHclStatus) HasManagedObject() bool {
	if o != nil && o.ManagedObject != nil {
		return true
	}

	return false
}

// SetManagedObject gets a reference to the given InventoryBaseRelationship and assigns it to the ManagedObject field.
func (o *CondHclStatus) SetManagedObject(v InventoryBaseRelationship) {
	o.ManagedObject = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *CondHclStatus) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatus) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *CondHclStatus) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *CondHclStatus) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o CondHclStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ComponentStatus != nil {
		toSerialize["ComponentStatus"] = o.ComponentStatus
	}
	if o.HardwareStatus != nil {
		toSerialize["HardwareStatus"] = o.HardwareStatus
	}
	if o.HclFirmwareVersion != nil {
		toSerialize["HclFirmwareVersion"] = o.HclFirmwareVersion
	}
	if o.HclModel != nil {
		toSerialize["HclModel"] = o.HclModel
	}
	if o.HclOsVendor != nil {
		toSerialize["HclOsVendor"] = o.HclOsVendor
	}
	if o.HclOsVersion != nil {
		toSerialize["HclOsVersion"] = o.HclOsVersion
	}
	if o.HclProcessor != nil {
		toSerialize["HclProcessor"] = o.HclProcessor
	}
	if o.InvFirmwareVersion != nil {
		toSerialize["InvFirmwareVersion"] = o.InvFirmwareVersion
	}
	if o.InvModel != nil {
		toSerialize["InvModel"] = o.InvModel
	}
	if o.InvOsVendor != nil {
		toSerialize["InvOsVendor"] = o.InvOsVendor
	}
	if o.InvOsVersion != nil {
		toSerialize["InvOsVersion"] = o.InvOsVersion
	}
	if o.InvProcessor != nil {
		toSerialize["InvProcessor"] = o.InvProcessor
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.ServerReason != nil {
		toSerialize["ServerReason"] = o.ServerReason
	}
	if o.SoftwareStatus != nil {
		toSerialize["SoftwareStatus"] = o.SoftwareStatus
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Details != nil {
		toSerialize["Details"] = o.Details
	}
	if o.ManagedObject != nil {
		toSerialize["ManagedObject"] = o.ManagedObject
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableCondHclStatus struct {
	value *CondHclStatus
	isSet bool
}

func (v NullableCondHclStatus) Get() *CondHclStatus {
	return v.value
}

func (v *NullableCondHclStatus) Set(val *CondHclStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCondHclStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCondHclStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondHclStatus(val *CondHclStatus) *NullableCondHclStatus {
	return &NullableCondHclStatus{value: val, isSet: true}
}

func (v NullableCondHclStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondHclStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
