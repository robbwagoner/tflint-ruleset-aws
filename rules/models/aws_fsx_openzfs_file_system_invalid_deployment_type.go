// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule checks the pattern is valid
type AwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule returns new rule with default attributes
func NewAwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule() *AwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule {
	return &AwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule{
		resourceType:  "aws_fsx_openzfs_file_system",
		attributeName: "deployment_type",
		enum: []string{
			"SINGLE_AZ_1",
		},
	}
}

// Name returns the rule name
func (r *AwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule) Name() string {
	return "aws_fsx_openzfs_file_system_invalid_deployment_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxOpenzfsFileSystemInvalidDeploymentTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as deployment_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
