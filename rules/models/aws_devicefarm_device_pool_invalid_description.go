// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDevicefarmDevicePoolInvalidDescriptionRule checks the pattern is valid
type AwsDevicefarmDevicePoolInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsDevicefarmDevicePoolInvalidDescriptionRule returns new rule with default attributes
func NewAwsDevicefarmDevicePoolInvalidDescriptionRule() *AwsDevicefarmDevicePoolInvalidDescriptionRule {
	return &AwsDevicefarmDevicePoolInvalidDescriptionRule{
		resourceType:  "aws_devicefarm_device_pool",
		attributeName: "description",
		max:           16384,
	}
}

// Name returns the rule name
func (r *AwsDevicefarmDevicePoolInvalidDescriptionRule) Name() string {
	return "aws_devicefarm_device_pool_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDevicefarmDevicePoolInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDevicefarmDevicePoolInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDevicefarmDevicePoolInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDevicefarmDevicePoolInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 16384 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
