// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsKmsReplicaKeyInvalidDescriptionRule checks the pattern is valid
type AwsKmsReplicaKeyInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsKmsReplicaKeyInvalidDescriptionRule returns new rule with default attributes
func NewAwsKmsReplicaKeyInvalidDescriptionRule() *AwsKmsReplicaKeyInvalidDescriptionRule {
	return &AwsKmsReplicaKeyInvalidDescriptionRule{
		resourceType:  "aws_kms_replica_key",
		attributeName: "description",
		max:           8192,
	}
}

// Name returns the rule name
func (r *AwsKmsReplicaKeyInvalidDescriptionRule) Name() string {
	return "aws_kms_replica_key_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKmsReplicaKeyInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKmsReplicaKeyInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKmsReplicaKeyInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKmsReplicaKeyInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 8192 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
