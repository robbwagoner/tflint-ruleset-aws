// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactDomainInvalidEncryptionKeyRule checks the pattern is valid
type AwsCodeartifactDomainInvalidEncryptionKeyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodeartifactDomainInvalidEncryptionKeyRule returns new rule with default attributes
func NewAwsCodeartifactDomainInvalidEncryptionKeyRule() *AwsCodeartifactDomainInvalidEncryptionKeyRule {
	return &AwsCodeartifactDomainInvalidEncryptionKeyRule{
		resourceType:  "aws_codeartifact_domain",
		attributeName: "encryption_key",
		max:           1011,
		min:           1,
		pattern:       regexp.MustCompile(`^\S+$`),
	}
}

// Name returns the rule name
func (r *AwsCodeartifactDomainInvalidEncryptionKeyRule) Name() string {
	return "aws_codeartifact_domain_invalid_encryption_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactDomainInvalidEncryptionKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactDomainInvalidEncryptionKeyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactDomainInvalidEncryptionKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactDomainInvalidEncryptionKeyRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"encryption_key must be 1011 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"encryption_key must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\S+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
