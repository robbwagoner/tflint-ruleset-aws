// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppstreamFleetInvalidDescriptionRule checks the pattern is valid
type AwsAppstreamFleetInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsAppstreamFleetInvalidDescriptionRule returns new rule with default attributes
func NewAwsAppstreamFleetInvalidDescriptionRule() *AwsAppstreamFleetInvalidDescriptionRule {
	return &AwsAppstreamFleetInvalidDescriptionRule{
		resourceType:  "aws_appstream_fleet",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsAppstreamFleetInvalidDescriptionRule) Name() string {
	return "aws_appstream_fleet_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppstreamFleetInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppstreamFleetInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppstreamFleetInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppstreamFleetInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 256 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
