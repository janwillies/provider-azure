/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this MaintenanceAssignmentDedicatedHost
func (mg *MaintenanceAssignmentDedicatedHost) GetTerraformResourceType() string {
	return "azurerm_maintenance_assignment_dedicated_host"
}

// GetConnectionDetailsMapping for this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetInitParameters for this MaintenanceAssignmentDedicatedHost
func (tr *MaintenanceAssignmentDedicatedHost) SetInitParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.InitProvider)
}

// LateInitialize this MaintenanceAssignmentDedicatedHost using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MaintenanceAssignmentDedicatedHost) LateInitialize(attrs []byte) (bool, error) {
	params := &MaintenanceAssignmentDedicatedHostParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MaintenanceAssignmentDedicatedHost) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MaintenanceAssignmentVirtualMachine
func (mg *MaintenanceAssignmentVirtualMachine) GetTerraformResourceType() string {
	return "azurerm_maintenance_assignment_virtual_machine"
}

// GetConnectionDetailsMapping for this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetInitParameters for this MaintenanceAssignmentVirtualMachine
func (tr *MaintenanceAssignmentVirtualMachine) SetInitParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.InitProvider)
}

// LateInitialize this MaintenanceAssignmentVirtualMachine using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MaintenanceAssignmentVirtualMachine) LateInitialize(attrs []byte) (bool, error) {
	params := &MaintenanceAssignmentVirtualMachineParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MaintenanceAssignmentVirtualMachine) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MaintenanceConfiguration
func (mg *MaintenanceConfiguration) GetTerraformResourceType() string {
	return "azurerm_maintenance_configuration"
}

// GetConnectionDetailsMapping for this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetInitParameters for this MaintenanceConfiguration
func (tr *MaintenanceConfiguration) SetInitParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.InitProvider)
}

// LateInitialize this MaintenanceConfiguration using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MaintenanceConfiguration) LateInitialize(attrs []byte) (bool, error) {
	params := &MaintenanceConfigurationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MaintenanceConfiguration) GetTerraformSchemaVersion() int {
	return 1
}
