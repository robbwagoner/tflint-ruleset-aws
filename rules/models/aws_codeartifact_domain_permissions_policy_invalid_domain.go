// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactDomainPermissionsPolicyInvalidDomainRule checks the pattern is valid
type AwsCodeartifactDomainPermissionsPolicyInvalidDomainRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodeartifactDomainPermissionsPolicyInvalidDomainRule returns new rule with default attributes
func NewAwsCodeartifactDomainPermissionsPolicyInvalidDomainRule() *AwsCodeartifactDomainPermissionsPolicyInvalidDomainRule {
	return &AwsCodeartifactDomainPermissionsPolicyInvalidDomainRule{
		resourceType:  "aws_codeartifact_domain_permissions_policy",
		attributeName: "domain",
		max:           50,
		min:           2,
		pattern:       regexp.MustCompile(`^[a-z][a-z0-9\-]{0,48}[a-z0-9]$`),
	}
}

// Name returns the rule name
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidDomainRule) Name() string {
	return "aws_codeartifact_domain_permissions_policy_invalid_domain"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidDomainRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidDomainRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidDomainRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactDomainPermissionsPolicyInvalidDomainRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"domain must be 50 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"domain must be 2 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z][a-z0-9\-]{0,48}[a-z0-9]$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
