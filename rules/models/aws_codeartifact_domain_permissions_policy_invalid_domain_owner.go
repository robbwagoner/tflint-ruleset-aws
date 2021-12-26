// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule checks the pattern is valid
type AwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule returns new rule with default attributes
func NewAwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule() *AwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule {
	return &AwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule{
		resourceType:  "aws_codeartifact_domain_permissions_policy",
		attributeName: "domain_owner",
		max:           12,
		min:           12,
		pattern:       regexp.MustCompile(`^[0-9]{12}$`),
	}
}

// Name returns the rule name
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule) Name() string {
	return "aws_codeartifact_domain_permissions_policy_invalid_domain_owner"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidDomainOwnerRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"domain_owner must be 12 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"domain_owner must be 12 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9]{12}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
