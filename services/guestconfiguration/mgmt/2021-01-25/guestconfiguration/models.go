package guestconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/guestconfiguration/mgmt/2021-01-25/guestconfiguration"

// Assignment guest configuration assignment is an association between a machine and guest configuration.
type Assignment struct {
	autorest.Response `json:"-"`
	// Properties - Properties of the Guest configuration assignment.
	Properties *AssignmentProperties `json:"properties,omitempty"`
	// SystemData - READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty"`
	// ID - READ-ONLY; ARM resource id of the guest configuration assignment.
	ID *string `json:"id,omitempty"`
	// Name - Name of the guest configuration assignment.
	Name *string `json:"name,omitempty"`
	// Location - Region where the VM is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Assignment.
func (a Assignment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.Properties != nil {
		objectMap["properties"] = a.Properties
	}
	if a.Name != nil {
		objectMap["name"] = a.Name
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	return json.Marshal(objectMap)
}

// AssignmentInfo information about the guest configuration assignment.
type AssignmentInfo struct {
	// Name - READ-ONLY; Name of the guest configuration assignment.
	Name *string `json:"name,omitempty"`
	// Configuration - Information about the configuration.
	Configuration *ConfigurationInfo `json:"configuration,omitempty"`
}

// MarshalJSON is the custom marshaler for AssignmentInfo.
func (ai AssignmentInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ai.Configuration != nil {
		objectMap["configuration"] = ai.Configuration
	}
	return json.Marshal(objectMap)
}

// AssignmentList the response of the list guest configuration assignment operation.
type AssignmentList struct {
	autorest.Response `json:"-"`
	// Value - Result of the list guest configuration assignment operation.
	Value *[]Assignment `json:"value,omitempty"`
}

// AssignmentProperties guest configuration assignment properties.
type AssignmentProperties struct {
	// TargetResourceID - READ-ONLY; VM resource Id.
	TargetResourceID *string `json:"targetResourceId,omitempty"`
	// GuestConfiguration - The guest configuration to assign.
	GuestConfiguration *Navigation `json:"guestConfiguration,omitempty"`
	// ComplianceStatus - READ-ONLY; A value indicating compliance status of the machine for the assigned guest configuration. Possible values include: 'ComplianceStatusCompliant', 'ComplianceStatusNonCompliant', 'ComplianceStatusPending'
	ComplianceStatus ComplianceStatus `json:"complianceStatus,omitempty"`
	// LastComplianceStatusChecked - READ-ONLY; Date and time when last compliance status was checked.
	LastComplianceStatusChecked *date.Time `json:"lastComplianceStatusChecked,omitempty"`
	// LatestReportID - READ-ONLY; Id of the latest report for the guest configuration assignment.
	LatestReportID *string `json:"latestReportId,omitempty"`
	// LatestAssignmentReport - Last reported guest configuration assignment report.
	LatestAssignmentReport *AssignmentReport `json:"latestAssignmentReport,omitempty"`
	// Context - The source which initiated the guest configuration assignment. Ex: Azure Policy
	Context *string `json:"context,omitempty"`
	// AssignmentHash - READ-ONLY; Combined hash of the configuration package and parameters.
	AssignmentHash *string `json:"assignmentHash,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state, which only appears in the response. Possible values include: 'ProvisioningStateSucceeded', 'ProvisioningStateFailed', 'ProvisioningStateCanceled', 'ProvisioningStateCreated'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// ResourceType - READ-ONLY; Type of the resource - VMSS / VM
	ResourceType *string `json:"resourceType,omitempty"`
	// VmssVMList - The list of VM Compliance data for VMSS
	VmssVMList *[]VMSSVMInfo `json:"vmssVMList,omitempty"`
}

// MarshalJSON is the custom marshaler for AssignmentProperties.
func (ap AssignmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ap.GuestConfiguration != nil {
		objectMap["guestConfiguration"] = ap.GuestConfiguration
	}
	if ap.LatestAssignmentReport != nil {
		objectMap["latestAssignmentReport"] = ap.LatestAssignmentReport
	}
	if ap.Context != nil {
		objectMap["context"] = ap.Context
	}
	if ap.VmssVMList != nil {
		objectMap["vmssVMList"] = ap.VmssVMList
	}
	return json.Marshal(objectMap)
}

// AssignmentReport ...
type AssignmentReport struct {
	// ID - READ-ONLY; ARM resource id of the report for the guest configuration assignment.
	ID *string `json:"id,omitempty"`
	// ReportID - READ-ONLY; GUID that identifies the guest configuration assignment report under a subscription, resource group.
	ReportID *string `json:"reportId,omitempty"`
	// Assignment - Configuration details of the guest configuration assignment.
	Assignment *AssignmentInfo `json:"assignment,omitempty"`
	// VM - Information about the VM.
	VM *VMInfo `json:"vm,omitempty"`
	// StartTime - READ-ONLY; Start date and time of the guest configuration assignment compliance status check.
	StartTime *date.Time `json:"startTime,omitempty"`
	// EndTime - READ-ONLY; End date and time of the guest configuration assignment compliance status check.
	EndTime *date.Time `json:"endTime,omitempty"`
	// ComplianceStatus - READ-ONLY; A value indicating compliance status of the machine for the assigned guest configuration. Possible values include: 'ComplianceStatusCompliant', 'ComplianceStatusNonCompliant', 'ComplianceStatusPending'
	ComplianceStatus ComplianceStatus `json:"complianceStatus,omitempty"`
	// OperationType - READ-ONLY; Type of report, Consistency or Initial. Possible values include: 'TypeConsistency', 'TypeInitial'
	OperationType Type `json:"operationType,omitempty"`
	// Resources - The list of resources for which guest configuration assignment compliance is checked.
	Resources *[]AssignmentReportResource `json:"resources,omitempty"`
}

// MarshalJSON is the custom marshaler for AssignmentReport.
func (ar AssignmentReport) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ar.Assignment != nil {
		objectMap["assignment"] = ar.Assignment
	}
	if ar.VM != nil {
		objectMap["vm"] = ar.VM
	}
	if ar.Resources != nil {
		objectMap["resources"] = ar.Resources
	}
	return json.Marshal(objectMap)
}

// AssignmentReportDetails details of the guest configuration assignment report.
type AssignmentReportDetails struct {
	// ComplianceStatus - READ-ONLY; A value indicating compliance status of the machine for the assigned guest configuration. Possible values include: 'ComplianceStatusCompliant', 'ComplianceStatusNonCompliant', 'ComplianceStatusPending'
	ComplianceStatus ComplianceStatus `json:"complianceStatus,omitempty"`
	// StartTime - READ-ONLY; Start date and time of the guest configuration assignment compliance status check.
	StartTime *date.Time `json:"startTime,omitempty"`
	// EndTime - READ-ONLY; End date and time of the guest configuration assignment compliance status check.
	EndTime *date.Time `json:"endTime,omitempty"`
	// JobID - READ-ONLY; GUID of the report.
	JobID *string `json:"jobId,omitempty"`
	// OperationType - READ-ONLY; Type of report, Consistency or Initial. Possible values include: 'TypeConsistency', 'TypeInitial'
	OperationType Type `json:"operationType,omitempty"`
	// Resources - The list of resources for which guest configuration assignment compliance is checked.
	Resources *[]AssignmentReportResource `json:"resources,omitempty"`
}

// MarshalJSON is the custom marshaler for AssignmentReportDetails.
func (ard AssignmentReportDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ard.Resources != nil {
		objectMap["resources"] = ard.Resources
	}
	return json.Marshal(objectMap)
}

// AssignmentReportList list of guest configuration assignment reports.
type AssignmentReportList struct {
	autorest.Response `json:"-"`
	// Value - List of reports for the guest configuration. Report contains information such as compliance status, reason and more.
	Value *[]AssignmentReportType `json:"value,omitempty"`
}

// AssignmentReportProperties report for the guest configuration assignment. Report contains information
// such as compliance status, reason, and more.
type AssignmentReportProperties struct {
	// ComplianceStatus - READ-ONLY; A value indicating compliance status of the machine for the assigned guest configuration. Possible values include: 'ComplianceStatusCompliant', 'ComplianceStatusNonCompliant', 'ComplianceStatusPending'
	ComplianceStatus ComplianceStatus `json:"complianceStatus,omitempty"`
	// ReportID - READ-ONLY; GUID that identifies the guest configuration assignment report under a subscription, resource group.
	ReportID *string `json:"reportId,omitempty"`
	// Assignment - Configuration details of the guest configuration assignment.
	Assignment *AssignmentInfo `json:"assignment,omitempty"`
	// VM - Information about the VM.
	VM *VMInfo `json:"vm,omitempty"`
	// StartTime - READ-ONLY; Start date and time of the guest configuration assignment compliance status check.
	StartTime *date.Time `json:"startTime,omitempty"`
	// EndTime - READ-ONLY; End date and time of the guest configuration assignment compliance status check.
	EndTime *date.Time `json:"endTime,omitempty"`
	// Details - Details of the assignment report.
	Details *AssignmentReportDetails `json:"details,omitempty"`
	// VmssResourceID - READ-ONLY; Azure resource Id of the VMSS.
	VmssResourceID *string `json:"vmssResourceId,omitempty"`
}

// MarshalJSON is the custom marshaler for AssignmentReportProperties.
func (arp AssignmentReportProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if arp.Assignment != nil {
		objectMap["assignment"] = arp.Assignment
	}
	if arp.VM != nil {
		objectMap["vm"] = arp.VM
	}
	if arp.Details != nil {
		objectMap["details"] = arp.Details
	}
	return json.Marshal(objectMap)
}

// AssignmentReportResource the guest configuration assignment resource.
type AssignmentReportResource struct {
	// ComplianceStatus - READ-ONLY; A value indicating compliance status of the machine for the assigned guest configuration. Possible values include: 'ComplianceStatusCompliant', 'ComplianceStatusNonCompliant', 'ComplianceStatusPending'
	ComplianceStatus ComplianceStatus `json:"complianceStatus,omitempty"`
	// ResourceID - READ-ONLY; Name of the guest configuration assignment resource setting.
	ResourceID *string `json:"resourceId,omitempty"`
	// Reasons - Compliance reason and reason code for a resource.
	Reasons *[]AssignmentReportResourceComplianceReason `json:"reasons,omitempty"`
	// Properties - READ-ONLY; Properties of a guest configuration assignment resource.
	Properties interface{} `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AssignmentReportResource.
func (arr AssignmentReportResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if arr.Reasons != nil {
		objectMap["reasons"] = arr.Reasons
	}
	return json.Marshal(objectMap)
}

// AssignmentReportResourceComplianceReason reason and code for the compliance of the guest configuration
// assignment resource.
type AssignmentReportResourceComplianceReason struct {
	// Phrase - READ-ONLY; Reason for the compliance of the guest configuration assignment resource.
	Phrase *string `json:"phrase,omitempty"`
	// Code - READ-ONLY; Code for the compliance of the guest configuration assignment resource.
	Code *string `json:"code,omitempty"`
}

// MarshalJSON is the custom marshaler for AssignmentReportResourceComplianceReason.
func (arrcr AssignmentReportResourceComplianceReason) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AssignmentReportType report for the guest configuration assignment. Report contains information such as
// compliance status, reason, and more.
type AssignmentReportType struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; ARM resource id of the report for the guest configuration assignment.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; GUID that identifies the guest configuration assignment report under a subscription, resource group.
	Name *string `json:"name,omitempty"`
	// Properties - Properties of the guest configuration report.
	Properties *AssignmentReportProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AssignmentReportType.
func (art AssignmentReportType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if art.Properties != nil {
		objectMap["properties"] = art.Properties
	}
	return json.Marshal(objectMap)
}

// ConfigurationInfo information about the configuration.
type ConfigurationInfo struct {
	// Name - READ-ONLY; Name of the configuration.
	Name *string `json:"name,omitempty"`
	// Version - READ-ONLY; Version of the configuration.
	Version *string `json:"version,omitempty"`
}

// MarshalJSON is the custom marshaler for ConfigurationInfo.
func (ci ConfigurationInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ConfigurationParameter represents a configuration parameter.
type ConfigurationParameter struct {
	// Name - Name of the configuration parameter.
	Name *string `json:"name,omitempty"`
	// Value - Value of the configuration parameter.
	Value *string `json:"value,omitempty"`
}

// ConfigurationSetting configuration setting of LCM (Local Configuration Manager).
type ConfigurationSetting struct {
	// ConfigurationMode - Specifies how the LCM(Local Configuration Manager) actually applies the configuration to the target nodes. Possible values are ApplyOnly, ApplyAndMonitor, and ApplyAndAutoCorrect. Possible values include: 'ConfigurationModeApplyOnly', 'ConfigurationModeApplyAndMonitor', 'ConfigurationModeApplyAndAutoCorrect'
	ConfigurationMode ConfigurationMode `json:"configurationMode,omitempty"`
	// AllowModuleOverwrite - If true - new configurations downloaded from the pull service are allowed to overwrite the old ones on the target node. Otherwise, false
	AllowModuleOverwrite *bool `json:"allowModuleOverwrite,omitempty"`
	// ActionAfterReboot - Specifies what happens after a reboot during the application of a configuration. The possible values are ContinueConfiguration and StopConfiguration. Possible values include: 'ActionAfterRebootContinueConfiguration', 'ActionAfterRebootStopConfiguration'
	ActionAfterReboot ActionAfterReboot `json:"actionAfterReboot,omitempty"`
	// RefreshFrequencyMins - The time interval, in minutes, at which the LCM checks a pull service to get updated configurations. This value is ignored if the LCM is not configured in pull mode. The default value is 30.
	RefreshFrequencyMins *float64 `json:"refreshFrequencyMins,omitempty"`
	// RebootIfNeeded - Set this to true to automatically reboot the node after a configuration that requires reboot is applied. Otherwise, you will have to manually reboot the node for any configuration that requires it. The default value is false. To use this setting when a reboot condition is enacted by something other than DSC (such as Windows Installer), combine this setting with the xPendingReboot module.
	RebootIfNeeded *bool `json:"rebootIfNeeded,omitempty"`
	// ConfigurationModeFrequencyMins - How often, in minutes, the current configuration is checked and applied. This property is ignored if the ConfigurationMode property is set to ApplyOnly. The default value is 15.
	ConfigurationModeFrequencyMins *float64 `json:"configurationModeFrequencyMins,omitempty"`
}

// ErrorResponse error response of an operation failure
type ErrorResponse struct {
	Error *ErrorResponseError `json:"error,omitempty"`
}

// ErrorResponseError ...
type ErrorResponseError struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Detail error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// Navigation guest configuration is an artifact that encapsulates DSC configuration and its dependencies.
// The artifact is a zip file containing DSC configuration (as MOF) and dependent resources and other
// dependencies like modules.
type Navigation struct {
	// Kind - Kind of the guest configuration. For example:DSC. Possible values include: 'KindDSC'
	Kind Kind `json:"kind,omitempty"`
	// Name - Name of the guest configuration.
	Name *string `json:"name,omitempty"`
	// Version - Version of the guest configuration.
	Version *string `json:"version,omitempty"`
	// ContentURI - Uri of the storage where guest configuration package is uploaded.
	ContentURI *string `json:"contentUri,omitempty"`
	// ContentHash - Combined hash of the guest configuration package and configuration parameters.
	ContentHash *string `json:"contentHash,omitempty"`
	// AssignmentType - Specifies the assignment type and execution of the configuration. Possible values are Audit, DeployAndAutoCorrect, ApplyAndAutoCorrect and ApplyAndMonitor. Possible values include: 'AssignmentTypeAudit', 'AssignmentTypeDeployAndAutoCorrect', 'AssignmentTypeApplyAndAutoCorrect', 'AssignmentTypeApplyAndMonitor'
	AssignmentType AssignmentType `json:"assignmentType,omitempty"`
	// ConfigurationParameter - The configuration parameters for the guest configuration.
	ConfigurationParameter *[]ConfigurationParameter `json:"configurationParameter,omitempty"`
	// ConfigurationSetting - The configuration setting for the guest configuration.
	ConfigurationSetting *ConfigurationSetting `json:"configurationSetting,omitempty"`
}

// Operation guestConfiguration REST API operation
type Operation struct {
	// Name - Operation name: For ex. providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/write or read
	Name *string `json:"name,omitempty"`
	// Display - Provider, Resource, Operation and description values.
	Display *OperationDisplay `json:"display,omitempty"`
	// OperationProperties - Provider, Resource, Operation and description values.
	*OperationProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.Name != nil {
		objectMap["name"] = o.Name
	}
	if o.Display != nil {
		objectMap["display"] = o.Display
	}
	if o.OperationProperties != nil {
		objectMap["properties"] = o.OperationProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Operation struct.
func (o *Operation) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				o.Name = &name
			}
		case "display":
			if v != nil {
				var display OperationDisplay
				err = json.Unmarshal(*v, &display)
				if err != nil {
					return err
				}
				o.Display = &display
			}
		case "properties":
			if v != nil {
				var operationProperties OperationProperties
				err = json.Unmarshal(*v, &operationProperties)
				if err != nil {
					return err
				}
				o.OperationProperties = &operationProperties
			}
		}
	}

	return nil
}

// OperationDisplay provider, Resource, Operation and description values.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.GuestConfiguration
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed:  For ex.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description about operation.
	Description *string `json:"description,omitempty"`
}

// OperationList the response model for the list of Automation operations
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - List of Automation operations supported by the Automation resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// OperationProperties provider, Resource, Operation and description values.
type OperationProperties struct {
	// StatusCode - Service provider: Microsoft.GuestConfiguration
	StatusCode *string `json:"statusCode,omitempty"`
}

// ProxyResource ARM proxy resource.
type ProxyResource struct {
	// ID - READ-ONLY; ARM resource id of the guest configuration assignment.
	ID *string `json:"id,omitempty"`
	// Name - Name of the guest configuration assignment.
	Name *string `json:"name,omitempty"`
	// Location - Region where the VM is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ProxyResource.
func (pr ProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pr.Name != nil {
		objectMap["name"] = pr.Name
	}
	if pr.Location != nil {
		objectMap["location"] = pr.Location
	}
	return json.Marshal(objectMap)
}

// Resource the core properties of ARM resources
type Resource struct {
	// ID - READ-ONLY; ARM resource id of the guest configuration assignment.
	ID *string `json:"id,omitempty"`
	// Name - Name of the guest configuration assignment.
	Name *string `json:"name,omitempty"`
	// Location - Region where the VM is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	return json.Marshal(objectMap)
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'CreatedByTypeUser', 'CreatedByTypeApplication', 'CreatedByTypeManagedIdentity', 'CreatedByTypeKey'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'CreatedByTypeUser', 'CreatedByTypeApplication', 'CreatedByTypeManagedIdentity', 'CreatedByTypeKey'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The timestamp of resource last modification (UTC)
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// ID - READ-ONLY; ARM resource id of the guest configuration assignment.
	ID *string `json:"id,omitempty"`
	// Name - Name of the guest configuration assignment.
	Name *string `json:"name,omitempty"`
	// Location - Region where the VM is located.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Name != nil {
		objectMap["name"] = tr.Name
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}

// VMInfo information about the VM.
type VMInfo struct {
	// ID - READ-ONLY; Azure resource Id of the VM.
	ID *string `json:"id,omitempty"`
	// UUID - READ-ONLY; UUID(Universally Unique Identifier) of the VM.
	UUID *string `json:"uuid,omitempty"`
}

// MarshalJSON is the custom marshaler for VMInfo.
func (vi VMInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// VMSSVMInfo information about VMSS VM
type VMSSVMInfo struct {
	// VMID - READ-ONLY; UUID of the VM.
	VMID *string `json:"vmId,omitempty"`
	// VMResourceID - READ-ONLY; Azure resource Id of the VM.
	VMResourceID *string `json:"vmResourceId,omitempty"`
	// ComplianceStatus - READ-ONLY; A value indicating compliance status of the machine for the assigned guest configuration. Possible values include: 'ComplianceStatusCompliant', 'ComplianceStatusNonCompliant', 'ComplianceStatusPending'
	ComplianceStatus ComplianceStatus `json:"complianceStatus,omitempty"`
	// LatestReportID - READ-ONLY; Id of the latest report for the guest configuration assignment.
	LatestReportID *string `json:"latestReportId,omitempty"`
	// LastComplianceChecked - READ-ONLY; Date and time when last compliance status was checked.
	LastComplianceChecked *date.Time `json:"lastComplianceChecked,omitempty"`
}

// MarshalJSON is the custom marshaler for VMSSVMInfo.
func (vi VMSSVMInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}
