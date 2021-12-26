// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule checks the pattern is valid
type AwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule returns new rule with default attributes
func NewAwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule() *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule {
	return &AwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule{
		resourceType:  "aws_codeartifact_domain_permissions_policy",
		attributeName: "policy_document",
		max:           5120,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule) Name() string {
	return "aws_codeartifact_domain_permissions_policy_invalid_policy_document"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidPolicyDocumentRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"policy_document must be 5120 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"policy_document must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
